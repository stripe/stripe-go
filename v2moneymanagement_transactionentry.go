//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The category of this TransactionEntry's parent Transaction.
type V2MoneyManagementTransactionEntryTransactionDetailsCategory string

// List of values that V2MoneyManagementTransactionEntryTransactionDetailsCategory can take
const (
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryAdjustment               V2MoneyManagementTransactionEntryTransactionDetailsCategory = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryCurrencyConversion       V2MoneyManagementTransactionEntryTransactionDetailsCategory = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryInboundTransfer          V2MoneyManagementTransactionEntryTransactionDetailsCategory = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryInboundTransferReversal  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "inbound_transfer_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundPayment          V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundPaymentReversal  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_payment_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundTransfer         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundTransferReversal V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_transfer_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedCredit           V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedCreditReversal   V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_credit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedDebit            V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedDebitReversal    V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_debit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeFee                V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_fee"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeFeeTax             V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_fee_tax"
)

// Open Enum. The type of flow that created the Transaction. When available, the flow's ID is provided in the corresponding field.
type V2MoneyManagementTransactionEntryTransactionDetailsFlowType string

// List of values that V2MoneyManagementTransactionEntryTransactionDetailsFlowType can take
const (
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeAdjustment         V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeCurrencyConversion V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeFeeBatch           V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "fee_batch"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeInboundTransfer    V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundPayment    V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundTransfer   V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedCredit     V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedDebit      V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_debit"
)

// The change to the FinancialAccount's balance that takes effect at this TransactionEntry's `effective_at` timestamp.
type V2MoneyManagementTransactionEntryBalanceImpact struct {
	// The change to the available balance.
	Available Amount `json:"available"`
	// The change to the `inbound_pending` balance.
	InboundPending Amount `json:"inbound_pending"`
	// The change to the `outbound_pending` balance.
	OutboundPending Amount `json:"outbound_pending"`
}

// Details about the flow object that created the Transaction this entry belongs to, such as an
// InboundTransfer or OutboundPayment. `transaction_details.flow.type` identifies the kind of object.
// The corresponding field contains its ID, when available: for an InboundTransfer,
// `transaction_details.flow.inbound_transfer` contains the InboundTransfer's ID. To list entries for
// the parent Transaction, pass the TransactionEntry's `transaction` ID to the `transaction` parameter
// on List Transaction Entries.
type V2MoneyManagementTransactionEntryTransactionDetailsFlow struct {
	// If applicable, the ID of the Adjustment that created this Transaction.
	Adjustment string `json:"adjustment,omitempty"`
	// When available, the ID of the CurrencyConversion that created this Transaction.
	CurrencyConversion string `json:"currency_conversion,omitempty"`
	// If applicable, the ID of the FeeBatch that created this Transaction.
	FeeBatch string `json:"fee_batch,omitempty"`
	// If applicable, the ID of the InboundTransfer that created this Transaction.
	InboundTransfer string `json:"inbound_transfer,omitempty"`
	// If applicable, the ID of the OutboundPayment that created this Transaction.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// If applicable, the ID of the OutboundTransfer that created this Transaction.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// If applicable, the ID of the ReceivedCredit that created this Transaction.
	ReceivedCredit string `json:"received_credit,omitempty"`
	// If applicable, the ID of the ReceivedDebit that created this Transaction.
	ReceivedDebit string `json:"received_debit,omitempty"`
	// Open Enum. The type of flow that created the Transaction. When available, the flow's ID is provided in the corresponding field.
	Type V2MoneyManagementTransactionEntryTransactionDetailsFlowType `json:"type"`
}

// Details copied from the Transaction that this TransactionEntry belongs to.
type V2MoneyManagementTransactionEntryTransactionDetails struct {
	// Open Enum. The category of this TransactionEntry's parent Transaction.
	Category V2MoneyManagementTransactionEntryTransactionDetailsCategory `json:"category"`
	// The ID of the FinancialAccount affected by this Transaction.
	FinancialAccount string `json:"financial_account"`
	// Details about the flow object that created the Transaction this entry belongs to, such as an
	// InboundTransfer or OutboundPayment. `transaction_details.flow.type` identifies the kind of object.
	// The corresponding field contains its ID, when available: for an InboundTransfer,
	// `transaction_details.flow.inbound_transfer` contains the InboundTransfer's ID. To list entries for
	// the parent Transaction, pass the TransactionEntry's `transaction` ID to the `transaction` parameter
	// on List Transaction Entries.
	Flow *V2MoneyManagementTransactionEntryTransactionDetailsFlow `json:"flow,omitempty"`
}

// TransactionEntries model funds moving between availability states for a given Transaction.
type V2MoneyManagementTransactionEntry struct {
	APIResource
	// The change to the FinancialAccount's balance that takes effect at this TransactionEntry's `effective_at` timestamp.
	BalanceImpact *V2MoneyManagementTransactionEntryBalanceImpact `json:"balance_impact"`
	// Time at which the object was created. Represented as an RFC 3339 timestamp in UTC with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Time at which the TransactionEntry affects the FinancialAccount balance. This can be in the past or future.
	EffectiveAt time.Time `json:"effective_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Transaction that this TransactionEntry belongs to.
	Transaction string `json:"transaction"`
	// Details copied from the Transaction that this TransactionEntry belongs to.
	TransactionDetails *V2MoneyManagementTransactionEntryTransactionDetails `json:"transaction_details"`
}
