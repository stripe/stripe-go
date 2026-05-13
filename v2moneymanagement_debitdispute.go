//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The bank network the dispute was originated on.
type V2MoneyManagementDebitDisputeBankTransferNetwork string

// List of values that V2MoneyManagementDebitDisputeBankTransferNetwork can take
const (
	V2MoneyManagementDebitDisputeBankTransferNetworkACH V2MoneyManagementDebitDisputeBankTransferNetwork = "ach"
)

// The reason for the dispute.
type V2MoneyManagementDebitDisputeBankTransferReason string

// List of values that V2MoneyManagementDebitDisputeBankTransferReason can take
const (
	V2MoneyManagementDebitDisputeBankTransferReasonIncorrectAmountOrDate V2MoneyManagementDebitDisputeBankTransferReason = "incorrect_amount_or_date"
	V2MoneyManagementDebitDisputeBankTransferReasonUnauthorized          V2MoneyManagementDebitDisputeBankTransferReason = "unauthorized"
)

// The status of the DebitDispute.
type V2MoneyManagementDebitDisputeStatus string

// List of values that V2MoneyManagementDebitDisputeStatus can take
const (
	V2MoneyManagementDebitDisputeStatusFailed    V2MoneyManagementDebitDisputeStatus = "failed"
	V2MoneyManagementDebitDisputeStatusSubmitted V2MoneyManagementDebitDisputeStatus = "submitted"
	V2MoneyManagementDebitDisputeStatusSucceeded V2MoneyManagementDebitDisputeStatus = "succeeded"
)

// The reason for the failure of the DebitDispute.
type V2MoneyManagementDebitDisputeStatusDetailsFailedReason string

// List of values that V2MoneyManagementDebitDisputeStatusDetailsFailedReason can take
const (
	V2MoneyManagementDebitDisputeStatusDetailsFailedReasonUnknown V2MoneyManagementDebitDisputeStatusDetailsFailedReason = "unknown"
)

// The type of the DebitDispute.
type V2MoneyManagementDebitDisputeType string

// List of values that V2MoneyManagementDebitDisputeType can take
const (
	V2MoneyManagementDebitDisputeTypeBankTransfer V2MoneyManagementDebitDisputeType = "bank_transfer"
)

// Details about the bank transfer dispute. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementDebitDisputeBankTransfer struct {
	// The bank network the dispute was originated on.
	Network V2MoneyManagementDebitDisputeBankTransferNetwork `json:"network"`
	// The reason for the dispute.
	Reason V2MoneyManagementDebitDisputeBankTransferReason `json:"reason,omitempty"`
	// The statement descriptor set by the originator of the debit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}

// Information that elaborates on the `failed` status of a DebitDispute.
// It is only present when the DebitDispute status is `failed`.
type V2MoneyManagementDebitDisputeStatusDetailsFailed struct {
	// The reason for the failure of the DebitDispute.
	Reason V2MoneyManagementDebitDisputeStatusDetailsFailedReason `json:"reason"`
}

// Detailed information about the status of the DebitDispute.
type V2MoneyManagementDebitDisputeStatusDetails struct {
	// Information that elaborates on the `failed` status of a DebitDispute.
	// It is only present when the DebitDispute status is `failed`.
	Failed *V2MoneyManagementDebitDisputeStatusDetailsFailed `json:"failed"`
}

// The time at which the DebitDispute transitioned to a particular status.
type V2MoneyManagementDebitDisputeStatusTransitions struct {
	// The time when the DebitDispute was marked as `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2026-04-23T13:22:18.123Z`.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// The time when the DebitDispute was marked as `succeeded`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2026-04-23T13:22:18.123Z`.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// A DebitDispute represents a dispute raised against a received debit.
type V2MoneyManagementDebitDispute struct {
	APIResource
	// The amount of the DebitDispute.
	Amount Amount `json:"amount"`
	// Details about the bank transfer dispute. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementDebitDisputeBankTransfer `json:"bank_transfer,omitempty"`
	// Time at which the DebitDispute was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2026-03-23T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// The Financial Account associated with the DebitDispute.
	FinancialAccount string `json:"financial_account"`
	// The ID of a DebitDispute.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the ReceivedDebit that was disputed.
	ReceivedDebit string `json:"received_debit"`
	// The status of the DebitDispute.
	Status V2MoneyManagementDebitDisputeStatus `json:"status"`
	// Detailed information about the status of the DebitDispute.
	StatusDetails *V2MoneyManagementDebitDisputeStatusDetails `json:"status_details,omitempty"`
	// The time at which the DebitDispute transitioned to a particular status.
	StatusTransitions *V2MoneyManagementDebitDisputeStatusTransitions `json:"status_transitions,omitempty"`
	// The type of the DebitDispute.
	Type V2MoneyManagementDebitDisputeType `json:"type"`
}
