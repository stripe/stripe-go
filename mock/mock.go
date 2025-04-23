package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/client"
	. "github.com/max-cape/stripe-go-test/testing"
	"github.com/stretchr/testify/assert"
)

type Assertion func(*testing.T, *http.Request)

func Server[T any](t *testing.T, method, path string, req T, resp func(T) []byte, asserts ...Assertion) (*httptest.Server, *client.API) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, assert := range asserts {
			assert(t, r)
		}
		assert.Equal(t, r.Method, method)
		assert.Equal(t, r.URL.Path, path)
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

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
	sc := client.New(TestAPIKey, backends)
	return server, sc
}
