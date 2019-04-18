// Package creditnote provides the /credit_notes APIs
package creditnote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CreditNoteList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/credit_notes", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
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

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
