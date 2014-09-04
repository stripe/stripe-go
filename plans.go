package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

// PlanInterval is the list of allowed values for a plan's interval.
// Allowed values are "day", "week", "month", "year".
type PlanInternval string

const (
	Day   PlanInternval = "day"
	Week  PlanInternval = "week"
	Month PlanInternval = "month"
	Year  PlanInternval = "year"
)

// PlanParams is the set of parameters that can be used when creating or updating a plan.
// For more details see https://stripe.com/docs/api#create_plan and https://stripe.com/docs/api#update_plan.
type PlanParams struct {
	Params
	Id, Name                   string
	Currency                   Currency
	Amount                     uint64
	Interval                   PlanInternval
	IntervalCount, TrialPeriod uint64
	Statement                  string
}

// PlanListParams is the set of parameters that can be used when listing plans.
// For more details see https://stripe.com/docs/api#list_plans.
type PlanListParams struct {
	ListParams
}

// Plan is the resource representing a Stripe plan.
// For more details see https://stripe.com/docs/api#plans.
type Plan struct {
	Id            string            `json:"id"`
	Live          bool              `json:"livemode"`
	Amount        uint64            `json:"amount"`
	Created       int64             `json:"created"`
	Currency      Currency          `json:"currency"`
	Interval      PlanInternval     `json:"interval"`
	IntervalCount uint64            `json:"interval_count"`
	Name          string            `json:"name"`
	Meta          map[string]string `json:"metadata"`
	TrialPeriod   uint64            `json:"trial_period_days"`
	Statement     string            `json:"statement_description"`
}

// PlanList is a list object for plans.
type PlanList struct {
	ListResponse
	Values []*Plan `json:"data"`
}

// PlanClient is the client used to invoke /plans APIs.
type PlanClient struct {
	api   Api
	token string
}

// Create POSTs a new plan.
// For more details see https://stripe.com/docs/api#create_plan.
func (c *PlanClient) Create(params *PlanParams) (*Plan, error) {
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

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	plan := &Plan{}
	err := c.api.Call("POST", "/plans", c.token, body, plan)

	return plan, err
}

// Get returns the details of a plan.
// For more details see https://stripe.com/docs/api#retrieve_plan.
func (c *PlanClient) Get(id string) (*Plan, error) {
	plan := &Plan{}
	err := c.api.Call("GET", "/plans/"+id, c.token, nil, plan)

	return plan, err
}

// Update updates a plan's properties.
// For more details see https://stripe.com/docs/api#update_plan.
func (c *PlanClient) Update(id string, params *PlanParams) (*Plan, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Name) > 0 {
			body.Add("name", params.Name)
		}

		if len(params.Statement) > 0 {
			body.Add("statement_description", params.Statement)
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}
	}

	plan := &Plan{}
	err := c.api.Call("POST", "/plans/"+id, c.token, body, plan)

	return plan, err
}

// Delete removes a plan.
// For more details see https://stripe.com/docs/api#delete_plan.
func (c *PlanClient) Delete(id string) error {
	return c.api.Call("DELETE", "/plans/"+id, c.token, nil, nil)
}

// List returns a list of plans.
// For more details see https://stripe.com/docs/api#list_plans.
func (c *PlanClient) List(params *PlanListParams) (*PlanList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

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
	}

	list := &PlanList{}
	err := c.api.Call("GET", "/plans", c.token, body, list)

	return list, err
}
