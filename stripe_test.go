package stripe

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

// A shortcut for a leveled logger that spits out all debug information (useful in tests).
var debugLeveledLogger = &LeveledLogger{
	Level: LevelDebug,
}

// For tests that produce a lot of logging or alarming error logs on a
// successful run (thereby making `go test . -test.v` quite noisy), use this
// null leveled logger instead of the debug one above.
var nullLeveledLogger = &LeveledLogger{
	Level: LevelNull,
}

//
// ---
//

func TestBearerAuth(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
}

func TestContext(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	p := &Params{Context: context.Background()}

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	// We assume that contexts are sufficiently tested in the standard library
	// and here we just check that the context sent in to `NewRequest` is
	// indeed properly set on the request that's returned.
	assert.Equal(t, p.Context, req.Context())
}

// Tests client retries.
//
// You can get pretty good visibility into what's going on by running just this
// test on verbose:
//
//     go test . -run TestDo_Retry -test.v
//
func TestDo_Retry(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."
	requestNum := 0

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)

		// The body should always be the same with every retry. We've
		// previously had regressions in this behavior as we switched to HTTP/2
		// and `Request` became non-reusable, so we want to check it with every
		// request.
		assert.Equal(t, "bar", r.Form.Get("foo"))

		switch requestNum {
		case 0:
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"error":"Conflict (this should be retried)."}`))

		case 1:
			response := testServerResponse{Message: message}

			data, err := json.Marshal(response)
			assert.NoError(t, err)

			_, err = w.Write(data)
			assert.NoError(t, err)

		default:
			assert.Fail(t, "Should not have reached request %v", requestNum)
		}

		requestNum++
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     nullLeveledLogger,
			MaxNetworkRetries: Int64(5),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	// Disable sleeping duration our tests.
	backend.SetNetworkRetriesSleep(false)

	request, err := backend.NewRequest(
		http.MethodPost,
		"/hello",
		"sk_test_123",
		"application/x-www-form-urlencoded",
		nil,
	)
	assert.NoError(t, err)

	bodyBuffer := bytes.NewBufferString("foo=bar")
	var response testServerResponse
	err = backend.Do(request, bodyBuffer, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)

	// We should have seen exactly two requests.
	assert.Equal(t, 2, requestNum)
}

func TestShouldRetry(t *testing.T) {
	MaxNetworkRetries := int64(3)

	c := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			MaxNetworkRetries: Int64(MaxNetworkRetries),
		},
	).(*BackendImplementation)

	// Exceeded maximum number of retries
	t.Run("DontRetryOnExceededRetries", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{},
			int(MaxNetworkRetries),
		)
		assert.False(t, shouldRetry)
	})

	// Canceled context -- don't retry
	t.Run("DontRetryOnCanceledContext", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", nil)
		assert.NoError(t, err)

		shouldRetry, _ := c.shouldRetry(
			nil,
			req,
			&http.Response{StatusCode: http.StatusOK},
			0,
		)
		assert.False(t, shouldRetry)
	})

	// Doesn't retry most Stripe errors (they must also match a status code
	// below to be retried)
	t.Run("DontRetryOnStripeError", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			&Error{Msg: "An error from Stripe"},
			&http.Request{},
			&http.Response{StatusCode: http.StatusBadRequest},
			0,
		)
		assert.False(t, shouldRetry)
	})

	// Don't retry too many redirects.
	t.Run("DontRetryOnTooManyRedirects", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			&url.Error{Err: fmt.Errorf("stopped after 5 redirects")},
			&http.Request{},
			nil,
			0,
		)
		assert.False(t, shouldRetry)
	})

	// Don't retry invalid protocol scheme.
	t.Run("DontRetryOnInvalidProtocolScheme", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			&url.Error{Err: fmt.Errorf("unsupported protocol scheme")},
			&http.Request{},
			nil,
			0,
		)
		assert.False(t, shouldRetry)
	})

	// Don't retry TLS certificate validation problems.
	t.Run("DontRetryOnCertificateError", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			&url.Error{Err: x509.UnknownAuthorityError{}},
			&http.Request{},
			nil,
			0,
		)
		assert.False(t, shouldRetry)
	})

	// Retries most non-Stripe errors
	t.Run("RetryOnNonStripeError", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			fmt.Errorf("an error"),
			&http.Request{},
			nil,
			0,
		)
		assert.True(t, shouldRetry)
	})

	// `Stripe-Should-Retry: false`
	t.Run("DontRetryOnStripeRetryHeaderFalse", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{
				Header: http.Header(map[string][]string{
					"Stripe-Should-Retry": {"false"},
				}),
				// Note we send status 409 here, which would normally be retried
				StatusCode: http.StatusConflict,
			},
			0,
		)
		assert.False(t, shouldRetry)
	})

	// `Stripe-Should-Retry: true`
	t.Run("RetryOnStripeRetryHeaderTrue", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{
				Header: http.Header(map[string][]string{
					"Stripe-Should-Retry": {"true"},
				}),
				// Note we send status 400 here, which would normally not be
				// retried
				StatusCode: http.StatusBadRequest,
			},
			0,
		)
		assert.True(t, shouldRetry)
	})

	// 409 Conflict
	t.Run("RetryOn409Conflict", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{StatusCode: http.StatusConflict},
			0,
		)
		assert.True(t, shouldRetry)
	})

	// 429 Too Many Requests -- retry on lock timeout
	t.Run("RetryOn429TooManyRequestsLockTimeout", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			&Error{Code: ErrorCodeLockTimeout},
			&http.Request{},
			&http.Response{StatusCode: http.StatusTooManyRequests},
			0,
		)
		assert.True(t, shouldRetry)
	})

	// 429 Too Many Requests -- don't retry normally
	t.Run("DontRetryOn429TooManyRequests", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{StatusCode: http.StatusTooManyRequests},
			0,
		)
		assert.False(t, shouldRetry)
	})

	// 500 Internal Server Error -- retry if non-POST
	t.Run("RetryOn500NonPost", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{Method: http.MethodGet},
			&http.Response{StatusCode: http.StatusInternalServerError},
			0,
		)
		assert.True(t, shouldRetry)
	})

	// 500 Internal Server Error -- don't retry POST
	t.Run("DontRetryOn500Post", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{Method: http.MethodPost},
			&http.Response{StatusCode: http.StatusInternalServerError},
			0,
		)
		assert.False(t, shouldRetry)
	})

	// 503 Service Unavailable
	t.Run("RetryOn503ServiceUnavailable", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{},
			&http.Response{StatusCode: http.StatusServiceUnavailable},
			0,
		)
		assert.True(t, shouldRetry)
	})
}

func TestDo_RetryOnTimeout(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	timeout := time.Second
	var counter uint32

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint32(&counter, 1)
		time.Sleep(timeout)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     nullLeveledLogger,
			MaxNetworkRetries: Int64(1),
			URL:               String(testServer.URL),
			HTTPClient:        &http.Client{Timeout: timeout},
		},
	).(*BackendImplementation)

	backend.SetNetworkRetriesSleep(false)

	request, err := backend.NewRequest(
		http.MethodPost,
		"/hello",
		"sk_test_123",
		"application/x-www-form-urlencoded",
		nil,
	)
	assert.NoError(t, err)

	var body = bytes.NewBufferString("foo=bar")
	var response testServerResponse

	err = backend.Do(request, body, &response)

	assert.Error(t, err)
	// timeout should not prevent retry
	assert.Equal(t, uint32(2), atomic.LoadUint32(&counter))
}

func TestDo_LastResponsePopulated(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."
	expectedResponse := testServerResponse{Message: message}
	rawJSON, err := json.Marshal(expectedResponse)
	assert.NoError(t, err)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Idempotency-Key", "key_123")
		w.Header().Set("Other-Header", "other_header")
		w.Header().Set("Request-Id", "req_123")

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(rawJSON)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	request, err := backend.NewRequest(
		http.MethodGet,
		"/hello",
		"sk_test_123",
		"application/x-www-form-urlencoded",
		nil,
	)
	assert.NoError(t, err)

	var resource testServerResponse
	err = backend.Do(request, nil, &resource)
	assert.NoError(t, err)
	assert.Equal(t, message, resource.Message)

	assert.Equal(t, "key_123", resource.LastResponse.IdempotencyKey)
	assert.Equal(t, "other_header", resource.LastResponse.Header.Get("Other-Header"))
	assert.Equal(t, rawJSON, resource.LastResponse.RawJSON)
	assert.Equal(t, "req_123", resource.LastResponse.RequestID)
	assert.Equal(t,
		fmt.Sprintf("%v %v", http.StatusCreated, http.StatusText(http.StatusCreated)),
		resource.LastResponse.Status)
	assert.Equal(t, http.StatusCreated, resource.LastResponse.StatusCode)
}

// Test that telemetry metrics are not sent by default
func TestDo_TelemetryDisabled(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."
	requestNum := 0

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// none of the requests should include telemetry metrics
		assert.Equal(t, r.Header.Get("X-Stripe-Client-Telemetry"), "")

		response := testServerResponse{Message: message}

		data, err := json.Marshal(response)
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)

		requestNum++
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	// When telemetry is enabled, the metrics for a request are sent with the
	// _next_ request via the `X-Stripe-Client-Telemetry header`. To test that
	// metrics aren't being sent, we need to fire off two requests in sequence.
	for i := 0; i < 2; i++ {
		request, err := backend.NewRequest(
			http.MethodGet,
			"/hello",
			"sk_test_123",
			"application/x-www-form-urlencoded",
			nil,
		)
		assert.NoError(t, err)

		var response testServerResponse
		err = backend.Do(request, nil, &response)

		assert.NoError(t, err)
		assert.Equal(t, message, response.Message)
	}

	// We should have seen exactly two requests.
	assert.Equal(t, 2, requestNum)
}

// Test that telemetry metrics are sent on subsequent requests when
// EnableTelemetry = true.
func TestDo_TelemetryEnabled(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	type requestMetrics struct {
		RequestDurationMS int    `json:"request_duration_ms"`
		RequestID         string `json:"request_id"`
	}

	type requestTelemetry struct {
		LastRequestMetrics requestMetrics `json:"last_request_metrics"`
	}

	message := "Hello, client."
	requestNum := 0

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestNum++

		telemetryStr := r.Header.Get("X-Stripe-Client-Telemetry")
		switch requestNum {
		case 1:
			// the first request should not receive any metrics
			assert.Equal(t, telemetryStr, "")
			time.Sleep(21 * time.Millisecond)
		case 2:
			assert.True(t, len(telemetryStr) > 0, "telemetryStr should not be empty")

			// the telemetry should properly unmarshal into RequestTelemetry
			var telemetry requestTelemetry
			err := json.Unmarshal([]byte(telemetryStr), &telemetry)
			assert.NoError(t, err)

			// the second request should include the metrics for the first request
			assert.Equal(t, telemetry.LastRequestMetrics.RequestID, "req_1")
			assert.True(t, telemetry.LastRequestMetrics.RequestDurationMS > 20,
				"request_duration_ms should be > 20ms")
		default:
			assert.Fail(t, "Should not have reached request %v", requestNum)
		}

		w.Header().Set("Request-Id", fmt.Sprintf("req_%d", requestNum))
		response := testServerResponse{Message: message}

		data, err := json.Marshal(response)
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	for i := 0; i < 2; i++ {
		request, err := backend.NewRequest(
			http.MethodGet,
			"/hello",
			"sk_test_123",
			"application/x-www-form-urlencoded",
			nil,
		)
		assert.NoError(t, err)

		var response testServerResponse
		err = backend.Do(request, nil, &response)

		assert.NoError(t, err)
		assert.Equal(t, message, response.Message)
	}

	// We should have seen exactly two requests.
	assert.Equal(t, 2, requestNum)
}

// This test does not perform any super valuable assertions - instead, it checks
// that our logic for buffering requestMetrics when EnableTelemetry = true does
// not trigger any data races. This test should pass when the -race flag is
// passed to `go test`.
func TestDo_TelemetryEnabledNoDataRace(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."
	var requestNum int32

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := atomic.AddInt32(&requestNum, 1)

		w.Header().Set("Request-Id", fmt.Sprintf("req_%d", reqID))
		response := testServerResponse{Message: message}

		data, err := json.Marshal(response)
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     nullLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	times := 20 // 20 > telemetryBufferSize, so some metrics could be discarded
	done := make(chan struct{})

	for i := 0; i < times; i++ {
		go func() {
			request, err := backend.NewRequest(
				http.MethodGet,
				"/hello",
				"sk_test_123",
				"application/x-www-form-urlencoded",
				nil,
			)
			assert.NoError(t, err)

			var response testServerResponse
			err = backend.Do(request, nil, &response)

			assert.NoError(t, err)
			assert.Equal(t, message, response.Message)

			done <- struct{}{}
		}()
	}

	for i := 0; i < times; i++ {
		<-done
	}

	assert.Equal(t, int32(times), requestNum)
}

func TestDo_Redaction(t *testing.T) {
	type testServerResponse struct {
		Error *Error `json:"error"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(402)
		data, err := json.Marshal(testServerResponse{Error: &Error{PaymentIntent: &PaymentIntent{ClientSecret: "SHOULDBEREDACTED"}}})
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)

	}))
	defer testServer.Close()

	var logs bytes.Buffer
	logger := &LeveledLogger{Level: LevelDebug, stderrOverride: &logs, stdoutOverride: &logs}

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     logger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	request, err := backend.NewRequest(
		http.MethodGet,
		"/hello",
		"sk_test_123",
		"application/x-www-form-urlencoded",
		nil,
	)
	assert.NoError(t, err)

	var response Charge
	err = backend.Do(request, nil, &response)
	assert.Error(t, err)

	assert.NotContains(t, logs.String(), "SHOULDBEREDACTED")
	assert.Contains(t, logs.String(), "REDACTED")
}

func TestDoStreaming(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		data := []byte("hello")

		var err error
		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	var logs bytes.Buffer
	logger := &LeveledLogger{Level: LevelDebug, stderrOverride: &logs, stdoutOverride: &logs}

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     logger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	type streamingResource struct {
		APIStream
	}

	response := streamingResource{}
	err := backend.CallStreaming(
		http.MethodGet,
		"/pdf",
		"sk_test_123",
		nil,
		&response,
	)
	assert.NoError(t, err)
	result, err := ioutil.ReadAll(response.LastResponse.Body)
	assert.NoError(t, err)
	err = response.LastResponse.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(result))
}

func TestDoStreaming_ParsableError(t *testing.T) {
	type testServerResponse struct {
		Error *Error `json:"error"`
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		var data []byte
		var err error
		data, err = json.Marshal(testServerResponse{Error: &Error{Msg: "Text of error"}})
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	var logs bytes.Buffer
	logger := &LeveledLogger{Level: LevelDebug, stderrOverride: &logs, stdoutOverride: &logs}

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     logger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	type streamingResource struct {
		APIStream
	}

	response := streamingResource{}
	err := backend.CallStreaming(
		http.MethodGet,
		"/pdf",
		"sk_test_123",
		nil,
		&response,
	)
	assert.NotNil(t, err)
	stripeErr, ok := err.(*Error)
	assert.True(t, ok)
	assert.Equal(t, stripeErr.Msg, "Text of error")
}

func TestDoStreaming_UnparsableError(t *testing.T) {
	type testServerResponse struct {
		Error *Error `json:"error"`
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		var data []byte
		var err error
		data = []byte("{invalid json}")

		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	var logs bytes.Buffer
	logger := &LeveledLogger{Level: LevelDebug, stderrOverride: &logs, stdoutOverride: &logs}

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			EnableTelemetry:   Bool(true),
			LeveledLogger:     logger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	type streamingResource struct {
		APIStream
	}

	response := streamingResource{}
	err := backend.CallStreaming(
		http.MethodGet,
		"/pdf",
		"sk_test_123",
		nil,
		&response,
	)
	assert.NotNil(t, err)
	_, ok := err.(*Error)
	assert.False(t, ok)
	assert.True(t, strings.Contains(err.Error(), "Couldn't deserialize JSON"))
}

func TestFormatURLPath(t *testing.T) {
	assert.Equal(t, "/v1/resources/1/subresources/2",
		FormatURLPath("/v1/resources/%s/subresources/%s", "1", "2"))

	// Tests that each parameter is escaped for use in URLs
	assert.Equal(t, "/v1/resources/%25",
		FormatURLPath("/v1/resources/%s", "%"))
}

func TestGetBackendWithConfig_Loggers(t *testing.T) {
	leveledLogger := &LeveledLogger{}

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger: leveledLogger,
		},
	).(*BackendImplementation)

	assert.Equal(t, leveledLogger, backend.LeveledLogger)
}

func TestGetBackendWithConfig_TrimV1Suffix(t *testing.T) {
	{
		backend := GetBackendWithConfig(
			APIBackend,
			&BackendConfig{
				URL: String("https://api.com/v1"),
			},
		).(*BackendImplementation)

		// The `/v1` suffix has been stripped.
		assert.Equal(t, "https://api.com", backend.URL)
	}

	// Also support trimming a `/v1/` with an extra trailing slash which is
	// probably an often seen mistake.
	{
		backend := GetBackendWithConfig(
			APIBackend,
			&BackendConfig{
				URL: String("https://api.com/v1/"),
			},
		).(*BackendImplementation)

		assert.Equal(t, "https://api.com", backend.URL)
	}

	// No-op otherwise.
	{
		backend := GetBackendWithConfig(
			APIBackend,
			&BackendConfig{
				URL: String("https://api.com"),
			},
		).(*BackendImplementation)

		assert.Equal(t, "https://api.com", backend.URL)
	}
}

func TestParseID(t *testing.T) {
	// JSON string
	{
		id, ok := ParseID([]byte(`"ch_123"`))
		assert.Equal(t, "ch_123", id)
		assert.True(t, ok)
	}

	// JSON object
	{
		id, ok := ParseID([]byte(`{"id":"ch_123"}`))
		assert.Equal(t, "", id)
		assert.False(t, ok)
	}

	// Other JSON scalar (this should never be used, but check the results anyway)
	{
		id, ok := ParseID([]byte(`123`))
		assert.Equal(t, "", id)
		assert.False(t, ok)
	}

	// Edge case that should never happen; found via fuzzing
	{
		id, ok := ParseID([]byte(`"`))
		assert.Equal(t, "", id)
		assert.False(t, ok)
	}
}

// TestMultipleAPICalls will fail the test run if a race condition is thrown while running multiple NewRequest calls.
func TestMultipleAPICalls(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := GetBackend(APIBackend).(*BackendImplementation)
			key := "apiKey"

			req, err := c.NewRequest("", "", key, "", nil)
			assert.NoError(t, err)

			assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
		}()
	}
	wg.Wait()
}

func TestIdempotencyKey(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	p := &Params{IdempotencyKey: String("idempotency-key")}

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, "idempotency-key", req.Header.Get("Idempotency-Key"))
}

func TestNewBackends(t *testing.T) {
	httpClient := &http.Client{}
	backends := NewBackends(httpClient)
	assert.Equal(t, httpClient, backends.API.(*BackendImplementation).HTTPClient)
	assert.Equal(t, httpClient, backends.Uploads.(*BackendImplementation).HTTPClient)
}

func TestStripeAccount(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	p := &Params{}
	p.SetStripeAccount("acct_123")

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, "acct_123", req.Header.Get("Stripe-Account"))
}

func TestUnmarshalJSONVerbose(t *testing.T) {
	type testServerResponse struct {
		Message string `json:"message"`
	}

	backend := GetBackend(APIBackend).(*BackendImplementation)

	// Valid JSON
	{
		type testServerResponse struct {
			Message string `json:"message"`
		}

		var sample testServerResponse
		err := backend.UnmarshalJSONVerbose(200, []byte(`{"message":"hello"}`), &sample)
		assert.NoError(t, err)
		assert.Equal(t, "hello", sample.Message)
	}

	// Invalid JSON (short)
	{
		body := `server error`

		var sample testServerResponse
		err := backend.UnmarshalJSONVerbose(200, []byte(body), &sample)
		assert.Regexp(t,
			fmt.Sprintf(`^Couldn't deserialize JSON \(response status: 200, body sample: '%s'\): invalid character`, body),
			err)
	}

	// Invalid JSON (long, and therefore truncated)
	{
		// Assembles a body that's at least as long as our maximum sample.
		// body is ~130 characters * 5.
		bodyText := `this is a really long body that will be truncated when added to the error message to protect against dumping huge responses in logs `
		body := bodyText + bodyText + bodyText + bodyText + bodyText

		var sample testServerResponse
		err := backend.UnmarshalJSONVerbose(200, []byte(body), &sample)
		assert.Regexp(t,
			fmt.Sprintf(`^Couldn't deserialize JSON \(response status: 200, body sample: '%s ...'\): invalid character`, body[0:500]),
			err)
	}
}

func TestUserAgent(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9]$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)
}

func TestUserAgentWithAppInfo(t *testing.T) {
	appInfo := &AppInfo{
		Name:      "MyAwesomePlugin",
		PartnerID: "partner_1234",
		URL:       "https://myawesomeplugin.info",
		Version:   "1.2.34",
	}
	SetAppInfo(appInfo)
	defer SetAppInfo(nil)

	c := GetBackend(APIBackend).(*BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	//
	// User-Agent
	//

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9] MyAwesomePlugin/1.2.34 \(https://myawesomeplugin.info\)$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)

	//
	// X-Stripe-Client-User-Agent
	//

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]interface{}
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	application := userAgent["application"].(map[string]interface{})

	assert.Equal(t, "MyAwesomePlugin", application["name"])
	assert.Equal(t, "partner_1234", application["partner_id"])
	assert.Equal(t, "https://myawesomeplugin.info", application["url"])
	assert.Equal(t, "1.2.34", application["version"])
}

func TestStripeClientUserAgent(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]string
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	//
	// Just test a few headers that we know to be stable.
	//

	assert.Empty(t, userAgent["application"])
	assert.Equal(t, "go", userAgent["lang"])
	assert.Equal(t, runtime.Version(), userAgent["lang_version"])

	// Anywhere these tests are running can reasonable be expected to have a
	// `uname` to run, so do this basic check.
	assert.NotEqual(t, UnknownPlatform, userAgent["lang_version"])
}

func TestStripeClientUserAgentWithAppInfo(t *testing.T) {
	appInfo := &AppInfo{
		Name:    "MyAwesomePlugin",
		URL:     "https://myawesomeplugin.info",
		Version: "1.2.34",
	}
	SetAppInfo(appInfo)
	defer SetAppInfo(nil)

	c := GetBackend(APIBackend).(*BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]interface{}
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	decodedAppInfo := userAgent["application"].(map[string]interface{})
	assert.Equal(t, appInfo.Name, decodedAppInfo["name"])
	assert.Equal(t, appInfo.URL, decodedAppInfo["url"])
	assert.Equal(t, appInfo.Version, decodedAppInfo["version"])
}

func TestResponseToError(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)

	// A test response that includes a status code and request ID.
	res := &http.Response{
		Header: http.Header{
			"Request-Id": []string{"request-id"},
		},
		StatusCode: 402,
	}

	// An error that contains expected fields which we're going to serialize to
	// JSON and inject into our conversion function.
	expectedErr := &Error{
		Code:  ErrorCodeMissing,
		Msg:   "That card was declined",
		Param: "expiry_date",
		Type:  ErrorTypeCard,
	}
	bytes, err := json.Marshal(expectedErr)
	assert.NoError(t, err)

	// Unpack the error that we just serialized so that we can inject a
	// type-specific field into it ("decline_code"). This will show up in a
	// field on a special CardError type which is attached to the common
	// Error.
	var raw map[string]string
	err = json.Unmarshal(bytes, &raw)
	assert.NoError(t, err)

	expectedDeclineCode := DeclineCodeInvalidCVC
	raw["decline_code"] = string(expectedDeclineCode)
	bytes, err = json.Marshal(raw)
	assert.NoError(t, err)

	// A generic Golang error.
	err = c.ResponseToError(res, wrapError(bytes))

	// An error containing Stripe-specific fields that we cast back from the
	// generic Golang error.
	stripeErr := err.(*Error)

	assert.Equal(t, expectedErr.Code, stripeErr.Code)
	assert.Equal(t, expectedErr.Msg, stripeErr.Msg)
	assert.Equal(t, expectedErr.Param, stripeErr.Param)
	assert.Equal(t, res.Header.Get("Request-Id"), stripeErr.RequestID)
	assert.Equal(t, res.StatusCode, stripeErr.HTTPStatusCode)
	assert.Equal(t, expectedErr.Type, stripeErr.Type)
	assert.Equal(t, expectedDeclineCode, stripeErr.DeclineCode)

	// Not exhaustive, but verify LastResponse is basically working as
	// expected.
	assert.Equal(t, res.Header.Get("Request-Id"), stripeErr.LastResponse.RequestID)
	assert.Equal(t, res.StatusCode, stripeErr.LastResponse.StatusCode)

	// Just a bogus type coercion to demonstrate how this code might be
	// written. Because we've assigned ErrorTypeCard as the error's type, Err
	// should always come out as a CardError.
	_, ok := stripeErr.Err.(*InvalidRequestError)
	assert.False(t, ok)

	cardErr, ok := stripeErr.Err.(*CardError)
	assert.True(t, ok)

	// For backwards compatibility, `DeclineCode` is also set on the
	// `CardError` structure.
	assert.Equal(t, expectedDeclineCode, cardErr.DeclineCode)
}

func TestStringSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := StringSlice(input)

	assert.Equal(t, "a", *result[0])
	assert.Equal(t, "b", *result[1])
	assert.Equal(t, "c", *result[2])

	assert.Equal(t, 0, len(StringSlice(nil)))
}

func TestInt64Slice(t *testing.T) {
	input := []int64{8, 7, 6}
	result := Int64Slice(input)

	assert.Equal(t, int64(8), *result[0])
	assert.Equal(t, int64(7), *result[1])
	assert.Equal(t, int64(6), *result[2])

	assert.Equal(t, 0, len(Int64Slice(nil)))
}

func TestFloat64Slice(t *testing.T) {
	input := []float64{8, 7, 6}
	result := Float64Slice(input)

	assert.Equal(t, float64(8), *result[0])
	assert.Equal(t, float64(7), *result[1])
	assert.Equal(t, float64(6), *result[2])

	assert.Equal(t, 0, len(Float64Slice(nil)))
}

func TestBoolSlice(t *testing.T) {
	input := []bool{true, false, true, false}
	result := BoolSlice(input)

	assert.Equal(t, true, *result[0])
	assert.Equal(t, false, *result[1])
	assert.Equal(t, true, *result[2])
	assert.Equal(t, false, *result[3])

	assert.Equal(t, 0, len(BoolSlice(nil)))
}

//
// ---
//

// A simple function that allows us to represent an error response from Stripe
// which comes wrapper in a JSON object with a single field of "error".
func wrapError(serialized []byte) []byte {
	return []byte(`{"error":` + string(serialized) + `}`)
}
