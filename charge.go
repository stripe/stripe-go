package stripe

import "encoding/json"

// Currency is the list of supported currencies.
// For more details see https://support.stripe.com/questions/which-currencies-does-stripe-support.
type Currency string

const (
	AED Currency = "aed" // United Arab Emirates Dirham
	AFN Currency = "afn" // Afghan Afghani
	ALL Currency = "all" // Albanian Lek
	AMD Currency = "amd" // Armenian Dram
	ANG Currency = "ang" // Netherlands Antillean Gulden
	AOA Currency = "aoa" // Angolan Kwanza
	ARS Currency = "ars" // Argentine Peso
	AUD Currency = "aud" // Australian Dollar
	AWG Currency = "awg" // Aruban Florin
	AZN Currency = "azn" // Azerbaijani Manat
	BAM Currency = "bam" // Bosnia & Herzegovina Convertible Mark
	BBD Currency = "bbd" // Barbadian Dollar
	BDT Currency = "bdt" // Bangladeshi Taka
	BGN Currency = "bgn" // Bulgarian Lev
	BIF Currency = "bif" // Burundian Franc
	BMD Currency = "bmd" // Bermudian Dollar
	BND Currency = "bnd" // Brunei Dollar
	BOB Currency = "bob" // Bolivian Boliviano
	BRL Currency = "brl" // Brazilian Real
	BSD Currency = "bsd" // Bahamian Dollar
	BWP Currency = "bwp" // Botswana Pula
	BZD Currency = "bzd" // Belize Dollar
	CAD Currency = "cad" // Canadian Dollar
	CDF Currency = "cdf" // Congolese Franc
	CHF Currency = "chf" // Swiss Franc
	CLP Currency = "clp" // Chilean Peso
	CNY Currency = "cny" // Chinese Renminbi Yuan
	COP Currency = "cop" // Colombian Peso
	CRC Currency = "crc" // Costa Rican Colón
	CVE Currency = "cve" // Cape Verdean Escudo
	CZK Currency = "czk" // Czech Koruna
	DJF Currency = "djf" // Djiboutian Franc
	DKK Currency = "dkk" // Danish Krone
	DOP Currency = "dop" // Dominican Peso
	DZD Currency = "dzd" // Algerian Dinar
	EEK Currency = "eek" // Estonian Kroon
	EGP Currency = "egp" // Egyptian Pound
	ETB Currency = "etb" // Ethiopian Birr
	EUR Currency = "eur" // Euro
	FJD Currency = "fjd" // Fijian Dollar
	FKP Currency = "fkp" // Falkland Islands Pound
	GBP Currency = "gbp" // British Pound
	GEL Currency = "gel" // Georgian Lari
	GIP Currency = "gip" // Gibraltar Pound
	GMD Currency = "gmd" // Gambian Dalasi
	GNF Currency = "gnf" // Guinean Franc
	GTQ Currency = "gtq" // Guatemalan Quetzal
	GYD Currency = "gyd" // Guyanese Dollar
	HKD Currency = "hkd" // Hong Kong Dollar
	HNL Currency = "hnl" // Honduran Lempira
	HRK Currency = "hrk" // Croatian Kuna
	HTG Currency = "htg" // Haitian Gourde
	HUF Currency = "huf" // Hungarian Forint
	IDR Currency = "idr" // Indonesian Rupiah
	ILS Currency = "ils" // Israeli New Sheqel
	INR Currency = "inr" // Indian Rupee
	ISK Currency = "isk" // Icelandic Króna
	JMD Currency = "jmd" // Jamaican Dollar
	JPY Currency = "jpy" // Japanese Yen
	KES Currency = "kes" // Kenyan Shilling
	KGS Currency = "kgs" // Kyrgyzstani Som
	KHR Currency = "khr" // Cambodian Riel
	KMF Currency = "kmf" // Comorian Franc
	KRW Currency = "krw" // South Korean Won
	KYD Currency = "kyd" // Cayman Islands Dollar
	KZT Currency = "kzt" // Kazakhstani Tenge
	LAK Currency = "lak" // Lao Kip
	LBP Currency = "lbp" // Lebanese Pound
	LKR Currency = "lkr" // Sri Lankan Rupee
	LRD Currency = "lrd" // Liberian Dollar
	LSL Currency = "lsl" // Lesotho Loti
	LTL Currency = "ltl" // Lithuanian Litas
	LVL Currency = "lvl" // Latvian Lats
	MAD Currency = "mad" // Moroccan Dirham
	MDL Currency = "mdl" // Moldovan Leu
	MGA Currency = "mga" // Malagasy Ariary
	MKD Currency = "mkd" // Macedonian Denar
	MNT Currency = "mnt" // Mongolian Tögrög
	MOP Currency = "mop" // Macanese Pataca
	MRO Currency = "mro" // Mauritanian Ouguiya
	MUR Currency = "mur" // Mauritian Rupee
	MVR Currency = "mvr" // Maldivian Rufiyaa
	MWK Currency = "mwk" // Malawian Kwacha
	MXN Currency = "mxn" // Mexican Peso
	MYR Currency = "myr" // Malaysian Ringgit
	MZN Currency = "mzn" // Mozambican Metical
	NAD Currency = "nad" // Namibian Dollar
	NGN Currency = "ngn" // Nigerian Naira
	NIO Currency = "nio" // Nicaraguan Córdoba
	NOK Currency = "nok" // Norwegian Krone
	NPR Currency = "npr" // Nepalese Rupee
	NZD Currency = "nzd" // New Zealand Dollar
	PAB Currency = "pab" // Panamanian Balboa
	PEN Currency = "pen" // Peruvian Nuevo Sol
	PGK Currency = "pgk" // Papua New Guinean Kina
	PHP Currency = "php" // Philippine Peso
	PKR Currency = "pkr" // Pakistani Rupee
	PLN Currency = "pln" // Polish Złoty
	PYG Currency = "pyg" // Paraguayan Guaraní
	QAR Currency = "qar" // Qatari Riyal
	RON Currency = "ron" // Romanian Leu
	RSD Currency = "rsd" // Serbian Dinar
	RUB Currency = "rub" // Russian Ruble
	RWF Currency = "rwf" // Rwandan Franc
	SAR Currency = "sar" // Saudi Riyal
	SBD Currency = "sbd" // Solomon Islands Dollar
	SCR Currency = "scr" // Seychellois Rupee
	SEK Currency = "sek" // Swedish Krona
	SGD Currency = "sgd" // Singapore Dollar
	SHP Currency = "shp" // Saint Helenian Pound
	SLL Currency = "sll" // Sierra Leonean Leone
	SOS Currency = "sos" // Somali Shilling
	SRD Currency = "srd" // Surinamese Dollar
	STD Currency = "std" // São Tomé and Príncipe Dobra
	SVC Currency = "svc" // Salvadoran Colón
	SZL Currency = "szl" // Swazi Lilangeni
	THB Currency = "thb" // Thai Baht
	TJS Currency = "tjs" // Tajikistani Somoni
	TOP Currency = "top" // Tongan Paʻanga
	TRY Currency = "try" // Turkish Lira
	TTD Currency = "ttd" // Trinidad and Tobago Dollar
	TWD Currency = "twd" // New Taiwan Dollar
	TZS Currency = "tzs" // Tanzanian Shilling
	UAH Currency = "uah" // Ukrainian Hryvnia
	UGX Currency = "ugx" // Ugandan Shilling
	USD Currency = "usd" // United States Dollar
	UYU Currency = "uyu" // Uruguayan Peso
	UZS Currency = "uzs" // Uzbekistani Som
	VEF Currency = "vef" // Venezuelan Bolívar
	VND Currency = "vnd" // Vietnamese Đồng
	VUV Currency = "vuv" // Vanuatu Vatu
	WST Currency = "wst" // Samoan Tala
	XAF Currency = "xaf" // Central African Cfa Franc
	XCD Currency = "xcd" // East Caribbean Dollar
	XOF Currency = "xof" // West African Cfa Franc
	XPF Currency = "xpf" // Cfp Franc
	YER Currency = "yer" // Yemeni Rial
	ZAR Currency = "zar" // South African Rand
	ZMW Currency = "zmw" // Zambian Kwacha
)

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
// For more details see https://stripe.com/docs/api#create_charge and https://stripe.com/docs/api#update_charge.
type ChargeParams struct {
	Params
	Amount                 uint64
	Currency               Currency
	Customer, Token        string
	Card                   *CardParams
	Desc, Statement, Email string
	NoCapture              bool
	Fee                    uint64
}

// ChargeListParams is the set of parameters that can be used when listing charges.
// For more details see https://stripe.com/docs/api#list_charges.
type ChargeListParams struct {
	ListParams
	Created  int64
	Customer string
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
// For more details see https://stripe.com/docs/api#charge_capture.
type CaptureParams struct {
	Params
	Amount, Fee uint64
	Email       string
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	Id             string            `json:"id"`
	Live           bool              `json:"livemode"`
	Amount         uint64            `json:"amount"`
	Captured       bool              `json:"captured"`
	Card           *Card             `json:"card"`
	Created        int64             `json:"created"`
	Currency       Currency          `json:"currency"`
	Paid           bool              `json:"paid"`
	Refunded       bool              `json:"refunded"`
	Refunds        *RefundList       `json:"refunds"`
	AmountRefunded uint64            `json:"amount_refunded"`
	Tx             *Transaction      `json:"balance_transaction"`
	Customer       *Customer         `json:"customer"`
	Desc           string            `json:"description"`
	Dispute        *Dispute          `json:"dispute"`
	FailMsg        string            `json:"failure_message"`
	FailCode       string            `json:"failure_code"`
	Invoice        *Invoice          `json:"invoice"`
	Meta           map[string]string `json:"metadata"`
	Email          string            `json:"receipt_email"`
	Statement      string            `json:"statement_description"`
}

// ChargeIter is a iterator for list responses.
type ChargeIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *ChargeIter) Next() (*Charge, error) {
	c, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return c.(*Charge), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *ChargeIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *ChargeIter) Meta() *ListMeta {
	return i.Iter.Meta()
}

func (c *Charge) UnmarshalJSON(data []byte) error {
	type charge Charge
	var cc charge
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Charge(cc)
	} else {
		// the id is surrounded by escaped \, so ignore those
		c.Id = string(data[1 : len(data)-1])
	}

	return nil
}
