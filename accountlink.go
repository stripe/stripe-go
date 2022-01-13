//
//
// File generated from our OpenAPI spec
//
//

package stripe

// AccountLinkType is the type of an account link.
type AccountLinkType string

// List of values that AccountLinkType can take.
const (
	AccountLinkTypeAccountOnboarding AccountLinkType = "account_onboarding"
	AccountLinkTypeAccountUpdate     AccountLinkType = "account_update"
)

// AccountLinkCollect describes what information the platform wants to collect with the account link.
type AccountLinkCollect string

// List of values that AccountLinkCollect can take.
const (
	AccountLinkCollectCurrentlyDue  AccountLinkCollect = "currently_due"
	AccountLinkCollectEventuallyDue AccountLinkCollect = "eventually_due"
)

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
type AccountLinkParams struct {
	Params     `form:"*"`
	Account    *string `form:"account"`
	Collect    *string `form:"collect"`
	RefreshURL *string `form:"refresh_url"`
	ReturnURL  *string `form:"return_url"`
	Type       *string `form:"type"`
}

// Account Links are the means by which a Connect platform grants a connected account permission to access
// Stripe-hosted applications, such as Connect Onboarding.
//
// Related guide: [Connect Onboarding](https://stripe.com/docs/connect/connect-onboarding).
type AccountLink struct {
	APIResource
	Created   int64  `json:"created"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
	URL       string `json:"url"`
}
