package stripe

// LoginLinkParams is the set of parameters that can be used when creating a login_link.
// For more details see https://stripe.com/docs/api#create_login_link.
type LoginLinkParams struct {
	Account string
}

// LoginLink is the resource representing a login link for Express accounts.
// For more details see https://stripe.com/docs/api#login_link_object
type LoginLink struct {
	Created int64  `json:"created"`
	Url     string `json:"url"`
}
