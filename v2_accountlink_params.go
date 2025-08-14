//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Indicates that the AccountLink provided should onboard an account.
type V2AccountLinkUseCaseAccountOnboardingParams struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2AccountLinkUseCaseAccountUpdateParams struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// The use case of the AccountLink.
type V2AccountLinkUseCaseParams struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2AccountLinkUseCaseAccountOnboardingParams `form:"account_onboarding" json:"account_onboarding,omitempty"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2AccountLinkUseCaseAccountUpdateParams `form:"account_update" json:"account_update,omitempty"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type *string `form:"type" json:"type"`
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
type V2AccountLinkParams struct {
	Params `form:"*"`
	// The ID of the Account to create link for.
	Account *string `form:"account" json:"account"`
	// The use case of the AccountLink.
	UseCase *V2AccountLinkUseCaseParams `form:"use_case" json:"use_case"`
}

// Indicates that the AccountLink provided should onboard an account.
type V2AccountLinkCreateUseCaseAccountOnboardingParams struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2AccountLinkCreateUseCaseAccountUpdateParams struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// The use case of the AccountLink.
type V2AccountLinkCreateUseCaseParams struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2AccountLinkCreateUseCaseAccountOnboardingParams `form:"account_onboarding" json:"account_onboarding,omitempty"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2AccountLinkCreateUseCaseAccountUpdateParams `form:"account_update" json:"account_update,omitempty"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type *string `form:"type" json:"type"`
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
type V2AccountLinkCreateParams struct {
	Params `form:"*"`
	// The ID of the Account to create link for.
	Account *string `form:"account" json:"account"`
	// The use case of the AccountLink.
	UseCase *V2AccountLinkCreateUseCaseParams `form:"use_case" json:"use_case"`
}
