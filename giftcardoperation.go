//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The failure code of the operation. Only present if the status is failed.
type GiftCardOperationFailureCode string

// List of values that GiftCardOperationFailureCode can take
const (
	GiftCardOperationFailureCodeActionNotSupported   GiftCardOperationFailureCode = "action_not_supported"
	GiftCardOperationFailureCodeCardAlreadyActivated GiftCardOperationFailureCode = "card_already_activated"
	GiftCardOperationFailureCodeCardExpired          GiftCardOperationFailureCode = "card_expired"
	GiftCardOperationFailureCodeCardNotActivated     GiftCardOperationFailureCode = "card_not_activated"
	GiftCardOperationFailureCodeDoNotHonor           GiftCardOperationFailureCode = "do_not_honor"
	GiftCardOperationFailureCodeGenericFailure       GiftCardOperationFailureCode = "generic_failure"
	GiftCardOperationFailureCodeInsufficientBalance  GiftCardOperationFailureCode = "insufficient_balance"
	GiftCardOperationFailureCodeInvalidAmount        GiftCardOperationFailureCode = "invalid_amount"
	GiftCardOperationFailureCodeInvalidCurrency      GiftCardOperationFailureCode = "invalid_currency"
	GiftCardOperationFailureCodeInvalidNumber        GiftCardOperationFailureCode = "invalid_number"
	GiftCardOperationFailureCodeInvalidPIN           GiftCardOperationFailureCode = "invalid_pin"
	GiftCardOperationFailureCodeInvalidTrackData     GiftCardOperationFailureCode = "invalid_track_data"
	GiftCardOperationFailureCodeLostCard             GiftCardOperationFailureCode = "lost_card"
	GiftCardOperationFailureCodeLostOrStolenCard     GiftCardOperationFailureCode = "lost_or_stolen_card"
	GiftCardOperationFailureCodePINRequired          GiftCardOperationFailureCode = "pin_required"
	GiftCardOperationFailureCodePINTriesExceeded     GiftCardOperationFailureCode = "pin_tries_exceeded"
	GiftCardOperationFailureCodeProcessingError      GiftCardOperationFailureCode = "processing_error"
	GiftCardOperationFailureCodeProviderUnavailable  GiftCardOperationFailureCode = "provider_unavailable"
	GiftCardOperationFailureCodeStolenCard           GiftCardOperationFailureCode = "stolen_card"
	GiftCardOperationFailureCodeSuspectedFraud       GiftCardOperationFailureCode = "suspected_fraud"
	GiftCardOperationFailureCodeTimeout              GiftCardOperationFailureCode = "timeout"
)

// The status of the operation.
type GiftCardOperationStatus string

// List of values that GiftCardOperationStatus can take
const (
	GiftCardOperationStatusFailed    GiftCardOperationStatus = "failed"
	GiftCardOperationStatusSucceeded GiftCardOperationStatus = "succeeded"
)

// The type of operation performed.
type GiftCardOperationType string

// List of values that GiftCardOperationType can take
const (
	GiftCardOperationTypeActivation     GiftCardOperationType = "activation"
	GiftCardOperationTypeActivationVoid GiftCardOperationType = "activation_void"
	GiftCardOperationTypeBalanceCheck   GiftCardOperationType = "balance_check"
	GiftCardOperationTypeCashout        GiftCardOperationType = "cashout"
	GiftCardOperationTypeCashoutVoid    GiftCardOperationType = "cashout_void"
	GiftCardOperationTypeDeactivation   GiftCardOperationType = "deactivation"
	GiftCardOperationTypeReload         GiftCardOperationType = "reload"
	GiftCardOperationTypeReloadVoid     GiftCardOperationType = "reload_void"
)

// Retrieves a third-party gift card operation object.
type GiftCardOperationParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardOperationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a third-party gift card operation object.
type GiftCardOperationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardOperationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationActivationBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card activation operation.
type GiftCardOperationActivation struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationActivationBalance `json:"balance"`
}

// Details about a gift card activation void operation.
type GiftCardOperationActivationVoid struct {
	// The operation that was voided.
	VoidedOperation string `json:"voided_operation"`
}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationBalanceCheckBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card balance check operation.
type GiftCardOperationBalanceCheck struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationBalanceCheckBalance `json:"balance"`
}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationCashoutBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// The balance before the operation.
type GiftCardOperationCashoutPreviousBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card cashout operation.
type GiftCardOperationCashout struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationCashoutBalance `json:"balance"`
	// The balance before the operation.
	PreviousBalance *GiftCardOperationCashoutPreviousBalance `json:"previous_balance"`
}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationCashoutVoidBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card cashout void operation.
type GiftCardOperationCashoutVoid struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationCashoutVoidBalance `json:"balance"`
	// The operation that was voided.
	VoidedOperation string `json:"voided_operation"`
}

// Details about a gift card deactivation operation.
type GiftCardOperationDeactivation struct{}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationReloadBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// The balance before the operation.
type GiftCardOperationReloadPreviousBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card reload operation.
type GiftCardOperationReload struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationReloadBalance `json:"balance"`
	// The balance before the operation.
	PreviousBalance *GiftCardOperationReloadPreviousBalance `json:"previous_balance"`
}

// The balance amount of a gift card, including currency and amount.
type GiftCardOperationReloadVoidBalance struct {
	// The balance amount.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// Details about a gift card reload void operation.
type GiftCardOperationReloadVoid struct {
	// The balance amount of a gift card, including currency and amount.
	Balance *GiftCardOperationReloadVoidBalance `json:"balance"`
	// The operation that was voided.
	VoidedOperation string `json:"voided_operation"`
}

// A GiftCardOperation represents an operation performed on a third-party gift card,
// such as activation, deactivation, reload, cashout, balance check, or void.
type GiftCardOperation struct {
	APIResource
	// Details about a gift card activation operation.
	Activation *GiftCardOperationActivation `json:"activation,omitempty"`
	// Details about a gift card activation void operation.
	ActivationVoid *GiftCardOperationActivationVoid `json:"activation_void,omitempty"`
	// Details about a gift card balance check operation.
	BalanceCheck *GiftCardOperationBalanceCheck `json:"balance_check,omitempty"`
	// Details about a gift card cashout operation.
	Cashout *GiftCardOperationCashout `json:"cashout,omitempty"`
	// Details about a gift card cashout void operation.
	CashoutVoid *GiftCardOperationCashoutVoid `json:"cashout_void,omitempty"`
	// The timestamp of when this operation was completed.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Details about a gift card deactivation operation.
	Deactivation *GiftCardOperationDeactivation `json:"deactivation,omitempty"`
	// The failure code of the operation. Only present if the status is failed.
	FailureCode GiftCardOperationFailureCode `json:"failure_code"`
	// The gift card this operation was performed on.
	GiftCard *GiftCard `json:"gift_card"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The connected account whose credentials were used to perform this operation.
	OnBehalfOf string `json:"on_behalf_of"`
	// Details about a gift card reload operation.
	Reload *GiftCardOperationReload `json:"reload,omitempty"`
	// Details about a gift card reload void operation.
	ReloadVoid *GiftCardOperationReloadVoid `json:"reload_void,omitempty"`
	// The status of the operation.
	Status GiftCardOperationStatus `json:"status"`
	// The type of operation performed.
	Type GiftCardOperationType `json:"type"`
}

// UnmarshalJSON handles deserialization of a GiftCardOperation.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (g *GiftCardOperation) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		g.ID = id
		return nil
	}

	type giftCardOperation GiftCardOperation
	var v giftCardOperation
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*g = GiftCardOperation(v)
	return nil
}
