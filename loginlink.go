//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a single-use login link for an Express account to access their Stripe dashboard.
//
// You may only create login links for [Express accounts](https://stripe.com/docs/connect/express-accounts) connected to your platform.
type LoginLinkParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Where to redirect the user after they log out of their dashboard.
	RedirectURL *string `form:"redirect_url"`
}
type LoginLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL for the login link.
	URL string `json:"url"`
}
