package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// SubscriptionStatus is the list of allowed values for the subscription's status.
// Allowed values are "trialing", "active", "past_due", "canceled", "unpaid".
type SubscriptionStatus string

const (
	Trialing SubscriptionStatus = "trialing"
	Active   SubscriptionStatus = "active"
	PastDue  SubscriptionStatus = "past_due"
	Canceled SubscriptionStatus = "canceled"
	Unpaid   SubscriptionStatus = "unpaid"
)

// SubParams is the set of parameters that can be used when creating or updating a subscription.
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubParams struct {
	Params
	Customer, Plan       string
	Coupon, Token        string
	TrialEnd             int64
	Card                 *CardParams
	Quantity             uint64
	FeePercent           float64
	NoProrate, EndCancel bool
}

// SubListParams is the set of parameters that can be used when listing active subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
type SubListParams struct {
	ListParams
	Customer string
}

// Subscription is the resource representing a Stripe subscription.
// For more details see https://stripe.com/docs/api#subscriptions.
type Subscription struct {
	Id          string             `json:"id"`
	EndCancel   bool               `json:"cancel_at_period_end"`
	Customer    *Customer          `json:"customer"`
	Plan        *Plan              `json:"plan"`
	Quantity    uint64             `json:"quantity"`
	Status      SubscriptionStatus `json:"status"`
	FeePercent  float64            `json:"application_fee_percent"`
	Canceled    int64              `json:"canceled_at"`
	PeriodEnd   int64              `json:"current_period_end"`
	PeriodStart int64              `json:"current_period_start"`
	Discount    *Discount          `json:"discount"`
	Ended       int64              `json:"ended_at"`
	Meta        map[string]string  `json:"metadata"`
	TrialEnd    int64              `json:"trial_end"`
	TrialStart  int64              `json:"trial_start"`
}

// SubscriptionList is a list object for subscriptions.
type SubscriptionList struct {
	ListResponse
	Values []*Subscription `json:"data"`
}

// SubscriptionClient is the client used to invoke /subscriptions APIs.
type SubscriptionClient struct {
	api   Api
	token string
}

// Create POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription.
func (c *SubscriptionClient) Create(params *SubParams) (*Subscription, error) {
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

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	token := c.token
	if params.FeePercent > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid sub params: an access token is required for application fees")
			return nil, err
		}

		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
		token = params.AccessToken
	}

	sub := &Subscription{}
	err := c.api.Call("POST", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), token, body, sub)

	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func (c *SubscriptionClient) Get(id string, params *SubParams) (*Subscription, error) {
	sub := &Subscription{}
	err := c.api.Call("GET", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.token, nil, sub)

	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func (c *SubscriptionClient) Update(id string, params *SubParams) (*Subscription, error) {
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
		if len(params.Card.Token) > 0 {
			body.Add("card", params.Card.Token)
		} else {
			params.Card.appendTo(body, true)
		}
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

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	token := c.token
	if params.FeePercent > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid sub params: an access token is required for application fees")
			return nil, err
		}

		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
		token = params.AccessToken
	}

	sub := &Subscription{}
	err := c.api.Call("POST", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), token, body, sub)

	return sub, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func (c *SubscriptionClient) Cancel(id string, params *SubParams) error {
	body := &url.Values{}

	if params.EndCancel {
		body.Add("at_period_end", strconv.FormatBool(true))
	}

	return c.api.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.token, body, nil)
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func (c *SubscriptionClient) List(params *SubListParams) (*SubscriptionList, error) {
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

	list := &SubscriptionList{}
	err := c.api.Call("GET", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), c.token, body, list)

	return list, err
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
	type sub Subscription
	var ss sub
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = Subscription(ss)
	} else {
		// the id is surrounded by escaped \, so ignore those
		s.Id = string(data[1 : len(data)-1])
	}

	return nil
}
