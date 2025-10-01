//
//
// File generated from our OpenAPI spec
//
//

// Package creditnote provides the /v1/credit_notes APIs
package creditnote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/form"
)

// Client is used to invoke /v1/credit_notes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice's amount_remaining (and amount_due), but not below zero.
// This amount is indicated by the credit note's pre_payment_amount. The excess amount is indicated by post_payment_amount, and it can result in any combination of the following:
//
// Refunds: create a new refund (using refund_amount) or link existing refunds (using refunds).
// Customer balance credit: credit the customer's balance (using credit_amount) which will be automatically applied to their next invoice when it's finalized.
// Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using out_of_band_amount).
//
// The sum of refunds, customer balance credits, and outside of Stripe credits must equal the post_payment_amount.
//
// You may issue multiple credit notes for an invoice. Each credit note may increment the invoice's pre_payment_credit_notes_amount,
// post_payment_credit_notes_amount, or both, depending on the invoice's amount_remaining at the time of credit note creation.
func New(params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().New(params)
}

// Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice's amount_remaining (and amount_due), but not below zero.
// This amount is indicated by the credit note's pre_payment_amount. The excess amount is indicated by post_payment_amount, and it can result in any combination of the following:
//
// Refunds: create a new refund (using refund_amount) or link existing refunds (using refunds).
// Customer balance credit: credit the customer's balance (using credit_amount) which will be automatically applied to their next invoice when it's finalized.
// Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using out_of_band_amount).
//
// The sum of refunds, customer balance credits, and outside of Stripe credits must equal the post_payment_amount.
//
// You may issue multiple credit notes for an invoice. Each credit note may increment the invoice's pre_payment_credit_notes_amount,
// post_payment_credit_notes_amount, or both, depending on the invoice's amount_remaining at the time of credit note creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(
		http.MethodPost, "/v1/credit_notes", c.Key, params, creditnote)
	return creditnote, err
}

// Retrieves the credit note object with the given identifier.
func Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Get(id, params)
}

// Retrieves the credit note object with the given identifier.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditnote)
	return creditnote, err
}

// Updates an existing credit note.
func Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Update(id, params)
}

// Updates an existing credit note.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// Get a preview of a credit note without creating it.
func Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	return getC().Preview(params)
}

// Get a preview of a credit note without creating it.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(
		http.MethodGet, "/v1/credit_notes/preview", c.Key, params, creditnote)
	return creditnote, err
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://docs.stripe.com/docs/billing/invoices/credit-notes#voiding).
func VoidCreditNote(id string, params *stripe.CreditNoteVoidCreditNoteParams) (*stripe.CreditNote, error) {
	return getC().VoidCreditNote(id, params)
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://docs.stripe.com/docs/billing/invoices/credit-notes#voiding).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VoidCreditNote(id string, params *stripe.CreditNoteVoidCreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s/void", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// Returns a list of credit notes.
func List(params *stripe.CreditNoteListParams) *Iter {
	return getC().List(params)
}

// Returns a list of credit notes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CreditNoteListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for credit notes.
type Iter struct {
	*stripe.Iter
}

// CreditNote returns the credit note which the iterator is currently pointing to.
func (i *Iter) CreditNote() *stripe.CreditNote {
	return i.Current().(*stripe.CreditNote)
}

// CreditNoteList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CreditNoteList() *stripe.CreditNoteList {
	return i.List().(*stripe.CreditNoteList)
}

// When retrieving a credit note, you'll get a lines property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLines(params *stripe.CreditNoteListLinesParams) *LineItemIter {
	return getC().ListLines(params)
}

// When retrieving a credit note, you'll get a lines property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLines(listParams *stripe.CreditNoteListLinesParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/credit_notes/%s/lines", stripe.StringValue(listParams.CreditNote))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// When retrieving a credit note preview, you'll get a lines property containing the first handful of those items. This URL you can retrieve the full (paginated) list of line items.
func PreviewLines(params *stripe.CreditNotePreviewLinesParams) *LineItemIter {
	return getC().PreviewLines(params)
}

// When retrieving a credit note preview, you'll get a lines property containing the first handful of those items. This URL you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) PreviewLines(listParams *stripe.CreditNotePreviewLinesParams) *LineItemIter {
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteLineItemList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes/preview/lines", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineItemIter is an iterator for credit note line items.
type LineItemIter struct {
	*stripe.Iter
}

// CreditNoteLineItem returns the credit note line item which the iterator is currently pointing to.
func (i *LineItemIter) CreditNoteLineItem() *stripe.CreditNoteLineItem {
	return i.Current().(*stripe.CreditNoteLineItem)
}

// CreditNoteLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) CreditNoteLineItemList() *stripe.CreditNoteLineItemList {
	return i.List().(*stripe.CreditNoteLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
