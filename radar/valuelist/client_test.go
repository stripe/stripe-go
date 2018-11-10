package valuelist

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestValueListDel(t *testing.T) {
	vl, err := Del("rsl_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestValueListGet(t *testing.T) {
	vl, err := Get("rsl_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestValueListList(t *testing.T) {
	i := List(&stripe.ValueListListParams{
		Alias:    stripe.String("alias"),
		Contains: stripe.String("value"),
	})

	// Verify that we can get at least one value list
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ValueList())
}

func TestValueListNew(t *testing.T) {
	vl, err := New(&stripe.ValueListParams{
		Alias:    stripe.String("alias"),
		Name:     stripe.String("name"),
		ItemType: stripe.String(string(stripe.ValueListItemTypeIPAddress)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}

func TestValueListUpdate(t *testing.T) {
	vl, err := Update("rsl_123", &stripe.ValueListParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, vl)
}
