//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. An enum representing the status of the FinancialAccount. This indicates whether or not the FinancialAccount can be used for any money movement flows.
type V2MoneyManagementFinancialAccountStatus string

// List of values that V2MoneyManagementFinancialAccountStatus can take
const (
	V2MoneyManagementFinancialAccountStatusClosed  V2MoneyManagementFinancialAccountStatus = "closed"
	V2MoneyManagementFinancialAccountStatusOpen    V2MoneyManagementFinancialAccountStatus = "open"
	V2MoneyManagementFinancialAccountStatusPending V2MoneyManagementFinancialAccountStatus = "pending"
)

// Closed Enum. The reason the FinancialAccount was closed.
type V2MoneyManagementFinancialAccountStatusDetailsClosedReason string

// List of values that V2MoneyManagementFinancialAccountStatusDetailsClosedReason can take
const (
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonAccountClosed    V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "account_closed"
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonClosedByPlatform V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "closed_by_platform"
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonOther            V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "other"
)

// The bank where funds are stored.
type V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityBankName string

// List of values that V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityBankName can take
const (
	V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityBankNameFifthThird V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityBankName = "fifth_third"
)

// The deposit insurance scheme.
type V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityType string

// List of values that V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityType can take
const (
	V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityTypeFdic            V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityType = "fdic"
	V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityTypeFdicPassthrough V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityType = "fdic_passthrough"
)

// Closed Enum. The type of the FinancialAccount. Details specific to the type appear in the matching hash.
// For example, a `storage` FinancialAccount includes a `storage` hash.
type V2MoneyManagementFinancialAccountType string

// List of values that V2MoneyManagementFinancialAccountType can take
const (
	V2MoneyManagementFinancialAccountTypeOther   V2MoneyManagementFinancialAccountType = "other"
	V2MoneyManagementFinancialAccountTypeStorage V2MoneyManagementFinancialAccountType = "storage"
)

// The FinancialAccount's balance, split into `available`, `inbound_pending`, and `outbound_pending` balances.
// Each balance is a hash keyed by lowercase currency codes, with an amount for each currency.
type V2MoneyManagementFinancialAccountBalance struct {
	// Balance available for money movement.
	Available map[string]Amount `json:"available"`
	// Balance of inbound funds that aren't yet available for money movement and will later move to the `available` balance.
	InboundPending map[string]Amount `json:"inbound_pending"`
	// Balance of funds being used for pending outbound money movements.
	OutboundPending map[string]Amount `json:"outbound_pending"`
}

// If the FinancialAccount's `type` is `other`, this hash identifies its actual type.
// Upgrade your API version to see that type in `type`.
type V2MoneyManagementFinancialAccountOther struct {
	// The FinancialAccount's actual type, represented as a string.
	// Upgrade your API version to see this value in `financial_account.type`.
	Type string `json:"type"`
}

// When configured, the forwarding settings for the closed FinancialAccount.
type V2MoneyManagementFinancialAccountStatusDetailsClosedForwardingSettings struct {
	// ID of the FinancialAccount that supplies funds to cover negative balances
	// during or after this FinancialAccount's closure.
	PaymentMethod string `json:"payment_method,omitempty"`
	// ID of the FinancialAccount that receives funds forwarded from this FinancialAccount when it closes.
	PayoutMethod string `json:"payout_method,omitempty"`
}

// Details about the FinancialAccount's closure. Only populated when `status` is `closed`.
type V2MoneyManagementFinancialAccountStatusDetailsClosed struct {
	// When configured, the forwarding settings for the closed FinancialAccount.
	ForwardingSettings *V2MoneyManagementFinancialAccountStatusDetailsClosedForwardingSettings `json:"forwarding_settings,omitempty"`
	// Closed Enum. The reason the FinancialAccount was closed.
	Reason V2MoneyManagementFinancialAccountStatusDetailsClosedReason `json:"reason"`
}

// Additional details related to the status of the FinancialAccount.
type V2MoneyManagementFinancialAccountStatusDetails struct {
	// Details about the FinancialAccount's closure. Only populated when `status` is `closed`.
	Closed *V2MoneyManagementFinancialAccountStatusDetailsClosed `json:"closed,omitempty"`
}

// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
type V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibility struct {
	// The bank where funds are stored.
	BankName V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityBankName `json:"bank_name"`
	// Currencies eligible for deposit insurance at this bank under this scheme.
	Currencies []Currency `json:"currencies"`
	// The deposit insurance scheme.
	Type V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityType `json:"type"`
}

// If the FinancialAccount's `type` is `storage`, this hash contains details specific to that type.
type V2MoneyManagementFinancialAccountStorage struct {
	// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
	DepositInsuranceEligibility []*V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibility `json:"deposit_insurance_eligibility,omitempty"`
	// The currencies that this FinancialAccount can hold.
	HoldsCurrencies []Currency `json:"holds_currencies"`
}

// A FinancialAccount represents a balance and can be the source or destination of money movement.
type V2MoneyManagementFinancialAccount struct {
	APIResource
	// The FinancialAccount's balance, split into `available`, `inbound_pending`, and `outbound_pending` balances.
	// Each balance is a hash keyed by lowercase currency codes, with an amount for each currency.
	Balance *V2MoneyManagementFinancialAccountBalance `json:"balance"`
	// Open Enum. The country where the LegalEntity associated with the FinancialAccount is based, represented as a two-letter country code.
	Country string `json:"country"`
	// Time at which the object was created. Represented as an RFC 3339 timestamp in UTC with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name appears in the Stripe Dashboard and embedded components.
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information
	// about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// If the FinancialAccount's `type` is `other`, this hash identifies its actual type.
	// Upgrade your API version to see that type in `type`.
	Other *V2MoneyManagementFinancialAccountOther `json:"other,omitempty"`
	// Closed Enum. An enum representing the status of the FinancialAccount. This indicates whether or not the FinancialAccount can be used for any money movement flows.
	Status V2MoneyManagementFinancialAccountStatus `json:"status"`
	// Additional details related to the status of the FinancialAccount.
	StatusDetails *V2MoneyManagementFinancialAccountStatusDetails `json:"status_details,omitempty"`
	// If the FinancialAccount's `type` is `storage`, this hash contains details specific to that type.
	Storage *V2MoneyManagementFinancialAccountStorage `json:"storage,omitempty"`
	// Closed Enum. The type of the FinancialAccount. Details specific to the type appear in the matching hash.
	// For example, a `storage` FinancialAccount includes a `storage` hash.
	Type V2MoneyManagementFinancialAccountType `json:"type"`
}
