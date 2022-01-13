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
	Params      `form:"*"`
	Account     *string `form:"-"` // Included in URL
	RedirectURL *string `form:"redirect_url"`
}
type LoginLink struct {
	APIResource
	Created int64  `json:"created"`
	Object  string `json:"object"`
	URL     string `json:"url"`
}
