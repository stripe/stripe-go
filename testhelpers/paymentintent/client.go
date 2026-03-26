//
//
// File generated from our OpenAPI spec
//
//

// Package paymentintent provides the /v1/payment_intents APIs
package paymentintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_intents APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Simulate an incoming crypto deposit for a testmode PaymentIntent with payment_method_options[crypto][mode]=deposit. The transaction_hash parameter determines whether the simulated deposit succeeds or fails. Learn more about [testing your integration](https://docs.stripe.com/docs/payments/deposit-mode-stablecoin-payments#test-your-integration).
func SimulateCryptoDeposit(id string, params *stripe.TestHelpersPaymentIntentSimulateCryptoDepositParams) (*stripe.PaymentIntent, error) {
	return getC().SimulateCryptoDeposit(id, params)
}

// Simulate an incoming crypto deposit for a testmode PaymentIntent with payment_method_options[crypto][mode]=deposit. The transaction_hash parameter determines whether the simulated deposit succeeds or fails. Learn more about [testing your integration](https://docs.stripe.com/docs/payments/deposit-mode-stablecoin-payments#test-your-integration).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SimulateCryptoDeposit(id string, params *stripe.TestHelpersPaymentIntentSimulateCryptoDepositParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/payment_intents/%s/simulate_crypto_deposit", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
