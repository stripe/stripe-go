// Package rawrequest provides sending stripe-flavored untyped HTTP calls
package rawrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v79"
)

func Get(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodGet, path, "", params)
}
func Post(path, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodPost, path, content, params)
}
func Delete(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodDelete, path, "", params)
}
