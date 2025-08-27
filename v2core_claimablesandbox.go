//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Country in which the account holder resides, or in which the business is legally established.
// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
type V2CoreClaimableSandboxPrefillCountry string

// List of values that V2CoreClaimableSandboxPrefillCountry can take
const (
	V2CoreClaimableSandboxPrefillCountryAd V2CoreClaimableSandboxPrefillCountry = "ad"
	V2CoreClaimableSandboxPrefillCountryAe V2CoreClaimableSandboxPrefillCountry = "ae"
	V2CoreClaimableSandboxPrefillCountryAf V2CoreClaimableSandboxPrefillCountry = "af"
	V2CoreClaimableSandboxPrefillCountryAg V2CoreClaimableSandboxPrefillCountry = "ag"
	V2CoreClaimableSandboxPrefillCountryAi V2CoreClaimableSandboxPrefillCountry = "ai"
	V2CoreClaimableSandboxPrefillCountryAl V2CoreClaimableSandboxPrefillCountry = "al"
	V2CoreClaimableSandboxPrefillCountryAm V2CoreClaimableSandboxPrefillCountry = "am"
	V2CoreClaimableSandboxPrefillCountryAo V2CoreClaimableSandboxPrefillCountry = "ao"
	V2CoreClaimableSandboxPrefillCountryAq V2CoreClaimableSandboxPrefillCountry = "aq"
	V2CoreClaimableSandboxPrefillCountryAR V2CoreClaimableSandboxPrefillCountry = "ar"
	V2CoreClaimableSandboxPrefillCountryAs V2CoreClaimableSandboxPrefillCountry = "as"
	V2CoreClaimableSandboxPrefillCountryAt V2CoreClaimableSandboxPrefillCountry = "at"
	V2CoreClaimableSandboxPrefillCountryAu V2CoreClaimableSandboxPrefillCountry = "au"
	V2CoreClaimableSandboxPrefillCountryAw V2CoreClaimableSandboxPrefillCountry = "aw"
	V2CoreClaimableSandboxPrefillCountryAx V2CoreClaimableSandboxPrefillCountry = "ax"
	V2CoreClaimableSandboxPrefillCountryAz V2CoreClaimableSandboxPrefillCountry = "az"
	V2CoreClaimableSandboxPrefillCountryBa V2CoreClaimableSandboxPrefillCountry = "ba"
	V2CoreClaimableSandboxPrefillCountryBb V2CoreClaimableSandboxPrefillCountry = "bb"
	V2CoreClaimableSandboxPrefillCountryBd V2CoreClaimableSandboxPrefillCountry = "bd"
	V2CoreClaimableSandboxPrefillCountryBe V2CoreClaimableSandboxPrefillCountry = "be"
	V2CoreClaimableSandboxPrefillCountryBf V2CoreClaimableSandboxPrefillCountry = "bf"
	V2CoreClaimableSandboxPrefillCountryBG V2CoreClaimableSandboxPrefillCountry = "bg"
	V2CoreClaimableSandboxPrefillCountryBh V2CoreClaimableSandboxPrefillCountry = "bh"
	V2CoreClaimableSandboxPrefillCountryBi V2CoreClaimableSandboxPrefillCountry = "bi"
	V2CoreClaimableSandboxPrefillCountryBj V2CoreClaimableSandboxPrefillCountry = "bj"
	V2CoreClaimableSandboxPrefillCountryBl V2CoreClaimableSandboxPrefillCountry = "bl"
	V2CoreClaimableSandboxPrefillCountryBm V2CoreClaimableSandboxPrefillCountry = "bm"
	V2CoreClaimableSandboxPrefillCountryBn V2CoreClaimableSandboxPrefillCountry = "bn"
	V2CoreClaimableSandboxPrefillCountryBo V2CoreClaimableSandboxPrefillCountry = "bo"
	V2CoreClaimableSandboxPrefillCountryBq V2CoreClaimableSandboxPrefillCountry = "bq"
	V2CoreClaimableSandboxPrefillCountryBr V2CoreClaimableSandboxPrefillCountry = "br"
	V2CoreClaimableSandboxPrefillCountryBs V2CoreClaimableSandboxPrefillCountry = "bs"
	V2CoreClaimableSandboxPrefillCountryBt V2CoreClaimableSandboxPrefillCountry = "bt"
	V2CoreClaimableSandboxPrefillCountryBv V2CoreClaimableSandboxPrefillCountry = "bv"
	V2CoreClaimableSandboxPrefillCountryBw V2CoreClaimableSandboxPrefillCountry = "bw"
	V2CoreClaimableSandboxPrefillCountryBy V2CoreClaimableSandboxPrefillCountry = "by"
	V2CoreClaimableSandboxPrefillCountryBz V2CoreClaimableSandboxPrefillCountry = "bz"
	V2CoreClaimableSandboxPrefillCountryCa V2CoreClaimableSandboxPrefillCountry = "ca"
	V2CoreClaimableSandboxPrefillCountryCc V2CoreClaimableSandboxPrefillCountry = "cc"
	V2CoreClaimableSandboxPrefillCountryCd V2CoreClaimableSandboxPrefillCountry = "cd"
	V2CoreClaimableSandboxPrefillCountryCf V2CoreClaimableSandboxPrefillCountry = "cf"
	V2CoreClaimableSandboxPrefillCountryCg V2CoreClaimableSandboxPrefillCountry = "cg"
	V2CoreClaimableSandboxPrefillCountryCh V2CoreClaimableSandboxPrefillCountry = "ch"
	V2CoreClaimableSandboxPrefillCountryCi V2CoreClaimableSandboxPrefillCountry = "ci"
	V2CoreClaimableSandboxPrefillCountryCk V2CoreClaimableSandboxPrefillCountry = "ck"
	V2CoreClaimableSandboxPrefillCountryCl V2CoreClaimableSandboxPrefillCountry = "cl"
	V2CoreClaimableSandboxPrefillCountryCm V2CoreClaimableSandboxPrefillCountry = "cm"
	V2CoreClaimableSandboxPrefillCountryCn V2CoreClaimableSandboxPrefillCountry = "cn"
	V2CoreClaimableSandboxPrefillCountryCo V2CoreClaimableSandboxPrefillCountry = "co"
	V2CoreClaimableSandboxPrefillCountryCr V2CoreClaimableSandboxPrefillCountry = "cr"
	V2CoreClaimableSandboxPrefillCountryCu V2CoreClaimableSandboxPrefillCountry = "cu"
	V2CoreClaimableSandboxPrefillCountryCv V2CoreClaimableSandboxPrefillCountry = "cv"
	V2CoreClaimableSandboxPrefillCountryCw V2CoreClaimableSandboxPrefillCountry = "cw"
	V2CoreClaimableSandboxPrefillCountryCx V2CoreClaimableSandboxPrefillCountry = "cx"
	V2CoreClaimableSandboxPrefillCountryCy V2CoreClaimableSandboxPrefillCountry = "cy"
	V2CoreClaimableSandboxPrefillCountryCz V2CoreClaimableSandboxPrefillCountry = "cz"
	V2CoreClaimableSandboxPrefillCountryDE V2CoreClaimableSandboxPrefillCountry = "de"
	V2CoreClaimableSandboxPrefillCountryDj V2CoreClaimableSandboxPrefillCountry = "dj"
	V2CoreClaimableSandboxPrefillCountryDk V2CoreClaimableSandboxPrefillCountry = "dk"
	V2CoreClaimableSandboxPrefillCountryDm V2CoreClaimableSandboxPrefillCountry = "dm"
	V2CoreClaimableSandboxPrefillCountryDo V2CoreClaimableSandboxPrefillCountry = "do"
	V2CoreClaimableSandboxPrefillCountryDz V2CoreClaimableSandboxPrefillCountry = "dz"
	V2CoreClaimableSandboxPrefillCountryEc V2CoreClaimableSandboxPrefillCountry = "ec"
	V2CoreClaimableSandboxPrefillCountryEe V2CoreClaimableSandboxPrefillCountry = "ee"
	V2CoreClaimableSandboxPrefillCountryEg V2CoreClaimableSandboxPrefillCountry = "eg"
	V2CoreClaimableSandboxPrefillCountryEh V2CoreClaimableSandboxPrefillCountry = "eh"
	V2CoreClaimableSandboxPrefillCountryEr V2CoreClaimableSandboxPrefillCountry = "er"
	V2CoreClaimableSandboxPrefillCountryES V2CoreClaimableSandboxPrefillCountry = "es"
	V2CoreClaimableSandboxPrefillCountryET V2CoreClaimableSandboxPrefillCountry = "et"
	V2CoreClaimableSandboxPrefillCountryFI V2CoreClaimableSandboxPrefillCountry = "fi"
	V2CoreClaimableSandboxPrefillCountryFj V2CoreClaimableSandboxPrefillCountry = "fj"
	V2CoreClaimableSandboxPrefillCountryFk V2CoreClaimableSandboxPrefillCountry = "fk"
	V2CoreClaimableSandboxPrefillCountryFm V2CoreClaimableSandboxPrefillCountry = "fm"
	V2CoreClaimableSandboxPrefillCountryFo V2CoreClaimableSandboxPrefillCountry = "fo"
	V2CoreClaimableSandboxPrefillCountryFR V2CoreClaimableSandboxPrefillCountry = "fr"
	V2CoreClaimableSandboxPrefillCountryGa V2CoreClaimableSandboxPrefillCountry = "ga"
	V2CoreClaimableSandboxPrefillCountryGB V2CoreClaimableSandboxPrefillCountry = "gb"
	V2CoreClaimableSandboxPrefillCountryGd V2CoreClaimableSandboxPrefillCountry = "gd"
	V2CoreClaimableSandboxPrefillCountryGe V2CoreClaimableSandboxPrefillCountry = "ge"
	V2CoreClaimableSandboxPrefillCountryGf V2CoreClaimableSandboxPrefillCountry = "gf"
	V2CoreClaimableSandboxPrefillCountryGg V2CoreClaimableSandboxPrefillCountry = "gg"
	V2CoreClaimableSandboxPrefillCountryGh V2CoreClaimableSandboxPrefillCountry = "gh"
	V2CoreClaimableSandboxPrefillCountryGi V2CoreClaimableSandboxPrefillCountry = "gi"
	V2CoreClaimableSandboxPrefillCountryGl V2CoreClaimableSandboxPrefillCountry = "gl"
	V2CoreClaimableSandboxPrefillCountryGm V2CoreClaimableSandboxPrefillCountry = "gm"
	V2CoreClaimableSandboxPrefillCountryGn V2CoreClaimableSandboxPrefillCountry = "gn"
	V2CoreClaimableSandboxPrefillCountryGp V2CoreClaimableSandboxPrefillCountry = "gp"
	V2CoreClaimableSandboxPrefillCountryGq V2CoreClaimableSandboxPrefillCountry = "gq"
	V2CoreClaimableSandboxPrefillCountryGr V2CoreClaimableSandboxPrefillCountry = "gr"
	V2CoreClaimableSandboxPrefillCountryGs V2CoreClaimableSandboxPrefillCountry = "gs"
	V2CoreClaimableSandboxPrefillCountryGt V2CoreClaimableSandboxPrefillCountry = "gt"
	V2CoreClaimableSandboxPrefillCountryGu V2CoreClaimableSandboxPrefillCountry = "gu"
	V2CoreClaimableSandboxPrefillCountryGw V2CoreClaimableSandboxPrefillCountry = "gw"
	V2CoreClaimableSandboxPrefillCountryGy V2CoreClaimableSandboxPrefillCountry = "gy"
	V2CoreClaimableSandboxPrefillCountryHk V2CoreClaimableSandboxPrefillCountry = "hk"
	V2CoreClaimableSandboxPrefillCountryHm V2CoreClaimableSandboxPrefillCountry = "hm"
	V2CoreClaimableSandboxPrefillCountryHn V2CoreClaimableSandboxPrefillCountry = "hn"
	V2CoreClaimableSandboxPrefillCountryHR V2CoreClaimableSandboxPrefillCountry = "hr"
	V2CoreClaimableSandboxPrefillCountryHt V2CoreClaimableSandboxPrefillCountry = "ht"
	V2CoreClaimableSandboxPrefillCountryHU V2CoreClaimableSandboxPrefillCountry = "hu"
	V2CoreClaimableSandboxPrefillCountryID V2CoreClaimableSandboxPrefillCountry = "id"
	V2CoreClaimableSandboxPrefillCountryIe V2CoreClaimableSandboxPrefillCountry = "ie"
	V2CoreClaimableSandboxPrefillCountryIl V2CoreClaimableSandboxPrefillCountry = "il"
	V2CoreClaimableSandboxPrefillCountryIm V2CoreClaimableSandboxPrefillCountry = "im"
	V2CoreClaimableSandboxPrefillCountryIn V2CoreClaimableSandboxPrefillCountry = "in"
	V2CoreClaimableSandboxPrefillCountryIo V2CoreClaimableSandboxPrefillCountry = "io"
	V2CoreClaimableSandboxPrefillCountryIq V2CoreClaimableSandboxPrefillCountry = "iq"
	V2CoreClaimableSandboxPrefillCountryIr V2CoreClaimableSandboxPrefillCountry = "ir"
	V2CoreClaimableSandboxPrefillCountryIs V2CoreClaimableSandboxPrefillCountry = "is"
	V2CoreClaimableSandboxPrefillCountryIT V2CoreClaimableSandboxPrefillCountry = "it"
	V2CoreClaimableSandboxPrefillCountryJe V2CoreClaimableSandboxPrefillCountry = "je"
	V2CoreClaimableSandboxPrefillCountryJm V2CoreClaimableSandboxPrefillCountry = "jm"
	V2CoreClaimableSandboxPrefillCountryJo V2CoreClaimableSandboxPrefillCountry = "jo"
	V2CoreClaimableSandboxPrefillCountryJP V2CoreClaimableSandboxPrefillCountry = "jp"
	V2CoreClaimableSandboxPrefillCountryKe V2CoreClaimableSandboxPrefillCountry = "ke"
	V2CoreClaimableSandboxPrefillCountryKg V2CoreClaimableSandboxPrefillCountry = "kg"
	V2CoreClaimableSandboxPrefillCountryKh V2CoreClaimableSandboxPrefillCountry = "kh"
	V2CoreClaimableSandboxPrefillCountryKi V2CoreClaimableSandboxPrefillCountry = "ki"
	V2CoreClaimableSandboxPrefillCountryKm V2CoreClaimableSandboxPrefillCountry = "km"
	V2CoreClaimableSandboxPrefillCountryKn V2CoreClaimableSandboxPrefillCountry = "kn"
	V2CoreClaimableSandboxPrefillCountryKp V2CoreClaimableSandboxPrefillCountry = "kp"
	V2CoreClaimableSandboxPrefillCountryKr V2CoreClaimableSandboxPrefillCountry = "kr"
	V2CoreClaimableSandboxPrefillCountryKw V2CoreClaimableSandboxPrefillCountry = "kw"
	V2CoreClaimableSandboxPrefillCountryKy V2CoreClaimableSandboxPrefillCountry = "ky"
	V2CoreClaimableSandboxPrefillCountryKz V2CoreClaimableSandboxPrefillCountry = "kz"
	V2CoreClaimableSandboxPrefillCountryLa V2CoreClaimableSandboxPrefillCountry = "la"
	V2CoreClaimableSandboxPrefillCountryLb V2CoreClaimableSandboxPrefillCountry = "lb"
	V2CoreClaimableSandboxPrefillCountryLc V2CoreClaimableSandboxPrefillCountry = "lc"
	V2CoreClaimableSandboxPrefillCountryLi V2CoreClaimableSandboxPrefillCountry = "li"
	V2CoreClaimableSandboxPrefillCountryLk V2CoreClaimableSandboxPrefillCountry = "lk"
	V2CoreClaimableSandboxPrefillCountryLr V2CoreClaimableSandboxPrefillCountry = "lr"
	V2CoreClaimableSandboxPrefillCountryLs V2CoreClaimableSandboxPrefillCountry = "ls"
	V2CoreClaimableSandboxPrefillCountryLT V2CoreClaimableSandboxPrefillCountry = "lt"
	V2CoreClaimableSandboxPrefillCountryLu V2CoreClaimableSandboxPrefillCountry = "lu"
	V2CoreClaimableSandboxPrefillCountryLV V2CoreClaimableSandboxPrefillCountry = "lv"
	V2CoreClaimableSandboxPrefillCountryLy V2CoreClaimableSandboxPrefillCountry = "ly"
	V2CoreClaimableSandboxPrefillCountryMa V2CoreClaimableSandboxPrefillCountry = "ma"
	V2CoreClaimableSandboxPrefillCountryMc V2CoreClaimableSandboxPrefillCountry = "mc"
	V2CoreClaimableSandboxPrefillCountryMd V2CoreClaimableSandboxPrefillCountry = "md"
	V2CoreClaimableSandboxPrefillCountryMe V2CoreClaimableSandboxPrefillCountry = "me"
	V2CoreClaimableSandboxPrefillCountryMf V2CoreClaimableSandboxPrefillCountry = "mf"
	V2CoreClaimableSandboxPrefillCountryMg V2CoreClaimableSandboxPrefillCountry = "mg"
	V2CoreClaimableSandboxPrefillCountryMh V2CoreClaimableSandboxPrefillCountry = "mh"
	V2CoreClaimableSandboxPrefillCountryMk V2CoreClaimableSandboxPrefillCountry = "mk"
	V2CoreClaimableSandboxPrefillCountryMl V2CoreClaimableSandboxPrefillCountry = "ml"
	V2CoreClaimableSandboxPrefillCountryMm V2CoreClaimableSandboxPrefillCountry = "mm"
	V2CoreClaimableSandboxPrefillCountryMn V2CoreClaimableSandboxPrefillCountry = "mn"
	V2CoreClaimableSandboxPrefillCountryMo V2CoreClaimableSandboxPrefillCountry = "mo"
	V2CoreClaimableSandboxPrefillCountryMp V2CoreClaimableSandboxPrefillCountry = "mp"
	V2CoreClaimableSandboxPrefillCountryMq V2CoreClaimableSandboxPrefillCountry = "mq"
	V2CoreClaimableSandboxPrefillCountryMr V2CoreClaimableSandboxPrefillCountry = "mr"
	V2CoreClaimableSandboxPrefillCountryMS V2CoreClaimableSandboxPrefillCountry = "ms"
	V2CoreClaimableSandboxPrefillCountryMT V2CoreClaimableSandboxPrefillCountry = "mt"
	V2CoreClaimableSandboxPrefillCountryMu V2CoreClaimableSandboxPrefillCountry = "mu"
	V2CoreClaimableSandboxPrefillCountryMv V2CoreClaimableSandboxPrefillCountry = "mv"
	V2CoreClaimableSandboxPrefillCountryMw V2CoreClaimableSandboxPrefillCountry = "mw"
	V2CoreClaimableSandboxPrefillCountryMX V2CoreClaimableSandboxPrefillCountry = "mx"
	V2CoreClaimableSandboxPrefillCountryMy V2CoreClaimableSandboxPrefillCountry = "my"
	V2CoreClaimableSandboxPrefillCountryMz V2CoreClaimableSandboxPrefillCountry = "mz"
	V2CoreClaimableSandboxPrefillCountryNa V2CoreClaimableSandboxPrefillCountry = "na"
	V2CoreClaimableSandboxPrefillCountryNc V2CoreClaimableSandboxPrefillCountry = "nc"
	V2CoreClaimableSandboxPrefillCountryNe V2CoreClaimableSandboxPrefillCountry = "ne"
	V2CoreClaimableSandboxPrefillCountryNf V2CoreClaimableSandboxPrefillCountry = "nf"
	V2CoreClaimableSandboxPrefillCountryNg V2CoreClaimableSandboxPrefillCountry = "ng"
	V2CoreClaimableSandboxPrefillCountryNi V2CoreClaimableSandboxPrefillCountry = "ni"
	V2CoreClaimableSandboxPrefillCountryNL V2CoreClaimableSandboxPrefillCountry = "nl"
	V2CoreClaimableSandboxPrefillCountryNo V2CoreClaimableSandboxPrefillCountry = "no"
	V2CoreClaimableSandboxPrefillCountryNp V2CoreClaimableSandboxPrefillCountry = "np"
	V2CoreClaimableSandboxPrefillCountryNr V2CoreClaimableSandboxPrefillCountry = "nr"
	V2CoreClaimableSandboxPrefillCountryNu V2CoreClaimableSandboxPrefillCountry = "nu"
	V2CoreClaimableSandboxPrefillCountryNz V2CoreClaimableSandboxPrefillCountry = "nz"
	V2CoreClaimableSandboxPrefillCountryOm V2CoreClaimableSandboxPrefillCountry = "om"
	V2CoreClaimableSandboxPrefillCountryPa V2CoreClaimableSandboxPrefillCountry = "pa"
	V2CoreClaimableSandboxPrefillCountryPe V2CoreClaimableSandboxPrefillCountry = "pe"
	V2CoreClaimableSandboxPrefillCountryPf V2CoreClaimableSandboxPrefillCountry = "pf"
	V2CoreClaimableSandboxPrefillCountryPg V2CoreClaimableSandboxPrefillCountry = "pg"
	V2CoreClaimableSandboxPrefillCountryPh V2CoreClaimableSandboxPrefillCountry = "ph"
	V2CoreClaimableSandboxPrefillCountryPk V2CoreClaimableSandboxPrefillCountry = "pk"
	V2CoreClaimableSandboxPrefillCountryPL V2CoreClaimableSandboxPrefillCountry = "pl"
	V2CoreClaimableSandboxPrefillCountryPm V2CoreClaimableSandboxPrefillCountry = "pm"
	V2CoreClaimableSandboxPrefillCountryPn V2CoreClaimableSandboxPrefillCountry = "pn"
	V2CoreClaimableSandboxPrefillCountryPr V2CoreClaimableSandboxPrefillCountry = "pr"
	V2CoreClaimableSandboxPrefillCountryPs V2CoreClaimableSandboxPrefillCountry = "ps"
	V2CoreClaimableSandboxPrefillCountryPT V2CoreClaimableSandboxPrefillCountry = "pt"
	V2CoreClaimableSandboxPrefillCountryPw V2CoreClaimableSandboxPrefillCountry = "pw"
	V2CoreClaimableSandboxPrefillCountryPy V2CoreClaimableSandboxPrefillCountry = "py"
	V2CoreClaimableSandboxPrefillCountryQa V2CoreClaimableSandboxPrefillCountry = "qa"
	V2CoreClaimableSandboxPrefillCountryQz V2CoreClaimableSandboxPrefillCountry = "qz"
	V2CoreClaimableSandboxPrefillCountryRe V2CoreClaimableSandboxPrefillCountry = "re"
	V2CoreClaimableSandboxPrefillCountryRO V2CoreClaimableSandboxPrefillCountry = "ro"
	V2CoreClaimableSandboxPrefillCountryRs V2CoreClaimableSandboxPrefillCountry = "rs"
	V2CoreClaimableSandboxPrefillCountryRU V2CoreClaimableSandboxPrefillCountry = "ru"
	V2CoreClaimableSandboxPrefillCountryRw V2CoreClaimableSandboxPrefillCountry = "rw"
	V2CoreClaimableSandboxPrefillCountrySa V2CoreClaimableSandboxPrefillCountry = "sa"
	V2CoreClaimableSandboxPrefillCountrySb V2CoreClaimableSandboxPrefillCountry = "sb"
	V2CoreClaimableSandboxPrefillCountrySc V2CoreClaimableSandboxPrefillCountry = "sc"
	V2CoreClaimableSandboxPrefillCountrySd V2CoreClaimableSandboxPrefillCountry = "sd"
	V2CoreClaimableSandboxPrefillCountrySe V2CoreClaimableSandboxPrefillCountry = "se"
	V2CoreClaimableSandboxPrefillCountrySg V2CoreClaimableSandboxPrefillCountry = "sg"
	V2CoreClaimableSandboxPrefillCountrySh V2CoreClaimableSandboxPrefillCountry = "sh"
	V2CoreClaimableSandboxPrefillCountrySi V2CoreClaimableSandboxPrefillCountry = "si"
	V2CoreClaimableSandboxPrefillCountrySj V2CoreClaimableSandboxPrefillCountry = "sj"
	V2CoreClaimableSandboxPrefillCountrySK V2CoreClaimableSandboxPrefillCountry = "sk"
	V2CoreClaimableSandboxPrefillCountrySL V2CoreClaimableSandboxPrefillCountry = "sl"
	V2CoreClaimableSandboxPrefillCountrySm V2CoreClaimableSandboxPrefillCountry = "sm"
	V2CoreClaimableSandboxPrefillCountrySn V2CoreClaimableSandboxPrefillCountry = "sn"
	V2CoreClaimableSandboxPrefillCountrySo V2CoreClaimableSandboxPrefillCountry = "so"
	V2CoreClaimableSandboxPrefillCountrySr V2CoreClaimableSandboxPrefillCountry = "sr"
	V2CoreClaimableSandboxPrefillCountrySs V2CoreClaimableSandboxPrefillCountry = "ss"
	V2CoreClaimableSandboxPrefillCountrySt V2CoreClaimableSandboxPrefillCountry = "st"
	V2CoreClaimableSandboxPrefillCountrySV V2CoreClaimableSandboxPrefillCountry = "sv"
	V2CoreClaimableSandboxPrefillCountrySx V2CoreClaimableSandboxPrefillCountry = "sx"
	V2CoreClaimableSandboxPrefillCountrySy V2CoreClaimableSandboxPrefillCountry = "sy"
	V2CoreClaimableSandboxPrefillCountrySz V2CoreClaimableSandboxPrefillCountry = "sz"
	V2CoreClaimableSandboxPrefillCountryTc V2CoreClaimableSandboxPrefillCountry = "tc"
	V2CoreClaimableSandboxPrefillCountryTd V2CoreClaimableSandboxPrefillCountry = "td"
	V2CoreClaimableSandboxPrefillCountryTf V2CoreClaimableSandboxPrefillCountry = "tf"
	V2CoreClaimableSandboxPrefillCountryTg V2CoreClaimableSandboxPrefillCountry = "tg"
	V2CoreClaimableSandboxPrefillCountryTH V2CoreClaimableSandboxPrefillCountry = "th"
	V2CoreClaimableSandboxPrefillCountryTj V2CoreClaimableSandboxPrefillCountry = "tj"
	V2CoreClaimableSandboxPrefillCountryTk V2CoreClaimableSandboxPrefillCountry = "tk"
	V2CoreClaimableSandboxPrefillCountryTl V2CoreClaimableSandboxPrefillCountry = "tl"
	V2CoreClaimableSandboxPrefillCountryTm V2CoreClaimableSandboxPrefillCountry = "tm"
	V2CoreClaimableSandboxPrefillCountryTn V2CoreClaimableSandboxPrefillCountry = "tn"
	V2CoreClaimableSandboxPrefillCountryTo V2CoreClaimableSandboxPrefillCountry = "to"
	V2CoreClaimableSandboxPrefillCountryTR V2CoreClaimableSandboxPrefillCountry = "tr"
	V2CoreClaimableSandboxPrefillCountryTt V2CoreClaimableSandboxPrefillCountry = "tt"
	V2CoreClaimableSandboxPrefillCountryTv V2CoreClaimableSandboxPrefillCountry = "tv"
	V2CoreClaimableSandboxPrefillCountryTw V2CoreClaimableSandboxPrefillCountry = "tw"
	V2CoreClaimableSandboxPrefillCountryTz V2CoreClaimableSandboxPrefillCountry = "tz"
	V2CoreClaimableSandboxPrefillCountryUa V2CoreClaimableSandboxPrefillCountry = "ua"
	V2CoreClaimableSandboxPrefillCountryUg V2CoreClaimableSandboxPrefillCountry = "ug"
	V2CoreClaimableSandboxPrefillCountryUm V2CoreClaimableSandboxPrefillCountry = "um"
	V2CoreClaimableSandboxPrefillCountryUS V2CoreClaimableSandboxPrefillCountry = "us"
	V2CoreClaimableSandboxPrefillCountryUy V2CoreClaimableSandboxPrefillCountry = "uy"
	V2CoreClaimableSandboxPrefillCountryUz V2CoreClaimableSandboxPrefillCountry = "uz"
	V2CoreClaimableSandboxPrefillCountryVa V2CoreClaimableSandboxPrefillCountry = "va"
	V2CoreClaimableSandboxPrefillCountryVc V2CoreClaimableSandboxPrefillCountry = "vc"
	V2CoreClaimableSandboxPrefillCountryVe V2CoreClaimableSandboxPrefillCountry = "ve"
	V2CoreClaimableSandboxPrefillCountryVg V2CoreClaimableSandboxPrefillCountry = "vg"
	V2CoreClaimableSandboxPrefillCountryVI V2CoreClaimableSandboxPrefillCountry = "vi"
	V2CoreClaimableSandboxPrefillCountryVn V2CoreClaimableSandboxPrefillCountry = "vn"
	V2CoreClaimableSandboxPrefillCountryVu V2CoreClaimableSandboxPrefillCountry = "vu"
	V2CoreClaimableSandboxPrefillCountryWf V2CoreClaimableSandboxPrefillCountry = "wf"
	V2CoreClaimableSandboxPrefillCountryWs V2CoreClaimableSandboxPrefillCountry = "ws"
	V2CoreClaimableSandboxPrefillCountryXx V2CoreClaimableSandboxPrefillCountry = "xx"
	V2CoreClaimableSandboxPrefillCountryYe V2CoreClaimableSandboxPrefillCountry = "ye"
	V2CoreClaimableSandboxPrefillCountryYt V2CoreClaimableSandboxPrefillCountry = "yt"
	V2CoreClaimableSandboxPrefillCountryZa V2CoreClaimableSandboxPrefillCountry = "za"
	V2CoreClaimableSandboxPrefillCountryZm V2CoreClaimableSandboxPrefillCountry = "zm"
	V2CoreClaimableSandboxPrefillCountryZw V2CoreClaimableSandboxPrefillCountry = "zw"
)

// Keys that can be used to set up an integration for this sandbox and operate on the account.
type V2CoreClaimableSandboxAPIKeys struct {
	// Used to communicate with [Stripe's MCP server](https://docs.stripe.com/mcp).
	// This allows LLM agents to securely operate on a Stripe account.
	Mcp string `json:"mcp"`
	// Publicly accessible in a web or mobile app client-side code.
	Publishable string `json:"publishable"`
	// Should be stored securely in server-side code (such as an environment variable).
	Secret string `json:"secret"`
}

// Values prefilled during the creation of the sandbox.
type V2CoreClaimableSandboxPrefill struct {
	// Country in which the account holder resides, or in which the business is legally established.
	// Use two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country V2CoreClaimableSandboxPrefillCountry `json:"country"`
	// Email that this sandbox is meant to be claimed by. Stripe will
	// send an email to this email address before the sandbox expires.
	Email string `json:"email"`
	// Name for the sandbox.
	Name string `json:"name"`
}

// A claimable sandbox represents a Stripe sandbox that is anonymous.
// When it is created, it can be prefilled with specific metadata, such as email, name, or country.
// Claimable sandboxes can be claimed through a URL. When a user claims a sandbox through this URL,
// it will prompt them to create a new Stripe account. Or, it will allow them to claim this sandbox in their
// existing Stripe account.
// Claimable sandboxes have 60 days to be claimed. After this expiration time has passed,
// if the sandbox is not claimed, it will be deleted.
type V2CoreClaimableSandbox struct {
	APIResource
	// Keys that can be used to set up an integration for this sandbox and operate on the account.
	APIKeys *V2CoreClaimableSandboxAPIKeys `json:"api_keys"`
	// URL for user to claim sandbox into their existing Stripe account.
	ClaimURL string `json:"claim_url"`
	// When the sandbox is created.
	Created time.Time `json:"created"`
	// Unique identifier for the Claimable sandbox.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Values prefilled during the creation of the sandbox.
	Prefill *V2CoreClaimableSandboxPrefill `json:"prefill"`
}
