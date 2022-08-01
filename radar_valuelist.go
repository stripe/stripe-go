//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of items in the value list. One of `card_fingerprint`, `card_bin`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, or `customer_id`.
type RadarValueListItemType string

// List of values that RadarValueListItemType can take
const (
	RadarValueListItemTypeCardBin             RadarValueListItemType = "card_bin"
	RadarValueListItemTypeCardFingerprint     RadarValueListItemType = "card_fingerprint"
	RadarValueListItemTypeCaseSensitiveString RadarValueListItemType = "case_sensitive_string"
	RadarValueListItemTypeCountry             RadarValueListItemType = "country"
	RadarValueListItemTypeCustomerID          RadarValueListItemType = "customer_id"
	RadarValueListItemTypeEmail               RadarValueListItemType = "email"
	RadarValueListItemTypeIPAddress           RadarValueListItemType = "ip_address"
	RadarValueListItemTypeString              RadarValueListItemType = "string"
)

// Returns a list of ValueList objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type RadarValueListListParams struct {
	ListParams `form:"*"`
	// The alias used to reference the value list when writing rules.
	Alias *string `form:"alias"`
	// A value contained within a value list - returns all value lists containing this value.
	Contains     *string           `form:"contains"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Creates a new ValueList object, which can then be referenced in rules.
type RadarValueListParams struct {
	Params `form:"*"`
	// The name of the value list for use in rules.
	Alias *string `form:"alias"`
	// Type of the items in the value list. One of `card_fingerprint`, `card_bin`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, or `customer_id`. Use `string` if the item type is unknown or mixed.
	ItemType *string `form:"item_type"`
	// The human-readable name of the value list.
	Name *string `form:"name"`
}

// Value lists allow you to group values together which can then be referenced in rules.
//
// Related guide: [Default Stripe Lists](https://stripe.com/docs/radar/lists#managing-list-items).
type RadarValueList struct {
	APIResource
	// The name of the value list for use in rules.
	Alias string `json:"alias"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name or email address of the user who created this value list.
	CreatedBy string `json:"created_by"`
	Deleted   bool   `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The type of items in the value list. One of `card_fingerprint`, `card_bin`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, or `customer_id`.
	ItemType RadarValueListItemType `json:"item_type"`
	// List of items contained within this value list.
	ListItems *RadarValueListItemList `json:"list_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The name of the value list.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// RadarValueListList is a list of ValueLists as retrieved from a list endpoint.
type RadarValueListList struct {
	APIResource
	ListMeta
	Data []*RadarValueList `json:"data"`
}
