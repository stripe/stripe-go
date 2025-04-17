//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List validation errors method
type PrivacyRedactionJobValidationErrorListParams struct {
	ListParams `form:"*"`
	Job        *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidationErrorListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve validation error method
type PrivacyRedactionJobValidationErrorParams struct {
	Params `form:"*"`
	Job    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidationErrorParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve validation error method
type PrivacyRedactionJobValidationErrorRetrieveParams struct {
	Params `form:"*"`
	Job    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidationErrorRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Validation errors
type PrivacyRedactionJobValidationError struct {
	APIResource
	Code           string            `json:"code"`
	ErroringObject map[string]string `json:"erroring_object"`
	// Unique identifier for the object.
	ID      string `json:"id"`
	Message string `json:"message"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// PrivacyRedactionJobValidationErrorList is a list of RedactionJobValidationErrors as retrieved from a list endpoint.
type PrivacyRedactionJobValidationErrorList struct {
	APIResource
	ListMeta
	Data []*PrivacyRedactionJobValidationError `json:"data"`
}
