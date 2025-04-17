//
//
// File generated from our OpenAPI spec
//
//

// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/account"
	"github.com/stripe/stripe-go/v82/accountlink"
	"github.com/stripe/stripe-go/v82/accountnotice"
	"github.com/stripe/stripe-go/v82/accountsession"
	"github.com/stripe/stripe-go/v82/applepaydomain"
	"github.com/stripe/stripe-go/v82/applicationfee"
	appssecret "github.com/stripe/stripe-go/v82/apps/secret"
	"github.com/stripe/stripe-go/v82/balance"
	"github.com/stripe/stripe-go/v82/balancesettings"
	"github.com/stripe/stripe-go/v82/balancetransaction"
	"github.com/stripe/stripe-go/v82/bankaccount"
	billingalert "github.com/stripe/stripe-go/v82/billing/alert"
	billingcreditbalancesummary "github.com/stripe/stripe-go/v82/billing/creditbalancesummary"
	billingcreditbalancetransaction "github.com/stripe/stripe-go/v82/billing/creditbalancetransaction"
	billingcreditgrant "github.com/stripe/stripe-go/v82/billing/creditgrant"
	billingmeter "github.com/stripe/stripe-go/v82/billing/meter"
	billingmeterevent "github.com/stripe/stripe-go/v82/billing/meterevent"
	billingmetereventadjustment "github.com/stripe/stripe-go/v82/billing/metereventadjustment"
	billingmetereventsummary "github.com/stripe/stripe-go/v82/billing/metereventsummary"
	billingportalconfiguration "github.com/stripe/stripe-go/v82/billingportal/configuration"
	billingportalsession "github.com/stripe/stripe-go/v82/billingportal/session"
	"github.com/stripe/stripe-go/v82/capability"
	capitalfinancingoffer "github.com/stripe/stripe-go/v82/capital/financingoffer"
	capitalfinancingsummary "github.com/stripe/stripe-go/v82/capital/financingsummary"
	capitalfinancingtransaction "github.com/stripe/stripe-go/v82/capital/financingtransaction"
	"github.com/stripe/stripe-go/v82/card"
	"github.com/stripe/stripe-go/v82/cashbalance"
	"github.com/stripe/stripe-go/v82/charge"
	checkoutsession "github.com/stripe/stripe-go/v82/checkout/session"
	climateorder "github.com/stripe/stripe-go/v82/climate/order"
	climateproduct "github.com/stripe/stripe-go/v82/climate/product"
	climatesupplier "github.com/stripe/stripe-go/v82/climate/supplier"
	"github.com/stripe/stripe-go/v82/confirmationtoken"
	"github.com/stripe/stripe-go/v82/countryspec"
	"github.com/stripe/stripe-go/v82/coupon"
	"github.com/stripe/stripe-go/v82/creditnote"
	"github.com/stripe/stripe-go/v82/customer"
	"github.com/stripe/stripe-go/v82/customerbalancetransaction"
	"github.com/stripe/stripe-go/v82/customercashbalancetransaction"
	"github.com/stripe/stripe-go/v82/customersession"
	"github.com/stripe/stripe-go/v82/dispute"
	entitlementsactiveentitlement "github.com/stripe/stripe-go/v82/entitlements/activeentitlement"
	entitlementsfeature "github.com/stripe/stripe-go/v82/entitlements/feature"
	"github.com/stripe/stripe-go/v82/ephemeralkey"
	"github.com/stripe/stripe-go/v82/event"
	"github.com/stripe/stripe-go/v82/externalaccount"
	"github.com/stripe/stripe-go/v82/feerefund"
	"github.com/stripe/stripe-go/v82/file"
	"github.com/stripe/stripe-go/v82/filelink"
	financialconnectionsaccount "github.com/stripe/stripe-go/v82/financialconnections/account"
	financialconnectionsaccountinferredbalance "github.com/stripe/stripe-go/v82/financialconnections/accountinferredbalance"
	financialconnectionsinstitution "github.com/stripe/stripe-go/v82/financialconnections/institution"
	financialconnectionssession "github.com/stripe/stripe-go/v82/financialconnections/session"
	financialconnectionstransaction "github.com/stripe/stripe-go/v82/financialconnections/transaction"
	forwardingrequest "github.com/stripe/stripe-go/v82/forwarding/request"
	"github.com/stripe/stripe-go/v82/fxquote"
	giftcardscard "github.com/stripe/stripe-go/v82/giftcards/card"
	giftcardstransaction "github.com/stripe/stripe-go/v82/giftcards/transaction"
	identityverificationreport "github.com/stripe/stripe-go/v82/identity/verificationreport"
	identityverificationsession "github.com/stripe/stripe-go/v82/identity/verificationsession"
	"github.com/stripe/stripe-go/v82/invoice"
	"github.com/stripe/stripe-go/v82/invoiceitem"
	"github.com/stripe/stripe-go/v82/invoicelineitem"
	"github.com/stripe/stripe-go/v82/invoicerenderingtemplate"
	issuingauthorization "github.com/stripe/stripe-go/v82/issuing/authorization"
	issuingcard "github.com/stripe/stripe-go/v82/issuing/card"
	issuingcardholder "github.com/stripe/stripe-go/v82/issuing/cardholder"
	issuingcreditunderwritingrecord "github.com/stripe/stripe-go/v82/issuing/creditunderwritingrecord"
	issuingdispute "github.com/stripe/stripe-go/v82/issuing/dispute"
	issuingdisputesettlementdetail "github.com/stripe/stripe-go/v82/issuing/disputesettlementdetail"
	issuingfraudliabilitydebit "github.com/stripe/stripe-go/v82/issuing/fraudliabilitydebit"
	issuingpersonalizationdesign "github.com/stripe/stripe-go/v82/issuing/personalizationdesign"
	issuingphysicalbundle "github.com/stripe/stripe-go/v82/issuing/physicalbundle"
	issuingtoken "github.com/stripe/stripe-go/v82/issuing/token"
	issuingtransaction "github.com/stripe/stripe-go/v82/issuing/transaction"
	"github.com/stripe/stripe-go/v82/loginlink"
	"github.com/stripe/stripe-go/v82/mandate"
	"github.com/stripe/stripe-go/v82/margin"
	"github.com/stripe/stripe-go/v82/oauth"
	"github.com/stripe/stripe-go/v82/order"
	"github.com/stripe/stripe-go/v82/paymentattemptrecord"
	"github.com/stripe/stripe-go/v82/paymentintent"
	"github.com/stripe/stripe-go/v82/paymentintentamountdetailslineitem"
	"github.com/stripe/stripe-go/v82/paymentlink"
	"github.com/stripe/stripe-go/v82/paymentmethod"
	"github.com/stripe/stripe-go/v82/paymentmethodconfiguration"
	"github.com/stripe/stripe-go/v82/paymentmethoddomain"
	"github.com/stripe/stripe-go/v82/paymentrecord"
	"github.com/stripe/stripe-go/v82/paymentsource"
	"github.com/stripe/stripe-go/v82/payout"
	"github.com/stripe/stripe-go/v82/person"
	"github.com/stripe/stripe-go/v82/plan"
	"github.com/stripe/stripe-go/v82/price"
	privacyredactionjob "github.com/stripe/stripe-go/v82/privacy/redactionjob"
	privacyredactionjobvalidationerror "github.com/stripe/stripe-go/v82/privacy/redactionjobvalidationerror"
	"github.com/stripe/stripe-go/v82/product"
	"github.com/stripe/stripe-go/v82/productfeature"
	"github.com/stripe/stripe-go/v82/promotioncode"
	"github.com/stripe/stripe-go/v82/quote"
	"github.com/stripe/stripe-go/v82/quotepreviewinvoice"
	"github.com/stripe/stripe-go/v82/quotepreviewsubscriptionschedule"
	radarearlyfraudwarning "github.com/stripe/stripe-go/v82/radar/earlyfraudwarning"
	radarvaluelist "github.com/stripe/stripe-go/v82/radar/valuelist"
	radarvaluelistitem "github.com/stripe/stripe-go/v82/radar/valuelistitem"
	"github.com/stripe/stripe-go/v82/refund"
	reportingreportrun "github.com/stripe/stripe-go/v82/reporting/reportrun"
	reportingreporttype "github.com/stripe/stripe-go/v82/reporting/reporttype"
	"github.com/stripe/stripe-go/v82/review"
	"github.com/stripe/stripe-go/v82/setupattempt"
	"github.com/stripe/stripe-go/v82/setupintent"
	"github.com/stripe/stripe-go/v82/shippingrate"
	sigmascheduledqueryrun "github.com/stripe/stripe-go/v82/sigma/scheduledqueryrun"
	"github.com/stripe/stripe-go/v82/source"
	"github.com/stripe/stripe-go/v82/sourcetransaction"
	"github.com/stripe/stripe-go/v82/subscription"
	"github.com/stripe/stripe-go/v82/subscriptionitem"
	"github.com/stripe/stripe-go/v82/subscriptionschedule"
	taxassociation "github.com/stripe/stripe-go/v82/tax/association"
	taxcalculation "github.com/stripe/stripe-go/v82/tax/calculation"
	taxform "github.com/stripe/stripe-go/v82/tax/form"
	taxregistration "github.com/stripe/stripe-go/v82/tax/registration"
	taxsettings "github.com/stripe/stripe-go/v82/tax/settings"
	taxtransaction "github.com/stripe/stripe-go/v82/tax/transaction"
	"github.com/stripe/stripe-go/v82/taxcode"
	"github.com/stripe/stripe-go/v82/taxid"
	"github.com/stripe/stripe-go/v82/taxrate"
	terminalconfiguration "github.com/stripe/stripe-go/v82/terminal/configuration"
	terminalconnectiontoken "github.com/stripe/stripe-go/v82/terminal/connectiontoken"
	terminallocation "github.com/stripe/stripe-go/v82/terminal/location"
	terminalreader "github.com/stripe/stripe-go/v82/terminal/reader"
	terminalreadercollecteddata "github.com/stripe/stripe-go/v82/terminal/readercollecteddata"
	testhelpersconfirmationtoken "github.com/stripe/stripe-go/v82/testhelpers/confirmationtoken"
	testhelperscustomer "github.com/stripe/stripe-go/v82/testhelpers/customer"
	testhelpersissuingauthorization "github.com/stripe/stripe-go/v82/testhelpers/issuing/authorization"
	testhelpersissuingcard "github.com/stripe/stripe-go/v82/testhelpers/issuing/card"
	testhelpersissuingpersonalizationdesign "github.com/stripe/stripe-go/v82/testhelpers/issuing/personalizationdesign"
	testhelpersissuingtransaction "github.com/stripe/stripe-go/v82/testhelpers/issuing/transaction"
	testhelpersrefund "github.com/stripe/stripe-go/v82/testhelpers/refund"
	testhelpersterminalreader "github.com/stripe/stripe-go/v82/testhelpers/terminal/reader"
	testhelperstestclock "github.com/stripe/stripe-go/v82/testhelpers/testclock"
	testhelperstreasuryinboundtransfer "github.com/stripe/stripe-go/v82/testhelpers/treasury/inboundtransfer"
	testhelperstreasuryoutboundpayment "github.com/stripe/stripe-go/v82/testhelpers/treasury/outboundpayment"
	testhelperstreasuryoutboundtransfer "github.com/stripe/stripe-go/v82/testhelpers/treasury/outboundtransfer"
	testhelperstreasuryreceivedcredit "github.com/stripe/stripe-go/v82/testhelpers/treasury/receivedcredit"
	testhelperstreasuryreceiveddebit "github.com/stripe/stripe-go/v82/testhelpers/treasury/receiveddebit"
	"github.com/stripe/stripe-go/v82/token"
	"github.com/stripe/stripe-go/v82/topup"
	"github.com/stripe/stripe-go/v82/transfer"
	"github.com/stripe/stripe-go/v82/transferreversal"
	treasurycreditreversal "github.com/stripe/stripe-go/v82/treasury/creditreversal"
	treasurydebitreversal "github.com/stripe/stripe-go/v82/treasury/debitreversal"
	treasuryfinancialaccount "github.com/stripe/stripe-go/v82/treasury/financialaccount"
	treasuryinboundtransfer "github.com/stripe/stripe-go/v82/treasury/inboundtransfer"
	treasuryoutboundpayment "github.com/stripe/stripe-go/v82/treasury/outboundpayment"
	treasuryoutboundtransfer "github.com/stripe/stripe-go/v82/treasury/outboundtransfer"
	treasuryreceivedcredit "github.com/stripe/stripe-go/v82/treasury/receivedcredit"
	treasuryreceiveddebit "github.com/stripe/stripe-go/v82/treasury/receiveddebit"
	treasurytransaction "github.com/stripe/stripe-go/v82/treasury/transaction"
	treasurytransactionentry "github.com/stripe/stripe-go/v82/treasury/transactionentry"
	v2billingmeterevent "github.com/stripe/stripe-go/v82/v2/billing/meterevent"
	v2billingmetereventadjustment "github.com/stripe/stripe-go/v82/v2/billing/metereventadjustment"
	v2billingmetereventsession "github.com/stripe/stripe-go/v82/v2/billing/metereventsession"
	v2billingmetereventstream "github.com/stripe/stripe-go/v82/v2/billing/metereventstream"
	v2coreaccount "github.com/stripe/stripe-go/v82/v2/core/account"
	v2coreaccountlink "github.com/stripe/stripe-go/v82/v2/core/accountlink"
	v2coreaccountsperson "github.com/stripe/stripe-go/v82/v2/core/accounts/person"
	v2coreevent "github.com/stripe/stripe-go/v82/v2/core/event"
	v2coreeventdestination "github.com/stripe/stripe-go/v82/v2/core/eventdestination"
	v2corevaultgbbankaccount "github.com/stripe/stripe-go/v82/v2/core/vault/gbbankaccount"
	v2corevaultusbankaccount "github.com/stripe/stripe-go/v82/v2/core/vault/usbankaccount"
	v2moneymanagementadjustment "github.com/stripe/stripe-go/v82/v2/moneymanagement/adjustment"
	v2moneymanagementfinancialaccount "github.com/stripe/stripe-go/v82/v2/moneymanagement/financialaccount"
	v2moneymanagementfinancialaddress "github.com/stripe/stripe-go/v82/v2/moneymanagement/financialaddress"
	v2moneymanagementinboundtransfer "github.com/stripe/stripe-go/v82/v2/moneymanagement/inboundtransfer"
	v2moneymanagementoutboundpayment "github.com/stripe/stripe-go/v82/v2/moneymanagement/outboundpayment"
	v2moneymanagementoutboundpaymentquote "github.com/stripe/stripe-go/v82/v2/moneymanagement/outboundpaymentquote"
	v2moneymanagementoutboundsetupintent "github.com/stripe/stripe-go/v82/v2/moneymanagement/outboundsetupintent"
	v2moneymanagementoutboundtransfer "github.com/stripe/stripe-go/v82/v2/moneymanagement/outboundtransfer"
	v2moneymanagementpayoutmethod "github.com/stripe/stripe-go/v82/v2/moneymanagement/payoutmethod"
	v2moneymanagementpayoutmethodsbankaccountspec "github.com/stripe/stripe-go/v82/v2/moneymanagement/payoutmethodsbankaccountspec"
	v2moneymanagementreceivedcredit "github.com/stripe/stripe-go/v82/v2/moneymanagement/receivedcredit"
	v2moneymanagementreceiveddebit "github.com/stripe/stripe-go/v82/v2/moneymanagement/receiveddebit"
	v2moneymanagementtransaction "github.com/stripe/stripe-go/v82/v2/moneymanagement/transaction"
	v2moneymanagementtransactionentry "github.com/stripe/stripe-go/v82/v2/moneymanagement/transactionentry"
	v2testhelpersfinancialaddress "github.com/stripe/stripe-go/v82/v2/testhelpers/financialaddress"
	"github.com/stripe/stripe-go/v82/webhookendpoint"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// AccountLinks is the client used to invoke /v1/account_links APIs.
	AccountLinks *accountlink.Client
	// AccountNotices is the client used to invoke /v1/account_notices APIs.
	AccountNotices *accountnotice.Client
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
	// BalanceSettings is the client used to invoke /v1/balance_settings APIs.
	BalanceSettings *balancesettings.Client
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
	// CapitalFinancingOffers is the client used to invoke /v1/capital/financing_offers APIs.
	CapitalFinancingOffers *capitalfinancingoffer.Client
	// CapitalFinancingSummary is the client used to invoke /v1/capital/financing_summary APIs.
	CapitalFinancingSummary *capitalfinancingsummary.Client
	// CapitalFinancingTransactions is the client used to invoke /v1/capital/financing_transactions APIs.
	CapitalFinancingTransactions *capitalfinancingtransaction.Client
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
	// ExternalAccounts is the client used to invoke /v1/external_accounts APIs.
	ExternalAccounts *externalaccount.Client
	// FeeRefunds is the client used to invoke /v1/application_fees/{id}/refunds APIs.
	FeeRefunds *feerefund.Client
	// FileLinks is the client used to invoke /v1/file_links APIs.
	FileLinks *filelink.Client
	// Files is the client used to invoke /v1/files APIs.
	Files *file.Client
	// FinancialConnectionsAccountInferredBalances is the client used to invoke /v1/financial_connections/accounts/{account}/inferred_balances APIs.
	FinancialConnectionsAccountInferredBalances *financialconnectionsaccountinferredbalance.Client
	// FinancialConnectionsAccounts is the client used to invoke /v1/financial_connections/accounts APIs.
	FinancialConnectionsAccounts *financialconnectionsaccount.Client
	// FinancialConnectionsInstitutions is the client used to invoke /v1/financial_connections/institutions APIs.
	FinancialConnectionsInstitutions *financialconnectionsinstitution.Client
	// FinancialConnectionsSessions is the client used to invoke /v1/financial_connections/sessions APIs.
	FinancialConnectionsSessions *financialconnectionssession.Client
	// FinancialConnectionsTransactions is the client used to invoke /v1/financial_connections/transactions APIs.
	FinancialConnectionsTransactions *financialconnectionstransaction.Client
	// ForwardingRequests is the client used to invoke /v1/forwarding/requests APIs.
	ForwardingRequests *forwardingrequest.Client
	// FxQuotes is the client used to invoke /v1/fx_quotes APIs.
	FxQuotes *fxquote.Client
	// GiftCardsCards is the client used to invoke /v1/gift_cards/cards APIs.
	GiftCardsCards *giftcardscard.Client
	// GiftCardsTransactions is the client used to invoke /v1/gift_cards/transactions APIs.
	GiftCardsTransactions *giftcardstransaction.Client
	// IdentityVerificationReports is the client used to invoke /v1/identity/verification_reports APIs.
	IdentityVerificationReports *identityverificationreport.Client
	// IdentityVerificationSessions is the client used to invoke /v1/identity/verification_sessions APIs.
	IdentityVerificationSessions *identityverificationsession.Client
	// InvoiceItems is the client used to invoke /v1/invoiceitems APIs.
	InvoiceItems *invoiceitem.Client
	// InvoiceLineItems is the client used to invoke /v1/invoices/{invoice}/lines APIs.
	InvoiceLineItems *invoicelineitem.Client
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
	// IssuingCreditUnderwritingRecords is the client used to invoke /v1/issuing/credit_underwriting_records APIs.
	IssuingCreditUnderwritingRecords *issuingcreditunderwritingrecord.Client
	// IssuingDisputes is the client used to invoke /v1/issuing/disputes APIs.
	IssuingDisputes *issuingdispute.Client
	// IssuingDisputeSettlementDetails is the client used to invoke /v1/issuing/dispute_settlement_details APIs.
	IssuingDisputeSettlementDetails *issuingdisputesettlementdetail.Client
	// IssuingFraudLiabilityDebits is the client used to invoke /v1/issuing/fraud_liability_debits APIs.
	IssuingFraudLiabilityDebits *issuingfraudliabilitydebit.Client
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
	// Margins is the client used to invoke /v1/billing/margins APIs.
	Margins *margin.Client
	// OAuth is the client used to invoke /oauth APIs
	OAuth *oauth.Client
	// Orders is the client used to invoke /v1/orders APIs.
	Orders *order.Client
	// PaymentAttemptRecords is the client used to invoke /v1/payment_attempt_records APIs.
	PaymentAttemptRecords *paymentattemptrecord.Client
	// PaymentIntentAmountDetailsLineItems is the client used to invoke /v1/payment_intents/{intent}/amount_details_line_items APIs.
	PaymentIntentAmountDetailsLineItems *paymentintentamountdetailslineitem.Client
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
	// PaymentRecords is the client used to invoke /v1/payment_records APIs.
	PaymentRecords *paymentrecord.Client
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
	// PrivacyRedactionJobs is the client used to invoke /v1/privacy/redaction_jobs APIs.
	PrivacyRedactionJobs *privacyredactionjob.Client
	// PrivacyRedactionJobValidationErrors is the client used to invoke /v1/privacy/redaction_jobs/{job}/validation_errors APIs.
	PrivacyRedactionJobValidationErrors *privacyredactionjobvalidationerror.Client
	// ProductFeatures is the client used to invoke /v1/products/{product}/features APIs.
	ProductFeatures *productfeature.Client
	// Products is the client used to invoke /v1/products APIs.
	Products *product.Client
	// PromotionCodes is the client used to invoke /v1/promotion_codes APIs.
	PromotionCodes *promotioncode.Client
	// QuotePreviewInvoices is the client used to invoke /v1/quotes/{quote}/preview_invoices APIs.
	QuotePreviewInvoices *quotepreviewinvoice.Client
	// QuotePreviewSubscriptionSchedules is the client used to invoke /v1/quotes/{quote}/preview_subscription_schedules APIs.
	QuotePreviewSubscriptionSchedules *quotepreviewsubscriptionschedule.Client
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
	// TaxAssociations is the client used to invoke association related APIs.
	TaxAssociations *taxassociation.Client
	// TaxCalculations is the client used to invoke /v1/tax/calculations APIs.
	TaxCalculations *taxcalculation.Client
	// TaxCodes is the client used to invoke /v1/tax_codes APIs.
	TaxCodes *taxcode.Client
	// TaxForms is the client used to invoke /v1/tax/forms APIs.
	TaxForms *taxform.Client
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
	// TerminalReaderCollectedData is the client used to invoke /v1/terminal/reader_collected_data APIs.
	TerminalReaderCollectedData *terminalreadercollecteddata.Client
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
	// V2CoreAccountLinks is the client used to invoke /v2/core/account_links APIs.
	V2CoreAccountLinks *v2coreaccountlink.Client
	// V2CoreAccounts is the client used to invoke /v2/core/accounts APIs.
	V2CoreAccounts *v2coreaccount.Client
	// V2CoreAccountsPersons is the client used to invoke /v2/core/accounts/{account_id}/persons APIs.
	V2CoreAccountsPersons *v2coreaccountsperson.Client
	// V2CoreEventDestinations is the client used to invoke /v2/core/event_destinations APIs.
	V2CoreEventDestinations *v2coreeventdestination.Client
	// V2CoreEvents is the client used to invoke /v2/core/events APIs.
	V2CoreEvents *v2coreevent.Client
	// V2CoreVaultGBBankAccounts is the client used to invoke /v2/core/vault/gb_bank_accounts APIs.
	V2CoreVaultGBBankAccounts *v2corevaultgbbankaccount.Client
	// V2CoreVaultUSBankAccounts is the client used to invoke /v2/core/vault/us_bank_accounts APIs.
	V2CoreVaultUSBankAccounts *v2corevaultusbankaccount.Client
	// V2MoneyManagementAdjustments is the client used to invoke /v2/money_management/adjustments APIs.
	V2MoneyManagementAdjustments *v2moneymanagementadjustment.Client
	// V2MoneyManagementFinancialAccounts is the client used to invoke /v2/money_management/financial_accounts APIs.
	V2MoneyManagementFinancialAccounts *v2moneymanagementfinancialaccount.Client
	// V2MoneyManagementFinancialAddresses is the client used to invoke /v2/money_management/financial_addresses APIs.
	V2MoneyManagementFinancialAddresses *v2moneymanagementfinancialaddress.Client
	// V2MoneyManagementInboundTransfers is the client used to invoke /v2/money_management/inbound_transfers APIs.
	V2MoneyManagementInboundTransfers *v2moneymanagementinboundtransfer.Client
	// V2MoneyManagementOutboundPaymentQuotes is the client used to invoke /v2/money_management/outbound_payment_quotes APIs.
	V2MoneyManagementOutboundPaymentQuotes *v2moneymanagementoutboundpaymentquote.Client
	// V2MoneyManagementOutboundPayments is the client used to invoke /v2/money_management/outbound_payments APIs.
	V2MoneyManagementOutboundPayments *v2moneymanagementoutboundpayment.Client
	// V2MoneyManagementOutboundSetupIntents is the client used to invoke /v2/money_management/outbound_setup_intents APIs.
	V2MoneyManagementOutboundSetupIntents *v2moneymanagementoutboundsetupintent.Client
	// V2MoneyManagementOutboundTransfers is the client used to invoke /v2/money_management/outbound_transfers APIs.
	V2MoneyManagementOutboundTransfers *v2moneymanagementoutboundtransfer.Client
	// V2MoneyManagementPayoutMethods is the client used to invoke /v2/money_management/payout_methods APIs.
	V2MoneyManagementPayoutMethods *v2moneymanagementpayoutmethod.Client
	// V2MoneyManagementPayoutMethodsBankAccountSpecs is the client used to invoke payoutmethodsbankaccountspec related APIs.
	V2MoneyManagementPayoutMethodsBankAccountSpecs *v2moneymanagementpayoutmethodsbankaccountspec.Client
	// V2MoneyManagementReceivedCredits is the client used to invoke /v2/money_management/received_credits APIs.
	V2MoneyManagementReceivedCredits *v2moneymanagementreceivedcredit.Client
	// V2MoneyManagementReceivedDebits is the client used to invoke /v2/money_management/received_debits APIs.
	V2MoneyManagementReceivedDebits *v2moneymanagementreceiveddebit.Client
	// V2MoneyManagementTransactionEntries is the client used to invoke /v2/money_management/transaction_entries APIs.
	V2MoneyManagementTransactionEntries *v2moneymanagementtransactionentry.Client
	// V2MoneyManagementTransactions is the client used to invoke /v2/money_management/transactions APIs.
	V2MoneyManagementTransactions *v2moneymanagementtransaction.Client
	// V2TestHelpersFinancialAddresses is the client used to invoke financialaddress related APIs.
	V2TestHelpersFinancialAddresses *v2testhelpersfinancialaddress.Client
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
	a.AccountNotices = &accountnotice.Client{B: backends.API, Key: key}
	a.Accounts = &account.Client{B: backends.API, Key: key}
	a.AccountSessions = &accountsession.Client{B: backends.API, Key: key}
	a.ApplePayDomains = &applepaydomain.Client{B: backends.API, Key: key}
	a.ApplicationFees = &applicationfee.Client{B: backends.API, Key: key}
	a.AppsSecrets = &appssecret.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.BalanceSettings = &balancesettings.Client{B: backends.API, Key: key}
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
	a.CapitalFinancingOffers = &capitalfinancingoffer.Client{B: backends.API, Key: key}
	a.CapitalFinancingSummary = &capitalfinancingsummary.Client{B: backends.API, Key: key}
	a.CapitalFinancingTransactions = &capitalfinancingtransaction.Client{B: backends.API, Key: key}
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
	a.ExternalAccounts = &externalaccount.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.FileLinks = &filelink.Client{B: backends.API, Key: key}
	a.Files = &file.Client{B: backends.API, BUploads: backends.Uploads, Key: key}
	a.FinancialConnectionsAccountInferredBalances = &financialconnectionsaccountinferredbalance.Client{B: backends.API, Key: key}
	a.FinancialConnectionsAccounts = &financialconnectionsaccount.Client{B: backends.API, Key: key}
	a.FinancialConnectionsInstitutions = &financialconnectionsinstitution.Client{B: backends.API, Key: key}
	a.FinancialConnectionsSessions = &financialconnectionssession.Client{B: backends.API, Key: key}
	a.FinancialConnectionsTransactions = &financialconnectionstransaction.Client{B: backends.API, Key: key}
	a.ForwardingRequests = &forwardingrequest.Client{B: backends.API, Key: key}
	a.FxQuotes = &fxquote.Client{B: backends.API, Key: key}
	a.GiftCardsCards = &giftcardscard.Client{B: backends.API, Key: key}
	a.GiftCardsTransactions = &giftcardstransaction.Client{B: backends.API, Key: key}
	a.IdentityVerificationReports = &identityverificationreport.Client{B: backends.API, Key: key}
	a.IdentityVerificationSessions = &identityverificationsession.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.InvoiceLineItems = &invoicelineitem.Client{B: backends.API, Key: key}
	a.InvoiceRenderingTemplates = &invoicerenderingtemplate.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.IssuingAuthorizations = &issuingauthorization.Client{B: backends.API, Key: key}
	a.IssuingCardholders = &issuingcardholder.Client{B: backends.API, Key: key}
	a.IssuingCards = &issuingcard.Client{B: backends.API, Key: key}
	a.IssuingCreditUnderwritingRecords = &issuingcreditunderwritingrecord.Client{B: backends.API, Key: key}
	a.IssuingDisputes = &issuingdispute.Client{B: backends.API, Key: key}
	a.IssuingDisputeSettlementDetails = &issuingdisputesettlementdetail.Client{B: backends.API, Key: key}
	a.IssuingFraudLiabilityDebits = &issuingfraudliabilitydebit.Client{B: backends.API, Key: key}
	a.IssuingPersonalizationDesigns = &issuingpersonalizationdesign.Client{B: backends.API, Key: key}
	a.IssuingPhysicalBundles = &issuingphysicalbundle.Client{B: backends.API, Key: key}
	a.IssuingTokens = &issuingtoken.Client{B: backends.API, Key: key}
	a.IssuingTransactions = &issuingtransaction.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Mandates = &mandate.Client{B: backends.API, Key: key}
	a.Margins = &margin.Client{B: backends.API, Key: key}
	a.OAuth = &oauth.Client{B: backends.Connect, Key: key}
	a.Orders = &order.Client{B: backends.API, Key: key}
	a.PaymentAttemptRecords = &paymentattemptrecord.Client{B: backends.API, Key: key}
	a.PaymentIntentAmountDetailsLineItems = &paymentintentamountdetailslineitem.Client{B: backends.API, Key: key}
	a.PaymentIntents = &paymentintent.Client{B: backends.API, Key: key}
	a.PaymentLinks = &paymentlink.Client{B: backends.API, Key: key}
	a.PaymentMethodConfigurations = &paymentmethodconfiguration.Client{B: backends.API, Key: key}
	a.PaymentMethodDomains = &paymentmethoddomain.Client{B: backends.API, Key: key}
	a.PaymentMethods = &paymentmethod.Client{B: backends.API, Key: key}
	a.PaymentRecords = &paymentrecord.Client{B: backends.API, Key: key}
	a.PaymentSources = &paymentsource.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Persons = &person.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Prices = &price.Client{B: backends.API, Key: key}
	a.PrivacyRedactionJobs = &privacyredactionjob.Client{B: backends.API, Key: key}
	a.PrivacyRedactionJobValidationErrors = &privacyredactionjobvalidationerror.Client{B: backends.API, Key: key}
	a.ProductFeatures = &productfeature.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.PromotionCodes = &promotioncode.Client{B: backends.API, Key: key}
	a.QuotePreviewInvoices = &quotepreviewinvoice.Client{B: backends.API, Key: key}
	a.QuotePreviewSubscriptionSchedules = &quotepreviewsubscriptionschedule.Client{B: backends.API, Key: key}
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
	a.TaxAssociations = &taxassociation.Client{B: backends.API, Key: key}
	a.TaxCalculations = &taxcalculation.Client{B: backends.API, Key: key}
	a.TaxCodes = &taxcode.Client{B: backends.API, Key: key}
	a.TaxForms = &taxform.Client{B: backends.API, BUploads: backends.Uploads, Key: key}
	a.TaxIDs = &taxid.Client{B: backends.API, Key: key}
	a.TaxRates = &taxrate.Client{B: backends.API, Key: key}
	a.TaxRegistrations = &taxregistration.Client{B: backends.API, Key: key}
	a.TaxSettings = &taxsettings.Client{B: backends.API, Key: key}
	a.TaxTransactions = &taxtransaction.Client{B: backends.API, Key: key}
	a.TerminalConfigurations = &terminalconfiguration.Client{B: backends.API, Key: key}
	a.TerminalConnectionTokens = &terminalconnectiontoken.Client{B: backends.API, Key: key}
	a.TerminalLocations = &terminallocation.Client{B: backends.API, Key: key}
	a.TerminalReaderCollectedData = &terminalreadercollecteddata.Client{B: backends.API, Key: key}
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
	a.V2CoreAccountLinks = &v2coreaccountlink.Client{B: backends.API, Key: key}
	a.V2CoreAccounts = &v2coreaccount.Client{B: backends.API, Key: key}
	a.V2CoreAccountsPersons = &v2coreaccountsperson.Client{B: backends.API, Key: key}
	a.V2CoreEventDestinations = &v2coreeventdestination.Client{B: backends.API, Key: key}
	a.V2CoreEvents = &v2coreevent.Client{B: backends.API, Key: key}
	a.V2CoreVaultGBBankAccounts = &v2corevaultgbbankaccount.Client{B: backends.API, Key: key}
	a.V2CoreVaultUSBankAccounts = &v2corevaultusbankaccount.Client{B: backends.API, Key: key}
	a.V2MoneyManagementAdjustments = &v2moneymanagementadjustment.Client{B: backends.API, Key: key}
	a.V2MoneyManagementFinancialAccounts = &v2moneymanagementfinancialaccount.Client{B: backends.API, Key: key}
	a.V2MoneyManagementFinancialAddresses = &v2moneymanagementfinancialaddress.Client{B: backends.API, Key: key}
	a.V2MoneyManagementInboundTransfers = &v2moneymanagementinboundtransfer.Client{B: backends.API, Key: key}
	a.V2MoneyManagementOutboundPaymentQuotes = &v2moneymanagementoutboundpaymentquote.Client{B: backends.API, Key: key}
	a.V2MoneyManagementOutboundPayments = &v2moneymanagementoutboundpayment.Client{B: backends.API, Key: key}
	a.V2MoneyManagementOutboundSetupIntents = &v2moneymanagementoutboundsetupintent.Client{B: backends.API, Key: key}
	a.V2MoneyManagementOutboundTransfers = &v2moneymanagementoutboundtransfer.Client{B: backends.API, Key: key}
	a.V2MoneyManagementPayoutMethods = &v2moneymanagementpayoutmethod.Client{B: backends.API, Key: key}
	a.V2MoneyManagementPayoutMethodsBankAccountSpecs = &v2moneymanagementpayoutmethodsbankaccountspec.Client{B: backends.API, Key: key}
	a.V2MoneyManagementReceivedCredits = &v2moneymanagementreceivedcredit.Client{B: backends.API, Key: key}
	a.V2MoneyManagementReceivedDebits = &v2moneymanagementreceiveddebit.Client{B: backends.API, Key: key}
	a.V2MoneyManagementTransactionEntries = &v2moneymanagementtransactionentry.Client{B: backends.API, Key: key}
	a.V2MoneyManagementTransactions = &v2moneymanagementtransaction.Client{B: backends.API, Key: key}
	a.V2TestHelpersFinancialAddresses = &v2testhelpersfinancialaddress.Client{B: backends.API, Key: key}
	a.WebhookEndpoints = &webhookendpoint.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *stripe.Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
