package stripe_test

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v84"
	. "github.com/stripe/stripe-go/v84/testing"
)

func TestBankAccountDelete_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Delete(context.TODO(), "ba_123", &stripe.BankAccountDeleteParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountDelete_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Delete(context.TODO(), "ba_123", &stripe.BankAccountDeleteParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountRetrieve_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Retrieve(context.TODO(), "ba_123", &stripe.BankAccountRetrieveParams{Account: stripe.String("acct_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountRetrieve_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Retrieve(context.TODO(), "ba_123", &stripe.BankAccountRetrieveParams{Customer: stripe.String("cus_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountList_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1BankAccounts.List(context.TODO(), &stripe.BankAccountListParams{Account: stripe.String("acct_123")})
	i.All(context.TODO())(func(ba *stripe.BankAccount, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, ba)
		return true
	})
}

func TestBankAccountList_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1BankAccounts.List(context.TODO(), &stripe.BankAccountListParams{Customer: stripe.String("cus_123")})
	i.All(context.TODO())(func(ba *stripe.BankAccount, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, ba)
		return true
	})
}

func TestBankAccountCreate_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Create(context.TODO(), &stripe.BankAccountCreateParams{
		Account:            stripe.String("acct_123"),
		Token:              stripe.String("tok_123"),
		DefaultForCurrency: stripe.Bool(false),
		Metadata: map[string]string{
			"key": "value",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountCreate_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Create(context.TODO(), &stripe.BankAccountCreateParams{
		Customer: stripe.String("cus_123"),
		Token:    stripe.String("tok_123"),
		Metadata: map[string]string{
			"key": "value",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountUpdate_ByCustomer(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Update(context.TODO(), "ba_123", &stripe.BankAccountUpdateParams{
		Customer:          stripe.String("cus_123"),
		AccountHolderName: stripe.String("Jenny Rosen"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}

func TestBankAccountUpdate_ByAccount(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	bankAccount, err := sc.V1BankAccounts.Update(context.TODO(), "ba_123", &stripe.BankAccountUpdateParams{
		Account:           stripe.String("acct_123"),
		AccountHolderName: stripe.String("Jenny Rosen"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAccount)
}
