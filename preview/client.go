// Package preview provides preview / private beta APIs
package preview

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
)

func GetDefaultRequestOptions(params stripe.RawParamsContainer) stripe.RawParamsContainer {
	rawParams := stripe.RawParams{
		Params:        *params.GetParams(),
		APIMode:       stripe.PreviewAPIMode,
		StripeContext: params.GetStripeContext(),
	}
	return &rawParams
}

func Get(path string, params stripe.RawParamsContainer) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodGet, path, "", GetDefaultRequestOptions(params))
}
func Post(path, content string, params stripe.RawParamsContainer) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodPost, path, content, params)
}
func Delete(path string, params stripe.RawParamsContainer) (*stripe.APIResponse, error) {
	return stripe.RawRequest(http.MethodDelete, path, "", params)
}
