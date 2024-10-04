// Package rawrequest provides sending stripe-flavored untyped HTTP calls
package rawrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v80"
)

// Client is used to invoke /quotes APIs.
type Client struct {
	B   stripe.RawRequestBackend
	Key string
}

func (c Client) RawRequest(method string, path string, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return c.B.RawRequest(method, path, c.Key, content, params)
}

func Get(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodGet, path, "", params)
}

func Post(path, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodPost, path, content, params)
}

func Delete(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodDelete, path, "", params)
}
