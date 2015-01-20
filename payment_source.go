package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/paymentsource"
)

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

// Interface to provide a human readable representation of a struct
type Displayer interface {
	Display() string
}

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BitcoinReceiver or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	Type            string `json:"object"`
	ID              string `json:"id"`
	Card            *Card
	BitcoinReceiver *BitcoinReceiver
}

// Human readable/displayable value for any payment type
func (s *PaymentSource) Display() string {
	switch s.Type {
	case paymentsource.BitcoinReceiver:
		return s.BitcoinReceiver.Display()
	case paymentsource.Card:
		return s.Card.Display()
	}

	return ""
}

// UnmarshalJSON handles deserialization of a PaymentSource.
// This custom unmarshaling is needed because the specific
// type of payment instrument it refers to is specified in the JSON
func (s *PaymentSource) UnmarshalJSON(data []byte) error {
	type source PaymentSource
	var ss source
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = PaymentSource(ss)

		switch s.Type {
		case paymentsource.BitcoinReceiver:
			s.BitcoinReceiver = &BitcoinReceiver{}
			json.Unmarshal(data, s.BitcoinReceiver)
		case paymentsource.Card:
			s.Card = &Card{}
			json.Unmarshal(data, s.Card)
		}
	}

	return nil
}
