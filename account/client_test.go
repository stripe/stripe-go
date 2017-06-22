package account

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestAccountDel(t *testing.T) {
	account, err := Del("acct_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountGet(t *testing.T) {
	account, err := Get()
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountGetByID(t *testing.T) {
	account, err := GetByID("acct_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountList(t *testing.T) {
	i := List(&stripe.AccountListParams{})

	// Verify that we can get at least one account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Account())
}

func TestAccountNew(t *testing.T) {
	account, err := New(&stripe.AccountParams{
		Type:                 stripe.AccountTypeCustom,
		Country:              "CA",
		BusinessUrl:          "www.stripe.com",
		BusinessName:         "Stripe",
		BusinessPrimaryColor: "#ffffff",
		DebitNegativeBal:     true,
		SupportEmail:         "foo@bar.com",
		SupportUrl:           "www.stripe.com",
		SupportPhone:         "4151234567",
		LegalEntity: &stripe.LegalEntity{
			Type:         stripe.Individual,
			BusinessName: "Stripe Go",
			AdditionalOwners: []stripe.Owner{
				{First: "Jane"},
			},
			DOB: stripe.DOB{
				Day:   1,
				Month: 2,
				Year:  1990,
			},
		},
		TOSAcceptance: &stripe.TOSAcceptanceParams{
			IP:        "127.0.0.1",
			Date:      1437578361,
			UserAgent: "Mozilla/5.0",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountReject(t *testing.T) {
	account, err := Reject("acct_123", &stripe.AccountRejectParams{
		Reason: "fraud",
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountUpdate(t *testing.T) {
	account, err := Update("acct_123", &stripe.AccountParams{
		Type:    stripe.AccountTypeCustom,
		Country: "CA",
		LegalEntity: &stripe.LegalEntity{
			Address: stripe.Address{
				Country: "CA",
				City:    "Montreal",
				Zip:     "H2Y 1C6",
				Line1:   "275, rue Notre-Dame Est",
				State:   "QC",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}
