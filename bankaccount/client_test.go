package bankaccount

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBankAccountDel_ByAccount(t *testing.T) {
	bankAccount, err := Del("ba_123", &stripe.BankAccountParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountDel_ByCustomer(t *testing.T) {
	bankAccount, err := Del("ba_123", &stripe.BankAccountParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountGet_ByAccount(t *testing.T) {
	bankAccount, err := Get("ba_123", &stripe.BankAccountParams{Account: stripe.String("acct_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountGet_ByCustomer(t *testing.T) {
	bankAccount, err := Get("ba_123", &stripe.BankAccountParams{Customer: stripe.String("cus_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountList_ByCustomer(t *testing.T) {
	i := List(&stripe.BankAccountListParams{Customer: stripe.String("cus_123")})

	// Verify that we can get at least one bank account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BankAccount())
	assert.NotNil(t, i.BankAccountList())
}

func TestBankAccountNew_ByAccount(t *testing.T) {
	bankAccount, err := New(&stripe.BankAccountParams{
		Account:            stripe.String("acct_123"),
		DefaultForCurrency: stripe.Bool(true),
		Token:              stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountNew_ByCustomer(t *testing.T) {
	bankAccount, err := New(&stripe.BankAccountParams{
		Customer: stripe.String("cus_123"),
		Token:    stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountUpdate_ByAccount(t *testing.T) {
	bankAccount, err := Update("ba_123", &stripe.BankAccountParams{
		Account:            stripe.String("acct_123"),
		DefaultForCurrency: stripe.Bool(true),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountUpdate_ByCustomer(t *testing.T) {
	bankAccount, err := Update("ba_123", &stripe.BankAccountParams{
		AccountHolderName: stripe.String("New Name"),
		Customer:          stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}
