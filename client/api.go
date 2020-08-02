// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/account"
	"github.com/stripe/stripe-go/v71/accountlink"
	"github.com/stripe/stripe-go/v71/applepaydomain"
	"github.com/stripe/stripe-go/v71/balance"
	"github.com/stripe/stripe-go/v71/balancetransaction"
	"github.com/stripe/stripe-go/v71/bankaccount"
	billingportalsession "github.com/stripe/stripe-go/v71/billingportal/session"
	"github.com/stripe/stripe-go/v71/bitcoinreceiver"
	"github.com/stripe/stripe-go/v71/bitcointransaction"
	"github.com/stripe/stripe-go/v71/capability"
	"github.com/stripe/stripe-go/v71/card"
	"github.com/stripe/stripe-go/v71/charge"
	checkoutsession "github.com/stripe/stripe-go/v71/checkout/session"
	"github.com/stripe/stripe-go/v71/countryspec"
	"github.com/stripe/stripe-go/v71/coupon"
	"github.com/stripe/stripe-go/v71/creditnote"
	"github.com/stripe/stripe-go/v71/customer"
	"github.com/stripe/stripe-go/v71/customerbalancetransaction"
	"github.com/stripe/stripe-go/v71/discount"
	"github.com/stripe/stripe-go/v71/dispute"
	"github.com/stripe/stripe-go/v71/ephemeralkey"
	"github.com/stripe/stripe-go/v71/event"
	"github.com/stripe/stripe-go/v71/exchangerate"
	"github.com/stripe/stripe-go/v71/fee"
	"github.com/stripe/stripe-go/v71/feerefund"
	"github.com/stripe/stripe-go/v71/file"
	"github.com/stripe/stripe-go/v71/filelink"
	"github.com/stripe/stripe-go/v71/invoice"
	"github.com/stripe/stripe-go/v71/invoiceitem"
	"github.com/stripe/stripe-go/v71/issuing/authorization"
	issuingcard "github.com/stripe/stripe-go/v71/issuing/card"
	"github.com/stripe/stripe-go/v71/issuing/cardholder"
	issuingdispute "github.com/stripe/stripe-go/v71/issuing/dispute"
	"github.com/stripe/stripe-go/v71/issuing/transaction"
	"github.com/stripe/stripe-go/v71/loginlink"
	"github.com/stripe/stripe-go/v71/mandate"
	"github.com/stripe/stripe-go/v71/oauth"
	"github.com/stripe/stripe-go/v71/order"
	"github.com/stripe/stripe-go/v71/orderreturn"
	"github.com/stripe/stripe-go/v71/paymentintent"
	"github.com/stripe/stripe-go/v71/paymentmethod"
	"github.com/stripe/stripe-go/v71/paymentsource"
	"github.com/stripe/stripe-go/v71/payout"
	"github.com/stripe/stripe-go/v71/person"
	"github.com/stripe/stripe-go/v71/plan"
	"github.com/stripe/stripe-go/v71/price"
	"github.com/stripe/stripe-go/v71/product"
	"github.com/stripe/stripe-go/v71/promotioncode"
	"github.com/stripe/stripe-go/v71/radar/earlyfraudwarning"
	"github.com/stripe/stripe-go/v71/radar/valuelist"
	"github.com/stripe/stripe-go/v71/radar/valuelistitem"
	"github.com/stripe/stripe-go/v71/recipient"
	"github.com/stripe/stripe-go/v71/refund"
	"github.com/stripe/stripe-go/v71/reporting/reportrun"
	"github.com/stripe/stripe-go/v71/reporting/reporttype"
	"github.com/stripe/stripe-go/v71/reversal"
	"github.com/stripe/stripe-go/v71/review"
	"github.com/stripe/stripe-go/v71/setupintent"
	"github.com/stripe/stripe-go/v71/sigma/scheduledqueryrun"
	"github.com/stripe/stripe-go/v71/sku"
	"github.com/stripe/stripe-go/v71/source"
	"github.com/stripe/stripe-go/v71/sourcetransaction"
	"github.com/stripe/stripe-go/v71/sub"
	"github.com/stripe/stripe-go/v71/subitem"
	"github.com/stripe/stripe-go/v71/subschedule"
	"github.com/stripe/stripe-go/v71/taxid"
	"github.com/stripe/stripe-go/v71/taxrate"
	terminalconnectiontoken "github.com/stripe/stripe-go/v71/terminal/connectiontoken"
	terminallocation "github.com/stripe/stripe-go/v71/terminal/location"
	terminalreader "github.com/stripe/stripe-go/v71/terminal/reader"
	"github.com/stripe/stripe-go/v71/threedsecure"
	"github.com/stripe/stripe-go/v71/token"
	"github.com/stripe/stripe-go/v71/topup"
	"github.com/stripe/stripe-go/v71/transfer"
	"github.com/stripe/stripe-go/v71/usagerecord"
	"github.com/stripe/stripe-go/v71/usagerecordsummary"
	"github.com/stripe/stripe-go/v71/webhookendpoint"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// Account is the client used to invoke /accounts APIs.
	Account *account.Client
	// AccountLink is the client used to invoke /account_links APIs.
	AccountLinks *accountlink.Client
	// ApplePayDomains is the client used to invoke /apple_pay/domains APIs.
	ApplePayDomains *applepaydomain.Client
	// Balance is the client used to invoke /balance APIs.
	Balance *balance.Client
	// BalanceTransaction is the client used to invoke /balance_transactions APIs.
	BalanceTransaction *balancetransaction.Client
	// BankAccounts is the client used to invoke bank account related APIs.
	BankAccounts *bankaccount.Client
	// BillingPortalSessions is the client used to invoke /billing_portal/sessions APIs.
	BillingPortalSessions *billingportalsession.Client
	// BitcoinReceivers is the client used to invoke /bitcoin/receivers APIs.
	BitcoinReceivers *bitcoinreceiver.Client
	// BitcoinTransactions is the client used to invoke /bitcoin/transactions APIs.
	BitcoinTransactions *bitcointransaction.Client
	// Capabilities is the client used to invoke capability related APIs.
	Capabilities *capability.Client
	// Cards is the client used to invoke card related APIs.
	Cards *card.Client
	// Charges is the client used to invoke /charges APIs.
	Charges *charge.Client
	// CheckoutSessions is the client used to invoke /checkout/sessions APIs.
	CheckoutSessions *checkoutsession.Client
	// CountrySpec is the client used to invoke /country_specs APIs.
	CountrySpec *countryspec.Client
	// Coupons is the client used to invoke /coupons APIs.
	Coupons *coupon.Client
	// CreditNotes is the client used to invoke /credit_notes APIs.
	CreditNotes *creditnote.Client
	// Customers is the client used to invoke /customers APIs.
	Customers *customer.Client
	// CustomerBalanceTransactions is the client used to invoke /customers/balance_transactions APIs.
	CustomerBalanceTransactions *customerbalancetransaction.Client
	// Discounts is the client used to invoke discount related APIs.
	Discounts *discount.Client
	// Disputes is the client used to invoke /disputes APIs.
	Disputes *dispute.Client
	// EphemeralKeys is the client used to invoke /ephemeral_keys APIs.
	EphemeralKeys *ephemeralkey.Client
	// Events is the client used to invoke /events APIs.
	Events *event.Client
	// ExchangeRates is the client used to invoke /exchange_rates APIs.
	ExchangeRates *exchangerate.Client
	// Fees is the client used to invoke /application_fees APIs.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	FeeRefunds *feerefund.Client
	// Files is the client used to invoke the /files APIs.
	Files *file.Client
	// FileLinks is the client used to invoke the /file_links APIs.
	FileLinks *filelink.Client
	// Invoices is the client used to invoke /invoices APIs.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	InvoiceItems *invoiceitem.Client
	// IssuingAuthorizations is the client used to invoke /issuing/authorizations APIs.
	IssuingAuthorizations *authorization.Client
	// IssuingCardholders is the client used to invoke /issuing/cardholders APIs.
	IssuingCardholders *cardholder.Client
	// IssuingCards is the client used to invoke /issuing/cards APIs.
	IssuingCards *issuingcard.Client
	// IssuingDisputes is the client used to invoke /issuing/disputes APIs.
	IssuingDisputes *issuingdispute.Client
	// IssuingTransactions is the client used to invoke /issuing/transactions APIs.
	IssuingTransactions *transaction.Client
	// LoginLinks is the client used to invoke login link related APIs.
	LoginLinks *loginlink.Client
	// Mandates is the client used to invoke mandates related APIs.
	Mandates *mandate.Client
	// OAuth is the client used to invoke /oauth APIs.
	OAuth *oauth.Client
	// Orders is the client used to invoke /orders APIs.
	Orders *order.Client
	// OrderReturns is the client used to invoke /order_returns APIs.
	OrderReturns *orderreturn.Client
	// PaymentIntents is the client used to invoke /payment_intents APIs.
	PaymentIntents *paymentintent.Client
	// PaymentMethods is the client used to invoke /payment_methods APIs.
	PaymentMethods *paymentmethod.Client
	// PaymentSource is used to invoke customer sources related APIs.
	PaymentSource *paymentsource.Client
	// Payouts is the client used to invoke /payouts APIs.
	Payouts *payout.Client
	// Persons is the client used to invoke /account/persons APIs.
	Persons *person.Client
	// Plans is the client used to invoke /plans APIs.
	Plans *plan.Client
	// Prices is the client used to invoke /prices APIs.
	Prices *price.Client
	// Products is the client used to invoke /products APIs.
	Products *product.Client
	// PromotionCodes is the client used to invoke /promotion_codes APIs.
	PromotionCodes *promotioncode.Client
	// RadarEarlyFraudWarnings is the client used to invoke /radar/early_fraud_warnings APIs.
	RadarEarlyFraudWarnings *earlyfraudwarning.Client
	// RadarValueLists is the client used to invoke /radar/value_lists APIs.
	RadarValueLists *valuelist.Client
	// RadarValueListItems is the client used to invoke /radar/value_list_items APIs.
	RadarValueListItems *valuelistitem.Client
	// Recipients is the client used to invoke /recipients APIs.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	Refunds *refund.Client
	// ReportRuns is the client used to invoke /reporting/report_runs APIs.
	ReportRuns *reportrun.Client
	// ReportTypes is the client used to invoke /reporting/report_types APIs.
	ReportTypes *reporttype.Client
	// Reversals is the client used to invoke /transfers/reversals APIs.
	Reversals *reversal.Client
	// Reviews is the client used to invoke /reviews APIs.
	Reviews *review.Client
	// SetupIntents is the client used to invoke /setup_intents APIs.
	SetupIntents *setupintent.Client
	// SigmaScheduledQueryRuns is the client used to invoke /sigma/scheduled_query_runs APIs.
	SigmaScheduledQueryRuns *scheduledqueryrun.Client
	// Skus is the client used to invoke /skus APIs.
	Skus *sku.Client
	// Sources is the client used to invoke /sources APIs.
	Sources *source.Client
	// SourceTransactions is the client used to invoke source transaction related APIs.
	SourceTransactions *sourcetransaction.Client
	// Subscriptions is the client used to invoke /subscriptions APIs.
	Subscriptions *sub.Client
	// SubscriptionItems is the client used to invoke subscription's items related APIs.
	SubscriptionItems *subitem.Client
	// SubscriptionSchedules is the client used to invoke subscription schedules related APIs.
	SubscriptionSchedules *subschedule.Client
	// TaxIDs is the client used to invoke /tax_ids APIs.
	TaxIDs *taxid.Client
	// TaxRates is the client used to invoke /tax_rates APIs.
	TaxRates *taxrate.Client
	// TerminalConnectionTokens is the client used to invoke /terminal/connectiontokens related APIs.
	TerminalConnectionTokens *terminalconnectiontoken.Client
	// TerminalLocations is the client used to invoke /terminal/locations related APIs.
	TerminalLocations *terminallocation.Client
	// TerminalReaders is the client used to invoke /terminal/readers related APIs.
	TerminalReaders *terminalreader.Client
	// ThreeDSecures is the client used to invoke /3d_secure related APIs.
	ThreeDSecures *threedsecure.Client
	// Tokens is the client used to invoke /tokens APIs.
	Tokens *token.Client
	// Topups is the client used to invoke /tokens APIs.
	Topups *topup.Client
	// Transfers is the client used to invoke /transfers APIs.
	Transfers *transfer.Client
	// UsageRecords is the client used to invoke usage record related APIs.
	UsageRecords *usagerecord.Client
	// UsageRecordsummaries is the client used to invoke usage record summary related APIs.
	UsageRecordSummaries *usagerecordsummary.Client
	// WebhookEndpoints is the client used to invoke usage record related APIs.
	WebhookEndpoints *webhookendpoint.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *stripe.Backends) {
	if backends == nil {
		backends = &stripe.Backends{
			API:     stripe.GetBackend(stripe.APIBackend),
			Connect: stripe.GetBackend(stripe.ConnectBackend),
			Uploads: stripe.GetBackend(stripe.UploadsBackend),
		}
	}

	a.Account = &account.Client{B: backends.API, Key: key}
	a.ApplePayDomains = &applepaydomain.Client{B: backends.API, Key: key}
	a.AccountLinks = &accountlink.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.BalanceTransaction = &balancetransaction.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.BillingPortalSessions = &billingportalsession.Client{B: backends.API, Key: key}
	a.BitcoinReceivers = &bitcoinreceiver.Client{B: backends.API, Key: key}
	a.BitcoinTransactions = &bitcointransaction.Client{B: backends.API, Key: key}
	a.Capabilities = &capability.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.CheckoutSessions = &checkoutsession.Client{B: backends.API, Key: key}
	a.CountrySpec = &countryspec.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.CreditNotes = &creditnote.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.CustomerBalanceTransactions = &customerbalancetransaction.Client{B: backends.API, Key: key}
	a.Discounts = &discount.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.EphemeralKeys = &ephemeralkey.Client{B: backends.API, Key: key}
	a.ExchangeRates = &exchangerate.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.Fees = &fee.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.Files = &file.Client{B: backends.Uploads, Key: key}
	a.FileLinks = &filelink.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.IssuingAuthorizations = &authorization.Client{B: backends.API, Key: key}
	a.IssuingCardholders = &cardholder.Client{B: backends.API, Key: key}
	a.IssuingCards = &issuingcard.Client{B: backends.API, Key: key}
	a.IssuingDisputes = &issuingdispute.Client{B: backends.API, Key: key}
	a.IssuingTransactions = &transaction.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Mandates = &mandate.Client{B: backends.API, Key: key}
	a.OAuth = &oauth.Client{B: backends.Connect, Key: key}
	a.OrderReturns = &orderreturn.Client{B: backends.API, Key: key}
	a.Orders = &order.Client{B: backends.API, Key: key}
	a.PaymentIntents = &paymentintent.Client{B: backends.API, Key: key}
	a.PaymentMethods = &paymentmethod.Client{B: backends.API, Key: key}
	a.PaymentSource = &paymentsource.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Persons = &person.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Prices = &price.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.PromotionCodes = &promotioncode.Client{B: backends.API, Key: key}
	a.RadarEarlyFraudWarnings = &earlyfraudwarning.Client{B: backends.API, Key: key}
	a.RadarValueLists = &valuelist.Client{B: backends.API, Key: key}
	a.RadarValueListItems = &valuelistitem.Client{B: backends.API, Key: key}
	a.Recipients = &recipient.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.ReportRuns = &reportrun.Client{B: backends.API, Key: key}
	a.ReportTypes = &reporttype.Client{B: backends.API, Key: key}
	a.Reversals = &reversal.Client{B: backends.API, Key: key}
	a.Reviews = &review.Client{B: backends.API, Key: key}
	a.SetupIntents = &setupintent.Client{B: backends.API, Key: key}
	a.SigmaScheduledQueryRuns = &scheduledqueryrun.Client{B: backends.API, Key: key}
	a.Skus = &sku.Client{B: backends.API, Key: key}
	a.Sources = &source.Client{B: backends.API, Key: key}
	a.SourceTransactions = &sourcetransaction.Client{B: backends.API, Key: key}
	a.Subscriptions = &sub.Client{B: backends.API, Key: key}
	a.SubscriptionItems = &subitem.Client{B: backends.API, Key: key}
	a.SubscriptionSchedules = &subschedule.Client{B: backends.API, Key: key}
	a.TaxIDs = &taxid.Client{B: backends.API, Key: key}
	a.TaxRates = &taxrate.Client{B: backends.API, Key: key}
	a.TerminalConnectionTokens = &terminalconnectiontoken.Client{B: backends.API, Key: key}
	a.TerminalLocations = &terminallocation.Client{B: backends.API, Key: key}
	a.TerminalReaders = &terminalreader.Client{B: backends.API, Key: key}
	a.ThreeDSecures = &threedsecure.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.Topups = &topup.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
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
