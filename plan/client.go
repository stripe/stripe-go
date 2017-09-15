// Package plan provides the /plans APIs
package plan

import (
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Day   stripe.PlanInterval = "day"
	Week  stripe.PlanInterval = "week"
	Month stripe.PlanInterval = "month"
	Year  stripe.PlanInterval = "year"
)

// Client is used to invoke /plans APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new plan.
// For more details see https://stripe.com/docs/api#create_plan.
func New(params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.PlanParams) (*stripe.Plan, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	plan := &stripe.Plan{}
	err := c.B.Call("POST", "/plans", c.Key, body, &params.Params, plan)

	return plan, err
}

// Get returns the details of a plan.
// For more details see https://stripe.com/docs/api#retrieve_plan.
func Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	plan := &stripe.Plan{}
	err := c.B.Call("GET", "/plans/"+url.QueryEscape(id), c.Key, body, commonParams, plan)

	return plan, err
}

// Update updates a plan's properties.
// For more details see https://stripe.com/docs/api#update_plan.
func Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	plan := &stripe.Plan{}
	err := c.B.Call("POST", "/plans/"+url.QueryEscape(id), c.Key, body, commonParams, plan)

	return plan, err
}

// Del removes a plan.
// For more details see https://stripe.com/docs/api#delete_plan.
func Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	plan := &stripe.Plan{}
	qid := url.QueryEscape(id) //Added query escape per commit 9821176
	err := c.B.Call("DELETE", fmt.Sprintf("/plans/%v", qid), c.Key, body, commonParams, plan)
	return plan, err
}

// List returns a list of plans.
// For more details see https://stripe.com/docs/api#list_plans.
func List(params *stripe.PlanListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.PlanListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.PlanList{}
		err := c.B.Call("GET", "/plans", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Plans.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Plan returns the most recent Plan
// visited by a call to Next.
func (i *Iter) Plan() *stripe.Plan {
	return i.Current().(*stripe.Plan)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
