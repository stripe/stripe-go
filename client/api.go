//
//
// File generated from our OpenAPI spec
//
//

// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/account"
	"github.com/max-cape/stripe-go-test/accountlink"
	"github.com/max-cape/stripe-go-test/accountsession"
	"github.com/max-cape/stripe-go-test/applepaydomain"
	"github.com/max-cape/stripe-go-test/applicationfee"
	appssecret "github.com/max-cape/stripe-go-test/apps/secret"
	"github.com/max-cape/stripe-go-test/balance"
	"github.com/max-cape/stripe-go-test/balancetransaction"
	"github.com/max-cape/stripe-go-test/bankaccount"
	billingalert "github.com/max-cape/stripe-go-test/billing/alert"
	billingcreditbalancesummary "github.com/max-cape/stripe-go-test/billing/creditbalancesummary"
	billingcreditbalancetransaction "github.com/max-cape/stripe-go-test/billing/creditbalancetransaction"
	billingcreditgrant "github.com/max-cape/stripe-go-test/billing/creditgrant"
	billingmeter "github.com/max-cape/stripe-go-test/billing/meter"
	billingmeterevent "github.com/max-cape/stripe-go-test/billing/meterevent"
	billingmetereventadjustment "github.com/max-cape/stripe-go-test/billing/metereventadjustment"
	billingmetereventsummary "github.com/max-cape/stripe-go-test/billing/metereventsummary"
	billingportalconfiguration "github.com/max-cape/stripe-go-test/billingportal/configuration"
	billingportalsession "github.com/max-cape/stripe-go-test/billingportal/session"
	"github.com/max-cape/stripe-go-test/capability"
	"github.com/max-cape/stripe-go-test/card"
	"github.com/max-cape/stripe-go-test/cashbalance"
	"github.com/max-cape/stripe-go-test/charge"
	checkoutsession "github.com/max-cape/stripe-go-test/checkout/session"
	climateorder "github.com/max-cape/stripe-go-test/climate/order"
	climateproduct "github.com/max-cape/stripe-go-test/climate/product"
	climatesupplier "github.com/max-cape/stripe-go-test/climate/supplier"
	"github.com/max-cape/stripe-go-test/confirmationtoken"
	"github.com/max-cape/stripe-go-test/countryspec"
	"github.com/max-cape/stripe-go-test/coupon"
	"github.com/max-cape/stripe-go-test/creditnote"
	"github.com/max-cape/stripe-go-test/customer"
	"github.com/max-cape/stripe-go-test/customerbalancetransaction"
	"github.com/max-cape/stripe-go-test/customercashbalancetransaction"
	"github.com/max-cape/stripe-go-test/customersession"
	"github.com/max-cape/stripe-go-test/dispute"
	entitlementsactiveentitlement "github.com/max-cape/stripe-go-test/entitlements/activeentitlement"
	entitlementsfeature "github.com/max-cape/stripe-go-test/entitlements/feature"
	"github.com/max-cape/stripe-go-test/ephemeralkey"
	"github.com/max-cape/stripe-go-test/event"
	"github.com/max-cape/stripe-go-test/feerefund"
	"github.com/max-cape/stripe-go-test/file"
	"github.com/max-cape/stripe-go-test/filelink"
	financialconnectionsaccount "github.com/max-cape/stripe-go-test/financialconnections/account"
	financialconnectionssession "github.com/max-cape/stripe-go-test/financialconnections/session"
	financialconnectionstransaction "github.com/max-cape/stripe-go-test/financialconnections/transaction"
	forwardingrequest "github.com/max-cape/stripe-go-test/forwarding/request"
	identityverificationreport "github.com/max-cape/stripe-go-test/identity/verificationreport"
	identityverificationsession "github.com/max-cape/stripe-go-test/identity/verificationsession"
	"github.com/max-cape/stripe-go-test/invoice"
	"github.com/max-cape/stripe-go-test/invoiceitem"
	"github.com/max-cape/stripe-go-test/invoicelineitem"
	"github.com/max-cape/stripe-go-test/invoicepayment"
	"github.com/max-cape/stripe-go-test/invoicerenderingtemplate"
	issuingauthorization "github.com/max-cape/stripe-go-test/issuing/authorization"
	issuingcard "github.com/max-cape/stripe-go-test/issuing/card"
	issuingcardholder "github.com/max-cape/stripe-go-test/issuing/cardholder"
	issuingdispute "github.com/max-cape/stripe-go-test/issuing/dispute"
	issuingpersonalizationdesign "github.com/max-cape/stripe-go-test/issuing/personalizationdesign"
	issuingphysicalbundle "github.com/max-cape/stripe-go-test/issuing/physicalbundle"
	issuingtoken "github.com/max-cape/stripe-go-test/issuing/token"
	issuingtransaction "github.com/max-cape/stripe-go-test/issuing/transaction"
	"github.com/max-cape/stripe-go-test/loginlink"
	"github.com/max-cape/stripe-go-test/mandate"
	"github.com/max-cape/stripe-go-test/oauth"
	"github.com/max-cape/stripe-go-test/paymentintent"
	"github.com/max-cape/stripe-go-test/paymentlink"
	"github.com/max-cape/stripe-go-test/paymentmethod"
	"github.com/max-cape/stripe-go-test/paymentmethodconfiguration"
	"github.com/max-cape/stripe-go-test/paymentmethoddomain"
	"github.com/max-cape/stripe-go-test/paymentsource"
	"github.com/max-cape/stripe-go-test/payout"
	"github.com/max-cape/stripe-go-test/person"
	"github.com/max-cape/stripe-go-test/plan"
	"github.com/max-cape/stripe-go-test/price"
	"github.com/max-cape/stripe-go-test/product"
	"github.com/max-cape/stripe-go-test/productfeature"
	"github.com/max-cape/stripe-go-test/promotioncode"
	"github.com/max-cape/stripe-go-test/quote"
	radarearlyfraudwarning "github.com/max-cape/stripe-go-test/radar/earlyfraudwarning"
	radarvaluelist "github.com/max-cape/stripe-go-test/radar/valuelist"
	radarvaluelistitem "github.com/max-cape/stripe-go-test/radar/valuelistitem"
	"github.com/max-cape/stripe-go-test/refund"
	reportingreportrun "github.com/max-cape/stripe-go-test/reporting/reportrun"
	reportingreporttype "github.com/max-cape/stripe-go-test/reporting/reporttype"
	"github.com/max-cape/stripe-go-test/review"
	"github.com/max-cape/stripe-go-test/setupattempt"
	"github.com/max-cape/stripe-go-test/setupintent"
	"github.com/max-cape/stripe-go-test/shippingrate"
	sigmascheduledqueryrun "github.com/max-cape/stripe-go-test/sigma/scheduledqueryrun"
	"github.com/max-cape/stripe-go-test/source"
	"github.com/max-cape/stripe-go-test/sourcetransaction"
	"github.com/max-cape/stripe-go-test/subscription"
	"github.com/max-cape/stripe-go-test/subscriptionitem"
	"github.com/max-cape/stripe-go-test/subscriptionschedule"
	taxcalculation "github.com/max-cape/stripe-go-test/tax/calculation"
	taxregistration "github.com/max-cape/stripe-go-test/tax/registration"
	taxsettings "github.com/max-cape/stripe-go-test/tax/settings"
	taxtransaction "github.com/max-cape/stripe-go-test/tax/transaction"
	"github.com/max-cape/stripe-go-test/taxcode"
	"github.com/max-cape/stripe-go-test/taxid"
	"github.com/max-cape/stripe-go-test/taxrate"
	terminalconfiguration "github.com/max-cape/stripe-go-test/terminal/configuration"
	terminalconnectiontoken "github.com/max-cape/stripe-go-test/terminal/connectiontoken"
	terminallocation "github.com/max-cape/stripe-go-test/terminal/location"
	terminalreader "github.com/max-cape/stripe-go-test/terminal/reader"
	testhelpersconfirmationtoken "github.com/max-cape/stripe-go-test/testhelpers/confirmationtoken"
	testhelperscustomer "github.com/max-cape/stripe-go-test/testhelpers/customer"
	testhelpersissuingauthorization "github.com/max-cape/stripe-go-test/testhelpers/issuing/authorization"
	testhelpersissuingcard "github.com/max-cape/stripe-go-test/testhelpers/issuing/card"
	testhelpersissuingpersonalizationdesign "github.com/max-cape/stripe-go-test/testhelpers/issuing/personalizationdesign"
	testhelpersissuingtransaction "github.com/max-cape/stripe-go-test/testhelpers/issuing/transaction"
	testhelpersrefund "github.com/max-cape/stripe-go-test/testhelpers/refund"
	testhelpersterminalreader "github.com/max-cape/stripe-go-test/testhelpers/terminal/reader"
	testhelperstestclock "github.com/max-cape/stripe-go-test/testhelpers/testclock"
	testhelperstreasuryinboundtransfer "github.com/max-cape/stripe-go-test/testhelpers/treasury/inboundtransfer"
	testhelperstreasuryoutboundpayment "github.com/max-cape/stripe-go-test/testhelpers/treasury/outboundpayment"
	testhelperstreasuryoutboundtransfer "github.com/max-cape/stripe-go-test/testhelpers/treasury/outboundtransfer"
	testhelperstreasuryreceivedcredit "github.com/max-cape/stripe-go-test/testhelpers/treasury/receivedcredit"
	testhelperstreasuryreceiveddebit "github.com/max-cape/stripe-go-test/testhelpers/treasury/receiveddebit"
	"github.com/max-cape/stripe-go-test/token"
	"github.com/max-cape/stripe-go-test/topup"
	"github.com/max-cape/stripe-go-test/transfer"
	"github.com/max-cape/stripe-go-test/transferreversal"
	treasurycreditreversal "github.com/max-cape/stripe-go-test/treasury/creditreversal"
	treasurydebitreversal "github.com/max-cape/stripe-go-test/treasury/debitreversal"
	treasuryfinancialaccount "github.com/max-cape/stripe-go-test/treasury/financialaccount"
	treasuryinboundtransfer "github.com/max-cape/stripe-go-test/treasury/inboundtransfer"
	treasuryoutboundpayment "github.com/max-cape/stripe-go-test/treasury/outboundpayment"
	treasuryoutboundtransfer "github.com/max-cape/stripe-go-test/treasury/outboundtransfer"
	treasuryreceivedcredit "github.com/max-cape/stripe-go-test/treasury/receivedcredit"
	treasuryreceiveddebit "github.com/max-cape/stripe-go-test/treasury/receiveddebit"
	treasurytransaction "github.com/max-cape/stripe-go-test/treasury/transaction"
	treasurytransactionentry "github.com/max-cape/stripe-go-test/treasury/transactionentry"
	v2billingmeterevent "github.com/max-cape/stripe-go-test/v2/billing/meterevent"
	v2billingmetereventadjustment "github.com/max-cape/stripe-go-test/v2/billing/metereventadjustment"
	v2billingmetereventsession "github.com/max-cape/stripe-go-test/v2/billing/metereventsession"
	v2billingmetereventstream "github.com/max-cape/stripe-go-test/v2/billing/metereventstream"
	v2coreevent "github.com/max-cape/stripe-go-test/v2/core/event"
	v2coreeventdestination "github.com/max-cape/stripe-go-test/v2/core/eventdestination"
	"github.com/max-cape/stripe-go-test/webhookendpoint"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// AccountLinks is the client used to invoke /v1/account_links APIs.
	AccountLinks *accountlink.Client
	// Accounts is the client used to invoke /v1/accounts APIs.
	Accounts *account.Client
	// AccountSessions is the client used to invoke /v1/account_sessions APIs.
	AccountSessions *accountsession.Client
	// ApplePayDomains is the client used to invoke /v1/apple_pay/domains APIs.
	ApplePayDomains *applepaydomain.Client
	// ApplicationFees is the client used to invoke /v1/application_fees APIs.
	ApplicationFees *applicationfee.Client
	// AppsSecrets is the client used to invoke /v1/apps/secrets APIs.
	AppsSecrets *appssecret.Client
	// Balance is the client used to invoke /v1/balance APIs.
	Balance *balance.Client
	// BalanceTransactions is the client used to invoke /v1/balance_transactions APIs.
	BalanceTransactions *balancetransaction.Client
	// BankAccounts is the client used to invoke /v1/accounts/{account}/external_accounts APIs.
	BankAccounts *bankaccount.Client
	// BillingAlerts is the client used to invoke /v1/billing/alerts APIs.
	BillingAlerts *billingalert.Client
	// BillingCreditBalanceSummary is the client used to invoke /v1/billing/credit_balance_summary APIs.
	BillingCreditBalanceSummary *billingcreditbalancesummary.Client
	// BillingCreditBalanceTransactions is the client used to invoke /v1/billing/credit_balance_transactions APIs.
	BillingCreditBalanceTransactions *billingcreditbalancetransaction.Client
	// BillingCreditGrants is the client used to invoke /v1/billing/credit_grants APIs.
	BillingCreditGrants *billingcreditgrant.Client
	// BillingMeterEventAdjustments is the client used to invoke /v1/billing/meter_event_adjustments APIs.
	BillingMeterEventAdjustments *billingmetereventadjustment.Client
	// BillingMeterEvents is the client used to invoke /v1/billing/meter_events APIs.
	BillingMeterEvents *billingmeterevent.Client
	// BillingMeterEventSummaries is the client used to invoke /v1/billing/meters/{id}/event_summaries APIs.
	BillingMeterEventSummaries *billingmetereventsummary.Client
	// BillingMeters is the client used to invoke /v1/billing/meters APIs.
	BillingMeters *billingmeter.Client
	// BillingPortalConfigurations is the client used to invoke /v1/billing_portal/configurations APIs.
	BillingPortalConfigurations *billingportalconfiguration.Client
	// BillingPortalSessions is the client used to invoke /v1/billing_portal/sessions APIs.
	BillingPortalSessions *billingportalsession.Client
	// Capabilities is the client used to invoke /v1/accounts/{account}/capabilities APIs.
	Capabilities *capability.Client
	// Cards is the client used to invoke /v1/accounts/{account}/external_accounts APIs.
	Cards *card.Client
	// CashBalances is the client used to invoke /v1/customers/{customer}/cash_balance APIs.
	CashBalances *cashbalance.Client
	// Charges is the client used to invoke /v1/charges APIs.
	Charges *charge.Client
	// CheckoutSessions is the client used to invoke /v1/checkout/sessions APIs.
	CheckoutSessions *checkoutsession.Client
	// ClimateOrders is the client used to invoke /v1/climate/orders APIs.
	ClimateOrders *climateorder.Client
	// ClimateProducts is the client used to invoke /v1/climate/products APIs.
	ClimateProducts *climateproduct.Client
	// ClimateSuppliers is the client used to invoke /v1/climate/suppliers APIs.
	ClimateSuppliers *climatesupplier.Client
	// ConfirmationTokens is the client used to invoke /v1/confirmation_tokens APIs.
	ConfirmationTokens *confirmationtoken.Client
	// CountrySpecs is the client used to invoke /v1/country_specs APIs.
	CountrySpecs *countryspec.Client
	// Coupons is the client used to invoke /v1/coupons APIs.
	Coupons *coupon.Client
	// CreditNotes is the client used to invoke /v1/credit_notes APIs.
	CreditNotes *creditnote.Client
	// CustomerBalanceTransactions is the client used to invoke /v1/customers/{customer}/balance_transactions APIs.
	CustomerBalanceTransactions *customerbalancetransaction.Client
	// CustomerCashBalanceTransactions is the client used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
	CustomerCashBalanceTransactions *customercashbalancetransaction.Client
	// Customers is the client used to invoke /v1/customers APIs.
	Customers *customer.Client
	// CustomerSessions is the client used to invoke /v1/customer_sessions APIs.
	CustomerSessions *customersession.Client
	// Disputes is the client used to invoke /v1/disputes APIs.
	Disputes *dispute.Client
	// EntitlementsActiveEntitlements is the client used to invoke /v1/entitlements/active_entitlements APIs.
	EntitlementsActiveEntitlements *entitlementsactiveentitlement.Client
	// EntitlementsFeatures is the client used to invoke /v1/entitlements/features APIs.
	EntitlementsFeatures *entitlementsfeature.Client
	// EphemeralKeys is the client used to invoke /v1/ephemeral_keys APIs.
	EphemeralKeys *ephemeralkey.Client
	// Events is the client used to invoke /v1/events APIs.
	Events *event.Client
	// FeeRefunds is the client used to invoke /v1/application_fees/{id}/refunds APIs.
	FeeRefunds *feerefund.Client
	// FileLinks is the client used to invoke /v1/file_links APIs.
	FileLinks *filelink.Client
	// Files is the client used to invoke /v1/files APIs.
	Files *file.Client
	// FinancialConnectionsAccounts is the client used to invoke /v1/financial_connections/accounts APIs.
	FinancialConnectionsAccounts *financialconnectionsaccount.Client
	// FinancialConnectionsSessions is the client used to invoke /v1/financial_connections/sessions APIs.
	FinancialConnectionsSessions *financialconnectionssession.Client
	// FinancialConnectionsTransactions is the client used to invoke /v1/financial_connections/transactions APIs.
	FinancialConnectionsTransactions *financialconnectionstransaction.Client
	// ForwardingRequests is the client used to invoke /v1/forwarding/requests APIs.
	ForwardingRequests *forwardingrequest.Client
	// IdentityVerificationReports is the client used to invoke /v1/identity/verification_reports APIs.
	IdentityVerificationReports *identityverificationreport.Client
	// IdentityVerificationSessions is the client used to invoke /v1/identity/verification_sessions APIs.
	IdentityVerificationSessions *identityverificationsession.Client
	// InvoiceItems is the client used to invoke /v1/invoiceitems APIs.
	InvoiceItems *invoiceitem.Client
	// InvoiceLineItems is the client used to invoke /v1/invoices/{invoice}/lines APIs.
	InvoiceLineItems *invoicelineitem.Client
	// InvoicePayments is the client used to invoke /v1/invoice_payments APIs.
	InvoicePayments *invoicepayment.Client
	// InvoiceRenderingTemplates is the client used to invoke /v1/invoice_rendering_templates APIs.
	InvoiceRenderingTemplates *invoicerenderingtemplate.Client
	// Invoices is the client used to invoke /v1/invoices APIs.
	Invoices *invoice.Client
	// IssuingAuthorizations is the client used to invoke /v1/issuing/authorizations APIs.
	IssuingAuthorizations *issuingauthorization.Client
	// IssuingCardholders is the client used to invoke /v1/issuing/cardholders APIs.
	IssuingCardholders *issuingcardholder.Client
	// IssuingCards is the client used to invoke /v1/issuing/cards APIs.
	IssuingCards *issuingcard.Client
	// IssuingDisputes is the client used to invoke /v1/issuing/disputes APIs.
	IssuingDisputes *issuingdispute.Client
	// IssuingPersonalizationDesigns is the client used to invoke /v1/issuing/personalization_designs APIs.
	IssuingPersonalizationDesigns *issuingpersonalizationdesign.Client
	// IssuingPhysicalBundles is the client used to invoke /v1/issuing/physical_bundles APIs.
	IssuingPhysicalBundles *issuingphysicalbundle.Client
	// IssuingTokens is the client used to invoke /v1/issuing/tokens APIs.
	IssuingTokens *issuingtoken.Client
	// IssuingTransactions is the client used to invoke /v1/issuing/transactions APIs.
	IssuingTransactions *issuingtransaction.Client
	// LoginLinks is the client used to invoke /v1/accounts/{account}/login_links APIs.
	LoginLinks *loginlink.Client
	// Mandates is the client used to invoke /v1/mandates APIs.
	Mandates *mandate.Client
	// OAuth is the client used to invoke /oauth APIs
	OAuth *oauth.Client
	// PaymentIntents is the client used to invoke /v1/payment_intents APIs.
	PaymentIntents *paymentintent.Client
	// PaymentLinks is the client used to invoke /v1/payment_links APIs.
	PaymentLinks *paymentlink.Client
	// PaymentMethodConfigurations is the client used to invoke /v1/payment_method_configurations APIs.
	PaymentMethodConfigurations *paymentmethodconfiguration.Client
	// PaymentMethodDomains is the client used to invoke /v1/payment_method_domains APIs.
	PaymentMethodDomains *paymentmethoddomain.Client
	// PaymentMethods is the client used to invoke /v1/payment_methods APIs.
	PaymentMethods *paymentmethod.Client
	// PaymentSources is the client used to invoke /v1/customers/{customer}/sources APIs.
	PaymentSources *paymentsource.Client
	// Payouts is the client used to invoke /v1/payouts APIs.
	Payouts *payout.Client
	// Persons is the client used to invoke /v1/accounts/{account}/persons APIs.
	Persons *person.Client
	// Plans is the client used to invoke /v1/plans APIs.
	Plans *plan.Client
	// Prices is the client used to invoke /v1/prices APIs.
	Prices *price.Client
	// ProductFeatures is the client used to invoke /v1/products/{product}/features APIs.
	ProductFeatures *productfeature.Client
	// Products is the client used to invoke /v1/products APIs.
	Products *product.Client
	// PromotionCodes is the client used to invoke /v1/promotion_codes APIs.
	PromotionCodes *promotioncode.Client
	// Quotes is the client used to invoke /v1/quotes APIs.
	Quotes *quote.Client
	// RadarEarlyFraudWarnings is the client used to invoke /v1/radar/early_fraud_warnings APIs.
	RadarEarlyFraudWarnings *radarearlyfraudwarning.Client
	// RadarValueListItems is the client used to invoke /v1/radar/value_list_items APIs.
	RadarValueListItems *radarvaluelistitem.Client
	// RadarValueLists is the client used to invoke /v1/radar/value_lists APIs.
	RadarValueLists *radarvaluelist.Client
	// Refunds is the client used to invoke /v1/refunds APIs.
	Refunds *refund.Client
	// ReportingReportRuns is the client used to invoke /v1/reporting/report_runs APIs.
	ReportingReportRuns *reportingreportrun.Client
	// ReportingReportTypes is the client used to invoke /v1/reporting/report_types APIs.
	ReportingReportTypes *reportingreporttype.Client
	// Reviews is the client used to invoke /v1/reviews APIs.
	Reviews *review.Client
	// SetupAttempts is the client used to invoke /v1/setup_attempts APIs.
	SetupAttempts *setupattempt.Client
	// SetupIntents is the client used to invoke /v1/setup_intents APIs.
	SetupIntents *setupintent.Client
	// ShippingRates is the client used to invoke /v1/shipping_rates APIs.
	ShippingRates *shippingrate.Client
	// SigmaScheduledQueryRuns is the client used to invoke /v1/sigma/scheduled_query_runs APIs.
	SigmaScheduledQueryRuns *sigmascheduledqueryrun.Client
	// Sources is the client used to invoke /v1/sources APIs.
	Sources *source.Client
	// SourceTransactions is the client used to invoke /v1/sources/{source}/source_transactions APIs.
	SourceTransactions *sourcetransaction.Client
	// SubscriptionItems is the client used to invoke /v1/subscription_items APIs.
	SubscriptionItems *subscriptionitem.Client
	// Subscriptions is the client used to invoke /v1/subscriptions APIs.
	Subscriptions *subscription.Client
	// SubscriptionSchedules is the client used to invoke /v1/subscription_schedules APIs.
	SubscriptionSchedules *subscriptionschedule.Client
	// TaxCalculations is the client used to invoke /v1/tax/calculations APIs.
	TaxCalculations *taxcalculation.Client
	// TaxCodes is the client used to invoke /v1/tax_codes APIs.
	TaxCodes *taxcode.Client
	// TaxIDs is the client used to invoke /v1/tax_ids APIs.
	TaxIDs *taxid.Client
	// TaxRates is the client used to invoke /v1/tax_rates APIs.
	TaxRates *taxrate.Client
	// TaxRegistrations is the client used to invoke /v1/tax/registrations APIs.
	TaxRegistrations *taxregistration.Client
	// TaxSettings is the client used to invoke /v1/tax/settings APIs.
	TaxSettings *taxsettings.Client
	// TaxTransactions is the client used to invoke /v1/tax/transactions APIs.
	TaxTransactions *taxtransaction.Client
	// TerminalConfigurations is the client used to invoke /v1/terminal/configurations APIs.
	TerminalConfigurations *terminalconfiguration.Client
	// TerminalConnectionTokens is the client used to invoke /v1/terminal/connection_tokens APIs.
	TerminalConnectionTokens *terminalconnectiontoken.Client
	// TerminalLocations is the client used to invoke /v1/terminal/locations APIs.
	TerminalLocations *terminallocation.Client
	// TerminalReaders is the client used to invoke /v1/terminal/readers APIs.
	TerminalReaders *terminalreader.Client
	// TestHelpersConfirmationTokens is the client used to invoke /v1/confirmation_tokens APIs.
	TestHelpersConfirmationTokens *testhelpersconfirmationtoken.Client
	// TestHelpersCustomers is the client used to invoke /v1/customers APIs.
	TestHelpersCustomers *testhelperscustomer.Client
	// TestHelpersIssuingAuthorizations is the client used to invoke /v1/issuing/authorizations APIs.
	TestHelpersIssuingAuthorizations *testhelpersissuingauthorization.Client
	// TestHelpersIssuingCards is the client used to invoke /v1/issuing/cards APIs.
	TestHelpersIssuingCards *testhelpersissuingcard.Client
	// TestHelpersIssuingPersonalizationDesigns is the client used to invoke /v1/issuing/personalization_designs APIs.
	TestHelpersIssuingPersonalizationDesigns *testhelpersissuingpersonalizationdesign.Client
	// TestHelpersIssuingTransactions is the client used to invoke /v1/issuing/transactions APIs.
	TestHelpersIssuingTransactions *testhelpersissuingtransaction.Client
	// TestHelpersRefunds is the client used to invoke /v1/refunds APIs.
	TestHelpersRefunds *testhelpersrefund.Client
	// TestHelpersTerminalReaders is the client used to invoke /v1/terminal/readers APIs.
	TestHelpersTerminalReaders *testhelpersterminalreader.Client
	// TestHelpersTestClocks is the client used to invoke /v1/test_helpers/test_clocks APIs.
	TestHelpersTestClocks *testhelperstestclock.Client
	// TestHelpersTreasuryInboundTransfers is the client used to invoke /v1/treasury/inbound_transfers APIs.
	TestHelpersTreasuryInboundTransfers *testhelperstreasuryinboundtransfer.Client
	// TestHelpersTreasuryOutboundPayments is the client used to invoke /v1/treasury/outbound_payments APIs.
	TestHelpersTreasuryOutboundPayments *testhelperstreasuryoutboundpayment.Client
	// TestHelpersTreasuryOutboundTransfers is the client used to invoke /v1/treasury/outbound_transfers APIs.
	TestHelpersTreasuryOutboundTransfers *testhelperstreasuryoutboundtransfer.Client
	// TestHelpersTreasuryReceivedCredits is the client used to invoke /v1/treasury/received_credits APIs.
	TestHelpersTreasuryReceivedCredits *testhelperstreasuryreceivedcredit.Client
	// TestHelpersTreasuryReceivedDebits is the client used to invoke /v1/treasury/received_debits APIs.
	TestHelpersTreasuryReceivedDebits *testhelperstreasuryreceiveddebit.Client
	// Tokens is the client used to invoke /v1/tokens APIs.
	Tokens *token.Client
	// Topups is the client used to invoke /v1/topups APIs.
	Topups *topup.Client
	// TransferReversals is the client used to invoke /v1/transfers/{id}/reversals APIs.
	TransferReversals *transferreversal.Client
	// Transfers is the client used to invoke /v1/transfers APIs.
	Transfers *transfer.Client
	// TreasuryCreditReversals is the client used to invoke /v1/treasury/credit_reversals APIs.
	TreasuryCreditReversals *treasurycreditreversal.Client
	// TreasuryDebitReversals is the client used to invoke /v1/treasury/debit_reversals APIs.
	TreasuryDebitReversals *treasurydebitreversal.Client
	// TreasuryFinancialAccounts is the client used to invoke /v1/treasury/financial_accounts APIs.
	TreasuryFinancialAccounts *treasuryfinancialaccount.Client
	// TreasuryInboundTransfers is the client used to invoke /v1/treasury/inbound_transfers APIs.
	TreasuryInboundTransfers *treasuryinboundtransfer.Client
	// TreasuryOutboundPayments is the client used to invoke /v1/treasury/outbound_payments APIs.
	TreasuryOutboundPayments *treasuryoutboundpayment.Client
	// TreasuryOutboundTransfers is the client used to invoke /v1/treasury/outbound_transfers APIs.
	TreasuryOutboundTransfers *treasuryoutboundtransfer.Client
	// TreasuryReceivedCredits is the client used to invoke /v1/treasury/received_credits APIs.
	TreasuryReceivedCredits *treasuryreceivedcredit.Client
	// TreasuryReceivedDebits is the client used to invoke /v1/treasury/received_debits APIs.
	TreasuryReceivedDebits *treasuryreceiveddebit.Client
	// TreasuryTransactionEntries is the client used to invoke /v1/treasury/transaction_entries APIs.
	TreasuryTransactionEntries *treasurytransactionentry.Client
	// TreasuryTransactions is the client used to invoke /v1/treasury/transactions APIs.
	TreasuryTransactions *treasurytransaction.Client
	// V2BillingMeterEventAdjustments is the client used to invoke /v2/billing/meter_event_adjustments APIs.
	V2BillingMeterEventAdjustments *v2billingmetereventadjustment.Client
	// V2BillingMeterEvents is the client used to invoke /v2/billing/meter_events APIs.
	V2BillingMeterEvents *v2billingmeterevent.Client
	// V2BillingMeterEventSessions is the client used to invoke /v2/billing/meter_event_session APIs.
	V2BillingMeterEventSessions *v2billingmetereventsession.Client
	// V2BillingMeterEventStreams is the client used to invoke /v2/billing/meter_event_stream APIs.
	V2BillingMeterEventStreams *v2billingmetereventstream.Client
	// V2CoreEventDestinations is the client used to invoke /v2/core/event_destinations APIs.
	V2CoreEventDestinations *v2coreeventdestination.Client
	// V2CoreEvents is the client used to invoke /v2/core/events APIs.
	V2CoreEvents *v2coreevent.Client
	// WebhookEndpoints is the client used to invoke /v1/webhook_endpoints APIs.
	WebhookEndpoints *webhookendpoint.Client
}

func (a *API) Init(key string, backends *stripe.Backends) {
	usage := []string{"stripe_client"}
	if backends == nil {
		backends = &stripe.Backends{
			API:         &stripe.UsageBackend{B: stripe.GetBackend(stripe.APIBackend), Usage: usage},
			Connect:     &stripe.UsageBackend{B: stripe.GetBackend(stripe.ConnectBackend), Usage: usage},
			Uploads:     &stripe.UsageBackend{B: stripe.GetBackend(stripe.UploadsBackend), Usage: usage},
			MeterEvents: &stripe.UsageBackend{B: stripe.GetBackend(stripe.MeterEventsBackend), Usage: usage},
		}
	}

	a.AccountLinks = &accountlink.Client{B: backends.API, Key: key}
	a.Accounts = &account.Client{B: backends.API, Key: key}
	a.AccountSessions = &accountsession.Client{B: backends.API, Key: key}
	a.ApplePayDomains = &applepaydomain.Client{B: backends.API, Key: key}
	a.ApplicationFees = &applicationfee.Client{B: backends.API, Key: key}
	a.AppsSecrets = &appssecret.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.BalanceTransactions = &balancetransaction.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.BillingAlerts = &billingalert.Client{B: backends.API, Key: key}
	a.BillingCreditBalanceSummary = &billingcreditbalancesummary.Client{B: backends.API, Key: key}
	a.BillingCreditBalanceTransactions = &billingcreditbalancetransaction.Client{B: backends.API, Key: key}
	a.BillingCreditGrants = &billingcreditgrant.Client{B: backends.API, Key: key}
	a.BillingMeterEventAdjustments = &billingmetereventadjustment.Client{B: backends.API, Key: key}
	a.BillingMeterEvents = &billingmeterevent.Client{B: backends.API, Key: key}
	a.BillingMeterEventSummaries = &billingmetereventsummary.Client{B: backends.API, Key: key}
	a.BillingMeters = &billingmeter.Client{B: backends.API, Key: key}
	a.BillingPortalConfigurations = &billingportalconfiguration.Client{B: backends.API, Key: key}
	a.BillingPortalSessions = &billingportalsession.Client{B: backends.API, Key: key}
	a.Capabilities = &capability.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.CashBalances = &cashbalance.Client{B: backends.API, Key: key}
	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.CheckoutSessions = &checkoutsession.Client{B: backends.API, Key: key}
	a.ClimateOrders = &climateorder.Client{B: backends.API, Key: key}
	a.ClimateProducts = &climateproduct.Client{B: backends.API, Key: key}
	a.ClimateSuppliers = &climatesupplier.Client{B: backends.API, Key: key}
	a.ConfirmationTokens = &confirmationtoken.Client{B: backends.API, Key: key}
	a.CountrySpecs = &countryspec.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.CreditNotes = &creditnote.Client{B: backends.API, Key: key}
	a.CustomerBalanceTransactions = &customerbalancetransaction.Client{B: backends.API, Key: key}
	a.CustomerCashBalanceTransactions = &customercashbalancetransaction.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.CustomerSessions = &customersession.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.EntitlementsActiveEntitlements = &entitlementsactiveentitlement.Client{B: backends.API, Key: key}
	a.EntitlementsFeatures = &entitlementsfeature.Client{B: backends.API, Key: key}
	a.EphemeralKeys = &ephemeralkey.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.FileLinks = &filelink.Client{B: backends.API, Key: key}
	a.Files = &file.Client{B: backends.API, BUploads: backends.Uploads, Key: key}
	a.FinancialConnectionsAccounts = &financialconnectionsaccount.Client{B: backends.API, Key: key}
	a.FinancialConnectionsSessions = &financialconnectionssession.Client{B: backends.API, Key: key}
	a.FinancialConnectionsTransactions = &financialconnectionstransaction.Client{B: backends.API, Key: key}
	a.ForwardingRequests = &forwardingrequest.Client{B: backends.API, Key: key}
	a.IdentityVerificationReports = &identityverificationreport.Client{B: backends.API, Key: key}
	a.IdentityVerificationSessions = &identityverificationsession.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.InvoiceLineItems = &invoicelineitem.Client{B: backends.API, Key: key}
	a.InvoicePayments = &invoicepayment.Client{B: backends.API, Key: key}
	a.InvoiceRenderingTemplates = &invoicerenderingtemplate.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.IssuingAuthorizations = &issuingauthorization.Client{B: backends.API, Key: key}
	a.IssuingCardholders = &issuingcardholder.Client{B: backends.API, Key: key}
	a.IssuingCards = &issuingcard.Client{B: backends.API, Key: key}
	a.IssuingDisputes = &issuingdispute.Client{B: backends.API, Key: key}
	a.IssuingPersonalizationDesigns = &issuingpersonalizationdesign.Client{B: backends.API, Key: key}
	a.IssuingPhysicalBundles = &issuingphysicalbundle.Client{B: backends.API, Key: key}
	a.IssuingTokens = &issuingtoken.Client{B: backends.API, Key: key}
	a.IssuingTransactions = &issuingtransaction.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Mandates = &mandate.Client{B: backends.API, Key: key}
	a.OAuth = &oauth.Client{B: backends.Connect, Key: key}
	a.PaymentIntents = &paymentintent.Client{B: backends.API, Key: key}
	a.PaymentLinks = &paymentlink.Client{B: backends.API, Key: key}
	a.PaymentMethodConfigurations = &paymentmethodconfiguration.Client{B: backends.API, Key: key}
	a.PaymentMethodDomains = &paymentmethoddomain.Client{B: backends.API, Key: key}
	a.PaymentMethods = &paymentmethod.Client{B: backends.API, Key: key}
	a.PaymentSources = &paymentsource.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Persons = &person.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Prices = &price.Client{B: backends.API, Key: key}
	a.ProductFeatures = &productfeature.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.PromotionCodes = &promotioncode.Client{B: backends.API, Key: key}
	a.Quotes = &quote.Client{B: backends.API, BUploads: backends.Uploads, Key: key}
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
	a.TaxRegistrations = &taxregistration.Client{B: backends.API, Key: key}
	a.TaxSettings = &taxsettings.Client{B: backends.API, Key: key}
	a.TaxTransactions = &taxtransaction.Client{B: backends.API, Key: key}
	a.TerminalConfigurations = &terminalconfiguration.Client{B: backends.API, Key: key}
	a.TerminalConnectionTokens = &terminalconnectiontoken.Client{B: backends.API, Key: key}
	a.TerminalLocations = &terminallocation.Client{B: backends.API, Key: key}
	a.TerminalReaders = &terminalreader.Client{B: backends.API, Key: key}
	a.TestHelpersConfirmationTokens = &testhelpersconfirmationtoken.Client{B: backends.API, Key: key}
	a.TestHelpersCustomers = &testhelperscustomer.Client{B: backends.API, Key: key}
	a.TestHelpersIssuingAuthorizations = &testhelpersissuingauthorization.Client{B: backends.API, Key: key}
	a.TestHelpersIssuingCards = &testhelpersissuingcard.Client{B: backends.API, Key: key}
	a.TestHelpersIssuingPersonalizationDesigns = &testhelpersissuingpersonalizationdesign.Client{B: backends.API, Key: key}
	a.TestHelpersIssuingTransactions = &testhelpersissuingtransaction.Client{B: backends.API, Key: key}
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
	a.V2BillingMeterEventAdjustments = &v2billingmetereventadjustment.Client{B: backends.API, Key: key}
	a.V2BillingMeterEvents = &v2billingmeterevent.Client{B: backends.API, Key: key}
	a.V2BillingMeterEventSessions = &v2billingmetereventsession.Client{B: backends.API, Key: key}
	a.V2BillingMeterEventStreams = &v2billingmetereventstream.Client{BMeterEvents: backends.MeterEvents, Key: key}
	a.V2CoreEventDestinations = &v2coreeventdestination.Client{B: backends.API, Key: key}
	a.V2CoreEvents = &v2coreevent.Client{B: backends.API, Key: key}
	a.WebhookEndpoints = &webhookendpoint.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *stripe.Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
