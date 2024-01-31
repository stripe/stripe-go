//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of feature.
type CustomerEntitlementType string

// List of values that CustomerEntitlementType can take
const (
	CustomerEntitlementTypeQuantity CustomerEntitlementType = "quantity"
	CustomerEntitlementTypeSwitch   CustomerEntitlementType = "switch"
)

// Retrieve a list of entitlements for a customer
type CustomerEntitlementListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CustomerEntitlementListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Contains information about entitlements relating to features with type=quantity. Required when the feature has type=quantity.
type CustomerEntitlementQuantity struct {
	// The total quantity available to the customer.
	TotalAvailable int64 `json:"total_available"`
}

// A entitlement for a customer describes access to a feature.
type CustomerEntitlement struct {
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
	// Contains information about entitlements relating to features with type=quantity. Required when the feature has type=quantity.
	Quantity *CustomerEntitlementQuantity `json:"quantity"`
	// The type of feature.
	Type CustomerEntitlementType `json:"type"`
}

// CustomerEntitlementList is a list of CustomerEntitlements as retrieved from a list endpoint.
type CustomerEntitlementList struct {
	APIResource
	ListMeta
	Data []*CustomerEntitlement `json:"data"`
}
