//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /v1/issuing/personalization_designs APIs
package personalizationdesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/issuing/personalization_designs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates the status of the specified testmode personalization design object to active.
func Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Activate(id, params)
}

// Updates the status of the specified testmode personalization design object to active.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/activate", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates the status of the specified testmode personalization design object to inactive.
func Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Deactivate(id, params)
}

// Updates the status of the specified testmode personalization design object to inactive.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/deactivate", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates the status of the specified testmode personalization design object to rejected.
func Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Reject(id, params)
}

// Updates the status of the specified testmode personalization design object to rejected.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/reject", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
