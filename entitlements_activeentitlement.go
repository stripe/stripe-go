//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieve a list of active entitlements for a customer
type EntitlementsActiveEntitlementListParams struct {
	ListParams `form:"*"`
	// The ID of the customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *EntitlementsActiveEntitlementListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// An active entitlement describes access to a feature for a customer.
type EntitlementsActiveEntitlement struct {
	// The feature that the customer is entitled to.
	Feature string `json:"feature"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// A unique key you provide as your own system identifier. This may be up to 80 characters.
	LookupKey string `json:"lookup_key"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// EntitlementsActiveEntitlementList is a list of ActiveEntitlements as retrieved from a list endpoint.
type EntitlementsActiveEntitlementList struct {
	APIResource
	ListMeta
	Data []*EntitlementsActiveEntitlement `json:"data"`
}