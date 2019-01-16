package valuelistitem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestRadarValueListItemDel(t *testing.T) {
	vli, err := Del("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestRadarValueListItemGet(t *testing.T) {
	vli, err := Get("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestRadarValueListItemList(t *testing.T) {
	i := List(&stripe.RadarValueListItemListParams{
		RadarValueList: stripe.String("rsl_123"),
	})

	// Verify that we can get at least one value list item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.RadarValueListItem())
}

func TestRadarValueListItemNew(t *testing.T) {
	vli, err := New(&stripe.RadarValueListItemParams{
		Value:          stripe.String("value"),
		RadarValueList: stripe.String("rsl_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}
