//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of account holder that this account belongs to.
type FinancialConnectionsConsentAccountHolderType string

// List of values that FinancialConnectionsConsentAccountHolderType can take
const (
	FinancialConnectionsConsentAccountHolderTypeAccount  FinancialConnectionsConsentAccountHolderType = "account"
	FinancialConnectionsConsentAccountHolderTypeCustomer FinancialConnectionsConsentAccountHolderType = "customer"
)

// Retrieves the details of a Financial Connections Consent.
type FinancialConnectionsConsentParams struct {
	Params `form:"*"`
	// The account holder for whom the Consent is issued.
	AccountHolder *FinancialConnectionsConsentAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The customer's preferred locale for the consent text, expressed as a BCP 47 language tag. If omitted, Stripe uses the default locale.
	Locale *string `form:"locale" json:"locale,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsConsentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder for whom the Consent is issued.
type FinancialConnectionsConsentAccountHolderParams struct {
	// The ID of the Account for whom the Consent is issued. Required when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// The ID of the Customer for whom the Consent is issued. Required when `type` is `customer` unless `customer_account` is provided.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of an Account representing the Customer for whom the Consent is issued. Required when `type` is `customer` unless `customer` is provided.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// The type of account holder for whom the Consent is issued.
	Type *string `form:"type" json:"type"`
}

// Retrieves the details of a Financial Connections Consent.
type FinancialConnectionsConsentRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsConsentRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder for whom the Consent is issued.
type FinancialConnectionsConsentCreateAccountHolderParams struct {
	// The ID of the Account for whom the Consent is issued. Required when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// The ID of the Customer for whom the Consent is issued. Required when `type` is `customer` unless `customer_account` is provided.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of an Account representing the Customer for whom the Consent is issued. Required when `type` is `customer` unless `customer` is provided.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// The type of account holder for whom the Consent is issued.
	Type *string `form:"type" json:"type"`
}

// Creates a Financial Connections Consent object for an account holder.
type FinancialConnectionsConsentCreateParams struct {
	Params `form:"*"`
	// The account holder for whom the Consent is issued.
	AccountHolder *FinancialConnectionsConsentCreateAccountHolderParams `form:"account_holder" json:"account_holder"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The customer's preferred locale for the consent text, expressed as a BCP 47 language tag. If omitted, Stripe uses the default locale.
	Locale *string `form:"locale" json:"locale,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsConsentCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type FinancialConnectionsConsentAccountHolder struct {
	// The ID of the Stripe account that this account belongs to. Only available when `account_holder.type` is `account`.
	Account *Account `json:"account,omitempty"`
	// The ID for an Account representing a customer that this account belongs to. Only available when `account_holder.type` is `customer`.
	Customer        *Customer `json:"customer,omitempty"`
	CustomerAccount string    `json:"customer_account,omitempty"`
	// Type of account holder that this account belongs to.
	Type FinancialConnectionsConsentAccountHolderType `json:"type"`
}

// Stripe-issued localized Financial Connections consent text.
type FinancialConnectionsConsent struct {
	APIResource
	AccountHolder *FinancialConnectionsConsentAccountHolder `json:"account_holder"`
	// The exact localized text that must be displayed before collecting affirmative consent.
	ConsentText string `json:"consent_text"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The exclusive time after which this Consent can no longer be used as launch evidence.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// The BCP 47 locale used to render `consent_text`.
	Locale string `json:"locale"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
