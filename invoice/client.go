// package invoice provides the /invoices APIs
package invoice

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is the client used to invoke /invoices APIs.
type Client struct {
	B   Backend
	Tok string
}

// Create POSTs new invoices.
// For more details see https://stripe.com/docs/api#create_invoice.
func Create(params *InvoiceParams) (*Invoice, error) {
	return getC().Create(params)
}

func (c Client) Create(params *InvoiceParams) (*Invoice, error) {
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

	params.AppendTo(body)

	token := c.Tok
	if params.Fee > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid invoice params: an access token is required for application fees")
			return nil, err
		}

		body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
		token = params.AccessToken
	}

	invoice := &Invoice{}
	err := c.B.Call("POST", "/invoices", token, body, invoice)

	return invoice, err
}

// Get returns the details of an invoice.
// For more details see https://stripe.com/docs/api#retrieve_invoice.
func Get(id string, params *InvoiceParams) (*Invoice, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *InvoiceParams) (*Invoice, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	invoice := &Invoice{}
	err := c.B.Call("GET", "/invoices/"+id, c.Tok, body, invoice)

	return invoice, err
}

// Pay pays an invoice.
// For more details see https://stripe.com/docs/api#pay_invoice.
func Pay(id string, params *InvoiceParams) (*Invoice, error) {
	return getC().Pay(id, params)
}

func (c Client) Pay(id string, params *InvoiceParams) (*Invoice, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	invoice := &Invoice{}
	err := c.B.Call("POST", fmt.Sprintf("/invoices/%v/pay", id), c.Tok, body, invoice)

	return invoice, err
}

// Update updates an invoice's properties.
// For more details see https://stripe.com/docs/api#update_invoice.
func Update(id string, params *InvoiceParams) (*Invoice, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *InvoiceParams) (*Invoice, error) {
	var body *url.Values
	token := c.Tok

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

		if params.Fee > 0 {
			if len(params.AccessToken) == 0 {
				err := errors.New("Invalid invoice params: an access token is required for application fees")
				return nil, err
			}

			body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
			token = params.AccessToken
		}

		params.AppendTo(body)
	}

	invoice := &Invoice{}
	err := c.B.Call("POST", "/invoices/"+id, token, body, invoice)

	return invoice, err
}

// GetNext returns the upcoming invoice's properties.
// For more details see https://stripe.com/docs/api#retrieve_customer_invoice.
func GetNext(params *InvoiceParams) (*Invoice, error) {
	return getC().GetNext(params)
}

func (c Client) GetNext(params *InvoiceParams) (*Invoice, error) {
	body := &url.Values{
		"customer": {params.Customer},
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	params.AppendTo(body)

	invoice := &Invoice{}
	err := c.B.Call("GET", "/invoices", c.Tok, body, invoice)

	return invoice, err
}

// List returns a list of invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
func List(params *InvoiceListParams) (*InvoiceList, error) {
	return getC().List(params)
}

func (c Client) List(params *InvoiceListParams) (*InvoiceList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		if params.Date > 0 {
			body.Add("date", strconv.FormatInt(params.Date, 10))
		}

		params.AppendTo(body)
	}

	list := &InvoiceList{}
	err := c.B.Call("GET", "/invoices", c.Tok, body, list)

	return list, err
}

// ListLines returns a list of line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
func ListLines(params *InvoiceLineListParams) (*InvoiceLineList, error) {
	return getC().ListLines(params)
}

func (c Client) ListLines(params *InvoiceLineListParams) (*InvoiceLineList, error) {
	body := &url.Values{}

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	params.AppendTo(body)

	list := &InvoiceLineList{}
	err := c.B.Call("GET", fmt.Sprintf("/invoices/%v/lines", params.Id), c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
