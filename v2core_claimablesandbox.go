//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The app channel that will be used when pre-installing your app on the claimable sandbox.
type V2CoreClaimableSandboxAppChannel string

// List of values that V2CoreClaimableSandboxAppChannel can take
const (
	V2CoreClaimableSandboxAppChannelPublic  V2CoreClaimableSandboxAppChannel = "public"
	V2CoreClaimableSandboxAppChannelTesting V2CoreClaimableSandboxAppChannel = "testing"
)

// Indicates whether the platform app is installed on the sandbox's livemode owner account.
type V2CoreClaimableSandboxOwnerDetailsAppInstallStatus string

// List of values that V2CoreClaimableSandboxOwnerDetailsAppInstallStatus can take
const (
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusInstalled         V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "installed"
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusPendingInstall    V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "pending_install"
	V2CoreClaimableSandboxOwnerDetailsAppInstallStatusPendingOnboarding V2CoreClaimableSandboxOwnerDetailsAppInstallStatus = "pending_onboarding"
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
type V2CoreClaimableSandboxOnboardingLinkDetails struct {
	// The timestamp the onboarding link expires.
	ExpiresAt time.Time `json:"expires_at"`
	// The URL the user will be redirected to if the onboarding link is expired or invalid.
	// The URL specified should attempt to generate a new onboarding link,
	// and re-direct the user to this new onboarding link so that they can proceed with the onboarding flow.
	RefreshURL string `json:"refresh_url"`
	// URL that will redirect the user to either claim or onboard the claimable sandbox depending on its status.
	URL string `json:"url"`
}

// Details about the livemode owner account of the sandbox.
// This will be null until the sandbox is claimed.
type V2CoreClaimableSandboxOwnerDetails struct {
	// The ID of the livemode Stripe account that owns the sandbox.
	// This field is only set when owner_details.app_install_status is `installed`.
	Account string `json:"account,omitempty"`
	// Indicates whether the platform app is installed on the sandbox's livemode owner account.
	AppInstallStatus V2CoreClaimableSandboxOwnerDetailsAppInstallStatus `json:"app_install_status"`
}

// Values prefilled during the creation of the sandbox. When a user claims the sandbox, they will be able to update these values.
type V2CoreClaimableSandboxPrefill struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// send an email to this email address before the sandbox expires.
	Email string `json:"email"`
	// Name for the sandbox.
	Name string `json:"name"`
}

// Keys that can be used to set up an integration for this sandbox and operate on the account. This will be present only in the create response, and will be null in subsequent retrieve responses.
type V2CoreClaimableSandboxSandboxDetailsAPIKeys struct {
	// Used to communicate with [Stripe's MCP server](https://docs.stripe.com/mcp).
	// This allows LLM agents to securely operate on a Stripe account.
	Mcp string `json:"mcp,omitempty"`
	// Publicly accessible in a web or mobile app client-side code.
	Publishable string `json:"publishable"`
	// Should be stored securely in server-side code (such as an environment variable).
	Secret string `json:"secret"`
}

// Data about the Stripe sandbox object.
type V2CoreClaimableSandboxSandboxDetails struct {
	// The sandbox's Stripe account ID.
	Account string `json:"account"`
	// Keys that can be used to set up an integration for this sandbox and operate on the account. This will be present only in the create response, and will be null in subsequent retrieve responses.
	APIKeys *V2CoreClaimableSandboxSandboxDetailsAPIKeys `json:"api_keys,omitempty"`
}

// A claimable sandbox represents a Stripe sandbox that is anonymous.
// When it is created, it can be prefilled with specific metadata, such as email, name, or country.
// Claimable sandboxes can be claimed through a URL. When a user claims a sandbox through this URL,
// it will prompt them to create a new Stripe account. Or, it will allow them to claim this sandbox in their
// existing Stripe account.
// Claimable sandboxes have 60 days to be claimed. After this expiration time has passed,
// if the sandbox is not claimed, it will be deleted.
type V2CoreClaimableSandbox struct {
	APIResource
	// The app channel that will be used when pre-installing your app on the claimable sandbox.
	AppChannel V2CoreClaimableSandboxAppChannel `json:"app_channel"`
	// The timestamp the sandbox was claimed. The value will be null if the sandbox status is not `claimed`.
	ClaimedAt time.Time `json:"claimed_at,omitempty"`
	// When the sandbox is created.
	Created time.Time `json:"created"`
	// The timestamp the sandbox will expire. The value will be null if the sandbox is `claimed`.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier for the Claimable sandbox.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the onboarding link.
	OnboardingLinkDetails *V2CoreClaimableSandboxOnboardingLinkDetails `json:"onboarding_link_details"`
	// Details about the livemode owner account of the sandbox.
	// This will be null until the sandbox is claimed.
	OwnerDetails *V2CoreClaimableSandboxOwnerDetails `json:"owner_details,omitempty"`
	// Values prefilled during the creation of the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxPrefill `json:"prefill"`
	// Data about the Stripe sandbox object.
	SandboxDetails *V2CoreClaimableSandboxSandboxDetails `json:"sandbox_details"`
	// Status of the sandbox.
	Status V2CoreClaimableSandboxStatus `json:"status"`
}
