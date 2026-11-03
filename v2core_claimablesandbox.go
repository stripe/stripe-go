//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The distribution channel used to pre-install your Stripe App on the claimable sandbox.
// Use `testing` to install your app's external test version or `public` to install its published version.
type V2CoreClaimableSandboxChannel string

// List of values that V2CoreClaimableSandboxChannel can take
const (
	V2CoreClaimableSandboxChannelPublic  V2CoreClaimableSandboxChannel = "public"
	V2CoreClaimableSandboxChannelTesting V2CoreClaimableSandboxChannel = "testing"
)

// Indicates whether the platform app is installed on the sandbox's live mode owner account.
type V2CoreClaimableSandboxOwnerDetailsAppInstallStatus string

// List of values that V2CoreClaimableSandboxOwnerDetailsAppInstallStatus can take
const (
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusInstalled         V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "installed"
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusPendingInstall    V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "pending_install"
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusPendingOnboarding V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "pending_onboarding"
)

// Current status of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatus string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatus can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatusActive  V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatus = "active"
	V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatusExpired V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatus = "expired"
)

// Type of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpType string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpType can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpTypeSecretKey V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpType = "secret_key"
)

// Current status of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatus string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatus can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatusActive  V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatus = "active"
	V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatusExpired V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatus = "expired"
)

// Type of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableType string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableType can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableTypePublishableKey V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableType = "publishable_key"
)

// Current status of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatus string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatus can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatusActive  V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatus = "active"
	V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatusExpired V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatus = "expired"
)

// Type of the API key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretType string

// List of values that V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretType can take
const (
	V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretTypeSecretKey V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretType = "secret_key"
)

// Status of the sandbox.
type V2CoreClaimableSandboxStatus string

// List of values that V2CoreClaimableSandboxStatus can take
const (
	V2CoreClaimableSandboxStatusClaimed   V2CoreClaimableSandboxStatus = "claimed"
	V2CoreClaimableSandboxStatusExpired   V2CoreClaimableSandboxStatus = "expired"
	V2CoreClaimableSandboxStatusLive      V2CoreClaimableSandboxStatus = "live"
	V2CoreClaimableSandboxStatusUnclaimed V2CoreClaimableSandboxStatus = "unclaimed"
)

// Details about the onboarding link.
// Null when the sandbox status is `live` or `expired`.
type V2CoreClaimableSandboxOnboardingLinkDetails struct {
	// Time at which the onboarding link expires. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	ExpiresAt time.Time `json:"expires_at"`
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL string `json:"refresh_url"`
	// The URL offered to the user as a way back to your app after they claim the sandbox.
	ReturnURL string `json:"return_url,omitempty"`
	// The URL that lets a user claim the sandbox or continue onboarding their Stripe account,
	// depending on the sandbox's status.
	URL string `json:"url"`
}

// Details about the live mode owner account of the sandbox.
// This will be null until the sandbox is claimed.
type V2CoreClaimableSandboxOwnerDetails struct {
	// The ID of the live mode Stripe account that owns the sandbox.
	// This field is only set when `owner_details.app_install_status` is `installed`.
	Account string `json:"account,omitempty"`
	// Indicates whether the platform app is installed on the sandbox's live mode owner account.
	AppInstallStatus V2CoreClaimableSandboxOwnerDetailsAppInstallStatus `json:"app_install_status,omitempty"`
}

// Values prefilled during the creation of the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxPrefill struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Email address of the user who will claim the sandbox.
	// Stripe sends this address an invitation to claim the sandbox and a reminder before the unclaimed sandbox expires.
	Email string `json:"email"`
	// Name for the sandbox.
	Name string `json:"name"`
}

// Token for the secret key. Present only in the create response;
// null in subsequent retrieve responses.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpSecretKey struct {
	// The plaintext token for the API key.
	Token string `json:"token,omitempty"`
}

// Object hash for the sandbox's MCP key.
// Null unless `enable_mcp_access: true` was passed in the creation request.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysMcp struct {
	// Time at which the API key was created. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Time at which the API key expires. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier of the API key.
	ID string `json:"id"`
	// Token for the secret key. Present only in the create response;
	// null in subsequent retrieve responses.
	SecretKey *V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpSecretKey `json:"secret_key,omitempty"`
	// Current status of the API key.
	Status V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpStatus `json:"status"`
	// Type of the API key.
	Type V2CoreClaimableSandboxSandboxDetailsAPIKeysMcpType `json:"type"`
}

// Token for the publishable key. Present only in the create response;
// null in subsequent retrieve responses.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishablePublishableKey struct {
	// The plaintext token for the API key.
	Token string `json:"token,omitempty"`
}

// Object hash for the sandbox's publishable key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishable struct {
	// Time at which the API key was created. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Time at which the API key expires. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier of the API key.
	ID string `json:"id"`
	// Token for the publishable key. Present only in the create response;
	// null in subsequent retrieve responses.
	PublishableKey *V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishablePublishableKey `json:"publishable_key,omitempty"`
	// Current status of the API key.
	Status V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableStatus `json:"status"`
	// Type of the API key.
	Type V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishableType `json:"type"`
}

// Token for the secret key. Present only in the create response;
// null in subsequent retrieve responses.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretSecretKey struct {
	// The plaintext token for the API key.
	Token string `json:"token,omitempty"`
}

// Object hash for the sandbox's secret key.
type V2CoreClaimableSandboxSandboxDetailsAPIKeysSecret struct {
	// Time at which the API key was created. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Time at which the API key expires. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier of the API key.
	ID string `json:"id"`
	// Token for the secret key. Present only in the create response;
	// null in subsequent retrieve responses.
	SecretKey *V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretSecretKey `json:"secret_key,omitempty"`
	// Current status of the API key.
	Status V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretStatus `json:"status"`
	// Type of the API key.
	Type V2CoreClaimableSandboxSandboxDetailsAPIKeysSecretType `json:"type"`
}

// Keys that can be used to set up an integration for this sandbox and operate on the account.
// Present in the create response. To fetch this field in the retrieve response,
// pass `sandbox_details.api_keys` in the `include` parameter.
type V2CoreClaimableSandboxSandboxDetailsAPIKeys struct {
	// Object hash for the sandbox's MCP key.
	// Null unless `enable_mcp_access: true` was passed in the creation request.
	Mcp *V2CoreClaimableSandboxSandboxDetailsAPIKeysMcp `json:"mcp,omitempty"`
	// Object hash for the sandbox's publishable key.
	Publishable *V2CoreClaimableSandboxSandboxDetailsAPIKeysPublishable `json:"publishable"`
	// Object hash for the sandbox's secret key.
	Secret *V2CoreClaimableSandboxSandboxDetailsAPIKeysSecret `json:"secret"`
}

// Data about the Stripe sandbox object.
type V2CoreClaimableSandboxSandboxDetails struct {
	// The sandbox's Stripe account ID.
	Account string `json:"account"`
	// Keys that can be used to set up an integration for this sandbox and operate on the account.
	// Present in the create response. To fetch this field in the retrieve response,
	// pass `sandbox_details.api_keys` in the `include` parameter.
	APIKeys *V2CoreClaimableSandboxSandboxDetailsAPIKeys `json:"api_keys,omitempty"`
}

// A claimable sandbox is a Stripe sandbox that a user can claim into a new or existing Stripe account
// through an onboarding link. You can prefill it with an email address, sandbox name, and country.
// Claimable sandboxes have 60 days to be claimed. After this expiration time has passed,
// if the sandbox is not claimed, it will be deleted.
type V2CoreClaimableSandbox struct {
	APIResource
	// The distribution channel used to pre-install your Stripe App on the claimable sandbox.
	// Use `testing` to install your app's external test version or `public` to install its published version.
	Channel V2CoreClaimableSandboxChannel `json:"channel"`
	// Time at which the sandbox was claimed. Set when the status is `claimed` or `live`.
	// Represented as an RFC 3339 timestamp in UTC with millisecond precision,
	// for example `2022-09-18T13:22:18.123Z`.
	ClaimedAt time.Time `json:"claimed_at,omitempty"`
	// Time at which the sandbox was created. Represented as an RFC 3339 timestamp in UTC
	// with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Time at which the unclaimed sandbox expires. Null when the status is `claimed` or `live`.
	// Represented as an RFC 3339 timestamp in UTC with millisecond precision,
	// for example `2022-09-18T13:22:18.123Z`.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier for the claimable sandbox.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the onboarding link.
	// Null when the sandbox status is `live` or `expired`.
	OnboardingLinkDetails *V2CoreClaimableSandboxOnboardingLinkDetails `json:"onboarding_link_details,omitempty"`
	// Details about the live mode owner account of the sandbox.
	// This will be null until the sandbox is claimed.
	OwnerDetails *V2CoreClaimableSandboxOwnerDetails `json:"owner_details,omitempty"`
	// Values prefilled during the creation of the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxPrefill `json:"prefill"`
	// Data about the Stripe sandbox object.
	SandboxDetails *V2CoreClaimableSandboxSandboxDetails `json:"sandbox_details"`
	// Status of the sandbox.
	Status V2CoreClaimableSandboxStatus `json:"status"`
}
