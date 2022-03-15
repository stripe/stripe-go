//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v72/form"
)

type PaymentSourceType string

// List of values that PaymentSourceType can take
const (
	PaymentSourceTypeAccount     PaymentSourceType = "account"
	PaymentSourceTypeBankAccount PaymentSourceType = "bank_account"
	PaymentSourceTypeCard        PaymentSourceType = "card"
	PaymentSourceTypeObject      PaymentSourceType = "source"
)

// List sources for a specified customer.
type SourceListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	// Filter sources according to a particular object type.
	Object *string `form:"object"`
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
// tokens
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
			Token: &p,
		}
	default:
		err = fmt.Errorf("Unsupported source type %s", p)
	}
	return sp, err
}

// SourceParams is a union struct used to describe an
// arbitrary payment source.
type SourceParams struct {
	Card  *CardParams `form:"-"`
	Token *string     `form:"source"`
}

// AppendTo implements custom encoding logic for SourceParams.
func (p *SourceParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Card != nil {
		p.Card.AppendToAsCardSourceOrExternalAccount(body, keyParts)
	}
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://stripe.com/docs/api#update_customer) to have a new default_source.
type CustomerSourceParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// The name of the person or business that owns the bank account.
	AccountHolderName *string `form:"account_holder_name"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType *string `form:"account_holder_type"`
	// City/District/Suburb/Town/Village.
	AddressCity *string `form:"address_city"`
	// Billing address country, if provided when creating card.
	AddressCountry *string `form:"address_country"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 *string `form:"address_line1"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 *string `form:"address_line2"`
	// State/County/Province/Region.
	AddressState *string `form:"address_state"`
	// ZIP or postal code.
	AddressZip *string `form:"address_zip"`
	// Two digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month"`
	// Four digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year"`
	// Cardholder name.
	Name  *string                   `form:"name"`
	Owner *PaymentSourceOwnerParams `form:"owner"`
	// Please refer to full [documentation](https://stripe.com/docs/api) instead.
	Source *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}
type PaymentSourceOwnerParams struct {
	// Owner's address.
	Address *AddressParams `form:"address"`
	// Owner's email address.
	Email *string `form:"email"`
	// Owner's full name.
	Name *string `form:"name"`
	// Owner's phone number.
	Phone *string `form:"phone"`
}

// Verify a specified bank account for a given customer.
type SourceVerifyParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.
	Amounts [2]int64  `form:"amounts"` // Amounts is used when verifying bank accounts
	Values  []*string `form:"values"`  // Values is used when verifying sources
}
type PaymentSource struct {
	APIResource
	BankAccount  *BankAccount      `json:"-"`
	Card         *Card             `json:"-"`
	Deleted      bool              `json:"deleted"`
	ID           string            `json:"id"`
	SourceObject *Source           `json:"-"`
	Type         PaymentSourceType `json:"object"`
}

// SourceList is a list of PaymentSources as retrieved from a list endpoint.
type SourceList struct {
	APIResource
	ListMeta
	Data []*PaymentSource `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentSource.
// This custom unmarshaling is needed because the specific
// type of payment instrument it refers to is specified in the JSON
func (s *PaymentSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type paymentSource PaymentSource
	var v paymentSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*s = PaymentSource(v)

	switch s.Type {
	case PaymentSourceTypeBankAccount:
		err = json.Unmarshal(data, &s.BankAccount)
	case PaymentSourceTypeCard:
		err = json.Unmarshal(data, &s.Card)
	case PaymentSourceTypeObject:
		err = json.Unmarshal(data, &s.SourceObject)
	}

	return err
}

// MarshalJSON handles serialization of a PaymentSource.
// This custom marshaling is needed because the specific type
// of payment instrument it represents is specified by the Type
func (s *PaymentSource) MarshalJSON() ([]byte, error) {
	var target interface{}

	switch s.Type {
	case PaymentSourceTypeCard:
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
	case PaymentSourceTypeAccount:
		target = struct {
			ID   string            `json:"id"`
			Type PaymentSourceType `json:"object"`
		}{
			ID:   s.ID,
			Type: s.Type,
		}
	case PaymentSourceTypeBankAccount:
		var customerID *string
		if s.BankAccount.Customer != nil {
			customerID = &s.BankAccount.Customer.ID
		}

		target = struct {
			*BankAccount
			Customer *string           `json:"customer"`
			Type     PaymentSourceType `json:"object"`
		}{
			BankAccount: s.BankAccount,
			Customer:    customerID,
			Type:        s.Type,
		}
	case "":
		target = s.ID
	}

	return json.Marshal(target)
}
