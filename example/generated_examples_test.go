// File generated from our OpenAPI spec
package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v74"
	account "github.com/stripe/stripe-go/v74/account"
	accountlink "github.com/stripe/stripe-go/v74/accountlink"
	apps_secret "github.com/stripe/stripe-go/v74/apps/secret"
	balancetransaction "github.com/stripe/stripe-go/v74/balancetransaction"
	billingportal_configuration "github.com/stripe/stripe-go/v74/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v74/billingportal/session"
	capability "github.com/stripe/stripe-go/v74/capability"
	cashbalance "github.com/stripe/stripe-go/v74/cashbalance"
	charge "github.com/stripe/stripe-go/v74/charge"
	checkout_session "github.com/stripe/stripe-go/v74/checkout/session"
	countryspec "github.com/stripe/stripe-go/v74/countryspec"
	coupon "github.com/stripe/stripe-go/v74/coupon"
	customer "github.com/stripe/stripe-go/v74/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v74/customerbalancetransaction"
	dispute "github.com/stripe/stripe-go/v74/dispute"
	event "github.com/stripe/stripe-go/v74/event"
	financialconnections_account "github.com/stripe/stripe-go/v74/financialconnections/account"
	financialconnections_session "github.com/stripe/stripe-go/v74/financialconnections/session"
	identity_verificationreport "github.com/stripe/stripe-go/v74/identity/verificationreport"
	identity_verificationsession "github.com/stripe/stripe-go/v74/identity/verificationsession"
	invoice "github.com/stripe/stripe-go/v74/invoice"
	invoiceitem "github.com/stripe/stripe-go/v74/invoiceitem"
	issuing_authorization "github.com/stripe/stripe-go/v74/issuing/authorization"
	issuing_card "github.com/stripe/stripe-go/v74/issuing/card"
	issuing_cardholder "github.com/stripe/stripe-go/v74/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v74/issuing/dispute"
	issuing_transaction "github.com/stripe/stripe-go/v74/issuing/transaction"
	mandate "github.com/stripe/stripe-go/v74/mandate"
	paymentintent "github.com/stripe/stripe-go/v74/paymentintent"
	paymentlink "github.com/stripe/stripe-go/v74/paymentlink"
	paymentmethod "github.com/stripe/stripe-go/v74/paymentmethod"
	payout "github.com/stripe/stripe-go/v74/payout"
	person "github.com/stripe/stripe-go/v74/person"
	plan "github.com/stripe/stripe-go/v74/plan"
	price "github.com/stripe/stripe-go/v74/price"
	product "github.com/stripe/stripe-go/v74/product"
	promotioncode "github.com/stripe/stripe-go/v74/promotioncode"
	quote "github.com/stripe/stripe-go/v74/quote"
	radar_earlyfraudwarning "github.com/stripe/stripe-go/v74/radar/earlyfraudwarning"
	refund "github.com/stripe/stripe-go/v74/refund"
	review "github.com/stripe/stripe-go/v74/review"
	setupattempt "github.com/stripe/stripe-go/v74/setupattempt"
	setupintent "github.com/stripe/stripe-go/v74/setupintent"
	shippingrate "github.com/stripe/stripe-go/v74/shippingrate"
	sigma_scheduledqueryrun "github.com/stripe/stripe-go/v74/sigma/scheduledqueryrun"
	source "github.com/stripe/stripe-go/v74/source"
	subscription "github.com/stripe/stripe-go/v74/subscription"
	subscriptionitem "github.com/stripe/stripe-go/v74/subscriptionitem"
	subscriptionschedule "github.com/stripe/stripe-go/v74/subscriptionschedule"
	tax_calculation "github.com/stripe/stripe-go/v74/tax/calculation"
	tax_transaction "github.com/stripe/stripe-go/v74/tax/transaction"
	taxcode "github.com/stripe/stripe-go/v74/taxcode"
	taxid "github.com/stripe/stripe-go/v74/taxid"
	taxrate "github.com/stripe/stripe-go/v74/taxrate"
	terminal_configuration "github.com/stripe/stripe-go/v74/terminal/configuration"
	terminal_connectiontoken "github.com/stripe/stripe-go/v74/terminal/connectiontoken"
	terminal_location "github.com/stripe/stripe-go/v74/terminal/location"
	terminal_reader "github.com/stripe/stripe-go/v74/terminal/reader"
	testhelpers_customer "github.com/stripe/stripe-go/v74/testhelpers/customer"
	testhelpers_issuing_card "github.com/stripe/stripe-go/v74/testhelpers/issuing/card"
	testhelpers_refund "github.com/stripe/stripe-go/v74/testhelpers/refund"
	testhelpers_testclock "github.com/stripe/stripe-go/v74/testhelpers/testclock"
	testhelpers_treasury_inboundtransfer "github.com/stripe/stripe-go/v74/testhelpers/treasury/inboundtransfer"
	testhelpers_treasury_outboundtransfer "github.com/stripe/stripe-go/v74/testhelpers/treasury/outboundtransfer"
	testhelpers_treasury_receivedcredit "github.com/stripe/stripe-go/v74/testhelpers/treasury/receivedcredit"
	testhelpers_treasury_receiveddebit "github.com/stripe/stripe-go/v74/testhelpers/treasury/receiveddebit"
	_ "github.com/stripe/stripe-go/v74/testing"
	token "github.com/stripe/stripe-go/v74/token"
	topup "github.com/stripe/stripe-go/v74/topup"
	transfer "github.com/stripe/stripe-go/v74/transfer"
	transferreversal "github.com/stripe/stripe-go/v74/transferreversal"
	treasury_creditreversal "github.com/stripe/stripe-go/v74/treasury/creditreversal"
	treasury_debitreversal "github.com/stripe/stripe-go/v74/treasury/debitreversal"
	treasury_financialaccount "github.com/stripe/stripe-go/v74/treasury/financialaccount"
	treasury_inboundtransfer "github.com/stripe/stripe-go/v74/treasury/inboundtransfer"
	treasury_outboundpayment "github.com/stripe/stripe-go/v74/treasury/outboundpayment"
	treasury_outboundtransfer "github.com/stripe/stripe-go/v74/treasury/outboundtransfer"
	treasury_receivedcredit "github.com/stripe/stripe-go/v74/treasury/receivedcredit"
	treasury_receiveddebit "github.com/stripe/stripe-go/v74/treasury/receiveddebit"
	treasury_transaction "github.com/stripe/stripe-go/v74/treasury/transaction"
	treasury_transactionentry "github.com/stripe/stripe-go/v74/treasury/transactionentry"
	usagerecord "github.com/stripe/stripe-go/v74/usagerecord"
	webhookendpoint "github.com/stripe/stripe-go/v74/webhookendpoint"
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
	result, err := apps_secret.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretDeleteWhere(t *testing.T) {
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

func TestAppsSecretFind(t *testing.T) {
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

func TestCheckoutSessionExpire(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("sess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionListLineItems(t *testing.T) {
	params := &stripe.CheckoutSessionListLineItemsParams{
		Session: stripe.String("sess_xyz"),
	}
	result := checkout_session.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestCashBalanceRetrieve(t *testing.T) {
	params := &stripe.CashBalanceParams{Customer: stripe.String("cus_123")}
	result, err := cashbalance.Get(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCashBalanceUpdate(t *testing.T) {
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

func TestCustomerCreateFundingInstructions(t *testing.T) {
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
	result, err := financialconnections_account.GetByID("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountDisconnect(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := financialconnections_account.Refresh("fca_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := financialconnections_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionRetrieve(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get("fcsess_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentCreate(t *testing.T) {
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

func TestPaymentIntentVerifyMicrodeposits(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, err := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := paymentlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinkRetrieve(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("pl_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := price.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := shippingrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationList(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTerminalConfigurationCreate(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationDelete(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationRetrieve(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("uc_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationUpdate(t *testing.T) {
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

func TestTestHelpersCustomerFundCashBalance(t *testing.T) {
	params := &stripe.TestHelpersCustomerFundCashBalanceParams{
		Amount:   stripe.Int64(30),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	result, err := testhelpers_customer.FundCashBalance("cus_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardDeliverCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardDeliverCardParams{}
	result, err := testhelpers_issuing_card.DeliverCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardFailCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardFailCardParams{}
	result, err := testhelpers_issuing_card.FailCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardReturnCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardReturnCardParams{}
	result, err := testhelpers_issuing_card.ReturnCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersIssuingCardShipCard(t *testing.T) {
	params := &stripe.TestHelpersIssuingCardShipCardParams{}
	result, err := testhelpers_issuing_card.ShipCard("card_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersRefundExpire(t *testing.T) {
	params := &stripe.TestHelpersRefundExpireParams{}
	result, err := testhelpers_refund.Expire("re_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockDelete(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockRetrieve(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockAdvance(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, err := testhelpers_testclock.Advance("clock_xyz", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransferFail(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferFailParams{
		FailureDetails: &stripe.TestHelpersTreasuryInboundTransferFailFailureDetailsParams{
			Code: stripe.String(string(stripe.TreasuryInboundTransferFailureDetailsCodeAccountClosed)),
		},
	}
	result, err := testhelpers_treasury_inboundtransfer.Fail("ibt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransferReturnInboundTransfer(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	result, err := testhelpers_treasury_inboundtransfer.ReturnInboundTransfer(
		"ibt_123",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryInboundTransferSucceed(t *testing.T) {
	params := &stripe.TestHelpersTreasuryInboundTransferSucceedParams{}
	result, err := testhelpers_treasury_inboundtransfer.Succeed(
		"ibt_123",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransferFail(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferFailParams{}
	result, err := testhelpers_treasury_outboundtransfer.Fail("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransferPost(t *testing.T) {
	params := &stripe.TestHelpersTreasuryOutboundTransferPostParams{}
	result, err := testhelpers_treasury_outboundtransfer.Post("obt_123", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTreasuryOutboundTransferReturnOutboundTransfer(
	t *testing.T,
) {
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

func TestTestHelpersTreasuryReceivedCreditCreate(t *testing.T) {
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

func TestTestHelpersTreasuryReceivedDebitCreate(t *testing.T) {
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

func TestTokenCreate(t *testing.T) {
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

func TestAccountLinkCreate(t *testing.T) {
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
	result, err := account.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountDelete(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.Del("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountRetrieve(t *testing.T) {
	params := &stripe.AccountParams{}
	result, err := account.GetByID("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountUpdate(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAccountReject(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, err := account.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := capability.Get("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCapabilityUpdate(t *testing.T) {
	params := &stripe.CapabilityParams{
		Requested: stripe.Bool(true),
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := capability.Update("card_payments", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Account:   stripe.String("acct_xxxxxxxxxxxxx"),
	}
	result, err := person.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPersonUpdate(t *testing.T) {
	params := &stripe.PersonParams{Account: stripe.String("acct_xxxxxxxxxxxxx")}
	params.AddMetadata("order_id", "6735")
	result, err := person.Update("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestAppsSecretList2(t *testing.T) {
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

func TestAppsSecretCreate2(t *testing.T) {
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

func TestBalanceTransactionList(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(3)
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.BalanceTransactionParams{}
	result, err := balancetransaction.Get("txn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := billingportal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationRetrieve(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{}
	result, err := billingportal_configuration.Get("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestBillingPortalConfigurationUpdate(t *testing.T) {
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

func TestBillingPortalSessionCreate(t *testing.T) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, err := billingportal_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Source:      &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := charge.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargeRetrieve(t *testing.T) {
	params := &stripe.ChargeParams{}
	result, err := charge.Get("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargeUpdate(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := charge.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestChargeCapture(t *testing.T) {
	params := &stripe.ChargeCaptureParams{}
	result, err := charge.Capture("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := checkout_session.Get("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCheckoutSessionExpire2(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, err := checkout_session.Expire("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := countryspec.Get("US", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := coupon.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponDelete(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Del("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponRetrieve(t *testing.T) {
	params := &stripe.CouponParams{}
	result, err := coupon.Get("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCouponUpdate(t *testing.T) {
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, err := coupon.Update("Z4OV52SU", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
	}
	result, err := customer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerDelete(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Del("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerRetrieve(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, err := customer.Get("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerUpdate(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, err := customer.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := customerbalancetransaction.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := customerbalancetransaction.Get("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCustomerBalanceTransactionUpdate(t *testing.T) {
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
		Type:     stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value:    stripe.String("DE123456789"),
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := taxid.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIDDelete(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Del("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxIDRetrieve(t *testing.T) {
	params := &stripe.TaxIDParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, err := taxid.Get("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := dispute.Get("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputeUpdate(t *testing.T) {
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := dispute.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDisputeClose(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, err := dispute.Close("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := event.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := financialconnections_account.GetByID(
		"fca_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsAccountDisconnect2(t *testing.T) {
	params := &stripe.FinancialConnectionsAccountDisconnectParams{}
	result, err := financialconnections_account.Disconnect(
		"fca_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := financialconnections_session.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFinancialConnectionsSessionRetrieve2(t *testing.T) {
	params := &stripe.FinancialConnectionsSessionParams{}
	result, err := financialconnections_session.Get(
		"fcsess_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := identity_verificationreport.Get("vr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := identity_verificationsession.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionRetrieve(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{}
	result, err := identity_verificationsession.Get("vs_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionUpdate(t *testing.T) {
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

func TestIdentityVerificationSessionCancel(t *testing.T) {
	params := &stripe.IdentityVerificationSessionCancelParams{}
	result, err := identity_verificationsession.Cancel(
		"vs_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIdentityVerificationSessionRedact(t *testing.T) {
	params := &stripe.IdentityVerificationSessionRedactParams{}
	result, err := identity_verificationsession.Redact(
		"vs_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := invoiceitem.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceItemDelete(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, err := invoiceitem.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceItemRetrieve(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, err := invoiceitem.Get("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceItemUpdate(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := invoiceitem.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := invoice.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceDelete(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceRetrieve(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, err := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceRetrieve2(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddExpand("customer")
	result, err := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceUpdate(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := invoice.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceFinalizeInvoice(t *testing.T) {
	params := &stripe.InvoiceFinalizeInvoiceParams{}
	result, err := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceMarkUncollectible(t *testing.T) {
	params := &stripe.InvoiceMarkUncollectibleParams{}
	result, err := invoice.MarkUncollectible("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoicePay(t *testing.T) {
	params := &stripe.InvoicePayParams{}
	result, err := invoice.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceSendInvoice(t *testing.T) {
	params := &stripe.InvoiceSendInvoiceParams{}
	result, err := invoice.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestInvoiceVoidInvoice(t *testing.T) {
	params := &stripe.InvoiceVoidInvoiceParams{}
	result, err := invoice.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := issuing_authorization.Get("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_authorization.Update("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationApprove(t *testing.T) {
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, err := issuing_authorization.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingAuthorizationDecline(t *testing.T) {
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, err := issuing_authorization.Decline("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := issuing_cardholder.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholderRetrieve(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	result, err := issuing_cardholder.Get("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardholderUpdate(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_cardholder.Update("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := issuing_card.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardRetrieve(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	result, err := issuing_card.Get("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingCardUpdate(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_card.Update("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := issuing_dispute.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputeRetrieve(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	result, err := issuing_dispute.Get("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingDisputeSubmit(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	result, err := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := issuing_transaction.Get("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestIssuingTransactionUpdate(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := issuing_transaction.Update("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestMandateRetrieve(t *testing.T) {
	params := &stripe.MandateParams{}
	result, err := mandate.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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

func TestPaymentIntentRetrieve(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	result, err := paymentintent.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentUpdate(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentApplyCustomerBalance(t *testing.T) {
	params := &stripe.PaymentIntentApplyCustomerBalanceParams{}
	result, err := paymentintent.ApplyCustomerBalance("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentCancel(t *testing.T) {
	params := &stripe.PaymentIntentCancelParams{}
	result, err := paymentintent.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentCapture(t *testing.T) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := paymentintent.Capture("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentConfirm(t *testing.T) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := paymentintent.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentIntentIncrementAuthorization(t *testing.T) {
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

func TestPaymentIntentVerifyMicrodeposits2(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := paymentlink.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinkRetrieve2(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, err := paymentlink.Get("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentLinkUpdate(t *testing.T) {
	params := &stripe.PaymentLinkParams{Active: stripe.Bool(false)}
	result, err := paymentlink.Update("plink_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := paymentmethod.Get("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodUpdate(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, err := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodAttach(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, err := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPaymentMethodDetach(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	result, err := paymentmethod.Detach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := payout.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutRetrieve(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Get("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutUpdate(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, err := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutCancel(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, err := payout.Cancel("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPayoutReverse(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	result, err := payout.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Product:  &stripe.PlanProductParams{Name: stripe.String("My product")},
	}
	result, err := plan.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlanDelete(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Del("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlanRetrieve(t *testing.T) {
	params := &stripe.PlanParams{}
	result, err := plan.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPlanUpdate(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, err := plan.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := price.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPriceRetrieve(t *testing.T) {
	params := &stripe.PriceParams{}
	result, err := price.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPriceUpdate(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := price.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := product.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductDelete(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Del("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductRetrieve(t *testing.T) {
	params := &stripe.ProductParams{}
	result, err := product.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProductUpdate(t *testing.T) {
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, err := product.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := promotioncode.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodeRetrieve(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	result, err := promotioncode.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestPromotionCodeUpdate(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, err := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := quote.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuoteRetrieve(t *testing.T) {
	params := &stripe.QuoteParams{}
	result, err := quote.Get("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuoteUpdate(t *testing.T) {
	params := &stripe.QuoteParams{}
	params.AddMetadata("order_id", "6735")
	result, err := quote.Update("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuoteAccept(t *testing.T) {
	params := &stripe.QuoteAcceptParams{}
	result, err := quote.Accept("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuoteCancel(t *testing.T) {
	params := &stripe.QuoteCancelParams{}
	result, err := quote.Cancel("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestQuoteFinalizeQuote(t *testing.T) {
	params := &stripe.QuoteFinalizeQuoteParams{}
	result, err := quote.FinalizeQuote("qt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := radar_earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := refund.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundRetrieve(t *testing.T) {
	params := &stripe.RefundParams{}
	result, err := refund.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundUpdate(t *testing.T) {
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, err := refund.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestRefundCancel(t *testing.T) {
	params := &stripe.RefundCancelParams{}
	result, err := refund.Cancel("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := review.Get("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestReviewApprove(t *testing.T) {
	params := &stripe.ReviewApproveParams{}
	result, err := review.Approve("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := setupintent.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentRetrieve(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	result, err := setupintent.Get("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentUpdate(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, err := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentCancel(t *testing.T) {
	params := &stripe.SetupIntentCancelParams{}
	result, err := setupintent.Cancel("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentConfirm(t *testing.T) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, err := setupintent.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSetupIntentVerifyMicrodeposits2(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{
		Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
	}
	result, err := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := shippingrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRateRetrieve(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	result, err := shippingrate.Get("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestShippingRateUpdate(t *testing.T) {
	params := &stripe.ShippingRateParams{}
	params.AddMetadata("order_id", "6735")
	result, err := shippingrate.Update("shr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := sigma_scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourceRetrieve(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourceRetrieve2(t *testing.T) {
	params := &stripe.SourceParams{}
	result, err := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.SourceParams{}
	params.AddMetadata("order_id", "6735")
	result, err := source.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemList(t *testing.T) {
	params := &stripe.SubscriptionItemListParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
	}
	result := subscriptionitem.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionItemCreate(t *testing.T) {
	params := &stripe.SubscriptionItemParams{
		Subscription: stripe.String("sub_xxxxxxxxxxxxx"),
		Price:        stripe.String("price_xxxxxxxxxxxxx"),
		Quantity:     stripe.Int64(2),
	}
	result, err := subscriptionitem.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemDelete(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Del("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemRetrieve(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	result, err := subscriptionitem.Get("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemUpdate(t *testing.T) {
	params := &stripe.SubscriptionItemParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscriptionitem.Update("si_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionItemUsageRecordSummaries(t *testing.T) {
	params := &stripe.SubscriptionItemUsageRecordSummariesParams{
		SubscriptionItem: stripe.String("si_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := subscriptionitem.UsageRecordSummaries(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestUsageRecordCreate(t *testing.T) {
	params := &stripe.UsageRecordParams{
		Quantity:         stripe.Int64(100),
		Timestamp:        stripe.Int64(1571252444),
		SubscriptionItem: stripe.String("si_xxxxxxxxxxxxx"),
	}
	result, err := usagerecord.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionScheduleList(t *testing.T) {
	params := &stripe.SubscriptionScheduleListParams{}
	params.Limit = stripe.Int64(3)
	result := subscriptionschedule.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionScheduleCreate(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		Customer:    stripe.String("cus_xxxxxxxxxxxxx"),
		StartDate:   stripe.Int64(1676070661),
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			&stripe.SubscriptionSchedulePhaseParams{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					&stripe.SubscriptionSchedulePhaseItemParams{
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

func TestSubscriptionScheduleRetrieve(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{}
	result, err := subscriptionschedule.Get("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionScheduleUpdate(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorRelease)),
	}
	result, err := subscriptionschedule.Update("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionScheduleCancel(t *testing.T) {
	params := &stripe.SubscriptionScheduleCancelParams{}
	result, err := subscriptionschedule.Cancel("sub_sched_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionScheduleRelease(t *testing.T) {
	params := &stripe.SubscriptionScheduleReleaseParams{}
	result, err := subscriptionschedule.Release(
		"sub_sched_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionList(t *testing.T) {
	params := &stripe.SubscriptionListParams{}
	params.Limit = stripe.Int64(3)
	result := subscription.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestSubscriptionCreate(t *testing.T) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Items: []*stripe.SubscriptionItemsParams{
			&stripe.SubscriptionItemsParams{
				Price: stripe.String("price_xxxxxxxxxxxxx"),
			},
		},
	}
	result, err := subscription.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionCancel(t *testing.T) {
	params := &stripe.SubscriptionCancelParams{}
	result, err := subscription.Cancel("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionRetrieve(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	result, err := subscription.Get("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionUpdate(t *testing.T) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	result, err := subscription.Update("sub_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSubscriptionSearch(t *testing.T) {
	params := &stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'active' AND metadata['order_id']:'6735'",
		},
	}
	result := subscription.Search(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
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
	result, err := taxcode.Get("txcd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := taxrate.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRateRetrieve(t *testing.T) {
	params := &stripe.TaxRateParams{}
	result, err := taxrate.Get("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxRateUpdate(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, err := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := terminal_configuration.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationDelete2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Del("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationRetrieve2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, err := terminal_configuration.Get("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConfigurationUpdate2(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		BBPOSWisePOSE: &stripe.TerminalConfigurationBBPOSWisePOSEParams{
			Splashscreen: stripe.String("file_xxxxxxxxxxxxx"),
		},
	}
	result, err := terminal_configuration.Update("tmc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalConnectionTokenCreate(t *testing.T) {
	params := &stripe.TerminalConnectionTokenParams{}
	result, err := terminal_connectiontoken.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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

func TestTerminalLocationDelete(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Del("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationRetrieve(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, err := terminal_location.Get("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalLocationUpdate(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, err := terminal_location.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := terminal_reader.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReaderDelete(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, err := terminal_reader.Del("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReaderUpdate(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, err := terminal_reader.Update("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReaderCancelAction(t *testing.T) {
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := terminal_reader.CancelAction("tmr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTerminalReaderProcessPaymentIntent(t *testing.T) {
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

func TestTerminalReaderProcessSetupIntent(t *testing.T) {
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
	result, err := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockDelete2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Del("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockRetrieve2(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, err := testhelpers_testclock.Get("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTestHelpersTestClockAdvance2(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(1675552261),
	}
	result, err := testhelpers_testclock.Advance("clock_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokenCreate2(t *testing.T) {
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

func TestTokenCreate3(t *testing.T) {
	params := &stripe.TokenParams{
		PII: &stripe.TokenPIIParams{IDNumber: stripe.String("000000000")},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokenCreate4(t *testing.T) {
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

func TestTokenCreate5(t *testing.T) {
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

func TestTokenCreate6(t *testing.T) {
	params := &stripe.TokenParams{
		CVCUpdate: &stripe.TokenCVCUpdateParams{CVC: stripe.String("123")},
	}
	result, err := token.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTokenRetrieve(t *testing.T) {
	params := &stripe.TokenParams{}
	result, err := token.Get("tok_xxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := topup.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupRetrieve(t *testing.T) {
	params := &stripe.TopupParams{}
	result, err := topup.Get("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTopupUpdate(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, err := topup.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := transfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransferRetrieve(t *testing.T) {
	params := &stripe.TransferParams{}
	result, err := transfer.Get("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransferUpdate(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, err := transfer.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransferReversalList(t *testing.T) {
	params := &stripe.TransferReversalListParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.Limit = stripe.Int64(3)
	result := transferreversal.List(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTransferReversalCreate(t *testing.T) {
	params := &stripe.TransferReversalParams{
		Amount: stripe.Int64(100),
		ID:     stripe.String("tr_xxxxxxxxxxxxx"),
	}
	result, err := transferreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTransferReversalUpdate(t *testing.T) {
	params := &stripe.TransferReversalParams{
		ID: stripe.String("tr_xxxxxxxxxxxxx"),
	}
	params.AddMetadata("order_id", "6735")
	result, err := transferreversal.Update("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_creditreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryCreditReversalRetrieve(t *testing.T) {
	params := &stripe.TreasuryCreditReversalParams{}
	result, err := treasury_creditreversal.Get("credrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_debitreversal.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryDebitReversalRetrieve(t *testing.T) {
	params := &stripe.TreasuryDebitReversalParams{}
	result, err := treasury_debitreversal.Get("debrev_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_financialaccount.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountRetrieve(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	result, err := treasury_financialaccount.Get("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountUpdate(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountParams{}
	params.AddMetadata("order_id", "6735")
	result, err := treasury_financialaccount.Update("fa_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryFinancialAccountRetrieveFeatures(t *testing.T) {
	params := &stripe.TreasuryFinancialAccountRetrieveFeaturesParams{}
	result, err := treasury_financialaccount.RetrieveFeatures(
		"fa_xxxxxxxxxxxxx",
		params,
	)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_inboundtransfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransferRetrieve(t *testing.T) {
	params := &stripe.TreasuryInboundTransferParams{}
	result, err := treasury_inboundtransfer.Get("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryInboundTransferCancel(t *testing.T) {
	params := &stripe.TreasuryInboundTransferCancelParams{}
	result, err := treasury_inboundtransfer.Cancel("ibt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
		Customer:                 stripe.String("cus_xxxxxxxxxxxxx"),
		DestinationPaymentMethod: stripe.String("pm_xxxxxxxxxxxxx"),
		Description:              stripe.String("OutboundPayment to a 3rd party"),
	}
	result, err := treasury_outboundpayment.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentRetrieve(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentParams{}
	result, err := treasury_outboundpayment.Get("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundPaymentCancel(t *testing.T) {
	params := &stripe.TreasuryOutboundPaymentCancelParams{}
	result, err := treasury_outboundpayment.Cancel("bot_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_outboundtransfer.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransferRetrieve(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferParams{}
	result, err := treasury_outboundtransfer.Get("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTreasuryOutboundTransferCancel(t *testing.T) {
	params := &stripe.TreasuryOutboundTransferCancelParams{}
	result, err := treasury_outboundtransfer.Cancel("obt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_receivedcredit.Get("rc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_receiveddebit.Get("rd_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_transactionentry.Get("trxne_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := treasury_transaction.Get("trxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
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
	result, err := webhookendpoint.New(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointDelete(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Del("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointRetrieve(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, err := webhookendpoint.Get("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, err := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxTransactionCreateFromCalculation(t *testing.T) {
	params := &stripe.TaxTransactionCreateFromCalculationParams{
		Calculation: stripe.String("xxx"),
		Reference:   stripe.String("yyy"),
	}
	result, err := tax_transaction.CreateFromCalculation(params)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestTaxCalculationListLineItems(t *testing.T) {
	params := &stripe.TaxCalculationListLineItemsParams{
		Calculation: stripe.String("xxx"),
	}
	result := tax_calculation.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestPaymentIntentCreate3(t *testing.T) {
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

func TestQuoteListLineItems(t *testing.T) {
	params := &stripe.QuoteListLineItemsParams{
		Quote: stripe.String("qt_xxxxxxxxxxxxx"),
	}
	result := quote.ListLineItems(params)
	assert.NotNil(t, result)
	assert.Nil(t, result.Err())
}

func TestTaxCalculationCreate(t *testing.T) {
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		LineItems: []*stripe.TaxCalculationLineItemParams{
			&stripe.TaxCalculationLineItemParams{
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
