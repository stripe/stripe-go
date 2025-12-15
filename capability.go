//
//
// File generated from our OpenAPI spec
//
//

package stripe

// This is typed as an enum for consistency with `requirements.disabled_reason`, but it safe to assume `future_requirements.disabled_reason` is null because fields in `future_requirements` will never disable the account.
type CapabilityFutureRequirementsDisabledReason string

// List of values that CapabilityFutureRequirementsDisabledReason can take
const (
	CapabilityFutureRequirementsDisabledReasonOther                       CapabilityFutureRequirementsDisabledReason = "other"
	CapabilityFutureRequirementsDisabledReasonPausedInactivity            CapabilityFutureRequirementsDisabledReason = "paused.inactivity"
	CapabilityFutureRequirementsDisabledReasonPendingOnboarding           CapabilityFutureRequirementsDisabledReason = "pending.onboarding"
	CapabilityFutureRequirementsDisabledReasonPendingReview               CapabilityFutureRequirementsDisabledReason = "pending.review"
	CapabilityFutureRequirementsDisabledReasonPlatformDisabled            CapabilityFutureRequirementsDisabledReason = "platform_disabled"
	CapabilityFutureRequirementsDisabledReasonPlatformPaused              CapabilityFutureRequirementsDisabledReason = "platform_paused"
	CapabilityFutureRequirementsDisabledReasonRejectedInactivity          CapabilityFutureRequirementsDisabledReason = "rejected.inactivity"
	CapabilityFutureRequirementsDisabledReasonRejectedOther               CapabilityFutureRequirementsDisabledReason = "rejected.other"
	CapabilityFutureRequirementsDisabledReasonRejectedUnsupportedBusiness CapabilityFutureRequirementsDisabledReason = "rejected.unsupported_business"
	CapabilityFutureRequirementsDisabledReasonRequirementsFieldsNeeded    CapabilityFutureRequirementsDisabledReason = "requirements.fields_needed"
)

// Description of why the capability is disabled. [Learn more about handling verification issues](https://docs.stripe.com/connect/handling-api-verification).
type CapabilityDisabledReason string

// List of values that CapabilityDisabledReason can take
const (
	CapabilityDisabledReasonOther                       CapabilityDisabledReason = "other"
	CapabilityDisabledReasonPausedInactivity            CapabilityDisabledReason = "paused.inactivity"
	CapabilityDisabledReasonPendingOnboarding           CapabilityDisabledReason = "pending.onboarding"
	CapabilityDisabledReasonPendingReview               CapabilityDisabledReason = "pending.review"
	CapabilityDisabledReasonPlatformDisabled            CapabilityDisabledReason = "platform_disabled"
	CapabilityDisabledReasonPlatformPaused              CapabilityDisabledReason = "platform_paused"
	CapabilityDisabledReasonRejectedInactivity          CapabilityDisabledReason = "rejected.inactivity"
	CapabilityDisabledReasonRejectedOther               CapabilityDisabledReason = "rejected.other"
	CapabilityDisabledReasonRejectedUnsupportedBusiness CapabilityDisabledReason = "rejected.unsupported_business"
	CapabilityDisabledReasonRequirementsFieldsNeeded    CapabilityDisabledReason = "requirements.fields_needed"
)

// The status of the capability.
type CapabilityStatus string

// List of values that CapabilityStatus can take
const (
	CapabilityStatusActive      CapabilityStatus = "active"
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
	// If a capability isn't permanent, you can remove it from the account by passing false. Some capabilities are permanent after they've been requested. Attempting to remove a permanent capability returns an error.
	Requested *bool `form:"requested"`
}

// AddExpand appends a new field to expand.
func (p *CapabilityParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves information about the specified Account Capability.
type CapabilityRetrieveParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapabilityRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates an existing Account Capability. Request or remove a capability by updating its requested parameter.
type CapabilityUpdateParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// To request a new capability for an account, pass true. There can be a delay before the requested capability becomes active. If the capability has any activation requirements, the response includes them in the `requirements` arrays.
	//
	// If a capability isn't permanent, you can remove it from the account by passing false. Some capabilities are permanent after they've been requested. Attempting to remove a permanent capability returns an error.
	Requested *bool `form:"requested"`
}

// AddExpand appends a new field to expand.
func (p *CapabilityUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Fields that are due and can be resolved by providing the corresponding alternative fields instead. Multiple alternatives can reference the same `original_fields_due`. When this happens, any of these alternatives can serve as a pathway for attempting to resolve the fields. Additionally, providing `original_fields_due` again also serves as a pathway for attempting to resolve the fields.
type CapabilityFutureRequirementsAlternative struct {
	// Fields that can be provided to resolve all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be resolved by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Details about validation and verification failures for `due` requirements that must be resolved.
type CapabilityFutureRequirementsError struct {
	// The code for the type of error.
	Code string `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}
type CapabilityFutureRequirements struct {
	// Fields that are due and can be resolved by providing the corresponding alternative fields instead. Multiple alternatives can reference the same `original_fields_due`. When this happens, any of these alternatives can serve as a pathway for attempting to resolve the fields. Additionally, providing `original_fields_due` again also serves as a pathway for attempting to resolve the fields.
	Alternatives []*CapabilityFutureRequirementsAlternative `json:"alternatives"`
	// Date on which `future_requirements` becomes the main `requirements` hash and `future_requirements` becomes empty. After the transition, `currently_due` requirements may immediately become `past_due`, but the account may also be given a grace period depending on the capability's enablement state prior to transitioning.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be resolved to keep the capability enabled. If not resolved by `future_requirements[current_deadline]`, these fields will transition to the main `requirements` hash.
	CurrentlyDue []string `json:"currently_due"`
	// This is typed as an enum for consistency with `requirements.disabled_reason`, but it safe to assume `future_requirements.disabled_reason` is null because fields in `future_requirements` will never disable the account.
	DisabledReason CapabilityFutureRequirementsDisabledReason `json:"disabled_reason"`
	// Details about validation and verification failures for `due` requirements that must be resolved.
	Errors []*CapabilityFutureRequirementsError `json:"errors"`
	// Fields you must collect when all thresholds are reached. As they become required, they appear in `currently_due` as well.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that haven't been resolved by `requirements.current_deadline`. These fields need to be resolved to enable the capability on the account. `future_requirements.past_due` is a subset of `requirements.past_due`.
	PastDue []string `json:"past_due"`
	// Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.
	PendingVerification []string `json:"pending_verification"`
}

// Fields that are due and can be resolved by providing the corresponding alternative fields instead. Multiple alternatives can reference the same `original_fields_due`. When this happens, any of these alternatives can serve as a pathway for attempting to resolve the fields. Additionally, providing `original_fields_due` again also serves as a pathway for attempting to resolve the fields.
type CapabilityRequirementsAlternative struct {
	// Fields that can be provided to resolve all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be resolved by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}
type CapabilityRequirements struct {
	// Fields that are due and can be resolved by providing the corresponding alternative fields instead. Multiple alternatives can reference the same `original_fields_due`. When this happens, any of these alternatives can serve as a pathway for attempting to resolve the fields. Additionally, providing `original_fields_due` again also serves as a pathway for attempting to resolve the fields.
	Alternatives []*CapabilityRequirementsAlternative `json:"alternatives"`
	// The date by which all required account information must be both submitted and verified. This includes fields listed in `currently_due` as well as those in `pending_verification`. If any required information is missing or unverified by this date, the account may be disabled. Note that `current_deadline` may change if additional `currently_due` requirements are requested.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be resolved to keep the capability enabled. If not resolved by `current_deadline`, these fields will appear in `past_due` as well, and the capability is disabled.
	CurrentlyDue []string `json:"currently_due"`
	// Description of why the capability is disabled. [Learn more about handling verification issues](https://docs.stripe.com/connect/handling-api-verification).
	DisabledReason CapabilityDisabledReason `json:"disabled_reason"`
	// Details about validation and verification failures for `due` requirements that must be resolved.
	Errors []*AccountRequirementsError `json:"errors"`
	// Fields you must collect when all thresholds are reached. As they become required, they appear in `currently_due` as well, and `current_deadline` becomes set.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that haven't been resolved by `current_deadline`. These fields need to be resolved to enable the capability on the account.
	PastDue []string `json:"past_due"`
	// Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.
	PendingVerification []string `json:"pending_verification"`
}

// This is an object representing a capability for a Stripe account.
//
// Related guide: [Account capabilities](https://docs.stripe.com/connect/account-capabilities)
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
	// The status of the capability.
	Status CapabilityStatus `json:"status"`
}

// CapabilityList is a list of Capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data"`
}
