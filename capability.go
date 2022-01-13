//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// If the capability is disabled, this string describes why. Can be `requirements.past_due`, `requirements.pending_verification`, `listed`, `platform_paused`, `rejected.fraud`, `rejected.listed`, `rejected.terms_of_service`, `rejected.other`, `under_review`, or `other`.
//
// `rejected.unsupported_business` means that the account's business is not supported by the capability. For example, payment methods may restrict the businesses they support in their terms of service:
//
// - [Afterpay Clearpay's terms of service](https://stripe.com/afterpay-clearpay/legal#restricted-businesses)
//
// If you believe that the rejection is in error, please contact support at https://support.stripe.com/contact/ for assistance.
type CapabilityDisabledReason string

// List of values that CapabilityDisabledReason can take
const (
	CapabilityDisabledReasonPendingOnboarding        CapabilityDisabledReason = "pending.onboarding"
	CapabilityDisabledReasonPendingReview            CapabilityDisabledReason = "pending.review"
	CapabilityDisabledReasonRejectedFraud            CapabilityDisabledReason = "rejected_fraud"
	CapabilityDisabledReasonRejectedListed           CapabilityDisabledReason = "rejected.listed"
	CapabilityDisabledReasonRejectedOther            CapabilityDisabledReason = "rejected.other"
	CapabilityDisabledReasonRequirementsFieldsNeeded CapabilityDisabledReason = "requirement.fields_needed"
)

// The status of the capability. Can be `active`, `inactive`, `pending`, or `unrequested`.
type CapabilityStatus string

// List of values that CapabilityStatus can take
const (
	CapabilityStatusActive      CapabilityStatus = "active"
	CapabilityStatusDisabled    CapabilityStatus = "disabled"
	CapabilityStatusInactive    CapabilityStatus = "inactive"
	CapabilityStatusPending     CapabilityStatus = "pending"
	CapabilityStatusUnrequested CapabilityStatus = "unrequested"
)

// Returns a list of capabilities associated with the account. The capabilities are returned sorted by creation date, with the most recent capability appearing first.
type CapabilityListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
}

// Retrieves information about the specified Account Capability.
type CapabilityParams struct {
	Params    `form:"*"`
	Account   *string `form:"-"` // Included in URL
	Requested *bool   `form:"requested"`
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

// This is an object representing a capability for a Stripe account.
//
// Related guide: [Account capabilities](https://stripe.com/docs/connect/account-capabilities).
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

// CapabilityList is a list of Capabilities as retrieved from a list endpoint.
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
