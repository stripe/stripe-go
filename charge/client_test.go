package charge

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestChargeCapture(t *testing.T) {
	charge, err := Capture("ch_123", nil)
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
}

func TestChargeNew(t *testing.T) {
	charge, err := New(&stripe.ChargeParams{
		Amount:   stripe.Int64(11700),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Level3: &stripe.ChargeLevel3Params{
			LineItems: []*stripe.ChargeLevel3LineItemsParams{
				{
					DiscountAmount:     stripe.Int64(200),
					ProductCode:        stripe.String("1234"),
					ProductDescription: stripe.String("description 1"),
					Quantity:           stripe.Int64(2),
					TaxAmount:          stripe.Int64(200),
					UnitCost:           stripe.Int64(1000),
				},
				{
					DiscountAmount:     stripe.Int64(300),
					ProductCode:        stripe.String("1235"),
					ProductDescription: stripe.String("description 2"),
					Quantity:           stripe.Int64(3),
					TaxAmount:          stripe.Int64(300),
					UnitCost:           stripe.Int64(3000),
				},
			},
			MerchantReference:  stripe.String("123"),
			ShippingAddressZip: stripe.String("94110"),
			ShippingAmount:     stripe.Int64(700),
		},
		Source: &stripe.SourceParams{Token: stripe.String("src_123")},
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

func TestChargeNew_WithSetSource(t *testing.T) {
	params := stripe.ChargeParams{
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	params.SetSource("tok_123")

	charge, err := New(&params)
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
