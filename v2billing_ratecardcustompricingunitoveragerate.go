//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A rate card custom pricing unit overage rate defines the conversion rate from a custom pricing unit
// to fiat currency when credits are exhausted. This enables automatic overage charges
// at a configurable per-unit rate.
type V2BillingRateCardCustomPricingUnitOverageRate struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the custom pricing unit this overage rate applies to.
	CustomPricingUnit string `json:"custom_pricing_unit"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A one-time item represents a product that can be charged as a one-time line item,
	// used for overage charges when custom pricing unit credits are exhausted.
	OneTimeItem *V2BillingOneTimeItem `json:"one_time_item"`
	// The ID of the Rate Card this overage rate belongs to.
	RateCard string `json:"rate_card"`
	// The ID of the Rate Card Version this overage rate was created on.
	RateCardVersion string `json:"rate_card_version"`
	// The per-unit amount to charge for overages, represented as a decimal string in minor currency units with at most 12 decimal places.
	UnitAmount string `json:"unit_amount"`
}
