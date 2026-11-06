//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The kind of information changed by this correction.
type TaxFormCorrectionCorrectionFormCorrectionType string

// List of values that TaxFormCorrectionCorrectionFormCorrectionType can take
const (
	TaxFormCorrectionCorrectionFormCorrectionTypePayeeChanged  TaxFormCorrectionCorrectionFormCorrectionType = "payee_changed"
	TaxFormCorrectionCorrectionFormCorrectionTypeTotalsChanged TaxFormCorrectionCorrectionFormCorrectionType = "totals_changed"
)

// The tax forms created by the correction.
type TaxFormCorrectionCorrectionForm struct {
	// The kind of information changed by this correction.
	CorrectionType TaxFormCorrectionCorrectionFormCorrectionType `json:"correction_type,omitempty"`
	// The corrected tax form.
	Form string `json:"form,omitempty"`
}

// Corrections to filed tax forms.
type TaxFormCorrection struct {
	APIResource
	// The tax forms created by the correction.
	CorrectionForms []*TaxFormCorrectionCorrectionForm `json:"correction_forms,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The original tax form that was corrected.
	OriginalForm string `json:"original_form,omitempty"`
}
