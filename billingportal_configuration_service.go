//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1BillingPortalConfigurationService is used to invoke /v1/billing_portal/configurations APIs.
type v1BillingPortalConfigurationService struct {
	B   Backend
	Key string
}

// Creates a configuration that describes the functionality and behavior of a PortalSession
func (c v1BillingPortalConfigurationService) Create(ctx context.Context, params *BillingPortalConfigurationCreateParams) (*BillingPortalConfiguration, error) {
	if params == nil {
		params = &BillingPortalConfigurationCreateParams{}
	}
	params.Context = ctx
	configuration := &BillingPortalConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing_portal/configurations", c.Key, params, configuration)
	return configuration, err
}

// Retrieves a configuration that describes the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) Retrieve(ctx context.Context, id string, params *BillingPortalConfigurationRetrieveParams) (*BillingPortalConfiguration, error) {
	if params == nil {
		params = &BillingPortalConfigurationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing_portal/configurations/%s", id)
	configuration := &BillingPortalConfiguration{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, configuration)
	return configuration, err
}

// Updates a configuration that describes the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) Update(ctx context.Context, id string, params *BillingPortalConfigurationUpdateParams) (*BillingPortalConfiguration, error) {
	if params == nil {
		params = &BillingPortalConfigurationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing_portal/configurations/%s", id)
	configuration := &BillingPortalConfiguration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, configuration)
	return configuration, err
}

// Returns a list of configurations that describe the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) List(ctx context.Context, listParams *BillingPortalConfigurationListParams) Seq2[*BillingPortalConfiguration, error] {
	if listParams == nil {
		listParams = &BillingPortalConfigurationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingPortalConfiguration, ListContainer, error) {
		list := &BillingPortalConfigurationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing_portal/configurations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
