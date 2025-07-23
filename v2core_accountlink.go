//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of AccountLink the user is requesting.
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
	V2CoreAccountLinkUseCaseAccountOnboardingConfigurationStorer    V2CoreAccountLinkUseCaseAccountOnboardingConfiguration = "storer"
)

// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
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
	V2CoreAccountLinkUseCaseAccountUpdateConfigurationStorer    V2CoreAccountLinkUseCaseAccountUpdateConfiguration = "storer"
)

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptions struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFields `json:"fields"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsFutureRequirements `json:"future_requirements"`
}

// Indicates that the AccountLink provided should onboard an account.
type V2CoreAccountLinkUseCaseAccountOnboarding struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptions `json:"collection_options"`
	// Open Enum. A v2/core/account can be configured to enable certain functionality. The configuration param targets the v2/core/account_link to collect information for the specified v2/core/account configuration/s.
	Configurations []V2CoreAccountLinkUseCaseAccountOnboardingConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url"`
}

// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
type V2CoreAccountLinkUseCaseAccountUpdateCollectionOptions struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify collection_options, the default value is currently_due.
	Fields V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFields `json:"fields"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsFutureRequirements `json:"future_requirements"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2CoreAccountLinkUseCaseAccountUpdate struct {
	// Specifies the requirements that Stripe collects from v2/core/accounts in the Onboarding flow.
	CollectionOptions *V2CoreAccountLinkUseCaseAccountUpdateCollectionOptions `json:"collection_options"`
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []V2CoreAccountLinkUseCaseAccountUpdateConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url"`
}

// The use case of AccountLink the user is requesting.
type V2CoreAccountLinkUseCase struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2CoreAccountLinkUseCaseAccountOnboarding `json:"account_onboarding"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2CoreAccountLinkUseCaseAccountUpdate `json:"account_update"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type V2CoreAccountLinkUseCaseType `json:"type"`
}

// AccountLinks are the means by which a Merchant grants an Account permission to access Stripe-hosted applications, such as Recipient Onboarding. This API is only available for users enrolled in the public preview for Accounts v2.
type V2CoreAccountLink struct {
	APIResource
	// The ID of the Account the link was created for.
	Account string `json:"account"`
	// The timestamp at which this AccountLink was created.
	Created time.Time `json:"created"`
	// The timestamp at which this AccountLink will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The URL for the AccountLink.
	URL string `json:"url"`
	// The use case of AccountLink the user is requesting.
	UseCase *V2CoreAccountLinkUseCase `json:"use_case"`
}
