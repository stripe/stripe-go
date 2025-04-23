package product

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestProductDel(t *testing.T) {
	product, err := Del("prod_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestProductGet(t *testing.T) {
	product, err := Get("prod_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestProductList(t *testing.T) {
	i := List(&stripe.ProductListParams{})

	// Verify that we can get at least one product
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Product())
	assert.NotNil(t, i.ProductList())
}

func TestProductNew(t *testing.T) {
	product, err := New(&stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String("Test Name"),
		Description: stripe.String("This is a description"),
		URL:         stripe.String("http://example.com"),
		Shippable:   stripe.Bool(true),
		PackageDimensions: &stripe.ProductPackageDimensionsParams{
			Height: stripe.Float64(2.234),
			Length: stripe.Float64(5.10),
			Width:  stripe.Float64(6.50),
			Weight: stripe.Float64(10),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestProductUpdate(t *testing.T) {
	product, err := Update("prod_123", &stripe.ProductParams{
		Name: stripe.String("Updated Name"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, product)
}
