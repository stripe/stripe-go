package stripe

// PlanInterval is the list of allowed values for a plan's interval.
// Allowed values are "day", "week", "month", "year".
type PlanInterval string

// Plan is the resource representing a Stripe plan.
// For more details see https://stripe.com/docs/api#plans.
type Plan struct {
	Amount        uint64            `json:"amount"`
	Created       int64             `json:"created"`
	Currency      Currency          `json:"currency"`
	Deleted       bool              `json:"deleted"`
	ID            string            `json:"id"`
	Interval      PlanInterval      `json:"interval"`
	IntervalCount uint64            `json:"interval_count"`
	Live          bool              `json:"livemode"`
	Meta          map[string]string `json:"metadata"`
	Name          string            `json:"name"`
	Statement     string            `json:"statement_descriptor"`
	TrialPeriod   uint64            `json:"trial_period_days"`
}

// PlanList is a list of plans as returned from a list endpoint.
type PlanList struct {
	ListMeta
	Values []*Plan `json:"data"`
}

// PlanListParams is the set of parameters that can be used when listing plans.
// For more details see https://stripe.com/docs/api#list_plans.
type PlanListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// PlanParams is the set of parameters that can be used when creating or updating a plan.
// For more details see https://stripe.com/docs/api#create_plan and https://stripe.com/docs/api#update_plan.
type PlanParams struct {
	Params        `form:"*"`
	Amount        uint64       `form:"amount"`
	Currency      Currency     `form:"currency"`
	ID            string       `form:"id"`
	Interval      PlanInterval `form:"interval"`
	IntervalCount uint64       `form:"interval_count"`
	Name          string       `form:"name"`
	Statement     string       `form:"statement_descriptor"`
	TrialPeriod   uint64       `form:"trial_period_days"`
}
