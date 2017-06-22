package subitem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSubItemDel(t *testing.T) {
	item, err := Del("si_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubItemGet(t *testing.T) {
	item, err := Get("si_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubItemList(t *testing.T) {
	i := List(&stripe.SubItemListParams{})

	// Verify that we can get at least one item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SubItem())
}

func TestSubItemNew(t *testing.T) {
	item, err := New(&stripe.SubItemParams{
		Quantity: 99,
		Plan:     "plan_123",
		Sub:      "sub_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestSubItemUpdate(t *testing.T) {
	item, err := Update("si_123", &stripe.SubItemParams{
		Quantity: 10,
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}
