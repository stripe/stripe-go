package product

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
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
}

func TestProductNew(t *testing.T) {
	active := true
	shippable := true

	product, err := New(&stripe.ProductParams{
		Active:      &active,
		Name:        stripe.String("Test Name"),
		Description: stripe.String("This is a description"),
		Caption:     stripe.String("This is a caption"),
		Attributes:  []string{"attr1", "attr2"},
		URL:         stripe.String("http://example.com"),
		Shippable:   &shippable,
		PackageDimensions: &stripe.PackageDimensionsParams{
			Height: stripe.Float64(2.234),
			Length: stripe.Float64(5.10),
			Width:  stripe.Float64(6.50),
			Weight: stripe.Float64(10),
		},
		Type: stripe.String(string(stripe.ProductTypeGood)),
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
