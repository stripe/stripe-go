package taxid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestTaxIDDel(t *testing.T) {
	taxid, err := Del("txi_123", &stripe.TaxIDParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, taxid)
}

func TestTaxIDGet(t *testing.T) {
	taxid, err := Get("txi_123", &stripe.TaxIDParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, taxid)
}

func TestTaxIDList(t *testing.T) {
	i := List(&stripe.TaxIDListParams{
		Customer: stripe.String("cus_123"),
	})

	// Verify that we can get at least one taxid
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TaxID())
	assert.NotNil(t, i.TaxIDList())
}

func TestTaxIDNew(t *testing.T) {
	taxid, err := New(&stripe.TaxIDParams{
		Customer: stripe.String("cus_123"),
		Type:     stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value:    stripe.String("11111"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, taxid)
}
