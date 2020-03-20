package stripe

import "encoding/json"

// IssuingCardBrand is the list of allowed values for an issuing card's brand.
type IssuingCardBrand string

// List of values that IssuingCardBrand can take.
const (
	issuingCardBrandAmex       IssuingCardBrand = "amex"
	issuingCardBrandDiscover   IssuingCardBrand = "Discover"
	issuingCardBrandDinersClub IssuingCardBrand = "Diners Club"
	issuingCardBrandJCB        IssuingCardBrand = "JCB"
	issuingCardBrandMasterCard IssuingCardBrand = "MasterCard"
	issuingCardBrandUnknown    IssuingCardBrand = "Unknown"
	issuingCardBrandUnionPay   IssuingCardBrand = "UnionPay"
	issuingCardBrandVisa       IssuingCardBrand = "Visa"
)

// IssuingCardPINStatus is the list of possible values for the status field of a Card PIN.
type IssuingCardPINStatus string

// List of values that IssuingCardPINStatus can take.
const (
	IssuingCardPINStatusActive  IssuingCardPINStatus = "active"
	IssuingCardPINStatusBlocked IssuingCardPINStatus = "blocked"
)

// IssuingCardReplacementReason is the list of possible values for the replacement reason on an
// issuing card.
type IssuingCardReplacementReason string

// List of values that IssuingCardReplacementReason can take.
const (
	IssuingCardReplacementReasonDamage     IssuingCardReplacementReason = "damage"
	IssuingCardReplacementReasonExpiration IssuingCardReplacementReason = "expiration"
	IssuingCardReplacementReasonLoss       IssuingCardReplacementReason = "loss"
	IssuingCardReplacementReasonTheft      IssuingCardReplacementReason = "theft"
)

// IssuingCardShippingCarrier is the carrier used for shipping an issuing card.
type IssuingCardShippingCarrier string

// List of values that IssuingCardShippingCarrier can take.
const (
	IssuingCardShippingCarrierFEDEX IssuingCardShippingCarrier = "fedex"
	IssuingCardShippingCarrierUSPS  IssuingCardShippingCarrier = "usps"
)

// IssuingCardShippingService is the shipment service for a card.
type IssuingCardShippingService string

// List of values that IssuingCardShippingService can take.
const (
	IssuingCardShippingServiceExpress   IssuingCardShippingService = "express"
	IssuingCardShippingServiceOvernight IssuingCardShippingService = "overnight"
	IssuingCardShippingServiceStandard  IssuingCardShippingService = "standard"
)

// IssuingCardShippingStatus is the list of possible values for the shipping status
// on an issuing card.
type IssuingCardShippingStatus string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardShippingStatusCanceled  IssuingCardShippingStatus = "canceled"
	IssuingCardShippingStatusDelivered IssuingCardShippingStatus = "delivered"
	IssuingCardShippingStatusFailure   IssuingCardShippingStatus = "failure"
	IssuingCardShippingStatusPending   IssuingCardShippingStatus = "pending"
	IssuingCardShippingStatusReturned  IssuingCardShippingStatus = "returned"
	IssuingCardShippingStatusShipped   IssuingCardShippingStatus = "shipped"
)

// IssuingCardShippingType is the list of possible values for the shipping type
// on an issuing card.
type IssuingCardShippingType string

// List of values that IssuingCardShippingType can take.
const (
	IssuingCardShippingTypeBulk       IssuingCardShippingType = "bulk"
	IssuingCardShippingTypeIndividual IssuingCardShippingType = "individual"
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

// IssuingCardAuthorizationControlsSpendingLimitInterval is the list of possible values for
// the interval of a given spending limit on an issuing card.
type IssuingCardAuthorizationControlsSpendingLimitInterval string

// List of values that IssuingCardAuthorizationControlsSpendingLimitInterval can take.
const (
	IssuingCardAuthorizationControlsSpendingLimitIntervalAllTime          IssuingCardAuthorizationControlsSpendingLimitInterval = "all_time"
	IssuingCardAuthorizationControlsSpendingLimitIntervalDaily            IssuingCardAuthorizationControlsSpendingLimitInterval = "daily"
	IssuingCardAuthorizationControlsSpendingLimitIntervalMonthly          IssuingCardAuthorizationControlsSpendingLimitInterval = "monthly"
	IssuingCardAuthorizationControlsSpendingLimitIntervalPerAuthorization IssuingCardAuthorizationControlsSpendingLimitInterval = "per_authorization"
	IssuingCardAuthorizationControlsSpendingLimitIntervalWeekly           IssuingCardAuthorizationControlsSpendingLimitInterval = "weekly"
	IssuingCardAuthorizationControlsSpendingLimitIntervalYearly           IssuingCardAuthorizationControlsSpendingLimitInterval = "yearly"
)

// IssuingCardAuthorizationControlsSpendingLimitParams is the set of parameters that can be used for
// a spending limit associated with a given issuing card.
type IssuingCardAuthorizationControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardAuthorizationControlsParams  is the set of parameters that can be used for
// a spending limit associated with a given issuing card.
type IssuingCardAuthorizationControlsParams struct {
	AllowedCategories []*string                                              `form:"allowed_categories"`
	BlockedCategories []*string                                              `form:"blocked_categories"`
	MaxApprovals      *int64                                                 `form:"max_approvals"`
	SpendingLimits    []*IssuingCardAuthorizationControlsSpendingLimitParams `form:"spending_limits"`
}

// IssuingCardShippingParams is the set of parameters that can be used for the shipping parameter.
type IssuingCardShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    string         `form:"name"`
	Service *string        `form:"service"`
	Type    *string        `form:"type"`
}

// IssuingCardParams is the set of parameters that can be used when creating or updating an issuing card.
type IssuingCardParams struct {
	Params                `form:"*"`
	AuthorizationControls *IssuingCardAuthorizationControlsParams `form:"authorization_controls"`
	Cardholder            *string                                 `form:"cardholder"`
	Currency              *string                                 `form:"currency"`
	ReplacementFor        *string                                 `form:"replacement_for"`
	ReplacementReason     *string                                 `form:"replacement_reason"`
	Status                *string                                 `form:"status"`
	Shipping              *IssuingCardShippingParams              `form:"shipping"`
	Type                  *string                                 `form:"type"`
}

// IssuingCardListParams is the set of parameters that can be used when listing issuing cards.
type IssuingCardListParams struct {
	ListParams   `form:"*"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	ExpMonth     *int64            `form:"exp_month"`
	ExpYear      *int64            `form:"exp_year"`
	Name         *string           `form:"name"`
	Last4        *string           `form:"last4"`
	Source       *string           `form:"source"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
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

// IssuingCardAuthorizationControlsSpendingLimit is the resource representing spending limits
// associated with an issuing card.
type IssuingCardAuthorizationControlsSpendingLimit struct {
	Amount     int64                                                 `json:"amount"`
	Categories []string                                              `json:"categories"`
	Interval   IssuingCardAuthorizationControlsSpendingLimitInterval `json:"interval"`
}

// IssuingCardAuthorizationControls is the resource representing authorization controls on an issuing card.
type IssuingCardAuthorizationControls struct {
	AllowedCategories      []string                                         `json:"allowed_categories"`
	BlockedCategories      []string                                         `json:"blocked_categories"`
	Currency               Currency                                         `json:"currency"`
	MaxApprovals           int64                                            `json:"max_approvals"`
	SpendingLimits         []*IssuingCardAuthorizationControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                         `json:"spending_limits_currency"`
}

// IssuingCardPIN contains data about the Card's PIN.
type IssuingCardPIN struct {
	Status IssuingCardPINStatus `json:"status"`
}

// IssuingCardShipping is the resource representing shipping on an issuing card.
type IssuingCardShipping struct {
	Address        *Address                   `json:"address"`
	Carrier        IssuingCardShippingCarrier `json:"carrier"`
	ETA            int64                      `json:"eta"`
	Name           string                     `json:"name"`
	Service        IssuingCardShippingService `json:"service"`
	Status         IssuingCardShippingStatus  `json:"status"`
	TrackingNumber string                     `json:"tracking_number"`
	TrackingURL    string                     `json:"tracking_url"`
	Type           IssuingCardShippingType    `json:"type"`
}

// IssuingCard is the resource representing a Stripe issuing card.
type IssuingCard struct {
	AuthorizationControls *IssuingCardAuthorizationControls `json:"authorization_controls"`
	Brand                 string                            `json:"brand"`
	Cardholder            *IssuingCardholder                `json:"cardholder"`
	Created               int64                             `json:"created"`
	Currency              Currency                          `json:"currency"`
	ExpMonth              int64                             `json:"exp_month"`
	ExpYear               int64                             `json:"exp_year"`
	ID                    string                            `json:"id"`
	Last4                 string                            `json:"last4"`
	Livemode              bool                              `json:"livemode"`
	Metadata              map[string]string                 `json:"metadata"`
	Name                  string                            `json:"name"`
	Object                string                            `json:"object"`
	PIN                   *IssuingCardPIN                   `json:"pin"`
	ReplacedBy            *IssuingCard                      `json:"replaced_by"`
	ReplacementFor        *IssuingCard                      `json:"replacement_for"`
	ReplacementReason     IssuingCardReplacementReason      `json:"replacement_reason"`
	Shipping              *IssuingCardShipping              `json:"shipping"`
	Status                IssuingCardStatus                 `json:"status"`
	Type                  IssuingCardType                   `json:"type"`
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
