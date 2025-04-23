package taxrate

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTaxRateGet(t *testing.T) {
	tr, err := Get("txr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
}

func TestTaxRateList(t *testing.T) {
	i := List(&stripe.TaxRateListParams{})

	// Verify that we can get at least one tr
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TaxRate())
	assert.NotNil(t, i.TaxRateList())
}

func TestTaxRateNew(t *testing.T) {
	tr, err := New(&stripe.TaxRateParams{
		DisplayName: stripe.String("name"),
		Inclusive:   stripe.Bool(false),
		Percentage:  stripe.Float64(10.15),
	})
	assert.Nil(t, err)
	assert.NotNil(t, tr)
}

func TestTaxRateUpdate(t *testing.T) {
	tr, err := Update("txr_123", &stripe.TaxRateParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, tr)
}
