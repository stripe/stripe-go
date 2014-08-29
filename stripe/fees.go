package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// AppFeeParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#refund_application_fee.
type AppFeeParams struct {
	Params
	Amount uint64
}

// AppFeeListParams is the set of parameters that can be used when listing application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
type AppFeeListParams struct {
	ListParams
	Created int64
	Charge  string
}

// AppFee is the resource representing a Stripe application fee.
// For more details see https://stripe.com/docs/api#application_fees.
type AppFee struct {
	Id             string         `json:"id"`
	Live           bool           `json:"livemode"`
	Account        *Account       `json:"account"`
	Amount         uint64         `json:"amount"`
	App            string         `json:"application"`
	Tx             *Transaction   `json:"balance_transaction"`
	Charge         *Charge        `json:"charge"`
	Created        int64          `json:"created"`
	Currency       Currency       `json:"currency"`
	Refunded       bool           `json:"refunded"`
	Refunds        *FeeRefundList `json:"refunds"`
	AmountRefunded uint64         `json:"amount_refunded"`
}

// AppFeeList is a list object for application fees.
type AppFeeList struct {
	ListResponse
	Values []*AppFee `json:"data"`
}

// AppFeeClient is the client used to invoke application_fees APIs.
type AppFeeClient struct {
	api   Api
	token string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func (c *AppFeeClient) Get(id string) (*AppFee, error) {
	fee := &AppFee{}
	err := c.api.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.token, nil, fee)

	return fee, err
}

// Refund refunds the application fee collected.
// For more details see https://stripe.com/docs/api#refund_application_fee.
func (c *AppFeeClient) Refund(id string, params *AppFeeParams) (*FeeRefund, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}
	}

	refund := &FeeRefund{}
	err := c.api.Call("POST", fmt.Sprintf("application_fees/%v/refunds", id), c.token, body, refund)

	return refund, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
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

func (f *AppFee) UnmarshalJSON(data []byte) error {
	type appfee AppFee
	var ff appfee
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = AppFee(ff)
	} else {
		// the id is surrounded by escaped \, so ignore those
		f.Id = string(data[1 : len(data)-1])
	}

	return nil
}
