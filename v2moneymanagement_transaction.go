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
	V2MoneyManagementTransactionCategoryAdjustment                              V2MoneyManagementTransactionCategory = "adjustment"
	V2MoneyManagementTransactionCategoryAdvance                                 V2MoneyManagementTransactionCategory = "advance"
	V2MoneyManagementTransactionCategoryAnticipationRepayment                   V2MoneyManagementTransactionCategory = "anticipation_repayment"
	V2MoneyManagementTransactionCategoryBalanceTransfer                         V2MoneyManagementTransactionCategory = "balance_transfer"
	V2MoneyManagementTransactionCategoryClimateOrderPurchase                    V2MoneyManagementTransactionCategory = "climate_order_purchase"
	V2MoneyManagementTransactionCategoryClimateOrderRefund                      V2MoneyManagementTransactionCategory = "climate_order_refund"
	V2MoneyManagementTransactionCategoryConnectCollectionTransfer               V2MoneyManagementTransactionCategory = "connect_collection_transfer"
	V2MoneyManagementTransactionCategoryConnectReservedFunds                    V2MoneyManagementTransactionCategory = "connect_reserved_funds"
	V2MoneyManagementTransactionCategoryContribution                            V2MoneyManagementTransactionCategory = "contribution"
	V2MoneyManagementTransactionCategoryCurrencyConversion                      V2MoneyManagementTransactionCategory = "currency_conversion"
	V2MoneyManagementTransactionCategoryDispute                                 V2MoneyManagementTransactionCategory = "dispute"
	V2MoneyManagementTransactionCategoryDisputeReversal                         V2MoneyManagementTransactionCategory = "dispute_reversal"
	V2MoneyManagementTransactionCategoryFinancingPaydown                        V2MoneyManagementTransactionCategory = "financing_paydown"
	V2MoneyManagementTransactionCategoryFinancingPaydownReversal                V2MoneyManagementTransactionCategory = "financing_paydown_reversal"
	V2MoneyManagementTransactionCategoryInboundPayment                          V2MoneyManagementTransactionCategory = "inbound_payment"
	V2MoneyManagementTransactionCategoryInboundPaymentFailure                   V2MoneyManagementTransactionCategory = "inbound_payment_failure"
	V2MoneyManagementTransactionCategoryInboundTransfer                         V2MoneyManagementTransactionCategory = "inbound_transfer"
	V2MoneyManagementTransactionCategoryInboundTransferReversal                 V2MoneyManagementTransactionCategory = "inbound_transfer_reversal"
	V2MoneyManagementTransactionCategoryIndiaMdrProcessingFee                   V2MoneyManagementTransactionCategory = "india_mdr_processing_fee"
	V2MoneyManagementTransactionCategoryIssuingDispute                          V2MoneyManagementTransactionCategory = "issuing_dispute"
	V2MoneyManagementTransactionCategoryIssuingDisputeFraudLiabilityDebit       V2MoneyManagementTransactionCategory = "issuing_dispute_fraud_liability_debit"
	V2MoneyManagementTransactionCategoryIssuingDisputeProvisionalCredit         V2MoneyManagementTransactionCategory = "issuing_dispute_provisional_credit"
	V2MoneyManagementTransactionCategoryIssuingDisputeProvisionalCreditReversal V2MoneyManagementTransactionCategory = "issuing_dispute_provisional_credit_reversal"
	V2MoneyManagementTransactionCategoryMinimumBalanceHold                      V2MoneyManagementTransactionCategory = "minimum_balance_hold"
	V2MoneyManagementTransactionCategoryNetworkCost                             V2MoneyManagementTransactionCategory = "network_cost"
	V2MoneyManagementTransactionCategoryObligation                              V2MoneyManagementTransactionCategory = "obligation"
	V2MoneyManagementTransactionCategoryOutboundPayment                         V2MoneyManagementTransactionCategory = "outbound_payment"
	V2MoneyManagementTransactionCategoryOutboundPaymentReversal                 V2MoneyManagementTransactionCategory = "outbound_payment_reversal"
	V2MoneyManagementTransactionCategoryOutboundTransfer                        V2MoneyManagementTransactionCategory = "outbound_transfer"
	V2MoneyManagementTransactionCategoryOutboundTransferReversal                V2MoneyManagementTransactionCategory = "outbound_transfer_reversal"
	V2MoneyManagementTransactionCategoryPartialCaptureReversal                  V2MoneyManagementTransactionCategory = "partial_capture_reversal"
	V2MoneyManagementTransactionCategoryPaymentMethodPassthroughFee             V2MoneyManagementTransactionCategory = "payment_method_passthrough_fee"
	V2MoneyManagementTransactionCategoryPaymentNetworkReservedFunds             V2MoneyManagementTransactionCategory = "payment_network_reserved_funds"
	V2MoneyManagementTransactionCategoryPlatformEarning                         V2MoneyManagementTransactionCategory = "platform_earning"
	V2MoneyManagementTransactionCategoryPlatformEarningRefund                   V2MoneyManagementTransactionCategory = "platform_earning_refund"
	V2MoneyManagementTransactionCategoryPlatformFee                             V2MoneyManagementTransactionCategory = "platform_fee"
	V2MoneyManagementTransactionCategoryReceivedCredit                          V2MoneyManagementTransactionCategory = "received_credit"
	V2MoneyManagementTransactionCategoryReceivedCreditReversal                  V2MoneyManagementTransactionCategory = "received_credit_reversal"
	V2MoneyManagementTransactionCategoryReceivedDebit                           V2MoneyManagementTransactionCategory = "received_debit"
	V2MoneyManagementTransactionCategoryReceivedDebitReversal                   V2MoneyManagementTransactionCategory = "received_debit_reversal"
	V2MoneyManagementTransactionCategoryRefund                                  V2MoneyManagementTransactionCategory = "refund"
	V2MoneyManagementTransactionCategoryRefundFailure                           V2MoneyManagementTransactionCategory = "refund_failure"
	V2MoneyManagementTransactionCategoryRiskReservedFunds                       V2MoneyManagementTransactionCategory = "risk_reserved_funds"
	V2MoneyManagementTransactionCategoryStripeBalancePaymentDebit               V2MoneyManagementTransactionCategory = "stripe_balance_payment_debit"
	V2MoneyManagementTransactionCategoryStripeBalancePaymentDebitReversal       V2MoneyManagementTransactionCategory = "stripe_balance_payment_debit_reversal"
	V2MoneyManagementTransactionCategoryStripeFee                               V2MoneyManagementTransactionCategory = "stripe_fee"
	V2MoneyManagementTransactionCategoryStripeFeeTax                            V2MoneyManagementTransactionCategory = "stripe_fee_tax"
	V2MoneyManagementTransactionCategoryTaxWithholding                          V2MoneyManagementTransactionCategory = "tax_withholding"
	V2MoneyManagementTransactionCategoryTransferReversal                        V2MoneyManagementTransactionCategory = "transfer_reversal"
	V2MoneyManagementTransactionCategoryUnreconciledCustomerFunds               V2MoneyManagementTransactionCategory = "unreconciled_customer_funds"
)

// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
type V2MoneyManagementTransactionFlowType string

// List of values that V2MoneyManagementTransactionFlowType can take
const (
	V2MoneyManagementTransactionFlowTypeAdjustment                   V2MoneyManagementTransactionFlowType = "adjustment"
	V2MoneyManagementTransactionFlowTypeApplicationFee               V2MoneyManagementTransactionFlowType = "application_fee"
	V2MoneyManagementTransactionFlowTypeApplicationFeeRefund         V2MoneyManagementTransactionFlowType = "application_fee_refund"
	V2MoneyManagementTransactionFlowTypeCharge                       V2MoneyManagementTransactionFlowType = "charge"
	V2MoneyManagementTransactionFlowTypeCurrencyConversion           V2MoneyManagementTransactionFlowType = "currency_conversion"
	V2MoneyManagementTransactionFlowTypeDispute                      V2MoneyManagementTransactionFlowType = "dispute"
	V2MoneyManagementTransactionFlowTypeFeeTransaction               V2MoneyManagementTransactionFlowType = "fee_transaction"
	V2MoneyManagementTransactionFlowTypeInboundTransfer              V2MoneyManagementTransactionFlowType = "inbound_transfer"
	V2MoneyManagementTransactionFlowTypeOutboundPayment              V2MoneyManagementTransactionFlowType = "outbound_payment"
	V2MoneyManagementTransactionFlowTypeOutboundTransfer             V2MoneyManagementTransactionFlowType = "outbound_transfer"
	V2MoneyManagementTransactionFlowTypePayout                       V2MoneyManagementTransactionFlowType = "payout"
	V2MoneyManagementTransactionFlowTypeReceivedCredit               V2MoneyManagementTransactionFlowType = "received_credit"
	V2MoneyManagementTransactionFlowTypeReceivedDebit                V2MoneyManagementTransactionFlowType = "received_debit"
	V2MoneyManagementTransactionFlowTypeRefund                       V2MoneyManagementTransactionFlowType = "refund"
	V2MoneyManagementTransactionFlowTypeReserveHold                  V2MoneyManagementTransactionFlowType = "reserve_hold"
	V2MoneyManagementTransactionFlowTypeReserveRelease               V2MoneyManagementTransactionFlowType = "reserve_release"
	V2MoneyManagementTransactionFlowTypeTopup                        V2MoneyManagementTransactionFlowType = "topup"
	V2MoneyManagementTransactionFlowTypeTransfer                     V2MoneyManagementTransactionFlowType = "transfer"
	V2MoneyManagementTransactionFlowTypeTransferReversal             V2MoneyManagementTransactionFlowType = "transfer_reversal"
	V2MoneyManagementTransactionFlowTypeTreasuryCreditReversal       V2MoneyManagementTransactionFlowType = "treasury_credit_reversal"
	V2MoneyManagementTransactionFlowTypeTreasuryDebitReversal        V2MoneyManagementTransactionFlowType = "treasury_debit_reversal"
	V2MoneyManagementTransactionFlowTypeTreasuryInboundTransfer      V2MoneyManagementTransactionFlowType = "treasury_inbound_transfer"
	V2MoneyManagementTransactionFlowTypeTreasuryIssuingAuthorization V2MoneyManagementTransactionFlowType = "treasury_issuing_authorization"
	V2MoneyManagementTransactionFlowTypeTreasuryOther                V2MoneyManagementTransactionFlowType = "treasury_other"
	V2MoneyManagementTransactionFlowTypeTreasuryOutboundPayment      V2MoneyManagementTransactionFlowType = "treasury_outbound_payment"
	V2MoneyManagementTransactionFlowTypeTreasuryOutboundTransfer     V2MoneyManagementTransactionFlowType = "treasury_outbound_transfer"
	V2MoneyManagementTransactionFlowTypeTreasuryReceivedCredit       V2MoneyManagementTransactionFlowType = "treasury_received_credit"
	V2MoneyManagementTransactionFlowTypeTreasuryReceivedDebit        V2MoneyManagementTransactionFlowType = "treasury_received_debit"
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

// Counterparty to this Transaction.
type V2MoneyManagementTransactionCounterparty struct {
	// Name of the counterparty.
	Name string `json:"name,omitempty"`
}

// Details about the Flow object that created the Transaction.
type V2MoneyManagementTransactionFlow struct {
	// If applicable, the ID of the Adjustment that created this Transaction.
	Adjustment string `json:"adjustment,omitempty"`
	// If applicable, the ID of the Application Fee that created this Transaction.
	ApplicationFee string `json:"application_fee,omitempty"`
	// If applicable, the ID of the Application Fee Refund that created this Transaction.
	ApplicationFeeRefund string `json:"application_fee_refund,omitempty"`
	// If applicable, the ID of the Charge that created this Transaction.
	Charge string `json:"charge,omitempty"`
	// In the future, this will be the ID of the currency conversion that created this Transaction. For now, this field is always null.
	CurrencyConversion string `json:"currency_conversion,omitempty"`
	// If applicable, the ID of the Dispute that created this Transaction.
	Dispute string `json:"dispute,omitempty"`
	// If applicable, the ID of the FeeTransaction that created this Transaction.
	FeeTransaction string `json:"fee_transaction,omitempty"`
	// If applicable, the ID of the InboundTransfer that created this Transaction.
	InboundTransfer string `json:"inbound_transfer,omitempty"`
	// If applicable, the ID of the OutboundPayment that created this Transaction.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// If applicable, the ID of the OutboundTransfer that created this Transaction.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// If applicable, the ID of the Payout that created this Transaction.
	Payout string `json:"payout,omitempty"`
	// If applicable, the ID of the ReceivedCredit that created this Transaction.
	ReceivedCredit string `json:"received_credit,omitempty"`
	// If applicable, the ID of the ReceivedDebit that created this Transaction.
	ReceivedDebit string `json:"received_debit,omitempty"`
	// If applicable, the ID of the Refund that created this Transaction.
	Refund string `json:"refund,omitempty"`
	// If applicable, the ID of the Reserve Hold that created this Transaction.
	ReserveHold string `json:"reserve_hold,omitempty"`
	// If applicable, the ID of the Reserve Release that created this Transaction.
	ReserveRelease string `json:"reserve_release,omitempty"`
	// If applicable, the ID of the Topup that created this Transaction.
	Topup string `json:"topup,omitempty"`
	// If applicable, the ID of the Transfer that created this Transaction.
	Transfer string `json:"transfer,omitempty"`
	// If applicable, the ID of the Transfer Reversal that created this Transaction.
	TransferReversal string `json:"transfer_reversal,omitempty"`
	// If applicable, the ID of the Treasury CreditReversal that created this Transaction.
	TreasuryCreditReversal string `json:"treasury_credit_reversal,omitempty"`
	// If applicable, the ID of the Treasury DebitReversal that created this Transaction.
	TreasuryDebitReversal string `json:"treasury_debit_reversal,omitempty"`
	// If applicable, the ID of the Treasury InboundTransfer that created this Transaction.
	TreasuryInboundTransfer string `json:"treasury_inbound_transfer,omitempty"`
	// If applicable, the ID of the Treasury IssuingAuthorization that created this Transaction.
	TreasuryIssuingAuthorization string `json:"treasury_issuing_authorization,omitempty"`
	// If applicable, the ID of the Treasury OutboundPayment that created this Transaction.
	TreasuryOutboundPayment string `json:"treasury_outbound_payment,omitempty"`
	// If applicable, the ID of the Treasury OutboundTransfer that created this Transaction.
	TreasuryOutboundTransfer string `json:"treasury_outbound_transfer,omitempty"`
	// If applicable, the ID of the Treasury ReceivedCredit that created this Transaction.
	TreasuryReceivedCredit string `json:"treasury_received_credit,omitempty"`
	// If applicable, the ID of the Treasury ReceivedDebit that created this Transaction.
	TreasuryReceivedDebit string `json:"treasury_received_debit,omitempty"`
	// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
	Type V2MoneyManagementTransactionFlowType `json:"type"`
}

// Timestamps for when the Transaction transitioned to a particular status.
type V2MoneyManagementTransactionStatusTransitions struct {
	// The time at which the Transaction became posted. Only present if status == posted.
	PostedAt time.Time `json:"posted_at,omitempty"`
	// The time at which the Transaction became void. Only present if status == void.
	VoidAt time.Time `json:"void_at,omitempty"`
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
	// Counterparty to this Transaction.
	Counterparty *V2MoneyManagementTransactionCounterparty `json:"counterparty,omitempty"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Description of this Transaction. When applicable, the description is copied from the Flow object at the time
	// of transaction creation.
	Description string `json:"description,omitempty"`
	// Indicates the FinancialAccount affected by this Transaction.
	FinancialAccount string `json:"financial_account"`
	// Details about the Flow object that created the Transaction.
	Flow *V2MoneyManagementTransactionFlow `json:"flow,omitempty"`
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
	// The v1 Treasury transaction associated with this transaction.
	TreasuryTransaction string `json:"treasury_transaction,omitempty"`
}
