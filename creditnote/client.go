//
//
// File generated from our OpenAPI spec
//
//

// Package creditnote provides the /credit_notes APIs
package creditnote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /credit_notes APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new credit note.
func New(params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().New(params)
}

// New creates a new credit note.
func (c Client) New(params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/credit_notes",
		c.Key,
		params,
		creditnote,
	)
	return creditnote, err
}

// Get returns the details of a credit note.
func Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Get(id, params)
}

// Get returns the details of a credit note.
func (c Client) Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditnote)
	return creditnote, err
}

// Update updates a credit note's properties.
func Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Update(id, params)
}

// Update updates a credit note's properties.
func (c Client) Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// Preview is the method for the `GET /v1/credit_notes/preview` API.
func Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	return getC().Preview(params)
}

// Preview is the method for the `GET /v1/credit_notes/preview` API.
func (c Client) Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(
		http.MethodGet,
		"/v1/credit_notes/preview",
		c.Key,
		params,
		creditnote,
	)
	return creditnote, err
}

// VoidCreditNote is the method for the `POST /v1/credit_notes/{id}/void` API.
func VoidCreditNote(id string, params *stripe.CreditNoteVoidParams) (*stripe.CreditNote, error) {
	return getC().VoidCreditNote(id, params)
}

// VoidCreditNote is the method for the `POST /v1/credit_notes/{id}/void` API.
func (c Client) VoidCreditNote(id string, params *stripe.CreditNoteVoidParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s/void", id)
	creditnote := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditnote)
	return creditnote, err
}

// List returns a list of credit notes.
func List(params *stripe.CreditNoteListParams) *Iter {
	return getC().List(params)
}

// List returns a list of credit notes.
func (c Client) List(listParams *stripe.CreditNoteListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes", c.Key, b, p, list)

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

// ListLines is the method for the `GET /v1/credit_notes/{credit_note}/lines` API.
func ListLines(params *stripe.CreditNoteLineItemListParams) *LineItemIter {
	return getC().ListLines(params)
}

// ListLines is the method for the `GET /v1/credit_notes/{credit_note}/lines` API.
func (c Client) ListLines(listParams *stripe.CreditNoteLineItemListParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/credit_notes/%s/lines",
		stripe.StringValue(listParams.ID),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// ListPreviewLines is the method for the `GET /v1/credit_notes/preview/lines` API.
func ListPreviewLines(params *stripe.CreditNoteLineItemListPreviewParams) *LineItemIter {
	return getC().ListPreviewLines(params)
}

// ListPreviewLines is the method for the `GET /v1/credit_notes/preview/lines` API.
func (c Client) ListPreviewLines(listParams *stripe.CreditNoteLineItemListPreviewParams) *LineItemIter {
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CreditNoteLineItemList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes/preview/lines", c.Key, b, p, list)

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
