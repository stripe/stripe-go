//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of link being generated.
type TerminalOnboardingLinkLinkType string

// List of values that TerminalOnboardingLinkLinkType can take
const (
	TerminalOnboardingLinkLinkTypeAppleTermsAndConditions TerminalOnboardingLinkLinkType = "apple_terms_and_conditions"
)

// The options associated with the Apple Terms and Conditions link type.
type TerminalOnboardingLinkLinkOptionsAppleTermsAndConditionsParams struct {
	// Whether the link should also support users relinking their Apple account.
	AllowRelinking *bool `form:"allow_relinking"`
	// The business name of the merchant accepting Apple's Terms and Conditions.
	MerchantDisplayName *string `form:"merchant_display_name"`
}

// Specific fields needed to generate the desired link type.
type TerminalOnboardingLinkLinkOptionsParams struct {
	// The options associated with the Apple Terms and Conditions link type.
	AppleTermsAndConditions *TerminalOnboardingLinkLinkOptionsAppleTermsAndConditionsParams `form:"apple_terms_and_conditions"`
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
type TerminalOnboardingLinkParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Specific fields needed to generate the desired link type.
	LinkOptions *TerminalOnboardingLinkLinkOptionsParams `form:"link_options"`
	// The type of link being generated.
	LinkType *string `form:"link_type"`
	// Stripe account ID to generate the link for.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// AddExpand appends a new field to expand.
func (p *TerminalOnboardingLinkParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The options associated with the Apple Terms and Conditions link type.
type TerminalOnboardingLinkCreateLinkOptionsAppleTermsAndConditionsParams struct {
	// Whether the link should also support users relinking their Apple account.
	AllowRelinking *bool `form:"allow_relinking"`
	// The business name of the merchant accepting Apple's Terms and Conditions.
	MerchantDisplayName *string `form:"merchant_display_name"`
}

// Specific fields needed to generate the desired link type.
type TerminalOnboardingLinkCreateLinkOptionsParams struct {
	// The options associated with the Apple Terms and Conditions link type.
	AppleTermsAndConditions *TerminalOnboardingLinkCreateLinkOptionsAppleTermsAndConditionsParams `form:"apple_terms_and_conditions"`
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
type TerminalOnboardingLinkCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Specific fields needed to generate the desired link type.
	LinkOptions *TerminalOnboardingLinkCreateLinkOptionsParams `form:"link_options"`
	// The type of link being generated.
	LinkType *string `form:"link_type"`
	// Stripe account ID to generate the link for.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// AddExpand appends a new field to expand.
func (p *TerminalOnboardingLinkCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The options associated with the Apple Terms and Conditions link type.
type TerminalOnboardingLinkLinkOptionsAppleTermsAndConditions struct {
	// Whether the link should also support users relinking their Apple account.
	AllowRelinking bool `json:"allow_relinking"`
	// The business name of the merchant accepting Apple's Terms and Conditions.
	MerchantDisplayName string `json:"merchant_display_name"`
}

// Link type options associated with the current onboarding link object.
type TerminalOnboardingLinkLinkOptions struct {
	// The options associated with the Apple Terms and Conditions link type.
	AppleTermsAndConditions *TerminalOnboardingLinkLinkOptionsAppleTermsAndConditions `json:"apple_terms_and_conditions"`
}

// Returns redirect links used for onboarding onto Tap to Pay on iPhone.
type TerminalOnboardingLink struct {
	APIResource
	// Link type options associated with the current onboarding link object.
	LinkOptions *TerminalOnboardingLinkLinkOptions `json:"link_options"`
	// The type of link being generated.
	LinkType TerminalOnboardingLinkLinkType `json:"link_type"`
	Object   string                         `json:"object"`
	// Stripe account ID to generate the link for.
	OnBehalfOf string `json:"on_behalf_of"`
	// The link passed back to the user for their onboarding.
	RedirectURL string `json:"redirect_url"`
}
