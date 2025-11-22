//
//
// File generated from our OpenAPI spec
//
//

// Package trialoffer provides the /v1/product_catalog/trial_offers APIs
package trialoffer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/product_catalog/trial_offers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a trial offer.
func New(params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	return getC().New(params)
}

// Creates a trial offer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	trialoffer := &stripe.ProductCatalogTrialOffer{}
	err := c.B.Call(
		http.MethodPost, "/v1/product_catalog/trial_offers", c.Key, params, trialoffer)
	return trialoffer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
