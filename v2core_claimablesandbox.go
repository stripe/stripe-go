//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Status of the sandbox. One of `unclaimed`, `expired`, `claimed`.
type V2CoreClaimableSandboxStatus string

// List of values that V2CoreClaimableSandboxStatus can take
const (
	V2CoreClaimableSandboxStatusClaimed   V2CoreClaimableSandboxStatus = "claimed"
	V2CoreClaimableSandboxStatusExpired   V2CoreClaimableSandboxStatus = "expired"
	V2CoreClaimableSandboxStatusUnclaimed V2CoreClaimableSandboxStatus = "unclaimed"
)

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
	// The livemode sandbox Stripe account ID. This field is only set if the user activates their sandbox
	// and chooses to install your platform's Stripe App in their live account.
	OwnerAccount string `json:"owner_account,omitempty"`
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
	// The timestamp the sandbox was claimed. The value will be null if the sandbox status is not `claimed`.
	ClaimedAt time.Time `json:"claimed_at,omitempty"`
	// URL for user to claim sandbox into their existing Stripe account.
	// The value will be null if the sandbox status is `claimed` or `expired`.
	ClaimURL string `json:"claim_url,omitempty"`
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
	// Values prefilled during the creation of the sandbox. When a user claims the sandbox, they will be able to update these values.
	Prefill *V2CoreClaimableSandboxPrefill `json:"prefill"`
	// Data about the Stripe sandbox object.
	SandboxDetails *V2CoreClaimableSandboxSandboxDetails `json:"sandbox_details"`
	// Status of the sandbox. One of `unclaimed`, `expired`, `claimed`.
	Status V2CoreClaimableSandboxStatus `json:"status"`
}
