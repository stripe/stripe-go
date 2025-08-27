//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Licensed Item objects in reverse chronological order of creation.
type V2BillingLicensedItemListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys,omitempty"`
}

// Stripe Tax details.
type V2BillingLicensedItemTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Create a Licensed Item object.
type V2BillingLicensedItemParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular billable item.
	// Maximum length of 200 characters.
	// To remove the lookup_key from the object, set it to null in the request.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Stripe Tax details.
	TaxDetails *V2BillingLicensedItemTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "seat" for Checkout to display "(price) per seat", or "environment" to display "(price) per environment".
	// Maximum length of 100 characters.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicensedItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Stripe Tax details.
type V2BillingLicensedItemCreateTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Create a Licensed Item object.
type V2BillingLicensedItemCreateParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular billable item.
	// Must be unique among billable items.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Stripe Tax details.
	TaxDetails *V2BillingLicensedItemCreateTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "seat" for Checkout to display "(price) per seat", or "environment" to display "(price) per environment".
	// Maximum length of 100 characters.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicensedItemCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Licensed Item object.
type V2BillingLicensedItemRetrieveParams struct {
	Params `form:"*"`
}

// Stripe Tax details.
type V2BillingLicensedItemUpdateTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Update a Licensed Item object. At least one of the fields is required.
type V2BillingLicensedItemUpdateParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular billable item.
	// Maximum length of 200 characters.
	// To remove the lookup_key from the object, set it to null in the request.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Stripe Tax details.
	TaxDetails *V2BillingLicensedItemUpdateTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "seat" for Checkout to display "(price) per seat", or "environment" to display "(price) per environment".
	// Maximum length of 100 characters.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicensedItemUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
