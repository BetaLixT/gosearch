// Package infra provides dependency set for the infrastructure layer
package infra

import (
	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"
	domtrace "github.com/BetaLixT/gosearch/pkg/domain/base/trace"
	"github.com/BetaLixT/gosearch/pkg/infra/config"
	"github.com/BetaLixT/gosearch/pkg/infra/lgr"
	"github.com/BetaLixT/gosearch/pkg/infra/redisdb"
	"github.com/BetaLixT/gosearch/pkg/infra/roachdb"
	"github.com/BetaLixT/gosearch/pkg/infra/snowflake"
	"github.com/BetaLixT/gosearch/pkg/infra/trace"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/appinsights"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/jaeger"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/promex"
	"github.com/BetaLixT/gosearch/pkg/infra/tracelib"

	"github.com/BetaLixT/gotred/v8"
	"github.com/BetaLixT/tsqlx"
	"github.com/google/wire"
)

// DependencySet dependency set for infra layer
var DependencySet = wire.NewSet(
	// Trace
	trace.NewTraceExporterList,
	config.NewTraceOptions,
	trace.NewTracer,
	jaeger.NewJaegerTraceExporter,
	config.NewJaegerExporterOptions,
	appinsights.NewTraceExporter,
	config.NewAppInsightsExporterOptions,
	promex.NewTraceExporter,

	config.NewInitializer,
	lgr.NewLoggerFactory,
	wire.Bind(
		new(logger.IFactory),
		new(*lgr.LoggerFactory),
	),

	// Databases
	roachdb.NewDatabaseContext,
	wire.Bind(
		new(tsqlx.ITracer),
		new(*tracelib.Tracer),
	),
	config.NewRoachDBOptions,
	redisdb.NewRedisContext,
	wire.Bind(
		new(gotred.ITracer),
		new(*tracelib.Tracer),
	),
	config.NewRedisOptions,
	snowflake.NewSnowflake,
	config.NewSnowflakeOptions,
	wire.Bind(
		new(domtrace.IRepository),
		new(*tracelib.Tracer),
	),
)
