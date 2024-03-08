//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /issuing/personalization_designs APIs
package personalizationdesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /issuing/personalization_designs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Activate is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/activate` API.
func Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Activate(id, params)
}

// Activate is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/activate` API.
func (c Client) Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/activate",
		id,
	)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Deactivate is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/deactivate` API.
func Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Deactivate(id, params)
}

// Deactivate is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/deactivate` API.
func (c Client) Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/deactivate",
		id,
	)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Reject is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/reject` API.
func Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Reject(id, params)
}

// Reject is the method for the `POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/reject` API.
func (c Client) Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/reject",
		id,
	)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
