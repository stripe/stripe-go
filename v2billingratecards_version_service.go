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

// v2BillingRateCardsVersionService is used to invoke version related APIs.
type v2BillingRateCardsVersionService struct {
	B   Backend
	Key string
}

// Retrieve a specific version of a Rate Card object.
func (c v2BillingRateCardsVersionService) Retrieve(ctx context.Context, id string, params *V2BillingRateCardsVersionRetrieveParams) (*V2BillingRateCardVersion, error) {
	if params == nil {
		params = &V2BillingRateCardsVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/versions/%s", StringValue(params.RateCardID), id)
	ratecardversion := &V2BillingRateCardVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardversion)
	return ratecardversion, err
}

// List the versions of a Rate Card object. Results are sorted in reverse chronological order (most recent first).
func (c v2BillingRateCardsVersionService) List(ctx context.Context, listParams *V2BillingRateCardsVersionListParams) Seq2[*V2BillingRateCardVersion, error] {
	if listParams == nil {
		listParams = &V2BillingRateCardsVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/versions", StringValue(listParams.RateCardID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingRateCardVersion], error) {
		page := &V2Page[*V2BillingRateCardVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
