//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The status of the AutomaticRule object.
type V2TaxAutomaticRuleStatus string

// List of values that V2TaxAutomaticRuleStatus can take
const (
	V2TaxAutomaticRuleStatusActive   V2TaxAutomaticRuleStatus = "active"
	V2TaxAutomaticRuleStatusInactive V2TaxAutomaticRuleStatus = "inactive"
)

// An AutomaticRule holds automatic Tax configuration for a BillableItem.
type V2TaxAutomaticRule struct {
	APIResource
	// The ID of the BillableItem.
	BillableItem string `json:"billable_item"`
	// The time at which the AutomaticRule object was created.
	Created time.Time `json:"created"`
	// The ID of the AutomaticRule object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the AutomaticRule object.
	Status V2TaxAutomaticRuleStatus `json:"status"`
	// A TaxCode object that will be used for automatic tax calculations.
	TaxCode string `json:"tax_code"`
}
