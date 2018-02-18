package sub

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSubscriptionCancel(t *testing.T) {
	subscription, err := Cancel("sub_123", nil)
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
		Customer:           stripe.String("cus_123"),
		Plan:               stripe.String("plan_123"),
		Quantity:           stripe.UInt64(10),
		TaxPercent:         stripe.Float64(20.0),
		BillingCycleAnchor: stripe.Int64(time.Now().AddDate(0, 0, 12).Unix()),
		Billing:            stripe.String("send_invoice"),
		DaysUntilDue:       stripe.UInt64(30),
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionNew_WithItems(t *testing.T) {
	subscription, err := New(&stripe.SubscriptionParams{
		Customer: stripe.String("cus_123"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan:     "gold",
				Quantity: stripe.UInt64(0),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionUpdate(t *testing.T) {
	subscription, err := Update("sub_123", &stripe.SubscriptionParams{
		Prorate:    stripe.Bool(true),
		Quantity:   stripe.UInt64(0),
		TaxPercent: stripe.Float64(0),
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}
