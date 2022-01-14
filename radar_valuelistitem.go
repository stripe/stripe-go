//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type RadarValueListItemListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Identifier for the parent value list this item belongs to.
	RadarValueList *string `form:"value_list"`
	// Return items belonging to the parent list whose value matches the specified value (using an "is like" match).
	Value *string `form:"value"`
}

// Creates a new ValueListItem object, which is added to the specified parent value list.
type RadarValueListItemParams struct {
	Params `form:"*"`
	// The identifier of the value list which the created item will be added to.
	RadarValueList *string `form:"value_list"`
	// The value of the item (whose type must match the type of the parent value list).
	Value *string `form:"value"`
}

// Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.
//
// Related guide: [Managing List Items](https://stripe.com/docs/radar/lists#managing-list-items).
type RadarValueListItem struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name or email address of the user who added this item to the value list.
	CreatedBy string `json:"created_by"`
	Deleted   bool   `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool   `json:"livemode"`
	Name     string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The identifier of the value list this item belongs to.
	RadarValueList string `json:"value_list"`
	// The value of the item.
	Value string `json:"value"`
}

// RadarValueListItemList is a list of ValueListItems as retrieved from a list endpoint.
type RadarValueListItemList struct {
	APIResource
	ListMeta
	Data []*RadarValueListItem `json:"data"`
}
