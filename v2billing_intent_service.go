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

// v2BillingIntentService is used to invoke intent related APIs.
type v2BillingIntentService struct {
	B   Backend
	Key string
}

// Create a BillingIntent.
func (c v2BillingIntentService) Create(ctx context.Context, params *V2BillingIntentCreateParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentCreateParams{}
	}
	params.Context = ctx
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodPost, "/v2/billing/intents", c.Key, params, intent)
	return intent, err
}

// Retrieve a BillingIntent.
func (c v2BillingIntentService) Retrieve(ctx context.Context, id string, params *V2BillingIntentRetrieveParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/intents/%s", id)
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, intent)
	return intent, err
}

// Cancel a BillingIntent.
func (c v2BillingIntentService) Cancel(ctx context.Context, id string, params *V2BillingIntentCancelParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/intents/%s/cancel", id)
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Commit a BillingIntent.
func (c v2BillingIntentService) Commit(ctx context.Context, id string, params *V2BillingIntentCommitParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentCommitParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/intents/%s/commit", id)
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Release a BillingIntent.
func (c v2BillingIntentService) ReleaseReservation(ctx context.Context, id string, params *V2BillingIntentReleaseReservationParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentReleaseReservationParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/intents/%s/release_reservation", id)
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Reserve a BillingIntent.
func (c v2BillingIntentService) Reserve(ctx context.Context, id string, params *V2BillingIntentReserveParams) (*V2BillingIntent, error) {
	if params == nil {
		params = &V2BillingIntentReserveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/intents/%s/reserve", id)
	intent := &V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// List BillingIntents.
func (c v2BillingIntentService) List(ctx context.Context, listParams *V2BillingIntentListParams) Seq2[*V2BillingIntent, error] {
	if listParams == nil {
		listParams = &V2BillingIntentListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/intents", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingIntent], error) {
		page := &V2Page[*V2BillingIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
