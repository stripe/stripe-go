//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulated data for the card_present payment method
type TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams struct {
	// Card Number
	Number *string `form:"number"`
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
type TestHelpersTerminalReaderPresentPaymentMethodParams struct {
	Params `form:"*"`
	// Simulated data for the card_present payment method
	CardPresent *TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams `form:"card_present"`
	// Simulated payment type
	Type *string `form:"type"`
}
