package charge

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestChargeCapture(t *testing.T) {
	charge, err := Capture("ch_123", &stripe.ChargeCaptureParams{
		Amount: stripe.Int64(123),
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeGet(t *testing.T) {
	charge, err := Get("ch_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeList(t *testing.T) {
	i := List(&stripe.ChargeListParams{})

	// Verify that we can get at least one charge
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Charge())
	assert.NotNil(t, i.ChargeList())
}

func TestChargeSearch(t *testing.T) {
	i := Search(&stripe.ChargeSearchParams{SearchParams: stripe.SearchParams{
		Query: "currency:\"USD\"",
	}})

	// Verify that we got a charge
	assert.Equal(t, *i.Meta().TotalCount, uint32(1))
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Charge())
	assert.False(t, i.Next())
}

func TestChargeNew(t *testing.T) {
	charge, err := New(&stripe.ChargeParams{
		Amount:   stripe.Int64(11700),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.PaymentSourceSourceParams{Token: stripe.String("src_123")},
		Shipping: &stripe.ShippingDetailsParams{
			Address: &stripe.AddressParams{
				Line1: stripe.String("line1"),
				City:  stripe.String("city"),
			},
			Carrier: stripe.String("carrier"),
			Name:    stripe.String("name"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeUpdate(t *testing.T) {
	charge, err := Update("ch_123", &stripe.ChargeParams{
		Description: stripe.String("Updated description"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}
