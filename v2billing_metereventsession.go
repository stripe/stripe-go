//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A Meter Event Session is an authentication session for the high-throughput meter event API. Meter Event Sessions provide temporary authentication tokens with expiration times, enabling secure and efficient bulk submission of usage events.
type V2BillingMeterEventSession struct {
	APIResource
	// The authentication token for this session.  Use this token when calling the
	// high-throughput meter event API.
	AuthenticationToken string `json:"authentication_token"`
	// The creation time of this session.
	Created time.Time `json:"created"`
	// The time at which this session will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// The unique id of this auth session.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
