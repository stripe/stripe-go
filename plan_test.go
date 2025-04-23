package stripe

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestPlan_Unmarshal(t *testing.T) {
	planData := map[string]interface{}{
		"id":             "pl_123",
		"object":         "plan",
		"amount":         0,
		"amount_decimal": "0.0123456789",
	}

	bytes, err := json.Marshal(&planData)
	assert.NoError(t, err)

	var plan Plan
	err = json.Unmarshal(bytes, &plan)
	assert.NoError(t, err)

	assert.Equal(t, 0.0123456789, plan.AmountDecimal)
}

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
	productParams := PlanProductParams{
		ID:                  String("ID"),
		Name:                String("Sapphire Elite"),
		StatementDescriptor: String("SAPPHIRE"),
	}
	tiers := []*PlanTierParams{
		{UnitAmount: Int64(123), UpTo: Int64(321)},
		{UnitAmount: Int64(123), UpToInf: Bool(true)}}
	testCases := []struct {
		field  string
		params *PlanParams
		want   interface{}
	}{
		{"amount", &PlanParams{Amount: Int64(123)}, strconv.FormatUint(123, 10)},
		{"currency", &PlanParams{Currency: String(string(CurrencyUSD))}, "usd"},
		{"id", &PlanParams{ID: String("sapphire-elite")}, "sapphire-elite"},
		{"interval", &PlanParams{Interval: String("month")}, "month"},
		{"interval_count", &PlanParams{IntervalCount: Int64(3)}, strconv.FormatUint(3, 10)},
		{"product[id]", &PlanParams{Product: &PlanProductParams{ID: String("prod_123abc")}}, "prod_123abc"},
		{"product[name]", &PlanParams{Product: &productParams}, "Sapphire Elite"},
		{"product[statement_descriptor]", &PlanParams{Product: &productParams}, "SAPPHIRE"},
		{"tiers_mode", &PlanParams{TiersMode: String(string(PlanTiersModeVolume))}, "volume"},
		{"tiers[0][unit_amount]", &PlanParams{Tiers: tiers}, strconv.FormatUint(123, 10)},
		{"tiers[0][up_to]", &PlanParams{Tiers: tiers}, strconv.FormatUint(321, 10)},
		{"tiers[1][unit_amount]", &PlanParams{Tiers: tiers}, strconv.FormatUint(123, 10)},
		{"tiers[1][up_to]", &PlanParams{Tiers: tiers}, "inf"},
		{"transform_usage[divide_by]", &PlanParams{TransformUsage: &PlanTransformUsageParams{DivideBy: Int64(123), Round: String("round_up")}}, strconv.FormatUint(123, 10)},
		{"transform_usage[round]", &PlanParams{TransformUsage: &PlanTransformUsageParams{DivideBy: Int64(123), Round: String("round_up")}}, "round_up"},
		{"trial_period_days", &PlanParams{TrialPeriodDays: Int64(123)}, strconv.FormatUint(123, 10)},
		{"usage_type", &PlanParams{UsageType: String("metered")}, "metered"},
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
