//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of AccountLink the user is requesting.
type V2AccountLinkUseCaseType string

// List of values that V2AccountLinkUseCaseType can take
const (
	V2AccountLinkUseCaseTypeAccountOnboarding V2AccountLinkUseCaseType = "account_onboarding"
	V2AccountLinkUseCaseTypeAccountUpdate     V2AccountLinkUseCaseType = "account_update"
)

// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
type V2AccountLinkUseCaseAccountOnboardingConfiguration string

// List of values that V2AccountLinkUseCaseAccountOnboardingConfiguration can take
const (
	V2AccountLinkUseCaseAccountOnboardingConfigurationRecipient V2AccountLinkUseCaseAccountOnboardingConfiguration = "recipient"
)

// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
type V2AccountLinkUseCaseAccountUpdateConfiguration string

// List of values that V2AccountLinkUseCaseAccountUpdateConfiguration can take
const (
	V2AccountLinkUseCaseAccountUpdateConfigurationRecipient V2AccountLinkUseCaseAccountUpdateConfiguration = "recipient"
)

// Indicates that the AccountLink provided should onboard an account.
type V2AccountLinkUseCaseAccountOnboarding struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []V2AccountLinkUseCaseAccountOnboardingConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url"`
}

// Indicates that the AccountLink provided should update a previously onboarded account.
type V2AccountLinkUseCaseAccountUpdate struct {
	// Open Enum. A v2/account can be configured to enable certain functionality. The configuration param targets the v2/account_link to collect information for the specified v2/account configuration/s.
	Configurations []V2AccountLinkUseCaseAccountUpdateConfiguration `json:"configurations"`
	// The URL the user will be redirected to if the AccountLink is expired, has been used, or is otherwise invalid. The URL you specify should attempt to generate a new AccountLink with the same parameters used to create the original AccountLink, then redirect the user to the new AccountLink's URL so they can continue the flow. If a new AccountLink cannot be generated or the redirect fails you should display a useful error to the user. Please make sure to implement authentication before redirecting the user in case this URL is leaked to a third party.
	RefreshURL string `json:"refresh_url"`
	// The URL that the user will be redirected to upon completing the linked flow.
	ReturnURL string `json:"return_url"`
}

// The use case of AccountLink the user is requesting.
type V2AccountLinkUseCase struct {
	// Indicates that the AccountLink provided should onboard an account.
	AccountOnboarding *V2AccountLinkUseCaseAccountOnboarding `json:"account_onboarding"`
	// Indicates that the AccountLink provided should update a previously onboarded account.
	AccountUpdate *V2AccountLinkUseCaseAccountUpdate `json:"account_update"`
	// Open Enum. The type of AccountLink the user is requesting.
	Type V2AccountLinkUseCaseType `json:"type"`
}

// AccountLinks are the means by which a Merchant grants an Account permission to access Stripe-hosted application, such as Recipient Onboarding.
type V2AccountLink struct {
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
	UseCase *V2AccountLinkUseCase `json:"use_case"`
}
