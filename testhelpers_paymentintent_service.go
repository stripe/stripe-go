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

// v1TestHelpersPaymentIntentService is used to invoke /v1/payment_intents APIs.
type v1TestHelpersPaymentIntentService struct {
	B   Backend
	Key string
}

// Simulate an incoming crypto deposit for a testmode PaymentIntent with payment_method_options[crypto][mode]=deposit. The transaction_hash parameter determines whether the simulated deposit succeeds or fails. Learn more about [testing your integration](https://docs.stripe.com/docs/payments/deposit-mode-stablecoin-payments#test-your-integration).
func (c v1TestHelpersPaymentIntentService) SimulateCryptoDeposit(ctx context.Context, id string, params *TestHelpersPaymentIntentSimulateCryptoDepositParams) (*PaymentIntent, error) {
	if params == nil {
		params = &TestHelpersPaymentIntentSimulateCryptoDepositParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/payment_intents/%s/simulate_crypto_deposit", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}
