//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The address to resolve.
type V2TaxOperationResolveAddressAddressParams struct {
	// The city.
	City *string `form:"city" json:"city,omitempty"`
	// The two-letter country code.
	Country *string `form:"country" json:"country"`
	// The first line of the street address.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// The postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// The state or province.
	State *string `form:"state" json:"state,omitempty"`
}

// Resolves an address to its tax precision level.
type V2TaxOperationResolveAddressParams struct {
	Params `form:"*"`
	// The address to resolve.
	Address *V2TaxOperationResolveAddressAddressParams `form:"address" json:"address"`
}
