package topup

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /topups APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new topups.
// For more details see https://stripe.com/docs/api#create_topup.
func New(params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.TopupParams) (*stripe.Topup, error) {
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodPost, "/topups", c.Key, params, topup)
	return topup, err
}

// Get returns the details of a topup.
// For more details see https://stripe.com/docs/api#retrieve_topup.
func Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	path := stripe.FormatURLPath("/topups/%s", id)
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, topup)
	return topup, err
}

// Update updates a topup's properties.
// For more details see https://stripe.com/docs/api#update_topup.
func Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	path := stripe.FormatURLPath("/topups/%s", id)
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// List returns a list of topups.
// For more details see https://stripe.com/docs/api#list_topups.
func List(params *stripe.TopupListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.TopupListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.TopupList{}
		err := c.B.CallRaw(http.MethodGet, "/topups", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Topups.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
