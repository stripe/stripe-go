//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsParams struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields *string `form:"fields" json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements *string `form:"future_requirements" json:"future_requirements,omitempty"`
}

// Indicates that the AccountLink provided should onboard an account.
type V2CoreAccountLinkUseCaseAccountOnboardingParams struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsParams `form:"collection_options" json:"collection_options,omitempty"`
	// Open Enum. A v2/core/account can be configured to enable certain functionality. The configuration param targets the v2/core/account_link to collect information for the specified v2/core/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsParams struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields *string `form:"fields" json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements *string `form:"future_requirements" json:"future_requirements,omitempty"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2CoreAccountLinkUseCaseAccountUpdateParams struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsParams `form:"collection_options" json:"collection_options,omitempty"`
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// The use case of the AccountLink.
type V2CoreAccountLinkUseCaseParams struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2CoreAccountLinkUseCaseAccountOnboardingParams `form:"account_onboarding" json:"account_onboarding,omitempty"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2CoreAccountLinkUseCaseAccountUpdateParams `form:"account_update" json:"account_update,omitempty"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type *string `form:"type" json:"type"`
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
type V2CoreAccountLinkParams struct {
	Params `form:"*"`
	// The ID of the Account to create link for.
	Account *string `form:"account" json:"account"`
	// The use case of the AccountLink.
	UseCase *V2CoreAccountLinkUseCaseParams `form:"use_case" json:"use_case"`
}

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkCreateUseCaseAccountOnboardingCollectionOptionsParams struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields *string `form:"fields" json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements *string `form:"future_requirements" json:"future_requirements,omitempty"`
}

// Indicates that the AccountLink provided should onboard an account.
type V2CoreAccountLinkCreateUseCaseAccountOnboardingParams struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkCreateUseCaseAccountOnboardingCollectionOptionsParams `form:"collection_options" json:"collection_options,omitempty"`
	// Open Enum. A v2/core/account can be configured to enable certain functionality. The configuration param targets the v2/core/account_link to collect information for the specified v2/core/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkCreateUseCaseAccountUpdateCollectionOptionsParams struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields *string `form:"fields" json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements *string `form:"future_requirements" json:"future_requirements,omitempty"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2CoreAccountLinkCreateUseCaseAccountUpdateParams struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkCreateUseCaseAccountUpdateCollectionOptionsParams `form:"collection_options" json:"collection_options,omitempty"`
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []*string `form:"configurations" json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL *string `form:"refresh_url" json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// The use case of the AccountLink.
type V2CoreAccountLinkCreateUseCaseParams struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2CoreAccountLinkCreateUseCaseAccountOnboardingParams `form:"account_onboarding" json:"account_onboarding,omitempty"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2CoreAccountLinkCreateUseCaseAccountUpdateParams `form:"account_update" json:"account_update,omitempty"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type *string `form:"type" json:"type"`
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
type V2CoreAccountLinkCreateParams struct {
	Params `form:"*"`
	// The ID of the Account to create link for.
	Account *string `form:"account" json:"account"`
	// The use case of the AccountLink.
	UseCase *V2CoreAccountLinkCreateUseCaseParams `form:"use_case" json:"use_case"`
}
