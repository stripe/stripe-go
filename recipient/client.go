package recipient

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /recipients APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Create POSTs a new recipient.
// For more details see https://stripe.com/docs/api#create_recipient.
func Create(params *RecipientParams) (*Recipient, error) {
	refresh()
	return c.Create(params)
}

func (c *Client) Create(params *RecipientParams) (*Recipient, error) {
	body := &url.Values{
		"name": {params.Name},
		"type": {string(params.Type)},
	}

	if params.Bank != nil {
		params.Bank.AppendDetails(body)
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.AppendDetails(body, true)
	}

	if len(params.TaxId) > 0 {
		body.Add("tax_id", params.TaxId)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	params.AppendTo(body)

	recipient := &Recipient{}
	err := c.B.Call("POST", "/recipients", c.Token, body, recipient)

	return recipient, err
}

// Get returns the details of a recipient.
// For more details see https://stripe.com/docs/api#retrieve_recipient.
func Get(id string, params *RecipientParams) (*Recipient, error) {
	refresh()
	return c.Get(id, params)
}

func (c *Client) Get(id string, params *RecipientParams) (*Recipient, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	recipient := &Recipient{}
	err := c.B.Call("GET", "/recipients/"+id, c.Token, body, recipient)

	return recipient, err
}

// Update updates a recipient's properties.
// For more details see https://stripe.com/docs/api#update_recipient.
func Update(id string, params *RecipientParams) (*Recipient, error) {
	refresh()
	return c.Update(id, params)
}

func (c *Client) Update(id string, params *RecipientParams) (*Recipient, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Name) > 0 {
			body.Add("name", params.Name)
		}

		if params.Bank != nil {
			params.Bank.AppendDetails(body)
		}

		if len(params.Token) > 0 {
			body.Add("card", params.Token)
		} else if params.Card != nil {
			params.Card.AppendDetails(body, true)
		}

		if len(params.TaxId) > 0 {
			body.Add("tax_id", params.TaxId)
		}

		if len(params.DefaultCard) > 0 {
			body.Add("default_card", params.DefaultCard)
		}

		if len(params.Email) > 0 {
			body.Add("email", params.Email)
		}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		params.AppendTo(body)
	}

	recipient := &Recipient{}
	err := c.B.Call("POST", "/recipients/"+id, c.Token, body, recipient)

	return recipient, err
}

// Delete removes a recipient.
// For more details see https://stripe.com/docs/api#delete_recipient.
func Delete(id string) error {
	refresh()
	return c.Delete(id)
}

func (c *Client) Delete(id string) error {
	return c.B.Call("DELETE", "/recipients/"+id, c.Token, nil, nil)
}

// List returns a list of recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
func List(params *RecipientListParams) (*RecipientList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *RecipientListParams) (*RecipientList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Verified {
			body.Add("verified", strconv.FormatBool(true))
		}

		params.AppendTo(body)
	}

	list := &RecipientList{}
	err := c.B.Call("GET", "/recipients", c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
