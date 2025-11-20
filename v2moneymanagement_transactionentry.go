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
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryAdjustment         V2MoneyManagementTransactionEntryTransactionDetailsCategory = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryCurrencyConversion V2MoneyManagementTransactionEntryTransactionDetailsCategory = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryInboundTransfer    V2MoneyManagementTransactionEntryTransactionDetailsCategory = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundPayment    V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryOutboundTransfer   V2MoneyManagementTransactionEntryTransactionDetailsCategory = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedCredit     V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReceivedDebit      V2MoneyManagementTransactionEntryTransactionDetailsCategory = "received_debit"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryReturn             V2MoneyManagementTransactionEntryTransactionDetailsCategory = "return"
	V2MoneyManagementTransactionEntryTransactionDetailsCategoryStripeFee          V2MoneyManagementTransactionEntryTransactionDetailsCategory = "stripe_fee"
)

// Open Enum. Type of the flow that created the Transaction. The field matching this value will contain the ID of the flow.
type V2MoneyManagementTransactionEntryTransactionDetailsFlowType string

// List of values that V2MoneyManagementTransactionEntryTransactionDetailsFlowType can take
const (
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeAdjustment         V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "adjustment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeCurrencyConversion V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "currency_conversion"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeFeeTransaction     V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "fee_transaction"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeInboundTransfer    V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "inbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundPayment    V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_payment"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeOutboundTransfer   V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "outbound_transfer"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedCredit     V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_credit"
	V2MoneyManagementTransactionEntryTransactionDetailsFlowTypeReceivedDebit      V2MoneyManagementTransactionEntryTransactionDetailsFlowType = "received_debit"
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
	// In the future, this will be the ID of the currency conversion that created this Transaction. For now, this field is always null.
	CurrencyConversion string `json:"currency_conversion,omitempty"`
	// If applicable, the ID of the FeeTransaction that created this Transaction.
	FeeTransaction string `json:"fee_transaction,omitempty"`
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
