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
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// Retrieves an Issuing Authorization object.
type IssuingAuthorizationParams struct {
	Params `form:"*"`
}

// Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
type IssuingAuthorizationApproveParams struct {
	Params `form:"*"`
	Amount *int64 `form:"amount"`
}

// Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
type IssuingAuthorizationDeclineParams struct {
	Params `form:"*"`
}

// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
type IssuingAuthorizationAmountDetails struct {
	ATMFee int64 `json:"atm_fee"`
}
type IssuingAuthorizationMerchantData struct {
	Category     string `json:"category"`
	CategoryCode string `json:"category_code"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Name         string `json:"name"`
	NetworkID    string `json:"network_id"`
	PostalCode   string `json:"postal_code"`
	State        string `json:"state"`
}

// The pending authorization request. This field will only be non-null during an `issuing_authorization.request` webhook.
type IssuingAuthorizationPendingRequest struct {
	Amount               int64                              `json:"amount"`
	AmountDetails        *IssuingAuthorizationAmountDetails `json:"amount_details"`
	Currency             Currency                           `json:"currency"`
	IsAmountControllable bool                               `json:"is_amount_controllable"`
	MerchantAmount       int64                              `json:"merchant_amount"`
	MerchantCurrency     Currency                           `json:"merchant_currency"`
}

// History of every time `pending_request` was approved/denied, either by you directly or by Stripe (e.g. based on your `spending_controls`). If the merchant changes the authorization by performing an [incremental authorization](https://stripe.com/docs/issuing/purchases/authorizations), you can look at this field to see the previous requests for the authorization.
type IssuingAuthorizationRequestHistory struct {
	Amount           int64                                    `json:"amount"`
	AmountDetails    *IssuingAuthorizationAmountDetails       `json:"amount_details"`
	Approved         bool                                     `json:"approved"`
	Created          int64                                    `json:"created"`
	Currency         Currency                                 `json:"currency"`
	MerchantAmount   int64                                    `json:"merchant_amount"`
	MerchantCurrency Currency                                 `json:"merchant_currency"`
	Reason           IssuingAuthorizationRequestHistoryReason `json:"reason"`
}
type IssuingAuthorizationVerificationData struct {
	AddressLine1Check      IssuingAuthorizationVerificationDataCheck `json:"address_line1_check"`
	AddressPostalCodeCheck IssuingAuthorizationVerificationDataCheck `json:"address_postal_code_check"`
	CVCCheck               IssuingAuthorizationVerificationDataCheck `json:"cvc_check"`
	ExpiryCheck            IssuingAuthorizationVerificationDataCheck `json:"expiry_check"`
}

// When an [issued card](https://stripe.com/docs/issuing) is used to make a purchase, an Issuing `Authorization`
// object is created. [Authorizations](https://stripe.com/docs/issuing/purchases/authorizations) must be approved for the
// purchase to be completed successfully.
//
// Related guide: [Issued Card Authorizations](https://stripe.com/docs/issuing/purchases/authorizations).
type IssuingAuthorization struct {
	APIResource
	Amount              int64                                   `json:"amount"`
	AmountDetails       *IssuingAuthorizationAmountDetails      `json:"amount_details"`
	Approved            bool                                    `json:"approved"`
	AuthorizationMethod IssuingAuthorizationAuthorizationMethod `json:"authorization_method"`
	BalanceTransactions []*BalanceTransaction                   `json:"balance_transactions"`
	Card                *IssuingCard                            `json:"card"`
	Cardholder          *IssuingCardholder                      `json:"cardholder"`
	Created             int64                                   `json:"created"`
	Currency            Currency                                `json:"currency"`
	ID                  string                                  `json:"id"`
	Livemode            bool                                    `json:"livemode"`
	MerchantAmount      int64                                   `json:"merchant_amount"`
	MerchantCurrency    Currency                                `json:"merchant_currency"`
	MerchantData        *IssuingAuthorizationMerchantData       `json:"merchant_data"`
	Metadata            map[string]string                       `json:"metadata"`
	Object              string                                  `json:"object"`
	PendingRequest      *IssuingAuthorizationPendingRequest     `json:"pending_request"`
	RequestHistory      []*IssuingAuthorizationRequestHistory   `json:"request_history"`
	Status              IssuingAuthorizationStatus              `json:"status"`
	Transactions        []*IssuingTransaction                   `json:"transactions"`
	VerificationData    *IssuingAuthorizationVerificationData   `json:"verification_data"`
	Wallet              IssuingAuthorizationWalletType          `json:"wallet"`
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
