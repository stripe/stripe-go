// File generated from our OpenAPI spec
package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v72"
	account "github.com/stripe/stripe-go/v72/account"
	accountlink "github.com/stripe/stripe-go/v72/accountlink"
	apps_secret "github.com/stripe/stripe-go/v72/apps/secret"
	balancetransaction "github.com/stripe/stripe-go/v72/balancetransaction"
	billingportal_configuration "github.com/stripe/stripe-go/v72/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v72/billingportal/session"
	capability "github.com/stripe/stripe-go/v72/capability"
	cashbalance "github.com/stripe/stripe-go/v72/cashbalance"
	charge "github.com/stripe/stripe-go/v72/charge"
	checkout_session "github.com/stripe/stripe-go/v72/checkout/session"
	countryspec "github.com/stripe/stripe-go/v72/countryspec"
	coupon "github.com/stripe/stripe-go/v72/coupon"
	customer "github.com/stripe/stripe-go/v72/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v72/customerbalancetransaction"
	dispute "github.com/stripe/stripe-go/v72/dispute"
	event "github.com/stripe/stripe-go/v72/event"
	financialconnections_account "github.com/stripe/stripe-go/v72/financialconnections/account"
	financialconnections_session "github.com/stripe/stripe-go/v72/financialconnections/session"
	identity_verificationreport "github.com/stripe/stripe-go/v72/identity/verificationreport"
	identity_verificationsession "github.com/stripe/stripe-go/v72/identity/verificationsession"
	invoice "github.com/stripe/stripe-go/v72/invoice"
	invoiceitem "github.com/stripe/stripe-go/v72/invoiceitem"
	issuing_authorization "github.com/stripe/stripe-go/v72/issuing/authorization"
	issuing_card "github.com/stripe/stripe-go/v72/issuing/card"
	issuing_cardholder "github.com/stripe/stripe-go/v72/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v72/issuing/dispute"
	issuing_transaction "github.com/stripe/stripe-go/v72/issuing/transaction"
	mandate "github.com/stripe/stripe-go/v72/mandate"
	paymentintent "github.com/stripe/stripe-go/v72/paymentintent"
	paymentlink "github.com/stripe/stripe-go/v72/paymentlink"
	paymentmethod "github.com/stripe/stripe-go/v72/paymentmethod"
	payout "github.com/stripe/stripe-go/v72/payout"
	person "github.com/stripe/stripe-go/v72/person"
	plan "github.com/stripe/stripe-go/v72/plan"
	price "github.com/stripe/stripe-go/v72/price"
	product "github.com/stripe/stripe-go/v72/product"
	promotioncode "github.com/stripe/stripe-go/v72/promotioncode"
	quote "github.com/stripe/stripe-go/v72/quote"
	radar_earlyfraudwarning "github.com/stripe/stripe-go/v72/radar/earlyfraudwarning"
	refund "github.com/stripe/stripe-go/v72/refund"
	reversal "github.com/stripe/stripe-go/v72/reversal"
	review "github.com/stripe/stripe-go/v72/review"
	setupattempt "github.com/stripe/stripe-go/v72/setupattempt"
	setupintent "github.com/stripe/stripe-go/v72/setupintent"
	shippingrate "github.com/stripe/stripe-go/v72/shippingrate"
	sigma_scheduledqueryrun "github.com/stripe/stripe-go/v72/sigma/scheduledqueryrun"
	sku "github.com/stripe/stripe-go/v72/sku"
	source "github.com/stripe/stripe-go/v72/source"
	taxcode "github.com/stripe/stripe-go/v72/taxcode"
	taxid "github.com/stripe/stripe-go/v72/taxid"
	taxrate "github.com/stripe/stripe-go/v72/taxrate"
	terminal_configuration "github.com/stripe/stripe-go/v72/terminal/configuration"
	terminal_connectiontoken "github.com/stripe/stripe-go/v72/terminal/connectiontoken"
	terminal_location "github.com/stripe/stripe-go/v72/terminal/location"
	terminal_reader "github.com/stripe/stripe-go/v72/terminal/reader"
	testhelpers_customer "github.com/stripe/stripe-go/v72/testhelpers/customer"
	testhelpers_issuing_card "github.com/stripe/stripe-go/v72/testhelpers/issuing/card"
	testhelpers_refund "github.com/stripe/stripe-go/v72/testhelpers/refund"
	testhelpers_testclock "github.com/stripe/stripe-go/v72/testhelpers/testclock"
	testhelpers_treasury_inboundtransfer "github.com/stripe/stripe-go/v72/testhelpers/treasury/inboundtransfer"
	testhelpers_treasury_outboundtransfer "github.com/stripe/stripe-go/v72/testhelpers/treasury/outboundtransfer"
	testhelpers_treasury_receivedcredit "github.com/stripe/stripe-go/v72/testhelpers/treasury/receivedcredit"
	testhelpers_treasury_receiveddebit "github.com/stripe/stripe-go/v72/testhelpers/treasury/receiveddebit"
	_ "github.com/stripe/stripe-go/v72/testing"
	topup "github.com/stripe/stripe-go/v72/topup"
	transfer "github.com/stripe/stripe-go/v72/transfer"
	treasury_creditreversal "github.com/stripe/stripe-go/v72/treasury/creditreversal"
	treasury_debitreversal "github.com/stripe/stripe-go/v72/treasury/debitreversal"
	treasury_financialaccount "github.com/stripe/stripe-go/v72/treasury/financialaccount"
	treasury_inboundtransfer "github.com/stripe/stripe-go/v72/treasury/inboundtransfer"
	treasury_outboundpayment "github.com/stripe/stripe-go/v72/treasury/outboundpayment"
	treasury_outboundtransfer "github.com/stripe/stripe-go/v72/treasury/outboundtransfer"
	treasury_receivedcredit "github.com/stripe/stripe-go/v72/treasury/receivedcredit"
	treasury_receiveddebit "github.com/stripe/stripe-go/v72/treasury/receiveddebit"
	treasury_transaction "github.com/stripe/stripe-go/v72/treasury/transaction"
	treasury_transactionentry "github.com/stripe/stripe-go/v72/treasury/transactionentry"
	usagerecord "github.com/stripe/stripe-go/v72/usagerecord"
	usagerecordsummary "github.com/stripe/stripe-go/v72/usagerecordsummary"
	webhookendpoint "github.com/stripe/stripe-go/v72/webhookendpoint"
)

func TestAppsSecretList(t *testing.T) {
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

func TestAppsSecretCreate(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("sec_123"),
		Payload: stripe.String("very secret string"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, _ := apps_secret.New(params)
	assert.NotNil(t, result)
}

func TestAppsSecretDeleteWhere(t *testing.T) {
	params := &stripe.AppsSecretDeleteWhereParams{
		Name: stripe.String("my-api-key"),
		Scope: &stripe.AppsSecretDeleteWhereScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, _ := apps_secret.DeleteWhere(params)
	assert.NotNil(t, result)
}

func TestAppsSecretFind(t *testing.T) {
	params := &stripe.AppsSecretFindParams{
		Name: stripe.String("sec_123"),
		Scope: &stripe.AppsSecretFindScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, _ := apps_secret.Find(params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionExpire(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, _ := checkout_session.Expire("sess_xyz", params)
	assert.NotNil(t, result)
}

func TestCashBalanceRetrieve(t *testing.T) {
	params := &stripe.CashBalanceParams{Customer: stripe.String("cus_123")}
	result, _ := cashbalance.Get(params)
	assert.NotNil(t, result)
}

func TestCashBalanceUpdate(t *testing.T) {
	params := &stripe.CashBalanceParams{
		Settings: &stripe.CashBalanceSettingsParams{
			ReconciliationMode: stripe.String(string(stripe.CashBalanceSettingsReconciliationModeManual)),
		},
		Customer: stripe.String("cus_123"),
	}
	result, _ := cashbalance.Update(params)
	assert.NotNil(t, result)
}

func TestCustomerCreateFundingInstructions(t *testing.T) {
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		FundingType: stripe.String("bank_transfer"),
	}
	result, _ := customer.CreateFundingInstructions("cus_123", params)
	assert.NotNil(t, result)
}

func TestCustomerListPaymentMethods(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xyz"),
	}
	result := customer.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountList(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListParams{}
	result := financialconnections_account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountRetrieve(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, _ := financialconnections_account.GetByID("fca_xyz", params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsAccountDisconnect(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, _ := financialconnections_account.Disconnect("fca_xyz", params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsAccountListOwners(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xyz"),
		Account:   stripe.String("fca_xyz"),
	}
	result := financialconnections_account.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountRefresh(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountRefreshParams{
		Features: []*string{stripe.String("balance")},
	}
	result, _ := financialconnections_account.Refresh("fca_xyz", params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsSessionCreate(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{
		AccountHolder: &stripe.FinancialConnectionsSessionAccountHolderParams{
			Type:     stripe.String(string(stripe.FinancialConnectionsSessionAccountHolderTypeCustomer)),
			Customer: stripe.String("cus_123"),
		},
		Permissions: []*string{
			stripe.String(string(stripe.FinancialConnectionsSessionPermissionBalances)),
		},
	}
	result, _ := financialconnections_session.New(params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsSessionRetrieve(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, _ := financialconnections_session.Get("fcsess_xyz", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCreate(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, _ := paymentintent.New(params)
	assert.NotNil(t, result)
}

func TestPaymentIntentVerifyMicrodeposits(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, _ := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentLinkCreate(t *testing.T) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			&stripe.PaymentLinkLineItemParams{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, _ := paymentlink.New(params)
	assert.NotNil(t, result)
}

func TestPaymentLinkRetrieve(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, _ := paymentlink.Get("pl_xyz", params)
	assert.NotNil(t, result)
}

func TestPaymentLinkListLineItems(t *testing.T) {
	params := &stripe.PaymentLinkListLineItemsParams{
		PaymentLink: stripe.String("pl_xyz"),
	}
	result := paymentlink.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPriceCreate(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		CurrencyOptions: map[string]*stripe.PriceCurrencyOptionsParams{
			"uah": &stripe.PriceCurrencyOptionsParams{UnitAmount: stripe.Int64(5000)},
			"eur": &stripe.PriceCurrencyOptionsParams{UnitAmount: stripe.Int64(1800)},
		},
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := price.New(params)
	assert.NotNil(t, result)
}

func TestSetupAttemptList(t *testing.T) {
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("si_xyz"),
	}
	params.Limit = stripe.Int64(3)
	result := setupattempt.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentVerifyMicrodeposits(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, _ := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestShippingRateList(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	result := shippingrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRateCreate(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, _ := shippingrate.New(params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationList(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationCreate(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.New(params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationDelete(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Del("uc_123", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationRetrieve(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Get("uc_123", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationUpdate(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		Tipping: &stripe.TerminalConfigurationTippingParams{
			USD: &stripe.TerminalConfigurationTippingUSDParams{
				FixedAmounts: []*int64{stripe.Int64(10)},
			},
		},
	}
	result, _ := terminal_configuration.Update("uc_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersCustomerFundCashBalance(t *testing.T) {
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	result, _ := testhelpers_customer.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersIssuingCardDeliverCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, _ := testhelpers_issuing_card.DeliverCard("card_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersIssuingCardFailCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, _ := testhelpers_issuing_card.FailCard("card_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersIssuingCardReturnCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, _ := testhelpers_issuing_card.ReturnCard("card_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersIssuingCardShipCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, _ := testhelpers_issuing_card.ShipCard("card_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersRefundExpire(t *testing.T) {
	params := &stripe.TestHelpersRefundExpireParams{}
	result, _ := testhelpers_refund.Expire("re_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockList(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClockCreate(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(123),
		Name:       stripe.String("cogsworth"),
	}
	result, _ := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockDelete(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Del("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockRetrieve(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Get("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockAdvance(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, _ := testhelpers_testclock.Advance("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryInboundTransferFail(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(string(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed)),
		},
	}
	result, _ := testhelpers_treasury_inboundtransfer.Fail("ibt_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryInboundTransferReturnInboundTransfer(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, _ := testhelpers_treasury_inboundtransfer.ReturnInboundTransfer(
		"ibt_123",
		params,
	)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryInboundTransferSucceed(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, _ := testhelpers_treasury_inboundtransfer.Succeed("ibt_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryOutboundTransferFail(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, _ := testhelpers_treasury_outboundtransfer.Fail("obt_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryOutboundTransferPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, _ := testhelpers_treasury_outboundtransfer.Post("obt_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryOutboundTransferReturnOutboundTransfer(
	t *testing.T,
) {
	params := &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams{
		ReturnedDetails: &stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams{
			Code: stripe.String(string(stripe.TreasuryOutboundTransferReturnedDetailsCodeAccountClosed)),
		},
	}
	result, _ := testhelpers_treasury_outboundtransfer.ReturnOutboundTransfer(
		"obt_123",
		params,
	)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryReceivedCreditCreate(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedCreditParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String(string(stripe.TreasuryReceivedCreditNetworkAch)),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := testhelpers_treasury_receivedcredit.New(params)
	assert.NotNil(t, result)
}

func TestTestHelpersTreasuryReceivedDebitCreate(t *testing.T) {
	params := &stripe.TestHelpersTreasuryReceivedDebitParams{
		FinancialAccount: stripe.String("fa_123"),
		Network:          stripe.String("ach"),
		Amount:           stripe.Int64(1234),
		Currency:         stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := testhelpers_treasury_receiveddebit.New(params)
	assert.NotNil(t, result)
}

func TestAccountLinkCreate(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, _ := accountlink.New(params)
	assert.NotNil(t, result)
}

func TestAccountList(t *testing.T) {
	params := &stripe.AccountListParams{}
	params.Limit = stripe.Int64(3)
	result := account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestAccountCreate(t *testing.T) {
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
	result, _ := account.New(params)
	assert.NotNil(t, result)
}

func TestAccountDelete(t *testing.T) {
	params := &stripe.AccountParams{}
	result, _ := account.Del("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountRetrieve(t *testing.T) {
	params := &stripe.AccountParams{}
	result, _ := account.GetByID("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountUpdate(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountReject(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, _ := account.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCapabilityList(t *testing.T) {
	params := &stripe.CapabilityListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result := capability.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCapabilityRetrieve(t *testing.T) {
	params := &stripe.CapabilityParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, _ := capability.Get("card_payments", params)
	assert.NotNil(t, result)
}

func TestCapabilityUpdate(t *testing.T) {
	params := &stripe.CapabilityParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, _ := capability.Update("card_payments", params)
	assert.NotNil(t, result)
}

func TestPersonList(t *testing.T) {
	params := &stripe.PersonListParams{
		Account: stripe.String("acct_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := person.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPersonCreate(t *testing.T) {
	params := &stripe.PersonParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
	}
	result, _ := person.New(params)
	assert.NotNil(t, result)
}

func TestPersonUpdate(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, _ := person.Update("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAppsSecretCreate2(t *testing.T) {
	params := &stripe.AppsSecretParams{
		Name:    stripe.String("my-api-key"),
		Payload: stripe.String("secret_key_xxxxxx"),
		Scope: &stripe.AppsSecretScopeParams{
			Type: stripe.String(string(stripe.AppsSecretScopeTypeAccount)),
		},
	}
	result, _ := apps_secret.New(params)
	assert.NotNil(t, result)
}

func TestBalanceTransactionList(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.BalanceTransactionParams{}
	result, _ := balancetransaction.Get("txn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationList(t *testing.T) {
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := billingportal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBillingPortalConfigurationCreate(t *testing.T) {
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
	result, _ := billingportal_configuration.New(params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationRetrieve(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{}
	result, _ := billingportal_configuration.Get("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationUpdate(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, _ := billingportal_configuration.Update("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBillingPortalSessionCreate(t *testing.T) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, _ := billingportal_session.New(params)
	assert.NotNil(t, result)
}

func TestChargeList(t *testing.T) {
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(3)
	result := charge.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestChargeCreate(t *testing.T) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Source:      &stripe.SourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs)"),
	}
	result, _ := charge.New(params)
	assert.NotNil(t, result)
}

func TestChargeRetrieve(t *testing.T) {
	params := &stripe.ChargeParams{}
	result, _ := charge.Get("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeUpdate(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := charge.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeCapture(t *testing.T) {
	params := &stripe.CaptureParams{}
	result, _ := charge.Capture("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeSearch(t *testing.T) {
	params := &stripe.ChargeSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "amount>999 AND metadata['order_id']:'6735'",
		},
	}
	result := charge.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionList(t *testing.T) {
	params := &stripe.CheckoutSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := checkout_session.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCheckoutSessionRetrieve(t *testing.T) {
	params := &stripe.CheckoutSessionParams{}
	result, _ := checkout_session.Get("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionExpire2(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, _ := checkout_session.Expire("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCountrySpecList(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Limit = stripe.Int64(3)
	result := countryspec.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCountrySpecRetrieve(t *testing.T) {
	params := &stripe.CountrySpecParams{}
	result, _ := countryspec.Get("US", params)
	assert.NotNil(t, result)
}

func TestCouponList(t *testing.T) {
	params := &stripe.CouponListParams{}
	params.Limit = stripe.Int64(3)
	result := coupon.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCouponCreate(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff:       stripe.Float64(25.5),
		Duration:         stripe.String(string(stripe.CouponDurationRepeating)),
		DurationInMonths: stripe.Int64(3),
	}
	result, _ := coupon.New(params)
	assert.NotNil(t, result)
}

func TestCouponDelete(t *testing.T) {
	params := &stripe.CouponParams{}
	result, _ := coupon.Del("Z4OV52SU", params)
	assert.NotNil(t, result)
}

func TestCouponRetrieve(t *testing.T) {
	params := &stripe.CouponParams{}
	result, _ := coupon.Get("Z4OV52SU", params)
	assert.NotNil(t, result)
}

func TestCouponUpdate(t *testing.T) {
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := coupon.Update("Z4OV52SU", params)
	assert.NotNil(t, result)
}

func TestCustomerList(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := customer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomerList2(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Limit = stripe.Int64(3)
	result := customer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomerCreate(t *testing.T) {
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	result, _ := customer.New(params)
	assert.NotNil(t, result)
}

func TestCustomerDelete(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, _ := customer.Del("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerRetrieve(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, _ := customer.Get("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerUpdate(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := customer.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionList(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := customerbalancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomerBalanceTransactionCreate(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := customerbalancetransaction.New(params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, _ := customerbalancetransaction.Get("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionUpdate(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, _ := customerbalancetransaction.Update("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerListPaymentMethods2(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type:     stripe.String("card"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result := customer.ListPaymentMethods(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxIDList(t *testing.T) {
	params := &stripe.TaxIDListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := taxid.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxIDCreate(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:  stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value: stripe.String("DE123456789"),
	}
	result, _ := taxid.New(params)
	assert.NotNil(t, result)
}

func TestTaxIDDelete(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := taxid.Del("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxIDRetrieve(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := taxid.Get("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerSearch(t *testing.T) {
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := customer.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCustomerSearch2(t *testing.T) {
	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "name:'fakename' AND metadata['foo']:'bar'",
		},
	}
	result := customer.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestDisputeList(t *testing.T) {
	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := dispute.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestDisputeRetrieve(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, _ := dispute.Get("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeUpdate(t *testing.T) {
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := dispute.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeClose(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, _ := dispute.Close("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestEventList(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Limit = stripe.Int64(3)
	result := event.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestEventRetrieve(t *testing.T) {
	params := &stripe.EventParams{}
	result, _ := event.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsAccountList2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListParams{
		AccountHolder: &stripe.FinancialConnectionsAccountListAccountHolderParams{
			Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		},
	}
	result := financialconnections_account.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsAccountRetrieve2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountParams{}
	result, _ := financialconnections_account.GetByID(
		"fca_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsAccountListOwners2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountListOwnersParams{
		Ownership: stripe.String("fcaowns_xxxxxxxxxxxxx"),
		Account:   stripe.String("fca_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := financialconnections_account.ListOwners(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestFinancialConnectionsSessionCreate2(t *testing.T) {
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
	result, _ := financialconnections_session.New(params)
	assert.NotNil(t, result)
}

func TestFinancialConnectionsSessionRetrieve2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, _ := financialconnections_session.Get("fcsess_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationReportList(t *testing.T) {
	params := &stripe.IdentityVerificationReportListParams{}
	params.Limit = stripe.Int64(3)
	result := identity_verificationreport.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationReportRetrieve(t *testing.T) {
	params := &stripe.IdentityVerificationReportParams{}
	result, _ := identity_verificationreport.Get("vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationSessionList(t *testing.T) {
	params := &stripe.IdentityVerificationSessionListParams{}
	params.Limit = stripe.Int64(3)
	result := identity_verificationsession.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIdentityVerificationSessionCreate(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeDocument)),
	}
	result, _ := identity_verificationsession.New(params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationSessionRetrieve(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{}
	result, _ := identity_verificationsession.Get("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationSessionUpdate(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{
		Type: stripe.String(string(stripe.IdentityVerificationSessionTypeIDNumber)),
	}
	result, _ := identity_verificationsession.Update("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationSessionCancel(t *testing.T) {
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, _ := identity_verificationsession.Cancel("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIdentityVerificationSessionRedact(t *testing.T) {
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, _ := identity_verificationsession.Redact("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemList(t *testing.T) {
	params := &stripe.InvoiceItemListParams{}
	params.Limit = stripe.Int64(3)
	result := invoiceitem.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoiceItemCreate(t *testing.T) {
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Price:    stripe.String("price_xxxxxxxxxxxxx"),
	}
	result, _ := invoiceitem.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceItemDelete(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, _ := invoiceitem.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemRetrieve(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, _ := invoiceitem.Get("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemUpdate(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoiceitem.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceList(t *testing.T) {
	params := &stripe.InvoiceListParams{}
	params.Limit = stripe.Int64(3)
	result := invoice.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestInvoiceCreate(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := invoice.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceDelete(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, _ := invoice.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceRetrieve(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, _ := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceUpdate(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoice.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceFinalizeInvoice(t *testing.T) {
	params := &stripe.InvoiceFinalizeParams{}
	result, _ := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceMarkUncollectible(t *testing.T) {
	params := &stripe.InvoiceMarkUncollectibleParams{}
	result, _ := invoice.MarkUncollectible("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoicePay(t *testing.T) {
	params := &stripe.InvoicePayParams{}
	result, _ := invoice.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceSendInvoice(t *testing.T) {
	params := &stripe.InvoiceSendParams{}
	result, _ := invoice.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceVoidInvoice(t *testing.T) {
	params := &stripe.InvoiceVoidParams{}
	result, _ := invoice.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceSearch(t *testing.T) {
	params := &stripe.InvoiceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "total>999 AND metadata['order_id']:'6735'",
		},
	}
	result := invoice.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingAuthorizationList(t *testing.T) {
	params := &stripe.IssuingAuthorizationListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_authorization.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingAuthorizationRetrieve(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	result, _ := issuing_authorization.Get("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_authorization.Update("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationApprove(t *testing.T) {
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, _ := issuing_authorization.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationDecline(t *testing.T) {
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, _ := issuing_authorization.Decline("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderList(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_cardholder.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardholderCreate(t *testing.T) {
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
	result, _ := issuing_cardholder.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderRetrieve(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	result, _ := issuing_cardholder.Get("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderUpdate(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_cardholder.Update("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardList(t *testing.T) {
	params := &stripe.IssuingCardListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_card.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingCardCreate(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String(string(stripe.IssuingCardTypeVirtual)),
	}
	result, _ := issuing_card.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardRetrieve(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	result, _ := issuing_card.Get("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardUpdate(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_card.Update("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeList(t *testing.T) {
	params := &stripe.IssuingDisputeListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_dispute.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingDisputeCreate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonFraudulent)),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, _ := issuing_dispute.New(params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeRetrieve(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	result, _ := issuing_dispute.Get("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeSubmit(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	result, _ := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionList(t *testing.T) {
	params := &stripe.IssuingTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := issuing_transaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestIssuingTransactionRetrieve(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	result, _ := issuing_transaction.Get("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionUpdate(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_transaction.Update("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestMandateRetrieve(t *testing.T) {
	params := &stripe.MandateParams{}
	result, _ := mandate.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentList(t *testing.T) {
	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := paymentintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentCreate2(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(2000),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := paymentintent.New(params)
	assert.NotNil(t, result)
}

func TestPaymentIntentRetrieve(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	result, _ := paymentintent.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentUpdate(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentApplyCustomerBalance(t *testing.T) {
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, _ := paymentintent.ApplyCustomerBalance("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCancel(t *testing.T) {
	params := &stripe.PaymentIntentCancelParams{}
	result, _ := paymentintent.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCapture(t *testing.T) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, _ := paymentintent.Capture("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentConfirm(t *testing.T) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := paymentintent.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentIncrementAuthorization(t *testing.T) {
	params := &stripe.PaymentIntentIncrementAuthorizationParams{
		Amount: stripe.Int64(2099),
	}
	result, _ := paymentintent.IncrementAuthorization("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentSearch(t *testing.T) {
	params := &stripe.PaymentIntentSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'succeeded' AND metadata['order_id']:'6735'",
		},
	}
	result := paymentintent.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinkList(t *testing.T) {
	params := &stripe.PaymentLinkListParams{}
	params.Limit = stripe.Int64(3)
	result := paymentlink.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentLinkCreate2(t *testing.T) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			&stripe.PaymentLinkLineItemParams{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, _ := paymentlink.New(params)
	assert.NotNil(t, result)
}

func TestPaymentLinkRetrieve2(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, _ := paymentlink.Get("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentLinkUpdate(t *testing.T) {
	params := &stripe.PaymentLinkParams{Active: stripe.Bool(false)}
	result, _ := paymentlink.Update("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodList(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentMethodRetrieve(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	result, _ := paymentmethod.Get("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodUpdate(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodAttach(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, _ := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodDetach(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	result, _ := paymentmethod.Detach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutList(t *testing.T) {
	params := &stripe.PayoutListParams{}
	params.Limit = stripe.Int64(3)
	result := payout.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPayoutCreate(t *testing.T) {
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := payout.New(params)
	assert.NotNil(t, result)
}

func TestPayoutRetrieve(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, _ := payout.Get("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutUpdate(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutCancel(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, _ := payout.Cancel("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutReverse(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	result, _ := payout.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanList(t *testing.T) {
	params := &stripe.PlanListParams{}
	params.Limit = stripe.Int64(3)
	result := plan.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPlanCreate(t *testing.T) {
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Interval: stripe.String(string(stripe.PlanIntervalMonth)),
		Product:  &stripe.PlanProductParams{ID: stripe.String("prod_xxxxxxxxxxxxx")},
	}
	result, _ := plan.New(params)
	assert.NotNil(t, result)
}

func TestPlanDelete(t *testing.T) {
	params := &stripe.PlanParams{}
	result, _ := plan.Del("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanRetrieve(t *testing.T) {
	params := &stripe.PlanParams{}
	result, _ := plan.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanUpdate(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := plan.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceList(t *testing.T) {
	params := &stripe.PriceListParams{}
	params.Limit = stripe.Int64(3)
	result := price.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPriceCreate2(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := price.New(params)
	assert.NotNil(t, result)
}

func TestPriceRetrieve(t *testing.T) {
	params := &stripe.PriceParams{}
	result, _ := price.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceUpdate(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := price.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceSearch(t *testing.T) {
	params := &stripe.PriceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := price.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductList(t *testing.T) {
	params := &stripe.ProductListParams{}
	params.Limit = stripe.Int64(3)
	result := product.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestProductCreate(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, _ := product.New(params)
	assert.NotNil(t, result)
}

func TestProductDelete(t *testing.T) {
	params := &stripe.ProductParams{}
	result, _ := product.Del("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductRetrieve(t *testing.T) {
	params := &stripe.ProductParams{}
	result, _ := product.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductUpdate(t *testing.T) {
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := product.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductSearch(t *testing.T) {
	params := &stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "active:'true' AND metadata['order_id']:'6735'",
		},
	}
	result := product.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodeList(t *testing.T) {
	params := &stripe.PromotionCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := promotioncode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPromotionCodeCreate(t *testing.T) {
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("Z4OV52SU")}
	result, _ := promotioncode.New(params)
	assert.NotNil(t, result)
}

func TestPromotionCodeRetrieve(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	result, _ := promotioncode.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPromotionCodeUpdate(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestQuoteList(t *testing.T) {
	params := &stripe.QuoteListParams{}
	params.Limit = stripe.Int64(3)
	result := quote.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestQuoteCreate(t *testing.T) {
	params := &stripe.QuoteParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		LineItems: []*stripe.QuoteLineItemParams{
			&stripe.QuoteLineItemParams{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(2),
			},
		},
	}
	result, _ := quote.New(params)
	assert.NotNil(t, result)
}

func TestQuoteRetrieve(t *testing.T) {
	params := &stripe.QuoteParams{}
	result, _ := quote.Get("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestQuoteUpdate(t *testing.T) {
	params := &stripe.QuoteParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := quote.Update("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestQuoteAccept(t *testing.T) {
	params := &stripe.QuoteAcceptParams{}
	result, _ := quote.Accept("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestQuoteCancel(t *testing.T) {
	params := &stripe.QuoteCancelParams{}
	result, _ := quote.Cancel("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestQuoteFinalizeQuote(t *testing.T) {
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, _ := quote.FinalizeQuote("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningList(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Limit = stripe.Int64(3)
	result := radar_earlyfraudwarning.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRadarEarlyFraudWarningRetrieve(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningParams{}
	result, _ := radar_earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundList(t *testing.T) {
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(3)
	result := refund.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestRefundCreate(t *testing.T) {
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, _ := refund.New(params)
	assert.NotNil(t, result)
}

func TestRefundRetrieve(t *testing.T) {
	params := &stripe.RefundParams{}
	result, _ := refund.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundUpdate(t *testing.T) {
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := refund.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundCancel(t *testing.T) {
	params := &stripe.RefundCancelParams{}
	result, _ := refund.Cancel("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReviewList(t *testing.T) {
	params := &stripe.ReviewListParams{}
	params.Limit = stripe.Int64(3)
	result := review.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReviewRetrieve(t *testing.T) {
	params := &stripe.ReviewParams{}
	result, _ := review.Get("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReviewApprove(t *testing.T) {
	params := &stripe.ReviewApproveParams{}
	result, _ := review.Approve("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentList(t *testing.T) {
	params := &stripe.SetupIntentListParams{}
	params.Limit = stripe.Int64(3)
	result := setupintent.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSetupIntentCreate(t *testing.T) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := setupintent.New(params)
	assert.NotNil(t, result)
}

func TestSetupIntentRetrieve(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	result, _ := setupintent.Get("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentUpdate(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, _ := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentCancel(t *testing.T) {
	params := &stripe.SetupIntentCancelParams{}
	result, _ := setupintent.Cancel("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentConfirm(t *testing.T) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := setupintent.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestShippingRateList2(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	params.Limit = stripe.Int64(3)
	result := shippingrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestShippingRateCreate2(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Ground shipping"),
		Type:        stripe.String("fixed_amount"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Amount:   stripe.Int64(500),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		},
	}
	result, _ := shippingrate.New(params)
	assert.NotNil(t, result)
}

func TestShippingRateRetrieve(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	result, _ := shippingrate.Get("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestShippingRateUpdate(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := shippingrate.Update("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunList(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Limit = stripe.Int64(3)
	result := sigma_scheduledqueryrun.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSigmaScheduledQueryRunRetrieve(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunParams{}
	result, _ := sigma_scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKUList(t *testing.T) {
	params := &stripe.SKUListParams{}
	params.Limit = stripe.Int64(3)
	result := sku.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSKUCreate(t *testing.T) {
	params := &stripe.SKUParams{
		Attributes: map[string]string{"size": "Medium", "gender": "Unisex"},
		Price:      stripe.Int64(1500),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Inventory: &stripe.InventoryParams{
			Type:     stripe.String(string(stripe.SKUInventoryTypeFinite)),
			Quantity: stripe.Int64(500),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := sku.New(params)
	assert.NotNil(t, result)
}

func TestSKUDelete(t *testing.T) {
	params := &stripe.SKUParams{}
	result, _ := sku.Del("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKURetrieve(t *testing.T) {
	params := &stripe.SKUParams{}
	result, _ := sku.Get("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKUUpdate(t *testing.T) {
	params := &stripe.SKUParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := sku.Update("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSourceRetrieve(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	result, _ := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSourceRetrieve2(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	result, _ := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := source.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestUsageRecordSummaryList(t *testing.T) {
	params := &stripe.UsageRecordSummaryListParams{
		SubscriptionItem: stripe.String("si_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := usagerecordsummary.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestUsageRecordCreate(t *testing.T) {
	params := &stripe.UsageRecordParams{
		Quantity:  stripe.Int64(100),
		Timestamp: stripe.Int64(1571252444),
	}
	result, _ := usagerecord.New(params)
	assert.NotNil(t, result)
}

func TestTaxCodeList(t *testing.T) {
	params := &stripe.TaxCodeListParams{}
	params.Limit = stripe.Int64(3)
	result := taxcode.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCodeRetrieve(t *testing.T) {
	params := &stripe.TaxCodeParams{}
	result, _ := taxcode.Get("txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxRateList(t *testing.T) {
	params := &stripe.TaxRateListParams{}
	params.Limit = stripe.Int64(3)
	result := taxrate.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxRateCreate(t *testing.T) {
	params := &stripe.TaxRateParams{
		DisplayName:  stripe.String("VAT"),
		Description:  stripe.String("VAT Germany"),
		Jurisdiction: stripe.String("DE"),
		Percentage:   stripe.Float64(16),
		Inclusive:    stripe.Bool(false),
	}
	result, _ := taxrate.New(params)
	assert.NotNil(t, result)
}

func TestTaxRateRetrieve(t *testing.T) {
	params := &stripe.TaxRateParams{}
	result, _ := taxrate.Get("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxRateUpdate(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, _ := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationList2(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationCreate2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, _ := terminal_configuration.New(params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationDelete2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Del("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationRetrieve2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Get("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationUpdate2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, _ := terminal_configuration.Update("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalConnectionTokenCreate(t *testing.T) {
	params := &stripe.TerminalConnectionTokenParams{}
	result, _ := terminal_connectiontoken.New(params)
	assert.NotNil(t, result)
}

func TestTerminalLocationList(t *testing.T) {
	params := &stripe.TerminalLocationListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_location.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalLocationCreate(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
		Address: &stripe.AccountAddressParams{
			Line1:      stripe.String("1234 Main Street"),
			City:       stripe.String("San Francisco"),
			Country:    stripe.String("US"),
			PostalCode: stripe.String("94111"),
		},
	}
	result, _ := terminal_location.New(params)
	assert.NotNil(t, result)
}

func TestTerminalLocationDelete(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, _ := terminal_location.Del("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalLocationRetrieve(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, _ := terminal_location.Get("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalLocationUpdate(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, _ := terminal_location.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderList(t *testing.T) {
	params := &stripe.TerminalReaderListParams{}
	params.Limit = stripe.Int64(3)
	result := terminal_reader.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalReaderCreate(t *testing.T) {
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, _ := terminal_reader.New(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderDelete(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, _ := terminal_reader.Del("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderUpdate(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, _ := terminal_reader.Update("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderCancelAction(t *testing.T) {
	params := &stripe.TerminalReaderCancelActionParams{}
	result, _ := terminal_reader.CancelAction("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderProcessPaymentIntent(t *testing.T) {
	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String("pi_xxxxxxxxxxxxx"),
	}
	result, _ := terminal_reader.ProcessPaymentIntent(
		"tmr_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockList2(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	params.Limit = stripe.Int64(3)
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTestHelpersTestClockCreate2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(1577836800),
	}
	result, _ := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockDelete2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Del("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockRetrieve2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Get("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockAdvance2(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1652390605),
	}
	result, _ := testhelpers_testclock.Advance("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTopupList(t *testing.T) {
	params := &stripe.TopupListParams{}
	params.Limit = stripe.Int64(3)
	result := topup.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTopupCreate(t *testing.T) {
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, _ := topup.New(params)
	assert.NotNil(t, result)
}

func TestTopupRetrieve(t *testing.T) {
	params := &stripe.TopupParams{}
	result, _ := topup.Get("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTopupUpdate(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := topup.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferList(t *testing.T) {
	params := &stripe.TransferListParams{}
	params.Limit = stripe.Int64(3)
	result := transfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransferCreate(t *testing.T) {
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, _ := transfer.New(params)
	assert.NotNil(t, result)
}

func TestTransferRetrieve(t *testing.T) {
	params := &stripe.TransferParams{}
	result, _ := transfer.Get("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferUpdate(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := transfer.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReversalList(t *testing.T) {
	params := &stripe.ReversalListParams{
		Transfer: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := reversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestReversalCreate(t *testing.T) {
	params := &stripe.ReversalParams{Amount: stripe.Int64(100)}
	result, _ := reversal.New(params)
	assert.NotNil(t, result)
}

func TestReversalRetrieve(t *testing.T) {
	params := &stripe.ReversalParams{Transfer: stripe.String("tr_xxxxxxxxxxxxx")}
	result, _ := reversal.Get("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReversalUpdate(t *testing.T) {
	params := &stripe.ReversalParams{Transfer: stripe.String("tr_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, _ := reversal.Update("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryCreditReversalList(t *testing.T) {
	params := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_creditreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryCreditReversalCreate(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{
		ReceivedCredit: stripe.String("rc_xxxxxxxxxxxxx"),
	}
	result, _ := treasury_creditreversal.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryCreditReversalRetrieve(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{}
	result, _ := treasury_creditreversal.Get("credrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryDebitReversalList(t *testing.T) {
	params := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_debitreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryDebitReversalCreate(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{
		ReceivedDebit: stripe.String("rd_xxxxxxxxxxxxx"),
	}
	result, _ := treasury_debitreversal.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryDebitReversalRetrieve(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{}
	result, _ := treasury_debitreversal.Get("debrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryFinancialAccountList(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountListParams{}
	params.Limit = stripe.Int64(3)
	result := treasury_financialaccount.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryFinancialAccountCreate(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{
		SupportedCurrencies: []*string{stripe.String("usd")},
		Features:            &stripe.TreasuryFinancialAccountFeaturesParams{},
	}
	result, _ := treasury_financialaccount.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryFinancialAccountRetrieve(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	result, _ := treasury_financialaccount.Get("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryFinancialAccountUpdate(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := treasury_financialaccount.Update("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryFinancialAccountRetrieveFeatures(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, _ := treasury_financialaccount.RetrieveFeatures(
		"fa_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
}

func TestTreasuryFinancialAccountUpdateFeatures(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountUpdateFeaturesParams{
		CardIssuing: &stripe.TreasuryFinancialAccountUpdateFeaturesCardIssuingParams{
			Requested: stripe.Bool(false),
		},
	}
	result, _ := treasury_financialaccount.UpdateFeatures(
		"fa_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
}

func TestTreasuryInboundTransferList(t *testing.T) {
	params := &stripe.TreasuryInboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_inboundtransfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryInboundTransferCreate(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{
		FinancialAccount:    stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:              stripe.Int64(10000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		OriginPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:         stripe.String("InboundTransfer from my bank account"),
	}
	result, _ := treasury_inboundtransfer.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryInboundTransferRetrieve(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{}
	result, _ := treasury_inboundtransfer.Get("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryInboundTransferCancel(t *testing.T) {
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, _ := treasury_inboundtransfer.Cancel("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundPaymentList(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_outboundpayment.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundPaymentCreate(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(10000),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Customer:                 stripe.String("cu_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, _ := treasury_outboundpayment.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundPaymentRetrieve(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{}
	result, _ := treasury_outboundpayment.Get("obp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundPaymentCancel(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, _ := treasury_outboundpayment.Cancel("obp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundTransferList(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_outboundtransfer.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryOutboundTransferCreate(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{
		FinancialAccount:         stripe.String("fa_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Amount:                   stripe.Int64(500),
		Currency:                 stripe.String(string(stripe.CurrencyUSD)),
		Description:              stripe.String("OutboundTransfer to my external bank account"),
	}
	result, _ := treasury_outboundtransfer.New(params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundTransferRetrieve(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{}
	result, _ := treasury_outboundtransfer.Get("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryOutboundTransferCancel(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, _ := treasury_outboundtransfer.Cancel("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryReceivedCreditList(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_receivedcredit.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedCreditRetrieve(t *testing.T) {
	params := &stripe.TreasuryReceivedCreditParams{}
	result, _ := treasury_receivedcredit.Get("rc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryReceivedDebitList(t *testing.T) {
	params := &stripe.TreasuryReceivedDebitListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_receiveddebit.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryReceivedDebitRetrieve(t *testing.T) {
	params := &stripe.TreasuryReceivedDebitParams{}
	result, _ := treasury_receiveddebit.Get("rd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryTransactionEntryList(t *testing.T) {
	params := &stripe.TreasuryTransactionEntryListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_transactionentry.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionEntryRetrieve(t *testing.T) {
	params := &stripe.TreasuryTransactionEntryParams{}
	result, _ := treasury_transactionentry.Get("trxne_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTreasuryTransactionList(t *testing.T) {
	params := &stripe.TreasuryTransactionListParams{
		FinancialAccount: stripe.String("fa_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := treasury_transaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTreasuryTransactionRetrieve(t *testing.T) {
	params := &stripe.TreasuryTransactionParams{}
	result, _ := treasury_transaction.Get("trxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointList(t *testing.T) {
	params := &stripe.WebhookEndpointListParams{}
	params.Limit = stripe.Int64(3)
	result := webhookendpoint.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestWebhookEndpointCreate(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/my/webhook/endpoint"),
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
	}
	result, _ := webhookendpoint.New(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointDelete(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, _ := webhookendpoint.Del("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointRetrieve(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, _ := webhookendpoint.Get("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, _ := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}
