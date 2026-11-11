//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The bank transfer network for this mandate.
type V2MoneyManagementSpendingMandateBankTransferNetwork string

// List of values that V2MoneyManagementSpendingMandateBankTransferNetwork can take
const (
	V2MoneyManagementSpendingMandateBankTransferNetworkBACS V2MoneyManagementSpendingMandateBankTransferNetwork = "bacs"
)

// The status of the SpendingMandate.
type V2MoneyManagementSpendingMandateStatus string

// List of values that V2MoneyManagementSpendingMandateStatus can take
const (
	V2MoneyManagementSpendingMandateStatusActive              V2MoneyManagementSpendingMandateStatus = "active"
	V2MoneyManagementSpendingMandateStatusCanceled            V2MoneyManagementSpendingMandateStatus = "canceled"
	V2MoneyManagementSpendingMandateStatusExpired             V2MoneyManagementSpendingMandateStatus = "expired"
	V2MoneyManagementSpendingMandateStatusPendingCancellation V2MoneyManagementSpendingMandateStatus = "pending_cancellation"
)

// The `canceled` status reason.
type V2MoneyManagementSpendingMandateStatusDetailsCanceledReason string

// List of values that V2MoneyManagementSpendingMandateStatusDetailsCanceledReason can take
const (
	V2MoneyManagementSpendingMandateStatusDetailsCanceledReasonCanceledByBeneficiary V2MoneyManagementSpendingMandateStatusDetailsCanceledReason = "canceled_by_beneficiary"
	V2MoneyManagementSpendingMandateStatusDetailsCanceledReasonCanceledByStripe      V2MoneyManagementSpendingMandateStatusDetailsCanceledReason = "canceled_by_stripe"
	V2MoneyManagementSpendingMandateStatusDetailsCanceledReasonUserAction            V2MoneyManagementSpendingMandateStatusDetailsCanceledReason = "user_action"
)

// The type of the SpendingMandate.
type V2MoneyManagementSpendingMandateType string

// List of values that V2MoneyManagementSpendingMandateType can take
const (
	V2MoneyManagementSpendingMandateTypeBankTransfer V2MoneyManagementSpendingMandateType = "bank_transfer"
)

// Bank transfer details. Present if `type` is `bank_transfer`.
type V2MoneyManagementSpendingMandateBankTransfer struct {
	// The name of the account holder that initiated the debit.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The financial address associated with this mandate.
	FinancialAddress string `json:"financial_address"`
	// The bank transfer network for this mandate.
	Network V2MoneyManagementSpendingMandateBankTransferNetwork `json:"network"`
	// The bank transfer reference provided by the bank. Only present when `network` is `bacs`.
	Reference string `json:"reference,omitempty"`
}

// If the mandate is canceled, this field provides more details on the cancellation reason.
type V2MoneyManagementSpendingMandateStatusDetailsCanceled struct {
	// The `canceled` status reason.
	Reason V2MoneyManagementSpendingMandateStatusDetailsCanceledReason `json:"reason"`
}

// Detailed information about the status of the SpendingMandate.
type V2MoneyManagementSpendingMandateStatusDetails struct {
	// If the mandate is canceled, this field provides more details on the cancellation reason.
	Canceled *V2MoneyManagementSpendingMandateStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Timestamps describing when the mandate changed status.
type V2MoneyManagementSpendingMandateStatusTransitions struct {
	// Timestamp describing when the SpendingMandate changed status to `active`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	// Timestamp describing when the SpendingMandate changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp describing when the SpendingMandate was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Timestamp describing when the SpendingMandate changed status to `expired`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Timestamp describing when the SpendingMandate changed status to `pending_cancellation`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	PendingCancellationAt time.Time `json:"pending_cancellation_at,omitempty"`
}

// A SpendingMandate authorizes a counterparty to pull money from a Financial Account.
type V2MoneyManagementSpendingMandate struct {
	APIResource
	// Bank transfer details. Present if `type` is `bank_transfer`.
	BankTransfer *V2MoneyManagementSpendingMandateBankTransfer `json:"bank_transfer,omitempty"`
	// The time at which the SpendingMandate was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	Created time.Time `json:"created"`
	// The currency in ISO 4217 format that debits will be collected in.
	Currency Currency `json:"currency"`
	// Financial account ID associated with this mandate.
	FinancialAccount string `json:"financial_account"`
	// The unique identifier for the SpendingMandate.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the SpendingMandate.
	Status V2MoneyManagementSpendingMandateStatus `json:"status"`
	// Detailed information about the status of the SpendingMandate.
	StatusDetails *V2MoneyManagementSpendingMandateStatusDetails `json:"status_details,omitempty"`
	// Timestamps describing when the mandate changed status.
	StatusTransitions *V2MoneyManagementSpendingMandateStatusTransitions `json:"status_transitions,omitempty"`
	// The type of the SpendingMandate.
	Type V2MoneyManagementSpendingMandateType `json:"type"`
}
