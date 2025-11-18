//
//
// File generated from our OpenAPI spec
//
//

// Package financingoffer provides the /v1/capital/financing_offers APIs
package financingoffer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/capital/financing_offers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a test financing offer for a connected account.
func New(params *stripe.TestHelpersCapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().New(params)
}

// Creates a test financing offer for a connected account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TestHelpersCapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/capital/financing_offers", c.Key, params, financingoffer)
	return financingoffer, err
}

// Refills a test financing offer for a connected account.
func Refill(id string, params *stripe.TestHelpersCapitalFinancingOfferRefillParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().Refill(id, params)
}

// Refills a test financing offer for a connected account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Refill(id string, params *stripe.TestHelpersCapitalFinancingOfferRefillParams) (*stripe.CapitalFinancingOffer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/capital/financing_offers/%s/refill", id)
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
