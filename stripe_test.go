package stripe

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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

	req, err := c.NewRequest("", "/v1/hello", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
}

func TestApiVersion(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, APIVersion, req.Header.Get("Stripe-Version"))
}

func TestCanSetBetaHeaders(t *testing.T) {
	defer cleanupBetaHeaders()
	AddBetaVersion("feature_in_beta", "v3")

	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, APIVersion+"; feature_in_beta=v3", req.Header.Get("Stripe-Version"))
}

func TestSetBetaVersionTwiceAsc(t *testing.T) {
	defer cleanupBetaHeaders()
	err := AddBetaVersion("feature_in_beta", "v3")
	assert.Nil(t, err)
	err = AddBetaVersion("feature_in_beta", "v5")
	assert.Nil(t, err)

	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, APIVersion+"; feature_in_beta=v5", req.Header.Get("Stripe-Version"))

	// clean up
	apiVersionWithBetaHeaders = APIVersion
}

func TestSetBetaVersionTwiceDesc(t *testing.T) {
	defer cleanupBetaHeaders()
	err := AddBetaVersion("feature_in_beta", "v5")
	assert.Nil(t, err)
	err = AddBetaVersion("feature_in_beta", "v3")
	assert.Nil(t, err)

	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, APIVersion+"; feature_in_beta=v5", req.Header.Get("Stripe-Version"))

	// clean up
	apiVersionWithBetaHeaders = APIVersion
}

func TestCannotSetSameBetaHeaderWithInvalidString(t *testing.T) {
	defer cleanupBetaHeaders()
	err := AddBetaVersion("feature_in_beta", "f3")
	assert.Equal(t, err.Error(), "beta version should start with 'v'")

	err = AddBetaVersion("feature_in_beta", "v3a")
	assert.Equal(t, err.Error(), "beta version should start with 'v' followed by a number")

	// clean up
	apiVersionWithBetaHeaders = APIVersion
}

func TestCanSetSecondBetaHeaders(t *testing.T) {
	defer cleanupBetaHeaders()
	AddBetaVersion("feature_in_beta", "v3")
	AddBetaVersion("second_feature_in_beta", "v2")

	c := GetBackend(APIBackend).(*BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, APIVersion+"; feature_in_beta=v3; second_feature_in_beta=v2", req.Header.Get("Stripe-Version"))

	// clean up
	apiVersionWithBetaHeaders = APIVersion
}

func TestContext(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	p := &Params{Context: context.Background()}

	req, err := c.NewRequest("", "/v1/hello", "", "", p)
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
//	go test . -run TestDo_Retry -test.v
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
		"/v1/hello",
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
			&Error{Code: ErrorCodeLockTimeout, HTTPStatusCode: http.StatusTooManyRequests},
			&http.Request{},
			&http.Response{},
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

	// 500 Internal Server Error -- retry
	t.Run("RetryOn500NonPost", func(t *testing.T) {
		shouldRetry, _ := c.shouldRetry(
			nil,
			&http.Request{Method: http.MethodGet},
			&http.Response{StatusCode: http.StatusInternalServerError},
			0,
		)
		assert.True(t, shouldRetry)
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
			HTTPClient:        &http.Client{Timeout: timeout / 2},
		},
	).(*BackendImplementation)

	backend.SetNetworkRetriesSleep(false)

	request, err := backend.NewRequest(
		http.MethodPost,
		"/v1/hello",
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
		"/v1/hello",
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
func TestCall_TelemetryDisabled(t *testing.T) {
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
		var response testServerResponse
		err := backend.Call(http.MethodGet, "/v1/hello", "sk_test_xyz", nil, &response)

		assert.NoError(t, err)
		assert.Equal(t, message, response.Message)
	}

	// We should have seen exactly two requests.
	assert.Equal(t, 2, requestNum)
}

// Test that telemetry metrics are sent on subsequent requests when
// EnableTelemetry = true.
func TestCall_TelemetryEnabled(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
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

			var telemetry requestTelemetry
			// the telemetry should properly unmarshal into RequestTelemetry
			err := json.Unmarshal([]byte(telemetryStr), &telemetry)
			assert.NoError(t, err)

			// the second request should include the metrics for the first request
			assert.Equal(t, telemetry.LastRequestMetrics.RequestID, "req_1")
			assert.True(t, *telemetry.LastRequestMetrics.RequestDurationMS > 20,
				"request_duration_ms should be > 20ms")

			// The telemetry in the second request should contain the
			// expected usage
			assert.ElementsMatch(t, telemetry.LastRequestMetrics.Usage, []string{"llama", "bufo"})
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

	type myCreateParams struct {
		Params `form:"*"`
		Foo    string `form:"foo"`
	}
	params := &myCreateParams{
		Foo: "bar",
	}
	params.InternalSetUsage([]string{"llama", "bufo"})
	for i := 0; i < 2; i++ {
		var response testServerResponse
		err := backend.Call(http.MethodGet, "/v1/hello", "sk_test_xyz", params, &response)

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
func TestCall_TelemetryEnabledNoDataRace(t *testing.T) {
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
			var response testServerResponse
			err := backend.Call(http.MethodGet, "/v1/hello", "sk_test_xyz", nil, &response)

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

// Test that a GET to a v2 endpoint works and is sent with
// the correct query parameters
func TestCall_V2PathGet(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet)
		assert.Equal(t, r.URL.Path, "/v2/hello")
		assert.Equal(t, r.URL.Query().Get("foo"), "bar")

		data, err := json.Marshal(testServerResponse{Message: message})
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
	)

	type myCreateParams struct {
		Params `form:"*"`
		Foo    string `form:"foo" json:"foo"`
	}
	params := &myCreateParams{
		Foo: "bar",
	}

	var response testServerResponse
	err := backend.Call(http.MethodGet, "/v2/hello", "sk_test_xyz", params, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)
}

// Test that a POST to a v2 endpoint works and is sent with
// the correct JSON contentType
func TestCall_V2PathPost(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/hello")
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")

		var p map[string]string
		err := json.NewDecoder(r.Body).Decode(&p)
		assert.NoError(t, err)
		assert.Equal(t, p["foo"], "bar")

		data, err := json.Marshal(testServerResponse{Message: message})
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
	)

	type myCreateParams struct {
		Params `form:"*"`
		Foo    string `form:"foo" json:"foo"`
	}
	params := &myCreateParams{
		Foo: "bar",
	}

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v2/hello", "sk_test_xyz", params, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)
}

// Test that a POST to a v2 endpoint works and is sent with an empty body
// when params are nil
func TestCall_V2PathPostNilParams(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/hello")
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, body, []byte{})

		data, err := json.Marshal(testServerResponse{Message: message})
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
	)

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v2/hello", "sk_test_xyz", nil, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)
}

// Test that UnsetFields causes explicit null values in v2 JSON POST body
func TestCall_V2PathPostUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	message := "Hello, client."

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v2/hello", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(body, &raw)
		assert.NoError(t, err)

		// "foo" should be present with value "bar"
		assert.Equal(t, `"bar"`, string(raw["foo"]))
		// "description" should be explicitly null
		assert.Equal(t, "null", string(raw["description"]))
		// "baz" (nil pointer, not in UnsetFields) should NOT be present
		_, bazPresent := raw["baz"]
		assert.False(t, bazPresent, "baz should not be present in JSON")

		data, err := json.Marshal(testServerResponse{Message: message})
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
	)

	type myUpdateParams struct {
		Params      `form:"*"`
		Foo         string   `form:"foo" json:"foo"`
		Description *string  `form:"description" json:"description,omitempty"`
		Baz         *string  `form:"baz" json:"baz,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	params := &myUpdateParams{
		Foo: "bar",
		// Description and Baz are both nil
	}
	params.UnsetFields = append(params.UnsetFields, "description") // Explicitly send null for description

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v2/hello", "sk_test_xyz", params, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)
}

// Test that UnsetFields sends field= (empty string) in v1 form-encoded requests
func TestCall_V1PathPostUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/hello", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		bodyStr := string(body)

		// "foo" should be present with value "bar"
		assert.Contains(t, bodyStr, "foo=bar")
		// "description" should be present as empty string (description=)
		assert.Contains(t, bodyStr, "description=")
		// but "description" should NOT have a value after the =
		// Parse to verify the value is empty
		vals, err := url.ParseQuery(bodyStr)
		assert.NoError(t, err)
		assert.Equal(t, "", vals.Get("description"))

		data, err := json.Marshal(testServerResponse{Message: "ok"})
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
	)

	type myParams struct {
		Params      `form:"*"`
		Foo         string   `form:"foo" json:"foo"`
		Description *string  `form:"description" json:"description,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	params := &myParams{
		Foo: "bar",
	}
	params.UnsetFields = append(params.UnsetFields, "description") // Should send description= in v1

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v1/hello", "sk_test_xyz", params, &response)

	assert.NoError(t, err)
}

// Test marshalV2JSON directly
func TestMarshalV2JSON(t *testing.T) {
	type testParams struct {
		Params      `form:"*"`
		Name        *string  `form:"name" json:"name,omitempty"`
		Description *string  `form:"description" json:"description,omitempty"`
		Amount      *int64   `form:"amount" json:"amount,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}

	t.Run("without UnsetFields", func(t *testing.T) {
		params := &testParams{
			Name: String("test"),
		}
		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		assert.Equal(t, `"test"`, string(raw["name"]))
		_, hasDesc := raw["description"]
		assert.False(t, hasDesc, "description should not be present")
	})

	t.Run("with UnsetFields", func(t *testing.T) {
		params := &testParams{
			Name: String("test"),
		}
		params.UnsetFields = append(params.UnsetFields, "description")
		params.UnsetFields = append(params.UnsetFields, "amount")

		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		assert.Equal(t, `"test"`, string(raw["name"]))
		assert.Equal(t, "null", string(raw["description"]))
		assert.Equal(t, "null", string(raw["amount"]))
	})

	t.Run("UnsetFields overrides set values", func(t *testing.T) {
		params := &testParams{
			Name:        String("test"),
			Description: String("a description"),
		}
		params.UnsetFields = append(params.UnsetFields, "description") // UnsetFields overrides even if set

		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		// UnsetFields takes precedence — user explicitly asked for null
		assert.Equal(t, "null", string(raw["description"]))
	})

	t.Run("nested struct UnsetFields", func(t *testing.T) {
		type nestedDetails struct {
			Comment     *string  `json:"comment,omitempty"`
			Feedback    *string  `json:"feedback,omitempty"`
			UnsetFields []string `form:"-" json:"-"`
		}
		type outerParams struct {
			Params      `form:"*"`
			Name        *string        `json:"name,omitempty"`
			Details     *nestedDetails `json:"details,omitempty"`
			UnsetFields []string       `form:"-" json:"-"`
		}

		params := &outerParams{
			Name: String("test"),
			Details: &nestedDetails{
				Feedback: String("too_expensive"),
			},
		}
		params.UnsetFields = append(params.UnsetFields, "name") // root-level empty
		params.Details.UnsetFields = append(params.Details.UnsetFields, "comment")

		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		// Root-level: "name" should be null
		assert.Equal(t, "null", string(raw["name"]))

		// Nested: "details.comment" should be null, "details.feedback" should be present
		var details map[string]json.RawMessage
		err = json.Unmarshal(raw["details"], &details)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(details["comment"]))
		assert.Equal(t, `"too_expensive"`, string(details["feedback"]))
	})

	t.Run("deeply nested UnsetFields", func(t *testing.T) {
		type innerSettings struct {
			Threshold   *int64   `json:"threshold,omitempty"`
			UnsetFields []string `form:"-" json:"-"`
		}
		type midLevel struct {
			Settings    *innerSettings `json:"settings,omitempty"`
			Mode        *string        `json:"mode,omitempty"`
			UnsetFields []string       `form:"-" json:"-"`
		}
		type rootParams struct {
			Params  `form:"*"`
			Details *midLevel `json:"details,omitempty"`
		}

		params := &rootParams{
			Details: &midLevel{
				Mode:     String("auto"),
				Settings: &innerSettings{},
			},
		}
		params.Details.UnsetFields = append(params.Details.UnsetFields, "mode")
		params.Details.Settings.UnsetFields = append(params.Details.Settings.UnsetFields, "threshold")

		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		var details map[string]json.RawMessage
		err = json.Unmarshal(raw["details"], &details)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(details["mode"]))

		var settings map[string]json.RawMessage
		err = json.Unmarshal(details["settings"], &settings)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(settings["threshold"]))
	})

	t.Run("slice element UnsetFields", func(t *testing.T) {
		type lineItem struct {
			Price       *string  `json:"price,omitempty"`
			Quantity    *int64   `json:"quantity,omitempty"`
			UnsetFields []string `form:"-" json:"-"`
		}
		type rootParams struct {
			Params    `form:"*"`
			LineItems []*lineItem `json:"line_items,omitempty"`
		}

		params := &rootParams{
			LineItems: []*lineItem{
				{Price: String("price_123"), Quantity: Int64(2)},
				{Price: String("price_456")},
			},
		}
		// Unset quantity on the second element
		params.LineItems[1].UnsetFields = append(params.LineItems[1].UnsetFields, "quantity")

		data, err := marshalV2JSON(params)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(data, &raw)
		assert.NoError(t, err)

		var items []json.RawMessage
		err = json.Unmarshal(raw["line_items"], &items)
		assert.NoError(t, err)
		assert.Len(t, items, 2)

		// First element: quantity is set, should not be null
		var item0 map[string]json.RawMessage
		err = json.Unmarshal(items[0], &item0)
		assert.NoError(t, err)
		assert.Equal(t, `"price_123"`, string(item0["price"]))
		assert.Equal(t, `2`, string(item0["quantity"]))

		// Second element: quantity should be null via UnsetFields
		var item1 map[string]json.RawMessage
		err = json.Unmarshal(items[1], &item1)
		assert.NoError(t, err)
		assert.Equal(t, `"price_456"`, string(item1["price"]))
		assert.Equal(t, "null", string(item1["quantity"]))
	})
}

// Test nested UnsetFields with v1 form encoding (bracket notation)
func TestCall_V1PathPostNestedUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		bodyStr := string(body)

		// Root-level: "description=" should be present
		vals, err := url.ParseQuery(bodyStr)
		assert.NoError(t, err)
		assert.Equal(t, "", vals.Get("description"))
		assert.Contains(t, bodyStr, "description=")

		// Nested: "details[comment]=" should be present (stripe form encoder
		// uses literal brackets, not percent-encoded)
		assert.Equal(t, "", vals.Get("details[comment]"))
		assert.Contains(t, bodyStr, "details[comment]=")

		// "details[feedback]" should have a value
		assert.Equal(t, "too_slow", vals.Get("details[feedback]"))

		data, _ := json.Marshal(testServerResponse{Message: "ok"})
		w.Write(data)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	)

	type nestedDetails struct {
		Comment     *string  `form:"comment" json:"comment,omitempty"`
		Feedback    *string  `form:"feedback" json:"feedback,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	type myParams struct {
		Params      `form:"*"`
		Description *string        `form:"description" json:"description,omitempty"`
		Details     *nestedDetails `form:"details" json:"details,omitempty"`
		UnsetFields []string       `form:"-" json:"-"`
	}
	params := &myParams{
		Details: &nestedDetails{
			Feedback: String("too_slow"),
		},
	}
	params.UnsetFields = append(params.UnsetFields, "description")
	params.Details.UnsetFields = append(params.Details.UnsetFields, "comment")

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v1/hello", "sk_test_xyz", params, &response)
	assert.NoError(t, err)
}

// Test nested UnsetFields with v2 JSON encoding (end-to-end via Call)
func TestCall_V2PathPostNestedUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(body, &raw)
		assert.NoError(t, err)

		// Root-level: "description" should be null
		assert.Equal(t, "null", string(raw["description"]))

		// Nested: "details.comment" should be null
		var details map[string]json.RawMessage
		err = json.Unmarshal(raw["details"], &details)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(details["comment"]))
		assert.Equal(t, `"too_slow"`, string(details["feedback"]))

		data, _ := json.Marshal(testServerResponse{Message: "ok"})
		w.Write(data)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	)

	type nestedDetails struct {
		Comment     *string  `form:"comment" json:"comment,omitempty"`
		Feedback    *string  `form:"feedback" json:"feedback,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	type myParams struct {
		Params      `form:"*"`
		Description *string        `form:"description" json:"description,omitempty"`
		Details     *nestedDetails `form:"details" json:"details,omitempty"`
		UnsetFields []string       `form:"-" json:"-"`
	}
	params := &myParams{
		Details: &nestedDetails{
			Feedback: String("too_slow"),
		},
	}
	params.UnsetFields = append(params.UnsetFields, "description")
	params.Details.UnsetFields = append(params.Details.UnsetFields, "comment")

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v2/hello", "sk_test_xyz", params, &response)
	assert.NoError(t, err)
}

// Test UnsetFields on slice element params with v1 form encoding
func TestCall_V1PathPostSliceElementUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		bodyStr := string(body)

		vals, err := url.ParseQuery(bodyStr)
		assert.NoError(t, err)

		// First item: price set, no unset
		assert.Equal(t, "price_123", vals.Get("line_items[0][price]"))

		// Second item: price set, tax_rates cleared via UnsetFields
		assert.Equal(t, "price_456", vals.Get("line_items[1][price]"))
		assert.Equal(t, "", vals.Get("line_items[1][tax_rates]"))
		assert.Contains(t, bodyStr, "line_items[1][tax_rates]=")

		data, _ := json.Marshal(testServerResponse{Message: "ok"})
		w.Write(data)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	)

	type lineItem struct {
		Price       *string  `form:"price" json:"price,omitempty"`
		TaxRates    *string  `form:"tax_rates" json:"tax_rates,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	type myParams struct {
		Params    `form:"*"`
		LineItems []*lineItem `form:"line_items" json:"line_items,omitempty"`
	}
	params := &myParams{
		LineItems: []*lineItem{
			{Price: String("price_123")},
			{Price: String("price_456")},
		},
	}
	params.LineItems[1].UnsetFields = append(params.LineItems[1].UnsetFields, "tax_rates")

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v1/hello", "sk_test_xyz", params, &response)
	assert.NoError(t, err)
}

// Test UnsetFields on slice element params with v2 JSON encoding
func TestCall_V2PathPostSliceElementUnsetFields(t *testing.T) {
	type testServerResponse struct {
		APIResource
		Message string `json:"message"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(body, &raw)
		assert.NoError(t, err)

		var items []json.RawMessage
		err = json.Unmarshal(raw["line_items"], &items)
		assert.NoError(t, err)
		assert.Len(t, items, 2)

		// First item: normal, no nulls
		var item0 map[string]json.RawMessage
		err = json.Unmarshal(items[0], &item0)
		assert.NoError(t, err)
		assert.Equal(t, `"price_123"`, string(item0["price"]))
		_, hasTax0 := item0["tax_rates"]
		assert.False(t, hasTax0)

		// Second item: tax_rates should be null
		var item1 map[string]json.RawMessage
		err = json.Unmarshal(items[1], &item1)
		assert.NoError(t, err)
		assert.Equal(t, `"price_456"`, string(item1["price"]))
		assert.Equal(t, "null", string(item1["tax_rates"]))

		data, _ := json.Marshal(testServerResponse{Message: "ok"})
		w.Write(data)
	}))
	defer testServer.Close()

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	)

	type lineItem struct {
		Price       *string  `form:"price" json:"price,omitempty"`
		TaxRates    *string  `form:"tax_rates" json:"tax_rates,omitempty"`
		UnsetFields []string `form:"-" json:"-"`
	}
	type myParams struct {
		Params    `form:"*"`
		LineItems []*lineItem `form:"line_items" json:"line_items,omitempty"`
	}
	params := &myParams{
		LineItems: []*lineItem{
			{Price: String("price_123")},
			{Price: String("price_456")},
		},
	}
	params.LineItems[1].UnsetFields = append(params.LineItems[1].UnsetFields, "tax_rates")

	var response testServerResponse
	err := backend.Call(http.MethodPost, "/v2/hello", "sk_test_xyz", params, &response)
	assert.NoError(t, err)
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
		"/v1/hello",
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
		"/v1/pdf",
		"sk_test_123",
		nil,
		&response,
	)
	assert.NoError(t, err)
	result, err := io.ReadAll(response.LastResponse.Body)
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
		"/v1/pdf",
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
		"/v1/pdf",
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

func TestGetBackendWithConfig(t *testing.T) {
	// No config overrides: https://github.com/stripe/stripe-go/issues/1894
	{
		config := &BackendConfig{}
		_ = GetBackendWithConfig(
			APIBackend,
			config,
		).(*BackendImplementation)

		// Config properties left unchanged
		assert.Nil(t, config.EnableTelemetry)
		assert.Nil(t, config.HTTPClient)
		assert.Nil(t, config.LeveledLogger)
		assert.Nil(t, config.MaxNetworkRetries)
		assert.Nil(t, config.URL)
	}
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

			req, err := c.NewRequest("", "/v1/hello", key, "", nil)
			assert.NoError(t, err)

			assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
		}()
	}
	wg.Wait()
}

func TestIdempotencyKey(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	p := &Params{IdempotencyKey: String("idempotency-key")}

	req, err := c.NewRequest("", "/v1/hello", "", "", p)
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

	req, err := c.NewRequest("", "/v1/hello", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, "acct_123", req.Header.Get("Stripe-Account"))
}

func TestErrorOnDuplicateMetadata(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	type myParams struct {
		Params   `form:"*"`
		Metadata map[string]string `form:"metadata"`
	}

	metadata := map[string]string{"foo": "bar"}
	resource := APIResource{}
	err := c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{Metadata: metadata}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{Params: Params{Metadata: metadata}}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{Metadata: metadata, Params: Params{Metadata: metadata}}, &resource)
	assert.Errorf(t, err, "You cannot specify both the (deprecated) .Params.Metadata and .Metadata in myParams")
}

func TestErrorOnDuplicateExpand(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	type myParams struct {
		Params `form:"*"`
		Expand []*string `form:"expand"`
	}

	expand := []*string{String("foo"), String("bar")}
	resource := APIResource{}
	err := c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{Expand: expand}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{
			Params: Params{Expand: expand},
		}, &resource)
	assert.NoError(t, err)

	err =
		c.Call("POST", "/v1/customers", "sk_test_xyz", &myParams{
			Expand: expand, Params: Params{Expand: expand}}, &resource)

	assert.Errorf(t, err, "You cannot specify both the (deprecated) .Params.Expand and .Expand in myParams")
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
		err := backend.unmarshalJSONVerbose(context.Background(), 200, []byte(`{"message":"hello"}`), &sample)
		assert.NoError(t, err)
		assert.Equal(t, "hello", sample.Message)
	}

	// Invalid JSON (short)
	{
		body := `server error`

		var sample testServerResponse
		err := backend.unmarshalJSONVerbose(context.Background(), 200, []byte(body), &sample)
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
		err := backend.unmarshalJSONVerbose(context.Background(), 200, []byte(body), &sample)
		assert.Regexp(t,
			fmt.Sprintf(`^Couldn't deserialize JSON \(response status: 200, body sample: '%s ...'\): invalid character`, body[0:500]),
			err)
	}
}

func TestUserAgent(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)

	req, err := c.NewRequest("", "/v1/hello", "", "", nil)
	assert.NoError(t, err)

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[.\-\w\d]+( AIAgent/\w+)?$`)

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

	req, err := c.NewRequest("", "/v1/hello", "", "", nil)
	assert.NoError(t, err)

	//
	// User-Agent
	//

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[.\-\w\d]+ MyAwesomePlugin/1.2.34 \(https://myawesomeplugin.info\)( AIAgent/\w+)?$`)

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

	req, err := c.NewRequest("", "/v1/hello", "", "", nil)
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

	// platform should be present when telemetry is enabled
	expectedPlatform := runtime.GOOS + " " + runtime.GOARCH
	assert.Equal(t, expectedPlatform, userAgent["platform"])
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

	req, err := c.NewRequest("", "/v1/hello", "", "", nil)
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

func TestDetectAIAgent(t *testing.T) {
	// Test detection of a specific agent
	t.Run("DetectsAgent", func(t *testing.T) {
		getEnv := func(key string) (string, bool) {
			if key == "CLAUDECODE" {
				return "1", true
			}
			return "", false
		}
		agent, ok := detectAIAgent(getEnv)
		assert.True(t, ok)
		assert.Equal(t, "claude_code", agent)
	})

	// Test no agent detected
	t.Run("NoAgent", func(t *testing.T) {
		getEnv := func(key string) (string, bool) { return "", false }
		agent, ok := detectAIAgent(getEnv)
		assert.False(t, ok)
		assert.Equal(t, "", agent)
	})

	// Test each agent individually
	t.Run("AllAgents", func(t *testing.T) {
		for k, expected := range aiAgents {
			getEnv := func(key string) (string, bool) {
				if key == k {
					return "1", true
				}
				return "", false
			}
			agent, ok := detectAIAgent(getEnv)
			assert.True(t, ok)
			assert.Equal(t, expected, agent, "Expected %s for env var %s", expected, k)
		}
	})
}

func TestUserAgentWithAIAgent(t *testing.T) {
	// Save and restore the original encodedUserAgent
	originalEncodedUserAgent := encodedUserAgent
	defer func() {
		encodedUserAgent = originalEncodedUserAgent
	}()

	// Simulate AI agent detection by manually setting the user agent
	encodedUserAgent = "Stripe/v1 GoBindings/" + clientversion + " AIAgent/claude_code"

	c := GetBackend(APIBackend).(*BackendImplementation)
	req, err := c.NewRequest("", "/v1/hello", "", "", nil)
	assert.NoError(t, err)

	assert.Contains(t, req.Header.Get("User-Agent"), "AIAgent/claude_code")
}

func TestStripeClientUserAgentWithAIAgent(t *testing.T) {
	// Save and restore the original state
	originalEncoded := encodedStripeUserAgent
	originalReady := encodedStripeUserAgentReady
	defer func() {
		encodedStripeUserAgent = originalEncoded
		encodedStripeUserAgentReady = originalReady
	}()

	// Reset so getEncodedStripeUserAgent re-computes
	encodedStripeUserAgentReady = &sync.Once{}

	encoded := getEncodedStripeUserAgent(true)
	var userAgent map[string]interface{}
	err := json.Unmarshal([]byte(encoded), &userAgent)
	assert.NoError(t, err)

	// The ai_agent field may or may not be present depending on the test
	// environment. If no AI agent env var is set, the key should be absent
	// (omitempty). We just verify the JSON is valid and the field type is
	// correct if present.
	if val, ok := userAgent["ai_agent"]; ok {
		_, isString := val.(string)
		assert.True(t, isString)
	}
}

func TestStripeClientUserAgentOmitsPlatformWithoutTelemetry(t *testing.T) {
	originalEncoded := encodedStripeUserAgent
	originalReady := encodedStripeUserAgentReady
	defer func() {
		encodedStripeUserAgent = originalEncoded
		encodedStripeUserAgentReady = originalReady
	}()

	encodedStripeUserAgentReady = &sync.Once{}

	encoded := getEncodedStripeUserAgent(false)
	var userAgent map[string]string
	err := json.Unmarshal([]byte(encoded), &userAgent)
	assert.NoError(t, err)

	assert.Equal(t, "go", userAgent["lang"])
	assert.Empty(t, userAgent["platform"])
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

func TestTimeValue_NilPointer(t *testing.T) {
	var nilTime *time.Time
	result := TimeValue(nilTime)
	assert.Equal(t, time.Time{}, result)
}

func TestTime(t *testing.T) {
	now := time.Now()
	stripeNow := Time(now)
	assert.Equal(t, &now, stripeNow)
	result := TimeValue(stripeNow)
	assert.Equal(t, now, result)
}

func TestString(t *testing.T) {
	input := "foo"
	result := String(input)
	assert.Equal(t, "foo", *result)
}

func TestString_Enum(t *testing.T) {
	type MyEnum string
	const (
		MyEnumFoo MyEnum = "foo"
		MyEnumBar MyEnum = "bar"
	)

	assert.Equal(t, "foo", *String(MyEnumFoo))
	assert.Equal(t, "foo", *String(string(MyEnumFoo)))
}

func TestRawRequestPreviewPost(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := io.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	type MyXYZ struct {
		DEF string `json:"def"`
	}
	type MyABC struct {
		Object string `json:"object"`
		XYZ    MyXYZ  `json:"xyz"`
	}
	params := &RawParams{Params: Params{}}
	response, err := backend.RawRequest(http.MethodPost, "/v2/abcs", "sk_test_xyz", `{"foo":"myFoo","bar":{"baz":false}}`, params)
	assert.NoError(t, err)
	myABC := &MyABC{}
	assert.Nil(t, params.Headers)
	assert.Equal(t, `{"foo":"myFoo","bar":{"baz":false}}`, body)
	assert.Equal(t, `/v2/abcs`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/json`, contentType)
	// assert.Equal(t, previewVersion, stripeVersion)
	err = json.Unmarshal(response.RawJSON, myABC)
	assert.NoError(t, err)
	assert.Equal(t, "jih", myABC.XYZ.DEF)
	assert.Equal(t, "abc", myABC.Object)
	defer testServer.Close()
}

func TestRawRequestStandardGet(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	var stripeVersion string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := io.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		stripeVersion = r.Header.Get("Stripe-Version")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	_, err := backend.RawRequest(http.MethodGet, "/v1/abc?foo=myFoo", "sk_test_xyz", ``, nil)
	assert.NoError(t, err)
	assert.Equal(t, ``, body)
	assert.Equal(t, `/v1/abc?foo=myFoo`, path)
	assert.Equal(t, `GET`, method)
	assert.Equal(t, `application/x-www-form-urlencoded`, contentType)
	assert.Equal(t, APIVersion, stripeVersion)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestRawRequestStandardPost(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	var stripeVersion string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := io.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		stripeVersion = r.Header.Get("Stripe-Version")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	_, err := backend.RawRequest(http.MethodPost, "/v1/abc", "sk_test_xyz", `foo=myFoo`, nil)
	assert.NoError(t, err)
	assert.Equal(t, `foo=myFoo`, body)
	assert.Equal(t, `/v1/abc`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/x-www-form-urlencoded`, contentType)
	assert.Equal(t, APIVersion, stripeVersion)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestRawRequestPreviewGet(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := io.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	params := &RawParams{Params: Params{}}
	_, err := backend.RawRequest(http.MethodGet, "/v2/abc?foo=myFoo", "sk_test_xyz", ``, params)
	assert.NoError(t, err)
	assert.Equal(t, ``, body)
	assert.Equal(t, `/v2/abc?foo=myFoo`, path)
	assert.Equal(t, `GET`, method)
	assert.Equal(t, `application/json`, contentType)
	// assert.Equal(t, previewVersion, stripeVersion)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestRawRequestWithAdditionalHeaders(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	var fooHeader string
	var stripeContext string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := io.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		fooHeader = r.Header.Get("foo")
		stripeContext = r.Header.Get("Stripe-Context")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
		},
	).(*BackendImplementation)

	headers := http.Header{}
	headers.Set("foo", "bar")
	params := &RawParams{Params: Params{Headers: headers}, StripeContext: "acct_123"}

	_, err := backend.RawRequest(http.MethodPost, "/v2/abc", "sk_test_xyz", `{"foo":"myFoo"}`, params)
	assert.NoError(t, err)
	assert.Equal(t, `{"foo":"myFoo"}`, body)
	assert.Equal(t, `/v2/abc`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/json`, contentType)
	assert.Equal(t, `bar`, fooHeader)
	assert.Equal(t, `acct_123`, stripeContext)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestRawRequestTelemetry(t *testing.T) {
	var telemetry []byte
	i := 0
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body.Close()
		telemetry = []byte(r.Header.Get("X-Stripe-Client-Telemetry"))
		i += 1
		w.Header().Add("Request-Id", fmt.Sprintf("req_%d", i))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     debugLeveledLogger,
			MaxNetworkRetries: Int64(0),
			URL:               String(testServer.URL),
			EnableTelemetry:   Bool(true),
		},
	).(*BackendImplementation)

	params := &RawParams{Params: Params{}}
	_, err := backend.RawRequest(http.MethodPost, "/v2/abcs", "sk_test_xyz", `{}`, params)
	assert.Empty(t, telemetry)
	assert.NoError(t, err)
	// Again, for the telemetry.
	_, err = backend.RawRequest(http.MethodPost, "/v2/abcs", "sk_test_xyz", `{}`, params)
	assert.NoError(t, err)
	metrics := struct {
		LastRequestMetrics requestMetrics `json:"last_request_metrics"`
	}{}
	json.Unmarshal(telemetry, &metrics)
	assert.Equal(t, []string{"raw_request"}, metrics.LastRequestMetrics.Usage)
	defer testServer.Close()
}

func TestAddBetaVersion(t *testing.T) {
	defer cleanupBetaHeaders()
	AddBetaVersion("feature_beta", "v3")
	expectedAPIVersion := APIVersion + "; feature_beta=v3"
	assert.Equal(t, expectedAPIVersion, apiVersionWithBetaHeaders)
	err := AddBetaVersion("feature_beta", "v3")
	assert.Nil(t, err)
}

//
// ---
//

// A simple function that allows us to represent an error response from Stripe
// which comes wrapper in a JSON object with a single field of "error".
func wrapError(serialized []byte) []byte {
	return []byte(`{"error":` + string(serialized) + `}`)
}

func cleanupBetaHeaders() {
	apiVersionWithBetaHeaders = APIVersion
}

func TestStripeContextWhenUnset(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	req, err := c.NewRequest("", "/v2/foo", "", "", nil)
	assert.NoError(t, err)
	assert.Empty(t, req.Header.Get("Stripe-Context"))
}

func TestStripeContextWhenSetWithV1(t *testing.T) {
	c := GetBackend(APIBackend).(*BackendImplementation)
	req, err := c.NewRequest("", "/v1/foo", "", "", nil)
	assert.NoError(t, err)
	assert.Empty(t, req.Header.Get("Stripe-Context"))
}

func TestStripeContextWhenSet(t *testing.T) {
	c := GetBackendWithConfig(APIBackend, &BackendConfig{StripeContext: ParseStripeContext("ctx").StringPtr()}).(*BackendImplementation)
	req, err := c.NewRequest("", "/v2/foo", "", "", nil)
	assert.NoError(t, err)
	assert.Equal(t, "ctx", req.Header.Get("Stripe-Context"))
}

func TestStripeContextWhenSetInParams(t *testing.T) {
	c := GetBackendWithConfig(APIBackend, &BackendConfig{StripeContext: String("ctx")}).(*BackendImplementation)
	req, err := c.NewRequest("", "/v2/foo", "", "", &Params{StripeContext: String("requestCtx")})
	assert.NoError(t, err)
	assert.Equal(t, "requestCtx", req.Header.Get("Stripe-Context"))
}

func TestStripeContextPrecedence(t *testing.T) {
	// Test precedence: BackendConfig.StripeContext vs Params.StripeContext

	// Case 1: Only BackendConfig.StripeContext is set
	{
		c := GetBackendWithConfig(APIBackend, &BackendConfig{StripeContext: String("ctx")}).(*BackendImplementation)
		req, err := c.NewRequest("", "/v2/foo", "", "", nil)
		assert.NoError(t, err)
		assert.Equal(t, "ctx", req.Header.Get("Stripe-Context"))
	}

	// Case 2: Only Params.StripeContext is set
	{
		c := GetBackendWithConfig(APIBackend, &BackendConfig{}).(*BackendImplementation)
		req, err := c.NewRequest("", "/v2/foo", "", "", &Params{StripeContext: String("requestCtx")})
		assert.NoError(t, err)
		assert.Equal(t, "requestCtx", req.Header.Get("Stripe-Context"))
	}

	// Case 3: Neither is set
	{
		c := GetBackendWithConfig(APIBackend, &BackendConfig{}).(*BackendImplementation)
		req, err := c.NewRequest("", "/v2/foo", "", "", nil)
		assert.NoError(t, err)
		assert.Empty(t, req.Header.Get("Stripe-Context"))
	}

	// Case 4: Both are set, Params.StripeContext should take precedence
	{
		c := GetBackendWithConfig(APIBackend, &BackendConfig{StripeContext: String("ctx")}).(*BackendImplementation)
		req, err := c.NewRequest("", "/v2/foo", "", "", &Params{StripeContext: String("requestCtx")})
		assert.NoError(t, err)
		assert.Equal(t, "requestCtx", req.Header.Get("Stripe-Context"))
	}

	// Case 5: Setting an empty StripeContext on the request should overwrite and clear the header
	{
		c := GetBackendWithConfig(APIBackend, &BackendConfig{StripeContext: String("ctx")}).(*BackendImplementation)
		ctx := NewStripeContext(nil)
		req, err := c.NewRequest("", "/v2/foo", "", "", &Params{StripeContext: ctx.StringPtr()})
		assert.NoError(t, err)
		assert.Empty(t, req.Header.Get("Stripe-Context"))
	}
}

func TestHandleV2ErrorWhenKnownError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, err := w.Write([]byte(`{"error":{"type":"temporary_session_expired","message":"Temporary session expired"}}`))
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
	)
	err := backend.Call(http.MethodGet, "/v2/hello", "sk_test_xyz", &Params{}, &APIResource{})
	assert.Error(t, err)
	stripeErr, ok := err.(*TemporarySessionExpiredError)
	assert.True(t, ok)
	assert.Equal(t, "Temporary session expired", stripeErr.Message)
}

func TestHandleV2ErrorWhenUnknownError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Request-Id", "req_123")
		w.WriteHeader(400)
		_, err := w.Write([]byte(`{"error":{"type":"unknown_type","message":"Some message"}}`))
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
	)
	err := backend.Call(http.MethodGet, "/v2/hello", "sk_test_xyz", &Params{}, &APIResource{})
	assert.Error(t, err)
	stripeErr, ok := err.(*V2RawError)
	assert.True(t, ok)
	assert.Equal(t, "Some message", stripeErr.Message)
	assert.Equal(t, "req_123", stripeErr.RequestID)
	assert.Equal(t, 400, stripeErr.HTTPStatusCode)
}

func TestHandleV2ErrorWhenUnknownErrorWithoutType(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Request-Id", "req_123")
		w.WriteHeader(400)
		_, err := w.Write([]byte(`{"error":{"message":"Some message"}}`))
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
	)
	err := backend.Call(http.MethodGet, "/v2/hello", "sk_test_xyz", &Params{}, &APIResource{})
	assert.Error(t, err)
	stripeErr, ok := err.(*V2RawError)
	assert.True(t, ok)
	assert.Equal(t, "Some message", stripeErr.Message)
	assert.Equal(t, "req_123", stripeErr.RequestID)
	assert.Equal(t, 400, stripeErr.HTTPStatusCode)
}

// TestHandleResponseBufferingErrors_NilResponse tests the segmentation fault
// described in https://github.com/stripe/stripe-go/issues/2211
// When http.Client.Do returns an error, the *http.Response may be nil.
// The handleResponseBufferingErrors method should handle this case without panicking.
func TestHandleResponseBufferingErrors_NilResponse(t *testing.T) {
	backend := GetBackendWithConfig(
		APIBackend,
		&BackendConfig{
			LeveledLogger:     nullLeveledLogger,
			MaxNetworkRetries: Int64(0),
		},
	).(*BackendImplementation)

	// Simulate a timeout error where http.Client.Do returns nil response
	var resp *http.Response = nil
	timeoutErr := fmt.Errorf("timeout error")

	_, err := backend.handleResponseBufferingErrors(resp, timeoutErr)
	assert.Error(t, err)
	assert.Equal(t, timeoutErr, err)
}

func BenchmarkCollectAllUnsetFields(b *testing.B) {
	// Minimal params — just the root struct with one unset field.
	b.Run("minimal", func(b *testing.B) {
		params := &PaymentIntentCreateParams{
			Amount:   Int64(1000),
			Currency: String("usd"),
		}
		params.AddUnsetField(PaymentIntentCreateParamsUnsetFieldMandateData)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			collectAllUnsetFields(reflect.ValueOf(params))
		}
	})

	// Deeply nested params — many populated sub-structs to stress the
	// recursive reflection walk.
	b.Run("deeply_nested", func(b *testing.B) {
		params := &PaymentIntentCreateParams{
			Amount:   Int64(1000),
			Currency: String("usd"),
			PaymentMethodOptions: &PaymentIntentCreatePaymentMethodOptionsParams{
				Card: &PaymentIntentCreatePaymentMethodOptionsCardParams{
					Installments: &PaymentIntentCreatePaymentMethodOptionsCardInstallmentsParams{
						Enabled: Bool(true),
					},
				},
				USBankAccount: &PaymentIntentCreatePaymentMethodOptionsUSBankAccountParams{
					FinancialConnections: &PaymentIntentCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams{
						Permissions: []*string{String("payment_method")},
					},
				},
			},
			PaymentMethodData: &PaymentIntentCreatePaymentMethodDataParams{
				BillingDetails: &PaymentIntentCreatePaymentMethodDataBillingDetailsParams{
					Address: &AddressParams{
						City:    String("SF"),
						Country: String("US"),
					},
				},
			},
			Shipping: &ShippingDetailsParams{
				Address: &AddressParams{
					City:    String("SF"),
					Country: String("US"),
				},
				Name: String("Test"),
			},
		}
		params.AddUnsetField(PaymentIntentCreateParamsUnsetFieldMandateData)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			collectAllUnsetFields(reflect.ValueOf(params))
		}
	})
}
