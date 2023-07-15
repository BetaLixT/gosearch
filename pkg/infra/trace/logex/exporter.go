// Package appinsights provides an app insights exporter for tracing
package logex

import (
	"context"
	"errors"
	"sync"

	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	trace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TraceExporter an application insights exporter
type TraceExporter sdktrace.SpanExporter

// NewTraceExporter constructs an app insights exporter
func New(lgrf logger.IFactory) TraceExporter {
	return &LoggerExporter{
		lgrf.Create(context.Background()),
		&sync.RWMutex{},
		false,
	}
}

type LoggerExporter struct {
	lgr    *zap.Logger
	mtx    *sync.RWMutex
	closed bool
}

// Exports an array of Open Telemetry spans to Application Insights
func (exp *LoggerExporter) ExportSpans(
	ctx context.Context,
	spans []sdktrace.ReadOnlySpan,
) error {
	exp.mtx.RLock()
	defer exp.mtx.RUnlock()

	if exp.closed {
		return errors.New("exporter closed")
	}

	for i := range spans {
		exp.process(spans[i])
	}
	return nil
}

// Exports an array of Open Telemetry spans to Application Insights
func (exp *LoggerExporter) Shutdown(
	ctx context.Context,
) error {
	exp.mtx.Lock()
	defer exp.mtx.Unlock()
	exp.closed = true

	return exp.lgr.Sync()
}

// Preprocesses the Otel span and dispatches it to app insights differently
// based on the span kind.
func (exp *LoggerExporter) process(sp sdktrace.ReadOnlySpan) {
	success := true
	if sp.Status().Code != codes.Ok {
		success = false
	}

	props := map[string]string{}

	rattr := sp.Resource().Attributes()
	for _, e := range rattr {
		props[string(e.Key)] = e.Value.AsString()
	}
	attr := sp.Attributes()
	for _, e := range attr {
		props[string(e.Key)] = e.Value.AsString()
	}

	msg := "trace"
	switch sp.SpanKind() {
	case trace.SpanKindUnspecified, trace.SpanKindInternal:
		msg = "trace internal process"
	case trace.SpanKindServer:
		msg = "trace request"
	case trace.SpanKindClient, trace.SpanKindProducer:
		msg = "trace dependency"
	case trace.SpanKindConsumer:
		msg = "trace event"
	}

	fields := make([]zap.Field, 0, 8+len(props))
	fields = append(
		fields,
		zap.String("name", sp.Name()),
		zap.Bool("success", success),
		zap.Time("startTime", sp.StartTime()),
		zap.Time("endTime", sp.EndTime()),
		zap.Duration("elapsed", sp.EndTime().Sub(sp.StartTime())),
		zap.String("tid", sp.SpanContext().TraceID().String()),
		zap.String("pid", sp.Parent().SpanID().String()),
		zap.String("rid", sp.SpanContext().TraceID().String()),
	)
	for key := range props {
		fields = append(fields, zap.String(key, props[key]))
	}

	exp.lgr.Info(
		msg,
		fields...,
	)
}
