package stripe

// ValueListItemParams is the set of parameters that can be used when creating a value list item.
type ValueListItemParams struct {
	Params    `form:"*"`
	Value     *string `form:"value"`
	ValueList *string `form:"value_list"`
}

// ValueListItemListParams is the set of parameters that can be used when listing value list items.
type ValueListItemListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	ValueList    *string           `form:"value_list"`
	Value        *string           `form:"value"`
}

// ValueListItem is the resource representing a value list item.
type ValueListItem struct {
	Created   int64  `json:"created"`
	CreatedBy string `json:"created_by"`
	Deleted   bool   `json:"deleted"`
	ID        string `json:"id"`
	Livemode  bool   `json:"livemode"`
	Name      string `json:"name"`
	Object    string `json:"object"`
	Value     string `json:"value"`
	ValueList string `json:"value_list"`
}

// ValueListItemList is a list of value list items as retrieved from a list endpoint.
type ValueListItemList struct {
	ListMeta
	Data []*ValueListItem `json:"data"`
}
