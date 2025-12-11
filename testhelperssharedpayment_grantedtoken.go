//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Limits on how this SharedPaymentGrantedToken can be used.
type TestHelpersSharedPaymentGrantedTokenUsageLimitsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt *int64 `form:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken
	MaxAmount *int64 `form:"max_amount"`
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// The Customer that the SharedPaymentGrantedToken belongs to. Should match the Customer that the PaymentMethod is attached to if any.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The PaymentMethod that is going to be shared by the SharedPaymentGrantedToken.
	PaymentMethod *string `form:"payment_method"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `form:"shared_metadata"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *TestHelpersSharedPaymentGrantedTokenUsageLimitsParams `form:"usage_limits"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Limits on how this SharedPaymentGrantedToken can be used.
type TestHelpersSharedPaymentGrantedTokenCreateUsageLimitsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt *int64 `form:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken
	MaxAmount *int64 `form:"max_amount"`
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenCreateParams struct {
	Params `form:"*"`
	// The Customer that the SharedPaymentGrantedToken belongs to. Should match the Customer that the PaymentMethod is attached to if any.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The PaymentMethod that is going to be shared by the SharedPaymentGrantedToken.
	PaymentMethod *string `form:"payment_method"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `form:"shared_metadata"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *TestHelpersSharedPaymentGrantedTokenCreateUsageLimitsParams `form:"usage_limits"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Revokes a test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to revoke SharedPaymentGrantedTokens for testing their integration
type TestHelpersSharedPaymentGrantedTokenUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersSharedPaymentGrantedTokenUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
