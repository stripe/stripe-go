// package client provides a Stripe client for invoking APIs across all resources
package client

import (
	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/balance"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/discount"
	"github.com/stripe/stripe-go/dispute"
	"github.com/stripe/stripe-go/event"
	"github.com/stripe/stripe-go/fee"
	"github.com/stripe/stripe-go/feerefund"
	"github.com/stripe/stripe-go/invoice"
	"github.com/stripe/stripe-go/invoiceitem"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/recipient"
	"github.com/stripe/stripe-go/refund"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/token"
	"github.com/stripe/stripe-go/transfer"
)

// Client is the Stripe client. It contains all the different resources available.
type Api struct {
	// key is the secret key used for authentication.
	key string
	// backend is the Backend implementation used to invoke Stripe APIs.
	backend Backend
	// Charges is the client used to invoke /charges APIs.
	// For more details see https://stripe.com/docs/api/#charges.
	Charges *charge.Client
	// Customers is the client used to invoke /customers APIs.
	// For more details see https://stripe.com/docs/api/#customers.
	Customers *customer.Client
	// Cards is the client used to invoke /cards APIs.
	// For more details see https://stripe.com/docs/api/#cards.
	Cards *card.Client
	// Subs is the client used to invoke /subscriptions APIs.
	// For more details see https://stripe.com/docs/api/#subscriptions.
	Subs *sub.Client
	// Plans is the client used to invoke /plans APIs.
	// For more details see https://stripe.com/docs/api/#plans.
	Plans *plan.Client
	// Coupons is the client used to invoke /coupons APIs.
	// For more details see https://stripe.com/docs/api/#coupons.
	Coupons *coupon.Client
	// Discounts is the client used to invoke discount-related APIs.
	// For mode details see https://stripe.com/docs/api/#discounts.
	Discounts *discount.Client
	// Invoices is the client used to invoke /invoices APIs.
	// For more details see https://stripe.com/docs/api/#invoices.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	// For more details see https://stripe.com/docs/api/#invoiceitems.
	InvoiceItems *invoiceitem.Client
	// Disputes is the client used to invoke dispute-related APIs.
	// For more details see https://stripe.com/docs/api/#disputes.
	Disputes *dispute.Client
	// Transfers is the client used to invoke /transfers APIs.
	// For more details see https://stripe.com/docs/api/#transfers.
	Transfers *transfer.Client
	// Recipients is the client used to invoke /recipients APIs.
	// For more details see https://stripe.com/docs/api/#recipients.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	// For more details see https://stripe.com/docs/api#refunds.
	Refunds *refund.Client
	// Fees is the client used to invoke /application_fees APIs.
	// For more details see https://stripe.com/docs/api/#application_fees.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	// For more details see https://stripe.com/docs/api/#fee_refundss.
	FeeRefunds *feerefund.Client
	// Account is the client used to invoke /account APIs.
	// For more details see https://stripe.com/docs/api/#account.
	Account *account.Client
	// Balance is the client used to invoke /balance and transaction-related APIs.
	// For more details see https://stripe.com/docs/api/#balance.
	Balance *balance.Client
	// Events is the client used to invoke /events APIs.
	// For more details see https://stripe.com/docs/api#events.
	Events *event.Client
	// Tokens is the client used to invoke /tokens APIs.
	// For more details see https://stripe.com/docs/api/#Tokens.
	Tokens *token.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *Api) Init(key string, backend Backend) {
	if backend == nil {
		backend = GetBackend()
	}

	a.backend = backend
	a.key = key

	a.Charges = &charge.Client{B: a.backend, Key: a.key}
	a.Customers = &customer.Client{B: a.backend, Key: a.key}
	a.Cards = &card.Client{B: a.backend, Key: a.key}
	a.Subs = &sub.Client{B: a.backend, Key: a.key}
	a.Plans = &plan.Client{B: a.backend, Key: a.key}
	a.Coupons = &coupon.Client{B: a.backend, Key: a.key}
	a.Discounts = &discount.Client{B: a.backend, Key: a.key}
	a.Invoices = &invoice.Client{B: a.backend, Key: a.key}
	a.InvoiceItems = &invoiceitem.Client{B: a.backend, Key: a.key}
	a.Disputes = &dispute.Client{B: a.backend, Key: a.key}
	a.Transfers = &transfer.Client{B: a.backend, Key: a.key}
	a.Recipients = &recipient.Client{B: a.backend, Key: a.key}
	a.Refunds = &refund.Client{B: a.backend, Key: a.key}
	a.Fees = &fee.Client{B: a.backend, Key: a.key}
	a.FeeRefunds = &feerefund.Client{B: a.backend, Key: a.key}
	a.Account = &account.Client{B: a.backend, Key: a.key}
	a.Balance = &balance.Client{B: a.backend, Key: a.key}
	a.Events = &event.Client{B: a.backend, Key: a.key}
	a.Tokens = &token.Client{B: a.backend, Key: a.key}
}
