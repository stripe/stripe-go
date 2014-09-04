package stripe

import (
	"fmt"
	"net/url"
)

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
	Exp         []string
	Meta        map[string]string
	AccessToken string
}

// Expand appends a new field to expand.
func (p *Params) Expand(f string) {
	p.Exp = append(p.Exp, f)
}

// appendTo adds the common parameters to the query string values.
func (p *Params) appendTo(body *url.Values) {
	for k, v := range p.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	for _, v := range p.Exp {
		body.Add("expand[]", v)
	}
}
