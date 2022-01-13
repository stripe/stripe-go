//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Inventory type. Possible values are `finite`, `bucket` (not quantified), and `infinite`.
type SKUInventoryType string

// List of values that SKUInventoryType can take
const (
	SKUInventoryTypeBucket   SKUInventoryType = "bucket"
	SKUInventoryTypeFinite   SKUInventoryType = "finite"
	SKUInventoryTypeInfinite SKUInventoryType = "infinite"
)

// An indicator of the inventory available. Possible values are `in_stock`, `limited`, and `out_of_stock`. Will be present if and only if `type` is `bucket`.
type SKUInventoryValue string

// List of values that SKUInventoryValue can take
const (
	SKUInventoryValueInStock    SKUInventoryValue = "in_stock"
	SKUInventoryValueLimited    SKUInventoryValue = "limited"
	SKUInventoryValueOutOfStock SKUInventoryValue = "out_of_stock"
)

// Retrieves the details of an existing SKU. Supply the unique SKU identifier from either a SKU creation request or from the product, and Stripe will return the corresponding SKU information.
type SKUParams struct {
	Params            `form:"*"`
	Active            *bool                    `form:"active"`
	Attributes        map[string]string        `form:"attributes"`
	Currency          *string                  `form:"currency"`
	Description       *string                  `form:"description"`
	ID                *string                  `form:"id"`
	Image             *string                  `form:"image"`
	Inventory         *InventoryParams         `form:"inventory"`
	PackageDimensions *PackageDimensionsParams `form:"package_dimensions"`
	Price             *int64                   `form:"price"`
	Product           *string                  `form:"product"`
}

// Description of the SKU's inventory.
type InventoryParams struct {
	Quantity *int64  `form:"quantity"`
	Type     *string `form:"type"`
	Value    *string `form:"value"`
}

// Returns a list of your SKUs. The SKUs are returned sorted by creation date, with the most recently created SKUs appearing first.
type SKUListParams struct {
	ListParams `form:"*"`
	Active     *bool             `form:"active"`
	Attributes map[string]string `form:"attributes"`
	IDs        []*string         `form:"ids"`
	InStock    *bool             `form:"in_stock"`
	Product    *string           `form:"product"`
}
type Inventory struct {
	Quantity int64             `json:"quantity"`
	Type     SKUInventoryType  `json:"type"`
	Value    SKUInventoryValue `json:"value"`
}

// Stores representations of [stock keeping units](http://en.wikipedia.org/wiki/Stock_keeping_unit).
// SKUs describe specific product variations, taking into account any combination of: attributes,
// currency, and cost. For example, a product may be a T-shirt, whereas a specific SKU represents
// the `size: large`, `color: red` version of that shirt.
//
// Can also be used to manage inventory.
//
// Related guide: [Tax, Shipping, and Inventory](https://stripe.com/docs/orders).
type SKU struct {
	APIResource
	Active            bool               `json:"active"`
	Attributes        map[string]string  `json:"attributes"`
	Created           int64              `json:"created"`
	Currency          Currency           `json:"currency"`
	Deleted           bool               `json:"deleted"`
	Description       string             `json:"description"`
	ID                string             `json:"id"`
	Image             string             `json:"image"`
	Inventory         *Inventory         `json:"inventory"`
	Livemode          bool               `json:"livemode"`
	Metadata          map[string]string  `json:"metadata"`
	Object            string             `json:"object"`
	PackageDimensions *PackageDimensions `json:"package_dimensions"`
	Price             int64              `json:"price"`
	Product           *Product           `json:"product"`
	Updated           int64              `json:"updated"`
}

// SKUList is a list of Skus as retrieved from a list endpoint.
type SKUList struct {
	APIResource
	ListMeta
	Data []*SKU `json:"data"`
}

// UnmarshalJSON handles deserialization of a SKU.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SKU) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type sKU SKU
	var v sKU
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SKU(v)
	return nil
}
