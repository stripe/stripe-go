package valuelist

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestRadarValueListDel(t *testing.T) {
	vl, err := Del("rsl_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestRadarValueListGet(t *testing.T) {
	vl, err := Get("rsl_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestRadarValueListList(t *testing.T) {
	i := List(&stripe.RadarValueListListParams{
		Alias:    stripe.String("alias"),
		Contains: stripe.String("value"),
	})

	// Verify that we can get at least one value list
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.RadarValueList())
	assert.NotNil(t, i.RadarValueListList())
}

func TestRadarValueListNew(t *testing.T) {
	vl, err := New(&stripe.RadarValueListParams{
		Alias:    stripe.String("alias"),
		Name:     stripe.String("name"),
		ItemType: stripe.String(string(stripe.RadarValueListItemTypeIPAddress)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestRadarValueListUpdate(t *testing.T) {
	vl, err := Update("rsl_123", &stripe.RadarValueListParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}
