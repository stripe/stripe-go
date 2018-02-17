package plan

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	_ "github.com/stripe/stripe-go/testing"
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
}

func TestPlanNew(t *testing.T) {
	plan, err := New(&stripe.PlanParams{
		Amount:   stripe.UInt64(1),
		Currency: stripe.String(string(currency.USD)),
		ID:       stripe.String("sapphire-elite"),
		Interval: stripe.String("month"),
		Product: &stripe.ProductParams{
			Name: stripe.String("Sapphire Elite"),
			Type: stripe.String(string(stripe.ProductTypeService)),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, plan)
}

func TestPlanNewWithProductID(t *testing.T) {
	productId := "prod_12345abc"
	plan, err := New(&stripe.PlanParams{
		Amount:    stripe.UInt64(1),
		Currency:  stripe.String(string(currency.USD)),
		ID:        stripe.String("sapphire-elite"),
		Interval:  stripe.String("month"),
		ProductID: &productId,
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
