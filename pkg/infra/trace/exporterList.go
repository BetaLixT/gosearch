package trace

import (
	"context"

	"github.com/BetaLixT/gex/pkg/domain/base/logger"

	"github.com/BetaLixT/gex/pkg/infra/trace/appinsights"
	"github.com/BetaLixT/gex/pkg/infra/trace/jaeger"
	"github.com/BetaLixT/gex/pkg/infra/trace/promex"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// NewTraceExporterList provides a list of exporters for tracing
func NewTraceExporterList(
	insexp appinsights.TraceExporter,
	jgrexp jaeger.TraceExporter,
	prmex promex.TraceExporter,
	lgrf logger.IFactory,
) *ExporterList {
	lgr := lgrf.Create(context.Background())
	exp := []sdktrace.SpanExporter{}

	if insexp != nil {
		exp = append(exp, insexp)
	} else {
		lgr.Warn("insights exporter not found")
	}
	if jgrexp != nil {
		exp = append(exp, jgrexp)
	} else {
		lgr.Warn("jeager exporter not found")
	}
	if len(exp) == 0 {
		panic("no tracing exporters found (float you <3)")
	}
	exp = append(exp, prmex)
	return &ExporterList{
		Exporters: exp,
	}
}
