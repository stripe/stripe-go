//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of next action.
type V2MoneyManagementOutboundSetupIntentNextActionType string

// List of values that V2MoneyManagementOutboundSetupIntentNextActionType can take
const (
	V2MoneyManagementOutboundSetupIntentNextActionTypeConfirmationOfPayee V2MoneyManagementOutboundSetupIntentNextActionType = "confirmation_of_payee"
)

// The Confirmation of Payee status.
type V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus string

// List of values that V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus can take
const (
	V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatusAwaitingAcknowledgement V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus = "awaiting_acknowledgement"
	V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatusConfirmed               V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus = "confirmed"
	V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatusUninitiated             V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus = "uninitiated"
)

// Closed Enum. Status of the outbound setup intent.
type V2MoneyManagementOutboundSetupIntentStatus string

// List of values that V2MoneyManagementOutboundSetupIntentStatus can take
const (
	V2MoneyManagementOutboundSetupIntentStatusCanceled             V2MoneyManagementOutboundSetupIntentStatus = "canceled"
	V2MoneyManagementOutboundSetupIntentStatusRequiresAction       V2MoneyManagementOutboundSetupIntentStatus = "requires_action"
	V2MoneyManagementOutboundSetupIntentStatusRequiresPayoutMethod V2MoneyManagementOutboundSetupIntentStatus = "requires_payout_method"
	V2MoneyManagementOutboundSetupIntentStatusSucceeded            V2MoneyManagementOutboundSetupIntentStatus = "succeeded"
)

// The intended money movement flow this payout method should be set up for, specified in params.
type V2MoneyManagementOutboundSetupIntentUsageIntent string

// List of values that V2MoneyManagementOutboundSetupIntentUsageIntent can take
const (
	V2MoneyManagementOutboundSetupIntentUsageIntentPayment  V2MoneyManagementOutboundSetupIntentUsageIntent = "payment"
	V2MoneyManagementOutboundSetupIntentUsageIntentTransfer V2MoneyManagementOutboundSetupIntentUsageIntent = "transfer"
)

// Confirmation of Payee details.
type V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayee struct {
	// The type of the credential.
	Object string `json:"object"`
	// The Confirmation of Payee status.
	Status V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayeeStatus `json:"status"`
}

// Specifies which actions needs to be taken next to continue setup of the credential.
type V2MoneyManagementOutboundSetupIntentNextAction struct {
	// Confirmation of Payee details.
	ConfirmationOfPayee *V2MoneyManagementOutboundSetupIntentNextActionConfirmationOfPayee `json:"confirmation_of_payee"`
	// The type of next action.
	Type V2MoneyManagementOutboundSetupIntentNextActionType `json:"type"`
}

// Use the OutboundSetupIntent API to create and setup usable payout methods.
type V2MoneyManagementOutboundSetupIntent struct {
	APIResource
	// Created timestamp.
	Created time.Time `json:"created"`
	// ID of the outbound setup intent.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Specifies which actions needs to be taken next to continue setup of the credential.
	NextAction *V2MoneyManagementOutboundSetupIntentNextAction `json:"next_action"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Information about the payout method that's created and linked to this outbound setup intent.
	PayoutMethod *V2MoneyManagementPayoutMethod `json:"payout_method"`
	// Closed Enum. Status of the outbound setup intent.
	Status V2MoneyManagementOutboundSetupIntentStatus `json:"status"`
	// The intended money movement flow this payout method should be set up for, specified in params.
	UsageIntent V2MoneyManagementOutboundSetupIntentUsageIntent `json:"usage_intent"`
}
