//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
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

// Retrieves the trial offer with the given ID.
func (c v1ProductCatalogTrialOfferService) Retrieve(ctx context.Context, id string, params *ProductCatalogTrialOfferRetrieveParams) (*ProductCatalogTrialOffer, error) {
	if params == nil {
		params = &ProductCatalogTrialOfferRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/product_catalog/trial_offers/%s", id)
	trialoffer := &ProductCatalogTrialOffer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, trialoffer)
	return trialoffer, err
}

// Returns a list of trial offers.
func (c v1ProductCatalogTrialOfferService) List(ctx context.Context, listParams *ProductCatalogTrialOfferListParams) *V1List[*ProductCatalogTrialOffer] {
	if listParams == nil {
		listParams = &ProductCatalogTrialOfferListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*ProductCatalogTrialOffer], error) {
		list := &v1Page[*ProductCatalogTrialOffer]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/product_catalog/trial_offers", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
