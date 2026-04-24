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
}

// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxPrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// notify this email address before the sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
type V2CoreClaimableSandboxParams struct {
	Params `form:"*"`
	// The app channel that will be used when pre-installing your app on the claimable sandbox.
	// This field defaults to `public` if omitted.
	AppChannel *string `form:"app_channel" json:"app_channel,omitempty"`
	// If true, returns a key that can be used with [Stripe's MCP server](https://docs.stripe.com/mcp).
	EnableMcpAccess *bool `form:"enable_mcp_access" json:"enable_mcp_access,omitempty"`
	// Details about the onboarding link.
	OnboardingLinkDetails *V2CoreClaimableSandboxOnboardingLinkDetailsParams `form:"onboarding_link_details" json:"onboarding_link_details,omitempty"`
	// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxPrefillParams `form:"prefill" json:"prefill,omitempty"`
}

// Details about the onboarding link.
// If omitted, the existing onboarding link details will be reused.
type V2CoreClaimableSandboxRenewOnboardingLinkOnboardingLinkDetailsParams struct {
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
}

// Renew the claimable sandbox onboarding link. This will invalidate any existing onboarding links.
// The endpoint only works on a claimable sandbox with status `unclaimed` or `claimed`.
type V2CoreClaimableSandboxRenewOnboardingLinkParams struct {
	Params `form:"*"`
	// Details about the onboarding link.
	// If omitted, the existing onboarding link details will be reused.
	OnboardingLinkDetails *V2CoreClaimableSandboxRenewOnboardingLinkOnboardingLinkDetailsParams `form:"onboarding_link_details" json:"onboarding_link_details,omitempty"`
}

// Details about the onboarding link.
type V2CoreClaimableSandboxCreateOnboardingLinkDetailsParams struct {
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
}

// Values that are prefilled when a user claims the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxCreatePrefillParams struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// notify this email address before the sandbox expires.
	Email *string `form:"email" json:"email"`
	// Name for the sandbox.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
type V2CoreClaimableSandboxCreateParams struct {
	Params `form:"*"`
	// The app channel that will be used when pre-installing your app on the claimable sandbox.
	// This field defaults to `public` if omitted.
	AppChannel *string `form:"app_channel" json:"app_channel,omitempty"`
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
}
