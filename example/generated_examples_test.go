//
//
// File generated from our OpenAPI spec
//
//

package example

import (
	"testing"

	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
	account "github.com/max-cape/stripe-go-test/account"
	accountlink "github.com/max-cape/stripe-go-test/accountlink"
	applicationfee "github.com/max-cape/stripe-go-test/applicationfee"
	apps_secret "github.com/max-cape/stripe-go-test/apps/secret"
	balancetransaction "github.com/max-cape/stripe-go-test/balancetransaction"
	billingportal_configuration "github.com/max-cape/stripe-go-test/billingportal/configuration"
	billingportal_session "github.com/max-cape/stripe-go-test/billingportal/session"
	capability "github.com/max-cape/stripe-go-test/capability"
	card "github.com/max-cape/stripe-go-test/card"
	cashbalance "github.com/max-cape/stripe-go-test/cashbalance"
	charge "github.com/max-cape/stripe-go-test/charge"
	checkout_session "github.com/max-cape/stripe-go-test/checkout/session"
	"github.com/max-cape/stripe-go-test/client"
	countryspec "github.com/max-cape/stripe-go-test/countryspec"
	coupon "github.com/max-cape/stripe-go-test/coupon"
	customer "github.com/max-cape/stripe-go-test/customer"
	customerbalancetransaction "github.com/max-cape/stripe-go-test/customerbalancetransaction"
	customercashbalancetransaction "github.com/max-cape/stripe-go-test/customercashbalancetransaction"
	customersession "github.com/max-cape/stripe-go-test/customersession"
	dispute "github.com/max-cape/stripe-go-test/dispute"
	event "github.com/max-cape/stripe-go-test/event"
	feerefund "github.com/max-cape/stripe-go-test/feerefund"
	financialconnections_account "github.com/max-cape/stripe-go-test/financialconnections/account"
	financialconnections_session "github.com/max-cape/stripe-go-test/financialconnections/session"
	financialconnections_transaction "github.com/max-cape/stripe-go-test/financialconnections/transaction"
	identity_verificationreport "github.com/max-cape/stripe-go-test/identity/verificationreport"
	identity_verificationsession "github.com/max-cape/stripe-go-test/identity/verificationsession"
	invoice "github.com/max-cape/stripe-go-test/invoice"
	invoiceitem "github.com/max-cape/stripe-go-test/invoiceitem"
	issuing_authorization "github.com/max-cape/stripe-go-test/issuing/authorization"
	issuing_card "github.com/max-cape/stripe-go-test/issuing/card"
	issuing_cardholder "github.com/max-cape/stripe-go-test/issuing/cardholder"
	issuing_dispute "github.com/max-cape/stripe-go-test/issuing/dispute"
	issuing_personalizationdesign "github.com/max-cape/stripe-go-test/issuing/personalizationdesign"
	issuing_physicalbundle "github.com/max-cape/stripe-go-test/issuing/physicalbundle"
	issuing_transaction "github.com/max-cape/stripe-go-test/issuing/transaction"
	loginlink "github.com/max-cape/stripe-go-test/loginlink"
	mandate "github.com/max-cape/stripe-go-test/mandate"
	paymentintent "github.com/max-cape/stripe-go-test/paymentintent"
	paymentlink "github.com/max-cape/stripe-go-test/paymentlink"
	paymentmethod "github.com/max-cape/stripe-go-test/paymentmethod"
	paymentmethodconfiguration "github.com/max-cape/stripe-go-test/paymentmethodconfiguration"
	paymentsource "github.com/max-cape/stripe-go-test/paymentsource"
	payout "github.com/max-cape/stripe-go-test/payout"
	person "github.com/max-cape/stripe-go-test/person"
	plan "github.com/max-cape/stripe-go-test/plan"
	price "github.com/max-cape/stripe-go-test/price"
	product "github.com/max-cape/stripe-go-test/product"
	promotioncode "github.com/max-cape/stripe-go-test/promotioncode"
	quote "github.com/max-cape/stripe-go-test/quote"
	radar_earlyfraudwarning "github.com/max-cape/stripe-go-test/radar/earlyfraudwarning"
	radar_valuelist "github.com/max-cape/stripe-go-test/radar/valuelist"
	radar_valuelistitem "github.com/max-cape/stripe-go-test/radar/valuelistitem"
	refund "github.com/max-cape/stripe-go-test/refund"
	reporting_reportrun "github.com/max-cape/stripe-go-test/reporting/reportrun"
	reporting_reporttype "github.com/max-cape/stripe-go-test/reporting/reporttype"
	review "github.com/max-cape/stripe-go-test/review"
	setupattempt "github.com/max-cape/stripe-go-test/setupattempt"
	setupintent "github.com/max-cape/stripe-go-test/setupintent"
	shippingrate "github.com/max-cape/stripe-go-test/shippingrate"
	sigma_scheduledqueryrun "github.com/max-cape/stripe-go-test/sigma/scheduledqueryrun"
	source "github.com/max-cape/stripe-go-test/source"
	subscription "github.com/max-cape/stripe-go-test/subscription"
	subscriptionitem "github.com/max-cape/stripe-go-test/subscriptionitem"
	subscriptionschedule "github.com/max-cape/stripe-go-test/subscriptionschedule"
	tax_calculation "github.com/max-cape/stripe-go-test/tax/calculation"
	tax_settings "github.com/max-cape/stripe-go-test/tax/settings"
	tax_transaction "github.com/max-cape/stripe-go-test/tax/transaction"
	taxcode "github.com/max-cape/stripe-go-test/taxcode"
	taxid "github.com/max-cape/stripe-go-test/taxid"
	taxrate "github.com/max-cape/stripe-go-test/taxrate"
	terminal_configuration "github.com/max-cape/stripe-go-test/terminal/configuration"
	terminal_connectiontoken "github.com/max-cape/stripe-go-test/terminal/connectiontoken"
	terminal_location "github.com/max-cape/stripe-go-test/terminal/location"
	terminal_reader "github.com/max-cape/stripe-go-test/terminal/reader"
	testhelpers_customer "github.com/max-cape/stripe-go-test/testhelpers/customer"
	testhelpers_issuing_authorization "github.com/max-cape/stripe-go-test/testhelpers/issuing/authorization"
	testhelpers_issuing_card "github.com/max-cape/stripe-go-test/testhelpers/issuing/card"
	testhelpers_issuing_personalizationdesign "github.com/max-cape/stripe-go-test/testhelpers/issuing/personalizationdesign"
	testhelpers_issuing_transaction "github.com/max-cape/stripe-go-test/testhelpers/issuing/transaction"
	testhelpers_refund "github.com/max-cape/stripe-go-test/testhelpers/refund"
	testhelpers_testclock "github.com/max-cape/stripe-go-test/testhelpers/testclock"
	testhelpers_treasury_inboundtransfer "github.com/max-cape/stripe-go-test/testhelpers/treasury/inboundtransfer"
	testhelpers_treasury_outboundtransfer "github.com/max-cape/stripe-go-test/testhelpers/treasury/outboundtransfer"
	testhelpers_treasury_receivedcredit "github.com/max-cape/stripe-go-test/testhelpers/treasury/receivedcredit"
	testhelpers_treasury_receiveddebit "github.com/max-cape/stripe-go-test/testhelpers/treasury/receiveddebit"
	. "github.com/max-cape/stripe-go-test/testing"
	token "github.com/max-cape/stripe-go-test/token"
	topup "github.com/max-cape/stripe-go-test/topup"
	transfer "github.com/max-cape/stripe-go-test/transfer"
	transferreversal "github.com/max-cape/stripe-go-test/transferreversal"
	treasury_creditreversal "github.com/max-cape/stripe-go-test/treasury/creditreversal"
	treasury_debitreversal "github.com/max-cape/stripe-go-test/treasury/debitreversal"
	treasury_financialaccount "github.com/max-cape/stripe-go-test/treasury/financialaccount"
	treasury_inboundtransfer "github.com/max-cape/stripe-go-test/treasury/inboundtransfer"
	treasury_outboundpayment "github.com/max-cape/stripe-go-test/treasury/outboundpayment"
	treasury_outboundtransfer "github.com/max-cape/stripe-go-test/treasury/outboundtransfer"
	treasury_receivedcredit "github.com/max-cape/stripe-go-test/treasury/receivedcredit"
	treasury_receiveddebit "github.com/max-cape/stripe-go-test/treasury/receiveddebit"
	treasury_transaction "github.com/max-cape/stripe-go-test/treasury/transaction"
	treasury_transactionentry "github.com/max-cape/stripe-go-test/treasury/transactionentry"
	webhookendpoint "github.com/max-cape/stripe-go-test/webhookendpoint"
	assert "github.com/stretchr/testify/require"
)

func TestAccountLinksPost(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, err := accountlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountLinksPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, err := sc.AccountLinks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsCapabilitiesGet(t *testing.T) {
	params := &stripe.CapabilityListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result := capability.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsCapabilitiesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CapabilityListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result := sc.Capabilities.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsCapabilitiesGet2(t *testing.T) {
	params := &stripe.CapabilityParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := capability.Get("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsCapabilitiesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CapabilityParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.Capabilities.Get("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsCapabilitiesPost(t *testing.T) {
	params := &stripe.CapabilityParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := capability.Update("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsCapabilitiesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CapabilityParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.Capabilities.Update("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsDelete(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.Del("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountParams{}
	result, err := sc.Accounts.Del("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsGet(t *testing.T) {
	params := &stripe.AccountListParams{}
	params.Limit = stripe.Int64(3)
	result := account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Accounts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsGet2(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.GetByID("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountParams{}
	result, err := sc.Accounts.GetByID("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsLoginLinksPost(t *testing.T) {
	params := &stripe.LoginLinkParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := loginlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsLoginLinksPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.LoginLinkParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.LoginLinks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsDelete(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := person.Del("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := sc.Persons.Del("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsGet(t *testing.T) {
	params := &stripe.PersonListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := person.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsPersonsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PersonListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.Persons.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountsPersonsGet2(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := person.Get("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := sc.Persons.Get("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsPost(t *testing.T) {
	params := &stripe.PersonParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := person.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PersonParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.Persons.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsPost2(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := person.Update("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPersonsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Persons.Update("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPost(t *testing.T) {
	params := &stripe.AccountParams{
		Type:    stripe.String(string(stripe.AccountTypeCustom)),
		Country: stripe.String("US"),
		Email:   stripe.String("jenny.rosen@example.com"),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
	}
	result, err := account.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountParams{
		Type:    stripe.String(string(stripe.AccountTypeCustom)),
		Country: stripe.String("US"),
		Email:   stripe.String("jenny.rosen@example.com"),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
	}
	result, err := sc.Accounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPost2(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Accounts.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsRejectPost(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, err := account.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsRejectPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, err := sc.Accounts.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesGet(t *testing.T) {
	params := &stripe.ApplicationFeeListParams{}
	params.Limit = stripe.Int64(3)
	result := applicationfee.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestApplicationFeesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ApplicationFeeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.ApplicationFees.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestApplicationFeesGet2(t *testing.T) {
	params := &stripe.ApplicationFeeParams{}
	result, err := applicationfee.Get("fee_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ApplicationFeeParams{}
	result, err := sc.ApplicationFees.Get("fee_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsGet(t *testing.T) {
	params := &stripe.FeeRefundListParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	params.Limit = stripe.Int64(3)
	result := feerefund.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestApplicationFeesRefundsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FeeRefundListParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	params.Limit = stripe.Int64(3)
	result := sc.FeeRefunds.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestApplicationFeesRefundsGet2(t *testing.T) {
	params := &stripe.FeeRefundParams{Fee: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := feerefund.Get("fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FeeRefundParams{Fee: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := sc.FeeRefunds.Get("fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsPost(t *testing.T) {
	params := &stripe.FeeRefundParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := feerefund.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FeeRefundParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := sc.FeeRefunds.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsPost2(t *testing.T) {
	params := &stripe.FeeRefundParams{Fee: stripe.String("fee_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := feerefund.Update("fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FeeRefundParams{Fee: stripe.String("fee_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := sc.FeeRefunds.Update("fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsDeletePost(t *testing.T) {
	params := &stripe.AppsSecretDeleteWhereParams{
		Name: stripe.String("my-api-key"),
		Scope: &stripe.AppsSecretDeleteWhereScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := apps_secret.DeleteWhere(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsDeletePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretDeleteWhereParams{
		Name: stripe.String("my-api-key"),
		Scope: &stripe.AppsSecretDeleteWhereScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := sc.AppsSecrets.DeleteWhere(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsFindGet(t *testing.T) {
	params := &stripe.AppsSecretFindParams{
		Name: stripe.String("sec_123"),
		Scope: &stripe.AppsSecretFindScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := apps_secret.Find(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsFindGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretFindParams{
		Name: stripe.String("sec_123"),
		Scope: &stripe.AppsSecretFindScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := sc.AppsSecrets.Find(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsGet(t *testing.T) {
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	params.Limit = stripe.Int64(2)
	result := apps_secret.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.AppsSecrets.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsGet2(t *testing.T) {
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	params.Limit = stripe.Int64(2)
	result := apps_secret.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.AppsSecrets.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsPost(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("sec_123"),
		Payload: stripe.String("very secret string"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := apps_secret.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("sec_123"),
		Payload: stripe.String("very secret string"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := sc.AppsSecrets.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPost2(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("my-api-key"),
		Payload: stripe.String("secret_key_xxxxxx"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := apps_secret.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("my-api-key"),
		Payload: stripe.String("secret_key_xxxxxx"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, err := sc.AppsSecrets.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBalanceTransactionsGet(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBalanceTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.BalanceTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBalanceTransactionsGet2(t *testing.T) {
	params := &stripe.BalanceTransactionParams{}
	result, err := balancetransaction.Get("txn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBalanceTransactionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BalanceTransactionParams{}
	result, err := sc.BalanceTransactions.Get("txn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsGet(t *testing.T) {
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := billingportal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBillingPortalConfigurationsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.BillingPortalConfigurations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBillingPortalConfigurationsGet2(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{}
	result, err := billingportal_configuration.Get("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BillingPortalConfigurationParams{}
	result, err := sc.BillingPortalConfigurations.Get(
		"bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsPost(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		Features: &stripe.BillingPortalConfigurationFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail)),
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID)),
				},
				Enabled: stripe.Bool(true),
			},
			InvoiceHistory: &stripe.BillingPortalConfigurationFeaturesInvoiceHistoryParams{
				Enabled: stripe.Bool(true),
			},
		},
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := billingportal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BillingPortalConfigurationParams{
		Features: &stripe.BillingPortalConfigurationFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail)),
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID)),
				},
				Enabled: stripe.Bool(true),
			},
			InvoiceHistory: &stripe.BillingPortalConfigurationFeaturesInvoiceHistoryParams{
				Enabled: stripe.Bool(true),
			},
		},
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := sc.BillingPortalConfigurations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsPost2(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := billingportal_configuration.Update(
		"bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := sc.BillingPortalConfigurations.Update(
		"bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalSessionsPost(t *testing.T) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, err := billingportal_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, err := sc.BillingPortalSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesCapturePost(t *testing.T) {
	params := &stripe.ChargeCaptureParams{}
	result, err := charge.Capture("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesCapturePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeCaptureParams{}
	result, err := sc.Charges.Capture("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesGet(t *testing.T) {
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(3)
	result := charge.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestChargesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Charges.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestChargesGet2(t *testing.T) {
	params := &stripe.ChargeParams{}
	result, err := charge.Get("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeParams{}
	result, err := sc.Charges.Get("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPost(t *testing.T) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := charge.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := sc.Charges.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPost2(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := charge.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Charges.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesSearchGet(t *testing.T) {
	params := &stripe.ChargeSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "amount>999 AND metadata['order_id']:'6735'",
		},
	}
	result := charge.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestChargesSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ChargeSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "amount>999 AND metadata['order_id']:'6735'",
		},
	}
	result := sc.Charges.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionsExpirePost(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("sess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsExpirePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := sc.CheckoutSessions.Expire("sess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsExpirePost2(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsExpirePost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := sc.CheckoutSessions.Expire("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsGet(t *testing.T) {
	params := &stripe.CheckoutSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := checkout_session.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.CheckoutSessions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionsGet2(t *testing.T) {
	params := &stripe.CheckoutSessionParams{}
	result, err := checkout_session.Get("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionParams{}
	result, err := sc.CheckoutSessions.Get("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsLineItemsGet(t *testing.T) {
	params := &stripe.CheckoutSessionListLineItemsParams{
		Session: stripe.String("sess_xyz"),
	}
	result := checkout_session.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionsLineItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionListLineItemsParams{
		Session: stripe.String("sess_xyz"),
	}
	result := sc.CheckoutSessions.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionsPost(t *testing.T) {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		ShippingOptions: []*stripe.CheckoutSessionShippingOptionParams{
			{
				ShippingRate: stripe.String("shr_standard"),
			},
			{
				ShippingRateData: &stripe.CheckoutSessionShippingOptionShippingRateDataParams{
					DisplayName: stripe.String("Standard"),
					DeliveryEstimate: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams{
						Minimum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(5),
						},
						Maximum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(7),
						},
					},
				},
			},
		},
	}
	result, err := checkout_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		ShippingOptions: []*stripe.CheckoutSessionShippingOptionParams{
			{
				ShippingRate: stripe.String("shr_standard"),
			},
			{
				ShippingRateData: &stripe.CheckoutSessionShippingOptionShippingRateDataParams{
					DisplayName: stripe.String("Standard"),
					DeliveryEstimate: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams{
						Minimum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(5),
						},
						Maximum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(7),
						},
					},
				},
			},
		},
	}
	result, err := sc.CheckoutSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsPost2(t *testing.T) {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	result, err := checkout_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	result, err := sc.CheckoutSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCountrySpecsGet(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Limit = stripe.Int64(3)
	result := countryspec.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCountrySpecsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CountrySpecListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.CountrySpecs.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCountrySpecsGet2(t *testing.T) {
	params := &stripe.CountrySpecParams{}
	result, err := countryspec.Get("US", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCountrySpecsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CountrySpecParams{}
	result, err := sc.CountrySpecs.Get("US", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsDelete(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Del("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponParams{}
	result, err := sc.Coupons.Del("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsGet(t *testing.T) {
	params := &stripe.CouponListParams{}
	params.Limit = stripe.Int64(3)
	result := coupon.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCouponsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Coupons.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCouponsGet2(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Get("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponParams{}
	result, err := sc.Coupons.Get("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPost(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff: stripe.Float64(25.5),
		Duration:   stripe.String(string(stripe.CouponDurationOnce)),
	}
	result, err := coupon.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponParams{
		PercentOff: stripe.Float64(25.5),
		Duration:   stripe.String(string(stripe.CouponDurationOnce)),
	}
	result, err := sc.Coupons.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPost2(t *testing.T) {
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, err := coupon.Update("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Coupons.Update("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerSessionsPost(t *testing.T) {
	params := &stripe.CustomerSessionParams{
		Customer: stripe.String("cus_123"),
		Components: &stripe.CustomerSessionComponentsParams{
			BuyButton: &stripe.CustomerSessionComponentsBuyButtonParams{
				Enabled: stripe.Bool(true),
			},
		},
	}
	result, err := customersession.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerSessionParams{
		Customer: stripe.String("cus_123"),
		Components: &stripe.CustomerSessionComponentsParams{
			BuyButton: &stripe.CustomerSessionComponentsBuyButtonParams{
				Enabled: stripe.Bool(true),
			},
		},
	}
	result, err := sc.CustomerSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsGet(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := customerbalancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersBalanceTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.CustomerBalanceTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersBalanceTransactionsGet2(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := customerbalancetransaction.Get("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.CustomerBalanceTransactions.Get(
		"cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPost(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := customerbalancetransaction.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.CustomerBalanceTransactions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPost2(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := customerbalancetransaction.Update(
		"cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.CustomerBalanceTransactions.Update(
		"cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalanceGet(t *testing.T) {
	params := &stripe.CashBalanceParams{Customer: stripe.String("cus_123")}
	result, err := cashbalance.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalanceGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CashBalanceParams{Customer: stripe.String("cus_123")}
	result, err := sc.CashBalances.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalancePost(t *testing.T) {
	params := &stripe.CashBalanceParams{
		Settings: &stripe.CashBalanceSettingsParams{
			ReconciliationMode: stripe.String(string(stripe.CashBalanceSettingsReconciliationModeManual)),
		},
		Customer: stripe.String("cus_123"),
	}
	result, err := cashbalance.Update(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalancePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CashBalanceParams{
		Settings: &stripe.CashBalanceSettingsParams{
			ReconciliationMode: stripe.String(string(stripe.CashBalanceSettingsReconciliationModeManual)),
		},
		Customer: stripe.String("cus_123"),
	}
	result, err := sc.CashBalances.Update(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalanceTransactionsGet(t *testing.T) {
	params := &stripe.CustomerCashBalanceTransactionListParams{
		Customer: stripe.String("cus_123"),
	}
	params.Limit = stripe.Int64(3)
	result := customercashbalancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersCashBalanceTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerCashBalanceTransactionListParams{
		Customer: stripe.String("cus_123"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.CustomerCashBalanceTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersDelete(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Del("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerParams{}
	result, err := sc.Customers.Del("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersFundingInstructionsPost(t *testing.T) {
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		FundingType: stripe.String("bank_transfer"),
	}
	result, err := customer.CreateFundingInstructions("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersFundingInstructionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		FundingType: stripe.String("bank_transfer"),
	}
	result, err := sc.Customers.CreateFundingInstructions("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersGet(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := customer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Customers.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersGet2(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := customer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Customers.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersGet3(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Get("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerParams{}
	result, err := sc.Customers.Get("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersPaymentMethodsGet(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xyz"),
	}
	result := customer.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersPaymentMethodsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xyz"),
	}
	result := sc.Customers.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersPaymentMethodsGet2(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result := customer.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersPaymentMethodsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result := sc.Customers.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersPost(t *testing.T) {
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := customer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := sc.Customers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersPost2(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, err := customer.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Customers.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSearchGet(t *testing.T) {
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := customer.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := sc.Customers.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSearchGet2(t *testing.T) {
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := customer.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSearchGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := sc.Customers.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSourcesDelete(t *testing.T) {
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := card.Del("ba_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := sc.Cards.Del("ba_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesDelete2(t *testing.T) {
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := card.Del("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesDelete2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := sc.Cards.Del("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesGet(t *testing.T) {
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("bank_account"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := paymentsource.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSourcesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("bank_account"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.PaymentSources.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSourcesGet2(t *testing.T) {
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := paymentsource.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSourcesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.PaymentSources.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersSourcesGet3(t *testing.T) {
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentsource.Get("ba_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.PaymentSources.Get("ba_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesGet4(t *testing.T) {
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentsource.Get("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesGet4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.PaymentSources.Get("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesPost(t *testing.T) {
	params := &stripe.CardParams{
		AccountHolderName: stripe.String("Kamil"),
		Customer:          stripe.String("cus_123"),
	}
	result, err := card.Update("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CardParams{
		AccountHolderName: stripe.String("Kamil"),
		Customer:          stripe.String("cus_123"),
	}
	result, err := sc.Cards.Update("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesPost5(t *testing.T) {
	params := &stripe.CardParams{
		Name:     stripe.String("Jenny Rosen"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := card.Update("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesPost5Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CardParams{
		Name:     stripe.String("Jenny Rosen"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.Cards.Update("card_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsDelete(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Del("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := sc.TaxIDs.Del("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsGet(t *testing.T) {
	params := &stripe.TaxIDListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := taxid.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersTaxIdsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TaxIDs.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomersTaxIdsGet2(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Get("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := sc.TaxIDs.Get("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsPost(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:     stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value:    stripe.String("DE123456789"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := taxid.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{
		Type:     stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value:    stripe.String("DE123456789"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.TaxIDs.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesClosePost(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, err := dispute.Close("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesClosePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.DisputeParams{}
	result, err := sc.Disputes.Close("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesGet(t *testing.T) {
	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := dispute.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestDisputesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Disputes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestDisputesGet2(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, err := dispute.Get("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.DisputeParams{}
	result, err := sc.Disputes.Get("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesPost(t *testing.T) {
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := dispute.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Disputes.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestEventsGet(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Limit = stripe.Int64(3)
	result := event.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestEventsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.EventListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Events.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestEventsGet2(t *testing.T) {
	params := &stripe.EventParams{}
	result, err := event.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestEventsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.EventParams{}
	result, err := sc.Events.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := sc.FinancialConnectionsAccounts.Disconnect("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPost2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect(
		"fca_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := sc.FinancialConnectionsAccounts.Disconnect(
		"fca_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListParams{}
	result := financialconnections_account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountListParams{}
	result := sc.FinancialConnectionsAccounts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := financialconnections_account.GetByID("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := sc.FinancialConnectionsAccounts.GetByID("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsGet3(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListParams{
		AccountHolder: &stripe.FinancialConnectionsAccountListAccountHolderParams{
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
	}
	result := financialconnections_account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountListParams{
		AccountHolder: &stripe.FinancialConnectionsAccountListAccountHolderParams{
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
	}
	result := sc.FinancialConnectionsAccounts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsGet4(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := financialconnections_account.GetByID(
		"fca_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsGet4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := sc.FinancialConnectionsAccounts.GetByID(
		"fca_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsOwnersGet(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xyz"),
		Account:   stripe.String("fca_xyz"),
	}
	result := financialconnections_account.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsOwnersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xyz"),
		Account:   stripe.String("fca_xyz"),
	}
	result := sc.FinancialConnectionsAccounts.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsOwnersGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xxxxxxxxxxxxx"),
		Account:   stripe.String("fca_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := financialconnections_account.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsOwnersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xxxxxxxxxxxxx"),
		Account:   stripe.String("fca_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.FinancialConnectionsAccounts.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsRefreshPost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountRefreshParams{
		Features: []*string{stripe.String("balance")},
	}
	result, err := financialconnections_account.Refresh("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsRefreshPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountRefreshParams{
		Features: []*string{stripe.String("balance")},
	}
	result, err := sc.FinancialConnectionsAccounts.Refresh("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsSubscribePost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountSubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := financialconnections_account.Subscribe("fa_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsSubscribePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountSubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := sc.FinancialConnectionsAccounts.Subscribe("fa_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsUnsubscribePost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountUnsubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := financialconnections_account.Unsubscribe("fa_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsUnsubscribePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsAccountUnsubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := sc.FinancialConnectionsAccounts.Unsubscribe("fa_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get("fcsess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := sc.FinancialConnectionsSessions.Get("fcsess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get(
		"fcsess_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := sc.FinancialConnectionsSessions.Get(
		"fcsess_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(string(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer)),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionBalances)),
		},
	}
	result, err := financialconnections_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(string(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer)),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionBalances)),
		},
	}
	result, err := sc.FinancialConnectionsSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(string(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer)),
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
		Permissions: []*string{
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionPaymentMethod)),
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionBalances)),
		},
		Filters: &stripe.FinancialConnectionsSessionFiltersParams{
			Countries: []*string{stripe.String("US")},
		},
	}
	result, err := financialconnections_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(string(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer)),
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
		Permissions: []*string{
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionPaymentMethod)),
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionBalances)),
		},
		Filters: &stripe.FinancialConnectionsSessionFiltersParams{
			Countries: []*string{stripe.String("US")},
		},
	}
	result, err := sc.FinancialConnectionsSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsTransactionsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsTransactionParams{}
	result, err := financialconnections_transaction.Get("tr_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsTransactionParams{}
	result, err := sc.FinancialConnectionsTransactions.Get("tr_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsTransactionsGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsTransactionListParams{
		Account: stripe.String("fca_xyz"),
	}
	result := financialconnections_transaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsTransactionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.FinancialConnectionsTransactionListParams{
		Account: stripe.String("fca_xyz"),
	}
	result := sc.FinancialConnectionsTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationReportsGet(t *testing.T) {
	params := &stripe.IdentityVerificationReportListParams{}
	params.Limit = stripe.Int64(3)
	result := identity_verificationreport.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationReportsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationReportListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IdentityVerificationReports.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationReportsGet2(t *testing.T) {
	params := &stripe.IdentityVerificationReportParams{}
	result, err := identity_verificationreport.Get("vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationReportsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationReportParams{}
	result, err := sc.IdentityVerificationReports.Get("vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsCancelPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := identity_verificationsession.Cancel(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := sc.IdentityVerificationSessions.Cancel(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsGet(t *testing.T) {
	params := &stripe.IdentityVerificationSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := identity_verificationsession.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationSessionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IdentityVerificationSessions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationSessionsGet2(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{}
	result, err := identity_verificationsession.Get("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionParams{}
	result, err := sc.IdentityVerificationSessions.Get(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeDocument)),
	}
	result, err := identity_verificationsession.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeDocument)),
	}
	result, err := sc.IdentityVerificationSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost2(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeIDNumber)),
	}
	result, err := identity_verificationsession.Update(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeIDNumber)),
	}
	result, err := sc.IdentityVerificationSessions.Update(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsRedactPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := identity_verificationsession.Redact(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsRedactPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := sc.IdentityVerificationSessions.Redact(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsDelete(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, err := invoiceitem.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceItemParams{}
	result, err := sc.InvoiceItems.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsPost(t *testing.T) {
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := invoiceitem.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.InvoiceItems.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsPost2(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := invoiceitem.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.InvoiceItems.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesDelete(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceParams{}
	result, err := sc.Invoices.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesFinalizePost(t *testing.T) {
	params := &stripe.InvoiceFinalizeInvoiceParams{}
	result, err := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesFinalizePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceFinalizeInvoiceParams{}
	result, err := sc.Invoices.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesGet(t *testing.T) {
	params := &stripe.InvoiceListParams{}
	params.Limit = stripe.Int64(3)
	result := invoice.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoicesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Invoices.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoicesGet2(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceParams{}
	result, err := sc.Invoices.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesGet3(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddExpand("customer")
	result, err := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceParams{}
	params.AddExpand("customer")
	result, err := sc.Invoices.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPayPost(t *testing.T) {
	params := &stripe.InvoicePayParams{}
	result, err := invoice.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPayPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoicePayParams{}
	result, err := sc.Invoices.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPost(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := invoice.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := sc.Invoices.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPost2(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := invoice.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Invoices.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesSearchGet(t *testing.T) {
	params := &stripe.InvoiceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "total>999 AND metadata['order_id']:'6735'",
		},
	}
	result := invoice.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoicesSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "total>999 AND metadata['order_id']:'6735'",
		},
	}
	result := sc.Invoices.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoicesSendPost(t *testing.T) {
	params := &stripe.InvoiceSendInvoiceParams{}
	result, err := invoice.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesSendPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceSendInvoiceParams{}
	result, err := sc.Invoices.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesVoidPost(t *testing.T) {
	params := &stripe.InvoiceVoidInvoiceParams{}
	result, err := invoice.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesVoidPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.InvoiceVoidInvoiceParams{}
	result, err := sc.Invoices.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsApprovePost(t *testing.T) {
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, err := issuing_authorization.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsApprovePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, err := sc.IssuingAuthorizations.Approve(
		"iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsDeclinePost(t *testing.T) {
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, err := issuing_authorization.Decline("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsDeclinePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, err := sc.IssuingAuthorizations.Decline(
		"iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsGet(t *testing.T) {
	params := &stripe.IssuingAuthorizationListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_authorization.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingAuthorizationsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingAuthorizationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IssuingAuthorizations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingAuthorizationsGet2(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	result, err := issuing_authorization.Get("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingAuthorizationParams{}
	result, err := sc.IssuingAuthorizations.Get("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsPost(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_authorization.Update("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.IssuingAuthorizations.Update("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersGet(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_cardholder.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardholdersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardholderListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IssuingCardholders.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardholdersGet2(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	result, err := issuing_cardholder.Get("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardholderParams{}
	result, err := sc.IssuingCardholders.Get("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersPost(t *testing.T) {
	params := &stripe.IssuingCardholderParams{
		Type:        stripe.String(string(stripe.IssuingCardholderTypeIndividual)),
		Name:        stripe.String("Jenny Rosen"),
		Email:       stripe.String("jenny.rosen@example.com"),
		PhoneNumber: stripe.String("+18888675309"),
		Billing: &stripe.IssuingCardholderBillingParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("1234 Main Street"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
				PostalCode: stripe.String("94111"),
			},
		},
	}
	result, err := issuing_cardholder.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardholderParams{
		Type:        stripe.String(string(stripe.IssuingCardholderTypeIndividual)),
		Name:        stripe.String("Jenny Rosen"),
		Email:       stripe.String("jenny.rosen@example.com"),
		PhoneNumber: stripe.String("+18888675309"),
		Billing: &stripe.IssuingCardholderBillingParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("1234 Main Street"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
				PostalCode: stripe.String("94111"),
			},
		},
	}
	result, err := sc.IssuingCardholders.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersPost2(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_cardholder.Update("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.IssuingCardholders.Update("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsGet(t *testing.T) {
	params := &stripe.IssuingCardListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_card.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IssuingCards.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardsGet2(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	result, err := issuing_card.Get("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardParams{}
	result, err := sc.IssuingCards.Get("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPost(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String(string(stripe.IssuingCardTypeVirtual)),
	}
	result, err := issuing_card.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String(string(stripe.IssuingCardTypeVirtual)),
	}
	result, err := sc.IssuingCards.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPost2(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_card.Update("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.IssuingCards.Update("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesGet(t *testing.T) {
	params := &stripe.IssuingDisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_dispute.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingDisputesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingDisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IssuingDisputes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingDisputesGet2(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	result, err := issuing_dispute.Get("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingDisputeParams{}
	result, err := sc.IssuingDisputes.Get("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesPost(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonFraudulent)),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, err := issuing_dispute.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonFraudulent)),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, err := sc.IssuingDisputes.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesSubmitPost(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	result, err := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesSubmitPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingDisputeSubmitParams{}
	result, err := sc.IssuingDisputes.Submit("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsGet(t *testing.T) {
	params := &stripe.IssuingPersonalizationDesignListParams{}
	result := issuing_personalizationdesign.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingPersonalizationDesignsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPersonalizationDesignListParams{}
	result := sc.IssuingPersonalizationDesigns.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingPersonalizationDesignsGet2(t *testing.T) {
	params := &stripe.IssuingPersonalizationDesignParams{}
	result, err := issuing_personalizationdesign.Get("pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPersonalizationDesignParams{}
	result, err := sc.IssuingPersonalizationDesigns.Get("pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsPost(t *testing.T) {
	params := &stripe.IssuingPersonalizationDesignParams{
		PhysicalBundle: stripe.String("pb_xyz"),
	}
	result, err := issuing_personalizationdesign.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPersonalizationDesignParams{
		PhysicalBundle: stripe.String("pb_xyz"),
	}
	result, err := sc.IssuingPersonalizationDesigns.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsPost2(t *testing.T) {
	params := &stripe.IssuingPersonalizationDesignParams{}
	result, err := issuing_personalizationdesign.Update("pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPersonalizationDesignsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPersonalizationDesignParams{}
	result, err := sc.IssuingPersonalizationDesigns.Update("pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPhysicalBundlesGet(t *testing.T) {
	params := &stripe.IssuingPhysicalBundleListParams{}
	result := issuing_physicalbundle.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingPhysicalBundlesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPhysicalBundleListParams{}
	result := sc.IssuingPhysicalBundles.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingPhysicalBundlesGet2(t *testing.T) {
	params := &stripe.IssuingPhysicalBundleParams{}
	result, err := issuing_physicalbundle.Get("pb_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingPhysicalBundlesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingPhysicalBundleParams{}
	result, err := sc.IssuingPhysicalBundles.Get("pb_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingTransactionsGet(t *testing.T) {
	params := &stripe.IssuingTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_transaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.IssuingTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingTransactionsGet2(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	result, err := issuing_transaction.Get("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingTransactionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingTransactionParams{}
	result, err := sc.IssuingTransactions.Get("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingTransactionsPost(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_transaction.Update("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingTransactionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.IssuingTransactions.Update("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestMandatesGet(t *testing.T) {
	params := &stripe.MandateParams{}
	result, err := mandate.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestMandatesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.MandateParams{}
	result, err := sc.Mandates.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsApplyCustomerBalancePost(t *testing.T) {
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, err := paymentintent.ApplyCustomerBalance("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsApplyCustomerBalancePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, err := sc.PaymentIntents.ApplyCustomerBalance(
		"pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCancelPost(t *testing.T) {
	params := &stripe.PaymentIntentCancelParams{}
	result, err := paymentintent.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentCancelParams{}
	result, err := sc.PaymentIntents.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCapturePost(t *testing.T) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := paymentintent.Capture("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCapturePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := sc.PaymentIntents.Capture("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsConfirmPost(t *testing.T) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := paymentintent.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsConfirmPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := sc.PaymentIntents.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsGet(t *testing.T) {
	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := paymentintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.PaymentIntents.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentsGet2(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	result, err := paymentintent.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentParams{}
	result, err := sc.PaymentIntents.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsIncrementAuthorizationPost(t *testing.T) {
	params := &stripe.PaymentIntentIncrementAuthorizationParams{
		Amount: stripe.Int64(2099),
	}
	result, err := paymentintent.IncrementAuthorization(
		"pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsIncrementAuthorizationPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentIncrementAuthorizationParams{
		Amount: stripe.Int64(2099),
	}
	result, err := sc.PaymentIntents.IncrementAuthorization(
		"pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := paymentintent.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost2(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := paymentintent.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost3(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.PaymentIntents.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost4(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(200),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("p24"),
			P24:  &stripe.PaymentMethodP24Params{Bank: stripe.String("blik")},
		},
	}
	result, err := paymentintent.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(200),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("p24"),
			P24:  &stripe.PaymentMethodP24Params{Bank: stripe.String("blik")},
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsSearchGet(t *testing.T) {
	params := &stripe.PaymentIntentSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'succeeded' AND metadata['order_id']:'6735'",
		},
	}
	result := paymentintent.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentsSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'succeeded' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.PaymentIntents.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentsVerifyMicrodepositsPost(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, err := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsVerifyMicrodepositsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, err := sc.PaymentIntents.VerifyMicrodeposits(
		"pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsVerifyMicrodepositsPost2(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsVerifyMicrodepositsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := sc.PaymentIntents.VerifyMicrodeposits(
		"pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksGet(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("pl_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkParams{}
	result, err := sc.PaymentLinks.Get("pl_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksGet2(t *testing.T) {
	params := &stripe.PaymentLinkListParams{}
	params.Limit = stripe.Int64(3)
	result := paymentlink.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinksGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.PaymentLinks.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinksGet3(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkParams{}
	result, err := sc.PaymentLinks.Get("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksLineItemsGet(t *testing.T) {
	params := &stripe.PaymentLinkListLineItemsParams{
		PaymentLink: stripe.String("pl_xyz"),
	}
	result := paymentlink.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinksLineItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkListLineItemsParams{
		PaymentLink: stripe.String("pl_xyz"),
	}
	result := sc.PaymentLinks.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinksPost(t *testing.T) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := paymentlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := sc.PaymentLinks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksPost2(t *testing.T) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := paymentlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := sc.PaymentLinks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksPost3(t *testing.T) {
	params := &stripe.PaymentLinkParams{Active: stripe.Bool(false)}
	result, err := paymentlink.Update("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinksPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentLinkParams{Active: stripe.Bool(false)}
	result, err := sc.PaymentLinks.Update("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsGet(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationListParams{
		Application: stripe.String("foo"),
	}
	result := paymentmethodconfiguration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodConfigurationsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodConfigurationListParams{
		Application: stripe.String("foo"),
	}
	result := sc.PaymentMethodConfigurations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodConfigurationsGet2(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{}
	result, err := paymentmethodconfiguration.Get("foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodConfigurationParams{}
	result, err := sc.PaymentMethodConfigurations.Get("foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceNone)),
			},
		},
		Affirm: &stripe.PaymentMethodConfigurationAffirmParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationAffirmDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationAffirmDisplayPreferencePreferenceNone)),
			},
		},
	}
	result, err := paymentmethodconfiguration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceNone)),
			},
		},
		Affirm: &stripe.PaymentMethodConfigurationAffirmParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationAffirmDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationAffirmDisplayPreferencePreferenceNone)),
			},
		},
	}
	result, err := sc.PaymentMethodConfigurations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost2(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceOn)),
			},
		},
	}
	result, err := paymentmethodconfiguration.Update("foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(string(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceOn)),
			},
		},
	}
	result, err := sc.PaymentMethodConfigurations.Update("foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsAttachPost(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsAttachPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.PaymentMethods.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsDetachPost(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	result, err := paymentmethod.Detach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsDetachPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodDetachParams{}
	result, err := sc.PaymentMethods.Detach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsGet(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	result := sc.PaymentMethods.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodsGet2(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	result, err := paymentmethod.Get("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodParams{}
	result, err := sc.PaymentMethods.Get("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsPost(t *testing.T) {
	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(8),
			ExpYear:  stripe.Int64(2024),
			CVC:      stripe.String("314"),
		},
	}
	result, err := paymentmethod.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(8),
			ExpYear:  stripe.Int64(2024),
			CVC:      stripe.String("314"),
		},
	}
	result, err := sc.PaymentMethods.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsPost2(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.PaymentMethods.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsCancelPost(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Cancel("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutParams{}
	result, err := sc.Payouts.Cancel("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsGet(t *testing.T) {
	params := &stripe.PayoutListParams{}
	params.Limit = stripe.Int64(3)
	result := payout.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPayoutsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Payouts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPayoutsGet2(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Get("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutParams{}
	result, err := sc.Payouts.Get("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPost(t *testing.T) {
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := payout.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := sc.Payouts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPost2(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, err := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Payouts.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsReversePost(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	result, err := payout.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsReversePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutReverseParams{}
	result, err := sc.Payouts.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansDelete(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Del("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PlanParams{}
	result, err := sc.Plans.Del("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansGet(t *testing.T) {
	params := &stripe.PlanListParams{}
	params.Limit = stripe.Int64(3)
	result := plan.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPlansGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PlanListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Plans.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPlansGet2(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PlanParams{}
	result, err := sc.Plans.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPost(t *testing.T) {
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Interval: stripe.String(string(stripe.PlanIntervalMonth)),
		Product:  &stripe.PlanProductParams{Name: stripe.String("My product")},
	}
	result, err := plan.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Interval: stripe.String(string(stripe.PlanIntervalMonth)),
		Product:  &stripe.PlanProductParams{Name: stripe.String("My product")},
	}
	result, err := sc.Plans.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPost2(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, err := plan.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Plans.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesGet(t *testing.T) {
	params := &stripe.PriceListParams{}
	params.Limit = stripe.Int64(3)
	result := price.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPricesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Prices.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPricesGet2(t *testing.T) {
	params := &stripe.PriceParams{}
	result, err := price.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceParams{}
	result, err := sc.Prices.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		CurrencyOptions: map[string]*stripe.PriceCurrencyOptionsParams{
			"uah": {UnitAmount: stripe.Int64(5000)},
			"eur": {UnitAmount: stripe.Int64(1800)},
		},
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := price.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		CurrencyOptions: map[string]*stripe.PriceCurrencyOptionsParams{
			"uah": {UnitAmount: stripe.Int64(5000)},
			"eur": {UnitAmount: stripe.Int64(1800)},
		},
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.Prices.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost2(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := price.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.Prices.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost3(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := price.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Prices.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesSearchGet(t *testing.T) {
	params := &stripe.PriceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := price.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPricesSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PriceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.Prices.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductsDelete(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Del("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductParams{}
	result, err := sc.Products.Del("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsGet(t *testing.T) {
	params := &stripe.ProductListParams{}
	params.Limit = stripe.Int64(3)
	result := product.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Products.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductsGet2(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductParams{}
	result, err := sc.Products.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsPost(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, err := product.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, err := sc.Products.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsPost2(t *testing.T) {
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, err := product.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Products.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsSearchGet(t *testing.T) {
	params := &stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := product.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductsSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.Products.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodesGet(t *testing.T) {
	params := &stripe.PromotionCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := promotioncode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PromotionCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.PromotionCodes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodesGet2(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	result, err := promotioncode.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PromotionCodeParams{}
	result, err := sc.PromotionCodes.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPost(t *testing.T) {
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("Z4OV52SU")}
	result, err := promotioncode.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("Z4OV52SU")}
	result, err := sc.PromotionCodes.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPost2(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.PromotionCodes.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesAcceptPost(t *testing.T) {
	params := &stripe.QuoteAcceptParams{}
	result, err := quote.Accept("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesAcceptPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteAcceptParams{}
	result, err := sc.Quotes.Accept("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesCancelPost(t *testing.T) {
	params := &stripe.QuoteCancelParams{}
	result, err := quote.Cancel("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteCancelParams{}
	result, err := sc.Quotes.Cancel("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesFinalizePost(t *testing.T) {
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, err := quote.FinalizeQuote("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesFinalizePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, err := sc.Quotes.FinalizeQuote("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesGet(t *testing.T) {
	params := &stripe.QuoteListParams{}
	params.Limit = stripe.Int64(3)
	result := quote.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Quotes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesGet2(t *testing.T) {
	params := &stripe.QuoteParams{}
	result, err := quote.Get("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteParams{}
	result, err := sc.Quotes.Get("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesLineItemsGet(t *testing.T) {
	params := &stripe.QuoteListLineItemsParams{
		Quote: stripe.String("qt_xxxxxxxxxxxxx"),
	}
	result := quote.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesLineItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteListLineItemsParams{
		Quote: stripe.String("qt_xxxxxxxxxxxxx"),
	}
	result := sc.Quotes.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesPdfGet(t *testing.T) {
	params := &stripe.QuotePDFParams{}
	result, err := quote.PDF("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPdfGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuotePDFParams{}
	result, err := sc.Quotes.PDF("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPost(t *testing.T) {
	params := &stripe.QuoteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		LineItems: []*stripe.QuoteLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
	}
	result, err := quote.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		LineItems: []*stripe.QuoteLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
	}
	result, err := sc.Quotes.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPost2(t *testing.T) {
	params := &stripe.QuoteParams{}
	params.AddMetadata("order_id", "6735")
	result, err := quote.Update("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Quotes.Update("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarEarlyFraudWarningsGet(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Limit = stripe.Int64(3)
	result := radar_earlyfraudwarning.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarEarlyFraudWarningsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.RadarEarlyFraudWarnings.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarEarlyFraudWarningsGet2(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningParams{}
	result, err := radar_earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarEarlyFraudWarningsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarEarlyFraudWarningParams{}
	result, err := sc.RadarEarlyFraudWarnings.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsDelete(t *testing.T) {
	params := &stripe.RadarValueListItemParams{}
	result, err := radar_valuelistitem.Del("rsli_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListItemParams{}
	result, err := sc.RadarValueListItems.Del("rsli_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsGet(t *testing.T) {
	params := &stripe.RadarValueListItemListParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := radar_valuelistitem.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarValueListItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListItemListParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.RadarValueListItems.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarValueListItemsGet2(t *testing.T) {
	params := &stripe.RadarValueListItemParams{}
	result, err := radar_valuelistitem.Get("rsli_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListItemParams{}
	result, err := sc.RadarValueListItems.Get("rsli_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsPost(t *testing.T) {
	params := &stripe.RadarValueListItemParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
		Value:     stripe.String("1.2.3.4"),
	}
	result, err := radar_valuelistitem.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListItemParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
		Value:     stripe.String("1.2.3.4"),
	}
	result, err := sc.RadarValueListItems.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsDelete(t *testing.T) {
	params := &stripe.RadarValueListParams{}
	result, err := radar_valuelist.Del("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListParams{}
	result, err := sc.RadarValueLists.Del("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsGet(t *testing.T) {
	params := &stripe.RadarValueListListParams{}
	params.Limit = stripe.Int64(3)
	result := radar_valuelist.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarValueListsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.RadarValueLists.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarValueListsGet2(t *testing.T) {
	params := &stripe.RadarValueListParams{}
	result, err := radar_valuelist.Get("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListParams{}
	result, err := sc.RadarValueLists.Get("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPost(t *testing.T) {
	params := &stripe.RadarValueListParams{
		Alias:    stripe.String("custom_ip_xxxxxxxxxxxxx"),
		Name:     stripe.String("Custom IP Blocklist"),
		ItemType: stripe.String(string(stripe.RadarValueListItemTypeIPAddress)),
	}
	result, err := radar_valuelist.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListParams{
		Alias:    stripe.String("custom_ip_xxxxxxxxxxxxx"),
		Name:     stripe.String("Custom IP Blocklist"),
		ItemType: stripe.String(string(stripe.RadarValueListItemTypeIPAddress)),
	}
	result, err := sc.RadarValueLists.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPost2(t *testing.T) {
	params := &stripe.RadarValueListParams{
		Name: stripe.String("Updated IP Block List"),
	}
	result, err := radar_valuelist.Update("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RadarValueListParams{
		Name: stripe.String("Updated IP Block List"),
	}
	result, err := sc.RadarValueLists.Update("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsCancelPost(t *testing.T) {
	params := &stripe.RefundCancelParams{}
	result, err := refund.Cancel("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RefundCancelParams{}
	result, err := sc.Refunds.Cancel("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsGet(t *testing.T) {
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(3)
	result := refund.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRefundsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Refunds.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRefundsGet2(t *testing.T) {
	params := &stripe.RefundParams{}
	result, err := refund.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RefundParams{}
	result, err := sc.Refunds.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsPost(t *testing.T) {
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, err := refund.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, err := sc.Refunds.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsPost2(t *testing.T) {
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, err := refund.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Refunds.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportRunsGet(t *testing.T) {
	params := &stripe.ReportingReportRunListParams{}
	params.Limit = stripe.Int64(3)
	result := reporting_reportrun.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportRunsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReportingReportRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.ReportingReportRuns.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportRunsGet2(t *testing.T) {
	params := &stripe.ReportingReportRunParams{}
	result, err := reporting_reportrun.Get("frr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportRunsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReportingReportRunParams{}
	result, err := sc.ReportingReportRuns.Get("frr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportRunsPost(t *testing.T) {
	params := &stripe.ReportingReportRunParams{
		ReportType: stripe.String("balance.summary.1"),
		Parameters: &stripe.ReportingReportRunParametersParams{
			IntervalStart: stripe.Int64(1522540800),
			IntervalEnd:   stripe.Int64(1525132800),
		},
	}
	result, err := reporting_reportrun.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportRunsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReportingReportRunParams{
		ReportType: stripe.String("balance.summary.1"),
		Parameters: &stripe.ReportingReportRunParametersParams{
			IntervalStart: stripe.Int64(1522540800),
			IntervalEnd:   stripe.Int64(1525132800),
		},
	}
	result, err := sc.ReportingReportRuns.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportTypesGet(t *testing.T) {
	params := &stripe.ReportingReportTypeListParams{}
	result := reporting_reporttype.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportTypesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReportingReportTypeListParams{}
	result := sc.ReportingReportTypes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportTypesGet2(t *testing.T) {
	params := &stripe.ReportingReportTypeParams{}
	result, err := reporting_reporttype.Get("balance.summary.1", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReportingReportTypesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReportingReportTypeParams{}
	result, err := sc.ReportingReportTypes.Get("balance.summary.1", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewsApprovePost(t *testing.T) {
	params := &stripe.ReviewApproveParams{}
	result, err := review.Approve("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewsApprovePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReviewApproveParams{}
	result, err := sc.Reviews.Approve("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewsGet(t *testing.T) {
	params := &stripe.ReviewListParams{}
	params.Limit = stripe.Int64(3)
	result := review.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReviewsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReviewListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Reviews.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReviewsGet2(t *testing.T) {
	params := &stripe.ReviewParams{}
	result, err := review.Get("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ReviewParams{}
	result, err := sc.Reviews.Get("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupAttemptsGet(t *testing.T) {
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("si_xyz"),
	}
	params.Limit = stripe.Int64(3)
	result := setupattempt.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupAttemptsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("si_xyz"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.SetupAttempts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentsCancelPost(t *testing.T) {
	params := &stripe.SetupIntentCancelParams{}
	result, err := setupintent.Cancel("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentCancelParams{}
	result, err := sc.SetupIntents.Cancel("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsConfirmPost(t *testing.T) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := setupintent.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsConfirmPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := sc.SetupIntents.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsGet(t *testing.T) {
	params := &stripe.SetupIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := setupintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.SetupIntents.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentsGet2(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	result, err := setupintent.Get("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentParams{}
	result, err := sc.SetupIntents.Get("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsPost(t *testing.T) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, err := setupintent.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, err := sc.SetupIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsPost2(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, err := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, err := sc.SetupIntents.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsVerifyMicrodepositsPost(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, err := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsVerifyMicrodepositsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, err := sc.SetupIntents.VerifyMicrodeposits(
		"seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsVerifyMicrodepositsPost2(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsVerifyMicrodepositsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := sc.SetupIntents.VerifyMicrodeposits(
		"seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesGet(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	result := shippingrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRatesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateListParams{}
	result := sc.ShippingRates.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRatesGet2(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	params.Limit = stripe.Int64(3)
	result := shippingrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRatesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.ShippingRates.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRatesGet3(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	result, err := shippingrate.Get("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateParams{}
	result, err := sc.ShippingRates.Get("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, err := shippingrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, err := sc.ShippingRates.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost2(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Ground shipping"),
		Type:        stripe.String("fixed_amount"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Amount:   stripe.Int64(500),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		},
	}
	result, err := shippingrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Ground shipping"),
		Type:        stripe.String("fixed_amount"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Amount:   stripe.Int64(500),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		},
	}
	result, err := sc.ShippingRates.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost3(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := shippingrate.Update("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.ShippingRateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.ShippingRates.Update("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSigmaScheduledQueryRunsGet(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sigma_scheduledqueryrun.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSigmaScheduledQueryRunsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.SigmaScheduledQueryRuns.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSigmaScheduledQueryRunsGet2(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunParams{}
	result, err := sigma_scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSigmaScheduledQueryRunsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SigmaScheduledQueryRunParams{}
	result, err := sc.SigmaScheduledQueryRuns.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGet(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SourceParams{}
	result, err := sc.Sources.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGet2(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SourceParams{}
	result, err := sc.Sources.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesPost(t *testing.T) {
	params := &stripe.SourceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := source.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SourceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Sources.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsDelete(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Del("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionItemParams{}
	result, err := sc.SubscriptionItems.Del("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsGet(t *testing.T) {
	params := &stripe.SubscriptionItemListParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
	}
	result := subscriptionitem.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionItemListParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
	}
	result := sc.SubscriptionItems.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionItemsGet2(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Get("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionItemParams{}
	result, err := sc.SubscriptionItems.Get("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsPost(t *testing.T) {
	params := &stripe.SubscriptionItemParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
		Price:        stripe.String("price_xxxxxxxxxxxxx"),
		Quantity:     stripe.Int64(2),
	}
	result, err := subscriptionitem.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionItemParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
		Price:        stripe.String("price_xxxxxxxxxxxxx"),
		Quantity:     stripe.Int64(2),
	}
	result, err := sc.SubscriptionItems.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsPost2(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscriptionitem.Update("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.SubscriptionItems.Update("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesCancelPost(t *testing.T) {
	params := &stripe.SubscriptionScheduleCancelParams{}
	result, err := subscriptionschedule.Cancel("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleCancelParams{}
	result, err := sc.SubscriptionSchedules.Cancel(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesGet(t *testing.T) {
	params := &stripe.SubscriptionScheduleListParams{}
	params.Limit = stripe.Int64(3)
	result := subscriptionschedule.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionSchedulesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.SubscriptionSchedules.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionSchedulesGet2(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{}
	result, err := subscriptionschedule.Get("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleParams{}
	result, err := sc.SubscriptionSchedules.Get(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		Customer:    stripe.String("cus_xxxxxxxxxxxxx"),
		StartDate:   stripe.Int64(1676070661),
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Price:    stripe.String("price_xxxxxxxxxxxxx"),
						Quantity: stripe.Int64(1),
					},
				},
				Iterations: stripe.Int64(12),
			},
		},
	}
	result, err := subscriptionschedule.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleParams{
		Customer:    stripe.String("cus_xxxxxxxxxxxxx"),
		StartDate:   stripe.Int64(1676070661),
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Price:    stripe.String("price_xxxxxxxxxxxxx"),
						Quantity: stripe.Int64(1),
					},
				},
				Iterations: stripe.Int64(12),
			},
		},
	}
	result, err := sc.SubscriptionSchedules.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost2(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
	}
	result, err := subscriptionschedule.Update("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
	}
	result, err := sc.SubscriptionSchedules.Update(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesReleasePost(t *testing.T) {
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := subscriptionschedule.Release(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesReleasePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := sc.SubscriptionSchedules.Release(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDelete(t *testing.T) {
	params := &stripe.SubscriptionCancelParams{}
	result, err := subscription.Cancel("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionCancelParams{}
	result, err := sc.Subscriptions.Cancel("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDiscountDelete(t *testing.T) {
	params := &stripe.SubscriptionDeleteDiscountParams{}
	result, err := subscription.DeleteDiscount("sub_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDiscountDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionDeleteDiscountParams{}
	result, err := sc.Subscriptions.DeleteDiscount("sub_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsGet(t *testing.T) {
	params := &stripe.SubscriptionListParams{}
	params.Limit = stripe.Int64(3)
	result := subscription.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Subscriptions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionsGet2(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	result, err := subscription.Get("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionParams{}
	result, err := sc.Subscriptions.Get("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsPost(t *testing.T) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_xxxxxxxxxxxxx"),
			},
		},
	}
	result, err := subscription.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_xxxxxxxxxxxxx"),
			},
		},
	}
	result, err := sc.Subscriptions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsPost2(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscription.Update("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Subscriptions.Update("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsSearchGet(t *testing.T) {
	params := &stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'active' AND metadata['order_id']:'6735'",
		},
	}
	result := subscription.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionsSearchGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'active' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.Subscriptions.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCalculationsLineItemsGet(t *testing.T) {
	params := &stripe.TaxCalculationListLineItemsParams{
		Calculation: stripe.String("xxx"),
	}
	result := tax_calculation.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCalculationsLineItemsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxCalculationListLineItemsParams{
		Calculation: stripe.String("xxx"),
	}
	result := sc.TaxCalculations.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCalculationsPost(t *testing.T) {
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		LineItems: []*stripe.TaxCalculationLineItemParams{
			{
				Amount:    stripe.Int64(1000),
				Reference: stripe.String("L1"),
			},
		},
		CustomerDetails: &stripe.TaxCalculationCustomerDetailsParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("354 Oyster Point Blvd"),
				City:       stripe.String("South San Francisco"),
				State:      stripe.String("CA"),
				PostalCode: stripe.String("94080"),
				Country:    stripe.String("US"),
			},
			AddressSource: stripe.String(string(stripe.TaxCalculationCustomerDetailsAddressSourceShipping)),
		},
	}
	result, err := tax_calculation.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCalculationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		LineItems: []*stripe.TaxCalculationLineItemParams{
			{
				Amount:    stripe.Int64(1000),
				Reference: stripe.String("L1"),
			},
		},
		CustomerDetails: &stripe.TaxCalculationCustomerDetailsParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("354 Oyster Point Blvd"),
				City:       stripe.String("South San Francisco"),
				State:      stripe.String("CA"),
				PostalCode: stripe.String("94080"),
				Country:    stripe.String("US"),
			},
			AddressSource: stripe.String(string(stripe.TaxCalculationCustomerDetailsAddressSourceShipping)),
		},
	}
	result, err := sc.TaxCalculations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCodesGet(t *testing.T) {
	params := &stripe.TaxCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := taxcode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCodesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TaxCodes.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCodesGet2(t *testing.T) {
	params := &stripe.TaxCodeParams{}
	result, err := taxcode.Get("txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCodesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxCodeParams{}
	result, err := sc.TaxCodes.Get("txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsDelete(t *testing.T) {
	params := &stripe.TaxIDParams{}
	result, err := taxid.Del("taxid_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{}
	result, err := sc.TaxIDs.Del("taxid_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsGet(t *testing.T) {
	params := &stripe.TaxIDListParams{}
	result := taxid.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxIdsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDListParams{}
	result := sc.TaxIDs.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxIdsGet2(t *testing.T) {
	params := &stripe.TaxIDParams{}
	result, err := taxid.Get("taxid_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{}
	result, err := sc.TaxIDs.Get("taxid_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsPost(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:  stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value: stripe.String("123"),
	}
	result, err := taxid.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{
		Type:  stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value: stripe.String("123"),
	}
	result, err := sc.TaxIDs.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesGet(t *testing.T) {
	params := &stripe.TaxRateListParams{}
	params.Limit = stripe.Int64(3)
	result := taxrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxRatesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxRateListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TaxRates.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxRatesGet2(t *testing.T) {
	params := &stripe.TaxRateParams{}
	result, err := taxrate.Get("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxRateParams{}
	result, err := sc.TaxRates.Get("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesPost(t *testing.T) {
	params := &stripe.TaxRateParams{
		DisplayName:  stripe.String("VAT"),
		Description:  stripe.String("VAT Germany"),
		Jurisdiction: stripe.String("DE"),
		Percentage:   stripe.Float64(16),
		Inclusive:    stripe.Bool(false),
	}
	result, err := taxrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxRateParams{
		DisplayName:  stripe.String("VAT"),
		Description:  stripe.String("VAT Germany"),
		Jurisdiction: stripe.String("DE"),
		Percentage:   stripe.Float64(16),
		Inclusive:    stripe.Bool(false),
	}
	result, err := sc.TaxRates.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesPost2(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, err := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRatesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, err := sc.TaxRates.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxSettingsGet(t *testing.T) {
	params := &stripe.TaxSettingsParams{}
	result, err := tax_settings.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxSettingsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxSettingsParams{}
	result, err := sc.TaxSettings.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxSettingsPost(t *testing.T) {
	params := &stripe.TaxSettingsParams{
		Defaults: &stripe.TaxSettingsDefaultsParams{
			TaxCode: stripe.String("txcd_10000000"),
		},
	}
	result, err := tax_settings.Update(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxSettingsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxSettingsParams{
		Defaults: &stripe.TaxSettingsDefaultsParams{
			TaxCode: stripe.String("txcd_10000000"),
		},
	}
	result, err := sc.TaxSettings.Update(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxTransactionsCreateFromCalculationPost(t *testing.T) {
	params := &stripe.TaxTransactionCreateFromCalculationParams{
		Calculation: stripe.String("xxx"),
		Reference:   stripe.String("yyy"),
	}
	result, err := tax_transaction.CreateFromCalculation(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxTransactionsCreateFromCalculationPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxTransactionCreateFromCalculationParams{
		Calculation: stripe.String("xxx"),
		Reference:   stripe.String("yyy"),
	}
	result, err := sc.TaxTransactions.CreateFromCalculation(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDelete(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{}
	result, err := sc.TerminalConfigurations.Del("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDelete2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDelete2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{}
	result, err := sc.TerminalConfigurations.Del("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsGet(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationListParams{}
	result := sc.TerminalConfigurations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationsGet2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{}
	result, err := sc.TerminalConfigurations.Get("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsGet3(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationsGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TerminalConfigurations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationsGet4(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsGet4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{}
	result, err := sc.TerminalConfigurations.Get("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{}
	result, err := sc.TerminalConfigurations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		Tipping: &stripe.TerminalConfigurationTippingParams{
			USD: &stripe.TerminalConfigurationTippingUSDParams{
				FixedAmounts: []*int64{stripe.Int64(10)},
			},
		},
	}
	result, err := terminal_configuration.Update("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{
		Tipping: &stripe.TerminalConfigurationTippingParams{
			USD: &stripe.TerminalConfigurationTippingUSDParams{
				FixedAmounts: []*int64{stripe.Int64(10)},
			},
		},
	}
	result, err := sc.TerminalConfigurations.Update("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost3(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := terminal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := sc.TerminalConfigurations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost4(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := terminal_configuration.Update("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := sc.TerminalConfigurations.Update("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConnectionTokensPost(t *testing.T) {
	params := &stripe.TerminalConnectionTokenParams{}
	result, err := terminal_connectiontoken.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConnectionTokensPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalConnectionTokenParams{}
	result, err := sc.TerminalConnectionTokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsDelete(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Del("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalLocationParams{}
	result, err := sc.TerminalLocations.Del("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsGet(t *testing.T) {
	params := &stripe.TerminalLocationListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_location.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalLocationsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalLocationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TerminalLocations.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalLocationsGet2(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Get("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalLocationParams{}
	result, err := sc.TerminalLocations.Get("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsPost(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
		Address: &stripe.AddressParams{
			Line1:      stripe.String("1234 Main Street"),
			City:       stripe.String("San Francisco"),
			PostalCode: stripe.String("94111"),
			State:      stripe.String("CA"),
			Country:    stripe.String("US"),
		},
	}
	result, err := terminal_location.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
		Address: &stripe.AddressParams{
			Line1:      stripe.String("1234 Main Street"),
			City:       stripe.String("San Francisco"),
			PostalCode: stripe.String("94111"),
			State:      stripe.String("CA"),
			Country:    stripe.String("US"),
		},
	}
	result, err := sc.TerminalLocations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsPost2(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, err := terminal_location.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, err := sc.TerminalLocations.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersCancelActionPost(t *testing.T) {
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := terminal_reader.CancelAction("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersCancelActionPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := sc.TerminalReaders.CancelAction("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersDelete(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, err := terminal_reader.Del("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderParams{}
	result, err := sc.TerminalReaders.Del("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersGet(t *testing.T) {
	params := &stripe.TerminalReaderListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_reader.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalReadersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TerminalReaders.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalReadersGet2(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, err := terminal_reader.Get("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderParams{}
	result, err := sc.TerminalReaders.Get("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersPost(t *testing.T) {
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, err := terminal_reader.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, err := sc.TerminalReaders.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersPost2(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, err := terminal_reader.Update("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, err := sc.TerminalReaders.Update("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessPaymentIntentPost(t *testing.T) {
	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String("pi_xxxxxxxxxxxxx"),
	}
	result, err := terminal_reader.ProcessPaymentIntent(
		"tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessPaymentIntentPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String("pi_xxxxxxxxxxxxx"),
	}
	result, err := sc.TerminalReaders.ProcessPaymentIntent(
		"tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessSetupIntentPost(t *testing.T) {
	params := &stripe.TerminalReaderProcessSetupIntentParams{
		SetupIntent:    stripe.String("seti_xxxxxxxxxxxxx"),
		AllowRedisplay: stripe.String("always"),
	}
	result, err := terminal_reader.ProcessSetupIntent(
		"tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessSetupIntentPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TerminalReaderProcessSetupIntentParams{
		SetupIntent:    stripe.String("seti_xxxxxxxxxxxxx"),
		AllowRedisplay: stripe.String("always"),
	}
	result, err := sc.TerminalReaders.ProcessSetupIntent(
		"tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersCustomersFundCashBalancePost(t *testing.T) {
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	result, err := testhelpers_customer.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersCustomersFundCashBalancePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	result, err := sc.TestHelpersCustomers.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsCapturePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationCaptureParams{
		CaptureAmount:      stripe.Int64(100),
		CloseAuthorization: stripe.Bool(true),
		PurchaseDetails: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1633651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1633651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := testhelpers_issuing_authorization.Capture(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsCapturePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingAuthorizationCaptureParams{
		CaptureAmount:      stripe.Int64(100),
		CloseAuthorization: stripe.Bool(true),
		PurchaseDetails: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1633651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1633651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingAuthorizationCapturePurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := sc.TestHelpersIssuingAuthorizations.Capture(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsExpirePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationExpireParams{}
	result, err := testhelpers_issuing_authorization.Expire(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsExpirePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingAuthorizationExpireParams{}
	result, err := sc.TestHelpersIssuingAuthorizations.Expire(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsIncrementPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationIncrementParams{
		IncrementAmount:      stripe.Int64(50),
		IsAmountControllable: stripe.Bool(true),
	}
	result, err := testhelpers_issuing_authorization.Increment(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsIncrementPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingAuthorizationIncrementParams{
		IncrementAmount:      stripe.Int64(50),
		IsAmountControllable: stripe.Bool(true),
	}
	result, err := sc.TestHelpersIssuingAuthorizations.Increment(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationParams{
		Amount: stripe.Int64(100),
		AmountDetails: &stripe.TestHelpersIssuingAuthorizationAmountDetailsParams{
			ATMFee:         stripe.Int64(10),
			CashbackAmount: stripe.Int64(5),
		},
		AuthorizationMethod:  stripe.String(string(stripe.IssuingAuthorizationAuthorizationMethodChip)),
		Card:                 stripe.String("foo"),
		Currency:             stripe.String(string(stripe.CurrencyUSD)),
		IsAmountControllable: stripe.Bool(true),
		MerchantData: &stripe.TestHelpersIssuingAuthorizationMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("bar"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("foo"),
			State:      stripe.String("bar"),
			TerminalID: stripe.String("foo"),
		},
		NetworkData: &stripe.TestHelpersIssuingAuthorizationNetworkDataParams{
			AcquiringInstitutionID: stripe.String("foo"),
		},
		VerificationData: &stripe.TestHelpersIssuingAuthorizationVerificationDataParams{
			AddressLine1Check:      stripe.String("mismatch"),
			AddressPostalCodeCheck: stripe.String("match"),
			CVCCheck:               stripe.String("match"),
			ExpiryCheck:            stripe.String("mismatch"),
		},
		Wallet: stripe.String("apple_pay"),
	}
	result, err := testhelpers_issuing_authorization.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingAuthorizationParams{
		Amount: stripe.Int64(100),
		AmountDetails: &stripe.TestHelpersIssuingAuthorizationAmountDetailsParams{
			ATMFee:         stripe.Int64(10),
			CashbackAmount: stripe.Int64(5),
		},
		AuthorizationMethod:  stripe.String(string(stripe.IssuingAuthorizationAuthorizationMethodChip)),
		Card:                 stripe.String("foo"),
		Currency:             stripe.String(string(stripe.CurrencyUSD)),
		IsAmountControllable: stripe.Bool(true),
		MerchantData: &stripe.TestHelpersIssuingAuthorizationMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("bar"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("foo"),
			State:      stripe.String("bar"),
			TerminalID: stripe.String("foo"),
		},
		NetworkData: &stripe.TestHelpersIssuingAuthorizationNetworkDataParams{
			AcquiringInstitutionID: stripe.String("foo"),
		},
		VerificationData: &stripe.TestHelpersIssuingAuthorizationVerificationDataParams{
			AddressLine1Check:      stripe.String("mismatch"),
			AddressPostalCodeCheck: stripe.String("match"),
			CVCCheck:               stripe.String("match"),
			ExpiryCheck:            stripe.String("mismatch"),
		},
		Wallet: stripe.String("apple_pay"),
	}
	result, err := sc.TestHelpersIssuingAuthorizations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsReversePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationReverseParams{
		ReverseAmount: stripe.Int64(20),
	}
	result, err := testhelpers_issuing_authorization.Reverse(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsReversePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingAuthorizationReverseParams{
		ReverseAmount: stripe.Int64(20),
	}
	result, err := sc.TestHelpersIssuingAuthorizations.Reverse(
		"example_authorization", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingDeliverPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, err := testhelpers_issuing_card.DeliverCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingDeliverPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, err := sc.TestHelpersIssuingCards.DeliverCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingFailPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, err := testhelpers_issuing_card.FailCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingFailPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, err := sc.TestHelpersIssuingCards.FailCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingReturnPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, err := testhelpers_issuing_card.ReturnCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingReturnPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, err := sc.TestHelpersIssuingCards.ReturnCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingShipPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, err := testhelpers_issuing_card.ShipCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingShipPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, err := sc.TestHelpersIssuingCards.ShipCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsActivatePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingPersonalizationDesignActivateParams{}
	result, err := testhelpers_issuing_personalizationdesign.Activate(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsActivatePostService(
	t *testing.T,
) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingPersonalizationDesignActivateParams{}
	result, err := sc.TestHelpersIssuingPersonalizationDesigns.Activate(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsDeactivatePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams{}
	result, err := testhelpers_issuing_personalizationdesign.Deactivate(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsDeactivatePostService(
	t *testing.T,
) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams{}
	result, err := sc.TestHelpersIssuingPersonalizationDesigns.Deactivate(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsRejectPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingPersonalizationDesignRejectParams{
		RejectionReasons: &stripe.TestHelpersIssuingPersonalizationDesignRejectRejectionReasonsParams{
			CardLogo: []*string{
				stripe.String(string(stripe.IssuingPersonalizationDesignRejectionReasonsCardLogoGeographicLocation)),
			},
		},
	}
	result, err := testhelpers_issuing_personalizationdesign.Reject(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsRejectPostService(
	t *testing.T,
) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingPersonalizationDesignRejectParams{
		RejectionReasons: &stripe.TestHelpersIssuingPersonalizationDesignRejectRejectionReasonsParams{
			CardLogo: []*string{
				stripe.String(string(stripe.IssuingPersonalizationDesignRejectionReasonsCardLogoGeographicLocation)),
			},
		},
	}
	result, err := sc.TestHelpersIssuingPersonalizationDesigns.Reject(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateForceCapturePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionCreateForceCaptureParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		MerchantData: &stripe.TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("US"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("10001"),
			State:      stripe.String("NY"),
			TerminalID: stripe.String("foo"),
		},
		PurchaseDetails: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1633651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1533651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := testhelpers_issuing_transaction.CreateForceCapture(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateForceCapturePostService(
	t *testing.T,
) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingTransactionCreateForceCaptureParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		MerchantData: &stripe.TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("US"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("10001"),
			State:      stripe.String("NY"),
			TerminalID: stripe.String("foo"),
		},
		PurchaseDetails: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1633651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1533651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := sc.TestHelpersIssuingTransactions.CreateForceCapture(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateUnlinkedRefundPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		MerchantData: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("bar"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("foo"),
			State:      stripe.String("bar"),
			TerminalID: stripe.String("foo"),
		},
		PurchaseDetails: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1533651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1533651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := testhelpers_issuing_transaction.CreateUnlinkedRefund(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateUnlinkedRefundPostService(
	t *testing.T,
) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		MerchantData: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("bar"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("foo"),
			State:      stripe.String("bar"),
			TerminalID: stripe.String("foo"),
		},
		PurchaseDetails: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams{
			Flight: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightParams{
				DepartureAt:   stripe.Int64(1533651200),
				PassengerName: stripe.String("John Doe"),
				Refundable:    stripe.Bool(true),
				Segments: []*stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightSegmentParams{
					{
						ArrivalAirportCode:   stripe.String("SFO"),
						Carrier:              stripe.String("Delta"),
						DepartureAirportCode: stripe.String("LAX"),
						FlightNumber:         stripe.String("DL100"),
						ServiceClass:         stripe.String("Economy"),
						StopoverAllowed:      stripe.Bool(true),
					},
				},
				TravelAgency: stripe.String("Orbitz"),
			},
			Fuel: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams{
				Type:            stripe.String("diesel"),
				Unit:            stripe.String("liter"),
				UnitCostDecimal: stripe.Float64(3.5),
				QuantityDecimal: stripe.Float64(10),
			},
			Lodging: &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsLodgingParams{
				CheckInAt: stripe.Int64(1533651200),
				Nights:    stripe.Int64(2),
			},
			Receipt: []*stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsReceiptParams{
				{
					Description: stripe.String("Room charge"),
					Quantity:    stripe.Float64(1),
					Total:       stripe.Int64(200),
					UnitCost:    stripe.Int64(200),
				},
			},
			Reference: stripe.String("foo"),
		},
	}
	result, err := sc.TestHelpersIssuingTransactions.CreateUnlinkedRefund(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsRefundPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionRefundParams{
		RefundAmount: stripe.Int64(50),
	}
	result, err := testhelpers_issuing_transaction.Refund(
		"example_transaction", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsRefundPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersIssuingTransactionRefundParams{
		RefundAmount: stripe.Int64(50),
	}
	result, err := sc.TestHelpersIssuingTransactions.Refund(
		"example_transaction", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersRefundsExpirePost(t *testing.T) {
	params := &stripe.TestHelpersRefundExpireParams{}
	result, err := testhelpers_refund.Expire("re_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersRefundsExpirePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersRefundExpireParams{}
	result, err := sc.TestHelpersRefunds.Expire("re_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksAdvancePost(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, err := testhelpers_testclock.Advance("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksAdvancePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, err := sc.TestHelpersTestClocks.Advance("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksAdvancePost2(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1675552261),
	}
	result, err := testhelpers_testclock.Advance("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksAdvancePost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1675552261),
	}
	result, err := sc.TestHelpersTestClocks.Advance(
		"clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDelete(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{}
	result, err := sc.TestHelpersTestClocks.Del("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDelete2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDelete2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{}
	result, err := sc.TestHelpersTestClocks.Del("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksGet(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClocksGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockListParams{}
	result := sc.TestHelpersTestClocks.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClocksGet2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{}
	result, err := sc.TestHelpersTestClocks.Get("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksGet3(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	params.Limit = stripe.Int64(3)
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClocksGet3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TestHelpersTestClocks.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClocksGet4(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksGet4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{}
	result, err := sc.TestHelpersTestClocks.Get("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksPost(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(123),
		Name:       stripe.String("cogsworth"),
	}
	result, err := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(123),
		Name:       stripe.String("cogsworth"),
	}
	result, err := sc.TestHelpersTestClocks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksPost2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(1577836800),
	}
	result, err := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(1577836800),
	}
	result, err := sc.TestHelpersTestClocks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersFailPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(string(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed)),
		},
	}
	result, err := testhelpers_treasury_inboundtransfer.Fail("ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersFailPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(string(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed)),
		},
	}
	result, err := sc.TestHelpersTreasuryInboundTransfers.Fail("ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersReturnPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, err := testhelpers_treasury_inboundtransfer.ReturnInboundTransfer(
		"ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersReturnPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, err := sc.TestHelpersTreasuryInboundTransfers.ReturnInboundTransfer(
		"ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersSucceedPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := testhelpers_treasury_inboundtransfer.Succeed(
		"ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersSucceedPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := sc.TestHelpersTreasuryInboundTransfers.Succeed(
		"ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersFailPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, err := testhelpers_treasury_outboundtransfer.Fail("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersFailPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, err := sc.TestHelpersTreasuryOutboundTransfers.Fail(
		"obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersPostPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, err := testhelpers_treasury_outboundtransfer.Post("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersPostPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, err := sc.TestHelpersTreasuryOutboundTransfers.Post(
		"obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersReturnPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams{
		ReturnedDetails: &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams{
			Code: stripe.String(string(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed)),
		},
	}
	result, err := testhelpers_treasury_outboundtransfer.ReturnOutboundTransfer(
		"obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersReturnPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams{
		ReturnedDetails: &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams{
			Code: stripe.String(string(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed)),
		},
	}
	result, err := sc.TestHelpersTreasuryOutboundTransfers.ReturnOutboundTransfer(
		"obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedCreditsPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedCreditParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(string(stripe.TreasuryReceivedCreditNetworkACH)),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := testhelpers_treasury_receivedcredit.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedCreditsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryReceivedCreditParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(string(stripe.TreasuryReceivedCreditNetworkACH)),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := sc.TestHelpersTreasuryReceivedCredits.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedDebitsPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedDebitParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String("ach"),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := testhelpers_treasury_receiveddebit.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedDebitsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryReceivedDebitParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String("ach"),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, err := sc.TestHelpersTreasuryReceivedDebits.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensGet(t *testing.T) {
	params := &stripe.TokenParams{}
	result, err := token.Get("tok_xxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{}
	result, err := sc.Tokens.Get("tok_xxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost(t *testing.T) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.String("5"),
			ExpYear:  stripe.String("2023"),
			CVC:      stripe.String("314"),
		},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.String("5"),
			ExpYear:  stripe.String("2023"),
			CVC:      stripe.String("314"),
		},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost2(t *testing.T) {
	params := &stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			Country:           stripe.String("US"),
			Currency:          stripe.String(string(stripe.CurrencyUSD)),
			AccountHolderName: stripe.String("Jenny Rosen"),
			AccountHolderType: stripe.String("individual"),
			RoutingNumber:     stripe.String("110000000"),
			AccountNumber:     stripe.String("000123456789"),
		},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			Country:           stripe.String("US"),
			Currency:          stripe.String(string(stripe.CurrencyUSD)),
			AccountHolderName: stripe.String("Jenny Rosen"),
			AccountHolderType: stripe.String("individual"),
			RoutingNumber:     stripe.String("110000000"),
			AccountNumber:     stripe.String("000123456789"),
		},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost3(t *testing.T) {
	params := &stripe.TokenParams{
		PII: &stripe.TokenPIIParams{IDNumber: stripe.String("000000000")},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost3Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		PII: &stripe.TokenPIIParams{IDNumber: stripe.String("000000000")},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost4(t *testing.T) {
	params := &stripe.TokenParams{
		Account: &stripe.TokenAccountParams{
			Individual: &stripe.PersonParams{
				FirstName: stripe.String("Jane"),
				LastName:  stripe.String("Doe"),
			},
			TOSShownAndAccepted: stripe.Bool(true),
		},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost4Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		Account: &stripe.TokenAccountParams{
			Individual: &stripe.PersonParams{
				FirstName: stripe.String("Jane"),
				LastName:  stripe.String("Doe"),
			},
			TOSShownAndAccepted: stripe.Bool(true),
		},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost5(t *testing.T) {
	params := &stripe.TokenParams{
		Person: &stripe.PersonParams{
			FirstName:    stripe.String("Jane"),
			LastName:     stripe.String("Doe"),
			Relationship: &stripe.PersonRelationshipParams{Owner: stripe.Bool(true)},
		},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost5Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		Person: &stripe.PersonParams{
			FirstName:    stripe.String("Jane"),
			LastName:     stripe.String("Doe"),
			Relationship: &stripe.PersonRelationshipParams{Owner: stripe.Bool(true)},
		},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost6(t *testing.T) {
	params := &stripe.TokenParams{
		CVCUpdate: &stripe.TokenCVCUpdateParams{CVC: stripe.String("123")},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost6Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TokenParams{
		CVCUpdate: &stripe.TokenCVCUpdateParams{CVC: stripe.String("123")},
	}
	result, err := sc.Tokens.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsCancelPost(t *testing.T) {
	params := &stripe.TopupParams{}
	result, err := topup.Cancel("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TopupParams{}
	result, err := sc.Topups.Cancel("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsGet(t *testing.T) {
	params := &stripe.TopupListParams{}
	params.Limit = stripe.Int64(3)
	result := topup.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTopupsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TopupListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Topups.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTopupsGet2(t *testing.T) {
	params := &stripe.TopupParams{}
	result, err := topup.Get("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TopupParams{}
	result, err := sc.Topups.Get("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPost(t *testing.T) {
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, err := topup.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, err := sc.Topups.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPost2(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, err := topup.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Topups.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersGet(t *testing.T) {
	params := &stripe.TransferListParams{}
	params.Limit = stripe.Int64(3)
	result := transfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransfersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.Transfers.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransfersGet2(t *testing.T) {
	params := &stripe.TransferParams{}
	result, err := transfer.Get("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferParams{}
	result, err := sc.Transfers.Get("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPost(t *testing.T) {
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, err := transfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, err := sc.Transfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPost2(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, err := transfer.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.Transfers.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsGet(t *testing.T) {
	params := &stripe.TransferReversalListParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := transferreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransfersReversalsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferReversalListParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TransferReversals.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransfersReversalsGet2(t *testing.T) {
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := transferreversal.Get("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := sc.TransferReversals.Get("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsPost(t *testing.T) {
	params := &stripe.TransferReversalParams{
		Amount: stripe.Int64(100),
		ID:     stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := transferreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferReversalParams{
		Amount: stripe.Int64(100),
		ID:     stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := sc.TransferReversals.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsPost2(t *testing.T) {
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := transferreversal.Update("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersReversalsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.TransferReversals.Update("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalsGet(t *testing.T) {
	params := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_creditreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryCreditReversalsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryCreditReversals.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryCreditReversalsGet2(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{}
	result, err := treasury_creditreversal.Get("credrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryCreditReversalParams{}
	result, err := sc.TreasuryCreditReversals.Get(
		"credrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalsPost(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{
		ReceivedCredit: stripe.String("rc_xxxxxxxxxxxxx"),
	}
	result, err := treasury_creditreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryCreditReversalParams{
		ReceivedCredit: stripe.String("rc_xxxxxxxxxxxxx"),
	}
	result, err := sc.TreasuryCreditReversals.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryDebitReversalsGet(t *testing.T) {
	params := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_debitreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryDebitReversalsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryDebitReversals.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryDebitReversalsGet2(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{}
	result, err := treasury_debitreversal.Get("debrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryDebitReversalsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryDebitReversalParams{}
	result, err := sc.TreasuryDebitReversals.Get("debrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryDebitReversalsPost(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{
		ReceivedDebit: stripe.String("rd_xxxxxxxxxxxxx"),
	}
	result, err := treasury_debitreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryDebitReversalsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryDebitReversalParams{
		ReceivedDebit: stripe.String("rd_xxxxxxxxxxxxx"),
	}
	result, err := sc.TreasuryDebitReversals.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsFeaturesGet(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, err := treasury_financialaccount.RetrieveFeatures(
		"fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsFeaturesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, err := sc.TreasuryFinancialAccounts.RetrieveFeatures(
		"fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsGet(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountListParams{}
	params.Limit = stripe.Int64(3)
	result := treasury_financialaccount.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryFinancialAccountsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryFinancialAccountListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryFinancialAccounts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryFinancialAccountsGet2(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	result, err := treasury_financialaccount.Get("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryFinancialAccountParams{}
	result, err := sc.TreasuryFinancialAccounts.Get("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsPost(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{
		SupportedCurrencies: []*string{stripe.String("usd")},
		Features:            &stripe.TreasuryFinancialAccountFeaturesParams{},
	}
	result, err := treasury_financialaccount.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryFinancialAccountParams{
		SupportedCurrencies: []*string{stripe.String("usd")},
		Features:            &stripe.TreasuryFinancialAccountFeaturesParams{},
	}
	result, err := sc.TreasuryFinancialAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsPost2(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := treasury_financialaccount.Update("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryFinancialAccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.TreasuryFinancialAccounts.Update(
		"fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersCancelPost(t *testing.T) {
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, err := treasury_inboundtransfer.Cancel("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, err := sc.TreasuryInboundTransfers.Cancel(
		"ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersGet(t *testing.T) {
	params := &stripe.TreasuryInboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_inboundtransfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryInboundTransfersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryInboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryInboundTransfers.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryInboundTransfersGet2(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{}
	result, err := treasury_inboundtransfer.Get("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryInboundTransferParams{}
	result, err := sc.TreasuryInboundTransfers.Get("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersPost(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{
		FinancialAccount:    stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:              stripe.Int64(10000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		OriginPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:         stripe.String("InboundTransfer from my bank account"),
	}
	result, err := treasury_inboundtransfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryInboundTransferParams{
		FinancialAccount:    stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:              stripe.Int64(10000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		OriginPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:         stripe.String("InboundTransfer from my bank account"),
	}
	result, err := sc.TreasuryInboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsCancelPost(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, err := treasury_outboundpayment.Cancel("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, err := sc.TreasuryOutboundPayments.Cancel(
		"bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsGet(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_outboundpayment.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundPaymentsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundPaymentListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryOutboundPayments.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundPaymentsGet2(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{}
	result, err := treasury_outboundpayment.Get("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundPaymentParams{}
	result, err := sc.TreasuryOutboundPayments.Get("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsPost(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(10000),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Customer:                 stripe.String("cus_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, err := treasury_outboundpayment.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundPaymentParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(10000),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Customer:                 stripe.String("cus_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, err := sc.TreasuryOutboundPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersCancelPost(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, err := treasury_outboundtransfer.Cancel("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersCancelPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, err := sc.TreasuryOutboundTransfers.Cancel(
		"obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersGet(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_outboundtransfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundTransfersGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryOutboundTransfers.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundTransfersGet2(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{}
	result, err := treasury_outboundtransfer.Get("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundTransferParams{}
	result, err := sc.TreasuryOutboundTransfers.Get("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersPost(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(500),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Description:              stripe.String("OutboundTransfer to my external bank account"),
	}
	result, err := treasury_outboundtransfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryOutboundTransferParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(500),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Description:              stripe.String("OutboundTransfer to my external bank account"),
	}
	result, err := sc.TreasuryOutboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryReceivedCreditsGet(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_receivedcredit.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedCreditsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryReceivedCredits.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedCreditsGet2(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditParams{}
	result, err := treasury_receivedcredit.Get("rc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryReceivedCreditsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryReceivedCreditParams{}
	result, err := sc.TreasuryReceivedCredits.Get("rc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryReceivedDebitsGet(t *testing.T) {
	params := &stripe.TreasuryReceivedDebitListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_receiveddebit.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedDebitsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryReceivedDebitListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryReceivedDebits.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedDebitsGet2(t *testing.T) {
	params := &stripe.TreasuryReceivedDebitParams{}
	result, err := treasury_receiveddebit.Get("rd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryReceivedDebitsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryReceivedDebitParams{}
	result, err := sc.TreasuryReceivedDebits.Get("rd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryTransactionEntriesGet(t *testing.T) {
	params := &stripe.TreasuryTransactionEntryListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_transactionentry.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionEntriesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryTransactionEntryListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryTransactionEntries.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionEntriesGet2(t *testing.T) {
	params := &stripe.TreasuryTransactionEntryParams{}
	result, err := treasury_transactionentry.Get("trxne_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryTransactionEntriesGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryTransactionEntryParams{}
	result, err := sc.TreasuryTransactionEntries.Get(
		"trxne_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryTransactionsGet(t *testing.T) {
	params := &stripe.TreasuryTransactionListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_transaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryTransactionListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.TreasuryTransactions.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionsGet2(t *testing.T) {
	params := &stripe.TreasuryTransactionParams{}
	result, err := treasury_transaction.Get("trxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryTransactionsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TreasuryTransactionParams{}
	result, err := sc.TreasuryTransactions.Get("trxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsDelete(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Del("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsDeleteService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.WebhookEndpointParams{}
	result, err := sc.WebhookEndpoints.Del("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsGet(t *testing.T) {
	params := &stripe.WebhookEndpointListParams{}
	params.Limit = stripe.Int64(3)
	result := webhookendpoint.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestWebhookEndpointsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.WebhookEndpointListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.WebhookEndpoints.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestWebhookEndpointsGet2(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Get("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsGet2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.WebhookEndpointParams{}
	result, err := sc.WebhookEndpoints.Get("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsPost(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/my/webhook/endpoint"),
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
	}
	result, err := webhookendpoint.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/my/webhook/endpoint"),
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
	}
	result, err := sc.WebhookEndpoints.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsPost2(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, err := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, err := sc.WebhookEndpoints.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCoreEventsGetService(t *testing.T) {
	params := &stripe.V2CoreEventParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events/ll_123", params, "{\"context\":\"context\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"livemode\":true,\"object\":\"v2.core.event\",\"reason\":{\"type\":\"request\",\"request\":{\"id\":\"obj_123\",\"idempotency_key\":\"idempotency_key\"}},\"type\":\"type\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEvents.Get("ll_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
