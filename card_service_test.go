package stripe_test

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
	. "github.com/stripe/stripe-go/v82/testing"
)

func TestCardDelete_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	card, err := sc.V1Cards.Delete(context.TODO(), "card_123", &stripe.CardDeleteParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardRetrieve_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	card, err := sc.V1Cards.Retrieve(context.TODO(), "card_123", &stripe.CardRetrieveParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardList_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1Cards.List(context.TODO(), &stripe.CardListParams{Customer: stripe.String("cus_123")})
	i(func(card *stripe.Card, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, card)
		return true
	})
}

func TestCardList_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1Cards.List(context.TODO(), &stripe.CardListParams{Account: stripe.String("acct_123")})
	i(func(card *stripe.Card, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, card)
		return true
	})
}

func TestCardCreate_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	card, err := sc.V1Cards.Create(context.TODO(), &stripe.CardCreateParams{
		Customer: stripe.String("cus_123"),
		Token:    stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardCreate_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	card, err := sc.V1Cards.Create(context.TODO(), &stripe.CardCreateParams{
		Account: stripe.String("acct_123"),
		Token:   stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardUpdate_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	card, err := sc.V1Cards.Update(context.TODO(), "card_123", &stripe.CardUpdateParams{
		Customer: stripe.String("cus_123"),
		Name:     stripe.String("New Name"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, card)
}
