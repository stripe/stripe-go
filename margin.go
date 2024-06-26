//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Retrieve a list of your margins.
type MarginListParams struct {
	ListParams `form:"*"`
	// Only return margins that are active or inactive, i.e., pass false to list all inactive margins.
	Active *bool `form:"active"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *MarginListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create a margin object to be used with invoices, invoice items, and invoice line items for a customer to represent a partner discount. A margin has a percent_off which is the percent that will be taken off the subtotal after all items and other discounts and promotions) of any invoices for a customer. Calculation of prorations do not include any partner margins applied on the original invoice item.
type MarginParams struct {
	Params `form:"*"`
	// Whether the margin can be applied to invoices, invoice items, or invoice line items or not. Defaults to `true`.
	Active *bool `form:"active"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Name of the margin, which is displayed to customers, such as on invoices.
	Name *string `form:"name"`
	// Percent that will be taken off the subtotal before tax (after all other discounts and promotions) of any invoice to which the margin is applied.
	PercentOff *float64 `form:"percent_off"`
}

// AddExpand appends a new field to expand.
func (p *MarginParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *MarginParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A (partner) margin represents a specific discount distributed in partner reseller programs to business partners who
// resell products and services and earn a discount (margin) for doing so.
type Margin struct {
	APIResource
	// Whether the margin can be applied to invoices, invoice items, or invoice line items. Defaults to `true`.
	Active bool `json:"active"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Name of the margin that's displayed on, for example, invoices.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Percent that will be taken off the subtotal before tax (after all other discounts and promotions) of any invoice to which the margin is applied.
	PercentOff float64 `json:"percent_off"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
}

// MarginList is a list of Margins as retrieved from a list endpoint.
type MarginList struct {
	APIResource
	ListMeta
	Data []*Margin `json:"data"`
}

// UnmarshalJSON handles deserialization of a Margin.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (m *Margin) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		m.ID = id
		return nil
	}

	type margin Margin
	var v margin
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*m = Margin(v)
	return nil
}
