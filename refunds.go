package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// RefundParams is the set of parameters that can be used when refunding a charge.
// For more details see https://stripe.com/docs/api#refund.
type RefundParams struct {
	Params
	Charge string
	Amount uint64
	Fee    bool
}

// RefundListParams is the set of parameters that can be used when listing refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
type RefundListParams struct {
	ListParams
	Charge string
}

// Refund is the resource representing a Stripe refund.
// For more details see https://stripe.com/docs/api#refunds.
type Refund struct {
	Id       string            `json:"id"`
	Amount   uint64            `json:"amount"`
	Created  int64             `json:"created"`
	Currency Currency          `json:"currency"`
	Tx       *Transaction      `json:"balance_transaction"`
	Charge   string            `json:"charge"`
	Meta     map[string]string `json:"metadata"`
}

// Refundist is a list object for refunds.
type RefundList struct {
	ListResponse
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
	body := &url.Values{}

	params.appendTo(body)

	refund := &Refund{}
	err := c.api.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.token, body, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func (c *RefundClient) Update(id string, params *RefundParams) (*Refund, error) {
	body := &url.Values{}

	params.appendTo(body)

	refund := &Refund{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/refunds/%v", params.Charge, id), c.token, body, refund)

	return refund, err
}

// List returns a list of refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
func (c *RefundClient) List(params *RefundListParams) (*RefundList, error) {
	body := &url.Values{}

	params.appendTo(body)

	list := &RefundList{}
	err := c.api.Call("GET", fmt.Sprintf("/charges/%v/refunds", params.Charge), c.token, body, list)

	return list, err
}

func (r *Refund) UnmarshalJSON(data []byte) error {
	type refund Refund
	var rr refund
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Refund(rr)
	} else {
		// the id is surrounded by escaped \, so ignore those
		r.Id = string(data[1 : len(data)-1])
	}

	return nil
}
