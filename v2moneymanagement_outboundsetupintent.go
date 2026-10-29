//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

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

// Use the OutboundSetupIntent API to create and setup usable payout methods.
type V2MoneyManagementOutboundSetupIntent struct {
	APIResource
	// Created timestamp.
	Created time.Time `json:"created"`
	// ID of the outbound setup intent.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Use the PayoutMethods API to list and interact with PayoutMethod objects.
	PayoutMethod *V2MoneyManagementPayoutMethod `json:"payout_method,omitempty"`
	// Closed Enum. Status of the outbound setup intent.
	Status V2MoneyManagementOutboundSetupIntentStatus `json:"status"`
	// The intended money movement flow this payout method should be set up for, specified in params.
	UsageIntent V2MoneyManagementOutboundSetupIntentUsageIntent `json:"usage_intent"`
}
