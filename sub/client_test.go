package sub

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSubscriptionCancel(t *testing.T) {
	subscription, err := Cancel("sub_123", &stripe.SubscriptionCancelParams{
		AtPeriodEnd: stripe.Bool(true),
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
}

func TestSubscriptionNew(t *testing.T) {
	subscription, err := New(&stripe.SubscriptionParams{
		Customer: stripe.String("cus_123"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan:     stripe.String("plan_123"),
				Quantity: stripe.Int64(10),
			},
		},
		TaxPercent:         stripe.Float64(20.0),
		BillingCycleAnchor: stripe.Int64(time.Now().AddDate(0, 0, 12).Unix()),
		Billing:            stripe.String(string(stripe.SubscriptionBillingSendInvoice)),
		DaysUntilDue:       stripe.Int64(30),
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
	subscription, err := Update("sub_123", &stripe.SubscriptionParams{
		Prorate:    stripe.Bool(true),
		TaxPercent: stripe.Float64(0),
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}
