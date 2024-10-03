//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of this amount. We currently only support `monetary` credits.
type BillingCreditGrantAmountType string

// List of values that BillingCreditGrantAmountType can take
const (
	BillingCreditGrantAmountTypeMonetary BillingCreditGrantAmountType = "monetary"
)

// The price type to which credit grants can apply to. We currently only support `metered` price type.
type BillingCreditGrantApplicabilityConfigScopePriceType string

// List of values that BillingCreditGrantApplicabilityConfigScopePriceType can take
const (
	BillingCreditGrantApplicabilityConfigScopePriceTypeMetered BillingCreditGrantApplicabilityConfigScopePriceType = "metered"
)

// The category of this credit grant.
type BillingCreditGrantCategory string

// List of values that BillingCreditGrantCategory can take
const (
	BillingCreditGrantCategoryPaid        BillingCreditGrantCategory = "paid"
	BillingCreditGrantCategoryPromotional BillingCreditGrantCategory = "promotional"
)

// Retrieve a list of credit grants
type BillingCreditGrantListParams struct {
	ListParams `form:"*"`
	// Only return credit grants for this customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The monetary amount.
type BillingCreditGrantAmountMonetaryParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency"`
	// A positive integer representing the amount of the credit grant.
	Value *int64 `form:"value"`
}

// Amount of this credit grant.
type BillingCreditGrantAmountParams struct {
	// The monetary amount.
	Monetary *BillingCreditGrantAmountMonetaryParams `form:"monetary"`
	// Specify the type of this amount. We currently only support `monetary` credits.
	Type *string `form:"type"`
}

// Specify the scope of this applicability config.
type BillingCreditGrantApplicabilityConfigScopeParams struct {
	// The price type to which credit grants can apply to. We currently only support `metered` price type.
	PriceType *string `form:"price_type"`
}

// Configuration specifying what this credit grant applies to.
type BillingCreditGrantApplicabilityConfigParams struct {
	// Specify the scope of this applicability config.
	Scope *BillingCreditGrantApplicabilityConfigScopeParams `form:"scope"`
}

// Creates a credit grant
type BillingCreditGrantParams struct {
	Params `form:"*"`
	// Amount of this credit grant.
	Amount *BillingCreditGrantAmountParams `form:"amount"`
	// Configuration specifying what this credit grant applies to.
	ApplicabilityConfig *BillingCreditGrantApplicabilityConfigParams `form:"applicability_config"`
	// The category of this credit grant.
	Category *string `form:"category"`
	// Id of the customer to whom the credit should be granted.
	Customer *string `form:"customer"`
	// The time when the credit becomes effective i.e when it is eligible to be used. Defaults to the current timestamp if not specified.
	EffectiveAt *int64 `form:"effective_at"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The time when the credit created by this credit grant will expire. If set to empty, the credit will never expire.
	ExpiresAt *int64 `form:"expires_at"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object (ex: cost basis) in a structured format.
	Metadata map[string]string `form:"metadata"`
	// A descriptive name shown in dashboard and on invoices.
	Name *string `form:"name"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BillingCreditGrantParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Expires a credit grant
type BillingCreditGrantExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Voids a credit grant
type BillingCreditGrantVoidGrantParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantVoidGrantParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The monetary amount.
type BillingCreditGrantAmountMonetary struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}
type BillingCreditGrantAmount struct {
	// The monetary amount.
	Monetary *BillingCreditGrantAmountMonetary `json:"monetary"`
	// The type of this amount. We currently only support `monetary` credits.
	Type BillingCreditGrantAmountType `json:"type"`
}
type BillingCreditGrantApplicabilityConfigScope struct {
	// The price type to which credit grants can apply to. We currently only support `metered` price type.
	PriceType BillingCreditGrantApplicabilityConfigScopePriceType `json:"price_type"`
}
type BillingCreditGrantApplicabilityConfig struct {
	Scope *BillingCreditGrantApplicabilityConfigScope `json:"scope"`
}

// A credit grant is a resource that records a grant of some credit to a customer.
type BillingCreditGrant struct {
	APIResource
	Amount              *BillingCreditGrantAmount              `json:"amount"`
	ApplicabilityConfig *BillingCreditGrantApplicabilityConfig `json:"applicability_config"`
	// The category of this credit grant.
	Category BillingCreditGrantCategory `json:"category"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Id of the customer to whom the credit was granted.
	Customer *Customer `json:"customer"`
	// The time when the credit becomes effective i.e when it is eligible to be used.
	EffectiveAt int64 `json:"effective_at"`
	// The time when the credit will expire. If not present, the credit will never expire.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A descriptive name shown in dashboard and on invoices.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the test clock this credit grant belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
	// The time when this credit grant was voided. If not present, the credit grant hasn't been voided.
	VoidedAt int64 `json:"voided_at"`
}

// BillingCreditGrantList is a list of CreditGrants as retrieved from a list endpoint.
type BillingCreditGrantList struct {
	APIResource
	ListMeta
	Data []*BillingCreditGrant `json:"data"`
}

// UnmarshalJSON handles deserialization of a BillingCreditGrant.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingCreditGrant) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingCreditGrant BillingCreditGrant
	var v billingCreditGrant
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingCreditGrant(v)
	return nil
}
