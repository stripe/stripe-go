//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Keys that can be used to set up an integration for this sandbox and operate on the account.
type V2CoreClaimableSandboxAPIKeys struct {
	// Used to communicate with [Stripe's MCP server](https://docs.stripe.com/mcp).
	// This allows LLM agents to securely operate on a Stripe account.
	Mcp string `json:"mcp,omitempty"`
	// Publicly accessible in a web or mobile app client-side code.
	Publishable string `json:"publishable"`
	// Should be stored securely in server-side code (such as an environment variable).
	Secret string `json:"secret"`
}

// Values prefilled during the creation of the sandbox.
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

// A claimable sandbox represents a Stripe sandbox that is anonymous.
// When it is created, it can be prefilled with specific metadata, such as email, name, or country.
// Claimable sandboxes can be claimed through a URL. When a user claims a sandbox through this URL,
// it will prompt them to create a new Stripe account. Or, it will allow them to claim this sandbox in their
// existing Stripe account.
// Claimable sandboxes have 60 days to be claimed. After this expiration time has passed,
// if the sandbox is not claimed, it will be deleted.
type V2CoreClaimableSandbox struct {
	APIResource
	// Keys that can be used to set up an integration for this sandbox and operate on the account.
	APIKeys *V2CoreClaimableSandboxAPIKeys `json:"api_keys"`
	// URL for user to claim sandbox into their existing Stripe account.
	ClaimURL string `json:"claim_url"`
	// When the sandbox is created.
	Created time.Time `json:"created"`
	// Unique identifier for the Claimable sandbox.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Values prefilled during the creation of the sandbox.
	Prefill *V2CoreClaimableSandboxPrefill `json:"prefill"`
}
