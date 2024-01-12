//
//
// File generated from our OpenAPI spec
//
//

package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v76"
	account "github.com/stripe/stripe-go/v76/account"
	accountlink "github.com/stripe/stripe-go/v76/accountlink"
	applicationfee "github.com/stripe/stripe-go/v76/applicationfee"
	apps_secret "github.com/stripe/stripe-go/v76/apps/secret"
	balancetransaction "github.com/stripe/stripe-go/v76/balancetransaction"
	billingportal_configuration "github.com/stripe/stripe-go/v76/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v76/billingportal/session"
	capability "github.com/stripe/stripe-go/v76/capability"
	card "github.com/stripe/stripe-go/v76/card"
	cashbalance "github.com/stripe/stripe-go/v76/cashbalance"
	charge "github.com/stripe/stripe-go/v76/charge"
	checkout_session "github.com/stripe/stripe-go/v76/checkout/session"
	countryspec "github.com/stripe/stripe-go/v76/countryspec"
	coupon "github.com/stripe/stripe-go/v76/coupon"
	customer "github.com/stripe/stripe-go/v76/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v76/customerbalancetransaction"
	customersession "github.com/stripe/stripe-go/v76/customersession"
	dispute "github.com/stripe/stripe-go/v76/dispute"
	event "github.com/stripe/stripe-go/v76/event"
	feerefund "github.com/stripe/stripe-go/v76/feerefund"
	financialconnections_account "github.com/stripe/stripe-go/v76/financialconnections/account"
	financialconnections_session "github.com/stripe/stripe-go/v76/financialconnections/session"
	financialconnections_transaction "github.com/stripe/stripe-go/v76/financialconnections/transaction"
	identity_verificationreport "github.com/stripe/stripe-go/v76/identity/verificationreport"
	identity_verificationsession "github.com/stripe/stripe-go/v76/identity/verificationsession"
	invoice "github.com/stripe/stripe-go/v76/invoice"
	invoiceitem "github.com/stripe/stripe-go/v76/invoiceitem"
	issuing_authorization "github.com/stripe/stripe-go/v76/issuing/authorization"
	issuing_card "github.com/stripe/stripe-go/v76/issuing/card"
	issuing_cardholder "github.com/stripe/stripe-go/v76/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v76/issuing/dispute"
	issuing_transaction "github.com/stripe/stripe-go/v76/issuing/transaction"
	loginlink "github.com/stripe/stripe-go/v76/loginlink"
	mandate "github.com/stripe/stripe-go/v76/mandate"
	paymentintent "github.com/stripe/stripe-go/v76/paymentintent"
	paymentlink "github.com/stripe/stripe-go/v76/paymentlink"
	paymentmethod "github.com/stripe/stripe-go/v76/paymentmethod"
	paymentmethodconfiguration "github.com/stripe/stripe-go/v76/paymentmethodconfiguration"
	paymentsource "github.com/stripe/stripe-go/v76/paymentsource"
	payout "github.com/stripe/stripe-go/v76/payout"
	person "github.com/stripe/stripe-go/v76/person"
	plan "github.com/stripe/stripe-go/v76/plan"
	price "github.com/stripe/stripe-go/v76/price"
	product "github.com/stripe/stripe-go/v76/product"
	promotioncode "github.com/stripe/stripe-go/v76/promotioncode"
	quote "github.com/stripe/stripe-go/v76/quote"
	radar_earlyfraudwarning "github.com/stripe/stripe-go/v76/radar/earlyfraudwarning"
	radar_valuelist "github.com/stripe/stripe-go/v76/radar/valuelist"
	radar_valuelistitem "github.com/stripe/stripe-go/v76/radar/valuelistitem"
	refund "github.com/stripe/stripe-go/v76/refund"
	reporting_reportrun "github.com/stripe/stripe-go/v76/reporting/reportrun"
	reporting_reporttype "github.com/stripe/stripe-go/v76/reporting/reporttype"
	review "github.com/stripe/stripe-go/v76/review"
	setupattempt "github.com/stripe/stripe-go/v76/setupattempt"
	setupintent "github.com/stripe/stripe-go/v76/setupintent"
	shippingrate "github.com/stripe/stripe-go/v76/shippingrate"
	sigma_scheduledqueryrun "github.com/stripe/stripe-go/v76/sigma/scheduledqueryrun"
	source "github.com/stripe/stripe-go/v76/source"
	subscription "github.com/stripe/stripe-go/v76/subscription"
	subscriptionitem "github.com/stripe/stripe-go/v76/subscriptionitem"
	subscriptionschedule "github.com/stripe/stripe-go/v76/subscriptionschedule"
	tax_calculation "github.com/stripe/stripe-go/v76/tax/calculation"
	tax_form "github.com/stripe/stripe-go/v76/tax/form"
	tax_settings "github.com/stripe/stripe-go/v76/tax/settings"
	tax_transaction "github.com/stripe/stripe-go/v76/tax/transaction"
	taxcode "github.com/stripe/stripe-go/v76/taxcode"
	taxid "github.com/stripe/stripe-go/v76/taxid"
	taxrate "github.com/stripe/stripe-go/v76/taxrate"
	terminal_configuration "github.com/stripe/stripe-go/v76/terminal/configuration"
	terminal_connectiontoken "github.com/stripe/stripe-go/v76/terminal/connectiontoken"
	terminal_location "github.com/stripe/stripe-go/v76/terminal/location"
	terminal_reader "github.com/stripe/stripe-go/v76/terminal/reader"
	testhelpers_customer "github.com/stripe/stripe-go/v76/testhelpers/customer"
	testhelpers_issuing_authorization "github.com/stripe/stripe-go/v76/testhelpers/issuing/authorization"
	testhelpers_issuing_card "github.com/stripe/stripe-go/v76/testhelpers/issuing/card"
	testhelpers_issuing_transaction "github.com/stripe/stripe-go/v76/testhelpers/issuing/transaction"
	testhelpers_refund "github.com/stripe/stripe-go/v76/testhelpers/refund"
	testhelpers_testclock "github.com/stripe/stripe-go/v76/testhelpers/testclock"
	testhelpers_treasury_inboundtransfer "github.com/stripe/stripe-go/v76/testhelpers/treasury/inboundtransfer"
	testhelpers_treasury_outboundtransfer "github.com/stripe/stripe-go/v76/testhelpers/treasury/outboundtransfer"
	testhelpers_treasury_receivedcredit "github.com/stripe/stripe-go/v76/testhelpers/treasury/receivedcredit"
	testhelpers_treasury_receiveddebit "github.com/stripe/stripe-go/v76/testhelpers/treasury/receiveddebit"
	_ "github.com/stripe/stripe-go/v76/testing"
	token "github.com/stripe/stripe-go/v76/token"
	topup "github.com/stripe/stripe-go/v76/topup"
	transfer "github.com/stripe/stripe-go/v76/transfer"
	transferreversal "github.com/stripe/stripe-go/v76/transferreversal"
	treasury_creditreversal "github.com/stripe/stripe-go/v76/treasury/creditreversal"
	treasury_debitreversal "github.com/stripe/stripe-go/v76/treasury/debitreversal"
	treasury_financialaccount "github.com/stripe/stripe-go/v76/treasury/financialaccount"
	treasury_inboundtransfer "github.com/stripe/stripe-go/v76/treasury/inboundtransfer"
	treasury_outboundpayment "github.com/stripe/stripe-go/v76/treasury/outboundpayment"
	treasury_outboundtransfer "github.com/stripe/stripe-go/v76/treasury/outboundtransfer"
	treasury_receivedcredit "github.com/stripe/stripe-go/v76/treasury/receivedcredit"
	treasury_receiveddebit "github.com/stripe/stripe-go/v76/treasury/receiveddebit"
	treasury_transaction "github.com/stripe/stripe-go/v76/treasury/transaction"
	treasury_transactionentry "github.com/stripe/stripe-go/v76/treasury/transactionentry"
	usagerecord "github.com/stripe/stripe-go/v76/usagerecord"
	webhookendpoint "github.com/stripe/stripe-go/v76/webhookendpoint"
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

func TestAccountsCapabilitiesGet(t *testing.T) {
	params := &stripe.CapabilityListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result := capability.List(params)
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

func TestAccountsCapabilitiesPost(t *testing.T) {
	params := &stripe.CapabilityParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := capability.Update("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsDelete(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.Del("acct_xxxxxxxxxxxxx", params)
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

func TestAccountsGet2(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.GetByID("acct_xxxxxxxxxxxxx", params)
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

func TestAccountsPersonsDelete(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := person.Del("person_xxxxxxxxxxxxx", params)
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

func TestAccountsPersonsGet2(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	result, err := person.Get("person_xxxxxxxxxxxxx", params)
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

func TestAccountsPersonsPost2(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := person.Update("person_xxxxxxxxxxxxx", params)
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

func TestAccountsPost2(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountsRejectPost(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, err := account.Reject("acct_xxxxxxxxxxxxx", params)
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

func TestApplicationFeesGet2(t *testing.T) {
	params := &stripe.ApplicationFeeParams{}
	result, err := applicationfee.Get("fee_xxxxxxxxxxxxx", params)
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

func TestApplicationFeesRefundsGet2(t *testing.T) {
	params := &stripe.FeeRefundParams{Fee: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := feerefund.Get("fr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestApplicationFeesRefundsPost(t *testing.T) {
	params := &stripe.FeeRefundParams{ID: stripe.String("fee_xxxxxxxxxxxxx")}
	result, err := feerefund.New(params)
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

func TestBalanceTransactionsGet(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBalanceTransactionsGet2(t *testing.T) {
	params := &stripe.BalanceTransactionParams{}
	result, err := balancetransaction.Get("txn_xxxxxxxxxxxxx", params)
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

func TestBillingPortalConfigurationsGet2(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{}
	result, err := billingportal_configuration.Get("bpc_xxxxxxxxxxxxx", params)
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

func TestBillingPortalConfigurationsPost2(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, err := billingportal_configuration.Update(
		"bpc_xxxxxxxxxxxxx",
		params,
	)
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

func TestChargesCapturePost(t *testing.T) {
	params := &stripe.ChargeCaptureParams{}
	result, err := charge.Capture("ch_xxxxxxxxxxxxx", params)
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

func TestChargesGet2(t *testing.T) {
	params := &stripe.ChargeParams{}
	result, err := charge.Get("ch_xxxxxxxxxxxxx", params)
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

func TestChargesPost2(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := charge.Update("ch_xxxxxxxxxxxxx", params)
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

func TestCheckoutSessionsExpirePost(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("sess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionsExpirePost2(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("cs_test_xxxxxxxxxxxxx", params)
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

func TestCheckoutSessionsGet2(t *testing.T) {
	params := &stripe.CheckoutSessionParams{}
	result, err := checkout_session.Get("cs_test_xxxxxxxxxxxxx", params)
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

func TestCountrySpecsGet(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Limit = stripe.Int64(3)
	result := countryspec.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCountrySpecsGet2(t *testing.T) {
	params := &stripe.CountrySpecParams{}
	result, err := countryspec.Get("US", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsDelete(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Del("Z4OV52SU", params)
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

func TestCouponsGet2(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Get("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponsPost(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff:       stripe.Float64(25.5),
		Duration:         stripe.String(string(stripe.CouponDurationRepeating)),
		DurationInMonths: stripe.Int64(3),
	}
	result, err := coupon.New(params)
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

func TestCustomersBalanceTransactionsGet(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := customerbalancetransaction.List(params)
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

func TestCustomersBalanceTransactionsPost2(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := customerbalancetransaction.Update(
		"cbtxn_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersCashBalanceGet(t *testing.T) {
	params := &stripe.CashBalanceParams{Customer: stripe.String("cus_123")}
	result, err := cashbalance.Get(params)
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

func TestCustomersDelete(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Del("cus_xxxxxxxxxxxxx", params)
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

func TestCustomersGet(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := customer.List(params)
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

func TestCustomersGet3(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Get("cus_xxxxxxxxxxxxx", params)
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

func TestCustomersPaymentMethodsGet2(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result := customer.ListPaymentMethods(params)
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

func TestCustomersPost2(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, err := customer.Update("cus_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesDelete(t *testing.T) {
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := card.Del("ba_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomersSourcesDelete2(t *testing.T) {
	params := &stripe.CardParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := card.Del("card_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesGet3(t *testing.T) {
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentsource.Get("ba_xxxxxxxxxxxxx", params)
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

func TestCustomersSourcesPost(t *testing.T) {
	params := &stripe.CardParams{
		AccountHolderName: stripe.String("Kamil"),
		Customer:          stripe.String("cus_123"),
	}
	result, err := card.Update("card_123", params)
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

func TestCustomersTaxIdsDelete(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Del("txi_xxxxxxxxxxxxx", params)
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

func TestCustomersTaxIdsGet2(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Get("txi_xxxxxxxxxxxxx", params)
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

func TestDisputesClosePost(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, err := dispute.Close("dp_xxxxxxxxxxxxx", params)
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

func TestDisputesGet2(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, err := dispute.Get("dp_xxxxxxxxxxxxx", params)
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

func TestEventsGet(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Limit = stripe.Int64(3)
	result := event.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestEventsGet2(t *testing.T) {
	params := &stripe.EventParams{}
	result, err := event.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsDisconnectPost2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect(
		"fca_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListParams{}
	result := financialconnections_account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountsGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := financialconnections_account.GetByID("fca_xyz", params)
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

func TestFinancialConnectionsAccountsGet4(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, err := financialconnections_account.GetByID(
		"fca_xxxxxxxxxxxxx",
		params,
	)
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

func TestFinancialConnectionsAccountsRefreshPost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountRefreshParams{
		Features: []*string{stripe.String("balance")},
	}
	result, err := financialconnections_account.Refresh("fca_xyz", params)
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

func TestFinancialConnectionsAccountsUnsubscribePost(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountUnsubscribeParams{
		Features: []*string{stripe.String("transactions")},
	}
	result, err := financialconnections_account.Unsubscribe("fa_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get("fcsess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionsGet2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get(
		"fcsess_xxxxxxxxxxxxx",
		params,
	)
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

func TestFinancialConnectionsTransactionsGet(t *testing.T) {
	params := &stripe.FinancialConnectionsTransactionParams{}
	result, err := financialconnections_transaction.Get("tr_123", params)
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

func TestIdentityVerificationReportsGet(t *testing.T) {
	params := &stripe.IdentityVerificationReportListParams{}
	params.Limit = stripe.Int64(3)
	result := identity_verificationreport.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationReportsGet2(t *testing.T) {
	params := &stripe.IdentityVerificationReportParams{}
	result, err := identity_verificationreport.Get("vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsCancelPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := identity_verificationsession.Cancel(
		"vs_xxxxxxxxxxxxx",
		params,
	)
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

func TestIdentityVerificationSessionsGet2(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{}
	result, err := identity_verificationsession.Get("vs_xxxxxxxxxxxxx", params)
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

func TestIdentityVerificationSessionsPost2(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeIDNumber)),
	}
	result, err := identity_verificationsession.Update(
		"vs_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionsRedactPost(t *testing.T) {
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := identity_verificationsession.Redact(
		"vs_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsDelete(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, err := invoiceitem.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceitemsPost(t *testing.T) {
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Price:    stripe.String("price_xxxxxxxxxxxxx"),
	}
	result, err := invoiceitem.New(params)
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

func TestInvoicesDelete(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesFinalizePost(t *testing.T) {
	params := &stripe.InvoiceFinalizeInvoiceParams{}
	result, err := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
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

func TestInvoicesGet2(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Get("in_xxxxxxxxxxxxx", params)
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

func TestInvoicesPayPost(t *testing.T) {
	params := &stripe.InvoicePayParams{}
	result, err := invoice.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesPost(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := invoice.New(params)
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

func TestInvoicesSendPost(t *testing.T) {
	params := &stripe.InvoiceSendInvoiceParams{}
	result, err := invoice.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicesVoidPost(t *testing.T) {
	params := &stripe.InvoiceVoidInvoiceParams{}
	result, err := invoice.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsApprovePost(t *testing.T) {
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, err := issuing_authorization.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationsDeclinePost(t *testing.T) {
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, err := issuing_authorization.Decline("iauth_xxxxxxxxxxxxx", params)
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

func TestIssuingAuthorizationsGet2(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	result, err := issuing_authorization.Get("iauth_xxxxxxxxxxxxx", params)
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

func TestIssuingCardholdersGet(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_cardholder.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardholdersGet2(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	result, err := issuing_cardholder.Get("ich_xxxxxxxxxxxxx", params)
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

func TestIssuingCardholdersPost2(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_cardholder.Update("ich_xxxxxxxxxxxxx", params)
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

func TestIssuingCardsGet2(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	result, err := issuing_card.Get("ic_xxxxxxxxxxxxx", params)
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

func TestIssuingCardsPost2(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_card.Update("ic_xxxxxxxxxxxxx", params)
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

func TestIssuingDisputesGet2(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	result, err := issuing_dispute.Get("idp_xxxxxxxxxxxxx", params)
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

func TestIssuingDisputesSubmitPost(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	result, err := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", params)
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

func TestIssuingTransactionsGet2(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	result, err := issuing_transaction.Get("ipi_xxxxxxxxxxxxx", params)
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

func TestMandatesGet(t *testing.T) {
	params := &stripe.MandateParams{}
	result, err := mandate.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsApplyCustomerBalancePost(t *testing.T) {
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, err := paymentintent.ApplyCustomerBalance("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCancelPost(t *testing.T) {
	params := &stripe.PaymentIntentCancelParams{}
	result, err := paymentintent.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsCapturePost(t *testing.T) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := paymentintent.Capture("pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsGet(t *testing.T) {
	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := paymentintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentsGet2(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	result, err := paymentintent.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentsIncrementAuthorizationPost(t *testing.T) {
	params := &stripe.PaymentIntentIncrementAuthorizationParams{
		Amount: stripe.Int64(2099),
	}
	result, err := paymentintent.IncrementAuthorization(
		"pi_xxxxxxxxxxxxx",
		params,
	)
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

func TestPaymentIntentsPost3(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
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

func TestPaymentIntentsVerifyMicrodepositsPost(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, err := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
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

func TestPaymentLinksGet(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("pl_xyz", params)
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

func TestPaymentLinksGet3(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("plink_xxxxxxxxxxxxx", params)
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

func TestPaymentLinksPost3(t *testing.T) {
	params := &stripe.PaymentLinkParams{Active: stripe.Bool(false)}
	result, err := paymentlink.Update("plink_xxxxxxxxxxxxx", params)
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

func TestPaymentMethodConfigurationsGet2(t *testing.T) {
	params := &stripe.PaymentMethodConfigurationParams{}
	result, err := paymentmethodconfiguration.Get("foo", params)
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

func TestPaymentMethodsAttachPost(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodsDetachPost(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	result, err := paymentmethod.Detach("pm_xxxxxxxxxxxxx", params)
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

func TestPaymentMethodsGet2(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	result, err := paymentmethod.Get("pm_xxxxxxxxxxxxx", params)
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

func TestPaymentMethodsPost2(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsCancelPost(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Cancel("po_xxxxxxxxxxxxx", params)
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

func TestPayoutsGet2(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Get("po_xxxxxxxxxxxxx", params)
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

func TestPayoutsPost2(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, err := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutsReversePost(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	result, err := payout.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlansDelete(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Del("price_xxxxxxxxxxxxx", params)
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

func TestPlansGet2(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Get("price_xxxxxxxxxxxxx", params)
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

func TestPlansPost2(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, err := plan.Update("price_xxxxxxxxxxxxx", params)
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

func TestPricesGet2(t *testing.T) {
	params := &stripe.PriceParams{}
	result, err := price.Get("price_xxxxxxxxxxxxx", params)
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

func TestPricesPost3(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := price.Update("price_xxxxxxxxxxxxx", params)
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

func TestProductsDelete(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Del("prod_xxxxxxxxxxxxx", params)
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

func TestProductsGet2(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductsPost(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, err := product.New(params)
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

func TestPromotionCodesGet(t *testing.T) {
	params := &stripe.PromotionCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := promotioncode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodesGet2(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	result, err := promotioncode.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodesPost(t *testing.T) {
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("Z4OV52SU")}
	result, err := promotioncode.New(params)
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

func TestQuotesAcceptPost(t *testing.T) {
	params := &stripe.QuoteAcceptParams{}
	result, err := quote.Accept("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesCancelPost(t *testing.T) {
	params := &stripe.QuoteCancelParams{}
	result, err := quote.Cancel("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuotesFinalizePost(t *testing.T) {
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, err := quote.FinalizeQuote("qt_xxxxxxxxxxxxx", params)
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

func TestQuotesGet2(t *testing.T) {
	params := &stripe.QuoteParams{}
	result, err := quote.Get("qt_xxxxxxxxxxxxx", params)
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

func TestQuotesPdfGet(t *testing.T) {
	params := &stripe.QuotePDFParams{}
	result, err := quote.PDF("qt_xxxxxxxxxxxxx", params)
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

func TestQuotesPost2(t *testing.T) {
	params := &stripe.QuoteParams{}
	params.AddMetadata("order_id", "6735")
	result, err := quote.Update("qt_xxxxxxxxxxxxx", params)
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

func TestRadarEarlyFraudWarningsGet(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Limit = stripe.Int64(3)
	result := radar_earlyfraudwarning.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarEarlyFraudWarningsGet2(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningParams{}
	result, err := radar_earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRadarValueListItemsDelete(t *testing.T) {
	params := &stripe.RadarValueListItemParams{}
	result, err := radar_valuelistitem.Del("rsli_xxxxxxxxxxxxx", params)
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

func TestRadarValueListItemsGet2(t *testing.T) {
	params := &stripe.RadarValueListItemParams{}
	result, err := radar_valuelistitem.Get("rsli_xxxxxxxxxxxxx", params)
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

func TestRadarValueListsDelete(t *testing.T) {
	params := &stripe.RadarValueListParams{}
	result, err := radar_valuelist.Del("rsl_xxxxxxxxxxxxx", params)
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

func TestRadarValueListsGet2(t *testing.T) {
	params := &stripe.RadarValueListParams{}
	result, err := radar_valuelist.Get("rsl_xxxxxxxxxxxxx", params)
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

func TestRadarValueListsPost2(t *testing.T) {
	params := &stripe.RadarValueListParams{
		Name: stripe.String("Updated IP Block List"),
	}
	result, err := radar_valuelist.Update("rsl_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsCancelPost(t *testing.T) {
	params := &stripe.RefundCancelParams{}
	result, err := refund.Cancel("re_xxxxxxxxxxxxx", params)
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

func TestRefundsGet2(t *testing.T) {
	params := &stripe.RefundParams{}
	result, err := refund.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundsPost(t *testing.T) {
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, err := refund.New(params)
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

func TestReportingReportRunsGet(t *testing.T) {
	params := &stripe.ReportingReportRunListParams{}
	params.Limit = stripe.Int64(3)
	result := reporting_reportrun.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportRunsGet2(t *testing.T) {
	params := &stripe.ReportingReportRunParams{}
	result, err := reporting_reportrun.Get("frr_xxxxxxxxxxxxx", params)
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

func TestReportingReportTypesGet(t *testing.T) {
	params := &stripe.ReportingReportTypeListParams{}
	result := reporting_reporttype.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReportingReportTypesGet2(t *testing.T) {
	params := &stripe.ReportingReportTypeParams{}
	result, err := reporting_reporttype.Get("balance.summary.1", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewsApprovePost(t *testing.T) {
	params := &stripe.ReviewApproveParams{}
	result, err := review.Approve("prv_xxxxxxxxxxxxx", params)
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

func TestReviewsGet2(t *testing.T) {
	params := &stripe.ReviewParams{}
	result, err := review.Get("prv_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsCancelPost(t *testing.T) {
	params := &stripe.SetupIntentCancelParams{}
	result, err := setupintent.Cancel("seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsGet(t *testing.T) {
	params := &stripe.SetupIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := setupintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentsGet2(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	result, err := setupintent.Get("seti_xxxxxxxxxxxxx", params)
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

func TestSetupIntentsPost2(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, err := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentsVerifyMicrodepositsPost(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, err := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
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

func TestShippingRatesGet(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	result := shippingrate.List(params)
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

func TestShippingRatesGet3(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	result, err := shippingrate.Get("shr_xxxxxxxxxxxxx", params)
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

func TestShippingRatesPost3(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := shippingrate.Update("shr_xxxxxxxxxxxxx", params)
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

func TestSigmaScheduledQueryRunsGet2(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunParams{}
	result, err := sigma_scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGet(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourcesGet2(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsDelete(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Del("si_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsGet2(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Get("si_xxxxxxxxxxxxx", params)
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

func TestSubscriptionItemsPost2(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscriptionitem.Update("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemsUsageRecordSummariesGet(t *testing.T) {
	params := &stripe.SubscriptionItemUsageRecordSummariesParams{
		SubscriptionItem: stripe.String("si_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := subscriptionitem.UsageRecordSummaries(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionItemsUsageRecordsPost(t *testing.T) {
	params := &stripe.UsageRecordParams{
		Quantity:         stripe.Int64(100),
		Timestamp:        stripe.Int64(1571252444),
		SubscriptionItem: stripe.String("si_xxxxxxxxxxxxx"),
	}
	result, err := usagerecord.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesCancelPost(t *testing.T) {
	params := &stripe.SubscriptionScheduleCancelParams{}
	result, err := subscriptionschedule.Cancel("sub_sched_xxxxxxxxxxxxx", params)
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

func TestSubscriptionSchedulesGet2(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{}
	result, err := subscriptionschedule.Get("sub_sched_xxxxxxxxxxxxx", params)
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

func TestSubscriptionSchedulesPost2(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
	}
	result, err := subscriptionschedule.Update("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSchedulesReleasePost(t *testing.T) {
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := subscriptionschedule.Release(
		"sub_sched_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDelete(t *testing.T) {
	params := &stripe.SubscriptionCancelParams{}
	result, err := subscription.Cancel("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionsDiscountDelete(t *testing.T) {
	params := &stripe.SubscriptionDeleteDiscountParams{}
	result, err := subscription.DeleteDiscount("sub_xyz", params)
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

func TestSubscriptionsGet2(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	result, err := subscription.Get("sub_xxxxxxxxxxxxx", params)
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

func TestSubscriptionsPost2(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscription.Update("sub_xxxxxxxxxxxxx", params)
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

func TestTaxCalculationsLineItemsGet(t *testing.T) {
	params := &stripe.TaxCalculationListLineItemsParams{
		Calculation: stripe.String("xxx"),
	}
	result := tax_calculation.ListLineItems(params)
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

func TestTaxCodesGet(t *testing.T) {
	params := &stripe.TaxCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := taxcode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCodesGet2(t *testing.T) {
	params := &stripe.TaxCodeParams{}
	result, err := taxcode.Get("txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxFormsPdfGet(t *testing.T) {
	params := &stripe.TaxFormPDFParams{}
	result, err := tax_form.PDF("form_xxxxxxxxxxxxx", params)
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

func TestTaxRatesGet2(t *testing.T) {
	params := &stripe.TaxRateParams{}
	result, err := taxrate.Get("txr_xxxxxxxxxxxxx", params)
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

func TestTaxRatesPost2(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, err := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxSettingsGet(t *testing.T) {
	params := &stripe.TaxSettingsParams{}
	result, err := tax_settings.Get(params)
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

func TestTaxTransactionsCreateFromCalculationPost(t *testing.T) {
	params := &stripe.TaxTransactionCreateFromCalculationParams{
		Calculation: stripe.String("xxx"),
		Reference:   stripe.String("yyy"),
	}
	result, err := tax_transaction.CreateFromCalculation(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDelete(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsDelete2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsGet(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationsGet2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("uc_123", params)
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

func TestTerminalConfigurationsGet4(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationsPost(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.New(params)
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

func TestTerminalConnectionTokensPost(t *testing.T) {
	params := &stripe.TerminalConnectionTokenParams{}
	result, err := terminal_connectiontoken.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationsDelete(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Del("tml_xxxxxxxxxxxxx", params)
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

func TestTerminalLocationsGet2(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Get("tml_xxxxxxxxxxxxx", params)
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

func TestTerminalLocationsPost2(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, err := terminal_location.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersCancelActionPost(t *testing.T) {
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := terminal_reader.CancelAction("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersDelete(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, err := terminal_reader.Del("tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersGet2(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, err := terminal_reader.Get("tmr_xxxxxxxxxxxxx", params)
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

func TestTerminalReadersPost2(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, err := terminal_reader.Update("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessPaymentIntentPost(t *testing.T) {
	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String("pi_xxxxxxxxxxxxx"),
	}
	result, err := terminal_reader.ProcessPaymentIntent(
		"tmr_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReadersProcessSetupIntentPost(t *testing.T) {
	params := &stripe.TerminalReaderProcessSetupIntentParams{
		SetupIntent:              stripe.String("seti_xxxxxxxxxxxxx"),
		CustomerConsentCollected: stripe.Bool(true),
	}
	result, err := terminal_reader.ProcessSetupIntent(
		"tmr_xxxxxxxxxxxxx",
		params,
	)
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
				VolumeDecimal:   stripe.Float64(10),
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
		"example_authorization",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsExpirePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationExpireParams{}
	result, err := testhelpers_issuing_authorization.Expire(
		"example_authorization",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingAuthorizationsIncrementPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationIncrementParams{
		IncrementAmount:      stripe.Int64(50),
		IsAmountControllable: stripe.Bool(true),
	}
	result, err := testhelpers_issuing_authorization.Increment(
		"example_authorization",
		params,
	)
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

func TestTestHelpersIssuingAuthorizationsReversePost(t *testing.T) {
	params := &stripe.TestHelpersIssuingAuthorizationReverseParams{
		ReverseAmount: stripe.Int64(20),
	}
	result, err := testhelpers_issuing_authorization.Reverse(
		"example_authorization",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingDeliverPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, err := testhelpers_issuing_card.DeliverCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingFailPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, err := testhelpers_issuing_card.FailCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingReturnPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, err := testhelpers_issuing_card.ReturnCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardsShippingShipPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, err := testhelpers_issuing_card.ShipCard("card_123", params)
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
				VolumeDecimal:   stripe.Float64(10),
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
				VolumeDecimal:   stripe.Float64(10),
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

func TestTestHelpersIssuingTransactionsRefundPost(t *testing.T) {
	params := &stripe.TestHelpersIssuingTransactionRefundParams{
		RefundAmount: stripe.Int64(50),
	}
	result, err := testhelpers_issuing_transaction.Refund(
		"example_transaction",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersRefundsExpirePost(t *testing.T) {
	params := &stripe.TestHelpersRefundExpireParams{}
	result, err := testhelpers_refund.Expire("re_123", params)
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

func TestTestHelpersTestClocksAdvancePost2(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1675552261),
	}
	result, err := testhelpers_testclock.Advance("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDelete(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksDelete2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClocksGet(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClocksGet2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xyz", params)
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

func TestTestHelpersTestClocksGet4(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xxxxxxxxxxxxx", params)
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

func TestTestHelpersTestClocksPost2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(1577836800),
	}
	result, err := testhelpers_testclock.New(params)
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

func TestTestHelpersTreasuryInboundTransfersReturnPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, err := testhelpers_treasury_inboundtransfer.ReturnInboundTransfer(
		"ibt_123",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransfersSucceedPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := testhelpers_treasury_inboundtransfer.Succeed(
		"ibt_123",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersFailPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, err := testhelpers_treasury_outboundtransfer.Fail("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransfersPostPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, err := testhelpers_treasury_outboundtransfer.Post("obt_123", params)
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
		"obt_123",
		params,
	)
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

func TestTokensGet(t *testing.T) {
	params := &stripe.TokenParams{}
	result, err := token.Get("tok_xxxx", params)
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

func TestTokensPost3(t *testing.T) {
	params := &stripe.TokenParams{
		PII: &stripe.TokenPIIParams{IDNumber: stripe.String("000000000")},
	}
	result, err := token.New(params)
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

func TestTokensPost6(t *testing.T) {
	params := &stripe.TokenParams{
		CVCUpdate: &stripe.TokenCVCUpdateParams{CVC: stripe.String("123")},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupsCancelPost(t *testing.T) {
	params := &stripe.TopupParams{}
	result, err := topup.Cancel("tu_xxxxxxxxxxxxx", params)
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

func TestTopupsGet2(t *testing.T) {
	params := &stripe.TopupParams{}
	result, err := topup.Get("tu_xxxxxxxxxxxxx", params)
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

func TestTopupsPost2(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, err := topup.Update("tu_xxxxxxxxxxxxx", params)
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

func TestTransfersGet2(t *testing.T) {
	params := &stripe.TransferParams{}
	result, err := transfer.Get("tr_xxxxxxxxxxxxx", params)
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

func TestTransfersPost2(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, err := transfer.Update("tr_xxxxxxxxxxxxx", params)
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

func TestTransfersReversalsGet2(t *testing.T) {
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := transferreversal.Get("trr_xxxxxxxxxxxxx", params)
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

func TestTransfersReversalsPost2(t *testing.T) {
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := transferreversal.Update("trr_xxxxxxxxxxxxx", params)
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

func TestTreasuryCreditReversalsGet2(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{}
	result, err := treasury_creditreversal.Get("credrev_xxxxxxxxxxxxx", params)
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

func TestTreasuryDebitReversalsGet(t *testing.T) {
	params := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_debitreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryDebitReversalsGet2(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{}
	result, err := treasury_debitreversal.Get("debrev_xxxxxxxxxxxxx", params)
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

func TestTreasuryFinancialAccountsFeaturesGet(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, err := treasury_financialaccount.RetrieveFeatures(
		"fa_xxxxxxxxxxxxx",
		params,
	)
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

func TestTreasuryFinancialAccountsGet2(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	result, err := treasury_financialaccount.Get("fa_xxxxxxxxxxxxx", params)
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

func TestTreasuryFinancialAccountsPost2(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := treasury_financialaccount.Update("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransfersCancelPost(t *testing.T) {
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, err := treasury_inboundtransfer.Cancel("ibt_xxxxxxxxxxxxx", params)
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

func TestTreasuryInboundTransfersGet2(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{}
	result, err := treasury_inboundtransfer.Get("ibt_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundPaymentsCancelPost(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, err := treasury_outboundpayment.Cancel("bot_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundPaymentsGet2(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{}
	result, err := treasury_outboundpayment.Get("bot_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundTransfersCancelPost(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, err := treasury_outboundtransfer.Cancel("obt_xxxxxxxxxxxxx", params)
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

func TestTreasuryOutboundTransfersGet2(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{}
	result, err := treasury_outboundtransfer.Get("obt_xxxxxxxxxxxxx", params)
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

func TestTreasuryReceivedCreditsGet(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_receivedcredit.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedCreditsGet2(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditParams{}
	result, err := treasury_receivedcredit.Get("rc_xxxxxxxxxxxxx", params)
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

func TestTreasuryReceivedDebitsGet2(t *testing.T) {
	params := &stripe.TreasuryReceivedDebitParams{}
	result, err := treasury_receiveddebit.Get("rd_xxxxxxxxxxxxx", params)
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

func TestTreasuryTransactionEntriesGet2(t *testing.T) {
	params := &stripe.TreasuryTransactionEntryParams{}
	result, err := treasury_transactionentry.Get("trxne_xxxxxxxxxxxxx", params)
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

func TestTreasuryTransactionsGet2(t *testing.T) {
	params := &stripe.TreasuryTransactionParams{}
	result, err := treasury_transaction.Get("trxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointsDelete(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Del("we_xxxxxxxxxxxxx", params)
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

func TestWebhookEndpointsGet2(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Get("we_xxxxxxxxxxxxx", params)
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

func TestWebhookEndpointsPost2(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, err := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
