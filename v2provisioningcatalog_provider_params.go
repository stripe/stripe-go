//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists providers available in the catalog.
type V2ProvisioningCatalogProviderListParams struct {
	Params `form:"*"`
	// Catalog partition to list providers from.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// When `true`, list development-only providers. When unset or `false`, development providers are
	// excluded.
	Development *bool `form:"development" json:"development,omitempty"`
	// Maximum number of providers to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}
