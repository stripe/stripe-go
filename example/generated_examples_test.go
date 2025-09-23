//
//
// File generated from our OpenAPI spec
//
//

package example

import (
	"testing"
	"time"

	"context"
	"net/http"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
	account "github.com/stripe/stripe-go/v82/account"
	accountlink "github.com/stripe/stripe-go/v82/accountlink"
	applicationfee "github.com/stripe/stripe-go/v82/applicationfee"
	apps_secret "github.com/stripe/stripe-go/v82/apps/secret"
	balancetransaction "github.com/stripe/stripe-go/v82/balancetransaction"
	billingportal_configuration "github.com/stripe/stripe-go/v82/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v82/billingportal/session"
	capability "github.com/stripe/stripe-go/v82/capability"
	card "github.com/stripe/stripe-go/v82/card"
	cashbalance "github.com/stripe/stripe-go/v82/cashbalance"
	charge "github.com/stripe/stripe-go/v82/charge"
	checkout_session "github.com/stripe/stripe-go/v82/checkout/session"
	"github.com/stripe/stripe-go/v82/client"
	countryspec "github.com/stripe/stripe-go/v82/countryspec"
	coupon "github.com/stripe/stripe-go/v82/coupon"
	customer "github.com/stripe/stripe-go/v82/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v82/customerbalancetransaction"
	customercashbalancetransaction "github.com/stripe/stripe-go/v82/customercashbalancetransaction"
	customersession "github.com/stripe/stripe-go/v82/customersession"
	dispute "github.com/stripe/stripe-go/v82/dispute"
	event "github.com/stripe/stripe-go/v82/event"
	feerefund "github.com/stripe/stripe-go/v82/feerefund"
	financialconnections_account "github.com/stripe/stripe-go/v82/financialconnections/account"
	financialconnections_session "github.com/stripe/stripe-go/v82/financialconnections/session"
	financialconnections_transaction "github.com/stripe/stripe-go/v82/financialconnections/transaction"
	identity_verificationreport "github.com/stripe/stripe-go/v82/identity/verificationreport"
	identity_verificationsession "github.com/stripe/stripe-go/v82/identity/verificationsession"
	invoice "github.com/stripe/stripe-go/v82/invoice"
	invoiceitem "github.com/stripe/stripe-go/v82/invoiceitem"
	issuing_authorization "github.com/stripe/stripe-go/v82/issuing/authorization"
	issuing_card "github.com/stripe/stripe-go/v82/issuing/card"
	issuing_cardholder "github.com/stripe/stripe-go/v82/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v82/issuing/dispute"
	issuing_personalizationdesign "github.com/stripe/stripe-go/v82/issuing/personalizationdesign"
	issuing_physicalbundle "github.com/stripe/stripe-go/v82/issuing/physicalbundle"
	issuing_transaction "github.com/stripe/stripe-go/v82/issuing/transaction"
	loginlink "github.com/stripe/stripe-go/v82/loginlink"
	mandate "github.com/stripe/stripe-go/v82/mandate"
	paymentintent "github.com/stripe/stripe-go/v82/paymentintent"
	paymentlink "github.com/stripe/stripe-go/v82/paymentlink"
	paymentmethod "github.com/stripe/stripe-go/v82/paymentmethod"
	paymentmethodconfiguration "github.com/stripe/stripe-go/v82/paymentmethodconfiguration"
	paymentsource "github.com/stripe/stripe-go/v82/paymentsource"
	payout "github.com/stripe/stripe-go/v82/payout"
	person "github.com/stripe/stripe-go/v82/person"
	plan "github.com/stripe/stripe-go/v82/plan"
	price "github.com/stripe/stripe-go/v82/price"
	product "github.com/stripe/stripe-go/v82/product"
	promotioncode "github.com/stripe/stripe-go/v82/promotioncode"
	quote "github.com/stripe/stripe-go/v82/quote"
	radar_earlyfraudwarning "github.com/stripe/stripe-go/v82/radar/earlyfraudwarning"
	radar_valuelist "github.com/stripe/stripe-go/v82/radar/valuelist"
	radar_valuelistitem "github.com/stripe/stripe-go/v82/radar/valuelistitem"
	refund "github.com/stripe/stripe-go/v82/refund"
	reporting_reportrun "github.com/stripe/stripe-go/v82/reporting/reportrun"
	reporting_reporttype "github.com/stripe/stripe-go/v82/reporting/reporttype"
	review "github.com/stripe/stripe-go/v82/review"
	setupattempt "github.com/stripe/stripe-go/v82/setupattempt"
	setupintent "github.com/stripe/stripe-go/v82/setupintent"
	shippingrate "github.com/stripe/stripe-go/v82/shippingrate"
	sigma_scheduledqueryrun "github.com/stripe/stripe-go/v82/sigma/scheduledqueryrun"
	source "github.com/stripe/stripe-go/v82/source"
	subscription "github.com/stripe/stripe-go/v82/subscription"
	subscriptionitem "github.com/stripe/stripe-go/v82/subscriptionitem"
	subscriptionschedule "github.com/stripe/stripe-go/v82/subscriptionschedule"
	tax_calculation "github.com/stripe/stripe-go/v82/tax/calculation"
	tax_form "github.com/stripe/stripe-go/v82/tax/form"
	tax_settings "github.com/stripe/stripe-go/v82/tax/settings"
	tax_transaction "github.com/stripe/stripe-go/v82/tax/transaction"
	taxcode "github.com/stripe/stripe-go/v82/taxcode"
	taxid "github.com/stripe/stripe-go/v82/taxid"
	taxrate "github.com/stripe/stripe-go/v82/taxrate"
	terminal_configuration "github.com/stripe/stripe-go/v82/terminal/configuration"
	terminal_connectiontoken "github.com/stripe/stripe-go/v82/terminal/connectiontoken"
	terminal_location "github.com/stripe/stripe-go/v82/terminal/location"
	terminal_reader "github.com/stripe/stripe-go/v82/terminal/reader"
	testhelpers_customer "github.com/stripe/stripe-go/v82/testhelpers/customer"
	testhelpers_issuing_authorization "github.com/stripe/stripe-go/v82/testhelpers/issuing/authorization"
	testhelpers_issuing_card "github.com/stripe/stripe-go/v82/testhelpers/issuing/card"
	testhelpers_issuing_personalizationdesign "github.com/stripe/stripe-go/v82/testhelpers/issuing/personalizationdesign"
	testhelpers_issuing_transaction "github.com/stripe/stripe-go/v82/testhelpers/issuing/transaction"
	testhelpers_refund "github.com/stripe/stripe-go/v82/testhelpers/refund"
	testhelpers_testclock "github.com/stripe/stripe-go/v82/testhelpers/testclock"
	testhelpers_treasury_inboundtransfer "github.com/stripe/stripe-go/v82/testhelpers/treasury/inboundtransfer"
	testhelpers_treasury_outboundtransfer "github.com/stripe/stripe-go/v82/testhelpers/treasury/outboundtransfer"
	testhelpers_treasury_receivedcredit "github.com/stripe/stripe-go/v82/testhelpers/treasury/receivedcredit"
	testhelpers_treasury_receiveddebit "github.com/stripe/stripe-go/v82/testhelpers/treasury/receiveddebit"
	. "github.com/stripe/stripe-go/v82/testing"
	token "github.com/stripe/stripe-go/v82/token"
	topup "github.com/stripe/stripe-go/v82/topup"
	transfer "github.com/stripe/stripe-go/v82/transfer"
	transferreversal "github.com/stripe/stripe-go/v82/transferreversal"
	treasury_creditreversal "github.com/stripe/stripe-go/v82/treasury/creditreversal"
	treasury_debitreversal "github.com/stripe/stripe-go/v82/treasury/debitreversal"
	treasury_financialaccount "github.com/stripe/stripe-go/v82/treasury/financialaccount"
	treasury_inboundtransfer "github.com/stripe/stripe-go/v82/treasury/inboundtransfer"
	treasury_outboundpayment "github.com/stripe/stripe-go/v82/treasury/outboundpayment"
	treasury_outboundtransfer "github.com/stripe/stripe-go/v82/treasury/outboundtransfer"
	treasury_receivedcredit "github.com/stripe/stripe-go/v82/treasury/receivedcredit"
	treasury_receiveddebit "github.com/stripe/stripe-go/v82/treasury/receiveddebit"
	treasury_transaction "github.com/stripe/stripe-go/v82/treasury/transaction"
	treasury_transactionentry "github.com/stripe/stripe-go/v82/treasury/transactionentry"
	webhookendpoint "github.com/stripe/stripe-go/v82/webhookendpoint"
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

func TestAccountLinksPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountLinkCreateParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, err := sc.V1AccountLinks.Create(context.TODO(), params)
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

func TestAccountsCapabilitiesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CapabilityListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result := sc.V1Capabilities.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestAccountsCapabilitiesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CapabilityRetrieveParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Capabilities.Retrieve(
		context.TODO(), "card_payments", params)
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

func TestAccountsCapabilitiesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CapabilityUpdateParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Capabilities.Update(
		context.TODO(), "card_payments", params)
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

func TestAccountsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountDeleteParams{}
	result, err := sc.V1Accounts.Delete(
		context.TODO(), "acct_xxxxxxxxxxxxx", params)
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

func TestAccountsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Accounts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestAccountsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountRetrieveParams{}
	result, err := sc.V1Accounts.GetByID(
		context.TODO(), "acct_xxxxxxxxxxxxx", params)
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

func TestAccountsLoginLinksPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.LoginLinkCreateParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1LoginLinks.Create(context.TODO(), params)
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

func TestAccountsPersonsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PersonDeleteParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Persons.Delete(
		context.TODO(), "person_xxxxxxxxxxxxx", params)
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

func TestAccountsPersonsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PersonListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1Persons.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestAccountsPersonsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PersonRetrieveParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Persons.Retrieve(
		context.TODO(), "person_xxxxxxxxxxxxx", params)
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

func TestAccountsPersonsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PersonCreateParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Persons.Create(context.TODO(), params)
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

func TestAccountsPersonsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PersonUpdateParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Persons.Update(
		context.TODO(), "person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsPost(t *testing.T) {
	params := &stripe.AccountParams{
		Type:    stripe.String(stripe.AccountTypeCustom),
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
		Type:    stripe.String(stripe.AccountTypeCustom),
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

func TestAccountsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountCreateParams{
		Type:    stripe.String(stripe.AccountTypeCustom),
		Country: stripe.String("US"),
		Email:   stripe.String("jenny.rosen@example.com"),
		Capabilities: &stripe.AccountCreateCapabilitiesParams{
			CardPayments: &stripe.AccountCreateCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCreateCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
	}
	result, err := sc.V1Accounts.Create(context.TODO(), params)
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

func TestAccountsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Accounts.Update(
		context.TODO(), "acct_xxxxxxxxxxxxx", params)
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

func TestAccountsRejectPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, err := sc.V1Accounts.Reject(
		context.TODO(), "acct_xxxxxxxxxxxxx", params)
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

func TestApplicationFeesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ApplicationFeeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1ApplicationFees.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestApplicationFeesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ApplicationFeeRetrieveParams{}
	result, err := sc.V1ApplicationFees.Retrieve(
		context.TODO(), "fee_xxxxxxxxxxxxx", params)
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

func TestApplicationFeesRefundsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FeeRefundListParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	params.Limit = stripe.Int64(3)
	result := sc.V1FeeRefunds.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestApplicationFeesRefundsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FeeRefundRetrieveParams{
		Fee: stripe.String("fee_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1FeeRefunds.Retrieve(
		context.TODO(), "fr_xxxxxxxxxxxxx", params)
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

func TestApplicationFeesRefundsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FeeRefundCreateParams{
		ID: stripe.String("fee_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1FeeRefunds.Create(context.TODO(), params)
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

func TestApplicationFeesRefundsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FeeRefundUpdateParams{
		Fee: stripe.String("fee_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1FeeRefunds.Update(
		context.TODO(), "fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsDeletePost(t *testing.T) {
	params := &stripe.AppsSecretDeleteWhereParams{
		Name: stripe.String("my-api-key"),
		Scope: &stripe.AppsSecretDeleteWhereScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.AppsSecrets.DeleteWhere(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsDeletePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretDeleteWhereParams{
		Name: stripe.String("my-api-key"),
		Scope: &stripe.AppsSecretDeleteWhereScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.V1AppsSecrets.DeleteWhere(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsFindGet(t *testing.T) {
	params := &stripe.AppsSecretFindParams{
		Name: stripe.String("sec_123"),
		Scope: &stripe.AppsSecretFindScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.AppsSecrets.Find(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsFindGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretFindParams{
		Name: stripe.String("sec_123"),
		Scope: &stripe.AppsSecretFindScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.V1AppsSecrets.Find(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsGet(t *testing.T) {
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.AppsSecrets.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.V1AppsSecrets.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestAppsSecretsGet2(t *testing.T) {
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.AppsSecrets.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAppsSecretsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretListParams{
		Scope: &stripe.AppsSecretListScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	params.Limit = stripe.Int64(2)
	result := sc.V1AppsSecrets.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestAppsSecretsPost(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("sec_123"),
		Payload: stripe.String("very secret string"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.AppsSecrets.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretCreateParams{
		Name:    stripe.String("sec_123"),
		Payload: stripe.String("very secret string"),
		Scope: &stripe.AppsSecretCreateScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.V1AppsSecrets.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPost2(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("my-api-key"),
		Payload: stripe.String("secret_key_xxxxxx"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
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
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.AppsSecrets.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.AppsSecretCreateParams{
		Name:    stripe.String("my-api-key"),
		Payload: stripe.String("secret_key_xxxxxx"),
		Scope: &stripe.AppsSecretCreateScopeParams{
			Type: stripe.String(stripe.AppsSecretScopeTypeAccount),
		},
	}
	result, err := sc.V1AppsSecrets.Create(context.TODO(), params)
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

func TestBalanceTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1BalanceTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestBalanceTransactionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BalanceTransactionRetrieveParams{}
	result, err := sc.V1BalanceTransactions.Retrieve(
		context.TODO(), "txn_xxxxxxxxxxxxx", params)
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

func TestBillingPortalConfigurationsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1BillingPortalConfigurations.List(context.TODO(), params)
	assert.NotNil(t, result)
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
	result, err := sc.BillingPortalConfigurations.Get("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BillingPortalConfigurationRetrieveParams{}
	result, err := sc.V1BillingPortalConfigurations.Retrieve(
		context.TODO(), "bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationsPost(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		Features: &stripe.BillingPortalConfigurationFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail),
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID),
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
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail),
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID),
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

func TestBillingPortalConfigurationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BillingPortalConfigurationCreateParams{
		Features: &stripe.BillingPortalConfigurationCreateFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationCreateFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail),
					stripe.String(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID),
				},
				Enabled: stripe.Bool(true),
			},
			InvoiceHistory: &stripe.BillingPortalConfigurationCreateFeaturesInvoiceHistoryParams{
				Enabled: stripe.Bool(true),
			},
		},
		BusinessProfile: &stripe.BillingPortalConfigurationCreateBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := sc.V1BillingPortalConfigurations.Create(context.TODO(), params)
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
	result, err := billingportal_configuration.Update("bpc_xxxxxxxxxxxxx", params)
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

func TestBillingPortalConfigurationsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BillingPortalConfigurationUpdateParams{
		BusinessProfile: &stripe.BillingPortalConfigurationUpdateBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := sc.V1BillingPortalConfigurations.Update(
		context.TODO(), "bpc_xxxxxxxxxxxxx", params)
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

func TestBillingPortalSessionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.BillingPortalSessionCreateParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, err := sc.V1BillingPortalSessions.Create(context.TODO(), params)
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

func TestChargesCapturePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeCaptureParams{}
	result, err := sc.V1Charges.Capture(
		context.TODO(), "ch_xxxxxxxxxxxxx", params)
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

func TestChargesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Charges.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestChargesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeRetrieveParams{}
	result, err := sc.V1Charges.Retrieve(
		context.TODO(), "ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPost(t *testing.T) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(stripe.CurrencyUSD),
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
		Currency:    stripe.String(stripe.CurrencyUSD),
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := sc.Charges.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeCreateParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(stripe.CurrencyUSD),
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := sc.V1Charges.Create(context.TODO(), params)
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

func TestChargesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Charges.Update(context.TODO(), "ch_xxxxxxxxxxxxx", params)
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

func TestChargesSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ChargeSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "amount>999 AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1Charges.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCheckoutSessionsExpirePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := sc.V1CheckoutSessions.Expire(
		context.TODO(), "sess_xyz", params)
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

func TestCheckoutSessionsExpirePost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := sc.V1CheckoutSessions.Expire(
		context.TODO(), "cs_test_xxxxxxxxxxxxx", params)
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

func TestCheckoutSessionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1CheckoutSessions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCheckoutSessionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionRetrieveParams{}
	result, err := sc.V1CheckoutSessions.Retrieve(
		context.TODO(), "cs_test_xxxxxxxxxxxxx", params)
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

func TestCheckoutSessionsLineItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionListLineItemsParams{
		Session: stripe.String("sess_xyz"),
	}
	result := sc.V1CheckoutSessions.ListLineItems(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionsPost(t *testing.T) {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
		Mode:       stripe.String(stripe.CheckoutSessionModePayment),
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
		Mode:       stripe.String(stripe.CheckoutSessionModePayment),
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

func TestCheckoutSessionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionCreateParams{
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
		Mode:       stripe.String(stripe.CheckoutSessionModePayment),
		ShippingOptions: []*stripe.CheckoutSessionCreateShippingOptionParams{
			{
				ShippingRate: stripe.String("shr_standard"),
			},
			{
				ShippingRateData: &stripe.CheckoutSessionCreateShippingOptionShippingRateDataParams{
					DisplayName: stripe.String("Standard"),
					DeliveryEstimate: &stripe.CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateParams{
						Minimum: &stripe.CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMinimumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(5),
						},
						Maximum: &stripe.CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMaximumParams{
							Unit:  stripe.String("day"),
							Value: stripe.Int64(7),
						},
					},
				},
			},
		},
	}
	result, err := sc.V1CheckoutSessions.Create(context.TODO(), params)
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
		Mode: stripe.String(stripe.CheckoutSessionModePayment),
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
		Mode: stripe.String(stripe.CheckoutSessionModePayment),
	}
	result, err := sc.CheckoutSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CheckoutSessionCreateParams{
		SuccessURL: stripe.String("https://example.com/success"),
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
		Mode: stripe.String(stripe.CheckoutSessionModePayment),
	}
	result, err := sc.V1CheckoutSessions.Create(context.TODO(), params)
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

func TestCountrySpecsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CountrySpecListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1CountrySpecs.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCountrySpecsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CountrySpecRetrieveParams{}
	result, err := sc.V1CountrySpecs.Retrieve(context.TODO(), "US", params)
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

func TestCouponsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CouponDeleteParams{}
	result, err := sc.V1Coupons.Delete(context.TODO(), "Z4OV52SU", params)
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

func TestCouponsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CouponListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Coupons.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCouponsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CouponRetrieveParams{}
	result, err := sc.V1Coupons.Retrieve(context.TODO(), "Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPost(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff: stripe.Float64(25.5),
		Duration:   stripe.String(stripe.CouponDurationOnce),
	}
	result, err := coupon.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.CouponParams{
		PercentOff: stripe.Float64(25.5),
		Duration:   stripe.String(stripe.CouponDurationOnce),
	}
	result, err := sc.Coupons.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CouponCreateParams{
		PercentOff: stripe.Float64(25.5),
		Duration:   stripe.String(stripe.CouponDurationOnce),
	}
	result, err := sc.V1Coupons.Create(context.TODO(), params)
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

func TestCouponsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CouponUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Coupons.Update(context.TODO(), "Z4OV52SU", params)
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

func TestCustomerSessionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerSessionCreateParams{
		Customer: stripe.String("cus_123"),
		Components: &stripe.CustomerSessionCreateComponentsParams{
			BuyButton: &stripe.CustomerSessionCreateComponentsBuyButtonParams{
				Enabled: stripe.Bool(true),
			},
		},
	}
	result, err := sc.V1CustomerSessions.Create(context.TODO(), params)
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

func TestCustomersBalanceTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1CustomerBalanceTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersBalanceTransactionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerBalanceTransactionRetrieveParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1CustomerBalanceTransactions.Retrieve(
		context.TODO(), "cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPost(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(stripe.CurrencyUSD),
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
		Currency: stripe.String(stripe.CurrencyUSD),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.CustomerBalanceTransactions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersBalanceTransactionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerBalanceTransactionCreateParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(stripe.CurrencyUSD),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1CustomerBalanceTransactions.Create(context.TODO(), params)
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

func TestCustomersBalanceTransactionsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerBalanceTransactionUpdateParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1CustomerBalanceTransactions.Update(
		context.TODO(), "cbtxn_xxxxxxxxxxxxx", params)
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

func TestCustomersCashBalanceGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CashBalanceRetrieveParams{
		Customer: stripe.String("cus_123"),
	}
	result, err := sc.V1CashBalances.Retrieve(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalancePost(t *testing.T) {
	params := &stripe.CashBalanceParams{
		Settings: &stripe.CashBalanceSettingsParams{
			ReconciliationMode: stripe.String(stripe.CashBalanceSettingsReconciliationModeManual),
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
			ReconciliationMode: stripe.String(stripe.CashBalanceSettingsReconciliationModeManual),
		},
		Customer: stripe.String("cus_123"),
	}
	result, err := sc.CashBalances.Update(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalancePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CashBalanceUpdateParams{
		Settings: &stripe.CashBalanceUpdateSettingsParams{
			ReconciliationMode: stripe.String(stripe.CashBalanceSettingsReconciliationModeManual),
		},
		Customer: stripe.String("cus_123"),
	}
	result, err := sc.V1CashBalances.Update(context.TODO(), params)
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

func TestCustomersCashBalanceTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerCashBalanceTransactionListParams{
		Customer: stripe.String("cus_123"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1CustomerCashBalanceTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerDeleteParams{}
	result, err := sc.V1Customers.Delete(
		context.TODO(), "cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersFundingInstructionsPost(t *testing.T) {
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(stripe.CurrencyUSD),
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
		Currency:    stripe.String(stripe.CurrencyUSD),
		FundingType: stripe.String("bank_transfer"),
	}
	result, err := sc.Customers.CreateFundingInstructions("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersFundingInstructionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(stripe.CurrencyUSD),
		FundingType: stripe.String("bank_transfer"),
	}
	result, err := sc.V1Customers.CreateFundingInstructions(
		context.TODO(), "cus_123", params)
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

func TestCustomersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Customers.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Customers.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerRetrieveParams{}
	result, err := sc.V1Customers.Retrieve(
		context.TODO(), "cus_xxxxxxxxxxxxx", params)
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

func TestCustomersPaymentMethodsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xyz"),
	}
	result := sc.V1Customers.ListPaymentMethods(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersPaymentMethodsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result := sc.V1Customers.ListPaymentMethods(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerCreateParams{
		Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := sc.V1Customers.Create(context.TODO(), params)
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

func TestCustomersPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Customers.Update(
		context.TODO(), "cus_xxxxxxxxxxxxx", params)
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

func TestCustomersSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := sc.V1Customers.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersSearchGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := sc.V1Customers.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersSourcesDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CardDeleteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Cards.Delete(context.TODO(), "ba_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesDelete2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CardDeleteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Cards.Delete(context.TODO(), "card_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("bank_account"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1PaymentSources.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersSourcesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentSourceListParams{
		Object:   stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1PaymentSources.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersSourcesGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentSourceRetrieveParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1PaymentSources.Retrieve(
		context.TODO(), "ba_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesGet4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentSourceRetrieveParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1PaymentSources.Retrieve(
		context.TODO(), "card_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CardUpdateParams{
		AccountHolderName: stripe.String("Kamil"),
		Customer:          stripe.String("cus_123"),
	}
	result, err := sc.V1Cards.Update(context.TODO(), "card_123", params)
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

func TestCustomersSourcesPost5Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CardUpdateParams{
		Name:     stripe.String("Jenny Rosen"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Cards.Update(context.TODO(), "card_xxxxxxxxxxxxx", params)
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

func TestCustomersTaxIdsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDDeleteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TaxIDs.Delete(context.TODO(), "txi_xxxxxxxxxxxxx", params)
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

func TestCustomersTaxIdsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TaxIDs.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestCustomersTaxIdsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDRetrieveParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TaxIDs.Retrieve(
		context.TODO(), "txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsPost(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:     stripe.String(stripe.TaxIDTypeEUVAT),
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
		Type:     stripe.String(stripe.TaxIDTypeEUVAT),
		Value:    stripe.String("DE123456789"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.TaxIDs.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersTaxIdsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDCreateParams{
		Type:     stripe.String(stripe.TaxIDTypeEUVAT),
		Value:    stripe.String("DE123456789"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TaxIDs.Create(context.TODO(), params)
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

func TestDisputesClosePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.DisputeCloseParams{}
	result, err := sc.V1Disputes.Close(context.TODO(), "dp_xxxxxxxxxxxxx", params)
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

func TestDisputesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Disputes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestDisputesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.DisputeRetrieveParams{}
	result, err := sc.V1Disputes.Retrieve(
		context.TODO(), "dp_xxxxxxxxxxxxx", params)
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

func TestDisputesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.DisputeUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Disputes.Update(
		context.TODO(), "dp_xxxxxxxxxxxxx", params)
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

func TestEventsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.EventListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Events.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestEventsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.EventRetrieveParams{}
	result, err := sc.V1Events.Retrieve(
		context.TODO(), "evt_xxxxxxxxxxxxx", params)
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

func TestFinancialConnectionsAccountsDisconnectPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := sc.V1FinancialConnectionsAccounts.Disconnect(
		context.TODO(), "fca_xyz", params)
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

func TestFinancialConnectionsAccountsDisconnectPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := sc.V1FinancialConnectionsAccounts.Disconnect(
		context.TODO(), "fca_xxxxxxxxxxxxx", params)
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

func TestFinancialConnectionsAccountsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountListParams{}
	result := sc.V1FinancialConnectionsAccounts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestFinancialConnectionsAccountsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountRetrieveParams{}
	result, err := sc.V1FinancialConnectionsAccounts.GetByID(
		context.TODO(), "fca_xyz", params)
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

func TestFinancialConnectionsAccountsGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountListParams{
		AccountHolder: &stripe.FinancialConnectionsAccountListAccountHolderParams{
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
	}
	result := sc.V1FinancialConnectionsAccounts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestFinancialConnectionsAccountsGet4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountRetrieveParams{}
	result, err := sc.V1FinancialConnectionsAccounts.GetByID(
		context.TODO(), "fca_xxxxxxxxxxxxx", params)
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

func TestFinancialConnectionsAccountsOwnersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xyz"),
		Account:   stripe.String("fca_xyz"),
	}
	result := sc.V1FinancialConnectionsAccounts.ListOwners(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestFinancialConnectionsAccountsOwnersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xxxxxxxxxxxxx"),
		Account:   stripe.String("fca_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1FinancialConnectionsAccounts.ListOwners(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestFinancialConnectionsAccountsRefreshPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountRefreshParams{
		Features: []*string{stripe.String("balance")},
	}
	result, err := sc.V1FinancialConnectionsAccounts.Refresh(
		context.TODO(), "fca_xyz", params)
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

func TestFinancialConnectionsAccountsSubscribePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountSubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := sc.V1FinancialConnectionsAccounts.Subscribe(
		context.TODO(), "fa_123", params)
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

func TestFinancialConnectionsAccountsUnsubscribePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsAccountUnsubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := sc.V1FinancialConnectionsAccounts.Unsubscribe(
		context.TODO(), "fa_123", params)
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

func TestFinancialConnectionsSessionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsSessionRetrieveParams{}
	result, err := sc.V1FinancialConnectionsSessions.Retrieve(
		context.TODO(), "fcsess_xyz", params)
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

func TestFinancialConnectionsSessionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsSessionRetrieveParams{}
	result, err := sc.V1FinancialConnectionsSessions.Retrieve(
		context.TODO(), "fcsess_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
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
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
		},
	}
	result, err := sc.FinancialConnectionsSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsSessionCreateParams{
		AccountHolder: &stripe.FinancialConnectionsSessionCreateAccountHolderParams{
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
		},
	}
	result, err := sc.V1FinancialConnectionsSessions.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionPaymentMethod),
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
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
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionPaymentMethod),
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
		},
		Filters: &stripe.FinancialConnectionsSessionFiltersParams{
			Countries: []*string{stripe.String("US")},
		},
	}
	result, err := sc.FinancialConnectionsSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsSessionCreateParams{
		AccountHolder: &stripe.FinancialConnectionsSessionCreateAccountHolderParams{
			Type:     stripe.String(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer),
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
		Permissions: []*string{
			stripe.String(stripe.FinancialConnectionsSessionPermissionPaymentMethod),
			stripe.String(stripe.FinancialConnectionsSessionPermissionBalances),
		},
		Filters: &stripe.FinancialConnectionsSessionCreateFiltersParams{
			Countries: []*string{stripe.String("US")},
		},
	}
	result, err := sc.V1FinancialConnectionsSessions.Create(
		context.TODO(), params)
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

func TestFinancialConnectionsTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsTransactionRetrieveParams{}
	result, err := sc.V1FinancialConnectionsTransactions.Retrieve(
		context.TODO(), "tr_123", params)
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

func TestFinancialConnectionsTransactionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.FinancialConnectionsTransactionListParams{
		Account: stripe.String("fca_xyz"),
	}
	result := sc.V1FinancialConnectionsTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIdentityVerificationReportsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationReportListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IdentityVerificationReports.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIdentityVerificationReportsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationReportRetrieveParams{}
	result, err := sc.V1IdentityVerificationReports.Retrieve(
		context.TODO(), "vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsCancelPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := identity_verificationsession.Cancel("vs_xxxxxxxxxxxxx", params)
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

func TestIdentityVerificationSessionsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := sc.V1IdentityVerificationSessions.Cancel(
		context.TODO(), "vs_xxxxxxxxxxxxx", params)
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

func TestIdentityVerificationSessionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IdentityVerificationSessions.List(context.TODO(), params)
	assert.NotNil(t, result)
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
	result, err := sc.IdentityVerificationSessions.Get("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionRetrieveParams{}
	result, err := sc.V1IdentityVerificationSessions.Retrieve(
		context.TODO(), "vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeDocument),
	}
	result, err := identity_verificationsession.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeDocument),
	}
	result, err := sc.IdentityVerificationSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionCreateParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeDocument),
	}
	result, err := sc.V1IdentityVerificationSessions.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost2(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeIDNumber),
	}
	result, err := identity_verificationsession.Update("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeIDNumber),
	}
	result, err := sc.IdentityVerificationSessions.Update(
		"vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionUpdateParams{
		Type: stripe.String(stripe.IdentityVerificationSessionTypeIDNumber),
	}
	result, err := sc.V1IdentityVerificationSessions.Update(
		context.TODO(), "vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsRedactPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := identity_verificationsession.Redact("vs_xxxxxxxxxxxxx", params)
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

func TestIdentityVerificationSessionsRedactPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := sc.V1IdentityVerificationSessions.Redact(
		context.TODO(), "vs_xxxxxxxxxxxxx", params)
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

func TestInvoiceitemsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceItemDeleteParams{}
	result, err := sc.V1InvoiceItems.Delete(
		context.TODO(), "ii_xxxxxxxxxxxxx", params)
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

func TestInvoiceitemsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceItemCreateParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1InvoiceItems.Create(context.TODO(), params)
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

func TestInvoiceitemsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceItemUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1InvoiceItems.Update(
		context.TODO(), "ii_xxxxxxxxxxxxx", params)
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

func TestInvoicesDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceDeleteParams{}
	result, err := sc.V1Invoices.Delete(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesFinalizePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceFinalizeInvoiceParams{}
	result, err := sc.V1Invoices.FinalizeInvoice(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Invoices.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestInvoicesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceRetrieveParams{}
	result, err := sc.V1Invoices.Retrieve(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceRetrieveParams{}
	params.AddExpand("customer")
	result, err := sc.V1Invoices.Retrieve(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesPayPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoicePayParams{}
	result, err := sc.V1Invoices.Pay(context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceCreateParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Invoices.Create(context.TODO(), params)
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

func TestInvoicesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Invoices.Update(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "total>999 AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1Invoices.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestInvoicesSendPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceSendInvoiceParams{}
	result, err := sc.V1Invoices.SendInvoice(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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

func TestInvoicesVoidPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.InvoiceVoidInvoiceParams{}
	result, err := sc.V1Invoices.VoidInvoice(
		context.TODO(), "in_xxxxxxxxxxxxx", params)
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
	result, err := sc.IssuingAuthorizations.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsApprovePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, err := sc.V1IssuingAuthorizations.Approve(
		context.TODO(), "iauth_xxxxxxxxxxxxx", params)
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
	result, err := sc.IssuingAuthorizations.Decline("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsDeclinePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, err := sc.V1IssuingAuthorizations.Decline(
		context.TODO(), "iauth_xxxxxxxxxxxxx", params)
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

func TestIssuingAuthorizationsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingAuthorizationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IssuingAuthorizations.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingAuthorizationsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingAuthorizationRetrieveParams{}
	result, err := sc.V1IssuingAuthorizations.Retrieve(
		context.TODO(), "iauth_xxxxxxxxxxxxx", params)
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

func TestIssuingAuthorizationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingAuthorizationUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1IssuingAuthorizations.Update(
		context.TODO(), "iauth_xxxxxxxxxxxxx", params)
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

func TestIssuingCardholdersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardholderListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IssuingCardholders.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingCardholdersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardholderRetrieveParams{}
	result, err := sc.V1IssuingCardholders.Retrieve(
		context.TODO(), "ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholdersPost(t *testing.T) {
	params := &stripe.IssuingCardholderParams{
		Type:        stripe.String(stripe.IssuingCardholderTypeIndividual),
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
		Type:        stripe.String(stripe.IssuingCardholderTypeIndividual),
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

func TestIssuingCardholdersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardholderCreateParams{
		Type:        stripe.String(stripe.IssuingCardholderTypeIndividual),
		Name:        stripe.String("Jenny Rosen"),
		Email:       stripe.String("jenny.rosen@example.com"),
		PhoneNumber: stripe.String("+18888675309"),
		Billing: &stripe.IssuingCardholderCreateBillingParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("1234 Main Street"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
				PostalCode: stripe.String("94111"),
			},
		},
	}
	result, err := sc.V1IssuingCardholders.Create(context.TODO(), params)
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

func TestIssuingCardholdersPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardholderUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1IssuingCardholders.Update(
		context.TODO(), "ich_xxxxxxxxxxxxx", params)
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

func TestIssuingCardsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IssuingCards.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingCardsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardRetrieveParams{}
	result, err := sc.V1IssuingCards.Retrieve(
		context.TODO(), "ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPost(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(stripe.CurrencyUSD),
		Type:       stripe.String(stripe.IssuingCardTypeVirtual),
	}
	result, err := issuing_card.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(stripe.CurrencyUSD),
		Type:       stripe.String(stripe.IssuingCardTypeVirtual),
	}
	result, err := sc.IssuingCards.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardCreateParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(stripe.CurrencyUSD),
		Type:       stripe.String(stripe.IssuingCardTypeVirtual),
	}
	result, err := sc.V1IssuingCards.Create(context.TODO(), params)
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

func TestIssuingCardsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingCardUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1IssuingCards.Update(
		context.TODO(), "ic_xxxxxxxxxxxxx", params)
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

func TestIssuingDisputesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingDisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IssuingDisputes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingDisputesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingDisputeRetrieveParams{}
	result, err := sc.V1IssuingDisputes.Retrieve(
		context.TODO(), "idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesPost(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(stripe.IssuingDisputeEvidenceReasonFraudulent),
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
			Reason: stripe.String(stripe.IssuingDisputeEvidenceReasonFraudulent),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, err := sc.IssuingDisputes.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingDisputeCreateParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeCreateEvidenceParams{
			Reason: stripe.String(stripe.IssuingDisputeEvidenceReasonFraudulent),
			Fraudulent: &stripe.IssuingDisputeCreateEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, err := sc.V1IssuingDisputes.Create(context.TODO(), params)
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

func TestIssuingDisputesSubmitPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingDisputeSubmitParams{}
	result, err := sc.V1IssuingDisputes.Submit(
		context.TODO(), "idp_xxxxxxxxxxxxx", params)
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

func TestIssuingPersonalizationDesignsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPersonalizationDesignListParams{}
	result := sc.V1IssuingPersonalizationDesigns.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingPersonalizationDesignsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPersonalizationDesignRetrieveParams{}
	result, err := sc.V1IssuingPersonalizationDesigns.Retrieve(
		context.TODO(), "pd_xyz", params)
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

func TestIssuingPersonalizationDesignsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPersonalizationDesignCreateParams{
		PhysicalBundle: stripe.String("pb_xyz"),
	}
	result, err := sc.V1IssuingPersonalizationDesigns.Create(
		context.TODO(), params)
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

func TestIssuingPersonalizationDesignsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPersonalizationDesignUpdateParams{}
	result, err := sc.V1IssuingPersonalizationDesigns.Update(
		context.TODO(), "pd_xyz", params)
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

func TestIssuingPhysicalBundlesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPhysicalBundleListParams{}
	result := sc.V1IssuingPhysicalBundles.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingPhysicalBundlesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingPhysicalBundleRetrieveParams{}
	result, err := sc.V1IssuingPhysicalBundles.Retrieve(
		context.TODO(), "pb_xyz", params)
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

func TestIssuingTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1IssuingTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestIssuingTransactionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingTransactionRetrieveParams{}
	result, err := sc.V1IssuingTransactions.Retrieve(
		context.TODO(), "ipi_xxxxxxxxxxxxx", params)
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

func TestIssuingTransactionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.IssuingTransactionUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1IssuingTransactions.Update(
		context.TODO(), "ipi_xxxxxxxxxxxxx", params)
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

func TestMandatesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.MandateRetrieveParams{}
	result, err := sc.V1Mandates.Retrieve(
		context.TODO(), "mandate_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsApplyCustomerBalancePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, err := sc.V1PaymentIntents.ApplyCustomerBalance(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentCancelParams{}
	result, err := sc.V1PaymentIntents.Cancel(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsCapturePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := sc.V1PaymentIntents.Capture(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsConfirmPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := sc.V1PaymentIntents.Confirm(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1PaymentIntents.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentIntentsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentRetrieveParams{}
	result, err := sc.V1PaymentIntents.Retrieve(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsIncrementAuthorizationPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentIncrementAuthorizationParams{
		Amount: stripe.Int64(2099),
	}
	result, err := sc.V1PaymentIntents.IncrementAuthorization(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(stripe.CurrencyEUR),
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
		Currency: stripe.String(stripe.CurrencyEUR),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentCreateParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(stripe.CurrencyEUR),
		AutomaticPaymentMethods: &stripe.PaymentIntentCreateAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.V1PaymentIntents.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost2(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(stripe.CurrencyUSD),
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
		Currency: stripe.String(stripe.CurrencyUSD),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentCreateParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(stripe.CurrencyUSD),
		AutomaticPaymentMethods: &stripe.PaymentIntentCreateAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := sc.V1PaymentIntents.Create(context.TODO(), params)
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

func TestPaymentIntentsPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1PaymentIntents.Update(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost4(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(200),
		Currency: stripe.String(stripe.CurrencyUSD),
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
		Currency: stripe.String(stripe.CurrencyUSD),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("p24"),
			P24:  &stripe.PaymentMethodP24Params{Bank: stripe.String("blik")},
		},
	}
	result, err := sc.PaymentIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsPost4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentCreateParams{
		Amount:   stripe.Int64(200),
		Currency: stripe.String(stripe.CurrencyUSD),
		PaymentMethodData: &stripe.PaymentIntentCreatePaymentMethodDataParams{
			Type: stripe.String("p24"),
			P24:  &stripe.PaymentMethodP24Params{Bank: stripe.String("blik")},
		},
	}
	result, err := sc.V1PaymentIntents.Create(context.TODO(), params)
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

func TestPaymentIntentsSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'succeeded' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1PaymentIntents.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentIntentsVerifyMicrodepositsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, err := sc.V1PaymentIntents.VerifyMicrodeposits(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsVerifyMicrodepositsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := sc.V1PaymentIntents.VerifyMicrodeposits(
		context.TODO(), "pi_xxxxxxxxxxxxx", params)
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

func TestPaymentLinksGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkRetrieveParams{}
	result, err := sc.V1PaymentLinks.Retrieve(context.TODO(), "pl_xyz", params)
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

func TestPaymentLinksGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1PaymentLinks.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentLinksGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkRetrieveParams{}
	result, err := sc.V1PaymentLinks.Retrieve(
		context.TODO(), "plink_xxxxxxxxxxxxx", params)
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

func TestPaymentLinksLineItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkListLineItemsParams{
		PaymentLink: stripe.String("pl_xyz"),
	}
	result := sc.V1PaymentLinks.ListLineItems(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentLinksPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkCreateParams{
		LineItems: []*stripe.PaymentLinkCreateLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := sc.V1PaymentLinks.Create(context.TODO(), params)
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

func TestPaymentLinksPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkCreateParams{
		LineItems: []*stripe.PaymentLinkCreateLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, err := sc.V1PaymentLinks.Create(context.TODO(), params)
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

func TestPaymentLinksPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentLinkUpdateParams{Active: stripe.Bool(false)}
	result, err := sc.V1PaymentLinks.Update(
		context.TODO(), "plink_xxxxxxxxxxxxx", params)
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

func TestPaymentMethodConfigurationsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodConfigurationListParams{
		Application: stripe.String("foo"),
	}
	result := sc.V1PaymentMethodConfigurations.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentMethodConfigurationsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodConfigurationRetrieveParams{}
	result, err := sc.V1PaymentMethodConfigurations.Retrieve(
		context.TODO(), "foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceNone),
			},
		},
		Affirm: &stripe.PaymentMethodConfigurationAffirmParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationAffirmDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationAffirmDisplayPreferencePreferenceNone),
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
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceNone),
			},
		},
		Affirm: &stripe.PaymentMethodConfigurationAffirmParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationAffirmDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationAffirmDisplayPreferencePreferenceNone),
			},
		},
	}
	result, err := sc.PaymentMethodConfigurations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodConfigurationCreateParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationCreateACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationCreateACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceNone),
			},
		},
		Affirm: &stripe.PaymentMethodConfigurationCreateAffirmParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationCreateAffirmDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationAffirmDisplayPreferencePreferenceNone),
			},
		},
	}
	result, err := sc.V1PaymentMethodConfigurations.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost2(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceOn),
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
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceOn),
			},
		},
	}
	result, err := sc.PaymentMethodConfigurations.Update("foo", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodConfigurationsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodConfigurationUpdateParams{
		ACSSDebit: &stripe.PaymentMethodConfigurationUpdateACSSDebitParams{
			DisplayPreference: &stripe.PaymentMethodConfigurationUpdateACSSDebitDisplayPreferenceParams{
				Preference: stripe.String(stripe.PaymentMethodConfigurationACSSDebitDisplayPreferencePreferenceOn),
			},
		},
	}
	result, err := sc.V1PaymentMethodConfigurations.Update(
		context.TODO(), "foo", params)
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

func TestPaymentMethodsAttachPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1PaymentMethods.Attach(
		context.TODO(), "pm_xxxxxxxxxxxxx", params)
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

func TestPaymentMethodsDetachPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodDetachParams{}
	result, err := sc.V1PaymentMethods.Detach(
		context.TODO(), "pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsGet(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(stripe.PaymentMethodTypeCard),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(stripe.PaymentMethodTypeCard),
	}
	result := sc.PaymentMethods.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(stripe.PaymentMethodTypeCard),
	}
	result := sc.V1PaymentMethods.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPaymentMethodsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodRetrieveParams{}
	result, err := sc.V1PaymentMethods.Retrieve(
		context.TODO(), "pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsPost(t *testing.T) {
	params := &stripe.PaymentMethodParams{
		Type: stripe.String(stripe.PaymentMethodTypeCard),
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
		Type: stripe.String(stripe.PaymentMethodTypeCard),
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

func TestPaymentMethodsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodCreateParams{
		Type: stripe.String(stripe.PaymentMethodTypeCard),
		Card: &stripe.PaymentMethodCreateCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(8),
			ExpYear:  stripe.Int64(2024),
			CVC:      stripe.String("314"),
		},
	}
	result, err := sc.V1PaymentMethods.Create(context.TODO(), params)
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

func TestPaymentMethodsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PaymentMethodUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1PaymentMethods.Update(
		context.TODO(), "pm_xxxxxxxxxxxxx", params)
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

func TestPayoutsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutCancelParams{}
	result, err := sc.V1Payouts.Cancel(context.TODO(), "po_xxxxxxxxxxxxx", params)
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

func TestPayoutsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Payouts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPayoutsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutRetrieveParams{}
	result, err := sc.V1Payouts.Retrieve(
		context.TODO(), "po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPost(t *testing.T) {
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(stripe.CurrencyUSD),
	}
	result, err := payout.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.Payouts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutCreateParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.V1Payouts.Create(context.TODO(), params)
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

func TestPayoutsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Payouts.Update(context.TODO(), "po_xxxxxxxxxxxxx", params)
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

func TestPayoutsReversePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PayoutReverseParams{}
	result, err := sc.V1Payouts.Reverse(
		context.TODO(), "po_xxxxxxxxxxxxx", params)
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

func TestPlansDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PlanDeleteParams{}
	result, err := sc.V1Plans.Delete(
		context.TODO(), "price_xxxxxxxxxxxxx", params)
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

func TestPlansGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PlanListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Plans.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPlansGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PlanRetrieveParams{}
	result, err := sc.V1Plans.Retrieve(
		context.TODO(), "price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPost(t *testing.T) {
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(stripe.CurrencyUSD),
		Interval: stripe.String(stripe.PlanIntervalMonth),
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
		Currency: stripe.String(stripe.CurrencyUSD),
		Interval: stripe.String(stripe.PlanIntervalMonth),
		Product:  &stripe.PlanProductParams{Name: stripe.String("My product")},
	}
	result, err := sc.Plans.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PlanCreateParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(stripe.CurrencyUSD),
		Interval: stripe.String(stripe.PlanIntervalMonth),
		Product:  &stripe.PlanCreateProductParams{Name: stripe.String("My product")},
	}
	result, err := sc.V1Plans.Create(context.TODO(), params)
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

func TestPlansPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PlanUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Plans.Update(
		context.TODO(), "price_xxxxxxxxxxxxx", params)
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

func TestPricesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Prices.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPricesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceRetrieveParams{}
	result, err := sc.V1Prices.Retrieve(
		context.TODO(), "price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(stripe.CurrencyUSD),
		CurrencyOptions: map[string]*stripe.PriceCurrencyOptionsParams{
			"uah": {UnitAmount: stripe.Int64(5000)},
			"eur": {UnitAmount: stripe.Int64(1800)},
		},
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
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
		Currency:   stripe.String(stripe.CurrencyUSD),
		CurrencyOptions: map[string]*stripe.PriceCurrencyOptionsParams{
			"uah": {UnitAmount: stripe.Int64(5000)},
			"eur": {UnitAmount: stripe.Int64(1800)},
		},
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.Prices.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceCreateParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(stripe.CurrencyUSD),
		CurrencyOptions: map[string]*stripe.PriceCreateCurrencyOptionsParams{
			"uah": {
				UnitAmount: stripe.Int64(5000),
			},
			"eur": {
				UnitAmount: stripe.Int64(1800),
			},
		},
		Recurring: &stripe.PriceCreateRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Prices.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost2(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(stripe.CurrencyUSD),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
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
		Currency:   stripe.String(stripe.CurrencyUSD),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.Prices.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPricesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceCreateParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(stripe.CurrencyUSD),
		Recurring: &stripe.PriceCreateRecurringParams{
			Interval: stripe.String(stripe.PriceRecurringIntervalMonth),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Prices.Create(context.TODO(), params)
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

func TestPricesPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Prices.Update(
		context.TODO(), "price_xxxxxxxxxxxxx", params)
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

func TestPricesSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PriceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1Prices.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestProductsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductDeleteParams{}
	result, err := sc.V1Products.Delete(
		context.TODO(), "prod_xxxxxxxxxxxxx", params)
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

func TestProductsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Products.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestProductsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductRetrieveParams{}
	result, err := sc.V1Products.Retrieve(
		context.TODO(), "prod_xxxxxxxxxxxxx", params)
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

func TestProductsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductCreateParams{Name: stripe.String("Gold Special")}
	result, err := sc.V1Products.Create(context.TODO(), params)
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

func TestProductsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Products.Update(
		context.TODO(), "prod_xxxxxxxxxxxxx", params)
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

func TestProductsSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1Products.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPromotionCodesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PromotionCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1PromotionCodes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestPromotionCodesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PromotionCodeRetrieveParams{}
	result, err := sc.V1PromotionCodes.Retrieve(
		context.TODO(), "promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPost(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.PromotionCodes.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.PromotionCodeUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1PromotionCodes.Update(
		context.TODO(), "promo_xxxxxxxxxxxxx", params)
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

func TestQuotesAcceptPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteAcceptParams{}
	result, err := sc.V1Quotes.Accept(context.TODO(), "qt_xxxxxxxxxxxxx", params)
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

func TestQuotesCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteCancelParams{}
	result, err := sc.V1Quotes.Cancel(context.TODO(), "qt_xxxxxxxxxxxxx", params)
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

func TestQuotesFinalizePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, err := sc.V1Quotes.FinalizeQuote(
		context.TODO(), "qt_xxxxxxxxxxxxx", params)
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

func TestQuotesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Quotes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestQuotesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteRetrieveParams{}
	result, err := sc.V1Quotes.Retrieve(
		context.TODO(), "qt_xxxxxxxxxxxxx", params)
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

func TestQuotesLineItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteListLineItemsParams{
		Quote: stripe.String("qt_xxxxxxxxxxxxx"),
	}
	result := sc.V1Quotes.ListLineItems(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestQuotesPdfGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuotePDFParams{}
	result, err := sc.V1Quotes.PDF(context.TODO(), "qt_xxxxxxxxxxxxx", params)
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

func TestQuotesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteCreateParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		LineItems: []*stripe.QuoteCreateLineItemParams{
			{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
	}
	result, err := sc.V1Quotes.Create(context.TODO(), params)
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

func TestQuotesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Quotes.Update(context.TODO(), "qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesPreviewInvoicesLinesGet(t *testing.T) {
	params := &stripe.QuoteListPreviewInvoiceLinesParams{
		Quote:          stripe.String("qt_xyz"),
		PreviewInvoice: stripe.String("in_xyz"),
	}
	result := quote.ListPreviewInvoiceLines(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesPreviewInvoicesLinesGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.QuoteListPreviewInvoiceLinesParams{
		Quote:          stripe.String("qt_xyz"),
		PreviewInvoice: stripe.String("in_xyz"),
	}
	result := sc.Quotes.ListPreviewInvoiceLines(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuotesPreviewInvoicesLinesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.QuoteListPreviewInvoiceLinesParams{
		Quote:          stripe.String("qt_xyz"),
		PreviewInvoice: stripe.String("in_xyz"),
	}
	result := sc.V1Quotes.ListPreviewInvoiceLines(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestRadarEarlyFraudWarningsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1RadarEarlyFraudWarnings.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestRadarEarlyFraudWarningsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarEarlyFraudWarningRetrieveParams{}
	result, err := sc.V1RadarEarlyFraudWarnings.Retrieve(
		context.TODO(), "issfr_xxxxxxxxxxxxx", params)
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

func TestRadarValueListItemsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListItemDeleteParams{}
	result, err := sc.V1RadarValueListItems.Delete(
		context.TODO(), "rsli_xxxxxxxxxxxxx", params)
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

func TestRadarValueListItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListItemListParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1RadarValueListItems.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestRadarValueListItemsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListItemRetrieveParams{}
	result, err := sc.V1RadarValueListItems.Retrieve(
		context.TODO(), "rsli_xxxxxxxxxxxxx", params)
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

func TestRadarValueListItemsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListItemCreateParams{
		ValueList: stripe.String("rsl_xxxxxxxxxxxxx"),
		Value:     stripe.String("1.2.3.4"),
	}
	result, err := sc.V1RadarValueListItems.Create(context.TODO(), params)
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

func TestRadarValueListsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListDeleteParams{}
	result, err := sc.V1RadarValueLists.Delete(
		context.TODO(), "rsl_xxxxxxxxxxxxx", params)
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

func TestRadarValueListsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1RadarValueLists.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestRadarValueListsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListRetrieveParams{}
	result, err := sc.V1RadarValueLists.Retrieve(
		context.TODO(), "rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPost(t *testing.T) {
	params := &stripe.RadarValueListParams{
		Alias:    stripe.String("custom_ip_xxxxxxxxxxxxx"),
		Name:     stripe.String("Custom IP Blocklist"),
		ItemType: stripe.String(stripe.RadarValueListItemTypeIPAddress),
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
		ItemType: stripe.String(stripe.RadarValueListItemTypeIPAddress),
	}
	result, err := sc.RadarValueLists.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListCreateParams{
		Alias:    stripe.String("custom_ip_xxxxxxxxxxxxx"),
		Name:     stripe.String("Custom IP Blocklist"),
		ItemType: stripe.String(stripe.RadarValueListItemTypeIPAddress),
	}
	result, err := sc.V1RadarValueLists.Create(context.TODO(), params)
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

func TestRadarValueListsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RadarValueListUpdateParams{
		Name: stripe.String("Updated IP Block List"),
	}
	result, err := sc.V1RadarValueLists.Update(
		context.TODO(), "rsl_xxxxxxxxxxxxx", params)
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

func TestRefundsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RefundCancelParams{}
	result, err := sc.V1Refunds.Cancel(context.TODO(), "re_xxxxxxxxxxxxx", params)
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

func TestRefundsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Refunds.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestRefundsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RefundRetrieveParams{}
	result, err := sc.V1Refunds.Retrieve(
		context.TODO(), "re_xxxxxxxxxxxxx", params)
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

func TestRefundsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RefundCreateParams{
		Charge: stripe.String("ch_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1Refunds.Create(context.TODO(), params)
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

func TestRefundsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.RefundUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Refunds.Update(context.TODO(), "re_xxxxxxxxxxxxx", params)
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

func TestReportingReportRunsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReportingReportRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1ReportingReportRuns.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestReportingReportRunsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReportingReportRunRetrieveParams{}
	result, err := sc.V1ReportingReportRuns.Retrieve(
		context.TODO(), "frr_xxxxxxxxxxxxx", params)
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

func TestReportingReportRunsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReportingReportRunCreateParams{
		ReportType: stripe.String("balance.summary.1"),
		Parameters: &stripe.ReportingReportRunCreateParametersParams{
			IntervalStart: stripe.Int64(1522540800),
			IntervalEnd:   stripe.Int64(1525132800),
		},
	}
	result, err := sc.V1ReportingReportRuns.Create(context.TODO(), params)
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

func TestReportingReportTypesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReportingReportTypeListParams{}
	result := sc.V1ReportingReportTypes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestReportingReportTypesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReportingReportTypeRetrieveParams{}
	result, err := sc.V1ReportingReportTypes.Retrieve(
		context.TODO(), "balance.summary.1", params)
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

func TestReviewsApprovePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReviewApproveParams{}
	result, err := sc.V1Reviews.Approve(
		context.TODO(), "prv_xxxxxxxxxxxxx", params)
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

func TestReviewsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReviewListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Reviews.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestReviewsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ReviewRetrieveParams{}
	result, err := sc.V1Reviews.Retrieve(
		context.TODO(), "prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupAttemptsGet(t *testing.T) {
	params := &stripe.SetupAttemptListParams{SetupIntent: stripe.String("si_xyz")}
	params.Limit = stripe.Int64(3)
	result := setupattempt.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupAttemptsGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SetupAttemptListParams{SetupIntent: stripe.String("si_xyz")}
	params.Limit = stripe.Int64(3)
	result := sc.SetupAttempts.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupAttemptsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupAttemptListParams{SetupIntent: stripe.String("si_xyz")}
	params.Limit = stripe.Int64(3)
	result := sc.V1SetupAttempts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestSetupIntentsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentCancelParams{}
	result, err := sc.V1SetupIntents.Cancel(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsConfirmPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := sc.V1SetupIntents.Confirm(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1SetupIntents.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestSetupIntentsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentRetrieveParams{}
	result, err := sc.V1SetupIntents.Retrieve(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentCreateParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, err := sc.V1SetupIntents.Create(context.TODO(), params)
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

func TestSetupIntentsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentUpdateParams{}
	params.AddMetadata("user_id", "3435453")
	result, err := sc.V1SetupIntents.Update(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsVerifyMicrodepositsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, err := sc.V1SetupIntents.VerifyMicrodeposits(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsVerifyMicrodepositsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SetupIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := sc.V1SetupIntents.VerifyMicrodeposits(
		context.TODO(), "seti_xxxxxxxxxxxxx", params)
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

func TestShippingRatesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateListParams{}
	result := sc.V1ShippingRates.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestShippingRatesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1ShippingRates.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestShippingRatesGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateRetrieveParams{}
	result, err := sc.V1ShippingRates.Retrieve(
		context.TODO(), "shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Currency: stripe.String(stripe.CurrencyUSD),
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
			Currency: stripe.String(stripe.CurrencyUSD),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, err := sc.ShippingRates.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateCreateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateCreateFixedAmountParams{
			Currency: stripe.String(stripe.CurrencyUSD),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, err := sc.V1ShippingRates.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost2(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Ground shipping"),
		Type:        stripe.String("fixed_amount"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Amount:   stripe.Int64(500),
			Currency: stripe.String(stripe.CurrencyUSD),
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
			Currency: stripe.String(stripe.CurrencyUSD),
		},
	}
	result, err := sc.ShippingRates.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRatesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateCreateParams{
		DisplayName: stripe.String("Ground shipping"),
		Type:        stripe.String("fixed_amount"),
		FixedAmount: &stripe.ShippingRateCreateFixedAmountParams{
			Amount:   stripe.Int64(500),
			Currency: stripe.String(stripe.CurrencyUSD),
		},
	}
	result, err := sc.V1ShippingRates.Create(context.TODO(), params)
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

func TestShippingRatesPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.ShippingRateUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1ShippingRates.Update(
		context.TODO(), "shr_xxxxxxxxxxxxx", params)
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

func TestSigmaScheduledQueryRunsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1SigmaScheduledQueryRuns.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestSigmaScheduledQueryRunsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SigmaScheduledQueryRunRetrieveParams{}
	result, err := sc.V1SigmaScheduledQueryRuns.Retrieve(
		context.TODO(), "sqr_xxxxxxxxxxxxx", params)
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

func TestSourcesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SourceRetrieveParams{}
	result, err := sc.V1Sources.Retrieve(
		context.TODO(), "src_xxxxxxxxxxxxx", params)
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

func TestSourcesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SourceRetrieveParams{}
	result, err := sc.V1Sources.Retrieve(
		context.TODO(), "src_xxxxxxxxxxxxx", params)
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

func TestSourcesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SourceUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Sources.Update(
		context.TODO(), "src_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionItemDeleteParams{}
	result, err := sc.V1SubscriptionItems.Delete(
		context.TODO(), "si_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionItemListParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
	}
	result := sc.V1SubscriptionItems.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestSubscriptionItemsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionItemRetrieveParams{}
	result, err := sc.V1SubscriptionItems.Retrieve(
		context.TODO(), "si_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionItemCreateParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
		Price:        stripe.String("price_xxxxxxxxxxxxx"),
		Quantity:     stripe.Int64(2),
	}
	result, err := sc.V1SubscriptionItems.Create(context.TODO(), params)
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

func TestSubscriptionItemsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionItemUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1SubscriptionItems.Update(
		context.TODO(), "si_xxxxxxxxxxxxx", params)
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

func TestSubscriptionSchedulesCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleCancelParams{}
	result, err := sc.V1SubscriptionSchedules.Cancel(
		context.TODO(), "sub_sched_xxxxxxxxxxxxx", params)
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

func TestSubscriptionSchedulesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1SubscriptionSchedules.List(context.TODO(), params)
	assert.NotNil(t, result)
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
	result, err := sc.SubscriptionSchedules.Get("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleRetrieveParams{}
	result, err := sc.V1SubscriptionSchedules.Retrieve(
		context.TODO(), "sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		Customer:    stripe.String("cus_xxxxxxxxxxxxx"),
		StartDate:   stripe.Int64(1676070661),
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Price:    stripe.String("price_xxxxxxxxxxxxx"),
						Quantity: stripe.Int64(1),
					},
				},
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
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Price:    stripe.String("price_xxxxxxxxxxxxx"),
						Quantity: stripe.Int64(1),
					},
				},
			},
		},
	}
	result, err := sc.SubscriptionSchedules.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleCreateParams{
		Customer:    stripe.String("cus_xxxxxxxxxxxxx"),
		StartDate:   stripe.Int64(1676070661),
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
		Phases: []*stripe.SubscriptionScheduleCreatePhaseParams{
			{
				Items: []*stripe.SubscriptionScheduleCreatePhaseItemParams{
					{
						Price:    stripe.String("price_xxxxxxxxxxxxx"),
						Quantity: stripe.Int64(1),
					},
				},
			},
		},
	}
	result, err := sc.V1SubscriptionSchedules.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost2(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
	}
	result, err := subscriptionschedule.Update("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost2Service(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
	}
	result, err := sc.SubscriptionSchedules.Update(
		"sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleUpdateParams{
		EndBehavior: stripe.String(stripe.SubscriptionScheduleEndBehaviorRelease),
	}
	result, err := sc.V1SubscriptionSchedules.Update(
		context.TODO(), "sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesReleasePost(t *testing.T) {
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := subscriptionschedule.Release("sub_sched_xxxxxxxxxxxxx", params)
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

func TestSubscriptionSchedulesReleasePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := sc.V1SubscriptionSchedules.Release(
		context.TODO(), "sub_sched_xxxxxxxxxxxxx", params)
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

func TestSubscriptionsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionCancelParams{}
	result, err := sc.V1Subscriptions.Cancel(
		context.TODO(), "sub_xxxxxxxxxxxxx", params)
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

func TestSubscriptionsDiscountDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionDeleteDiscountParams{}
	result, err := sc.V1Subscriptions.DeleteDiscount(
		context.TODO(), "sub_xyz", params)
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

func TestSubscriptionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Subscriptions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestSubscriptionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionRetrieveParams{}
	result, err := sc.V1Subscriptions.Retrieve(
		context.TODO(), "sub_xxxxxxxxxxxxx", params)
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

func TestSubscriptionsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionCreateParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Items: []*stripe.SubscriptionCreateItemParams{
			{
				Price: stripe.String("price_xxxxxxxxxxxxx"),
			},
		},
	}
	result, err := sc.V1Subscriptions.Create(context.TODO(), params)
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

func TestSubscriptionsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Subscriptions.Update(
		context.TODO(), "sub_xxxxxxxxxxxxx", params)
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

func TestSubscriptionsSearchGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'active' AND metadata['order_id']:'6735'",
		},
	}
	result := sc.V1Subscriptions.Search(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTaxCalculationsLineItemsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxCalculationListLineItemsParams{
		Calculation: stripe.String("xxx"),
	}
	result := sc.V1TaxCalculations.ListLineItems(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestTaxCalculationsPost(t *testing.T) {
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(stripe.CurrencyUSD),
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
			AddressSource: stripe.String(stripe.TaxCalculationCustomerDetailsAddressSourceShipping),
		},
	}
	result, err := tax_calculation.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCalculationsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(stripe.CurrencyUSD),
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
			AddressSource: stripe.String(stripe.TaxCalculationCustomerDetailsAddressSourceShipping),
		},
	}
	result, err := sc.TaxCalculations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCalculationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxCalculationCreateParams{
		Currency: stripe.String(stripe.CurrencyUSD),
		LineItems: []*stripe.TaxCalculationCreateLineItemParams{
			{
				Amount:    stripe.Int64(1000),
				Reference: stripe.String("L1"),
			},
		},
		CustomerDetails: &stripe.TaxCalculationCreateCustomerDetailsParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("354 Oyster Point Blvd"),
				City:       stripe.String("South San Francisco"),
				State:      stripe.String("CA"),
				PostalCode: stripe.String("94080"),
				Country:    stripe.String("US"),
			},
			AddressSource: stripe.String(stripe.TaxCalculationCustomerDetailsAddressSourceShipping),
		},
	}
	result, err := sc.V1TaxCalculations.Create(context.TODO(), params)
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

func TestTaxCodesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TaxCodes.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTaxCodesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxCodeRetrieveParams{}
	result, err := sc.V1TaxCodes.Retrieve(
		context.TODO(), "txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxFormsPdfGet(t *testing.T) {
	params := &stripe.TaxFormPDFParams{}
	result, err := tax_form.PDF("form_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxFormsPdfGetService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxFormPDFParams{}
	result, err := sc.TaxForms.PDF("form_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxFormsPdfGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxFormPDFParams{}
	result, err := sc.V1TaxForms.PDF(context.TODO(), "form_xxxxxxxxxxxxx", params)
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

func TestTaxIdsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDDeleteParams{}
	result, err := sc.V1TaxIDs.Delete(context.TODO(), "taxid_123", params)
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

func TestTaxIdsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDListParams{}
	result := sc.V1TaxIDs.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTaxIdsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDRetrieveParams{}
	result, err := sc.V1TaxIDs.Retrieve(context.TODO(), "taxid_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsPost(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:  stripe.String(stripe.TaxIDTypeEUVAT),
		Value: stripe.String("123"),
	}
	result, err := taxid.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TaxIDParams{
		Type:  stripe.String(stripe.TaxIDTypeEUVAT),
		Value: stripe.String("123"),
	}
	result, err := sc.TaxIDs.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIdsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxIDCreateParams{
		Type:  stripe.String(stripe.TaxIDTypeEUVAT),
		Value: stripe.String("123"),
	}
	result, err := sc.V1TaxIDs.Create(context.TODO(), params)
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

func TestTaxRatesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxRateListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TaxRates.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTaxRatesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxRateRetrieveParams{}
	result, err := sc.V1TaxRates.Retrieve(
		context.TODO(), "txr_xxxxxxxxxxxxx", params)
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

func TestTaxRatesPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxRateCreateParams{
		DisplayName:  stripe.String("VAT"),
		Description:  stripe.String("VAT Germany"),
		Jurisdiction: stripe.String("DE"),
		Percentage:   stripe.Float64(16),
		Inclusive:    stripe.Bool(false),
	}
	result, err := sc.V1TaxRates.Create(context.TODO(), params)
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

func TestTaxRatesPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxRateUpdateParams{Active: stripe.Bool(false)}
	result, err := sc.V1TaxRates.Update(
		context.TODO(), "txr_xxxxxxxxxxxxx", params)
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

func TestTaxSettingsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxSettingsRetrieveParams{}
	result, err := sc.V1TaxSettings.Retrieve(context.TODO(), params)
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

func TestTaxSettingsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxSettingsUpdateParams{
		Defaults: &stripe.TaxSettingsUpdateDefaultsParams{
			TaxCode: stripe.String("txcd_10000000"),
		},
	}
	result, err := sc.V1TaxSettings.Update(context.TODO(), params)
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

func TestTaxTransactionsCreateFromCalculationPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TaxTransactionCreateFromCalculationParams{
		Calculation: stripe.String("xxx"),
		Reference:   stripe.String("yyy"),
	}
	result, err := sc.V1TaxTransactions.CreateFromCalculation(
		context.TODO(), params)
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

func TestTerminalConfigurationsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationDeleteParams{}
	result, err := sc.V1TerminalConfigurations.Delete(
		context.TODO(), "uc_123", params)
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

func TestTerminalConfigurationsDelete2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationDeleteParams{}
	result, err := sc.V1TerminalConfigurations.Delete(
		context.TODO(), "tmc_xxxxxxxxxxxxx", params)
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

func TestTerminalConfigurationsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationListParams{}
	result := sc.V1TerminalConfigurations.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTerminalConfigurationsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationRetrieveParams{}
	result, err := sc.V1TerminalConfigurations.Retrieve(
		context.TODO(), "uc_123", params)
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

func TestTerminalConfigurationsGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TerminalConfigurations.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTerminalConfigurationsGet4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationRetrieveParams{}
	result, err := sc.V1TerminalConfigurations.Retrieve(
		context.TODO(), "tmc_xxxxxxxxxxxxx", params)
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

func TestTerminalConfigurationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationCreateParams{}
	result, err := sc.V1TerminalConfigurations.Create(context.TODO(), params)
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

func TestTerminalConfigurationsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationUpdateParams{
		Tipping: &stripe.TerminalConfigurationUpdateTippingParams{
			USD: &stripe.TerminalConfigurationUpdateTippingUSDParams{
				FixedAmounts: []*int64{stripe.Int64(10)},
			},
		},
	}
	result, err := sc.V1TerminalConfigurations.Update(
		context.TODO(), "uc_123", params)
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

func TestTerminalConfigurationsPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationCreateParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationCreateBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := sc.V1TerminalConfigurations.Create(context.TODO(), params)
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

func TestTerminalConfigurationsPost4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConfigurationUpdateParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationUpdateBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := sc.V1TerminalConfigurations.Update(
		context.TODO(), "tmc_xxxxxxxxxxxxx", params)
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

func TestTerminalConnectionTokensPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalConnectionTokenCreateParams{}
	result, err := sc.V1TerminalConnectionTokens.Create(context.TODO(), params)
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

func TestTerminalLocationsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalLocationDeleteParams{}
	result, err := sc.V1TerminalLocations.Delete(
		context.TODO(), "tml_xxxxxxxxxxxxx", params)
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

func TestTerminalLocationsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalLocationListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TerminalLocations.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTerminalLocationsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalLocationRetrieveParams{}
	result, err := sc.V1TerminalLocations.Retrieve(
		context.TODO(), "tml_xxxxxxxxxxxxx", params)
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

func TestTerminalLocationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalLocationCreateParams{
		DisplayName: stripe.String("My First Store"),
		Address: &stripe.AddressParams{
			Line1:      stripe.String("1234 Main Street"),
			City:       stripe.String("San Francisco"),
			PostalCode: stripe.String("94111"),
			State:      stripe.String("CA"),
			Country:    stripe.String("US"),
		},
	}
	result, err := sc.V1TerminalLocations.Create(context.TODO(), params)
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

func TestTerminalLocationsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalLocationUpdateParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, err := sc.V1TerminalLocations.Update(
		context.TODO(), "tml_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersCancelActionPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := sc.V1TerminalReaders.CancelAction(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderDeleteParams{}
	result, err := sc.V1TerminalReaders.Delete(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TerminalReaders.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTerminalReadersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderRetrieveParams{}
	result, err := sc.V1TerminalReaders.Retrieve(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderCreateParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, err := sc.V1TerminalReaders.Create(context.TODO(), params)
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

func TestTerminalReadersPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderUpdateParams{
		Label: stripe.String("Blue Rabbit"),
	}
	result, err := sc.V1TerminalReaders.Update(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersProcessPaymentIntentPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String("pi_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TerminalReaders.ProcessPaymentIntent(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessSetupIntentPost(t *testing.T) {
	params := &stripe.TerminalReaderProcessSetupIntentParams{
		SetupIntent:    stripe.String("seti_xxxxxxxxxxxxx"),
		AllowRedisplay: stripe.String("always"),
	}
	result, err := terminal_reader.ProcessSetupIntent("tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersProcessSetupIntentPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TerminalReaderProcessSetupIntentParams{
		SetupIntent:    stripe.String("seti_xxxxxxxxxxxxx"),
		AllowRedisplay: stripe.String("always"),
	}
	result, err := sc.V1TerminalReaders.ProcessSetupIntent(
		context.TODO(), "tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersCustomersFundCashBalancePost(t *testing.T) {
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(stripe.CurrencyEUR),
	}
	result, err := testhelpers_customer.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersCustomersFundCashBalancePostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(stripe.CurrencyEUR),
	}
	result, err := sc.TestHelpersCustomers.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersCustomersFundCashBalancePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(stripe.CurrencyEUR),
	}
	result, err := sc.V1TestHelpersCustomers.FundCashBalance(
		context.TODO(), "cus_123", params)
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

func TestTestHelpersIssuingAuthorizationsCapturePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
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
	result, err := sc.V1TestHelpersIssuingAuthorizations.Capture(
		context.TODO(), "example_authorization", params)
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

func TestTestHelpersIssuingAuthorizationsExpirePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingAuthorizationExpireParams{}
	result, err := sc.V1TestHelpersIssuingAuthorizations.Expire(
		context.TODO(), "example_authorization", params)
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

func TestTestHelpersIssuingAuthorizationsIncrementPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingAuthorizationIncrementParams{
		IncrementAmount:      stripe.Int64(50),
		IsAmountControllable: stripe.Bool(true),
	}
	result, err := sc.V1TestHelpersIssuingAuthorizations.Increment(
		context.TODO(), "example_authorization", params)
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
		AuthorizationMethod:  stripe.String(stripe.IssuingAuthorizationAuthorizationMethodChip),
		Card:                 stripe.String("foo"),
		Currency:             stripe.String(stripe.CurrencyUSD),
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
		AuthorizationMethod:  stripe.String(stripe.IssuingAuthorizationAuthorizationMethodChip),
		Card:                 stripe.String("foo"),
		Currency:             stripe.String(stripe.CurrencyUSD),
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

func TestTestHelpersIssuingAuthorizationsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingAuthorizationCreateParams{
		Amount: stripe.Int64(100),
		AmountDetails: &stripe.TestHelpersIssuingAuthorizationCreateAmountDetailsParams{
			ATMFee:         stripe.Int64(10),
			CashbackAmount: stripe.Int64(5),
		},
		AuthorizationMethod:  stripe.String(stripe.IssuingAuthorizationAuthorizationMethodChip),
		Card:                 stripe.String("foo"),
		Currency:             stripe.String(stripe.CurrencyUSD),
		IsAmountControllable: stripe.Bool(true),
		MerchantData: &stripe.TestHelpersIssuingAuthorizationCreateMerchantDataParams{
			Category:   stripe.String("ac_refrigeration_repair"),
			City:       stripe.String("foo"),
			Country:    stripe.String("bar"),
			Name:       stripe.String("foo"),
			NetworkID:  stripe.String("bar"),
			PostalCode: stripe.String("foo"),
			State:      stripe.String("bar"),
			TerminalID: stripe.String("foo"),
		},
		NetworkData: &stripe.TestHelpersIssuingAuthorizationCreateNetworkDataParams{
			AcquiringInstitutionID: stripe.String("foo"),
		},
		VerificationData: &stripe.TestHelpersIssuingAuthorizationCreateVerificationDataParams{
			AddressLine1Check:      stripe.String("mismatch"),
			AddressPostalCodeCheck: stripe.String("match"),
			CVCCheck:               stripe.String("match"),
			ExpiryCheck:            stripe.String("mismatch"),
		},
		Wallet: stripe.String("apple_pay"),
	}
	result, err := sc.V1TestHelpersIssuingAuthorizations.Create(
		context.TODO(), params)
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

func TestTestHelpersIssuingAuthorizationsReversePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingAuthorizationReverseParams{
		ReverseAmount: stripe.Int64(20),
	}
	result, err := sc.V1TestHelpersIssuingAuthorizations.Reverse(
		context.TODO(), "example_authorization", params)
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

func TestTestHelpersIssuingCardsShippingDeliverPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, err := sc.V1TestHelpersIssuingCards.DeliverCard(
		context.TODO(), "card_123", params)
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

func TestTestHelpersIssuingCardsShippingFailPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, err := sc.V1TestHelpersIssuingCards.FailCard(
		context.TODO(), "card_123", params)
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

func TestTestHelpersIssuingCardsShippingReturnPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, err := sc.V1TestHelpersIssuingCards.ReturnCard(
		context.TODO(), "card_123", params)
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

func TestTestHelpersIssuingCardsShippingShipPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, err := sc.V1TestHelpersIssuingCards.ShipCard(
		context.TODO(), "card_123", params)
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

func TestTestHelpersIssuingPersonalizationDesignsActivatePostClient(
	t *testing.T,
) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingPersonalizationDesignActivateParams{}
	result, err := sc.V1TestHelpersIssuingPersonalizationDesigns.Activate(
		context.TODO(), "pd_xyz", params)
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

func TestTestHelpersIssuingPersonalizationDesignsDeactivatePostClient(
	t *testing.T,
) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams{}
	result, err := sc.V1TestHelpersIssuingPersonalizationDesigns.Deactivate(
		context.TODO(), "pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsRejectPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingPersonalizationDesignRejectParams{
		RejectionReasons: &stripe.TestHelpersIssuingPersonalizationDesignRejectRejectionReasonsParams{
			CardLogo: []*string{
				stripe.String(stripe.IssuingPersonalizationDesignRejectionReasonsCardLogoGeographicLocation),
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
				stripe.String(stripe.IssuingPersonalizationDesignRejectionReasonsCardLogoGeographicLocation),
			},
		},
	}
	result, err := sc.TestHelpersIssuingPersonalizationDesigns.Reject(
		"pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingPersonalizationDesignsRejectPostClient(
	t *testing.T,
) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingPersonalizationDesignRejectParams{
		RejectionReasons: &stripe.TestHelpersIssuingPersonalizationDesignRejectRejectionReasonsParams{
			CardLogo: []*string{
				stripe.String(stripe.IssuingPersonalizationDesignRejectionReasonsCardLogoGeographicLocation),
			},
		},
	}
	result, err := sc.V1TestHelpersIssuingPersonalizationDesigns.Reject(
		context.TODO(), "pd_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateForceCapturePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionCreateForceCaptureParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(stripe.CurrencyUSD),
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
		Currency: stripe.String(stripe.CurrencyUSD),
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

func TestTestHelpersIssuingTransactionsCreateForceCapturePostClient(
	t *testing.T,
) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingTransactionCreateForceCaptureParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(stripe.CurrencyUSD),
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
	result, err := sc.V1TestHelpersIssuingTransactions.CreateForceCapture(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingTransactionsCreateUnlinkedRefundPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(stripe.CurrencyUSD),
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
		Currency: stripe.String(stripe.CurrencyUSD),
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

func TestTestHelpersIssuingTransactionsCreateUnlinkedRefundPostClient(
	t *testing.T,
) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams{
		Amount:   stripe.Int64(100),
		Card:     stripe.String("foo"),
		Currency: stripe.String(stripe.CurrencyUSD),
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
	result, err := sc.V1TestHelpersIssuingTransactions.CreateUnlinkedRefund(
		context.TODO(), params)
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

func TestTestHelpersIssuingTransactionsRefundPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersIssuingTransactionRefundParams{
		RefundAmount: stripe.Int64(50),
	}
	result, err := sc.V1TestHelpersIssuingTransactions.Refund(
		context.TODO(), "example_transaction", params)
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

func TestTestHelpersRefundsExpirePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersRefundExpireParams{}
	result, err := sc.V1TestHelpersRefunds.Expire(
		context.TODO(), "re_123", params)
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

func TestTestHelpersTestClocksAdvancePostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, err := sc.V1TestHelpersTestClocks.Advance(
		context.TODO(), "clock_xyz", params)
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
	result, err := sc.TestHelpersTestClocks.Advance("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksAdvancePost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1675552261),
	}
	result, err := sc.V1TestHelpersTestClocks.Advance(
		context.TODO(), "clock_xxxxxxxxxxxxx", params)
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

func TestTestHelpersTestClocksDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockDeleteParams{}
	result, err := sc.V1TestHelpersTestClocks.Delete(
		context.TODO(), "clock_xyz", params)
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

func TestTestHelpersTestClocksDelete2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockDeleteParams{}
	result, err := sc.V1TestHelpersTestClocks.Delete(
		context.TODO(), "clock_xxxxxxxxxxxxx", params)
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

func TestTestHelpersTestClocksGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockListParams{}
	result := sc.V1TestHelpersTestClocks.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTestHelpersTestClocksGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockRetrieveParams{}
	result, err := sc.V1TestHelpersTestClocks.Retrieve(
		context.TODO(), "clock_xyz", params)
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

func TestTestHelpersTestClocksGet3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TestHelpersTestClocks.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTestHelpersTestClocksGet4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockRetrieveParams{}
	result, err := sc.V1TestHelpersTestClocks.Retrieve(
		context.TODO(), "clock_xxxxxxxxxxxxx", params)
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

func TestTestHelpersTestClocksPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockCreateParams{
		FrozenTime: stripe.Int64(123),
		Name:       stripe.String("cogsworth"),
	}
	result, err := sc.V1TestHelpersTestClocks.Create(context.TODO(), params)
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

func TestTestHelpersTestClocksPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTestClockCreateParams{
		FrozenTime: stripe.Int64(1577836800),
	}
	result, err := sc.V1TestHelpersTestClocks.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersFailPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed),
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
			Code: stripe.String(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed),
		},
	}
	result, err := sc.TestHelpersTreasuryInboundTransfers.Fail("ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersFailPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed),
		},
	}
	result, err := sc.V1TestHelpersTreasuryInboundTransfers.Fail(
		context.TODO(), "ibt_123", params)
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

func TestTestHelpersTreasuryInboundTransfersReturnPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, err := sc.V1TestHelpersTreasuryInboundTransfers.ReturnInboundTransfer(
		context.TODO(), "ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersSucceedPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := testhelpers_treasury_inboundtransfer.Succeed("ibt_123", params)
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

func TestTestHelpersTreasuryInboundTransfersSucceedPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := sc.V1TestHelpersTreasuryInboundTransfers.Succeed(
		context.TODO(), "ibt_123", params)
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
	result, err := sc.TestHelpersTreasuryOutboundTransfers.Fail("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersFailPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, err := sc.V1TestHelpersTreasuryOutboundTransfers.Fail(
		context.TODO(), "obt_123", params)
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
	result, err := sc.TestHelpersTreasuryOutboundTransfers.Post("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersPostPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, err := sc.V1TestHelpersTreasuryOutboundTransfers.Post(
		context.TODO(), "obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersReturnPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams{
		ReturnedDetails: &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams{
			Code: stripe.String(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed),
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
			Code: stripe.String(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed),
		},
	}
	result, err := sc.TestHelpersTreasuryOutboundTransfers.ReturnOutboundTransfer(
		"obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersReturnPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams{
		ReturnedDetails: &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams{
			Code: stripe.String(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed),
		},
	}
	result, err := sc.V1TestHelpersTreasuryOutboundTransfers.ReturnOutboundTransfer(
		context.TODO(), "obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedCreditsPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedCreditParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(stripe.TreasuryReceivedCreditNetworkACH),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(stripe.CurrencyUSD),
	}
	result, err := testhelpers_treasury_receivedcredit.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedCreditsPostService(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	params := &stripe.TestHelpersTreasuryReceivedCreditParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(stripe.TreasuryReceivedCreditNetworkACH),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.TestHelpersTreasuryReceivedCredits.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedCreditsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryReceivedCreditCreateParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(stripe.TreasuryReceivedCreditNetworkACH),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.V1TestHelpersTreasuryReceivedCredits.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedDebitsPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedDebitParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String("ach"),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(stripe.CurrencyUSD),
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
		Currency:         stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.TestHelpersTreasuryReceivedDebits.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryReceivedDebitsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TestHelpersTreasuryReceivedDebitCreateParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String("ach"),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(stripe.CurrencyUSD),
	}
	result, err := sc.V1TestHelpersTreasuryReceivedDebits.Create(
		context.TODO(), params)
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

func TestTokensGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenRetrieveParams{}
	result, err := sc.V1Tokens.Retrieve(context.TODO(), "tok_xxxx", params)
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

func TestTokensPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		Card: &stripe.CardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.String("5"),
			ExpYear:  stripe.String("2023"),
			CVC:      stripe.String("314"),
		},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokensPost2(t *testing.T) {
	params := &stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			Country:           stripe.String("US"),
			Currency:          stripe.String(stripe.CurrencyUSD),
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
			Currency:          stripe.String(stripe.CurrencyUSD),
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

func TestTokensPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		BankAccount: &stripe.BankAccountParams{
			Country:           stripe.String("US"),
			Currency:          stripe.String(stripe.CurrencyUSD),
			AccountHolderName: stripe.String("Jenny Rosen"),
			AccountHolderType: stripe.String("individual"),
			RoutingNumber:     stripe.String("110000000"),
			AccountNumber:     stripe.String("000123456789"),
		},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
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

func TestTokensPost3Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		PII: &stripe.TokenCreatePIIParams{IDNumber: stripe.String("000000000")},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
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

func TestTokensPost4Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		Account: &stripe.TokenCreateAccountParams{
			Individual: &stripe.PersonParams{
				FirstName: stripe.String("Jane"),
				LastName:  stripe.String("Doe"),
			},
			TOSShownAndAccepted: stripe.Bool(true),
		},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
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

func TestTokensPost5Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		Person: &stripe.PersonParams{
			FirstName:    stripe.String("Jane"),
			LastName:     stripe.String("Doe"),
			Relationship: &stripe.PersonRelationshipParams{Owner: stripe.Bool(true)},
		},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
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

func TestTokensPost6Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TokenCreateParams{
		CVCUpdate: &stripe.TokenCreateCVCUpdateParams{CVC: stripe.String("123")},
	}
	result, err := sc.V1Tokens.Create(context.TODO(), params)
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

func TestTopupsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TopupCancelParams{}
	result, err := sc.V1Topups.Cancel(context.TODO(), "tu_xxxxxxxxxxxxx", params)
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

func TestTopupsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TopupListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Topups.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTopupsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TopupRetrieveParams{}
	result, err := sc.V1Topups.Retrieve(
		context.TODO(), "tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPost(t *testing.T) {
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(stripe.CurrencyUSD),
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
		Currency:            stripe.String(stripe.CurrencyUSD),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, err := sc.Topups.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TopupCreateParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(stripe.CurrencyUSD),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, err := sc.V1Topups.Create(context.TODO(), params)
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

func TestTopupsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TopupUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Topups.Update(context.TODO(), "tu_xxxxxxxxxxxxx", params)
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

func TestTransfersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1Transfers.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTransfersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferRetrieveParams{}
	result, err := sc.V1Transfers.Retrieve(
		context.TODO(), "tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPost(t *testing.T) {
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(stripe.CurrencyUSD),
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
		Currency:      stripe.String(stripe.CurrencyUSD),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, err := sc.Transfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransfersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferCreateParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(stripe.CurrencyUSD),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, err := sc.V1Transfers.Create(context.TODO(), params)
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

func TestTransfersPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1Transfers.Update(
		context.TODO(), "tr_xxxxxxxxxxxxx", params)
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

func TestTransfersReversalsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferReversalListParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TransferReversals.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTransfersReversalsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferReversalRetrieveParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TransferReversals.Retrieve(
		context.TODO(), "trr_xxxxxxxxxxxxx", params)
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

func TestTransfersReversalsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferReversalCreateParams{
		Amount: stripe.Int64(100),
		ID:     stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TransferReversals.Create(context.TODO(), params)
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

func TestTransfersReversalsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TransferReversalUpdateParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1TransferReversals.Update(
		context.TODO(), "trr_xxxxxxxxxxxxx", params)
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

func TestTreasuryCreditReversalsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryCreditReversals.List(context.TODO(), params)
	assert.NotNil(t, result)
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
	result, err := sc.TreasuryCreditReversals.Get("credrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryCreditReversalRetrieveParams{}
	result, err := sc.V1TreasuryCreditReversals.Retrieve(
		context.TODO(), "credrev_xxxxxxxxxxxxx", params)
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

func TestTreasuryCreditReversalsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryCreditReversalCreateParams{
		ReceivedCredit: stripe.String("rc_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TreasuryCreditReversals.Create(context.TODO(), params)
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

func TestTreasuryDebitReversalsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryDebitReversals.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryDebitReversalsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryDebitReversalRetrieveParams{}
	result, err := sc.V1TreasuryDebitReversals.Retrieve(
		context.TODO(), "debrev_xxxxxxxxxxxxx", params)
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

func TestTreasuryDebitReversalsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryDebitReversalCreateParams{
		ReceivedDebit: stripe.String("rd_xxxxxxxxxxxxx"),
	}
	result, err := sc.V1TreasuryDebitReversals.Create(context.TODO(), params)
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

func TestTreasuryFinancialAccountsFeaturesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, err := sc.V1TreasuryFinancialAccounts.RetrieveFeatures(
		context.TODO(), "fa_xxxxxxxxxxxxx", params)
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

func TestTreasuryFinancialAccountsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryFinancialAccountListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryFinancialAccounts.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryFinancialAccountsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryFinancialAccountRetrieveParams{}
	result, err := sc.V1TreasuryFinancialAccounts.Retrieve(
		context.TODO(), "fa_xxxxxxxxxxxxx", params)
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

func TestTreasuryFinancialAccountsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryFinancialAccountCreateParams{
		SupportedCurrencies: []*string{stripe.String("usd")},
		Features:            &stripe.TreasuryFinancialAccountCreateFeaturesParams{},
	}
	result, err := sc.V1TreasuryFinancialAccounts.Create(context.TODO(), params)
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
	result, err := sc.TreasuryFinancialAccounts.Update("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryFinancialAccountUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := sc.V1TreasuryFinancialAccounts.Update(
		context.TODO(), "fa_xxxxxxxxxxxxx", params)
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
	result, err := sc.TreasuryInboundTransfers.Cancel("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, err := sc.V1TreasuryInboundTransfers.Cancel(
		context.TODO(), "ibt_xxxxxxxxxxxxx", params)
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

func TestTreasuryInboundTransfersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryInboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryInboundTransfers.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryInboundTransfersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryInboundTransferRetrieveParams{}
	result, err := sc.V1TreasuryInboundTransfers.Retrieve(
		context.TODO(), "ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersPost(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{
		FinancialAccount:    stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:              stripe.Int64(10000),
		Currency:            stripe.String(stripe.CurrencyUSD),
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
		Currency:            stripe.String(stripe.CurrencyUSD),
		OriginPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:         stripe.String("InboundTransfer from my bank account"),
	}
	result, err := sc.TreasuryInboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryInboundTransferCreateParams{
		FinancialAccount:    stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:              stripe.Int64(10000),
		Currency:            stripe.String(stripe.CurrencyUSD),
		OriginPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:         stripe.String("InboundTransfer from my bank account"),
	}
	result, err := sc.V1TreasuryInboundTransfers.Create(context.TODO(), params)
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
	result, err := sc.TreasuryOutboundPayments.Cancel("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, err := sc.V1TreasuryOutboundPayments.Cancel(
		context.TODO(), "bot_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundPaymentsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundPaymentListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryOutboundPayments.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryOutboundPaymentsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundPaymentRetrieveParams{}
	result, err := sc.V1TreasuryOutboundPayments.Retrieve(
		context.TODO(), "bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsPost(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(10000),
		Currency:                 stripe.String(stripe.CurrencyUSD),
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
		Currency:                 stripe.String(stripe.CurrencyUSD),
		Customer:                 stripe.String("cus_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, err := sc.TreasuryOutboundPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundPaymentCreateParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(10000),
		Currency:                 stripe.String(stripe.CurrencyUSD),
		Customer:                 stripe.String("cus_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, err := sc.V1TreasuryOutboundPayments.Create(context.TODO(), params)
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

func TestTreasuryOutboundTransfersCancelPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, err := sc.V1TreasuryOutboundTransfers.Cancel(
		context.TODO(), "obt_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundTransfersGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryOutboundTransfers.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryOutboundTransfersGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundTransferRetrieveParams{}
	result, err := sc.V1TreasuryOutboundTransfers.Retrieve(
		context.TODO(), "obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersPost(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(500),
		Currency:                 stripe.String(stripe.CurrencyUSD),
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
		Currency:                 stripe.String(stripe.CurrencyUSD),
		Description:              stripe.String("OutboundTransfer to my external bank account"),
	}
	result, err := sc.TreasuryOutboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransfersPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryOutboundTransferCreateParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(500),
		Currency:                 stripe.String(stripe.CurrencyUSD),
		Description:              stripe.String("OutboundTransfer to my external bank account"),
	}
	result, err := sc.V1TreasuryOutboundTransfers.Create(context.TODO(), params)
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

func TestTreasuryReceivedCreditsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryReceivedCredits.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryReceivedCreditsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryReceivedCreditRetrieveParams{}
	result, err := sc.V1TreasuryReceivedCredits.Retrieve(
		context.TODO(), "rc_xxxxxxxxxxxxx", params)
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

func TestTreasuryReceivedDebitsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryReceivedDebitListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryReceivedDebits.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryReceivedDebitsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryReceivedDebitRetrieveParams{}
	result, err := sc.V1TreasuryReceivedDebits.Retrieve(
		context.TODO(), "rd_xxxxxxxxxxxxx", params)
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

func TestTreasuryTransactionEntriesGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryTransactionEntryListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryTransactionEntries.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryTransactionEntriesGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryTransactionEntryRetrieveParams{}
	result, err := sc.V1TreasuryTransactionEntries.Retrieve(
		context.TODO(), "trxne_xxxxxxxxxxxxx", params)
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

func TestTreasuryTransactionsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryTransactionListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := sc.V1TreasuryTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestTreasuryTransactionsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.TreasuryTransactionRetrieveParams{}
	result, err := sc.V1TreasuryTransactions.Retrieve(
		context.TODO(), "trxn_xxxxxxxxxxxxx", params)
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

func TestWebhookEndpointsDeleteClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.WebhookEndpointDeleteParams{}
	result, err := sc.V1WebhookEndpoints.Delete(
		context.TODO(), "we_xxxxxxxxxxxxx", params)
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

func TestWebhookEndpointsGetClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.WebhookEndpointListParams{}
	params.Limit = stripe.Int64(3)
	result := sc.V1WebhookEndpoints.List(context.TODO(), params)
	assert.NotNil(t, result)
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

func TestWebhookEndpointsGet2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.WebhookEndpointRetrieveParams{}
	result, err := sc.V1WebhookEndpoints.Retrieve(
		context.TODO(), "we_xxxxxxxxxxxxx", params)
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

func TestWebhookEndpointsPostClient(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.WebhookEndpointCreateParams{
		URL: stripe.String("https://example.com/my/webhook/endpoint"),
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
	}
	result, err := sc.V1WebhookEndpoints.Create(context.TODO(), params)
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

func TestWebhookEndpointsPost2Client(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.WebhookEndpointUpdateParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, err := sc.V1WebhookEndpoints.Update(
		context.TODO(), "we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCoreEventsGetService(t *testing.T) {
	params := &stripe.V2CoreEventParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events/ll_123", params, "{\"context\":\"context\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"reason\":{\"type\":\"request\",\"request\":{\"id\":\"obj_123\",\"idempotency_key\":\"idempotency_key\"}},\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEvents.Get("ll_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCoreEventsGetClient(t *testing.T) {
	params := &stripe.V2CoreEventRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events/ll_123", params, "{\"context\":\"context\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"reason\":{\"type\":\"request\",\"request\":{\"id\":\"obj_123\",\"idempotency_key\":\"idempotency_key\"}},\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEvents.Retrieve(context.TODO(), "ll_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventPostService(t *testing.T) {
	params := &stripe.V2BillingMeterEventParams{
		EventName: stripe.String("event_name"),
		Payload:   map[string]string{"key": "payload"},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_events", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"event_name\":\"event_name\",\"identifier\":\"identifier\",\"object\":\"v2.billing.meter_event\",\"payload\":{\"key\":\"payload\"},\"timestamp\":\"1970-01-01T15:18:46.294Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2BillingMeterEvents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventPostClient(t *testing.T) {
	params := &stripe.V2BillingMeterEventCreateParams{
		EventName: stripe.String("event_name"),
		Payload:   map[string]string{"key": "payload"},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_events", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"event_name\":\"event_name\",\"identifier\":\"identifier\",\"object\":\"v2.billing.meter_event\",\"payload\":{\"key\":\"payload\"},\"timestamp\":\"1970-01-01T15:18:46.294Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2BillingMeterEvents.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventAdjustmentPostService(t *testing.T) {
	params := &stripe.V2BillingMeterEventAdjustmentParams{
		Cancel: &stripe.V2BillingMeterEventAdjustmentCancelParams{
			Identifier: stripe.String("identifier"),
		},
		EventName: stripe.String("event_name"),
		Type:      stripe.String("cancel"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_adjustments", params, "{\"cancel\":{\"identifier\":\"identifier\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"event_name\":\"event_name\",\"id\":\"obj_123\",\"object\":\"v2.billing.meter_event_adjustment\",\"status\":\"complete\",\"type\":\"cancel\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2BillingMeterEventAdjustments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventAdjustmentPostClient(t *testing.T) {
	params := &stripe.V2BillingMeterEventAdjustmentCreateParams{
		Cancel: &stripe.V2BillingMeterEventAdjustmentCreateCancelParams{
			Identifier: stripe.String("identifier"),
		},
		EventName: stripe.String("event_name"),
		Type:      stripe.String("cancel"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_adjustments", params, "{\"cancel\":{\"identifier\":\"identifier\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"event_name\":\"event_name\",\"id\":\"obj_123\",\"object\":\"v2.billing.meter_event_adjustment\",\"status\":\"complete\",\"type\":\"cancel\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2BillingMeterEventAdjustments.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventSessionPostService(t *testing.T) {
	params := &stripe.V2BillingMeterEventSessionParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_session", params, "{\"authentication_token\":\"authentication_token\",\"created\":\"1970-01-12T21:42:34.472Z\",\"expires_at\":\"1970-01-10T15:36:51.170Z\",\"id\":\"obj_123\",\"object\":\"v2.billing.meter_event_session\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2BillingMeterEventSessions.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventSessionPostClient(t *testing.T) {
	params := &stripe.V2BillingMeterEventSessionCreateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_session", params, "{\"authentication_token\":\"authentication_token\",\"created\":\"1970-01-12T21:42:34.472Z\",\"expires_at\":\"1970-01-10T15:36:51.170Z\",\"id\":\"obj_123\",\"object\":\"v2.billing.meter_event_session\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2BillingMeterEventSessions.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2BillingMeterEventStreamPostService(t *testing.T) {
	params := &stripe.V2BillingMeterEventStreamParams{
		Events: []*stripe.V2BillingMeterEventStreamEventParams{
			{
				EventName:  stripe.String("event_name"),
				Identifier: stripe.String("identifier"),
				Payload:    map[string]string{"key": "payload"},
				Timestamp:  stripe.Time(time.Now()),
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_stream", params, "{}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	sc.V2BillingMeterEventStreams.New(params)
}

func TestV2BillingMeterEventStreamPostClient(t *testing.T) {
	params := &stripe.V2BillingMeterEventStreamCreateParams{
		Events: []*stripe.V2BillingMeterEventStreamCreateEventParams{
			{
				EventName:  stripe.String("event_name"),
				Identifier: stripe.String("identifier"),
				Payload:    map[string]string{"key": "payload"},
				Timestamp:  stripe.Time(time.Now()),
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_stream", params, "{}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	sc.V2BillingMeterEventStreams.Create(context.TODO(), params)
}

func TestV2CoreAccountGetService(t *testing.T) {
	params := &stripe.V2CoreAccountListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts", params, "{\"data\":[{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2CoreAccounts.All(params)
	assert.NotNil(t, result)
}

func TestV2CoreAccountGetClient(t *testing.T) {
	params := &stripe.V2CoreAccountListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts", params, "{\"data\":[{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2CoreAccounts.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2CoreAccountPostService(t *testing.T) {
	params := &stripe.V2CoreAccountParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountPostClient(t *testing.T) {
	params := &stripe.V2CoreAccountCreateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountGet2Service(t *testing.T) {
	params := &stripe.V2CoreAccountParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/id_123", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccounts.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountGet2Client(t *testing.T) {
	params := &stripe.V2CoreAccountRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/id_123", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccounts.Retrieve(context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountPost2Service(t *testing.T) {
	params := &stripe.V2CoreAccountParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/id_123", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccounts.Update("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountPost2Client(t *testing.T) {
	params := &stripe.V2CoreAccountUpdateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/id_123", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccounts.Update(context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountPost3Service(t *testing.T) {
	params := &stripe.V2CoreAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/id_123/close", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccounts.Close("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountPost3Client(t *testing.T) {
	params := &stripe.V2CoreAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/id_123/close", params, "{\"applied_configurations\":[\"storer\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccounts.Close(context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonGetService(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonListParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/account_id_123/persons", params, "{\"data\":[{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2CoreAccountsPersons.All(params)
	assert.NotNil(t, result)
}

func TestV2CoreAccountsPersonGetClient(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonListParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/account_id_123/persons", params, "{\"data\":[{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2CoreAccountsPersons.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2CoreAccountsPersonPostService(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/account_id_123/persons", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccountsPersons.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonPostClient(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonCreateParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/account_id_123/persons", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccountsPersons.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonDeleteService(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodDelete, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"id\":\"abc_123\",\"object\":\"some.object.tag\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccountsPersons.Del("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonDeleteClient(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonDeleteParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodDelete, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"id\":\"abc_123\",\"object\":\"some.object.tag\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccountsPersons.Delete(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonGet2Service(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccountsPersons.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonGet2Client(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonRetrieveParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccountsPersons.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonPost2Service(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccountsPersons.Update("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountsPersonPost2Client(t *testing.T) {
	params := &stripe.V2CoreAccountsPersonUpdateParams{
		AccountID: stripe.String("account_id_123"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/accounts/account_id_123/persons/id_123", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.account_person\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccountsPersons.Update(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountLinkPostService(t *testing.T) {
	params := &stripe.V2CoreAccountLinkParams{
		Account: stripe.String("account"),
		UseCase: &stripe.V2CoreAccountLinkUseCaseParams{
			Type: stripe.String("account_onboarding"),
			AccountOnboarding: &stripe.V2CoreAccountLinkUseCaseAccountOnboardingParams{
				CollectionOptions: &stripe.V2CoreAccountLinkUseCaseAccountOnboardingCollectionOptionsParams{
					Fields:             stripe.String("eventually_due"),
					FutureRequirements: stripe.String("include"),
				},
				Configurations: []*string{stripe.String("storer")},
				RefreshURL:     stripe.String("refresh_url"),
				ReturnURL:      stripe.String("return_url"),
			},
			AccountUpdate: &stripe.V2CoreAccountLinkUseCaseAccountUpdateParams{
				CollectionOptions: &stripe.V2CoreAccountLinkUseCaseAccountUpdateCollectionOptionsParams{
					Fields:             stripe.String("eventually_due"),
					FutureRequirements: stripe.String("include"),
				},
				Configurations: []*string{stripe.String("storer")},
				RefreshURL:     stripe.String("refresh_url"),
				ReturnURL:      stripe.String("return_url"),
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/account_links", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"expires_at\":\"1970-01-10T15:36:51.170Z\",\"object\":\"v2.core.account_link\",\"url\":\"url\",\"use_case\":{\"type\":\"account_onboarding\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreAccountLinks.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreAccountLinkPostClient(t *testing.T) {
	params := &stripe.V2CoreAccountLinkCreateParams{
		Account: stripe.String("account"),
		UseCase: &stripe.V2CoreAccountLinkCreateUseCaseParams{
			Type: stripe.String("account_onboarding"),
			AccountOnboarding: &stripe.V2CoreAccountLinkCreateUseCaseAccountOnboardingParams{
				CollectionOptions: &stripe.V2CoreAccountLinkCreateUseCaseAccountOnboardingCollectionOptionsParams{
					Fields:             stripe.String("eventually_due"),
					FutureRequirements: stripe.String("include"),
				},
				Configurations: []*string{stripe.String("storer")},
				RefreshURL:     stripe.String("refresh_url"),
				ReturnURL:      stripe.String("return_url"),
			},
			AccountUpdate: &stripe.V2CoreAccountLinkCreateUseCaseAccountUpdateParams{
				CollectionOptions: &stripe.V2CoreAccountLinkCreateUseCaseAccountUpdateCollectionOptionsParams{
					Fields:             stripe.String("eventually_due"),
					FutureRequirements: stripe.String("include"),
				},
				Configurations: []*string{stripe.String("storer")},
				RefreshURL:     stripe.String("refresh_url"),
				ReturnURL:      stripe.String("return_url"),
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/account_links", params, "{\"account\":\"account\",\"created\":\"1970-01-12T21:42:34.472Z\",\"expires_at\":\"1970-01-10T15:36:51.170Z\",\"object\":\"v2.core.account_link\",\"url\":\"url\",\"use_case\":{\"type\":\"account_onboarding\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreAccountLinks.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventGetService(t *testing.T) {
	params := &stripe.V2CoreEventListParams{ObjectID: stripe.String("object_id")}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2CoreEvents.All(params)
	assert.NotNil(t, result)
}

func TestV2CoreEventGetClient(t *testing.T) {
	params := &stripe.V2CoreEventListParams{ObjectID: stripe.String("object_id")}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2CoreEvents.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2CoreEventGet2Service(t *testing.T) {
	params := &stripe.V2CoreEventParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEvents.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventGet2Client(t *testing.T) {
	params := &stripe.V2CoreEventRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/events/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEvents.Retrieve(context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationGetService(t *testing.T) {
	params := &stripe.V2CoreEventDestinationListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/event_destinations", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2CoreEventDestinations.All(params)
	assert.NotNil(t, result)
}

func TestV2CoreEventDestinationGetClient(t *testing.T) {
	params := &stripe.V2CoreEventDestinationListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/event_destinations", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2CoreEventDestinations.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2CoreEventDestinationPostService(t *testing.T) {
	params := &stripe.V2CoreEventDestinationParams{
		EnabledEvents: []*string{stripe.String("enabled_events")},
		EventPayload:  stripe.String("thin"),
		Name:          stripe.String("name"),
		Type:          stripe.String("amazon_eventbridge"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPostClient(t *testing.T) {
	params := &stripe.V2CoreEventDestinationCreateParams{
		EnabledEvents: []*string{stripe.String("enabled_events")},
		EventPayload:  stripe.String("thin"),
		Name:          stripe.String("name"),
		Type:          stripe.String("amazon_eventbridge"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationDeleteService(t *testing.T) {
	params := &stripe.V2CoreEventDestinationParams{}
	testServer := MockServer(
		t, http.MethodDelete, "/v2/core/event_destinations/id_123", params, "{\"id\":\"abc_123\",\"object\":\"some.object.tag\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Del("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationDeleteClient(t *testing.T) {
	params := &stripe.V2CoreEventDestinationDeleteParams{}
	testServer := MockServer(
		t, http.MethodDelete, "/v2/core/event_destinations/id_123", params, "{\"id\":\"abc_123\",\"object\":\"some.object.tag\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Delete(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationGet2Service(t *testing.T) {
	params := &stripe.V2CoreEventDestinationParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/event_destinations/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationGet2Client(t *testing.T) {
	params := &stripe.V2CoreEventDestinationRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/event_destinations/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost2Service(t *testing.T) {
	params := &stripe.V2CoreEventDestinationParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Update("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost2Client(t *testing.T) {
	params := &stripe.V2CoreEventDestinationUpdateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Update(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost3Service(t *testing.T) {
	params := &stripe.V2CoreEventDestinationDisableParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/disable", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Disable("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost3Client(t *testing.T) {
	params := &stripe.V2CoreEventDestinationDisableParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/disable", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Disable(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost4Service(t *testing.T) {
	params := &stripe.V2CoreEventDestinationEnableParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/enable", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Enable("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost4Client(t *testing.T) {
	params := &stripe.V2CoreEventDestinationEnableParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/enable", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"enabled_events\":[\"enabled_events\"],\"event_payload\":\"thin\",\"id\":\"obj_123\",\"name\":\"name\",\"object\":\"v2.core.event_destination\",\"status\":\"disabled\",\"type\":\"amazon_eventbridge\",\"updated\":\"1970-01-03T17:07:10.277Z\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Enable(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost5Service(t *testing.T) {
	params := &stripe.V2CoreEventDestinationPingParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/ping", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreEventDestinations.Ping("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreEventDestinationPost5Client(t *testing.T) {
	params := &stripe.V2CoreEventDestinationPingParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/event_destinations/id_123/ping", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.core.event\",\"type\":\"type\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreEventDestinations.Ping("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPostService(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountParams{
		AccountNumber: stripe.String("account_number"),
		SortCode:      stripe.String("sort_code"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultGBBankAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPostClient(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountCreateParams{
		AccountNumber: stripe.String("account_number"),
		SortCode:      stripe.String("sort_code"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultGBBankAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountGetService(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/vault/gb_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultGBBankAccounts.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountGetClient(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/vault/gb_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultGBBankAccounts.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost2Service(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountAcknowledgeConfirmationOfPayeeParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/acknowledge_confirmation_of_payee", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultGBBankAccounts.AcknowledgeConfirmationOfPayee(
		"id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost2Client(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountAcknowledgeConfirmationOfPayeeParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/acknowledge_confirmation_of_payee", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultGBBankAccounts.AcknowledgeConfirmationOfPayee(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost3Service(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/archive", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultGBBankAccounts.Archive("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost3Client(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/archive", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultGBBankAccounts.Archive(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost4Service(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountInitiateConfirmationOfPayeeParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/initiate_confirmation_of_payee", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultGBBankAccounts.InitiateConfirmationOfPayee(
		"id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultGbBankAccountPost4Client(t *testing.T) {
	params := &stripe.V2CoreVaultGBBankAccountInitiateConfirmationOfPayeeParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/gb_bank_accounts/id_123/initiate_confirmation_of_payee", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"confirmation_of_payee\":{\"result\":{\"created\":\"1970-01-12T21:42:34.472Z\",\"match_result\":\"unavailable\",\"matched\":{},\"message\":\"message\",\"provided\":{\"business_type\":\"personal\",\"name\":\"name\"}},\"status\":\"awaiting_acknowledgement\"},\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.gb_bank_account\",\"sort_code\":\"sort_code\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultGBBankAccounts.InitiateConfirmationOfPayee(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPostService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPostClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountCreateParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountGetService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/vault/us_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountGetClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/core/vault/us_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPost2Service(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.Update("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPost2Client(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountUpdateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Update(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPost3Service(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123/archive", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.Archive("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2CoreVaultUsBankAccountPost3Client(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123/archive", params, "{\"archived\":true,\"bank_account_type\":\"savings\",\"bank_name\":\"bank_name\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"last4\":\"last4\",\"object\":\"v2.core.vault.us_bank_account\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Archive(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementAdjustmentGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementAdjustmentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/adjustments", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.adjustment\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementAdjustments.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementAdjustmentGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementAdjustmentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/adjustments", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.adjustment\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementAdjustments.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementAdjustmentGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementAdjustmentParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/adjustments/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.adjustment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementAdjustments.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementAdjustmentGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementAdjustmentRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/adjustments/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.adjustment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementAdjustments.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_accounts", params, "{\"data\":[{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementFinancialAccounts.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementFinancialAccountGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_accounts", params, "{\"data\":[{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementFinancialAccounts.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementFinancialAccountPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCreateParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_accounts/id_123", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_accounts/id_123", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountPost2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts/id_123/close", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.Close("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAccountPost2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts/id_123/close", params, "{\"balance\":{\"available\":{\"key\":{\"currency\":\"USD\",\"value\":35}},\"inbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":11}},\"outbound_pending\":{\"key\":{\"currency\":\"USD\",\"value\":60}}},\"country\":\"country\",\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_account\",\"status\":\"closed\",\"type\":\"other\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Close(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAddressGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_addresses", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementFinancialAddresses.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementFinancialAddressGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_addresses", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementFinancialAddresses.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementFinancialAddressPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressParams{
		Currency:         stripe.String(stripe.CurrencyUSD),
		FinancialAccount: stripe.String("financial_account"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_addresses", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAddresses.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAddressPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressCreateParams{
		Currency:         stripe.String(stripe.CurrencyUSD),
		FinancialAccount: stripe.String("financial_account"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_addresses", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAddresses.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAddressGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_addresses/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAddresses.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementFinancialAddressGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/financial_addresses/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"currency\":\"usd\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.financial_address\",\"status\":\"failed\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAddresses.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementInboundTransferGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/inbound_transfers", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementInboundTransfers.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementInboundTransferGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/inbound_transfers", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementInboundTransfers.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementInboundTransferPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementInboundTransferFromParams{
			Currency:      stripe.String(stripe.CurrencyUSD),
			PaymentMethod: stripe.String("payment_method"),
		},
		To: &stripe.V2MoneyManagementInboundTransferToParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/inbound_transfers", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementInboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementInboundTransferPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementInboundTransferCreateFromParams{
			Currency:      stripe.String(stripe.CurrencyUSD),
			PaymentMethod: stripe.String("payment_method"),
		},
		To: &stripe.V2MoneyManagementInboundTransferCreateToParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/inbound_transfers", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementInboundTransfers.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementInboundTransferGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/inbound_transfers/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementInboundTransfers.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementInboundTransferGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementInboundTransferRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/inbound_transfers/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"description\":\"description\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"payment_method\":{\"type\":\"type\"}},\"id\":\"obj_123\",\"object\":\"v2.money_management.inbound_transfer\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"financial_account\":\"financial_account\"},\"transfer_history\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"level\":\"canonical\",\"type\":\"bank_debit_failed\"}],\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementInboundTransfers.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payments", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementOutboundPayments.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundPaymentGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payments", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementOutboundPayments.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundPaymentPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
			Recipient:    stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentCreateFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentCreateToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
			Recipient:    stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payments/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payments/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentPost2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentPost2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment\",\"recipient_notification\":{\"setting\":\"configured\"},\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentQuotePostService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentQuoteParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentQuoteFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentQuoteToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
			Recipient:    stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payment_quotes", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"estimated_fees\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"type\":\"cross_border_payout_fee\"}],\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"fx_quote\":{\"lock_duration\":\"five_minutes\",\"lock_expires_at\":\"1970-01-18T15:15:29.586Z\",\"lock_status\":\"active\",\"rates\":{\"key\":{\"exchange_rate\":\"exchange_rate\"}},\"to_currency\":\"usd\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment_quote\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPaymentQuotes.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentQuotePostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentQuoteCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentQuoteCreateFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentQuoteCreateToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
			Recipient:    stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payment_quotes", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"estimated_fees\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"type\":\"cross_border_payout_fee\"}],\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"fx_quote\":{\"lock_duration\":\"five_minutes\",\"lock_expires_at\":\"1970-01-18T15:15:29.586Z\",\"lock_status\":\"active\",\"rates\":{\"key\":{\"exchange_rate\":\"exchange_rate\"}},\"to_currency\":\"usd\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment_quote\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPaymentQuotes.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentQuoteGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentQuoteParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payment_quotes/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"estimated_fees\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"type\":\"cross_border_payout_fee\"}],\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"fx_quote\":{\"lock_duration\":\"five_minutes\",\"lock_expires_at\":\"1970-01-18T15:15:29.586Z\",\"lock_status\":\"active\",\"rates\":{\"key\":{\"exchange_rate\":\"exchange_rate\"}},\"to_currency\":\"usd\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment_quote\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPaymentQuotes.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundPaymentQuoteGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentQuoteRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_payment_quotes/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"estimated_fees\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"type\":\"cross_border_payout_fee\"}],\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"fx_quote\":{\"lock_duration\":\"five_minutes\",\"lock_expires_at\":\"1970-01-18T15:15:29.586Z\",\"lock_status\":\"active\",\"rates\":{\"key\":{\"exchange_rate\":\"exchange_rate\"}},\"to_currency\":\"usd\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_payment_quote\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\",\"recipient\":\"recipient\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPaymentQuotes.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_setup_intents", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementOutboundSetupIntents.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundSetupIntentGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_setup_intents", params, "{\"data\":[{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementOutboundSetupIntents.List(
		context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundSetupIntentPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundSetupIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentCreateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_setup_intents/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_setup_intents/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentPost2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Update(
		"id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentPost2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentUpdateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents/id_123", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Update(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentPost3Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents/id_123/cancel", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Cancel(
		"id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundSetupIntentPost3Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents/id_123/cancel", params, "{\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_setup_intent\",\"payout_method\":{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true},\"status\":\"requires_payout_method\",\"usage_intent\":\"payment\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_transfers", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementOutboundTransfers.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundTransferGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_transfers", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementOutboundTransfers.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementOutboundTransferPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundTransferFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundTransferToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_transfers", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundTransfers.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundTransferCreateFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundTransferCreateToParams{
			Currency:     stripe.String(stripe.CurrencyUSD),
			PayoutMethod: stripe.String("payout_method"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_transfers", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundTransfers.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_transfers/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundTransfers.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/outbound_transfers/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundTransfers.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferPost2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_transfers/id_123/cancel", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundTransfers.Cancel("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementOutboundTransferPost2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundTransferCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_transfers/id_123/cancel", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"cancelable\":true,\"created\":\"1970-01-12T21:42:34.472Z\",\"from\":{\"debited\":{\"currency\":\"USD\",\"value\":55},\"financial_account\":\"financial_account\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.outbound_transfer\",\"statement_descriptor\":\"statement_descriptor\",\"status\":\"canceled\",\"to\":{\"credited\":{\"currency\":\"USD\",\"value\":68},\"payout_method\":\"payout_method\"},\"trace_id\":{\"status\":\"pending\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundTransfers.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods", params, "{\"data\":[{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementPayoutMethods.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementPayoutMethodGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods", params, "{\"data\":[{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementPayoutMethods.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementPayoutMethodGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods/id_123", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementPayoutMethods.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods/id_123", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementPayoutMethods.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodPostService(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/payout_methods/id_123/archive", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementPayoutMethods.Archive("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodPostClient(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/payout_methods/id_123/archive", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementPayoutMethods.Archive(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodPost2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodUnarchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/payout_methods/id_123/unarchive", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementPayoutMethods.Unarchive("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodPost2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodUnarchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/payout_methods/id_123/unarchive", params, "{\"available_payout_speeds\":[\"standard\"],\"created\":\"1970-01-12T21:42:34.472Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.payout_method\",\"type\":\"bank_account\",\"usage_status\":{\"payments\":\"requires_action\",\"transfers\":\"invalid\"},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementPayoutMethods.Unarchive(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodsBankAccountSpecGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodsBankAccountSpecParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods_bank_account_spec", params, "{\"countries\":{\"key\":{\"fields\":[{\"local_name\":\"local_name\",\"local_name_human\":{\"content\":\"content\",\"localization_key\":\"localization_key\"},\"max_length\":1111390753,\"min_length\":711577229,\"placeholder\":\"placeholder\",\"stripe_name\":\"stripe_name\",\"validation_regex\":\"validation_regex\"}]}},\"object\":\"v2.money_management.payout_methods_bank_account_spec\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementPayoutMethodsBankAccountSpecs.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementPayoutMethodsBankAccountSpecGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementPayoutMethodsBankAccountSpecRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/payout_methods_bank_account_spec", params, "{\"countries\":{\"key\":{\"fields\":[{\"local_name\":\"local_name\",\"local_name_human\":{\"content\":\"content\",\"localization_key\":\"localization_key\"},\"max_length\":1111390753,\"min_length\":711577229,\"placeholder\":\"placeholder\",\"stripe_name\":\"stripe_name\",\"validation_regex\":\"validation_regex\"}]}},\"object\":\"v2.money_management.payout_methods_bank_account_spec\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementPayoutMethodsBankAccountSpecs.Retrieve(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementReceivedCreditGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedCreditListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_credits", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_credit\",\"status\":\"returned\",\"type\":\"balance_transfer\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementReceivedCredits.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementReceivedCreditGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedCreditListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_credits", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_credit\",\"status\":\"returned\",\"type\":\"balance_transfer\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementReceivedCredits.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementReceivedCreditGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedCreditParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_credits/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_credit\",\"status\":\"returned\",\"type\":\"balance_transfer\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementReceivedCredits.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementReceivedCreditGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedCreditRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_credits/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_credit\",\"status\":\"returned\",\"type\":\"balance_transfer\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementReceivedCredits.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementReceivedDebitGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedDebitListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_debits", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_debit\",\"status\":\"canceled\",\"type\":\"bank_transfer\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementReceivedDebits.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementReceivedDebitGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedDebitListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_debits", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_debit\",\"status\":\"canceled\",\"type\":\"bank_transfer\",\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementReceivedDebits.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementReceivedDebitGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedDebitParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_debits/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_debit\",\"status\":\"canceled\",\"type\":\"bank_transfer\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementReceivedDebits.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementReceivedDebitGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementReceivedDebitRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/received_debits/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"id\":\"obj_123\",\"object\":\"v2.money_management.received_debit\",\"status\":\"canceled\",\"type\":\"bank_transfer\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementReceivedDebits.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementTransactionGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transactions", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"category\":\"received_debit\",\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction\",\"status\":\"pending\",\"status_transitions\":{},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementTransactions.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementTransactionGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transactions", params, "{\"data\":[{\"amount\":{\"currency\":\"USD\",\"value\":96},\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"category\":\"received_debit\",\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction\",\"status\":\"pending\",\"status_transitions\":{},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementTransactions.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementTransactionGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transactions/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"category\":\"received_debit\",\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction\",\"status\":\"pending\",\"status_transitions\":{},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementTransactions.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementTransactionGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transactions/id_123", params, "{\"amount\":{\"currency\":\"USD\",\"value\":96},\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"category\":\"received_debit\",\"created\":\"1970-01-12T21:42:34.472Z\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"},\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction\",\"status\":\"pending\",\"status_transitions\":{},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementTransactions.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementTransactionEntryGetService(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionEntryListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transaction_entries", params, "{\"data\":[{\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction_entry\",\"transaction\":\"transaction\",\"transaction_details\":{\"category\":\"received_debit\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"}},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2MoneyManagementTransactionEntries.All(params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementTransactionEntryGetClient(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionEntryListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transaction_entries", params, "{\"data\":[{\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction_entry\",\"transaction\":\"transaction\",\"transaction_details\":{\"category\":\"received_debit\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"}},\"livemode\":true}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2MoneyManagementTransactionEntries.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2MoneyManagementTransactionEntryGet2Service(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionEntryParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transaction_entries/id_123", params, "{\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction_entry\",\"transaction\":\"transaction\",\"transaction_details\":{\"category\":\"received_debit\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"}},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementTransactionEntries.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2MoneyManagementTransactionEntryGet2Client(t *testing.T) {
	params := &stripe.V2MoneyManagementTransactionEntryRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/money_management/transaction_entries/id_123", params, "{\"balance_impact\":{\"available\":{\"currency\":\"USD\",\"value\":35},\"inbound_pending\":{\"currency\":\"USD\",\"value\":11},\"outbound_pending\":{\"currency\":\"USD\",\"value\":60}},\"created\":\"1970-01-12T21:42:34.472Z\",\"effective_at\":\"1970-01-03T20:38:28.043Z\",\"id\":\"obj_123\",\"object\":\"v2.money_management.transaction_entry\",\"transaction\":\"transaction\",\"transaction_details\":{\"category\":\"received_debit\",\"financial_account\":\"financial_account\",\"flow\":{\"type\":\"fee_transaction\"}},\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementTransactionEntries.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentGetService(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/payments/off_session_payments", params, "{\"data\":[{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result := sc.V2PaymentsOffSessionPayments.All(params)
	assert.NotNil(t, result)
}

func TestV2PaymentsOffSessionPaymentGetClient(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentListParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/payments/off_session_payments", params, "{\"data\":[{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}],\"next_page_url\":null,\"previous_page_url\":null}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result := sc.V2PaymentsOffSessionPayments.List(context.TODO(), params)
	assert.NotNil(t, result)
}

func TestV2PaymentsOffSessionPaymentPostService(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentParams{
		Amount:        &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		Cadence:       stripe.String("unscheduled"),
		Customer:      stripe.String("customer"),
		PaymentMethod: stripe.String("payment_method"),
	}
	params.AddMetadata("key", "metadata")
	testServer := MockServer(
		t, http.MethodPost, "/v2/payments/off_session_payments", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2PaymentsOffSessionPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentPostClient(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentCreateParams{
		Amount:        &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		Cadence:       stripe.String("unscheduled"),
		Customer:      stripe.String("customer"),
		PaymentMethod: stripe.String("payment_method"),
	}
	params.AddMetadata("key", "metadata")
	testServer := MockServer(
		t, http.MethodPost, "/v2/payments/off_session_payments", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2PaymentsOffSessionPayments.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentGet2Service(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/payments/off_session_payments/id_123", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2PaymentsOffSessionPayments.Get("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentGet2Client(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentRetrieveParams{}
	testServer := MockServer(
		t, http.MethodGet, "/v2/payments/off_session_payments/id_123", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2PaymentsOffSessionPayments.Retrieve(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentPost2Service(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/payments/off_session_payments/id_123/cancel", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2PaymentsOffSessionPayments.Cancel("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2PaymentsOffSessionPaymentPost2Client(t *testing.T) {
	params := &stripe.V2PaymentsOffSessionPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/payments/off_session_payments/id_123/cancel", params, "{\"amount_requested\":{\"currency\":\"USD\",\"value\":47},\"cadence\":\"unscheduled\",\"compartment_id\":\"compartment_id\",\"created\":\"1970-01-12T21:42:34.472Z\",\"customer\":\"customer\",\"id\":\"obj_123\",\"livemode\":true,\"metadata\":{\"key\":\"metadata\"},\"object\":\"v2.payments.off_session_payment\",\"payment_method\":\"payment_method\",\"retry_details\":{\"attempts\":542738246,\"retry_strategy\":\"none\"},\"status\":\"pending\"}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2PaymentsOffSessionPayments.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2TestHelpersFinancialAddressPostService(t *testing.T) {
	params := &stripe.V2TestHelpersFinancialAddressCreditParams{
		Amount:  &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		Network: stripe.String("rtp"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/test_helpers/financial_addresses/id_123/credit", params, "{\"object\":\"financial_address_credit_simulation\",\"status\":\"status\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2TestHelpersFinancialAddresses.Credit("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2TestHelpersFinancialAddressPostClient(t *testing.T) {
	params := &stripe.V2TestHelpersFinancialAddressCreditParams{
		Amount:  &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		Network: stripe.String("rtp"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/test_helpers/financial_addresses/id_123/credit", params, "{\"object\":\"financial_address_credit_simulation\",\"status\":\"status\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2TestHelpersFinancialAddresses.Credit(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2TestHelpersFinancialAddressPost2Service(t *testing.T) {
	params := &stripe.V2TestHelpersFinancialAddressGenerateMicrodepositsParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/test_helpers/financial_addresses/id_123/generate_microdeposits", params, "{\"amounts\":[{\"currency\":\"USD\",\"value\":1}],\"object\":\"financial_address_generated_microdeposits\",\"status\":\"accepted\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2TestHelpersFinancialAddresses.GenerateMicrodeposits(
		"id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestV2TestHelpersFinancialAddressPost2Client(t *testing.T) {
	params := &stripe.V2TestHelpersFinancialAddressGenerateMicrodepositsParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/test_helpers/financial_addresses/id_123/generate_microdeposits", params, "{\"amounts\":[{\"currency\":\"USD\",\"value\":1}],\"object\":\"financial_address_generated_microdeposits\",\"status\":\"accepted\",\"livemode\":true}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2TestHelpersFinancialAddresses.GenerateMicrodeposits(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAlreadyCanceledErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"error\":{\"type\":\"already_canceled\",\"code\":\"outbound_payment_already_canceled\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAlreadyCanceledErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"error\":{\"type\":\"already_canceled\",\"code\":\"outbound_payment_already_canceled\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAlreadyExistsErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"error\":{\"type\":\"already_exists\",\"code\":\"already_exists\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAlreadyExistsErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCreateParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"error\":{\"type\":\"already_exists\",\"code\":\"already_exists\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBlockedByStripeErrorService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"blocked_by_stripe\",\"code\":\"blocked_payout_method_bank_account\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBlockedByStripeErrorClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountCreateParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"blocked_by_stripe\",\"code\":\"blocked_payout_method_bank_account\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestControlledByDashboardErrorService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123/archive", params, "{\"error\":{\"type\":\"controlled_by_dashboard\",\"code\":\"bank_account_cannot_be_archived\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.Archive("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestControlledByDashboardErrorClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountArchiveParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts/id_123/archive", params, "{\"error\":{\"type\":\"controlled_by_dashboard\",\"code\":\"bank_account_cannot_be_archived\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Archive(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFeatureNotEnabledErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"error\":{\"type\":\"feature_not_enabled\",\"code\":\"recipient_feature_not_active\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFeatureNotEnabledErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCreateParams{
		Type: stripe.String("storage"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts", params, "{\"error\":{\"type\":\"feature_not_enabled\",\"code\":\"recipient_feature_not_active\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialAccountNotOpenErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressParams{
		Currency:         stripe.String(stripe.CurrencyUSD),
		FinancialAccount: stripe.String("financial_account"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_addresses", params, "{\"error\":{\"type\":\"financial_account_not_open\",\"code\":\"financial_account_not_in_open_status\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAddresses.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialAccountNotOpenErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAddressCreateParams{
		Currency:         stripe.String(stripe.CurrencyUSD),
		FinancialAccount: stripe.String("financial_account"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_addresses", params, "{\"error\":{\"type\":\"financial_account_not_open\",\"code\":\"financial_account_not_in_open_status\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAddresses.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInsufficientFundsErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentToParams{
			Recipient: stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"error\":{\"type\":\"insufficient_funds\",\"code\":\"insufficient_funds\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInsufficientFundsErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentCreateFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentCreateToParams{
			Recipient: stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"error\":{\"type\":\"insufficient_funds\",\"code\":\"insufficient_funds\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvalidPaymentMethodErrorService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"invalid_payment_method\",\"code\":\"invalid_us_bank_account\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvalidPaymentMethodErrorClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountCreateParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"invalid_payment_method\",\"code\":\"invalid_us_bank_account\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvalidPayoutMethodErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents", params, "{\"error\":{\"type\":\"invalid_payout_method\",\"code\":\"invalid_payout_method\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundSetupIntents.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvalidPayoutMethodErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentCreateParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_setup_intents", params, "{\"error\":{\"type\":\"invalid_payout_method\",\"code\":\"invalid_payout_method\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundSetupIntents.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestNonZeroBalanceErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts/id_123/close", params, "{\"error\":{\"type\":\"non_zero_balance\",\"code\":\"closing_financial_account_with_non_zero_balances\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementFinancialAccounts.Close("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestNonZeroBalanceErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementFinancialAccountCloseParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/financial_accounts/id_123/close", params, "{\"error\":{\"type\":\"non_zero_balance\",\"code\":\"closing_financial_account_with_non_zero_balances\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementFinancialAccounts.Close(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestNotCancelableErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"error\":{\"type\":\"not_cancelable\",\"code\":\"outbound_payment_not_cancelable\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel("id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestNotCancelableErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCancelParams{}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments/id_123/cancel", params, "{\"error\":{\"type\":\"not_cancelable\",\"code\":\"outbound_payment_not_cancelable\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Cancel(
		context.TODO(), "id_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotaExceededErrorService(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"quota_exceeded\",\"code\":\"archived_payout_method_card\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2CoreVaultUSBankAccounts.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotaExceededErrorClient(t *testing.T) {
	params := &stripe.V2CoreVaultUSBankAccountCreateParams{
		AccountNumber: stripe.String("account_number"),
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/core/vault/us_bank_accounts", params, "{\"error\":{\"type\":\"quota_exceeded\",\"code\":\"archived_payout_method_card\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2CoreVaultUSBankAccounts.Create(context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRecipientNotNotifiableErrorService(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentToParams{
			Recipient: stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"error\":{\"type\":\"recipient_not_notifiable\",\"code\":\"recipient_email_does_not_exist\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	result, err := sc.V2MoneyManagementOutboundPayments.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRecipientNotNotifiableErrorClient(t *testing.T) {
	params := &stripe.V2MoneyManagementOutboundPaymentCreateParams{
		Amount: &stripe.Amount{Value: 96, Currency: stripe.CurrencyUSD},
		From: &stripe.V2MoneyManagementOutboundPaymentCreateFromParams{
			Currency:         stripe.String(stripe.CurrencyUSD),
			FinancialAccount: stripe.String("financial_account"),
		},
		To: &stripe.V2MoneyManagementOutboundPaymentCreateToParams{
			Recipient: stripe.String("recipient"),
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/money_management/outbound_payments", params, "{\"error\":{\"type\":\"recipient_not_notifiable\",\"code\":\"recipient_email_does_not_exist\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	result, err := sc.V2MoneyManagementOutboundPayments.Create(
		context.TODO(), params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTemporarySessionExpiredErrorService(t *testing.T) {
	params := &stripe.V2BillingMeterEventStreamParams{
		Events: []*stripe.V2BillingMeterEventStreamEventParams{
			{
				EventName: stripe.String("event_name"),
				Payload:   map[string]string{"key": "payload"},
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_stream", params, "{\"error\":{\"type\":\"temporary_session_expired\",\"code\":\"billing_meter_event_session_expired\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := client.New(TestAPIKey, backends)
	sc.V2BillingMeterEventStreams.New(params)
}

func TestTemporarySessionExpiredErrorClient(t *testing.T) {
	params := &stripe.V2BillingMeterEventStreamCreateParams{
		Events: []*stripe.V2BillingMeterEventStreamCreateEventParams{
			{
				EventName: stripe.String("event_name"),
				Payload:   map[string]string{"key": "payload"},
			},
		},
	}
	testServer := MockServer(
		t, http.MethodPost, "/v2/billing/meter_event_stream", params, "{\"error\":{\"type\":\"temporary_session_expired\",\"code\":\"billing_meter_event_session_expired\"}}")
	defer testServer.Close()
	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{URL: &testServer.URL})
	sc := stripe.NewClient(TestAPIKey, stripe.WithBackends(backends))
	sc.V2BillingMeterEventStreams.Create(context.TODO(), params)
}
