// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/channelmeter/stripe-go"
	"github.com/channelmeter/stripe-go/account"
	"github.com/channelmeter/stripe-go/accountlink"
	"github.com/channelmeter/stripe-go/applepaydomain"
	"github.com/channelmeter/stripe-go/balance"
	"github.com/channelmeter/stripe-go/balancetransaction"
	"github.com/channelmeter/stripe-go/bankaccount"
	"github.com/channelmeter/stripe-go/bitcoinreceiver"
	"github.com/channelmeter/stripe-go/bitcointransaction"
	"github.com/channelmeter/stripe-go/capability"
	"github.com/channelmeter/stripe-go/card"
	"github.com/channelmeter/stripe-go/charge"
	checkoutsession "github.com/channelmeter/stripe-go/checkout/session"
	"github.com/channelmeter/stripe-go/countryspec"
	"github.com/channelmeter/stripe-go/coupon"
	"github.com/channelmeter/stripe-go/creditnote"
	"github.com/channelmeter/stripe-go/customer"
	"github.com/channelmeter/stripe-go/customerbalancetransaction"
	"github.com/channelmeter/stripe-go/discount"
	"github.com/channelmeter/stripe-go/dispute"
	"github.com/channelmeter/stripe-go/ephemeralkey"
	"github.com/channelmeter/stripe-go/event"
	"github.com/channelmeter/stripe-go/exchangerate"
	"github.com/channelmeter/stripe-go/fee"
	"github.com/channelmeter/stripe-go/feerefund"
	"github.com/channelmeter/stripe-go/file"
	"github.com/channelmeter/stripe-go/filelink"
	"github.com/channelmeter/stripe-go/invoice"
	"github.com/channelmeter/stripe-go/invoiceitem"
	"github.com/channelmeter/stripe-go/issuing/authorization"
	issuingcard "github.com/channelmeter/stripe-go/issuing/card"
	"github.com/channelmeter/stripe-go/issuing/cardholder"
	issuingdispute "github.com/channelmeter/stripe-go/issuing/dispute"
	"github.com/channelmeter/stripe-go/issuing/transaction"
	"github.com/channelmeter/stripe-go/loginlink"
	"github.com/channelmeter/stripe-go/mandate"
	"github.com/channelmeter/stripe-go/oauth"
	"github.com/channelmeter/stripe-go/order"
	"github.com/channelmeter/stripe-go/orderreturn"
	"github.com/channelmeter/stripe-go/paymentintent"
	"github.com/channelmeter/stripe-go/paymentmethod"
	"github.com/channelmeter/stripe-go/paymentsource"
	"github.com/channelmeter/stripe-go/payout"
	"github.com/channelmeter/stripe-go/person"
	"github.com/channelmeter/stripe-go/plan"
	"github.com/channelmeter/stripe-go/product"
	"github.com/channelmeter/stripe-go/radar/earlyfraudwarning"
	"github.com/channelmeter/stripe-go/radar/valuelist"
	"github.com/channelmeter/stripe-go/radar/valuelistitem"
	"github.com/channelmeter/stripe-go/recipient"
	"github.com/channelmeter/stripe-go/refund"
	"github.com/channelmeter/stripe-go/reporting/reportrun"
	"github.com/channelmeter/stripe-go/reporting/reporttype"
	"github.com/channelmeter/stripe-go/reversal"
	"github.com/channelmeter/stripe-go/review"
	"github.com/channelmeter/stripe-go/setupintent"
	"github.com/channelmeter/stripe-go/sigma/scheduledqueryrun"
	"github.com/channelmeter/stripe-go/sku"
	"github.com/channelmeter/stripe-go/source"
	"github.com/channelmeter/stripe-go/sourcetransaction"
	"github.com/channelmeter/stripe-go/sub"
	"github.com/channelmeter/stripe-go/subitem"
	"github.com/channelmeter/stripe-go/subschedule"
	"github.com/channelmeter/stripe-go/taxid"
	"github.com/channelmeter/stripe-go/taxrate"
	terminalconnectiontoken "github.com/channelmeter/stripe-go/terminal/connectiontoken"
	terminallocation "github.com/channelmeter/stripe-go/terminal/location"
	terminalreader "github.com/channelmeter/stripe-go/terminal/reader"
	"github.com/channelmeter/stripe-go/threedsecure"
	"github.com/channelmeter/stripe-go/token"
	"github.com/channelmeter/stripe-go/topup"
	"github.com/channelmeter/stripe-go/transfer"
	"github.com/channelmeter/stripe-go/usagerecord"
	"github.com/channelmeter/stripe-go/usagerecordsummary"
	"github.com/channelmeter/stripe-go/webhookendpoint"
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
	// CheckoutSessions is the client used to invoke /checkout_sessions APIs.
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
	// Products is the client used to invoke /products APIs.
	Products *product.Client
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
	a.Products = &product.Client{B: backends.API, Key: key}
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
