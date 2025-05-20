//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
type LoginLinkParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *LoginLinkParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
type LoginLinkCreateParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *LoginLinkCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Login Links are single-use URLs that takes an Express account to the login page for their Stripe dashboard.
// A Login Link differs from an [Account Link](https://stripe.com/docs/api/account_links) in that it takes the user directly to their [Express dashboard for the specified account](https://stripe.com/docs/connect/integrate-express-dashboard#create-login-link)
type LoginLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL for the login link.
	URL string `json:"url"`
}
