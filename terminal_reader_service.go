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

// v1TerminalReaderService is used to invoke /v1/terminal/readers APIs.
type v1TerminalReaderService struct {
	B   Backend
	Key string
}

// Creates a new Reader object.
func (c v1TerminalReaderService) Create(ctx context.Context, params *TerminalReaderCreateParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderCreateParams{}
	}
	params.Context = ctx
	reader := &TerminalReader{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/readers", c.Key, params, reader)
	return reader, err
}

// Retrieves a Reader object.
func (c v1TerminalReaderService) Retrieve(ctx context.Context, id string, params *TerminalReaderRetrieveParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reader)
	return reader, err
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1TerminalReaderService) Update(ctx context.Context, id string, params *TerminalReaderUpdateParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Deletes a Reader object.
func (c v1TerminalReaderService) Delete(ctx context.Context, id string, params *TerminalReaderDeleteParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, reader)
	return reader, err
}

// Cancels the current reader action.
func (c v1TerminalReaderService) CancelAction(ctx context.Context, id string, params *TerminalReaderCancelActionParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderCancelActionParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/cancel_action", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates an input collection flow on a Reader.
func (c v1TerminalReaderService) CollectInputs(ctx context.Context, id string, params *TerminalReaderCollectInputsParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderCollectInputsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/collect_inputs", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation.
func (c v1TerminalReaderService) CollectPaymentMethod(ctx context.Context, id string, params *TerminalReaderCollectPaymentMethodParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderCollectPaymentMethodParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/collect_payment_method", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Finalizes a payment on a Reader.
func (c v1TerminalReaderService) ConfirmPaymentIntent(ctx context.Context, id string, params *TerminalReaderConfirmPaymentIntentParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderConfirmPaymentIntentParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/confirm_payment_intent", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a payment flow on a Reader.
func (c v1TerminalReaderService) ProcessPaymentIntent(ctx context.Context, id string, params *TerminalReaderProcessPaymentIntentParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderProcessPaymentIntentParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/process_payment_intent", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a setup intent flow on a Reader.
func (c v1TerminalReaderService) ProcessSetupIntent(ctx context.Context, id string, params *TerminalReaderProcessSetupIntentParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderProcessSetupIntentParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/process_setup_intent", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Initiates a refund on a Reader
func (c v1TerminalReaderService) RefundPayment(ctx context.Context, id string, params *TerminalReaderRefundPaymentParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderRefundPaymentParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/refund_payment", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Sets reader display to show cart details.
func (c v1TerminalReaderService) SetReaderDisplay(ctx context.Context, id string, params *TerminalReaderSetReaderDisplayParams) (*TerminalReader, error) {
	if params == nil {
		params = &TerminalReaderSetReaderDisplayParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/readers/%s/set_reader_display", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Returns a list of Reader objects.
func (c v1TerminalReaderService) List(ctx context.Context, listParams *TerminalReaderListParams) Seq2[*TerminalReader, error] {
	if listParams == nil {
		listParams = &TerminalReaderListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TerminalReader, ListContainer, error) {
		list := &TerminalReaderList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/terminal/readers", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
