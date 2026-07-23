//
//
// File generated from our OpenAPI spec
//
//

// Package depositaddress provides the /v1/crypto/deposit_addresses APIs
package depositaddress

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/crypto/deposit_addresses APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
func New(params *stripe.CryptoDepositAddressParams) (*stripe.CryptoDepositAddress, error) {
	return getC().New(params)
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.CryptoDepositAddressParams) (*stripe.CryptoDepositAddress, error) {
	depositaddress := &stripe.CryptoDepositAddress{}
	err := c.B.Call(
		http.MethodPost, "/v1/crypto/deposit_addresses", c.Key, params, depositaddress)
	return depositaddress, err
}

// Retrieves the details of an existing crypto deposit address by ID.
func Get(id string, params *stripe.CryptoDepositAddressParams) (*stripe.CryptoDepositAddress, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing crypto deposit address by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CryptoDepositAddressParams) (*stripe.CryptoDepositAddress, error) {
	path := stripe.FormatURLPath("/v1/crypto/deposit_addresses/%s", id)
	depositaddress := &stripe.CryptoDepositAddress{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, depositaddress)
	return depositaddress, err
}

// Lists crypto deposit addresses for the authenticated merchant.
// Supports cursor-based pagination and optional filtering by customer, network, or on-chain address.
func List(params *stripe.CryptoDepositAddressListParams) *Iter {
	return getC().List(params)
}

// Lists crypto deposit addresses for the authenticated merchant.
// Supports cursor-based pagination and optional filtering by customer, network, or on-chain address.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CryptoDepositAddressListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CryptoDepositAddressList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/crypto/deposit_addresses", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for crypto deposit addresses.
type Iter struct {
	*stripe.Iter
}

// CryptoDepositAddress returns the crypto deposit address which the iterator is currently pointing to.
func (i *Iter) CryptoDepositAddress() *stripe.CryptoDepositAddress {
	return i.Current().(*stripe.CryptoDepositAddress)
}

// CryptoDepositAddressList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CryptoDepositAddressList() *stripe.CryptoDepositAddressList {
	return i.List().(*stripe.CryptoDepositAddressList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
