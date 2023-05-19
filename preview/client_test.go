package preview

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v74"
	_ "github.com/stripe/stripe-go/v74/testing"
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

func TestPreviewRequestWithAdditionalHeaders(t *testing.T) {
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

	type myParams struct {
		stripe.Params    `form:"-"`
		stripe.RawParams `form:"-"`
	}

	params := myParams{stripe.Params{Headers: http.Header{"foo": []string{"bar"}}}, stripe.RawParams{StripeContext: "acct_123"}}

	_, err := Get("/v1/abc", &params)
	assert.NoError(t, err)

	assert.Equal(t, "", params.RawParams.APIMode)
	assert.Equal(t, ``, body)
	assert.Equal(t, `/v1/abc`, path)
	assert.Equal(t, `GET`, method)
	assert.Equal(t, ``, contentType)
	assert.Equal(t, `bar`, fooHeader)
	assert.Equal(t, `acct_123`, stripeContext)
	assert.NoError(t, err)
	defer testServer.Close()
}
