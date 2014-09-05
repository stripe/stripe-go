package plan

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /plans APIs.
type Client struct {
	B   Backend
	Tok string
}

// Create POSTs a new plan.
// For more details see https://stripe.com/docs/api#create_plan.
func Create(params *PlanParams) (*Plan, error) {
	return getC().Create(params)
}

func (c Client) Create(params *PlanParams) (*Plan, error) {
	body := &url.Values{
		"id":       {params.Id},
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
		body.Add("statement_description", params.Statement)
	}

	params.AppendTo(body)

	plan := &Plan{}
	err := c.B.Call("POST", "/plans", c.Tok, body, plan)

	return plan, err
}

// Get returns the details of a plan.
// For more details see https://stripe.com/docs/api#retrieve_plan.
func Get(id string, params *PlanParams) (*Plan, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *PlanParams) (*Plan, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	plan := &Plan{}
	err := c.B.Call("GET", "/plans/"+id, c.Tok, body, plan)

	return plan, err
}

// Update updates a plan's properties.
// For more details see https://stripe.com/docs/api#update_plan.
func Update(id string, params *PlanParams) (*Plan, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *PlanParams) (*Plan, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Name) > 0 {
			body.Add("name", params.Name)
		}

		if len(params.Statement) > 0 {
			body.Add("statement_description", params.Statement)
		}

		params.AppendTo(body)
	}

	plan := &Plan{}
	err := c.B.Call("POST", "/plans/"+id, c.Tok, body, plan)

	return plan, err
}

// Delete removes a plan.
// For more details see https://stripe.com/docs/api#delete_plan.
func Delete(id string) error {
	return getC().Delete(id)
}

func (c Client) Delete(id string) error {
	return c.B.Call("DELETE", "/plans/"+id, c.Tok, nil, nil)
}

// List returns a list of plans.
// For more details see https://stripe.com/docs/api#list_plans.
func List(params *PlanListParams) (*PlanList, error) {
	return getC().List(params)
}

func (c Client) List(params *PlanListParams) (*PlanList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		params.AppendTo(body)
	}

	list := &PlanList{}
	err := c.B.Call("GET", "/plans", c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
