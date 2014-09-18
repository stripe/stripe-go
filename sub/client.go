// package sub provides the /subscriptions APIs
package sub

import (
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   Backend
	Key string
}

// New POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription.
func New(params *SubParams) (*Sub, error) {
	return getC().New(params)
}

func (c Client) New(params *SubParams) (*Sub, error) {
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

	token := c.Key
	if params.FeePercent > 0 {
		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
	}

	params.AppendTo(body)

	sub := &Sub{}
	err := c.B.Call("POST", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), token, body, sub)

	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *SubParams) (*Sub, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *SubParams) (*Sub, error) {
	body := &url.Values{}

	params.AppendTo(body)

	sub := &Sub{}
	err := c.B.Call("GET", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.Key, body, sub)

	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *SubParams) (*Sub, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *SubParams) (*Sub, error) {
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

	token := c.Key
	if params.FeePercent > 0 {
		body.Add("application_fee_percent", strconv.FormatFloat(params.FeePercent, 'f', 2, 64))
	}

	params.AppendTo(body)

	sub := &Sub{}
	err := c.B.Call("POST", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), token, body, sub)

	return sub, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Cancel(id string, params *SubParams) error {
	return getC().Cancel(id, params)
}

func (c Client) Cancel(id string, params *SubParams) error {
	body := &url.Values{}

	if params.EndCancel {
		body.Add("at_period_end", strconv.FormatBool(true))
	}

	params.AppendTo(body)

	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v", params.Customer, id), c.Key, body, nil)
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func List(params *SubListParams) *SubIter {
	return getC().List(params)
}

func (c Client) List(params *SubListParams) *SubIter {
	body := &url.Values{}
	var lp *ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &SubIter{GetIter(lp, body, func(b url.Values) ([]interface{}, ListMeta, error) {
		list := &SubList{}
		err := c.B.Call("GET", fmt.Sprintf("/customers/%v/subscriptions", params.Customer), c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func getC() Client {
	return Client{GetBackend(), Key}
}
