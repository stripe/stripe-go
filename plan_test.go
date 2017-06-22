package stripe

import (
	"strconv"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestPlanListParams_AppendTo(t *testing.T) {
	testCases := []struct {
		field  string
		params *PlanListParams
		want   interface{}
	}{
		{"created", &PlanListParams{Created: 123}, strconv.FormatInt(123, 10)},
		{
			"created[gt]",
			&PlanListParams{CreatedRange: &RangeQueryParams{GreaterThan: 123}},
			strconv.FormatInt(123, 10),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			body := &form.Values{}
			form.AppendTo(body, tc.params)
			values := body.ToValues()
			assert.Equal(t, tc.want, values.Get(tc.field))
		})
	}
}

func TestPlanListParams_AppendTo_Empty(t *testing.T) {
	body := &form.Values{}
	params := &PlanListParams{}
	form.AppendTo(body, params)
	assert.True(t, body.Empty())
}

func TestPlanParams_AppendTo(t *testing.T) {
	testCases := []struct {
		field  string
		params *PlanParams
		want   interface{}
	}{
		{"amount", &PlanParams{Amount: 123}, strconv.FormatUint(123, 10)},
		{"currency", &PlanParams{Currency: "USD"}, "USD"},
		{"id", &PlanParams{ID: "sapphire-elite"}, "sapphire-elite"},
		{"interval", &PlanParams{Interval: "month"}, "month"},
		{"interval_count", &PlanParams{IntervalCount: 3}, strconv.FormatUint(3, 10)},
		{"name", &PlanParams{Name: "Sapphire Elite"}, "Sapphire Elite"},
		{"statement_descriptor", &PlanParams{Statement: "Sapphire Elite"}, "Sapphire Elite"},
		{"trial_period_days", &PlanParams{TrialPeriod: 123}, strconv.FormatUint(123, 10)},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			body := &form.Values{}
			form.AppendTo(body, tc.params)
			values := body.ToValues()
			assert.Equal(t, tc.want, values.Get(tc.field))
		})
	}
}

func TestPlanParams_AppendTo_Empty(t *testing.T) {
	body := &form.Values{}
	params := &PlanParams{}
	form.AppendTo(body, params)
	assert.True(t, body.Empty())
}
