package taxcode

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTaxCodeGet(t *testing.T) {
	tr, err := Get("txcd_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
}

func TestTaxCodeList(t *testing.T) {
	i := List(&stripe.TaxCodeListParams{})

	// Verify that we can get at least one tr
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TaxCode())
	assert.NotNil(t, i.TaxCodeList())
}
