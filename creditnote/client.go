// Package creditnote provides the /credit_notes APIs
package creditnote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is the client used to invoke /credit_notes APIs.
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
	cn := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, "/v1/credit_notes", c.Key, params, cn)
	return cn, err
}

// Get returns the details of a credit note.
func Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Get(id, params)
}

// Get returns the details of a credit note.
func (c Client) Get(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	cn := &stripe.CreditNote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cn)
	return cn, err
}

// Update updates a credit note.
func Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	return getC().Update(id, params)
}

// Update updates a credit note.
func (c Client) Update(id string, params *stripe.CreditNoteParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s", id)
	cn := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cn)
	return cn, err
}

// List returns a list of credit notes.
func List(params *stripe.CreditNoteListParams) *Iter {
	return getC().List(params)
}

// List returns a list of credit notes.
func (c Client) List(listParams *stripe.CreditNoteListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.CreditNoteList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// ListLines returns a list of credit note line items on a credit note.
func ListLines(params *stripe.CreditNoteLineItemListParams) *LineItemIter {
	return getC().ListLines(params)
}

// ListLines returns a list of credit note line items on a credit note.
func (c Client) ListLines(listParams *stripe.CreditNoteLineItemListParams) *LineItemIter {
	path := stripe.FormatURLPath("/v1/credit_notes/%s/lines", stripe.StringValue(listParams.ID))
	return &LineItemIter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.CreditNoteLineItemList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// ListPreviewLines returns a list of lines on a previewed credit note.
func ListPreviewLines(params *stripe.CreditNoteLineItemListPreviewParams) *LineItemIter {
	return getC().ListPreviewLines(params)
}

// ListPreviewLines returns a list of lines on a previewed credit note.
func (c Client) ListPreviewLines(listParams *stripe.CreditNoteLineItemListPreviewParams) *LineItemIter {
	return &LineItemIter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.CreditNoteLineItemList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes/preview/lines", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Preview previews a credit note.
func Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	return getC().Preview(params)
}

// Preview previews a credit note.
func (c Client) Preview(params *stripe.CreditNotePreviewParams) (*stripe.CreditNote, error) {
	cn := &stripe.CreditNote{}
	err := c.B.Call(http.MethodGet, "/v1/credit_notes/preview", c.Key, params, cn)
	return cn, err
}

// VoidCreditNote voids a credit note.
func VoidCreditNote(id string, params *stripe.CreditNoteVoidParams) (*stripe.CreditNote, error) {
	return getC().VoidCreditNote(id, params)
}

// VoidCreditNote voids a credit note.
func (c Client) VoidCreditNote(id string, params *stripe.CreditNoteVoidParams) (*stripe.CreditNote, error) {
	path := stripe.FormatURLPath("/v1/credit_notes/%s/void", id)
	cn := &stripe.CreditNote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cn)
	return cn, err
}

// Iter is an iterator for credit notes.
type Iter struct {
	*stripe.Iter
}

// CreditNote returns the cn which the iterator is currently pointing to.
func (i *Iter) CreditNote() *stripe.CreditNote {
	return i.Current().(*stripe.CreditNote)
}

// CreditNoteList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CreditNoteList() *stripe.CreditNoteList {
	return i.List().(*stripe.CreditNoteList)
}

// LineItemIter is an iterator for credit note line items on a credit note.
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
func (i *Iter) CreditNoteLineItemList() *stripe.CreditNoteLineItemList {
	return i.List().(*stripe.CreditNoteLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
