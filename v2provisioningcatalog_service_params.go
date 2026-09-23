//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists services available in the catalog.
type V2ProvisioningCatalogServiceListParams struct {
	Params `form:"*"`
	// Catalog partition to list services from.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// When `true`, list development-only services. When unset or `false`, development services are
	// excluded.
	Development *bool `form:"development" json:"development,omitempty"`
	// Maximum number of services to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filters services to those offered by the provider with this name.
	ProviderName *string `form:"provider_name" json:"provider_name,omitempty"`
}
