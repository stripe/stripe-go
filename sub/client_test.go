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
		Customer:           "cus_123",
		Plan:               "plan_123",
		Quantity:           10,
		TaxPercent:         20.0,
		BillingCycleAnchor: time.Now().AddDate(0, 0, 12).Unix(),
		Billing:            "send_invoice",
		DaysUntilDue:       30,
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionNew_WithItems(t *testing.T) {
	subscription, err := New(&stripe.SubscriptionParams{
		Customer: "cus_123",
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan:         "gold",
				QuantityZero: true,
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscriptionUpdate(t *testing.T) {
	subscription, err := Update("sub_123", &stripe.SubscriptionParams{
		NoProrate:      true,
		QuantityZero:   true,
		TaxPercentZero: true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
}
