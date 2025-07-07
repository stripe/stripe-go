//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists the schemas available to the authorized merchant in Stripe's data products
type SigmaSchemaListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand  []*string `form:"expand"`
	Product *string   `form:"product"`
}

// AddExpand appends a new field to expand.
func (p *SigmaSchemaListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type SigmaSchemaTableColumn struct {
	Comment    string `json:"comment"`
	ForeignKey string `json:"foreign_key"`
	Name       string `json:"name"`
	PrimaryKey bool   `json:"primary_key"`
	Type       string `json:"type"`
}
type SigmaSchemaTable struct {
	Columns []*SigmaSchemaTableColumn `json:"columns"`
	Comment string                    `json:"comment"`
	Name    string                    `json:"name"`
	Section string                    `json:"section"`
}

// Contains information about the tables in a reporting Schema.
type SigmaSchema struct {
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string              `json:"object"`
	Tables []*SigmaSchemaTable `json:"tables"`
}

// SigmaSchemaList is a list of Schemas as retrieved from a list endpoint.
type SigmaSchemaList struct {
	APIResource
	ListMeta
	Data []*SigmaSchema `json:"data"`
}
