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

// v1CryptoDepositAddressService is used to invoke /v1/crypto/deposit_addresses APIs.
type v1CryptoDepositAddressService struct {
	B   Backend
	Key string
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
func (c v1CryptoDepositAddressService) Create(ctx context.Context, params *CryptoDepositAddressCreateParams) (*CryptoDepositAddress, error) {
	if params == nil {
		params = &CryptoDepositAddressCreateParams{}
	}
	params.Context = ctx
	depositaddress := &CryptoDepositAddress{}
	err := c.B.Call(
		http.MethodPost, "/v1/crypto/deposit_addresses", c.Key, params, depositaddress)
	return depositaddress, err
}

// Retrieves the details of an existing crypto deposit address by ID.
func (c v1CryptoDepositAddressService) Retrieve(ctx context.Context, id string, params *CryptoDepositAddressRetrieveParams) (*CryptoDepositAddress, error) {
	if params == nil {
		params = &CryptoDepositAddressRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/crypto/deposit_addresses/%s", id)
	depositaddress := &CryptoDepositAddress{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, depositaddress)
	return depositaddress, err
}

// Lists crypto deposit addresses for the authenticated merchant.
// Supports cursor-based pagination and optional filtering by customer, network, or on-chain address.
func (c v1CryptoDepositAddressService) List(ctx context.Context, listParams *CryptoDepositAddressListParams) *V1List[*CryptoDepositAddress] {
	if listParams == nil {
		listParams = &CryptoDepositAddressListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*CryptoDepositAddress], error) {
		list := &v1Page[*CryptoDepositAddress]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/crypto/deposit_addresses", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
