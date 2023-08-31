//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapabilityListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves information about the specified Account Capability.
type CapabilityParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// To request a new capability for an account, pass true. There can be a delay before the requested capability becomes active. If the capability has any activation requirements, the response includes them in the `requirements` arrays.
	//
	// If a capability isn't permanent, you can remove it from the account by passing false. Most capabilities are permanent after they've been requested. Attempting to remove a permanent capability returns an error.
	Requested *bool `form:"requested"`
}

// AddExpand appends a new field to expand.
func (p *CapabilityParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityFutureRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type CapabilityFutureRequirementsError struct {
	// The code for the type of error.
	Code string `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}
type CapabilityFutureRequirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*CapabilityFutureRequirementsAlternative `json:"alternatives"`
	// Date on which `future_requirements` merges with the main `requirements` hash and `future_requirements` becomes empty. After the transition, `currently_due` requirements may immediately become `past_due`, but the account may also be given a grace period depending on the capability's enablement state prior to transitioning.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be collected to keep the capability enabled. If not collected by `future_requirements[current_deadline]`, these fields will transition to the main `requirements` hash.
	CurrentlyDue []string `json:"currently_due"`
	// This is typed as a string for consistency with `requirements.disabled_reason`, but it safe to assume `future_requirements.disabled_reason` is empty because fields in `future_requirements` will never disable the account.
	DisabledReason string `json:"disabled_reason"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*CapabilityFutureRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by `requirements.current_deadline`. These fields need to be collected to enable the capability on the account. New fields will never appear here; `future_requirements.past_due` will always be a subset of `requirements.past_due`.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due` or `currently_due`.
	PendingVerification []string `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}
type CapabilityRequirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*CapabilityRequirementsAlternative `json:"alternatives"`
	// Date by which the fields in `currently_due` must be collected to keep the capability enabled for the account. These fields may disable the capability sooner if the next threshold is reached before they are collected.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be collected to keep the capability enabled. If not collected by `current_deadline`, these fields appear in `past_due` as well, and the capability is disabled.
	CurrentlyDue []string `json:"currently_due"`
	// If the capability is disabled, this string describes why. Can be `requirements.past_due`, `requirements.pending_verification`, `listed`, `platform_paused`, `rejected.fraud`, `rejected.listed`, `rejected.terms_of_service`, `rejected.other`, `under_review`, or `other`.
	//
	// `rejected.unsupported_business` means that the account's business is not supported by the capability. For example, payment methods may restrict the businesses they support in their terms of service:
	//
	// - [Afterpay Clearpay's terms of service](https://stripe.com/afterpay-clearpay/legal#restricted-businesses)
	//
	// If you believe that the rejection is in error, please contact support at https://support.stripe.com/contact/ for assistance.
	DisabledReason CapabilityDisabledReason `json:"disabled_reason"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*AccountRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well, and `current_deadline` becomes set.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by `current_deadline`. These fields need to be collected to enable the capability on the account.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due`, `currently_due`, or `past_due`.
	PendingVerification []string `json:"pending_verification"`
}

// This is an object representing a capability for a Stripe account.
//
// Related guide: [Account capabilities](https://stripe.com/docs/connect/account-capabilities)
type Capability struct {
	APIResource
	// The account for which the capability enables functionality.
	Account            *Account                      `json:"account"`
	FutureRequirements *CapabilityFutureRequirements `json:"future_requirements"`
	// The identifier for the capability.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether the capability has been requested.
	Requested bool `json:"requested"`
	// Time at which the capability was requested. Measured in seconds since the Unix epoch.
	RequestedAt  int64                   `json:"requested_at"`
	Requirements *CapabilityRequirements `json:"requirements"`
	// The status of the capability. Can be `active`, `inactive`, `pending`, or `unrequested`.
	Status CapabilityStatus `json:"status"`
}

// CapabilityList is a list of Capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data"`
}
