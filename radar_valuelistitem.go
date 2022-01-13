//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type RadarValueListItemListParams struct {
	ListParams     `form:"*"`
	Created        *int64            `form:"created"`
	CreatedRange   *RangeQueryParams `form:"created"`
	RadarValueList *string           `form:"value_list"`
	Value          *string           `form:"value"`
}

// Creates a new ValueListItem object, which is added to the specified parent value list.
type RadarValueListItemParams struct {
	Params         `form:"*"`
	RadarValueList *string `form:"value_list"`
	Value          *string `form:"value"`
}

// Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.
//
// Related guide: [Managing List Items](https://stripe.com/docs/radar/lists#managing-list-items).
type RadarValueListItem struct {
	APIResource
	Created        int64  `json:"created"`
	CreatedBy      string `json:"created_by"`
	Deleted        bool   `json:"deleted"`
	ID             string `json:"id"`
	Livemode       bool   `json:"livemode"`
	Name           string `json:"name"`
	Object         string `json:"object"`
	RadarValueList string `json:"value_list"`
	Value          string `json:"value"`
}

// RadarValueListItemList is a list of ValueListItems as retrieved from a list endpoint.
type RadarValueListItemList struct {
	APIResource
	ListMeta
	Data []*RadarValueListItem `json:"data"`
}
