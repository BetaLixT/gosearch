// Package infra provides dependency set for the infrastructure layer
package infra

import (
	"github.com/BetaLixT/gex/pkg/domain/base/logger"
	domtrace "github.com/BetaLixT/gex/pkg/domain/base/trace"
	"github.com/BetaLixT/gex/pkg/infra/config"
	"github.com/BetaLixT/gex/pkg/infra/lgr"
	"github.com/BetaLixT/gex/pkg/infra/redisdb"
	"github.com/BetaLixT/gex/pkg/infra/roachdb"
	"github.com/BetaLixT/gex/pkg/infra/snowflake"
	"github.com/BetaLixT/gex/pkg/infra/trace"
	"github.com/BetaLixT/gex/pkg/infra/trace/appinsights"
	"github.com/BetaLixT/gex/pkg/infra/trace/jaeger"
	"github.com/BetaLixT/gex/pkg/infra/trace/promex"
	"github.com/BetaLixT/gex/pkg/infra/tracelib"

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
