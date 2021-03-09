//
//
// File generated from our OpenAPI spec
//
//

package stripe

// LoginLinkParams is the set of parameters that can be used when creating a login_link.
// For more details see https://stripe.com/docs/api#create_login_link.
type LoginLinkParams struct {
	Params      `form:"*"`
	Account     *string `form:"-"` // Included in URL
	RedirectURL *string `form:"redirect_url"`
}

// LoginLink is the resource representing a login link for Express accounts.
// For more details see https://stripe.com/docs/api#login_link_object
type LoginLink struct {
	APIResource
	Created int64  `json:"created"`
	Object  string `json:"object"`
	URL     string `json:"url"`
}
