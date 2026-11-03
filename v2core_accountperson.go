//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

// Purpose of additional address.
type V2CoreAccountPersonAdditionalAddressPurpose string

// List of values that V2CoreAccountPersonAdditionalAddressPurpose can take
const (
	V2CoreAccountPersonAdditionalAddressPurposeRegistered V2CoreAccountPersonAdditionalAddressPurpose = "registered"
)

// The purpose or type of the additional name.
type V2CoreAccountPersonAdditionalNamePurpose string

// List of values that V2CoreAccountPersonAdditionalNamePurpose can take
const (
	V2CoreAccountPersonAdditionalNamePurposeAlias  V2CoreAccountPersonAdditionalNamePurpose = "alias"
	V2CoreAccountPersonAdditionalNamePurposeMaiden V2CoreAccountPersonAdditionalNamePurpose = "maiden"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsCompanyAuthorizationType string

// List of values that V2CoreAccountPersonDocumentsCompanyAuthorizationType can take
const (
	V2CoreAccountPersonDocumentsCompanyAuthorizationTypeFiles V2CoreAccountPersonDocumentsCompanyAuthorizationType = "files"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsPassportType string

// List of values that V2CoreAccountPersonDocumentsPassportType can take
const (
	V2CoreAccountPersonDocumentsPassportTypeFiles V2CoreAccountPersonDocumentsPassportType = "files"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CoreAccountPersonDocumentsPrimaryVerificationType string

// List of values that V2CoreAccountPersonDocumentsPrimaryVerificationType can take
const (
	V2CoreAccountPersonDocumentsPrimaryVerificationTypeFrontBack V2CoreAccountPersonDocumentsPrimaryVerificationType = "front_back"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CoreAccountPersonDocumentsSecondaryVerificationType string

// List of values that V2CoreAccountPersonDocumentsSecondaryVerificationType can take
const (
	V2CoreAccountPersonDocumentsSecondaryVerificationTypeFrontBack V2CoreAccountPersonDocumentsSecondaryVerificationType = "front_back"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsVisaType string

// List of values that V2CoreAccountPersonDocumentsVisaType can take
const (
	V2CoreAccountPersonDocumentsVisaTypeFiles V2CoreAccountPersonDocumentsVisaType = "files"
)

// The ID number type of an individual.
type V2CoreAccountPersonIDNumberType string

// List of values that V2CoreAccountPersonIDNumberType can take
const (
	V2CoreAccountPersonIDNumberTypeAeEid       V2CoreAccountPersonIDNumberType = "ae_eid"
	V2CoreAccountPersonIDNumberTypeAePp        V2CoreAccountPersonIDNumberType = "ae_pp"
	V2CoreAccountPersonIDNumberTypeAlPp        V2CoreAccountPersonIDNumberType = "al_pp"
	V2CoreAccountPersonIDNumberTypeAmPp        V2CoreAccountPersonIDNumberType = "am_pp"
	V2CoreAccountPersonIDNumberTypeAoNif       V2CoreAccountPersonIDNumberType = "ao_nif"
	V2CoreAccountPersonIDNumberTypeARCuil      V2CoreAccountPersonIDNumberType = "ar_cuil"
	V2CoreAccountPersonIDNumberTypeARDni       V2CoreAccountPersonIDNumberType = "ar_dni"
	V2CoreAccountPersonIDNumberTypeARPp        V2CoreAccountPersonIDNumberType = "ar_pp"
	V2CoreAccountPersonIDNumberTypeAtPp        V2CoreAccountPersonIDNumberType = "at_pp"
	V2CoreAccountPersonIDNumberTypeAtStn       V2CoreAccountPersonIDNumberType = "at_stn"
	V2CoreAccountPersonIDNumberTypeAuPp        V2CoreAccountPersonIDNumberType = "au_pp"
	V2CoreAccountPersonIDNumberTypeAzPIN       V2CoreAccountPersonIDNumberType = "az_pin"
	V2CoreAccountPersonIDNumberTypeAzPp        V2CoreAccountPersonIDNumberType = "az_pp"
	V2CoreAccountPersonIDNumberTypeAzTin       V2CoreAccountPersonIDNumberType = "az_tin"
	V2CoreAccountPersonIDNumberTypeBdBrc       V2CoreAccountPersonIDNumberType = "bd_brc"
	V2CoreAccountPersonIDNumberTypeBdEtin      V2CoreAccountPersonIDNumberType = "bd_etin"
	V2CoreAccountPersonIDNumberTypeBdNid       V2CoreAccountPersonIDNumberType = "bd_nid"
	V2CoreAccountPersonIDNumberTypeBdPp        V2CoreAccountPersonIDNumberType = "bd_pp"
	V2CoreAccountPersonIDNumberTypeBeNrn       V2CoreAccountPersonIDNumberType = "be_nrn"
	V2CoreAccountPersonIDNumberTypeBePp        V2CoreAccountPersonIDNumberType = "be_pp"
	V2CoreAccountPersonIDNumberTypeBGPp        V2CoreAccountPersonIDNumberType = "bg_pp"
	V2CoreAccountPersonIDNumberTypeBGUcn       V2CoreAccountPersonIDNumberType = "bg_ucn"
	V2CoreAccountPersonIDNumberTypeBhPp        V2CoreAccountPersonIDNumberType = "bh_pp"
	V2CoreAccountPersonIDNumberTypeBjPp        V2CoreAccountPersonIDNumberType = "bj_pp"
	V2CoreAccountPersonIDNumberTypeBnNric      V2CoreAccountPersonIDNumberType = "bn_nric"
	V2CoreAccountPersonIDNumberTypeBoPp        V2CoreAccountPersonIDNumberType = "bo_pp"
	V2CoreAccountPersonIDNumberTypeBRCPF       V2CoreAccountPersonIDNumberType = "br_cpf"
	V2CoreAccountPersonIDNumberTypeBrPp        V2CoreAccountPersonIDNumberType = "br_pp"
	V2CoreAccountPersonIDNumberTypeBrRg        V2CoreAccountPersonIDNumberType = "br_rg"
	V2CoreAccountPersonIDNumberTypeBsPp        V2CoreAccountPersonIDNumberType = "bs_pp"
	V2CoreAccountPersonIDNumberTypeBsTin       V2CoreAccountPersonIDNumberType = "bs_tin"
	V2CoreAccountPersonIDNumberTypeByPp        V2CoreAccountPersonIDNumberType = "by_pp"
	V2CoreAccountPersonIDNumberTypeCaPp        V2CoreAccountPersonIDNumberType = "ca_pp"
	V2CoreAccountPersonIDNumberTypeCaSin       V2CoreAccountPersonIDNumberType = "ca_sin"
	V2CoreAccountPersonIDNumberTypeChOasi      V2CoreAccountPersonIDNumberType = "ch_oasi"
	V2CoreAccountPersonIDNumberTypeChPp        V2CoreAccountPersonIDNumberType = "ch_pp"
	V2CoreAccountPersonIDNumberTypeCiPp        V2CoreAccountPersonIDNumberType = "ci_pp"
	V2CoreAccountPersonIDNumberTypeClPp        V2CoreAccountPersonIDNumberType = "cl_pp"
	V2CoreAccountPersonIDNumberTypeClRut       V2CoreAccountPersonIDNumberType = "cl_rut"
	V2CoreAccountPersonIDNumberTypeCmPp        V2CoreAccountPersonIDNumberType = "cm_pp"
	V2CoreAccountPersonIDNumberTypeCnPp        V2CoreAccountPersonIDNumberType = "cn_pp"
	V2CoreAccountPersonIDNumberTypeCoNuip      V2CoreAccountPersonIDNumberType = "co_nuip"
	V2CoreAccountPersonIDNumberTypeCoPp        V2CoreAccountPersonIDNumberType = "co_pp"
	V2CoreAccountPersonIDNumberTypeCrCi        V2CoreAccountPersonIDNumberType = "cr_ci"
	V2CoreAccountPersonIDNumberTypeCrCpf       V2CoreAccountPersonIDNumberType = "cr_cpf"
	V2CoreAccountPersonIDNumberTypeCrDimex     V2CoreAccountPersonIDNumberType = "cr_dimex"
	V2CoreAccountPersonIDNumberTypeCrNite      V2CoreAccountPersonIDNumberType = "cr_nite"
	V2CoreAccountPersonIDNumberTypeCrPp        V2CoreAccountPersonIDNumberType = "cr_pp"
	V2CoreAccountPersonIDNumberTypeCyPp        V2CoreAccountPersonIDNumberType = "cy_pp"
	V2CoreAccountPersonIDNumberTypeCyTic       V2CoreAccountPersonIDNumberType = "cy_tic"
	V2CoreAccountPersonIDNumberTypeCzDic       V2CoreAccountPersonIDNumberType = "cz_dic"
	V2CoreAccountPersonIDNumberTypeCzPp        V2CoreAccountPersonIDNumberType = "cz_pp"
	V2CoreAccountPersonIDNumberTypeCzRc        V2CoreAccountPersonIDNumberType = "cz_rc"
	V2CoreAccountPersonIDNumberTypeDEPp        V2CoreAccountPersonIDNumberType = "de_pp"
	V2CoreAccountPersonIDNumberTypeDEStn       V2CoreAccountPersonIDNumberType = "de_stn"
	V2CoreAccountPersonIDNumberTypeDkCpr       V2CoreAccountPersonIDNumberType = "dk_cpr"
	V2CoreAccountPersonIDNumberTypeDkPp        V2CoreAccountPersonIDNumberType = "dk_pp"
	V2CoreAccountPersonIDNumberTypeDoCie       V2CoreAccountPersonIDNumberType = "do_cie"
	V2CoreAccountPersonIDNumberTypeDoPp        V2CoreAccountPersonIDNumberType = "do_pp"
	V2CoreAccountPersonIDNumberTypeDORCN       V2CoreAccountPersonIDNumberType = "do_rcn"
	V2CoreAccountPersonIDNumberTypeDzPp        V2CoreAccountPersonIDNumberType = "dz_pp"
	V2CoreAccountPersonIDNumberTypeEcCi        V2CoreAccountPersonIDNumberType = "ec_ci"
	V2CoreAccountPersonIDNumberTypeEcPp        V2CoreAccountPersonIDNumberType = "ec_pp"
	V2CoreAccountPersonIDNumberTypeEeIk        V2CoreAccountPersonIDNumberType = "ee_ik"
	V2CoreAccountPersonIDNumberTypeEePp        V2CoreAccountPersonIDNumberType = "ee_pp"
	V2CoreAccountPersonIDNumberTypeEgPp        V2CoreAccountPersonIDNumberType = "eg_pp"
	V2CoreAccountPersonIDNumberTypeESNif       V2CoreAccountPersonIDNumberType = "es_nif"
	V2CoreAccountPersonIDNumberTypeESPp        V2CoreAccountPersonIDNumberType = "es_pp"
	V2CoreAccountPersonIDNumberTypeETTin       V2CoreAccountPersonIDNumberType = "et_tin"
	V2CoreAccountPersonIDNumberTypeFIHetu      V2CoreAccountPersonIDNumberType = "fi_hetu"
	V2CoreAccountPersonIDNumberTypeFIPp        V2CoreAccountPersonIDNumberType = "fi_pp"
	V2CoreAccountPersonIDNumberTypeFRNir       V2CoreAccountPersonIDNumberType = "fr_nir"
	V2CoreAccountPersonIDNumberTypeFRPp        V2CoreAccountPersonIDNumberType = "fr_pp"
	V2CoreAccountPersonIDNumberTypeFRSpi       V2CoreAccountPersonIDNumberType = "fr_spi"
	V2CoreAccountPersonIDNumberTypeGaNif       V2CoreAccountPersonIDNumberType = "ga_nif"
	V2CoreAccountPersonIDNumberTypeGBNino      V2CoreAccountPersonIDNumberType = "gb_nino"
	V2CoreAccountPersonIDNumberTypeGBPp        V2CoreAccountPersonIDNumberType = "gb_pp"
	V2CoreAccountPersonIDNumberTypeGePp        V2CoreAccountPersonIDNumberType = "ge_pp"
	V2CoreAccountPersonIDNumberTypeGhPp        V2CoreAccountPersonIDNumberType = "gh_pp"
	V2CoreAccountPersonIDNumberTypeGrAfm       V2CoreAccountPersonIDNumberType = "gr_afm"
	V2CoreAccountPersonIDNumberTypeGrPp        V2CoreAccountPersonIDNumberType = "gr_pp"
	V2CoreAccountPersonIDNumberTypeGtCui       V2CoreAccountPersonIDNumberType = "gt_cui"
	V2CoreAccountPersonIDNumberTypeGtNit       V2CoreAccountPersonIDNumberType = "gt_nit"
	V2CoreAccountPersonIDNumberTypeGtPp        V2CoreAccountPersonIDNumberType = "gt_pp"
	V2CoreAccountPersonIDNumberTypeHkID        V2CoreAccountPersonIDNumberType = "hk_id"
	V2CoreAccountPersonIDNumberTypeHkPp        V2CoreAccountPersonIDNumberType = "hk_pp"
	V2CoreAccountPersonIDNumberTypeHnDni       V2CoreAccountPersonIDNumberType = "hn_dni"
	V2CoreAccountPersonIDNumberTypeHnPp        V2CoreAccountPersonIDNumberType = "hn_pp"
	V2CoreAccountPersonIDNumberTypeHROIB       V2CoreAccountPersonIDNumberType = "hr_oib"
	V2CoreAccountPersonIDNumberTypeHRPp        V2CoreAccountPersonIDNumberType = "hr_pp"
	V2CoreAccountPersonIDNumberTypeHtPp        V2CoreAccountPersonIDNumberType = "ht_pp"
	V2CoreAccountPersonIDNumberTypeHUAd        V2CoreAccountPersonIDNumberType = "hu_ad"
	V2CoreAccountPersonIDNumberTypeHUPp        V2CoreAccountPersonIDNumberType = "hu_pp"
	V2CoreAccountPersonIDNumberTypeIDNik       V2CoreAccountPersonIDNumberType = "id_nik"
	V2CoreAccountPersonIDNumberTypeIDPp        V2CoreAccountPersonIDNumberType = "id_pp"
	V2CoreAccountPersonIDNumberTypeIePp        V2CoreAccountPersonIDNumberType = "ie_pp"
	V2CoreAccountPersonIDNumberTypeIePpsn      V2CoreAccountPersonIDNumberType = "ie_ppsn"
	V2CoreAccountPersonIDNumberTypeIlPp        V2CoreAccountPersonIDNumberType = "il_pp"
	V2CoreAccountPersonIDNumberTypeInPan       V2CoreAccountPersonIDNumberType = "in_pan"
	V2CoreAccountPersonIDNumberTypeInPp        V2CoreAccountPersonIDNumberType = "in_pp"
	V2CoreAccountPersonIDNumberTypeIqPp        V2CoreAccountPersonIDNumberType = "iq_pp"
	V2CoreAccountPersonIDNumberTypeIsKt        V2CoreAccountPersonIDNumberType = "is_kt"
	V2CoreAccountPersonIDNumberTypeIsPp        V2CoreAccountPersonIDNumberType = "is_pp"
	V2CoreAccountPersonIDNumberTypeITCf        V2CoreAccountPersonIDNumberType = "it_cf"
	V2CoreAccountPersonIDNumberTypeITPp        V2CoreAccountPersonIDNumberType = "it_pp"
	V2CoreAccountPersonIDNumberTypeJePp        V2CoreAccountPersonIDNumberType = "je_pp"
	V2CoreAccountPersonIDNumberTypeJmPp        V2CoreAccountPersonIDNumberType = "jm_pp"
	V2CoreAccountPersonIDNumberTypeJoPp        V2CoreAccountPersonIDNumberType = "jo_pp"
	V2CoreAccountPersonIDNumberTypeJPInc       V2CoreAccountPersonIDNumberType = "jp_inc"
	V2CoreAccountPersonIDNumberTypeJPPp        V2CoreAccountPersonIDNumberType = "jp_pp"
	V2CoreAccountPersonIDNumberTypeKEPIN       V2CoreAccountPersonIDNumberType = "ke_pin"
	V2CoreAccountPersonIDNumberTypeKePp        V2CoreAccountPersonIDNumberType = "ke_pp"
	V2CoreAccountPersonIDNumberTypeKrPp        V2CoreAccountPersonIDNumberType = "kr_pp"
	V2CoreAccountPersonIDNumberTypeKwPp        V2CoreAccountPersonIDNumberType = "kw_pp"
	V2CoreAccountPersonIDNumberTypeKzIIN       V2CoreAccountPersonIDNumberType = "kz_iin"
	V2CoreAccountPersonIDNumberTypeKzPp        V2CoreAccountPersonIDNumberType = "kz_pp"
	V2CoreAccountPersonIDNumberTypeLbPp        V2CoreAccountPersonIDNumberType = "lb_pp"
	V2CoreAccountPersonIDNumberTypeLiPeid      V2CoreAccountPersonIDNumberType = "li_peid"
	V2CoreAccountPersonIDNumberTypeLiPp        V2CoreAccountPersonIDNumberType = "li_pp"
	V2CoreAccountPersonIDNumberTypeLkPp        V2CoreAccountPersonIDNumberType = "lk_pp"
	V2CoreAccountPersonIDNumberTypeLTAk        V2CoreAccountPersonIDNumberType = "lt_ak"
	V2CoreAccountPersonIDNumberTypeLTPp        V2CoreAccountPersonIDNumberType = "lt_pp"
	V2CoreAccountPersonIDNumberTypeLuNif       V2CoreAccountPersonIDNumberType = "lu_nif"
	V2CoreAccountPersonIDNumberTypeLuPp        V2CoreAccountPersonIDNumberType = "lu_pp"
	V2CoreAccountPersonIDNumberTypeLVPk        V2CoreAccountPersonIDNumberType = "lv_pk"
	V2CoreAccountPersonIDNumberTypeLVPp        V2CoreAccountPersonIDNumberType = "lv_pp"
	V2CoreAccountPersonIDNumberTypeMaPp        V2CoreAccountPersonIDNumberType = "ma_pp"
	V2CoreAccountPersonIDNumberTypeMdCf        V2CoreAccountPersonIDNumberType = "md_cf"
	V2CoreAccountPersonIDNumberTypeMdPp        V2CoreAccountPersonIDNumberType = "md_pp"
	V2CoreAccountPersonIDNumberTypeMgNif       V2CoreAccountPersonIDNumberType = "mg_nif"
	V2CoreAccountPersonIDNumberTypeMkPIN       V2CoreAccountPersonIDNumberType = "mk_pin"
	V2CoreAccountPersonIDNumberTypeMkPp        V2CoreAccountPersonIDNumberType = "mk_pp"
	V2CoreAccountPersonIDNumberTypeMmPp        V2CoreAccountPersonIDNumberType = "mm_pp"
	V2CoreAccountPersonIDNumberTypeMnPp        V2CoreAccountPersonIDNumberType = "mn_pp"
	V2CoreAccountPersonIDNumberTypeMTPp        V2CoreAccountPersonIDNumberType = "mt_pp"
	V2CoreAccountPersonIDNumberTypeMuPp        V2CoreAccountPersonIDNumberType = "mu_pp"
	V2CoreAccountPersonIDNumberTypeMXCurp      V2CoreAccountPersonIDNumberType = "mx_curp"
	V2CoreAccountPersonIDNumberTypeMXPp        V2CoreAccountPersonIDNumberType = "mx_pp"
	V2CoreAccountPersonIDNumberTypeMXRFC       V2CoreAccountPersonIDNumberType = "mx_rfc"
	V2CoreAccountPersonIDNumberTypeMyNric      V2CoreAccountPersonIDNumberType = "my_nric"
	V2CoreAccountPersonIDNumberTypeMyPp        V2CoreAccountPersonIDNumberType = "my_pp"
	V2CoreAccountPersonIDNumberTypeMzNuit      V2CoreAccountPersonIDNumberType = "mz_nuit"
	V2CoreAccountPersonIDNumberTypeNgNin       V2CoreAccountPersonIDNumberType = "ng_nin"
	V2CoreAccountPersonIDNumberTypeNgPp        V2CoreAccountPersonIDNumberType = "ng_pp"
	V2CoreAccountPersonIDNumberTypeNLBsn       V2CoreAccountPersonIDNumberType = "nl_bsn"
	V2CoreAccountPersonIDNumberTypeNLPp        V2CoreAccountPersonIDNumberType = "nl_pp"
	V2CoreAccountPersonIDNumberTypeNoNin       V2CoreAccountPersonIDNumberType = "no_nin"
	V2CoreAccountPersonIDNumberTypeNoPp        V2CoreAccountPersonIDNumberType = "no_pp"
	V2CoreAccountPersonIDNumberTypeNpPp        V2CoreAccountPersonIDNumberType = "np_pp"
	V2CoreAccountPersonIDNumberTypeNzIrd       V2CoreAccountPersonIDNumberType = "nz_ird"
	V2CoreAccountPersonIDNumberTypeNzPp        V2CoreAccountPersonIDNumberType = "nz_pp"
	V2CoreAccountPersonIDNumberTypePaPp        V2CoreAccountPersonIDNumberType = "pa_pp"
	V2CoreAccountPersonIDNumberTypePeDni       V2CoreAccountPersonIDNumberType = "pe_dni"
	V2CoreAccountPersonIDNumberTypePePp        V2CoreAccountPersonIDNumberType = "pe_pp"
	V2CoreAccountPersonIDNumberTypePhPp        V2CoreAccountPersonIDNumberType = "ph_pp"
	V2CoreAccountPersonIDNumberTypePkCnic      V2CoreAccountPersonIDNumberType = "pk_cnic"
	V2CoreAccountPersonIDNumberTypePkPp        V2CoreAccountPersonIDNumberType = "pk_pp"
	V2CoreAccountPersonIDNumberTypePkSnic      V2CoreAccountPersonIDNumberType = "pk_snic"
	V2CoreAccountPersonIDNumberTypePLPesel     V2CoreAccountPersonIDNumberType = "pl_pesel"
	V2CoreAccountPersonIDNumberTypePLPp        V2CoreAccountPersonIDNumberType = "pl_pp"
	V2CoreAccountPersonIDNumberTypePrPp        V2CoreAccountPersonIDNumberType = "pr_pp"
	V2CoreAccountPersonIDNumberTypePsPp        V2CoreAccountPersonIDNumberType = "ps_pp"
	V2CoreAccountPersonIDNumberTypePTNif       V2CoreAccountPersonIDNumberType = "pt_nif"
	V2CoreAccountPersonIDNumberTypePTPp        V2CoreAccountPersonIDNumberType = "pt_pp"
	V2CoreAccountPersonIDNumberTypePyCi        V2CoreAccountPersonIDNumberType = "py_ci"
	V2CoreAccountPersonIDNumberTypePyPp        V2CoreAccountPersonIDNumberType = "py_pp"
	V2CoreAccountPersonIDNumberTypeROCnp       V2CoreAccountPersonIDNumberType = "ro_cnp"
	V2CoreAccountPersonIDNumberTypeROPp        V2CoreAccountPersonIDNumberType = "ro_pp"
	V2CoreAccountPersonIDNumberTypeRsPp        V2CoreAccountPersonIDNumberType = "rs_pp"
	V2CoreAccountPersonIDNumberTypeRUPp        V2CoreAccountPersonIDNumberType = "ru_pp"
	V2CoreAccountPersonIDNumberTypeRwTin       V2CoreAccountPersonIDNumberType = "rw_tin"
	V2CoreAccountPersonIDNumberTypeSaPp        V2CoreAccountPersonIDNumberType = "sa_pp"
	V2CoreAccountPersonIDNumberTypeSaPtin      V2CoreAccountPersonIDNumberType = "sa_ptin"
	V2CoreAccountPersonIDNumberTypeSaTin       V2CoreAccountPersonIDNumberType = "sa_tin"
	V2CoreAccountPersonIDNumberTypeSePIN       V2CoreAccountPersonIDNumberType = "se_pin"
	V2CoreAccountPersonIDNumberTypeSePp        V2CoreAccountPersonIDNumberType = "se_pp"
	V2CoreAccountPersonIDNumberTypeSgFin       V2CoreAccountPersonIDNumberType = "sg_fin"
	V2CoreAccountPersonIDNumberTypeSGNRIC      V2CoreAccountPersonIDNumberType = "sg_nric"
	V2CoreAccountPersonIDNumberTypeSgPp        V2CoreAccountPersonIDNumberType = "sg_pp"
	V2CoreAccountPersonIDNumberTypeSiPp        V2CoreAccountPersonIDNumberType = "si_pp"
	V2CoreAccountPersonIDNumberTypeSiPtin      V2CoreAccountPersonIDNumberType = "si_ptin"
	V2CoreAccountPersonIDNumberTypeSKDic       V2CoreAccountPersonIDNumberType = "sk_dic"
	V2CoreAccountPersonIDNumberTypeSKPp        V2CoreAccountPersonIDNumberType = "sk_pp"
	V2CoreAccountPersonIDNumberTypeSVPp        V2CoreAccountPersonIDNumberType = "sv_pp"
	V2CoreAccountPersonIDNumberTypeTHLc        V2CoreAccountPersonIDNumberType = "th_lc"
	V2CoreAccountPersonIDNumberTypeTHPIN       V2CoreAccountPersonIDNumberType = "th_pin"
	V2CoreAccountPersonIDNumberTypeTHPp        V2CoreAccountPersonIDNumberType = "th_pp"
	V2CoreAccountPersonIDNumberTypeTnPp        V2CoreAccountPersonIDNumberType = "tn_pp"
	V2CoreAccountPersonIDNumberTypeTRPp        V2CoreAccountPersonIDNumberType = "tr_pp"
	V2CoreAccountPersonIDNumberTypeTRTIN       V2CoreAccountPersonIDNumberType = "tr_tin"
	V2CoreAccountPersonIDNumberTypeTtBir       V2CoreAccountPersonIDNumberType = "tt_bir"
	V2CoreAccountPersonIDNumberTypeTwPp        V2CoreAccountPersonIDNumberType = "tw_pp"
	V2CoreAccountPersonIDNumberTypeUaPp        V2CoreAccountPersonIDNumberType = "ua_pp"
	V2CoreAccountPersonIDNumberTypeUgPp        V2CoreAccountPersonIDNumberType = "ug_pp"
	V2CoreAccountPersonIDNumberTypeUSItin      V2CoreAccountPersonIDNumberType = "us_itin"
	V2CoreAccountPersonIDNumberTypeUSItinLast4 V2CoreAccountPersonIDNumberType = "us_itin_last_4"
	V2CoreAccountPersonIDNumberTypeUSPp        V2CoreAccountPersonIDNumberType = "us_pp"
	V2CoreAccountPersonIDNumberTypeUSSSN       V2CoreAccountPersonIDNumberType = "us_ssn"
	V2CoreAccountPersonIDNumberTypeUSSSNLast4  V2CoreAccountPersonIDNumberType = "us_ssn_last_4"
	V2CoreAccountPersonIDNumberTypeUyDni       V2CoreAccountPersonIDNumberType = "uy_dni"
	V2CoreAccountPersonIDNumberTypeUyPp        V2CoreAccountPersonIDNumberType = "uy_pp"
	V2CoreAccountPersonIDNumberTypeUzPp        V2CoreAccountPersonIDNumberType = "uz_pp"
	V2CoreAccountPersonIDNumberTypeUzTin       V2CoreAccountPersonIDNumberType = "uz_tin"
	V2CoreAccountPersonIDNumberTypeVePp        V2CoreAccountPersonIDNumberType = "ve_pp"
	V2CoreAccountPersonIDNumberTypeVnPp        V2CoreAccountPersonIDNumberType = "vn_pp"
	V2CoreAccountPersonIDNumberTypeZaID        V2CoreAccountPersonIDNumberType = "za_id"
	V2CoreAccountPersonIDNumberTypeZaPp        V2CoreAccountPersonIDNumberType = "za_pp"
)

// The person's gender (International regulations require either "male" or "female").
type V2CoreAccountPersonLegalGender string

// List of values that V2CoreAccountPersonLegalGender can take
const (
	V2CoreAccountPersonLegalGenderFemale V2CoreAccountPersonLegalGender = "female"
	V2CoreAccountPersonLegalGenderMale   V2CoreAccountPersonLegalGender = "male"
)

// The person's political exposure.
type V2CoreAccountPersonPoliticalExposure string

// List of values that V2CoreAccountPersonPoliticalExposure can take
const (
	V2CoreAccountPersonPoliticalExposureExisting V2CoreAccountPersonPoliticalExposure = "existing"
	V2CoreAccountPersonPoliticalExposureNone     V2CoreAccountPersonPoliticalExposure = "none"
)

// Additional addresses associated with the person.
type V2CoreAccountPersonAdditionalAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// Purpose of additional address.
	Purpose V2CoreAccountPersonAdditionalAddressPurpose `json:"purpose"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or district.
	Town string `json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CoreAccountPersonAdditionalName struct {
	// The individual's full name.
	FullName string `json:"full_name,omitempty"`
	// The individual's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose V2CoreAccountPersonAdditionalNamePurpose `json:"purpose"`
	// The individual's last or family name.
	Surname string `json:"surname,omitempty"`
}

// Stripe terms of service agreement.
type V2CoreAccountPersonAdditionalTermsOfServiceAccount struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date time.Time `json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP string `json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent string `json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountPersonAdditionalTermsOfService struct {
	// Stripe terms of service agreement.
	Account *V2CoreAccountPersonAdditionalTermsOfServiceAccount `json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountPersonAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or district.
	Town string `json:"town,omitempty"`
}

// The person's date of birth.
type V2CoreAccountPersonDateOfBirth struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountPersonDocumentsCompanyAuthorization struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/files/create) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsCompanyAuthorizationType `json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountPersonDocumentsPassport struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/files/create) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsPassportType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/files/create) tokens for the front and back of the verification document.
type V2CoreAccountPersonDocumentsPrimaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/files/create) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/files/create) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountPersonDocumentsPrimaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/files/create) tokens for the front and back of the verification document.
	FrontBack *V2CoreAccountPersonDocumentsPrimaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CoreAccountPersonDocumentsPrimaryVerificationType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/files/create) tokens for the front and back of the verification document.
type V2CoreAccountPersonDocumentsSecondaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/files/create) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/files/create) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountPersonDocumentsSecondaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/files/create) tokens for the front and back of the verification document.
	FrontBack *V2CoreAccountPersonDocumentsSecondaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CoreAccountPersonDocumentsSecondaryVerificationType `json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountPersonDocumentsVisa struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/files/create) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsVisaType `json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountPersonDocuments struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountPersonDocumentsCompanyAuthorization `json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountPersonDocumentsPassport `json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountPersonDocumentsPrimaryVerification `json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountPersonDocumentsSecondaryVerification `json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountPersonDocumentsVisa `json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountPersonIDNumber struct {
	// The ID number type of an individual.
	Type V2CoreAccountPersonIDNumberType `json:"type"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountPersonRelationship struct {
	// Whether the individual is an authorizer of the Account's identity.
	Authorizer bool `json:"authorizer,omitempty"`
	// Whether the individual is a director of the Account's identity. Directors are typically members of the governing board of the company or are responsible for making sure that the company meets its regulatory obligations.
	Director bool `json:"director,omitempty"`
	// Whether the individual has significant responsibility to control, manage, or direct the organization.
	Executive bool `json:"executive,omitempty"`
	// Whether the individual is the legal guardian of the Account's representative.
	LegalGuardian bool `json:"legal_guardian,omitempty"`
	// Whether the individual is an owner of the Account's identity.
	Owner bool `json:"owner,omitempty"`
	// The percentage of the Account's identity that the individual owns.
	PercentOwnership decimal.Decimal `json:"percent_ownership,omitempty"`
	// Whether the individual is authorized as the primary representative of the Account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.
	Representative bool `json:"representative,omitempty"`
	// The individual's title (e.g., CEO, Support Engineer).
	Title string `json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountPersonScriptAddressesKana struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or district.
	Town string `json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountPersonScriptAddressesKanji struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or district.
	Town string `json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CoreAccountPersonScriptAddresses struct {
	// Kana Address.
	Kana *V2CoreAccountPersonScriptAddressesKana `json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountPersonScriptAddressesKanji `json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountPersonScriptNamesKana struct {
	// The person's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The person's last or family name.
	Surname string `json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountPersonScriptNamesKanji struct {
	// The person's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The person's last or family name.
	Surname string `json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountPersonScriptNames struct {
	// Persons name in kana script.
	Kana *V2CoreAccountPersonScriptNamesKana `json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountPersonScriptNamesKanji `json:"kanji,omitempty"`
}

// A Person represents an individual associated with an Account's identity (for example, an owner, director, executive, or representative). Use Persons to provide and update identity information for verification and compliance.
type V2CoreAccountPerson struct {
	APIResource
	// The account ID which the individual belongs to.
	Account string `json:"account"`
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountPersonAdditionalAddress `json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountPersonAdditionalName `json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountPersonAdditionalTermsOfService `json:"additional_terms_of_service,omitempty"`
	// The person's residential address.
	Address *V2CoreAccountPersonAddress `json:"address,omitempty"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountPersonDateOfBirth `json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountPersonDocuments `json:"documents,omitempty"`
	// The person's email address.
	Email string `json:"email,omitempty"`
	// The person's first name.
	GivenName string `json:"given_name,omitempty"`
	// Unique identifier for the Person.
	ID string `json:"id"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountPersonIDNumber `json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender V2CoreAccountPersonLegalGender `json:"legal_gender,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The countries where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []string `json:"nationalities,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The person's phone number.
	Phone string `json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure V2CoreAccountPersonPoliticalExposure `json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountPersonRelationship `json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountPersonScriptAddresses `json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountPersonScriptNames `json:"script_names,omitempty"`
	// The person's last name.
	Surname string `json:"surname,omitempty"`
	// Time at which the object was last updated. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Updated time.Time `json:"updated"`
}
