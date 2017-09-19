package stripe

import (
	"encoding/json"
	"strconv"

	"github.com/stripe/stripe-go/form"
)

// cardSource is a string that's used to build card form parameters. It's a
// constant just to make mistakes less likely.
const cardSource = "source"

// CardBrand is the list of allowed values for the card's brand.
// Allowed values are "Unknown", "Visa", "American Express", "MasterCard", "Discover"
// "JCB", "Diners Club".
type CardBrand string

// CardFunding is the list of allowed values for the card's funding.
// Allowed values are "credit", "debit", "prepaid", "unknown".
type CardFunding string

// TokenizationMethod is the list of allowed values for the card's tokenization method.
// Allowed values are "apple_pay", "android_pay".
type TokenizationMethod string

// Verification is the list of allowed verification responses.
// Allowed values are "pass", "fail", "unchecked", "unavailable".
type Verification string

// CardParams is the set of parameters that can be used when creating or updating a card.
// For more details see https://stripe.com/docs/api#create_card and https://stripe.com/docs/api#update_card.
//
// Note that while form annotations are used for tokenization and updates,
// cards have some unusual logic on creates that necessitates manual handling
// of all parameters. See AppendToAsCardSourceOrExternalAccount.
type CardParams struct {
	Params             `form:"*"`
	Account            string `form:"-"`
	AddressCity        string `form:"address_city"`
	AddressCountry     string `form:"address_country"`
	AddressLine1       string `form:"address_line1"`
	AddressLine2       string `form:"address_line2"`
	AddressState       string `form:"address_state"`
	AddressZip         string `form:"address_zip"`
	CVC                string `form:"cvc"`
	Currency           string `form:"currency"`
	Customer           string `form:"-"`
	DefaultForCurrency bool   `form:"default_for_currency"`
	ExpMonth           string `form:"exp_month"`
	ExpYear            string `form:"exp_year"`
	Name               string `form:"name"`
	Number             string `form:"number"`
	Recipient          string `form:"-"`
	Token              string `form:"-"`

	// ID is used when tokenizing a card for shared customers
	ID string `form:"*"`
}

// AppendToAsCardSourceOrExternalAccount appends the given CardParams as either a
// card or external account.
//
// It may look like an AppendTo from the form package, but it's not, and is
// only used in the special case where we use `card.New`. It's needed because
// we have some weird encoding logic here that can't be handled by the form
// package (and it's special enough that it wouldn't be desirable to have it do
// so).
//
// This is not a pattern that we want to push forward, and this largely exists
// because the cards endpoint is a little unusual. There is one other resource
// like it, which is bank account.
func (c *CardParams) AppendToAsCardSourceOrExternalAccount(body *form.Values, keyParts []string) {
	if c.DefaultForCurrency {
		body.Add(form.FormatKey(append(keyParts, "default_for_currency")), strconv.FormatBool(c.DefaultForCurrency))
	}

	if len(c.Token) > 0 {
		if len(c.Account) > 0 {
			body.Add(form.FormatKey(append(keyParts, "external_account")), c.Token)
		} else {
			body.Add(form.FormatKey(append(keyParts, cardSource)), c.Token)
		}
	}

	if len(c.Number) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "object")), "card")
		body.Add(form.FormatKey(append(keyParts, cardSource, "number")), c.Number)
	}

	if len(c.CVC) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "cvc")), c.CVC)
	}

	if len(c.Currency) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "currency")), c.Currency)
	}

	if len(c.ExpMonth) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "exp_month")), c.ExpMonth)
	}

	if len(c.ExpYear) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "exp_year")), c.ExpYear)
	}

	if len(c.Name) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "name")), c.Name)
	}

	if len(c.AddressCity) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_city")), c.AddressCity)
	}

	if len(c.AddressCountry) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_country")), c.AddressCountry)
	}

	if len(c.AddressLine1) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_line1")), c.AddressLine1)
	}

	if len(c.AddressLine2) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_line2")), c.AddressLine2)
	}

	if len(c.AddressState) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_state")), c.AddressState)
	}

	if len(c.AddressZip) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_zip")), c.AddressZip)
	}
}

// CardListParams is the set of parameters that can be used when listing cards.
// For more details see https://stripe.com/docs/api#list_cards.
type CardListParams struct {
	ListParams `form:"*"`
	Account    string `form:"-"`
	Customer   string `form:"-"`
	Recipient  string `form:"-"`
}

// Card is the resource representing a Stripe credit/debit card.
// For more details see https://stripe.com/docs/api#cards.
type Card struct {
	AddressCity        string       `json:"address_city"`
	AddressCountry     string       `json:"address_country"`
	AddressLine1       string       `json:"address_line1"`
	AddressLine1Check  Verification `json:"address_line1_check"`
	AddressLine2       string       `json:"address_line2"`
	AddressState       string       `json:"address_state"`
	AddressZip         string       `json:"address_zip"`
	AddressZipCheck    Verification `json:"address_zip_check"`
	Brand              CardBrand    `json:"brand"`
	CVCCheck           Verification `json:"cvc_check"`
	Country            string       `json:"country"`
	Currency           Currency     `json:"currency"`
	Customer           *Customer    `json:"customer"`
	DefaultForCurrency bool         `json:"default_for_currency"`
	Deleted            bool         `json:"deleted"`

	// Description is a succinct summary of the card's information.
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	Description string `json:"description"`

	DynamicLast4 string      `json:"dynamic_last4"`
	ExpMonth     uint8       `json:"exp_month"`
	ExpYear      uint16      `json:"exp_year"`
	Fingerprint  string      `json:"fingerprint"`
	Funding      CardFunding `json:"funding"`
	ID           string      `json:"id"`

	// IIN is the card's "Issuer Identification Number".
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	IIN string `json:"iin"`

	// Issuer is a bank or financial institution that provides the card.
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	Issuer string `json:"issuer"`

	Last4              string             `json:"last4"`
	Metadata           map[string]string  `json:"metadata"`
	Name               string             `json:"name"`
	Recipient          *Recipient         `json:"recipient"`
	ThreeDSecure       *ThreeDSecure      `json:"three_d_secure"`
	TokenizationMethod TokenizationMethod `json:"tokenization_method"`
}

// CardList is a list object for cards.
type CardList struct {
	ListMeta
	Data []*Card `json:"data"`
}

// UnmarshalJSON handles deserialization of a Card.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Card) UnmarshalJSON(data []byte) error {
	type card Card
	var cc card
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Card(cc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}

	return nil
}
