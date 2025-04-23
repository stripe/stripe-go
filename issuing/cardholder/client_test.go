package cardholder

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
	assert.NotNil(t, i.IssuingCardholderList())
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
		SpendingControls: &stripe.IssuingCardholderSpendingControlsParams{
			AllowedCategories: stripe.StringSlice([]string{
				"fast_food_restaurants",
				"miscellaneous_food_stores",
			}),
			SpendingLimits: []*stripe.IssuingCardholderSpendingControlsSpendingLimitParams{
				{
					Amount:   stripe.Int64(1000),
					Interval: stripe.String(string(stripe.IssuingCardholderSpendingControlsSpendingLimitIntervalWeekly)),
				},
			},
		},
		Type: stripe.String(string(stripe.IssuingCardholderTypeIndividual)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, cardholder)
	assert.Equal(t, "issuing.cardholder", cardholder.Object)
}

// IssuingCardholderSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardholderSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing cardholder
type IssuingCardholderSpendingControlsParams struct {
	AllowedCategories      []*string                                               `form:"allowed_categories"`
	BlockedCategories      []*string                                               `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                                 `form:"spending_limits_currency"`
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
