package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

// RefundParams is the set of parameters that can be used when refunding a refund.
// For more details see https://stripe.com/docs/api#refund_refund.
type RefundParams struct {
	Charge string
	Amount uint64
	Fee    bool
	Meta   map[string]string
}

// RefundListParams is the set of parameters that can be used when listing refunds.
// For more details see https://stripe.com/docs/api#list_refunds..
type RefundListParams struct {
	Charge     string
	Filters    Filters
	Start, End string
	Limit      uint64
}

// Refund is the resource representing a Stripe refund.
// For more details see https://stripe.com/docs/api#refunds.
type Refund struct {
	Id       string            `json:"id"`
	Amount   uint64            `json:"amount"`
	Created  int64             `json:"created"`
	Currency Currency          `json:"currency"`
	Tx       string            `json:"balance_transaction"`
	Charge   string            `json:"charge"`
	Meta     map[string]string `json:"metadata"`
}

// Refundist is a list object for refunds.
type RefundList struct {
	Count  uint16    `json:"total_count"`
	More   bool      `json:"has_more"`
	Url    string    `json:"url"`
	Values []*Refund `json:"data"`
}

// RefundClient is the  client used to invoke /refunds APIs.
type RefundClient struct {
	api   Api
	token string
}

// Get returns the details of a refund.
// For more details see https://stripe.com/docs/api#retrieve_refund.
func (c *RefundClient) Get(id string, params *RefundParams) (*Refund, error) {
	refund := &Refund{}
	err := c.api.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.token, nil, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func (c *RefundClient) Update(id string, params *RefundParams) (*Refund, error) {
	body := &url.Values{}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	refund := &Refund{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.token, body, refund)

	return refund, err
}

// List returns a list of refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
func (c *RefundClient) List(params *RefundListParams) (*RefundList, error) {
	body := &url.Values{}
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

	list := &RefundList{}
	err := c.api.Call("GET", fmt.Sprintf("/charges/%v/refunds", params.Charge), c.token, body, list)

	return list, err
}
