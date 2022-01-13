//
//
// File generated from our OpenAPI spec
//
//

package stripe

// To connect to a reader the Stripe Terminal SDK needs to retrieve a short-lived connection token from Stripe, proxied through your server. On your backend, add an endpoint that creates and returns a connection token.
type TerminalConnectionTokenParams struct {
	Params   `form:"*"`
	Location string `form:"location"`
}

// A Connection Token is used by the Stripe Terminal SDK to connect to a reader.
//
// Related guide: [Fleet Management](https://stripe.com/docs/terminal/fleet/locations).
type TerminalConnectionToken struct {
	APIResource
	Location string `json:"location"`
	Object   string `json:"object"`
	Secret   string `json:"secret"`
}
