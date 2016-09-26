// Package sub provides the /subscriptions APIs
package subitem

import (
	"fmt"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription_item.
func New(params *stripe.SubItemParams) (*stripe.SubItem, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.SubItemParams) (*stripe.SubItem, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params
	token := c.Key

	if params != nil {
		body = &stripe.RequestValues{}
		body.Add("subscription", params.Sub)

		if len(params.Plan) > 0 {
			body.Add("plan", params.Plan)
		}
		if params.Quantity > 0 {
			body.Add("quantity", strconv.FormatUint(params.Quantity, 10))
		} else if params.QuantityZero {
			body.Add("quantity", "0")
		}

		commonParams = &params.Params
		params.AppendTo(body)
	}

	item := &stripe.SubItem{}
	err := c.B.Call("POST", "/subscription_items", token, body, commonParams, item)
	return item, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		params.AppendTo(body)
		commonParams = &params.Params
	}

	item := &stripe.SubItem{}
	err := c.B.Call("GET", fmt.Sprintf("/subscription_items/%v", id), c.Key, body, commonParams, item)

	return item, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params
	token := c.Key

	if params != nil {
		body = &stripe.RequestValues{}

		if len(params.Plan) > 0 {
			body.Add("plan", params.Plan)
		}

		if params.NoProrate {
			body.Add("prorate", strconv.FormatBool(false))
		}

		if params.Quantity > 0 {
			body.Add("quantity", strconv.FormatUint(params.Quantity, 10))
		} else if params.QuantityZero {
			body.Add("quantity", "0")
		}

		if params.ProrationDate > 0 {
			body.Add("proration_date", strconv.FormatInt(params.ProrationDate, 10))
		}

		commonParams = &params.Params
		params.AppendTo(body)
	}

	subi := &stripe.SubItem{}
	err := c.B.Call("POST", fmt.Sprintf("/subscription_items/%v", id), token, body, commonParams, subi)

	return subi, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Del(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.SubItemParams) (*stripe.SubItem, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		params.AppendTo(body)
		commonParams = &params.Params
	}

	item := &stripe.SubItem{}
	err := c.B.Call("DELETE", fmt.Sprintf("/subscription_items/%v", id), c.Key, body, commonParams, item)

	return item, err
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func List(params *stripe.SubListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.SubListParams) *Iter {
	var body *stripe.RequestValues
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		if len(params.Plan) > 0 {
			body.Add("plan", params.Plan)
		}

		params.AppendTo(body)

		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *stripe.RequestValues) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SubList{}
		err := c.B.Call("GET", "/subscription_items", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Subs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Sub returns the most recent Sub
// visited by a call to Next.
func (i *Iter) SubItem() *stripe.SubItem {
	return i.Current().(*stripe.SubItem)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
