package bankaccount

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestBankAccountDel_ByAccount(t *testing.T) {
	bankAcount, err := Del("ba_123", &stripe.BankAccountParams{
		AccountID: "acct_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountDel_ByCustomer(t *testing.T) {
	bankAcount, err := Del("ba_123", &stripe.BankAccountParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountGet_ByAccount(t *testing.T) {
	bankAcount, err := Get("ba_123", &stripe.BankAccountParams{AccountID: "acct_123"})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountGet_ByCustomer(t *testing.T) {
	bankAcount, err := Get("ba_123", &stripe.BankAccountParams{Customer: "cus_123"})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountList_ByAccount(t *testing.T) {
	i := List(&stripe.BankAccountListParams{Customer: "acct_123"})

	// Verify that we can get at least one bank account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BankAccount())
}

func TestBankAccountList_ByCustomer(t *testing.T) {
	i := List(&stripe.BankAccountListParams{Customer: "cus_123"})

	// Verify that we can get at least one bank account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BankAccount())
}

func TestBankAccountNew_ByAccount(t *testing.T) {
	bankAcount, err := New(&stripe.BankAccountParams{
		AccountID: "acct_123",
		Default:   true,
		Token:     "tok_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountNew_ByCustomer(t *testing.T) {
	bankAcount, err := New(&stripe.BankAccountParams{
		Customer: "cus_123",
		Default:  true,
		Token:    "tok_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountUpdate_ByAccount(t *testing.T) {
	bankAcount, err := Update("ba_123", &stripe.BankAccountParams{
		AccountID: "acct_123",
		Default:   true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountUpdate_ByCustomer(t *testing.T) {
	bankAcount, err := Update("ba_123", &stripe.BankAccountParams{
		Customer: "cus_123",
		Default:  true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}
