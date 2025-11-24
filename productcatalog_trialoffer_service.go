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

// v1ProductCatalogTrialOfferService is used to invoke /v1/product_catalog/trial_offers APIs.
type v1ProductCatalogTrialOfferService struct {
	B   Backend
	Key string
}

// Creates a trial offer.
func (c v1ProductCatalogTrialOfferService) Create(ctx context.Context, params *ProductCatalogTrialOfferCreateParams) (*ProductCatalogTrialOffer, error) {
	if params == nil {
		params = &ProductCatalogTrialOfferCreateParams{}
	}
	params.Context = ctx
	trialoffer := &ProductCatalogTrialOffer{}
	err := c.B.Call(
		http.MethodPost, "/v1/product_catalog/trial_offers", c.Key, params, trialoffer)
	return trialoffer, err
}
