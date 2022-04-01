//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulated card present data
type TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams struct {
	// Card Number
	Number *string `form:"number"`
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
type TestHelpersTerminalReaderPresentPaymentMethodParams struct {
	Params `form:"*"`
	// Simulated card present data
	CardPresent *TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams `form:"card_present"`
	// Simulated payment type
	Type *string `form:"type"`
}
