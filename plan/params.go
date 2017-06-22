package plan

import (
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// PlanParams is the set of parameters that can be used when creating or updating a plan.
// For more details see https://stripe.com/docs/api#create_plan and https://stripe.com/docs/api#update_plan.
type PlanParams struct {
	stripe.Params
	ID, Name                   string
	Currency                   stripe.Currency
	Amount                     uint64
	Interval                   stripe.PlanInterval
	IntervalCount, TrialPeriod uint64
	Statement                  string
}

func (p *PlanParams) AppendTo(values *stripe.RequestValues) {
	if p.Amount > 0 {
		values.Add("amount", strconv.FormatUint(p.Amount, 10))
	}

	if p.Currency != "" {
		values.Add("currency", string(p.Currency))
	}

	if p.ID != "" {
		values.Add("id", p.ID)
	}

	if p.Interval != "" {
		values.Add("interval", string(p.Interval))
	}

	if p.IntervalCount > 0 {
		values.Add("interval_count", strconv.FormatUint(p.IntervalCount, 10))
	}

	if p.Name != "" {
		values.Add("name", p.Name)
	}

	if len(p.Statement) > 0 {
		values.Add("statement_descriptor", p.Statement)
	}

	if p.TrialPeriod > 0 {
		values.Add("trial_period_days", strconv.FormatUint(p.TrialPeriod, 10))
	}
}

// PlanListParams is the set of parameters that can be used when listing plans.
// For more details see https://stripe.com/docs/api#list_plans.
type PlanListParams struct {
	stripe.ListParams
}
