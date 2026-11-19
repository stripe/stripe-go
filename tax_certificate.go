//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Deletes a Tax Certificate without deleting its uploaded file. You can reuse the file to create another certificate, and you can still retrieve the deleted certificate.
type TaxCertificateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of a [file](https://docs.stripe.com/api/files/object) uploaded with purpose `tax_certificate`. The file must belong to the same account and mode as the tax exemption. You can't use a deleted file or a file attached to another certificate that hasn't been deleted.
	File *string `form:"file" json:"file,omitempty"`
	// The ID of an existing [customer tax exemption](https://docs.stripe.com/api/tax/exemptions/object) that hasn't been deleted.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Lists a tax exemption's certificates, ordered from newest to oldest. The list excludes deleted certificates.
type TaxCertificateListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of the [customer tax exemption](https://docs.stripe.com/api/tax/exemptions/object) whose certificates to list. The tax exemption must not be deleted.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deletes a Tax Certificate without deleting its uploaded file. You can reuse the file to create another certificate, and you can still retrieve the deleted certificate.
type TaxCertificateDeleteParams struct {
	Params `form:"*"`
}

// Retrieves a Tax Certificate. Deleted certificates remain retrievable with all their fields and deleted set to true.
type TaxCertificateRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a Tax Certificate associated with a tax exemption. Before creating the certificate, [upload a file](https://docs.stripe.com/api/files/create) with purpose tax_certificate.
type TaxCertificateCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The ID of a [file](https://docs.stripe.com/api/files/object) uploaded with purpose `tax_certificate`. The file must belong to the same account and mode as the tax exemption. You can't use a deleted file or a file attached to another certificate that hasn't been deleted.
	File *string `form:"file" json:"file"`
	// The ID of an existing [customer tax exemption](https://docs.stripe.com/api/tax/exemptions/object) that hasn't been deleted.
	TaxExemption *string `form:"tax_exemption" json:"tax_exemption"`
}

// AddExpand appends a new field to expand.
func (p *TaxCertificateCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A Tax Certificate associates an uploaded exemption or resale certificate file with a customer's tax exemption. A tax exemption can have multiple certificates. Deleting a tax exemption also deletes its certificates, but retains their uploaded files.
type TaxCertificate struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// If `true`, the certificate has been deleted. If `false`, it hasn't been deleted.
	Deleted bool `json:"deleted"`
	// The ID of the uploaded certificate [file](https://docs.stripe.com/api/files/object).
	File *File `json:"file"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the [customer tax exemption](https://docs.stripe.com/api/tax/exemptions/object) associated with this certificate.
	TaxExemption *CustomerTaxExemption `json:"tax_exemption"`
}

// TaxCertificateList is a list of Certificates as retrieved from a list endpoint.
type TaxCertificateList struct {
	APIResource
	ListMeta
	Data []*TaxCertificate `json:"data"`
}
