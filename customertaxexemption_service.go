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

// v1CustomerTaxExemptionService is used to invoke /v1/customers/{customer}/tax_exemptions APIs.
type v1CustomerTaxExemptionService struct {
	B   Backend
	Key string
}

// Create a location specific tax exemption for a customer.
func (c v1CustomerTaxExemptionService) Create(ctx context.Context, params *CustomerTaxExemptionCreateParams) (*CustomerTaxExemption, error) {
	if params == nil {
		params = &CustomerTaxExemptionCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/tax_exemptions", StringValue(params.Customer))
	customertaxexemption := &CustomerTaxExemption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// Retrieve a location specific tax exemption for a customer.
func (c v1CustomerTaxExemptionService) Retrieve(ctx context.Context, id string, params *CustomerTaxExemptionRetrieveParams) (*CustomerTaxExemption, error) {
	if params == nil {
		params = &CustomerTaxExemptionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/tax_exemptions/%s", StringValue(params.Customer), id)
	customertaxexemption := &CustomerTaxExemption{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// Delete a location specific tax exemption for a customer.
func (c v1CustomerTaxExemptionService) Delete(ctx context.Context, id string, params *CustomerTaxExemptionDeleteParams) (*CustomerTaxExemption, error) {
	if params == nil {
		params = &CustomerTaxExemptionDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/tax_exemptions/%s", StringValue(params.Customer), id)
	customertaxexemption := &CustomerTaxExemption{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// List all location specific tax exemptions for a customer.
func (c v1CustomerTaxExemptionService) List(ctx context.Context, listParams *CustomerTaxExemptionListParams) *V1List[*CustomerTaxExemption] {
	if listParams == nil {
		listParams = &CustomerTaxExemptionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/tax_exemptions", StringValue(listParams.Customer))
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*CustomerTaxExemption], error) {
		list := &v1Page[*CustomerTaxExemption]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
