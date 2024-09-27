// Package rawrequest provides sending stripe-flavored untyped HTTP calls
package rawrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v79"
)

func getDefaultRequestOptions(params *stripe.RawParams) *stripe.RawParams {
	if params == nil {
		params = &stripe.RawParams{}
	}

	rawParams := stripe.RawParams{
		Params:        params.Params,
		APIMode:       stripe.V2APIMode, // most of this usage is probably v2, so default to that mode
		StripeContext: params.StripeContext,
	}

	return &rawParams
}

func Get(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodGet, path, "", getDefaultRequestOptions(params))
}
func Post(path, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodPost, path, content, getDefaultRequestOptions(params))
}
func Delete(path string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodDelete, path, "", getDefaultRequestOptions(params))
}
