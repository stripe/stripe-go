package product

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /products APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func New(params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().New(params)
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func (c Client) New(params *stripe.ProductParams) (*stripe.Product, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	p := &stripe.Product{}
	err := c.B.Call("POST", "/products", c.Key, body, commonParams, p)

	return p, err
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Update(id, params)
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func (c Client) Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	p := &stripe.Product{}
	err := c.B.Call("POST", "/products/"+id, c.Key, body, commonParams, p)

	return p, err
}

// Get returns the details of an product
// For more details see https://stripe.com/docs/api#retrieve_product.
func Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	p := &stripe.Product{}
	err := c.B.Call("GET", "/products/"+id, c.Key, body, commonParams, p)

	return p, err
}

// List returns a list of products.
// For more details see https://stripe.com/docs/api#list_products
func List(params *stripe.ProductListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.ProductListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ProductList{}
		err := c.B.Call("GET", "/products", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Products.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Product returns the most recent Product
// visited by a call to Next.
func (i *Iter) Product() *stripe.Product {
	return i.Current().(*stripe.Product)
}

// Delete deletes a product
// For more details see https://stripe.com/docs/api#delete_product.
func Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Del(id, params)
}

// Delete deletes a product.
// For more details see https://stripe.com/docs/api#delete_product.
func (c Client) Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	p := &stripe.Product{}
	err := c.B.Call("DELETE", "/products/"+id, c.Key, body, commonParams, p)

	return p, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
