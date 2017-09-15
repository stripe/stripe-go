package stripe

import (
	"encoding/json"
	"strconv"

	"github.com/stripe/stripe-go/form"
)

// CardBrand is the list of allowed values for the card's brand.
// Allowed values are "Unknown", "Visa", "American Express", "MasterCard", "Discover"
// "JCB", "Diners Club".
type CardBrand string

// Verification is the list of allowed verification responses.
// Allowed values are "pass", "fail", "unchecked", "unavailable".
type Verification string

// CardFunding is the list of allowed values for the card's funding.
// Allowed values are "credit", "debit", "prepaid", "unknown".
type CardFunding string

// TokenizationMethod is the list of allowed values for the card's tokenization method.
// Allowed values are "apple_pay", "android_pay".
type TokenizationMethod string

// CardParams is the set of parameters that can be used when creating or updating a card.
// For more details see https://stripe.com/docs/api#create_card and https://stripe.com/docs/api#update_card.
//
// Note that while form annotations are used for tokenization and updates,
// cards have some unusual logic on creates that necessitates manual handling
// of all parameters. See AppendToAsCardSourceOrExternalAccount.
type CardParams struct {
	Params    `form:"*"`
	Token     string `form:"-"`
	Default   bool   `form:"default_for_currency"`
	Account   string `form:"-"`
	Customer  string `form:"-"`
	Recipient string `form:"-"`
	Name      string `form:"name"`
	Number    string `form:"number"`
	Month     string `form:"exp_month"`
	Year      string `form:"exp_year"`
	CVC       string `form:"-"`
	Currency  string `form:"currency"`
	Address1  string `form:"address_line1"`
	Address2  string `form:"address_line2"`
	City      string `form:"address_city"`
	State     string `form:"address_state"`
	Zip       string `form:"address_zip"`
	Country   string `form:"address_country"`
}

// cardSource is a string that's used to build card form parameters. It's a
// constant just to make mistakes less likely.
const cardSource = "source"

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
	if len(c.Token) > 0 && len(c.Account) > 0 {
		body.Add(form.FormatKey(append(keyParts, "external_account")), c.Token)

	}

	if c.Default {
		body.Add(form.FormatKey(append(keyParts, "default_for_currency")), strconv.FormatBool(c.Default))
	}

	if len(c.Token) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource)), c.Token)
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

	if c.Default {
		body.Add(form.FormatKey(append(keyParts, cardSource, "default_for_currency")), strconv.FormatBool(c.Default))
	}

	if len(c.Month) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "exp_month")), c.Month)
	}

	if len(c.Year) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "exp_year")), c.Year)
	}

	if len(c.Name) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "name")), c.Name)
	}

	if len(c.Address1) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_line1")), c.Address1)
	}

	if len(c.Address2) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_line2")), c.Address2)
	}

	if len(c.City) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_city")), c.City)
	}

	if len(c.State) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_state")), c.State)
	}

	if len(c.Zip) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_zip")), c.Zip)
	}

	if len(c.Country) > 0 {
		body.Add(form.FormatKey(append(keyParts, cardSource, "address_country")), c.Country)
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
	ID                 string             `json:"id"`
	Month              uint8              `json:"exp_month"`
	Year               uint16             `json:"exp_year"`
	Fingerprint        string             `json:"fingerprint"`
	Funding            CardFunding        `json:"funding"`
	LastFour           string             `json:"last4"`
	Brand              CardBrand          `json:"brand"`
	Currency           Currency           `json:"currency"`
	Default            bool               `json:"default_for_currency"`
	City               string             `json:"address_city"`
	Country            string             `json:"address_country"`
	Address1           string             `json:"address_line1"`
	Address1Check      Verification       `json:"address_line1_check"`
	Address2           string             `json:"address_line2"`
	State              string             `json:"address_state"`
	Zip                string             `json:"address_zip"`
	ZipCheck           Verification       `json:"address_zip_check"`
	CardCountry        string             `json:"country"`
	Customer           *Customer          `json:"customer"`
	CVCCheck           Verification       `json:"cvc_check"`
	Meta               map[string]string  `json:"metadata"`
	Name               string             `json:"name"`
	Recipient          *Recipient         `json:"recipient"`
	DynLastFour        string             `json:"dynamic_last4"`
	Deleted            bool               `json:"deleted"`
	ThreeDSecure       *ThreeDSecure      `json:"three_d_secure"`
	TokenizationMethod TokenizationMethod `json:"tokenization_method"`

	// Description is a succinct summary of the card's information.
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	Description string `json:"description"`

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
}

// CardList is a list object for cards.
type CardList struct {
	ListMeta
	Values []*Card `json:"data"`
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
