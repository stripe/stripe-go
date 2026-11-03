//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. A category used to classify the Transaction.
type V2MoneyManagementTransactionCategory string

// List of values that V2MoneyManagementTransactionCategory can take
const (
	V2MoneyManagementTransactionCategoryAdjustment               V2MoneyManagementTransactionCategory = "adjustment"
	V2MoneyManagementTransactionCategoryCurrencyConversion       V2MoneyManagementTransactionCategory = "currency_conversion"
	V2MoneyManagementTransactionCategoryInboundTransfer          V2MoneyManagementTransactionCategory = "inbound_transfer"
	V2MoneyManagementTransactionCategoryInboundTransferReversal  V2MoneyManagementTransactionCategory = "inbound_transfer_reversal"
	V2MoneyManagementTransactionCategoryOutboundPayment          V2MoneyManagementTransactionCategory = "outbound_payment"
	V2MoneyManagementTransactionCategoryOutboundPaymentReversal  V2MoneyManagementTransactionCategory = "outbound_payment_reversal"
	V2MoneyManagementTransactionCategoryOutboundTransfer         V2MoneyManagementTransactionCategory = "outbound_transfer"
	V2MoneyManagementTransactionCategoryOutboundTransferReversal V2MoneyManagementTransactionCategory = "outbound_transfer_reversal"
	V2MoneyManagementTransactionCategoryReceivedCredit           V2MoneyManagementTransactionCategory = "received_credit"
	V2MoneyManagementTransactionCategoryReceivedCreditReversal   V2MoneyManagementTransactionCategory = "received_credit_reversal"
	V2MoneyManagementTransactionCategoryReceivedDebit            V2MoneyManagementTransactionCategory = "received_debit"
	V2MoneyManagementTransactionCategoryReceivedDebitReversal    V2MoneyManagementTransactionCategory = "received_debit_reversal"
	V2MoneyManagementTransactionCategoryStripeFee                V2MoneyManagementTransactionCategory = "stripe_fee"
	V2MoneyManagementTransactionCategoryStripeFeeTax             V2MoneyManagementTransactionCategory = "stripe_fee_tax"
)

// Open Enum. The type of flow that created the Transaction. When available, the flow's ID is provided in the corresponding field.
type V2MoneyManagementTransactionFlowType string

// List of values that V2MoneyManagementTransactionFlowType can take
const (
	V2MoneyManagementTransactionFlowTypeAdjustment         V2MoneyManagementTransactionFlowType = "adjustment"
	V2MoneyManagementTransactionFlowTypeCurrencyConversion V2MoneyManagementTransactionFlowType = "currency_conversion"
	V2MoneyManagementTransactionFlowTypeFeeBatch           V2MoneyManagementTransactionFlowType = "fee_batch"
	V2MoneyManagementTransactionFlowTypeInboundTransfer    V2MoneyManagementTransactionFlowType = "inbound_transfer"
	V2MoneyManagementTransactionFlowTypeOutboundPayment    V2MoneyManagementTransactionFlowType = "outbound_payment"
	V2MoneyManagementTransactionFlowTypeOutboundTransfer   V2MoneyManagementTransactionFlowType = "outbound_transfer"
	V2MoneyManagementTransactionFlowTypeReceivedCredit     V2MoneyManagementTransactionFlowType = "received_credit"
	V2MoneyManagementTransactionFlowTypeReceivedDebit      V2MoneyManagementTransactionFlowType = "received_debit"
)

// Current availability of the regulatory receipt.
type V2MoneyManagementTransactionRegulatoryReceiptStatus string

// List of values that V2MoneyManagementTransactionRegulatoryReceiptStatus can take
const (
	V2MoneyManagementTransactionRegulatoryReceiptStatusAvailable     V2MoneyManagementTransactionRegulatoryReceiptStatus = "available"
	V2MoneyManagementTransactionRegulatoryReceiptStatusNotApplicable V2MoneyManagementTransactionRegulatoryReceiptStatus = "not_applicable"
	V2MoneyManagementTransactionRegulatoryReceiptStatusPending       V2MoneyManagementTransactionRegulatoryReceiptStatus = "pending"
	V2MoneyManagementTransactionRegulatoryReceiptStatusURLExpired    V2MoneyManagementTransactionRegulatoryReceiptStatus = "url_expired"
)

// Closed Enum. The current status of the Transaction.
// A Transaction is `pending` if either `balance_impact.inbound_pending` or `balance_impact.outbound_pending` is non-zero.
// A Transaction is `posted` if only `balance_impact.available` is non-zero.
// A Transaction is `void` if there is no balance impact.
// `posted` and `void` are terminal states, and no additional TransactionEntries will be added to the Transaction.
type V2MoneyManagementTransactionStatus string

// List of values that V2MoneyManagementTransactionStatus can take
const (
	V2MoneyManagementTransactionStatusPending V2MoneyManagementTransactionStatus = "pending"
	V2MoneyManagementTransactionStatusPosted  V2MoneyManagementTransactionStatus = "posted"
	V2MoneyManagementTransactionStatusVoid    V2MoneyManagementTransactionStatus = "void"
)

// The change to the FinancialAccount's balance. This is the sum of the balance impacts of this Transaction's
// TransactionEntries with `effective_at` timestamps in the past.
type V2MoneyManagementTransactionBalanceImpact struct {
	// The change to the available balance.
	Available Amount `json:"available"`
	// The change to the `inbound_pending` balance.
	InboundPending Amount `json:"inbound_pending"`
	// The change to the `outbound_pending` balance.
	OutboundPending Amount `json:"outbound_pending"`
}

// Details about the counterparty to this Transaction.
type V2MoneyManagementTransactionCounterparty struct {
	// Name of the counterparty.
	Name string `json:"name,omitempty"`
}

// Details about the flow object that created the Transaction, such as an InboundTransfer or OutboundPayment.
// `flow.type` identifies the kind of object. The corresponding field contains its ID, when available:
// for an InboundTransfer, `flow.inbound_transfer` contains the InboundTransfer's ID. Pass that ID to
// the `flow` parameter on List Transactions.
type V2MoneyManagementTransactionFlow struct {
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
	Type V2MoneyManagementTransactionFlowType `json:"type"`
}

// Hosted transaction receipt that is provided when money movement is considered regulated under Stripe's money transmission licenses. If not applicable, `regulatory_receipt.status` will be `not_applicable` and no URL will be provided unless you opt in by calling `/refresh_regulatory_receipt`.
type V2MoneyManagementTransactionRegulatoryReceipt struct {
	// Current availability of the regulatory receipt.
	Status V2MoneyManagementTransactionRegulatoryReceiptStatus `json:"status"`
	// Hosted URL for the receipt.
	URL string `json:"url,omitempty"`
	// Time until which `url` is valid.
	URLExpiresAt time.Time `json:"url_expires_at,omitempty"`
}

// Timestamps for when the Transaction transitioned to each status.
type V2MoneyManagementTransactionStatusTransitions struct {
	// The time at which the Transaction became `posted`. Only present when `status` is `posted`.
	PostedAt time.Time `json:"posted_at,omitempty"`
	// The time at which the Transaction became `void`. Only present when `status` is `void`.
	VoidAt time.Time `json:"void_at,omitempty"`
}

// A Transaction records how a money movement flow affects your FinancialAccount balance.
// A flow can create one or more Transactions.
type V2MoneyManagementTransaction struct {
	APIResource
	// The amount of the Transaction.
	Amount Amount `json:"amount"`
	// The change to the FinancialAccount's balance. This is the sum of the balance impacts of this Transaction's
	// TransactionEntries with `effective_at` timestamps in the past.
	BalanceImpact *V2MoneyManagementTransactionBalanceImpact `json:"balance_impact"`
	// Open Enum. A category used to classify the Transaction.
	Category V2MoneyManagementTransactionCategory `json:"category"`
	// Details about the counterparty to this Transaction.
	Counterparty *V2MoneyManagementTransactionCounterparty `json:"counterparty,omitempty"`
	// Time at which the object was created. Represented as an RFC 3339 timestamp in UTC with millisecond precision, for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// A description of the Transaction. When applicable, the description is copied from the flow object when
	// the Transaction is created. You can update it using the Update Transaction endpoint.
	Description string `json:"description,omitempty"`
	// The ID of the FinancialAccount affected by this Transaction.
	FinancialAccount string `json:"financial_account"`
	// Details about the flow object that created the Transaction, such as an InboundTransfer or OutboundPayment.
	// `flow.type` identifies the kind of object. The corresponding field contains its ID, when available:
	// for an InboundTransfer, `flow.inbound_transfer` contains the InboundTransfer's ID. Pass that ID to
	// the `flow` parameter on List Transactions.
	Flow *V2MoneyManagementTransactionFlow `json:"flow,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information
	// about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Hosted transaction receipt that is provided when money movement is considered regulated under Stripe's money transmission licenses. If not applicable, `regulatory_receipt.status` will be `not_applicable` and no URL will be provided unless you opt in by calling `/refresh_regulatory_receipt`.
	RegulatoryReceipt *V2MoneyManagementTransactionRegulatoryReceipt `json:"regulatory_receipt"`
	// Closed Enum. The current status of the Transaction.
	// A Transaction is `pending` if either `balance_impact.inbound_pending` or `balance_impact.outbound_pending` is non-zero.
	// A Transaction is `posted` if only `balance_impact.available` is non-zero.
	// A Transaction is `void` if there is no balance impact.
	// `posted` and `void` are terminal states, and no additional TransactionEntries will be added to the Transaction.
	Status V2MoneyManagementTransactionStatus `json:"status"`
	// Timestamps for when the Transaction transitioned to each status.
	StatusTransitions *V2MoneyManagementTransactionStatusTransitions `json:"status_transitions"`
}
