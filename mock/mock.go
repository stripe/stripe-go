// Package mock provides a test helper for spinning up a local HTTP server that
// simulates the Stripe API, for use in unit tests that don't require
// stripe-mock.
package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/v86"
	stripetest "github.com/stripe/stripe-go/v86/testing"
)

type Assertion func(*testing.T, *http.Request)

func Server[T any](t *testing.T, method, path string, req T, resp func(T) []byte, asserts ...Assertion) (*httptest.Server, *stripe.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, assert := range asserts {
			assert(t, r)
		}
		assert.Equal(t, r.Method, method)
		assert.Equal(t, r.URL.Path, path)
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+stripetest.TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		if len(body) > 0 {
			assert.NoError(t, json.Unmarshal(body, req))
		}
		data := resp(req)
		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{
			URL: stripe.String(server.URL),
		},
	)
	sc := stripe.NewClient(stripetest.TestAPIKey, stripe.WithBackends(backends))
	return server, sc
}
