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
	productParams := PlanProductParams{
		ID:                  "ID",
		Name:                "Sapphire Elite",
		StatementDescriptor: "SAPPHIRE",
	}
	productID := "prod_123abc"
	tiers := []*PlanTierParams{{Amount: 123, UpTo: 321}, {Amount: 123, UpToInf: true}}
	testCases := []struct {
		field  string
		params *PlanParams
		want   interface{}
	}{
		{"amount", &PlanParams{}, ""},
		{"amount", &PlanParams{Amount: 0, AmountZero: false}, ""},
		{"amount", &PlanParams{Amount: 0, AmountZero: true}, strconv.FormatUint(0, 10)},
		{"amount", &PlanParams{Amount: 123}, strconv.FormatUint(123, 10)},
		{"currency", &PlanParams{Currency: "USD"}, "USD"},
		{"id", &PlanParams{ID: "sapphire-elite"}, "sapphire-elite"},
		{"interval_count", &PlanParams{IntervalCount: 3}, strconv.FormatUint(3, 10)},
		{"interval", &PlanParams{Interval: "month"}, "month"},
		{"product", &PlanParams{ProductID: &productID}, "prod_123abc"},
		{"product[id]", &PlanParams{Product: &productParams}, "ID"},
		{"product[name]", &PlanParams{Product: &productParams}, "Sapphire Elite"},
		{"product[statement_descriptor]", &PlanParams{Product: &productParams}, "SAPPHIRE"},
		{"tiers_mode", &PlanParams{TiersMode: "volume"}, "volume"},
		{"tiers[0][amount]", &PlanParams{Tiers: tiers}, strconv.FormatUint(123, 10)},
		{"tiers[0][up_to]", &PlanParams{Tiers: tiers}, strconv.FormatUint(321, 10)},
		{"tiers[1][amount]", &PlanParams{Tiers: tiers}, strconv.FormatUint(123, 10)},
		{"tiers[1][up_to]", &PlanParams{Tiers: tiers}, "inf"},
		{"transform_usage[bucket_size]", &PlanParams{TransformUsage: &PlanTransformUsageParams{DivideBy: 123, Round: "round_up"}}, strconv.FormatUint(123, 10)},
		{"transform_usage[round]", &PlanParams{TransformUsage: &PlanTransformUsageParams{DivideBy: 123, Round: "round_up"}}, "round_up"},
		{"trial_period_days", &PlanParams{TrialPeriod: 123}, strconv.FormatUint(123, 10)},
		{"usage_type", &PlanParams{UsageType: "metered"}, "metered"},
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
