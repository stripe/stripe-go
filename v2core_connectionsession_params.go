//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Create a new ConnectionSession.
type V2CoreConnectionSessionParams struct {
	Params `form:"*"`
	// The Account the ConnectionSession will create a connection for.
	Account *string `form:"account" json:"account,omitempty"`
	// The Connection types that the ConnectionSession is allowed to establish.
	AllowedConnectionTypes []*string `form:"allowed_connection_types" json:"allowed_connection_types,omitempty"`
	// The access that should be collected with the ConnectionSession.
	RequestedAccess []*string `form:"requested_access" json:"requested_access,omitempty"`
}

// Create a new ConnectionSession.
type V2CoreConnectionSessionCreateParams struct {
	Params `form:"*"`
	// The Account the ConnectionSession will create a connection for.
	Account *string `form:"account" json:"account"`
	// The Connection types that the ConnectionSession is allowed to establish.
	AllowedConnectionTypes []*string `form:"allowed_connection_types" json:"allowed_connection_types,omitempty"`
	// The access that should be collected with the ConnectionSession.
	RequestedAccess []*string `form:"requested_access" json:"requested_access,omitempty"`
}

// Retrieve a ConnectionSession.
type V2CoreConnectionSessionRetrieveParams struct {
	Params `form:"*"`
}
