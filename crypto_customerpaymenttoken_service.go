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

// v1CryptoCustomerPaymentTokenService is used to invoke /v1/crypto/customers/{id}/payment_tokens APIs.
type v1CryptoCustomerPaymentTokenService struct {
	B   Backend
	Key string
}

// Lists the Payment Tokens for a Crypto Customer.
func (c v1CryptoCustomerPaymentTokenService) List(ctx context.Context, listParams *CryptoCustomerPaymentTokenListParams) *V1List[*CryptoCustomerPaymentToken] {
	if listParams == nil {
		listParams = &CryptoCustomerPaymentTokenListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/crypto/customers/%s/payment_tokens", StringValue(listParams.ID))
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*CryptoCustomerPaymentToken], error) {
		list := &v1Page[*CryptoCustomerPaymentToken]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
