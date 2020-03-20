package cardholder

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestIssuingCardholderGet(t *testing.T) {
	cardholder, err := Get("ich_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, cardholder)
	assert.Equal(t, "issuing.cardholder", cardholder.Object)
}

func TestIssuingCardholderList(t *testing.T) {
	i := List(&stripe.IssuingCardholderListParams{})

	// Verify that we can get at least one cardholder
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingCardholder())
	assert.Equal(t, "issuing.cardholder", i.IssuingCardholder().Object)
}

func TestIssuingCardholderNew(t *testing.T) {
	cardholder, err := New(&stripe.IssuingCardholderParams{
		Billing: &stripe.IssuingCardholderBillingParams{
			Address: &stripe.AddressParams{
				Country:    stripe.String("US"),
				Line1:      stripe.String("line1"),
				City:       stripe.String("city"),
				PostalCode: stripe.String("90210"),
				State:      stripe.String("CA"),
			},
		},
		Individual: &stripe.IssuingCardholderIndividualParams{
			DOB: &stripe.IssuingCardholderIndividualDOBParams{
				Day:   stripe.Int64(1),
				Month: stripe.Int64(1),
				Year:  stripe.Int64(1980),
			},
			FirstName: stripe.String("Jenny"),
			LastName:  stripe.String("Rosen"),
			Verification: &stripe.IssuingCardholderIndividualVerificationParams{
				Document: &stripe.IssuingCardholderIndividualVerificationDocumentParams{
					Back:  stripe.String("file_back"),
					Front: stripe.String("file_front"),
				},
			},
		},
		Name: stripe.String("cardholder name"),
		Type: stripe.String(string(stripe.IssuingCardholderTypeIndividual)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, cardholder)
	assert.Equal(t, "issuing.cardholder", cardholder.Object)
}

func TestIssuingCardholderUpdate(t *testing.T) {
	cardholder, err := Update("ich_123", &stripe.IssuingCardholderParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, cardholder)
	assert.Equal(t, "issuing.cardholder", cardholder.Object)
}
