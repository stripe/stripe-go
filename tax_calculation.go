//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of customer address provided.
type TaxCalculationCustomerDetailsAddressSource string

// List of values that TaxCalculationCustomerDetailsAddressSource can take
const (
	TaxCalculationCustomerDetailsAddressSourceBilling  TaxCalculationCustomerDetailsAddressSource = "billing"
	TaxCalculationCustomerDetailsAddressSourceShipping TaxCalculationCustomerDetailsAddressSource = "shipping"
)

// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
type TaxCalculationCustomerDetailsTaxIDType string

// List of values that TaxCalculationCustomerDetailsTaxIDType can take
const (
	TaxCalculationCustomerDetailsTaxIDTypeAETRN    TaxCalculationCustomerDetailsTaxIDType = "ae_trn"
	TaxCalculationCustomerDetailsTaxIDTypeAUABN    TaxCalculationCustomerDetailsTaxIDType = "au_abn"
	TaxCalculationCustomerDetailsTaxIDTypeAUARN    TaxCalculationCustomerDetailsTaxIDType = "au_arn"
	TaxCalculationCustomerDetailsTaxIDTypeBGUIC    TaxCalculationCustomerDetailsTaxIDType = "bg_uic"
	TaxCalculationCustomerDetailsTaxIDTypeBRCNPJ   TaxCalculationCustomerDetailsTaxIDType = "br_cnpj"
	TaxCalculationCustomerDetailsTaxIDTypeBRCPF    TaxCalculationCustomerDetailsTaxIDType = "br_cpf"
	TaxCalculationCustomerDetailsTaxIDTypeCABN     TaxCalculationCustomerDetailsTaxIDType = "ca_bn"
	TaxCalculationCustomerDetailsTaxIDTypeCAGSTHST TaxCalculationCustomerDetailsTaxIDType = "ca_gst_hst"
	TaxCalculationCustomerDetailsTaxIDTypeCAPSTBC  TaxCalculationCustomerDetailsTaxIDType = "ca_pst_bc"
	TaxCalculationCustomerDetailsTaxIDTypeCAPSTMB  TaxCalculationCustomerDetailsTaxIDType = "ca_pst_mb"
	TaxCalculationCustomerDetailsTaxIDTypeCAPSTSK  TaxCalculationCustomerDetailsTaxIDType = "ca_pst_sk"
	TaxCalculationCustomerDetailsTaxIDTypeCAQST    TaxCalculationCustomerDetailsTaxIDType = "ca_qst"
	TaxCalculationCustomerDetailsTaxIDTypeCHVAT    TaxCalculationCustomerDetailsTaxIDType = "ch_vat"
	TaxCalculationCustomerDetailsTaxIDTypeCLTIN    TaxCalculationCustomerDetailsTaxIDType = "cl_tin"
	TaxCalculationCustomerDetailsTaxIDTypeEGTIN    TaxCalculationCustomerDetailsTaxIDType = "eg_tin"
	TaxCalculationCustomerDetailsTaxIDTypeESCIF    TaxCalculationCustomerDetailsTaxIDType = "es_cif"
	TaxCalculationCustomerDetailsTaxIDTypeEUOSSVAT TaxCalculationCustomerDetailsTaxIDType = "eu_oss_vat"
	TaxCalculationCustomerDetailsTaxIDTypeEUVAT    TaxCalculationCustomerDetailsTaxIDType = "eu_vat"
	TaxCalculationCustomerDetailsTaxIDTypeGBVAT    TaxCalculationCustomerDetailsTaxIDType = "gb_vat"
	TaxCalculationCustomerDetailsTaxIDTypeGEVAT    TaxCalculationCustomerDetailsTaxIDType = "ge_vat"
	TaxCalculationCustomerDetailsTaxIDTypeHKBR     TaxCalculationCustomerDetailsTaxIDType = "hk_br"
	TaxCalculationCustomerDetailsTaxIDTypeHUTIN    TaxCalculationCustomerDetailsTaxIDType = "hu_tin"
	TaxCalculationCustomerDetailsTaxIDTypeIDNPWP   TaxCalculationCustomerDetailsTaxIDType = "id_npwp"
	TaxCalculationCustomerDetailsTaxIDTypeILVAT    TaxCalculationCustomerDetailsTaxIDType = "il_vat"
	TaxCalculationCustomerDetailsTaxIDTypeINGST    TaxCalculationCustomerDetailsTaxIDType = "in_gst"
	TaxCalculationCustomerDetailsTaxIDTypeISVAT    TaxCalculationCustomerDetailsTaxIDType = "is_vat"
	TaxCalculationCustomerDetailsTaxIDTypeJPCN     TaxCalculationCustomerDetailsTaxIDType = "jp_cn"
	TaxCalculationCustomerDetailsTaxIDTypeJPRN     TaxCalculationCustomerDetailsTaxIDType = "jp_rn"
	TaxCalculationCustomerDetailsTaxIDTypeJPTRN    TaxCalculationCustomerDetailsTaxIDType = "jp_trn"
	TaxCalculationCustomerDetailsTaxIDTypeKEPIN    TaxCalculationCustomerDetailsTaxIDType = "ke_pin"
	TaxCalculationCustomerDetailsTaxIDTypeKRBRN    TaxCalculationCustomerDetailsTaxIDType = "kr_brn"
	TaxCalculationCustomerDetailsTaxIDTypeLIUID    TaxCalculationCustomerDetailsTaxIDType = "li_uid"
	TaxCalculationCustomerDetailsTaxIDTypeMXRFC    TaxCalculationCustomerDetailsTaxIDType = "mx_rfc"
	TaxCalculationCustomerDetailsTaxIDTypeMYFRP    TaxCalculationCustomerDetailsTaxIDType = "my_frp"
	TaxCalculationCustomerDetailsTaxIDTypeMYITN    TaxCalculationCustomerDetailsTaxIDType = "my_itn"
	TaxCalculationCustomerDetailsTaxIDTypeMYSST    TaxCalculationCustomerDetailsTaxIDType = "my_sst"
	TaxCalculationCustomerDetailsTaxIDTypeNOVAT    TaxCalculationCustomerDetailsTaxIDType = "no_vat"
	TaxCalculationCustomerDetailsTaxIDTypeNZGST    TaxCalculationCustomerDetailsTaxIDType = "nz_gst"
	TaxCalculationCustomerDetailsTaxIDTypePHTIN    TaxCalculationCustomerDetailsTaxIDType = "ph_tin"
	TaxCalculationCustomerDetailsTaxIDTypeRUINN    TaxCalculationCustomerDetailsTaxIDType = "ru_inn"
	TaxCalculationCustomerDetailsTaxIDTypeRUKPP    TaxCalculationCustomerDetailsTaxIDType = "ru_kpp"
	TaxCalculationCustomerDetailsTaxIDTypeSAVAT    TaxCalculationCustomerDetailsTaxIDType = "sa_vat"
	TaxCalculationCustomerDetailsTaxIDTypeSGGST    TaxCalculationCustomerDetailsTaxIDType = "sg_gst"
	TaxCalculationCustomerDetailsTaxIDTypeSGUEN    TaxCalculationCustomerDetailsTaxIDType = "sg_uen"
	TaxCalculationCustomerDetailsTaxIDTypeSITIN    TaxCalculationCustomerDetailsTaxIDType = "si_tin"
	TaxCalculationCustomerDetailsTaxIDTypeTHVAT    TaxCalculationCustomerDetailsTaxIDType = "th_vat"
	TaxCalculationCustomerDetailsTaxIDTypeTRTIN    TaxCalculationCustomerDetailsTaxIDType = "tr_tin"
	TaxCalculationCustomerDetailsTaxIDTypeTWVAT    TaxCalculationCustomerDetailsTaxIDType = "tw_vat"
	TaxCalculationCustomerDetailsTaxIDTypeUAVAT    TaxCalculationCustomerDetailsTaxIDType = "ua_vat"
	TaxCalculationCustomerDetailsTaxIDTypeUnknown  TaxCalculationCustomerDetailsTaxIDType = "unknown"
	TaxCalculationCustomerDetailsTaxIDTypeUSEIN    TaxCalculationCustomerDetailsTaxIDType = "us_ein"
	TaxCalculationCustomerDetailsTaxIDTypeZAVAT    TaxCalculationCustomerDetailsTaxIDType = "za_vat"
)

// The taxability override used for taxation.
type TaxCalculationCustomerDetailsTaxabilityOverride string

// List of values that TaxCalculationCustomerDetailsTaxabilityOverride can take
const (
	TaxCalculationCustomerDetailsTaxabilityOverrideCustomerExempt TaxCalculationCustomerDetailsTaxabilityOverride = "customer_exempt"
	TaxCalculationCustomerDetailsTaxabilityOverrideNone           TaxCalculationCustomerDetailsTaxabilityOverride = "none"
	TaxCalculationCustomerDetailsTaxabilityOverrideReverseCharge  TaxCalculationCustomerDetailsTaxabilityOverride = "reverse_charge"
)

// Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes.
type TaxCalculationShippingCostTaxBehavior string

// List of values that TaxCalculationShippingCostTaxBehavior can take
const (
	TaxCalculationShippingCostTaxBehaviorExclusive TaxCalculationShippingCostTaxBehavior = "exclusive"
	TaxCalculationShippingCostTaxBehaviorInclusive TaxCalculationShippingCostTaxBehavior = "inclusive"
)

// Indicates the level of the jurisdiction imposing the tax.
type TaxCalculationShippingCostTaxBreakdownJurisdictionLevel string

// List of values that TaxCalculationShippingCostTaxBreakdownJurisdictionLevel can take
const (
	TaxCalculationShippingCostTaxBreakdownJurisdictionLevelCity     TaxCalculationShippingCostTaxBreakdownJurisdictionLevel = "city"
	TaxCalculationShippingCostTaxBreakdownJurisdictionLevelCountry  TaxCalculationShippingCostTaxBreakdownJurisdictionLevel = "country"
	TaxCalculationShippingCostTaxBreakdownJurisdictionLevelCounty   TaxCalculationShippingCostTaxBreakdownJurisdictionLevel = "county"
	TaxCalculationShippingCostTaxBreakdownJurisdictionLevelDistrict TaxCalculationShippingCostTaxBreakdownJurisdictionLevel = "district"
	TaxCalculationShippingCostTaxBreakdownJurisdictionLevelState    TaxCalculationShippingCostTaxBreakdownJurisdictionLevel = "state"
)

// Indicates whether the jurisdiction was determined by the origin (merchant's address) or destination (customer's address).
type TaxCalculationShippingCostTaxBreakdownSourcing string

// List of values that TaxCalculationShippingCostTaxBreakdownSourcing can take
const (
	TaxCalculationShippingCostTaxBreakdownSourcingDestination TaxCalculationShippingCostTaxBreakdownSourcing = "destination"
	TaxCalculationShippingCostTaxBreakdownSourcingOrigin      TaxCalculationShippingCostTaxBreakdownSourcing = "origin"
)

// The tax type, such as `vat` or `sales_tax`.
type TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType string

// List of values that TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType can take
const (
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeGST      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "gst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeHST      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "hst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeIGST     TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "igst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeJCT      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "jct"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeLeaseTax TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "lease_tax"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypePST      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "pst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeQST      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "qst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeRST      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "rst"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeSalesTax TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "sales_tax"
	TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxTypeVAT      TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType = "vat"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type TaxCalculationShippingCostTaxBreakdownTaxabilityReason string

// List of values that TaxCalculationShippingCostTaxBreakdownTaxabilityReason can take
const (
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonCustomerExempt       TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "customer_exempt"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonNotCollecting        TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "not_collecting"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonNotSubjectToTax      TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "not_subject_to_tax"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonNotSupported         TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "not_supported"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonPortionProductExempt TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "portion_product_exempt"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonPortionReducedRated  TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "portion_reduced_rated"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonPortionStandardRated TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "portion_standard_rated"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonProductExempt        TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "product_exempt"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonProductExemptHoliday TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "product_exempt_holiday"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonProportionallyRated  TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "proportionally_rated"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonReducedRated         TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "reduced_rated"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonReverseCharge        TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "reverse_charge"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonStandardRated        TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "standard_rated"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonTaxableBasisReduced  TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "taxable_basis_reduced"
	TaxCalculationShippingCostTaxBreakdownTaxabilityReasonZeroRated            TaxCalculationShippingCostTaxBreakdownTaxabilityReason = "zero_rated"
)

// The tax type, such as `vat` or `sales_tax`.
type TaxCalculationTaxBreakdownTaxRateDetailsTaxType string

// List of values that TaxCalculationTaxBreakdownTaxRateDetailsTaxType can take
const (
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeGST      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "gst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeHST      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "hst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeIGST     TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "igst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeJCT      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "jct"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeLeaseTax TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "lease_tax"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypePST      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "pst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeQST      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "qst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeRST      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "rst"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeSalesTax TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "sales_tax"
	TaxCalculationTaxBreakdownTaxRateDetailsTaxTypeVAT      TaxCalculationTaxBreakdownTaxRateDetailsTaxType = "vat"
)

// The customer's tax IDs.
type TaxCalculationCustomerDetailsTaxIDParams struct {
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `eg_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ph_tin`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tr_tin`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`
	Type *string `form:"type"`
	// Value of the tax ID.
	Value *string `form:"value"`
}

// Details about the customer, including address and tax IDs.
type TaxCalculationCustomerDetailsParams struct {
	// The customer's postal address (for example, home or business location).
	Address *AddressParams `form:"address"`
	// The type of customer address provided.
	AddressSource *string `form:"address_source"`
	// The customer's IP address (IPv4 or IPv6).
	IPAddress *string `form:"ip_address"`
	// Overrides the tax calculation result to allow you to not collect tax from your customer. Use this if you've manually checked your customer's tax exemptions. Prefer providing the customer's `tax_ids` where possible, which automatically determines whether `reverse_charge` applies.
	TaxabilityOverride *string `form:"taxability_override"`
	// The customer's tax IDs.
	TaxIDs []*TaxCalculationCustomerDetailsTaxIDParams `form:"tax_ids"`
}

// A list of items the customer is purchasing.
type TaxCalculationLineItemParams struct {
	// A positive integer in cents representing the line item's total price. If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes are calculated on top of this amount.
	Amount *int64 `form:"amount"`
	// If provided, the product's `tax_code` will be used as the line item's `tax_code`.
	Product *string `form:"product"`
	// The number of units of the item being purchased. Used to calculate the per-unit price from the total `amount` for the line. For example, if `amount=100` and `quantity=4`, the calculated unit price is 25.
	Quantity *int64 `form:"quantity"`
	// A custom identifier for this line item, which must be unique across the line items in the calculation. The reference helps identify each line item in exported [tax reports](https://stripe.com/docs/tax/reports).
	Reference *string `form:"reference"`
	// Specifies whether the `amount` includes taxes. Defaults to `exclusive`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID to use for this line item. If not provided, we will use the tax code from the provided `product` param. If neither `tax_code` nor `product` is provided, we will use the default tax code from your Tax Settings.
	TaxCode *string `form:"tax_code"`
}

// Shipping cost details to be used for the calculation.
type TaxCalculationShippingCostParams struct {
	// A positive integer in cents representing the shipping charge. If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes are calculated on top of this amount.
	Amount *int64 `form:"amount"`
	// If provided, the [shipping rate](https://stripe.com/docs/api/shipping_rates/object)'s `amount`, `tax_code` and `tax_behavior` are used. If you provide a shipping rate, then you cannot pass the `amount`, `tax_code`, or `tax_behavior` parameters.
	ShippingRate *string `form:"shipping_rate"`
	// Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes. Defaults to `exclusive`.
	TaxBehavior *string `form:"tax_behavior"`
	// The [tax code](https://stripe.com/docs/tax/tax-categories) used to calculate tax on shipping. If not provided, the default shipping tax code from your [Tax Settings](https://stripe.com/settings/tax) is used.
	TaxCode *string `form:"tax_code"`
}

// Calculates tax based on input and returns a Tax Calculation object.
type TaxCalculationParams struct {
	Params `form:"*"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of an existing customer to use for this calculation. If provided, the customer's address and tax IDs are copied to `customer_details`.
	Customer *string `form:"customer"`
	// Details about the customer, including address and tax IDs.
	CustomerDetails *TaxCalculationCustomerDetailsParams `form:"customer_details"`
	// A list of items the customer is purchasing.
	LineItems []*TaxCalculationLineItemParams `form:"line_items"`
	// Shipping cost details to be used for the calculation.
	ShippingCost *TaxCalculationShippingCostParams `form:"shipping_cost"`
	// Timestamp of date at which the tax rules and rates in effect applies for the calculation. Measured in seconds since the Unix epoch. Can be up to 48 hours in the past, and up to 48 hours in the future.
	TaxDate *int64 `form:"tax_date"`
}

// Retrieves the line items of a persisted tax calculation as a collection.
type TaxCalculationListLineItemsParams struct {
	ListParams  `form:"*"`
	Calculation *string `form:"-"` // Included in URL
}

// The customer's tax IDs (for example, EU VAT numbers).
type TaxCalculationCustomerDetailsTaxID struct {
	// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
	Type TaxCalculationCustomerDetailsTaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}
type TaxCalculationCustomerDetails struct {
	// The customer's postal address (for example, home or business location).
	Address *Address `json:"address"`
	// The type of customer address provided.
	AddressSource TaxCalculationCustomerDetailsAddressSource `json:"address_source"`
	// The customer's IP address (IPv4 or IPv6).
	IPAddress string `json:"ip_address"`
	// The taxability override used for taxation.
	TaxabilityOverride TaxCalculationCustomerDetailsTaxabilityOverride `json:"taxability_override"`
	// The customer's tax IDs (for example, EU VAT numbers).
	TaxIDs []*TaxCalculationCustomerDetailsTaxID `json:"tax_ids"`
}
type TaxCalculationShippingCostTaxBreakdownJurisdiction struct {
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// A human-readable name for the jurisdiction imposing the tax.
	DisplayName string `json:"display_name"`
	// Indicates the level of the jurisdiction imposing the tax.
	Level TaxCalculationShippingCostTaxBreakdownJurisdictionLevel `json:"level"`
	// [ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix. For example, "NY" for New York, United States.
	State string `json:"state"`
}

// Details regarding the rate for this tax. This field will be `null` when the tax is not imposed, for example if the product is exempt from tax.
type TaxCalculationShippingCostTaxBreakdownTaxRateDetails struct {
	// A localized display name for tax type, intended to be human-readable. For example, "Local Sales and Use Tax", "Value-added tax (VAT)", or "Umsatzsteuer (USt.)".
	DisplayName string `json:"display_name"`
	// The tax rate percentage as a string. For example, 8.5% is represented as "8.5".
	PercentageDecimal string `json:"percentage_decimal"`
	// The tax type, such as `vat` or `sales_tax`.
	TaxType TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType `json:"tax_type"`
}

// Detailed account of taxes relevant to shipping cost.
type TaxCalculationShippingCostTaxBreakdown struct {
	// The amount of tax, in integer cents.
	Amount       int64                                               `json:"amount"`
	Jurisdiction *TaxCalculationShippingCostTaxBreakdownJurisdiction `json:"jurisdiction"`
	// Indicates whether the jurisdiction was determined by the origin (merchant's address) or destination (customer's address).
	Sourcing TaxCalculationShippingCostTaxBreakdownSourcing `json:"sourcing"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason TaxCalculationShippingCostTaxBreakdownTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in integer cents.
	TaxableAmount int64 `json:"taxable_amount"`
	// Details regarding the rate for this tax. This field will be `null` when the tax is not imposed, for example if the product is exempt from tax.
	TaxRateDetails *TaxCalculationShippingCostTaxBreakdownTaxRateDetails `json:"tax_rate_details"`
}

// The shipping cost details for the calculation.
type TaxCalculationShippingCost struct {
	// The shipping amount in integer cents. If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes were calculated on top of this amount.
	Amount int64 `json:"amount"`
	// The amount of tax calculated for shipping, in integer cents.
	AmountTax int64 `json:"amount_tax"`
	// The ID of an existing [ShippingRate](https://stripe.com/docs/api/shipping_rates/object)
	ShippingRate string `json:"shipping_rate"`
	// Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes.
	TaxBehavior TaxCalculationShippingCostTaxBehavior `json:"tax_behavior"`
	// Detailed account of taxes relevant to shipping cost.
	TaxBreakdown []*TaxCalculationShippingCostTaxBreakdown `json:"tax_breakdown"`
	// The [tax code](https://stripe.com/docs/tax/tax-categories) ID used for shipping.
	TaxCode string `json:"tax_code"`
}
type TaxCalculationTaxBreakdownTaxRateDetails struct {
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// The tax rate percentage as a string. For example, 8.5% is represented as `"8.5"`.
	PercentageDecimal string `json:"percentage_decimal"`
	// State, county, province, or region.
	State string `json:"state"`
	// The tax type, such as `vat` or `sales_tax`.
	TaxType TaxCalculationTaxBreakdownTaxRateDetailsTaxType `json:"tax_type"`
}

// Breakdown of individual tax amounts that add up to the total.
type TaxCalculationTaxBreakdown struct {
	// The amount of tax, in integer cents.
	Amount int64 `json:"amount"`
	// Specifies whether the tax amount is included in the line item amount.
	Inclusive bool `json:"inclusive"`
	// The amount on which tax is calculated, in integer cents.
	TaxableAmount  int64                                     `json:"taxable_amount"`
	TaxRateDetails *TaxCalculationTaxBreakdownTaxRateDetails `json:"tax_rate_details"`
}

// A Tax Calculation allows you to calculate the tax to collect from your customer.
//
// Related guide: [Calculate tax in your custom payment flow](https://stripe.com/docs/tax/custom)
type TaxCalculation struct {
	APIResource
	// Total after taxes.
	AmountTotal int64 `json:"amount_total"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of an existing [Customer](https://stripe.com/docs/api/customers/object) used for the resource.
	Customer        string                         `json:"customer"`
	CustomerDetails *TaxCalculationCustomerDetails `json:"customer_details"`
	// Timestamp of date at which the tax calculation will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the calculation.
	ID string `json:"id"`
	// The list of items the customer is purchasing.
	LineItems *TaxCalculationLineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The shipping cost details for the calculation.
	ShippingCost *TaxCalculationShippingCost `json:"shipping_cost"`
	// The amount of tax to be collected on top of the line item prices.
	TaxAmountExclusive int64 `json:"tax_amount_exclusive"`
	// The amount of tax already included in the line item prices.
	TaxAmountInclusive int64 `json:"tax_amount_inclusive"`
	// Breakdown of individual tax amounts that add up to the total.
	TaxBreakdown []*TaxCalculationTaxBreakdown `json:"tax_breakdown"`
	// Timestamp of date at which the tax rules and rates in effect applies for the calculation.
	TaxDate int64 `json:"tax_date"`
}
