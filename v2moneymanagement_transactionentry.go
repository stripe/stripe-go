//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum for now, and will be turned into an Open Enum soon. A descriptive category used to classify the Transaction.
type V2MoneyManagementTransactionEntryTransactionDetailsCategory string

// List of values that V2MoneyManagementTransactionEntryTransactionDetailsCategory can take
const (
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryAdjustment                              V2MoneyManagementTransactionEntryTransactionDetailsCategory = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryAdvance                                 V2MoneyManagementTransactionEntryTransactionDetailsCategory = "advance"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryAnticipationRepayment                   V2MoneyManagementTransactionEntryTransactionDetailsCategory = "anticipation_repayment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryBalanceTransfer                         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "balance_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryCharge                                  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "charge"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryChargeFailure                           V2MoneyManagementTransactionEntryTransactionDetailsCategory = "charge_failure"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryClimateOrderPurchase                    V2MoneyManagementTransactionEntryTransactionDetailsCategory = "climate_order_purchase"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryClimateOrderRefund                      V2MoneyManagementTransactionEntryTransactionDetailsCategory = "climate_order_refund"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryConnectCollectionTransfer               V2MoneyManagementTransactionEntryTransactionDetailsCategory = "connect_collection_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryConnectReservedFunds                    V2MoneyManagementTransactionEntryTransactionDetailsCategory = "connect_reserved_funds"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryContribution                            V2MoneyManagementTransactionEntryTransactionDetailsCategory = "contribution"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryCurrencyConversion                      V2MoneyManagementTransactionEntryTransactionDetailsCategory = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryDisputeReversal                         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "dispute_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryFinancingPaydown                        V2MoneyManagementTransactionEntryTransactionDetailsCategory = "financing_paydown"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryFinancingPaydownReversal                V2MoneyManagementTransactionEntryTransactionDetailsCategory = "financing_paydown_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryInboundTransfer                         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryInboundTransferReversal                 V2MoneyManagementTransactionEntryTransactionDetailsCategory = "inbound_transfer_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryIssuingDispute                          V2MoneyManagementTransactionEntryTransactionDetailsCategory = "issuing_dispute"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryIssuingDisputeFraudLiabilityDebit       V2MoneyManagementTransactionEntryTransactionDetailsCategory = "issuing_dispute_fraud_liability_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryIssuingDisputeProvisionalCredit         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "issuing_dispute_provisional_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryIssuingDisputeProvisionalCreditReversal V2MoneyManagementTransactionEntryTransactionDetailsCategory = "issuing_dispute_provisional_credit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryMinimumBalanceHold                      V2MoneyManagementTransactionEntryTransactionDetailsCategory = "minimum_balance_hold"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryNetworkCost                             V2MoneyManagementTransactionEntryTransactionDetailsCategory = "network_cost"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryObligation                              V2MoneyManagementTransactionEntryTransactionDetailsCategory = "obligation"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundPayment                         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundPaymentReversal                 V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_payment_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundTransfer                        V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundTransferReversal                V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_transfer_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryPartialCaptureReversal                  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "partial_capture_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryPaymentNetworkReservedFunds             V2MoneyManagementTransactionEntryTransactionDetailsCategory = "payment_network_reserved_funds"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryPlatformEarning                         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "platform_earning"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryPlatformEarningRefund                   V2MoneyManagementTransactionEntryTransactionDetailsCategory = "platform_earning_refund"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryPlatformFee                             V2MoneyManagementTransactionEntryTransactionDetailsCategory = "platform_fee"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedCredit                          V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedCreditReversal                  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_credit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedDebit                           V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedDebitReversal                   V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_debit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryRefundFailure                           V2MoneyManagementTransactionEntryTransactionDetailsCategory = "refund_failure"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReturn                                  V2MoneyManagementTransactionEntryTransactionDetailsCategory = "return"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryRiskReservedFunds                       V2MoneyManagementTransactionEntryTransactionDetailsCategory = "risk_reserved_funds"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeBalancePaymentDebit               V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_balance_payment_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeBalancePaymentDebitReversal       V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_balance_payment_debit_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeFee                               V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_fee"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeFeeTax                            V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_fee_tax"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryTransferReversal                        V2MoneyManagementTransactionEntryTransactionDetailsCategory = "transfer_reversal"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryUnreconciledCustomerFunds               V2MoneyManagementTransactionEntryTransactionDetailsCategory = "unreconciled_customer_funds"
)

// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
type V2MoneyManagementTransactionEntryTransactionDetailsFlowType string

// List of values that V2MoneyManagementTransactionEntryTransactionDetailsFlowType can take
const (
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeAdjustment           V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeApplicationFee       V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "application_fee"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeApplicationFeeRefund V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "application_fee_refund"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeCharge               V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "charge"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeCurrencyConversion   V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeDispute              V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "dispute"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeFeeTransaction       V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "fee_transaction"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeInboundTransfer      V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundPayment      V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundTransfer     V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypePayout               V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "payout"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedCredit       V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedDebit        V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeRefund               V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "refund"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReserveHold          V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "reserve_hold"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReserveRelease       V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "reserve_release"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeTopup                V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "topup"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeTransfer             V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeTransferReversal     V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "transfer_reversal"
)

// Impact to the available balance.
type V2MoneyManagementTransactionEntryBalanceImpactAvailable struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Impact to the inbound_pending balance.
type V2MoneyManagementTransactionEntryBalanceImpactInboundPending struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Impact to the outbound_pending balance.
type V2MoneyManagementTransactionEntryBalanceImpactOutboundPending struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The delta to the FinancialAccount's balance.
type V2MoneyManagementTransactionEntryBalanceImpact struct {
	// Impact to the available balance.
	Available *V2MoneyManagementTransactionEntryBalanceImpactAvailable `json:"available"`
	// Impact to the inbound_pending balance.
	InboundPending *V2MoneyManagementTransactionEntryBalanceImpactInboundPending `json:"inbound_pending"`
	// Impact to the outbound_pending balance.
	OutboundPending *V2MoneyManagementTransactionEntryBalanceImpactOutboundPending `json:"outbound_pending"`
}

// Details about the Flow object that created the Transaction.
type V2MoneyManagementTransactionEntryTransactionDetailsFlow struct {
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
	// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
	Type V2MoneyManagementTransactionEntryTransactionDetailsFlowType `json:"type"`
}

// Details copied from the transaction that this TransactionEntry belongs to.
type V2MoneyManagementTransactionEntryTransactionDetails struct {
	// Closed Enum for now, and will be turned into an Open Enum soon. A descriptive category used to classify the Transaction.
	Category V2MoneyManagementTransactionEntryTransactionDetailsCategory `json:"category"`
	// Indicates the FinancialAccount affected by this Transaction.
	FinancialAccount string `json:"financial_account"`
	// Details about the Flow object that created the Transaction.
	Flow *V2MoneyManagementTransactionEntryTransactionDetailsFlow `json:"flow"`
}

// TransactionEntries represent individual money movements across different states within a Transaction.
type V2MoneyManagementTransactionEntry struct {
	APIResource
	// The delta to the FinancialAccount's balance.
	BalanceImpact *V2MoneyManagementTransactionEntryBalanceImpact `json:"balance_impact"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Time at which the entry impacted (or will impact if it's in the future) the FinancialAccount balance.
	EffectiveAt time.Time `json:"effective_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The Transaction that this TransactionEntry belongs to.
	Transaction string `json:"transaction"`
	// Details copied from the transaction that this TransactionEntry belongs to.
	TransactionDetails *V2MoneyManagementTransactionEntryTransactionDetails `json:"transaction_details"`
}
