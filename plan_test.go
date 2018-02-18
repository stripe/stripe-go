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
		{"created", &PlanListParams{Created: Int64(123)}, strconv.FormatInt(123, 10)},
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
	productParams := ProductParams{
		Name:                String("Sapphire Elite"),
		StatementDescriptor: String("SAPPHIRE"),
		Type:                String(string(ProductTypeService)),
	}
	productId := "prod_123abc"
	testCases := []struct {
		field  string
		params *PlanParams
		want   interface{}
	}{
		{"amount", &PlanParams{Amount: UInt64(123)}, strconv.FormatUint(123, 10)},
		{"currency", &PlanParams{Currency: String("usd")}, "usd"},
		{"id", &PlanParams{ID: String("sapphire-elite")}, "sapphire-elite"},
		{"interval", &PlanParams{Interval: String("month")}, "month"},
		{"interval_count", &PlanParams{IntervalCount: UInt64(3)}, strconv.FormatUint(3, 10)},
		{"product[name]", &PlanParams{Product: &productParams}, "Sapphire Elite"},
		{"product[statement_descriptor]", &PlanParams{Product: &productParams}, "SAPPHIRE"},
		{"product", &PlanParams{ProductID: &productId}, "prod_123abc"},
		{"trial_period_days", &PlanParams{TrialPeriodDays: UInt64(123)}, strconv.FormatUint(123, 10)},
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
