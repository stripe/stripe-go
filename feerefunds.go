package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// FeeRefundParams is the set of parameters that can be used when refunding a fee.
// For more details see https://stripe.com/docs/api#fee_refund.
type FeeRefundParams struct {
	Fee    string
	Amount uint64
	Meta   map[string]string
}

// FeeRefundListParams is the set of parameters that can be used when listing fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
type FeeRefundListParams struct {
	ListParams
	Fee string
}

// FeeRefund is the resource representing a Stripe fee refund.
// For more details see https://stripe.com/docs/api#fee_refunds.
type FeeRefund struct {
	Id       string            `json:"id"`
	Amount   uint64            `json:"amount"`
	Created  int64             `json:"created"`
	Currency Currency          `json:"currency"`
	Tx       *Transaction      `json:"balance_transaction"`
	Fee      string            `json:"fee"`
	Meta     map[string]string `json:"metadata"`
}

// FeeRefundist is a list object for fee refunds.
type FeeRefundList struct {
	ListResponse
	Values []*FeeRefund `json:"data"`
}

// FeeRefundClient is the  client used to invoke /application_fees/refunds APIs.
type FeeRefundClient struct {
	api   Api
	token string
}

// Get returns the details of a fee refund.
// For more details see https://stripe.com/docs/api#retrieve_fee_refund.
func (c *FeeRefundClient) Get(id string, params *FeeRefundParams) (*FeeRefund, error) {
	refund := &FeeRefund{}
	err := c.api.Call("GET", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.token, nil, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func (c *FeeRefundClient) Update(id string, params *FeeRefundParams) (*FeeRefund, error) {
	body := &url.Values{}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	refund := &FeeRefund{}
	err := c.api.Call("POST", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.token, body, refund)

	return refund, err
}

// List returns a list of fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
func (c *FeeRefundClient) List(params *FeeRefundListParams) (*FeeRefundList, error) {
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

	list := &FeeRefundList{}
	err := c.api.Call("GET", fmt.Sprintf("/application_fees/%v/refunds", params.Fee), c.token, body, list)

	return list, err
}

func (f *FeeRefund) UnmarshalJSON(data []byte) error {
	type feerefund FeeRefund
	var ff feerefund
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = FeeRefund(ff)
	} else {
		// the id is surrounded by escaped \, so ignore those
		f.Id = string(data[1 : len(data)-1])
	}

	return nil
}
