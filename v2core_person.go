//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CorePersonAdditionalAddressCountry string

// List of values that V2CorePersonAdditionalAddressCountry can take
const (
	V2CorePersonAdditionalAddressCountryAd V2CorePersonAdditionalAddressCountry = "ad"
	V2CorePersonAdditionalAddressCountryAe V2CorePersonAdditionalAddressCountry = "ae"
	V2CorePersonAdditionalAddressCountryAf V2CorePersonAdditionalAddressCountry = "af"
	V2CorePersonAdditionalAddressCountryAg V2CorePersonAdditionalAddressCountry = "ag"
	V2CorePersonAdditionalAddressCountryAi V2CorePersonAdditionalAddressCountry = "ai"
	V2CorePersonAdditionalAddressCountryAl V2CorePersonAdditionalAddressCountry = "al"
	V2CorePersonAdditionalAddressCountryAm V2CorePersonAdditionalAddressCountry = "am"
	V2CorePersonAdditionalAddressCountryAo V2CorePersonAdditionalAddressCountry = "ao"
	V2CorePersonAdditionalAddressCountryAq V2CorePersonAdditionalAddressCountry = "aq"
	V2CorePersonAdditionalAddressCountryAR V2CorePersonAdditionalAddressCountry = "ar"
	V2CorePersonAdditionalAddressCountryAs V2CorePersonAdditionalAddressCountry = "as"
	V2CorePersonAdditionalAddressCountryAt V2CorePersonAdditionalAddressCountry = "at"
	V2CorePersonAdditionalAddressCountryAu V2CorePersonAdditionalAddressCountry = "au"
	V2CorePersonAdditionalAddressCountryAw V2CorePersonAdditionalAddressCountry = "aw"
	V2CorePersonAdditionalAddressCountryAx V2CorePersonAdditionalAddressCountry = "ax"
	V2CorePersonAdditionalAddressCountryAz V2CorePersonAdditionalAddressCountry = "az"
	V2CorePersonAdditionalAddressCountryBa V2CorePersonAdditionalAddressCountry = "ba"
	V2CorePersonAdditionalAddressCountryBb V2CorePersonAdditionalAddressCountry = "bb"
	V2CorePersonAdditionalAddressCountryBd V2CorePersonAdditionalAddressCountry = "bd"
	V2CorePersonAdditionalAddressCountryBe V2CorePersonAdditionalAddressCountry = "be"
	V2CorePersonAdditionalAddressCountryBf V2CorePersonAdditionalAddressCountry = "bf"
	V2CorePersonAdditionalAddressCountryBG V2CorePersonAdditionalAddressCountry = "bg"
	V2CorePersonAdditionalAddressCountryBh V2CorePersonAdditionalAddressCountry = "bh"
	V2CorePersonAdditionalAddressCountryBi V2CorePersonAdditionalAddressCountry = "bi"
	V2CorePersonAdditionalAddressCountryBj V2CorePersonAdditionalAddressCountry = "bj"
	V2CorePersonAdditionalAddressCountryBl V2CorePersonAdditionalAddressCountry = "bl"
	V2CorePersonAdditionalAddressCountryBm V2CorePersonAdditionalAddressCountry = "bm"
	V2CorePersonAdditionalAddressCountryBn V2CorePersonAdditionalAddressCountry = "bn"
	V2CorePersonAdditionalAddressCountryBo V2CorePersonAdditionalAddressCountry = "bo"
	V2CorePersonAdditionalAddressCountryBq V2CorePersonAdditionalAddressCountry = "bq"
	V2CorePersonAdditionalAddressCountryBr V2CorePersonAdditionalAddressCountry = "br"
	V2CorePersonAdditionalAddressCountryBs V2CorePersonAdditionalAddressCountry = "bs"
	V2CorePersonAdditionalAddressCountryBt V2CorePersonAdditionalAddressCountry = "bt"
	V2CorePersonAdditionalAddressCountryBv V2CorePersonAdditionalAddressCountry = "bv"
	V2CorePersonAdditionalAddressCountryBw V2CorePersonAdditionalAddressCountry = "bw"
	V2CorePersonAdditionalAddressCountryBy V2CorePersonAdditionalAddressCountry = "by"
	V2CorePersonAdditionalAddressCountryBz V2CorePersonAdditionalAddressCountry = "bz"
	V2CorePersonAdditionalAddressCountryCa V2CorePersonAdditionalAddressCountry = "ca"
	V2CorePersonAdditionalAddressCountryCc V2CorePersonAdditionalAddressCountry = "cc"
	V2CorePersonAdditionalAddressCountryCd V2CorePersonAdditionalAddressCountry = "cd"
	V2CorePersonAdditionalAddressCountryCf V2CorePersonAdditionalAddressCountry = "cf"
	V2CorePersonAdditionalAddressCountryCg V2CorePersonAdditionalAddressCountry = "cg"
	V2CorePersonAdditionalAddressCountryCh V2CorePersonAdditionalAddressCountry = "ch"
	V2CorePersonAdditionalAddressCountryCi V2CorePersonAdditionalAddressCountry = "ci"
	V2CorePersonAdditionalAddressCountryCk V2CorePersonAdditionalAddressCountry = "ck"
	V2CorePersonAdditionalAddressCountryCl V2CorePersonAdditionalAddressCountry = "cl"
	V2CorePersonAdditionalAddressCountryCm V2CorePersonAdditionalAddressCountry = "cm"
	V2CorePersonAdditionalAddressCountryCn V2CorePersonAdditionalAddressCountry = "cn"
	V2CorePersonAdditionalAddressCountryCo V2CorePersonAdditionalAddressCountry = "co"
	V2CorePersonAdditionalAddressCountryCr V2CorePersonAdditionalAddressCountry = "cr"
	V2CorePersonAdditionalAddressCountryCu V2CorePersonAdditionalAddressCountry = "cu"
	V2CorePersonAdditionalAddressCountryCv V2CorePersonAdditionalAddressCountry = "cv"
	V2CorePersonAdditionalAddressCountryCw V2CorePersonAdditionalAddressCountry = "cw"
	V2CorePersonAdditionalAddressCountryCx V2CorePersonAdditionalAddressCountry = "cx"
	V2CorePersonAdditionalAddressCountryCy V2CorePersonAdditionalAddressCountry = "cy"
	V2CorePersonAdditionalAddressCountryCz V2CorePersonAdditionalAddressCountry = "cz"
	V2CorePersonAdditionalAddressCountryDE V2CorePersonAdditionalAddressCountry = "de"
	V2CorePersonAdditionalAddressCountryDj V2CorePersonAdditionalAddressCountry = "dj"
	V2CorePersonAdditionalAddressCountryDk V2CorePersonAdditionalAddressCountry = "dk"
	V2CorePersonAdditionalAddressCountryDm V2CorePersonAdditionalAddressCountry = "dm"
	V2CorePersonAdditionalAddressCountryDo V2CorePersonAdditionalAddressCountry = "do"
	V2CorePersonAdditionalAddressCountryDz V2CorePersonAdditionalAddressCountry = "dz"
	V2CorePersonAdditionalAddressCountryEc V2CorePersonAdditionalAddressCountry = "ec"
	V2CorePersonAdditionalAddressCountryEe V2CorePersonAdditionalAddressCountry = "ee"
	V2CorePersonAdditionalAddressCountryEg V2CorePersonAdditionalAddressCountry = "eg"
	V2CorePersonAdditionalAddressCountryEh V2CorePersonAdditionalAddressCountry = "eh"
	V2CorePersonAdditionalAddressCountryEr V2CorePersonAdditionalAddressCountry = "er"
	V2CorePersonAdditionalAddressCountryES V2CorePersonAdditionalAddressCountry = "es"
	V2CorePersonAdditionalAddressCountryET V2CorePersonAdditionalAddressCountry = "et"
	V2CorePersonAdditionalAddressCountryFI V2CorePersonAdditionalAddressCountry = "fi"
	V2CorePersonAdditionalAddressCountryFj V2CorePersonAdditionalAddressCountry = "fj"
	V2CorePersonAdditionalAddressCountryFk V2CorePersonAdditionalAddressCountry = "fk"
	V2CorePersonAdditionalAddressCountryFm V2CorePersonAdditionalAddressCountry = "fm"
	V2CorePersonAdditionalAddressCountryFo V2CorePersonAdditionalAddressCountry = "fo"
	V2CorePersonAdditionalAddressCountryFR V2CorePersonAdditionalAddressCountry = "fr"
	V2CorePersonAdditionalAddressCountryGa V2CorePersonAdditionalAddressCountry = "ga"
	V2CorePersonAdditionalAddressCountryGB V2CorePersonAdditionalAddressCountry = "gb"
	V2CorePersonAdditionalAddressCountryGd V2CorePersonAdditionalAddressCountry = "gd"
	V2CorePersonAdditionalAddressCountryGe V2CorePersonAdditionalAddressCountry = "ge"
	V2CorePersonAdditionalAddressCountryGf V2CorePersonAdditionalAddressCountry = "gf"
	V2CorePersonAdditionalAddressCountryGg V2CorePersonAdditionalAddressCountry = "gg"
	V2CorePersonAdditionalAddressCountryGh V2CorePersonAdditionalAddressCountry = "gh"
	V2CorePersonAdditionalAddressCountryGi V2CorePersonAdditionalAddressCountry = "gi"
	V2CorePersonAdditionalAddressCountryGl V2CorePersonAdditionalAddressCountry = "gl"
	V2CorePersonAdditionalAddressCountryGm V2CorePersonAdditionalAddressCountry = "gm"
	V2CorePersonAdditionalAddressCountryGn V2CorePersonAdditionalAddressCountry = "gn"
	V2CorePersonAdditionalAddressCountryGp V2CorePersonAdditionalAddressCountry = "gp"
	V2CorePersonAdditionalAddressCountryGq V2CorePersonAdditionalAddressCountry = "gq"
	V2CorePersonAdditionalAddressCountryGr V2CorePersonAdditionalAddressCountry = "gr"
	V2CorePersonAdditionalAddressCountryGs V2CorePersonAdditionalAddressCountry = "gs"
	V2CorePersonAdditionalAddressCountryGt V2CorePersonAdditionalAddressCountry = "gt"
	V2CorePersonAdditionalAddressCountryGu V2CorePersonAdditionalAddressCountry = "gu"
	V2CorePersonAdditionalAddressCountryGw V2CorePersonAdditionalAddressCountry = "gw"
	V2CorePersonAdditionalAddressCountryGy V2CorePersonAdditionalAddressCountry = "gy"
	V2CorePersonAdditionalAddressCountryHk V2CorePersonAdditionalAddressCountry = "hk"
	V2CorePersonAdditionalAddressCountryHm V2CorePersonAdditionalAddressCountry = "hm"
	V2CorePersonAdditionalAddressCountryHn V2CorePersonAdditionalAddressCountry = "hn"
	V2CorePersonAdditionalAddressCountryHR V2CorePersonAdditionalAddressCountry = "hr"
	V2CorePersonAdditionalAddressCountryHt V2CorePersonAdditionalAddressCountry = "ht"
	V2CorePersonAdditionalAddressCountryHU V2CorePersonAdditionalAddressCountry = "hu"
	V2CorePersonAdditionalAddressCountryID V2CorePersonAdditionalAddressCountry = "id"
	V2CorePersonAdditionalAddressCountryIe V2CorePersonAdditionalAddressCountry = "ie"
	V2CorePersonAdditionalAddressCountryIl V2CorePersonAdditionalAddressCountry = "il"
	V2CorePersonAdditionalAddressCountryIm V2CorePersonAdditionalAddressCountry = "im"
	V2CorePersonAdditionalAddressCountryIn V2CorePersonAdditionalAddressCountry = "in"
	V2CorePersonAdditionalAddressCountryIo V2CorePersonAdditionalAddressCountry = "io"
	V2CorePersonAdditionalAddressCountryIq V2CorePersonAdditionalAddressCountry = "iq"
	V2CorePersonAdditionalAddressCountryIr V2CorePersonAdditionalAddressCountry = "ir"
	V2CorePersonAdditionalAddressCountryIs V2CorePersonAdditionalAddressCountry = "is"
	V2CorePersonAdditionalAddressCountryIT V2CorePersonAdditionalAddressCountry = "it"
	V2CorePersonAdditionalAddressCountryJe V2CorePersonAdditionalAddressCountry = "je"
	V2CorePersonAdditionalAddressCountryJm V2CorePersonAdditionalAddressCountry = "jm"
	V2CorePersonAdditionalAddressCountryJo V2CorePersonAdditionalAddressCountry = "jo"
	V2CorePersonAdditionalAddressCountryJP V2CorePersonAdditionalAddressCountry = "jp"
	V2CorePersonAdditionalAddressCountryKe V2CorePersonAdditionalAddressCountry = "ke"
	V2CorePersonAdditionalAddressCountryKg V2CorePersonAdditionalAddressCountry = "kg"
	V2CorePersonAdditionalAddressCountryKh V2CorePersonAdditionalAddressCountry = "kh"
	V2CorePersonAdditionalAddressCountryKi V2CorePersonAdditionalAddressCountry = "ki"
	V2CorePersonAdditionalAddressCountryKm V2CorePersonAdditionalAddressCountry = "km"
	V2CorePersonAdditionalAddressCountryKn V2CorePersonAdditionalAddressCountry = "kn"
	V2CorePersonAdditionalAddressCountryKp V2CorePersonAdditionalAddressCountry = "kp"
	V2CorePersonAdditionalAddressCountryKr V2CorePersonAdditionalAddressCountry = "kr"
	V2CorePersonAdditionalAddressCountryKw V2CorePersonAdditionalAddressCountry = "kw"
	V2CorePersonAdditionalAddressCountryKy V2CorePersonAdditionalAddressCountry = "ky"
	V2CorePersonAdditionalAddressCountryKz V2CorePersonAdditionalAddressCountry = "kz"
	V2CorePersonAdditionalAddressCountryLa V2CorePersonAdditionalAddressCountry = "la"
	V2CorePersonAdditionalAddressCountryLb V2CorePersonAdditionalAddressCountry = "lb"
	V2CorePersonAdditionalAddressCountryLc V2CorePersonAdditionalAddressCountry = "lc"
	V2CorePersonAdditionalAddressCountryLi V2CorePersonAdditionalAddressCountry = "li"
	V2CorePersonAdditionalAddressCountryLk V2CorePersonAdditionalAddressCountry = "lk"
	V2CorePersonAdditionalAddressCountryLr V2CorePersonAdditionalAddressCountry = "lr"
	V2CorePersonAdditionalAddressCountryLs V2CorePersonAdditionalAddressCountry = "ls"
	V2CorePersonAdditionalAddressCountryLT V2CorePersonAdditionalAddressCountry = "lt"
	V2CorePersonAdditionalAddressCountryLu V2CorePersonAdditionalAddressCountry = "lu"
	V2CorePersonAdditionalAddressCountryLV V2CorePersonAdditionalAddressCountry = "lv"
	V2CorePersonAdditionalAddressCountryLy V2CorePersonAdditionalAddressCountry = "ly"
	V2CorePersonAdditionalAddressCountryMa V2CorePersonAdditionalAddressCountry = "ma"
	V2CorePersonAdditionalAddressCountryMc V2CorePersonAdditionalAddressCountry = "mc"
	V2CorePersonAdditionalAddressCountryMd V2CorePersonAdditionalAddressCountry = "md"
	V2CorePersonAdditionalAddressCountryMe V2CorePersonAdditionalAddressCountry = "me"
	V2CorePersonAdditionalAddressCountryMf V2CorePersonAdditionalAddressCountry = "mf"
	V2CorePersonAdditionalAddressCountryMg V2CorePersonAdditionalAddressCountry = "mg"
	V2CorePersonAdditionalAddressCountryMh V2CorePersonAdditionalAddressCountry = "mh"
	V2CorePersonAdditionalAddressCountryMk V2CorePersonAdditionalAddressCountry = "mk"
	V2CorePersonAdditionalAddressCountryMl V2CorePersonAdditionalAddressCountry = "ml"
	V2CorePersonAdditionalAddressCountryMm V2CorePersonAdditionalAddressCountry = "mm"
	V2CorePersonAdditionalAddressCountryMn V2CorePersonAdditionalAddressCountry = "mn"
	V2CorePersonAdditionalAddressCountryMo V2CorePersonAdditionalAddressCountry = "mo"
	V2CorePersonAdditionalAddressCountryMp V2CorePersonAdditionalAddressCountry = "mp"
	V2CorePersonAdditionalAddressCountryMq V2CorePersonAdditionalAddressCountry = "mq"
	V2CorePersonAdditionalAddressCountryMr V2CorePersonAdditionalAddressCountry = "mr"
	V2CorePersonAdditionalAddressCountryMS V2CorePersonAdditionalAddressCountry = "ms"
	V2CorePersonAdditionalAddressCountryMT V2CorePersonAdditionalAddressCountry = "mt"
	V2CorePersonAdditionalAddressCountryMu V2CorePersonAdditionalAddressCountry = "mu"
	V2CorePersonAdditionalAddressCountryMv V2CorePersonAdditionalAddressCountry = "mv"
	V2CorePersonAdditionalAddressCountryMw V2CorePersonAdditionalAddressCountry = "mw"
	V2CorePersonAdditionalAddressCountryMX V2CorePersonAdditionalAddressCountry = "mx"
	V2CorePersonAdditionalAddressCountryMy V2CorePersonAdditionalAddressCountry = "my"
	V2CorePersonAdditionalAddressCountryMz V2CorePersonAdditionalAddressCountry = "mz"
	V2CorePersonAdditionalAddressCountryNa V2CorePersonAdditionalAddressCountry = "na"
	V2CorePersonAdditionalAddressCountryNc V2CorePersonAdditionalAddressCountry = "nc"
	V2CorePersonAdditionalAddressCountryNe V2CorePersonAdditionalAddressCountry = "ne"
	V2CorePersonAdditionalAddressCountryNf V2CorePersonAdditionalAddressCountry = "nf"
	V2CorePersonAdditionalAddressCountryNg V2CorePersonAdditionalAddressCountry = "ng"
	V2CorePersonAdditionalAddressCountryNi V2CorePersonAdditionalAddressCountry = "ni"
	V2CorePersonAdditionalAddressCountryNL V2CorePersonAdditionalAddressCountry = "nl"
	V2CorePersonAdditionalAddressCountryNo V2CorePersonAdditionalAddressCountry = "no"
	V2CorePersonAdditionalAddressCountryNp V2CorePersonAdditionalAddressCountry = "np"
	V2CorePersonAdditionalAddressCountryNr V2CorePersonAdditionalAddressCountry = "nr"
	V2CorePersonAdditionalAddressCountryNu V2CorePersonAdditionalAddressCountry = "nu"
	V2CorePersonAdditionalAddressCountryNz V2CorePersonAdditionalAddressCountry = "nz"
	V2CorePersonAdditionalAddressCountryOm V2CorePersonAdditionalAddressCountry = "om"
	V2CorePersonAdditionalAddressCountryPa V2CorePersonAdditionalAddressCountry = "pa"
	V2CorePersonAdditionalAddressCountryPe V2CorePersonAdditionalAddressCountry = "pe"
	V2CorePersonAdditionalAddressCountryPf V2CorePersonAdditionalAddressCountry = "pf"
	V2CorePersonAdditionalAddressCountryPg V2CorePersonAdditionalAddressCountry = "pg"
	V2CorePersonAdditionalAddressCountryPh V2CorePersonAdditionalAddressCountry = "ph"
	V2CorePersonAdditionalAddressCountryPk V2CorePersonAdditionalAddressCountry = "pk"
	V2CorePersonAdditionalAddressCountryPL V2CorePersonAdditionalAddressCountry = "pl"
	V2CorePersonAdditionalAddressCountryPm V2CorePersonAdditionalAddressCountry = "pm"
	V2CorePersonAdditionalAddressCountryPn V2CorePersonAdditionalAddressCountry = "pn"
	V2CorePersonAdditionalAddressCountryPr V2CorePersonAdditionalAddressCountry = "pr"
	V2CorePersonAdditionalAddressCountryPs V2CorePersonAdditionalAddressCountry = "ps"
	V2CorePersonAdditionalAddressCountryPT V2CorePersonAdditionalAddressCountry = "pt"
	V2CorePersonAdditionalAddressCountryPw V2CorePersonAdditionalAddressCountry = "pw"
	V2CorePersonAdditionalAddressCountryPy V2CorePersonAdditionalAddressCountry = "py"
	V2CorePersonAdditionalAddressCountryQa V2CorePersonAdditionalAddressCountry = "qa"
	V2CorePersonAdditionalAddressCountryQz V2CorePersonAdditionalAddressCountry = "qz"
	V2CorePersonAdditionalAddressCountryRe V2CorePersonAdditionalAddressCountry = "re"
	V2CorePersonAdditionalAddressCountryRO V2CorePersonAdditionalAddressCountry = "ro"
	V2CorePersonAdditionalAddressCountryRs V2CorePersonAdditionalAddressCountry = "rs"
	V2CorePersonAdditionalAddressCountryRU V2CorePersonAdditionalAddressCountry = "ru"
	V2CorePersonAdditionalAddressCountryRw V2CorePersonAdditionalAddressCountry = "rw"
	V2CorePersonAdditionalAddressCountrySa V2CorePersonAdditionalAddressCountry = "sa"
	V2CorePersonAdditionalAddressCountrySb V2CorePersonAdditionalAddressCountry = "sb"
	V2CorePersonAdditionalAddressCountrySc V2CorePersonAdditionalAddressCountry = "sc"
	V2CorePersonAdditionalAddressCountrySd V2CorePersonAdditionalAddressCountry = "sd"
	V2CorePersonAdditionalAddressCountrySe V2CorePersonAdditionalAddressCountry = "se"
	V2CorePersonAdditionalAddressCountrySg V2CorePersonAdditionalAddressCountry = "sg"
	V2CorePersonAdditionalAddressCountrySh V2CorePersonAdditionalAddressCountry = "sh"
	V2CorePersonAdditionalAddressCountrySi V2CorePersonAdditionalAddressCountry = "si"
	V2CorePersonAdditionalAddressCountrySj V2CorePersonAdditionalAddressCountry = "sj"
	V2CorePersonAdditionalAddressCountrySK V2CorePersonAdditionalAddressCountry = "sk"
	V2CorePersonAdditionalAddressCountrySL V2CorePersonAdditionalAddressCountry = "sl"
	V2CorePersonAdditionalAddressCountrySm V2CorePersonAdditionalAddressCountry = "sm"
	V2CorePersonAdditionalAddressCountrySn V2CorePersonAdditionalAddressCountry = "sn"
	V2CorePersonAdditionalAddressCountrySo V2CorePersonAdditionalAddressCountry = "so"
	V2CorePersonAdditionalAddressCountrySr V2CorePersonAdditionalAddressCountry = "sr"
	V2CorePersonAdditionalAddressCountrySs V2CorePersonAdditionalAddressCountry = "ss"
	V2CorePersonAdditionalAddressCountrySt V2CorePersonAdditionalAddressCountry = "st"
	V2CorePersonAdditionalAddressCountrySV V2CorePersonAdditionalAddressCountry = "sv"
	V2CorePersonAdditionalAddressCountrySx V2CorePersonAdditionalAddressCountry = "sx"
	V2CorePersonAdditionalAddressCountrySy V2CorePersonAdditionalAddressCountry = "sy"
	V2CorePersonAdditionalAddressCountrySz V2CorePersonAdditionalAddressCountry = "sz"
	V2CorePersonAdditionalAddressCountryTc V2CorePersonAdditionalAddressCountry = "tc"
	V2CorePersonAdditionalAddressCountryTd V2CorePersonAdditionalAddressCountry = "td"
	V2CorePersonAdditionalAddressCountryTf V2CorePersonAdditionalAddressCountry = "tf"
	V2CorePersonAdditionalAddressCountryTg V2CorePersonAdditionalAddressCountry = "tg"
	V2CorePersonAdditionalAddressCountryTH V2CorePersonAdditionalAddressCountry = "th"
	V2CorePersonAdditionalAddressCountryTj V2CorePersonAdditionalAddressCountry = "tj"
	V2CorePersonAdditionalAddressCountryTk V2CorePersonAdditionalAddressCountry = "tk"
	V2CorePersonAdditionalAddressCountryTl V2CorePersonAdditionalAddressCountry = "tl"
	V2CorePersonAdditionalAddressCountryTm V2CorePersonAdditionalAddressCountry = "tm"
	V2CorePersonAdditionalAddressCountryTn V2CorePersonAdditionalAddressCountry = "tn"
	V2CorePersonAdditionalAddressCountryTo V2CorePersonAdditionalAddressCountry = "to"
	V2CorePersonAdditionalAddressCountryTR V2CorePersonAdditionalAddressCountry = "tr"
	V2CorePersonAdditionalAddressCountryTt V2CorePersonAdditionalAddressCountry = "tt"
	V2CorePersonAdditionalAddressCountryTv V2CorePersonAdditionalAddressCountry = "tv"
	V2CorePersonAdditionalAddressCountryTw V2CorePersonAdditionalAddressCountry = "tw"
	V2CorePersonAdditionalAddressCountryTz V2CorePersonAdditionalAddressCountry = "tz"
	V2CorePersonAdditionalAddressCountryUa V2CorePersonAdditionalAddressCountry = "ua"
	V2CorePersonAdditionalAddressCountryUg V2CorePersonAdditionalAddressCountry = "ug"
	V2CorePersonAdditionalAddressCountryUm V2CorePersonAdditionalAddressCountry = "um"
	V2CorePersonAdditionalAddressCountryUS V2CorePersonAdditionalAddressCountry = "us"
	V2CorePersonAdditionalAddressCountryUy V2CorePersonAdditionalAddressCountry = "uy"
	V2CorePersonAdditionalAddressCountryUz V2CorePersonAdditionalAddressCountry = "uz"
	V2CorePersonAdditionalAddressCountryVa V2CorePersonAdditionalAddressCountry = "va"
	V2CorePersonAdditionalAddressCountryVc V2CorePersonAdditionalAddressCountry = "vc"
	V2CorePersonAdditionalAddressCountryVe V2CorePersonAdditionalAddressCountry = "ve"
	V2CorePersonAdditionalAddressCountryVg V2CorePersonAdditionalAddressCountry = "vg"
	V2CorePersonAdditionalAddressCountryVI V2CorePersonAdditionalAddressCountry = "vi"
	V2CorePersonAdditionalAddressCountryVn V2CorePersonAdditionalAddressCountry = "vn"
	V2CorePersonAdditionalAddressCountryVu V2CorePersonAdditionalAddressCountry = "vu"
	V2CorePersonAdditionalAddressCountryWf V2CorePersonAdditionalAddressCountry = "wf"
	V2CorePersonAdditionalAddressCountryWs V2CorePersonAdditionalAddressCountry = "ws"
	V2CorePersonAdditionalAddressCountryXx V2CorePersonAdditionalAddressCountry = "xx"
	V2CorePersonAdditionalAddressCountryYe V2CorePersonAdditionalAddressCountry = "ye"
	V2CorePersonAdditionalAddressCountryYt V2CorePersonAdditionalAddressCountry = "yt"
	V2CorePersonAdditionalAddressCountryZa V2CorePersonAdditionalAddressCountry = "za"
	V2CorePersonAdditionalAddressCountryZm V2CorePersonAdditionalAddressCountry = "zm"
	V2CorePersonAdditionalAddressCountryZw V2CorePersonAdditionalAddressCountry = "zw"
)

// Purpose of additional address.
type V2CorePersonAdditionalAddressPurpose string

// List of values that V2CorePersonAdditionalAddressPurpose can take
const (
	V2CorePersonAdditionalAddressPurposeRegistered V2CorePersonAdditionalAddressPurpose = "registered"
)

// The purpose or type of the additional name.
type V2CorePersonAdditionalNamePurpose string

// List of values that V2CorePersonAdditionalNamePurpose can take
const (
	V2CorePersonAdditionalNamePurposeAlias  V2CorePersonAdditionalNamePurpose = "alias"
	V2CorePersonAdditionalNamePurposeMaiden V2CorePersonAdditionalNamePurpose = "maiden"
)

// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CorePersonAddressCountry string

// List of values that V2CorePersonAddressCountry can take
const (
	V2CorePersonAddressCountryAd V2CorePersonAddressCountry = "ad"
	V2CorePersonAddressCountryAe V2CorePersonAddressCountry = "ae"
	V2CorePersonAddressCountryAf V2CorePersonAddressCountry = "af"
	V2CorePersonAddressCountryAg V2CorePersonAddressCountry = "ag"
	V2CorePersonAddressCountryAi V2CorePersonAddressCountry = "ai"
	V2CorePersonAddressCountryAl V2CorePersonAddressCountry = "al"
	V2CorePersonAddressCountryAm V2CorePersonAddressCountry = "am"
	V2CorePersonAddressCountryAo V2CorePersonAddressCountry = "ao"
	V2CorePersonAddressCountryAq V2CorePersonAddressCountry = "aq"
	V2CorePersonAddressCountryAR V2CorePersonAddressCountry = "ar"
	V2CorePersonAddressCountryAs V2CorePersonAddressCountry = "as"
	V2CorePersonAddressCountryAt V2CorePersonAddressCountry = "at"
	V2CorePersonAddressCountryAu V2CorePersonAddressCountry = "au"
	V2CorePersonAddressCountryAw V2CorePersonAddressCountry = "aw"
	V2CorePersonAddressCountryAx V2CorePersonAddressCountry = "ax"
	V2CorePersonAddressCountryAz V2CorePersonAddressCountry = "az"
	V2CorePersonAddressCountryBa V2CorePersonAddressCountry = "ba"
	V2CorePersonAddressCountryBb V2CorePersonAddressCountry = "bb"
	V2CorePersonAddressCountryBd V2CorePersonAddressCountry = "bd"
	V2CorePersonAddressCountryBe V2CorePersonAddressCountry = "be"
	V2CorePersonAddressCountryBf V2CorePersonAddressCountry = "bf"
	V2CorePersonAddressCountryBG V2CorePersonAddressCountry = "bg"
	V2CorePersonAddressCountryBh V2CorePersonAddressCountry = "bh"
	V2CorePersonAddressCountryBi V2CorePersonAddressCountry = "bi"
	V2CorePersonAddressCountryBj V2CorePersonAddressCountry = "bj"
	V2CorePersonAddressCountryBl V2CorePersonAddressCountry = "bl"
	V2CorePersonAddressCountryBm V2CorePersonAddressCountry = "bm"
	V2CorePersonAddressCountryBn V2CorePersonAddressCountry = "bn"
	V2CorePersonAddressCountryBo V2CorePersonAddressCountry = "bo"
	V2CorePersonAddressCountryBq V2CorePersonAddressCountry = "bq"
	V2CorePersonAddressCountryBr V2CorePersonAddressCountry = "br"
	V2CorePersonAddressCountryBs V2CorePersonAddressCountry = "bs"
	V2CorePersonAddressCountryBt V2CorePersonAddressCountry = "bt"
	V2CorePersonAddressCountryBv V2CorePersonAddressCountry = "bv"
	V2CorePersonAddressCountryBw V2CorePersonAddressCountry = "bw"
	V2CorePersonAddressCountryBy V2CorePersonAddressCountry = "by"
	V2CorePersonAddressCountryBz V2CorePersonAddressCountry = "bz"
	V2CorePersonAddressCountryCa V2CorePersonAddressCountry = "ca"
	V2CorePersonAddressCountryCc V2CorePersonAddressCountry = "cc"
	V2CorePersonAddressCountryCd V2CorePersonAddressCountry = "cd"
	V2CorePersonAddressCountryCf V2CorePersonAddressCountry = "cf"
	V2CorePersonAddressCountryCg V2CorePersonAddressCountry = "cg"
	V2CorePersonAddressCountryCh V2CorePersonAddressCountry = "ch"
	V2CorePersonAddressCountryCi V2CorePersonAddressCountry = "ci"
	V2CorePersonAddressCountryCk V2CorePersonAddressCountry = "ck"
	V2CorePersonAddressCountryCl V2CorePersonAddressCountry = "cl"
	V2CorePersonAddressCountryCm V2CorePersonAddressCountry = "cm"
	V2CorePersonAddressCountryCn V2CorePersonAddressCountry = "cn"
	V2CorePersonAddressCountryCo V2CorePersonAddressCountry = "co"
	V2CorePersonAddressCountryCr V2CorePersonAddressCountry = "cr"
	V2CorePersonAddressCountryCu V2CorePersonAddressCountry = "cu"
	V2CorePersonAddressCountryCv V2CorePersonAddressCountry = "cv"
	V2CorePersonAddressCountryCw V2CorePersonAddressCountry = "cw"
	V2CorePersonAddressCountryCx V2CorePersonAddressCountry = "cx"
	V2CorePersonAddressCountryCy V2CorePersonAddressCountry = "cy"
	V2CorePersonAddressCountryCz V2CorePersonAddressCountry = "cz"
	V2CorePersonAddressCountryDE V2CorePersonAddressCountry = "de"
	V2CorePersonAddressCountryDj V2CorePersonAddressCountry = "dj"
	V2CorePersonAddressCountryDk V2CorePersonAddressCountry = "dk"
	V2CorePersonAddressCountryDm V2CorePersonAddressCountry = "dm"
	V2CorePersonAddressCountryDo V2CorePersonAddressCountry = "do"
	V2CorePersonAddressCountryDz V2CorePersonAddressCountry = "dz"
	V2CorePersonAddressCountryEc V2CorePersonAddressCountry = "ec"
	V2CorePersonAddressCountryEe V2CorePersonAddressCountry = "ee"
	V2CorePersonAddressCountryEg V2CorePersonAddressCountry = "eg"
	V2CorePersonAddressCountryEh V2CorePersonAddressCountry = "eh"
	V2CorePersonAddressCountryEr V2CorePersonAddressCountry = "er"
	V2CorePersonAddressCountryES V2CorePersonAddressCountry = "es"
	V2CorePersonAddressCountryET V2CorePersonAddressCountry = "et"
	V2CorePersonAddressCountryFI V2CorePersonAddressCountry = "fi"
	V2CorePersonAddressCountryFj V2CorePersonAddressCountry = "fj"
	V2CorePersonAddressCountryFk V2CorePersonAddressCountry = "fk"
	V2CorePersonAddressCountryFm V2CorePersonAddressCountry = "fm"
	V2CorePersonAddressCountryFo V2CorePersonAddressCountry = "fo"
	V2CorePersonAddressCountryFR V2CorePersonAddressCountry = "fr"
	V2CorePersonAddressCountryGa V2CorePersonAddressCountry = "ga"
	V2CorePersonAddressCountryGB V2CorePersonAddressCountry = "gb"
	V2CorePersonAddressCountryGd V2CorePersonAddressCountry = "gd"
	V2CorePersonAddressCountryGe V2CorePersonAddressCountry = "ge"
	V2CorePersonAddressCountryGf V2CorePersonAddressCountry = "gf"
	V2CorePersonAddressCountryGg V2CorePersonAddressCountry = "gg"
	V2CorePersonAddressCountryGh V2CorePersonAddressCountry = "gh"
	V2CorePersonAddressCountryGi V2CorePersonAddressCountry = "gi"
	V2CorePersonAddressCountryGl V2CorePersonAddressCountry = "gl"
	V2CorePersonAddressCountryGm V2CorePersonAddressCountry = "gm"
	V2CorePersonAddressCountryGn V2CorePersonAddressCountry = "gn"
	V2CorePersonAddressCountryGp V2CorePersonAddressCountry = "gp"
	V2CorePersonAddressCountryGq V2CorePersonAddressCountry = "gq"
	V2CorePersonAddressCountryGr V2CorePersonAddressCountry = "gr"
	V2CorePersonAddressCountryGs V2CorePersonAddressCountry = "gs"
	V2CorePersonAddressCountryGt V2CorePersonAddressCountry = "gt"
	V2CorePersonAddressCountryGu V2CorePersonAddressCountry = "gu"
	V2CorePersonAddressCountryGw V2CorePersonAddressCountry = "gw"
	V2CorePersonAddressCountryGy V2CorePersonAddressCountry = "gy"
	V2CorePersonAddressCountryHk V2CorePersonAddressCountry = "hk"
	V2CorePersonAddressCountryHm V2CorePersonAddressCountry = "hm"
	V2CorePersonAddressCountryHn V2CorePersonAddressCountry = "hn"
	V2CorePersonAddressCountryHR V2CorePersonAddressCountry = "hr"
	V2CorePersonAddressCountryHt V2CorePersonAddressCountry = "ht"
	V2CorePersonAddressCountryHU V2CorePersonAddressCountry = "hu"
	V2CorePersonAddressCountryID V2CorePersonAddressCountry = "id"
	V2CorePersonAddressCountryIe V2CorePersonAddressCountry = "ie"
	V2CorePersonAddressCountryIl V2CorePersonAddressCountry = "il"
	V2CorePersonAddressCountryIm V2CorePersonAddressCountry = "im"
	V2CorePersonAddressCountryIn V2CorePersonAddressCountry = "in"
	V2CorePersonAddressCountryIo V2CorePersonAddressCountry = "io"
	V2CorePersonAddressCountryIq V2CorePersonAddressCountry = "iq"
	V2CorePersonAddressCountryIr V2CorePersonAddressCountry = "ir"
	V2CorePersonAddressCountryIs V2CorePersonAddressCountry = "is"
	V2CorePersonAddressCountryIT V2CorePersonAddressCountry = "it"
	V2CorePersonAddressCountryJe V2CorePersonAddressCountry = "je"
	V2CorePersonAddressCountryJm V2CorePersonAddressCountry = "jm"
	V2CorePersonAddressCountryJo V2CorePersonAddressCountry = "jo"
	V2CorePersonAddressCountryJP V2CorePersonAddressCountry = "jp"
	V2CorePersonAddressCountryKe V2CorePersonAddressCountry = "ke"
	V2CorePersonAddressCountryKg V2CorePersonAddressCountry = "kg"
	V2CorePersonAddressCountryKh V2CorePersonAddressCountry = "kh"
	V2CorePersonAddressCountryKi V2CorePersonAddressCountry = "ki"
	V2CorePersonAddressCountryKm V2CorePersonAddressCountry = "km"
	V2CorePersonAddressCountryKn V2CorePersonAddressCountry = "kn"
	V2CorePersonAddressCountryKp V2CorePersonAddressCountry = "kp"
	V2CorePersonAddressCountryKr V2CorePersonAddressCountry = "kr"
	V2CorePersonAddressCountryKw V2CorePersonAddressCountry = "kw"
	V2CorePersonAddressCountryKy V2CorePersonAddressCountry = "ky"
	V2CorePersonAddressCountryKz V2CorePersonAddressCountry = "kz"
	V2CorePersonAddressCountryLa V2CorePersonAddressCountry = "la"
	V2CorePersonAddressCountryLb V2CorePersonAddressCountry = "lb"
	V2CorePersonAddressCountryLc V2CorePersonAddressCountry = "lc"
	V2CorePersonAddressCountryLi V2CorePersonAddressCountry = "li"
	V2CorePersonAddressCountryLk V2CorePersonAddressCountry = "lk"
	V2CorePersonAddressCountryLr V2CorePersonAddressCountry = "lr"
	V2CorePersonAddressCountryLs V2CorePersonAddressCountry = "ls"
	V2CorePersonAddressCountryLT V2CorePersonAddressCountry = "lt"
	V2CorePersonAddressCountryLu V2CorePersonAddressCountry = "lu"
	V2CorePersonAddressCountryLV V2CorePersonAddressCountry = "lv"
	V2CorePersonAddressCountryLy V2CorePersonAddressCountry = "ly"
	V2CorePersonAddressCountryMa V2CorePersonAddressCountry = "ma"
	V2CorePersonAddressCountryMc V2CorePersonAddressCountry = "mc"
	V2CorePersonAddressCountryMd V2CorePersonAddressCountry = "md"
	V2CorePersonAddressCountryMe V2CorePersonAddressCountry = "me"
	V2CorePersonAddressCountryMf V2CorePersonAddressCountry = "mf"
	V2CorePersonAddressCountryMg V2CorePersonAddressCountry = "mg"
	V2CorePersonAddressCountryMh V2CorePersonAddressCountry = "mh"
	V2CorePersonAddressCountryMk V2CorePersonAddressCountry = "mk"
	V2CorePersonAddressCountryMl V2CorePersonAddressCountry = "ml"
	V2CorePersonAddressCountryMm V2CorePersonAddressCountry = "mm"
	V2CorePersonAddressCountryMn V2CorePersonAddressCountry = "mn"
	V2CorePersonAddressCountryMo V2CorePersonAddressCountry = "mo"
	V2CorePersonAddressCountryMp V2CorePersonAddressCountry = "mp"
	V2CorePersonAddressCountryMq V2CorePersonAddressCountry = "mq"
	V2CorePersonAddressCountryMr V2CorePersonAddressCountry = "mr"
	V2CorePersonAddressCountryMS V2CorePersonAddressCountry = "ms"
	V2CorePersonAddressCountryMT V2CorePersonAddressCountry = "mt"
	V2CorePersonAddressCountryMu V2CorePersonAddressCountry = "mu"
	V2CorePersonAddressCountryMv V2CorePersonAddressCountry = "mv"
	V2CorePersonAddressCountryMw V2CorePersonAddressCountry = "mw"
	V2CorePersonAddressCountryMX V2CorePersonAddressCountry = "mx"
	V2CorePersonAddressCountryMy V2CorePersonAddressCountry = "my"
	V2CorePersonAddressCountryMz V2CorePersonAddressCountry = "mz"
	V2CorePersonAddressCountryNa V2CorePersonAddressCountry = "na"
	V2CorePersonAddressCountryNc V2CorePersonAddressCountry = "nc"
	V2CorePersonAddressCountryNe V2CorePersonAddressCountry = "ne"
	V2CorePersonAddressCountryNf V2CorePersonAddressCountry = "nf"
	V2CorePersonAddressCountryNg V2CorePersonAddressCountry = "ng"
	V2CorePersonAddressCountryNi V2CorePersonAddressCountry = "ni"
	V2CorePersonAddressCountryNL V2CorePersonAddressCountry = "nl"
	V2CorePersonAddressCountryNo V2CorePersonAddressCountry = "no"
	V2CorePersonAddressCountryNp V2CorePersonAddressCountry = "np"
	V2CorePersonAddressCountryNr V2CorePersonAddressCountry = "nr"
	V2CorePersonAddressCountryNu V2CorePersonAddressCountry = "nu"
	V2CorePersonAddressCountryNz V2CorePersonAddressCountry = "nz"
	V2CorePersonAddressCountryOm V2CorePersonAddressCountry = "om"
	V2CorePersonAddressCountryPa V2CorePersonAddressCountry = "pa"
	V2CorePersonAddressCountryPe V2CorePersonAddressCountry = "pe"
	V2CorePersonAddressCountryPf V2CorePersonAddressCountry = "pf"
	V2CorePersonAddressCountryPg V2CorePersonAddressCountry = "pg"
	V2CorePersonAddressCountryPh V2CorePersonAddressCountry = "ph"
	V2CorePersonAddressCountryPk V2CorePersonAddressCountry = "pk"
	V2CorePersonAddressCountryPL V2CorePersonAddressCountry = "pl"
	V2CorePersonAddressCountryPm V2CorePersonAddressCountry = "pm"
	V2CorePersonAddressCountryPn V2CorePersonAddressCountry = "pn"
	V2CorePersonAddressCountryPr V2CorePersonAddressCountry = "pr"
	V2CorePersonAddressCountryPs V2CorePersonAddressCountry = "ps"
	V2CorePersonAddressCountryPT V2CorePersonAddressCountry = "pt"
	V2CorePersonAddressCountryPw V2CorePersonAddressCountry = "pw"
	V2CorePersonAddressCountryPy V2CorePersonAddressCountry = "py"
	V2CorePersonAddressCountryQa V2CorePersonAddressCountry = "qa"
	V2CorePersonAddressCountryQz V2CorePersonAddressCountry = "qz"
	V2CorePersonAddressCountryRe V2CorePersonAddressCountry = "re"
	V2CorePersonAddressCountryRO V2CorePersonAddressCountry = "ro"
	V2CorePersonAddressCountryRs V2CorePersonAddressCountry = "rs"
	V2CorePersonAddressCountryRU V2CorePersonAddressCountry = "ru"
	V2CorePersonAddressCountryRw V2CorePersonAddressCountry = "rw"
	V2CorePersonAddressCountrySa V2CorePersonAddressCountry = "sa"
	V2CorePersonAddressCountrySb V2CorePersonAddressCountry = "sb"
	V2CorePersonAddressCountrySc V2CorePersonAddressCountry = "sc"
	V2CorePersonAddressCountrySd V2CorePersonAddressCountry = "sd"
	V2CorePersonAddressCountrySe V2CorePersonAddressCountry = "se"
	V2CorePersonAddressCountrySg V2CorePersonAddressCountry = "sg"
	V2CorePersonAddressCountrySh V2CorePersonAddressCountry = "sh"
	V2CorePersonAddressCountrySi V2CorePersonAddressCountry = "si"
	V2CorePersonAddressCountrySj V2CorePersonAddressCountry = "sj"
	V2CorePersonAddressCountrySK V2CorePersonAddressCountry = "sk"
	V2CorePersonAddressCountrySL V2CorePersonAddressCountry = "sl"
	V2CorePersonAddressCountrySm V2CorePersonAddressCountry = "sm"
	V2CorePersonAddressCountrySn V2CorePersonAddressCountry = "sn"
	V2CorePersonAddressCountrySo V2CorePersonAddressCountry = "so"
	V2CorePersonAddressCountrySr V2CorePersonAddressCountry = "sr"
	V2CorePersonAddressCountrySs V2CorePersonAddressCountry = "ss"
	V2CorePersonAddressCountrySt V2CorePersonAddressCountry = "st"
	V2CorePersonAddressCountrySV V2CorePersonAddressCountry = "sv"
	V2CorePersonAddressCountrySx V2CorePersonAddressCountry = "sx"
	V2CorePersonAddressCountrySy V2CorePersonAddressCountry = "sy"
	V2CorePersonAddressCountrySz V2CorePersonAddressCountry = "sz"
	V2CorePersonAddressCountryTc V2CorePersonAddressCountry = "tc"
	V2CorePersonAddressCountryTd V2CorePersonAddressCountry = "td"
	V2CorePersonAddressCountryTf V2CorePersonAddressCountry = "tf"
	V2CorePersonAddressCountryTg V2CorePersonAddressCountry = "tg"
	V2CorePersonAddressCountryTH V2CorePersonAddressCountry = "th"
	V2CorePersonAddressCountryTj V2CorePersonAddressCountry = "tj"
	V2CorePersonAddressCountryTk V2CorePersonAddressCountry = "tk"
	V2CorePersonAddressCountryTl V2CorePersonAddressCountry = "tl"
	V2CorePersonAddressCountryTm V2CorePersonAddressCountry = "tm"
	V2CorePersonAddressCountryTn V2CorePersonAddressCountry = "tn"
	V2CorePersonAddressCountryTo V2CorePersonAddressCountry = "to"
	V2CorePersonAddressCountryTR V2CorePersonAddressCountry = "tr"
	V2CorePersonAddressCountryTt V2CorePersonAddressCountry = "tt"
	V2CorePersonAddressCountryTv V2CorePersonAddressCountry = "tv"
	V2CorePersonAddressCountryTw V2CorePersonAddressCountry = "tw"
	V2CorePersonAddressCountryTz V2CorePersonAddressCountry = "tz"
	V2CorePersonAddressCountryUa V2CorePersonAddressCountry = "ua"
	V2CorePersonAddressCountryUg V2CorePersonAddressCountry = "ug"
	V2CorePersonAddressCountryUm V2CorePersonAddressCountry = "um"
	V2CorePersonAddressCountryUS V2CorePersonAddressCountry = "us"
	V2CorePersonAddressCountryUy V2CorePersonAddressCountry = "uy"
	V2CorePersonAddressCountryUz V2CorePersonAddressCountry = "uz"
	V2CorePersonAddressCountryVa V2CorePersonAddressCountry = "va"
	V2CorePersonAddressCountryVc V2CorePersonAddressCountry = "vc"
	V2CorePersonAddressCountryVe V2CorePersonAddressCountry = "ve"
	V2CorePersonAddressCountryVg V2CorePersonAddressCountry = "vg"
	V2CorePersonAddressCountryVI V2CorePersonAddressCountry = "vi"
	V2CorePersonAddressCountryVn V2CorePersonAddressCountry = "vn"
	V2CorePersonAddressCountryVu V2CorePersonAddressCountry = "vu"
	V2CorePersonAddressCountryWf V2CorePersonAddressCountry = "wf"
	V2CorePersonAddressCountryWs V2CorePersonAddressCountry = "ws"
	V2CorePersonAddressCountryXx V2CorePersonAddressCountry = "xx"
	V2CorePersonAddressCountryYe V2CorePersonAddressCountry = "ye"
	V2CorePersonAddressCountryYt V2CorePersonAddressCountry = "yt"
	V2CorePersonAddressCountryZa V2CorePersonAddressCountry = "za"
	V2CorePersonAddressCountryZm V2CorePersonAddressCountry = "zm"
	V2CorePersonAddressCountryZw V2CorePersonAddressCountry = "zw"
)

// The format of the document. Currently supports `files` only.
type V2CorePersonDocumentsCompanyAuthorizationType string

// List of values that V2CorePersonDocumentsCompanyAuthorizationType can take
const (
	V2CorePersonDocumentsCompanyAuthorizationTypeFiles V2CorePersonDocumentsCompanyAuthorizationType = "files"
)

// The format of the document. Currently supports `files` only.
type V2CorePersonDocumentsPassportType string

// List of values that V2CorePersonDocumentsPassportType can take
const (
	V2CorePersonDocumentsPassportTypeFiles V2CorePersonDocumentsPassportType = "files"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CorePersonDocumentsPrimaryVerificationType string

// List of values that V2CorePersonDocumentsPrimaryVerificationType can take
const (
	V2CorePersonDocumentsPrimaryVerificationTypeFrontBack V2CorePersonDocumentsPrimaryVerificationType = "front_back"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CorePersonDocumentsSecondaryVerificationType string

// List of values that V2CorePersonDocumentsSecondaryVerificationType can take
const (
	V2CorePersonDocumentsSecondaryVerificationTypeFrontBack V2CorePersonDocumentsSecondaryVerificationType = "front_back"
)

// The format of the document. Currently supports `files` only.
type V2CorePersonDocumentsVisaType string

// List of values that V2CorePersonDocumentsVisaType can take
const (
	V2CorePersonDocumentsVisaTypeFiles V2CorePersonDocumentsVisaType = "files"
)

// The ID number type of an individual.
type V2CorePersonIDNumberType string

// List of values that V2CorePersonIDNumberType can take
const (
	V2CorePersonIDNumberTypeAeEid       V2CorePersonIDNumberType = "ae_eid"
	V2CorePersonIDNumberTypeAoNif       V2CorePersonIDNumberType = "ao_nif"
	V2CorePersonIDNumberTypeAzTin       V2CorePersonIDNumberType = "az_tin"
	V2CorePersonIDNumberTypeBdBrc       V2CorePersonIDNumberType = "bd_brc"
	V2CorePersonIDNumberTypeBdEtin      V2CorePersonIDNumberType = "bd_etin"
	V2CorePersonIDNumberTypeBdNid       V2CorePersonIDNumberType = "bd_nid"
	V2CorePersonIDNumberTypeBRCPF       V2CorePersonIDNumberType = "br_cpf"
	V2CorePersonIDNumberTypeCrCpf       V2CorePersonIDNumberType = "cr_cpf"
	V2CorePersonIDNumberTypeCrDimex     V2CorePersonIDNumberType = "cr_dimex"
	V2CorePersonIDNumberTypeCrNite      V2CorePersonIDNumberType = "cr_nite"
	V2CorePersonIDNumberTypeDEStn       V2CorePersonIDNumberType = "de_stn"
	V2CorePersonIDNumberTypeDORCN       V2CorePersonIDNumberType = "do_rcn"
	V2CorePersonIDNumberTypeGtNit       V2CorePersonIDNumberType = "gt_nit"
	V2CorePersonIDNumberTypeHkID        V2CorePersonIDNumberType = "hk_id"
	V2CorePersonIDNumberTypeKzIIN       V2CorePersonIDNumberType = "kz_iin"
	V2CorePersonIDNumberTypeMXRFC       V2CorePersonIDNumberType = "mx_rfc"
	V2CorePersonIDNumberTypeMyNric      V2CorePersonIDNumberType = "my_nric"
	V2CorePersonIDNumberTypeMzNuit      V2CorePersonIDNumberType = "mz_nuit"
	V2CorePersonIDNumberTypeNLBsn       V2CorePersonIDNumberType = "nl_bsn"
	V2CorePersonIDNumberTypePeDni       V2CorePersonIDNumberType = "pe_dni"
	V2CorePersonIDNumberTypePkCnic      V2CorePersonIDNumberType = "pk_cnic"
	V2CorePersonIDNumberTypePkSnic      V2CorePersonIDNumberType = "pk_snic"
	V2CorePersonIDNumberTypeSaTin       V2CorePersonIDNumberType = "sa_tin"
	V2CorePersonIDNumberTypeSgFin       V2CorePersonIDNumberType = "sg_fin"
	V2CorePersonIDNumberTypeSGNRIC      V2CorePersonIDNumberType = "sg_nric"
	V2CorePersonIDNumberTypeTHLc        V2CorePersonIDNumberType = "th_lc"
	V2CorePersonIDNumberTypeTHPIN       V2CorePersonIDNumberType = "th_pin"
	V2CorePersonIDNumberTypeUSItin      V2CorePersonIDNumberType = "us_itin"
	V2CorePersonIDNumberTypeUSItinLast4 V2CorePersonIDNumberType = "us_itin_last_4"
	V2CorePersonIDNumberTypeUSSSN       V2CorePersonIDNumberType = "us_ssn"
	V2CorePersonIDNumberTypeUSSSNLast4  V2CorePersonIDNumberType = "us_ssn_last_4"
)

// The person's gender (International regulations require either "male" or "female").
type V2CorePersonLegalGender string

// List of values that V2CorePersonLegalGender can take
const (
	V2CorePersonLegalGenderFemale V2CorePersonLegalGender = "female"
	V2CorePersonLegalGenderMale   V2CorePersonLegalGender = "male"
)

// The countries where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CorePersonNationality string

// List of values that V2CorePersonNationality can take
const (
	V2CorePersonNationalityAd V2CorePersonNationality = "ad"
	V2CorePersonNationalityAe V2CorePersonNationality = "ae"
	V2CorePersonNationalityAf V2CorePersonNationality = "af"
	V2CorePersonNationalityAg V2CorePersonNationality = "ag"
	V2CorePersonNationalityAi V2CorePersonNationality = "ai"
	V2CorePersonNationalityAl V2CorePersonNationality = "al"
	V2CorePersonNationalityAm V2CorePersonNationality = "am"
	V2CorePersonNationalityAo V2CorePersonNationality = "ao"
	V2CorePersonNationalityAq V2CorePersonNationality = "aq"
	V2CorePersonNationalityAR V2CorePersonNationality = "ar"
	V2CorePersonNationalityAs V2CorePersonNationality = "as"
	V2CorePersonNationalityAt V2CorePersonNationality = "at"
	V2CorePersonNationalityAu V2CorePersonNationality = "au"
	V2CorePersonNationalityAw V2CorePersonNationality = "aw"
	V2CorePersonNationalityAx V2CorePersonNationality = "ax"
	V2CorePersonNationalityAz V2CorePersonNationality = "az"
	V2CorePersonNationalityBa V2CorePersonNationality = "ba"
	V2CorePersonNationalityBb V2CorePersonNationality = "bb"
	V2CorePersonNationalityBd V2CorePersonNationality = "bd"
	V2CorePersonNationalityBe V2CorePersonNationality = "be"
	V2CorePersonNationalityBf V2CorePersonNationality = "bf"
	V2CorePersonNationalityBG V2CorePersonNationality = "bg"
	V2CorePersonNationalityBh V2CorePersonNationality = "bh"
	V2CorePersonNationalityBi V2CorePersonNationality = "bi"
	V2CorePersonNationalityBj V2CorePersonNationality = "bj"
	V2CorePersonNationalityBl V2CorePersonNationality = "bl"
	V2CorePersonNationalityBm V2CorePersonNationality = "bm"
	V2CorePersonNationalityBn V2CorePersonNationality = "bn"
	V2CorePersonNationalityBo V2CorePersonNationality = "bo"
	V2CorePersonNationalityBq V2CorePersonNationality = "bq"
	V2CorePersonNationalityBr V2CorePersonNationality = "br"
	V2CorePersonNationalityBs V2CorePersonNationality = "bs"
	V2CorePersonNationalityBt V2CorePersonNationality = "bt"
	V2CorePersonNationalityBv V2CorePersonNationality = "bv"
	V2CorePersonNationalityBw V2CorePersonNationality = "bw"
	V2CorePersonNationalityBy V2CorePersonNationality = "by"
	V2CorePersonNationalityBz V2CorePersonNationality = "bz"
	V2CorePersonNationalityCa V2CorePersonNationality = "ca"
	V2CorePersonNationalityCc V2CorePersonNationality = "cc"
	V2CorePersonNationalityCd V2CorePersonNationality = "cd"
	V2CorePersonNationalityCf V2CorePersonNationality = "cf"
	V2CorePersonNationalityCg V2CorePersonNationality = "cg"
	V2CorePersonNationalityCh V2CorePersonNationality = "ch"
	V2CorePersonNationalityCi V2CorePersonNationality = "ci"
	V2CorePersonNationalityCk V2CorePersonNationality = "ck"
	V2CorePersonNationalityCl V2CorePersonNationality = "cl"
	V2CorePersonNationalityCm V2CorePersonNationality = "cm"
	V2CorePersonNationalityCn V2CorePersonNationality = "cn"
	V2CorePersonNationalityCo V2CorePersonNationality = "co"
	V2CorePersonNationalityCr V2CorePersonNationality = "cr"
	V2CorePersonNationalityCu V2CorePersonNationality = "cu"
	V2CorePersonNationalityCv V2CorePersonNationality = "cv"
	V2CorePersonNationalityCw V2CorePersonNationality = "cw"
	V2CorePersonNationalityCx V2CorePersonNationality = "cx"
	V2CorePersonNationalityCy V2CorePersonNationality = "cy"
	V2CorePersonNationalityCz V2CorePersonNationality = "cz"
	V2CorePersonNationalityDE V2CorePersonNationality = "de"
	V2CorePersonNationalityDj V2CorePersonNationality = "dj"
	V2CorePersonNationalityDk V2CorePersonNationality = "dk"
	V2CorePersonNationalityDm V2CorePersonNationality = "dm"
	V2CorePersonNationalityDo V2CorePersonNationality = "do"
	V2CorePersonNationalityDz V2CorePersonNationality = "dz"
	V2CorePersonNationalityEc V2CorePersonNationality = "ec"
	V2CorePersonNationalityEe V2CorePersonNationality = "ee"
	V2CorePersonNationalityEg V2CorePersonNationality = "eg"
	V2CorePersonNationalityEh V2CorePersonNationality = "eh"
	V2CorePersonNationalityEr V2CorePersonNationality = "er"
	V2CorePersonNationalityES V2CorePersonNationality = "es"
	V2CorePersonNationalityET V2CorePersonNationality = "et"
	V2CorePersonNationalityFI V2CorePersonNationality = "fi"
	V2CorePersonNationalityFj V2CorePersonNationality = "fj"
	V2CorePersonNationalityFk V2CorePersonNationality = "fk"
	V2CorePersonNationalityFm V2CorePersonNationality = "fm"
	V2CorePersonNationalityFo V2CorePersonNationality = "fo"
	V2CorePersonNationalityFR V2CorePersonNationality = "fr"
	V2CorePersonNationalityGa V2CorePersonNationality = "ga"
	V2CorePersonNationalityGB V2CorePersonNationality = "gb"
	V2CorePersonNationalityGd V2CorePersonNationality = "gd"
	V2CorePersonNationalityGe V2CorePersonNationality = "ge"
	V2CorePersonNationalityGf V2CorePersonNationality = "gf"
	V2CorePersonNationalityGg V2CorePersonNationality = "gg"
	V2CorePersonNationalityGh V2CorePersonNationality = "gh"
	V2CorePersonNationalityGi V2CorePersonNationality = "gi"
	V2CorePersonNationalityGl V2CorePersonNationality = "gl"
	V2CorePersonNationalityGm V2CorePersonNationality = "gm"
	V2CorePersonNationalityGn V2CorePersonNationality = "gn"
	V2CorePersonNationalityGp V2CorePersonNationality = "gp"
	V2CorePersonNationalityGq V2CorePersonNationality = "gq"
	V2CorePersonNationalityGr V2CorePersonNationality = "gr"
	V2CorePersonNationalityGs V2CorePersonNationality = "gs"
	V2CorePersonNationalityGt V2CorePersonNationality = "gt"
	V2CorePersonNationalityGu V2CorePersonNationality = "gu"
	V2CorePersonNationalityGw V2CorePersonNationality = "gw"
	V2CorePersonNationalityGy V2CorePersonNationality = "gy"
	V2CorePersonNationalityHk V2CorePersonNationality = "hk"
	V2CorePersonNationalityHm V2CorePersonNationality = "hm"
	V2CorePersonNationalityHn V2CorePersonNationality = "hn"
	V2CorePersonNationalityHR V2CorePersonNationality = "hr"
	V2CorePersonNationalityHt V2CorePersonNationality = "ht"
	V2CorePersonNationalityHU V2CorePersonNationality = "hu"
	V2CorePersonNationalityID V2CorePersonNationality = "id"
	V2CorePersonNationalityIe V2CorePersonNationality = "ie"
	V2CorePersonNationalityIl V2CorePersonNationality = "il"
	V2CorePersonNationalityIm V2CorePersonNationality = "im"
	V2CorePersonNationalityIn V2CorePersonNationality = "in"
	V2CorePersonNationalityIo V2CorePersonNationality = "io"
	V2CorePersonNationalityIq V2CorePersonNationality = "iq"
	V2CorePersonNationalityIr V2CorePersonNationality = "ir"
	V2CorePersonNationalityIs V2CorePersonNationality = "is"
	V2CorePersonNationalityIT V2CorePersonNationality = "it"
	V2CorePersonNationalityJe V2CorePersonNationality = "je"
	V2CorePersonNationalityJm V2CorePersonNationality = "jm"
	V2CorePersonNationalityJo V2CorePersonNationality = "jo"
	V2CorePersonNationalityJP V2CorePersonNationality = "jp"
	V2CorePersonNationalityKe V2CorePersonNationality = "ke"
	V2CorePersonNationalityKg V2CorePersonNationality = "kg"
	V2CorePersonNationalityKh V2CorePersonNationality = "kh"
	V2CorePersonNationalityKi V2CorePersonNationality = "ki"
	V2CorePersonNationalityKm V2CorePersonNationality = "km"
	V2CorePersonNationalityKn V2CorePersonNationality = "kn"
	V2CorePersonNationalityKp V2CorePersonNationality = "kp"
	V2CorePersonNationalityKr V2CorePersonNationality = "kr"
	V2CorePersonNationalityKw V2CorePersonNationality = "kw"
	V2CorePersonNationalityKy V2CorePersonNationality = "ky"
	V2CorePersonNationalityKz V2CorePersonNationality = "kz"
	V2CorePersonNationalityLa V2CorePersonNationality = "la"
	V2CorePersonNationalityLb V2CorePersonNationality = "lb"
	V2CorePersonNationalityLc V2CorePersonNationality = "lc"
	V2CorePersonNationalityLi V2CorePersonNationality = "li"
	V2CorePersonNationalityLk V2CorePersonNationality = "lk"
	V2CorePersonNationalityLr V2CorePersonNationality = "lr"
	V2CorePersonNationalityLs V2CorePersonNationality = "ls"
	V2CorePersonNationalityLT V2CorePersonNationality = "lt"
	V2CorePersonNationalityLu V2CorePersonNationality = "lu"
	V2CorePersonNationalityLV V2CorePersonNationality = "lv"
	V2CorePersonNationalityLy V2CorePersonNationality = "ly"
	V2CorePersonNationalityMa V2CorePersonNationality = "ma"
	V2CorePersonNationalityMc V2CorePersonNationality = "mc"
	V2CorePersonNationalityMd V2CorePersonNationality = "md"
	V2CorePersonNationalityMe V2CorePersonNationality = "me"
	V2CorePersonNationalityMf V2CorePersonNationality = "mf"
	V2CorePersonNationalityMg V2CorePersonNationality = "mg"
	V2CorePersonNationalityMh V2CorePersonNationality = "mh"
	V2CorePersonNationalityMk V2CorePersonNationality = "mk"
	V2CorePersonNationalityMl V2CorePersonNationality = "ml"
	V2CorePersonNationalityMm V2CorePersonNationality = "mm"
	V2CorePersonNationalityMn V2CorePersonNationality = "mn"
	V2CorePersonNationalityMo V2CorePersonNationality = "mo"
	V2CorePersonNationalityMp V2CorePersonNationality = "mp"
	V2CorePersonNationalityMq V2CorePersonNationality = "mq"
	V2CorePersonNationalityMr V2CorePersonNationality = "mr"
	V2CorePersonNationalityMS V2CorePersonNationality = "ms"
	V2CorePersonNationalityMT V2CorePersonNationality = "mt"
	V2CorePersonNationalityMu V2CorePersonNationality = "mu"
	V2CorePersonNationalityMv V2CorePersonNationality = "mv"
	V2CorePersonNationalityMw V2CorePersonNationality = "mw"
	V2CorePersonNationalityMX V2CorePersonNationality = "mx"
	V2CorePersonNationalityMy V2CorePersonNationality = "my"
	V2CorePersonNationalityMz V2CorePersonNationality = "mz"
	V2CorePersonNationalityNa V2CorePersonNationality = "na"
	V2CorePersonNationalityNc V2CorePersonNationality = "nc"
	V2CorePersonNationalityNe V2CorePersonNationality = "ne"
	V2CorePersonNationalityNf V2CorePersonNationality = "nf"
	V2CorePersonNationalityNg V2CorePersonNationality = "ng"
	V2CorePersonNationalityNi V2CorePersonNationality = "ni"
	V2CorePersonNationalityNL V2CorePersonNationality = "nl"
	V2CorePersonNationalityNo V2CorePersonNationality = "no"
	V2CorePersonNationalityNp V2CorePersonNationality = "np"
	V2CorePersonNationalityNr V2CorePersonNationality = "nr"
	V2CorePersonNationalityNu V2CorePersonNationality = "nu"
	V2CorePersonNationalityNz V2CorePersonNationality = "nz"
	V2CorePersonNationalityOm V2CorePersonNationality = "om"
	V2CorePersonNationalityPa V2CorePersonNationality = "pa"
	V2CorePersonNationalityPe V2CorePersonNationality = "pe"
	V2CorePersonNationalityPf V2CorePersonNationality = "pf"
	V2CorePersonNationalityPg V2CorePersonNationality = "pg"
	V2CorePersonNationalityPh V2CorePersonNationality = "ph"
	V2CorePersonNationalityPk V2CorePersonNationality = "pk"
	V2CorePersonNationalityPL V2CorePersonNationality = "pl"
	V2CorePersonNationalityPm V2CorePersonNationality = "pm"
	V2CorePersonNationalityPn V2CorePersonNationality = "pn"
	V2CorePersonNationalityPr V2CorePersonNationality = "pr"
	V2CorePersonNationalityPs V2CorePersonNationality = "ps"
	V2CorePersonNationalityPT V2CorePersonNationality = "pt"
	V2CorePersonNationalityPw V2CorePersonNationality = "pw"
	V2CorePersonNationalityPy V2CorePersonNationality = "py"
	V2CorePersonNationalityQa V2CorePersonNationality = "qa"
	V2CorePersonNationalityQz V2CorePersonNationality = "qz"
	V2CorePersonNationalityRe V2CorePersonNationality = "re"
	V2CorePersonNationalityRO V2CorePersonNationality = "ro"
	V2CorePersonNationalityRs V2CorePersonNationality = "rs"
	V2CorePersonNationalityRU V2CorePersonNationality = "ru"
	V2CorePersonNationalityRw V2CorePersonNationality = "rw"
	V2CorePersonNationalitySa V2CorePersonNationality = "sa"
	V2CorePersonNationalitySb V2CorePersonNationality = "sb"
	V2CorePersonNationalitySc V2CorePersonNationality = "sc"
	V2CorePersonNationalitySd V2CorePersonNationality = "sd"
	V2CorePersonNationalitySe V2CorePersonNationality = "se"
	V2CorePersonNationalitySg V2CorePersonNationality = "sg"
	V2CorePersonNationalitySh V2CorePersonNationality = "sh"
	V2CorePersonNationalitySi V2CorePersonNationality = "si"
	V2CorePersonNationalitySj V2CorePersonNationality = "sj"
	V2CorePersonNationalitySK V2CorePersonNationality = "sk"
	V2CorePersonNationalitySL V2CorePersonNationality = "sl"
	V2CorePersonNationalitySm V2CorePersonNationality = "sm"
	V2CorePersonNationalitySn V2CorePersonNationality = "sn"
	V2CorePersonNationalitySo V2CorePersonNationality = "so"
	V2CorePersonNationalitySr V2CorePersonNationality = "sr"
	V2CorePersonNationalitySs V2CorePersonNationality = "ss"
	V2CorePersonNationalitySt V2CorePersonNationality = "st"
	V2CorePersonNationalitySV V2CorePersonNationality = "sv"
	V2CorePersonNationalitySx V2CorePersonNationality = "sx"
	V2CorePersonNationalitySy V2CorePersonNationality = "sy"
	V2CorePersonNationalitySz V2CorePersonNationality = "sz"
	V2CorePersonNationalityTc V2CorePersonNationality = "tc"
	V2CorePersonNationalityTd V2CorePersonNationality = "td"
	V2CorePersonNationalityTf V2CorePersonNationality = "tf"
	V2CorePersonNationalityTg V2CorePersonNationality = "tg"
	V2CorePersonNationalityTH V2CorePersonNationality = "th"
	V2CorePersonNationalityTj V2CorePersonNationality = "tj"
	V2CorePersonNationalityTk V2CorePersonNationality = "tk"
	V2CorePersonNationalityTl V2CorePersonNationality = "tl"
	V2CorePersonNationalityTm V2CorePersonNationality = "tm"
	V2CorePersonNationalityTn V2CorePersonNationality = "tn"
	V2CorePersonNationalityTo V2CorePersonNationality = "to"
	V2CorePersonNationalityTR V2CorePersonNationality = "tr"
	V2CorePersonNationalityTt V2CorePersonNationality = "tt"
	V2CorePersonNationalityTv V2CorePersonNationality = "tv"
	V2CorePersonNationalityTw V2CorePersonNationality = "tw"
	V2CorePersonNationalityTz V2CorePersonNationality = "tz"
	V2CorePersonNationalityUa V2CorePersonNationality = "ua"
	V2CorePersonNationalityUg V2CorePersonNationality = "ug"
	V2CorePersonNationalityUm V2CorePersonNationality = "um"
	V2CorePersonNationalityUS V2CorePersonNationality = "us"
	V2CorePersonNationalityUy V2CorePersonNationality = "uy"
	V2CorePersonNationalityUz V2CorePersonNationality = "uz"
	V2CorePersonNationalityVa V2CorePersonNationality = "va"
	V2CorePersonNationalityVc V2CorePersonNationality = "vc"
	V2CorePersonNationalityVe V2CorePersonNationality = "ve"
	V2CorePersonNationalityVg V2CorePersonNationality = "vg"
	V2CorePersonNationalityVI V2CorePersonNationality = "vi"
	V2CorePersonNationalityVn V2CorePersonNationality = "vn"
	V2CorePersonNationalityVu V2CorePersonNationality = "vu"
	V2CorePersonNationalityWf V2CorePersonNationality = "wf"
	V2CorePersonNationalityWs V2CorePersonNationality = "ws"
	V2CorePersonNationalityXx V2CorePersonNationality = "xx"
	V2CorePersonNationalityYe V2CorePersonNationality = "ye"
	V2CorePersonNationalityYt V2CorePersonNationality = "yt"
	V2CorePersonNationalityZa V2CorePersonNationality = "za"
	V2CorePersonNationalityZm V2CorePersonNationality = "zm"
	V2CorePersonNationalityZw V2CorePersonNationality = "zw"
)

// The person's political exposure.
type V2CorePersonPoliticalExposure string

// List of values that V2CorePersonPoliticalExposure can take
const (
	V2CorePersonPoliticalExposureExisting V2CorePersonPoliticalExposure = "existing"
	V2CorePersonPoliticalExposureNone     V2CorePersonPoliticalExposure = "none"
)

// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CorePersonScriptAddressesKanaCountry string

// List of values that V2CorePersonScriptAddressesKanaCountry can take
const (
	V2CorePersonScriptAddressesKanaCountryAd V2CorePersonScriptAddressesKanaCountry = "ad"
	V2CorePersonScriptAddressesKanaCountryAe V2CorePersonScriptAddressesKanaCountry = "ae"
	V2CorePersonScriptAddressesKanaCountryAf V2CorePersonScriptAddressesKanaCountry = "af"
	V2CorePersonScriptAddressesKanaCountryAg V2CorePersonScriptAddressesKanaCountry = "ag"
	V2CorePersonScriptAddressesKanaCountryAi V2CorePersonScriptAddressesKanaCountry = "ai"
	V2CorePersonScriptAddressesKanaCountryAl V2CorePersonScriptAddressesKanaCountry = "al"
	V2CorePersonScriptAddressesKanaCountryAm V2CorePersonScriptAddressesKanaCountry = "am"
	V2CorePersonScriptAddressesKanaCountryAo V2CorePersonScriptAddressesKanaCountry = "ao"
	V2CorePersonScriptAddressesKanaCountryAq V2CorePersonScriptAddressesKanaCountry = "aq"
	V2CorePersonScriptAddressesKanaCountryAR V2CorePersonScriptAddressesKanaCountry = "ar"
	V2CorePersonScriptAddressesKanaCountryAs V2CorePersonScriptAddressesKanaCountry = "as"
	V2CorePersonScriptAddressesKanaCountryAt V2CorePersonScriptAddressesKanaCountry = "at"
	V2CorePersonScriptAddressesKanaCountryAu V2CorePersonScriptAddressesKanaCountry = "au"
	V2CorePersonScriptAddressesKanaCountryAw V2CorePersonScriptAddressesKanaCountry = "aw"
	V2CorePersonScriptAddressesKanaCountryAx V2CorePersonScriptAddressesKanaCountry = "ax"
	V2CorePersonScriptAddressesKanaCountryAz V2CorePersonScriptAddressesKanaCountry = "az"
	V2CorePersonScriptAddressesKanaCountryBa V2CorePersonScriptAddressesKanaCountry = "ba"
	V2CorePersonScriptAddressesKanaCountryBb V2CorePersonScriptAddressesKanaCountry = "bb"
	V2CorePersonScriptAddressesKanaCountryBd V2CorePersonScriptAddressesKanaCountry = "bd"
	V2CorePersonScriptAddressesKanaCountryBe V2CorePersonScriptAddressesKanaCountry = "be"
	V2CorePersonScriptAddressesKanaCountryBf V2CorePersonScriptAddressesKanaCountry = "bf"
	V2CorePersonScriptAddressesKanaCountryBG V2CorePersonScriptAddressesKanaCountry = "bg"
	V2CorePersonScriptAddressesKanaCountryBh V2CorePersonScriptAddressesKanaCountry = "bh"
	V2CorePersonScriptAddressesKanaCountryBi V2CorePersonScriptAddressesKanaCountry = "bi"
	V2CorePersonScriptAddressesKanaCountryBj V2CorePersonScriptAddressesKanaCountry = "bj"
	V2CorePersonScriptAddressesKanaCountryBl V2CorePersonScriptAddressesKanaCountry = "bl"
	V2CorePersonScriptAddressesKanaCountryBm V2CorePersonScriptAddressesKanaCountry = "bm"
	V2CorePersonScriptAddressesKanaCountryBn V2CorePersonScriptAddressesKanaCountry = "bn"
	V2CorePersonScriptAddressesKanaCountryBo V2CorePersonScriptAddressesKanaCountry = "bo"
	V2CorePersonScriptAddressesKanaCountryBq V2CorePersonScriptAddressesKanaCountry = "bq"
	V2CorePersonScriptAddressesKanaCountryBr V2CorePersonScriptAddressesKanaCountry = "br"
	V2CorePersonScriptAddressesKanaCountryBs V2CorePersonScriptAddressesKanaCountry = "bs"
	V2CorePersonScriptAddressesKanaCountryBt V2CorePersonScriptAddressesKanaCountry = "bt"
	V2CorePersonScriptAddressesKanaCountryBv V2CorePersonScriptAddressesKanaCountry = "bv"
	V2CorePersonScriptAddressesKanaCountryBw V2CorePersonScriptAddressesKanaCountry = "bw"
	V2CorePersonScriptAddressesKanaCountryBy V2CorePersonScriptAddressesKanaCountry = "by"
	V2CorePersonScriptAddressesKanaCountryBz V2CorePersonScriptAddressesKanaCountry = "bz"
	V2CorePersonScriptAddressesKanaCountryCa V2CorePersonScriptAddressesKanaCountry = "ca"
	V2CorePersonScriptAddressesKanaCountryCc V2CorePersonScriptAddressesKanaCountry = "cc"
	V2CorePersonScriptAddressesKanaCountryCd V2CorePersonScriptAddressesKanaCountry = "cd"
	V2CorePersonScriptAddressesKanaCountryCf V2CorePersonScriptAddressesKanaCountry = "cf"
	V2CorePersonScriptAddressesKanaCountryCg V2CorePersonScriptAddressesKanaCountry = "cg"
	V2CorePersonScriptAddressesKanaCountryCh V2CorePersonScriptAddressesKanaCountry = "ch"
	V2CorePersonScriptAddressesKanaCountryCi V2CorePersonScriptAddressesKanaCountry = "ci"
	V2CorePersonScriptAddressesKanaCountryCk V2CorePersonScriptAddressesKanaCountry = "ck"
	V2CorePersonScriptAddressesKanaCountryCl V2CorePersonScriptAddressesKanaCountry = "cl"
	V2CorePersonScriptAddressesKanaCountryCm V2CorePersonScriptAddressesKanaCountry = "cm"
	V2CorePersonScriptAddressesKanaCountryCn V2CorePersonScriptAddressesKanaCountry = "cn"
	V2CorePersonScriptAddressesKanaCountryCo V2CorePersonScriptAddressesKanaCountry = "co"
	V2CorePersonScriptAddressesKanaCountryCr V2CorePersonScriptAddressesKanaCountry = "cr"
	V2CorePersonScriptAddressesKanaCountryCu V2CorePersonScriptAddressesKanaCountry = "cu"
	V2CorePersonScriptAddressesKanaCountryCv V2CorePersonScriptAddressesKanaCountry = "cv"
	V2CorePersonScriptAddressesKanaCountryCw V2CorePersonScriptAddressesKanaCountry = "cw"
	V2CorePersonScriptAddressesKanaCountryCx V2CorePersonScriptAddressesKanaCountry = "cx"
	V2CorePersonScriptAddressesKanaCountryCy V2CorePersonScriptAddressesKanaCountry = "cy"
	V2CorePersonScriptAddressesKanaCountryCz V2CorePersonScriptAddressesKanaCountry = "cz"
	V2CorePersonScriptAddressesKanaCountryDE V2CorePersonScriptAddressesKanaCountry = "de"
	V2CorePersonScriptAddressesKanaCountryDj V2CorePersonScriptAddressesKanaCountry = "dj"
	V2CorePersonScriptAddressesKanaCountryDk V2CorePersonScriptAddressesKanaCountry = "dk"
	V2CorePersonScriptAddressesKanaCountryDm V2CorePersonScriptAddressesKanaCountry = "dm"
	V2CorePersonScriptAddressesKanaCountryDo V2CorePersonScriptAddressesKanaCountry = "do"
	V2CorePersonScriptAddressesKanaCountryDz V2CorePersonScriptAddressesKanaCountry = "dz"
	V2CorePersonScriptAddressesKanaCountryEc V2CorePersonScriptAddressesKanaCountry = "ec"
	V2CorePersonScriptAddressesKanaCountryEe V2CorePersonScriptAddressesKanaCountry = "ee"
	V2CorePersonScriptAddressesKanaCountryEg V2CorePersonScriptAddressesKanaCountry = "eg"
	V2CorePersonScriptAddressesKanaCountryEh V2CorePersonScriptAddressesKanaCountry = "eh"
	V2CorePersonScriptAddressesKanaCountryEr V2CorePersonScriptAddressesKanaCountry = "er"
	V2CorePersonScriptAddressesKanaCountryES V2CorePersonScriptAddressesKanaCountry = "es"
	V2CorePersonScriptAddressesKanaCountryET V2CorePersonScriptAddressesKanaCountry = "et"
	V2CorePersonScriptAddressesKanaCountryFI V2CorePersonScriptAddressesKanaCountry = "fi"
	V2CorePersonScriptAddressesKanaCountryFj V2CorePersonScriptAddressesKanaCountry = "fj"
	V2CorePersonScriptAddressesKanaCountryFk V2CorePersonScriptAddressesKanaCountry = "fk"
	V2CorePersonScriptAddressesKanaCountryFm V2CorePersonScriptAddressesKanaCountry = "fm"
	V2CorePersonScriptAddressesKanaCountryFo V2CorePersonScriptAddressesKanaCountry = "fo"
	V2CorePersonScriptAddressesKanaCountryFR V2CorePersonScriptAddressesKanaCountry = "fr"
	V2CorePersonScriptAddressesKanaCountryGa V2CorePersonScriptAddressesKanaCountry = "ga"
	V2CorePersonScriptAddressesKanaCountryGB V2CorePersonScriptAddressesKanaCountry = "gb"
	V2CorePersonScriptAddressesKanaCountryGd V2CorePersonScriptAddressesKanaCountry = "gd"
	V2CorePersonScriptAddressesKanaCountryGe V2CorePersonScriptAddressesKanaCountry = "ge"
	V2CorePersonScriptAddressesKanaCountryGf V2CorePersonScriptAddressesKanaCountry = "gf"
	V2CorePersonScriptAddressesKanaCountryGg V2CorePersonScriptAddressesKanaCountry = "gg"
	V2CorePersonScriptAddressesKanaCountryGh V2CorePersonScriptAddressesKanaCountry = "gh"
	V2CorePersonScriptAddressesKanaCountryGi V2CorePersonScriptAddressesKanaCountry = "gi"
	V2CorePersonScriptAddressesKanaCountryGl V2CorePersonScriptAddressesKanaCountry = "gl"
	V2CorePersonScriptAddressesKanaCountryGm V2CorePersonScriptAddressesKanaCountry = "gm"
	V2CorePersonScriptAddressesKanaCountryGn V2CorePersonScriptAddressesKanaCountry = "gn"
	V2CorePersonScriptAddressesKanaCountryGp V2CorePersonScriptAddressesKanaCountry = "gp"
	V2CorePersonScriptAddressesKanaCountryGq V2CorePersonScriptAddressesKanaCountry = "gq"
	V2CorePersonScriptAddressesKanaCountryGr V2CorePersonScriptAddressesKanaCountry = "gr"
	V2CorePersonScriptAddressesKanaCountryGs V2CorePersonScriptAddressesKanaCountry = "gs"
	V2CorePersonScriptAddressesKanaCountryGt V2CorePersonScriptAddressesKanaCountry = "gt"
	V2CorePersonScriptAddressesKanaCountryGu V2CorePersonScriptAddressesKanaCountry = "gu"
	V2CorePersonScriptAddressesKanaCountryGw V2CorePersonScriptAddressesKanaCountry = "gw"
	V2CorePersonScriptAddressesKanaCountryGy V2CorePersonScriptAddressesKanaCountry = "gy"
	V2CorePersonScriptAddressesKanaCountryHk V2CorePersonScriptAddressesKanaCountry = "hk"
	V2CorePersonScriptAddressesKanaCountryHm V2CorePersonScriptAddressesKanaCountry = "hm"
	V2CorePersonScriptAddressesKanaCountryHn V2CorePersonScriptAddressesKanaCountry = "hn"
	V2CorePersonScriptAddressesKanaCountryHR V2CorePersonScriptAddressesKanaCountry = "hr"
	V2CorePersonScriptAddressesKanaCountryHt V2CorePersonScriptAddressesKanaCountry = "ht"
	V2CorePersonScriptAddressesKanaCountryHU V2CorePersonScriptAddressesKanaCountry = "hu"
	V2CorePersonScriptAddressesKanaCountryID V2CorePersonScriptAddressesKanaCountry = "id"
	V2CorePersonScriptAddressesKanaCountryIe V2CorePersonScriptAddressesKanaCountry = "ie"
	V2CorePersonScriptAddressesKanaCountryIl V2CorePersonScriptAddressesKanaCountry = "il"
	V2CorePersonScriptAddressesKanaCountryIm V2CorePersonScriptAddressesKanaCountry = "im"
	V2CorePersonScriptAddressesKanaCountryIn V2CorePersonScriptAddressesKanaCountry = "in"
	V2CorePersonScriptAddressesKanaCountryIo V2CorePersonScriptAddressesKanaCountry = "io"
	V2CorePersonScriptAddressesKanaCountryIq V2CorePersonScriptAddressesKanaCountry = "iq"
	V2CorePersonScriptAddressesKanaCountryIr V2CorePersonScriptAddressesKanaCountry = "ir"
	V2CorePersonScriptAddressesKanaCountryIs V2CorePersonScriptAddressesKanaCountry = "is"
	V2CorePersonScriptAddressesKanaCountryIT V2CorePersonScriptAddressesKanaCountry = "it"
	V2CorePersonScriptAddressesKanaCountryJe V2CorePersonScriptAddressesKanaCountry = "je"
	V2CorePersonScriptAddressesKanaCountryJm V2CorePersonScriptAddressesKanaCountry = "jm"
	V2CorePersonScriptAddressesKanaCountryJo V2CorePersonScriptAddressesKanaCountry = "jo"
	V2CorePersonScriptAddressesKanaCountryJP V2CorePersonScriptAddressesKanaCountry = "jp"
	V2CorePersonScriptAddressesKanaCountryKe V2CorePersonScriptAddressesKanaCountry = "ke"
	V2CorePersonScriptAddressesKanaCountryKg V2CorePersonScriptAddressesKanaCountry = "kg"
	V2CorePersonScriptAddressesKanaCountryKh V2CorePersonScriptAddressesKanaCountry = "kh"
	V2CorePersonScriptAddressesKanaCountryKi V2CorePersonScriptAddressesKanaCountry = "ki"
	V2CorePersonScriptAddressesKanaCountryKm V2CorePersonScriptAddressesKanaCountry = "km"
	V2CorePersonScriptAddressesKanaCountryKn V2CorePersonScriptAddressesKanaCountry = "kn"
	V2CorePersonScriptAddressesKanaCountryKp V2CorePersonScriptAddressesKanaCountry = "kp"
	V2CorePersonScriptAddressesKanaCountryKr V2CorePersonScriptAddressesKanaCountry = "kr"
	V2CorePersonScriptAddressesKanaCountryKw V2CorePersonScriptAddressesKanaCountry = "kw"
	V2CorePersonScriptAddressesKanaCountryKy V2CorePersonScriptAddressesKanaCountry = "ky"
	V2CorePersonScriptAddressesKanaCountryKz V2CorePersonScriptAddressesKanaCountry = "kz"
	V2CorePersonScriptAddressesKanaCountryLa V2CorePersonScriptAddressesKanaCountry = "la"
	V2CorePersonScriptAddressesKanaCountryLb V2CorePersonScriptAddressesKanaCountry = "lb"
	V2CorePersonScriptAddressesKanaCountryLc V2CorePersonScriptAddressesKanaCountry = "lc"
	V2CorePersonScriptAddressesKanaCountryLi V2CorePersonScriptAddressesKanaCountry = "li"
	V2CorePersonScriptAddressesKanaCountryLk V2CorePersonScriptAddressesKanaCountry = "lk"
	V2CorePersonScriptAddressesKanaCountryLr V2CorePersonScriptAddressesKanaCountry = "lr"
	V2CorePersonScriptAddressesKanaCountryLs V2CorePersonScriptAddressesKanaCountry = "ls"
	V2CorePersonScriptAddressesKanaCountryLT V2CorePersonScriptAddressesKanaCountry = "lt"
	V2CorePersonScriptAddressesKanaCountryLu V2CorePersonScriptAddressesKanaCountry = "lu"
	V2CorePersonScriptAddressesKanaCountryLV V2CorePersonScriptAddressesKanaCountry = "lv"
	V2CorePersonScriptAddressesKanaCountryLy V2CorePersonScriptAddressesKanaCountry = "ly"
	V2CorePersonScriptAddressesKanaCountryMa V2CorePersonScriptAddressesKanaCountry = "ma"
	V2CorePersonScriptAddressesKanaCountryMc V2CorePersonScriptAddressesKanaCountry = "mc"
	V2CorePersonScriptAddressesKanaCountryMd V2CorePersonScriptAddressesKanaCountry = "md"
	V2CorePersonScriptAddressesKanaCountryMe V2CorePersonScriptAddressesKanaCountry = "me"
	V2CorePersonScriptAddressesKanaCountryMf V2CorePersonScriptAddressesKanaCountry = "mf"
	V2CorePersonScriptAddressesKanaCountryMg V2CorePersonScriptAddressesKanaCountry = "mg"
	V2CorePersonScriptAddressesKanaCountryMh V2CorePersonScriptAddressesKanaCountry = "mh"
	V2CorePersonScriptAddressesKanaCountryMk V2CorePersonScriptAddressesKanaCountry = "mk"
	V2CorePersonScriptAddressesKanaCountryMl V2CorePersonScriptAddressesKanaCountry = "ml"
	V2CorePersonScriptAddressesKanaCountryMm V2CorePersonScriptAddressesKanaCountry = "mm"
	V2CorePersonScriptAddressesKanaCountryMn V2CorePersonScriptAddressesKanaCountry = "mn"
	V2CorePersonScriptAddressesKanaCountryMo V2CorePersonScriptAddressesKanaCountry = "mo"
	V2CorePersonScriptAddressesKanaCountryMp V2CorePersonScriptAddressesKanaCountry = "mp"
	V2CorePersonScriptAddressesKanaCountryMq V2CorePersonScriptAddressesKanaCountry = "mq"
	V2CorePersonScriptAddressesKanaCountryMr V2CorePersonScriptAddressesKanaCountry = "mr"
	V2CorePersonScriptAddressesKanaCountryMS V2CorePersonScriptAddressesKanaCountry = "ms"
	V2CorePersonScriptAddressesKanaCountryMT V2CorePersonScriptAddressesKanaCountry = "mt"
	V2CorePersonScriptAddressesKanaCountryMu V2CorePersonScriptAddressesKanaCountry = "mu"
	V2CorePersonScriptAddressesKanaCountryMv V2CorePersonScriptAddressesKanaCountry = "mv"
	V2CorePersonScriptAddressesKanaCountryMw V2CorePersonScriptAddressesKanaCountry = "mw"
	V2CorePersonScriptAddressesKanaCountryMX V2CorePersonScriptAddressesKanaCountry = "mx"
	V2CorePersonScriptAddressesKanaCountryMy V2CorePersonScriptAddressesKanaCountry = "my"
	V2CorePersonScriptAddressesKanaCountryMz V2CorePersonScriptAddressesKanaCountry = "mz"
	V2CorePersonScriptAddressesKanaCountryNa V2CorePersonScriptAddressesKanaCountry = "na"
	V2CorePersonScriptAddressesKanaCountryNc V2CorePersonScriptAddressesKanaCountry = "nc"
	V2CorePersonScriptAddressesKanaCountryNe V2CorePersonScriptAddressesKanaCountry = "ne"
	V2CorePersonScriptAddressesKanaCountryNf V2CorePersonScriptAddressesKanaCountry = "nf"
	V2CorePersonScriptAddressesKanaCountryNg V2CorePersonScriptAddressesKanaCountry = "ng"
	V2CorePersonScriptAddressesKanaCountryNi V2CorePersonScriptAddressesKanaCountry = "ni"
	V2CorePersonScriptAddressesKanaCountryNL V2CorePersonScriptAddressesKanaCountry = "nl"
	V2CorePersonScriptAddressesKanaCountryNo V2CorePersonScriptAddressesKanaCountry = "no"
	V2CorePersonScriptAddressesKanaCountryNp V2CorePersonScriptAddressesKanaCountry = "np"
	V2CorePersonScriptAddressesKanaCountryNr V2CorePersonScriptAddressesKanaCountry = "nr"
	V2CorePersonScriptAddressesKanaCountryNu V2CorePersonScriptAddressesKanaCountry = "nu"
	V2CorePersonScriptAddressesKanaCountryNz V2CorePersonScriptAddressesKanaCountry = "nz"
	V2CorePersonScriptAddressesKanaCountryOm V2CorePersonScriptAddressesKanaCountry = "om"
	V2CorePersonScriptAddressesKanaCountryPa V2CorePersonScriptAddressesKanaCountry = "pa"
	V2CorePersonScriptAddressesKanaCountryPe V2CorePersonScriptAddressesKanaCountry = "pe"
	V2CorePersonScriptAddressesKanaCountryPf V2CorePersonScriptAddressesKanaCountry = "pf"
	V2CorePersonScriptAddressesKanaCountryPg V2CorePersonScriptAddressesKanaCountry = "pg"
	V2CorePersonScriptAddressesKanaCountryPh V2CorePersonScriptAddressesKanaCountry = "ph"
	V2CorePersonScriptAddressesKanaCountryPk V2CorePersonScriptAddressesKanaCountry = "pk"
	V2CorePersonScriptAddressesKanaCountryPL V2CorePersonScriptAddressesKanaCountry = "pl"
	V2CorePersonScriptAddressesKanaCountryPm V2CorePersonScriptAddressesKanaCountry = "pm"
	V2CorePersonScriptAddressesKanaCountryPn V2CorePersonScriptAddressesKanaCountry = "pn"
	V2CorePersonScriptAddressesKanaCountryPr V2CorePersonScriptAddressesKanaCountry = "pr"
	V2CorePersonScriptAddressesKanaCountryPs V2CorePersonScriptAddressesKanaCountry = "ps"
	V2CorePersonScriptAddressesKanaCountryPT V2CorePersonScriptAddressesKanaCountry = "pt"
	V2CorePersonScriptAddressesKanaCountryPw V2CorePersonScriptAddressesKanaCountry = "pw"
	V2CorePersonScriptAddressesKanaCountryPy V2CorePersonScriptAddressesKanaCountry = "py"
	V2CorePersonScriptAddressesKanaCountryQa V2CorePersonScriptAddressesKanaCountry = "qa"
	V2CorePersonScriptAddressesKanaCountryQz V2CorePersonScriptAddressesKanaCountry = "qz"
	V2CorePersonScriptAddressesKanaCountryRe V2CorePersonScriptAddressesKanaCountry = "re"
	V2CorePersonScriptAddressesKanaCountryRO V2CorePersonScriptAddressesKanaCountry = "ro"
	V2CorePersonScriptAddressesKanaCountryRs V2CorePersonScriptAddressesKanaCountry = "rs"
	V2CorePersonScriptAddressesKanaCountryRU V2CorePersonScriptAddressesKanaCountry = "ru"
	V2CorePersonScriptAddressesKanaCountryRw V2CorePersonScriptAddressesKanaCountry = "rw"
	V2CorePersonScriptAddressesKanaCountrySa V2CorePersonScriptAddressesKanaCountry = "sa"
	V2CorePersonScriptAddressesKanaCountrySb V2CorePersonScriptAddressesKanaCountry = "sb"
	V2CorePersonScriptAddressesKanaCountrySc V2CorePersonScriptAddressesKanaCountry = "sc"
	V2CorePersonScriptAddressesKanaCountrySd V2CorePersonScriptAddressesKanaCountry = "sd"
	V2CorePersonScriptAddressesKanaCountrySe V2CorePersonScriptAddressesKanaCountry = "se"
	V2CorePersonScriptAddressesKanaCountrySg V2CorePersonScriptAddressesKanaCountry = "sg"
	V2CorePersonScriptAddressesKanaCountrySh V2CorePersonScriptAddressesKanaCountry = "sh"
	V2CorePersonScriptAddressesKanaCountrySi V2CorePersonScriptAddressesKanaCountry = "si"
	V2CorePersonScriptAddressesKanaCountrySj V2CorePersonScriptAddressesKanaCountry = "sj"
	V2CorePersonScriptAddressesKanaCountrySK V2CorePersonScriptAddressesKanaCountry = "sk"
	V2CorePersonScriptAddressesKanaCountrySL V2CorePersonScriptAddressesKanaCountry = "sl"
	V2CorePersonScriptAddressesKanaCountrySm V2CorePersonScriptAddressesKanaCountry = "sm"
	V2CorePersonScriptAddressesKanaCountrySn V2CorePersonScriptAddressesKanaCountry = "sn"
	V2CorePersonScriptAddressesKanaCountrySo V2CorePersonScriptAddressesKanaCountry = "so"
	V2CorePersonScriptAddressesKanaCountrySr V2CorePersonScriptAddressesKanaCountry = "sr"
	V2CorePersonScriptAddressesKanaCountrySs V2CorePersonScriptAddressesKanaCountry = "ss"
	V2CorePersonScriptAddressesKanaCountrySt V2CorePersonScriptAddressesKanaCountry = "st"
	V2CorePersonScriptAddressesKanaCountrySV V2CorePersonScriptAddressesKanaCountry = "sv"
	V2CorePersonScriptAddressesKanaCountrySx V2CorePersonScriptAddressesKanaCountry = "sx"
	V2CorePersonScriptAddressesKanaCountrySy V2CorePersonScriptAddressesKanaCountry = "sy"
	V2CorePersonScriptAddressesKanaCountrySz V2CorePersonScriptAddressesKanaCountry = "sz"
	V2CorePersonScriptAddressesKanaCountryTc V2CorePersonScriptAddressesKanaCountry = "tc"
	V2CorePersonScriptAddressesKanaCountryTd V2CorePersonScriptAddressesKanaCountry = "td"
	V2CorePersonScriptAddressesKanaCountryTf V2CorePersonScriptAddressesKanaCountry = "tf"
	V2CorePersonScriptAddressesKanaCountryTg V2CorePersonScriptAddressesKanaCountry = "tg"
	V2CorePersonScriptAddressesKanaCountryTH V2CorePersonScriptAddressesKanaCountry = "th"
	V2CorePersonScriptAddressesKanaCountryTj V2CorePersonScriptAddressesKanaCountry = "tj"
	V2CorePersonScriptAddressesKanaCountryTk V2CorePersonScriptAddressesKanaCountry = "tk"
	V2CorePersonScriptAddressesKanaCountryTl V2CorePersonScriptAddressesKanaCountry = "tl"
	V2CorePersonScriptAddressesKanaCountryTm V2CorePersonScriptAddressesKanaCountry = "tm"
	V2CorePersonScriptAddressesKanaCountryTn V2CorePersonScriptAddressesKanaCountry = "tn"
	V2CorePersonScriptAddressesKanaCountryTo V2CorePersonScriptAddressesKanaCountry = "to"
	V2CorePersonScriptAddressesKanaCountryTR V2CorePersonScriptAddressesKanaCountry = "tr"
	V2CorePersonScriptAddressesKanaCountryTt V2CorePersonScriptAddressesKanaCountry = "tt"
	V2CorePersonScriptAddressesKanaCountryTv V2CorePersonScriptAddressesKanaCountry = "tv"
	V2CorePersonScriptAddressesKanaCountryTw V2CorePersonScriptAddressesKanaCountry = "tw"
	V2CorePersonScriptAddressesKanaCountryTz V2CorePersonScriptAddressesKanaCountry = "tz"
	V2CorePersonScriptAddressesKanaCountryUa V2CorePersonScriptAddressesKanaCountry = "ua"
	V2CorePersonScriptAddressesKanaCountryUg V2CorePersonScriptAddressesKanaCountry = "ug"
	V2CorePersonScriptAddressesKanaCountryUm V2CorePersonScriptAddressesKanaCountry = "um"
	V2CorePersonScriptAddressesKanaCountryUS V2CorePersonScriptAddressesKanaCountry = "us"
	V2CorePersonScriptAddressesKanaCountryUy V2CorePersonScriptAddressesKanaCountry = "uy"
	V2CorePersonScriptAddressesKanaCountryUz V2CorePersonScriptAddressesKanaCountry = "uz"
	V2CorePersonScriptAddressesKanaCountryVa V2CorePersonScriptAddressesKanaCountry = "va"
	V2CorePersonScriptAddressesKanaCountryVc V2CorePersonScriptAddressesKanaCountry = "vc"
	V2CorePersonScriptAddressesKanaCountryVe V2CorePersonScriptAddressesKanaCountry = "ve"
	V2CorePersonScriptAddressesKanaCountryVg V2CorePersonScriptAddressesKanaCountry = "vg"
	V2CorePersonScriptAddressesKanaCountryVI V2CorePersonScriptAddressesKanaCountry = "vi"
	V2CorePersonScriptAddressesKanaCountryVn V2CorePersonScriptAddressesKanaCountry = "vn"
	V2CorePersonScriptAddressesKanaCountryVu V2CorePersonScriptAddressesKanaCountry = "vu"
	V2CorePersonScriptAddressesKanaCountryWf V2CorePersonScriptAddressesKanaCountry = "wf"
	V2CorePersonScriptAddressesKanaCountryWs V2CorePersonScriptAddressesKanaCountry = "ws"
	V2CorePersonScriptAddressesKanaCountryXx V2CorePersonScriptAddressesKanaCountry = "xx"
	V2CorePersonScriptAddressesKanaCountryYe V2CorePersonScriptAddressesKanaCountry = "ye"
	V2CorePersonScriptAddressesKanaCountryYt V2CorePersonScriptAddressesKanaCountry = "yt"
	V2CorePersonScriptAddressesKanaCountryZa V2CorePersonScriptAddressesKanaCountry = "za"
	V2CorePersonScriptAddressesKanaCountryZm V2CorePersonScriptAddressesKanaCountry = "zm"
	V2CorePersonScriptAddressesKanaCountryZw V2CorePersonScriptAddressesKanaCountry = "zw"
)

// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CorePersonScriptAddressesKanjiCountry string

// List of values that V2CorePersonScriptAddressesKanjiCountry can take
const (
	V2CorePersonScriptAddressesKanjiCountryAd V2CorePersonScriptAddressesKanjiCountry = "ad"
	V2CorePersonScriptAddressesKanjiCountryAe V2CorePersonScriptAddressesKanjiCountry = "ae"
	V2CorePersonScriptAddressesKanjiCountryAf V2CorePersonScriptAddressesKanjiCountry = "af"
	V2CorePersonScriptAddressesKanjiCountryAg V2CorePersonScriptAddressesKanjiCountry = "ag"
	V2CorePersonScriptAddressesKanjiCountryAi V2CorePersonScriptAddressesKanjiCountry = "ai"
	V2CorePersonScriptAddressesKanjiCountryAl V2CorePersonScriptAddressesKanjiCountry = "al"
	V2CorePersonScriptAddressesKanjiCountryAm V2CorePersonScriptAddressesKanjiCountry = "am"
	V2CorePersonScriptAddressesKanjiCountryAo V2CorePersonScriptAddressesKanjiCountry = "ao"
	V2CorePersonScriptAddressesKanjiCountryAq V2CorePersonScriptAddressesKanjiCountry = "aq"
	V2CorePersonScriptAddressesKanjiCountryAR V2CorePersonScriptAddressesKanjiCountry = "ar"
	V2CorePersonScriptAddressesKanjiCountryAs V2CorePersonScriptAddressesKanjiCountry = "as"
	V2CorePersonScriptAddressesKanjiCountryAt V2CorePersonScriptAddressesKanjiCountry = "at"
	V2CorePersonScriptAddressesKanjiCountryAu V2CorePersonScriptAddressesKanjiCountry = "au"
	V2CorePersonScriptAddressesKanjiCountryAw V2CorePersonScriptAddressesKanjiCountry = "aw"
	V2CorePersonScriptAddressesKanjiCountryAx V2CorePersonScriptAddressesKanjiCountry = "ax"
	V2CorePersonScriptAddressesKanjiCountryAz V2CorePersonScriptAddressesKanjiCountry = "az"
	V2CorePersonScriptAddressesKanjiCountryBa V2CorePersonScriptAddressesKanjiCountry = "ba"
	V2CorePersonScriptAddressesKanjiCountryBb V2CorePersonScriptAddressesKanjiCountry = "bb"
	V2CorePersonScriptAddressesKanjiCountryBd V2CorePersonScriptAddressesKanjiCountry = "bd"
	V2CorePersonScriptAddressesKanjiCountryBe V2CorePersonScriptAddressesKanjiCountry = "be"
	V2CorePersonScriptAddressesKanjiCountryBf V2CorePersonScriptAddressesKanjiCountry = "bf"
	V2CorePersonScriptAddressesKanjiCountryBG V2CorePersonScriptAddressesKanjiCountry = "bg"
	V2CorePersonScriptAddressesKanjiCountryBh V2CorePersonScriptAddressesKanjiCountry = "bh"
	V2CorePersonScriptAddressesKanjiCountryBi V2CorePersonScriptAddressesKanjiCountry = "bi"
	V2CorePersonScriptAddressesKanjiCountryBj V2CorePersonScriptAddressesKanjiCountry = "bj"
	V2CorePersonScriptAddressesKanjiCountryBl V2CorePersonScriptAddressesKanjiCountry = "bl"
	V2CorePersonScriptAddressesKanjiCountryBm V2CorePersonScriptAddressesKanjiCountry = "bm"
	V2CorePersonScriptAddressesKanjiCountryBn V2CorePersonScriptAddressesKanjiCountry = "bn"
	V2CorePersonScriptAddressesKanjiCountryBo V2CorePersonScriptAddressesKanjiCountry = "bo"
	V2CorePersonScriptAddressesKanjiCountryBq V2CorePersonScriptAddressesKanjiCountry = "bq"
	V2CorePersonScriptAddressesKanjiCountryBr V2CorePersonScriptAddressesKanjiCountry = "br"
	V2CorePersonScriptAddressesKanjiCountryBs V2CorePersonScriptAddressesKanjiCountry = "bs"
	V2CorePersonScriptAddressesKanjiCountryBt V2CorePersonScriptAddressesKanjiCountry = "bt"
	V2CorePersonScriptAddressesKanjiCountryBv V2CorePersonScriptAddressesKanjiCountry = "bv"
	V2CorePersonScriptAddressesKanjiCountryBw V2CorePersonScriptAddressesKanjiCountry = "bw"
	V2CorePersonScriptAddressesKanjiCountryBy V2CorePersonScriptAddressesKanjiCountry = "by"
	V2CorePersonScriptAddressesKanjiCountryBz V2CorePersonScriptAddressesKanjiCountry = "bz"
	V2CorePersonScriptAddressesKanjiCountryCa V2CorePersonScriptAddressesKanjiCountry = "ca"
	V2CorePersonScriptAddressesKanjiCountryCc V2CorePersonScriptAddressesKanjiCountry = "cc"
	V2CorePersonScriptAddressesKanjiCountryCd V2CorePersonScriptAddressesKanjiCountry = "cd"
	V2CorePersonScriptAddressesKanjiCountryCf V2CorePersonScriptAddressesKanjiCountry = "cf"
	V2CorePersonScriptAddressesKanjiCountryCg V2CorePersonScriptAddressesKanjiCountry = "cg"
	V2CorePersonScriptAddressesKanjiCountryCh V2CorePersonScriptAddressesKanjiCountry = "ch"
	V2CorePersonScriptAddressesKanjiCountryCi V2CorePersonScriptAddressesKanjiCountry = "ci"
	V2CorePersonScriptAddressesKanjiCountryCk V2CorePersonScriptAddressesKanjiCountry = "ck"
	V2CorePersonScriptAddressesKanjiCountryCl V2CorePersonScriptAddressesKanjiCountry = "cl"
	V2CorePersonScriptAddressesKanjiCountryCm V2CorePersonScriptAddressesKanjiCountry = "cm"
	V2CorePersonScriptAddressesKanjiCountryCn V2CorePersonScriptAddressesKanjiCountry = "cn"
	V2CorePersonScriptAddressesKanjiCountryCo V2CorePersonScriptAddressesKanjiCountry = "co"
	V2CorePersonScriptAddressesKanjiCountryCr V2CorePersonScriptAddressesKanjiCountry = "cr"
	V2CorePersonScriptAddressesKanjiCountryCu V2CorePersonScriptAddressesKanjiCountry = "cu"
	V2CorePersonScriptAddressesKanjiCountryCv V2CorePersonScriptAddressesKanjiCountry = "cv"
	V2CorePersonScriptAddressesKanjiCountryCw V2CorePersonScriptAddressesKanjiCountry = "cw"
	V2CorePersonScriptAddressesKanjiCountryCx V2CorePersonScriptAddressesKanjiCountry = "cx"
	V2CorePersonScriptAddressesKanjiCountryCy V2CorePersonScriptAddressesKanjiCountry = "cy"
	V2CorePersonScriptAddressesKanjiCountryCz V2CorePersonScriptAddressesKanjiCountry = "cz"
	V2CorePersonScriptAddressesKanjiCountryDE V2CorePersonScriptAddressesKanjiCountry = "de"
	V2CorePersonScriptAddressesKanjiCountryDj V2CorePersonScriptAddressesKanjiCountry = "dj"
	V2CorePersonScriptAddressesKanjiCountryDk V2CorePersonScriptAddressesKanjiCountry = "dk"
	V2CorePersonScriptAddressesKanjiCountryDm V2CorePersonScriptAddressesKanjiCountry = "dm"
	V2CorePersonScriptAddressesKanjiCountryDo V2CorePersonScriptAddressesKanjiCountry = "do"
	V2CorePersonScriptAddressesKanjiCountryDz V2CorePersonScriptAddressesKanjiCountry = "dz"
	V2CorePersonScriptAddressesKanjiCountryEc V2CorePersonScriptAddressesKanjiCountry = "ec"
	V2CorePersonScriptAddressesKanjiCountryEe V2CorePersonScriptAddressesKanjiCountry = "ee"
	V2CorePersonScriptAddressesKanjiCountryEg V2CorePersonScriptAddressesKanjiCountry = "eg"
	V2CorePersonScriptAddressesKanjiCountryEh V2CorePersonScriptAddressesKanjiCountry = "eh"
	V2CorePersonScriptAddressesKanjiCountryEr V2CorePersonScriptAddressesKanjiCountry = "er"
	V2CorePersonScriptAddressesKanjiCountryES V2CorePersonScriptAddressesKanjiCountry = "es"
	V2CorePersonScriptAddressesKanjiCountryET V2CorePersonScriptAddressesKanjiCountry = "et"
	V2CorePersonScriptAddressesKanjiCountryFI V2CorePersonScriptAddressesKanjiCountry = "fi"
	V2CorePersonScriptAddressesKanjiCountryFj V2CorePersonScriptAddressesKanjiCountry = "fj"
	V2CorePersonScriptAddressesKanjiCountryFk V2CorePersonScriptAddressesKanjiCountry = "fk"
	V2CorePersonScriptAddressesKanjiCountryFm V2CorePersonScriptAddressesKanjiCountry = "fm"
	V2CorePersonScriptAddressesKanjiCountryFo V2CorePersonScriptAddressesKanjiCountry = "fo"
	V2CorePersonScriptAddressesKanjiCountryFR V2CorePersonScriptAddressesKanjiCountry = "fr"
	V2CorePersonScriptAddressesKanjiCountryGa V2CorePersonScriptAddressesKanjiCountry = "ga"
	V2CorePersonScriptAddressesKanjiCountryGB V2CorePersonScriptAddressesKanjiCountry = "gb"
	V2CorePersonScriptAddressesKanjiCountryGd V2CorePersonScriptAddressesKanjiCountry = "gd"
	V2CorePersonScriptAddressesKanjiCountryGe V2CorePersonScriptAddressesKanjiCountry = "ge"
	V2CorePersonScriptAddressesKanjiCountryGf V2CorePersonScriptAddressesKanjiCountry = "gf"
	V2CorePersonScriptAddressesKanjiCountryGg V2CorePersonScriptAddressesKanjiCountry = "gg"
	V2CorePersonScriptAddressesKanjiCountryGh V2CorePersonScriptAddressesKanjiCountry = "gh"
	V2CorePersonScriptAddressesKanjiCountryGi V2CorePersonScriptAddressesKanjiCountry = "gi"
	V2CorePersonScriptAddressesKanjiCountryGl V2CorePersonScriptAddressesKanjiCountry = "gl"
	V2CorePersonScriptAddressesKanjiCountryGm V2CorePersonScriptAddressesKanjiCountry = "gm"
	V2CorePersonScriptAddressesKanjiCountryGn V2CorePersonScriptAddressesKanjiCountry = "gn"
	V2CorePersonScriptAddressesKanjiCountryGp V2CorePersonScriptAddressesKanjiCountry = "gp"
	V2CorePersonScriptAddressesKanjiCountryGq V2CorePersonScriptAddressesKanjiCountry = "gq"
	V2CorePersonScriptAddressesKanjiCountryGr V2CorePersonScriptAddressesKanjiCountry = "gr"
	V2CorePersonScriptAddressesKanjiCountryGs V2CorePersonScriptAddressesKanjiCountry = "gs"
	V2CorePersonScriptAddressesKanjiCountryGt V2CorePersonScriptAddressesKanjiCountry = "gt"
	V2CorePersonScriptAddressesKanjiCountryGu V2CorePersonScriptAddressesKanjiCountry = "gu"
	V2CorePersonScriptAddressesKanjiCountryGw V2CorePersonScriptAddressesKanjiCountry = "gw"
	V2CorePersonScriptAddressesKanjiCountryGy V2CorePersonScriptAddressesKanjiCountry = "gy"
	V2CorePersonScriptAddressesKanjiCountryHk V2CorePersonScriptAddressesKanjiCountry = "hk"
	V2CorePersonScriptAddressesKanjiCountryHm V2CorePersonScriptAddressesKanjiCountry = "hm"
	V2CorePersonScriptAddressesKanjiCountryHn V2CorePersonScriptAddressesKanjiCountry = "hn"
	V2CorePersonScriptAddressesKanjiCountryHR V2CorePersonScriptAddressesKanjiCountry = "hr"
	V2CorePersonScriptAddressesKanjiCountryHt V2CorePersonScriptAddressesKanjiCountry = "ht"
	V2CorePersonScriptAddressesKanjiCountryHU V2CorePersonScriptAddressesKanjiCountry = "hu"
	V2CorePersonScriptAddressesKanjiCountryID V2CorePersonScriptAddressesKanjiCountry = "id"
	V2CorePersonScriptAddressesKanjiCountryIe V2CorePersonScriptAddressesKanjiCountry = "ie"
	V2CorePersonScriptAddressesKanjiCountryIl V2CorePersonScriptAddressesKanjiCountry = "il"
	V2CorePersonScriptAddressesKanjiCountryIm V2CorePersonScriptAddressesKanjiCountry = "im"
	V2CorePersonScriptAddressesKanjiCountryIn V2CorePersonScriptAddressesKanjiCountry = "in"
	V2CorePersonScriptAddressesKanjiCountryIo V2CorePersonScriptAddressesKanjiCountry = "io"
	V2CorePersonScriptAddressesKanjiCountryIq V2CorePersonScriptAddressesKanjiCountry = "iq"
	V2CorePersonScriptAddressesKanjiCountryIr V2CorePersonScriptAddressesKanjiCountry = "ir"
	V2CorePersonScriptAddressesKanjiCountryIs V2CorePersonScriptAddressesKanjiCountry = "is"
	V2CorePersonScriptAddressesKanjiCountryIT V2CorePersonScriptAddressesKanjiCountry = "it"
	V2CorePersonScriptAddressesKanjiCountryJe V2CorePersonScriptAddressesKanjiCountry = "je"
	V2CorePersonScriptAddressesKanjiCountryJm V2CorePersonScriptAddressesKanjiCountry = "jm"
	V2CorePersonScriptAddressesKanjiCountryJo V2CorePersonScriptAddressesKanjiCountry = "jo"
	V2CorePersonScriptAddressesKanjiCountryJP V2CorePersonScriptAddressesKanjiCountry = "jp"
	V2CorePersonScriptAddressesKanjiCountryKe V2CorePersonScriptAddressesKanjiCountry = "ke"
	V2CorePersonScriptAddressesKanjiCountryKg V2CorePersonScriptAddressesKanjiCountry = "kg"
	V2CorePersonScriptAddressesKanjiCountryKh V2CorePersonScriptAddressesKanjiCountry = "kh"
	V2CorePersonScriptAddressesKanjiCountryKi V2CorePersonScriptAddressesKanjiCountry = "ki"
	V2CorePersonScriptAddressesKanjiCountryKm V2CorePersonScriptAddressesKanjiCountry = "km"
	V2CorePersonScriptAddressesKanjiCountryKn V2CorePersonScriptAddressesKanjiCountry = "kn"
	V2CorePersonScriptAddressesKanjiCountryKp V2CorePersonScriptAddressesKanjiCountry = "kp"
	V2CorePersonScriptAddressesKanjiCountryKr V2CorePersonScriptAddressesKanjiCountry = "kr"
	V2CorePersonScriptAddressesKanjiCountryKw V2CorePersonScriptAddressesKanjiCountry = "kw"
	V2CorePersonScriptAddressesKanjiCountryKy V2CorePersonScriptAddressesKanjiCountry = "ky"
	V2CorePersonScriptAddressesKanjiCountryKz V2CorePersonScriptAddressesKanjiCountry = "kz"
	V2CorePersonScriptAddressesKanjiCountryLa V2CorePersonScriptAddressesKanjiCountry = "la"
	V2CorePersonScriptAddressesKanjiCountryLb V2CorePersonScriptAddressesKanjiCountry = "lb"
	V2CorePersonScriptAddressesKanjiCountryLc V2CorePersonScriptAddressesKanjiCountry = "lc"
	V2CorePersonScriptAddressesKanjiCountryLi V2CorePersonScriptAddressesKanjiCountry = "li"
	V2CorePersonScriptAddressesKanjiCountryLk V2CorePersonScriptAddressesKanjiCountry = "lk"
	V2CorePersonScriptAddressesKanjiCountryLr V2CorePersonScriptAddressesKanjiCountry = "lr"
	V2CorePersonScriptAddressesKanjiCountryLs V2CorePersonScriptAddressesKanjiCountry = "ls"
	V2CorePersonScriptAddressesKanjiCountryLT V2CorePersonScriptAddressesKanjiCountry = "lt"
	V2CorePersonScriptAddressesKanjiCountryLu V2CorePersonScriptAddressesKanjiCountry = "lu"
	V2CorePersonScriptAddressesKanjiCountryLV V2CorePersonScriptAddressesKanjiCountry = "lv"
	V2CorePersonScriptAddressesKanjiCountryLy V2CorePersonScriptAddressesKanjiCountry = "ly"
	V2CorePersonScriptAddressesKanjiCountryMa V2CorePersonScriptAddressesKanjiCountry = "ma"
	V2CorePersonScriptAddressesKanjiCountryMc V2CorePersonScriptAddressesKanjiCountry = "mc"
	V2CorePersonScriptAddressesKanjiCountryMd V2CorePersonScriptAddressesKanjiCountry = "md"
	V2CorePersonScriptAddressesKanjiCountryMe V2CorePersonScriptAddressesKanjiCountry = "me"
	V2CorePersonScriptAddressesKanjiCountryMf V2CorePersonScriptAddressesKanjiCountry = "mf"
	V2CorePersonScriptAddressesKanjiCountryMg V2CorePersonScriptAddressesKanjiCountry = "mg"
	V2CorePersonScriptAddressesKanjiCountryMh V2CorePersonScriptAddressesKanjiCountry = "mh"
	V2CorePersonScriptAddressesKanjiCountryMk V2CorePersonScriptAddressesKanjiCountry = "mk"
	V2CorePersonScriptAddressesKanjiCountryMl V2CorePersonScriptAddressesKanjiCountry = "ml"
	V2CorePersonScriptAddressesKanjiCountryMm V2CorePersonScriptAddressesKanjiCountry = "mm"
	V2CorePersonScriptAddressesKanjiCountryMn V2CorePersonScriptAddressesKanjiCountry = "mn"
	V2CorePersonScriptAddressesKanjiCountryMo V2CorePersonScriptAddressesKanjiCountry = "mo"
	V2CorePersonScriptAddressesKanjiCountryMp V2CorePersonScriptAddressesKanjiCountry = "mp"
	V2CorePersonScriptAddressesKanjiCountryMq V2CorePersonScriptAddressesKanjiCountry = "mq"
	V2CorePersonScriptAddressesKanjiCountryMr V2CorePersonScriptAddressesKanjiCountry = "mr"
	V2CorePersonScriptAddressesKanjiCountryMS V2CorePersonScriptAddressesKanjiCountry = "ms"
	V2CorePersonScriptAddressesKanjiCountryMT V2CorePersonScriptAddressesKanjiCountry = "mt"
	V2CorePersonScriptAddressesKanjiCountryMu V2CorePersonScriptAddressesKanjiCountry = "mu"
	V2CorePersonScriptAddressesKanjiCountryMv V2CorePersonScriptAddressesKanjiCountry = "mv"
	V2CorePersonScriptAddressesKanjiCountryMw V2CorePersonScriptAddressesKanjiCountry = "mw"
	V2CorePersonScriptAddressesKanjiCountryMX V2CorePersonScriptAddressesKanjiCountry = "mx"
	V2CorePersonScriptAddressesKanjiCountryMy V2CorePersonScriptAddressesKanjiCountry = "my"
	V2CorePersonScriptAddressesKanjiCountryMz V2CorePersonScriptAddressesKanjiCountry = "mz"
	V2CorePersonScriptAddressesKanjiCountryNa V2CorePersonScriptAddressesKanjiCountry = "na"
	V2CorePersonScriptAddressesKanjiCountryNc V2CorePersonScriptAddressesKanjiCountry = "nc"
	V2CorePersonScriptAddressesKanjiCountryNe V2CorePersonScriptAddressesKanjiCountry = "ne"
	V2CorePersonScriptAddressesKanjiCountryNf V2CorePersonScriptAddressesKanjiCountry = "nf"
	V2CorePersonScriptAddressesKanjiCountryNg V2CorePersonScriptAddressesKanjiCountry = "ng"
	V2CorePersonScriptAddressesKanjiCountryNi V2CorePersonScriptAddressesKanjiCountry = "ni"
	V2CorePersonScriptAddressesKanjiCountryNL V2CorePersonScriptAddressesKanjiCountry = "nl"
	V2CorePersonScriptAddressesKanjiCountryNo V2CorePersonScriptAddressesKanjiCountry = "no"
	V2CorePersonScriptAddressesKanjiCountryNp V2CorePersonScriptAddressesKanjiCountry = "np"
	V2CorePersonScriptAddressesKanjiCountryNr V2CorePersonScriptAddressesKanjiCountry = "nr"
	V2CorePersonScriptAddressesKanjiCountryNu V2CorePersonScriptAddressesKanjiCountry = "nu"
	V2CorePersonScriptAddressesKanjiCountryNz V2CorePersonScriptAddressesKanjiCountry = "nz"
	V2CorePersonScriptAddressesKanjiCountryOm V2CorePersonScriptAddressesKanjiCountry = "om"
	V2CorePersonScriptAddressesKanjiCountryPa V2CorePersonScriptAddressesKanjiCountry = "pa"
	V2CorePersonScriptAddressesKanjiCountryPe V2CorePersonScriptAddressesKanjiCountry = "pe"
	V2CorePersonScriptAddressesKanjiCountryPf V2CorePersonScriptAddressesKanjiCountry = "pf"
	V2CorePersonScriptAddressesKanjiCountryPg V2CorePersonScriptAddressesKanjiCountry = "pg"
	V2CorePersonScriptAddressesKanjiCountryPh V2CorePersonScriptAddressesKanjiCountry = "ph"
	V2CorePersonScriptAddressesKanjiCountryPk V2CorePersonScriptAddressesKanjiCountry = "pk"
	V2CorePersonScriptAddressesKanjiCountryPL V2CorePersonScriptAddressesKanjiCountry = "pl"
	V2CorePersonScriptAddressesKanjiCountryPm V2CorePersonScriptAddressesKanjiCountry = "pm"
	V2CorePersonScriptAddressesKanjiCountryPn V2CorePersonScriptAddressesKanjiCountry = "pn"
	V2CorePersonScriptAddressesKanjiCountryPr V2CorePersonScriptAddressesKanjiCountry = "pr"
	V2CorePersonScriptAddressesKanjiCountryPs V2CorePersonScriptAddressesKanjiCountry = "ps"
	V2CorePersonScriptAddressesKanjiCountryPT V2CorePersonScriptAddressesKanjiCountry = "pt"
	V2CorePersonScriptAddressesKanjiCountryPw V2CorePersonScriptAddressesKanjiCountry = "pw"
	V2CorePersonScriptAddressesKanjiCountryPy V2CorePersonScriptAddressesKanjiCountry = "py"
	V2CorePersonScriptAddressesKanjiCountryQa V2CorePersonScriptAddressesKanjiCountry = "qa"
	V2CorePersonScriptAddressesKanjiCountryQz V2CorePersonScriptAddressesKanjiCountry = "qz"
	V2CorePersonScriptAddressesKanjiCountryRe V2CorePersonScriptAddressesKanjiCountry = "re"
	V2CorePersonScriptAddressesKanjiCountryRO V2CorePersonScriptAddressesKanjiCountry = "ro"
	V2CorePersonScriptAddressesKanjiCountryRs V2CorePersonScriptAddressesKanjiCountry = "rs"
	V2CorePersonScriptAddressesKanjiCountryRU V2CorePersonScriptAddressesKanjiCountry = "ru"
	V2CorePersonScriptAddressesKanjiCountryRw V2CorePersonScriptAddressesKanjiCountry = "rw"
	V2CorePersonScriptAddressesKanjiCountrySa V2CorePersonScriptAddressesKanjiCountry = "sa"
	V2CorePersonScriptAddressesKanjiCountrySb V2CorePersonScriptAddressesKanjiCountry = "sb"
	V2CorePersonScriptAddressesKanjiCountrySc V2CorePersonScriptAddressesKanjiCountry = "sc"
	V2CorePersonScriptAddressesKanjiCountrySd V2CorePersonScriptAddressesKanjiCountry = "sd"
	V2CorePersonScriptAddressesKanjiCountrySe V2CorePersonScriptAddressesKanjiCountry = "se"
	V2CorePersonScriptAddressesKanjiCountrySg V2CorePersonScriptAddressesKanjiCountry = "sg"
	V2CorePersonScriptAddressesKanjiCountrySh V2CorePersonScriptAddressesKanjiCountry = "sh"
	V2CorePersonScriptAddressesKanjiCountrySi V2CorePersonScriptAddressesKanjiCountry = "si"
	V2CorePersonScriptAddressesKanjiCountrySj V2CorePersonScriptAddressesKanjiCountry = "sj"
	V2CorePersonScriptAddressesKanjiCountrySK V2CorePersonScriptAddressesKanjiCountry = "sk"
	V2CorePersonScriptAddressesKanjiCountrySL V2CorePersonScriptAddressesKanjiCountry = "sl"
	V2CorePersonScriptAddressesKanjiCountrySm V2CorePersonScriptAddressesKanjiCountry = "sm"
	V2CorePersonScriptAddressesKanjiCountrySn V2CorePersonScriptAddressesKanjiCountry = "sn"
	V2CorePersonScriptAddressesKanjiCountrySo V2CorePersonScriptAddressesKanjiCountry = "so"
	V2CorePersonScriptAddressesKanjiCountrySr V2CorePersonScriptAddressesKanjiCountry = "sr"
	V2CorePersonScriptAddressesKanjiCountrySs V2CorePersonScriptAddressesKanjiCountry = "ss"
	V2CorePersonScriptAddressesKanjiCountrySt V2CorePersonScriptAddressesKanjiCountry = "st"
	V2CorePersonScriptAddressesKanjiCountrySV V2CorePersonScriptAddressesKanjiCountry = "sv"
	V2CorePersonScriptAddressesKanjiCountrySx V2CorePersonScriptAddressesKanjiCountry = "sx"
	V2CorePersonScriptAddressesKanjiCountrySy V2CorePersonScriptAddressesKanjiCountry = "sy"
	V2CorePersonScriptAddressesKanjiCountrySz V2CorePersonScriptAddressesKanjiCountry = "sz"
	V2CorePersonScriptAddressesKanjiCountryTc V2CorePersonScriptAddressesKanjiCountry = "tc"
	V2CorePersonScriptAddressesKanjiCountryTd V2CorePersonScriptAddressesKanjiCountry = "td"
	V2CorePersonScriptAddressesKanjiCountryTf V2CorePersonScriptAddressesKanjiCountry = "tf"
	V2CorePersonScriptAddressesKanjiCountryTg V2CorePersonScriptAddressesKanjiCountry = "tg"
	V2CorePersonScriptAddressesKanjiCountryTH V2CorePersonScriptAddressesKanjiCountry = "th"
	V2CorePersonScriptAddressesKanjiCountryTj V2CorePersonScriptAddressesKanjiCountry = "tj"
	V2CorePersonScriptAddressesKanjiCountryTk V2CorePersonScriptAddressesKanjiCountry = "tk"
	V2CorePersonScriptAddressesKanjiCountryTl V2CorePersonScriptAddressesKanjiCountry = "tl"
	V2CorePersonScriptAddressesKanjiCountryTm V2CorePersonScriptAddressesKanjiCountry = "tm"
	V2CorePersonScriptAddressesKanjiCountryTn V2CorePersonScriptAddressesKanjiCountry = "tn"
	V2CorePersonScriptAddressesKanjiCountryTo V2CorePersonScriptAddressesKanjiCountry = "to"
	V2CorePersonScriptAddressesKanjiCountryTR V2CorePersonScriptAddressesKanjiCountry = "tr"
	V2CorePersonScriptAddressesKanjiCountryTt V2CorePersonScriptAddressesKanjiCountry = "tt"
	V2CorePersonScriptAddressesKanjiCountryTv V2CorePersonScriptAddressesKanjiCountry = "tv"
	V2CorePersonScriptAddressesKanjiCountryTw V2CorePersonScriptAddressesKanjiCountry = "tw"
	V2CorePersonScriptAddressesKanjiCountryTz V2CorePersonScriptAddressesKanjiCountry = "tz"
	V2CorePersonScriptAddressesKanjiCountryUa V2CorePersonScriptAddressesKanjiCountry = "ua"
	V2CorePersonScriptAddressesKanjiCountryUg V2CorePersonScriptAddressesKanjiCountry = "ug"
	V2CorePersonScriptAddressesKanjiCountryUm V2CorePersonScriptAddressesKanjiCountry = "um"
	V2CorePersonScriptAddressesKanjiCountryUS V2CorePersonScriptAddressesKanjiCountry = "us"
	V2CorePersonScriptAddressesKanjiCountryUy V2CorePersonScriptAddressesKanjiCountry = "uy"
	V2CorePersonScriptAddressesKanjiCountryUz V2CorePersonScriptAddressesKanjiCountry = "uz"
	V2CorePersonScriptAddressesKanjiCountryVa V2CorePersonScriptAddressesKanjiCountry = "va"
	V2CorePersonScriptAddressesKanjiCountryVc V2CorePersonScriptAddressesKanjiCountry = "vc"
	V2CorePersonScriptAddressesKanjiCountryVe V2CorePersonScriptAddressesKanjiCountry = "ve"
	V2CorePersonScriptAddressesKanjiCountryVg V2CorePersonScriptAddressesKanjiCountry = "vg"
	V2CorePersonScriptAddressesKanjiCountryVI V2CorePersonScriptAddressesKanjiCountry = "vi"
	V2CorePersonScriptAddressesKanjiCountryVn V2CorePersonScriptAddressesKanjiCountry = "vn"
	V2CorePersonScriptAddressesKanjiCountryVu V2CorePersonScriptAddressesKanjiCountry = "vu"
	V2CorePersonScriptAddressesKanjiCountryWf V2CorePersonScriptAddressesKanjiCountry = "wf"
	V2CorePersonScriptAddressesKanjiCountryWs V2CorePersonScriptAddressesKanjiCountry = "ws"
	V2CorePersonScriptAddressesKanjiCountryXx V2CorePersonScriptAddressesKanjiCountry = "xx"
	V2CorePersonScriptAddressesKanjiCountryYe V2CorePersonScriptAddressesKanjiCountry = "ye"
	V2CorePersonScriptAddressesKanjiCountryYt V2CorePersonScriptAddressesKanjiCountry = "yt"
	V2CorePersonScriptAddressesKanjiCountryZa V2CorePersonScriptAddressesKanjiCountry = "za"
	V2CorePersonScriptAddressesKanjiCountryZm V2CorePersonScriptAddressesKanjiCountry = "zm"
	V2CorePersonScriptAddressesKanjiCountryZw V2CorePersonScriptAddressesKanjiCountry = "zw"
)

// Additional addresses associated with the person.
type V2CorePersonAdditionalAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country V2CorePersonAdditionalAddressCountry `json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// Purpose of additional address.
	Purpose V2CorePersonAdditionalAddressPurpose `json:"purpose"`
	// State, county, province, or region.
	State string `json:"state"`
	// Town or cho-me.
	Town string `json:"town"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CorePersonAdditionalName struct {
	// The individual's full name.
	FullName string `json:"full_name"`
	// The individual's first or given name.
	GivenName string `json:"given_name"`
	// The purpose or type of the additional name.
	Purpose V2CorePersonAdditionalNamePurpose `json:"purpose"`
	// The individual's last or family name.
	Surname string `json:"surname"`
}

// Stripe terms of service agreement.
type V2CorePersonAdditionalTermsOfServiceAccount struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date time.Time `json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP string `json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent string `json:"user_agent"`
}

// Attestations of accepted terms of service agreements.
type V2CorePersonAdditionalTermsOfService struct {
	// Stripe terms of service agreement.
	Account *V2CorePersonAdditionalTermsOfServiceAccount `json:"account"`
}

// The person's residential address.
type V2CorePersonAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country V2CorePersonAddressCountry `json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// State, county, province, or region.
	State string `json:"state"`
	// Town or cho-me.
	Town string `json:"town"`
}

// The person's date of birth.
type V2CorePersonDateOfBirth struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CorePersonDocumentsCompanyAuthorization struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CorePersonDocumentsCompanyAuthorizationType `json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CorePersonDocumentsPassport struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CorePersonDocumentsPassportType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
type V2CorePersonDocumentsPrimaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CorePersonDocumentsPrimaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
	FrontBack *V2CorePersonDocumentsPrimaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CorePersonDocumentsPrimaryVerificationType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
type V2CorePersonDocumentsSecondaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CorePersonDocumentsSecondaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
	FrontBack *V2CorePersonDocumentsSecondaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CorePersonDocumentsSecondaryVerificationType `json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CorePersonDocumentsVisa struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CorePersonDocumentsVisaType `json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CorePersonDocuments struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CorePersonDocumentsCompanyAuthorization `json:"company_authorization"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CorePersonDocumentsPassport `json:"passport"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CorePersonDocumentsPrimaryVerification `json:"primary_verification"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CorePersonDocumentsSecondaryVerification `json:"secondary_verification"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CorePersonDocumentsVisa `json:"visa"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CorePersonIDNumber struct {
	// The ID number type of an individual.
	Type V2CorePersonIDNumberType `json:"type"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CorePersonRelationship struct {
	// Whether the individual is an authorizer of the Account's legal entity.
	Authorizer bool `json:"authorizer"`
	// Whether the individual is a director of the Account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director bool `json:"director"`
	// Whether the individual has significant responsibility to control, manage, or direct the organization.
	Executive bool `json:"executive"`
	// Whether the individual is the legal guardian of the Account's representative.
	LegalGuardian bool `json:"legal_guardian"`
	// Whether the individual is an owner of the Account's legal entity.
	Owner bool `json:"owner"`
	// The percent owned by the individual of the Account's legal entity.
	PercentOwnership string `json:"percent_ownership"`
	// Whether the individual is authorized as the primary representative of the Account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.
	Representative bool `json:"representative"`
	// The individual's title (e.g., CEO, Support Engineer).
	Title string `json:"title"`
}

// Kana Address.
type V2CorePersonScriptAddressesKana struct {
	// City, district, suburb, town, or village.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country V2CorePersonScriptAddressesKanaCountry `json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// State, county, province, or region.
	State string `json:"state"`
	// Town or cho-me.
	Town string `json:"town"`
}

// Kanji Address.
type V2CorePersonScriptAddressesKanji struct {
	// City, district, suburb, town, or village.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country V2CorePersonScriptAddressesKanjiCountry `json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// State, county, province, or region.
	State string `json:"state"`
	// Town or cho-me.
	Town string `json:"town"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CorePersonScriptAddresses struct {
	// Kana Address.
	Kana *V2CorePersonScriptAddressesKana `json:"kana"`
	// Kanji Address.
	Kanji *V2CorePersonScriptAddressesKanji `json:"kanji"`
}

// Persons name in kana script.
type V2CorePersonScriptNamesKana struct {
	// The person's first or given name.
	GivenName string `json:"given_name"`
	// The person's last or family name.
	Surname string `json:"surname"`
}

// Persons name in kanji script.
type V2CorePersonScriptNamesKanji struct {
	// The person's first or given name.
	GivenName string `json:"given_name"`
	// The person's last or family name.
	Surname string `json:"surname"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CorePersonScriptNames struct {
	// Persons name in kana script.
	Kana *V2CorePersonScriptNamesKana `json:"kana"`
	// Persons name in kanji script.
	Kanji *V2CorePersonScriptNamesKanji `json:"kanji"`
}

// Person retrieval response schema.
type V2CorePerson struct {
	APIResource
	// The account ID which the individual belongs to.
	Account string `json:"account"`
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CorePersonAdditionalAddress `json:"additional_addresses"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CorePersonAdditionalName `json:"additional_names"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CorePersonAdditionalTermsOfService `json:"additional_terms_of_service"`
	// The person's residential address.
	Address *V2CorePersonAddress `json:"address"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// The person's date of birth.
	DateOfBirth *V2CorePersonDateOfBirth `json:"date_of_birth"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CorePersonDocuments `json:"documents"`
	// The person's email address.
	Email string `json:"email"`
	// The person's first name.
	GivenName string `json:"given_name"`
	// Unique identifier for the Person.
	ID string `json:"id"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CorePersonIDNumber `json:"id_numbers"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender V2CorePersonLegalGender `json:"legal_gender"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The countries where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []V2CorePersonNationality `json:"nationalities"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The person's phone number.
	Phone string `json:"phone"`
	// The person's political exposure.
	PoliticalExposure V2CorePersonPoliticalExposure `json:"political_exposure"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CorePersonRelationship `json:"relationship"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CorePersonScriptAddresses `json:"script_addresses"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CorePersonScriptNames `json:"script_names"`
	// The person's last name.
	Surname string `json:"surname"`
	// Time at which the object was last updated. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Updated time.Time `json:"updated"`
}
