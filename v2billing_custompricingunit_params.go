//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Custom Pricing Unit objects.
type V2BillingCustomPricingUnitListParams struct {
	Params `form:"*"`
	// Filter for active/inactive custom pricing units. Mutually exclusive with `lookup_keys`.
	Active *bool `form:"active" json:"active,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys. Mutually exclusive with `active`.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys,omitempty"`
}

// Create a Custom Pricing Unit object.
type V2BillingCustomPricingUnitParams struct {
	Params `form:"*"`
	// Whether the Custom Pricing Unit is active.
	Active *bool `form:"active" json:"active,omitempty"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 10 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular custom pricing unit item.
	// Must be unique among items.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCustomPricingUnitParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Create a Custom Pricing Unit object.
type V2BillingCustomPricingUnitCreateParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 10 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular custom pricing unit item.
	// Must be unique among items.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCustomPricingUnitCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Custom Pricing Unit object.
type V2BillingCustomPricingUnitRetrieveParams struct {
	Params `form:"*"`
}

// Update a Custom Pricing Unit object.
type V2BillingCustomPricingUnitUpdateParams struct {
	Params `form:"*"`
	// Whether the Custom Pricing Unit is active.
	Active *bool `form:"active" json:"active,omitempty"`
	// Description that customers will see in the invoice line item.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular CustomPricingUnit item.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCustomPricingUnitUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
