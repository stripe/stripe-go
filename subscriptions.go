package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type SubscriptionStatus string

const (
	Trialing SubscriptionStatus = "trialing"
	Active   SubscriptionStatus = "active"
	PastDue  SubscriptionStatus = "past_due"
	Canceled SubscriptionStatus = "canceled"
	Unpaid   SubscriptionStatus = "unpaid"
)

type SubscriptionParams struct {
	Customer             string
	Plan                 string
	Coupon, Token        string
	TrialEnd             int64
	Card                 *CardParams
	Quantity             uint64
	Fee                  float64
	Meta                 map[string]string
	NoProrate, EndCancel bool
}

type Subscription struct {
	Id          string             `json:"id"`
	EndCancel   bool               `json"cancel_at_period_end"`
	Customer    string             `json:"customer"`
	Plan        *Plan              `json:"plan"`
	Quantity    uint64             `json:"quantity"`
	Status      SubscriptionStatus `json:"status"`
	Fee         float64            `json:"application_fee_percent"`
	Canceled    int64              `json:"canceled_at"`
	PeriodEnd   int64              `json:"current_period_end"`
	PeriodStart int64              `json:"current_period_start"`
	Discount    *Discount          `json:"discount"`
	Ended       int64              `json:"ended_at"`
	Meta        map[string]string  `json:"metadata"`
	TrialEnd    int64              `json:"trial_end"`
	TrialStart  int64              `json:"trial_start"`
}

type SubscriptionList struct {
	Count  uint16          `json:"total_count"`
	More   bool            `json:"has_more"`
	Url    string          `json:"url"`
	Values []*Subscription `json:"data"`
}

type SubscriptionClient struct {
	api   Api
	token string
}

func (c *SubscriptionClient) Create(params *SubscriptionParams) (*Subscription, error) {
	body := &url.Values{
		"plan": {params.Plan},
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.Coupon) > 0 {
		body.Add("coupon", params.Coupon)
	}

	if params.TrialEnd > 0 {
		body.Add("trial_end", strconv.FormatInt(params.TrialEnd, 10))
	}

	if params.Quantity > 0 {
		body.Add("quantity", strconv.FormatUint(params.Quantity, 10))
	}

	if params.Fee > 0 {
		body.Add("application_fee_percent", strconv.FormatFloat(params.Fee, 'f', 2, 64))
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), c.token, body)

	if err != nil {
		return nil, err
	}

	var sub Subscription
	err = json.Unmarshal(res, &sub)
	return &sub, err
}

func (c *SubscriptionClient) Get(id string, params *SubscriptionParams) (*Subscription, error) {
	res, err := c.api.Call("GET", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.token, nil)

	if err != nil {
		return nil, err
	}

	var sub Subscription
	err = json.Unmarshal(res, &sub)
	return &sub, err
}

func (c *SubscriptionClient) Update(id string, params *SubscriptionParams) (*Subscription, error) {
	body := &url.Values{}

	if len(params.Plan) > 0 {
		body.Add("plan", params.Plan)
	}

	if params.NoProrate {
		body.Add("prorate", strconv.FormatBool(false))
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.Coupon) > 0 {
		body.Add("coupon", params.Coupon)
	}

	if params.TrialEnd > 0 {
		body.Add("trial_end", strconv.FormatInt(params.TrialEnd, 10))
	}

	if params.Quantity > 0 {
		body.Add("quantity", strconv.FormatUint(params.Quantity, 10))
	}

	if params.Fee > 0 {
		body.Add("application_fee_percent", strconv.FormatFloat(params.Fee, 'f', 2, 64))
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.token, body)

	if err != nil {
		return nil, err
	}

	var sub Subscription
	err = json.Unmarshal(res, &sub)
	return &sub, err
}

func (c *SubscriptionClient) Cancel(id string, params *SubscriptionParams) error {
	body := &url.Values{}

	if params.EndCancel {
		body.Add("at_period_end", strconv.FormatBool(true))
	}

	_, err := c.api.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.token, body)

	return err
}
