//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. A descriptive category used to classify the Transaction.
type V2MoneyManagementTransactionCategory string

// List of values that V2MoneyManagementTransactionCategory can take
const (
	V2MoneyManagementTransactionCategoryAdjustment         V2MoneyManagementTransactionCategory = "adjustment"
	V2MoneyManagementTransactionCategoryCurrencyConversion V2MoneyManagementTransactionCategory = "currency_conversion"
	V2MoneyManagementTransactionCategoryInboundTransfer    V2MoneyManagementTransactionCategory = "inbound_transfer"
	V2MoneyManagementTransactionCategoryOutboundPayment    V2MoneyManagementTransactionCategory = "outbound_payment"
	V2MoneyManagementTransactionCategoryOutboundTransfer   V2MoneyManagementTransactionCategory = "outbound_transfer"
	V2MoneyManagementTransactionCategoryReceivedCredit     V2MoneyManagementTransactionCategory = "received_credit"
	V2MoneyManagementTransactionCategoryReceivedDebit      V2MoneyManagementTransactionCategory = "received_debit"
	V2MoneyManagementTransactionCategoryReturn             V2MoneyManagementTransactionCategory = "return"
	V2MoneyManagementTransactionCategoryStripeFee          V2MoneyManagementTransactionCategory = "stripe_fee"
)

// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
type V2MoneyManagementTransactionFlowType string

// List of values that V2MoneyManagementTransactionFlowType can take
const (
	V2MoneyManagementTransactionFlowTypeAdjustment         V2MoneyManagementTransactionFlowType = "adjustment"
	V2MoneyManagementTransactionFlowTypeCurrencyConversion V2MoneyManagementTransactionFlowType = "currency_conversion"
	V2MoneyManagementTransactionFlowTypeFeeTransaction     V2MoneyManagementTransactionFlowType = "fee_transaction"
	V2MoneyManagementTransactionFlowTypeInboundTransfer    V2MoneyManagementTransactionFlowType = "inbound_transfer"
	V2MoneyManagementTransactionFlowTypeOutboundPayment    V2MoneyManagementTransactionFlowType = "outbound_payment"
	V2MoneyManagementTransactionFlowTypeOutboundTransfer   V2MoneyManagementTransactionFlowType = "outbound_transfer"
	V2MoneyManagementTransactionFlowTypeReceivedCredit     V2MoneyManagementTransactionFlowType = "received_credit"
	V2MoneyManagementTransactionFlowTypeReceivedDebit      V2MoneyManagementTransactionFlowType = "received_debit"
)

// Closed Enum. Current status of the Transaction.
// A Transaction is `pending` if either `balance_impact.inbound_pending` or `balance_impact.outbound_pending` is non-zero.
// A Transaction is `posted` if only `balance_impact.available` is non-zero.
// A Transaction is `void` if there is no balance impact.
// `posted` and `void` are terminal states, and no additional entries will be added to the Transaction.
type V2MoneyManagementTransactionStatus string

// List of values that V2MoneyManagementTransactionStatus can take
const (
	V2MoneyManagementTransactionStatusPending V2MoneyManagementTransactionStatus = "pending"
	V2MoneyManagementTransactionStatusPosted  V2MoneyManagementTransactionStatus = "posted"
	V2MoneyManagementTransactionStatusVoid    V2MoneyManagementTransactionStatus = "void"
)

// The delta to the FinancialAccount's balance. The balance_impact for the Transaction is equal to sum of its
// TransactionEntries that have `effective_at`s in the past.
type V2MoneyManagementTransactionBalanceImpact struct {
	// Impact to the available balance.
	Available Amount `json:"available"`
	// Impact to the inbound_pending balance.
	InboundPending Amount `json:"inbound_pending"`
	// Impact to the outbound_pending balance.
	OutboundPending Amount `json:"outbound_pending"`
}

// Details about the Flow object that created the Transaction.
type V2MoneyManagementTransactionFlow struct {
	// If applicable, the ID of the Adjustment that created this Transaction.
	Adjustment string `json:"adjustment"`
	// In the future, this will be the ID of the currency conversion that created this Transaction. For now, this field is always null.
	CurrencyConversion string `json:"currency_conversion"`
	// If applicable, the ID of the FeeTransaction that created this Transaction.
	FeeTransaction string `json:"fee_transaction"`
	// If applicable, the ID of the InboundTransfer that created this Transaction.
	InboundTransfer string `json:"inbound_transfer"`
	// If applicable, the ID of the OutboundPayment that created this Transaction.
	OutboundPayment string `json:"outbound_payment"`
	// If applicable, the ID of the OutboundTransfer that created this Transaction.
	OutboundTransfer string `json:"outbound_transfer"`
	// If applicable, the ID of the ReceivedCredit that created this Transaction.
	ReceivedCredit string `json:"received_credit"`
	// If applicable, the ID of the ReceivedDebit that created this Transaction.
	ReceivedDebit string `json:"received_debit"`
	// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
	Type V2MoneyManagementTransactionFlowType `json:"type"`
}

// Timestamps for when the Transaction transitioned to a particular status.
type V2MoneyManagementTransactionStatusTransitions struct {
	// The time at which the Transaction became posted. Only present if status == posted.
	PostedAt time.Time `json:"posted_at"`
	// The time at which the Transaction became void. Only present if status == void.
	VoidAt time.Time `json:"void_at"`
}

// Use Transactions to view changes to your FinancialAccount balance over time. Every flow that moves money, such as OutboundPayments or ReceivedCredits, will have one or more Transactions that describes how the flow impacted your balance. Note that while the FinancialAccount balance will always be up to date, be aware that Transactions and TransactionEntries are created shortly after to reflect changes.
type V2MoneyManagementTransaction struct {
	APIResource
	// The amount of the Transaction.
	Amount Amount `json:"amount"`
	// The delta to the FinancialAccount's balance. The balance_impact for the Transaction is equal to sum of its
	// TransactionEntries that have `effective_at`s in the past.
	BalanceImpact *V2MoneyManagementTransactionBalanceImpact `json:"balance_impact"`
	// Open Enum. A descriptive category used to classify the Transaction.
	Category V2MoneyManagementTransactionCategory `json:"category"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Indicates the FinancialAccount affected by this Transaction.
	FinancialAccount string `json:"financial_account"`
	// Details about the Flow object that created the Transaction.
	Flow *V2MoneyManagementTransactionFlow `json:"flow"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. Current status of the Transaction.
	// A Transaction is `pending` if either `balance_impact.inbound_pending` or `balance_impact.outbound_pending` is non-zero.
	// A Transaction is `posted` if only `balance_impact.available` is non-zero.
	// A Transaction is `void` if there is no balance impact.
	// `posted` and `void` are terminal states, and no additional entries will be added to the Transaction.
	Status V2MoneyManagementTransactionStatus `json:"status"`
	// Timestamps for when the Transaction transitioned to a particular status.
	StatusTransitions *V2MoneyManagementTransactionStatusTransitions `json:"status_transitions"`
}
