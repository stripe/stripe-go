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

// Specifies the payee type.
type TaxFormPayeeType string

// List of values that TaxFormPayeeType can take
const (
	TaxFormPayeeTypeAccount           TaxFormPayeeType = "account"
	TaxFormPayeeTypeExternalReference TaxFormPayeeType = "external_reference"
)

// Whether the tax form is a mutable draft or a finalized form.
type TaxFormStatus string

// List of values that TaxFormStatus can take
const (
	TaxFormStatusDraft     TaxFormStatus = "draft"
	TaxFormStatusFinalized TaxFormStatus = "finalized"
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
	Account *string `form:"account" json:"account,omitempty"`
	// The external reference to the payee whose forms will be retrieved.
	ExternalReference *string `form:"external_reference" json:"external_reference,omitempty"`
	// Specifies the payee type.
	Type *string `form:"type" json:"type,omitempty"`
}

// Returns a list of tax forms which were previously created. The tax forms are returned in sorted order, with the oldest tax forms appearing first.
type TaxFormListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The payee whose volume is represented on the tax form.
	Payee *TaxFormListPayeeParams `form:"payee" json:"payee"`
	// Filter forms by draft or finalized status.
	Status *string `form:"status" json:"status,omitempty"`
	// An optional filter on the list, based on the object `type` field. Without the filter, the list includes all current and future tax form types. If your integration expects only one type of tax form in the response, make sure to provide a type value in the request.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
type TaxFormParams struct {
	Params `form:"*"`
	// The ID of the filed tax form to correct.
	ID *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Replacement payee details. Cannot be combined with numerical correction fields. Omitted fields remain unchanged.
	PayeeDetails *TaxFormPayeeDetailsParams `form:"payee_details" json:"payee_details,omitempty"`
	// Whether a second TIN notice was received. Supported for K and MISC forms only; cannot be supplied without a substantive correction.
	SecondTinNotice *bool `form:"second_tin_notice" json:"second_tin_notice,omitempty"`
	// Replacement numerical values for a 1099-K. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099K *TaxFormUS1099KParams `form:"us_1099_k" json:"us_1099_k,omitempty"`
	// Replacement numerical values for a 1099-MISC. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099MISC *TaxFormUS1099MISCParams `form:"us_1099_misc" json:"us_1099_misc,omitempty"`
	// Replacement numerical values for a 1099-NEC. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099Nec *TaxFormUS1099NecParams `form:"us_1099_nec" json:"us_1099_nec,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Replacement payee details. Cannot be combined with numerical correction fields. Omitted fields remain unchanged.
type TaxFormPayeeDetailsParams struct {
	// The payee's replacement address fields. Omitted fields remain unchanged.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The payee's merchant category code.
	MerchantCategoryCode *string `form:"merchant_category_code" json:"merchant_category_code,omitempty"`
	// The first line of the payee's name.
	NameLine1 *string `form:"name_line_1" json:"name_line_1,omitempty"`
	// The second line of the payee's name. An empty string clears this field.
	NameLine2 *string `form:"name_line_2" json:"name_line_2,omitempty"`
	// The payee's US taxpayer identification number.
	TaxIDNumber *string `form:"tax_id_number" json:"tax_id_number,omitempty"`
	// The payee's Treasury tipped occupation code, where supported by the reporting year and form type.
	TreasuryTippedOccupationCode *string `form:"treasury_tipped_occupation_code" json:"treasury_tipped_occupation_code,omitempty"`
}

// Replacement card-not-present transaction amount.
type TaxFormUS1099KCardNotPresentTransactionsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormUS1099KCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormUS1099KFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement amounts for all twelve months, from January through December.
type TaxFormUS1099KMonthlyVolumeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement payment transaction count.
type TaxFormUS1099KPaymentTransactionsCountParams struct {
	// The replacement non-negative transaction count.
	Count *int64 `form:"count" json:"count"`
}

// Replacement state income tax withheld.
type TaxFormUS1099KStateIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-K. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormUS1099KParams struct {
	// Replacement card-not-present transaction amount.
	CardNotPresentTransactions *TaxFormUS1099KCardNotPresentTransactionsParams `form:"card_not_present_transactions" json:"card_not_present_transactions,omitempty"`
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormUS1099KCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormUS1099KFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement amounts for all twelve months, from January through December.
	MonthlyVolumes []*TaxFormUS1099KMonthlyVolumeParams `form:"monthly_volumes" json:"monthly_volumes,omitempty"`
	// Replacement payment transaction count.
	PaymentTransactionsCount *TaxFormUS1099KPaymentTransactionsCountParams `form:"payment_transactions_count" json:"payment_transactions_count,omitempty"`
	// Replacement state income tax withheld.
	StateIncomeTaxWithheld *TaxFormUS1099KStateIncomeTaxWithheldParams `form:"state_income_tax_withheld" json:"state_income_tax_withheld,omitempty"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormUS1099MISCCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement crop insurance proceeds.
type TaxFormUS1099MISCCropInsuranceProceedsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement excess golden parachute payments, where supported by the reporting year.
type TaxFormUS1099MISCExcessGoldenParachutePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormUS1099MISCFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement fish purchased for resale.
type TaxFormUS1099MISCFishPurchasedForResaleParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement fishing boat proceeds.
type TaxFormUS1099MISCFishingBoatProceedsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement gross proceeds paid to an attorney.
type TaxFormUS1099MISCGrossProceedsPaidToAnAttorneyParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement medical and health care payments.
type TaxFormUS1099MISCMedicalAndHealthCarePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement nonqualified deferred compensation.
type TaxFormUS1099MISCNonqualifiedDeferredCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement other income.
type TaxFormUS1099MISCOtherIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement overtime compensation, where supported by the reporting year.
type TaxFormUS1099MISCOvertimeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement rents.
type TaxFormUS1099MISCRentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement royalties.
type TaxFormUS1099MISCRoyaltiesParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement section 409A deferrals.
type TaxFormUS1099MISCSection409aDeferralsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state income.
type TaxFormUS1099MISCStateIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state tax withheld.
type TaxFormUS1099MISCStateTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement substitute payments.
type TaxFormUS1099MISCSubstitutePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-MISC. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormUS1099MISCParams struct {
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormUS1099MISCCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement crop insurance proceeds.
	CropInsuranceProceeds *TaxFormUS1099MISCCropInsuranceProceedsParams `form:"crop_insurance_proceeds" json:"crop_insurance_proceeds,omitempty"`
	// Replacement direct-sales-for-resale indicator.
	DirectSalesForResale *bool `form:"direct_sales_for_resale" json:"direct_sales_for_resale,omitempty"`
	// Replacement excess golden parachute payments, where supported by the reporting year.
	ExcessGoldenParachutePayments *TaxFormUS1099MISCExcessGoldenParachutePaymentsParams `form:"excess_golden_parachute_payments" json:"excess_golden_parachute_payments,omitempty"`
	// Replacement FATCA filing requirement; cannot be supplied without a substantive correction.
	FatcaFilingRequired *bool `form:"fatca_filing_required" json:"fatca_filing_required,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormUS1099MISCFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement fishing boat proceeds.
	FishingBoatProceeds *TaxFormUS1099MISCFishingBoatProceedsParams `form:"fishing_boat_proceeds" json:"fishing_boat_proceeds,omitempty"`
	// Replacement fish purchased for resale.
	FishPurchasedForResale *TaxFormUS1099MISCFishPurchasedForResaleParams `form:"fish_purchased_for_resale" json:"fish_purchased_for_resale,omitempty"`
	// Replacement gross proceeds paid to an attorney.
	GrossProceedsPaidToAnAttorney *TaxFormUS1099MISCGrossProceedsPaidToAnAttorneyParams `form:"gross_proceeds_paid_to_an_attorney" json:"gross_proceeds_paid_to_an_attorney,omitempty"`
	// Replacement medical and health care payments.
	MedicalAndHealthCarePayments *TaxFormUS1099MISCMedicalAndHealthCarePaymentsParams `form:"medical_and_health_care_payments" json:"medical_and_health_care_payments,omitempty"`
	// Replacement nonqualified deferred compensation.
	NonqualifiedDeferredCompensation *TaxFormUS1099MISCNonqualifiedDeferredCompensationParams `form:"nonqualified_deferred_compensation" json:"nonqualified_deferred_compensation,omitempty"`
	// Replacement other income.
	OtherIncome *TaxFormUS1099MISCOtherIncomeParams `form:"other_income" json:"other_income,omitempty"`
	// Replacement overtime compensation, where supported by the reporting year.
	OvertimeCompensation *TaxFormUS1099MISCOvertimeCompensationParams `form:"overtime_compensation" json:"overtime_compensation,omitempty"`
	// Replacement rents.
	Rents *TaxFormUS1099MISCRentsParams `form:"rents" json:"rents,omitempty"`
	// Replacement royalties.
	Royalties *TaxFormUS1099MISCRoyaltiesParams `form:"royalties" json:"royalties,omitempty"`
	// Replacement section 409A deferrals.
	Section409aDeferrals *TaxFormUS1099MISCSection409aDeferralsParams `form:"section_409a_deferrals" json:"section_409a_deferrals,omitempty"`
	// Replacement state income.
	StateIncome *TaxFormUS1099MISCStateIncomeParams `form:"state_income" json:"state_income,omitempty"`
	// Replacement state tax withheld.
	StateTaxWithheld *TaxFormUS1099MISCStateTaxWithheldParams `form:"state_tax_withheld" json:"state_tax_withheld,omitempty"`
	// Replacement substitute payments.
	SubstitutePayments *TaxFormUS1099MISCSubstitutePaymentsParams `form:"substitute_payments" json:"substitute_payments,omitempty"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormUS1099NecCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormUS1099NecFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement nonemployee compensation.
type TaxFormUS1099NecNonemployeeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement overtime compensation, where supported by the reporting year.
type TaxFormUS1099NecOvertimeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state income.
type TaxFormUS1099NecStateIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state tax withheld.
type TaxFormUS1099NecStateTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-NEC. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormUS1099NecParams struct {
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormUS1099NecCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement direct-sales indicator.
	DirectSalesIndicator *bool `form:"direct_sales_indicator" json:"direct_sales_indicator,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormUS1099NecFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement nonemployee compensation.
	NonemployeeCompensation *TaxFormUS1099NecNonemployeeCompensationParams `form:"nonemployee_compensation" json:"nonemployee_compensation,omitempty"`
	// Replacement overtime compensation, where supported by the reporting year.
	OvertimeCompensation *TaxFormUS1099NecOvertimeCompensationParams `form:"overtime_compensation" json:"overtime_compensation,omitempty"`
	// Replacement state income.
	StateIncome *TaxFormUS1099NecStateIncomeParams `form:"state_income" json:"state_income,omitempty"`
	// Replacement state tax withheld.
	StateTaxWithheld *TaxFormUS1099NecStateTaxWithheldParams `form:"state_tax_withheld" json:"state_tax_withheld,omitempty"`
}

// Download the PDF for a tax form.
type TaxFormPDFParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormPDFParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
type TaxFormRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Replacement payee details. Cannot be combined with numerical correction fields. Omitted fields remain unchanged.
type TaxFormCreatePayeeDetailsParams struct {
	// The payee's replacement address fields. Omitted fields remain unchanged.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The payee's merchant category code.
	MerchantCategoryCode *string `form:"merchant_category_code" json:"merchant_category_code,omitempty"`
	// The first line of the payee's name.
	NameLine1 *string `form:"name_line_1" json:"name_line_1,omitempty"`
	// The second line of the payee's name. An empty string clears this field.
	NameLine2 *string `form:"name_line_2" json:"name_line_2,omitempty"`
	// The payee's US taxpayer identification number.
	TaxIDNumber *string `form:"tax_id_number" json:"tax_id_number,omitempty"`
	// The payee's Treasury tipped occupation code, where supported by the reporting year and form type.
	TreasuryTippedOccupationCode *string `form:"treasury_tipped_occupation_code" json:"treasury_tipped_occupation_code,omitempty"`
}

// Replacement card-not-present transaction amount.
type TaxFormCreateUS1099KCardNotPresentTransactionsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormCreateUS1099KCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormCreateUS1099KFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement amounts for all twelve months, from January through December.
type TaxFormCreateUS1099KMonthlyVolumeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement payment transaction count.
type TaxFormCreateUS1099KPaymentTransactionsCountParams struct {
	// The replacement non-negative transaction count.
	Count *int64 `form:"count" json:"count"`
}

// Replacement state income tax withheld.
type TaxFormCreateUS1099KStateIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-K. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormCreateUS1099KParams struct {
	// Replacement card-not-present transaction amount.
	CardNotPresentTransactions *TaxFormCreateUS1099KCardNotPresentTransactionsParams `form:"card_not_present_transactions" json:"card_not_present_transactions,omitempty"`
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormCreateUS1099KCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormCreateUS1099KFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement amounts for all twelve months, from January through December.
	MonthlyVolumes []*TaxFormCreateUS1099KMonthlyVolumeParams `form:"monthly_volumes" json:"monthly_volumes,omitempty"`
	// Replacement payment transaction count.
	PaymentTransactionsCount *TaxFormCreateUS1099KPaymentTransactionsCountParams `form:"payment_transactions_count" json:"payment_transactions_count,omitempty"`
	// Replacement state income tax withheld.
	StateIncomeTaxWithheld *TaxFormCreateUS1099KStateIncomeTaxWithheldParams `form:"state_income_tax_withheld" json:"state_income_tax_withheld,omitempty"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormCreateUS1099MISCCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement crop insurance proceeds.
type TaxFormCreateUS1099MISCCropInsuranceProceedsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement excess golden parachute payments, where supported by the reporting year.
type TaxFormCreateUS1099MISCExcessGoldenParachutePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormCreateUS1099MISCFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement fish purchased for resale.
type TaxFormCreateUS1099MISCFishPurchasedForResaleParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement fishing boat proceeds.
type TaxFormCreateUS1099MISCFishingBoatProceedsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement gross proceeds paid to an attorney.
type TaxFormCreateUS1099MISCGrossProceedsPaidToAnAttorneyParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement medical and health care payments.
type TaxFormCreateUS1099MISCMedicalAndHealthCarePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement nonqualified deferred compensation.
type TaxFormCreateUS1099MISCNonqualifiedDeferredCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement other income.
type TaxFormCreateUS1099MISCOtherIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement overtime compensation, where supported by the reporting year.
type TaxFormCreateUS1099MISCOvertimeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement rents.
type TaxFormCreateUS1099MISCRentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement royalties.
type TaxFormCreateUS1099MISCRoyaltiesParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement section 409A deferrals.
type TaxFormCreateUS1099MISCSection409aDeferralsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state income.
type TaxFormCreateUS1099MISCStateIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state tax withheld.
type TaxFormCreateUS1099MISCStateTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement substitute payments.
type TaxFormCreateUS1099MISCSubstitutePaymentsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-MISC. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormCreateUS1099MISCParams struct {
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormCreateUS1099MISCCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement crop insurance proceeds.
	CropInsuranceProceeds *TaxFormCreateUS1099MISCCropInsuranceProceedsParams `form:"crop_insurance_proceeds" json:"crop_insurance_proceeds,omitempty"`
	// Replacement direct-sales-for-resale indicator.
	DirectSalesForResale *bool `form:"direct_sales_for_resale" json:"direct_sales_for_resale,omitempty"`
	// Replacement excess golden parachute payments, where supported by the reporting year.
	ExcessGoldenParachutePayments *TaxFormCreateUS1099MISCExcessGoldenParachutePaymentsParams `form:"excess_golden_parachute_payments" json:"excess_golden_parachute_payments,omitempty"`
	// Replacement FATCA filing requirement; cannot be supplied without a substantive correction.
	FatcaFilingRequired *bool `form:"fatca_filing_required" json:"fatca_filing_required,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormCreateUS1099MISCFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement fishing boat proceeds.
	FishingBoatProceeds *TaxFormCreateUS1099MISCFishingBoatProceedsParams `form:"fishing_boat_proceeds" json:"fishing_boat_proceeds,omitempty"`
	// Replacement fish purchased for resale.
	FishPurchasedForResale *TaxFormCreateUS1099MISCFishPurchasedForResaleParams `form:"fish_purchased_for_resale" json:"fish_purchased_for_resale,omitempty"`
	// Replacement gross proceeds paid to an attorney.
	GrossProceedsPaidToAnAttorney *TaxFormCreateUS1099MISCGrossProceedsPaidToAnAttorneyParams `form:"gross_proceeds_paid_to_an_attorney" json:"gross_proceeds_paid_to_an_attorney,omitempty"`
	// Replacement medical and health care payments.
	MedicalAndHealthCarePayments *TaxFormCreateUS1099MISCMedicalAndHealthCarePaymentsParams `form:"medical_and_health_care_payments" json:"medical_and_health_care_payments,omitempty"`
	// Replacement nonqualified deferred compensation.
	NonqualifiedDeferredCompensation *TaxFormCreateUS1099MISCNonqualifiedDeferredCompensationParams `form:"nonqualified_deferred_compensation" json:"nonqualified_deferred_compensation,omitempty"`
	// Replacement other income.
	OtherIncome *TaxFormCreateUS1099MISCOtherIncomeParams `form:"other_income" json:"other_income,omitempty"`
	// Replacement overtime compensation, where supported by the reporting year.
	OvertimeCompensation *TaxFormCreateUS1099MISCOvertimeCompensationParams `form:"overtime_compensation" json:"overtime_compensation,omitempty"`
	// Replacement rents.
	Rents *TaxFormCreateUS1099MISCRentsParams `form:"rents" json:"rents,omitempty"`
	// Replacement royalties.
	Royalties *TaxFormCreateUS1099MISCRoyaltiesParams `form:"royalties" json:"royalties,omitempty"`
	// Replacement section 409A deferrals.
	Section409aDeferrals *TaxFormCreateUS1099MISCSection409aDeferralsParams `form:"section_409a_deferrals" json:"section_409a_deferrals,omitempty"`
	// Replacement state income.
	StateIncome *TaxFormCreateUS1099MISCStateIncomeParams `form:"state_income" json:"state_income,omitempty"`
	// Replacement state tax withheld.
	StateTaxWithheld *TaxFormCreateUS1099MISCStateTaxWithheldParams `form:"state_tax_withheld" json:"state_tax_withheld,omitempty"`
	// Replacement substitute payments.
	SubstitutePayments *TaxFormCreateUS1099MISCSubstitutePaymentsParams `form:"substitute_payments" json:"substitute_payments,omitempty"`
}

// Replacement cash tips, where supported by the reporting year.
type TaxFormCreateUS1099NecCashTipsParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement federal income tax withheld.
type TaxFormCreateUS1099NecFederalIncomeTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement nonemployee compensation.
type TaxFormCreateUS1099NecNonemployeeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement overtime compensation, where supported by the reporting year.
type TaxFormCreateUS1099NecOvertimeCompensationParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state income.
type TaxFormCreateUS1099NecStateIncomeParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement state tax withheld.
type TaxFormCreateUS1099NecStateTaxWithheldParams struct {
	// The replacement amount in major currency units, with at most two decimal places. Must be non-negative.
	VolumeDecimal *string `form:"volume_decimal" json:"volume_decimal"`
}

// Replacement numerical values for a 1099-NEC. Supply only the group matching the original form. Omitted fields remain unchanged.
type TaxFormCreateUS1099NecParams struct {
	// Replacement cash tips, where supported by the reporting year.
	CashTips *TaxFormCreateUS1099NecCashTipsParams `form:"cash_tips" json:"cash_tips,omitempty"`
	// Replacement direct-sales indicator.
	DirectSalesIndicator *bool `form:"direct_sales_indicator" json:"direct_sales_indicator,omitempty"`
	// Replacement federal income tax withheld.
	FederalIncomeTaxWithheld *TaxFormCreateUS1099NecFederalIncomeTaxWithheldParams `form:"federal_income_tax_withheld" json:"federal_income_tax_withheld,omitempty"`
	// Replacement nonemployee compensation.
	NonemployeeCompensation *TaxFormCreateUS1099NecNonemployeeCompensationParams `form:"nonemployee_compensation" json:"nonemployee_compensation,omitempty"`
	// Replacement overtime compensation, where supported by the reporting year.
	OvertimeCompensation *TaxFormCreateUS1099NecOvertimeCompensationParams `form:"overtime_compensation" json:"overtime_compensation,omitempty"`
	// Replacement state income.
	StateIncome *TaxFormCreateUS1099NecStateIncomeParams `form:"state_income" json:"state_income,omitempty"`
	// Replacement state tax withheld.
	StateTaxWithheld *TaxFormCreateUS1099NecStateTaxWithheldParams `form:"state_tax_withheld" json:"state_tax_withheld,omitempty"`
}

// Creates a correction for a filed tax form. Supply either payee details or numerical values for the tax form type. Created correction forms are not filed automatically.
type TaxFormCreateParams struct {
	Params `form:"*"`
	// The ID of the filed tax form to correct.
	ID *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Replacement payee details. Cannot be combined with numerical correction fields. Omitted fields remain unchanged.
	PayeeDetails *TaxFormCreatePayeeDetailsParams `form:"payee_details" json:"payee_details,omitempty"`
	// Whether a second TIN notice was received. Supported for K and MISC forms only; cannot be supplied without a substantive correction.
	SecondTinNotice *bool `form:"second_tin_notice" json:"second_tin_notice,omitempty"`
	// Replacement numerical values for a 1099-K. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099K *TaxFormCreateUS1099KParams `form:"us_1099_k" json:"us_1099_k,omitempty"`
	// Replacement numerical values for a 1099-MISC. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099MISC *TaxFormCreateUS1099MISCParams `form:"us_1099_misc" json:"us_1099_misc,omitempty"`
	// Replacement numerical values for a 1099-NEC. Supply only the group matching the original form. Omitted fields remain unchanged.
	US1099Nec *TaxFormCreateUS1099NecParams `form:"us_1099_nec" json:"us_1099_nec,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFormCreateParams) AddExpand(f string) {
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
	// Specifies the payee type.
	Type TaxFormPayeeType `json:"type"`
}
type TaxFormUS1099KCardNotPresentTransactions struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099KCashTips struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099KFederalIncomeTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}

// The gross amounts for each month, ordered from January through December.
type TaxFormUS1099KMonthlyVolume struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099KPaymentTransactionsCount struct {
	// The effective number of transactions.
	Count int64 `json:"count,omitempty"`
	// The signed adjustment included in the effective count. Only present for drafts.
	Delta int64 `json:"delta,omitempty"`
}
type TaxFormUS1099KStateIncomeTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099K struct {
	CardNotPresentTransactions *TaxFormUS1099KCardNotPresentTransactions `json:"card_not_present_transactions,omitempty"`
	CashTips                   *TaxFormUS1099KCashTips                   `json:"cash_tips,omitempty"`
	// The currency of the amounts on the form. Always `usd`.
	Currency                 Currency                                `json:"currency,omitempty"`
	FederalIncomeTaxWithheld *TaxFormUS1099KFederalIncomeTaxWithheld `json:"federal_income_tax_withheld,omitempty"`
	// The gross amount of payment transactions, as a decimal string in USD.
	GrossAmountOfTransactionsDecimal string `json:"gross_amount_of_transactions_decimal,omitempty"`
	// The gross amounts for each month, ordered from January through December.
	MonthlyVolumes           []*TaxFormUS1099KMonthlyVolume          `json:"monthly_volumes,omitempty"`
	PaymentTransactionsCount *TaxFormUS1099KPaymentTransactionsCount `json:"payment_transactions_count,omitempty"`
	// Year represented by the information reported on the tax form.
	ReportingYear          int64                                 `json:"reporting_year"`
	StateIncomeTaxWithheld *TaxFormUS1099KStateIncomeTaxWithheld `json:"state_income_tax_withheld,omitempty"`
}
type TaxFormUS1099MISCCashTips struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCCropInsuranceProceeds struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCExcessGoldenParachutePayments struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCFederalIncomeTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCFishPurchasedForResale struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCFishingBoatProceeds struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCGrossProceedsPaidToAnAttorney struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCMedicalAndHealthCarePayments struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCNonqualifiedDeferredCompensation struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCOtherIncome struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCOvertimeCompensation struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCRents struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCRoyalties struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCSection409aDeferrals struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCStateIncome struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCStateTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISCSubstitutePayments struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099MISC struct {
	CashTips              *TaxFormUS1099MISCCashTips              `json:"cash_tips,omitempty"`
	CropInsuranceProceeds *TaxFormUS1099MISCCropInsuranceProceeds `json:"crop_insurance_proceeds,omitempty"`
	// The currency of the amounts on the form. Always `usd`.
	Currency Currency `json:"currency,omitempty"`
	// Whether direct sales of at least $5,000 of consumer products were made for resale.
	DirectSalesForResale          bool                                            `json:"direct_sales_for_resale,omitempty"`
	ExcessGoldenParachutePayments *TaxFormUS1099MISCExcessGoldenParachutePayments `json:"excess_golden_parachute_payments,omitempty"`
	// Whether the FATCA filing requirement applies.
	FatcaFilingRequired              bool                                               `json:"fatca_filing_required,omitempty"`
	FederalIncomeTaxWithheld         *TaxFormUS1099MISCFederalIncomeTaxWithheld         `json:"federal_income_tax_withheld,omitempty"`
	FishingBoatProceeds              *TaxFormUS1099MISCFishingBoatProceeds              `json:"fishing_boat_proceeds,omitempty"`
	FishPurchasedForResale           *TaxFormUS1099MISCFishPurchasedForResale           `json:"fish_purchased_for_resale,omitempty"`
	GrossProceedsPaidToAnAttorney    *TaxFormUS1099MISCGrossProceedsPaidToAnAttorney    `json:"gross_proceeds_paid_to_an_attorney,omitempty"`
	MedicalAndHealthCarePayments     *TaxFormUS1099MISCMedicalAndHealthCarePayments     `json:"medical_and_health_care_payments,omitempty"`
	NonqualifiedDeferredCompensation *TaxFormUS1099MISCNonqualifiedDeferredCompensation `json:"nonqualified_deferred_compensation,omitempty"`
	OtherIncome                      *TaxFormUS1099MISCOtherIncome                      `json:"other_income,omitempty"`
	OvertimeCompensation             *TaxFormUS1099MISCOvertimeCompensation             `json:"overtime_compensation,omitempty"`
	Rents                            *TaxFormUS1099MISCRents                            `json:"rents,omitempty"`
	// Year represented by the information reported on the tax form.
	ReportingYear        int64                                  `json:"reporting_year"`
	Royalties            *TaxFormUS1099MISCRoyalties            `json:"royalties,omitempty"`
	Section409aDeferrals *TaxFormUS1099MISCSection409aDeferrals `json:"section_409a_deferrals,omitempty"`
	StateIncome          *TaxFormUS1099MISCStateIncome          `json:"state_income,omitempty"`
	StateTaxWithheld     *TaxFormUS1099MISCStateTaxWithheld     `json:"state_tax_withheld,omitempty"`
	SubstitutePayments   *TaxFormUS1099MISCSubstitutePayments   `json:"substitute_payments,omitempty"`
}
type TaxFormUS1099NecCashTips struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099NecFederalIncomeTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099NecNonemployeeCompensation struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099NecOvertimeCompensation struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099NecStateIncome struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099NecStateTaxWithheld struct {
	// The signed adjustment included in the effective amount, as a decimal string. Only present for drafts.
	DeltaDecimal string `json:"delta_decimal,omitempty"`
	// The effective amount in the form's currency, as a decimal string.
	VolumeDecimal string `json:"volume_decimal,omitempty"`
}
type TaxFormUS1099Nec struct {
	CashTips *TaxFormUS1099NecCashTips `json:"cash_tips,omitempty"`
	// The currency of the amounts on the form. Always `usd`.
	Currency Currency `json:"currency,omitempty"`
	// Whether direct sales of at least $5,000 of consumer products were made for resale.
	DirectSalesIndicator bool `json:"direct_sales_indicator,omitempty"`
	// Whether the FATCA filing requirement applies.
	FatcaFilingRequirement   bool                                      `json:"fatca_filing_requirement,omitempty"`
	FederalIncomeTaxWithheld *TaxFormUS1099NecFederalIncomeTaxWithheld `json:"federal_income_tax_withheld,omitempty"`
	NonemployeeCompensation  *TaxFormUS1099NecNonemployeeCompensation  `json:"nonemployee_compensation,omitempty"`
	OvertimeCompensation     *TaxFormUS1099NecOvertimeCompensation     `json:"overtime_compensation,omitempty"`
	// Year represented by the information reported on the tax form.
	ReportingYear    int64                             `json:"reporting_year"`
	StateIncome      *TaxFormUS1099NecStateIncome      `json:"state_income,omitempty"`
	StateTaxWithheld *TaxFormUS1099NecStateTaxWithheld `json:"state_tax_withheld,omitempty"`
}

// Tax forms are legal documents which are delivered to one or more tax authorities for information reporting purposes.
//
// Related guide: [US tax reporting for Connect platforms](https://stripe.com/docs/connect/tax-reporting)
type TaxForm struct {
	APIResource
	AuSerr *TaxFormAuSerr `json:"au_serr,omitempty"`
	CaMrdp *TaxFormCaMrdp `json:"ca_mrdp,omitempty"`
	// The form that corrects this form, if any.
	CorrectedBy *TaxForm `json:"corrected_by"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64          `json:"created"`
	EUDac7  *TaxFormEUDac7 `json:"eu_dac7,omitempty"`
	// A list of tax filing statuses. Note that a filing status will only be included if the form has been filed directly with the jurisdiction's tax authority.
	FilingStatuses []*TaxFormFilingStatus `json:"filing_statuses"`
	GBMrdp         *TaxFormGBMrdp         `json:"gb_mrdp,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool           `json:"livemode"`
	NzMrdp   *TaxFormNzMrdp `json:"nz_mrdp,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string        `json:"object"`
	Payee  *TaxFormPayee `json:"payee"`
	// Whether the tax form is a mutable draft or a finalized form.
	Status TaxFormStatus `json:"status,omitempty"`
	// The type of the tax form. An additional hash is included on the tax form with a name matching this value. It contains additional information specific to the tax form type.
	Type       TaxFormType        `json:"type"`
	US1099K    *TaxFormUS1099K    `json:"us_1099_k,omitempty"`
	US1099MISC *TaxFormUS1099MISC `json:"us_1099_misc,omitempty"`
	US1099Nec  *TaxFormUS1099Nec  `json:"us_1099_nec,omitempty"`
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
