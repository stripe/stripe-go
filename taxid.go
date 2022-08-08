//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`. Note that some legacy tax IDs have type `unknown`
type TaxIDType string

// List of values that TaxIDType can take
const (
	TaxIDTypeAETRN    TaxIDType = "ae_trn"
	TaxIDTypeAUABN    TaxIDType = "au_abn"
	TaxIDTypeAUARN    TaxIDType = "au_arn"
	TaxIDTypeBGUIC    TaxIDType = "bg_uic"
	TaxIDTypeBRCNPJ   TaxIDType = "br_cnpj"
	TaxIDTypeBRCPF    TaxIDType = "br_cpf"
	TaxIDTypeCABN     TaxIDType = "ca_bn"
	TaxIDTypeCAGSTHST TaxIDType = "ca_gst_hst"
	TaxIDTypeCAPSTBC  TaxIDType = "ca_pst_bc"
	TaxIDTypeCAPSTMB  TaxIDType = "ca_pst_mb"
	TaxIDTypeCAPSTSK  TaxIDType = "ca_pst_sk"
	TaxIDTypeCAQST    TaxIDType = "ca_qst"
	TaxIDTypeCHVAT    TaxIDType = "ch_vat"
	TaxIDTypeCLTIN    TaxIDType = "cl_tin"
	TaxIDTypeESCIF    TaxIDType = "es_cif"
	TaxIDTypeEUOSSVAT TaxIDType = "eu_oss_vat"
	TaxIDTypeEUVAT    TaxIDType = "eu_vat"
	TaxIDTypeGBVAT    TaxIDType = "gb_vat"
	TaxIDTypeGEVAT    TaxIDType = "ge_vat"
	TaxIDTypeHKBR     TaxIDType = "hk_br"
	TaxIDTypeHUTIN    TaxIDType = "hu_tin"
	TaxIDTypeIDNPWP   TaxIDType = "id_npwp"
	TaxIDTypeILVAT    TaxIDType = "il_vat"
	TaxIDTypeINGST    TaxIDType = "in_gst"
	TaxIDTypeISVAT    TaxIDType = "is_vat"
	TaxIDTypeJPCN     TaxIDType = "jp_cn"
	TaxIDTypeJPRN     TaxIDType = "jp_rn"
	TaxIDTypeKRBRN    TaxIDType = "kr_brn"
	TaxIDTypeLIUID    TaxIDType = "li_uid"
	TaxIDTypeMXRFC    TaxIDType = "mx_rfc"
	TaxIDTypeMYFRP    TaxIDType = "my_frp"
	TaxIDTypeMYITN    TaxIDType = "my_itn"
	TaxIDTypeMYSST    TaxIDType = "my_sst"
	TaxIDTypeNOVAT    TaxIDType = "no_vat"
	TaxIDTypeNZGST    TaxIDType = "nz_gst"
	TaxIDTypeRUINN    TaxIDType = "ru_inn"
	TaxIDTypeRUKPP    TaxIDType = "ru_kpp"
	TaxIDTypeSAVAT    TaxIDType = "sa_vat"
	TaxIDTypeSGGST    TaxIDType = "sg_gst"
	TaxIDTypeSGUEN    TaxIDType = "sg_uen"
	TaxIDTypeSITIN    TaxIDType = "si_tin"
	TaxIDTypeTHVAT    TaxIDType = "th_vat"
	TaxIDTypeTWVAT    TaxIDType = "tw_vat"
	TaxIDTypeUAVAT    TaxIDType = "ua_vat"
	TaxIDTypeUnknown  TaxIDType = "unknown"
	TaxIDTypeUSEIN    TaxIDType = "us_ein"
	TaxIDTypeZAVAT    TaxIDType = "za_vat"
)

// Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.
type TaxIDVerificationStatus string

// List of values that TaxIDVerificationStatus can take
const (
	TaxIDVerificationStatusPending     TaxIDVerificationStatus = "pending"
	TaxIDVerificationStatusUnavailable TaxIDVerificationStatus = "unavailable"
	TaxIDVerificationStatusUnverified  TaxIDVerificationStatus = "unverified"
	TaxIDVerificationStatusVerified    TaxIDVerificationStatus = "verified"
)

// Creates a new TaxID object for a customer.
type TaxIDParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`
	Type *string `form:"type"`
	// Value of the tax ID.
	Value *string `form:"value"`
}

// Returns a list of tax IDs for a customer.
type TaxIDListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
}

// Tax ID verification information.
type TaxIDVerification struct {
	// Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.
	Status TaxIDVerificationStatus `json:"status"`
	// Verified address.
	VerifiedAddress string `json:"verified_address"`
	// Verified name.
	VerifiedName string `json:"verified_name"`
}

// You can add one or multiple tax IDs to a [customer](https://stripe.com/docs/api/customers).
// A customer's tax IDs are displayed on invoices and credit notes issued for the customer.
//
// Related guide: [Customer Tax Identification Numbers](https://stripe.com/docs/billing/taxes/tax-ids).
type TaxID struct {
	APIResource
	// Two-letter ISO code representing the country of the tax ID.
	Country string `json:"country"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of the customer.
	Customer *Customer `json:"customer"`
	Deleted  bool      `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`. Note that some legacy tax IDs have type `unknown`
	Type TaxIDType `json:"type"`
	// Value of the tax ID.
	Value string `json:"value"`
	// Tax ID verification information.
	Verification *TaxIDVerification `json:"verification"`
}

// TaxIDList is a list of TaxIds as retrieved from a list endpoint.
type TaxIDList struct {
	APIResource
	ListMeta
	Data []*TaxID `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxID.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TaxID) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type taxID TaxID
	var v taxID
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TaxID(v)
	return nil
}
