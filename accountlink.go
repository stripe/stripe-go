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
	Params `form:"*"`
	// The identifier of the account to create an account link for.
	Account *string `form:"account"`
	// Which information the platform needs to collect from the user. One of `currently_due` or `eventually_due`. Default is `currently_due`.
	Collect *string `form:"collect"`
	// The URL the user will be redirected to if the account link is expired, has been previously-visited, or is otherwise invalid. The URL you specify should attempt to generate a new account link with the same parameters used to create the original account link, then redirect the user to the new account link's URL so they can continue with Connect Onboarding. If a new account link cannot be generated or the redirect fails you should display a useful error to the user.
	RefreshURL *string `form:"refresh_url"`
	// The URL that the user will be redirected to upon leaving or completing the linked flow.
	ReturnURL *string `form:"return_url"`
	// The type of account link the user is requesting. Possible values are `account_onboarding` or `account_update`.
	Type *string `form:"type"`
}

// Account Links are the means by which a Connect platform grants a connected account permission to access
// Stripe-hosted applications, such as Connect Onboarding.
//
// Related guide: [Connect Onboarding](https://stripe.com/docs/connect/connect-onboarding).
type AccountLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The timestamp at which this account link will expire.
	ExpiresAt int64 `json:"expires_at"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL for the account link.
	URL string `json:"url"`
}
