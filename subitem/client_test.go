package subitem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSubscriptionItemDel(t *testing.T) {
	item, err := Del("si_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubscriptionItemGet(t *testing.T) {
	item, err := Get("si_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubscriptionItemList(t *testing.T) {
	i := List(&stripe.SubscriptionItemListParams{})

	// Verify that we can get at least one item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SubscriptionItem())
}

func TestSubscriptionItemNew(t *testing.T) {
	item, err := New(&stripe.SubscriptionItemParams{
		Quantity:     99,
		Plan:         "plan_123",
		Subscription: "sub_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubscriptionItemUpdate(t *testing.T) {
	item, err := Update("si_123", &stripe.SubscriptionItemParams{
		Quantity: 10,
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}
