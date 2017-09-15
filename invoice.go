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
	Billing        InvoiceBilling `form:"billing"`
	Closed         bool           `form:"closed"`
	Customer       string         `form:"customer"`
	DaysUntilDue   uint64         `form:"days_until_due"`
	Desc           string         `form:"description"`
	DueDate        int64          `form:"due_date"`
	Fee            uint64         `form:"application_fee"`
	FeeZero        bool           `form:"application_fee,zero"`
	Forgive        bool           `form:"forgiven"`
	NoClosed       bool           `form:"closed,invert"`
	Paid           bool           `form:"paid"`
	Statement      string         `form:"statement_descriptor"`
	Sub            string         `form:"subscription"`
	TaxPercent     float64        `form:"tax_percent"`
	TaxPercentZero bool           `form:"tax_percent,zero"`

	// These are all for exclusive use by GetNext.

	SubItems         []*SubItemsParams `form:"subscription_items,indexed"`
	SubNoProrate     bool              `form:"subscription_prorate,invert"`
	SubPlan          string            `form:"subscription_plan"`
	SubProrationDate int64             `form:"subscription_proration_date"`
	SubQuantity      uint64            `form:"subscription_quantity"`
	SubQuantityZero  bool              `form:"subscription_quantity,zero"`
	SubTrialEnd      int64             `form:"subscription_trial_end"`
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams `form:"*"`
	Billing    InvoiceBilling    `form:"billing"`
	Customer   string            `form:"customer"`
	Date       int64             `form:"date"`
	DateRange  *RangeQueryParams `form:"date"`
	DueDate    int64             `form:"due_date"`
	Sub        string            `form:"subscription"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*"`

	Customer string `form:"customer"`

	// ID is the invoice ID to list invoice lines for.
	ID string `form:"-"` // Goes in the URL

	Sub string `form:"subscription"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	Amount        int64             `json:"amount_due"`
	Attempted     bool              `json:"attempted"`
	Attempts      uint64            `json:"attempt_count"`
	Billing       InvoiceBilling    `json:"billing"`
	Charge        *Charge           `json:"charge"`
	Closed        bool              `json:"closed"`
	Currency      Currency          `json:"currency"`
	Customer      *Customer         `json:"customer"`
	Date          int64             `json:"date"`
	Desc          string            `json:"description"`
	Discount      *Discount         `json:"discount"`
	DueDate       int64             `json:"due_date"`
	End           int64             `json:"period_end"`
	EndBalance    int64             `json:"ending_balance"`
	Fee           uint64            `json:"application_fee"`
	Forgive       bool              `json:"forgiven"`
	ID            string            `json:"id"`
	Lines         *InvoiceLineList  `json:"lines"`
	Live          bool              `json:"livemode"`
	Meta          map[string]string `json:"metadata"`
	NextAttempt   int64             `json:"next_payment_attempt"`
	Number        string            `json:"number"`
	Paid          bool              `json:"paid"`
	ReceiptNumber string            `json:"receipt_number"`
	Start         int64             `json:"period_start"`
	StartBalance  int64             `json:"starting_balance"`
	Statement     string            `json:"statement_descriptor"`
	Sub           string            `json:"subscription"`
	Subtotal      int64             `json:"subtotal"`
	Tax           int64             `json:"tax"`
	TaxPercent    float64           `json:"tax_percent"`
	Total         int64             `json:"total"`
	Webhook       int64             `json:"webhooks_delivered_at"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	ListMeta
	Values []*Invoice `json:"data"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	Amount       int64             `json:"amount"`
	Currency     Currency          `json:"currency"`
	Desc         string            `json:"description"`
	Discountable bool              `json:"discountable"`
	ID           string            `json:"id"`
	Live         bool              `json:"live_mode"`
	Meta         map[string]string `json:"metadata"`
	Period       *Period           `json:"period"`
	Plan         *Plan             `json:"plan"`
	Proration    bool              `json:"proration"`
	Quantity     int64             `json:"quantity"`
	Sub          string            `json:"subscription"`
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
