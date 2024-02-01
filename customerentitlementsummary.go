//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieve the entitlement summary for a customer
type CustomerEntitlementSummaryParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CustomerEntitlementSummaryParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A summary of a customer's entitlements.
type CustomerEntitlementSummary struct {
	APIResource
	// The customer that is entitled to this feature.
	Customer string `json:"customer"`
	// The list of entitlements this customer has.
	Entitlements *CustomerEntitlementList `json:"entitlements"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
