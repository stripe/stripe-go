//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of payment plans.
type PaymentPlanListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return payment plans associated with the given invoice.
	Invoice *string `form:"invoice" json:"invoice,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentPlanListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details of the invoice this payment plan collects on.
type PaymentPlanCollectsOnInvoiceDetailsParams struct {
	// The ID of the invoice.
	Invoice *string `form:"invoice" json:"invoice"`
}

// The invoice(s) this payment plan collects on. Currently must contain exactly one invoice entry.
type PaymentPlanCollectsOnParams struct {
	// Details of the invoice this payment plan collects on.
	InvoiceDetails *PaymentPlanCollectsOnInvoiceDetailsParams `form:"invoice_details" json:"invoice_details"`
	// The type of object this plan collects on. Currently always `invoice_details`.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'relative'.
type PaymentPlanScheduleAmountsDueAmountDueDateRelativeParams struct {
	// The number of intervals after finalization.
	Count *int64 `form:"count" json:"count"`
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
}

// When this installment is due.
type PaymentPlanScheduleAmountsDueAmountDueDateParams struct {
	// Unix timestamp. Required when type is 'absolute'.
	Absolute *int64 `form:"absolute" json:"absolute,omitempty"`
	// Required when type is 'relative'.
	Relative *PaymentPlanScheduleAmountsDueAmountDueDateRelativeParams `form:"relative" json:"relative,omitempty"`
	// Either 'absolute' or 'relative'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'fixed_amount'.
type PaymentPlanScheduleAmountsDueAmountFixedAmountParams struct {
	// The installment amount in minor units.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter ISO currency code.
	Currency *string `form:"currency" json:"currency"`
}

// The list of installment entries.
type PaymentPlanScheduleAmountsDueAmountParams struct {
	// Optional description for this installment.
	Description *string `form:"description" json:"description,omitempty"`
	// When this installment is due.
	DueDate *PaymentPlanScheduleAmountsDueAmountDueDateParams `form:"due_date" json:"due_date,omitempty"`
	// Required when type is 'fixed_amount'.
	FixedAmount *PaymentPlanScheduleAmountsDueAmountFixedAmountParams `form:"fixed_amount" json:"fixed_amount,omitempty"`
	// Optional stable identifier for the installment entry.
	ID *string `form:"id" json:"id,omitempty"`
	// The installment percentage of the total. Required when type is 'percentage'.
	Percentage *float64 `form:"percentage" json:"percentage,omitempty"`
	// Either 'fixed_amount' or 'percentage'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'amounts_due'.
type PaymentPlanScheduleAmountsDueParams struct {
	// The list of installment entries.
	Amounts []*PaymentPlanScheduleAmountsDueAmountParams `form:"amounts" json:"amounts"`
}

// The schedule defining how to split the invoice total into installments.
type PaymentPlanScheduleParams struct {
	// Required when type is 'amounts_due'.
	AmountsDue *PaymentPlanScheduleAmountsDueParams `form:"amounts_due" json:"amounts_due"`
	// The schedule type. Currently only 'amounts_due' is supported.
	Type *string `form:"type" json:"type"`
}

// Creates a payment plan that splits a single invoice obligation into installments with their own due dates and amounts.
type PaymentPlanParams struct {
	Params `form:"*"`
	// The invoice(s) this payment plan collects on. Currently must contain exactly one invoice entry.
	CollectsOn []*PaymentPlanCollectsOnParams `form:"collects_on" json:"collects_on,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The schedule defining how to split the invoice total into installments.
	Schedule *PaymentPlanScheduleParams `form:"schedule" json:"schedule,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentPlanParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentPlanParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details of the invoice this payment plan collects on.
type PaymentPlanCreateCollectsOnInvoiceDetailsParams struct {
	// The ID of the invoice.
	Invoice *string `form:"invoice" json:"invoice"`
}

// The invoice(s) this payment plan collects on. Currently must contain exactly one invoice entry.
type PaymentPlanCreateCollectsOnParams struct {
	// Details of the invoice this payment plan collects on.
	InvoiceDetails *PaymentPlanCreateCollectsOnInvoiceDetailsParams `form:"invoice_details" json:"invoice_details"`
	// The type of object this plan collects on. Currently always `invoice_details`.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'relative'.
type PaymentPlanCreateScheduleAmountsDueAmountDueDateRelativeParams struct {
	// The number of intervals after finalization.
	Count *int64 `form:"count" json:"count"`
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
}

// When this installment is due.
type PaymentPlanCreateScheduleAmountsDueAmountDueDateParams struct {
	// Unix timestamp. Required when type is 'absolute'.
	Absolute *int64 `form:"absolute" json:"absolute,omitempty"`
	// Required when type is 'relative'.
	Relative *PaymentPlanCreateScheduleAmountsDueAmountDueDateRelativeParams `form:"relative" json:"relative,omitempty"`
	// Either 'absolute' or 'relative'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'fixed_amount'.
type PaymentPlanCreateScheduleAmountsDueAmountFixedAmountParams struct {
	// The installment amount in minor units.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter ISO currency code.
	Currency *string `form:"currency" json:"currency"`
}

// The list of installment entries.
type PaymentPlanCreateScheduleAmountsDueAmountParams struct {
	// Optional description for this installment.
	Description *string `form:"description" json:"description,omitempty"`
	// When this installment is due.
	DueDate *PaymentPlanCreateScheduleAmountsDueAmountDueDateParams `form:"due_date" json:"due_date,omitempty"`
	// Required when type is 'fixed_amount'.
	FixedAmount *PaymentPlanCreateScheduleAmountsDueAmountFixedAmountParams `form:"fixed_amount" json:"fixed_amount,omitempty"`
	// Optional stable identifier for the installment entry.
	ID *string `form:"id" json:"id,omitempty"`
	// The installment percentage of the total. Required when type is 'percentage'.
	Percentage *float64 `form:"percentage" json:"percentage,omitempty"`
	// Either 'fixed_amount' or 'percentage'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'amounts_due'.
type PaymentPlanCreateScheduleAmountsDueParams struct {
	// The list of installment entries.
	Amounts []*PaymentPlanCreateScheduleAmountsDueAmountParams `form:"amounts" json:"amounts"`
}

// The schedule defining how to split the invoice total into installments.
type PaymentPlanCreateScheduleParams struct {
	// Required when type is 'amounts_due'.
	AmountsDue *PaymentPlanCreateScheduleAmountsDueParams `form:"amounts_due" json:"amounts_due"`
	// The schedule type. Currently only 'amounts_due' is supported.
	Type *string `form:"type" json:"type"`
}

// Creates a payment plan that splits a single invoice obligation into installments with their own due dates and amounts.
type PaymentPlanCreateParams struct {
	Params `form:"*"`
	// The invoice(s) this payment plan collects on. Currently must contain exactly one invoice entry.
	CollectsOn []*PaymentPlanCreateCollectsOnParams `form:"collects_on" json:"collects_on"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The schedule defining how to split the invoice total into installments.
	Schedule *PaymentPlanCreateScheduleParams `form:"schedule" json:"schedule"`
}

// AddExpand appends a new field to expand.
func (p *PaymentPlanCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentPlanCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the payment plan with the given ID.
type PaymentPlanRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentPlanRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Required when type is 'relative'.
type PaymentPlanUpdateScheduleAmountsDueAmountDueDateRelativeParams struct {
	// The number of intervals after finalization.
	Count *int64 `form:"count" json:"count"`
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
}

// When this installment is due.
type PaymentPlanUpdateScheduleAmountsDueAmountDueDateParams struct {
	// Unix timestamp. Required when type is 'absolute'.
	Absolute *int64 `form:"absolute" json:"absolute,omitempty"`
	// Required when type is 'relative'.
	Relative *PaymentPlanUpdateScheduleAmountsDueAmountDueDateRelativeParams `form:"relative" json:"relative,omitempty"`
	// Either 'absolute' or 'relative'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'fixed_amount'.
type PaymentPlanUpdateScheduleAmountsDueAmountFixedAmountParams struct {
	// The installment amount in minor units.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter ISO currency code.
	Currency *string `form:"currency" json:"currency"`
}

// The list of installment entries.
type PaymentPlanUpdateScheduleAmountsDueAmountParams struct {
	// Optional description for this installment.
	Description *string `form:"description" json:"description,omitempty"`
	// When this installment is due.
	DueDate *PaymentPlanUpdateScheduleAmountsDueAmountDueDateParams `form:"due_date" json:"due_date,omitempty"`
	// Required when type is 'fixed_amount'.
	FixedAmount *PaymentPlanUpdateScheduleAmountsDueAmountFixedAmountParams `form:"fixed_amount" json:"fixed_amount,omitempty"`
	// Optional stable identifier for the installment entry.
	ID *string `form:"id" json:"id,omitempty"`
	// The installment percentage of the total. Required when type is 'percentage'.
	Percentage *float64 `form:"percentage" json:"percentage,omitempty"`
	// Either 'fixed_amount' or 'percentage'.
	Type *string `form:"type" json:"type"`
}

// Required when type is 'amounts_due'.
type PaymentPlanUpdateScheduleAmountsDueParams struct {
	// The list of installment entries.
	Amounts []*PaymentPlanUpdateScheduleAmountsDueAmountParams `form:"amounts" json:"amounts"`
}

// The new schedule for this payment plan.
type PaymentPlanUpdateScheduleParams struct {
	// Required when type is 'amounts_due'.
	AmountsDue *PaymentPlanUpdateScheduleAmountsDueParams `form:"amounts_due" json:"amounts_due"`
	// The schedule type. Currently only 'amounts_due' is supported.
	Type *string `form:"type" json:"type"`
}

// Updates the schedule or metadata of an existing payment plan. Only unpaid installments can be updated.
type PaymentPlanUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The new schedule for this payment plan.
	Schedule *PaymentPlanUpdateScheduleParams `form:"schedule" json:"schedule,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentPlanUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentPlanUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type PaymentPlanCollectsOnInvoiceDetails struct {
	// The ID of the invoice this plan collects against.
	Invoice string `json:"invoice"`
}

// The list of objects this payment plan collects against.
type PaymentPlanCollectsOn struct {
	InvoiceDetails *PaymentPlanCollectsOnInvoiceDetails `json:"invoice_details"`
	// The type of object this plan collects against. Currently always `invoice_details`.
	Type string `json:"type"`
}

// The list of installments derived from the schedule. Each installment tracks an individual payment obligation.
type PaymentPlanInstallment struct {
	// Amount owed for this installment, in the smallest currency unit.
	AmountDue int64 `json:"amount_due"`
	// Amount forgiven for this installment, in the smallest currency unit.
	AmountForgiven int64 `json:"amount_forgiven"`
	// Amount already paid toward this installment, in the smallest currency unit.
	AmountPaid int64 `json:"amount_paid"`
	// Three-letter ISO currency code.
	Currency Currency `json:"currency"`
	// A description of this installment.
	Description string `json:"description"`
	// Unix timestamp when this installment is due. Omitted for installments with no due date.
	DueDate int64 `json:"due_date,omitempty"`
	// Unique identifier for the installment.
	ID string `json:"id,omitempty"`
	// Unix timestamp when this installment was paid.
	PaidAt int64 `json:"paid_at,omitempty"`
	// The status of this installment. One of `open`, `paid`, `past_due`, or `canceled`.
	Status string `json:"status"`
}
type PaymentPlanScheduleAmountsDueAmountDueDateRelative struct {
	// The number of intervals after the invoice is finalized that this entry is due.
	Count int64 `json:"count"`
	// The interval unit: `day`, `week`, `month`, or `year`.
	Interval string `json:"interval"`
}
type PaymentPlanScheduleAmountsDueAmountDueDate struct {
	// Unix timestamp of the due date. Present when type is `absolute`.
	Absolute int64                                               `json:"absolute,omitempty"`
	Relative *PaymentPlanScheduleAmountsDueAmountDueDateRelative `json:"relative,omitempty"`
	// The type of due date. Either `absolute` or `relative`.
	Type string `json:"type"`
}
type PaymentPlanScheduleAmountsDueAmountFixedAmount struct {
	// Fixed amount for this entry, in the smallest currency unit.
	Amount int64 `json:"amount"`
	// Three-letter ISO currency code.
	Currency Currency `json:"currency"`
}

// The list of installment schedule entries.
type PaymentPlanScheduleAmountsDueAmount struct {
	// A description of this schedule entry.
	Description string                                          `json:"description"`
	DueDate     *PaymentPlanScheduleAmountsDueAmountDueDate     `json:"due_date,omitempty"`
	FixedAmount *PaymentPlanScheduleAmountsDueAmountFixedAmount `json:"fixed_amount,omitempty"`
	// Unique identifier for this schedule entry.
	ID string `json:"id,omitempty"`
	// Percentage of the invoice total for this entry (0–100). Present when type is `percentage`.
	Percentage float64 `json:"percentage,omitempty"`
	// The type of this schedule entry. Either `fixed_amount` or `percentage`.
	Type string `json:"type"`
}
type PaymentPlanScheduleAmountsDue struct {
	// The list of installment schedule entries.
	Amounts []*PaymentPlanScheduleAmountsDueAmount `json:"amounts"`
}
type PaymentPlanSchedule struct {
	AmountsDue *PaymentPlanScheduleAmountsDue `json:"amounts_due"`
	// The type of schedule. Currently always `amounts_due`.
	Type string `json:"type"`
}

// A Payment Plan splits a single invoice obligation into multiple installments,
// each with its own due date and amount. Payment Plans are associated with a
// finalized or draft invoice and track how much has been collected against
// each installment.
type PaymentPlan struct {
	APIResource
	// The list of objects this payment plan collects against.
	CollectsOn []*PaymentPlanCollectsOn `json:"collects_on"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The list of installments derived from the schedule. Each installment tracks an individual payment obligation.
	Installments []*PaymentPlanInstallment `json:"installments"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object   string               `json:"object"`
	Schedule *PaymentPlanSchedule `json:"schedule"`
}

// PaymentPlanList is a list of PaymentPlans as retrieved from a list endpoint.
type PaymentPlanList struct {
	APIResource
	ListMeta
	Data []*PaymentPlan `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentPlan.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentPlan) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentPlan PaymentPlan
	var v paymentPlan
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentPlan(v)
	return nil
}
