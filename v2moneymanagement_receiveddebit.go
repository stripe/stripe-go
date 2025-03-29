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
	V2MoneyManagementReceivedDebitTypeCardSpend     V2MoneyManagementReceivedDebitType = "card_spend"
	V2MoneyManagementReceivedDebitTypeExternalDebit V2MoneyManagementReceivedDebitType = "external_debit"
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
	CanceledAt time.Time `json:"canceled_at"`
	// The time when the ReceivedDebit was marked as `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	FailedAt time.Time `json:"failed_at"`
	// The time when the ReceivedDebit was marked as `succeeded`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	SucceededAt time.Time `json:"succeeded_at"`
}

// The payment method used to originate the debit.
type V2MoneyManagementReceivedDebitBankTransferUSBankAccount struct {
	// The name of the bank the debit originated from.
	BankName string `json:"bank_name"`
	// Open Enum. The bank network the debit was originated on.
	Network V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork `json:"network"`
	// The routing number of the bank that originated the debit.
	RoutingNumber string `json:"routing_number"`
}

// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementReceivedDebitBankTransfer struct {
	// The Financial Address that was debited.
	FinancialAddress string `json:"financial_address"`
	// Open Enum. The type of the payment method used to originate the debit.
	PaymentMethodType V2MoneyManagementReceivedDebitBankTransferPaymentMethodType `json:"payment_method_type"`
	// The statement descriptor set by the originator of the debit.
	StatementDescriptor string `json:"statement_descriptor"`
	// The payment method used to originate the debit.
	USBankAccount *V2MoneyManagementReceivedDebitBankTransferUSBankAccount `json:"us_bank_account"`
}

// The Issuing Authorization for this card_spend. Contains the reference id and the amount.
type V2MoneyManagementReceivedDebitCardSpendAuthorization struct {
	// Amount associated with this issuing authorization.
	Amount Amount `json:"amount"`
	// The reference to the v1 issuing authorization ID.
	IssuingAuthorizationV1 string `json:"issuing_authorization_v1"`
}

// The list of card spend transactions. These contain the transaction reference ID and the amount.
type V2MoneyManagementReceivedDebitCardSpendCardTransaction struct {
	// Amount associated with this issuing transaction.
	Amount Amount `json:"amount"`
	// The reference to the v1 issuing transaction ID.
	IssuingTransactionV1 string `json:"issuing_transaction_v1"`
}

// This object stores details about the issuing transactions that resulted in the ReceivedDebit. Present if `type` field value is `card_spend`.
type V2MoneyManagementReceivedDebitCardSpend struct {
	// The Issuing Authorization for this card_spend. Contains the reference id and the amount.
	Authorization *V2MoneyManagementReceivedDebitCardSpendAuthorization `json:"authorization"`
	// The list of card spend transactions. These contain the transaction reference ID and the amount.
	CardTransactions []*V2MoneyManagementReceivedDebitCardSpendCardTransaction `json:"card_transactions"`
	// The reference to the card object that resulted in the debit.
	CardV1ID string `json:"card_v1_id"`
}

// ReceivedDebit resource
type V2MoneyManagementReceivedDebit struct {
	APIResource
	// Amount and currency of the ReceivedDebit.
	Amount Amount `json:"amount"`
	// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedDebitBankTransfer `json:"bank_transfer"`
	// This object stores details about the issuing transactions that resulted in the ReceivedDebit. Present if `type` field value is `card_spend`.
	CardSpend *V2MoneyManagementReceivedDebitCardSpend `json:"card_spend"`
	// The time at which the ReceivedDebit was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Freeform string sent by the originator of the ReceivedDebit.
	Description string `json:"description"`
	// Financial Account on which funds for ReceivedDebit were debited.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the ReceivedDebit.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A link to the Stripe-hosted receipt for this ReceivedDebit.
	ReceiptURL string `json:"receipt_url"`
	// Open Enum. The status of the ReceivedDebit.
	Status V2MoneyManagementReceivedDebitStatus `json:"status"`
	// Detailed information about the status of the ReceivedDebit.
	StatusDetails *V2MoneyManagementReceivedDebitStatusDetails `json:"status_details"`
	// The time at which the ReceivedDebit transitioned to a particular status.
	StatusTransitions *V2MoneyManagementReceivedDebitStatusTransitions `json:"status_transitions"`
	// Open Enum. The type of the ReceivedDebit.
	Type V2MoneyManagementReceivedDebitType `json:"type"`
}
