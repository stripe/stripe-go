//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	Expand []*string `form:"expand" json:"expand,omitempty"`
	Status *string   `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls how individual line items are displayed on the invoice, visible to customers.
type InvoiceRenderingTemplateLineItemPolicyParams struct {
	// A CEL expression that takes a line item and returns a boolean indicating whether to display it on the invoice.
	Display *string `form:"display" json:"display,omitempty"`
}

// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
type InvoiceRenderingTemplateSummaryItemPolicyParams struct {
	// Controls how each summary group is rendered on the invoice. `all` displays the summary line item, its constituent line items, and a group subtotal. `none` displays only the summary line item. Defaults to `all`.
	DisplayItems *string `form:"display_items" json:"display_items,omitempty"`
	// A CEL expression that takes a line item and returns a boolean indicating whether to include it in this grouping policy. If omitted, includes all ungrouped line items.
	Filter *string `form:"filter" json:"filter,omitempty"`
	// A CEL expression that takes a line item and returns a string. Line items with the same result are grouped into a summary item, using that result as its group title.
	GroupBy *string `form:"group_by" json:"group_by,omitempty"`
}

// Creates an invoice rendering template.
type InvoiceRenderingTemplateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls how individual line items are displayed on the invoice, visible to customers.
	LineItemPolicy *InvoiceRenderingTemplateLineItemPolicyParams `form:"line_item_policy" json:"line_item_policy,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Nickname for the template that is hidden from customers
	Nickname *string `form:"nickname" json:"nickname,omitempty"`
	// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
	SummaryItemPolicies []*InvoiceRenderingTemplateSummaryItemPolicyParams `form:"summary_item_policies" json:"summary_item_policies,omitempty"`
	Version             *int64                                             `form:"version" json:"version,omitempty"`
	UnsetFields         []InvoiceRenderingTemplateParamsUnsetField         `form:"-" json:"-"`
}

// InvoiceRenderingTemplateParamsUnsetField is the list of fields that can be cleared/unset on InvoiceRenderingTemplateParams.
type InvoiceRenderingTemplateParamsUnsetField string

const (
	InvoiceRenderingTemplateParamsUnsetFieldLineItemPolicy      InvoiceRenderingTemplateParamsUnsetField = "line_item_policy"
	InvoiceRenderingTemplateParamsUnsetFieldMetadata            InvoiceRenderingTemplateParamsUnsetField = "metadata"
	InvoiceRenderingTemplateParamsUnsetFieldSummaryItemPolicies InvoiceRenderingTemplateParamsUnsetField = "summary_item_policies"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *InvoiceRenderingTemplateParams) AddUnsetField(field InvoiceRenderingTemplateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceRenderingTemplateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Updates the status of an invoice rendering template to ‘archived' so no new Stripe objects (customers, invoices, etc.) can reference it. The template can also no longer be updated. However, if the template is already set on a Stripe object, it will continue to be applied on invoices generated by it.
type InvoiceRenderingTemplateArchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateArchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Unarchive an invoice rendering template so it can be used on new Stripe objects again.
type InvoiceRenderingTemplateUnarchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateUnarchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls how individual line items are displayed on the invoice, visible to customers.
type InvoiceRenderingTemplateCreateLineItemPolicyParams struct {
	// A CEL expression that takes a line item and returns a boolean indicating whether to display it on the invoice.
	Display *string `form:"display" json:"display,omitempty"`
}

// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
type InvoiceRenderingTemplateCreateSummaryItemPolicyParams struct {
	// Controls how each summary group is rendered on the invoice. `all` displays the summary line item, its constituent line items, and a group subtotal. `none` displays only the summary line item. Defaults to `all`.
	DisplayItems *string `form:"display_items" json:"display_items,omitempty"`
	// A CEL expression that takes a line item and returns a boolean indicating whether to include it in this grouping policy. If omitted, includes all ungrouped line items.
	Filter *string `form:"filter" json:"filter,omitempty"`
	// A CEL expression that takes a line item and returns a string. Line items with the same result are grouped into a summary item, using that result as its group title.
	GroupBy *string `form:"group_by" json:"group_by,omitempty"`
}

// Creates an invoice rendering template.
type InvoiceRenderingTemplateCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls how individual line items are displayed on the invoice, visible to customers.
	LineItemPolicy *InvoiceRenderingTemplateCreateLineItemPolicyParams `form:"line_item_policy" json:"line_item_policy,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Nickname for the template that is hidden from customers
	Nickname *string `form:"nickname" json:"nickname,omitempty"`
	// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
	SummaryItemPolicies []*InvoiceRenderingTemplateCreateSummaryItemPolicyParams `form:"summary_item_policies" json:"summary_item_policies,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceRenderingTemplateCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves an invoice rendering template with the given ID. It by default returns the latest version of the template. Optionally, specify a version to see previous versions.
type InvoiceRenderingTemplateRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand  []*string `form:"expand" json:"expand,omitempty"`
	Version *int64    `form:"version" json:"version,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls how individual line items are displayed on the invoice, visible to customers.
type InvoiceRenderingTemplateUpdateLineItemPolicyParams struct {
	// A CEL expression that takes a line item and returns a boolean indicating whether to display it on the invoice.
	Display *string `form:"display" json:"display,omitempty"`
}

// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
type InvoiceRenderingTemplateUpdateSummaryItemPolicyParams struct {
	// Controls how each summary group is rendered on the invoice. `all` displays the summary line item, its constituent line items, and a group subtotal. `none` displays only the summary line item. Defaults to `all`.
	DisplayItems *string `form:"display_items" json:"display_items,omitempty"`
	// A CEL expression that takes a line item and returns a boolean indicating whether to include it in this grouping policy. If omitted, includes all ungrouped line items.
	Filter *string `form:"filter" json:"filter,omitempty"`
	// A CEL expression that takes a line item and returns a string. Line items with the same result are grouped into a summary item, using that result as its group title.
	GroupBy *string `form:"group_by" json:"group_by,omitempty"`
}

// Updates an existing invoice rendering template.
type InvoiceRenderingTemplateUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls how individual line items are displayed on the invoice, visible to customers.
	LineItemPolicy *InvoiceRenderingTemplateUpdateLineItemPolicyParams `form:"line_item_policy" json:"line_item_policy,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Nickname for the template that is hidden from customers
	Nickname *string `form:"nickname" json:"nickname,omitempty"`
	// Controls the grouping of line items into summary items which are displayed on the invoice, visible to customers. The policies are applied in order - only ungrouped line items from the previous policy are considered by the following policies.
	SummaryItemPolicies []*InvoiceRenderingTemplateUpdateSummaryItemPolicyParams `form:"summary_item_policies" json:"summary_item_policies,omitempty"`
	UnsetFields         []InvoiceRenderingTemplateUpdateParamsUnsetField         `form:"-" json:"-"`
}

// InvoiceRenderingTemplateUpdateParamsUnsetField is the list of fields that can be cleared/unset on InvoiceRenderingTemplateUpdateParams.
type InvoiceRenderingTemplateUpdateParamsUnsetField string

const (
	InvoiceRenderingTemplateUpdateParamsUnsetFieldLineItemPolicy      InvoiceRenderingTemplateUpdateParamsUnsetField = "line_item_policy"
	InvoiceRenderingTemplateUpdateParamsUnsetFieldMetadata            InvoiceRenderingTemplateUpdateParamsUnsetField = "metadata"
	InvoiceRenderingTemplateUpdateParamsUnsetFieldSummaryItemPolicies InvoiceRenderingTemplateUpdateParamsUnsetField = "summary_item_policies"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *InvoiceRenderingTemplateUpdateParams) AddUnsetField(field InvoiceRenderingTemplateUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *InvoiceRenderingTemplateUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceRenderingTemplateUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The policy that controls how individual line items are rendered
type InvoiceRenderingTemplateLineItemPolicy struct {
	// A CEL expression that takes in a line item and returns a string indicating whether the line item should be 'hidden_if_zero' or 'visible'.
	//  If 'hidden_if_zero', the line item will be hidden if the line item's amount is zero on the invoice.
	//  If 'visible', the line item will be shown regardless of the line item's amount.
	Display string `json:"display"`
}

// The policies that control how line items are grouped into summary items
type InvoiceRenderingTemplateSummaryItemPolicy struct {
	// Either `all` or `none`. For all summary items created from this policy, if `all`, every line item will be shown under the summary item header. If `none`, the line items will be hidden and only the summary item header will be shown.
	DisplayItems string `json:"display_items"`
	// A CEL expression that takes in a line item and returns a boolean indicating whether the line item should be considered when grouping.
	Filter string `json:"filter"`
	// A CEL expression that takes in a line item and returns a string. Line items with the same string results are put in the same summary items. The string result will become the description of the summary item.
	GroupBy string `json:"group_by"`
}

// Invoice Rendering Templates are used to configure how invoices are rendered on surfaces like the PDF. Invoice Rendering Templates
// can be created from within the Dashboard, and they can be used over the API when creating invoices.
type InvoiceRenderingTemplate struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The policy that controls how individual line items are rendered
	LineItemPolicy *InvoiceRenderingTemplateLineItemPolicy `json:"line_item_policy,omitempty"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A brief description of the template, hidden from customers
	Nickname string `json:"nickname"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the template, one of `active` or `archived`.
	Status InvoiceRenderingTemplateStatus `json:"status"`
	// The policies that control how line items are grouped into summary items
	SummaryItemPolicies []*InvoiceRenderingTemplateSummaryItemPolicy `json:"summary_item_policies,omitempty"`
	// Version of this template; version increases by one when an update on the template changes any field that controls invoice rendering
	Version int64 `json:"version"`
}

// InvoiceRenderingTemplateList is a list of InvoiceRenderingTemplates as retrieved from a list endpoint.
type InvoiceRenderingTemplateList struct {
	APIResource
	ListMeta
	Data []*InvoiceRenderingTemplate `json:"data"`
}
