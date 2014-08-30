package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// InvoiceLineType is the list of allowed values for the invoice line's type.
// Allowed values are "invoiceitem", "subscription".
type InvoiceLineType string

const (
	TypeInvoiceItem  InvoiceLineType = "invoiceitem"
	TypeSubscription InvoiceLineType = "subscription"
)

// InvoiceParams is the set of parameters that can be used when creating or updating an invoice.
// For more details see https://stripe.com/docs/api#create_invoice, https://stripe.com/docs/api#update_invoice.
type InvoiceParams struct {
	Params
	Customer             string
	Desc, Statement, Sub string
	Fee                  uint64
	Closed, Forgive      bool
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams
	Date     int64
	Customer string
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams
	Id            string
	Customer, Sub string
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	Id           string            `json:"id"`
	Live         bool              `json:"livemode"`
	Amount       int64             `json:"amount_due"`
	Attempts     uint64            `json:"attempt_count"`
	Attempted    bool              `json:"attempted"`
	Closed       bool              `json:"closed"`
	Currency     Currency          `json:"currency"`
	Customer     *Customer         `json:"customer"`
	Date         int64             `json:"date"`
	Forgive      bool              `json:"forgiven"`
	Lines        *InvoiceLineList  `json:"lines"`
	Paid         bool              `json:"paid"`
	End          int64             `json:"period_end"`
	Start        int64             `json:"period_start"`
	StartBalance int64             `json:"starting_balance"`
	Subtotal     int64             `json:"subtotal"`
	Total        int64             `json:"total"`
	Fee          uint64            `json:"application_fee"`
	Charge       *Charge           `json:"charge"`
	Desc         string            `json:"description"`
	Discount     *Discount         `json:"discount"`
	EndBalance   int64             `json:"ending_balance"`
	NextAttempt  int64             `json:"next_payment_attempt"`
	Statement    string            `json:"statement_description"`
	Sub          string            `json:"subscription"`
	Webhook      int64             `json:"webhooks_delivered_at"`
	Meta         map[string]string `json:"metadata"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	Id        string            `json:"id"`
	Live      bool              `json:"live_mode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Period    *Period           `json:"period"`
	Proration bool              `json:"proration"`
	Type      InvoiceLineType   `json:"type"`
	Desc      string            `json:"description"`
	Meta      map[string]string `json:"metadata"`
	Plan      *Plan             `json:"plan"`
	Quantity  int64             `json:"quantity"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// InvoiceList is a list object for invoices.
type InvoiceList struct {
	ListResponse
	Values []*Invoice `json:"data"`
}

// InvoiceLineList is a list object for invoice line items.
type InvoiceLineList struct {
	ListResponse
	Values []*InvoiceLine `json:"data"`
}

// InvoiceClient is the client used to invoke /invoices APIs.
type InvoiceClient struct {
	api   Api
	token string
}

// Create POSTs new invoices.
// For more details see https://stripe.com/docs/api#create_invoice.
func (c *InvoiceClient) Create(params *InvoiceParams) (*Invoice, error) {
	body := &url.Values{
		"customer": {params.Customer},
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	token := c.token
	if params.Fee > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid invoice params: an access token is required for application fees")
			return nil, err
		}

		body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
		token = params.AccessToken
	}

	invoice := &Invoice{}
	err := c.api.Call("POST", "/invoices", token, body, invoice)

	return invoice, err
}

// Get returns the details of an invoice.
// For more details see https://stripe.com/docs/api#retrieve_invoice.
func (c *InvoiceClient) Get(id string) (*Invoice, error) {
	invoice := &Invoice{}
	err := c.api.Call("GET", "/invoices/"+id, c.token, nil, invoice)

	return invoice, err
}

// Pay pays an invoice.
// For more details see https://stripe.com/docs/api#pay_invoice.
func (c *InvoiceClient) Pay(id string) (*Invoice, error) {
	invoice := &Invoice{}
	err := c.api.Call("POST", fmt.Sprintf("/invoices/%v/pay", id), c.token, nil, invoice)

	return invoice, err
}

// Update updates an invoice's properties.
// For more details see https://stripe.com/docs/api#update_invoice.
func (c *InvoiceClient) Update(id string, params *InvoiceParams) (*Invoice, error) {
	var body *url.Values
	token := c.token

	if params != nil {
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		if len(params.Statement) > 0 {
			body.Add("statement_description", params.Statement)
		}

		if len(params.Sub) > 0 {
			body.Add("subscription", params.Sub)
		}

		if params.Closed {
			body.Add("closed", strconv.FormatBool(true))
		}

		if params.Forgive {
			body.Add("forgiven", strconv.FormatBool(true))
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}

		if params.Fee > 0 {
			if len(params.AccessToken) == 0 {
				err := errors.New("Invalid invoice params: an access token is required for application fees")
				return nil, err
			}

			body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
			token = params.AccessToken
		}
	}

	invoice := &Invoice{}
	err := c.api.Call("POST", "/invoices/"+id, token, body, invoice)

	return invoice, err
}

// GetNext returns the upcoming invoice's properties.
// For more details see https://stripe.com/docs/api#retrieve_customer_invoice.
func (c *InvoiceClient) GetNext(params *InvoiceParams) (*Invoice, error) {
	body := &url.Values{
		"customer": {params.Customer},
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	invoice := &Invoice{}
	err := c.api.Call("GET", "/invoices", c.token, body, invoice)

	return invoice, err
}

// List returns a list of invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
func (c *InvoiceClient) List(params *InvoiceListParams) (*InvoiceList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		if params.Date > 0 {
			body.Add("date", strconv.FormatInt(params.Date, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Start) > 0 {
			body.Add("starting_after", params.Start)
		}

		if len(params.End) > 0 {
			body.Add("ending_before", params.End)
		}

		if params.Limit > 0 {
			if params.Limit > 100 {
				params.Limit = 100
			}

			body.Add("limit", strconv.FormatUint(params.Limit, 10))
		}
	}

	list := &InvoiceList{}
	err := c.api.Call("GET", "/invoices", c.token, body, list)

	return list, err
}

// ListLines returns a list of line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
func (c *InvoiceClient) ListLines(params *InvoiceLineListParams) (*InvoiceLineList, error) {
	body := &url.Values{}

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	if len(params.Filters.f) > 0 {
		params.Filters.appendTo(body)
	}

	if len(params.Start) > 0 {
		body.Add("starting_after", params.Start)
	}

	if len(params.End) > 0 {
		body.Add("ending_before", params.End)
	}

	if params.Limit > 0 {
		if params.Limit > 100 {
			params.Limit = 100
		}

		body.Add("limit", strconv.FormatUint(params.Limit, 10))
	}

	list := &InvoiceLineList{}
	err := c.api.Call("GET", fmt.Sprintf("/invoices/%v/lines", params.Id), c.token, body, list)

	return list, err
}

func (i *Invoice) UnmarshalJSON(data []byte) error {
	type invoice Invoice
	var ii invoice
	err := json.Unmarshal(data, &ii)
	if err == nil {
		*i = Invoice(ii)
	} else {
		// the id is surrounded by escaped \, so ignore those
		i.Id = string(data[1 : len(data)-1])
	}

	return nil
}
