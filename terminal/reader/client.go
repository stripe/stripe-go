//
//
// File generated from our OpenAPI spec
//
//

// Package reader provides the /terminal/readers APIs
package reader

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /terminal/readers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new terminal reader.
func New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().New(params)
}

// New creates a new terminal reader.
func (c Client) New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	reader := &stripe.TerminalReader{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/terminal/readers",
		c.Key,
		params,
		reader,
	)
	return reader, err
}

// Get returns the details of a terminal reader.
func Get(id string, params *stripe.TerminalReaderGetParams) (*stripe.TerminalReader, error) {
	return getC().Get(id, params)
}

// Get returns the details of a terminal reader.
func (c Client) Get(id string, params *stripe.TerminalReaderGetParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reader)
	return reader, err
}

// Update updates a terminal reader's properties.
func Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().Update(id, params)
}

// Update updates a terminal reader's properties.
func (c Client) Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Del removes a terminal reader.
func Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().Del(id, params)
}

// Del removes a terminal reader.
func (c Client) Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, reader)
	return reader, err
}

// CancelAction is the method for the `POST /v1/terminal/readers/{reader}/cancel_action` API.
func CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	return getC().CancelAction(id, params)
}

// CancelAction is the method for the `POST /v1/terminal/readers/{reader}/cancel_action` API.
func (c Client) CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/cancel_action", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// ProcessPaymentIntent is the method for the `POST /v1/terminal/readers/{reader}/process_payment_intent` API.
func ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	return getC().ProcessPaymentIntent(id, params)
}

// ProcessPaymentIntent is the method for the `POST /v1/terminal/readers/{reader}/process_payment_intent` API.
func (c Client) ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/process_payment_intent",
		id,
	)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// ProcessSetupIntent is the method for the `POST /v1/terminal/readers/{reader}/process_setup_intent` API.
func ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	return getC().ProcessSetupIntent(id, params)
}

// ProcessSetupIntent is the method for the `POST /v1/terminal/readers/{reader}/process_setup_intent` API.
func (c Client) ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/process_setup_intent",
		id,
	)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// SetReaderDisplay is the method for the `POST /v1/terminal/readers/{reader}/set_reader_display` API.
func SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	return getC().SetReaderDisplay(id, params)
}

// SetReaderDisplay is the method for the `POST /v1/terminal/readers/{reader}/set_reader_display` API.
func (c Client) SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/set_reader_display", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// List returns a list of terminal readers.
func List(params *stripe.TerminalReaderListParams) *Iter {
	return getC().List(params)
}

// List returns a list of terminal readers.
func (c Client) List(listParams *stripe.TerminalReaderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalReaderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/terminal/readers", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for terminal readers.
type Iter struct {
	*stripe.Iter
}

// TerminalReader returns the terminal reader which the iterator is currently pointing to.
func (i *Iter) TerminalReader() *stripe.TerminalReader {
	return i.Current().(*stripe.TerminalReader)
}

// TerminalReaderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalReaderList() *stripe.TerminalReaderList {
	return i.List().(*stripe.TerminalReaderList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
