//
//
// File generated from our OpenAPI spec
//
//

package stripe

// RadarValueListItemParams is the set of parameters that can be used when creating a value list item.
type RadarValueListItemParams struct {
	Params         `form:"*"`
	RadarValueList *string `form:"value_list"`
	Value          *string `form:"value"`
}

// RadarValueListItemListParams is the set of parameters that can be used when listing value list items.
type RadarValueListItemListParams struct {
	ListParams     `form:"*"`
	Created        *int64            `form:"created"`
	CreatedRange   *RangeQueryParams `form:"created"`
	RadarValueList *string           `form:"value_list"`
	Value          *string           `form:"value"`
}

// RadarValueListItem is the resource representing a value list item.
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

// RadarValueListItemList is a list of value list items as retrieved from a list endpoint.
type RadarValueListItemList struct {
	APIResource
	ListMeta
	Data []*RadarValueListItem `json:"data"`
}
