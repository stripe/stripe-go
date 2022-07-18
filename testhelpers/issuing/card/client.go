//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /issuing/cards APIs
package card

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /issuing/cards APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// DeliverCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/deliver` API.
func DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	return getC().DeliverCard(id, params)
}

// DeliverCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/deliver` API.
func (c Client) DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/deliver",
		id,
	)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// FailCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/fail` API.
func FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	return getC().FailCard(id, params)
}

// FailCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/fail` API.
func (c Client) FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/fail",
		id,
	)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// ReturnCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/return` API.
func ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	return getC().ReturnCard(id, params)
}

// ReturnCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/return` API.
func (c Client) ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/return",
		id,
	)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// ShipCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/ship` API.
func ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	return getC().ShipCard(id, params)
}

// ShipCard is the method for the `POST /v1/test_helpers/issuing/cards/{card}/shipping/ship` API.
func (c Client) ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/ship",
		id,
	)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
