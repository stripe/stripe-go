//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The bank transfer network for this mandate.
type V2MoneyManagementReceivedDebitMandateBankTransferNetwork string

// List of values that V2MoneyManagementReceivedDebitMandateBankTransferNetwork can take
const (
	V2MoneyManagementReceivedDebitMandateBankTransferNetworkBACS V2MoneyManagementReceivedDebitMandateBankTransferNetwork = "bacs"
)

// The status of the ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateStatus string

// List of values that V2MoneyManagementReceivedDebitMandateStatus can take
const (
	V2MoneyManagementReceivedDebitMandateStatusActive              V2MoneyManagementReceivedDebitMandateStatus = "active"
	V2MoneyManagementReceivedDebitMandateStatusCanceled            V2MoneyManagementReceivedDebitMandateStatus = "canceled"
	V2MoneyManagementReceivedDebitMandateStatusExpired             V2MoneyManagementReceivedDebitMandateStatus = "expired"
	V2MoneyManagementReceivedDebitMandateStatusPendingCancellation V2MoneyManagementReceivedDebitMandateStatus = "pending_cancellation"
)

// The `canceled` status reason.
type V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason string

// List of values that V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason can take
const (
	V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReasonCanceledByBeneficiary V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason = "canceled_by_beneficiary"
	V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReasonCanceledByStripe      V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason = "canceled_by_stripe"
	V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReasonUserAction            V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason = "user_action"
)

// The type of the ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateType string

// List of values that V2MoneyManagementReceivedDebitMandateType can take
const (
	V2MoneyManagementReceivedDebitMandateTypeBankTransfer V2MoneyManagementReceivedDebitMandateType = "bank_transfer"
)

// This object stores details about the originating bank transfer that resulted in the ReceivedDebitMandate. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementReceivedDebitMandateBankTransfer struct {
	// The name of the account holder that initiated the debit.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The financial address associated with this mandate.
	FinancialAddress string `json:"financial_address"`
	// The bank transfer network for this mandate.
	Network V2MoneyManagementReceivedDebitMandateBankTransferNetwork `json:"network"`
	// The bank transfer reference provided by the bank.
	Reference string `json:"reference,omitempty"`
}

// If the mandate is canceled, this field provides more details on the cancellation reason.
type V2MoneyManagementReceivedDebitMandateStatusDetailsCanceled struct {
	// The `canceled` status reason.
	Reason V2MoneyManagementReceivedDebitMandateStatusDetailsCanceledReason `json:"reason"`
}

// Detailed information that elaborates on the specific status of the ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateStatusDetails struct {
	// If the mandate is canceled, this field provides more details on the cancellation reason.
	Canceled *V2MoneyManagementReceivedDebitMandateStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Timestamps describing when the mandate changed status.
type V2MoneyManagementReceivedDebitMandateStatusTransitions struct {
	// Timestamp describing when the ReceivedDebitMandate changed status to `active`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	// Timestamp describing when the ReceivedDebitMandate changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision for example: 2026-06-03T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp describing when the ReceivedDebitMandate was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision for example: 2026-06-03T13:22:18.123Z.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Timestamp describing when the ReceivedDebitMandate changed status to `expired`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2026-06-03T13:22:18.123Z.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Timestamp describing when the ReceivedDebitMandate changed status to `pending_cancellation`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision.
	PendingCancellationAt time.Time `json:"pending_cancellation_at,omitempty"`
}

// A ReceivedDebitMandate represents an authorization from a third party to debit a financial account on a recurring basis.
type V2MoneyManagementReceivedDebitMandate struct {
	APIResource
	// This object stores details about the originating bank transfer that resulted in the ReceivedDebitMandate. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedDebitMandateBankTransfer `json:"bank_transfer,omitempty"`
	// The time at which the ReceivedDebitMandate was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2026-06-03T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// The currency of the ReceivedDebitMandate in ISO 4217 format. This is the currency that debits will be collected in.
	Currency Currency `json:"currency"`
	// Financial account ID associated with this mandate.
	FinancialAccount string `json:"financial_account"`
	// The unique identifier for the ReceivedDebitMandate.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the ReceivedDebitMandate.
	Status V2MoneyManagementReceivedDebitMandateStatus `json:"status"`
	// Detailed information that elaborates on the specific status of the ReceivedDebitMandate.
	StatusDetails *V2MoneyManagementReceivedDebitMandateStatusDetails `json:"status_details,omitempty"`
	// Timestamps describing when the mandate changed status.
	StatusTransitions *V2MoneyManagementReceivedDebitMandateStatusTransitions `json:"status_transitions,omitempty"`
	// The type of the ReceivedDebitMandate.
	Type V2MoneyManagementReceivedDebitMandateType `json:"type"`
}
