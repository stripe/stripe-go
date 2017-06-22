package plan

import (
	"strconv"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
)

func TestPlanParamsAppendTo(t *testing.T) {
	var body *stripe.RequestValues
	var params PlanParams

	// Should run even with an empty set of parameters
	body = &stripe.RequestValues{}
	params.AppendTo(body)
	assert.True(t, body.Empty())

	body = &stripe.RequestValues{}
	params = PlanParams{
		Amount:        99,
		Currency:      currency.USD,
		ID:            "test_plan",
		Interval:      Month,
		IntervalCount: 3,
		Name:          "Test Plan",
		Statement:     "Test Plan",
		TrialPeriod:   30,
	}
	params.AppendTo(body)
	values := body.ToValues()
	assert.Equal(t, strconv.FormatUint(params.Amount, 10), values.Get("amount"))
	assert.Equal(t, string(params.Currency), values.Get("currency"))
	assert.Equal(t, params.ID, values.Get("id"))
	assert.Equal(t, string(params.Interval), values.Get("interval"))
	assert.Equal(t, strconv.FormatUint(params.IntervalCount, 10), values.Get("interval_count"))
	assert.Equal(t, params.Name, values.Get("name"))
	assert.Equal(t, params.Statement, values.Get("statement_descriptor"))
	assert.Equal(t, strconv.FormatUint(params.TrialPeriod, 10), values.Get("trial_period_days"))
}
