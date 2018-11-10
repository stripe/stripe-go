package stripe

// ValueListItemType is the possible values for a type of value list items.
type ValueListItemType string

// List of values that ValueListItemType can take.
const (
	ValueListItemTypeCardBin             ValueListItemType = "card_bin"
	ValueListItemTypeCardFingerprint     ValueListItemType = "card_fingerprint"
	ValueListItemTypeCountry             ValueListItemType = "country"
	ValueListItemTypeEmail               ValueListItemType = "email"
	ValueListItemTypeIPAddress           ValueListItemType = "ip_address"
	ValueListItemTypeString              ValueListItemType = "string"
	ValueListItemTypeCaseSensitiveString ValueListItemType = "case_sensitive_string"
)

// ValueListParams is the set of parameters that can be used when creating a value list.
type ValueListParams struct {
	Params   `form:"*"`
	Alias    *string `form:"alias"`
	ItemType *string `form:"item_type"`
	Name     *string `form:"name"`
}

// ValueListListParams is the set of parameters that can be used when listing value lists.
type ValueListListParams struct {
	ListParams   `form:"*"`
	Alias        *string           `form:"alias"`
	Contains     *string           `form:"contains"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// ValueList is the resource representing a value list.
type ValueList struct {
	Alias     string             `json:"alias"`
	Created   int64              `json:"created"`
	CreatedBy string             `json:"created_by"`
	Deleted   bool               `json:"deleted"`
	ID        string             `json:"id"`
	ItemType  ValueListItemType  `json:"item_type"`
	ListItems *ValueListItemList `json:"list_items"`
	Livemode  bool               `json:"livemode"`
	Metadata  map[string]string  `json:"metadata"`
	Name      string             `json:"name"`
	Object    string             `json:"object"`
	Updated   int64              `json:"updated"`
	UpdatedBy string             `json:"updated_by"`
}

// ValueListList is a list of value lists as retrieved from a list endpoint.
type ValueListList struct {
	ListMeta
	Data []*ValueList `json:"data"`
}
