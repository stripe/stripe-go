//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. If applicable, the type of flow linked to this Adjustment. The field matching this value will contain the ID of the flow.
type V2MoneyManagementAdjustmentAdjustedFlowType string

// List of values that V2MoneyManagementAdjustmentAdjustedFlowType can take
const (
	V2MoneyManagementAdjustmentAdjustedFlowTypeAdjustment       V2MoneyManagementAdjustmentAdjustedFlowType = "adjustment"
	V2MoneyManagementAdjustmentAdjustedFlowTypeInboundTransfer  V2MoneyManagementAdjustmentAdjustedFlowType = "inbound_transfer"
	V2MoneyManagementAdjustmentAdjustedFlowTypeOutboundPayment  V2MoneyManagementAdjustmentAdjustedFlowType = "outbound_payment"
	V2MoneyManagementAdjustmentAdjustedFlowTypeOutboundTransfer V2MoneyManagementAdjustmentAdjustedFlowType = "outbound_transfer"
	V2MoneyManagementAdjustmentAdjustedFlowTypeReceivedCredit   V2MoneyManagementAdjustmentAdjustedFlowType = "received_credit"
	V2MoneyManagementAdjustmentAdjustedFlowTypeReceivedDebit    V2MoneyManagementAdjustmentAdjustedFlowType = "received_debit"
)

// Details about the original flow object linked to this Adjustment, if applicable.
// `adjusted_flow.type` identifies the kind of object. The corresponding field contains its ID:
// for an InboundTransfer, `adjusted_flow.inbound_transfer` contains the InboundTransfer's ID.
// Pass that ID to the `adjusted_flow` parameter on List Adjustments.
type V2MoneyManagementAdjustmentAdjustedFlow struct {
	// If applicable, the ID of the original Adjustment linked to this Adjustment.
	Adjustment string `json:"adjustment,omitempty"`
	// If applicable, the ID of the InboundTransfer linked to this Adjustment.
	InboundTransfer string `json:"inbound_transfer,omitempty"`
	// If applicable, the ID of the OutboundPayment linked to this Adjustment.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// If applicable, the ID of the OutboundTransfer linked to this Adjustment.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// If applicable, the ID of the ReceivedCredit linked to this Adjustment.
	ReceivedCredit string `json:"received_credit,omitempty"`
	// If applicable, the ID of the ReceivedDebit linked to this Adjustment.
	ReceivedDebit string `json:"received_debit,omitempty"`
	// Open Enum. If applicable, the type of flow linked to this Adjustment. The field matching this value will contain the ID of the flow.
	Type V2MoneyManagementAdjustmentAdjustedFlowType `json:"type"`
}

// An Adjustment records a credit or debit that Stripe makes to your Financial Account balance,
// for example to correct technical or operational errors.
type V2MoneyManagementAdjustment struct {
	APIResource
	// Details about the original flow object linked to this Adjustment, if applicable.
	// `adjusted_flow.type` identifies the kind of object. The corresponding field contains its ID:
	// for an InboundTransfer, `adjusted_flow.inbound_transfer` contains the InboundTransfer's ID.
	// Pass that ID to the `adjusted_flow` parameter on List Adjustments.
	AdjustedFlow *V2MoneyManagementAdjustmentAdjustedFlow `json:"adjusted_flow,omitempty"`
	// The amount of the Adjustment.
	Amount Amount `json:"amount"`
	// Time at which the object was created. Represented as an RFC 3339 timestamp in UTC with millisecond precision,
	// for example `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// A description of the Adjustment and its purpose.
	Description string `json:"description,omitempty"`
	// The ID of the FinancialAccount whose balance this Adjustment affects.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A link to the Stripe-hosted receipt provided when money movement is regulated under Stripe's money transmission licenses.
	// You can use this link to view the receipt for 60 days after the Adjustment is created.
	// After that, the link expires and `receipt_url` is `null`.
	ReceiptURL string `json:"receipt_url,omitempty"`
}
