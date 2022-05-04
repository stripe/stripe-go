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
	ListParams `form:"*"`
	// Only return cardholders that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return cardholders that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return cardholders that have the given email address.
	Email *string `form:"email"`
	// Only return cardholders that have the given phone number.
	PhoneNumber *string `form:"phone_number"`
	// Only return cardholders that have the given status. One of `active`, `inactive`, or `blocked`.
	Status *string `form:"status"`
	// Only return cardholders that have the given type. One of `individual` or `company`.
	Type *string `form:"type"`
}

// The cardholder's billing address.
type IssuingCardholderBillingParams struct {
	// The cardholder's billing address.
	Address *AddressParams `form:"address"`
}

// Additional information about a `company` cardholder.
type IssuingCardholderCompanyParams struct {
	// The entity's business ID number.
	TaxID *string `form:"tax_id"`
}

// The date of birth of this cardholder.
type IssuingCardholderIndividualDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// An identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocumentParams struct {
	// The back of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Back *string `form:"back"`
	// The front of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Front *string `form:"front"`
}

// Government-issued ID document for this cardholder.
type IssuingCardholderIndividualVerificationParams struct {
	// An identifying document, either a passport or local ID card.
	Document *IssuingCardholderIndividualVerificationDocumentParams `form:"document"`
}

// Additional information about an `individual` cardholder.
type IssuingCardholderIndividualParams struct {
	// The date of birth of this cardholder.
	DOB *IssuingCardholderIndividualDOBParams `form:"dob"`
	// The first name of this cardholder.
	FirstName *string `form:"first_name"`
	// The last name of this cardholder.
	LastName *string `form:"last_name"`
	// Government-issued ID document for this cardholder.
	Verification *IssuingCardholderIndividualVerificationParams `form:"verification"`
}

// Limit spending with amount-based rules that apply across this cardholder's cards.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	// Maximum amount allowed to spend per interval.
	Amount *int64 `form:"amount"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.
	Categories []*string `form:"categories"`
	// Interval (or event) to which the amount applies.
	Interval *string `form:"interval"`
}

// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardholderSpendingControlsParams struct {
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.
	AllowedCategories []*string `form:"allowed_categories"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.
	BlockedCategories []*string `form:"blocked_categories"`
	// Limit spending with amount-based rules that apply across this cardholder's cards.
	SpendingLimits []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	// Currency of amounts within `spending_limits`. Defaults to your merchant country's currency.
	SpendingLimitsCurrency *string `form:"spending_limits_currency"`
}

// Creates a new Issuing Cardholder object that can be issued cards.
type IssuingCardholderParams struct {
	Params `form:"*"`
	// The cardholder's billing address.
	Billing *IssuingCardholderBillingParams `form:"billing"`
	// Additional information about a `company` cardholder.
	Company *IssuingCardholderCompanyParams `form:"company"`
	// The cardholder's email address.
	Email *string `form:"email"`
	// Additional information about an `individual` cardholder.
	Individual *IssuingCardholderIndividualParams `form:"individual"`
	// The cardholder's name. This will be printed on cards issued to them. The maximum length of this field is 24 characters.
	Name *string `form:"name"`
	// The cardholder's phone number. This is required for all cardholders who will be creating EU cards. See the [3D Secure documentation](https://stripe.com/docs/issuing/3d-secure) for more details.
	PhoneNumber *string `form:"phone_number"`
	// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
	SpendingControls *IssuingCardholderSpendingControlsParams `form:"spending_controls"`
	// Specifies whether to permit authorizations on this cardholder's cards.
	Status *string `form:"status"`
	// One of `individual` or `company`.
	Type *string `form:"type"`
}
type IssuingCardholderBilling struct {
	Address *Address `json:"address"`
}

// Additional information about a `company` cardholder.
type IssuingCardholderCompany struct {
	// Whether the company's business ID number was provided.
	TaxIDProvided bool `json:"tax_id_provided"`
}

// The date of birth of this cardholder.
type IssuingCardholderIndividualDOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}

// An identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocument struct {
	// The back of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Back *File `json:"back"`
	// The front of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Front *File `json:"front"`
}

// Government-issued ID document for this cardholder.
type IssuingCardholderIndividualVerification struct {
	// An identifying document, either a passport or local ID card.
	Document *IssuingCardholderIndividualVerificationDocument `json:"document"`
}

// Additional information about an `individual` cardholder.
type IssuingCardholderIndividual struct {
	// The date of birth of this cardholder.
	DOB *IssuingCardholderIndividualDOB `json:"dob"`
	// The first name of this cardholder.
	FirstName string `json:"first_name"`
	// The last name of this cardholder.
	LastName string `json:"last_name"`
	// Government-issued ID document for this cardholder.
	Verification *IssuingCardholderIndividualVerification `json:"verification"`
}
type IssuingCardholderRequirements struct {
	// If `disabled_reason` is present, all cards will decline authorizations with `cardholder_verification_required` reason.
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason"`
	// Array of fields that need to be collected in order to verify and re-enable the cardholder.
	PastDue []string `json:"past_due"`
}

// Limit spending with amount-based rules that apply across this cardholder's cards.
type IssuingCardholderSpendingControlsSpendingLimit struct {
	// Maximum amount allowed to spend per interval. This amount is in the card's currency and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.
	Categories []string `json:"categories"`
	// Interval (or event) to which the amount applies.
	Interval IssuingCardholderSpendingControlsSpendingLimitInterval `json:"interval"`
}

// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardholderSpendingControls struct {
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.
	AllowedCategories []string `json:"allowed_categories"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.
	BlockedCategories []string `json:"blocked_categories"`
	// Limit spending with amount-based rules that apply across this cardholder's cards.
	SpendingLimits []*IssuingCardholderSpendingControlsSpendingLimit `json:"spending_limits"`
	// Currency of the amounts within `spending_limits`.
	SpendingLimitsCurrency Currency `json:"spending_limits_currency"`
}

// An Issuing `Cardholder` object represents an individual or business entity who is [issued](https://stripe.com/docs/issuing) cards.
//
// Related guide: [How to create a Cardholder](https://stripe.com/docs/issuing/cards#create-cardholder)
type IssuingCardholder struct {
	APIResource
	Billing *IssuingCardholderBilling `json:"billing"`
	// Additional information about a `company` cardholder.
	Company *IssuingCardholderCompany `json:"company"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The cardholder's email address.
	Email string `json:"email"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Additional information about an `individual` cardholder.
	Individual *IssuingCardholderIndividual `json:"individual"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The cardholder's name. This will be printed on cards issued to them.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The cardholder's phone number. This is required for all cardholders who will be creating EU cards. See the [3D Secure documentation](https://stripe.com/docs/issuing/3d-secure#when-is-3d-secure-applied) for more details.
	PhoneNumber  string                         `json:"phone_number"`
	Requirements *IssuingCardholderRequirements `json:"requirements"`
	// Rules that control spending across this cardholder's cards. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
	SpendingControls *IssuingCardholderSpendingControls `json:"spending_controls"`
	// Specifies whether to permit authorizations on this cardholder's cards.
	Status IssuingCardholderStatus `json:"status"`
	// One of `individual` or `company`.
	Type IssuingCardholderType `json:"type"`
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
