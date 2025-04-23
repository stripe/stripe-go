// Package rawrequest provides sending stripe-flavored untyped HTTP calls
package rawrequest

import (
	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke raw requests against the specified backend
type Client struct {
	B   stripe.RawRequestBackend
	Key string
}

func (c Client) RawRequest(method string, path string, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return c.B.RawRequest(method, path, c.Key, content, params)
}
