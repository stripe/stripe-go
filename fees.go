package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

type AppFeeParams struct {
	Amount uint64
}

type AppFeeListParams struct {
	Created            int64
	Filters            Filters
	Charge, Start, End string
	Limit              uint64
}

type AppFee struct {
	Id             string    `json:"id"`
	Live           bool      `json:"livemode"`
	Account        string    `json:"account"`
	Amount         uint64    `json:"amount"`
	App            string    `json:"application"`
	Tx             string    `json:"balance_transaction"`
	Charge         string    `json:"charge"`
	Created        int64     `json:"created"`
	Currency       Currency  `json:"currency"`
	Refunded       bool      `json:"refunded"`
	Refunds        []*Refund `json:"refunds"`
	AmountRefunded uint64    `json:"amount_refunded"`
}

type AppFeeList struct {
	Count  uint16    `json:"total_count"`
	More   bool      `json:"has_more"`
	Url    string    `json:"url"`
	Values []*AppFee `json:"data"`
}

type AppFeeClient struct {
	api   Api
	token string
}

func (c *AppFeeClient) Get(id string) (*AppFee, error) {
	fee := &AppFee{}
	err := c.api.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.token, nil, fee)

	return fee, err
}

func (c *AppFeeClient) Refund(id string, params *AppFeeParams) (*AppFee, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}
	}

	fee := &AppFee{}
	err := c.api.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.token, body, fee)

	return fee, err
}

func (c *AppFeeClient) List(params *AppFeeListParams) (*AppFeeList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Charge) > 0 {
			body.Add("charge", params.Charge)
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

	list := &AppFeeList{}
	err := c.api.Call("GET", "/application_fees", c.token, body, list)

	return list, err
}
