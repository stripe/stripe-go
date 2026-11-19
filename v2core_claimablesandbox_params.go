//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Details about the onboarding link.
type V2CoreClaimableSandboxOnboardingLinkDetailsParams struct {
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL offered to the user as a way back to your app after they claim the sandbox.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxPrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email address of the user who will claim the sandbox.
	// Stripe sends this address an invitation to claim the sandbox and a reminder before the unclaimed sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create a claimable sandbox with prefilled data. The response includes an onboarding link that lets
// a user claim the sandbox into a new or existing Stripe account.
type V2CoreClaimableSandboxParams struct {
	Params `form:"*"`
	// The distribution channel used to pre-install your Stripe App on the claimable sandbox.
	// Use `testing` to install your app's external test version or `public` to install its published version.
	// This field defaults to `public` if omitted.
	Channel *string `form:"channel" json:"channel,omitempty"`
	// If true, returns a key that can be used with [Stripe's MCP server](https://docs.stripe.com/mcp).
	EnableMcpAccess *bool `form:"enable_mcp_access" json:"enable_mcp_access,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Details about the onboarding link.
	OnboardingLinkDetails *V2CoreClaimableSandboxOnboardingLinkDetailsParams `form:"onboarding_link_details" json:"onboarding_link_details,omitempty"`
	// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxPrefillParams `form:"prefill" json:"prefill,omitempty"`
}

// Details about the onboarding link. If omitted, existing details are reused.
// Omitted URL fields retain their existing values.
type V2CoreClaimableSandboxRenewOnboardingLinkOnboardingLinkDetailsParams struct {
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// If omitted, the existing refresh URL is retained.
	// The URL should generate a new onboarding link and redirect the user to it.
	RefreshURL *string `form:"refresh_url" json:"refresh_url,omitempty"`
	// The URL offered to the user as a way back to your app after they claim the sandbox.
	// If omitted, the existing return URL is retained. Set to null to clear it.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Renew the claimable sandbox onboarding link. This invalidates the existing onboarding link.
// The endpoint only works on a claimable sandbox with status `unclaimed` or `claimed`.
type V2CoreClaimableSandboxRenewOnboardingLinkParams struct {
	Params `form:"*"`
	// Details about the onboarding link. If omitted, existing details are reused.
	// Omitted URL fields retain their existing values.
	OnboardingLinkDetails *V2CoreClaimableSandboxRenewOnboardingLinkOnboardingLinkDetailsParams `form:"onboarding_link_details" json:"onboarding_link_details,omitempty"`
}

// Details about the onboarding link.
type V2CoreClaimableSandboxCreateOnboardingLinkDetailsParams struct {
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL offered to the user as a way back to your app after they claim the sandbox.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxCreatePrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email address of the user who will claim the sandbox.
	// Stripe sends this address an invitation to claim the sandbox and a reminder before the unclaimed sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create a claimable sandbox with prefilled data. The response includes an onboarding link that lets
// a user claim the sandbox into a new or existing Stripe account.
type V2CoreClaimableSandboxCreateParams struct {
	Params `form:"*"`
	// The distribution channel used to pre-install your Stripe App on the claimable sandbox.
	// Use `testing` to install your app's external test version or `public` to install its published version.
	// This field defaults to `public` if omitted.
	Channel *string `form:"channel" json:"channel,omitempty"`
	// If true, returns a key that can be used with [Stripe's MCP server](https://docs.stripe.com/mcp).
	EnableMcpAccess *bool `form:"enable_mcp_access" json:"enable_mcp_access"`
	// Details about the onboarding link.
	OnboardingLinkDetails *V2CoreClaimableSandboxCreateOnboardingLinkDetailsParams `form:"onboarding_link_details" json:"onboarding_link_details"`
	// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxCreatePrefillParams `form:"prefill" json:"prefill"`
}

// Retrieves the details of a claimable sandbox that was previously been created.
// Supply the unique claimable sandbox ID that was returned from your creation request,
// and Stripe will return the corresponding sandbox information.
type V2CoreClaimableSandboxRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}
