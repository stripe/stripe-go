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
	Params         `form:"*"`
	Customer       string         `form:"customer"`
	Desc           string         `form:"description"`
	Statement      string         `form:"statement_descriptor"`
	Sub            string         `form:"subscription"`
	Fee            uint64         `form:"application_fee"`
	FeeZero        bool           `form:"application_fee,zero"`
	Closed         bool           `form:"closed"`
	NoClosed       bool           `form:"closed,invert"`
	Forgive        bool           `form:"forgiven"`
	TaxPercent     float64        `form:"tax_percent"`
	TaxPercentZero bool           `form:"tax_percent,zero"`
	Billing        InvoiceBilling `form:"billing"`
	DueDate        int64          `form:"due_date"`
	DaysUntilDue   uint64         `form:"days_until_due"`
	Paid           bool           `form:"paid"`

	// These are all for exclusive use by GetNext.

	SubPlan          string            `form:"subscription_plan"`
	SubNoProrate     bool              `form:"subscription_prorate,invert"`
	SubProrationDate int64             `form:"subscription_proration_date"`
	SubQuantity      uint64            `form:"subscription_quantity"`
	SubQuantityZero  bool              `form:"subscription_quantity,zero"`
	SubTrialEnd      int64             `form:"subscription_trial_end"`
	SubItems         []*SubItemsParams `form:"subscription_items,indexed"`
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams `form:"*"`
	Date       int64             `form:"date"`
	DateRange  *RangeQueryParams `form:"date"`
	Customer   string            `form:"customer"`
	Sub        string            `form:"subscription"`
	Billing    InvoiceBilling    `form:"billing"`
	DueDate    int64             `form:"due_date"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*"`

	// ID is the invoice ID to list invoice lines for.
	ID string `form:"-"` // Goes in the URL

	Customer string `form:"customer"`
	Sub      string `form:"subscription"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	ID            string            `json:"id"`
	Live          bool              `json:"livemode"`
	Amount        int64             `json:"amount_due"`
	Attempts      uint64            `json:"attempt_count"`
	Attempted     bool              `json:"attempted"`
	Closed        bool              `json:"closed"`
	Currency      Currency          `json:"currency"`
	Customer      *Customer         `json:"customer"`
	Date          int64             `json:"date"`
	Forgive       bool              `json:"forgiven"`
	Lines         *InvoiceLineList  `json:"lines"`
	Paid          bool              `json:"paid"`
	End           int64             `json:"period_end"`
	Start         int64             `json:"period_start"`
	StartBalance  int64             `json:"starting_balance"`
	Subtotal      int64             `json:"subtotal"`
	Total         int64             `json:"total"`
	Tax           int64             `json:"tax"`
	TaxPercent    float64           `json:"tax_percent"`
	Fee           uint64            `json:"application_fee"`
	Charge        *Charge           `json:"charge"`
	Desc          string            `json:"description"`
	Discount      *Discount         `json:"discount"`
	ReceiptNumber string            `json:"receipt_number"`
	EndBalance    int64             `json:"ending_balance"`
	NextAttempt   int64             `json:"next_payment_attempt"`
	Statement     string            `json:"statement_descriptor"`
	Sub           string            `json:"subscription"`
	Webhook       int64             `json:"webhooks_delivered_at"`
	Meta          map[string]string `json:"metadata"`
	Billing       InvoiceBilling    `json:"billing"`
	DueDate       int64             `json:"due_date"`
	Number        string            `json:"number"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	ListMeta
	Values []*Invoice `json:"data"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	ID           string            `json:"id"`
	Live         bool              `json:"live_mode"`
	Amount       int64             `json:"amount"`
	Currency     Currency          `json:"currency"`
	Period       *Period           `json:"period"`
	Proration    bool              `json:"proration"`
	Type         InvoiceLineType   `json:"type"`
	Desc         string            `json:"description"`
	Meta         map[string]string `json:"metadata"`
	Plan         *Plan             `json:"plan"`
	Quantity     int64             `json:"quantity"`
	Sub          string            `json:"subscription"`
	Discountable bool              `json:"discountable"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// InvoiceLineList is a list object for invoice line items.
type InvoiceLineList struct {
	ListMeta
	Values []*InvoiceLine `json:"data"`
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
