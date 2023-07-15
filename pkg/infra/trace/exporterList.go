package trace

import (
	"context"

	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"

	"github.com/BetaLixT/gosearch/pkg/infra/trace/appinsights"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/jaeger"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/logex"
	"github.com/BetaLixT/gosearch/pkg/infra/trace/promex"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// NewTraceExporterList provides a list of exporters for tracing
func NewTraceExporterList(
	insexp appinsights.TraceExporter,
	jgrexp jaeger.TraceExporter,
	lgexp logex.TraceExporter,
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
		lgr.Warn("not tracing exporters found, console trace exporter will be used")
		exp = append(exp, lgexp)
	}
	exp = append(exp, prmex)
	return &ExporterList{
		Exporters: exp,
	}
}
