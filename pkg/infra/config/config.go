// Package config provides configuration for the infra layer
package config

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/BetaLixT/gex/pkg/domain/base/logger"
	"github.com/BetaLixT/gex/pkg/domain/common"
	"github.com/BetaLixT/gex/pkg/infra/redisdb"
	"github.com/BetaLixT/gex/pkg/infra/roachdb"
	"github.com/BetaLixT/gex/pkg/infra/snowflake"
	"github.com/BetaLixT/gex/pkg/infra/trace"
	"github.com/BetaLixT/gex/pkg/infra/trace/appinsights"
	"github.com/BetaLixT/gex/pkg/infra/trace/jaeger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Initializer including this as your dependency ensures that configs have
// been loaded from all sources before extracting from environment
type Initializer struct {
	lgrf logger.IFactory
}

// NewInitializer loads configs from the .env file
func NewInitializer(
	lgrf logger.IFactory,
) *Initializer {
	c := &Initializer{
		lgrf: lgrf,
	}
	c.LoadConfigCustom("./cfg/.env")
	return c
}

// LoadConfigCustom loads config from a given file
func (c *Initializer) LoadConfigCustom(loc string) {
	err := godotenv.Load(loc)
	if err != nil {
		lgr := c.lgrf.Create(context.Background())
		lgr.Warn(
			"failed to load .env file, probably missing",
			zap.String("location", loc),
		)
	}
}

// NewRedisOptions provides redis options
func NewRedisOptions(c *Initializer) *redisdb.Options {
	address := os.Getenv("RedisAddress")
	if address == "" {
		panic("missing redis address config")
	}
	password := os.Getenv("RedisPassword")
	if password == "" {
		lgr := c.lgrf.Create(context.Background())
		lgr.Warn("redis password missing")
	}
	tls := os.Getenv("RedisTls") == "true"
	databaseNumber := os.Getenv("RedisDatabase")
	db, err := strconv.Atoi(databaseNumber)
	if err != nil {
		db = 0
		lgr := c.lgrf.Create(context.Background())
		lgr.Warn("no database number was provided for redis, using default")
	}

	return &redisdb.Options{
		Address:     address,
		Password:    password,
		ServiceName: address,
		TLS:         tls,
		Database:    db,
	}
}

// NewSnowflakeOptions provides snowflake options
func NewSnowflakeOptions(_ *Initializer) *snowflake.Options {
	nn := os.Getenv("SnowflakeNodeNumber")
	if nn == "" {
		panic("missing snowflake node number")
	}
	nni, err := strconv.ParseInt(
		nn,
		10,
		64,
	)
	if err != nil {
		panic("failed to parse node number")
	}

	return &snowflake.Options{
		NodeNumber: nni,
	}
}

// NewPSQLDBOptions provides psqldb options
func NewRoachDBOptions(_ *Initializer) *roachdb.DatabaseOptions {
	cons := os.Getenv("DatabaseConnectionString")
	if cons == "" {
		panic("missing database connection string config")
	}

	split := strings.Split(cons, " ")
	name := "psql-database"
	for idx := range split {
		if strings.HasPrefix(split[idx], "host=") {
			name = strings.TrimPrefix(split[idx], "host=")
		}
	}
	return &roachdb.DatabaseOptions{
		ConnectionString:    cons,
		DatabaseServiceName: name,
	}
}

// NewAppInsightsExporterOptions provides app insights exporter options
func NewAppInsightsExporterOptions(
	c *Initializer,
) *appinsights.ExporterOptions {
	inskey := os.Getenv("InsightsInstrumentationKey")
	lgr := c.lgrf.Create(context.Background())
	if inskey == "" {
		lgr.Warn("missing insights instrumentation key")
	}
	return &appinsights.ExporterOptions{
		InstrKey: inskey,
	}
}

// NewJaegerExporterOptions provides jaeger exporter options
func NewJaegerExporterOptions(c *Initializer) *jaeger.ExporterOptions {
	endpoint := os.Getenv("JaegerEndpoint")
	lgr := c.lgrf.Create(context.Background())
	if endpoint == "" {
		lgr.Warn("missing jaeger endpoint")
	}
	return &jaeger.ExporterOptions{
		Endpoint: endpoint,
	}
}

// NewTraceOptions provides trace options
func NewTraceOptions(_ *Initializer) (*trace.Options, error) {
	cnf := &trace.Options{
		ServiceName: common.ServiceName,
	}
	return cnf, nil
}
