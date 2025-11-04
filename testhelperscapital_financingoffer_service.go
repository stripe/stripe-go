//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TestHelpersCapitalFinancingOfferService is used to invoke /v1/capital/financing_offers APIs.
type v1TestHelpersCapitalFinancingOfferService struct {
	B   Backend
	Key string
}

// Creates a test financing offer for a connected account.
func (c v1TestHelpersCapitalFinancingOfferService) Create(ctx context.Context, params *TestHelpersCapitalFinancingOfferCreateParams) (*CapitalFinancingOffer, error) {
	if params == nil {
		params = &TestHelpersCapitalFinancingOfferCreateParams{}
	}
	params.Context = ctx
	financingoffer := &CapitalFinancingOffer{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/capital/financing_offers", c.Key, params, financingoffer)
	return financingoffer, err
}

// Refills a test financing offer for a connected account.
func (c v1TestHelpersCapitalFinancingOfferService) Refill(ctx context.Context, id string, params *TestHelpersCapitalFinancingOfferRefillParams) (*CapitalFinancingOffer, error) {
	if params == nil {
		params = &TestHelpersCapitalFinancingOfferRefillParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/capital/financing_offers/%s/refill", id)
	financingoffer := &CapitalFinancingOffer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}
