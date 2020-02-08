package stripe

import "encoding/json"

// TaxIDType is the list of allowed values for the tax id's type..
type TaxIDType string

// List of values that TaxIDType can take.
const (
	TaxIDTypeAUABN   TaxIDType = "au_abn"
	TaxIDTypeCABN    TaxIDType = "ca_bn"
	TaxIDTypeCAQST   TaxIDType = "ca_qst"
	TaxIDTypeCHVAT   TaxIDType = "ch_vat"
	TaxIDTypeESCIF   TaxIDType = "es_cif"
	TaxIDTypeEUVAT   TaxIDType = "eu_vat"
	TaxIDTypeHKBR    TaxIDType = "hk_br"
	TaxIDTypeINGST   TaxIDType = "in_gst"
	TaxIDTypeJPCN    TaxIDType = "jp_cn"
	TaxIDTypeKRBRN   TaxIDType = "kr_brn"
	TaxIDTypeLIUID   TaxIDType = "li_uid"
	TaxIDTypeMXRFC   TaxIDType = "mx_rfc"
	TaxIDTypeMYITN   TaxIDType = "my_itn"
	TaxIDTypeNOVAT   TaxIDType = "no_vat"
	TaxIDTypeNZGST   TaxIDType = "nz_gst"
	TaxIDTypeRUINN   TaxIDType = "ru_inn"
	TaxIDTypeSGUEN   TaxIDType = "sg_uen"
	TaxIDTypeTHVAT   TaxIDType = "th_vat"
	TaxIDTypeTWVAT   TaxIDType = "tw_vat"
	TaxIDTypeUSEIN   TaxIDType = "us_ein"
	TaxIDTypeZAVAT   TaxIDType = "za_vat"
	TaxIDTypeUnknown TaxIDType = "unknown"
)

// TaxIDVerificationStatus is the list of allowed values for the tax id's verification status..
type TaxIDVerificationStatus string

// List of values that TaxIDDuration can take.
const (
	TaxIDVerificationStatusPending     TaxIDVerificationStatus = "pending"
	TaxIDVerificationStatusUnavailable TaxIDVerificationStatus = "unavailable"
	TaxIDVerificationStatusUnverified  TaxIDVerificationStatus = "unverified"
	TaxIDVerificationStatusVerified    TaxIDVerificationStatus = "verified"
)

// TaxIDParams is the set of parameters that can be used when creating a tax id.
// For more details see https://stripe.com/docs/api/customers/create_tax_id
type TaxIDParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"`
	Type     *string `form:"type"`
	Value    *string `form:"value"`
}

// TaxIDListParams is the set of parameters that can be used when listing tax ids.
// For more detail see https://stripe.com/docs/api/customers/tax_ids
type TaxIDListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"`
}

// TaxIDVerification represents the verification details of a customer's tax id.
type TaxIDVerification struct {
	Status          TaxIDVerificationStatus `json:"status"`
	VerifiedAddress string                  `json:"verified_address"`
	VerifiedName    string                  `json:"verified_name"`
}

// TaxID is the resource representing a customer's tax id.
// For more details see https://stripe.com/docs/api/customers/tax_id_object
type TaxID struct {
	Country      string             `json:"country"`
	Created      int64              `json:"created"`
	Customer     *Customer          `json:"customer"`
	Deleted      bool               `json:"deleted"`
	ID           string             `json:"id"`
	Livemode     bool               `json:"livemode"`
	Object       string             `json:"object"`
	Type         TaxIDType          `json:"type"`
	Value        string             `json:"value"`
	Verification *TaxIDVerification `json:"verification"`
}

// TaxIDList is a list of tax ids as retrieved from a list endpoint.
type TaxIDList struct {
	ListMeta
	Data []*TaxID `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxID.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *TaxID) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type taxid TaxID
	var v taxid
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = TaxID(v)
	return nil
}
