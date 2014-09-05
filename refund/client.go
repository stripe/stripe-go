package refund

import (
	"fmt"
	"net/url"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /refunds APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Get returns the details of a refund.
// For more details see https://stripe.com/docs/api#retrieve_refund.
func Get(id string, params *RefundParams) (*Refund, error) {
	refresh()
	return c.Get(id, params)
}

func (c *Client) Get(id string, params *RefundParams) (*Refund, error) {
	body := &url.Values{}

	params.AppendTo(body)

	refund := &Refund{}
	err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.Token, body, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *RefundParams) (*Refund, error) {
	refresh()
	return c.Update(id, params)
}

func (c *Client) Update(id string, params *RefundParams) (*Refund, error) {
	body := &url.Values{}

	params.AppendTo(body)

	refund := &Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.Token, body, refund)

	return refund, err
}

// List returns a list of refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
func List(params *RefundListParams) (*RefundList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *RefundListParams) (*RefundList, error) {
	body := &url.Values{}

	params.AppendTo(body)

	list := &RefundList{}
	err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds", params.Charge), c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
