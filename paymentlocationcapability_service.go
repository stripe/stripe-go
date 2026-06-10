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

// v1PaymentLocationCapabilityService is used to invoke /v1/payment_location_capabilities APIs.
type v1PaymentLocationCapabilityService struct {
	B   Backend
	Key string
}

// Retrieves a payment_location capability
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

// Updates a payment_location capability. Request or remove a payment_location capability by updating its requested parameter.
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

// List all payment location capabilities associated with the payment location.
func (c v1PaymentLocationCapabilityService) List(ctx context.Context, listParams *PaymentLocationCapabilityListParams) *V1List[*PaymentLocationCapability] {
	if listParams == nil {
		listParams = &PaymentLocationCapabilityListParams{}
	}
	p := listParams.GetParams()
	p.Context = ctx
	queryParams := &form.Values{}
	form.AppendTo(queryParams, listParams)
	return newV1List(ctx, nil, func(ctx context.Context, _ *Params, _ *form.Values) (*v1Page[*PaymentLocationCapability], error) {
		list := &v1Page[*PaymentLocationCapability]{}
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_location_capabilities", c.Key, []byte(queryParams.Encode()), p, list)
		return list, err
	})
}
