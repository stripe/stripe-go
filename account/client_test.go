package account

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
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
	assert.NotNil(t, i.AccountList())
}

func TestAccountNew(t *testing.T) {
	account, err := New(&stripe.AccountParams{
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			Name:         stripe.String("name"),
			SupportEmail: stripe.String("foo@bar.com"),
			SupportURL:   stripe.String("www.stripe.com"),
			SupportPhone: stripe.String("4151234567"),
		},
		BusinessType: stripe.String(string(stripe.AccountBusinessTypeCompany)),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Company: &stripe.AccountCompanyParams{
			DirectorsProvided: stripe.Bool(true),
			Name:              stripe.String("company_name"),
			Verification: &stripe.AccountCompanyVerificationParams{
				Document: &stripe.AccountCompanyVerificationDocumentParams{
					Back:  stripe.String("file_123"),
					Front: stripe.String("file_abc"),
				},
			},
		},
		Country: stripe.String("CA"),
		ExternalAccount: &stripe.AccountExternalAccountParams{
			Token: stripe.String("tok_123"),
		},
		Settings: &stripe.AccountSettingsParams{
			Branding: &stripe.AccountSettingsBrandingParams{
				Icon: stripe.String("file_123"),
				Logo: stripe.String("file_234"),
			},
			CardPayments: &stripe.AccountSettingsCardPaymentsParams{
				DeclineOn: &stripe.AccountDeclineSettingsParams{
					AVSFailure: stripe.Bool(true),
					CVCFailure: stripe.Bool(true),
				},
				StatementDescriptorPrefix: stripe.String("prefix"),
			},
			Payments: &stripe.AccountSettingsPaymentsParams{
				StatementDescriptor: stripe.String("descriptor"),
			},
			Payouts: &stripe.AccountSettingsPayoutsParams{
				DebitNegativeBalances: stripe.Bool(true),
				Schedule: &stripe.PayoutScheduleParams{
					DelayDaysMinimum: stripe.Bool(true),
					Interval:         stripe.String(string(stripe.PayoutIntervalManual)),
				},
				StatementDescriptor: stripe.String("payout_descriptor"),
			},
		},
		Type: stripe.String(string(stripe.AccountTypeCustom)),
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
		Company: &stripe.AccountCompanyParams{
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
