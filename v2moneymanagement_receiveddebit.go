//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The status of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatus string

// List of values that V2MoneyManagementReceivedDebitStatus can take
const (
	V2MoneyManagementReceivedDebitStatusCanceled  V2MoneyManagementReceivedDebitStatus = "canceled"
	V2MoneyManagementReceivedDebitStatusFailed    V2MoneyManagementReceivedDebitStatus = "failed"
	V2MoneyManagementReceivedDebitStatusPending   V2MoneyManagementReceivedDebitStatus = "pending"
	V2MoneyManagementReceivedDebitStatusReturned  V2MoneyManagementReceivedDebitStatus = "returned"
	V2MoneyManagementReceivedDebitStatusSucceeded V2MoneyManagementReceivedDebitStatus = "succeeded"
)

// Open Enum. The reason for the failure of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatusDetailsFailedReason string

// List of values that V2MoneyManagementReceivedDebitStatusDetailsFailedReason can take
const (
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonFinancialAddressInactive V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "financial_address_inactive"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonInsufficientFunds        V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "insufficient_funds"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonStripeRejected           V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "stripe_rejected"
)

// Open Enum. The type of the ReceivedDebit.
type V2MoneyManagementReceivedDebitType string

// List of values that V2MoneyManagementReceivedDebitType can take
const (
	V2MoneyManagementReceivedDebitTypeBankTransfer  V2MoneyManagementReceivedDebitType = "bank_transfer"
	V2MoneyManagementReceivedDebitTypeExternalDebit V2MoneyManagementReceivedDebitType = "external_debit"
)

// Open Enum. Indicates the origin type through which this debit was initiated.
type V2MoneyManagementReceivedDebitBankTransferOriginType string

// List of values that V2MoneyManagementReceivedDebitBankTransferOriginType can take
const (
	V2MoneyManagementReceivedDebitBankTransferOriginTypeUSBankAccount V2MoneyManagementReceivedDebitBankTransferOriginType = "us_bank_account"
)

// Open Enum. The type of the payment method used to originate the debit.
type V2MoneyManagementReceivedDebitBankTransferPaymentMethodType string

// List of values that V2MoneyManagementReceivedDebitBankTransferPaymentMethodType can take
const (
	V2MoneyManagementReceivedDebitBankTransferPaymentMethodTypeUSBankAccount V2MoneyManagementReceivedDebitBankTransferPaymentMethodType = "us_bank_account"
)

// Open Enum. The bank network the debit was originated on.
type V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork string

// List of values that V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork can take
const (
	V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetworkACH V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork = "ach"
)

// Information that elaborates on the `failed` status of a ReceivedDebit.
// It is only present when the ReceivedDebit status is `failed`.
type V2MoneyManagementReceivedDebitStatusDetailsFailed struct {
	// Open Enum. The reason for the failure of the ReceivedDebit.
	Reason V2MoneyManagementReceivedDebitStatusDetailsFailedReason `json:"reason"`
}

// Detailed information about the status of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatusDetails struct {
	// Information that elaborates on the `failed` status of a ReceivedDebit.
	// It is only present when the ReceivedDebit status is `failed`.
	Failed *V2MoneyManagementReceivedDebitStatusDetailsFailed `json:"failed"`
}

// The time at which the ReceivedDebit transitioned to a particular status.
type V2MoneyManagementReceivedDebitStatusTransitions struct {
	// The time when the ReceivedDebit was marked as `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// The time when the ReceivedDebit was marked as `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// The time when the ReceivedDebit was marked as `succeeded`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// The payment method used to originate the debit.
type V2MoneyManagementReceivedDebitBankTransferUSBankAccount struct {
	// The name of the bank the debit originated from.
	BankName string `json:"bank_name,omitempty"`
	// Open Enum. The bank network the debit was originated on.
	Network V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork `json:"network"`
	// The routing number of the bank that originated the debit.
	RoutingNumber string `json:"routing_number,omitempty"`
}

// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementReceivedDebitBankTransfer struct {
	// The Financial Address that was debited.
	FinancialAddress string `json:"financial_address"`
	// Open Enum. Indicates the origin type through which this debit was initiated.
	OriginType V2MoneyManagementReceivedDebitBankTransferOriginType `json:"origin_type"`
	// Open Enum. The type of the payment method used to originate the debit.
	PaymentMethodType V2MoneyManagementReceivedDebitBankTransferPaymentMethodType `json:"payment_method_type"`
	// The statement descriptor set by the originator of the debit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// The payment method used to originate the debit.
	USBankAccount *V2MoneyManagementReceivedDebitBankTransferUSBankAccount `json:"us_bank_account"`
}

// ReceivedDebit resource
type V2MoneyManagementReceivedDebit struct {
	APIResource
	// Amount and currency of the ReceivedDebit.
	Amount Amount `json:"amount"`
	// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedDebitBankTransfer `json:"bank_transfer,omitempty"`
	// The time at which the ReceivedDebit was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Freeform string sent by the originator of the ReceivedDebit.
	Description string `json:"description,omitempty"`
	// Financial Account on which funds for ReceivedDebit were debited.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the ReceivedDebit.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A link to the Stripe-hosted receipt for this ReceivedDebit.
	ReceiptURL string `json:"receipt_url,omitempty"`
	// Open Enum. The status of the ReceivedDebit.
	Status V2MoneyManagementReceivedDebitStatus `json:"status"`
	// Detailed information about the status of the ReceivedDebit.
	StatusDetails *V2MoneyManagementReceivedDebitStatusDetails `json:"status_details,omitempty"`
	// The time at which the ReceivedDebit transitioned to a particular status.
	StatusTransitions *V2MoneyManagementReceivedDebitStatusTransitions `json:"status_transitions,omitempty"`
	// Open Enum. The type of the ReceivedDebit.
	Type V2MoneyManagementReceivedDebitType `json:"type"`
}
