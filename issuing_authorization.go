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
	IssuingAuthorizationRequestHistoryReasonAuthorizationControls IssuingAuthorizationRequestHistoryReason = "authorization_controls"
	IssuingAuthorizationRequestHistoryReasonCardActive            IssuingAuthorizationRequestHistoryReason = "card_active"
	IssuingAuthorizationRequestHistoryReasonCardInactive          IssuingAuthorizationRequestHistoryReason = "card_inactive"
	IssuingAuthorizationRequestHistoryReasonInsufficientFunds     IssuingAuthorizationRequestHistoryReason = "insufficient_funds"
	IssuingAuthorizationRequestHistoryReasonWebhookApproved       IssuingAuthorizationRequestHistoryReason = "webhook_approved"
	IssuingAuthorizationRequestHistoryReasonWebhookDeclined       IssuingAuthorizationRequestHistoryReason = "webhook_declined"
	IssuingAuthorizationRequestHistoryReasonWebhookTimeout        IssuingAuthorizationRequestHistoryReason = "webhook_timeout"
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

// IssuingAuthorizationWalletProviderType is the list of possible values for the authorization's wallet provider.
type IssuingAuthorizationWalletProviderType string

// List of values that IssuingAuthorizationWalletProviderType can take.
const (
	IssuingAuthorizationWalletProviderTypeApplePay   IssuingAuthorizationWalletProviderType = "apple_pay"
	IssuingAuthorizationWalletProviderTypeGooglePay  IssuingAuthorizationWalletProviderType = "google_pay"
	IssuingAuthorizationWalletProviderTypeSamsungPay IssuingAuthorizationWalletProviderType = "samsung_pay"
)

// IssuingAuthorizationParams is the set of parameters that can be used when updating an issuing authorization.
type IssuingAuthorizationParams struct {
	Params     `form:"*"`
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
type IssuingAuthorizationAuthorizationControls struct {
	AllowedCategories []string `json:"allowed_categories"`
	BlockedCategories []string `json:"blocked_categories"`
	Currency          Currency `json:"currency"`
	MaxAmount         int64    `json:"max_amount"`
	MaxApprovals      int64    `json:"max_approvals"`
}

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControl is the resource representing an
// authorizaton control that caused the authorization to fail.
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControl struct {
	Entity IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity `json:"entity"`
	Name   IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName   `json:"name"`
}

// IssuingAuthorizationRequestHistory is the resource representing a request history on an issuing authorization.
type IssuingAuthorizationRequestHistory struct {
	Approved                      bool                                                              `json:"approved"`
	AuthorizedAmount              int64                                                             `json:"authorized_amount"`
	AuthorizedCurrency            Currency                                                          `json:"authorized_currency"`
	Created                       int64                                                             `json:"created"`
	HeldAmount                    int64                                                             `json:"held_amount"`
	HeldCurrency                  Currency                                                          `json:"held_currency"`
	Reason                        IssuingAuthorizationRequestHistoryReason                          `json:"reason"`
	ViolatedAuthorizationControls []*IssuingAuthorizationRequestHistoryViolatedAuthorizationControl `json:"violated_authorization_controls"`
}

// IssuingAuthorizationVerificationData is the resource representing verification data on an issuing authorization.
type IssuingAuthorizationVerificationData struct {
	AddressLine1Check IssuingAuthorizationVerificationDataCheck          `json:"address_line1_check"`
	AddressZipCheck   IssuingAuthorizationVerificationDataCheck          `json:"address_zip_check"`
	Authentication    IssuingAuthorizationVerificationDataAuthentication `json:"authentication"`
	CVCCheck          IssuingAuthorizationVerificationDataCheck          `json:"cvc_check"`
}

// IssuingAuthorization is the resource representing a Stripe issuing authorization.
type IssuingAuthorization struct {
	Approved                 bool                                    `json:"approved"`
	AuthorizationMethod      IssuingAuthorizationAuthorizationMethod `json:"authorization_method"`
	AuthorizedAmount         int64                                   `json:"authorized_amount"`
	AuthorizedCurrency       Currency                                `json:"authorized_currency"`
	BalanceTransactions      []*BalanceTransaction                   `json:"balance_transactions"`
	Card                     *IssuingCard                            `json:"card"`
	Cardholder               *IssuingCardholder                      `json:"cardholder"`
	Created                  int64                                   `json:"created"`
	HeldAmount               int64                                   `json:"held_amount"`
	HeldCurrency             Currency                                `json:"held_currency"`
	ID                       string                                  `json:"id"`
	IsHeldAmountControllable bool                                    `json:"is_held_amount_controllable"`
	Livemode                 bool                                    `json:"livemode"`
	MerchantData             *IssuingMerchantData                    `json:"merchant_data"`
	Metadata                 map[string]string                       `json:"metadata"`
	Object                   string                                  `json:"object"`
	PendingAuthorizedAmount  int64                                   `json:"pending_authorized_amount"`
	PendingHeldAmount        int64                                   `json:"pending_held_amount"`
	RequestHistory           []*IssuingAuthorizationRequestHistory   `json:"request_history"`
	Status                   IssuingAuthorizationStatus              `json:"status"`
	Transactions             []*IssuingTransaction                   `json:"transactions"`
	VerificationData         *IssuingAuthorizationVerificationData   `json:"verification_data"`
	WalletProvider           IssuingAuthorizationWalletProviderType  `json:"wallet_provider"`
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
