package sub

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Create POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription.
func Create(params *SubParams) (*Subscription, error) {
	refresh()
	return c.Create(params)
}

func (c *Client) Create(params *SubParams) (*Subscription, error) {
	body := &url.Values{
		"plan": {params.Plan},
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.AppendDetails(body, true)
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

	token := c.Token
	if params.FeePercent > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid sub params: an access Token is required for application fees")
			return nil, err
		}

		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
		token = params.AccessToken
	}

	params.AppendTo(body)

	sub := &Subscription{}
	err := c.B.Call("POST", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), token, body, sub)

	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *SubParams) (*Subscription, error) {
	refresh()
	return c.Get(id, params)
}

func (c *Client) Get(id string, params *SubParams) (*Subscription, error) {
	body := &url.Values{}

	params.AppendTo(body)

	sub := &Subscription{}
	err := c.B.Call("GET", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.Token, body, sub)

	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *SubParams) (*Subscription, error) {
	refresh()
	return c.Update(id, params)
}

func (c *Client) Update(id string, params *SubParams) (*Subscription, error) {
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
			params.Card.AppendDetails(body, true)
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

	token := c.Token
	if params.FeePercent > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid sub params: an access Token is required for application fees")
			return nil, err
		}

		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
		token = params.AccessToken
	}

	params.AppendTo(body)

	sub := &Subscription{}
	err := c.B.Call("POST", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), token, body, sub)

	return sub, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Cancel(id string, params *SubParams) error {
	refresh()
	return c.Cancel(id, params)
}

func (c *Client) Cancel(id string, params *SubParams) error {
	body := &url.Values{}

	if params.EndCancel {
		body.Add("at_period_end", strconv.FormatBool(true))
	}

	params.AppendTo(body)

	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.Token, body, nil)
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func List(params *SubListParams) (*SubscriptionList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *SubListParams) (*SubscriptionList, error) {
	body := &url.Values{}

	params.AppendTo(body)

	list := &SubscriptionList{}
	err := c.B.Call("GET", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
