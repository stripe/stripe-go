//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Default [tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#tax-behavior) used to specify whether the price is considered inclusive of taxes or exclusive of taxes. If the item's price has a tax behavior set, it will take precedence over the default tax behavior.
type TaxSettingsDefaultsTaxBehavior string

// List of values that TaxSettingsDefaultsTaxBehavior can take
const (
	TaxSettingsDefaultsTaxBehaviorExclusive          TaxSettingsDefaultsTaxBehavior = "exclusive"
	TaxSettingsDefaultsTaxBehaviorInclusive          TaxSettingsDefaultsTaxBehavior = "inclusive"
	TaxSettingsDefaultsTaxBehaviorInferredByCurrency TaxSettingsDefaultsTaxBehavior = "inferred_by_currency"
)

// The role of this location address.
type TaxSettingsLocationRole string

// List of values that TaxSettingsLocationRole can take
const (
	TaxSettingsLocationRoleHeadOffice TaxSettingsLocationRole = "head_office"
)

// Retrieves Tax Settings for a merchant.
type TaxSettingsParams struct {
	Params `form:"*"`
	// Default configuration to be used on Stripe Tax calculations.
	Defaults *TaxSettingsDefaultsParams `form:"defaults"`
	// The place where your business is located.
	HeadOffice *TaxSettingsHeadOfficeParams `form:"head_office"`
	// The deprecated places where your business is located.
	Locations []*TaxSettingsLocationParams `form:"locations"`
}

// Default configuration to be used on Stripe Tax calculations.
type TaxSettingsDefaultsParams struct {
	// Specifies the default [tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#tax-behavior) to be used when the item's price has unspecified tax behavior. One of inclusive, exclusive, or inferred_by_currency. Once specified, it cannot be changed back to null.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
}

// The place where your business is located.
type TaxSettingsHeadOfficeParams struct {
	// The location of the business for tax purposes.
	Address *AddressParams `form:"address"`
}

// The deprecated places where your business is located.
type TaxSettingsLocationParams struct {
	// The location of the business for tax purposes.
	Address *AddressParams `form:"address"`
	// The role of this location address.
	Role *string `form:"role"`
}
type TaxSettingsDefaults struct {
	// Default [tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#tax-behavior) used to specify whether the price is considered inclusive of taxes or exclusive of taxes. If the item's price has a tax behavior set, it will take precedence over the default tax behavior.
	TaxBehavior TaxSettingsDefaultsTaxBehavior `json:"tax_behavior"`
	// Default [tax code](https://stripe.com/docs/tax/tax-categories) used to classify your products and prices.
	TaxCode string `json:"tax_code"`
}

// The place where your business is located.
type TaxSettingsHeadOffice struct {
	Address *Address `json:"address"`
}

// The deprecated places where your business is located.
type TaxSettingsLocation struct {
	Address *Address `json:"address"`
	// The role of this location address.
	Role TaxSettingsLocationRole `json:"role"`
}

// You can use Tax `Settings` to manage configurations used by Stripe Tax calculations.
//
// Related guide: [Using the Settings API](https://stripe.com/docs/tax/settings-api)
type TaxSettings struct {
	APIResource
	Defaults *TaxSettingsDefaults `json:"defaults"`
	// The place where your business is located.
	HeadOffice *TaxSettingsHeadOffice `json:"head_office"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The deprecated places where your business is located.
	Locations []*TaxSettingsLocation `json:"locations"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
