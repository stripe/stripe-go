//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulate an incoming crypto deposit for a testmode PaymentIntent with payment_method_options[crypto][mode]=deposit. The transaction_hash parameter determines whether the simulated deposit succeeds or fails. Learn more about [testing your integration](https://docs.stripe.com/docs/payments/deposit-mode-stablecoin-payments#test-your-integration).
type TestHelpersPaymentIntentSimulateCryptoDepositParams struct {
	Params `form:"*"`
	// The buyer's wallet address from which the crypto deposit is originating.
	BuyerWallet *string `form:"buyer_wallet"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The blockchain network of the simulated crypto deposit.
	Network *string `form:"network"`
	// The token currency of the simulated crypto deposit.
	TokenCurrency *string `form:"token_currency"`
	// A testmode transaction hash that determines the outcome of the simulated deposit.
	TransactionHash *string `form:"transaction_hash"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersPaymentIntentSimulateCryptoDepositParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
