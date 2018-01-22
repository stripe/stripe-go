package stripe

import "encoding/json"

// ProductType is the type of a product.
type ProductType string

const (
	// ProductTypeGood is a constant that indicates a product represents a physical good,
	// which may be sold through the Stripe Relay API.
	ProductTypeGood ProductType = "good"

	// ProductTypeService is a constant that indicates a product represents a service
	// which is provided on a recurring basis and is priced with a Stripe plan.
	ProductTypeService ProductType = "service"
)

// PackageDimensions represents the dimension of a product or a sku from the
// perspective of shipping.
type PackageDimensions struct {
	Height float64 `json:"height" form:"height"`
	Length float64 `json:"length" form:"length"`
	Weight float64 `json:"weight" form:"weight"`
	Width  float64 `json:"width" form:"width"`
}

// ProductParams is the set of parameters that can be used
// when creating or updating a product.
// For more details, see https://stripe.com/docs/api#create_product
// and https://stripe.com/docs/api#update_product.
type ProductParams struct {
	Params              `form:"*"`
	Active              *bool              `form:"active"`
	Attrs               []string           `form:"attributes"`
	Caption             string             `form:"caption"`
	DeactivateOn        []string           `form:"deactivate_on"`
	Desc                string             `form:"description"`
	ID                  string             `form:"id"`
	Images              []string           `form:"images"`
	Name                string             `form:"name"`
	PackageDimensions   *PackageDimensions `form:"package_dimensions"`
	Shippable           *bool              `form:"shippable"`
	StatementDescriptor string             `form:"statement_descriptor"`
	Type                ProductType        `form:"type"`
	URL                 string             `form:"url"`
}

// Product is the resource representing a Stripe product.
// For more details see https://stripe.com/docs/api#products.
type Product struct {
	Active              bool               `json:"active"`
	Attrs               []string           `json:"attributes"`
	Caption             string             `json:"caption"`
	Created             int64              `json:"created"`
	DeactivateOn        []string           `json:"deactivate_on"`
	Desc                string             `json:"description"`
	ID                  string             `json:"id"`
	Images              []string           `json:"images"`
	Live                bool               `json:"livemode"`
	Meta                map[string]string  `json:"metadata"`
	Name                string             `json:"name"`
	PackageDimensions   *PackageDimensions `json:"package_dimensions"`
	Shippable           bool               `json:"shippable"`
	Skus                *SKUList           `json:"skus"`
	StatementDescriptor string             `json:"statement_descriptor"`
	URL                 string             `json:"url"`
	Updated             int64              `json:"updated"`
	Type                ProductType        `json:"type"`
}

// ProductList is a list of products as retrieved from a list endpoint.
type ProductList struct {
	ListMeta
	Values []*Product `json:"data"`
}

// ProductListParams is the set of parameters that can be used when
// listing products. For more details, see:
// https://stripe.com/docs/api#list_products.
type ProductListParams struct {
	ListParams `form:"*"`
	Active     *bool    `form:"active"`
	IDs        []string `form:"ids"`
	Shippable  *bool    `form:"shippable"`
	URL        string   `form:"url"`
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
