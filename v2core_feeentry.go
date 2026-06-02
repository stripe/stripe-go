//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of entity that charged this fee.
type V2CoreFeeEntryChargedByType string

// List of values that V2CoreFeeEntryChargedByType can take
const (
	V2CoreFeeEntryChargedByTypeApplication V2CoreFeeEntryChargedByType = "application"
	V2CoreFeeEntryChargedByTypeNetwork     V2CoreFeeEntryChargedByType = "network"
	V2CoreFeeEntryChargedByTypeStripe      V2CoreFeeEntryChargedByType = "stripe"
)

// The reason this fee entry was created.
type V2CoreFeeEntryReason string

// List of values that V2CoreFeeEntryReason can take
const (
	V2CoreFeeEntryReasonOther         V2CoreFeeEntryReason = "other"
	V2CoreFeeEntryReasonProcessingFee V2CoreFeeEntryReason = "processing_fee"
	V2CoreFeeEntryReasonRefund        V2CoreFeeEntryReason = "refund"
	V2CoreFeeEntryReasonRefundFailure V2CoreFeeEntryReason = "refund_failure"
	V2CoreFeeEntryReasonReprice       V2CoreFeeEntryReason = "reprice"
	V2CoreFeeEntryReasonTierTrueUp    V2CoreFeeEntryReason = "tier_true_up"
)

// The category of this fee.
type V2CoreFeeEntryType string

// List of values that V2CoreFeeEntryType can take
const (
	V2CoreFeeEntryTypeApplicationFee V2CoreFeeEntryType = "application_fee"
	V2CoreFeeEntryTypePassthroughFee V2CoreFeeEntryType = "passthrough_fee"
	V2CoreFeeEntryTypeStripeFee      V2CoreFeeEntryType = "stripe_fee"
)

// Details for a fee charged by a Connect application.
type V2CoreFeeEntryChargedByApplication struct {
	// Human-readable product name, e.g. "Card payments - Stripe fee".
	FeatureName string `json:"feature_name,omitempty"`
}

// Details for a fee charged by the payment network.
type V2CoreFeeEntryChargedByNetwork struct {
	// Human-readable product name, e.g. "Card payments - Stripe fee".
	FeatureName string `json:"feature_name,omitempty"`
}

// Details for a fee charged by Stripe.
type V2CoreFeeEntryChargedByStripe struct {
	// Human-readable product name, e.g. "Card payments - Stripe fee".
	FeatureName string `json:"feature_name,omitempty"`
}

// The entity that assessed this fee.
type V2CoreFeeEntryChargedBy struct {
	// Details for a fee charged by a Connect application.
	Application *V2CoreFeeEntryChargedByApplication `json:"application,omitempty"`
	// Details for a fee charged by the payment network.
	Network *V2CoreFeeEntryChargedByNetwork `json:"network,omitempty"`
	// Details for a fee charged by Stripe.
	Stripe *V2CoreFeeEntryChargedByStripe `json:"stripe,omitempty"`
	// The type of entity that charged this fee.
	Type V2CoreFeeEntryChargedByType `json:"type"`
}

// The usage event that caused this fee to be assessed.
type V2CoreFeeEntryIncurredBy struct {
	// The account that incurred the usage (may differ from the billing account).
	Account string `json:"account,omitempty"`
	// Public API object id, e.g. ch_xxx.
	ID string `json:"id"`
	// Timestamp of when the usage event occurred.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
	// Public API object type: "charge", "payment", "refund", "dispute", "payout", etc.
	Type string `json:"type"`
}

// The tax portion of the fee, if applicable.
type V2CoreFeeEntryTax struct {
	// The tax amount calculated for this fee.
	Amount Amount `json:"amount"`
}

// A FeeEntry is the atomic, append-only record of an assessed fee.
type V2CoreFeeEntry struct {
	APIResource
	// The fee amount.
	Amount Amount `json:"amount"`
	// The entity that assessed this fee.
	ChargedBy *V2CoreFeeEntryChargedBy `json:"charged_by"`
	// Timestamp of when this fee entry was created.
	Created time.Time `json:"created"`
	// The ID of the FeeBatch that collected this fee, if any.
	FeeBatch string `json:"fee_batch,omitempty"`
	// Unique identifier for the FeeEntry.
	ID string `json:"id"`
	// The usage event that caused this fee to be assessed.
	IncurredBy *V2CoreFeeEntryIncurredBy `json:"incurred_by"`
	// Has the value `true` if the object exists in live mode, or `false` if in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The reason this fee entry was created.
	Reason V2CoreFeeEntryReason `json:"reason"`
	// The tax portion of the fee, if applicable.
	Tax *V2CoreFeeEntryTax `json:"tax,omitempty"`
	// The category of this fee.
	Type V2CoreFeeEntryType `json:"type"`
}
