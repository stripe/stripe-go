package price

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPriceGet(t *testing.T) {
	price, err := Get("gold", nil)
	assert.Nil(t, err)
	assert.NotNil(t, price)
}

func TestPriceList(t *testing.T) {
	params := &stripe.PriceListParams{
		Active: stripe.Bool(true),
		LookupKeys: stripe.StringSlice([]string{
			"Key1",
			"Key2",
		}),
		Recurring: &stripe.PriceListRecurringParams{
			Interval:  stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			UsageType: stripe.String(string(stripe.PriceRecurringUsageTypeLicensed)),
		},
	}
	i := List(params)

	// Verify that we can get at least one price
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Price())
	assert.NotNil(t, i.PriceList())
}

func TestPriceNew(t *testing.T) {
	price, err := New(&stripe.PriceParams{
		BillingScheme: stripe.String(string(stripe.PriceBillingSchemeTiered)),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval:      stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			IntervalCount: stripe.Int64(6),
			UsageType:     stripe.String(string(stripe.PriceRecurringUsageTypeLicensed)),
		},
		ProductData: &stripe.PriceProductDataParams{
			Name:                stripe.String("Sapphire Elite"),
			StatementDescriptor: stripe.String("statement descriptor"),
		},
		Tiers: []*stripe.PriceTierParams{
			{UnitAmount: stripe.Int64(500), UpTo: stripe.Int64(5)},
			{UnitAmount: stripe.Int64(400), UpTo: stripe.Int64(10)},
			{UnitAmount: stripe.Int64(300), UpTo: stripe.Int64(15)},
			{UnitAmount: stripe.Int64(200), UpTo: stripe.Int64(20)},
			{UnitAmount: stripe.Int64(200), UpToInf: stripe.Bool(true)},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, price)
}

func TestPriceUpdate(t *testing.T) {
	price, err := Update("gold", &stripe.PriceParams{
		Nickname: stripe.String("Updated nickame"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, price)
}
