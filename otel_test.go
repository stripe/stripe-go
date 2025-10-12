package stripe

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func TestSetTracerProvider(t *testing.T) {
	// Test setting tracer provider
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))

	SetTracerProvider(tp)
	assert.NotNil(t, tracer)
	assert.Equal(t, tp, tracerProvider)

	// Test setting to nil
	SetTracerProvider(nil)
	assert.Nil(t, tracer)
	assert.Nil(t, tracerProvider)
}

func TestGetTracerProvider(t *testing.T) {
	// Initially should be nil
	assert.Nil(t, GetTracerProvider())

	// Set a provider
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	SetTracerProvider(tp)

	// Should return the same provider
	assert.Equal(t, tp, GetTracerProvider())

	// Clean up
	SetTracerProvider(nil)
}

func TestStartHTTPSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	SetTracerProvider(tp)
	defer SetTracerProvider(nil)

	ctx := context.Background()
	info := otelSpanInfo{
		method:      "POST",
		url:         "https://api.stripe.com/v1/charges/ch_123",
		host:        "api.stripe.com",
		urlTemplate: "/v1/charges/{id}",
		backend:     APIBackend,
	}

	ctx, span := startHTTPSpan(ctx, info)
	assert.NotNil(t, span)
	assert.True(t, span.IsRecording())

	span.End()

	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Equal(t, "HTTP POST", spans[0].Name)

	// Check attributes
	attrs := spans[0].Attributes
	assertHasAttribute(t, attrs, string(semconv.HTTPRequestMethodKey), "POST")
	assertHasAttribute(t, attrs, "url.template", "/v1/charges/{id}")
	assertHasAttribute(t, attrs, "stripe.backend", "api")
}

func TestStartHTTPSpanWithoutTracer(t *testing.T) {
	// Ensure tracer is nil
	SetTracerProvider(nil)

	ctx := context.Background()
	info := otelSpanInfo{
		method:      "GET",
		url:         "https://api.stripe.com/v1/charges",
		host:        "api.stripe.com",
		urlTemplate: "/v1/charges",
		backend:     APIBackend,
	}

	ctx, span := startHTTPSpan(ctx, info)
	// Should return a non-recording span
	assert.NotNil(t, span)
	assert.False(t, span.IsRecording())
}

func TestEndHTTPSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	SetTracerProvider(tp)
	defer SetTracerProvider(nil)

	ctx := context.Background()
	info := otelSpanInfo{
		method:      "GET",
		url:         "https://api.stripe.com/v1/charges",
		host:        "api.stripe.com",
		urlTemplate: "",
		backend:     APIBackend,
	}

	ctx, span := startHTTPSpan(ctx, info)

	// Simulate a successful response
	// Since we don't have a real http.Response, we'll test the nil case
	endHTTPSpan(span, nil, nil)

	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
}

func TestExtractURLTemplate(t *testing.T) {
	// Test with nil params
	assert.Equal(t, "", extractURLTemplate(nil))

	// Test with params but no URLTemplate
	params := &Params{}
	assert.Equal(t, "", extractURLTemplate(params))

	// Test with URLTemplate set
	template := "/v1/charges/{id}"
	params.URLTemplate = &template
	assert.Equal(t, "/v1/charges/{id}", extractURLTemplate(params))
}

func TestContextWithURLTemplate(t *testing.T) {
	ctx := context.Background()

	// Test with empty template
	newCtx := contextWithURLTemplate(ctx, "")
	assert.Equal(t, ctx, newCtx)

	// Test with valid template
	template := "/v1/charges/{id}"
	newCtx = contextWithURLTemplate(ctx, template)
	assert.NotEqual(t, ctx, newCtx)

	// Verify we can extract it back
	extracted := urlTemplateFromContext(newCtx)
	assert.Equal(t, template, extracted)
}

func TestURLTemplateFromContext(t *testing.T) {
	ctx := context.Background()

	// Test with context that doesn't have template
	assert.Equal(t, "", urlTemplateFromContext(ctx))

	// Test with template in context
	template := "/v1/customers/{customer_id}"
	ctx = contextWithURLTemplate(ctx, template)
	assert.Equal(t, template, urlTemplateFromContext(ctx))
}

func TestFormatURLPathWithTemplate(t *testing.T) {
	params := &Params{}

	// Test single parameter
	path := FormatURLPathWithTemplate(params, "/v1/charges/{id}", "ch_123")
	assert.Equal(t, "/v1/charges/ch_123", path)
	assert.NotNil(t, params.URLTemplate)
	assert.Equal(t, "/v1/charges/{id}", *params.URLTemplate)

	// Test multiple parameters
	params2 := &Params{}
	path2 := FormatURLPathWithTemplate(params2, "/v1/customers/{customer_id}/sources/{id}", "cus_123", "card_456")
	assert.Equal(t, "/v1/customers/cus_123/sources/card_456", path2)
	assert.Equal(t, "/v1/customers/{customer_id}/sources/{id}", *params2.URLTemplate)

	// Test with nil params (should not panic)
	path3 := FormatURLPathWithTemplate(nil, "/v1/charges/{id}", "ch_123")
	assert.Equal(t, "/v1/charges/ch_123", path3)

	// Test panic on parameter mismatch
	assert.Panics(t, func() {
		FormatURLPathWithTemplate(&Params{}, "/v1/charges/{id}", "ch_123", "extra")
	})

	// Test panic on too few parameters
	assert.Panics(t, func() {
		FormatURLPathWithTemplate(&Params{}, "/v1/customers/{customer_id}/sources/{id}", "cus_123")
	})
}

// Helper function to check if an attribute exists with a given value
func assertHasAttribute(t *testing.T, attrs []attribute.KeyValue, key string, expected string) {
	for _, attr := range attrs {
		if string(attr.Key) == key {
			assert.Equal(t, expected, attr.Value.AsString())
			return
		}
	}
	t.Errorf("Attribute %s not found", key)
}
