//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// The status of the template, one of `active` or `archived`.
type InvoiceRenderingTemplateStatus string

// List of values that InvoiceRenderingTemplateStatus can take
const (
	InvoiceRenderingTemplateStatusActive   InvoiceRenderingTemplateStatus = "active"
	InvoiceRenderingTemplateStatusArchived InvoiceRenderingTemplateStatus = "archived"
)

// List all templates, ordered by creation date, with the most recently created template appearing first.
type InvoiceRenderingTemplateListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	Status *string   `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an invoice rendering template with the given ID. It by default returns the latest version of the template. Optionally, specify a version to see previous versions.
type InvoiceRenderingTemplateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand  []*string `form:"expand"`
	Version *int64    `form:"version"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the status of an invoice rendering template to ‘archived' so no new Stripe objects (customers, invoices, etc.) can reference it. The template can also no longer be updated. However, if the template is already set on a Stripe object, it will continue to be applied on invoices generated by it.
type InvoiceRenderingTemplateArchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateArchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Unarchive an invoice rendering template so it can be used on new Stripe objects again.
type InvoiceRenderingTemplateUnarchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateUnarchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Invoice Rendering Templates are used to configure how invoices are rendered on surfaces like the PDF. Invoice Rendering Templates
// can be created from within the Dashboard, and they can be used over the API when creating invoices.
type InvoiceRenderingTemplate struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A brief description of the template, hidden from customers
	Nickname string `json:"nickname"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the template, one of `active` or `archived`.
	Status InvoiceRenderingTemplateStatus `json:"status"`
	// Version of this template; version increases by one when an update on the template changes any field that controls invoice rendering
	Version int64 `json:"version"`
}

// InvoiceRenderingTemplateList is a list of InvoiceRenderingTemplates as retrieved from a list endpoint.
type InvoiceRenderingTemplateList struct {
	APIResource
	ListMeta
	Data []*InvoiceRenderingTemplate `json:"data"`
}

// UnmarshalJSON handles deserialization of an InvoiceRenderingTemplate.
// This custom unmarshaling is needed to handle the time fields correctly.
func (i *InvoiceRenderingTemplate) UnmarshalJSON(data []byte) error {
	type invoiceRenderingTemplate InvoiceRenderingTemplate
	v := struct {
		Created int64 `json:"created"`
		*invoiceRenderingTemplate
	}{
		invoiceRenderingTemplate: (*invoiceRenderingTemplate)(i),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of an InvoiceRenderingTemplate.
// This custom marshaling is needed to handle the time fields correctly.
func (i InvoiceRenderingTemplate) MarshalJSON() ([]byte, error) {
	type invoiceRenderingTemplate InvoiceRenderingTemplate
	v := struct {
		Created int64 `json:"created"`
		invoiceRenderingTemplate
	}{
		invoiceRenderingTemplate: (invoiceRenderingTemplate)(i),
		Created:                  i.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
