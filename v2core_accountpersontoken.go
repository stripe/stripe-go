//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Person Tokens are single-use tokens which tokenize person information, and are used for creating or updating a Person.
type V2CoreAccountPersonToken struct {
	APIResource
	// Time at which the token was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Time at which the token will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// Unique identifier for the token.
	ID string `json:"id"`
	// Has the value `true` if the token exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Determines if the token has already been used (tokens can only be used once).
	Used bool `json:"used"`
}
