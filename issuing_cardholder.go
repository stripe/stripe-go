package stripe

import "encoding/json"

// IssuingCardholderRequirementsDisabledReason is the possible values for the disabled reason on an
// issuing cardholder.
type IssuingCardholderRequirementsDisabledReason string

// List of values that IssuingCardholderRequirementsDisabledReason can take.
const (
	IssuingCardholderRequirementsDisabledReasonListed         IssuingCardholderRequirementsDisabledReason = "listed"
	IssuingCardholderRequirementsDisabledReasonRejectedListed IssuingCardholderRequirementsDisabledReason = "rejected.listed"
	IssuingCardholderRequirementsDisabledReasonUnderReview    IssuingCardholderRequirementsDisabledReason = "under_review"
)

// IssuingCardholderSpendingControlsSpendingLimitInterval is the list of possible values for the interval
// for a spending limit on an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardholderSpendingControlsSpendingLimitIntervalAllTime          IssuingCardholderSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardholderSpendingControlsSpendingLimitIntervalDaily            IssuingCardholderSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardholderSpendingControlsSpendingLimitIntervalMonthly          IssuingCardholderSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardholderSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardholderSpendingControlsSpendingLimitIntervalWeekly           IssuingCardholderSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalYearly           IssuingCardholderSpendingControlsSpendingLimitInterval = "yearly"
)

// IssuingCardholderStatus is the possible values for status on an issuing cardholder.
type IssuingCardholderStatus string

// List of values that IssuingCardholderStatus can take.
const (
	IssuingCardholderStatusActive   IssuingCardholderStatus = "active"
	IssuingCardholderStatusInactive IssuingCardholderStatus = "inactive"

	// This value is deprecated
	IssuingCardholderStatusPending IssuingCardholderStatus = "pending"
)

// IssuingCardholderType is the type of an issuing cardholder.
type IssuingCardholderType string

// List of values that IssuingCardholderType can take.
const (
	IssuingCardholderTypeCompany    IssuingCardholderType = "company"
	IssuingCardholderTypeIndividual IssuingCardholderType = "individual"

	// This value is deprecated. Use IssuingCardholderTypeCompany instead
	IssuingCardholderTypeBusinessEntity IssuingCardholderType = "business_entity"
)

// IssuingBillingParams is the set of parameters that can be used for billing with the Issuing APIs.
type IssuingBillingParams struct {
	Address *AddressParams `form:"address"`

	// This parameter is deprecated
	Name *string `form:"name"`
}

// IssuingCardholderCompanyParams represents additional information about a
// `business_entity` cardholder.
type IssuingCardholderCompanyParams struct {
	TaxID *string `form:"tax_id"`
}

// IssuingCardholderIndividualVerificationDocumentParams represents an
// identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// IssuingCardholderIndividualVerificationParams represents government-issued ID
// document for this cardholder.
type IssuingCardholderIndividualVerificationParams struct {
	Document *IssuingCardholderIndividualVerificationDocumentParams `form:"document"`
}

// IssuingCardholderIndividualDOBParams represents the date of birth of the
// cardholder individual.
type IssuingCardholderIndividualDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// IssuingCardholderIndividualParams represents additional information about an
// `individual` cardholder.
type IssuingCardholderIndividualParams struct {
	DOB          *IssuingCardholderIndividualDOBParams          `form:"dob"`
	FirstName    *string                                        `form:"first_name"`
	LastName     *string                                        `form:"last_name"`
	Verification *IssuingCardholderIndividualVerificationParams `form:"verification"`
}

// IssuingCardholderSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardholderSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing cardholder
type IssuingCardholderSpendingControlsParams struct {
	AllowedCategories      []*string                                               `form:"allowed_categories"`
	BlockedCategories      []*string                                               `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                                 `form:"spending_limits_currency"`
}

// IssuingCardholderParams is the set of parameters that can be used when creating or updating an issuing cardholder.
type IssuingCardholderParams struct {
	Params           `form:"*"`
	Billing          *IssuingBillingParams                    `form:"billing"`
	Company          *IssuingCardholderCompanyParams          `form:"company"`
	Email            *string                                  `form:"email"`
	Individual       *IssuingCardholderIndividualParams       `form:"individual"`
	Name             *string                                  `form:"name"`
	PhoneNumber      *string                                  `form:"phone_number"`
	SpendingControls *IssuingCardholderSpendingControlsParams `form:"spending_controls"`
	Status           *string                                  `form:"status"`
	Type             *string                                  `form:"type"`

	// The following parameters are deprecated
	IsDefault             *bool                        `form:"is_default"`
	AuthorizationControls *AuthorizationControlsParams `form:"authorization_controls"`
}

// IssuingCardholderListParams is the set of parameters that can be used when listing issuing cardholders.
type IssuingCardholderListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
	PhoneNumber  *string           `form:"phone_number"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`

	// The property is considered deprecated.
	// TODO remove in the next major version
	IsDefault *bool `form:"is_default"`
}

// IssuingBilling is the resource representing the billing hash with the Issuing APIs.
type IssuingBilling struct {
	Address *Address `json:"address"`

	// This property is deprecated
	Name string `json:"name"`
}

// IssuingCardholderRequirements contains the verification requirements for the cardholder.
type IssuingCardholderRequirements struct {
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason"`
	PastDue        []string                                    `json:"past_due"`
}

// IssuingCardholderIndividualVerificationDocument represents an identifying
// document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocument struct {
	Back  *File `json:"back"`
	Front *File `json:"front"`
}

// IssuingCardholderIndividualVerification represents the Government-issued ID
// document for this cardholder
type IssuingCardholderIndividualVerification struct {
	Document *IssuingCardholderIndividualVerificationDocument `json:"document"`
}

// IssuingCardholderIndividualDOB represents the date of birth of the issuing card hoder
// individual.
type IssuingCardholderIndividualDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// IssuingCardholderIndividual represents additional information about an
// individual cardholder.
type IssuingCardholderIndividual struct {
	DOB          *IssuingCardholderIndividualDOB          `json:"dob"`
	FirstName    string                                   `json:"first_name"`
	LastName     string                                   `json:"last_name"`
	Verification *IssuingCardholderIndividualVerification `json:"verification"`
}

// IssuingCardholderCompany represents additional information about a
// business_entity cardholder.
type IssuingCardholderCompany struct {
	TaxIDProvided bool `json:"tax_id_provided"`
}

// IssuingCardholderSpendingControlsSpendingLimit is the resource representing a spending limit
// for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimit struct {
	Amount     int64                                                  `json:"amount"`
	Categories []string                                               `json:"categories"`
	Interval   IssuingCardholderSpendingControlsSpendingLimitInterval `json:"interval"`
}

// IssuingCardholderSpendingControls is the resource representing spending controls
// for an issuing cardholder.
type IssuingCardholderSpendingControls struct {
	AllowedCategories      []string                                          `json:"allowed_categories"`
	BlockedCategories      []string                                          `json:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                          `json:"spending_limits_currency"`
}

// IssuingCardholder is the resource representing a Stripe issuing cardholder.
type IssuingCardholder struct {
	Billing          *IssuingBilling                    `json:"billing"`
	Company          *IssuingCardholderCompany          `json:"company"`
	Created          int64                              `json:"created"`
	Email            string                             `json:"email"`
	ID               string                             `json:"id"`
	Individual       *IssuingCardholderIndividual       `json:"individual"`
	Livemode         bool                               `json:"livemode"`
	Metadata         map[string]string                  `json:"metadata"`
	Name             string                             `json:"name"`
	Object           string                             `json:"object"`
	PhoneNumber      string                             `json:"phone_number"`
	Requirements     *IssuingCardholderRequirements     `json:"requirements"`
	SpendingControls *IssuingCardholderSpendingControls `json:"spending_controls"`
	Status           IssuingCardholderStatus            `json:"status"`
	Type             IssuingCardholderType              `json:"type"`

	// The following property is deprecated, use SpendingControls instead
	AuthorizationControls *IssuingCardAuthorizationControls `json:"authorization_controls"`
}

// IssuingCardholderList is a list of issuing cardholders as retrieved from a list endpoint.
type IssuingCardholderList struct {
	ListMeta
	Data []*IssuingCardholder `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCardholder.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCardholder) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCardholder IssuingCardholder
	var v issuingCardholder
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCardholder(v)
	return nil
}
