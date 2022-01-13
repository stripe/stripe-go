//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// If `disabled_reason` is present, all cards will decline authorizations with `cardholder_verification_required` reason.
type IssuingCardholderRequirementsDisabledReason string

// List of values that IssuingCardholderRequirementsDisabledReason can take
const (
	IssuingCardholderRequirementsDisabledReasonListed         IssuingCardholderRequirementsDisabledReason = "listed"
	IssuingCardholderRequirementsDisabledReasonRejectedListed IssuingCardholderRequirementsDisabledReason = "rejected.listed"
	IssuingCardholderRequirementsDisabledReasonUnderReview    IssuingCardholderRequirementsDisabledReason = "under_review"
)

// Interval (or event) to which the amount applies.
type IssuingCardholderSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardholderSpendingControlsSpendingLimitInterval can take
const (
	IssuingCardholderSpendingControlsSpendingLimitIntervalAllTime          IssuingCardholderSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardholderSpendingControlsSpendingLimitIntervalDaily            IssuingCardholderSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardholderSpendingControlsSpendingLimitIntervalMonthly          IssuingCardholderSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardholderSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardholderSpendingControlsSpendingLimitIntervalWeekly           IssuingCardholderSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalYearly           IssuingCardholderSpendingControlsSpendingLimitInterval = "yearly"
)

// Specifies whether to permit authorizations on this cardholder's cards.
type IssuingCardholderStatus string

// List of values that IssuingCardholderStatus can take
const (
	IssuingCardholderStatusActive   IssuingCardholderStatus = "active"
	IssuingCardholderStatusBlocked  IssuingCardholderStatus = "blocked"
	IssuingCardholderStatusInactive IssuingCardholderStatus = "inactive"
)

// One of `individual` or `company`.
type IssuingCardholderType string

// List of values that IssuingCardholderType can take
const (
	IssuingCardholderTypeCompany    IssuingCardholderType = "company"
	IssuingCardholderTypeIndividual IssuingCardholderType = "individual"
)

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingCardholderListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
	PhoneNumber  *string           `form:"phone_number"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
}

// The cardholder's billing address.
type IssuingCardholderBillingParams struct {
	Address *AddressParams `form:"address"`
}

// Additional information about a `company` cardholder.
type IssuingCardholderCompanyParams struct {
	TaxID *string `form:"tax_id"`
}

// The date of birth of this cardholder.
type IssuingCardholderIndividualDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// An identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// Government-issued ID document for this cardholder.
type IssuingCardholderIndividualVerificationParams struct {
	Document *IssuingCardholderIndividualVerificationDocumentParams `form:"document"`
}

// Additional information about an `individual` cardholder.
type IssuingCardholderIndividualParams struct {
	DOB          *IssuingCardholderIndividualDOBParams          `form:"dob"`
	FirstName    *string                                        `form:"first_name"`
	LastName     *string                                        `form:"last_name"`
	Verification *IssuingCardholderIndividualVerificationParams `form:"verification"`
}

// Limit spending with amount-based rules that apply across this cardholder's cards.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardholderSpendingControlsParams struct {
	AllowedCategories      []*string                                               `form:"allowed_categories"`
	BlockedCategories      []*string                                               `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                                 `form:"spending_limits_currency"`
}

// Creates a new Issuing Cardholder object that can be issued cards.
type IssuingCardholderParams struct {
	Params           `form:"*"`
	Billing          *IssuingCardholderBillingParams          `form:"billing"`
	Company          *IssuingCardholderCompanyParams          `form:"company"`
	Email            *string                                  `form:"email"`
	Individual       *IssuingCardholderIndividualParams       `form:"individual"`
	Name             *string                                  `form:"name"`
	PhoneNumber      *string                                  `form:"phone_number"`
	SpendingControls *IssuingCardholderSpendingControlsParams `form:"spending_controls"`
	Status           *string                                  `form:"status"`
	Type             *string                                  `form:"type"`
}
type IssuingCardholderBilling struct {
	Address *Address `json:"address"`
}

// Additional information about a `company` cardholder.
type IssuingCardholderCompany struct {
	TaxIDProvided bool `json:"tax_id_provided"`
}

// The date of birth of this cardholder.
type IssuingCardholderIndividualDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// An identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocument struct {
	Back  *File `json:"back"`
	Front *File `json:"front"`
}

// Government-issued ID document for this cardholder.
type IssuingCardholderIndividualVerification struct {
	Document *IssuingCardholderIndividualVerificationDocument `json:"document"`
}

// Additional information about an `individual` cardholder.
type IssuingCardholderIndividual struct {
	DOB          *IssuingCardholderIndividualDOB          `json:"dob"`
	FirstName    string                                   `json:"first_name"`
	LastName     string                                   `json:"last_name"`
	Verification *IssuingCardholderIndividualVerification `json:"verification"`
}
type IssuingCardholderRequirements struct {
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason"`
	PastDue        []string                                    `json:"past_due"`
}

// Limit spending with amount-based rules that apply across this cardholder's cards.
type IssuingCardholderSpendingControlsSpendingLimit struct {
	Amount     int64                                                  `json:"amount"`
	Categories []string                                               `json:"categories"`
	Interval   IssuingCardholderSpendingControlsSpendingLimitInterval `json:"interval"`
}

// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardholderSpendingControls struct {
	AllowedCategories      []string                                          `json:"allowed_categories"`
	BlockedCategories      []string                                          `json:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                          `json:"spending_limits_currency"`
}

// An Issuing `Cardholder` object represents an individual or business entity who is [issued](https://stripe.com/docs/issuing) cards.
//
// Related guide: [How to create a Cardholder](https://stripe.com/docs/issuing/cards#create-cardholder)
type IssuingCardholder struct {
	APIResource
	Billing          *IssuingCardholderBilling          `json:"billing"`
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
}

// IssuingCardholderList is a list of Cardholders as retrieved from a list endpoint.
type IssuingCardholderList struct {
	APIResource
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
