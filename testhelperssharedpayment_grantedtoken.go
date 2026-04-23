//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Limits on how this SharedPaymentGrantedToken can be used.
type TestHelpersSharedPaymentGrantedTokenUsageLimitsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at,omitempty"`
	// Max amount that can be captured using this SharedPaymentToken
	MaxAmount *int64 `form:"max_amount" json:"max_amount"`
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// The Customer that the SharedPaymentGrantedToken belongs to. Should match the Customer that the PaymentMethod is attached to if any.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The PaymentMethod that is going to be shared by the SharedPaymentGrantedToken.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *TestHelpersSharedPaymentGrantedTokenUsageLimitsParams `form:"usage_limits" json:"usage_limits"`
	UnsetFields []TestHelpersSharedPaymentGrantedTokenParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersSharedPaymentGrantedTokenParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersSharedPaymentGrantedTokenParams.
type TestHelpersSharedPaymentGrantedTokenParamsUnsetField string

const (
	TestHelpersSharedPaymentGrantedTokenParamsUnsetFieldSharedMetadata TestHelpersSharedPaymentGrantedTokenParamsUnsetField = "shared_metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersSharedPaymentGrantedTokenParams) AddUnsetField(field TestHelpersSharedPaymentGrantedTokenParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Revokes a test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to revoke SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenRevokeParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenRevokeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Limits on how this SharedPaymentGrantedToken can be used.
type TestHelpersSharedPaymentGrantedTokenCreateUsageLimitsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at,omitempty"`
	// Max amount that can be captured using this SharedPaymentToken
	MaxAmount *int64 `form:"max_amount" json:"max_amount"`
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenCreateParams struct {
	Params `form:"*"`
	// The Customer that the SharedPaymentGrantedToken belongs to. Should match the Customer that the PaymentMethod is attached to if any.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The PaymentMethod that is going to be shared by the SharedPaymentGrantedToken.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *TestHelpersSharedPaymentGrantedTokenCreateUsageLimitsParams `form:"usage_limits" json:"usage_limits"`
	UnsetFields []TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersSharedPaymentGrantedTokenCreateParams.
type TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetField string

const (
	TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetFieldSharedMetadata TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetField = "shared_metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersSharedPaymentGrantedTokenCreateParams) AddUnsetField(field TestHelpersSharedPaymentGrantedTokenCreateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
