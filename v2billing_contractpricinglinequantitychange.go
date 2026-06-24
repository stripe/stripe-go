//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of pricing.
type V2BillingContractPricingLineQuantityChangePricingType string

// List of values that V2BillingContractPricingLineQuantityChangePricingType can take
const (
	V2BillingContractPricingLineQuantityChangePricingTypePrice V2BillingContractPricingLineQuantityChangePricingType = "price"
)

// The pricing configuration for the associated pricing line.
type V2BillingContractPricingLineQuantityChangePricing struct {
	// The ID of the V1 price. Present when `type` is `price`.
	Price string `json:"price,omitempty"`
	// The type of pricing.
	Type V2BillingContractPricingLineQuantityChangePricingType `json:"type"`
}

// A quantity change object for a pricing line, returned by ListContractPricingLineQuantityChanges.
type V2BillingContractPricingLineQuantityChange struct {
	APIResource
	// The timestamp when this quantity change object was created.
	Created time.Time `json:"created"`
	// The timestamp when this quantity change takes effect.
	EffectiveAt time.Time `json:"effective_at"`
	// The ID of the quantity change object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The pricing configuration for the associated pricing line.
	Pricing *V2BillingContractPricingLineQuantityChangePricing `json:"pricing"`
	// The ID of the pricing line associated with this quantity change.
	PricingLine string `json:"pricing_line"`
	// The quantity at the effective time.
	Quantity float64 `json:"quantity,string"`
}
