package stripe

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/form"
)

// SourceParams is a union struct used to describe an
// arbitrary payment source.
type SourceParams struct {
	Card  *CardParams `form:"-"`
	Token string      `form:"source"`
}

func (p *SourceParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Card != nil {
		p.Card.AppendToAsCardSourceOrExternalAccount(body, keyParts)
	}
}

// CustomerSourceParams are used to manipulate a given Stripe
// Customer object's payment sources.
// For more details see https://stripe.com/docs/api#sources
type CustomerSourceParams struct {
	Params   `form:"*"`
	Customer string        `form:"-"` // Goes in the URL
	Source   *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}

// SourceVerifyParams are used to verify a customer source
// For more details see https://stripe.com/docs/guides/ach-beta
type SourceVerifyParams struct {
	Params   `form:"*"`
	Amounts  [2]int64 `form:"amounts"` // Amounts is used when verifying bank accounts
	Customer string   `form:"-"`       // Goes in the URL
	Values   []string `form:"values"`  // Values is used when verifying sources
}

// SetSource adds valid sources to a CustomerSourceParams object,
// returning an error for unsupported sources.
func (cp *CustomerSourceParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	cp.Source = source
	return err
}

// SourceParamsFor creates SourceParams objects around supported
// payment sources, returning errors if not.
//
// Currently supported source types are Card (CardParams) and
// Tokens/IDs (string), where Tokens could be single use card
// tokens or bitcoin receiver ids
func SourceParamsFor(obj interface{}) (*SourceParams, error) {
	var sp *SourceParams
	var err error
	switch p := obj.(type) {
	case *CardParams:
		sp = &SourceParams{
			Card: p,
		}
	case string:
		sp = &SourceParams{
			Token: p,
		}
	default:
		err = fmt.Errorf("Unsupported source type %s", p)
	}
	return sp, err
}

// PaymentSourceType consts represent valid payment sources.
type PaymentSourceType string

const (
	// PaymentSourceAccount is a constant representing a payment source which is
	// an account.
	PaymentSourceAccount PaymentSourceType = "account"

	// PaymentSourceBankAccount is a constant representing a payment source
	// which is a bank account.
	PaymentSourceBankAccount PaymentSourceType = "bank_account"

	// PaymentSourceBitcoinReceiver is a constant representing a payment source
	// which is a Bitcoin receiver.
	PaymentSourceBitcoinReceiver PaymentSourceType = "bitcoin_receiver"

	// PaymentSourceCard is a constant representing a payment source which is a
	// card.
	PaymentSourceCard PaymentSourceType = "card"

	// PaymentSourceObject is a constant representing a payment source which
	// is a top level source object (/v1/sources).
	PaymentSourceObject PaymentSourceType = "source"
)

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BitcoinReceiver or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	BankAccount     *BankAccount      `json:"-"`
	BitcoinReceiver *BitcoinReceiver  `json:"-"`
	Card            *Card             `json:"-"`
	Deleted         bool              `json:"deleted"`
	ID              string            `json:"id"`
	SourceObject    *Source           `json:"-"`
	Type            PaymentSourceType `json:"object"`
}

// SourceList is a list object for cards.
type SourceList struct {
	ListMeta
	Values []*PaymentSource `json:"data"`
}

// SourceListParams are used to enumerate the payment sources that are attached
// to a Customer.
type SourceListParams struct {
	ListParams `form:"*"`
	Customer   string `form:"-"` // Handled in URL
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
		case PaymentSourceBankAccount:
			json.Unmarshal(data, &s.BankAccount)
		case PaymentSourceBitcoinReceiver:
			json.Unmarshal(data, &s.BitcoinReceiver)
		case PaymentSourceCard:
			json.Unmarshal(data, &s.Card)
		case PaymentSourceObject:
			json.Unmarshal(data, &s.SourceObject)
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a PaymentSource.
// This custom marshaling is needed because the specific type
// of payment instrument it represents is specified by the PaymentSourceType
func (s *PaymentSource) MarshalJSON() ([]byte, error) {
	var target interface{}

	switch s.Type {
	case PaymentSourceBitcoinReceiver:
		target = struct {
			*BitcoinReceiver
			Type PaymentSourceType `json:"object"`
		}{
			BitcoinReceiver: s.BitcoinReceiver,
			Type:            s.Type,
		}
	case PaymentSourceCard:
		var customerID *string
		if s.Card.Customer != nil {
			customerID = &s.Card.Customer.ID
		}

		target = struct {
			*Card
			Customer *string           `json:"customer"`
			Type     PaymentSourceType `json:"object"`
		}{
			Card:     s.Card,
			Customer: customerID,
			Type:     s.Type,
		}
	case PaymentSourceAccount:
		target = struct {
			ID   string            `json:"id"`
			Type PaymentSourceType `json:"object"`
		}{
			ID:   s.ID,
			Type: s.Type,
		}
	case "":
		target = s.ID
	}

	return json.Marshal(target)
}
