package subscription

import (
	"testing"
	"time"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSubscriptionCancel(t *testing.T) {
	subscription, err := Cancel("sub_123", &stripe.SubscriptionCancelParams{
		InvoiceNow: stripe.Bool(true),
		Prorate:    stripe.Bool(true),
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionGet(t *testing.T) {
	subscription, err := Get("sub_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionList(t *testing.T) {
	i := List(&stripe.SubscriptionListParams{})

	// Verify that we can get at least one subscription
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Subscription())
	assert.NotNil(t, i.SubscriptionList())
}

func TestSubscriptionNew(t *testing.T) {
	subscription, err := New(&stripe.SubscriptionParams{
		AddInvoiceItems: []*stripe.SubscriptionAddInvoiceItemParams{
			{
				Price:    stripe.String("price_123"),
				Quantity: stripe.Int64(2),
			},
			{
				PriceData: &stripe.InvoiceItemPriceDataParams{
					Currency:   stripe.String(string(stripe.CurrencyUSD)),
					UnitAmount: stripe.Int64(1000),
					Product:    stripe.String("prod_123"),
				},
				Quantity: stripe.Int64(4),
			},
		},
		BillingCycleAnchor: stripe.Int64(time.Now().AddDate(0, 0, 12).Unix()),
		CollectionMethod:   stripe.String(string(stripe.SubscriptionCollectionMethodChargeAutomatically)),
		Customer:           stripe.String("cus_123"),
		DaysUntilDue:       stripe.Int64(30),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    stripe.String("price_ABC"),
				Quantity: stripe.Int64(10),
			},
			{
				PriceData: &stripe.SubscriptionItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					Product:  stripe.String("prod_ABC"),
					Recurring: &stripe.SubscriptionItemPriceDataRecurringParams{
						Interval:      stripe.String(string(stripe.PriceRecurringIntervalMonth)),
						IntervalCount: stripe.Int64(6),
					},
					UnitAmount: stripe.Int64(1000),
				},
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionNew_WithItems(t *testing.T) {
	subscription, err := New(&stripe.SubscriptionParams{
		Customer: stripe.String("cus_123"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan:     stripe.String("gold"),
				Quantity: stripe.Int64(0),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionUpdate(t *testing.T) {
	params := &stripe.SubscriptionParams{
		ProrationBehavior: stripe.String("none"),
	}
	subscription, err := Update("sub_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestDeleteDiscount(t *testing.T) {
	discount, err := DeleteDiscount("sub_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, discount)
}
