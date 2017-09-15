package sku

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSKUDel(t *testing.T) {
	sku, err := Del("sku_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}

func TestSKUGet(t *testing.T) {
	sku, err := Get("sku_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}

func TestSKUList(t *testing.T) {
	i := List(&stripe.SKUListParams{})

	// Verify that we can get at least one sku
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SKU())
}

func TestSKUNew(t *testing.T) {
	active := true
	sku, err := New(&stripe.SKUParams{
		Active:    &active,
		Attrs:     map[string]string{"attr1": "val1", "attr2": "val2"},
		Price:     499,
		Currency:  "usd",
		Inventory: stripe.Inventory{Type: "bucket", Value: "limited"},
		Product:   "prod_123",
		Image:     "http://example.com/foo.png",
	})
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}

func TestSKUUpdate(t *testing.T) {
	sku, err := Update("sku_123", &stripe.SKUParams{
		Inventory: stripe.Inventory{Type: "bucket", Value: "in_stock"},
	})
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}
