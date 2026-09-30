//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current provider-issued access configuration for a Provisioning Resource.
type V2ProvisioningResourceAccessConfiguration struct {
	APIResource
	// Provider-defined configuration names mapped to their secret string values.
	Configuration map[string]string `json:"configuration"`
	// Time at which this credential generation became current.
	Created time.Time `json:"created"`
	// Time at which these credentials cease to be valid, when supplied by the Provider.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Whether the referenced Resource uses Stripe live-mode objects.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Provisioning Resource to which this access configuration belongs.
	Resource string `json:"resource"`
}
