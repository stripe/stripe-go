//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The Level of the HistoryEntry.
type V2MoneyManagementInboundTransferTransferHistoryLevel string

// List of values that V2MoneyManagementInboundTransferTransferHistoryLevel can take
const (
	V2MoneyManagementInboundTransferTransferHistoryLevelCanonical V2MoneyManagementInboundTransferTransferHistoryLevel = "canonical"
	V2MoneyManagementInboundTransferTransferHistoryLevelDebug     V2MoneyManagementInboundTransferTransferHistoryLevel = "debug"
)

// Open Enum. The type of the HistoryEntry.
type V2MoneyManagementInboundTransferTransferHistoryType string

// List of values that V2MoneyManagementInboundTransferTransferHistoryType can take
const (
	V2MoneyManagementInboundTransferTransferHistoryTypeBankDebitFailed     V2MoneyManagementInboundTransferTransferHistoryType = "bank_debit_failed"
	V2MoneyManagementInboundTransferTransferHistoryTypeBankDebitProcessing V2MoneyManagementInboundTransferTransferHistoryType = "bank_debit_processing"
	V2MoneyManagementInboundTransferTransferHistoryTypeBankDebitQueued     V2MoneyManagementInboundTransferTransferHistoryType = "bank_debit_queued"
	V2MoneyManagementInboundTransferTransferHistoryTypeBankDebitReturned   V2MoneyManagementInboundTransferTransferHistoryType = "bank_debit_returned"
	V2MoneyManagementInboundTransferTransferHistoryTypeBankDebitSucceeded  V2MoneyManagementInboundTransferTransferHistoryType = "bank_debit_succeeded"
)

// Open Enum. The return reason for the failed InboundTransfer.
type V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason string

// List of values that V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason can take
const (
	V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReasonBankAccountClosed            V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason = "bank_account_closed"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReasonBankAccountNotFound          V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason = "bank_account_not_found"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReasonBankDebitCouldNotBeProcessed V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason = "bank_debit_could_not_be_processed"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReasonBankDebitNotAuthorized       V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason = "bank_debit_not_authorized"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReasonInsufficientFunds            V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason = "insufficient_funds"
)

// Open Enum. The return reason for the returned InboundTransfer.
type V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason string

// List of values that V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason can take
const (
	V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReasonBankAccountClosed            V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason = "bank_account_closed"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReasonBankAccountNotFound          V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason = "bank_account_not_found"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReasonBankDebitCouldNotBeProcessed V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason = "bank_debit_could_not_be_processed"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReasonBankDebitNotAuthorized       V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason = "bank_debit_not_authorized"
	V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReasonInsufficientFunds            V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason = "insufficient_funds"
)

// The amount in specified currency that will land in the FinancialAccount balance.
type V2MoneyManagementInboundTransferAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The amount in specified currency that was debited from the Payment Method.
type V2MoneyManagementInboundTransferFromDebited struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The Payment Method object used to create the InboundTransfer.
type V2MoneyManagementInboundTransferFromPaymentMethod struct {
	// The type of object this destination represents. For a us bank account, we expect us_bank_account.
	Type string `json:"type"`
	// The destination US bank account identifier. eg "usba_***".
	USBankAccount string `json:"us_bank_account,omitempty"`
}

// A nested object containing information about the origin of the InboundTransfer.
type V2MoneyManagementInboundTransferFrom struct {
	// The amount in specified currency that was debited from the Payment Method.
	Debited *V2MoneyManagementInboundTransferFromDebited `json:"debited"`
	// The Payment Method object used to create the InboundTransfer.
	PaymentMethod *V2MoneyManagementInboundTransferFromPaymentMethod `json:"payment_method"`
}

// The amount by which the FinancialAccount balance is credited.
type V2MoneyManagementInboundTransferToCredited struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// A nested object containing information about the destination of the InboundTransfer.
type V2MoneyManagementInboundTransferTo struct {
	// The amount by which the FinancialAccount balance is credited.
	Credited *V2MoneyManagementInboundTransferToCredited `json:"credited"`
	// The FinancialAccount that funds will land in.
	FinancialAccount string `json:"financial_account"`
}

// The history entry for a failed InboundTransfer.
type V2MoneyManagementInboundTransferTransferHistoryBankDebitFailed struct {
	// Open Enum. The return reason for the failed InboundTransfer.
	FailureReason V2MoneyManagementInboundTransferTransferHistoryBankDebitFailedFailureReason `json:"failure_reason"`
}

// The history entry for a returned InboundTransfer.
type V2MoneyManagementInboundTransferTransferHistoryBankDebitReturned struct {
	// Open Enum. The return reason for the returned InboundTransfer.
	ReturnReason V2MoneyManagementInboundTransferTransferHistoryBankDebitReturnedReturnReason `json:"return_reason"`
}

// A list of history objects, representing changes in the state of the InboundTransfer.
type V2MoneyManagementInboundTransferTransferHistory struct {
	// The history entry for a failed InboundTransfer.
	BankDebitFailed *V2MoneyManagementInboundTransferTransferHistoryBankDebitFailed `json:"bank_debit_failed,omitempty"`
	// The history entry for a processing InboundTransfer.
	BankDebitProcessing map[string]any `json:"bank_debit_processing,omitempty"`
	// The history entry for a queued InboundTransfer.
	BankDebitQueued map[string]any `json:"bank_debit_queued,omitempty"`
	// The history entry for a returned InboundTransfer.
	BankDebitReturned *V2MoneyManagementInboundTransferTransferHistoryBankDebitReturned `json:"bank_debit_returned,omitempty"`
	// The history entry for a succeeded InboundTransfer.
	BankDebitSucceeded map[string]any `json:"bank_debit_succeeded,omitempty"`
	// Creation time of the HistoryEntry in RFC 3339 format and UTC.
	Created time.Time `json:"created"`
	// Effective at time of the HistoryEntry in RFC 3339 format and UTC.
	EffectiveAt time.Time `json:"effective_at"`
	// A unique ID for the HistoryEntry.
	ID string `json:"id"`
	// Open Enum. The Level of the HistoryEntry.
	Level V2MoneyManagementInboundTransferTransferHistoryLevel `json:"level"`
	// Open Enum. The type of the HistoryEntry.
	Type V2MoneyManagementInboundTransferTransferHistoryType `json:"type"`
}

// An InboundTransfer object, representing a money movement from a
// user owned PaymentMethod to a FinancialAccount belonging to the same user.
type V2MoneyManagementInboundTransfer struct {
	APIResource
	// The amount in specified currency that will land in the FinancialAccount balance.
	Amount *V2MoneyManagementInboundTransferAmount `json:"amount"`
	// Creation time of the InboundTransfer. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// A freeform text field provided by user, containing metadata.
	Description string `json:"description"`
	// A nested object containing information about the origin of the InboundTransfer.
	From *V2MoneyManagementInboundTransferFrom `json:"from"`
	// Unique identifier for the InboundTransfer.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A hosted transaction receipt URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	ReceiptURL string `json:"receipt_url,omitempty"`
	// A nested object containing information about the destination of the InboundTransfer.
	To *V2MoneyManagementInboundTransferTo `json:"to"`
	// A list of history objects, representing changes in the state of the InboundTransfer.
	TransferHistory []*V2MoneyManagementInboundTransferTransferHistory `json:"transfer_history"`
}
