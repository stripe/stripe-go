//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The reason the customer is exempt.
type CustomerTaxExemptionReason string

// List of values that CustomerTaxExemptionReason can take
const (
	CustomerTaxExemptionReasonDirectPayPermit CustomerTaxExemptionReason = "direct_pay_permit"
	CustomerTaxExemptionReasonLegacyContract  CustomerTaxExemptionReason = "legacy_contract"
	CustomerTaxExemptionReasonOther           CustomerTaxExemptionReason = "other"
)

// The group of tax codes this exemption applies to. Exemptions created without one have `all`.
type CustomerTaxExemptionUSTaxCodeGroup string

// List of values that CustomerTaxExemptionUSTaxCodeGroup can take
const (
	CustomerTaxExemptionUSTaxCodeGroupAll  CustomerTaxExemptionUSTaxCodeGroup = "all"
	CustomerTaxExemptionUSTaxCodeGroupSaas CustomerTaxExemptionUSTaxCodeGroup = "saas"
)

// Delete a location specific tax exemption for a customer.
type CustomerTaxExemptionParams struct {
	Params `form:"*"`
	// Unique identifier for the object.
	Customer *string `form:"-"` // Included in URL
	// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
	Ca *CustomerTaxExemptionCaParams `form:"ca" json:"ca,omitempty"`
	// Two-letter ISO country code for the exemption location.
	Country *string `form:"country" json:"country,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption becomes effective. Must be no more than one year after today's UTC date (inclusive).
	EffectiveDate *string `form:"effective_date" json:"effective_date,omitempty"`
	// Spain-specific exemption details. Optional when country is ES; must be absent otherwise.
	ES *CustomerTaxExemptionESParams `form:"es" json:"es,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate *string `form:"expiration_date" json:"expiration_date,omitempty"`
	// The reason the customer is exempt. Required when `us.tax_code_group` is `saas`.
	Reason *string `form:"reason" json:"reason,omitempty"`
	// Additional details about the exemption reason.
	ReasonDetails *CustomerTaxExemptionReasonDetailsParams `form:"reason_details" json:"reason_details,omitempty"`
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
	// Unique identifier for the object.
	Customer *string `form:"-"` // Included in URL
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

// Spain-specific exemption details. Optional when country is ES; must be absent otherwise.
type CustomerTaxExemptionESParams struct {
	// Two-letter Spanish subdivision code (ISO 3166-2).
	State *string `form:"state" json:"state"`
}

// Details for the `other` exemption reason.
type CustomerTaxExemptionReasonDetailsOtherParams struct {
	// A free-text description of the exemption reason. Maximum 5000 characters.
	ReasonText *string `form:"reason_text" json:"reason_text"`
}

// Additional details about the exemption reason.
type CustomerTaxExemptionReasonDetailsParams struct {
	// Details for the `other` exemption reason.
	Other *CustomerTaxExemptionReasonDetailsOtherParams `form:"other" json:"other,omitempty"`
}

// US-specific exemption details. Required when country is US; must be absent otherwise.
type CustomerTaxExemptionUSParams struct {
	// Two-letter US state code (ISO 3166-2).
	State *string `form:"state" json:"state"`
	// The group of tax codes this exemption applies to. Defaults to `all`. `saas` is only available when `state` is `CA`.
	TaxCodeGroup *string `form:"tax_code_group" json:"tax_code_group,omitempty"`
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

// Spain-specific exemption details. Optional when country is ES; must be absent otherwise.
type CustomerTaxExemptionCreateESParams struct {
	// Two-letter Spanish subdivision code (ISO 3166-2).
	State *string `form:"state" json:"state"`
}

// Details for the `other` exemption reason.
type CustomerTaxExemptionCreateReasonDetailsOtherParams struct {
	// A free-text description of the exemption reason. Maximum 5000 characters.
	ReasonText *string `form:"reason_text" json:"reason_text"`
}

// Additional details about the exemption reason.
type CustomerTaxExemptionCreateReasonDetailsParams struct {
	// Details for the `other` exemption reason.
	Other *CustomerTaxExemptionCreateReasonDetailsOtherParams `form:"other" json:"other,omitempty"`
}

// US-specific exemption details. Required when country is US; must be absent otherwise.
type CustomerTaxExemptionCreateUSParams struct {
	// Two-letter US state code (ISO 3166-2).
	State *string `form:"state" json:"state"`
	// The group of tax codes this exemption applies to. Defaults to `all`. `saas` is only available when `state` is `CA`.
	TaxCodeGroup *string `form:"tax_code_group" json:"tax_code_group,omitempty"`
}

// Create a location specific tax exemption for a customer.
type CustomerTaxExemptionCreateParams struct {
	Params `form:"*"`
	// Unique identifier for the object.
	Customer *string `form:"-"` // Included in URL
	// Canada-specific exemption details. Required when country is CA; must be absent otherwise.
	Ca *CustomerTaxExemptionCreateCaParams `form:"ca" json:"ca,omitempty"`
	// Two-letter ISO country code for the exemption location.
	Country *string `form:"country" json:"country"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption becomes effective. Must be no more than one year after today's UTC date (inclusive).
	EffectiveDate *string `form:"effective_date" json:"effective_date"`
	// Spain-specific exemption details. Optional when country is ES; must be absent otherwise.
	ES *CustomerTaxExemptionCreateESParams `form:"es" json:"es,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate *string `form:"expiration_date" json:"expiration_date,omitempty"`
	// The reason the customer is exempt. Required when `us.tax_code_group` is `saas`.
	Reason *string `form:"reason" json:"reason,omitempty"`
	// Additional details about the exemption reason.
	ReasonDetails *CustomerTaxExemptionCreateReasonDetailsParams `form:"reason_details" json:"reason_details,omitempty"`
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
type CustomerTaxExemptionES struct {
	// Two-letter Spanish subdivision code (ISO 3166-2). Absent for country-wide Spain exemptions.
	State string `json:"state,omitempty"`
}
type CustomerTaxExemptionReasonDetailsOther struct {
	// A free-text description of the exemption reason.
	ReasonText string `json:"reason_text,omitempty"`
}

// Additional details about the exemption reason.
type CustomerTaxExemptionReasonDetails struct {
	Other *CustomerTaxExemptionReasonDetailsOther `json:"other,omitempty"`
}
type CustomerTaxExemptionUS struct {
	// Two-letter US state code (ISO 3166-2).
	State string `json:"state"`
	// The group of tax codes this exemption applies to. Exemptions created without one have `all`.
	TaxCodeGroup CustomerTaxExemptionUSTaxCodeGroup `json:"tax_code_group,omitempty"`
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
	EffectiveDate string                  `json:"effective_date"`
	ES            *CustomerTaxExemptionES `json:"es,omitempty"`
	// ISO 8601 date (YYYY-MM-DD) when the exemption expires.
	ExpirationDate string `json:"expiration_date"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The reason the customer is exempt.
	Reason CustomerTaxExemptionReason `json:"reason,omitempty"`
	// Additional details about the exemption reason.
	ReasonDetails *CustomerTaxExemptionReasonDetails `json:"reason_details,omitempty"`
	US            *CustomerTaxExemptionUS            `json:"us,omitempty"`
}

// CustomerTaxExemptionList is a list of CustomerTaxExemptions as retrieved from a list endpoint.
type CustomerTaxExemptionList struct {
	APIResource
	ListMeta
	Data []*CustomerTaxExemption `json:"data"`
}

// UnmarshalJSON handles deserialization of a CustomerTaxExemption.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CustomerTaxExemption) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type customerTaxExemption CustomerTaxExemption
	var v customerTaxExemption
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CustomerTaxExemption(v)
	return nil
}
