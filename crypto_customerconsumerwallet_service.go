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

// v1CryptoCustomerConsumerWalletService is used to invoke /v1/crypto/customers/{id}/crypto_consumer_wallets APIs.
type v1CryptoCustomerConsumerWalletService struct {
	B   Backend
	Key string
}

// Lists the Consumer Wallets for a Crypto Customer.
func (c v1CryptoCustomerConsumerWalletService) List(ctx context.Context, listParams *CryptoCustomerConsumerWalletListParams) *V1List[*CryptoCustomerConsumerWallet] {
	if listParams == nil {
		listParams = &CryptoCustomerConsumerWalletListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/crypto/customers/%s/crypto_consumer_wallets", StringValue(
			listParams.ID))
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*CryptoCustomerConsumerWallet], error) {
		list := &v1Page[*CryptoCustomerConsumerWallet]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
