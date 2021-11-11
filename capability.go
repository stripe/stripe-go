//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// CapabilityDisabledReason describes why a capability is disabled.
type CapabilityDisabledReason string

// List of values that CapabilityDisabledReason can take.
const (
	CapabilityDisabledReasonPendingOnboarding        CapabilityDisabledReason = "pending.onboarding"
	CapabilityDisabledReasonPendingReview            CapabilityDisabledReason = "pending.review"
	CapabilityDisabledReasonRejectedFraud            CapabilityDisabledReason = "rejected_fraud"
	CapabilityDisabledReasonRejectedListed           CapabilityDisabledReason = "rejected.listed"
	CapabilityDisabledReasonRejectedOther            CapabilityDisabledReason = "rejected.other"
	CapabilityDisabledReasonRequirementsFieldsNeeded CapabilityDisabledReason = "requirement.fields_needed"
)

// CapabilityStatus describes the different statuses for a capability's status.
type CapabilityStatus string

// List of values that CapabilityStatus can take.
const (
	CapabilityStatusActive      CapabilityStatus = "active"
	CapabilityStatusDisabled    CapabilityStatus = "disabled"
	CapabilityStatusInactive    CapabilityStatus = "inactive"
	CapabilityStatusPending     CapabilityStatus = "pending"
	CapabilityStatusUnrequested CapabilityStatus = "unrequested"
)

// CapabilityParams is the set of parameters that can be used when updating a capability.
// For more details see https://stripe.com/docs/api/capabilities/update
type CapabilityParams struct {
	Params    `form:"*"`
	Account   *string `form:"-"` // Included in URL
	Requested *bool   `form:"requested"`
}

// CapabilityListParams is the set of parameters that can be used when listing capabilities.
// For more detail see https://stripe.com/docs/api/capabilities/list
type CapabilityListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type CapabilityFutureRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}
type CapabilityFutureRequirements struct {
	Alternatives        []*CapabilityFutureRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                      `json:"current_deadline"`
	CurrentlyDue        []string                                   `json:"currently_due"`
	DisabledReason      string                                     `json:"disabled_reason"`
	Errors              []*CapabilityFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                                   `json:"eventually_due"`
	PastDue             []string                                   `json:"past_due"`
	PendingVerification []string                                   `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// CapabilityRequirements represents information that needs to be collected for a capability.
type CapabilityRequirements struct {
	Alternatives        []*CapabilityRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                `json:"current_deadline"`
	CurrentlyDue        []string                             `json:"currently_due"`
	DisabledReason      CapabilityDisabledReason             `json:"disabled_reason"`
	Errors              []*AccountRequirementsError          `json:"errors"`
	EventuallyDue       []string                             `json:"eventually_due"`
	PastDue             []string                             `json:"past_due"`
	PendingVerification []string                             `json:"pending_verification"`
}

// Capability is the resource representing a Stripe capability.
// For more details see https://stripe.com/docs/api/capabilities
type Capability struct {
	APIResource
	Account            *Account                      `json:"account"`
	FutureRequirements *CapabilityFutureRequirements `json:"future_requirements"`
	ID                 string                        `json:"id"`
	Object             string                        `json:"object"`
	Requested          bool                          `json:"requested"`
	RequestedAt        int64                         `json:"requested_at"`
	Requirements       *CapabilityRequirements       `json:"requirements"`
	Status             CapabilityStatus              `json:"status"`
}

// CapabilityList is a list of capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data"`
}

// UnmarshalJSON handles deserialization of a Capability.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Capability) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type capability Capability
	var v capability
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Capability(v)
	return nil
}
