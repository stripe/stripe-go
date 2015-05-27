package stripe

import "encoding/json"

// ProductParams is the set of parameters that can be used
// when creating or updating a product.
// For more details, see https://stripe.com/docs/api#create_product
// and https://stripe.com/docs/api#update_product.
type ProductParams struct {
	Params
	ID        string
	Active    *bool
	Name      string
	Caption   string
	Desc      string
	Attrs     []string
	Images    []string `json:"images"`
	Shippable *bool
}

// Product is the resource representing a Stripe product.
// For more details see https://stripe.com/docs/api#products.
type Product struct {
	ID        string            `json:"id"`
	Created   int64             `json:"created"`
	Updated   int64             `json:"updated"`
	Live      bool              `json:"livemode"`
	Active    bool              `json:"active"`
	Name      string            `json:"name"`
	Caption   string            `json:"caption"`
	Desc      string            `json:"description"`
	Attrs     []string          `json:"attributes"`
	Shippable bool              `json:"shippable"`
	Images    []string          `json:"images"`
	Meta      map[string]string `json:"metdata"`
	Skus      *SKUList          `json:"skus"`
}

// ProductListParams is the set of parameters that can be used when
// listing products. For more details, see:
// https://stripe.com/docs/api#list_products.
type ProductListParams struct {
	ListParams
	Created int64
}

// UnmarshalJSON handles deserialization of a Product.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Product) UnmarshalJSON(data []byte) error {
	type product Product
	var pr product
	err := json.Unmarshal(data, &pr)
	if err == nil {
		*p = Product(pr)
	} else {
		p.ID = string(data[1 : len(data)-1])
	}

	return nil
}
