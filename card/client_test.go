package card

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestCardDel(t *testing.T) {
	card, err := Del("card_123", &stripe.CardParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardDel_RequiresParams(t *testing.T) {
	_, err := Del("card_123", nil)
	assert.Error(t, err, "params should not be nil")
}

func TestCardGet(t *testing.T) {
	card, err := Get("card_123", &stripe.CardParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardGet_RequiresParams(t *testing.T) {
	_, err := Get("card_123", nil)
	assert.Error(t, err, "params should not be nil")
}

func TestCardList_ByCustomer(t *testing.T) {
	i := List(&stripe.CardListParams{Customer: stripe.String("cus_123")})

	// Verify that we can get at least one card
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Card())
	assert.NotNil(t, i.CardList())
}

func TestCardList_RequiresParams(t *testing.T) {
	i := List(nil)
	assert.False(t, i.Next())
	assert.Error(t, i.Err(), "params should not be nil")
}

func TestCardNew(t *testing.T) {
	card, err := New(&stripe.CardParams{
		Customer: stripe.String("cus_123"),
		Token:    stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardNew_RequiresParams(t *testing.T) {
	_, err := New(nil)
	assert.Error(t, err, "params should not be nil")
}

func TestCardUpdate(t *testing.T) {
	card, err := Update("card_123", &stripe.CardParams{
		Customer: stripe.String("cus_123"),
		Name:     stripe.String("New Name"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardUpdate_RequiresParams(t *testing.T) {
	_, err := Update("card_123", nil)
	assert.Error(t, err, "params should not be nil")
}
