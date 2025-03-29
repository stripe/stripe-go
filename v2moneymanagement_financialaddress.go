//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of Credentials that are provisioned for the FinancialAddress.
type V2MoneyManagementFinancialAddressCredentialsType string

// List of values that V2MoneyManagementFinancialAddressCredentialsType can take
const (
	V2MoneyManagementFinancialAddressCredentialsTypeGBBankAccount V2MoneyManagementFinancialAddressCredentialsType = "gb_bank_account"
	V2MoneyManagementFinancialAddressCredentialsTypeUSBankAccount V2MoneyManagementFinancialAddressCredentialsType = "us_bank_account"
)

// Open Enum. The currency the FinancialAddress supports.
type V2MoneyManagementFinancialAddressCurrency string

// List of values that V2MoneyManagementFinancialAddressCurrency can take
const (
	V2MoneyManagementFinancialAddressCurrencyAed  V2MoneyManagementFinancialAddressCurrency = "aed"
	V2MoneyManagementFinancialAddressCurrencyAfn  V2MoneyManagementFinancialAddressCurrency = "afn"
	V2MoneyManagementFinancialAddressCurrencyAll  V2MoneyManagementFinancialAddressCurrency = "all"
	V2MoneyManagementFinancialAddressCurrencyAmd  V2MoneyManagementFinancialAddressCurrency = "amd"
	V2MoneyManagementFinancialAddressCurrencyAng  V2MoneyManagementFinancialAddressCurrency = "ang"
	V2MoneyManagementFinancialAddressCurrencyAoa  V2MoneyManagementFinancialAddressCurrency = "aoa"
	V2MoneyManagementFinancialAddressCurrencyArs  V2MoneyManagementFinancialAddressCurrency = "ars"
	V2MoneyManagementFinancialAddressCurrencyAUD  V2MoneyManagementFinancialAddressCurrency = "aud"
	V2MoneyManagementFinancialAddressCurrencyAwg  V2MoneyManagementFinancialAddressCurrency = "awg"
	V2MoneyManagementFinancialAddressCurrencyAzn  V2MoneyManagementFinancialAddressCurrency = "azn"
	V2MoneyManagementFinancialAddressCurrencyBam  V2MoneyManagementFinancialAddressCurrency = "bam"
	V2MoneyManagementFinancialAddressCurrencyBbd  V2MoneyManagementFinancialAddressCurrency = "bbd"
	V2MoneyManagementFinancialAddressCurrencyBdt  V2MoneyManagementFinancialAddressCurrency = "bdt"
	V2MoneyManagementFinancialAddressCurrencyBgn  V2MoneyManagementFinancialAddressCurrency = "bgn"
	V2MoneyManagementFinancialAddressCurrencyBhd  V2MoneyManagementFinancialAddressCurrency = "bhd"
	V2MoneyManagementFinancialAddressCurrencyBif  V2MoneyManagementFinancialAddressCurrency = "bif"
	V2MoneyManagementFinancialAddressCurrencyBmd  V2MoneyManagementFinancialAddressCurrency = "bmd"
	V2MoneyManagementFinancialAddressCurrencyBnd  V2MoneyManagementFinancialAddressCurrency = "bnd"
	V2MoneyManagementFinancialAddressCurrencyBob  V2MoneyManagementFinancialAddressCurrency = "bob"
	V2MoneyManagementFinancialAddressCurrencyBov  V2MoneyManagementFinancialAddressCurrency = "bov"
	V2MoneyManagementFinancialAddressCurrencyBrl  V2MoneyManagementFinancialAddressCurrency = "brl"
	V2MoneyManagementFinancialAddressCurrencyBsd  V2MoneyManagementFinancialAddressCurrency = "bsd"
	V2MoneyManagementFinancialAddressCurrencyBtn  V2MoneyManagementFinancialAddressCurrency = "btn"
	V2MoneyManagementFinancialAddressCurrencyBwp  V2MoneyManagementFinancialAddressCurrency = "bwp"
	V2MoneyManagementFinancialAddressCurrencyByn  V2MoneyManagementFinancialAddressCurrency = "byn"
	V2MoneyManagementFinancialAddressCurrencyByr  V2MoneyManagementFinancialAddressCurrency = "byr"
	V2MoneyManagementFinancialAddressCurrencyBzd  V2MoneyManagementFinancialAddressCurrency = "bzd"
	V2MoneyManagementFinancialAddressCurrencyCAD  V2MoneyManagementFinancialAddressCurrency = "cad"
	V2MoneyManagementFinancialAddressCurrencyCdf  V2MoneyManagementFinancialAddressCurrency = "cdf"
	V2MoneyManagementFinancialAddressCurrencyChe  V2MoneyManagementFinancialAddressCurrency = "che"
	V2MoneyManagementFinancialAddressCurrencyCHF  V2MoneyManagementFinancialAddressCurrency = "chf"
	V2MoneyManagementFinancialAddressCurrencyChw  V2MoneyManagementFinancialAddressCurrency = "chw"
	V2MoneyManagementFinancialAddressCurrencyClf  V2MoneyManagementFinancialAddressCurrency = "clf"
	V2MoneyManagementFinancialAddressCurrencyClp  V2MoneyManagementFinancialAddressCurrency = "clp"
	V2MoneyManagementFinancialAddressCurrencyCny  V2MoneyManagementFinancialAddressCurrency = "cny"
	V2MoneyManagementFinancialAddressCurrencyCop  V2MoneyManagementFinancialAddressCurrency = "cop"
	V2MoneyManagementFinancialAddressCurrencyCou  V2MoneyManagementFinancialAddressCurrency = "cou"
	V2MoneyManagementFinancialAddressCurrencyCrc  V2MoneyManagementFinancialAddressCurrency = "crc"
	V2MoneyManagementFinancialAddressCurrencyCuc  V2MoneyManagementFinancialAddressCurrency = "cuc"
	V2MoneyManagementFinancialAddressCurrencyCup  V2MoneyManagementFinancialAddressCurrency = "cup"
	V2MoneyManagementFinancialAddressCurrencyCve  V2MoneyManagementFinancialAddressCurrency = "cve"
	V2MoneyManagementFinancialAddressCurrencyCZK  V2MoneyManagementFinancialAddressCurrency = "czk"
	V2MoneyManagementFinancialAddressCurrencyDjf  V2MoneyManagementFinancialAddressCurrency = "djf"
	V2MoneyManagementFinancialAddressCurrencyDKK  V2MoneyManagementFinancialAddressCurrency = "dkk"
	V2MoneyManagementFinancialAddressCurrencyDop  V2MoneyManagementFinancialAddressCurrency = "dop"
	V2MoneyManagementFinancialAddressCurrencyDzd  V2MoneyManagementFinancialAddressCurrency = "dzd"
	V2MoneyManagementFinancialAddressCurrencyEek  V2MoneyManagementFinancialAddressCurrency = "eek"
	V2MoneyManagementFinancialAddressCurrencyEgp  V2MoneyManagementFinancialAddressCurrency = "egp"
	V2MoneyManagementFinancialAddressCurrencyErn  V2MoneyManagementFinancialAddressCurrency = "ern"
	V2MoneyManagementFinancialAddressCurrencyEtb  V2MoneyManagementFinancialAddressCurrency = "etb"
	V2MoneyManagementFinancialAddressCurrencyEUR  V2MoneyManagementFinancialAddressCurrency = "eur"
	V2MoneyManagementFinancialAddressCurrencyFjd  V2MoneyManagementFinancialAddressCurrency = "fjd"
	V2MoneyManagementFinancialAddressCurrencyFkp  V2MoneyManagementFinancialAddressCurrency = "fkp"
	V2MoneyManagementFinancialAddressCurrencyGBP  V2MoneyManagementFinancialAddressCurrency = "gbp"
	V2MoneyManagementFinancialAddressCurrencyGel  V2MoneyManagementFinancialAddressCurrency = "gel"
	V2MoneyManagementFinancialAddressCurrencyGhc  V2MoneyManagementFinancialAddressCurrency = "ghc"
	V2MoneyManagementFinancialAddressCurrencyGhs  V2MoneyManagementFinancialAddressCurrency = "ghs"
	V2MoneyManagementFinancialAddressCurrencyGip  V2MoneyManagementFinancialAddressCurrency = "gip"
	V2MoneyManagementFinancialAddressCurrencyGmd  V2MoneyManagementFinancialAddressCurrency = "gmd"
	V2MoneyManagementFinancialAddressCurrencyGnf  V2MoneyManagementFinancialAddressCurrency = "gnf"
	V2MoneyManagementFinancialAddressCurrencyGtq  V2MoneyManagementFinancialAddressCurrency = "gtq"
	V2MoneyManagementFinancialAddressCurrencyGyd  V2MoneyManagementFinancialAddressCurrency = "gyd"
	V2MoneyManagementFinancialAddressCurrencyHKD  V2MoneyManagementFinancialAddressCurrency = "hkd"
	V2MoneyManagementFinancialAddressCurrencyHnl  V2MoneyManagementFinancialAddressCurrency = "hnl"
	V2MoneyManagementFinancialAddressCurrencyHrk  V2MoneyManagementFinancialAddressCurrency = "hrk"
	V2MoneyManagementFinancialAddressCurrencyHtg  V2MoneyManagementFinancialAddressCurrency = "htg"
	V2MoneyManagementFinancialAddressCurrencyHuf  V2MoneyManagementFinancialAddressCurrency = "huf"
	V2MoneyManagementFinancialAddressCurrencyIdr  V2MoneyManagementFinancialAddressCurrency = "idr"
	V2MoneyManagementFinancialAddressCurrencyIls  V2MoneyManagementFinancialAddressCurrency = "ils"
	V2MoneyManagementFinancialAddressCurrencyInr  V2MoneyManagementFinancialAddressCurrency = "inr"
	V2MoneyManagementFinancialAddressCurrencyIqd  V2MoneyManagementFinancialAddressCurrency = "iqd"
	V2MoneyManagementFinancialAddressCurrencyIrr  V2MoneyManagementFinancialAddressCurrency = "irr"
	V2MoneyManagementFinancialAddressCurrencyIsk  V2MoneyManagementFinancialAddressCurrency = "isk"
	V2MoneyManagementFinancialAddressCurrencyJmd  V2MoneyManagementFinancialAddressCurrency = "jmd"
	V2MoneyManagementFinancialAddressCurrencyJod  V2MoneyManagementFinancialAddressCurrency = "jod"
	V2MoneyManagementFinancialAddressCurrencyJPY  V2MoneyManagementFinancialAddressCurrency = "jpy"
	V2MoneyManagementFinancialAddressCurrencyKes  V2MoneyManagementFinancialAddressCurrency = "kes"
	V2MoneyManagementFinancialAddressCurrencyKgs  V2MoneyManagementFinancialAddressCurrency = "kgs"
	V2MoneyManagementFinancialAddressCurrencyKhr  V2MoneyManagementFinancialAddressCurrency = "khr"
	V2MoneyManagementFinancialAddressCurrencyKmf  V2MoneyManagementFinancialAddressCurrency = "kmf"
	V2MoneyManagementFinancialAddressCurrencyKpw  V2MoneyManagementFinancialAddressCurrency = "kpw"
	V2MoneyManagementFinancialAddressCurrencyKrw  V2MoneyManagementFinancialAddressCurrency = "krw"
	V2MoneyManagementFinancialAddressCurrencyKwd  V2MoneyManagementFinancialAddressCurrency = "kwd"
	V2MoneyManagementFinancialAddressCurrencyKyd  V2MoneyManagementFinancialAddressCurrency = "kyd"
	V2MoneyManagementFinancialAddressCurrencyKzt  V2MoneyManagementFinancialAddressCurrency = "kzt"
	V2MoneyManagementFinancialAddressCurrencyLak  V2MoneyManagementFinancialAddressCurrency = "lak"
	V2MoneyManagementFinancialAddressCurrencyLbp  V2MoneyManagementFinancialAddressCurrency = "lbp"
	V2MoneyManagementFinancialAddressCurrencyLkr  V2MoneyManagementFinancialAddressCurrency = "lkr"
	V2MoneyManagementFinancialAddressCurrencyLrd  V2MoneyManagementFinancialAddressCurrency = "lrd"
	V2MoneyManagementFinancialAddressCurrencyLsl  V2MoneyManagementFinancialAddressCurrency = "lsl"
	V2MoneyManagementFinancialAddressCurrencyLtl  V2MoneyManagementFinancialAddressCurrency = "ltl"
	V2MoneyManagementFinancialAddressCurrencyLvl  V2MoneyManagementFinancialAddressCurrency = "lvl"
	V2MoneyManagementFinancialAddressCurrencyLyd  V2MoneyManagementFinancialAddressCurrency = "lyd"
	V2MoneyManagementFinancialAddressCurrencyMad  V2MoneyManagementFinancialAddressCurrency = "mad"
	V2MoneyManagementFinancialAddressCurrencyMdl  V2MoneyManagementFinancialAddressCurrency = "mdl"
	V2MoneyManagementFinancialAddressCurrencyMga  V2MoneyManagementFinancialAddressCurrency = "mga"
	V2MoneyManagementFinancialAddressCurrencyMkd  V2MoneyManagementFinancialAddressCurrency = "mkd"
	V2MoneyManagementFinancialAddressCurrencyMmk  V2MoneyManagementFinancialAddressCurrency = "mmk"
	V2MoneyManagementFinancialAddressCurrencyMnt  V2MoneyManagementFinancialAddressCurrency = "mnt"
	V2MoneyManagementFinancialAddressCurrencyMop  V2MoneyManagementFinancialAddressCurrency = "mop"
	V2MoneyManagementFinancialAddressCurrencyMro  V2MoneyManagementFinancialAddressCurrency = "mro"
	V2MoneyManagementFinancialAddressCurrencyMru  V2MoneyManagementFinancialAddressCurrency = "mru"
	V2MoneyManagementFinancialAddressCurrencyMur  V2MoneyManagementFinancialAddressCurrency = "mur"
	V2MoneyManagementFinancialAddressCurrencyMvr  V2MoneyManagementFinancialAddressCurrency = "mvr"
	V2MoneyManagementFinancialAddressCurrencyMwk  V2MoneyManagementFinancialAddressCurrency = "mwk"
	V2MoneyManagementFinancialAddressCurrencyMxn  V2MoneyManagementFinancialAddressCurrency = "mxn"
	V2MoneyManagementFinancialAddressCurrencyMxv  V2MoneyManagementFinancialAddressCurrency = "mxv"
	V2MoneyManagementFinancialAddressCurrencyMYR  V2MoneyManagementFinancialAddressCurrency = "myr"
	V2MoneyManagementFinancialAddressCurrencyMzn  V2MoneyManagementFinancialAddressCurrency = "mzn"
	V2MoneyManagementFinancialAddressCurrencyNad  V2MoneyManagementFinancialAddressCurrency = "nad"
	V2MoneyManagementFinancialAddressCurrencyNgn  V2MoneyManagementFinancialAddressCurrency = "ngn"
	V2MoneyManagementFinancialAddressCurrencyNio  V2MoneyManagementFinancialAddressCurrency = "nio"
	V2MoneyManagementFinancialAddressCurrencyNOK  V2MoneyManagementFinancialAddressCurrency = "nok"
	V2MoneyManagementFinancialAddressCurrencyNpr  V2MoneyManagementFinancialAddressCurrency = "npr"
	V2MoneyManagementFinancialAddressCurrencyNZD  V2MoneyManagementFinancialAddressCurrency = "nzd"
	V2MoneyManagementFinancialAddressCurrencyOmr  V2MoneyManagementFinancialAddressCurrency = "omr"
	V2MoneyManagementFinancialAddressCurrencyPab  V2MoneyManagementFinancialAddressCurrency = "pab"
	V2MoneyManagementFinancialAddressCurrencyPen  V2MoneyManagementFinancialAddressCurrency = "pen"
	V2MoneyManagementFinancialAddressCurrencyPgk  V2MoneyManagementFinancialAddressCurrency = "pgk"
	V2MoneyManagementFinancialAddressCurrencyPhp  V2MoneyManagementFinancialAddressCurrency = "php"
	V2MoneyManagementFinancialAddressCurrencyPkr  V2MoneyManagementFinancialAddressCurrency = "pkr"
	V2MoneyManagementFinancialAddressCurrencyPLN  V2MoneyManagementFinancialAddressCurrency = "pln"
	V2MoneyManagementFinancialAddressCurrencyPyg  V2MoneyManagementFinancialAddressCurrency = "pyg"
	V2MoneyManagementFinancialAddressCurrencyQar  V2MoneyManagementFinancialAddressCurrency = "qar"
	V2MoneyManagementFinancialAddressCurrencyRon  V2MoneyManagementFinancialAddressCurrency = "ron"
	V2MoneyManagementFinancialAddressCurrencyRsd  V2MoneyManagementFinancialAddressCurrency = "rsd"
	V2MoneyManagementFinancialAddressCurrencyRub  V2MoneyManagementFinancialAddressCurrency = "rub"
	V2MoneyManagementFinancialAddressCurrencyRwf  V2MoneyManagementFinancialAddressCurrency = "rwf"
	V2MoneyManagementFinancialAddressCurrencySar  V2MoneyManagementFinancialAddressCurrency = "sar"
	V2MoneyManagementFinancialAddressCurrencySbd  V2MoneyManagementFinancialAddressCurrency = "sbd"
	V2MoneyManagementFinancialAddressCurrencyScr  V2MoneyManagementFinancialAddressCurrency = "scr"
	V2MoneyManagementFinancialAddressCurrencySdg  V2MoneyManagementFinancialAddressCurrency = "sdg"
	V2MoneyManagementFinancialAddressCurrencySEK  V2MoneyManagementFinancialAddressCurrency = "sek"
	V2MoneyManagementFinancialAddressCurrencySGD  V2MoneyManagementFinancialAddressCurrency = "sgd"
	V2MoneyManagementFinancialAddressCurrencyShp  V2MoneyManagementFinancialAddressCurrency = "shp"
	V2MoneyManagementFinancialAddressCurrencySle  V2MoneyManagementFinancialAddressCurrency = "sle"
	V2MoneyManagementFinancialAddressCurrencySll  V2MoneyManagementFinancialAddressCurrency = "sll"
	V2MoneyManagementFinancialAddressCurrencySos  V2MoneyManagementFinancialAddressCurrency = "sos"
	V2MoneyManagementFinancialAddressCurrencySrd  V2MoneyManagementFinancialAddressCurrency = "srd"
	V2MoneyManagementFinancialAddressCurrencySsp  V2MoneyManagementFinancialAddressCurrency = "ssp"
	V2MoneyManagementFinancialAddressCurrencyStd  V2MoneyManagementFinancialAddressCurrency = "std"
	V2MoneyManagementFinancialAddressCurrencyStn  V2MoneyManagementFinancialAddressCurrency = "stn"
	V2MoneyManagementFinancialAddressCurrencySvc  V2MoneyManagementFinancialAddressCurrency = "svc"
	V2MoneyManagementFinancialAddressCurrencySyp  V2MoneyManagementFinancialAddressCurrency = "syp"
	V2MoneyManagementFinancialAddressCurrencySzl  V2MoneyManagementFinancialAddressCurrency = "szl"
	V2MoneyManagementFinancialAddressCurrencyThb  V2MoneyManagementFinancialAddressCurrency = "thb"
	V2MoneyManagementFinancialAddressCurrencyTjs  V2MoneyManagementFinancialAddressCurrency = "tjs"
	V2MoneyManagementFinancialAddressCurrencyTmt  V2MoneyManagementFinancialAddressCurrency = "tmt"
	V2MoneyManagementFinancialAddressCurrencyTnd  V2MoneyManagementFinancialAddressCurrency = "tnd"
	V2MoneyManagementFinancialAddressCurrencyTop  V2MoneyManagementFinancialAddressCurrency = "top"
	V2MoneyManagementFinancialAddressCurrencyTry  V2MoneyManagementFinancialAddressCurrency = "try"
	V2MoneyManagementFinancialAddressCurrencyTtd  V2MoneyManagementFinancialAddressCurrency = "ttd"
	V2MoneyManagementFinancialAddressCurrencyTwd  V2MoneyManagementFinancialAddressCurrency = "twd"
	V2MoneyManagementFinancialAddressCurrencyTzs  V2MoneyManagementFinancialAddressCurrency = "tzs"
	V2MoneyManagementFinancialAddressCurrencyUah  V2MoneyManagementFinancialAddressCurrency = "uah"
	V2MoneyManagementFinancialAddressCurrencyUgx  V2MoneyManagementFinancialAddressCurrency = "ugx"
	V2MoneyManagementFinancialAddressCurrencyUSD  V2MoneyManagementFinancialAddressCurrency = "usd"
	V2MoneyManagementFinancialAddressCurrencyUsdc V2MoneyManagementFinancialAddressCurrency = "usdc"
	V2MoneyManagementFinancialAddressCurrencyUsn  V2MoneyManagementFinancialAddressCurrency = "usn"
	V2MoneyManagementFinancialAddressCurrencyUyi  V2MoneyManagementFinancialAddressCurrency = "uyi"
	V2MoneyManagementFinancialAddressCurrencyUyu  V2MoneyManagementFinancialAddressCurrency = "uyu"
	V2MoneyManagementFinancialAddressCurrencyUzs  V2MoneyManagementFinancialAddressCurrency = "uzs"
	V2MoneyManagementFinancialAddressCurrencyVef  V2MoneyManagementFinancialAddressCurrency = "vef"
	V2MoneyManagementFinancialAddressCurrencyVes  V2MoneyManagementFinancialAddressCurrency = "ves"
	V2MoneyManagementFinancialAddressCurrencyVnd  V2MoneyManagementFinancialAddressCurrency = "vnd"
	V2MoneyManagementFinancialAddressCurrencyVuv  V2MoneyManagementFinancialAddressCurrency = "vuv"
	V2MoneyManagementFinancialAddressCurrencyWst  V2MoneyManagementFinancialAddressCurrency = "wst"
	V2MoneyManagementFinancialAddressCurrencyXaf  V2MoneyManagementFinancialAddressCurrency = "xaf"
	V2MoneyManagementFinancialAddressCurrencyXcd  V2MoneyManagementFinancialAddressCurrency = "xcd"
	V2MoneyManagementFinancialAddressCurrencyXcg  V2MoneyManagementFinancialAddressCurrency = "xcg"
	V2MoneyManagementFinancialAddressCurrencyXof  V2MoneyManagementFinancialAddressCurrency = "xof"
	V2MoneyManagementFinancialAddressCurrencyXpf  V2MoneyManagementFinancialAddressCurrency = "xpf"
	V2MoneyManagementFinancialAddressCurrencyYer  V2MoneyManagementFinancialAddressCurrency = "yer"
	V2MoneyManagementFinancialAddressCurrencyZar  V2MoneyManagementFinancialAddressCurrency = "zar"
	V2MoneyManagementFinancialAddressCurrencyZmk  V2MoneyManagementFinancialAddressCurrency = "zmk"
	V2MoneyManagementFinancialAddressCurrencyZmw  V2MoneyManagementFinancialAddressCurrency = "zmw"
	V2MoneyManagementFinancialAddressCurrencyZwd  V2MoneyManagementFinancialAddressCurrency = "zwd"
	V2MoneyManagementFinancialAddressCurrencyZwg  V2MoneyManagementFinancialAddressCurrency = "zwg"
	V2MoneyManagementFinancialAddressCurrencyZwl  V2MoneyManagementFinancialAddressCurrency = "zwl"
)

// Closed Enum. An enum representing the status of the FinancialAddress. This indicates whether or not the FinancialAddress can be used for any money movement flows.
type V2MoneyManagementFinancialAddressStatus string

// List of values that V2MoneyManagementFinancialAddressStatus can take
const (
	V2MoneyManagementFinancialAddressStatusActive   V2MoneyManagementFinancialAddressStatus = "active"
	V2MoneyManagementFinancialAddressStatusArchived V2MoneyManagementFinancialAddressStatus = "archived"
	V2MoneyManagementFinancialAddressStatusFailed   V2MoneyManagementFinancialAddressStatus = "failed"
	V2MoneyManagementFinancialAddressStatusPending  V2MoneyManagementFinancialAddressStatus = "pending"
)

// The credentials of the UK Bank Account for the FinancialAddress. This contains unique banking details such as the sort code, account number, etc. of a UK bank account.
type V2MoneyManagementFinancialAddressCredentialsGBBankAccount struct {
	// The account holder name to be used during bank transference.
	AccountHolderName string `json:"account_holder_name"`
	// The account number of the UK Bank Account.
	AccountNumber string `json:"account_number"`
	// The last four digits of the UK Bank Account number. This will always be returned.
	// To view the full account number when retrieving or listing FinancialAddresses, use the `include` request parameter.
	Last4 string `json:"last4"`
	// The sort code of the UK Bank Account.
	SortCode string `json:"sort_code"`
}

// The credentials of the US Bank Account for the FinancialAddress. This contains unique banking details such as the routing number, account number, etc. of a US bank account.
type V2MoneyManagementFinancialAddressCredentialsUSBankAccount struct {
	// The account number of the US Bank Account.
	AccountNumber string `json:"account_number"`
	// The name of the Bank.
	BankName string `json:"bank_name"`
	// The last four digits of the US Bank Account number. This will always be returned.
	// To view the full account number when retrieving or listing FinancialAddresses, use the `include` request parameter.
	Last4 string `json:"last4"`
	// The routing number of the US Bank Account.
	RoutingNumber string `json:"routing_number"`
	// The swift code of the bank or financial institution.
	SwiftCode string `json:"swift_code"`
}

// Object indicates the type of credentials that have been allocated and attached to the FinancialAddress.
// It contains all necessary banking details with which to perform money movements with the FinancialAddress.
// This field is only available for FinancialAddresses with an active status.
type V2MoneyManagementFinancialAddressCredentials struct {
	// The credentials of the UK Bank Account for the FinancialAddress. This contains unique banking details such as the sort code, account number, etc. of a UK bank account.
	GBBankAccount *V2MoneyManagementFinancialAddressCredentialsGBBankAccount `json:"gb_bank_account"`
	// Open Enum. The type of Credentials that are provisioned for the FinancialAddress.
	Type V2MoneyManagementFinancialAddressCredentialsType `json:"type"`
	// The credentials of the US Bank Account for the FinancialAddress. This contains unique banking details such as the routing number, account number, etc. of a US bank account.
	USBankAccount *V2MoneyManagementFinancialAddressCredentialsUSBankAccount `json:"us_bank_account"`
}

// A FinancialAddress contains information needed to transfer money to a Financial Account. A Financial Account can have more than one Financial Address.
type V2MoneyManagementFinancialAddress struct {
	APIResource
	// The creation timestamp of the FinancialAddress.
	Created time.Time `json:"created"`
	// Object indicates the type of credentials that have been allocated and attached to the FinancialAddress.
	// It contains all necessary banking details with which to perform money movements with the FinancialAddress.
	// This field is only available for FinancialAddresses with an active status.
	Credentials *V2MoneyManagementFinancialAddressCredentials `json:"credentials"`
	// Open Enum. The currency the FinancialAddress supports.
	Currency V2MoneyManagementFinancialAddressCurrency `json:"currency"`
	// A ID of the FinancialAccount this FinancialAddress corresponds to.
	FinancialAccount string `json:"financial_account"`
	// The ID of a FinancialAddress.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. An enum representing the status of the FinancialAddress. This indicates whether or not the FinancialAddress can be used for any money movement flows.
	Status V2MoneyManagementFinancialAddressStatus `json:"status"`
}
