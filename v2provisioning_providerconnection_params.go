//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists the provider connections for the account.
type V2ProvisioningProviderConnectionListParams struct {
	Params `form:"*"`
	// Maximum number of provider connections to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Unlinks a provider connection so it can no longer be used to create resources.
type V2ProvisioningProviderConnectionUnlinkParams struct {
	Params `form:"*"`
}
