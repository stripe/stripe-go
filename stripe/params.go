package stripe

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	Start, End string
	Limit      uint64
	Filters    Filters
}

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	Expand      []string
	Meta        map[string]string
	AccessToken string
}
