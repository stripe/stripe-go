//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
type AccountSessionParams struct {
	Params `form:"*"`
	// The identifier of the account to create an Account Session for.
	Account *string `form:"account"`
}

// An AccountSession allows a Connect platform to grant access to a connected account in Connect Elements.
//
// We recommend that you create an AccountSession each time you need to display an embedded UI
// to your user. Do not save AccountSessions to your database as they expire relatively
// quickly, and cannot be used more than once.
//
// Related guide: [Connect Elements](https://stripe.com/docs/connect/get-started-connect-elements).
type AccountSession struct {
	APIResource
	// The ID of the account the AccountSession was created for
	Account string `json:"account"`
	// The client secret of this AccountSession. Used on the client to set up secure access to the given `account`.
	//
	// The client secret can be used to provide access to `account` from your frontend. It should not be stored, logged, or exposed to anyone other than the connected account. Make sure that you have TLS enabled on any page that includes the client secret.
	//
	// Refer to our docs to [setup Connect Elements](https://stripe.com/docs/connect/get-started-connect-elements) and learn about how `client_secret` should be handled.
	ClientSecret string `json:"client_secret"`
	// The timestamp at which this AccountSession will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
