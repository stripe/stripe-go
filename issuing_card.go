//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The reason why the card was canceled.
type IssuingCardCancellationReason string

// List of values that IssuingCardCancellationReason can take
const (
	IssuingCardCancellationReasonLost   IssuingCardCancellationReason = "lost"
	IssuingCardCancellationReasonStolen IssuingCardCancellationReason = "stolen"
)

// The reason why the previous card needed to be replaced.
type IssuingCardReplacementReason string

// List of values that IssuingCardReplacementReason can take
const (
	IssuingCardReplacementReasonDamaged IssuingCardReplacementReason = "damaged"
	IssuingCardReplacementReasonExpired IssuingCardReplacementReason = "expired"
	IssuingCardReplacementReasonLost    IssuingCardReplacementReason = "lost"
	IssuingCardReplacementReasonStolen  IssuingCardReplacementReason = "stolen"
)

// The delivery company that shipped a card.
type IssuingCardShippingCarrier string

// List of values that IssuingCardShippingCarrier can take
const (
	IssuingCardShippingCarrierDHL       IssuingCardShippingCarrier = "dhl"
	IssuingCardShippingCarrierFEDEX     IssuingCardShippingCarrier = "fedex"
	IssuingCardShippingCarrierRoyalMail IssuingCardShippingCarrier = "royal_mail"
	IssuingCardShippingCarrierUSPS      IssuingCardShippingCarrier = "usps"
)

// Shipment service, such as `standard` or `express`.
type IssuingCardShippingService string

// List of values that IssuingCardShippingService can take
const (
	IssuingCardShippingServiceExpress  IssuingCardShippingService = "express"
	IssuingCardShippingServicePriority IssuingCardShippingService = "priority"
	IssuingCardShippingServiceStandard IssuingCardShippingService = "standard"
)

// The delivery status of the card.
type IssuingCardShippingStatus string

// List of values that IssuingCardShippingStatus can take
const (
	IssuingCardShippingStatusCanceled  IssuingCardShippingStatus = "canceled"
	IssuingCardShippingStatusDelivered IssuingCardShippingStatus = "delivered"
	IssuingCardShippingStatusFailure   IssuingCardShippingStatus = "failure"
	IssuingCardShippingStatusPending   IssuingCardShippingStatus = "pending"
	IssuingCardShippingStatusReturned  IssuingCardShippingStatus = "returned"
	IssuingCardShippingStatusShipped   IssuingCardShippingStatus = "shipped"
)

// Packaging options.
type IssuingCardShippingType string

// List of values that IssuingCardShippingType can take
const (
	IssuingCardShippingTypeBulk       IssuingCardShippingType = "bulk"
	IssuingCardShippingTypeIndividual IssuingCardShippingType = "individual"
)

// Interval (or event) to which the amount applies.
type IssuingCardSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardSpendingControlsSpendingLimitInterval can take
const (
	IssuingCardSpendingControlsSpendingLimitIntervalAllTime          IssuingCardSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardSpendingControlsSpendingLimitIntervalDaily            IssuingCardSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardSpendingControlsSpendingLimitIntervalMonthly          IssuingCardSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardSpendingControlsSpendingLimitIntervalWeekly           IssuingCardSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardSpendingControlsSpendingLimitIntervalYearly           IssuingCardSpendingControlsSpendingLimitInterval = "yearly"
)

// Whether authorizations can be approved on this card.
type IssuingCardStatus string

// List of values that IssuingCardStatus can take
const (
	IssuingCardStatusActive   IssuingCardStatus = "active"
	IssuingCardStatusCanceled IssuingCardStatus = "canceled"
	IssuingCardStatusInactive IssuingCardStatus = "inactive"
)

// The type of the card.
type IssuingCardType string

// List of values that IssuingCardType can take
const (
	IssuingCardTypePhysical IssuingCardType = "physical"
	IssuingCardTypeVirtual  IssuingCardType = "virtual"
)

// Reason the card is ineligible for Apple Pay
type IssuingCardWalletsApplePayIneligibleReason string

// List of values that IssuingCardWalletsApplePayIneligibleReason can take
const (
	IssuingCardWalletsApplePayIneligibleReasonMissingAgreement         IssuingCardWalletsApplePayIneligibleReason = "missing_agreement"
	IssuingCardWalletsApplePayIneligibleReasonMissingCardholderContact IssuingCardWalletsApplePayIneligibleReason = "missing_cardholder_contact"
	IssuingCardWalletsApplePayIneligibleReasonUnsupportedRegion        IssuingCardWalletsApplePayIneligibleReason = "unsupported_region"
)

// Reason the card is ineligible for Google Pay
type IssuingCardWalletsGooglePayIneligibleReason string

// List of values that IssuingCardWalletsGooglePayIneligibleReason can take
const (
	IssuingCardWalletsGooglePayIneligibleReasonMissingAgreement         IssuingCardWalletsGooglePayIneligibleReason = "missing_agreement"
	IssuingCardWalletsGooglePayIneligibleReasonMissingCardholderContact IssuingCardWalletsGooglePayIneligibleReason = "missing_cardholder_contact"
	IssuingCardWalletsGooglePayIneligibleReasonUnsupportedRegion        IssuingCardWalletsGooglePayIneligibleReason = "unsupported_region"
)

// Returns a list of Issuing Card objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
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

// The address where the card will be shipped.
type IssuingCardShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    string         `form:"name"`
	Service *string        `form:"service"`
	Type    *string        `form:"type"`
}

// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
type IssuingCardSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// Rules that control spending for this card. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardSpendingControlsParams struct {
	AllowedCategories      []*string                                         `form:"allowed_categories"`
	BlockedCategories      []*string                                         `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                           `form:"spending_limits_currency"`
}

// Creates an Issuing Card object.
type IssuingCardParams struct {
	Params            `form:"*"`
	Cardholder        *string                            `form:"cardholder"`
	Currency          *string                            `form:"currency"`
	ReplacementFor    *string                            `form:"replacement_for"`
	ReplacementReason *string                            `form:"replacement_reason"`
	Shipping          *IssuingCardShippingParams         `form:"shipping"`
	SpendingControls  *IssuingCardSpendingControlsParams `form:"spending_controls"`
	Status            *string                            `form:"status"`
	Type              *string                            `form:"type"`
	// The following parameter is only supported when updating a card
	CancellationReason *string `form:"cancellation_reason"`
}

// Where and how the card will be shipped.
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

// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
type IssuingCardSpendingControlsSpendingLimit struct {
	Amount     int64                                            `json:"amount"`
	Categories []string                                         `json:"categories"`
	Interval   IssuingCardSpendingControlsSpendingLimitInterval `json:"interval"`
}
type IssuingCardSpendingControls struct {
	AllowedCategories      []string                                    `json:"allowed_categories"`
	BlockedCategories      []string                                    `json:"blocked_categories"`
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimit `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                    `json:"spending_limits_currency"`
}
type IssuingCardWalletsApplePay struct {
	Eligible         bool                                       `json:"eligible"`
	IneligibleReason IssuingCardWalletsApplePayIneligibleReason `json:"ineligible_reason"`
}
type IssuingCardWalletsGooglePay struct {
	Eligible         bool                                        `json:"eligible"`
	IneligibleReason IssuingCardWalletsGooglePayIneligibleReason `json:"ineligible_reason"`
}

// Information relating to digital wallets (like Apple Pay and Google Pay).
type IssuingCardWallets struct {
	ApplePay                 *IssuingCardWalletsApplePay  `json:"apple_pay"`
	GooglePay                *IssuingCardWalletsGooglePay `json:"google_pay"`
	PrimaryAccountIdentifier string                       `json:"primary_account_identifier"`
}

// You can [create physical or virtual cards](https://stripe.com/docs/issuing/cards) that are issued to cardholders.
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
	Wallets            *IssuingCardWallets           `json:"wallets"`
}

// IssuingCardList is a list of Cards as retrieved from a list endpoint.
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
