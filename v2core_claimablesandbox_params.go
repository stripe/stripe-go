//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Values that are prefilled when a user claims the sandbox.
type V2CoreClaimableSandboxPrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// notify this email address before the sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox. If not provided, this will be generated.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
type V2CoreClaimableSandboxParams struct {
	Params `form:"*"`
	// If true, returns a key that can be used with [Stripe's MCP server](https://docs.stripe.com/mcp).
	EnableMcpAccess *bool `form:"enable_mcp_access" json:"enable_mcp_access"`
	// Values that are prefilled when a user claims the sandbox.
	Prefill *V2CoreClaimableSandboxPrefillParams `form:"prefill" json:"prefill"`
}

// Values that are prefilled when a user claims the sandbox.
type V2CoreClaimableSandboxCreatePrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// notify this email address before the sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox. If not provided, this will be generated.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
type V2CoreClaimableSandboxCreateParams struct {
	Params `form:"*"`
	// If true, returns a key that can be used with [Stripe's MCP server](https://docs.stripe.com/mcp).
	EnableMcpAccess *bool `form:"enable_mcp_access" json:"enable_mcp_access"`
	// Values that are prefilled when a user claims the sandbox.
	Prefill *V2CoreClaimableSandboxCreatePrefillParams `form:"prefill" json:"prefill"`
}
