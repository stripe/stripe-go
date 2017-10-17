package stripe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestErrorError(t *testing.T) {
	err := &Error{Type: "foo", Msg: "bar"}
	assert.Equal(t, `{"message":"bar","type":"foo"}`, err.Error())
}

func TestErrorResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Request-Id", "req_123")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, `{"error":{"message":"bar","type":"`+InvalidRequest+`"}}`)
	}))
	defer ts.Close()

	SetBackend("api", &BackendConfiguration{
		APIBackend,
		ts.URL,
		&http.Client{},
	})

	err := GetBackend(APIBackend).Call("GET", "/v1/account", "sk_test_badKey", nil, nil, nil)
	assert.Error(t, err)

	stripeErr := err.(*Error)
	assert.Equal(t, InvalidRequest, stripeErr.Type)
	assert.Equal(t, "req_123", stripeErr.RequestID)
	assert.Equal(t, 401, stripeErr.HTTPStatusCode)
}
