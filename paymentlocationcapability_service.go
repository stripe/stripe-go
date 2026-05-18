//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v85/form"
)

// v1PaymentLocationCapabilityService is used to invoke /v1/payment_location_capabilities APIs.
type v1PaymentLocationCapabilityService struct {
	B   Backend
	Key string
}

// Retrieves information about the specified Payment Location Capability.
func (c v1PaymentLocationCapabilityService) Retrieve(ctx context.Context, id string, params *PaymentLocationCapabilityRetrieveParams) (*PaymentLocationCapability, error) {
	if params == nil {
		params = &PaymentLocationCapabilityRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_location_capabilities/%s", id)
	paymentlocationcapability := &PaymentLocationCapability{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, paymentlocationcapability)
	return paymentlocationcapability, err
}

// Updates a specified Payment Location Capability. Request or remove a payment location capability by updating its requested parameter.
func (c v1PaymentLocationCapabilityService) Update(ctx context.Context, id string, params *PaymentLocationCapabilityUpdateParams) (*PaymentLocationCapability, error) {
	if params == nil {
		params = &PaymentLocationCapabilityUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_location_capabilities/%s", id)
	paymentlocationcapability := &PaymentLocationCapability{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, paymentlocationcapability)
	return paymentlocationcapability, err
}

// Returns a list of PaymentLocationCapability objects associated with the location.
func (c v1PaymentLocationCapabilityService) List(ctx context.Context, listParams *PaymentLocationCapabilityListParams) *V1List[*PaymentLocationCapability] {
	if listParams == nil {
		listParams = &PaymentLocationCapabilityListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*PaymentLocationCapability], error) {
		list := &v1Page[*PaymentLocationCapability]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_location_capabilities", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
