package sku

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
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
	assert.NotNil(t, i.SKUList())
}

func TestSKUNew(t *testing.T) {
	sku, err := New(&stripe.SKUParams{
		Active:     stripe.Bool(true),
		Attributes: map[string]string{"attr1": "val1", "attr2": "val2"},
		Price:      stripe.Int64(499),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Inventory: &stripe.InventoryParams{
			Type:  stripe.String(string(stripe.SKUInventoryTypeBucket)),
			Value: stripe.String(string(stripe.SKUInventoryValueLimited)),
		},
		Product: stripe.String("prod_123"),
		Image:   stripe.String("http://example.com/foo.png"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}

func TestSKUUpdate(t *testing.T) {
	sku, err := Update("sku_123", &stripe.SKUParams{
		Inventory: &stripe.InventoryParams{
			Type:  stripe.String(string(stripe.SKUInventoryTypeBucket)),
			Value: stripe.String(string(stripe.SKUInventoryValueInStock)),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, sku)
}
