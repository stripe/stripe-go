//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates the level of the jurisdiction where the form was filed.
type TaxFormFilingStatusJurisdictionLevel string

// List of values that TaxFormFilingStatusJurisdictionLevel can take
const (
	TaxFormFilingStatusJurisdictionLevelCountry TaxFormFilingStatusJurisdictionLevel = "country"
	TaxFormFilingStatusJurisdictionLevelState   TaxFormFilingStatusJurisdictionLevel = "state"
)

// The current status of the filed form.
type TaxFormFilingStatusValue string

// List of values that TaxFormFilingStatusValue can take
const (
	TaxFormFilingStatusValueAccepted TaxFormFilingStatusValue = "accepted"
	TaxFormFilingStatusValueFiled    TaxFormFilingStatusValue = "filed"
	TaxFormFilingStatusValueRejected TaxFormFilingStatusValue = "rejected"
)

// Either `account` or `external_reference`.
type TaxFormPayeeType string

// List of values that TaxFormPayeeType can take
const (
	TaxFormPayeeTypeAccount           TaxFormPayeeType = "account"
	TaxFormPayeeTypeExternalReference TaxFormPayeeType = "external_reference"
)

// The type of the tax form. An additional hash is included on the tax form with a name matching this value. It contains additional information specific to the tax form type.
type TaxFormType string

// List of values that TaxFormType can take
const (
	TaxFormTypeAuSerr     TaxFormType = "au_serr"
	TaxFormTypeCaMrdp     TaxFormType = "ca_mrdp"
	TaxFormTypeEUDac7     TaxFormType = "eu_dac7"
	TaxFormTypeGBMrdp     TaxFormType = "gb_mrdp"
	TaxFormTypeNzMrdp     TaxFormType = "nz_mrdp"
	TaxFormTypeUS1099K    TaxFormType = "us_1099_k"
	TaxFormTypeUS1099MISC TaxFormType = "us_1099_misc"
	TaxFormTypeUS1099Nec  TaxFormType = "us_1099_nec"
)

// The payee whose volume is represented on the tax form.
type TaxFormListPayeeParams struct {
	// The ID of the Stripe account whose forms will be retrieved.
	Account *string `form:"account"`
	// The external reference to the payee whose forms will be retrieved.
	ExternalReference *string `form:"external_reference"`
	// Specifies the payee type. Either `account` or `external_reference`.
	Type *string `form:"type"`
}

// Returns a list of tax forms which were previously created. The tax forms are returned in sorted order, with the oldest tax forms appearing first.
type TaxFormListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The payee whose volume is represented on the tax form.
	Payee *TaxFormListPayeeParams `form:"payee"`
	// An optional filter on the list, based on the object `type` field. Without the filter, the list includes all current and future tax form types. If your integration expects only one type of tax form in the response, make sure to provide a type value in the request.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
type TaxFormParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Download the PDF for a tax form.
type TaxFormPDFParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormPDFParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
type TaxFormRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type TaxFormAuSerr struct {
	// End date of the period represented by the information reported on the tax form.
	ReportingPeriodEndDate string `json:"reporting_period_end_date"`
	// Start date of the period represented by the information reported on the tax form.
	ReportingPeriodStartDate string `json:"reporting_period_start_date"`
}
type TaxFormCaMrdp struct {
	// End date of the period represented by the information reported on the tax form.
	ReportingPeriodEndDate string `json:"reporting_period_end_date"`
	// Start date of the period represented by the information reported on the tax form.
	ReportingPeriodStartDate string `json:"reporting_period_start_date"`
}
type TaxFormEUDac7 struct {
	// End date of the period represented by the information reported on the tax form.
	ReportingPeriodEndDate string `json:"reporting_period_end_date"`
	// Start date of the period represented by the information reported on the tax form.
	ReportingPeriodStartDate string `json:"reporting_period_start_date"`
}
type TaxFormFilingStatusJurisdiction struct {
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Indicates the level of the jurisdiction where the form was filed.
	Level TaxFormFilingStatusJurisdictionLevel `json:"level"`
	// [ISO 3166-2 U.S. state code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix, if any. For example, "NY" for New York, United States. Null for non-U.S. forms.
	State string `json:"state"`
}

// A list of tax filing statuses. Note that a filing status will only be included if the form has been filed directly with the jurisdiction's tax authority.
type TaxFormFilingStatus struct {
	// Time when the filing status was updated.
	EffectiveAt  int64                            `json:"effective_at"`
	Jurisdiction *TaxFormFilingStatusJurisdiction `json:"jurisdiction"`
	// The current status of the filed form.
	Value TaxFormFilingStatusValue `json:"value"`
}
type TaxFormGBMrdp struct {
	// End date of the period represented by the information reported on the tax form.
	ReportingPeriodEndDate string `json:"reporting_period_end_date"`
	// Start date of the period represented by the information reported on the tax form.
	ReportingPeriodStartDate string `json:"reporting_period_start_date"`
}
type TaxFormNzMrdp struct {
	// End date of the period represented by the information reported on the tax form.
	ReportingPeriodEndDate string `json:"reporting_period_end_date"`
	// Start date of the period represented by the information reported on the tax form.
	ReportingPeriodStartDate string `json:"reporting_period_start_date"`
}
type TaxFormPayee struct {
	// The ID of the payee's Stripe account.
	Account *Account `json:"account"`
	// The external reference to this payee.
	ExternalReference string `json:"external_reference"`
	// Either `account` or `external_reference`.
	Type TaxFormPayeeType `json:"type"`
}
type TaxFormUS1099K struct {
	// Year represented by the information reported on the tax form.
	ReportingYear int64 `json:"reporting_year"`
}
type TaxFormUS1099MISC struct {
	// Year represented by the information reported on the tax form.
	ReportingYear int64 `json:"reporting_year"`
}
type TaxFormUS1099Nec struct {
	// Year represented by the information reported on the tax form.
	ReportingYear int64 `json:"reporting_year"`
}

// Tax forms are legal documents which are delivered to one or more tax authorities for information reporting purposes.
//
// Related guide: [US tax reporting for Connect platforms](https://stripe.com/docs/connect/tax-reporting)
type TaxForm struct {
	APIResource
	AuSerr *TaxFormAuSerr `json:"au_serr"`
	CaMrdp *TaxFormCaMrdp `json:"ca_mrdp"`
	// The form that corrects this form, if any.
	CorrectedBy *TaxForm `json:"corrected_by"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64          `json:"created"`
	EUDac7  *TaxFormEUDac7 `json:"eu_dac7"`
	// A list of tax filing statuses. Note that a filing status will only be included if the form has been filed directly with the jurisdiction's tax authority.
	FilingStatuses []*TaxFormFilingStatus `json:"filing_statuses"`
	GBMrdp         *TaxFormGBMrdp         `json:"gb_mrdp"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool           `json:"livemode"`
	NzMrdp   *TaxFormNzMrdp `json:"nz_mrdp"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string        `json:"object"`
	Payee  *TaxFormPayee `json:"payee"`
	// The type of the tax form. An additional hash is included on the tax form with a name matching this value. It contains additional information specific to the tax form type.
	Type       TaxFormType        `json:"type"`
	US1099K    *TaxFormUS1099K    `json:"us_1099_k"`
	US1099MISC *TaxFormUS1099MISC `json:"us_1099_misc"`
	US1099Nec  *TaxFormUS1099Nec  `json:"us_1099_nec"`
}

// TaxFormList is a list of Forms as retrieved from a list endpoint.
type TaxFormList struct {
	APIResource
	ListMeta
	Data []*TaxForm `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxForm.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TaxForm) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type taxForm TaxForm
	var v taxForm
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TaxForm(v)
	return nil
}
