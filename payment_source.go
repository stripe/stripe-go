package stripe

import (
	"encoding/json"
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

// Displayer provides a human readable representation of a struct
type Displayer interface {
	Display() string
}

// PaymentSourceType consts represent valid payment sources
type PaymentSourceType string

const (
	PaymentSourceBitcoinReceiver PaymentSourceType = "bitcoin_receiver"
	PaymentSourceCard            PaymentSourceType = "card"
)

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BitcoinReceiver or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	Type            PaymentSourceType `json:"object"`
	ID              string            `json:"id"`
	Card            *Card             `json:"-"`
	BitcoinReceiver *BitcoinReceiver  `json:"-"`
}

// Display human readable representation of source.
func (s *PaymentSource) Display() string {
	switch s.Type {
	case PaymentSourceBitcoinReceiver:
		return s.BitcoinReceiver.Display()
	case PaymentSourceCard:
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
		case PaymentSourceBitcoinReceiver:
			json.Unmarshal(data, &s.BitcoinReceiver)
		case PaymentSourceCard:
			json.Unmarshal(data, &s.Card)
		}
	}

	return nil
}
