package stripe

import "encoding/json"

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
)

// IssuingCardShippingCarrier is the list of possible values for the shipping carrier
// on an issuing card.
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
	IssuingCardShippingServiceExpress  IssuingCardShippingService = "express"
	IssuingCardShippingServicePriority IssuingCardShippingService = "priority"
	IssuingCardShippingServiceStandard IssuingCardShippingService = "standard"
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
)

// IssuingCardType is the type of an issuing card.
type IssuingCardType string

// List of values that IssuingCardType can take.
const (
	IssuingCardTypePhysical IssuingCardType = "physical"
	IssuingCardTypeVirtual  IssuingCardType = "virtual"
)

// IssuingCardShippingParams is the set of parameters that can be used for the shipping parameter.
type IssuingCardShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    string         `form:"name"`
	Service *string        `form:"service"`
	Type    *string        `form:"type"`
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
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                           `form:"spending_limits_currency"`
}

// IssuingCardParams is the set of parameters that can be used when creating or updating an issuing card.
type IssuingCardParams struct {
	Params            `form:"*"`
	Cardholder        *string                            `form:"cardholder"`
	Currency          *string                            `form:"currency"`
	ReplacementFor    *string                            `form:"replacement_for"`
	ReplacementReason *string                            `form:"replacement_reason"`
	SpendingControls  *IssuingCardSpendingControlsParams `form:"spending_controls"`
	Shipping          *IssuingCardShippingParams         `form:"shipping"`
	Status            *string                            `form:"status"`
	Type              *string                            `form:"type"`

	// The following parameter is only supported when updating a card
	CancellationReason *string `form:"cancellation_reason"`
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
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                    `json:"spending_limits_currency"`
}

// IssuingCard is the resource representing a Stripe issuing card.
type IssuingCard struct {
	APIResource
	Brand              string                        `json:"brand"`
	CancellationReason IssuingCardCancellationReason `json:"cancellation_reason"`
	Cardholder         *IssuingCardholder            `json:"cardholder"`
	Created            int64                         `json:"created"`
	Currency           Currency                      `json:"currency"`
	CVC                string                        `json:"cvc"`
	ExpMonth           int64                         `json:"exp_month"`
	ExpYear            int64                         `json:"exp_year"`
	ID                 string                        `json:"id"`
	Last4              string                        `json:"last4"`
	Livemode           bool                          `json:"livemode"`
	Metadata           map[string]string             `json:"metadata"`
	Number             string                        `json:"number"`
	Object             string                        `json:"object"`
	ReplacedBy         *IssuingCard                  `json:"replaced_by"`
	ReplacementFor     *IssuingCard                  `json:"replacement_for"`
	ReplacementReason  IssuingCardReplacementReason  `json:"replacement_reason"`
	Shipping           *IssuingCardShipping          `json:"shipping"`
	SpendingControls   *IssuingCardSpendingControls  `json:"spending_controls"`
	Status             IssuingCardStatus             `json:"status"`
	Type               IssuingCardType               `json:"type"`
}

// IssuingCardList is a list of issuing cards as retrieved from a list endpoint.
type IssuingCardList struct {
	APIResource
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
