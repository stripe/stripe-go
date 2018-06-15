package token

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestTokenGet(t *testing.T) {
	token, err := Get("tok_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_WithBankAccount(t *testing.T) {
	token, err := New(&stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			Country:       stripe.String("US"),
			RoutingNumber: stripe.String("110000000"),
			AccountNumber: stripe.String("000123456789"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_WithCard(t *testing.T) {
	token, err := New(&stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String("4242424242424242"), // raw PAN as we're testing token creation
			ExpMonth: stripe.String("10"),
			ExpYear:  stripe.String("20"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_WithPII(t *testing.T) {
	token, err := New(&stripe.TokenParams{
		PII: &stripe.PIIParams{
			PersonalIDNumber: stripe.String("000000000"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_SharedCustomerCard(t *testing.T) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			ID: "card_123",
		},
		Customer: stripe.String("cus_123"),
	}
	params.SetStripeAccount("acct_123")
	token, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_SharedCustomerBankAccount(t *testing.T) {
	params := &stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			ID: stripe.String("ba_123"),
		},
		Customer: stripe.String("cus_123"),
	}
	params.SetStripeAccount("acct_123")
	token, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
