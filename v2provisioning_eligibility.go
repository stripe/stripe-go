//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Whether a project is eligible to provision resources with a provider, and any
// outstanding KYC requirements that must be satisfied first.
type V2ProvisioningEligibility struct {
	APIResource
	// Whether the project is eligible to provision resources with the provider.
	IsEligible bool `json:"is_eligible"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Outstanding requirements that must be satisfied before the project is eligible, if any.
	Requirements []string `json:"requirements"`
}
