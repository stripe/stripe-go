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
		Bank: &stripe.BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_WithCard(t *testing.T) {
	token, err := New(&stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: "4242424242424242", // raw PAN as we're testing token creation
			Month:  "10",
			Year:   "20",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_WithPII(t *testing.T) {
	token, err := New(&stripe.TokenParams{
		PII: &stripe.PIIParams{
			PersonalIDNumber: "000000000",
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
		Customer: "cus_123",
	}
	params.SetStripeAccount("acct_123")
	token, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestTokenNew_SharedCustomerBankAccount(t *testing.T) {
	params := &stripe.TokenParams{
		Bank: &stripe.BankAccountParams{
			ID: "ba_123",
		},
		Customer: "cus_123",
	}
	params.SetStripeAccount("acct_123")
	token, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
