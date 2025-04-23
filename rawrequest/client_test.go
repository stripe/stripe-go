package rawrequest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
		req, _ := ioutil.ReadAll(r.Body)
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
		req, _ := ioutil.ReadAll(r.Body)
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
