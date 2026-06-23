package stripe

// Client is the Stripe client. It contains all the different services available.
type Client struct {
	// stripeClientStruct: The beginning of the section generated from our OpenAPI spec
	// stripeClientStruct: The end of the section generated from our OpenAPI spec

	V2CoreEvents *v2CoreEventService

	backend Backend
	key     string
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
	// enough information on the Client to make API calls
	client.backend = backends.API
	client.V2CoreEvents = &v2CoreEventService{B: backends.API, Key: key}
	client.key = key

	// stripeClientInit: The beginning of the section generated from our OpenAPI spec
	client.OAuth = &oauthService{B: backends.Connect, Key: key}
	client.V1AccountLinks = &v1AccountLinkService{B: backends.API, Key: key}
	client.V1Accounts = &v1AccountService{B: backends.API, Key: key}
	client.V1AccountSessions = &v1AccountSessionService{B: backends.API, Key: key}
	client.V1ApplePayDomains = &v1ApplePayDomainService{B: backends.API, Key: key}
	client.V1ApplicationFees = &v1ApplicationFeeService{B: backends.API, Key: key}
	client.V1AppsSecrets = &v1AppsSecretService{B: backends.API, Key: key}
	client.V1Balance = &v1BalanceService{B: backends.API, Key: key}
	client.V1BalanceSettings = &v1BalanceSettingsService{B: backends.API, Key: key}
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
	client.V1PaymentAttemptRecords = &v1PaymentAttemptRecordService{B: backends.API, Key: key}
	client.V1PaymentIntentAmountDetailsLineItems = &v1PaymentIntentAmountDetailsLineItemService{B: backends.API, Key: key}
	client.V1PaymentIntents = &v1PaymentIntentService{B: backends.API, Key: key}
	client.V1PaymentLinks = &v1PaymentLinkService{B: backends.API, Key: key}
	client.V1PaymentMethodConfigurations = &v1PaymentMethodConfigurationService{B: backends.API, Key: key}
	client.V1PaymentMethodDomains = &v1PaymentMethodDomainService{B: backends.API, Key: key}
	client.V1PaymentMethods = &v1PaymentMethodService{B: backends.API, Key: key}
	client.V1PaymentRecords = &v1PaymentRecordService{B: backends.API, Key: key}
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
	client.V1RadarPaymentEvaluations = &v1RadarPaymentEvaluationService{B: backends.API, Key: key}
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
	client.V1TaxAssociations = &v1TaxAssociationService{B: backends.API, Key: key}
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
	client.V1TerminalOnboardingLinks = &v1TerminalOnboardingLinkService{B: backends.API, Key: key}
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
	client.V2CommerceProductCatalogImports = &v2CommerceProductCatalogImportService{B: backends.API, Key: key}
	client.V2CoreAccountLinks = &v2CoreAccountLinkService{B: backends.API, Key: key}
	client.V2CoreAccounts = &v2CoreAccountService{B: backends.API, Key: key}
	client.V2CoreAccountsPersons = &v2CoreAccountsPersonService{B: backends.API, Key: key}
	client.V2CoreAccountsPersonTokens = &v2CoreAccountsPersonTokenService{B: backends.API, Key: key}
	client.V2CoreAccountTokens = &v2CoreAccountTokenService{B: backends.API, Key: key}
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

// ParseEventNotification parses a Stripe event from the payload and verifies its signature.
// It returns a union of all possible event notification types that implement EventNotificationContainer.
func (c *Client) ParseEventNotification(payload []byte, header string, secret string, opts ...WebhookOption) (EventNotificationContainer, error) {
	if err := ValidatePayload(payload, header, secret, opts...); err != nil {
		return nil, err
	}

	return EventNotificationFromJSON(payload, *c)
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
