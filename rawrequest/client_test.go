package rawrequest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v79"
	_ "github.com/stripe/stripe-go/v79/testing"
)

func stubAPIBackend(testServer *httptest.Server) {
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

	stripe.SetBackend(stripe.APIBackend, backend)
}

func TestGetDefaultRequestOptionsCopiesParams(t *testing.T) {
	headers := http.Header{}
	headers.Set("foo", "bar")

	params := &stripe.RawParams{Params: stripe.Params{Headers: headers}, StripeContext: "acct_123"}

	defaultRequestOptions := getDefaultRequestOptions(params)

	assert.Equal(t, stripe.V2APIMode, defaultRequestOptions.APIMode)
	assert.Equal(t, "bar", defaultRequestOptions.Headers.Get("foo"))
	assert.Equal(t, "acct_123", defaultRequestOptions.StripeContext)
}

func TestGetDefaultRequestOptionsOverridesAPIMode(t *testing.T) {
	params := &stripe.RawParams{APIMode: stripe.V1APIMode}

	defaultRequestOptions := getDefaultRequestOptions(params)

	assert.Equal(t, stripe.V1APIMode, params.APIMode) // original params should not be modified
	assert.Equal(t, stripe.V2APIMode, defaultRequestOptions.APIMode)
}

func TestGetDefaultRequestOptionsWithNilParams(t *testing.T) {
	var params *stripe.RawParams
	defaultRequestOptions := getDefaultRequestOptions(params)

	assert.Nil(t, params) // original params should not be modified
	assert.Equal(t, stripe.V2APIMode, defaultRequestOptions.APIMode)
}

func TestV2PostRequest(t *testing.T) {
	var body string
	var path string
	var method string
	var contentType string
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		body = string(req)
		path = r.URL.RequestURI()
		method = r.Method
		contentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"object": "abc", "xyz": {"def": "jih"}}`))
	}))

	stubAPIBackend(testServer)

	var params *stripe.RawParams

	_, err := Post("/v1/abc", `{"xyz": {"def": "jih"}}`, params)
	assert.NoError(t, err)

	assert.Nil(t, params) // original params should not be modified
	assert.Equal(t, `{"xyz": {"def": "jih"}}`, body)
	assert.Equal(t, `/v1/abc`, path)
	assert.Equal(t, `POST`, method)
	assert.Equal(t, `application/json`, contentType)
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
		req, _ := ioutil.ReadAll(r.Body)
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

	stubAPIBackend(testServer)

	headers := http.Header{}
	headers.Set("foo", "bar")
	params := &stripe.RawParams{Params: stripe.Params{Headers: headers}, StripeContext: "acct_123"}

	_, err := Get("/v1/abc", params)
	assert.NoError(t, err)

	assert.Equal(t, stripe.APIMode(""), params.APIMode) // original params should not be modified
	assert.Equal(t, ``, body)
	assert.Equal(t, `/v1/abc`, path)
	assert.Equal(t, `GET`, method)
	assert.Equal(t, `application/json`, contentType)
	assert.Equal(t, `bar`, fooHeader)
	assert.Equal(t, `acct_123`, stripeContext)
	assert.NoError(t, err)
	defer testServer.Close()
}
