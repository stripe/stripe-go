package card

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestIssuingCardDetails(t *testing.T) {
	cardDetails, err := Details("ic_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, cardDetails)
	assert.Equal(t, "issuing.card_details", cardDetails.Object)
}

func TestIssuingCardGet(t *testing.T) {
	card, err := Get("ic_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, "issuing.card", card.Object)
}

func TestIssuingCardList(t *testing.T) {
	i := List(&stripe.IssuingCardListParams{})

	// Verify that we can get at least one card
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingCard())
	assert.Equal(t, "issuing.card", i.IssuingCard().Object)
}

func TestIssuingCardNew(t *testing.T) {
	params := &stripe.IssuingCardParams{
		AuthorizationControls: &stripe.AuthorizationControlsParams{
			SpendingLimits: []*stripe.IssuingAuthorizationControlsSpendingLimitsParams{
				{
					Amount: stripe.Int64(1000),
					Categories: stripe.StringSlice([]string{
						"financial_institutions",
					}),
					Interval: stripe.String(string(stripe.IssuingSpendingLimitIntervalAllTime)),
				},
			},
		},
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Type:     stripe.String(string(stripe.IssuingCardTypeVirtual)),
	}
	card, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, "issuing.card", card.Object)
}

func TestIssuingCardUpdate(t *testing.T) {
	card, err := Update("ic_123", &stripe.IssuingCardParams{
		Status: stripe.String(string(stripe.IssuingCardStatusInactive)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, "issuing.card", card.Object)
}
