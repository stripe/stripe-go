//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
	"strconv"
)

// If `address_line1` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`.
type CardVerification string

// List of values that CardVerification can take
const (
	CardVerificationFail        CardVerification = "fail"
	CardVerificationPass        CardVerification = "pass"
	CardVerificationUnavailable CardVerification = "unavailable"
	CardVerificationUnchecked   CardVerification = "unchecked"
)

// A set of available payout methods for this card. Only values from this set should be passed as the `method` when creating a payout.
type CardAvailablePayoutMethod string

// List of values that CardAvailablePayoutMethod can take
const (
	CardAvailablePayoutMethodInstant  CardAvailablePayoutMethod = "instant"
	CardAvailablePayoutMethodStandard CardAvailablePayoutMethod = "standard"
)

// Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`, `Visa`, or `Unknown`.
type CardBrand string

// List of values that CardBrand can take
const (
	CardBrandAmex       CardBrand = "American Express"
	CardBrandDiscover   CardBrand = "Discover"
	CardBrandDinersClub CardBrand = "Diners Club"
	CardBrandJCB        CardBrand = "JCB"
	CardBrandMasterCard CardBrand = "MasterCard"
	CardBrandUnknown    CardBrand = "Unknown"
	CardBrandUnionPay   CardBrand = "UnionPay"
	CardBrandVisa       CardBrand = "Visa"
)

// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
type CardFunding string

// List of values that CardFunding can take
const (
	CardFundingCredit  CardFunding = "credit"
	CardFundingDebit   CardFunding = "debit"
	CardFundingPrepaid CardFunding = "prepaid"
	CardFundingUnknown CardFunding = "unknown"
)

// If the card number is tokenized, this is the method that was used. Can be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or null.
type CardTokenizationMethod string

// List of values that CardTokenizationMethod can take
const (
	TokenizationMethodAndroidPay CardTokenizationMethod = "android_pay"
	TokenizationMethodApplePay   CardTokenizationMethod = "apple_pay"
)

type CardOwnerParams struct {
	// Owner's address.
	Address *AddressParams `form:"address"`
	// Owner's email address.
	Email *string `form:"email"`
	// Owner's full name.
	Name *string `form:"name"`
	// Owner's phone number.
	Phone *string `form:"phone"`
}

// cardSource is a string that's used to build card form parameters. It's a
// constant just to make mistakes less likely.
const cardSource = "source"

// Update a specified source for a given customer.
type CardParams struct {
	Params   `form:"*"`
	Account  *string `form:"-"` // Included in URL
	Token    *string `form:"-"` // Included in URL
	Customer *string `form:"-"` // Included in URL
	// The name of the person or business that owns the bank account.
	AccountHolderName *string `form:"account_holder_name"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType *string `form:"account_holder_type"`
	AccountType       *string `form:"account_type"`
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
	AddressZip         *string `form:"address_zip"`
	Currency           *string `form:"currency"`
	CVC                *string `form:"cvc"`
	DefaultForCurrency *bool   `form:"default_for_currency"`
	// Two digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month"`
	// Four digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year"`
	// Cardholder name.
	Name   *string          `form:"name"`
	Number *string          `form:"number"`
	Owner  *CardOwnerParams `form:"owner"`
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
	// Rather than being called in addition to `AppendTo`, this function
	// *replaces* `AppendTo`, so we must also make sure to handle the encoding
	// of `Params` so metadata and the like is included in the encoded payload.
	form.AppendToPrefixed(body, c.Params, keyParts)

	if c.DefaultForCurrency != nil {
		body.Add(
			form.FormatKey(append(keyParts, "default_for_currency")),
			strconv.FormatBool(BoolValue(c.DefaultForCurrency)),
		)
	}

	if c.Token != nil {
		if c.Account != nil {
			body.Add(form.FormatKey(append(keyParts, "external_account")), StringValue(c.Token))
		} else {
			body.Add(form.FormatKey(append(keyParts, cardSource)), StringValue(c.Token))
		}
	}

	if c.Number != nil {
		body.Add(form.FormatKey(append(keyParts, cardSource, "object")), "card")
		body.Add(form.FormatKey(append(keyParts, cardSource, "number")), StringValue(c.Number))
	}
	if c.CVC != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "cvc")),
			StringValue(c.CVC),
		)
	}
	if c.Currency != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "currency")),
			StringValue(c.Currency),
		)
	}
	if c.ExpMonth != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "exp_month")),
			StringValue(c.ExpMonth),
		)
	}
	if c.ExpYear != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "exp_year")),
			StringValue(c.ExpYear),
		)
	}
	if c.Name != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "name")),
			StringValue(c.Name),
		)
	}
	if c.AddressCity != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_city")),
			StringValue(c.AddressCity),
		)
	}
	if c.AddressCountry != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_country")),
			StringValue(c.AddressCountry),
		)
	}
	if c.AddressLine1 != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_line1")),
			StringValue(c.AddressLine1),
		)
	}
	if c.AddressLine2 != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_line2")),
			StringValue(c.AddressLine2),
		)
	}
	if c.AddressState != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_state")),
			StringValue(c.AddressState),
		)
	}
	if c.AddressZip != nil {
		body.Add(
			form.FormatKey(append(keyParts, cardSource, "address_zip")),
			StringValue(c.AddressZip),
		)
	}
}

type CardListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
	Customer   *string `form:"-"` // Included in URL
}

// AppendTo implements custom encoding logic for CardListParams
// so that we can send the special required `object` field up along with the
// other specified parameters.
func (p *CardListParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Account != nil || p.Customer != nil {
		body.Add(form.FormatKey(append(keyParts, "object")), "card")
	}
}

// You can store multiple cards on a customer in order to charge the customer
// later. You can also store multiple debit cards on a recipient in order to
// transfer to those cards later.
//
// Related guide: [Card Payments with Sources](https://stripe.com/docs/sources/cards).
type Card struct {
	APIResource
	// The account this card belongs to. This attribute will not be in the card object if the card belongs to a customer or recipient instead.
	Account *Account `json:"account"`
	// City/District/Suburb/Town/Village.
	AddressCity string `json:"address_city"`
	// Billing address country, if provided when creating card.
	AddressCountry string `json:"address_country"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 string `json:"address_line1"`
	// If `address_line1` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check CardVerification `json:"address_line1_check"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 string `json:"address_line2"`
	// State/County/Province/Region.
	AddressState string `json:"address_state"`
	// ZIP or postal code.
	AddressZip string `json:"address_zip"`
	// If `address_zip` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressZipCheck CardVerification `json:"address_zip_check"`
	// A set of available payout methods for this card. Only values from this set should be passed as the `method` when creating a payout.
	AvailablePayoutMethods []CardAvailablePayoutMethod `json:"available_payout_methods"`
	// Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`, `Visa`, or `Unknown`.
	Brand CardBrand `json:"brand"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Three-letter [ISO code for currency](https://stripe.com/docs/payouts). Only applicable on accounts (not customers or recipients). The card can be used as a transfer destination for funds in this currency.
	Currency Currency `json:"currency"`
	// The customer that this card belongs to. This attribute will not be in the card object if the card belongs to an account or recipient instead.
	Customer *Customer `json:"customer"`
	// If a CVC was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A result of unchecked indicates that CVC was provided but hasn't been checked yet. Checks are typically performed when attaching a card to a Customer object, or when creating a charge. For more details, see [Check if a card is valid without a charge](https://support.stripe.com/questions/check-if-a-card-is-valid-without-a-charge).
	CVCCheck CardVerification `json:"cvc_check"`
	// Whether this card is the default external account for its currency.
	DefaultForCurrency bool `json:"default_for_currency"`
	Deleted            bool `json:"deleted"`
	// Description is a succinct summary of the card's information.
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string `json:"dynamic_last4"`
	// Two-digit number representing the card's expiration month.
	ExpMonth uint8 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear uint16 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding CardFunding `json:"funding"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// IIN is the card's "Issuer Identification Number".
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// Issuer is a bank or financial institution that provides the card.
	//
	// Please note that this field is for internal use only and is not returned
	// as part of standard API requests.
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Cardholder name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// For external accounts, possible values are `new` and `errored`. If a transfer fails, the status is set to `errored` and transfers are stopped until account details are updated.
	Status string `json:"status"`
	// If the card number is tokenized, this is the method that was used. Can be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or null.
	TokenizationMethod CardTokenizationMethod `json:"tokenization_method"`
}

// CardList is a list of Cards as retrieved from a list endpoint.
type CardList struct {
	APIResource
	ListMeta
	Data []*Card `json:"data"`
}

// UnmarshalJSON handles deserialization of a Card.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Card) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type card Card
	var v card
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Card(v)
	return nil
}
