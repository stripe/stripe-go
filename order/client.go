package product

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe-internal/stripe-go"
)

// Client is used to invoke /products APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func New(params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().New(params)
}

// New POSTs a new product.
// For more details see https://stripe.com/docs/api#create_product.
func (c Client) New(params *stripe.OrderParams) (*stripe.Order, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &url.Values{}

		// Required fields
		body.Add("currency", params.Currency)

		if params.Customer != "" {
			body.Add("customer", params.Customer)
		}

		if params.Email != "" {
			body.Add("email", params.Email)
		}

		if params.Email != "" {
			body.Add("email", params.Email)
		}

		if len(params.Items) > 0 {
			for _, item := range params.Items {
				body.Add("items[][description]", item.Description)
				body.Add("items[][type]", item.Type)
				body.Add("items[][amount]", strconv.FormatInt(item.Amount, 10))
				body.Add("items[][currency]", item.Currency)
				if item.Parent != "" {
					body.Add("items[][parent]", item.Parent)
				}
				if item.Quantity != nil {
					body.Add("items[][quantity]", strconv.FormatInt(*item.Quantity, 10))
				}
			}
		}

		body.Add("shipping[address][line1]", params.Shipping.Address.Line1)
		if params.Shipping.Address.Line2 != "" {
			body.Add("shipping[address][line2]", params.Shipping.Address.Line2)
		}
		if params.Shipping.Address.City != "" {
			body.Add("shipping[address][city]", params.Shipping.Address.City)
		}
		if params.Shipping.Address.Country != "" {
			body.Add("shipping[address][country]", params.Shipping.Address.Country)
		}
		if params.Shipping.Address.Zip != "" {
			body.Add("shipping[address][postal_code]", params.Shipping.Address.Zip)
		}
		if params.Shipping.Address.State != "" {
			body.Add("shipping[address][state]", params.Shipping.Address.State)
		}

		if params.Shipping.Name != "" {
			body.Add("shipping[address][name]", params.Shipping.Name)
		}
		if params.Shipping.Phone != "" {
			body.Add("shipping[address][phone]", params.Shipping.Phone)
		}

		params.AppendTo(body)
	}

	p := &stripe.Order{}
	err := c.B.Call("POST", "/orders", c.Key, body, commonParams, p)

	return p, err
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Update(id, params)
}

// Update updates a product's properties.
// For more details see https://stripe.com/docs/api#update_product.
func (c Client) Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &url.Values{}

		body.Add("id", params.ID)

		params.AppendTo(body)
	}

	p := &stripe.Order{}
	err := c.B.Call("POST", "/products/"+id, c.Key, body, commonParams, p)

	return p, err
}

// Get returns the details of an product
// For more details see https://stripe.com/docs/api#retrieve_product.
func Get(id string) (*stripe.Order, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*stripe.Order, error) {
	product := &stripe.Order{}
	err := c.B.Call("GET", "/products/"+id, c.Key, nil, nil, product)

	return product, err
}

// List returns a list of products.
// For more details see https://stripe.com/docs/api#list_products
func List(params *stripe.OrderListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.OrderListParams) *Iter {
	type productList struct {
		stripe.ListMeta
		Values []*stripe.Order `json:"data"`
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
		list := &productList{}
		err := c.B.Call("GET", "/products", c.Key, &b, nil, list)

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
func (i *Iter) Product() *stripe.Order {
	return i.Current().(*stripe.Order)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
