//
//
// File generated from our OpenAPI spec
//
//

// Package customerpaymenttoken provides the /v1/crypto/customers/{id}/payment_tokens APIs
package customerpaymenttoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/crypto/customers/{id}/payment_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists the Payment Tokens for a Crypto Customer.
func List(params *stripe.CryptoCustomerPaymentTokenListParams) *Iter {
	return getC().List(params)
}

// Lists the Payment Tokens for a Crypto Customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CryptoCustomerPaymentTokenListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/crypto/customers/%s/payment_tokens", stripe.StringValue(listParams.ID))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CryptoCustomerPaymentTokenList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for crypto customer payment tokens.
type Iter struct {
	*stripe.Iter
}

// CryptoCustomerPaymentToken returns the crypto customer payment token which the iterator is currently pointing to.
func (i *Iter) CryptoCustomerPaymentToken() *stripe.CryptoCustomerPaymentToken {
	return i.Current().(*stripe.CryptoCustomerPaymentToken)
}

// CryptoCustomerPaymentTokenList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CryptoCustomerPaymentTokenList() *stripe.CryptoCustomerPaymentTokenList {
	return i.List().(*stripe.CryptoCustomerPaymentTokenList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
