package plan

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPlanDel(t *testing.T) {
	plan, err := Del("gold", nil)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
}

func TestPlanGet(t *testing.T) {
	plan, err := Get("gold", nil)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
}

func TestPlanList(t *testing.T) {
	i := List(&stripe.PlanListParams{})

	// Verify that we can get at least one plan
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Plan())
	assert.NotNil(t, i.PlanList())
}

func TestPlanNew(t *testing.T) {
	plan, err := New(&stripe.PlanParams{
		AmountDecimal: stripe.Float64(0.0123456789),
		BillingScheme: stripe.String(string(stripe.PlanBillingSchemeTiered)),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		ID:            stripe.String("sapphire-elite"),
		Interval:      stripe.String(string(stripe.PlanIntervalMonth)),
		Product: &stripe.PlanProductParams{
			ID:                  stripe.String("plan_id"),
			Name:                stripe.String("Sapphire Elite"),
			StatementDescriptor: stripe.String("statement descriptor"),
		},
		Tiers: []*stripe.PlanTierParams{
			{UnitAmount: stripe.Int64(500), UpTo: stripe.Int64(5)},
			{UnitAmount: stripe.Int64(400), UpTo: stripe.Int64(10)},
			{UnitAmount: stripe.Int64(300), UpTo: stripe.Int64(15)},
			{UnitAmount: stripe.Int64(200), UpTo: stripe.Int64(20)},
			{UnitAmount: stripe.Int64(200), UpToInf: stripe.Bool(true)},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, plan)
}

func TestPlanUpdate(t *testing.T) {
	plan, err := Update("gold", &stripe.PlanParams{
		Nickname: stripe.String("Updated nickame"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, plan)
}
