//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1CreditNoteService is used to invoke /v1/credit_notes APIs.
type v1CreditNoteService struct {
	B   Backend
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
func (c v1CreditNoteService) Create(ctx context.Context, params *CreditNoteCreateParams) (*CreditNote, error) {
	if params == nil {
		params = &CreditNoteCreateParams{}
	}
	params.Context = ctx
	creditnote := &CreditNote{}
	err := c.B.Call(
		http.MethodPost, "/v1/credit_notes", c.Key, params, creditnote)
	return creditnote, err
}

// Retrieves the credit note object with the given identifier.
func (c v1CreditNoteService) Retrieve(ctx context.Context, id string, params *CreditNoteRetrieveParams) (*CreditNote, error) {
	if params == nil {
		params = &CreditNoteRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &CreditNote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditnote)
	return creditnote, err
}

// Updates an existing credit note.
func (c v1CreditNoteService) Update(ctx context.Context, id string, params *CreditNoteUpdateParams) (*CreditNote, error) {
	if params == nil {
		params = &CreditNoteUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// Get a preview of a credit note without creating it.
func (c v1CreditNoteService) Preview(ctx context.Context, params *CreditNotePreviewParams) (*CreditNote, error) {
	if params == nil {
		params = &CreditNotePreviewParams{}
	}
	params.Context = ctx
	creditnote := &CreditNote{}
	err := c.B.Call(
		http.MethodGet, "/v1/credit_notes/preview", c.Key, params, creditnote)
	return creditnote, err
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://docs.stripe.com/docs/billing/invoices/credit-notes#voiding).
func (c v1CreditNoteService) VoidCreditNote(ctx context.Context, id string, params *CreditNoteVoidCreditNoteParams) (*CreditNote, error) {
	if params == nil {
		params = &CreditNoteVoidCreditNoteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/credit_notes/%s/void", id)
	creditnote := &CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// Returns a list of credit notes.
func (c v1CreditNoteService) List(ctx context.Context, listParams *CreditNoteListParams) Seq2[*CreditNote, error] {
	if listParams == nil {
		listParams = &CreditNoteListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CreditNote, ListContainer, error) {
		list := &CreditNoteList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a credit note, you'll get a lines property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1CreditNoteService) ListLines(ctx context.Context, listParams *CreditNoteListLinesParams) Seq2[*CreditNoteLineItem, error] {
	if listParams == nil {
		listParams = &CreditNoteListLinesParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/credit_notes/%s/lines", StringValue(listParams.CreditNote))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CreditNoteLineItem, ListContainer, error) {
		list := &CreditNoteLineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a credit note preview, you'll get a lines property containing the first handful of those items. This URL you can retrieve the full (paginated) list of line items.
func (c v1CreditNoteService) PreviewLines(ctx context.Context, listParams *CreditNotePreviewLinesParams) Seq2[*CreditNoteLineItem, error] {
	if listParams == nil {
		listParams = &CreditNotePreviewLinesParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CreditNoteLineItem, ListContainer, error) {
		list := &CreditNoteLineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes/preview/lines", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
