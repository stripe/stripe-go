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

// v1PaymentLocationCapabilityService is used to invoke paymentlocationcapability related APIs.
type v1PaymentLocationCapabilityService struct {
	B   Backend
	Key string
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
