package subscriptionitem

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
	i := List(&stripe.SubscriptionItemListParams{
		Subscription: stripe.String("sub_123"),
	})

	// Verify that we can get at least one item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SubscriptionItem())
	assert.NotNil(t, i.SubscriptionItemList())
}

func TestSubscriptionItemNew(t *testing.T) {
	item, err := New(&stripe.SubscriptionItemParams{
		Quantity:     stripe.Int64(99),
		Price:        stripe.String("price_123"),
		Subscription: stripe.String("sub_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubscriptionItemUpdate(t *testing.T) {
	item, err := Update("si_123", &stripe.SubscriptionItemParams{
		Quantity: stripe.Int64(10),
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}
