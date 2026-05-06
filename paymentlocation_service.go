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

// Retrieve a Payment Location.
func (c v1PaymentLocationService) Retrieve(ctx context.Context, id string, params *PaymentLocationRetrieveParams) (*PaymentLocation, error) {
	if params == nil {
		params = &PaymentLocationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &PaymentLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}

// Update a Payment Location.
func (c v1PaymentLocationService) Update(ctx context.Context, id string, params *PaymentLocationUpdateParams) (*PaymentLocation, error) {
	if params == nil {
		params = &PaymentLocationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &PaymentLocation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}

// Delete a Payment Location.
func (c v1PaymentLocationService) Delete(ctx context.Context, id string, params *PaymentLocationDeleteParams) (*PaymentLocation, error) {
	if params == nil {
		params = &PaymentLocationDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &PaymentLocation{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}
