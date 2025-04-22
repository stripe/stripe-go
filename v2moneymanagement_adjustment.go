//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. If applicable, the type of flow linked to this Adjustment. The field matching this value will contain the ID of the flow.
type V2MoneyManagementAdjustmentAdjustedFlowType string

// List of values that V2MoneyManagementAdjustmentAdjustedFlowType can take
const (
	V2MoneyManagementAdjustmentAdjustedFlowTypeAdjustment       V2MoneyManagementAdjustmentAdjustedFlowType = "adjustment"
	V2MoneyManagementAdjustmentAdjustedFlowTypeBalanceExchange  V2MoneyManagementAdjustmentAdjustedFlowType = "balance_exchange"
	V2MoneyManagementAdjustmentAdjustedFlowTypeInboundPayment   V2MoneyManagementAdjustmentAdjustedFlowType = "inbound_payment"
	V2MoneyManagementAdjustmentAdjustedFlowTypeInboundTransfer  V2MoneyManagementAdjustmentAdjustedFlowType = "inbound_transfer"
	V2MoneyManagementAdjustmentAdjustedFlowTypeOutboundPayment  V2MoneyManagementAdjustmentAdjustedFlowType = "outbound_payment"
	V2MoneyManagementAdjustmentAdjustedFlowTypeOutboundTransfer V2MoneyManagementAdjustmentAdjustedFlowType = "outbound_transfer"
	V2MoneyManagementAdjustmentAdjustedFlowTypeReceivedCredit   V2MoneyManagementAdjustmentAdjustedFlowType = "received_credit"
	V2MoneyManagementAdjustmentAdjustedFlowTypeReceivedDebit    V2MoneyManagementAdjustmentAdjustedFlowType = "received_debit"
)

// If applicable, contains information about the original flow linked to this Adjustment.
type V2MoneyManagementAdjustmentAdjustedFlow struct {
	// If applicable, the ID of the Adjustment linked to this Adjustment.
	Adjustment string `json:"adjustment"`
	// If applicable, the ID of the InboundTransfer linked to this Adjustment.
	InboundTransfer string `json:"inbound_transfer"`
	// If applicable, the ID of the OutboundPayment linked to this Adjustment.
	OutboundPayment string `json:"outbound_payment"`
	// If applicable, the ID of the OutboundTransfer linked to this Adjustment.
	OutboundTransfer string `json:"outbound_transfer"`
	// If applicable, the ID of the ReceivedCredit linked to this Adjustment.
	ReceivedCredit string `json:"received_credit"`
	// If applicable, the ID of the ReceivedDebit linked to this Adjustment.
	ReceivedDebit string `json:"received_debit"`
	// Closed Enum. If applicable, the type of flow linked to this Adjustment. The field matching this value will contain the ID of the flow.
	Type V2MoneyManagementAdjustmentAdjustedFlowType `json:"type"`
}

// Adjustments represent Stripe-initiated credits or debits to a user balance. They might be used to amend balances due to technical or operational error.
type V2MoneyManagementAdjustment struct {
	APIResource
	// If applicable, contains information about the original flow linked to this Adjustment.
	AdjustedFlow *V2MoneyManagementAdjustmentAdjustedFlow `json:"adjusted_flow"`
	// The amount of the Adjustment.
	Amount Amount `json:"amount"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Description of the Adjustment and what it was used for.
	Description string `json:"description"`
	// The FinancialAccount that this adjustment is for.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A link to the Stripe-hosted receipt that is provided when money movement is considered regulated under Stripe's money transmission licenses. The receipt link remains active for 60 days from the Adjustment creation date. After this period, the link will expire and the receipt url value will be null.
	ReceiptURL string `json:"receipt_url"`
}
