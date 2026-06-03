//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of the license pricing.
type V2BillingContractLicensePricingQuantityChangeLicensePricingType string

// List of values that V2BillingContractLicensePricingQuantityChangeLicensePricingType can take
const (
	V2BillingContractLicensePricingQuantityChangeLicensePricingTypeLicenseFee V2BillingContractLicensePricingQuantityChangeLicensePricingType = "license_fee"
	V2BillingContractLicensePricingQuantityChangeLicensePricingTypePrice      V2BillingContractLicensePricingQuantityChangeLicensePricingType = "price"
)

// A license pricing quantity change object returned by ListContractLicenseQuantityChanges.
type V2BillingContractLicensePricingQuantityChange struct {
	APIResource
	// The timestamp when this quantity change object was created.
	Created time.Time `json:"created"`
	// The timestamp when this quantity change takes effect.
	EffectiveAt time.Time `json:"effective_at"`
	// The ID of the quantity change object.
	ID string `json:"id"`
	// The ID of the license pricing.
	LicensePricingID string `json:"license_pricing_id"`
	// The type of the license pricing.
	LicensePricingType V2BillingContractLicensePricingQuantityChangeLicensePricingType `json:"license_pricing_type"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the pricing line associated with this quantity change.
	PricingLine string `json:"pricing_line"`
	// The quantity at the effective time.
	Quantity int64 `json:"quantity"`
}
