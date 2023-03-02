//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Whether this card bundle can be used to create cards.
type IssuingCardBundleStatus string

// List of values that IssuingCardBundleStatus can take
const (
	IssuingCardBundleStatusActive   IssuingCardBundleStatus = "active"
	IssuingCardBundleStatusInactive IssuingCardBundleStatus = "inactive"
	IssuingCardBundleStatusReview   IssuingCardBundleStatus = "review"
)

// Whether this card bundle is a standard Stripe offering or custom-made for you.
type IssuingCardBundleType string

// List of values that IssuingCardBundleType can take
const (
	IssuingCardBundleTypeCustom   IssuingCardBundleType = "custom"
	IssuingCardBundleTypeStandard IssuingCardBundleType = "standard"
)

// Returns a list of card bundle objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingCardBundleListParams struct {
	ListParams `form:"*"`
	// Only return card bundles with the given status.
	Status *string `form:"status"`
	// Only return card bundles with the given type.
	Type *string `form:"type"`
}

// Retrieves a card bundle object.
type IssuingCardBundleParams struct {
	Params `form:"*"`
}

// A Card Bundle represents the bundle of physical items - card stock, carrier letter, and envelope - that is shipped to a cardholder when you create a physical card.
type IssuingCardBundle struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Friendly display name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether this card bundle can be used to create cards.
	Status IssuingCardBundleStatus `json:"status"`
	// Whether this card bundle is a standard Stripe offering or custom-made for you.
	Type IssuingCardBundleType `json:"type"`
}

// IssuingCardBundleList is a list of CardBundles as retrieved from a list endpoint.
type IssuingCardBundleList struct {
	APIResource
	ListMeta
	Data []*IssuingCardBundle `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCardBundle.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCardBundle) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCardBundle IssuingCardBundle
	var v issuingCardBundle
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCardBundle(v)
	return nil
}
