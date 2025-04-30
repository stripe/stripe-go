package stripe

import (
	"encoding/json"
)

// Client is the Stripe client. It contains all the different services available.
type Client struct {
	// stripeClientStruct: The beginning of the section generated from our OpenAPI spec

	// OAuth is the service used to invoke /oauth APIs
	OAuth *oauthService
	// V1AccountLinks is the service used to invoke /v1/account_links APIs.
	V1AccountLinks *v1AccountLinkService
	// V1Accounts is the service used to invoke /v1/accounts APIs.
	V1Accounts *v1AccountService
	// V1AccountSessions is the service used to invoke /v1/account_sessions APIs.
	V1AccountSessions *v1AccountSessionService
	// V1ApplePayDomains is the service used to invoke /v1/apple_pay/domains APIs.
	V1ApplePayDomains *v1ApplePayDomainService
	// V1ApplicationFees is the service used to invoke /v1/application_fees APIs.
	V1ApplicationFees *v1ApplicationFeeService
	// V1AppsSecrets is the service used to invoke /v1/apps/secrets APIs.
	V1AppsSecrets *v1AppsSecretService
	// V1Balance is the service used to invoke /v1/balance APIs.
	V1Balance *v1BalanceService
	// V1BalanceTransactions is the service used to invoke /v1/balance_transactions APIs.
	V1BalanceTransactions *v1BalanceTransactionService
	// V1BankAccounts is the service used to invoke /v1/accounts/{account}/external_accounts APIs.
	V1BankAccounts *v1BankAccountService
	// V1BillingAlerts is the service used to invoke /v1/billing/alerts APIs.
	V1BillingAlerts *v1BillingAlertService
	// V1BillingCreditBalanceSummary is the service used to invoke /v1/billing/credit_balance_summary APIs.
	V1BillingCreditBalanceSummary *v1BillingCreditBalanceSummaryService
	// V1BillingCreditBalanceTransactions is the service used to invoke /v1/billing/credit_balance_transactions APIs.
	V1BillingCreditBalanceTransactions *v1BillingCreditBalanceTransactionService
	// V1BillingCreditGrants is the service used to invoke /v1/billing/credit_grants APIs.
	V1BillingCreditGrants *v1BillingCreditGrantService
	// V1BillingMeterEventAdjustments is the service used to invoke /v1/billing/meter_event_adjustments APIs.
	V1BillingMeterEventAdjustments *v1BillingMeterEventAdjustmentService
	// V1BillingMeterEvents is the service used to invoke /v1/billing/meter_events APIs.
	V1BillingMeterEvents *v1BillingMeterEventService
	// V1BillingMeterEventSummaries is the service used to invoke /v1/billing/meters/{id}/event_summaries APIs.
	V1BillingMeterEventSummaries *v1BillingMeterEventSummaryService
	// V1BillingMeters is the service used to invoke /v1/billing/meters APIs.
	V1BillingMeters *v1BillingMeterService
	// V1BillingPortalConfigurations is the service used to invoke /v1/billing_portal/configurations APIs.
	V1BillingPortalConfigurations *v1BillingPortalConfigurationService
	// V1BillingPortalSessions is the service used to invoke /v1/billing_portal/sessions APIs.
	V1BillingPortalSessions *v1BillingPortalSessionService
	// V1Capabilities is the service used to invoke /v1/accounts/{account}/capabilities APIs.
	V1Capabilities *v1CapabilityService
	// V1Cards is the service used to invoke /v1/accounts/{account}/external_accounts APIs.
	V1Cards *v1CardService
	// V1CashBalances is the service used to invoke /v1/customers/{customer}/cash_balance APIs.
	V1CashBalances *v1CashBalanceService
	// V1Charges is the service used to invoke /v1/charges APIs.
	V1Charges *v1ChargeService
	// V1CheckoutSessions is the service used to invoke /v1/checkout/sessions APIs.
	V1CheckoutSessions *v1CheckoutSessionService
	// V1ClimateOrders is the service used to invoke /v1/climate/orders APIs.
	V1ClimateOrders *v1ClimateOrderService
	// V1ClimateProducts is the service used to invoke /v1/climate/products APIs.
	V1ClimateProducts *v1ClimateProductService
	// V1ClimateSuppliers is the service used to invoke /v1/climate/suppliers APIs.
	V1ClimateSuppliers *v1ClimateSupplierService
	// V1ConfirmationTokens is the service used to invoke /v1/confirmation_tokens APIs.
	V1ConfirmationTokens *v1ConfirmationTokenService
	// V1CountrySpecs is the service used to invoke /v1/country_specs APIs.
	V1CountrySpecs *v1CountrySpecService
	// V1Coupons is the service used to invoke /v1/coupons APIs.
	V1Coupons *v1CouponService
	// V1CreditNotes is the service used to invoke /v1/credit_notes APIs.
	V1CreditNotes *v1CreditNoteService
	// V1CustomerBalanceTransactions is the service used to invoke /v1/customers/{customer}/balance_transactions APIs.
	V1CustomerBalanceTransactions *v1CustomerBalanceTransactionService
	// V1CustomerCashBalanceTransactions is the service used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
	V1CustomerCashBalanceTransactions *v1CustomerCashBalanceTransactionService
	// V1Customers is the service used to invoke /v1/customers APIs.
	V1Customers *v1CustomerService
	// V1CustomerSessions is the service used to invoke /v1/customer_sessions APIs.
	V1CustomerSessions *v1CustomerSessionService
	// V1Disputes is the service used to invoke /v1/disputes APIs.
	V1Disputes *v1DisputeService
	// V1EntitlementsActiveEntitlements is the service used to invoke /v1/entitlements/active_entitlements APIs.
	V1EntitlementsActiveEntitlements *v1EntitlementsActiveEntitlementService
	// V1EntitlementsFeatures is the service used to invoke /v1/entitlements/features APIs.
	V1EntitlementsFeatures *v1EntitlementsFeatureService
	// V1EphemeralKeys is the service used to invoke /v1/ephemeral_keys APIs.
	V1EphemeralKeys *v1EphemeralKeyService
	// V1Events is the service used to invoke /v1/events APIs.
	V1Events *v1EventService
	// V1FeeRefunds is the service used to invoke /v1/application_fees/{id}/refunds APIs.
	V1FeeRefunds *v1FeeRefundService
	// V1FileLinks is the service used to invoke /v1/file_links APIs.
	V1FileLinks *v1FileLinkService
	// V1Files is the service used to invoke /v1/files APIs.
	V1Files *v1FileService
	// V1FinancialConnectionsAccounts is the service used to invoke /v1/financial_connections/accounts APIs.
	V1FinancialConnectionsAccounts *v1FinancialConnectionsAccountService
	// V1FinancialConnectionsSessions is the service used to invoke /v1/financial_connections/sessions APIs.
	V1FinancialConnectionsSessions *v1FinancialConnectionsSessionService
	// V1FinancialConnectionsTransactions is the service used to invoke /v1/financial_connections/transactions APIs.
	V1FinancialConnectionsTransactions *v1FinancialConnectionsTransactionService
	// V1ForwardingRequests is the service used to invoke /v1/forwarding/requests APIs.
	V1ForwardingRequests *v1ForwardingRequestService
	// V1IdentityVerificationReports is the service used to invoke /v1/identity/verification_reports APIs.
	V1IdentityVerificationReports *v1IdentityVerificationReportService
	// V1IdentityVerificationSessions is the service used to invoke /v1/identity/verification_sessions APIs.
	V1IdentityVerificationSessions *v1IdentityVerificationSessionService
	// V1InvoiceItems is the service used to invoke /v1/invoiceitems APIs.
	V1InvoiceItems *v1InvoiceItemService
	// V1InvoiceLineItems is the service used to invoke /v1/invoices/{invoice}/lines APIs.
	V1InvoiceLineItems *v1InvoiceLineItemService
	// V1InvoicePayments is the service used to invoke /v1/invoice_payments APIs.
	V1InvoicePayments *v1InvoicePaymentService
	// V1InvoiceRenderingTemplates is the service used to invoke /v1/invoice_rendering_templates APIs.
	V1InvoiceRenderingTemplates *v1InvoiceRenderingTemplateService
	// V1Invoices is the service used to invoke /v1/invoices APIs.
	V1Invoices *v1InvoiceService
	// V1IssuingAuthorizations is the service used to invoke /v1/issuing/authorizations APIs.
	V1IssuingAuthorizations *v1IssuingAuthorizationService
	// V1IssuingCardholders is the service used to invoke /v1/issuing/cardholders APIs.
	V1IssuingCardholders *v1IssuingCardholderService
	// V1IssuingCards is the service used to invoke /v1/issuing/cards APIs.
	V1IssuingCards *v1IssuingCardService
	// V1IssuingDisputes is the service used to invoke /v1/issuing/disputes APIs.
	V1IssuingDisputes *v1IssuingDisputeService
	// V1IssuingPersonalizationDesigns is the service used to invoke /v1/issuing/personalization_designs APIs.
	V1IssuingPersonalizationDesigns *v1IssuingPersonalizationDesignService
	// V1IssuingPhysicalBundles is the service used to invoke /v1/issuing/physical_bundles APIs.
	V1IssuingPhysicalBundles *v1IssuingPhysicalBundleService
	// V1IssuingTokens is the service used to invoke /v1/issuing/tokens APIs.
	V1IssuingTokens *v1IssuingTokenService
	// V1IssuingTransactions is the service used to invoke /v1/issuing/transactions APIs.
	V1IssuingTransactions *v1IssuingTransactionService
	// V1LoginLinks is the service used to invoke /v1/accounts/{account}/login_links APIs.
	V1LoginLinks *v1LoginLinkService
	// V1Mandates is the service used to invoke /v1/mandates APIs.
	V1Mandates *v1MandateService
	// V1PaymentIntents is the service used to invoke /v1/payment_intents APIs.
	V1PaymentIntents *v1PaymentIntentService
	// V1PaymentLinks is the service used to invoke /v1/payment_links APIs.
	V1PaymentLinks *v1PaymentLinkService
	// V1PaymentMethodConfigurations is the service used to invoke /v1/payment_method_configurations APIs.
	V1PaymentMethodConfigurations *v1PaymentMethodConfigurationService
	// V1PaymentMethodDomains is the service used to invoke /v1/payment_method_domains APIs.
	V1PaymentMethodDomains *v1PaymentMethodDomainService
	// V1PaymentMethods is the service used to invoke /v1/payment_methods APIs.
	V1PaymentMethods *v1PaymentMethodService
	// V1PaymentSources is the service used to invoke /v1/customers/{customer}/sources APIs.
	V1PaymentSources *v1PaymentSourceService
	// V1Payouts is the service used to invoke /v1/payouts APIs.
	V1Payouts *v1PayoutService
	// V1Persons is the service used to invoke /v1/accounts/{account}/persons APIs.
	V1Persons *v1PersonService
	// V1Plans is the service used to invoke /v1/plans APIs.
	V1Plans *v1PlanService
	// V1Prices is the service used to invoke /v1/prices APIs.
	V1Prices *v1PriceService
	// V1ProductFeatures is the service used to invoke /v1/products/{product}/features APIs.
	V1ProductFeatures *v1ProductFeatureService
	// V1Products is the service used to invoke /v1/products APIs.
	V1Products *v1ProductService
	// V1PromotionCodes is the service used to invoke /v1/promotion_codes APIs.
	V1PromotionCodes *v1PromotionCodeService
	// V1Quotes is the service used to invoke /v1/quotes APIs.
	V1Quotes *v1QuoteService
	// V1RadarEarlyFraudWarnings is the service used to invoke /v1/radar/early_fraud_warnings APIs.
	V1RadarEarlyFraudWarnings *v1RadarEarlyFraudWarningService
	// V1RadarValueListItems is the service used to invoke /v1/radar/value_list_items APIs.
	V1RadarValueListItems *v1RadarValueListItemService
	// V1RadarValueLists is the service used to invoke /v1/radar/value_lists APIs.
	V1RadarValueLists *v1RadarValueListService
	// V1Refunds is the service used to invoke /v1/refunds APIs.
	V1Refunds *v1RefundService
	// V1ReportingReportRuns is the service used to invoke /v1/reporting/report_runs APIs.
	V1ReportingReportRuns *v1ReportingReportRunService
	// V1ReportingReportTypes is the service used to invoke /v1/reporting/report_types APIs.
	V1ReportingReportTypes *v1ReportingReportTypeService
	// V1Reviews is the service used to invoke /v1/reviews APIs.
	V1Reviews *v1ReviewService
	// V1SetupAttempts is the service used to invoke /v1/setup_attempts APIs.
	V1SetupAttempts *v1SetupAttemptService
	// V1SetupIntents is the service used to invoke /v1/setup_intents APIs.
	V1SetupIntents *v1SetupIntentService
	// V1ShippingRates is the service used to invoke /v1/shipping_rates APIs.
	V1ShippingRates *v1ShippingRateService
	// V1SigmaScheduledQueryRuns is the service used to invoke /v1/sigma/scheduled_query_runs APIs.
	V1SigmaScheduledQueryRuns *v1SigmaScheduledQueryRunService
	// V1Sources is the service used to invoke /v1/sources APIs.
	V1Sources *v1SourceService
	// V1SourceTransactions is the service used to invoke /v1/sources/{source}/source_transactions APIs.
	V1SourceTransactions *v1SourceTransactionService
	// V1SubscriptionItems is the service used to invoke /v1/subscription_items APIs.
	V1SubscriptionItems *v1SubscriptionItemService
	// V1Subscriptions is the service used to invoke /v1/subscriptions APIs.
	V1Subscriptions *v1SubscriptionService
	// V1SubscriptionSchedules is the service used to invoke /v1/subscription_schedules APIs.
	V1SubscriptionSchedules *v1SubscriptionScheduleService
	// V1TaxCalculations is the service used to invoke /v1/tax/calculations APIs.
	V1TaxCalculations *v1TaxCalculationService
	// V1TaxCodes is the service used to invoke /v1/tax_codes APIs.
	V1TaxCodes *v1TaxCodeService
	// V1TaxIDs is the service used to invoke /v1/tax_ids APIs.
	V1TaxIDs *v1TaxIDService
	// V1TaxRates is the service used to invoke /v1/tax_rates APIs.
	V1TaxRates *v1TaxRateService
	// V1TaxRegistrations is the service used to invoke /v1/tax/registrations APIs.
	V1TaxRegistrations *v1TaxRegistrationService
	// V1TaxSettings is the service used to invoke /v1/tax/settings APIs.
	V1TaxSettings *v1TaxSettingsService
	// V1TaxTransactions is the service used to invoke /v1/tax/transactions APIs.
	V1TaxTransactions *v1TaxTransactionService
	// V1TerminalConfigurations is the service used to invoke /v1/terminal/configurations APIs.
	V1TerminalConfigurations *v1TerminalConfigurationService
	// V1TerminalConnectionTokens is the service used to invoke /v1/terminal/connection_tokens APIs.
	V1TerminalConnectionTokens *v1TerminalConnectionTokenService
	// V1TerminalLocations is the service used to invoke /v1/terminal/locations APIs.
	V1TerminalLocations *v1TerminalLocationService
	// V1TerminalReaders is the service used to invoke /v1/terminal/readers APIs.
	V1TerminalReaders *v1TerminalReaderService
	// V1TestHelpersConfirmationTokens is the service used to invoke /v1/confirmation_tokens APIs.
	V1TestHelpersConfirmationTokens *v1TestHelpersConfirmationTokenService
	// V1TestHelpersCustomers is the service used to invoke /v1/customers APIs.
	V1TestHelpersCustomers *v1TestHelpersCustomerService
	// V1TestHelpersIssuingAuthorizations is the service used to invoke /v1/issuing/authorizations APIs.
	V1TestHelpersIssuingAuthorizations *v1TestHelpersIssuingAuthorizationService
	// V1TestHelpersIssuingCards is the service used to invoke /v1/issuing/cards APIs.
	V1TestHelpersIssuingCards *v1TestHelpersIssuingCardService
	// V1TestHelpersIssuingPersonalizationDesigns is the service used to invoke /v1/issuing/personalization_designs APIs.
	V1TestHelpersIssuingPersonalizationDesigns *v1TestHelpersIssuingPersonalizationDesignService
	// V1TestHelpersIssuingTransactions is the service used to invoke /v1/issuing/transactions APIs.
	V1TestHelpersIssuingTransactions *v1TestHelpersIssuingTransactionService
	// V1TestHelpersRefunds is the service used to invoke /v1/refunds APIs.
	V1TestHelpersRefunds *v1TestHelpersRefundService
	// V1TestHelpersTerminalReaders is the service used to invoke /v1/terminal/readers APIs.
	V1TestHelpersTerminalReaders *v1TestHelpersTerminalReaderService
	// V1TestHelpersTestClocks is the service used to invoke /v1/test_helpers/test_clocks APIs.
	V1TestHelpersTestClocks *v1TestHelpersTestClockService
	// V1TestHelpersTreasuryInboundTransfers is the service used to invoke /v1/treasury/inbound_transfers APIs.
	V1TestHelpersTreasuryInboundTransfers *v1TestHelpersTreasuryInboundTransferService
	// V1TestHelpersTreasuryOutboundPayments is the service used to invoke /v1/treasury/outbound_payments APIs.
	V1TestHelpersTreasuryOutboundPayments *v1TestHelpersTreasuryOutboundPaymentService
	// V1TestHelpersTreasuryOutboundTransfers is the service used to invoke /v1/treasury/outbound_transfers APIs.
	V1TestHelpersTreasuryOutboundTransfers *v1TestHelpersTreasuryOutboundTransferService
	// V1TestHelpersTreasuryReceivedCredits is the service used to invoke /v1/treasury/received_credits APIs.
	V1TestHelpersTreasuryReceivedCredits *v1TestHelpersTreasuryReceivedCreditService
	// V1TestHelpersTreasuryReceivedDebits is the service used to invoke /v1/treasury/received_debits APIs.
	V1TestHelpersTreasuryReceivedDebits *v1TestHelpersTreasuryReceivedDebitService
	// V1Tokens is the service used to invoke /v1/tokens APIs.
	V1Tokens *v1TokenService
	// V1Topups is the service used to invoke /v1/topups APIs.
	V1Topups *v1TopupService
	// V1TransferReversals is the service used to invoke /v1/transfers/{id}/reversals APIs.
	V1TransferReversals *v1TransferReversalService
	// V1Transfers is the service used to invoke /v1/transfers APIs.
	V1Transfers *v1TransferService
	// V1TreasuryCreditReversals is the service used to invoke /v1/treasury/credit_reversals APIs.
	V1TreasuryCreditReversals *v1TreasuryCreditReversalService
	// V1TreasuryDebitReversals is the service used to invoke /v1/treasury/debit_reversals APIs.
	V1TreasuryDebitReversals *v1TreasuryDebitReversalService
	// V1TreasuryFinancialAccounts is the service used to invoke /v1/treasury/financial_accounts APIs.
	V1TreasuryFinancialAccounts *v1TreasuryFinancialAccountService
	// V1TreasuryInboundTransfers is the service used to invoke /v1/treasury/inbound_transfers APIs.
	V1TreasuryInboundTransfers *v1TreasuryInboundTransferService
	// V1TreasuryOutboundPayments is the service used to invoke /v1/treasury/outbound_payments APIs.
	V1TreasuryOutboundPayments *v1TreasuryOutboundPaymentService
	// V1TreasuryOutboundTransfers is the service used to invoke /v1/treasury/outbound_transfers APIs.
	V1TreasuryOutboundTransfers *v1TreasuryOutboundTransferService
	// V1TreasuryReceivedCredits is the service used to invoke /v1/treasury/received_credits APIs.
	V1TreasuryReceivedCredits *v1TreasuryReceivedCreditService
	// V1TreasuryReceivedDebits is the service used to invoke /v1/treasury/received_debits APIs.
	V1TreasuryReceivedDebits *v1TreasuryReceivedDebitService
	// V1TreasuryTransactionEntries is the service used to invoke /v1/treasury/transaction_entries APIs.
	V1TreasuryTransactionEntries *v1TreasuryTransactionEntryService
	// V1TreasuryTransactions is the service used to invoke /v1/treasury/transactions APIs.
	V1TreasuryTransactions *v1TreasuryTransactionService
	// V1WebhookEndpoints is the service used to invoke /v1/webhook_endpoints APIs.
	V1WebhookEndpoints *v1WebhookEndpointService
	// V2BillingMeterEventAdjustments is the service used to invoke /v2/billing/meter_event_adjustments APIs.
	V2BillingMeterEventAdjustments *v2BillingMeterEventAdjustmentService
	// V2BillingMeterEvents is the service used to invoke /v2/billing/meter_events APIs.
	V2BillingMeterEvents *v2BillingMeterEventService
	// V2BillingMeterEventSessions is the service used to invoke /v2/billing/meter_event_session APIs.
	V2BillingMeterEventSessions *v2BillingMeterEventSessionService
	// V2BillingMeterEventStreams is the service used to invoke /v2/billing/meter_event_stream APIs.
	V2BillingMeterEventStreams *v2BillingMeterEventStreamService
	// V2CoreEventDestinations is the service used to invoke /v2/core/event_destinations APIs.
	V2CoreEventDestinations *v2CoreEventDestinationService
	// V2CoreEvents is the service used to invoke /v2/core/events APIs.
	V2CoreEvents *v2CoreEventService
	// stripeClientStruct: The end of the section generated from our OpenAPI spec
}

// NewClient creates a new Stripe [Client] with the given API key.
func NewClient(key string, opts ...ClientOption) *Client {
	usage := []string{"stripe_client_new"}
	client := &Client{}
	cfg := clientConfig{key: key, usage: usage}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&cfg)
	}
	initClient(client, cfg)
	return client
}

func initClient(client *Client, cfg clientConfig) {
	if cfg.backends == nil {
		cfg.backends = &Backends{
			API:         &UsageBackend{B: GetBackend(APIBackend), Usage: cfg.usage},
			Connect:     &UsageBackend{B: GetBackend(ConnectBackend), Usage: cfg.usage},
			Uploads:     &UsageBackend{B: GetBackend(UploadsBackend), Usage: cfg.usage},
			MeterEvents: &UsageBackend{B: GetBackend(MeterEventsBackend), Usage: cfg.usage},
		}
	}
	backends := cfg.backends
	key := cfg.key

	// stripeClientInit: The beginning of the section generated from our OpenAPI spec
	client.OAuth = &oauthService{B: backends.Connect, Key: key}
	client.V1AccountLinks = &v1AccountLinkService{B: backends.API, Key: key}
	client.V1Accounts = &v1AccountService{B: backends.API, Key: key}
	client.V1AccountSessions = &v1AccountSessionService{B: backends.API, Key: key}
	client.V1ApplePayDomains = &v1ApplePayDomainService{B: backends.API, Key: key}
	client.V1ApplicationFees = &v1ApplicationFeeService{B: backends.API, Key: key}
	client.V1AppsSecrets = &v1AppsSecretService{B: backends.API, Key: key}
	client.V1Balance = &v1BalanceService{B: backends.API, Key: key}
	client.V1BalanceTransactions = &v1BalanceTransactionService{B: backends.API, Key: key}
	client.V1BankAccounts = &v1BankAccountService{B: backends.API, Key: key}
	client.V1BillingAlerts = &v1BillingAlertService{B: backends.API, Key: key}
	client.V1BillingCreditBalanceSummary = &v1BillingCreditBalanceSummaryService{B: backends.API, Key: key}
	client.V1BillingCreditBalanceTransactions = &v1BillingCreditBalanceTransactionService{B: backends.API, Key: key}
	client.V1BillingCreditGrants = &v1BillingCreditGrantService{B: backends.API, Key: key}
	client.V1BillingMeterEventAdjustments = &v1BillingMeterEventAdjustmentService{B: backends.API, Key: key}
	client.V1BillingMeterEvents = &v1BillingMeterEventService{B: backends.API, Key: key}
	client.V1BillingMeterEventSummaries = &v1BillingMeterEventSummaryService{B: backends.API, Key: key}
	client.V1BillingMeters = &v1BillingMeterService{B: backends.API, Key: key}
	client.V1BillingPortalConfigurations = &v1BillingPortalConfigurationService{B: backends.API, Key: key}
	client.V1BillingPortalSessions = &v1BillingPortalSessionService{B: backends.API, Key: key}
	client.V1Capabilities = &v1CapabilityService{B: backends.API, Key: key}
	client.V1Cards = &v1CardService{B: backends.API, Key: key}
	client.V1CashBalances = &v1CashBalanceService{B: backends.API, Key: key}
	client.V1Charges = &v1ChargeService{B: backends.API, Key: key}
	client.V1CheckoutSessions = &v1CheckoutSessionService{B: backends.API, Key: key}
	client.V1ClimateOrders = &v1ClimateOrderService{B: backends.API, Key: key}
	client.V1ClimateProducts = &v1ClimateProductService{B: backends.API, Key: key}
	client.V1ClimateSuppliers = &v1ClimateSupplierService{B: backends.API, Key: key}
	client.V1ConfirmationTokens = &v1ConfirmationTokenService{B: backends.API, Key: key}
	client.V1CountrySpecs = &v1CountrySpecService{B: backends.API, Key: key}
	client.V1Coupons = &v1CouponService{B: backends.API, Key: key}
	client.V1CreditNotes = &v1CreditNoteService{B: backends.API, Key: key}
	client.V1CustomerBalanceTransactions = &v1CustomerBalanceTransactionService{B: backends.API, Key: key}
	client.V1CustomerCashBalanceTransactions = &v1CustomerCashBalanceTransactionService{B: backends.API, Key: key}
	client.V1Customers = &v1CustomerService{B: backends.API, Key: key}
	client.V1CustomerSessions = &v1CustomerSessionService{B: backends.API, Key: key}
	client.V1Disputes = &v1DisputeService{B: backends.API, Key: key}
	client.V1EntitlementsActiveEntitlements = &v1EntitlementsActiveEntitlementService{B: backends.API, Key: key}
	client.V1EntitlementsFeatures = &v1EntitlementsFeatureService{B: backends.API, Key: key}
	client.V1EphemeralKeys = &v1EphemeralKeyService{B: backends.API, Key: key}
	client.V1Events = &v1EventService{B: backends.API, Key: key}
	client.V1FeeRefunds = &v1FeeRefundService{B: backends.API, Key: key}
	client.V1FileLinks = &v1FileLinkService{B: backends.API, Key: key}
	client.V1Files = &v1FileService{B: backends.API, BUploads: backends.Uploads, Key: key}
	client.V1FinancialConnectionsAccounts = &v1FinancialConnectionsAccountService{B: backends.API, Key: key}
	client.V1FinancialConnectionsSessions = &v1FinancialConnectionsSessionService{B: backends.API, Key: key}
	client.V1FinancialConnectionsTransactions = &v1FinancialConnectionsTransactionService{B: backends.API, Key: key}
	client.V1ForwardingRequests = &v1ForwardingRequestService{B: backends.API, Key: key}
	client.V1IdentityVerificationReports = &v1IdentityVerificationReportService{B: backends.API, Key: key}
	client.V1IdentityVerificationSessions = &v1IdentityVerificationSessionService{B: backends.API, Key: key}
	client.V1InvoiceItems = &v1InvoiceItemService{B: backends.API, Key: key}
	client.V1InvoiceLineItems = &v1InvoiceLineItemService{B: backends.API, Key: key}
	client.V1InvoicePayments = &v1InvoicePaymentService{B: backends.API, Key: key}
	client.V1InvoiceRenderingTemplates = &v1InvoiceRenderingTemplateService{B: backends.API, Key: key}
	client.V1Invoices = &v1InvoiceService{B: backends.API, Key: key}
	client.V1IssuingAuthorizations = &v1IssuingAuthorizationService{B: backends.API, Key: key}
	client.V1IssuingCardholders = &v1IssuingCardholderService{B: backends.API, Key: key}
	client.V1IssuingCards = &v1IssuingCardService{B: backends.API, Key: key}
	client.V1IssuingDisputes = &v1IssuingDisputeService{B: backends.API, Key: key}
	client.V1IssuingPersonalizationDesigns = &v1IssuingPersonalizationDesignService{B: backends.API, Key: key}
	client.V1IssuingPhysicalBundles = &v1IssuingPhysicalBundleService{B: backends.API, Key: key}
	client.V1IssuingTokens = &v1IssuingTokenService{B: backends.API, Key: key}
	client.V1IssuingTransactions = &v1IssuingTransactionService{B: backends.API, Key: key}
	client.V1LoginLinks = &v1LoginLinkService{B: backends.API, Key: key}
	client.V1Mandates = &v1MandateService{B: backends.API, Key: key}
	client.V1PaymentIntents = &v1PaymentIntentService{B: backends.API, Key: key}
	client.V1PaymentLinks = &v1PaymentLinkService{B: backends.API, Key: key}
	client.V1PaymentMethodConfigurations = &v1PaymentMethodConfigurationService{B: backends.API, Key: key}
	client.V1PaymentMethodDomains = &v1PaymentMethodDomainService{B: backends.API, Key: key}
	client.V1PaymentMethods = &v1PaymentMethodService{B: backends.API, Key: key}
	client.V1PaymentSources = &v1PaymentSourceService{B: backends.API, Key: key}
	client.V1Payouts = &v1PayoutService{B: backends.API, Key: key}
	client.V1Persons = &v1PersonService{B: backends.API, Key: key}
	client.V1Plans = &v1PlanService{B: backends.API, Key: key}
	client.V1Prices = &v1PriceService{B: backends.API, Key: key}
	client.V1ProductFeatures = &v1ProductFeatureService{B: backends.API, Key: key}
	client.V1Products = &v1ProductService{B: backends.API, Key: key}
	client.V1PromotionCodes = &v1PromotionCodeService{B: backends.API, Key: key}
	client.V1Quotes = &v1QuoteService{B: backends.API, BUploads: backends.Uploads, Key: key}
	client.V1RadarEarlyFraudWarnings = &v1RadarEarlyFraudWarningService{B: backends.API, Key: key}
	client.V1RadarValueListItems = &v1RadarValueListItemService{B: backends.API, Key: key}
	client.V1RadarValueLists = &v1RadarValueListService{B: backends.API, Key: key}
	client.V1Refunds = &v1RefundService{B: backends.API, Key: key}
	client.V1ReportingReportRuns = &v1ReportingReportRunService{B: backends.API, Key: key}
	client.V1ReportingReportTypes = &v1ReportingReportTypeService{B: backends.API, Key: key}
	client.V1Reviews = &v1ReviewService{B: backends.API, Key: key}
	client.V1SetupAttempts = &v1SetupAttemptService{B: backends.API, Key: key}
	client.V1SetupIntents = &v1SetupIntentService{B: backends.API, Key: key}
	client.V1ShippingRates = &v1ShippingRateService{B: backends.API, Key: key}
	client.V1SigmaScheduledQueryRuns = &v1SigmaScheduledQueryRunService{B: backends.API, Key: key}
	client.V1Sources = &v1SourceService{B: backends.API, Key: key}
	client.V1SourceTransactions = &v1SourceTransactionService{B: backends.API, Key: key}
	client.V1SubscriptionItems = &v1SubscriptionItemService{B: backends.API, Key: key}
	client.V1Subscriptions = &v1SubscriptionService{B: backends.API, Key: key}
	client.V1SubscriptionSchedules = &v1SubscriptionScheduleService{B: backends.API, Key: key}
	client.V1TaxCalculations = &v1TaxCalculationService{B: backends.API, Key: key}
	client.V1TaxCodes = &v1TaxCodeService{B: backends.API, Key: key}
	client.V1TaxIDs = &v1TaxIDService{B: backends.API, Key: key}
	client.V1TaxRates = &v1TaxRateService{B: backends.API, Key: key}
	client.V1TaxRegistrations = &v1TaxRegistrationService{B: backends.API, Key: key}
	client.V1TaxSettings = &v1TaxSettingsService{B: backends.API, Key: key}
	client.V1TaxTransactions = &v1TaxTransactionService{B: backends.API, Key: key}
	client.V1TerminalConfigurations = &v1TerminalConfigurationService{B: backends.API, Key: key}
	client.V1TerminalConnectionTokens = &v1TerminalConnectionTokenService{B: backends.API, Key: key}
	client.V1TerminalLocations = &v1TerminalLocationService{B: backends.API, Key: key}
	client.V1TerminalReaders = &v1TerminalReaderService{B: backends.API, Key: key}
	client.V1TestHelpersConfirmationTokens = &v1TestHelpersConfirmationTokenService{B: backends.API, Key: key}
	client.V1TestHelpersCustomers = &v1TestHelpersCustomerService{B: backends.API, Key: key}
	client.V1TestHelpersIssuingAuthorizations = &v1TestHelpersIssuingAuthorizationService{B: backends.API, Key: key}
	client.V1TestHelpersIssuingCards = &v1TestHelpersIssuingCardService{B: backends.API, Key: key}
	client.V1TestHelpersIssuingPersonalizationDesigns = &v1TestHelpersIssuingPersonalizationDesignService{B: backends.API, Key: key}
	client.V1TestHelpersIssuingTransactions = &v1TestHelpersIssuingTransactionService{B: backends.API, Key: key}
	client.V1TestHelpersRefunds = &v1TestHelpersRefundService{B: backends.API, Key: key}
	client.V1TestHelpersTerminalReaders = &v1TestHelpersTerminalReaderService{B: backends.API, Key: key}
	client.V1TestHelpersTestClocks = &v1TestHelpersTestClockService{B: backends.API, Key: key}
	client.V1TestHelpersTreasuryInboundTransfers = &v1TestHelpersTreasuryInboundTransferService{B: backends.API, Key: key}
	client.V1TestHelpersTreasuryOutboundPayments = &v1TestHelpersTreasuryOutboundPaymentService{B: backends.API, Key: key}
	client.V1TestHelpersTreasuryOutboundTransfers = &v1TestHelpersTreasuryOutboundTransferService{B: backends.API, Key: key}
	client.V1TestHelpersTreasuryReceivedCredits = &v1TestHelpersTreasuryReceivedCreditService{B: backends.API, Key: key}
	client.V1TestHelpersTreasuryReceivedDebits = &v1TestHelpersTreasuryReceivedDebitService{B: backends.API, Key: key}
	client.V1Tokens = &v1TokenService{B: backends.API, Key: key}
	client.V1Topups = &v1TopupService{B: backends.API, Key: key}
	client.V1TransferReversals = &v1TransferReversalService{B: backends.API, Key: key}
	client.V1Transfers = &v1TransferService{B: backends.API, Key: key}
	client.V1TreasuryCreditReversals = &v1TreasuryCreditReversalService{B: backends.API, Key: key}
	client.V1TreasuryDebitReversals = &v1TreasuryDebitReversalService{B: backends.API, Key: key}
	client.V1TreasuryFinancialAccounts = &v1TreasuryFinancialAccountService{B: backends.API, Key: key}
	client.V1TreasuryInboundTransfers = &v1TreasuryInboundTransferService{B: backends.API, Key: key}
	client.V1TreasuryOutboundPayments = &v1TreasuryOutboundPaymentService{B: backends.API, Key: key}
	client.V1TreasuryOutboundTransfers = &v1TreasuryOutboundTransferService{B: backends.API, Key: key}
	client.V1TreasuryReceivedCredits = &v1TreasuryReceivedCreditService{B: backends.API, Key: key}
	client.V1TreasuryReceivedDebits = &v1TreasuryReceivedDebitService{B: backends.API, Key: key}
	client.V1TreasuryTransactionEntries = &v1TreasuryTransactionEntryService{B: backends.API, Key: key}
	client.V1TreasuryTransactions = &v1TreasuryTransactionService{B: backends.API, Key: key}
	client.V1WebhookEndpoints = &v1WebhookEndpointService{B: backends.API, Key: key}
	client.V2BillingMeterEventAdjustments = &v2BillingMeterEventAdjustmentService{B: backends.API, Key: key}
	client.V2BillingMeterEvents = &v2BillingMeterEventService{B: backends.API, Key: key}
	client.V2BillingMeterEventSessions = &v2BillingMeterEventSessionService{B: backends.API, Key: key}
	client.V2BillingMeterEventStreams = &v2BillingMeterEventStreamService{BMeterEvents: backends.MeterEvents, Key: key}
	client.V2CoreEventDestinations = &v2CoreEventDestinationService{B: backends.API, Key: key}
	client.V2CoreEvents = &v2CoreEventService{B: backends.API, Key: key}
	// stripeClientInit: The end of the section generated from our OpenAPI spec
}

type clientConfig struct {
	backends *Backends
	usage    []string
	key      string
}

// ClientOption allows for functional options to be passed to the NewClient constructor.
type ClientOption func(*clientConfig)

// WithBackends allows for setting a custom [*Backends] struct when creating a new client.
// This is useful for testing or when you want to use a different backend constructed
// from [NewBackendsWithConfig].
func WithBackends(backends *Backends) ClientOption {
	return func(c *clientConfig) {
		c.backends = backends
	}
}

// ParseThinEvent parses a Stripe event from the payload and verifies its signature.
// It returns a ThinEvent object and an error if the parsing or verification fails.
func (c *Client) ParseThinEvent(payload []byte, header string, secret string, opts ...WebhookOption) (*ThinEvent, error) {
	if err := ValidatePayload(payload, header, secret, opts...); err != nil {
		return nil, err
	}
	var event ThinEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// ConstructEvent initializes an Event object from a JSON webhook payload, validating
// the Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable, if the
// signature doesn't match, or if the timestamp for the signature is older than
// WebhookDefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// APIVersion constant.
func (c *Client) ConstructEvent(payload []byte, header string, secret string, opts ...WebhookOption) (Event, error) {
	return ConstructEvent(payload, header, secret, opts...)
}
