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

// v2BillingServiceActionService is used to invoke serviceaction related APIs.
type v2BillingServiceActionService struct {
	B   Backend
	Key string
}

// Create a ServiceAction object.
func (c v2BillingServiceActionService) Create(ctx context.Context, params *V2BillingServiceActionCreateParams) (*V2BillingServiceAction, error) {
	if params == nil {
		params = &V2BillingServiceActionCreateParams{}
	}
	params.Context = ctx
	serviceaction := &V2BillingServiceAction{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/service_actions", c.Key, params, serviceaction)
	return serviceaction, err
}

// Retrieve a ServiceAction object.
func (c v2BillingServiceActionService) Retrieve(ctx context.Context, id string, params *V2BillingServiceActionRetrieveParams) (*V2BillingServiceAction, error) {
	if params == nil {
		params = &V2BillingServiceActionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/service_actions/%s", id)
	serviceaction := &V2BillingServiceAction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, serviceaction)
	return serviceaction, err
}
