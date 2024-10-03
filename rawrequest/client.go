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

func Get(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodGet, path, "", params)
}

func (c Client) Get(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return c.B.RawRequest(http.MethodGet, path, c.Key, "", params)
}

func Post(path, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodPost, path, content, params)
}

func (c Client) Post(path, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return c.B.RawRequest(http.MethodPost, path, c.Key, content, params)
}

func Delete(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodDelete, path, "", params)
}

func (c Client) Delete(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return c.B.RawRequest(http.MethodDelete, path, c.Key, "", params)
}
