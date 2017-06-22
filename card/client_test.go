package card

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestCardDel(t *testing.T) {
	card, err := Del("card_123", &stripe.CardParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardGet(t *testing.T) {
	card, err := Get("card_123", &stripe.CardParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardListByCustomer(t *testing.T) {
	i := List(&stripe.CardListParams{Customer: "cus_123"})

	// Verify that we can get at least one card
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Card())
}

func TestCardNew(t *testing.T) {
	card, err := New(&stripe.CardParams{
		Customer: "cus_123",
		Token:    "tok_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardUpdate(t *testing.T) {
	card, err := Update("card_123", &stripe.CardParams{
		Customer: "cus_123",
		Default:  true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}
