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
		Type:                  stripe.String(string(stripe.AccountTypeCustom)),
		Country:               stripe.String("CA"),
		BusinessURL:           stripe.String("www.stripe.com"),
		BusinessName:          stripe.String("Stripe"),
		BusinessPrimaryColor:  stripe.String("#ffffff"),
		DebitNegativeBalances: stripe.Bool(true),
		SupportEmail:          stripe.String("foo@bar.com"),
		SupportURL:            stripe.String("www.stripe.com"),
		SupportPhone:          stripe.String("4151234567"),
		LegalEntity: &stripe.LegalEntityParams{
			Type:         stripe.String(string(stripe.LegalEntityTypeIndividual)),
			BusinessName: stripe.String("Stripe Go"),
			AdditionalOwners: []*stripe.AdditionalOwnerParams{
				{
					FirstName: stripe.String("Jane"),
					Verification: &stripe.IdentityVerificationParams{
						Document:     stripe.String("file_345"),
						DocumentBack: stripe.String("file_567"),
					},
				},
			},
			DOB: &stripe.DOBParams{
				Day:   stripe.Int64(1),
				Month: stripe.Int64(2),
				Year:  stripe.Int64(1990),
			},
			Verification: &stripe.IdentityVerificationParams{
				Document:     stripe.String("file_123"),
				DocumentBack: stripe.String("file_234"),
			},
		},
		TOSAcceptance: &stripe.TOSAcceptanceParams{
			IP:        stripe.String("127.0.0.1"),
			Date:      stripe.Int64(1437578361),
			UserAgent: stripe.String("Mozilla/5.0"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountReject(t *testing.T) {
	account, err := Reject("acct_123", &stripe.AccountRejectParams{
		Reason: stripe.String(string(stripe.AccountRejectReasonFraud)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountUpdate(t *testing.T) {
	account, err := Update("acct_123", &stripe.AccountParams{
		LegalEntity: &stripe.LegalEntityParams{
			Address: &stripe.AccountAddressParams{
				Country:    stripe.String("CA"),
				City:       stripe.String("Montreal"),
				PostalCode: stripe.String("H2Y 1C6"),
				Line1:      stripe.String("275, rue Notre-Dame Est"),
				State:      stripe.String("QC"),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}
