//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The identifier for the capability.
type PaymentLocationCapabilityCapability string

// List of values that PaymentLocationCapabilityCapability can take
const (
	PaymentLocationCapabilityCapabilityFRMealVouchersConecsPayments PaymentLocationCapabilityCapability = "fr_meal_vouchers_conecs_payments"
)

// Description of why the capability is disabled.
type PaymentLocationCapabilityRequirementsDisabledReason string

// List of values that PaymentLocationCapabilityRequirementsDisabledReason can take
const (
	PaymentLocationCapabilityRequirementsDisabledReasonAccountCapabilityRequired   PaymentLocationCapabilityRequirementsDisabledReason = "account.capability_required"
	PaymentLocationCapabilityRequirementsDisabledReasonPendingOnboarding           PaymentLocationCapabilityRequirementsDisabledReason = "pending.onboarding"
	PaymentLocationCapabilityRequirementsDisabledReasonPendingReview               PaymentLocationCapabilityRequirementsDisabledReason = "pending.review"
	PaymentLocationCapabilityRequirementsDisabledReasonRejectedOther               PaymentLocationCapabilityRequirementsDisabledReason = "rejected.other"
	PaymentLocationCapabilityRequirementsDisabledReasonRejectedUnsupportedBusiness PaymentLocationCapabilityRequirementsDisabledReason = "rejected.unsupported_business"
	PaymentLocationCapabilityRequirementsDisabledReasonRequirementsFieldsNeeded    PaymentLocationCapabilityRequirementsDisabledReason = "requirements.fields_needed"
)

// The code for the type of error.
type PaymentLocationCapabilityRequirementsErrorCode string

// List of values that PaymentLocationCapabilityRequirementsErrorCode can take
const (
	PaymentLocationCapabilityRequirementsErrorCodeInformationMissing PaymentLocationCapabilityRequirementsErrorCode = "information_missing"
	PaymentLocationCapabilityRequirementsErrorCodeInvalidValueOther  PaymentLocationCapabilityRequirementsErrorCode = "invalid_value_other"
)

// The status of the capability.
type PaymentLocationCapabilityStatus string

// List of values that PaymentLocationCapabilityStatus can take
const (
	PaymentLocationCapabilityStatusActive      PaymentLocationCapabilityStatus = "active"
	PaymentLocationCapabilityStatusInactive    PaymentLocationCapabilityStatus = "inactive"
	PaymentLocationCapabilityStatusPending     PaymentLocationCapabilityStatus = "pending"
	PaymentLocationCapabilityStatusUnrequested PaymentLocationCapabilityStatus = "unrequested"
)

// Returns a list of PaymentLocationCapability objects associated with the location.
type PaymentLocationCapabilityListParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The location for which the capabilities enable functionality.
	Location *string `form:"location" json:"location"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationCapabilityListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves information about the specified Payment Location Capability.
type PaymentLocationCapabilityParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The payment location for which the capability enables functionality.
	Location *string `form:"location" json:"location"`
	// To request a new capability for the location, set this to `true`. You can remove it from the location by passing `false`.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationCapabilityParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves information about the specified Payment Location Capability.
type PaymentLocationCapabilityRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The payment location for which the capability enables functionality.
	Location *string `form:"location" json:"location"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationCapabilityRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates a specified Payment Location Capability. Request or remove a payment location capability by updating its requested parameter.
type PaymentLocationCapabilityUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The location for which the capability enables functionality.
	Location *string `form:"location" json:"location"`
	// To request a new capability for the location, set this to `true`. You can remove it from the location by passing `false`.
	Requested *bool `form:"requested" json:"requested"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationCapabilityUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PaymentLocationCapabilityRequirementsError struct {
	// The code for the type of error.
	Code PaymentLocationCapabilityRequirementsErrorCode `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}
type PaymentLocationCapabilityRequirements struct {
	// Fields that need to be collected to keep the capability enabled.
	CurrentlyDue []string `json:"currently_due"`
	// Description of why the capability is disabled.
	DisabledReason PaymentLocationCapabilityRequirementsDisabledReason `json:"disabled_reason"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*PaymentLocationCapabilityRequirementsError `json:"errors"`
}

// A Payment Location Capability represents a capability for a Stripe account at a Payment Location.
type PaymentLocationCapability struct {
	APIResource
	// The account for which the capability enables functionality.
	Account string `json:"account"`
	// The identifier for the capability.
	Capability PaymentLocationCapabilityCapability `json:"capability"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// The payment location for which the capability enables functionality.
	Location string `json:"location"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether the capability has been requested.
	Requested bool `json:"requested"`
	// Time at which the capability was requested. Measured in seconds since the Unix epoch.
	RequestedAt  int64                                  `json:"requested_at"`
	Requirements *PaymentLocationCapabilityRequirements `json:"requirements"`
	// The status of the capability.
	Status PaymentLocationCapabilityStatus `json:"status"`
}

// PaymentLocationCapabilityList is a list of PaymentLocationCapabilities as retrieved from a list endpoint.
type PaymentLocationCapabilityList struct {
	APIResource
	ListMeta
	Data []*PaymentLocationCapability `json:"data"`
}
