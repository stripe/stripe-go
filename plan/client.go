// Package plan provides the /plans APIs
package plan

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	plan := &stripe.Plan{}
	err := c.B.Call("POST", "/plans", c.Key, params, plan)
	return plan, err
}

// Get returns the details of a plan.
// For more details see https://stripe.com/docs/api#retrieve_plan.
func Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	path := stripe.FormatURLPath("/plans/%s", id)
	plan := &stripe.Plan{}
	err := c.B.Call("GET", path, c.Key, params, plan)
	return plan, err
}

// Update updates a plan's properties.
// For more details see https://stripe.com/docs/api#update_plan.
func Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	path := stripe.FormatURLPath("/plans/%s", id)
	plan := &stripe.Plan{}
	err := c.B.Call("POST", path, c.Key, params, plan)
	return plan, err
}

// Del removes a plan.
// For more details see https://stripe.com/docs/api#delete_plan.
func Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	path := stripe.FormatURLPath("/plans/%s", id)
	plan := &stripe.Plan{}
	err := c.B.Call("DELETE", path, c.Key, params, plan)
	return plan, err
}

// List returns a list of plans.
// For more details see https://stripe.com/docs/api#list_plans.
func List(params *stripe.PlanListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.PlanListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.PlanList{}
		err := c.B.CallRaw("GET", "/plans", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
