//
//
// File generated from our OpenAPI spec
//
//

// Package carddesign provides the /issuing/card_designs APIs
package carddesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
)

// Client is used to invoke /issuing/card_designs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// ActivateTestmode is the method for the `POST /v1/test_helpers/issuing/card_designs/{card_design}/status/activate` API.
func ActivateTestmode(id string, params *stripe.TestHelpersIssuingCardDesignActivateTestmodeParams) (*stripe.IssuingCardDesign, error) {
	return getC().ActivateTestmode(id, params)
}

// ActivateTestmode is the method for the `POST /v1/test_helpers/issuing/card_designs/{card_design}/status/activate` API.
func (c Client) ActivateTestmode(id string, params *stripe.TestHelpersIssuingCardDesignActivateTestmodeParams) (*stripe.IssuingCardDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/card_designs/%s/status/activate",
		id,
	)
	carddesign := &stripe.IssuingCardDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, carddesign)
	return carddesign, err
}

// DeactivateTestmode is the method for the `POST /v1/test_helpers/issuing/card_designs/{card_design}/status/deactivate` API.
func DeactivateTestmode(id string, params *stripe.TestHelpersIssuingCardDesignDeactivateTestmodeParams) (*stripe.IssuingCardDesign, error) {
	return getC().DeactivateTestmode(id, params)
}

// DeactivateTestmode is the method for the `POST /v1/test_helpers/issuing/card_designs/{card_design}/status/deactivate` API.
func (c Client) DeactivateTestmode(id string, params *stripe.TestHelpersIssuingCardDesignDeactivateTestmodeParams) (*stripe.IssuingCardDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/card_designs/%s/status/deactivate",
		id,
	)
	carddesign := &stripe.IssuingCardDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, carddesign)
	return carddesign, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
