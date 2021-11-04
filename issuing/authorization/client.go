//
//
// File generated from our OpenAPI spec
//
//

// Package authorization provides the /issuing/authorizations APIs
package authorization

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /issuing/authorizations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuing authorization.
func Get(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing authorization.
func (c Client) Get(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath("/v1/issuing/authorizations/%s", id)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, authorization)
	return authorization, err
}

// Update updates an issuing authorization's properties.
func Update(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	return getC().Update(id, params)
}

// Update updates an issuing authorization's properties.
func (c Client) Update(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath("/v1/issuing/authorizations/%s", id)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Approve is the method for the `POST /v1/issuing/authorizations/{authorization}/approve` API.
func Approve(id string, params *stripe.IssuingAuthorizationApproveParams) (*stripe.IssuingAuthorization, error) {
	return getC().Approve(id, params)
}

// Approve is the method for the `POST /v1/issuing/authorizations/{authorization}/approve` API.
func (c Client) Approve(id string, params *stripe.IssuingAuthorizationApproveParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath("/v1/issuing/authorizations/%s/approve", id)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Decline is the method for the `POST /v1/issuing/authorizations/{authorization}/decline` API.
func Decline(id string, params *stripe.IssuingAuthorizationDeclineParams) (*stripe.IssuingAuthorization, error) {
	return getC().Decline(id, params)
}

// Decline is the method for the `POST /v1/issuing/authorizations/{authorization}/decline` API.
func (c Client) Decline(id string, params *stripe.IssuingAuthorizationDeclineParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath("/v1/issuing/authorizations/%s/decline", id)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// List returns a list of issuing authorizations.
func List(params *stripe.IssuingAuthorizationListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing authorizations.
func (c Client) List(listParams *stripe.IssuingAuthorizationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingAuthorizationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/authorizations", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing authorizations.
type Iter struct {
	*stripe.Iter
}

// IssuingAuthorization returns the issuing authorization which the iterator is currently pointing to.
func (i *Iter) IssuingAuthorization() *stripe.IssuingAuthorization {
	return i.Current().(*stripe.IssuingAuthorization)
}

// IssuingAuthorizationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingAuthorizationList() *stripe.IssuingAuthorizationList {
	return i.List().(*stripe.IssuingAuthorizationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
