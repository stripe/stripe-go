package stripe

import (
	"encoding/json"
	"net/url"
)

// SourceParams is the set of parameters that can be used to describe the source
// object used to make a Charge. For example, both Card and BitcoinReceiver objects
// can be described by SourceParams.
// For more details see https://stripe.com/docs/api#create_charge
type SourceParams struct {
	Params
	ID       string
	Token    string
	Card     *CardParams
	Customer string
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

// SourceList is a list object for cards.
type SourceList struct {
	ListMeta
	Values []*PaymentSource `json:"data"`
}

// AppendDetails adds the source's details to the query string values.
// When creating a new source, the parameters are passed as a dictionary, but
// on updates they are simply the parameter name.
func (s *SourceParams) AppendDetails(values *url.Values, creating bool) {
	if s.Card != nil {
		s.Card.AppendDetails(values, creating)
	}

	if len(s.ID) > 0 {
		values.Add("source", s.ID)
	}

	if len(s.Token) > 0 {
		values.Add("source", s.Token)
	}
}

// PaymentSourceListParams are used to enumerate the payment sources
// that are attached to a Customer.
type SourceListParams struct {
	ListParams
	Customer string
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
	} else {
		// the id is surrounded by "\" characters, so strip them
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}
