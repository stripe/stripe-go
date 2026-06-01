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

// v1DelegatedCheckoutOrderService is used to invoke /v1/delegated_checkout/orders APIs.
type v1DelegatedCheckoutOrderService struct {
	B   Backend
	Key string
}

// Retrieves a delegated checkout order.
func (c v1DelegatedCheckoutOrderService) Retrieve(ctx context.Context, id string, params *DelegatedCheckoutOrderRetrieveParams) (*DelegatedCheckoutOrder, error) {
	if params == nil {
		params = &DelegatedCheckoutOrderRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/delegated_checkout/orders/%s", id)
	order := &DelegatedCheckoutOrder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}
