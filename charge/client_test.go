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

func TestChargeCloseDispute(t *testing.T) {
	charge, err := CloseDispute("ch_123")
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

func TestChargeMarkFraudulent(t *testing.T) {
	charge, err := MarkFraudulent("ch_123")
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeNew(t *testing.T) {
	charge, err := New(&stripe.ChargeParams{
		Amount:   123,
		Currency: "usd",
		Source:   &stripe.SourceParams{Token: "src_123"},
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeNew_WithSetSource(t *testing.T) {
	params := stripe.ChargeParams{
		Amount:   123,
		Currency: "usd",
	}
	params.SetSource("tok_123")

	charge, err := New(&params)
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeMarkSafe(t *testing.T) {
	charge, err := MarkSafe("ch_123")
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeUpdate(t *testing.T) {
	charge, err := Update("ch_123", &stripe.ChargeParams{
		Desc: "Updated description",
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}

func TestChargeUpdateDispute(t *testing.T) {
	charge, err := UpdateDispute("ch_123", &stripe.DisputeParams{
		Evidence: &stripe.DisputeEvidenceParams{
			ProductDesc: "original description",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, charge)
}
