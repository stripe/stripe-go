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
	Params `form:"*"`
	// Whether the SKU is available for purchase. Default to `true`.
	Active *bool `form:"active"`
	// A dictionary of attributes and values for the attributes defined by the product. If, for example, a product's attributes are `["size", "gender"]`, a valid SKU has the following dictionary of attributes: `{"size": "Medium", "gender": "Unisex"}`.
	Attributes map[string]string `form:"attributes"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	// The identifier for the SKU. Must be unique. If not provided, an identifier will be randomly generated.
	ID *string `form:"id"`
	// The URL of an image for this SKU, meant to be displayable to the customer.
	Image *string `form:"image"`
	// Description of the SKU's inventory.
	Inventory *InventoryParams `form:"inventory"`
	// The dimensions of this SKU for shipping purposes.
	PackageDimensions *PackageDimensionsParams `form:"package_dimensions"`
	// The cost of the item as a nonnegative integer in the smallest currency unit (that is, 100 cents to charge $1.00, or 100 to charge ¥100, Japanese Yen being a zero-decimal currency).
	Price *int64 `form:"price"`
	// The ID of the product this SKU is associated with. Must be a product with type `good`.
	Product *string `form:"product"`
}

// Description of the SKU's inventory.
type InventoryParams struct {
	// The count of inventory available. Required if `type` is `finite`.
	Quantity *int64 `form:"quantity"`
	// Inventory type. Possible values are `finite`, `bucket` (not quantified), and `infinite`.
	Type *string `form:"type"`
	// An indicator of the inventory available. Possible values are `in_stock`, `limited`, and `out_of_stock`. Will be present if and only if `type` is `bucket`.
	Value *string `form:"value"`
}

// Returns a list of your SKUs. The SKUs are returned sorted by creation date, with the most recently created SKUs appearing first.
type SKUListParams struct {
	ListParams `form:"*"`
	// Only return SKUs that are active or inactive (e.g., pass `false` to list all inactive products).
	Active *bool `form:"active"`
	// Only return SKUs that have the specified key-value pairs in this partially constructed dictionary. Can be specified only if `product` is also supplied. For instance, if the associated product has attributes `["color", "size"]`, passing in `attributes[color]=red` returns all the SKUs for this product that have `color` set to `red`.
	Attributes map[string]string `form:"attributes"`
	// Only return SKUs with the given IDs.
	IDs []*string `form:"ids"`
	// Only return SKUs that are either in stock or out of stock (e.g., pass `false` to list all SKUs that are out of stock). If no value is provided, all SKUs are returned.
	InStock *bool `form:"in_stock"`
	// The ID of the product whose SKUs will be retrieved. Must be a product with type `good`.
	Product *string `form:"product"`
}
type Inventory struct {
	// The count of inventory available. Will be present if and only if `type` is `finite`.
	Quantity int64 `json:"quantity"`
	// Inventory type. Possible values are `finite`, `bucket` (not quantified), and `infinite`.
	Type SKUInventoryType `json:"type"`
	// An indicator of the inventory available. Possible values are `in_stock`, `limited`, and `out_of_stock`. Will be present if and only if `type` is `bucket`.
	Value SKUInventoryValue `json:"value"`
}

// Stores representations of [stock keeping units](http://en.wikipedia.org/wiki/Stock_keeping_unit).
// SKUs describe specific product variations, taking into account any combination of: attributes,
// currency, and cost. For example, a product may be a T-shirt, whereas a specific SKU represents
// the `size: large`, `color: red` version of that shirt.
//
// Can also be used to manage inventory.
type SKU struct {
	APIResource
	// Whether the SKU is available for purchase.
	Active bool `json:"active"`
	// A dictionary of attributes and values for the attributes defined by the product. If, for example, a product's attributes are `["size", "gender"]`, a valid SKU has the following dictionary of attributes: `{"size": "Medium", "gender": "Unisex"}`.
	Attributes map[string]string `json:"attributes"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency    Currency `json:"currency"`
	Deleted     bool     `json:"deleted"`
	Description string   `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The URL of an image for this SKU, meant to be displayable to the customer.
	Image     string     `json:"image"`
	Inventory *Inventory `json:"inventory"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The dimensions of this SKU for shipping purposes.
	PackageDimensions *PackageDimensions `json:"package_dimensions"`
	// The cost of the item as a positive integer in the smallest currency unit (that is, 100 cents to charge $1.00, or 100 to charge ¥100, Japanese Yen being a zero-decimal currency).
	Price int64 `json:"price"`
	// The ID of the product this SKU is associated with. The product must be currently active.
	Product *Product `json:"product"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
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

// SKUList is a list of Skus as retrieved from a list endpoint.
type SKUList struct {
	APIResource
	ListMeta
	Data []*SKU `json:"data"`
}
