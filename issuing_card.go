package stripe

import "encoding/json"

// IssuingCardPINStatus is the list of possible values for the status field of a Card PIN.
type IssuingCardPINStatus string

// List of values that IssuingCardPINStatus can take.
const (
	IssuingCardPINStatusActive  IssuingCardPINStatus = "active"
	IssuingCardPINStatusBlocked IssuingCardPINStatus = "blocked"
)

// IssuingCardCancellationReason is the list of possible values for the cancellation reason
// on an issuing card.
type IssuingCardCancellationReason string

// List of values that IssuingCardReplacementReason can take.
const (
	IssuingCardCancellationReasonLost   IssuingCardCancellationReason = "lost"
	IssuingCardCancellationReasonStolen IssuingCardCancellationReason = "stolen"
)

// IssuingCardReplacementReason is the list of possible values for the replacement reason on an
// issuing card.
type IssuingCardReplacementReason string

// List of values that IssuingCardReplacementReason can take.
const (
	IssuingCardReplacementReasonDamaged IssuingCardReplacementReason = "damaged"
	IssuingCardReplacementReasonExpired IssuingCardReplacementReason = "expired"
	IssuingCardReplacementReasonLost    IssuingCardReplacementReason = "lost"
	IssuingCardReplacementReasonStolen  IssuingCardReplacementReason = "stolen"

	// The following values are deprecated and will be removed in the next major version.
	IssuingCardReplacementReasonDamage     IssuingCardReplacementReason = "damage"
	IssuingCardReplacementReasonExpiration IssuingCardReplacementReason = "expiration"
	IssuingCardReplacementReasonLoss       IssuingCardReplacementReason = "loss"
	IssuingCardReplacementReasonTheft      IssuingCardReplacementReason = "theft"
)

// IssuingCardShippingStatus is the list of possible values for the shipping status
// on an issuing card.
type IssuingCardShippingStatus string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardShippingTypeDelivered IssuingCardShippingStatus = "delivered"
	IssuingCardShippingTypeFailure   IssuingCardShippingStatus = "failure"
	IssuingCardShippingTypePending   IssuingCardShippingStatus = "pending"
	IssuingCardShippingTypeReturned  IssuingCardShippingStatus = "returned"
	IssuingCardShippingTypeShipped   IssuingCardShippingStatus = "shipped"
)

// IssuingCardShippingService is the shipment service for a card.
type IssuingCardShippingService string

// List of values that IssuingCardShippingService can take.
const (
	IssuingCardShippingServiceExpress  IssuingCardShippingService = "express"
	IssuingCardShippingServicePriority IssuingCardShippingService = "priority"
	IssuingCardShippingServiceStandard IssuingCardShippingService = "standard"

	// The following value is deprecated, use IssuingCardShippingServicePriority instead
	IssuingCardShippingServiceOvernight IssuingCardShippingService = "overnight"
)

// IssuingCardShippingSpeed is the shipment speed for a card.
// This is deprecated, use IssuingCardShippingService instead
type IssuingCardShippingSpeed string

// List of values that IssuingCardShippingSpeed can take
const (
	IssuingCardShippingSpeedExpress   IssuingCardShippingSpeed = "express"
	IssuingCardShippingSpeedOvernight IssuingCardShippingSpeed = "overnight"
	IssuingCardShippingSpeedStandard  IssuingCardShippingSpeed = "standard"
)

// IssuingCardShippingType is the list of possible values for the shipping type
// on an issuing card.
type IssuingCardShippingType string

// List of values that IssuingCardShippingType can take.
const (
	IssuingCardShippingTypeBulk       IssuingCardShippingType = "bulk"
	IssuingCardShippingTypeIndividual IssuingCardShippingType = "individual"
)

// IssuingCardSpendingControlsSpendingLimitInterval is the list of possible values for the interval
// for a spending limit on an issuing card.
type IssuingCardSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardSpendingControlsSpendingLimitIntervalAllTime          IssuingCardSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardSpendingControlsSpendingLimitIntervalDaily            IssuingCardSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardSpendingControlsSpendingLimitIntervalMonthly          IssuingCardSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardSpendingControlsSpendingLimitIntervalWeekly           IssuingCardSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardSpendingControlsSpendingLimitIntervalYearly           IssuingCardSpendingControlsSpendingLimitInterval = "yearly"
)

// IssuingCardStatus is the list of possible values for status on an issuing card.
type IssuingCardStatus string

// List of values that IssuingCardStatus can take.
const (
	IssuingCardStatusActive   IssuingCardStatus = "active"
	IssuingCardStatusCanceled IssuingCardStatus = "canceled"
	IssuingCardStatusInactive IssuingCardStatus = "inactive"
	IssuingCardStatusPending  IssuingCardStatus = "pending"
)

// IssuingCardType is the type of an issuing card.
type IssuingCardType string

// List of values that IssuingCardType can take.
const (
	IssuingCardTypePhysical IssuingCardType = "physical"
	IssuingCardTypeVirtual  IssuingCardType = "virtual"
)

// IssuingSpendingLimitInterval is the list of possible values for the interval of a given
// spending limit on an issuing card or cardholder.
// This is deprecated, use IssuingCardSpendingControlsSpendingLimitInterval instead
type IssuingSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingSpendingLimitIntervalAllTime          IssuingSpendingLimitInterval = "all_time"
	IssuingSpendingLimitIntervalDaily            IssuingSpendingLimitInterval = "daily"
	IssuingSpendingLimitIntervalMonthly          IssuingSpendingLimitInterval = "monthly"
	IssuingSpendingLimitIntervalPerAuthorization IssuingSpendingLimitInterval = "per_authorization"
	IssuingSpendingLimitIntervalWeekly           IssuingSpendingLimitInterval = "weekly"
	IssuingSpendingLimitIntervalYearly           IssuingSpendingLimitInterval = "yearly"
)

// IssuingAuthorizationControlsSpendingLimitsParams is the set of parameters that can be used for
// the spending limits associated with a given issuing card or cardholder.
// This is deprecated and will be removed in the next major version.
type IssuingAuthorizationControlsSpendingLimitsParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// AuthorizationControlsParams is the set of parameters that can be used for the shipping parameter.
// This is deprecated and will be removed in the next major version.
type AuthorizationControlsParams struct {
	AllowedCategories []*string                                           `form:"allowed_categories"`
	BlockedCategories []*string                                           `form:"blocked_categories"`
	MaxApprovals      *int64                                              `form:"max_approvals"`
	SpendingLimits    []*IssuingAuthorizationControlsSpendingLimitsParams `form:"spending_limits"`

	// The following parameter only applies to Cardholder
	SpendingLimitsCurrency *string `form:"spending_limits_currency"`

	// The following parameter is deprecated
	MaxAmount *int64 `form:"max_amount"`
}

// IssuingCardShippingParams is the set of parameters that can be used for the shipping parameter.
type IssuingCardShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    string         `form:"name"`
	Service *string        `form:"service"`
	Type    *string        `form:"type"`

	// This parameter is deprecated. Use Service instead.
	// TODO remove in the next major version
	Speed *string `form:"speed"`
}

// IssuingCardSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing card.
type IssuingCardSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing card
type IssuingCardSpendingControlsParams struct {
	AllowedCategories      []*string                                         `form:"allowed_categories"`
	BlockedCategories      []*string                                         `form:"blocked_categories"`
	MaxApprovals           *int64                                            `form:"max_approvals"`
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                           `form:"spending_limits_currency"`
}

// IssuingCardParams is the set of parameters that can be used when creating or updating an issuing card.
type IssuingCardParams struct {
	Params             `form:"*"`
	Billing            *IssuingBillingParams              `form:"billing"`
	CancellationReason *string                            `form:"cancellation_reason"`
	Cardholder         *string                            `form:"cardholder"`
	Currency           *string                            `form:"currency"`
	ReplacementFor     *string                            `form:"replacement_for"`
	ReplacementReason  *string                            `form:"replacement_reason"`
	SpendingControls   *IssuingCardSpendingControlsParams `form:"spending_controls"`
	Status             *string                            `form:"status"`
	Shipping           *IssuingCardShippingParams         `form:"shipping"`
	Type               *string                            `form:"type"`

	// The following parameter is deprecated, use SpendingControls instead.
	AuthorizationControls *AuthorizationControlsParams `form:"authorization_controls"`

	// The following parameter is deprecated
	Name *string `form:"name"`
}

// IssuingCardListParams is the set of parameters that can be used when listing issuing cards.
type IssuingCardListParams struct {
	ListParams   `form:"*"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	ExpMonth     *int64            `form:"exp_month"`
	ExpYear      *int64            `form:"exp_year"`
	Last4        *string           `form:"last4"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`

	// The following parameters are deprecated
	Name   *string `form:"name"`
	Source *string `form:"source"`
}

// IssuingCardDetails is the resource representing issuing card details.
type IssuingCardDetails struct {
	Card     *IssuingCard `json:"card"`
	CVC      string       `json:"cvc"`
	ExpMonth *string      `form:"exp_month"`
	ExpYear  *string      `form:"exp_year"`
	Number   string       `json:"number"`
	Object   string       `json:"object"`
}

// IssuingAuthorizationControlsSpendingLimits is the resource representing spending limits
// associated with a card or cardholder.
type IssuingAuthorizationControlsSpendingLimits struct {
	Amount     int64                        `json:"amount"`
	Categories []string                     `json:"categories"`
	Interval   IssuingSpendingLimitInterval `json:"interval"`
}

// IssuingCardAuthorizationControls is the resource representing authorization controls on an issuing card.
// TODO: Add the Cardholder version to "un-share" between Card and Cardholder in the next major version.
type IssuingCardAuthorizationControls struct {
	AllowedCategories      []string                                      `json:"allowed_categories"`
	BlockedCategories      []string                                      `json:"blocked_categories"`
	MaxApprovals           int64                                         `json:"max_approvals"`
	SpendingLimits         []*IssuingAuthorizationControlsSpendingLimits `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                      `json:"spending_limits_currency"`

	// The properties below are deprecated and can only be used for an issuing card.
	// TODO remove in the next major version
	Currency  Currency `json:"currency"`
	MaxAmount int64    `json:"max_amount"`
}

// IssuingCardPIN contains data about the Card's PIN.
type IssuingCardPIN struct {
	Status IssuingCardPINStatus `json:"status"`
}

// IssuingCardShipping is the resource representing shipping on an issuing card.
type IssuingCardShipping struct {
	Address        *Address                   `json:"address"`
	Carrier        string                     `json:"carrier"`
	ETA            int64                      `json:"eta"`
	Name           string                     `json:"name"`
	Phone          string                     `json:"phone"`
	Service        IssuingCardShippingService `json:"service"`
	Status         IssuingCardShippingStatus  `json:"status"`
	TrackingNumber string                     `json:"tracking_number"`
	TrackingURL    string                     `json:"tracking_url"`
	Type           IssuingCardShippingType    `json:"type"`

	// The property is deprecated. Use AddressPostalCodeCheck instead.
	// TODO remove in the next major version
	Speed IssuingCardShippingSpeed `json:"speed"`
}

// IssuingCardSpendingControlsSpendingLimit is the resource representing a spending limit
// for an issuing card.
type IssuingCardSpendingControlsSpendingLimit struct {
	Amount     int64                                            `json:"amount"`
	Categories []string                                         `json:"categories"`
	Interval   IssuingCardSpendingControlsSpendingLimitInterval `json:"interval"`
}

// IssuingCardSpendingControls is the resource representing spending controls
// for an issuing card.
type IssuingCardSpendingControls struct {
	AllowedCategories      []string                                    `json:"allowed_categories"`
	BlockedCategories      []string                                    `json:"blocked_categories"`
	MaxApprovals           int64                                       `json:"max_approvals"`
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                    `json:"spending_limits_currency"`
}

// IssuingCard is the resource representing a Stripe issuing card.
type IssuingCard struct {
	Billing            *IssuingBilling               `json:"billing"`
	Brand              string                        `json:"brand"`
	CancellationReason IssuingCardCancellationReason `json:"cancellation_reason"`
	Cardholder         *IssuingCardholder            `json:"cardholder"`
	Created            int64                         `json:"created"`
	ExpMonth           int64                         `json:"exp_month"`
	ExpYear            int64                         `json:"exp_year"`
	Last4              string                        `json:"last4"`
	ID                 string                        `json:"id"`
	Livemode           bool                          `json:"livemode"`
	Metadata           map[string]string             `json:"metadata"`
	Object             string                        `json:"object"`
	PIN                *IssuingCardPIN               `json:"pin"`
	ReplacedBy         *IssuingCard                  `json:"replaced_by"`
	ReplacementFor     *IssuingCard                  `json:"replacement_for"`
	ReplacementReason  IssuingCardReplacementReason  `json:"replacement_reason"`
	Shipping           *IssuingCardShipping          `json:"shipping"`
	SpendingControls   *IssuingCardSpendingControls  `json:"spending_controls"`
	Status             IssuingCardStatus             `json:"status"`
	Type               IssuingCardType               `json:"type"`

	// The following property is deprecated, use SpendingControls instead.
	AuthorizationControls *IssuingCardAuthorizationControls `json:"authorization_controls"`

	// The following property is deprecated, use Cardholder.Name instead.
	Name string `json:"name"`
}

// IssuingCardList is a list of issuing cards as retrieved from a list endpoint.
type IssuingCardList struct {
	ListMeta
	Data []*IssuingCard `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCard.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCard) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCard IssuingCard
	var v issuingCard
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCard(v)
	return nil
}
