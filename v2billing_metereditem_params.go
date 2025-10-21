//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Metered Item objects in reverse chronological order of creation.
type V2BillingMeteredItemListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys,flat_array" json:"lookup_keys,omitempty"`
}

// Optional array of Meter segments to filter event dimension keys for billing.
type V2BillingMeteredItemMeterSegmentConditionParams struct {
	// A Meter dimension.
	Dimension *string `form:"dimension" json:"dimension"`
	// To count usage towards this metered item, the dimension must have this value.
	Value *string `form:"value" json:"value"`
}

// Stripe Tax details.
type V2BillingMeteredItemTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Create a Metered Item object.
type V2BillingMeteredItemParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Optional array of Meter dimensions to group event dimension keys for invoice line items.
	InvoicePresentationDimensions []*string `form:"invoice_presentation_dimensions,flat_array" json:"invoice_presentation_dimensions,omitempty"`
	// An internal key you can use to search for a particular billable item.
	// Maximum length of 200 characters.
	// To remove the lookup_key from the object, set it to null in the request.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// ID of the Meter that measures usage for this Metered Item.
	Meter *string `form:"meter" json:"meter,omitempty"`
	// Optional array of Meter segments to filter event dimension keys for billing.
	MeterSegmentConditions []*V2BillingMeteredItemMeterSegmentConditionParams `form:"meter_segment_conditions,flat_array" json:"meter_segment_conditions,omitempty"`
	// Stripe Tax details.
	TaxDetails *V2BillingMeteredItemTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "CPU-hour" for Checkout to display "(price) per CPU-hour", or "1 million events" to display "(price) per 1
	// million events".
	// Maximum length of 100 characters.
	// To remove the unit_label from the object, set it to null in the request.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingMeteredItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Optional array of Meter segments to filter event dimension keys for billing.
type V2BillingMeteredItemCreateMeterSegmentConditionParams struct {
	// A Meter dimension.
	Dimension *string `form:"dimension" json:"dimension"`
	// To count usage towards this metered item, the dimension must have this value.
	Value *string `form:"value" json:"value"`
}

// Stripe Tax details.
type V2BillingMeteredItemCreateTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Create a Metered Item object.
type V2BillingMeteredItemCreateParams struct {
	Params `form:"*"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Optional array of Meter dimensions to group event dimension keys for invoice line items.
	InvoicePresentationDimensions []*string `form:"invoice_presentation_dimensions,flat_array" json:"invoice_presentation_dimensions,omitempty"`
	// An internal key you can use to search for a particular billable item.
	// Must be unique among billable items.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// ID of the Meter that measures usage for this Metered Item.
	Meter *string `form:"meter" json:"meter"`
	// Optional array of Meter segments to filter event dimension keys for billing.
	MeterSegmentConditions []*V2BillingMeteredItemCreateMeterSegmentConditionParams `form:"meter_segment_conditions,flat_array" json:"meter_segment_conditions,omitempty"`
	// Stripe Tax details.
	TaxDetails *V2BillingMeteredItemCreateTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "CPU-hour" for Checkout to display "(price) per CPU-hour", or "1 million events" to display "(price) per 1
	// million events".
	// Maximum length of 100 characters.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingMeteredItemCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Metered Item object.
type V2BillingMeteredItemRetrieveParams struct {
	Params `form:"*"`
}

// Stripe Tax details.
type V2BillingMeteredItemUpdateTaxDetailsParams struct {
	// Product tax code (PTC).
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Update a Metered Item object. At least one of the fields is required.
type V2BillingMeteredItemUpdateParams struct {
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
	TaxDetails *V2BillingMeteredItemUpdateTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "CPU-hour" for Checkout to display "(price) per CPU-hour", or "1 million events" to display "(price) per 1
	// million events".
	// Maximum length of 100 characters.
	// To remove the unit_label from the object, set it to null in the request.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingMeteredItemUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
