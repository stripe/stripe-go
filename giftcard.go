//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The brand of the gift card.
type GiftCardBrand string

// List of values that GiftCardBrand can take
const (
	GiftCardBrandFiservValuelink GiftCardBrand = "fiserv_valuelink"
	GiftCardBrandGivex           GiftCardBrand = "givex"
	GiftCardBrandSvs             GiftCardBrand = "svs"
)

// Retrieves a third-party gift card object.
type GiftCardParams struct {
	Params `form:"*"`
	// The brand of the gift card.
	Brand *string `form:"brand" json:"brand,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Two-digit number representing the gift card's expiration month.
	ExpMonth *int64 `form:"exp_month" json:"exp_month,omitempty"`
	// Four-digit number representing the gift card's expiration year.
	ExpYear *int64 `form:"exp_year" json:"exp_year,omitempty"`
	// The gift card number.
	Number *string `form:"number" json:"number,omitempty"`
	// The gift card PIN.
	PIN *string `form:"pin" json:"pin,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The initial balance to set on the gift card.
type GiftCardActivateBalanceParams struct {
	// The initial balance amount to be loaded when activating the gift card, in the smallest currency unit
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
}

// Activates a third-party gift card and optionally sets its balance.
type GiftCardActivateParams struct {
	Params `form:"*"`
	// The initial balance to set on the gift card.
	Balance *GiftCardActivateBalanceParams `form:"balance" json:"balance,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The Stripe account ID to process the gift card operation on behalf of.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardActivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cashout a third-party gift card by zeroing its balance.
type GiftCardCashoutParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The Stripe account ID to process the gift card operation on behalf of.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardCashoutParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Checks the balance of a third-party gift card.
type GiftCardCheckBalanceParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The Stripe account ID to process the gift card operation on behalf of.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardCheckBalanceParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reloads a third-party gift card by adding the specified amount to its balance.
type GiftCardReloadParams struct {
	Params `form:"*"`
	// The amount to add to the gift card balance, in the smallest currency unit.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The Stripe account ID to process the gift card operation on behalf of.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardReloadParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Voids a previously performed gift card operation.
type GiftCardVoidOperationParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The Stripe account ID to process the gift card operation on behalf of.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// The ID of the gift card operation to void.
	Operation *string `form:"operation" json:"operation"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardVoidOperationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a third-party gift card object.
type GiftCardRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a gift card object.
type GiftCardCreateParams struct {
	Params `form:"*"`
	// The brand of the gift card.
	Brand *string `form:"brand" json:"brand"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Two-digit number representing the gift card's expiration month.
	ExpMonth *int64 `form:"exp_month" json:"exp_month,omitempty"`
	// Four-digit number representing the gift card's expiration year.
	ExpYear *int64 `form:"exp_year" json:"exp_year,omitempty"`
	// The gift card number.
	Number *string `form:"number" json:"number,omitempty"`
	// The gift card PIN.
	PIN *string `form:"pin" json:"pin,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Represents third-party gift cards that can be used as a payment method through Stripe.
type GiftCard struct {
	APIResource
	// The brand of the gift card.
	Brand GiftCardBrand `json:"brand"`
	// The expiration month of the gift card.
	ExpMonth int64 `json:"exp_month"`
	// The expiration year of the gift card.
	ExpYear int64 `json:"exp_year"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The last four digits of the gift card number.
	Last4 string `json:"last4"`
	// The last operation performed on this gift card.
	LastOperation *GiftCardOperation `json:"last_operation"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// UnmarshalJSON handles deserialization of a GiftCard.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (g *GiftCard) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		g.ID = id
		return nil
	}

	type giftCard GiftCard
	var v giftCard
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*g = GiftCard(v)
	return nil
}
