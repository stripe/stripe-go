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
	ListParams   `form:"*"`
	Alias        *string           `form:"alias"`
	Contains     *string           `form:"contains"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Creates a new ValueList object, which can then be referenced in rules.
type RadarValueListParams struct {
	Params   `form:"*"`
	Alias    *string `form:"alias"`
	ItemType *string `form:"item_type"`
	Name     *string `form:"name"`
}

// Value lists allow you to group values together which can then be referenced in rules.
//
// Related guide: [Default Stripe Lists](https://stripe.com/docs/radar/lists#managing-list-items).
type RadarValueList struct {
	APIResource
	Alias     string                  `json:"alias"`
	Created   int64                   `json:"created"`
	CreatedBy string                  `json:"created_by"`
	Deleted   bool                    `json:"deleted"`
	ID        string                  `json:"id"`
	ItemType  RadarValueListItemType  `json:"item_type"`
	ListItems *RadarValueListItemList `json:"list_items"`
	Livemode  bool                    `json:"livemode"`
	Metadata  map[string]string       `json:"metadata"`
	Name      string                  `json:"name"`
	Object    string                  `json:"object"`
	Updated   int64                   `json:"updated"`
	UpdatedBy string                  `json:"updated_by"`
}

// RadarValueListList is a list of ValueLists as retrieved from a list endpoint.
type RadarValueListList struct {
	APIResource
	ListMeta
	Data []*RadarValueList `json:"data"`
}
