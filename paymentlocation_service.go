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

// v1PaymentLocationService is used to invoke /v1/payment_locations APIs.
type v1PaymentLocationService struct {
	B   Backend
	Key string
}

// Create a Payment Location.
func (c v1PaymentLocationService) Create(ctx context.Context, params *PaymentLocationCreateParams) (*PaymentLocation, error) {
	if params == nil {
		params = &PaymentLocationCreateParams{}
	}
	params.Context = ctx
	paymentlocation := &PaymentLocation{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_locations", c.Key, params, paymentlocation)
	return paymentlocation, err
}
