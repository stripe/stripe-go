package stripe

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// Test that BackendImplementation uses ContextLeveledLoggerInterface
func TestBackendUsesContextLogger(t *testing.T) {
	var stdout, stderr bytes.Buffer
	contextLogger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelInfo,
		},
	}

	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     contextLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)
	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-123")

	// Verify context logger is used and context is passed
	backend.logInfof(ctx, "test message")
	assert.Equal(t, "[INFO] test message\n", stdout.String())
	assert.Equal(t, ctx, contextLogger.lastContext, "Context should be passed to logger")
}

// Test that BackendImplementation works with regular LeveledLoggerInterface
func TestBackendWorksWithRegularLogger(t *testing.T) {
	var stdout, stderr bytes.Buffer
	regularLogger := &LeveledLogger{
		stdoutOverride: &stdout,
		stderrOverride: &stderr,
		Level:          LevelInfo,
	}

	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     regularLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)

	// Verify logging still works with regular logger
	backend.logInfof(context.Background(), "test message")
	assert.Equal(t, "[INFO] test message\n", stdout.String(), "Regular logger should still work")
}

// Test context propagation through helper methods
func TestContextPropagationThroughHelpers(t *testing.T) {
	var stdout, stderr bytes.Buffer
	contextLogger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelDebug,
		},
	}

	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     contextLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)

	// Create context with test value
	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-trace-123")

	// Test all helper methods
	backend.logDebugf(ctx, "debug message")
	assert.Equal(t, ctx, contextLogger.lastContext, "logDebugf should pass context")
	clearBuffers(&stdout, &stderr)

	backend.logInfof(ctx, "info message")
	assert.Equal(t, ctx, contextLogger.lastContext, "logInfof should pass context")
	clearBuffers(&stdout, &stderr)

	backend.logWarnf(ctx, "warn message")
	assert.Equal(t, ctx, contextLogger.lastContext, "logWarnf should pass context")
	clearBuffers(&stdout, &stderr)

	backend.logErrorf(ctx, "error message")
	assert.Equal(t, ctx, contextLogger.lastContext, "logErrorf should pass context")
}

// Test that regular logger fallback works
func TestRegularLoggerFallback(t *testing.T) {
	var stdout, stderr bytes.Buffer
	regularLogger := &LeveledLogger{
		stdoutOverride: &stdout,
		stderrOverride: &stderr,
		Level:          LevelInfo,
	}

	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     regularLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)

	// Create context - should be ignored since logger doesn't support it
	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-trace-456")

	// Call helper methods - should fall back to regular logger
	backend.logInfof(ctx, "test message")

	// Verify logging still works (without context support)
	assert.Equal(t, "[INFO] test message\n", stdout.String(), "Should fall back to regular logger")
}

// Test logError helper with context
func TestLogErrorWithContext(t *testing.T) {
	var stdout, stderr bytes.Buffer
	contextLogger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelInfo, // Need Info level to capture 402 status logging
		},
	}

	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     contextLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)
	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "error-trace")

	// Create a Stripe error
	stripeErr := &Error{
		Type:           ErrorTypeInvalidRequest,
		Msg:            "Test error",
		HTTPStatusCode: 400,
	}

	// Test error logging with 400 status
	backend.logError(ctx, 400, stripeErr)
	assert.Contains(t, stderr.String(), "Request error from Stripe")
	assert.Equal(t, ctx, contextLogger.lastContext, "Context should be passed to error logger")
	clearBuffers(&stdout, &stderr)

	// Test error logging with 402 status (should use Info level)
	backend.logError(ctx, 402, stripeErr)
	assert.Contains(t, stdout.String(), "User-compelled request error")
	assert.Equal(t, ctx, contextLogger.lastContext, "Context should be passed to info logger for 402")
}

// Test backward compatibility: old code without context should still work
func TestBackwardCompatibility(t *testing.T) {
	var stdout, stderr bytes.Buffer

	// Test 1: Regular logger (old interface) still works
	regularLogger := &LeveledLogger{
		stdoutOverride: &stdout,
		stderrOverride: &stderr,
		Level:          LevelInfo,
	}

	// Old code path - directly using LeveledLogger
	regularLogger.Infof("old style logging")
	assert.Equal(t, "[INFO] old style logging\n", stdout.String(), "Old LeveledLogger should work unchanged")
	clearBuffers(&stdout, &stderr)

	// Test 2: Backend with regular logger works
	config := &BackendConfig{
		HTTPClient:        &http.Client{},
		LeveledLogger:     regularLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)

	// Old code using backend (internally uses helper methods but with regular logger)
	// Cast LeveledLogger back to LeveledLoggerInterface for direct access
	if logger, ok := backend.LeveledLogger.(LeveledLoggerInterface); ok {
		logger.Infof("backend logging")
	}
	assert.Equal(t, "[INFO] backend logging\n", stdout.String(), "Backend with regular logger should work")
}

// mockRoundTripper is a mock HTTP transport for testing
type mockRoundTripper struct {
	response *http.Response
	err      error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.response != nil {
		// Create a response with a body
		body := io.NopCloser(bytes.NewBufferString(`{"id":"test_123"}`))
		return &http.Response{
			StatusCode: 200,
			Body:       body,
			Header:     http.Header{},
			Request:    req,
		}, nil
	}
	return m.response, m.err
}

// Test context propagation in a real request scenario
func TestContextPropagationInRequest(t *testing.T) {
	var stdout, stderr bytes.Buffer
	contextLogger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelInfo,
		},
	}

	// Create backend with context logger
	httpClient := &http.Client{
		Transport: &mockRoundTripper{
			response: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(`{"id":"cus_123"}`)),
				Header:     http.Header{},
			},
		},
	}

	config := &BackendConfig{
		HTTPClient:        httpClient,
		LeveledLogger:     contextLogger,
		MaxNetworkRetries: Int64(0),
		URL:               String("https://api.stripe.com"),
	}

	backend := newBackendImplementation(APIBackend, config).(*BackendImplementation)

	// Create a request with context
	ctx := context.WithValue(context.Background(), contextKey("request_id"), "req-789")
	params := &Params{Context: ctx}

	req, err := backend.NewRequest(http.MethodGet, "/v1/customers/cus_123", "sk_test_123", "application/x-www-form-urlencoded", params)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	// Verify context was attached to request
	assert.Equal(t, ctx, req.Context(), "Context should be attached to http.Request")

	// Simulate logging that would happen during request
	backend.logInfof(req.Context(), "Requesting GET /v1/customers")

	// Verify context was passed to logger
	assert.Equal(t, ctx, contextLogger.lastContext, "Request context should be passed to logger")
}
