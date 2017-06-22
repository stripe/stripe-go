package stripe

import (
	"strconv"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPlanParamsAppendTo(t *testing.T) {
	var body *RequestValues
	var params PlanParams

	// Should run even with an empty set of parameters
	body = &RequestValues{}
	params.AppendTo(body)
	assert.True(t, body.Empty())

	body = &RequestValues{}
	params = PlanParams{
		Amount:        99,
		Currency:      Currency("usd"),
		ID:            "test_plan",
		Interval:      PlanInterval("month"),
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
