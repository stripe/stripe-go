//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Delete a location specific tax exemption for a customer.
type CustomerTaxExemptionParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
	Ca *CustomerTaxExemptionCaParams `form:"ca" json:"ca,omitempty"`
	// Two-letter ISO country code for the exemption location.
	Country *string `form:"country" json:"country,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption becomes effective. Must be no more than one year after today's UTC date (inclusive).
	EffectiveDate *string `form:"effective_date" json:"effective_date,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate *string `form:"expiration_date" json:"expiration_date,omitempty"`
	// US-specific exemption details. Required when country is US; must be absent otherwise.
	US *CustomerTaxExemptionUSParams `form:"us" json:"us,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CustomerTaxExemptionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// List all location specific tax exemptions for a customer.
type CustomerTaxExemptionListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	// Filter by two-letter ISO country code (ISO 3166-1 alpha-2).
	Country *string `form:"country" json:"country,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CustomerTaxExemptionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
type CustomerTaxExemptionCaParams struct {
	// Two-letter Canadian province code (ISO 3166-2). Required when tax_type is pst, qst, or rst.
	State *string `form:"state" json:"state,omitempty"`
	// The type of Canadian tax (gst_hst, PST, QST, RST).
	TaxType *string `form:"tax_type" json:"tax_type"`
}

// US-specific exemption details. Required when country is US; must be absent otherwise.
type CustomerTaxExemptionUSParams struct {
	// Two-letter US state code (ISO 3166-2).
	State *string `form:"state" json:"state"`
}

// Delete a location specific tax exemption for a customer.
type CustomerTaxExemptionDeleteParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
}

// Retrieve a location specific tax exemption for a customer.
type CustomerTaxExemptionRetrieveParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CustomerTaxExemptionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
type CustomerTaxExemptionCreateCaParams struct {
	// Two-letter Canadian province code (ISO 3166-2). Required when tax_type is pst, qst, or rst.
	State *string `form:"state" json:"state,omitempty"`
	// The type of Canadian tax (gst_hst, PST, QST, RST).
	TaxType *string `form:"tax_type" json:"tax_type"`
}

// US-specific exemption details. Required when country is US; must be absent otherwise.
type CustomerTaxExemptionCreateUSParams struct {
	// Two-letter US state code (ISO 3166-2).
	State *string `form:"state" json:"state"`
}

// Create a location specific tax exemption for a customer.
type CustomerTaxExemptionCreateParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
	Ca *CustomerTaxExemptionCreateCaParams `form:"ca" json:"ca,omitempty"`
	// Two-letter ISO country code for the exemption location.
	Country *string `form:"country" json:"country"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption becomes effective. Must be no more than one year after today's UTC date (inclusive).
	EffectiveDate *string `form:"effective_date" json:"effective_date"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate *string `form:"expiration_date" json:"expiration_date,omitempty"`
	// US-specific exemption details. Required when country is US; must be absent otherwise.
	US *CustomerTaxExemptionCreateUSParams `form:"us" json:"us,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CustomerTaxExemptionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type CustomerTaxExemptionCa struct {
	// Two-letter Canadian province code (ISO 3166-2). Null for country-wide GST/HST exemptions.
	State string `json:"state"`
	// The type of Canadian tax (gst_hst, PST, QST, RST).
	TaxType string `json:"tax_type"`
}
type CustomerTaxExemptionUS struct {
	// Two-letter US state code (ISO 3166-2).
	State string `json:"state"`
}

// Location specific customer tax exemptions.
type CustomerTaxExemption struct {
	APIResource
	Ca *CustomerTaxExemptionCa `json:"ca,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of the customer this tax exemption belongs to.
	Customer string `json:"customer"`
	// Present and true when the exemption has been deleted.
	Deleted bool `json:"deleted,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption becomes effective.
	EffectiveDate string `json:"effective_date"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate string `json:"expiration_date"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string                  `json:"object"`
	US     *CustomerTaxExemptionUS `json:"us,omitempty"`
}

// CustomerTaxExemptionList is a list of CustomerTaxExemptions as retrieved from a list endpoint.
type CustomerTaxExemptionList struct {
	APIResource
	ListMeta
	Data []*CustomerTaxExemption `json:"data"`
}
