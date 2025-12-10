//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of Account Link the user is requesting.
type V2CoreAccountLinkUseCaseType string

// List of values that V2CoreAccountLinkUseCaseType can take
const (
	V2CoreAccountLinkUseCaseTypeAccountOnboarding V2CoreAccountLinkUseCaseType = "account_onboarding"
	V2CoreAccountLinkUseCaseTypeAccountUpdate     V2CoreAccountLinkUseCaseType = "account_update"
)

// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
type V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields string

// List of values that V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields can take
const (
	V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFieldsCurrentlyDue  V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields = "currently_due"
	V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFieldsEventuallyDue V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields = "eventually_due"
)

// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
type V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements string

// List of values that V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements can take
const (
	V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirementsInclude V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements = "include"
	V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirementsOmit    V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements = "omit"
)

// Open Enum. A v2/core/account can be configured to enable certain functionality. The configuration param targets the v2/core/account_link to collect information for the specified v2/core/account configuration/s.
type V2CoreAccountLinkUseCaseAccountOnboardingConfiguration string

// List of values that V2CoreAccountLinkUseCaseAccountOnboardingConfiguration can take
const (
	V2CoreAccountLinkUseCaseAccountOnboardingConfigurationCustomer  V2CoreAccountLinkUseCaseAccountOnboardingConfiguration = "customer"
	V2CoreAccountLinkUseCaseAccountOnboardingConfigurationMerchant  V2CoreAccountLinkUseCaseAccountOnboardingConfiguration = "merchant"
	V2CoreAccountLinkUseCaseAccountOnboardingConfigurationRecipient V2CoreAccountLinkUseCaseAccountOnboardingConfiguration = "recipient"
)

// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). The default value is `currently_due`.
type V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields string

// List of values that V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields can take
const (
	V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFieldsCurrentlyDue  V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields = "currently_due"
	V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFieldsEventuallyDue V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields = "eventually_due"
)

// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
type V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements string

// List of values that V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements can take
const (
	V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirementsInclude V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements = "include"
	V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirementsOmit    V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements = "omit"
)

// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
type V2CoreAccountLinkUseCaseAccountUpdateConfiguration string

// List of values that V2CoreAccountLinkUseCaseAccountUpdateConfiguration can take
const (
	V2CoreAccountLinkUseCaseAccountUpdateConfigurationCustomer  V2CoreAccountLinkUseCaseAccountUpdateConfiguration = "customer"
	V2CoreAccountLinkUseCaseAccountUpdateConfigurationMerchant  V2CoreAccountLinkUseCaseAccountUpdateConfiguration = "merchant"
	V2CoreAccountLinkUseCaseAccountUpdateConfigurationRecipient V2CoreAccountLinkUseCaseAccountUpdateConfiguration = "recipient"
)

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptions struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields `json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements `json:"future_requirements,omitempty"`
}

// Hash containing configuration options for an Account Link object that onboards a new account.
type V2CoreAccountLinkUseCaseAccountOnboarding struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptions `json:"collection_options,omitempty"`
	// Open Enum. A v2/core/account can be configured to enable certain functionality. The configuration param targets the v2/core/account_link to collect information for the specified v2/core/account configuration/s.
	Configurations []V2CoreAccountLinkUseCaseAccountOnboardingConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url,omitempty"`
}

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountUpdateCollectionOptions struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). The default value is `currently_due`.
	Fields V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields `json:"fields,omitempty"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements `json:"future_requirements,omitempty"`
}

// Hash containing configuration options for an Account Link that updates an existing account.
type V2CoreAccountLinkUseCaseAccountUpdate struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountUpdateCollectionOptions `json:"collection_options,omitempty"`
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []V2CoreAccountLinkUseCaseAccountUpdateConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the Account Link is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new Account Link with the same parameters used to create the original Account Link, then redirect the user to the new Account Link URL so they can continue the flow. Make sure to authenticate the user before redirecting to the new Account Link, in case the URL leaks to a third party. If a new Account Link can't be generated, or if the redirect fails, you should display a useful error to the user.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url,omitempty"`
}

// Hash containing usage options.
type V2CoreAccountLinkUseCase struct {
	// Hash containing configuration options for an Account Link object that onboards a new account.
	AccountOnboarding *V2CoreAccountLinkUseCaseAccountOnboarding `json:"account_onboarding,omitempty"`
	// Hash containing configuration options for an Account Link that updates an existing account.
	AccountUpdate *V2CoreAccountLinkUseCaseAccountUpdate `json:"account_update,omitempty"`
	// Open Enum. The type of Account Link the user is requesting.
	Type V2CoreAccountLinkUseCaseType `json:"type"`
}

// Account Links let a platform create a temporary, single-use URL that an account can use to access a Stripe-hosted flow for collecting or updating required information.
type V2CoreAccountLink struct {
	APIResource
	// The ID of the connected account this Account Link applies to.
	Account string `json:"account"`
	// The timestamp at which this Account Link was created.
	Created time.Time `json:"created"`
	// The timestamp at which this Account Link will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The URL at which the account can access the Stripe-hosted flow.
	URL string `json:"url"`
	// Hash containing usage options.
	UseCase *V2CoreAccountLinkUseCase `json:"use_case"`
}
