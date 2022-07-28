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
	ListParams `form:"*"`
	// Only return cards belonging to the Cardholder with the provided ID.
	Cardholder *string `form:"cardholder"`
	// Only return cards that were issued during the given date interval.
	Created *int64 `form:"created"`
	// Only return cards that were issued during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return cards that have the given expiration month.
	ExpMonth *int64 `form:"exp_month"`
	// Only return cards that have the given expiration year.
	ExpYear *int64 `form:"exp_year"`
	// Only return cards that have the given last four digits.
	Last4 *string `form:"last4"`
	// Only return cards that have the given status. One of `active`, `inactive`, or `canceled`.
	Status *string `form:"status"`
	// Only return cards that have the given type. One of `virtual` or `physical`.
	Type *string `form:"type"`
}

// The address where the card will be shipped.
type IssuingCardShippingParams struct {
	// The address that the card is shipped to.
	Address *AddressParams `form:"address"`
	// The name printed on the shipping label when shipping the card.
	Name string `form:"name"`
	// Shipment service.
	Service *string `form:"service"`
	// Packaging options.
	Type *string `form:"type"`
}

// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
type IssuingCardSpendingControlsSpendingLimitParams struct {
	// Maximum amount allowed to spend per interval.
	Amount *int64 `form:"amount"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.
	Categories []*string `form:"categories"`
	// Interval (or event) to which the amount applies.
	Interval *string `form:"interval"`
}

// Rules that control spending for this card. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
type IssuingCardSpendingControlsParams struct {
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.
	AllowedCategories []*string `form:"allowed_categories"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.
	BlockedCategories []*string `form:"blocked_categories"`
	// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                           `form:"spending_limits_currency"`
}

// Creates an Issuing Card object.
type IssuingCardParams struct {
	Params `form:"*"`
	// The [Cardholder](https://stripe.com/docs/api#issuing_cardholder_object) object with which the card will be associated.
	Cardholder *string `form:"cardholder"`
	// The currency for the card.
	Currency         *string `form:"currency"`
	FinancialAccount *string `form:"financial_account"`
	// The desired new PIN for this card.
	PIN *IssuingCardPINParams `form:"pin"`
	// The card this is meant to be a replacement for (if any).
	ReplacementFor *string `form:"replacement_for"`
	// If `replacement_for` is specified, this should indicate why that card is being replaced.
	ReplacementReason *string `form:"replacement_reason"`
	// The address where the card will be shipped.
	Shipping *IssuingCardShippingParams `form:"shipping"`
	// Rules that control spending for this card. Refer to our [documentation](https://stripe.com/docs/issuing/controls/spending-controls) for more details.
	SpendingControls *IssuingCardSpendingControlsParams `form:"spending_controls"`
	// Dictates whether authorizations can be approved on this card. If this card is being canceled because it was lost or stolen, this information should be provided as `cancellation_reason`.
	Status *string `form:"status"`
	// The type of card to issue. Possible values are `physical` or `virtual`.
	Type *string `form:"type"`
	// The following parameter is only supported when updating a card
	// Reason why the `status` of this card is `canceled`.
	CancellationReason *string `form:"cancellation_reason"`
}

// The desired new PIN for this card.
type IssuingCardPINParams struct {
	// The card's desired new PIN, encrypted under Stripe's public key.
	EncryptedNumber *string `form:"encrypted_number"`
}

// Where and how the card will be shipped.
type IssuingCardShipping struct {
	Address *Address `json:"address"`
	// The delivery company that shipped a card.
	Carrier IssuingCardShippingCarrier `json:"carrier"`
	// A unix timestamp representing a best estimate of when the card will be delivered.
	ETA int64 `json:"eta"`
	// Recipient name.
	Name string `json:"name"`
	// Shipment service, such as `standard` or `express`.
	Service IssuingCardShippingService `json:"service"`
	// The delivery status of the card.
	Status IssuingCardShippingStatus `json:"status"`
	// A tracking number for a card shipment.
	TrackingNumber string `json:"tracking_number"`
	// A link to the shipping carrier's site where you can view detailed information about a card shipment.
	TrackingURL string `json:"tracking_url"`
	// Packaging options.
	Type IssuingCardShippingType `json:"type"`
}

// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
type IssuingCardSpendingControlsSpendingLimit struct {
	// Maximum amount allowed to spend per interval. This amount is in the card's currency and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.
	Categories []string `json:"categories"`
	// Interval (or event) to which the amount applies.
	Interval IssuingCardSpendingControlsSpendingLimitInterval `json:"interval"`
}
type IssuingCardSpendingControls struct {
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.
	AllowedCategories []string `json:"allowed_categories"`
	// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.
	BlockedCategories []string `json:"blocked_categories"`
	// Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).
	SpendingLimits []*IssuingCardSpendingControlsSpendingLimit `json:"spending_limits"`
	// Currency of the amounts within `spending_limits`. Always the same as the currency of the card.
	SpendingLimitsCurrency Currency `json:"spending_limits_currency"`
}
type IssuingCardWalletsApplePay struct {
	// Apple Pay Eligibility
	Eligible bool `json:"eligible"`
	// Reason the card is ineligible for Apple Pay
	IneligibleReason IssuingCardWalletsApplePayIneligibleReason `json:"ineligible_reason"`
}
type IssuingCardWalletsGooglePay struct {
	// Google Pay Eligibility
	Eligible bool `json:"eligible"`
	// Reason the card is ineligible for Google Pay
	IneligibleReason IssuingCardWalletsGooglePayIneligibleReason `json:"ineligible_reason"`
}

// Information relating to digital wallets (like Apple Pay and Google Pay).
type IssuingCardWallets struct {
	ApplePay  *IssuingCardWalletsApplePay  `json:"apple_pay"`
	GooglePay *IssuingCardWalletsGooglePay `json:"google_pay"`
	// Unique identifier for a card used with digital wallets
	PrimaryAccountIdentifier string `json:"primary_account_identifier"`
}

// You can [create physical or virtual cards](https://stripe.com/docs/issuing/cards) that are issued to cardholders.
type IssuingCard struct {
	APIResource
	// The brand of the card.
	Brand string `json:"brand"`
	// The reason why the card was canceled.
	CancellationReason IssuingCardCancellationReason `json:"cancellation_reason"`
	// An Issuing `Cardholder` object represents an individual or business entity who is [issued](https://stripe.com/docs/issuing) cards.
	//
	// Related guide: [How to create a Cardholder](https://stripe.com/docs/issuing/cards#create-cardholder)
	Cardholder *IssuingCardholder `json:"cardholder"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Supported currencies are `usd` in the US, `eur` in the EU, and `gbp` in the UK.
	Currency Currency `json:"currency"`
	// The card's CVC. For security reasons, this is only available for virtual cards, and will be omitted unless you explicitly request it with [the `expand` parameter](https://stripe.com/docs/api/expanding_objects). Additionally, it's only available via the ["Retrieve a card" endpoint](https://stripe.com/docs/api/issuing/cards/retrieve), not via "List all cards" or any other endpoint.
	CVC string `json:"cvc"`
	// The expiration month of the card.
	ExpMonth int64 `json:"exp_month"`
	// The expiration year of the card.
	ExpYear int64 `json:"exp_year"`
	// The financial account this card is attached to.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The full unredacted card number. For security reasons, this is only available for virtual cards, and will be omitted unless you explicitly request it with [the `expand` parameter](https://stripe.com/docs/api/expanding_objects). Additionally, it's only available via the ["Retrieve a card" endpoint](https://stripe.com/docs/api/issuing/cards/retrieve), not via "List all cards" or any other endpoint.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The latest card that replaces this card, if any.
	ReplacedBy *IssuingCard `json:"replaced_by"`
	// The card this card replaces, if any.
	ReplacementFor *IssuingCard `json:"replacement_for"`
	// The reason why the previous card needed to be replaced.
	ReplacementReason IssuingCardReplacementReason `json:"replacement_reason"`
	// Where and how the card will be shipped.
	Shipping         *IssuingCardShipping         `json:"shipping"`
	SpendingControls *IssuingCardSpendingControls `json:"spending_controls"`
	// Whether authorizations can be approved on this card.
	Status IssuingCardStatus `json:"status"`
	// The type of the card.
	Type IssuingCardType `json:"type"`
	// Information relating to digital wallets (like Apple Pay and Google Pay).
	Wallets *IssuingCardWallets `json:"wallets"`
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
