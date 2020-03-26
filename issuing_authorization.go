package stripe

import "encoding/json"

// IssuingAuthorizationAuthorizationMethod is the list of possible values for the authorization method
// on an issuing authorization.
type IssuingAuthorizationAuthorizationMethod string

// List of values that IssuingAuthorizationAuthorizationMethod can take.
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

// IssuingAuthorizationRequestHistoryReason is the list of possible values for the request history
// reason on an issuing authorization.
type IssuingAuthorizationRequestHistoryReason string

// List of values that IssuingAuthorizationRequestHistoryReason can take.
const (
	IssuingAuthorizationRequestHistoryReasonAccountComplianceDisabled      IssuingAuthorizationRequestHistoryReason = "account_compliance_disabled"
	IssuingAuthorizationRequestHistoryReasonAccountInactive                IssuingAuthorizationRequestHistoryReason = "account_inactive"
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

	// The following value is deprecated. Use IssuingAuthorizationRequestHistoryReasonSpendingControls instead.
	IssuingAuthorizationRequestHistoryReasonAuthorizationControls IssuingAuthorizationRequestHistoryReason = "authorization_controls"

	// The following values are deprecated. Use IssuingAuthorizationRequestHistoryReasonVerificationFailed instead
	IssuingAuthorizationRequestHistoryReasonAuthenticationFailed IssuingAuthorizationRequestHistoryReason = "authentication_failed"
	IssuingAuthorizationRequestHistoryReasonIncorrectCVC         IssuingAuthorizationRequestHistoryReason = "incorrect_cvc"
	IssuingAuthorizationRequestHistoryReasonIncorrectExpiry      IssuingAuthorizationRequestHistoryReason = "incorrect_expiry"
)

// IssuingAuthorizationStatus is the possible values for status for an issuing authorization.
type IssuingAuthorizationStatus string

// List of values that IssuingAuthorizationStatus can take.
const (
	IssuingAuthorizationStatusClosed   IssuingAuthorizationStatus = "closed"
	IssuingAuthorizationStatusPending  IssuingAuthorizationStatus = "pending"
	IssuingAuthorizationStatusReversed IssuingAuthorizationStatus = "reversed"
)

// IssuingAuthorizationVerificationDataAuthentication is the list of possible values for the result
// of an authentication on an issuing authorization.
type IssuingAuthorizationVerificationDataAuthentication string

// List of values that IssuingAuthorizationVerificationDataCheck can take.
const (
	IssuingAuthorizationVerificationDataAuthenticationExempt  IssuingAuthorizationVerificationDataAuthentication = "exempt"
	IssuingAuthorizationVerificationDataAuthenticationFailure IssuingAuthorizationVerificationDataAuthentication = "failure"
	IssuingAuthorizationVerificationDataAuthenticationNone    IssuingAuthorizationVerificationDataAuthentication = "none"
	IssuingAuthorizationVerificationDataAuthenticationSuccess IssuingAuthorizationVerificationDataAuthentication = "success"
)

// IssuingAuthorizationVerificationDataCheck is the list of possible values for result of a check
// for verification data on an issuing authorization.
type IssuingAuthorizationVerificationDataCheck string

// List of values that IssuingAuthorizationVerificationDataCheck can take.
const (
	IssuingAuthorizationVerificationDataCheckMatch       IssuingAuthorizationVerificationDataCheck = "match"
	IssuingAuthorizationVerificationDataCheckMismatch    IssuingAuthorizationVerificationDataCheck = "mismatch"
	IssuingAuthorizationVerificationDataCheckNotProvided IssuingAuthorizationVerificationDataCheck = "not_provided"
)

// IssuingAuthorizationVerificationDataThreeDSecureResult is the list of possible values for result of 3DS.
type IssuingAuthorizationVerificationDataThreeDSecureResult string

// List of values that IssuingAuthorizationVerificationDataThreeDSecureResult can take.
const (
	IssuingAuthorizationVerificationDataThreeDSecureResultAttemptAcknowledged IssuingAuthorizationVerificationDataThreeDSecureResult = "attempt_acknowledged"
	IssuingAuthorizationVerificationDataThreeDSecureResultAuthenticated       IssuingAuthorizationVerificationDataThreeDSecureResult = "authenticated"
	IssuingAuthorizationVerificationDataThreeDSecureResultFailed              IssuingAuthorizationVerificationDataThreeDSecureResult = "failed"
)

// IssuingAuthorizationWalletType is the list of possible values for the authorization's wallet provider.
type IssuingAuthorizationWalletType string

// List of values that IssuingAuthorizationWalletType can take.
const (
	IssuingAuthorizationWalletTypeApplePay   IssuingAuthorizationWalletType = "apple_pay"
	IssuingAuthorizationWalletTypeGooglePay  IssuingAuthorizationWalletType = "google_pay"
	IssuingAuthorizationWalletTypeSamsungPay IssuingAuthorizationWalletType = "samsung_pay"
)

// IssuingAuthorizationWalletProviderType is the list of possible values for the authorization's wallet provider.
// TODO remove in the next major version
type IssuingAuthorizationWalletProviderType string

// List of values that IssuingAuthorizationWalletProviderType can take.
const (
	IssuingAuthorizationWalletProviderTypeApplePay   IssuingAuthorizationWalletProviderType = "apple_pay"
	IssuingAuthorizationWalletProviderTypeGooglePay  IssuingAuthorizationWalletProviderType = "google_pay"
	IssuingAuthorizationWalletProviderTypeSamsungPay IssuingAuthorizationWalletProviderType = "samsung_pay"
)

// IssuingAuthorizationParams is the set of parameters that can be used when updating an issuing authorization.
type IssuingAuthorizationParams struct {
	Params `form:"*"`
	Amount *int64 `form:"amount"`

	// The following parameter is deprecated, use Amount instead.
	// TODO: remove in a future major version
	HeldAmount *int64 `form:"held_amount"`
}

// IssuingAuthorizationListParams is the set of parameters that can be used when listing issuing authorizations.
type IssuingAuthorizationListParams struct {
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// IssuingAuthorizationAuthorizationControls is the resource representing authorization controls on an issuing authorization.
// This is deprecated and will be removed in the next major version
type IssuingAuthorizationAuthorizationControls struct {
	AllowedCategories []string `json:"allowed_categories"`
	BlockedCategories []string `json:"blocked_categories"`
	Currency          Currency `json:"currency"`
	MaxAmount         int64    `json:"max_amount"`
	MaxApprovals      int64    `json:"max_approvals"`
}

// IssuingAuthorizationPendingRequest is the resource representing details about the pending authorization request.
type IssuingAuthorizationPendingRequest struct {
	Amount               int64    `json:"amount"`
	Currency             Currency `json:"currency"`
	IsAmountControllable bool     `json:"is_amount_controllable"`
	MerchantAmount       int64    `json:"merchant_amount"`
	MerchantCurrency     Currency `json:"merchant_currency"`
}

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControl is the resource representing an
// authorizaton control that caused the authorization to fail.
// This is deprecated and will be removed in the next major version
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControl struct {
	Entity IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity `json:"entity"`
	Name   IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName   `json:"name"`
}

// IssuingAuthorizationRequestHistory is the resource representing a request history on an issuing authorization.
type IssuingAuthorizationRequestHistory struct {
	Amount           int64                                    `json:"amount"`
	Approved         bool                                     `json:"approved"`
	Created          int64                                    `json:"created"`
	Currency         Currency                                 `json:"currency"`
	MerchantAmount   int64                                    `json:"merchant_amount"`
	MerchantCurrency Currency                                 `json:"merchant_currency"`
	Reason           IssuingAuthorizationRequestHistoryReason `json:"reason"`

	// The following properties are deprecated
	// TODO: remove in the next major version
	AuthorizedAmount              int64                                                             `json:"authorized_amount"`
	AuthorizedCurrency            Currency                                                          `json:"authorized_currency"`
	HeldAmount                    int64                                                             `json:"held_amount"`
	HeldCurrency                  Currency                                                          `json:"held_currency"`
	ViolatedAuthorizationControls []*IssuingAuthorizationRequestHistoryViolatedAuthorizationControl `json:"violated_authorization_controls"`
}

// IssuingAuthorizationVerificationDataThreeDSecure is the resource representing 3DS results.
type IssuingAuthorizationVerificationDataThreeDSecure struct {
	Result IssuingAuthorizationVerificationDataThreeDSecureResult `json:"result"`
}

// IssuingAuthorizationVerificationData is the resource representing verification data on an issuing authorization.
type IssuingAuthorizationVerificationData struct {
	AddressLine1Check      IssuingAuthorizationVerificationDataCheck         `json:"address_line1_check"`
	AddressPostalCodeCheck IssuingAuthorizationVerificationDataCheck         `json:"address_postal_code_check"`
	CVCCheck               IssuingAuthorizationVerificationDataCheck         `json:"cvc_check"`
	ExpiryCheck            IssuingAuthorizationVerificationDataCheck         `json:"expiry_check"`
	ThreeDSecure           *IssuingAuthorizationVerificationDataThreeDSecure `json:"three_d_secure"`

	// The property is considered deprecated. Use AddressPostalCodeCheck instead.
	// TODO remove in the next major version
	AddressZipCheck IssuingAuthorizationVerificationDataCheck `json:"address_zip_check"`

	// The property is considered deprecated. Use ThreeDSecure instead.
	// TODO remove in the next major version
	Authentication IssuingAuthorizationVerificationDataAuthentication `json:"authentication"`
}

// IssuingAuthorization is the resource representing a Stripe issuing authorization.
type IssuingAuthorization struct {
	Amount              int64                                   `json:"amount"`
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
	MerchantData        *IssuingMerchantData                    `json:"merchant_data"`
	Metadata            map[string]string                       `json:"metadata"`
	Object              string                                  `json:"object"`
	PendingRequest      *IssuingAuthorizationPendingRequest     `json:"pending_request"`
	RequestHistory      []*IssuingAuthorizationRequestHistory   `json:"request_history"`
	Status              IssuingAuthorizationStatus              `json:"status"`
	Transactions        []*IssuingTransaction                   `json:"transactions"`
	VerificationData    *IssuingAuthorizationVerificationData   `json:"verification_data"`
	Wallet              IssuingAuthorizationWalletType          `json:"wallet"`

	// This property is deprecated and we recommend that you use Wallet instead.
	// TODO: remove in the next major version
	WalletProvider IssuingAuthorizationWalletProviderType `json:"wallet_provider"`

	// The following properties are considered deprecated
	// TODO: remove in the next major version
	AuthorizedAmount         int64    `json:"authorized_amount"`
	AuthorizedCurrency       Currency `json:"authorized_currency"`
	HeldAmount               int64    `json:"held_amount"`
	HeldCurrency             Currency `json:"held_currency"`
	IsHeldAmountControllable bool     `json:"is_held_amount_controllable"`
	PendingAuthorizedAmount  int64    `json:"pending_authorized_amount"`
	PendingHeldAmount        int64    `json:"pending_held_amount"`
}

// IssuingMerchantData is the resource representing merchant data on Issuing APIs.
type IssuingMerchantData struct {
	Category   string `json:"category"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Name       string `json:"name"`
	NetworkID  string `json:"network_id"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	URL        string `json:"url"`
}

// IssuingAuthorizationList is a list of issuing authorizations as retrieved from a list endpoint.
type IssuingAuthorizationList struct {
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
