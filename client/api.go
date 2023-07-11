//
//
// File generated from our OpenAPI spec
//
//

// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/account"
	"github.com/stripe/stripe-go/v74/accountlink"
	"github.com/stripe/stripe-go/v74/applepaydomain"
	"github.com/stripe/stripe-go/v74/applicationfee"
	appssecret "github.com/stripe/stripe-go/v74/apps/secret"
	"github.com/stripe/stripe-go/v74/balance"
	"github.com/stripe/stripe-go/v74/balancetransaction"
	"github.com/stripe/stripe-go/v74/bankaccount"
	billingportalconfiguration "github.com/stripe/stripe-go/v74/billingportal/configuration"
	billingportalsession "github.com/stripe/stripe-go/v74/billingportal/session"
	"github.com/stripe/stripe-go/v74/capability"
	"github.com/stripe/stripe-go/v74/card"
	"github.com/stripe/stripe-go/v74/cashbalance"
	"github.com/stripe/stripe-go/v74/charge"
	checkoutsession "github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/countryspec"
	"github.com/stripe/stripe-go/v74/coupon"
	"github.com/stripe/stripe-go/v74/creditnote"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/customerbalancetransaction"
	"github.com/stripe/stripe-go/v74/customercashbalancetransaction"
	"github.com/stripe/stripe-go/v74/dispute"
	"github.com/stripe/stripe-go/v74/ephemeralkey"
	"github.com/stripe/stripe-go/v74/event"
	"github.com/stripe/stripe-go/v74/feerefund"
	"github.com/stripe/stripe-go/v74/file"
	"github.com/stripe/stripe-go/v74/filelink"
	financialconnectionsaccount "github.com/stripe/stripe-go/v74/financialconnections/account"
	financialconnectionssession "github.com/stripe/stripe-go/v74/financialconnections/session"
	identityverificationreport "github.com/stripe/stripe-go/v74/identity/verificationreport"
	identityverificationsession "github.com/stripe/stripe-go/v74/identity/verificationsession"
	"github.com/stripe/stripe-go/v74/invoice"
	"github.com/stripe/stripe-go/v74/invoiceitem"
	issuingauthorization "github.com/stripe/stripe-go/v74/issuing/authorization"
	issuingcard "github.com/stripe/stripe-go/v74/issuing/card"
	issuingcardholder "github.com/stripe/stripe-go/v74/issuing/cardholder"
	issuingdispute "github.com/stripe/stripe-go/v74/issuing/dispute"
	issuingtransaction "github.com/stripe/stripe-go/v74/issuing/transaction"
	"github.com/stripe/stripe-go/v74/loginlink"
	"github.com/stripe/stripe-go/v74/mandate"
	"github.com/stripe/stripe-go/v74/oauth"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/paymentlink"
	"github.com/stripe/stripe-go/v74/paymentmethod"
	"github.com/stripe/stripe-go/v74/paymentsource"
	"github.com/stripe/stripe-go/v74/payout"
	"github.com/stripe/stripe-go/v74/person"
	"github.com/stripe/stripe-go/v74/plan"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/promotioncode"
	"github.com/stripe/stripe-go/v74/quote"
	radarearlyfraudwarning "github.com/stripe/stripe-go/v74/radar/earlyfraudwarning"
	radarvaluelist "github.com/stripe/stripe-go/v74/radar/valuelist"
	radarvaluelistitem "github.com/stripe/stripe-go/v74/radar/valuelistitem"
	"github.com/stripe/stripe-go/v74/refund"
	reportingreportrun "github.com/stripe/stripe-go/v74/reporting/reportrun"
	reportingreporttype "github.com/stripe/stripe-go/v74/reporting/reporttype"
	"github.com/stripe/stripe-go/v74/review"
	"github.com/stripe/stripe-go/v74/setupattempt"
	"github.com/stripe/stripe-go/v74/setupintent"
	"github.com/stripe/stripe-go/v74/shippingrate"
	sigmascheduledqueryrun "github.com/stripe/stripe-go/v74/sigma/scheduledqueryrun"
	"github.com/stripe/stripe-go/v74/source"
	"github.com/stripe/stripe-go/v74/sourcetransaction"
	"github.com/stripe/stripe-go/v74/subscription"
	"github.com/stripe/stripe-go/v74/subscriptionitem"
	"github.com/stripe/stripe-go/v74/subscriptionschedule"
	taxcalculation "github.com/stripe/stripe-go/v74/tax/calculation"
	taxsettings "github.com/stripe/stripe-go/v74/tax/settings"
	taxtransaction "github.com/stripe/stripe-go/v74/tax/transaction"
	"github.com/stripe/stripe-go/v74/taxcode"
	"github.com/stripe/stripe-go/v74/taxid"
	"github.com/stripe/stripe-go/v74/taxrate"
	terminalconfiguration "github.com/stripe/stripe-go/v74/terminal/configuration"
	terminalconnectiontoken "github.com/stripe/stripe-go/v74/terminal/connectiontoken"
	terminallocation "github.com/stripe/stripe-go/v74/terminal/location"
	terminalreader "github.com/stripe/stripe-go/v74/terminal/reader"
	testhelperscustomer "github.com/stripe/stripe-go/v74/testhelpers/customer"
	testhelpersissuingcard "github.com/stripe/stripe-go/v74/testhelpers/issuing/card"
	testhelpersrefund "github.com/stripe/stripe-go/v74/testhelpers/refund"
	testhelpersterminalreader "github.com/stripe/stripe-go/v74/testhelpers/terminal/reader"
	testhelperstestclock "github.com/stripe/stripe-go/v74/testhelpers/testclock"
	testhelperstreasuryinboundtransfer "github.com/stripe/stripe-go/v74/testhelpers/treasury/inboundtransfer"
	testhelperstreasuryoutboundpayment "github.com/stripe/stripe-go/v74/testhelpers/treasury/outboundpayment"
	testhelperstreasuryoutboundtransfer "github.com/stripe/stripe-go/v74/testhelpers/treasury/outboundtransfer"
	testhelperstreasuryreceivedcredit "github.com/stripe/stripe-go/v74/testhelpers/treasury/receivedcredit"
	testhelperstreasuryreceiveddebit "github.com/stripe/stripe-go/v74/testhelpers/treasury/receiveddebit"
	"github.com/stripe/stripe-go/v74/token"
	"github.com/stripe/stripe-go/v74/topup"
	"github.com/stripe/stripe-go/v74/transfer"
	"github.com/stripe/stripe-go/v74/transferreversal"
	treasurycreditreversal "github.com/stripe/stripe-go/v74/treasury/creditreversal"
	treasurydebitreversal "github.com/stripe/stripe-go/v74/treasury/debitreversal"
	treasuryfinancialaccount "github.com/stripe/stripe-go/v74/treasury/financialaccount"
	treasuryinboundtransfer "github.com/stripe/stripe-go/v74/treasury/inboundtransfer"
	treasuryoutboundpayment "github.com/stripe/stripe-go/v74/treasury/outboundpayment"
	treasuryoutboundtransfer "github.com/stripe/stripe-go/v74/treasury/outboundtransfer"
	treasuryreceivedcredit "github.com/stripe/stripe-go/v74/treasury/receivedcredit"
	treasuryreceiveddebit "github.com/stripe/stripe-go/v74/treasury/receiveddebit"
	treasurytransaction "github.com/stripe/stripe-go/v74/treasury/transaction"
	treasurytransactionentry "github.com/stripe/stripe-go/v74/treasury/transactionentry"
	"github.com/stripe/stripe-go/v74/usagerecord"
	"github.com/stripe/stripe-go/v74/usagerecordsummary"
	"github.com/stripe/stripe-go/v74/webhookendpoint"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// AccountLinks is the client used to invoke /account_links APIs.
	AccountLinks *accountlink.Client
	// Accounts is the client used to invoke /accounts APIs.
	Accounts *account.Client
	// ApplePayDomains is the client used to invoke /apple_pay/domains APIs.
	ApplePayDomains *applepaydomain.Client
	// ApplicationFees is the client used to invoke /application_fees APIs.
	ApplicationFees *applicationfee.Client
	// AppsSecrets is the client used to invoke /apps/secrets APIs.
	AppsSecrets *appssecret.Client
	// Balance is the client used to invoke /balance APIs.
	Balance *balance.Client
	// BalanceTransactions is the client used to invoke /balance_transactions APIs.
	BalanceTransactions *balancetransaction.Client
	// BankAccounts is the client used to invoke bankaccount related APIs.
	BankAccounts *bankaccount.Client
	// BillingPortalConfigurations is the client used to invoke /billing_portal/configurations APIs.
	BillingPortalConfigurations *billingportalconfiguration.Client
	// BillingPortalSessions is the client used to invoke /billing_portal/sessions APIs.
	BillingPortalSessions *billingportalsession.Client
	// Capabilities is the client used to invoke /accounts/{account}/capabilities APIs.
	Capabilities *capability.Client
	// Cards is the client used to invoke card related APIs.
	Cards *card.Client
	// CashBalances is the client used to invoke /customers/{customer}/cash_balance APIs.
	CashBalances *cashbalance.Client
	// Charges is the client used to invoke /charges APIs.
	Charges *charge.Client
	// CheckoutSessions is the client used to invoke /checkout/sessions APIs.
	CheckoutSessions *checkoutsession.Client
	// CountrySpecs is the client used to invoke /country_specs APIs.
	CountrySpecs *countryspec.Client
	// Coupons is the client used to invoke /coupons APIs.
	Coupons *coupon.Client
	// CreditNotes is the client used to invoke /credit_notes APIs.
	CreditNotes *creditnote.Client
	// CustomerBalanceTransactions is the client used to invoke /customers/{customer}/balance_transactions APIs.
	CustomerBalanceTransactions *customerbalancetransaction.Client
	// CustomerCashBalanceTransactions is the client used to invoke /customers/{customer}/cash_balance_transactions APIs.
	CustomerCashBalanceTransactions *customercashbalancetransaction.Client
	// Customers is the client used to invoke /customers APIs.
	Customers *customer.Client
	// Disputes is the client used to invoke /disputes APIs.
	Disputes *dispute.Client
	// EphemeralKeys is the client used to invoke /ephemeral_keys APIs.
	EphemeralKeys *ephemeralkey.Client
	// Events is the client used to invoke /events APIs.
	Events *event.Client
	// FeeRefunds is the client used to invoke /application_fees/{id}/refunds APIs.
	FeeRefunds *feerefund.Client
	// FileLinks is the client used to invoke /file_links APIs.
	FileLinks *filelink.Client
	// Files is the client used to invoke /files APIs.
	Files *file.Client
	// FinancialConnectionsAccounts is the client used to invoke /financial_connections/accounts APIs.
	FinancialConnectionsAccounts *financialconnectionsaccount.Client
	// FinancialConnectionsSessions is the client used to invoke /financial_connections/sessions APIs.
	FinancialConnectionsSessions *financialconnectionssession.Client
	// IdentityVerificationReports is the client used to invoke /identity/verification_reports APIs.
	IdentityVerificationReports *identityverificationreport.Client
	// IdentityVerificationSessions is the client used to invoke /identity/verification_sessions APIs.
	IdentityVerificationSessions *identityverificationsession.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	InvoiceItems *invoiceitem.Client
	// Invoices is the client used to invoke /invoices APIs.
	Invoices *invoice.Client
	// IssuingAuthorizations is the client used to invoke /issuing/authorizations APIs.
	IssuingAuthorizations *issuingauthorization.Client
	// IssuingCardholders is the client used to invoke /issuing/cardholders APIs.
	IssuingCardholders *issuingcardholder.Client
	// IssuingCards is the client used to invoke /issuing/cards APIs.
	IssuingCards *issuingcard.Client
	// IssuingDisputes is the client used to invoke /issuing/disputes APIs.
	IssuingDisputes *issuingdispute.Client
	// IssuingTransactions is the client used to invoke /issuing/transactions APIs.
	IssuingTransactions *issuingtransaction.Client
	// LoginLinks is the client used to invoke /accounts/{account}/login_links APIs.
	LoginLinks *loginlink.Client
	// Mandates is the client used to invoke /mandates APIs.
	Mandates *mandate.Client
	// OAuth is the client used to invoke /oauth APIs
	OAuth *oauth.Client
	// PaymentIntents is the client used to invoke /payment_intents APIs.
	PaymentIntents *paymentintent.Client
	// PaymentLinks is the client used to invoke /payment_links APIs.
	PaymentLinks *paymentlink.Client
	// PaymentMethods is the client used to invoke /payment_methods APIs.
	PaymentMethods *paymentmethod.Client
	// PaymentSources is the client used to invoke /customers/{customer}/sources APIs.
	PaymentSources *paymentsource.Client
	// Payouts is the client used to invoke /payouts APIs.
	Payouts *payout.Client
	// Persons is the client used to invoke /accounts/{account}/persons APIs.
	Persons *person.Client
	// Plans is the client used to invoke /plans APIs.
	Plans *plan.Client
	// Prices is the client used to invoke /prices APIs.
	Prices *price.Client
	// Products is the client used to invoke /products APIs.
	Products *product.Client
	// PromotionCodes is the client used to invoke /promotion_codes APIs.
	PromotionCodes *promotioncode.Client
	// Quotes is the client used to invoke /quotes APIs.
	Quotes *quote.Client
	// RadarEarlyFraudWarnings is the client used to invoke /radar/early_fraud_warnings APIs.
	RadarEarlyFraudWarnings *radarearlyfraudwarning.Client
	// RadarValueListItems is the client used to invoke /radar/value_list_items APIs.
	RadarValueListItems *radarvaluelistitem.Client
	// RadarValueLists is the client used to invoke /radar/value_lists APIs.
	RadarValueLists *radarvaluelist.Client
	// Refunds is the client used to invoke /refunds APIs.
	Refunds *refund.Client
	// ReportingReportRuns is the client used to invoke /reporting/report_runs APIs.
	ReportingReportRuns *reportingreportrun.Client
	// ReportingReportTypes is the client used to invoke /reporting/report_types APIs.
	ReportingReportTypes *reportingreporttype.Client
	// Reviews is the client used to invoke /reviews APIs.
	Reviews *review.Client
	// SetupAttempts is the client used to invoke /setup_attempts APIs.
	SetupAttempts *setupattempt.Client
	// SetupIntents is the client used to invoke /setup_intents APIs.
	SetupIntents *setupintent.Client
	// ShippingRates is the client used to invoke /shipping_rates APIs.
	ShippingRates *shippingrate.Client
	// SigmaScheduledQueryRuns is the client used to invoke /sigma/scheduled_query_runs APIs.
	SigmaScheduledQueryRuns *sigmascheduledqueryrun.Client
	// Sources is the client used to invoke /sources APIs.
	Sources *source.Client
	// SourceTransactions is the client used to invoke sourcetransaction related APIs.
	SourceTransactions *sourcetransaction.Client
	// SubscriptionItems is the client used to invoke /subscription_items APIs.
	SubscriptionItems *subscriptionitem.Client
	// Subscriptions is the client used to invoke /subscriptions APIs.
	Subscriptions *subscription.Client
	// SubscriptionSchedules is the client used to invoke /subscription_schedules APIs.
	SubscriptionSchedules *subscriptionschedule.Client
	// TaxCalculations is the client used to invoke /tax/calculations APIs.
	TaxCalculations *taxcalculation.Client
	// TaxCodes is the client used to invoke /tax_codes APIs.
	TaxCodes *taxcode.Client
	// TaxIDs is the client used to invoke /customers/{customer}/tax_ids APIs.
	TaxIDs *taxid.Client
	// TaxRates is the client used to invoke /tax_rates APIs.
	TaxRates *taxrate.Client
	// TaxSettings is the client used to invoke /tax/settings APIs.
	TaxSettings *taxsettings.Client
	// TaxTransactions is the client used to invoke /tax/transactions APIs.
	TaxTransactions *taxtransaction.Client
	// TerminalConfigurations is the client used to invoke /terminal/configurations APIs.
	TerminalConfigurations *terminalconfiguration.Client
	// TerminalConnectionTokens is the client used to invoke /terminal/connection_tokens APIs.
	TerminalConnectionTokens *terminalconnectiontoken.Client
	// TerminalLocations is the client used to invoke /terminal/locations APIs.
	TerminalLocations *terminallocation.Client
	// TerminalReaders is the client used to invoke /terminal/readers APIs.
	TerminalReaders *terminalreader.Client
	// TestHelpersCustomers is the client used to invoke /customers APIs.
	TestHelpersCustomers *testhelperscustomer.Client
	// TestHelpersIssuingCards is the client used to invoke /issuing/cards APIs.
	TestHelpersIssuingCards *testhelpersissuingcard.Client
	// TestHelpersRefunds is the client used to invoke /refunds APIs.
	TestHelpersRefunds *testhelpersrefund.Client
	// TestHelpersTerminalReaders is the client used to invoke /terminal/readers APIs.
	TestHelpersTerminalReaders *testhelpersterminalreader.Client
	// TestHelpersTestClocks is the client used to invoke /test_helpers/test_clocks APIs.
	TestHelpersTestClocks *testhelperstestclock.Client
	// TestHelpersTreasuryInboundTransfers is the client used to invoke /treasury/inbound_transfers APIs.
	TestHelpersTreasuryInboundTransfers *testhelperstreasuryinboundtransfer.Client
	// TestHelpersTreasuryOutboundPayments is the client used to invoke /treasury/outbound_payments APIs.
	TestHelpersTreasuryOutboundPayments *testhelperstreasuryoutboundpayment.Client
	// TestHelpersTreasuryOutboundTransfers is the client used to invoke /treasury/outbound_transfers APIs.
	TestHelpersTreasuryOutboundTransfers *testhelperstreasuryoutboundtransfer.Client
	// TestHelpersTreasuryReceivedCredits is the client used to invoke /treasury/received_credits APIs.
	TestHelpersTreasuryReceivedCredits *testhelperstreasuryreceivedcredit.Client
	// TestHelpersTreasuryReceivedDebits is the client used to invoke /treasury/received_debits APIs.
	TestHelpersTreasuryReceivedDebits *testhelperstreasuryreceiveddebit.Client
	// Tokens is the client used to invoke /tokens APIs.
	Tokens *token.Client
	// Topups is the client used to invoke /topups APIs.
	Topups *topup.Client
	// TransferReversals is the client used to invoke /transfers/{id}/reversals APIs.
	TransferReversals *transferreversal.Client
	// Transfers is the client used to invoke /transfers APIs.
	Transfers *transfer.Client
	// TreasuryCreditReversals is the client used to invoke /treasury/credit_reversals APIs.
	TreasuryCreditReversals *treasurycreditreversal.Client
	// TreasuryDebitReversals is the client used to invoke /treasury/debit_reversals APIs.
	TreasuryDebitReversals *treasurydebitreversal.Client
	// TreasuryFinancialAccounts is the client used to invoke /treasury/financial_accounts APIs.
	TreasuryFinancialAccounts *treasuryfinancialaccount.Client
	// TreasuryInboundTransfers is the client used to invoke /treasury/inbound_transfers APIs.
	TreasuryInboundTransfers *treasuryinboundtransfer.Client
	// TreasuryOutboundPayments is the client used to invoke /treasury/outbound_payments APIs.
	TreasuryOutboundPayments *treasuryoutboundpayment.Client
	// TreasuryOutboundTransfers is the client used to invoke /treasury/outbound_transfers APIs.
	TreasuryOutboundTransfers *treasuryoutboundtransfer.Client
	// TreasuryReceivedCredits is the client used to invoke /treasury/received_credits APIs.
	TreasuryReceivedCredits *treasuryreceivedcredit.Client
	// TreasuryReceivedDebits is the client used to invoke /treasury/received_debits APIs.
	TreasuryReceivedDebits *treasuryreceiveddebit.Client
	// TreasuryTransactionEntries is the client used to invoke /treasury/transaction_entries APIs.
	TreasuryTransactionEntries *treasurytransactionentry.Client
	// TreasuryTransactions is the client used to invoke /treasury/transactions APIs.
	TreasuryTransactions *treasurytransaction.Client
	// UsageRecords is the client used to invoke /subscription_items/{subscription_item}/usage_records APIs.
	UsageRecords *usagerecord.Client
	// UsageRecordSummaries is the client used to invoke /subscription_items/{subscription_item}/usage_record_summaries APIs.
	UsageRecordSummaries *usagerecordsummary.Client
	// WebhookEndpoints is the client used to invoke /webhook_endpoints APIs.
	WebhookEndpoints *webhookendpoint.Client
}

func (a *API) Init(key string, backends *stripe.Backends) {

	if backends == nil {
		backends = &stripe.Backends{
			API:     stripe.GetBackend(stripe.APIBackend),
			Connect: stripe.GetBackend(stripe.ConnectBackend),
			Uploads: stripe.GetBackend(stripe.UploadsBackend),
		}
	}

	a.AccountLinks = &accountlink.Client{B: backends.API, Key: key}
	a.Accounts = &account.Client{B: backends.API, Key: key}
	a.ApplePayDomains = &applepaydomain.Client{B: backends.API, Key: key}
	a.ApplicationFees = &applicationfee.Client{B: backends.API, Key: key}
	a.AppsSecrets = &appssecret.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.BalanceTransactions = &balancetransaction.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.BillingPortalConfigurations = &billingportalconfiguration.Client{B: backends.API, Key: key}
	a.BillingPortalSessions = &billingportalsession.Client{B: backends.API, Key: key}
	a.Capabilities = &capability.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.CashBalances = &cashbalance.Client{B: backends.API, Key: key}
	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.CheckoutSessions = &checkoutsession.Client{B: backends.API, Key: key}
	a.CountrySpecs = &countryspec.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.CreditNotes = &creditnote.Client{B: backends.API, Key: key}
	a.CustomerBalanceTransactions = &customerbalancetransaction.Client{B: backends.API, Key: key}
	a.CustomerCashBalanceTransactions = &customercashbalancetransaction.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.EphemeralKeys = &ephemeralkey.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.FileLinks = &filelink.Client{B: backends.API, Key: key}
	a.Files = &file.Client{B: backends.Uploads, Key: key}
	a.FinancialConnectionsAccounts = &financialconnectionsaccount.Client{B: backends.API, Key: key}
	a.FinancialConnectionsSessions = &financialconnectionssession.Client{B: backends.API, Key: key}
	a.IdentityVerificationReports = &identityverificationreport.Client{B: backends.API, Key: key}
	a.IdentityVerificationSessions = &identityverificationsession.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.IssuingAuthorizations = &issuingauthorization.Client{B: backends.API, Key: key}
	a.IssuingCardholders = &issuingcardholder.Client{B: backends.API, Key: key}
	a.IssuingCards = &issuingcard.Client{B: backends.API, Key: key}
	a.IssuingDisputes = &issuingdispute.Client{B: backends.API, Key: key}
	a.IssuingTransactions = &issuingtransaction.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Mandates = &mandate.Client{B: backends.API, Key: key}
	a.OAuth = &oauth.Client{B: backends.Connect, Key: key}
	a.PaymentIntents = &paymentintent.Client{B: backends.API, Key: key}
	a.PaymentLinks = &paymentlink.Client{B: backends.API, Key: key}
	a.PaymentMethods = &paymentmethod.Client{B: backends.API, Key: key}
	a.PaymentSources = &paymentsource.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Persons = &person.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Prices = &price.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.PromotionCodes = &promotioncode.Client{B: backends.API, Key: key}
	a.Quotes = &quote.Client{B: backends.API, PDFBackend: backends.Uploads, Key: key}
	a.RadarEarlyFraudWarnings = &radarearlyfraudwarning.Client{B: backends.API, Key: key}
	a.RadarValueListItems = &radarvaluelistitem.Client{B: backends.API, Key: key}
	a.RadarValueLists = &radarvaluelist.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.ReportingReportRuns = &reportingreportrun.Client{B: backends.API, Key: key}
	a.ReportingReportTypes = &reportingreporttype.Client{B: backends.API, Key: key}
	a.Reviews = &review.Client{B: backends.API, Key: key}
	a.SetupAttempts = &setupattempt.Client{B: backends.API, Key: key}
	a.SetupIntents = &setupintent.Client{B: backends.API, Key: key}
	a.ShippingRates = &shippingrate.Client{B: backends.API, Key: key}
	a.SigmaScheduledQueryRuns = &sigmascheduledqueryrun.Client{B: backends.API, Key: key}
	a.Sources = &source.Client{B: backends.API, Key: key}
	a.SourceTransactions = &sourcetransaction.Client{B: backends.API, Key: key}
	a.SubscriptionItems = &subscriptionitem.Client{B: backends.API, Key: key}
	a.Subscriptions = &subscription.Client{B: backends.API, Key: key}
	a.SubscriptionSchedules = &subscriptionschedule.Client{B: backends.API, Key: key}
	a.TaxCalculations = &taxcalculation.Client{B: backends.API, Key: key}
	a.TaxCodes = &taxcode.Client{B: backends.API, Key: key}
	a.TaxIDs = &taxid.Client{B: backends.API, Key: key}
	a.TaxRates = &taxrate.Client{B: backends.API, Key: key}
	a.TaxSettings = &taxsettings.Client{B: backends.API, Key: key}
	a.TaxTransactions = &taxtransaction.Client{B: backends.API, Key: key}
	a.TerminalConfigurations = &terminalconfiguration.Client{B: backends.API, Key: key}
	a.TerminalConnectionTokens = &terminalconnectiontoken.Client{B: backends.API, Key: key}
	a.TerminalLocations = &terminallocation.Client{B: backends.API, Key: key}
	a.TerminalReaders = &terminalreader.Client{B: backends.API, Key: key}
	a.TestHelpersCustomers = &testhelperscustomer.Client{B: backends.API, Key: key}
	a.TestHelpersIssuingCards = &testhelpersissuingcard.Client{B: backends.API, Key: key}
	a.TestHelpersRefunds = &testhelpersrefund.Client{B: backends.API, Key: key}
	a.TestHelpersTerminalReaders = &testhelpersterminalreader.Client{B: backends.API, Key: key}
	a.TestHelpersTestClocks = &testhelperstestclock.Client{B: backends.API, Key: key}
	a.TestHelpersTreasuryInboundTransfers = &testhelperstreasuryinboundtransfer.Client{B: backends.API, Key: key}
	a.TestHelpersTreasuryOutboundPayments = &testhelperstreasuryoutboundpayment.Client{B: backends.API, Key: key}
	a.TestHelpersTreasuryOutboundTransfers = &testhelperstreasuryoutboundtransfer.Client{B: backends.API, Key: key}
	a.TestHelpersTreasuryReceivedCredits = &testhelperstreasuryreceivedcredit.Client{B: backends.API, Key: key}
	a.TestHelpersTreasuryReceivedDebits = &testhelperstreasuryreceiveddebit.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.Topups = &topup.Client{B: backends.API, Key: key}
	a.TransferReversals = &transferreversal.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
	a.TreasuryCreditReversals = &treasurycreditreversal.Client{B: backends.API, Key: key}
	a.TreasuryDebitReversals = &treasurydebitreversal.Client{B: backends.API, Key: key}
	a.TreasuryFinancialAccounts = &treasuryfinancialaccount.Client{B: backends.API, Key: key}
	a.TreasuryInboundTransfers = &treasuryinboundtransfer.Client{B: backends.API, Key: key}
	a.TreasuryOutboundPayments = &treasuryoutboundpayment.Client{B: backends.API, Key: key}
	a.TreasuryOutboundTransfers = &treasuryoutboundtransfer.Client{B: backends.API, Key: key}
	a.TreasuryReceivedCredits = &treasuryreceivedcredit.Client{B: backends.API, Key: key}
	a.TreasuryReceivedDebits = &treasuryreceiveddebit.Client{B: backends.API, Key: key}
	a.TreasuryTransactionEntries = &treasurytransactionentry.Client{B: backends.API, Key: key}
	a.TreasuryTransactions = &treasurytransaction.Client{B: backends.API, Key: key}
	a.UsageRecords = &usagerecord.Client{B: backends.API, Key: key}
	a.UsageRecordSummaries = &usagerecordsummary.Client{B: backends.API, Key: key}
	a.WebhookEndpoints = &webhookendpoint.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *stripe.Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
