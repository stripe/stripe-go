//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.
type ProductType string

// List of values that ProductType can take
const (
	ProductTypeGood    ProductType = "good"
	ProductTypeService ProductType = "service"
)

// The dimensions of this product for shipping purposes.
type PackageDimensionsParams struct {
	Height *float64 `form:"height"`
	Length *float64 `form:"length"`
	Weight *float64 `form:"weight"`
	Width  *float64 `form:"width"`
}

// Creates a new product object.
type ProductParams struct {
	Params              `form:"*"`
	Active              *bool                    `form:"active"`
	Attributes          []*string                `form:"attributes"`
	Caption             *string                  `form:"caption"`
	DeactivateOn        []*string                `form:"deactivate_on"`
	Description         *string                  `form:"description"`
	ID                  *string                  `form:"id"`
	Images              []*string                `form:"images"`
	Name                *string                  `form:"name"`
	PackageDimensions   *PackageDimensionsParams `form:"package_dimensions"`
	Shippable           *bool                    `form:"shippable"`
	StatementDescriptor *string                  `form:"statement_descriptor"`
	TaxCode             *string                  `form:"tax_code"`
	Type                *string                  `form:"type"`
	UnitLabel           *string                  `form:"unit_label"`
	URL                 *string                  `form:"url"`
}

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
type ProductListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	IDs          []*string         `form:"ids"`
	Shippable    *bool             `form:"shippable"`
	Type         *string           `form:"type"`
	URL          *string           `form:"url"`
}

// The dimensions of this product for shipping purposes.
type PackageDimensions struct {
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Weight float64 `json:"weight"`
	Width  float64 `json:"width"`
}

// Products describe the specific goods or services you offer to your customers.
// For example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.
// They can be used in conjunction with [Prices](https://stripe.com/docs/api#prices) to configure pricing in Checkout and Subscriptions.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription) or accept [one-time payments with Checkout](https://stripe.com/docs/payments/checkout/client#create-products) and more about [Products and Prices](https://stripe.com/docs/billing/prices-guide)
type Product struct {
	APIResource
	Active              bool               `json:"active"`
	Attributes          []string           `json:"attributes"`
	Caption             string             `json:"caption"`
	Created             int64              `json:"created"`
	DeactivateOn        []string           `json:"deactivate_on"`
	Deleted             bool               `json:"deleted"`
	Description         string             `json:"description"`
	ID                  string             `json:"id"`
	Images              []string           `json:"images"`
	Livemode            bool               `json:"livemode"`
	Metadata            map[string]string  `json:"metadata"`
	Name                string             `json:"name"`
	Object              string             `json:"object"`
	PackageDimensions   *PackageDimensions `json:"package_dimensions"`
	Shippable           bool               `json:"shippable"`
	StatementDescriptor string             `json:"statement_descriptor"`
	TaxCode             *TaxCode           `json:"tax_code"`
	Type                ProductType        `json:"type"`
	UnitLabel           string             `json:"unit_label"`
	Updated             int64              `json:"updated"`
	URL                 string             `json:"url"`
}

// ProductList is a list of Products as retrieved from a list endpoint.
type ProductList struct {
	APIResource
	ListMeta
	Data []*Product `json:"data"`
}

// UnmarshalJSON handles deserialization of a Product.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Product) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type product Product
	var v product
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Product(v)
	return nil
}
