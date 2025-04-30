//
//
// File generated from our OpenAPI spec
//
//

// Package reader provides the /v1/terminal/readers APIs
package reader

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/terminal/readers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Reader object.
func New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().New(params)
}

// Creates a new Reader object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	reader := &stripe.TerminalReader{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/readers", c.Key, params, reader)
	return reader, err
}

// Retrieves a Reader object.
func Get(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().Get(id, params)
}

// Retrieves a Reader object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reader)
	return reader, err
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().Update(id, params)
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Deletes a Reader object.
func Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	return getC().Del(id, params)
}

// Deletes a Reader object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, reader)
	return reader, err
}

// Cancels the current reader action.
func CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	return getC().CancelAction(id, params)
}

// Cancels the current reader action.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/cancel_action", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates an input collection flow on a Reader.
func CollectInputs(id string, params *stripe.TerminalReaderCollectInputsParams) (*stripe.TerminalReader, error) {
	return getC().CollectInputs(id, params)
}

// Initiates an input collection flow on a Reader.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CollectInputs(id string, params *stripe.TerminalReaderCollectInputsParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/collect_inputs", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation.
func CollectPaymentMethod(id string, params *stripe.TerminalReaderCollectPaymentMethodParams) (*stripe.TerminalReader, error) {
	return getC().CollectPaymentMethod(id, params)
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CollectPaymentMethod(id string, params *stripe.TerminalReaderCollectPaymentMethodParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/collect_payment_method", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Finalizes a payment on a Reader.
func ConfirmPaymentIntent(id string, params *stripe.TerminalReaderConfirmPaymentIntentParams) (*stripe.TerminalReader, error) {
	return getC().ConfirmPaymentIntent(id, params)
}

// Finalizes a payment on a Reader.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ConfirmPaymentIntent(id string, params *stripe.TerminalReaderConfirmPaymentIntentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/confirm_payment_intent", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a payment flow on a Reader.
func ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	return getC().ProcessPaymentIntent(id, params)
}

// Initiates a payment flow on a Reader.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/process_payment_intent", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a setup intent flow on a Reader.
func ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	return getC().ProcessSetupIntent(id, params)
}

// Initiates a setup intent flow on a Reader.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/terminal/readers/%s/process_setup_intent", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a refund on a Reader
func RefundPayment(id string, params *stripe.TerminalReaderRefundPaymentParams) (*stripe.TerminalReader, error) {
	return getC().RefundPayment(id, params)
}

// Initiates a refund on a Reader
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RefundPayment(id string, params *stripe.TerminalReaderRefundPaymentParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/refund_payment", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Sets reader display to show cart details.
func SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	return getC().SetReaderDisplay(id, params)
}

// Sets reader display to show cart details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath("/v1/terminal/readers/%s/set_reader_display", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Returns a list of Reader objects.
func List(params *stripe.TerminalReaderListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Reader objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TerminalReaderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalReaderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/terminal/readers", c.Key, []byte(b.Encode()), p, list)

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
