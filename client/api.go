// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/balance"
	"github.com/stripe/stripe-go/bankaccount"
	"github.com/stripe/stripe-go/bitcoinreceiver"
	"github.com/stripe/stripe-go/bitcointransaction"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/countryspec"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/discount"
	"github.com/stripe/stripe-go/dispute"
	"github.com/stripe/stripe-go/ephemeralkey"
	"github.com/stripe/stripe-go/event"
	"github.com/stripe/stripe-go/exchangerate"
	"github.com/stripe/stripe-go/fee"
	"github.com/stripe/stripe-go/feerefund"
	"github.com/stripe/stripe-go/fileupload"
	"github.com/stripe/stripe-go/invoice"
	"github.com/stripe/stripe-go/invoiceitem"
	"github.com/stripe/stripe-go/loginlink"
	"github.com/stripe/stripe-go/order"
	"github.com/stripe/stripe-go/orderreturn"
	"github.com/stripe/stripe-go/paymentsource"
	"github.com/stripe/stripe-go/payout"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/recipient"
	"github.com/stripe/stripe-go/refund"
	"github.com/stripe/stripe-go/reversal"
	"github.com/stripe/stripe-go/sku"
	"github.com/stripe/stripe-go/source"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/subitem"
	"github.com/stripe/stripe-go/token"
	"github.com/stripe/stripe-go/transfer"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// Charges is the client used to invoke /charges APIs.
	// For more details see https://stripe.com/docs/api#charges.
	Charges *charge.Client
	// Customers is the client used to invoke /customers APIs.
	// For more details see https://stripe.com/docs/api#customers.
	Customers *customer.Client
	// Cards is the client used to invoke /cards APIs.
	// For more details see https://stripe.com/docs/api#cards.
	Cards *card.Client
	// Subs is the client used to invoke /subscriptions APIs.
	// For more details see https://stripe.com/docs/api#subscriptions.
	Subs *sub.Client
	// SubItems is the client used to invoke subscription's items-related APIs.
	// For more details see https://stripe.com/docs/api#subscription_items.
	SubItems *subitem.Client
	// Plans is the client used to invoke /plans APIs.
	// For more details see https://stripe.com/docs/api#plans.
	Plans *plan.Client
	// Coupons is the client used to invoke /coupons APIs.
	// For more details see https://stripe.com/docs/api#coupons.
	Coupons *coupon.Client
	// Discounts is the client used to invoke discount-related APIs.
	// For mode details see https://stripe.com/docs/api#discounts.
	Discounts *discount.Client
	// Invoices is the client used to invoke /invoices APIs.
	// For more details see https://stripe.com/docs/api#invoices.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	// For more details see https://stripe.com/docs/api#invoiceitems.
	InvoiceItems *invoiceitem.Client
	// LoginLinks is the client used to invoke /v1/accounts/<account_id>/login_links APIs.
	// For more details see https://stripe.com/docs/api#login_link_object.
	LoginLinks *loginlink.Client
	// Disputes is the client used to invoke dispute-related APIs.
	// For more details see https://stripe.com/docs/api#disputes.
	Disputes *dispute.Client
	// Transfers is the client used to invoke /transfers APIs.
	// For more details see https://stripe.com/docs/api#transfers.
	Transfers *transfer.Client
	// Payouts is the client used to invoke /payouts APIs.
	// For more details see https://stripe.com/docs/api#payouts.
	Payouts *payout.Client
	// Recipients is the client used to invoke /recipients APIs.
	// For more details see https://stripe.com/docs/api#recipients.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	// For more details see https://stripe.com/docs/api#refunds.
	Refunds *refund.Client
	// Fees is the client used to invoke /application_fees APIs.
	// For more details see https://stripe.com/docs/api#application_fees.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	// For more details see https://stripe.com/docs/api#fee_refundss.
	FeeRefunds *feerefund.Client
	// Account is the client used to invoke /account APIs.
	// For more details see https://stripe.com/docs/api#account.
	Account *account.Client
	// CountrySpec is the client used to invoke /country_specs APIs.
	// For more details see https://stripe.com/docs/api#country_specs.
	CountrySpec *countryspec.Client
	// Balance is the client used to invoke /balance and transaction-related APIs.
	// For more details see https://stripe.com/docs/api#balance.
	Balance *balance.Client
	// EphemeralKeys is the client used to invoke /ephemeral_keys APIs.
	// For more details see https://stripe.com/docs/api#ephemeral_keys.
	EphemeralKeys *ephemeralkey.Client
	// Events is the client used to invoke /events APIs.
	// For more details see https://stripe.com/docs/api#events.
	Events *event.Client
	// Tokens is the client used to invoke /tokens APIs.
	// For more details see https://stripe.com/docs/api#tokens.
	Tokens *token.Client
	// FileUploads is the client used to invoke the uploads /files APIs.
	// For more details see https://stripe.com/docs/api#file_uploads.
	FileUploads *fileupload.Client
	// BitcoinReceivers is the client used to invoke /bitcoin/receivers APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinReceivers *bitcoinreceiver.Client
	// BitcoinTransactions is the client used to invoke /bitcoin/transactions APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinTransactions *bitcointransaction.Client
	// Reversals is the client used to invoke /transfers/reversals APIs.
	Reversals *reversal.Client
	// BankAccounts is the client used to invoke /accounts/bank_accounts APIs.
	BankAccounts *bankaccount.Client
	// Products is the client used to invoke /products APIs.
	// For more details see https://stripe.com/docs/api#products.
	Products *product.Client
	// Orders is the client used to invoke /orders APIs.
	// For more details see https://stripe.com/docs/api#orders.
	Orders *order.Client
	// OrderReturns is the client used to invoke /order_returns APIs.
	// For more details, see https://stripe.com/docs/api#order_returns.
	OrderReturns *orderreturn.Client
	// Skus is the client used to invoke /skus APIs.
	// For more details see https://stripe.com/docs/api#skus.
	Skus *sku.Client
	// Sources is the client used to invoke /sources APIs.
	// For more details see https://stripe.com/docs/api#sources.
	Sources *source.Client
	// PaymentSource is used to invoke /sources APIs.
	// For more details see https://stripe.com/docs/api.
	PaymentSource *paymentsource.Client
	// ExchangeRates is the client used to invoke /exchange_rates APIs.
	ExchangeRates *exchangerate.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *Backends) {
	if backends == nil {
		backends = &Backends{API: GetBackend(APIBackend), Uploads: GetBackend(UploadsBackend)}
	}

	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.Subs = &sub.Client{B: backends.API, Key: key}
	a.SubItems = &subitem.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.Discounts = &discount.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Recipients = &recipient.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.Fees = &fee.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.Account = &account.Client{B: backends.API, Key: key}
	a.CountrySpec = &countryspec.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.EphemeralKeys = &ephemeralkey.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.FileUploads = &fileupload.Client{B: backends.Uploads, Key: key}
	a.BitcoinReceivers = &bitcoinreceiver.Client{B: backends.API, Key: key}
	a.BitcoinTransactions = &bitcointransaction.Client{B: backends.API, Key: key}
	a.Reversals = &reversal.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.Orders = &order.Client{B: backends.API, Key: key}
	a.OrderReturns = &orderreturn.Client{B: backends.API, Key: key}
	a.Skus = &sku.Client{B: backends.API, Key: key}
	a.Sources = &source.Client{B: backends.API, Key: key}
	a.PaymentSource = &paymentsource.Client{B: backends.API, Key: key}
	a.ExchangeRates = &exchangerate.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
