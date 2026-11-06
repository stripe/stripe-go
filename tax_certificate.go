//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Delete a tax certificate associated with a tax exemption.
type TaxCertificateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of a file uploaded with purpose `tax_certificate`.
	File *string `form:"file" json:"file,omitempty"`
	// The ID of an existing tax exemption.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// List certificates for an existing tax exemption.
type TaxCertificateListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of an existing tax exemption.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Delete a tax certificate associated with a tax exemption.
type TaxCertificateDeleteParams struct {
	Params `form:"*"`
}

// Retrieve a tax certificate associated with a tax exemption.
type TaxCertificateRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create a tax certificate associated with a tax exemption.
type TaxCertificateCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of a file uploaded with purpose `tax_certificate`.
	File *string `form:"file" json:"file"`
	// The ID of an existing tax exemption.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Tax exemption and resale certificates associated with tax exemptions.
type TaxCertificate struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Whether the certificate has been deleted.
	Deleted bool `json:"deleted"`
	// The uploaded certificate file.
	File *File `json:"file"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The tax exemption associated with this certificate.
	TaxExemption *CustomerTaxExemption `json:"tax_exemption"`
}

// TaxCertificateList is a list of Certificates as retrieved from a list endpoint.
type TaxCertificateList struct {
	APIResource
	ListMeta
	Data []*TaxCertificate `json:"data"`
}
