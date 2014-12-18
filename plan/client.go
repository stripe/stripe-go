// Package plan provides the /plans APIs
package plan

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
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
	body := &url.Values{
		"id":       {params.ID},
		"name":     {params.Name},
		"amount":   {strconv.FormatUint(params.Amount, 10)},
		"currency": {string(params.Currency)},
		"interval": {string(params.Interval)},
	}

	if params.IntervalCount > 0 {
		body.Add("interval_count", strconv.FormatUint(params.IntervalCount, 10))
	}

	if params.TrialPeriod > 0 {
		body.Add("trial_period_days", strconv.FormatUint(params.TrialPeriod, 10))
	}

	if len(params.Statement) > 0 {
		body.Add("statement_descriptor", params.Statement)
	}
	params.AppendTo(body)

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
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}

	plan := &stripe.Plan{}
	err := c.B.Call("GET", "/plans/"+id, c.Key, body, commonParams, plan)

	return plan, err
}

// Update updates a plan's properties.
// For more details see https://stripe.com/docs/api#update_plan.
func Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if len(params.Name) > 0 {
			body.Add("name", params.Name)
		}

		if len(params.Statement) > 0 {
			body.Add("statement_descriptor", params.Statement)
		}

		params.AppendTo(body)
	}

	plan := &stripe.Plan{}
	err := c.B.Call("POST", "/plans/"+id, c.Key, body, commonParams, plan)

	return plan, err
}

// Del removes a plan.
// For more details see https://stripe.com/docs/api#delete_plan.
func Del(id string) error {
	return getC().Del(id)
}

func (c Client) Del(id string) error {
	return c.B.Call("DELETE", "/plans/"+id, c.Key, nil, nil, nil)
}

// List returns a list of plans.
// For more details see https://stripe.com/docs/api#list_plans.
func List(params *stripe.PlanListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.PlanListParams) *Iter {
	type planList struct {
		stripe.ListMeta
		Values []*stripe.Plan `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &planList{}
		err := c.B.Call("GET", "/plans", c.Key, &b, nil, list)

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
