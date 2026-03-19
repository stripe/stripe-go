//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of promotion.
type PromotionCodePromotionType string

// List of values that PromotionCodePromotionType can take
const (
	PromotionCodePromotionTypeCoupon PromotionCodePromotionType = "coupon"
)

// Returns a list of your promotion codes.
type PromotionCodeListParams struct {
	ListParams `form:"*"`
	// Filter promotion codes by whether they are active.
	Active *bool `form:"active"`
	// Only return promotion codes that have this case-insensitive code.
	Code *string `form:"code"`
	// Only return promotion codes for this coupon.
	Coupon *string `form:"coupon"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return promotion codes that are restricted to this customer.
	Customer *string `form:"customer"`
	// Only return promotion codes that are restricted to this account representing the customer.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PromotionCodeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The promotion referenced by this promotion code.
type PromotionCodePromotionParams struct {
	// If promotion `type` is `coupon`, the coupon for this promotion code.
	Coupon *string `form:"coupon"`
	// Specifies the type of promotion.
	Type *string `form:"type"`
}

// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PromotionCodeRestrictionsCurrencyOptionsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Settings that restrict the redemption of the promotion code.
type PromotionCodeRestrictionsParams struct {
	// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*PromotionCodeRestrictionsCurrencyOptionsParams `form:"currency_options"`
	// A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices
	FirstTimeTransaction *bool `form:"first_time_transaction"`
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
	// Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount
	MinimumAmountCurrency *string `form:"minimum_amount_currency"`
}

// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
type PromotionCodeParams struct {
	Params `form:"*"`
	// Whether the promotion code is currently active. A promotion code can only be reactivated when the coupon is still valid and the promotion code is otherwise redeemable.
	Active *bool `form:"active"`
	// The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).
	//
	// If left blank, we will generate one automatically.
	Code *string `form:"code"`
	// The customer who can use this promotion code. If not set, all customers can use the promotion code.
	Customer *string `form:"customer"`
	// The account representing the customer who can use this promotion code. If not set, all customers can use the promotion code.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.
	ExpiresAt *int64 `form:"expires_at"`
	// A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`.
	MaxRedemptions *int64 `form:"max_redemptions"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The promotion referenced by this promotion code.
	Promotion *PromotionCodePromotionParams `form:"promotion"`
	// Settings that restrict the redemption of the promotion code.
	Restrictions *PromotionCodeRestrictionsParams `form:"restrictions"`
}

// AddExpand appends a new field to expand.
func (p *PromotionCodeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PromotionCodeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The promotion referenced by this promotion code.
type PromotionCodeCreatePromotionParams struct {
	// If promotion `type` is `coupon`, the coupon for this promotion code.
	Coupon *string `form:"coupon"`
	// Specifies the type of promotion.
	Type *string `form:"type"`
}

// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PromotionCodeCreateRestrictionsCurrencyOptionsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Settings that restrict the redemption of the promotion code.
type PromotionCodeCreateRestrictionsParams struct {
	// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*PromotionCodeCreateRestrictionsCurrencyOptionsParams `form:"currency_options"`
	// A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices
	FirstTimeTransaction *bool `form:"first_time_transaction"`
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
	// Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount
	MinimumAmountCurrency *string `form:"minimum_amount_currency"`
}

// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
type PromotionCodeCreateParams struct {
	Params `form:"*"`
	// Whether the promotion code is currently active.
	Active *bool `form:"active"`
	// The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).
	//
	// If left blank, we will generate one automatically.
	Code *string `form:"code"`
	// The customer who can use this promotion code. If not set, all customers can use the promotion code.
	Customer *string `form:"customer"`
	// The account representing the customer who can use this promotion code. If not set, all customers can use the promotion code.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.
	ExpiresAt *int64 `form:"expires_at"`
	// A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`.
	MaxRedemptions *int64 `form:"max_redemptions"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The promotion referenced by this promotion code.
	Promotion *PromotionCodeCreatePromotionParams `form:"promotion"`
	// Settings that restrict the redemption of the promotion code.
	Restrictions *PromotionCodeCreateRestrictionsParams `form:"restrictions"`
}

// AddExpand appends a new field to expand.
func (p *PromotionCodeCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PromotionCodeCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://docs.stripe.com/docs/api/promotion_codes/list) with the desired code.
type PromotionCodeRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PromotionCodeRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PromotionCodeUpdateRestrictionsCurrencyOptionsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Settings that restrict the redemption of the promotion code.
type PromotionCodeUpdateRestrictionsParams struct {
	// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*PromotionCodeUpdateRestrictionsCurrencyOptionsParams `form:"currency_options"`
}

// Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.
type PromotionCodeUpdateParams struct {
	Params `form:"*"`
	// Whether the promotion code is currently active. A promotion code can only be reactivated when the coupon is still valid and the promotion code is otherwise redeemable.
	Active *bool `form:"active"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Settings that restrict the redemption of the promotion code.
	Restrictions *PromotionCodeUpdateRestrictionsParams `form:"restrictions"`
}

// AddExpand appends a new field to expand.
func (p *PromotionCodeUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PromotionCodeUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type PromotionCodePromotion struct {
	// If promotion `type` is `coupon`, the coupon for this promotion.
	Coupon *Coupon `json:"coupon"`
	// The type of promotion.
	Type PromotionCodePromotionType `json:"type"`
}

// Promotion code restrictions defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PromotionCodeRestrictionsCurrencyOptions struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictions struct {
	// Promotion code restrictions defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*PromotionCodeRestrictionsCurrencyOptions `json:"currency_options"`
	// A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices
	FirstTimeTransaction bool `json:"first_time_transaction"`
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
	// Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount
	MinimumAmountCurrency Currency `json:"minimum_amount_currency"`
}

// A Promotion Code represents a customer-redeemable code for an underlying promotion.
// You can create multiple codes for a single promotion.
//
// If you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.
// Customers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.
type PromotionCode struct {
	APIResource
	// Whether the promotion code is currently active. A promotion code is only active if the coupon is also valid.
	Active bool `json:"active"`
	// The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for each customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).
	Code string `json:"code"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The customer who can use this promotion code.
	Customer *Customer `json:"customer"`
	// The account representing the customer who can use this promotion code.
	CustomerAccount string `json:"customer_account"`
	// Date at which the promotion code can no longer be redeemed.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Maximum number of times this promotion code can be redeemed.
	MaxRedemptions int64 `json:"max_redemptions"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object       string                     `json:"object"`
	Promotion    *PromotionCodePromotion    `json:"promotion"`
	Restrictions *PromotionCodeRestrictions `json:"restrictions"`
	// Number of times this promotion code has been used.
	TimesRedeemed int64 `json:"times_redeemed"`
}

// PromotionCodeList is a list of PromotionCodes as retrieved from a list endpoint.
type PromotionCodeList struct {
	APIResource
	ListMeta
	Data []*PromotionCode `json:"data"`
}

// UnmarshalJSON handles deserialization of a PromotionCode.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PromotionCode) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type promotionCode PromotionCode
	var v promotionCode
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PromotionCode(v)
	return nil
}
