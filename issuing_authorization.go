//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// How the card details were provided.
type IssuingAuthorizationAuthorizationMethod string

// List of values that IssuingAuthorizationAuthorizationMethod can take
const (
	IssuingAuthorizationAuthorizationMethodChip        IssuingAuthorizationAuthorizationMethod = "chip"
	IssuingAuthorizationAuthorizationMethodContactless IssuingAuthorizationAuthorizationMethod = "contactless"
	IssuingAuthorizationAuthorizationMethodKeyedIn     IssuingAuthorizationAuthorizationMethod = "keyed_in"
	IssuingAuthorizationAuthorizationMethodOnline      IssuingAuthorizationAuthorizationMethod = "online"
	IssuingAuthorizationAuthorizationMethodSwipe       IssuingAuthorizationAuthorizationMethod = "swipe"
)

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity is the list of possible values
// for the entity that owns the authorization control.
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity string

// List of values that IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity can take.
const (
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityAccount    IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "account"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityCard       IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "card"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityCardholder IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "cardholder"
)

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName is the list of possible values
// for the name associated with the authorization control.
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName string

// List of values that IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName can take.
const (
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameAllowedCategories IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "allowed_categories"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameBlockedCategories IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "blocked_categories"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameMaxAmount         IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "max_amount"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameMaxApprovals      IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "max_approvals"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameSpendingLimits    IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "spending_limits"
)

// The reason for the approval or decline.
type IssuingAuthorizationRequestHistoryReason string

// List of values that IssuingAuthorizationRequestHistoryReason can take
const (
	IssuingAuthorizationRequestHistoryReasonAccountDisabled                IssuingAuthorizationRequestHistoryReason = "account_disabled"
	IssuingAuthorizationRequestHistoryReasonCardActive                     IssuingAuthorizationRequestHistoryReason = "card_active"
	IssuingAuthorizationRequestHistoryReasonCardInactive                   IssuingAuthorizationRequestHistoryReason = "card_inactive"
	IssuingAuthorizationRequestHistoryReasonCardholderInactive             IssuingAuthorizationRequestHistoryReason = "cardholder_inactive"
	IssuingAuthorizationRequestHistoryReasonCardholderVerificationRequired IssuingAuthorizationRequestHistoryReason = "cardholder_verification_required"
	IssuingAuthorizationRequestHistoryReasonInsufficientFunds              IssuingAuthorizationRequestHistoryReason = "insufficient_funds"
	IssuingAuthorizationRequestHistoryReasonNotAllowed                     IssuingAuthorizationRequestHistoryReason = "not_allowed"
	IssuingAuthorizationRequestHistoryReasonSpendingControls               IssuingAuthorizationRequestHistoryReason = "spending_controls"
	IssuingAuthorizationRequestHistoryReasonSuspectedFraud                 IssuingAuthorizationRequestHistoryReason = "suspected_fraud"
	IssuingAuthorizationRequestHistoryReasonVerificationFailed             IssuingAuthorizationRequestHistoryReason = "verification_failed"
	IssuingAuthorizationRequestHistoryReasonWebhookApproved                IssuingAuthorizationRequestHistoryReason = "webhook_approved"
	IssuingAuthorizationRequestHistoryReasonWebhookDeclined                IssuingAuthorizationRequestHistoryReason = "webhook_declined"
	IssuingAuthorizationRequestHistoryReasonWebhookTimeout                 IssuingAuthorizationRequestHistoryReason = "webhook_timeout"
)

// The current status of the authorization in its lifecycle.
type IssuingAuthorizationStatus string

// List of values that IssuingAuthorizationStatus can take
const (
	IssuingAuthorizationStatusClosed   IssuingAuthorizationStatus = "closed"
	IssuingAuthorizationStatusPending  IssuingAuthorizationStatus = "pending"
	IssuingAuthorizationStatusReversed IssuingAuthorizationStatus = "reversed"
)

// Whether the cardholder provided an address first line and if it matched the cardholder's `billing.address.line1`.
type IssuingAuthorizationVerificationDataCheck string

// List of values that IssuingAuthorizationVerificationDataCheck can take
const (
	IssuingAuthorizationVerificationDataCheckMatch       IssuingAuthorizationVerificationDataCheck = "match"
	IssuingAuthorizationVerificationDataCheckMismatch    IssuingAuthorizationVerificationDataCheck = "mismatch"
	IssuingAuthorizationVerificationDataCheckNotProvided IssuingAuthorizationVerificationDataCheck = "not_provided"
)

// The digital wallet used for this authorization. One of `apple_pay`, `google_pay`, or `samsung_pay`.
type IssuingAuthorizationWalletType string

// List of values that IssuingAuthorizationWalletType can take
const (
	IssuingAuthorizationWalletTypeApplePay   IssuingAuthorizationWalletType = "apple_pay"
	IssuingAuthorizationWalletTypeGooglePay  IssuingAuthorizationWalletType = "google_pay"
	IssuingAuthorizationWalletTypeSamsungPay IssuingAuthorizationWalletType = "samsung_pay"
)

// Returns a list of Issuing Authorization objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingAuthorizationListParams struct {
	ListParams `form:"*"`
	// Only return authorizations that belong to the given card.
	Card *string `form:"card"`
	// Only return authorizations that belong to the given cardholder.
	Cardholder *string `form:"cardholder"`
	// Only return authorizations that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return authorizations that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return authorizations with the given status. One of `pending`, `closed`, or `reversed`.
	Status *string `form:"status"`
}

// Retrieves an Issuing Authorization object.
type IssuingAuthorizationParams struct {
	Params `form:"*"`
}

// Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
type IssuingAuthorizationApproveParams struct {
	Params `form:"*"`
	// If the authorization's `pending_request.is_amount_controllable` property is `true`, you may provide this value to control how much to hold for the authorization. Must be positive (use [`decline`](https://stripe.com/docs/api/issuing/authorizations/decline) to decline an authorization request).
	Amount *int64 `form:"amount"`
}

// Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
type IssuingAuthorizationDeclineParams struct {
	Params `form:"*"`
}

// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
type IssuingAuthorizationAmountDetails struct {
	// The fee charged by the ATM for the cash withdrawal.
	ATMFee int64 `json:"atm_fee"`
}
type IssuingAuthorizationMerchantData struct {
	// A categorization of the seller's type of business. See our [merchant categories guide](https://stripe.com/docs/issuing/merchant-categories) for a list of possible values.
	Category string `json:"category"`
	// The merchant category code for the seller's business
	CategoryCode string `json:"category_code"`
	// City where the seller is located
	City string `json:"city"`
	// Country where the seller is located
	Country string `json:"country"`
	// Name of the seller
	Name string `json:"name"`
	// Identifier assigned to the seller by the card brand
	NetworkID string `json:"network_id"`
	// Postal code where the seller is located
	PostalCode string `json:"postal_code"`
	// State where the seller is located
	State string `json:"state"`
}

// The pending authorization request. This field will only be non-null during an `issuing_authorization.request` webhook.
type IssuingAuthorizationPendingRequest struct {
	// The additional amount Stripe will hold if the authorization is approved, in the card's [currency](https://stripe.com/docs/api#issuing_authorization_object-pending-request-currency) and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	AmountDetails *IssuingAuthorizationAmountDetails `json:"amount_details"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// If set `true`, you may provide [amount](https://stripe.com/docs/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization.
	IsAmountControllable bool `json:"is_amount_controllable"`
	// The amount the merchant is requesting to be authorized in the `merchant_currency`. The amount is in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	MerchantAmount int64 `json:"merchant_amount"`
	// The local currency the merchant is requesting to authorize.
	MerchantCurrency Currency `json:"merchant_currency"`
}

// History of every time `pending_request` was approved/denied, either by you directly or by Stripe (e.g. based on your `spending_controls`). If the merchant changes the authorization by performing an [incremental authorization](https://stripe.com/docs/issuing/purchases/authorizations), you can look at this field to see the previous requests for the authorization.
type IssuingAuthorizationRequestHistory struct {
	// The `pending_request.amount` at the time of the request, presented in your card's currency and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal). Stripe held this amount from your account to fund the authorization if the request was approved.
	Amount int64 `json:"amount"`
	// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	AmountDetails *IssuingAuthorizationAmountDetails `json:"amount_details"`
	// Whether this request was approved.
	Approved bool `json:"approved"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The `pending_request.merchant_amount` at the time of the request, presented in the `merchant_currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	MerchantAmount int64 `json:"merchant_amount"`
	// The currency that was collected by the merchant and presented to the cardholder for the authorization. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	MerchantCurrency Currency `json:"merchant_currency"`
	// The reason for the approval or decline.
	Reason IssuingAuthorizationRequestHistoryReason `json:"reason"`
}

// [Treasury](https://stripe.com/docs/api/treasury) details related to this authorization if it was created on a [FinancialAccount](https://stripe.com/docs/api/treasury/financial_accounts).
type IssuingAuthorizationTreasury struct {
	// The array of [ReceivedCredits](https://stripe.com/docs/api/treasury/received_credits) associated with this authorization
	ReceivedCredits []string `json:"received_credits"`
	// The array of [ReceivedDebits](https://stripe.com/docs/api/treasury/received_debits) associated with this authorization
	ReceivedDebits []string `json:"received_debits"`
	// The Treasury [Transaction](https://stripe.com/docs/api/treasury/transactions) associated with this authorization
	Transaction string `json:"transaction"`
}
type IssuingAuthorizationVerificationData struct {
	// Whether the cardholder provided an address first line and if it matched the cardholder's `billing.address.line1`.
	AddressLine1Check IssuingAuthorizationVerificationDataCheck `json:"address_line1_check"`
	// Whether the cardholder provided a postal code and if it matched the cardholder's `billing.address.postal_code`.
	AddressPostalCodeCheck IssuingAuthorizationVerificationDataCheck `json:"address_postal_code_check"`
	// Whether the cardholder provided a CVC and if it matched Stripe's record.
	CVCCheck IssuingAuthorizationVerificationDataCheck `json:"cvc_check"`
	// Whether the cardholder provided an expiry date and if it matched Stripe's record.
	ExpiryCheck IssuingAuthorizationVerificationDataCheck `json:"expiry_check"`
}

// When an [issued card](https://stripe.com/docs/issuing) is used to make a purchase, an Issuing `Authorization`
// object is created. [Authorizations](https://stripe.com/docs/issuing/purchases/authorizations) must be approved for the
// purchase to be completed successfully.
//
// Related guide: [Issued Card Authorizations](https://stripe.com/docs/issuing/purchases/authorizations).
type IssuingAuthorization struct {
	APIResource
	// The total amount that was authorized or rejected. This amount is in the card's currency and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	AmountDetails *IssuingAuthorizationAmountDetails `json:"amount_details"`
	// Whether the authorization has been approved.
	Approved bool `json:"approved"`
	// How the card details were provided.
	AuthorizationMethod IssuingAuthorizationAuthorizationMethod `json:"authorization_method"`
	// List of balance transactions associated with this authorization.
	BalanceTransactions []*BalanceTransaction `json:"balance_transactions"`
	// You can [create physical or virtual cards](https://stripe.com/docs/issuing/cards) that are issued to cardholders.
	Card *IssuingCard `json:"card"`
	// The cardholder to whom this authorization belongs.
	Cardholder *IssuingCardholder `json:"cardholder"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The total amount that was authorized or rejected. This amount is in the `merchant_currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	MerchantAmount int64 `json:"merchant_amount"`
	// The currency that was presented to the cardholder for the authorization. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	MerchantCurrency Currency                          `json:"merchant_currency"`
	MerchantData     *IssuingAuthorizationMerchantData `json:"merchant_data"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The pending authorization request. This field will only be non-null during an `issuing_authorization.request` webhook.
	PendingRequest *IssuingAuthorizationPendingRequest `json:"pending_request"`
	// History of every time `pending_request` was approved/denied, either by you directly or by Stripe (e.g. based on your `spending_controls`). If the merchant changes the authorization by performing an [incremental authorization](https://stripe.com/docs/issuing/purchases/authorizations), you can look at this field to see the previous requests for the authorization.
	RequestHistory []*IssuingAuthorizationRequestHistory `json:"request_history"`
	// The current status of the authorization in its lifecycle.
	Status IssuingAuthorizationStatus `json:"status"`
	// List of [transactions](https://stripe.com/docs/api/issuing/transactions) associated with this authorization.
	Transactions []*IssuingTransaction `json:"transactions"`
	// [Treasury](https://stripe.com/docs/api/treasury) details related to this authorization if it was created on a [FinancialAccount](https://stripe.com/docs/api/treasury/financial_accounts).
	Treasury         *IssuingAuthorizationTreasury         `json:"treasury"`
	VerificationData *IssuingAuthorizationVerificationData `json:"verification_data"`
	// The digital wallet used for this authorization. One of `apple_pay`, `google_pay`, or `samsung_pay`.
	Wallet IssuingAuthorizationWalletType `json:"wallet"`
}

// IssuingAuthorizationList is a list of Authorizations as retrieved from a list endpoint.
type IssuingAuthorizationList struct {
	APIResource
	ListMeta
	Data []*IssuingAuthorization `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingAuthorization.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingAuthorization) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingAuthorization IssuingAuthorization
	var v issuingAuthorization
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingAuthorization(v)
	return nil
}
