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

// v1CryptoCustomerService is used to invoke /v1/crypto/customers APIs.
type v1CryptoCustomerService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Crypto Customer.
func (c v1CryptoCustomerService) Retrieve(ctx context.Context, id string, params *CryptoCustomerRetrieveParams) (*CryptoCustomer, error) {
	if params == nil {
		params = &CryptoCustomerRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/crypto/customers/%s", id)
	customer := &CryptoCustomer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customer)
	return customer, err
}
