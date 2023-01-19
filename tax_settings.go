//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	// The places where your business is located.
	Locations []*TaxSettingsLocationParams `form:"locations"`
}

// Default configuration to be used on Stripe Tax calculations.
type TaxSettingsDefaultsParams struct {
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
}

// The places where your business is located.
type TaxSettingsLocationParams struct {
	// The location of the business for tax purposes.
	Address *AddressParams `form:"address"`
	// The role of this location address.
	Role *string `form:"role"`
}
type TaxSettingsDefaults struct {
	// Default [tax code](https://stripe.com/docs/tax/tax-categories) used to classify your products and prices.
	TaxCode string `json:"tax_code"`
}

// The places where your business is located.
type TaxSettingsLocation struct {
	Address *Address `json:"address"`
	// The role of this location address.
	Role TaxSettingsLocationRole `json:"role"`
}

// You can use Tax `Settings` to manage configurations used by Stripe Tax calculations.
//
// Related guide: [Account settings](https://stripe.com/docs/tax/connect/settings).
type TaxSettings struct {
	APIResource
	Defaults *TaxSettingsDefaults `json:"defaults"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The places where your business is located.
	Locations []*TaxSettingsLocation `json:"locations"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
