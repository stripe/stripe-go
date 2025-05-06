//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TestHelpersIssuingCardService is used to invoke /v1/issuing/cards APIs.
type v1TestHelpersIssuingCardService struct {
	B   Backend
	Key string
}

// Updates the shipping status of the specified Issuing Card object to delivered.
func (c v1TestHelpersIssuingCardService) DeliverCard(ctx context.Context, id string, params *TestHelpersIssuingCardDeliverCardParams) (*IssuingCard, error) {
	if params == nil {
		params = &TestHelpersIssuingCardDeliverCardParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/deliver", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to failure.
func (c v1TestHelpersIssuingCardService) FailCard(ctx context.Context, id string, params *TestHelpersIssuingCardFailCardParams) (*IssuingCard, error) {
	if params == nil {
		params = &TestHelpersIssuingCardFailCardParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/cards/%s/shipping/fail", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to returned.
func (c v1TestHelpersIssuingCardService) ReturnCard(ctx context.Context, id string, params *TestHelpersIssuingCardReturnCardParams) (*IssuingCard, error) {
	if params == nil {
		params = &TestHelpersIssuingCardReturnCardParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/cards/%s/shipping/return", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to shipped.
func (c v1TestHelpersIssuingCardService) ShipCard(ctx context.Context, id string, params *TestHelpersIssuingCardShipCardParams) (*IssuingCard, error) {
	if params == nil {
		params = &TestHelpersIssuingCardShipCardParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/cards/%s/shipping/ship", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
func (c v1TestHelpersIssuingCardService) SubmitCard(ctx context.Context, id string, params *TestHelpersIssuingCardSubmitCardParams) (*IssuingCard, error) {
	if params == nil {
		params = &TestHelpersIssuingCardSubmitCardParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/cards/%s/shipping/submit", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}
