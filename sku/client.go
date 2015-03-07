package sku

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /skus APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an sku
// For more details see https://stripe.com/docs/api#retrieve_sku.
func Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	sku := &stripe.SKU{}
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}
	err := c.B.Call("GET", "/skus/"+id, c.Key, body, commonParams, sku)

	return sku, err
}

// List returns a list of skus.
// For more details see https://stripe.com/docs/api#list_skus
func List(params *stripe.SKUListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.SKUListParams) *Iter {
	type skuList struct {
		stripe.ListMeta
		Values []*stripe.SKU `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &skuList{}
		err := c.B.Call("GET", "/skus", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of SKUs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// SKU returns the most recent SKU
// visited by a call to Next.
func (i *Iter) SKU() *stripe.SKU {
	return i.Current().(*stripe.SKU)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
