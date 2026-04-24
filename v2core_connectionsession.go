//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The Connection types that the Connection Session is allowed to establish.
type V2CoreConnectionSessionAllowedConnectionType string

// List of values that V2CoreConnectionSessionAllowedConnectionType can take
const (
	V2CoreConnectionSessionAllowedConnectionTypeLink V2CoreConnectionSessionAllowedConnectionType = "link"
)

// The access granted to the Account by the Connection.
type V2CoreConnectionSessionConnectionGrantedAccess string

// List of values that V2CoreConnectionSessionConnectionGrantedAccess can take
const (
	V2CoreConnectionSessionConnectionGrantedAccessPayoutMethods V2CoreConnectionSessionConnectionGrantedAccess = "payout_methods"
)

// The type of the Connection.
type V2CoreConnectionSessionConnectionType string

// List of values that V2CoreConnectionSessionConnectionType can take
const (
	V2CoreConnectionSessionConnectionTypeLink V2CoreConnectionSessionConnectionType = "link"
)

// The access that is collected with the Connection Session.
type V2CoreConnectionSessionRequestedAccess string

// List of values that V2CoreConnectionSessionRequestedAccess can take
const (
	V2CoreConnectionSessionRequestedAccessPayoutMethods V2CoreConnectionSessionRequestedAccess = "payout_methods"
)

// The Connection created by the ConnectionSession.
type V2CoreConnectionSessionConnection struct {
	// The access granted to the Account by the Connection.
	GrantedAccess []V2CoreConnectionSessionConnectionGrantedAccess `json:"granted_access,omitempty"`
	// The type of the Connection.
	Type V2CoreConnectionSessionConnectionType `json:"type"`
}

// A short-lived, single-use session used to launch client-side Link onboarding
// that connects an Account to a Link consumer account.
type V2CoreConnectionSession struct {
	APIResource
	// The Account this Connection Session was created for.
	Account string `json:"account"`
	// The Connection types that the Connection Session is allowed to establish.
	AllowedConnectionTypes []V2CoreConnectionSessionAllowedConnectionType `json:"allowed_connection_types,omitempty"`
	// The client secret of this Connection Session. Used on the client to set up secure access to the given Account.
	ClientSecret string `json:"client_secret"`
	// The Connection created by the ConnectionSession.
	Connection *V2CoreConnectionSessionConnection `json:"connection,omitempty"`
	// Time at which the ConnectionSession was created.
	Created time.Time `json:"created"`
	// The unique identifier for this ConnectionSession.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The access that is collected with the Connection Session.
	RequestedAccess []V2CoreConnectionSessionRequestedAccess `json:"requested_access,omitempty"`
}
