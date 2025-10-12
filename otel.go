package stripe

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

// tracerProvider is the global OpenTelemetry TracerProvider for Stripe instrumentation.
// Set this using SetTracerProvider to enable OpenTelemetry tracing.
var tracerProvider trace.TracerProvider

// tracer is the OpenTelemetry tracer instance for stripe-go
var tracer trace.Tracer

const (
	instrumentationName    = "github.com/stripe/stripe-go/v83"
	instrumentationVersion = clientversion
)

// SetTracerProvider sets the OpenTelemetry TracerProvider for Stripe instrumentation.
// Set to nil to disable OpenTelemetry tracing.
//
// Example:
//
//	import "go.opentelemetry.io/otel"
//	stripe.SetTracerProvider(otel.GetTracerProvider())
func SetTracerProvider(tp trace.TracerProvider) {
	tracerProvider = tp
	if tp != nil {
		tracer = tp.Tracer(instrumentationName, trace.WithInstrumentationVersion(instrumentationVersion))
	} else {
		tracer = nil
	}
}

// GetTracerProvider returns the configured OpenTelemetry TracerProvider.
func GetTracerProvider() trace.TracerProvider {
	return tracerProvider
}

// otelSpanInfo holds information needed to create an OpenTelemetry span
type otelSpanInfo struct {
	method      string
	url         string
	host        string
	urlTemplate string
	backend     SupportedBackend
}

// startHTTPSpan creates an OpenTelemetry span for an HTTP request
func startHTTPSpan(ctx context.Context, info otelSpanInfo) (context.Context, trace.Span) {
	if tracer == nil {
		return ctx, trace.SpanFromContext(ctx)
	}

	spanName := fmt.Sprintf("HTTP %s", info.method)

	attrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String(info.method),
		semconv.URLFull(info.url),
		semconv.ServerAddress(info.host),
		attribute.String("stripe.api_version", APIVersion),
		attribute.String("stripe.backend", string(info.backend)),
	}

	// Add URL template if available (preferred for cardinality reasons)
	if info.urlTemplate != "" {
		attrs = append(attrs, attribute.String("url.template", info.urlTemplate))
	}

	ctx, span := tracer.Start(ctx, spanName,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(attrs...),
	)

	return ctx, span
}

// endHTTPSpan updates the span with response information and ends it
func endHTTPSpan(span trace.Span, res *http.Response, err error) {
	if span == nil || !span.IsRecording() {
		return
	}

	defer span.End()

	if res != nil {
		span.SetAttributes(
			semconv.HTTPResponseStatusCode(res.StatusCode),
			attribute.String("stripe.request_id", res.Header.Get("Request-Id")),
		)
	}

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	} else {
		span.SetStatus(codes.Ok, "")
	}
}

// extractURLTemplate gets the URL template from params if available
func extractURLTemplate(params *Params) string {
	if params != nil && params.URLTemplate != nil {
		return *params.URLTemplate
	}
	return ""
}

// contextWithURLTemplate adds the URL template to the context for span creation
func contextWithURLTemplate(ctx context.Context, template string) context.Context {
	if template == "" {
		return ctx
	}
	type urlTemplateKey struct{}
	return context.WithValue(ctx, urlTemplateKey{}, template)
}

// urlTemplateFromContext extracts the URL template from context
func urlTemplateFromContext(ctx context.Context) string {
	type urlTemplateKey struct{}
	if val := ctx.Value(urlTemplateKey{}); val != nil {
		if template, ok := val.(string); ok {
			return template
		}
	}
	return ""
}
