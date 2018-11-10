package valuelistitem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestValueListItemDel(t *testing.T) {
	vli, err := Del("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestValueListItemGet(t *testing.T) {
	vli, err := Get("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestValueListItemList(t *testing.T) {
	i := List(&stripe.ValueListItemListParams{
		ValueList: stripe.String("rsl_123"),
	})

	// Verify that we can get at least one value list item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ValueListItem())
}

func TestValueListItemNew(t *testing.T) {
	vli, err := New(&stripe.ValueListItemParams{
		Value:     stripe.String("value"),
		ValueList: stripe.String("rsl_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}
