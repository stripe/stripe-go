//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of this object.
type V2MoneyManagementInboundFundingHistoryType string

// List of values that V2MoneyManagementInboundFundingHistoryType can take
const (
	V2MoneyManagementInboundFundingHistoryTypeCheckFailed     V2MoneyManagementInboundFundingHistoryType = "check_failed"
	V2MoneyManagementInboundFundingHistoryTypeCheckProcessing V2MoneyManagementInboundFundingHistoryType = "check_processing"
	V2MoneyManagementInboundFundingHistoryTypeCheckQueued     V2MoneyManagementInboundFundingHistoryType = "check_queued"
	V2MoneyManagementInboundFundingHistoryTypeCheckReturned   V2MoneyManagementInboundFundingHistoryType = "check_returned"
	V2MoneyManagementInboundFundingHistoryTypeCheckSucceeded  V2MoneyManagementInboundFundingHistoryType = "check_succeeded"
)

// The current status of the InboundFunding.
type V2MoneyManagementInboundFundingStatus string

// List of values that V2MoneyManagementInboundFundingStatus can take
const (
	V2MoneyManagementInboundFundingStatusFailed     V2MoneyManagementInboundFundingStatus = "failed"
	V2MoneyManagementInboundFundingStatusProcessing V2MoneyManagementInboundFundingStatus = "processing"
	V2MoneyManagementInboundFundingStatusQueued     V2MoneyManagementInboundFundingStatus = "queued"
	V2MoneyManagementInboundFundingStatusReturned   V2MoneyManagementInboundFundingStatus = "returned"
	V2MoneyManagementInboundFundingStatusSucceeded  V2MoneyManagementInboundFundingStatus = "succeeded"
)

// The payment method that funds the InboundFunding.
type V2MoneyManagementInboundFundingFromPaymentMethod struct {
	// The type of this object.
	Type string `json:"type"`
	// The ID of the paper check.
	USPaperCheck string `json:"us_paper_check,omitempty"`
}

// The source of the funds.
type V2MoneyManagementInboundFundingFrom struct {
	// The amount debited from the source.
	Debited Amount `json:"debited"`
	// The payment method that funds the InboundFunding.
	PaymentMethod *V2MoneyManagementInboundFundingFromPaymentMethod `json:"payment_method"`
}

// Details of the transition to failed.
type V2MoneyManagementInboundFundingHistoryCheckFailed struct {
	// The raw banking-partner failure code, when available.
	FailureCode string `json:"failure_code,omitempty"`
	// The reason the check failed.
	FailureReason string `json:"failure_reason,omitempty"`
}

// Details of the transition to processing.
type V2MoneyManagementInboundFundingHistoryCheckProcessing struct{}

// Details of the transition to queued.
type V2MoneyManagementInboundFundingHistoryCheckQueued struct{}

// Details of the transition to returned.
type V2MoneyManagementInboundFundingHistoryCheckReturned struct {
	// The raw banking-partner return code, when available.
	ReturnCode string `json:"return_code,omitempty"`
	// The reason the check was returned.
	ReturnReason string `json:"return_reason,omitempty"`
	// The ID of the reversal Transaction.
	ReversalTransaction string `json:"reversal_transaction,omitempty"`
}

// Details of the transition to succeeded.
type V2MoneyManagementInboundFundingHistoryCheckSucceeded struct{}

// The check lifecycle history.
type V2MoneyManagementInboundFundingHistory struct {
	// Details of the transition to failed.
	CheckFailed *V2MoneyManagementInboundFundingHistoryCheckFailed `json:"check_failed,omitempty"`
	// Details of the transition to processing.
	CheckProcessing *V2MoneyManagementInboundFundingHistoryCheckProcessing `json:"check_processing,omitempty"`
	// Details of the transition to queued.
	CheckQueued *V2MoneyManagementInboundFundingHistoryCheckQueued `json:"check_queued,omitempty"`
	// Details of the transition to returned.
	CheckReturned *V2MoneyManagementInboundFundingHistoryCheckReturned `json:"check_returned,omitempty"`
	// Details of the transition to succeeded.
	CheckSucceeded *V2MoneyManagementInboundFundingHistoryCheckSucceeded `json:"check_succeeded,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Time at which the history entry took effect.
	EffectiveAt time.Time `json:"effective_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The type of this object.
	Type V2MoneyManagementInboundFundingHistoryType `json:"type"`
}

// The destination for the funds.
type V2MoneyManagementInboundFundingTo struct {
	// The amount credited to the destination.
	Credited Amount `json:"credited"`
	// The ID of the destination FinancialAccount.
	FinancialAccount string `json:"financial_account"`
}

// An InboundFunding represents an other-to-me pull movement into a FinancialAccount.
type V2MoneyManagementInboundFunding struct {
	APIResource
	// The amount of the InboundFunding.
	Amount Amount `json:"amount"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// An arbitrary string attached to the InboundFunding.
	Description string `json:"description,omitempty"`
	// The source of the funds.
	From *V2MoneyManagementInboundFundingFrom `json:"from"`
	// The check lifecycle history.
	History []*V2MoneyManagementInboundFundingHistory `json:"history"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether the InboundFunding exists in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current status of the InboundFunding.
	Status V2MoneyManagementInboundFundingStatus `json:"status"`
	// The destination for the funds.
	To *V2MoneyManagementInboundFundingTo `json:"to"`
}
