package rawrequest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v83"
	_ "github.com/stripe/stripe-go/v83/testing"
)

func createTestClient(testServer *httptest.Server) Client {
	backend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			LeveledLogger: &stripe.LeveledLogger{
				Level: stripe.LevelDebug,
			},
			MaxNetworkRetries: stripe.Int64(0),
			URL:               stripe.String(testServer.URL),
		},
	).(*stripe.BackendImplementation)

	return Client{B: backend, Key: stripe.Key}
	// stripe.SetBackend(stripe.APIBackend, backend)
}

func TestV2PostRequest(t *testing.T) {
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

	client := createTestClient(testServer)

	var params *stripe.RawParams

	_, err := client.RawRequest(http.MethodPost, "/v2/abc", `{"xyz": {"def": "jih"}}`, params)
	assert.NoError(t, err)

	assert.Nil(t, params) // original params should not be modified
	assert.Equal(t, `{"xyz": {"def": "jih"}}`, body)
	assert.Equal(t, `/v2/abc`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/json`, contentType)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestRawV1PostRequest(t *testing.T) {
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

	client := createTestClient(testServer)

	var params *stripe.RawParams

	_, err := client.RawRequest(http.MethodPost, "/v1/abc", `abc=123&a[name]=nested`, params)
	assert.NoError(t, err)

	assert.Nil(t, params) // original params should not be modified
	assert.Equal(t, `abc=123&a[name]=nested`, body)
	assert.Equal(t, `/v1/abc`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/x-www-form-urlencoded`, contentType)
	assert.NoError(t, err)
	defer testServer.Close()
}

func TestV2GetRequestWithAdditionalHeaders(t *testing.T) {
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

	client := createTestClient(testServer)

	headers := http.Header{}
	headers.Set("foo", "bar")
	params := &stripe.RawParams{Params: stripe.Params{Headers: headers}, StripeContext: "acct_123"}

	_, err := client.RawRequest(http.MethodGet, "/v2/abc", "", params)
	assert.NoError(t, err)

	assert.Equal(t, ``, body)
	assert.Equal(t, `/v2/abc`, path)
	assert.Equal(t, `GET`, method)
	assert.Equal(t, `application/json`, contentType)
	assert.Equal(t, `bar`, fooHeader)
	assert.Equal(t, `acct_123`, stripeContext)
	assert.NoError(t, err)
	defer testServer.Close()
}

// TestRawRequestTimeout_NilResponsePanic tests the segmentation fault described in
// https://github.com/stripe/stripe-go/issues/2211
// This test reproduces the actual scenario from the issue where a RawRequest
// times out, causing http.Client.Do to return a nil response with an error,
// which then leads to a nil pointer dereference in handleResponseBufferingErrors.
func TestRawRequestTimeout_NilResponsePanic(t *testing.T) {
	// Create a test server that deliberately times out by never responding
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a slow endpoint that exceeds the timeout
		// Don't write any response, causing the client to timeout
		select {}
	}))
	defer testServer.Close()

	// Create a backend with a very short timeout
	backend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			LeveledLogger: &stripe.LeveledLogger{
				Level: stripe.LevelNull,
			},
			MaxNetworkRetries: stripe.Int64(0),
			URL:               stripe.String(testServer.URL),
			HTTPClient: &http.Client{
				Timeout: 1 * time.Nanosecond,
			},
		},
	).(*stripe.BackendImplementation)
	client := Client{B: backend, Key: "sk_test_123"}

	_, err := client.RawRequest(http.MethodPost, "/v2/billing/meter_event_stream", `{"events":[]}`, nil)
	assert.Error(t, err)
}
