package stripe

// SourceParams is the set of parameters that can be used to describe the source
// object used to make a Charge. For example, both Card and BitcoinReceiver objects
// can be described by SourceParams.
// For more details see https://stripe.com/docs/api#create_charge
type SourceParams struct {
	Params
	ID    string
	Token string
	Card  *CardParams
}

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BitcoinReceiver or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	Type            string
	Card            *Card
	BitcoinReceiver *BitcoinReceiver
}
