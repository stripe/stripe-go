// Package review provides the /reviews APIs
package review

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /reviews APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Approve approves a review.
func Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	return getC().Approve(id, params)
}

// Approve approves a review.
func (c Client) Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	path := stripe.FormatURLPath("/v1/reviews/%s/approve", id)
	review := &stripe.Review{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, review)
	return review, err
}

// Get returns the details of a review.
func Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	return getC().Get(id, params)
}

// Get returns the details of a review.
func (c Client) Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	path := stripe.FormatURLPath("/v1/reviews/%s", id)
	review := &stripe.Review{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, review)
	return review, err
}

// List returns a list of reviews.
func List(params *stripe.ReviewListParams) *Iter {
	return getC().List(params)
}

// List returns a list of reviews.
func (c Client) List(listParams *stripe.ReviewListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.ReviewList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/reviews", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for reviews.
type Iter struct {
	*stripe.Iter
}

// Review returns the review which the iterator is currently pointing to.
func (i *Iter) Review() *stripe.Review {
	return i.Current().(*stripe.Review)
}

// ReviewList returns the current list object which the iterator is currently
// using. List objects will change as new API calls are made to continue
// pagination.
func (i *Iter) ReviewList() *stripe.ReviewList {
	return i.List().(*stripe.ReviewList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
