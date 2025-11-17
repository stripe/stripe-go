//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List Billing Profiles.
type V2BillingProfileListParams struct {
	Params `form:"*"`
	// Filter billing profiles by a customer. Mutually exclusive
	// with `lookup_keys` and `default_payment_method`.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Filter billing profiles by a default payment method. Mutually exclusive
	// with `customer` and `lookup_keys`.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter billing profiles by lookup keys. Mutually exclusive
	// with `customer` and `default_payment_method`.
	// You can specify up to 10 lookup_keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys"`
	// Filter billing profiles by status. Can be combined
	// with all other filters. If not provided, all billing profiles will be returned.
	Status *string `form:"status" json:"status,omitempty"`
}

// Create a BillingProfile object.
type V2BillingProfileParams struct {
	Params `form:"*"`
	// The ID of the customer object.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of the payment method object.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A customer-facing name for the billing profile.
	// Maximum length of 250 characters.
	// To remove the display_name from the object, set it to null in the request.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular billing profile. It must be unique among billing profiles for a given customer.
	// Maximum length of 200 characters.
	// To remove the lookup_key from the object, set it to null in the request.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingProfileParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}

// Create a BillingProfile object.
type V2BillingProfileCreateParams struct {
	Params `form:"*"`
	// The ID of the customer object.
	Customer *string `form:"customer" json:"customer"`
	// The ID of the payment method object.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A customer-facing name for the billing profile.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular billing profile. It must be unique among billing profiles for a given customer.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingProfileCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a BillingProfile object.
type V2BillingProfileRetrieveParams struct {
	Params `form:"*"`
}

// Update a BillingProfile object.
type V2BillingProfileUpdateParams struct {
	Params `form:"*"`
	// The ID of the payment method object.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A customer-facing name for the billing profile.
	// Maximum length of 250 characters.
	// To remove the display_name from the object, set it to null in the request.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// An internal key you can use to search for a particular billing profile. It must be unique among billing profiles for a given customer.
	// Maximum length of 200 characters.
	// To remove the lookup_key from the object, set it to null in the request.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingProfileUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}
