//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Delete an apple pay domain.
type ApplePayDomainParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name" json:"domain_name,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// List apple pay domains.
type ApplePayDomainListParams struct {
	ListParams `form:"*"`
	DomainName *string `form:"domain_name" json:"domain_name,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Delete an apple pay domain.
type ApplePayDomainDeleteParams struct {
	Params `form:"*"`
}

// Retrieve an apple pay domain.
type ApplePayDomainRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create an apple pay domain.
type ApplePayDomainCreateParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name" json:"domain_name"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ApplePayDomain struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created    int64  `json:"created"`
	Deleted    bool   `json:"deleted,omitempty"`
	DomainName string `json:"domain_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// ApplePayDomainList is a list of ApplePayDomains as retrieved from a list endpoint.
type ApplePayDomainList struct {
	APIResource
	ListMeta
	Data []*ApplePayDomain `json:"data"`
}
