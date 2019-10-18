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

// IssuingCardholderStatus is the possible values for status on an issuing cardholder.
type IssuingCardholderStatus string

// List of values that IssuingCardholderStatus can take.
const (
	IssuingCardholderStatusActive   IssuingCardholderStatus = "active"
	IssuingCardholderStatusInactive IssuingCardholderStatus = "inactive"
	IssuingCardholderStatusPending  IssuingCardholderStatus = "pending"
)

// IssuingCardholderType is the type of an issuing cardholder.
type IssuingCardholderType string

// List of values that IssuingCardholderType can take.
const (
	IssuingCardholderTypeBusinessEntity IssuingCardholderType = "business_entity"
	IssuingCardholderTypeIndividual     IssuingCardholderType = "individual"
)

// IssuingBillingParams is the set of parameters that can be used for billing with the Issuing APIs.
type IssuingBillingParams struct {
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
}

// IssuingCardholderParams is the set of parameters that can be used when creating or updating an issuing cardholder.
type IssuingCardholderParams struct {
	Params                `form:"*"`
	AuthorizationControls *AuthorizationControlsParams `form:"authorization_controls"`
	Billing               *IssuingBillingParams        `form:"billing"`
	Email                 *string                      `form:"email"`
	IsDefault             *bool                        `form:"is_default"`
	Name                  *string                      `form:"name"`
	PhoneNumber           *string                      `form:"phone_number"`
	Status                *string                      `form:"status"`
	Type                  *string                      `form:"type"`
}

// IssuingCardholderListParams is the set of parameters that can be used when listing issuing cardholders.
type IssuingCardholderListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
	IsDefault    *bool             `form:"is_default"`
	PhoneNumber  *string           `form:"phone_number"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
}

// IssuingBilling is the resource representing the billing hash with the Issuing APIs.
type IssuingBilling struct {
	Address *Address `json:"address"`
	Name    string   `json:"name"`
}

// IssuingCardholderRequirements contains the verification requirements for the cardholder.
type IssuingCardholderRequirements struct {
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason"`
	PastDue        []string                                    `json:"past_due"`
}

// IssuingCardholder is the resource representing a Stripe issuing cardholder.
type IssuingCardholder struct {
	AuthorizationControls *IssuingCardAuthorizationControls `json:"authorization_controls"`
	Billing               *IssuingBilling                   `json:"billing"`
	Created               int64                             `json:"created"`
	Email                 string                            `json:"email"`
	ID                    string                            `json:"id"`
	Livemode              bool                              `json:"livemode"`
	Metadata              map[string]string                 `json:"metadata"`
	Name                  string                            `json:"name"`
	Object                string                            `json:"object"`
	PhoneNumber           string                            `json:"phone_number"`
	Requirements          *IssuingCardholderRequirements    `json:"requirements"`
	Status                IssuingCardholderStatus           `json:"status"`
	Type                  IssuingCardholderType             `json:"type"`
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
