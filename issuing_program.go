//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all of the programs the given Issuing user has access to.
type IssuingProgramListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingProgramListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create a Program object.
type IssuingProgramParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If true, makes the specified program the default for the given account.
	IsDefault *bool `form:"is_default"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The program to use as the parent for the new program to create.
	PlatformProgram *string `form:"platform_program"`
}

// AddExpand appends a new field to expand.
func (p *IssuingProgramParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingProgramParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a Program object.
type IssuingProgramCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If true, makes the specified program the default for the given account.
	IsDefault *bool `form:"is_default"`
	// The program to use as the parent for the new program to create.
	PlatformProgram *string `form:"platform_program"`
}

// AddExpand appends a new field to expand.
func (p *IssuingProgramCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the program specified by the given id.
type IssuingProgramRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingProgramRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates a Program object.
type IssuingProgramUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If true, makes the specified program the default.
	IsDefault *bool `form:"is_default"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *IssuingProgramUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingProgramUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// An Issuing `Program` represents a card program that the user has access to.
type IssuingProgram struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether or not this is the "default" issuing program new cards are created on. Only one active `is_default` program at the same time.
	IsDefault bool `json:"is_default"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The platform's Issuing Program for which this program is associated.
	PlatformProgram string `json:"platform_program"`
}

// IssuingProgramList is a list of Programs as retrieved from a list endpoint.
type IssuingProgramList struct {
	APIResource
	ListMeta
	Data []*IssuingProgram `json:"data"`
}
