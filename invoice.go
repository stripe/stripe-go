package stripe

import "encoding/json"

// InvoiceLineType is the list of allowed values for the invoice line's type.
// Allowed values are "invoiceitem", "subscription".
type InvoiceLineType string

// InvoiceBilling is the type of billing method for this invoice.
// Currently supported values are "send_invoice" and "charge_automatically".
type InvoiceBilling string

// InvoiceParams is the set of parameters that can be used when creating or updating an invoice.
// For more details see https://stripe.com/docs/api#create_invoice, https://stripe.com/docs/api#update_invoice.
type InvoiceParams struct {
	Params              `form:"*"`
	ApplicationFee      uint64         `form:"application_fee"`
	ApplicationFeeZero  bool           `form:"application_fee,zero"`
	Billing             InvoiceBilling `form:"billing"`
	Closed              bool           `form:"closed"`
	Customer            string         `form:"customer"`
	DaysUntilDue        uint64         `form:"days_until_due"`
	Description         string         `form:"description"`
	DueDate             int64          `form:"due_date"`
	Forgiven            bool           `form:"forgiven"`
	NoClosed            bool           `form:"closed,invert"`
	Paid                bool           `form:"paid"`
	StatementDescriptor string         `form:"statement_descriptor"`
	Subscription        string         `form:"subscription"`
	TaxPercent          float64        `form:"tax_percent"`
	TaxPercentZero      bool           `form:"tax_percent,zero"`

	// These are all for exclusive use by GetNext.

	SubscriptionItems         []*SubscriptionItemsParams `form:"subscription_items,indexed"`
	SubscriptionNoProrate     bool                       `form:"subscription_prorate,invert"`
	SubscriptionPlan          string                     `form:"subscription_plan"`
	SubscriptionProrationDate int64                      `form:"subscription_proration_date"`
	SubscriptionQuantity      uint64                     `form:"subscription_quantity"`
	SubscriptionQuantityZero  bool                       `form:"subscription_quantity,zero"`
	SubscriptionTrialEnd      int64                      `form:"subscription_trial_end"`
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams   `form:"*"`
	Billing      InvoiceBilling    `form:"billing"`
	Customer     string            `form:"customer"`
	Date         int64             `form:"date"`
	DateRange    *RangeQueryParams `form:"date"`
	DueDate      int64             `form:"due_date"`
	Subscription string            `form:"subscription"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*"`

	Customer string `form:"customer"`

	// ID is the invoice ID to list invoice lines for.
	ID string `form:"-"` // Goes in the URL

	Subscription string `form:"subscription"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	AmountDue           int64             `json:"amount_due"`
	ApplicationFee      uint64            `json:"application_fee"`
	AttemptCount        uint64            `json:"attempt_count"`
	Attempted           bool              `json:"attempted"`
	Billing             InvoiceBilling    `json:"billing"`
	Charge              *Charge           `json:"charge"`
	Closed              bool              `json:"closed"`
	Currency            Currency          `json:"currency"`
	Customer            *Customer         `json:"customer"`
	Date                int64             `json:"date"`
	Description         string            `json:"description"`
	Discount            *Discount         `json:"discount"`
	DueDate             int64             `json:"due_date"`
	EndingBalance       int64             `json:"ending_balance"`
	Forgiven            bool              `json:"forgiven"`
	ID                  string            `json:"id"`
	Lines               *InvoiceLineList  `json:"lines"`
	Livemode            bool              `json:"livemode"`
	Metadata            map[string]string `json:"metadata"`
	NextPaymentAttempt  int64             `json:"next_payment_attempt"`
	Number              string            `json:"number"`
	Paid                bool              `json:"paid"`
	PeriodEnd           int64             `json:"period_end"`
	PeriodStart         int64             `json:"period_start"`
	ReceiptNumber       string            `json:"receipt_number"`
	StartingBalance     int64             `json:"starting_balance"`
	StatementDescriptor string            `json:"statement_descriptor"`
	Subscription        string            `json:"subscription"`
	Subtotal            int64             `json:"subtotal"`
	Tax                 int64             `json:"tax"`
	TaxPercent          float64           `json:"tax_percent"`
	Total               int64             `json:"total"`
	WebhooksDeliveredAt int64             `json:"webhooks_delivered_at"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	ListMeta
	Data []*Invoice `json:"data"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	Amount       int64             `json:"amount"`
	Currency     Currency          `json:"currency"`
	Description  string            `json:"description"`
	Discountable bool              `json:"discountable"`
	ID           string            `json:"id"`
	Livemode     bool              `json:"live_mode"`
	Metadata     map[string]string `json:"metadata"`
	Period       *Period           `json:"period"`
	Plan         *Plan             `json:"plan"`
	Proration    bool              `json:"proration"`
	Quantity     int64             `json:"quantity"`
	Subscription string            `json:"subscription"`
	Type         InvoiceLineType   `json:"type"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

// InvoiceLineList is a list object for invoice line items.
type InvoiceLineList struct {
	ListMeta
	Data []*InvoiceLine `json:"data"`
}

// InvoicePayParams is the set of parameters that can be used when
// paying invoices. For more details, see:
// https://stripe.com/docs/api#pay_invoice.
type InvoicePayParams struct {
	Params `form:"*"`
	Source string `form:"source"`
}

// UnmarshalJSON handles deserialization of an Invoice.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Invoice) UnmarshalJSON(data []byte) error {
	type invoice Invoice
	var ii invoice
	err := json.Unmarshal(data, &ii)
	if err == nil {
		*i = Invoice(ii)
	} else {
		// the id is surrounded by "\" characters, so strip them
		i.ID = string(data[1 : len(data)-1])
	}

	return nil
}
