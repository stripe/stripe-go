package account_test

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/client"
	. "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestAccountDel(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.Del("acct_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountGet(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.Get()
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountGetByID(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.GetByID("acct_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountList(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	i := sc.Accounts.List(&stripe.AccountListParams{})

	// Verify that we can get at least one account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Account())
	assert.NotNil(t, i.AccountList())
}

func TestAccountNew(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.New(&stripe.AccountParams{
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
		Documents: &stripe.AccountDocumentsParams{
			CompanyLicense: &stripe.AccountDocumentsCompanyLicenseParams{
				Files: []*string{stripe.String("file_xyz")},
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
				DeclineOn: &stripe.AccountSettingsCardPaymentsDeclineOnParams{
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
				Schedule: &stripe.AccountSettingsPayoutsScheduleParams{
					DelayDaysMinimum: stripe.Bool(true),
					Interval:         stripe.String(string(stripe.AccountSettingsPayoutsScheduleIntervalManual)),
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
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.Reject("acct_123", &stripe.AccountRejectParams{
		Reason: stripe.String("fraud"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestAccountUpdate(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	account, err := sc.Accounts.Update("acct_123", &stripe.AccountParams{
		Company: &stripe.AccountCompanyParams{
			Address: &stripe.AddressParams{
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
