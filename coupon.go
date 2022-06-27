//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// One of `forever`, `once`, and `repeating`. Describes how long a customer who applies this coupon will get the discount.
type CouponDuration string

// List of values that CouponDuration can take
const (
	CouponDurationForever   CouponDuration = "forever"
	CouponDurationOnce      CouponDuration = "once"
	CouponDurationRepeating CouponDuration = "repeating"
)

// Returns a list of your coupons.
type CouponListParams struct {
	ListParams `form:"*"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
}

// A hash containing directions for what this Coupon will apply discounts to.
type CouponAppliesToParams struct {
	// An array of Product IDs that this Coupon will apply to.
	Products []*string `form:"products"`
}

// Coupon defined in AED.
type CouponCurrencyOptionsAedParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AFN.
type CouponCurrencyOptionsAfnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ALL.
type CouponCurrencyOptionsAllParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AMD.
type CouponCurrencyOptionsAmdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ANG.
type CouponCurrencyOptionsAngParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AOA.
type CouponCurrencyOptionsAoaParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ARS.
type CouponCurrencyOptionsArsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AUD.
type CouponCurrencyOptionsAUDParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AWG.
type CouponCurrencyOptionsAwgParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in AZN.
type CouponCurrencyOptionsAznParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BAM.
type CouponCurrencyOptionsBamParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BBD.
type CouponCurrencyOptionsBbdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BDT.
type CouponCurrencyOptionsBdtParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BGN.
type CouponCurrencyOptionsBgnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BHD.
type CouponCurrencyOptionsBhdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BIF.
type CouponCurrencyOptionsBifParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BMD.
type CouponCurrencyOptionsBmdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BND.
type CouponCurrencyOptionsBndParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BOB.
type CouponCurrencyOptionsBobParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BRL.
type CouponCurrencyOptionsBrlParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BSD.
type CouponCurrencyOptionsBsdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BTN.
type CouponCurrencyOptionsBtnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BWP.
type CouponCurrencyOptionsBwpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BYN.
type CouponCurrencyOptionsBynParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in BZD.
type CouponCurrencyOptionsBzdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CAD.
type CouponCurrencyOptionsCADParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CDF.
type CouponCurrencyOptionsCdfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CHF.
type CouponCurrencyOptionsCHFParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CLP.
type CouponCurrencyOptionsClpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CNY.
type CouponCurrencyOptionsCnyParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in COP.
type CouponCurrencyOptionsCopParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CRC.
type CouponCurrencyOptionsCrcParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CVE.
type CouponCurrencyOptionsCveParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in CZK.
type CouponCurrencyOptionsCZKParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in DJF.
type CouponCurrencyOptionsDjfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in DKK.
type CouponCurrencyOptionsDKKParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in DOP.
type CouponCurrencyOptionsDopParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in DZD.
type CouponCurrencyOptionsDzdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in EGP.
type CouponCurrencyOptionsEgpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ETB.
type CouponCurrencyOptionsEtbParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in EUR.
type CouponCurrencyOptionsEURParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in FJD.
type CouponCurrencyOptionsFjdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in FKP.
type CouponCurrencyOptionsFkpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GBP.
type CouponCurrencyOptionsGBPParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GEL.
type CouponCurrencyOptionsGelParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GHS.
type CouponCurrencyOptionsGhsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GIP.
type CouponCurrencyOptionsGipParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GMD.
type CouponCurrencyOptionsGmdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GNF.
type CouponCurrencyOptionsGnfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GTQ.
type CouponCurrencyOptionsGtqParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in GYD.
type CouponCurrencyOptionsGydParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in HKD.
type CouponCurrencyOptionsHKDParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in HNL.
type CouponCurrencyOptionsHnlParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in HRK.
type CouponCurrencyOptionsHrkParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in HTG.
type CouponCurrencyOptionsHtgParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in HUF.
type CouponCurrencyOptionsHufParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in IDR.
type CouponCurrencyOptionsIdrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ILS.
type CouponCurrencyOptionsIlsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in INR.
type CouponCurrencyOptionsInrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ISK.
type CouponCurrencyOptionsIskParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in JMD.
type CouponCurrencyOptionsJmdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in JOD.
type CouponCurrencyOptionsJodParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in JPY.
type CouponCurrencyOptionsJpyParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KES.
type CouponCurrencyOptionsKesParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KGS.
type CouponCurrencyOptionsKgsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KHR.
type CouponCurrencyOptionsKhrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KMF.
type CouponCurrencyOptionsKmfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KRW.
type CouponCurrencyOptionsKrwParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KWD.
type CouponCurrencyOptionsKwdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KYD.
type CouponCurrencyOptionsKydParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in KZT.
type CouponCurrencyOptionsKztParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in LAK.
type CouponCurrencyOptionsLakParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in LBP.
type CouponCurrencyOptionsLbpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in LKR.
type CouponCurrencyOptionsLkrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in LRD.
type CouponCurrencyOptionsLrdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in LSL.
type CouponCurrencyOptionsLslParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MAD.
type CouponCurrencyOptionsMadParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MDL.
type CouponCurrencyOptionsMdlParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MGA.
type CouponCurrencyOptionsMgaParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MKD.
type CouponCurrencyOptionsMkdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MMK.
type CouponCurrencyOptionsMmkParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MNT.
type CouponCurrencyOptionsMntParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MOP.
type CouponCurrencyOptionsMopParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MRO.
type CouponCurrencyOptionsMroParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MUR.
type CouponCurrencyOptionsMurParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MVR.
type CouponCurrencyOptionsMvrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MWK.
type CouponCurrencyOptionsMwkParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MXN.
type CouponCurrencyOptionsMxnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MYR.
type CouponCurrencyOptionsMYRParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in MZN.
type CouponCurrencyOptionsMznParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NAD.
type CouponCurrencyOptionsNadParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NGN.
type CouponCurrencyOptionsNgnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NIO.
type CouponCurrencyOptionsNioParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NOK.
type CouponCurrencyOptionsNOKParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NPR.
type CouponCurrencyOptionsNprParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in NZD.
type CouponCurrencyOptionsNZDParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in OMR.
type CouponCurrencyOptionsOmrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PAB.
type CouponCurrencyOptionsPabParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PEN.
type CouponCurrencyOptionsPenParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PGK.
type CouponCurrencyOptionsPgkParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PHP.
type CouponCurrencyOptionsPhpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PKR.
type CouponCurrencyOptionsPkrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PLN.
type CouponCurrencyOptionsPlnParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in PYG.
type CouponCurrencyOptionsPygParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in QAR.
type CouponCurrencyOptionsQarParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in RON.
type CouponCurrencyOptionsRonParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in RSD.
type CouponCurrencyOptionsRsdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in RUB.
type CouponCurrencyOptionsRubParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in RWF.
type CouponCurrencyOptionsRwfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SAR.
type CouponCurrencyOptionsSarParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SBD.
type CouponCurrencyOptionsSbdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SCR.
type CouponCurrencyOptionsScrParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SEK.
type CouponCurrencyOptionsSEKParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SGD.
type CouponCurrencyOptionsSGDParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SHP.
type CouponCurrencyOptionsShpParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SLL.
type CouponCurrencyOptionsSllParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SOS.
type CouponCurrencyOptionsSosParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SRD.
type CouponCurrencyOptionsSrdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in STD.
type CouponCurrencyOptionsStdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in SZL.
type CouponCurrencyOptionsSzlParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in THB.
type CouponCurrencyOptionsThbParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TJS.
type CouponCurrencyOptionsTjsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TND.
type CouponCurrencyOptionsTndParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TOP.
type CouponCurrencyOptionsTopParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TRY.
type CouponCurrencyOptionsTryParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TTD.
type CouponCurrencyOptionsTtdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TWD.
type CouponCurrencyOptionsTwdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in TZS.
type CouponCurrencyOptionsTzsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in UAH.
type CouponCurrencyOptionsUahParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in UGX.
type CouponCurrencyOptionsUgxParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in USD.
type CouponCurrencyOptionsUSDParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in USDC.
type CouponCurrencyOptionsUsdcParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in UYU.
type CouponCurrencyOptionsUyuParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in UZS.
type CouponCurrencyOptionsUzsParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in VND.
type CouponCurrencyOptionsVndParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in VUV.
type CouponCurrencyOptionsVuvParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in WST.
type CouponCurrencyOptionsWstParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in XAF.
type CouponCurrencyOptionsXafParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in XCD.
type CouponCurrencyOptionsXcdParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in XOF.
type CouponCurrencyOptionsXofParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in XPF.
type CouponCurrencyOptionsXpfParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in YER.
type CouponCurrencyOptionsYerParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ZAR.
type CouponCurrencyOptionsZarParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupon defined in ZMW.
type CouponCurrencyOptionsZmwParams struct {
	// A positive integer representing the amount to subtract from an invoice total.
	AmountOff *int64 `form:"amount_off"`
}

// Coupons defined in each available currency option (only supported if `amount_off` is passed). Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type CouponCurrencyOptionsParams struct {
	// Coupon defined in AED.
	Aed *CouponCurrencyOptionsAedParams `form:"aed"`
	// Coupon defined in AFN.
	Afn *CouponCurrencyOptionsAfnParams `form:"afn"`
	// Coupon defined in ALL.
	All *CouponCurrencyOptionsAllParams `form:"all"`
	// Coupon defined in AMD.
	Amd *CouponCurrencyOptionsAmdParams `form:"amd"`
	// Coupon defined in ANG.
	Ang *CouponCurrencyOptionsAngParams `form:"ang"`
	// Coupon defined in AOA.
	Aoa *CouponCurrencyOptionsAoaParams `form:"aoa"`
	// Coupon defined in ARS.
	Ars *CouponCurrencyOptionsArsParams `form:"ars"`
	// Coupon defined in AUD.
	AUD *CouponCurrencyOptionsAUDParams `form:"aud"`
	// Coupon defined in AWG.
	Awg *CouponCurrencyOptionsAwgParams `form:"awg"`
	// Coupon defined in AZN.
	Azn *CouponCurrencyOptionsAznParams `form:"azn"`
	// Coupon defined in BAM.
	Bam *CouponCurrencyOptionsBamParams `form:"bam"`
	// Coupon defined in BBD.
	Bbd *CouponCurrencyOptionsBbdParams `form:"bbd"`
	// Coupon defined in BDT.
	Bdt *CouponCurrencyOptionsBdtParams `form:"bdt"`
	// Coupon defined in BGN.
	Bgn *CouponCurrencyOptionsBgnParams `form:"bgn"`
	// Coupon defined in BHD.
	Bhd *CouponCurrencyOptionsBhdParams `form:"bhd"`
	// Coupon defined in BIF.
	Bif *CouponCurrencyOptionsBifParams `form:"bif"`
	// Coupon defined in BMD.
	Bmd *CouponCurrencyOptionsBmdParams `form:"bmd"`
	// Coupon defined in BND.
	Bnd *CouponCurrencyOptionsBndParams `form:"bnd"`
	// Coupon defined in BOB.
	Bob *CouponCurrencyOptionsBobParams `form:"bob"`
	// Coupon defined in BRL.
	Brl *CouponCurrencyOptionsBrlParams `form:"brl"`
	// Coupon defined in BSD.
	Bsd *CouponCurrencyOptionsBsdParams `form:"bsd"`
	// Coupon defined in BTN.
	Btn *CouponCurrencyOptionsBtnParams `form:"btn"`
	// Coupon defined in BWP.
	Bwp *CouponCurrencyOptionsBwpParams `form:"bwp"`
	// Coupon defined in BYN.
	Byn *CouponCurrencyOptionsBynParams `form:"byn"`
	// Coupon defined in BZD.
	Bzd *CouponCurrencyOptionsBzdParams `form:"bzd"`
	// Coupon defined in CAD.
	CAD *CouponCurrencyOptionsCADParams `form:"cad"`
	// Coupon defined in CDF.
	Cdf *CouponCurrencyOptionsCdfParams `form:"cdf"`
	// Coupon defined in CHF.
	CHF *CouponCurrencyOptionsCHFParams `form:"chf"`
	// Coupon defined in CLP.
	Clp *CouponCurrencyOptionsClpParams `form:"clp"`
	// Coupon defined in CNY.
	Cny *CouponCurrencyOptionsCnyParams `form:"cny"`
	// Coupon defined in COP.
	Cop *CouponCurrencyOptionsCopParams `form:"cop"`
	// Coupon defined in CRC.
	Crc *CouponCurrencyOptionsCrcParams `form:"crc"`
	// Coupon defined in CVE.
	Cve *CouponCurrencyOptionsCveParams `form:"cve"`
	// Coupon defined in CZK.
	CZK *CouponCurrencyOptionsCZKParams `form:"czk"`
	// Coupon defined in DJF.
	Djf *CouponCurrencyOptionsDjfParams `form:"djf"`
	// Coupon defined in DKK.
	DKK *CouponCurrencyOptionsDKKParams `form:"dkk"`
	// Coupon defined in DOP.
	Dop *CouponCurrencyOptionsDopParams `form:"dop"`
	// Coupon defined in DZD.
	Dzd *CouponCurrencyOptionsDzdParams `form:"dzd"`
	// Coupon defined in EGP.
	Egp *CouponCurrencyOptionsEgpParams `form:"egp"`
	// Coupon defined in ETB.
	Etb *CouponCurrencyOptionsEtbParams `form:"etb"`
	// Coupon defined in EUR.
	EUR *CouponCurrencyOptionsEURParams `form:"eur"`
	// Coupon defined in FJD.
	Fjd *CouponCurrencyOptionsFjdParams `form:"fjd"`
	// Coupon defined in FKP.
	Fkp *CouponCurrencyOptionsFkpParams `form:"fkp"`
	// Coupon defined in GBP.
	GBP *CouponCurrencyOptionsGBPParams `form:"gbp"`
	// Coupon defined in GEL.
	Gel *CouponCurrencyOptionsGelParams `form:"gel"`
	// Coupon defined in GHS.
	Ghs *CouponCurrencyOptionsGhsParams `form:"ghs"`
	// Coupon defined in GIP.
	Gip *CouponCurrencyOptionsGipParams `form:"gip"`
	// Coupon defined in GMD.
	Gmd *CouponCurrencyOptionsGmdParams `form:"gmd"`
	// Coupon defined in GNF.
	Gnf *CouponCurrencyOptionsGnfParams `form:"gnf"`
	// Coupon defined in GTQ.
	Gtq *CouponCurrencyOptionsGtqParams `form:"gtq"`
	// Coupon defined in GYD.
	Gyd *CouponCurrencyOptionsGydParams `form:"gyd"`
	// Coupon defined in HKD.
	HKD *CouponCurrencyOptionsHKDParams `form:"hkd"`
	// Coupon defined in HNL.
	Hnl *CouponCurrencyOptionsHnlParams `form:"hnl"`
	// Coupon defined in HRK.
	Hrk *CouponCurrencyOptionsHrkParams `form:"hrk"`
	// Coupon defined in HTG.
	Htg *CouponCurrencyOptionsHtgParams `form:"htg"`
	// Coupon defined in HUF.
	Huf *CouponCurrencyOptionsHufParams `form:"huf"`
	// Coupon defined in IDR.
	Idr *CouponCurrencyOptionsIdrParams `form:"idr"`
	// Coupon defined in ILS.
	Ils *CouponCurrencyOptionsIlsParams `form:"ils"`
	// Coupon defined in INR.
	Inr *CouponCurrencyOptionsInrParams `form:"inr"`
	// Coupon defined in ISK.
	Isk *CouponCurrencyOptionsIskParams `form:"isk"`
	// Coupon defined in JMD.
	Jmd *CouponCurrencyOptionsJmdParams `form:"jmd"`
	// Coupon defined in JOD.
	Jod *CouponCurrencyOptionsJodParams `form:"jod"`
	// Coupon defined in JPY.
	Jpy *CouponCurrencyOptionsJpyParams `form:"jpy"`
	// Coupon defined in KES.
	Kes *CouponCurrencyOptionsKesParams `form:"kes"`
	// Coupon defined in KGS.
	Kgs *CouponCurrencyOptionsKgsParams `form:"kgs"`
	// Coupon defined in KHR.
	Khr *CouponCurrencyOptionsKhrParams `form:"khr"`
	// Coupon defined in KMF.
	Kmf *CouponCurrencyOptionsKmfParams `form:"kmf"`
	// Coupon defined in KRW.
	Krw *CouponCurrencyOptionsKrwParams `form:"krw"`
	// Coupon defined in KWD.
	Kwd *CouponCurrencyOptionsKwdParams `form:"kwd"`
	// Coupon defined in KYD.
	Kyd *CouponCurrencyOptionsKydParams `form:"kyd"`
	// Coupon defined in KZT.
	Kzt *CouponCurrencyOptionsKztParams `form:"kzt"`
	// Coupon defined in LAK.
	Lak *CouponCurrencyOptionsLakParams `form:"lak"`
	// Coupon defined in LBP.
	Lbp *CouponCurrencyOptionsLbpParams `form:"lbp"`
	// Coupon defined in LKR.
	Lkr *CouponCurrencyOptionsLkrParams `form:"lkr"`
	// Coupon defined in LRD.
	Lrd *CouponCurrencyOptionsLrdParams `form:"lrd"`
	// Coupon defined in LSL.
	Lsl *CouponCurrencyOptionsLslParams `form:"lsl"`
	// Coupon defined in MAD.
	Mad *CouponCurrencyOptionsMadParams `form:"mad"`
	// Coupon defined in MDL.
	Mdl *CouponCurrencyOptionsMdlParams `form:"mdl"`
	// Coupon defined in MGA.
	Mga *CouponCurrencyOptionsMgaParams `form:"mga"`
	// Coupon defined in MKD.
	Mkd *CouponCurrencyOptionsMkdParams `form:"mkd"`
	// Coupon defined in MMK.
	Mmk *CouponCurrencyOptionsMmkParams `form:"mmk"`
	// Coupon defined in MNT.
	Mnt *CouponCurrencyOptionsMntParams `form:"mnt"`
	// Coupon defined in MOP.
	Mop *CouponCurrencyOptionsMopParams `form:"mop"`
	// Coupon defined in MRO.
	Mro *CouponCurrencyOptionsMroParams `form:"mro"`
	// Coupon defined in MUR.
	Mur *CouponCurrencyOptionsMurParams `form:"mur"`
	// Coupon defined in MVR.
	Mvr *CouponCurrencyOptionsMvrParams `form:"mvr"`
	// Coupon defined in MWK.
	Mwk *CouponCurrencyOptionsMwkParams `form:"mwk"`
	// Coupon defined in MXN.
	Mxn *CouponCurrencyOptionsMxnParams `form:"mxn"`
	// Coupon defined in MYR.
	MYR *CouponCurrencyOptionsMYRParams `form:"myr"`
	// Coupon defined in MZN.
	Mzn *CouponCurrencyOptionsMznParams `form:"mzn"`
	// Coupon defined in NAD.
	Nad *CouponCurrencyOptionsNadParams `form:"nad"`
	// Coupon defined in NGN.
	Ngn *CouponCurrencyOptionsNgnParams `form:"ngn"`
	// Coupon defined in NIO.
	Nio *CouponCurrencyOptionsNioParams `form:"nio"`
	// Coupon defined in NOK.
	NOK *CouponCurrencyOptionsNOKParams `form:"nok"`
	// Coupon defined in NPR.
	Npr *CouponCurrencyOptionsNprParams `form:"npr"`
	// Coupon defined in NZD.
	NZD *CouponCurrencyOptionsNZDParams `form:"nzd"`
	// Coupon defined in OMR.
	Omr *CouponCurrencyOptionsOmrParams `form:"omr"`
	// Coupon defined in PAB.
	Pab *CouponCurrencyOptionsPabParams `form:"pab"`
	// Coupon defined in PEN.
	Pen *CouponCurrencyOptionsPenParams `form:"pen"`
	// Coupon defined in PGK.
	Pgk *CouponCurrencyOptionsPgkParams `form:"pgk"`
	// Coupon defined in PHP.
	Php *CouponCurrencyOptionsPhpParams `form:"php"`
	// Coupon defined in PKR.
	Pkr *CouponCurrencyOptionsPkrParams `form:"pkr"`
	// Coupon defined in PLN.
	Pln *CouponCurrencyOptionsPlnParams `form:"pln"`
	// Coupon defined in PYG.
	Pyg *CouponCurrencyOptionsPygParams `form:"pyg"`
	// Coupon defined in QAR.
	Qar *CouponCurrencyOptionsQarParams `form:"qar"`
	// Coupon defined in RON.
	Ron *CouponCurrencyOptionsRonParams `form:"ron"`
	// Coupon defined in RSD.
	Rsd *CouponCurrencyOptionsRsdParams `form:"rsd"`
	// Coupon defined in RUB.
	Rub *CouponCurrencyOptionsRubParams `form:"rub"`
	// Coupon defined in RWF.
	Rwf *CouponCurrencyOptionsRwfParams `form:"rwf"`
	// Coupon defined in SAR.
	Sar *CouponCurrencyOptionsSarParams `form:"sar"`
	// Coupon defined in SBD.
	Sbd *CouponCurrencyOptionsSbdParams `form:"sbd"`
	// Coupon defined in SCR.
	Scr *CouponCurrencyOptionsScrParams `form:"scr"`
	// Coupon defined in SEK.
	SEK *CouponCurrencyOptionsSEKParams `form:"sek"`
	// Coupon defined in SGD.
	SGD *CouponCurrencyOptionsSGDParams `form:"sgd"`
	// Coupon defined in SHP.
	Shp *CouponCurrencyOptionsShpParams `form:"shp"`
	// Coupon defined in SLL.
	Sll *CouponCurrencyOptionsSllParams `form:"sll"`
	// Coupon defined in SOS.
	Sos *CouponCurrencyOptionsSosParams `form:"sos"`
	// Coupon defined in SRD.
	Srd *CouponCurrencyOptionsSrdParams `form:"srd"`
	// Coupon defined in STD.
	Std *CouponCurrencyOptionsStdParams `form:"std"`
	// Coupon defined in SZL.
	Szl *CouponCurrencyOptionsSzlParams `form:"szl"`
	// Coupon defined in THB.
	Thb *CouponCurrencyOptionsThbParams `form:"thb"`
	// Coupon defined in TJS.
	Tjs *CouponCurrencyOptionsTjsParams `form:"tjs"`
	// Coupon defined in TND.
	Tnd *CouponCurrencyOptionsTndParams `form:"tnd"`
	// Coupon defined in TOP.
	Top *CouponCurrencyOptionsTopParams `form:"top"`
	// Coupon defined in TRY.
	Try *CouponCurrencyOptionsTryParams `form:"try"`
	// Coupon defined in TTD.
	Ttd *CouponCurrencyOptionsTtdParams `form:"ttd"`
	// Coupon defined in TWD.
	Twd *CouponCurrencyOptionsTwdParams `form:"twd"`
	// Coupon defined in TZS.
	Tzs *CouponCurrencyOptionsTzsParams `form:"tzs"`
	// Coupon defined in UAH.
	Uah *CouponCurrencyOptionsUahParams `form:"uah"`
	// Coupon defined in UGX.
	Ugx *CouponCurrencyOptionsUgxParams `form:"ugx"`
	// Coupon defined in USD.
	USD *CouponCurrencyOptionsUSDParams `form:"usd"`
	// Coupon defined in USDC.
	Usdc *CouponCurrencyOptionsUsdcParams `form:"usdc"`
	// Coupon defined in UYU.
	Uyu *CouponCurrencyOptionsUyuParams `form:"uyu"`
	// Coupon defined in UZS.
	Uzs *CouponCurrencyOptionsUzsParams `form:"uzs"`
	// Coupon defined in VND.
	Vnd *CouponCurrencyOptionsVndParams `form:"vnd"`
	// Coupon defined in VUV.
	Vuv *CouponCurrencyOptionsVuvParams `form:"vuv"`
	// Coupon defined in WST.
	Wst *CouponCurrencyOptionsWstParams `form:"wst"`
	// Coupon defined in XAF.
	Xaf *CouponCurrencyOptionsXafParams `form:"xaf"`
	// Coupon defined in XCD.
	Xcd *CouponCurrencyOptionsXcdParams `form:"xcd"`
	// Coupon defined in XOF.
	Xof *CouponCurrencyOptionsXofParams `form:"xof"`
	// Coupon defined in XPF.
	Xpf *CouponCurrencyOptionsXpfParams `form:"xpf"`
	// Coupon defined in YER.
	Yer *CouponCurrencyOptionsYerParams `form:"yer"`
	// Coupon defined in ZAR.
	Zar *CouponCurrencyOptionsZarParams `form:"zar"`
	// Coupon defined in ZMW.
	Zmw *CouponCurrencyOptionsZmwParams `form:"zmw"`
}

// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
//
// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
type CouponParams struct {
	Params `form:"*"`
	// A positive integer representing the amount to subtract from an invoice total (required if `percent_off` is not passed).
	AmountOff *int64 `form:"amount_off"`
	// A hash containing directions for what this Coupon will apply discounts to.
	AppliesTo *CouponAppliesToParams `form:"applies_to"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `amount_off` parameter (required if `amount_off` is passed).
	Currency *string `form:"currency"`
	// Coupons defined in each available currency option (only supported if the coupon is amount-based). Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *CouponCurrencyOptionsParams `form:"currency_options"`
	// Specifies how long the discount will be in effect if used on a subscription. Can be `forever`, `once`, or `repeating`. Defaults to `once`.
	Duration *string `form:"duration"`
	// Required only if `duration` is `repeating`, in which case it must be a positive integer that specifies the number of months the discount will be in effect.
	DurationInMonths *int64 `form:"duration_in_months"`
	// Unique string of your choice that will be used to identify this coupon when applying it to a customer. If you don't want to specify a particular code, you can leave the ID blank and we'll generate a random code for you.
	ID *string `form:"id"`
	// A positive integer specifying the number of times the coupon can be redeemed before it's no longer valid. For example, you might have a 50% off coupon that the first 20 readers of your blog can use.
	MaxRedemptions *int64 `form:"max_redemptions"`
	// Name of the coupon displayed to customers on, for instance invoices, or receipts. By default the `id` is shown if `name` is not set.
	Name *string `form:"name"`
	// A positive float larger than 0, and smaller or equal to 100, that represents the discount the coupon will apply (required if `amount_off` is not passed).
	PercentOff *float64 `form:"percent_off"`
	// Unix timestamp specifying the last time at which the coupon can be redeemed. After the redeem_by date, the coupon can no longer be applied to new customers.
	RedeemBy *int64 `form:"redeem_by"`
}
type CouponAppliesTo struct {
	// A list of product IDs this coupon applies to
	Products []string `json:"products"`
}
type CouponCurrencyOptionsAed struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAfn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAll struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAmd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAng struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAoa struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsArs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAUD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAwg struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsAzn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBam struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBbd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBdt struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBgn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBhd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBif struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBmd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBnd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBob struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBrl struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBsd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBtn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBwp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsByn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsBzd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCAD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCdf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCHF struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsClp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCny struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCop struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCrc struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCve struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsCZK struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsDjf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsDKK struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsDop struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsDzd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsEgp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsEtb struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsEUR struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsFjd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsFkp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGBP struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGel struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGhs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGip struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGmd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGnf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGtq struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsGyd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsHKD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsHnl struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsHrk struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsHtg struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsHuf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsIdr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsIls struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsInr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsIsk struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsJmd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsJod struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsJpy struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKes struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKgs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKhr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKmf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKrw struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKwd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKyd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsKzt struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsLak struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsLbp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsLkr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsLrd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsLsl struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMad struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMdl struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMga struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMkd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMmk struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMnt struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMop struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMro struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMur struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMvr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMwk struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMxn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMYR struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsMzn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNad struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNgn struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNio struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNOK struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNpr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsNZD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsOmr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPab struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPen struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPgk struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPhp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPkr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPln struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsPyg struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsQar struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsRon struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsRsd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsRub struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsRwf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSar struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSbd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsScr struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSEK struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSGD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsShp struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSll struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSos struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSrd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsStd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsSzl struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsThb struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTjs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTnd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTop struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTry struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTtd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTwd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsTzs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUah struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUgx struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUSD struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUsdc struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUyu struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsUzs struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsVnd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsVuv struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsWst struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsXaf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsXcd struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsXof struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsXpf struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsYer struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsZar struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptionsZmw struct {
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64 `json:"amount_off"`
}
type CouponCurrencyOptions struct {
	Aed  *CouponCurrencyOptionsAed  `json:"aed"`
	Afn  *CouponCurrencyOptionsAfn  `json:"afn"`
	All  *CouponCurrencyOptionsAll  `json:"all"`
	Amd  *CouponCurrencyOptionsAmd  `json:"amd"`
	Ang  *CouponCurrencyOptionsAng  `json:"ang"`
	Aoa  *CouponCurrencyOptionsAoa  `json:"aoa"`
	Ars  *CouponCurrencyOptionsArs  `json:"ars"`
	AUD  *CouponCurrencyOptionsAUD  `json:"aud"`
	Awg  *CouponCurrencyOptionsAwg  `json:"awg"`
	Azn  *CouponCurrencyOptionsAzn  `json:"azn"`
	Bam  *CouponCurrencyOptionsBam  `json:"bam"`
	Bbd  *CouponCurrencyOptionsBbd  `json:"bbd"`
	Bdt  *CouponCurrencyOptionsBdt  `json:"bdt"`
	Bgn  *CouponCurrencyOptionsBgn  `json:"bgn"`
	Bhd  *CouponCurrencyOptionsBhd  `json:"bhd"`
	Bif  *CouponCurrencyOptionsBif  `json:"bif"`
	Bmd  *CouponCurrencyOptionsBmd  `json:"bmd"`
	Bnd  *CouponCurrencyOptionsBnd  `json:"bnd"`
	Bob  *CouponCurrencyOptionsBob  `json:"bob"`
	Brl  *CouponCurrencyOptionsBrl  `json:"brl"`
	Bsd  *CouponCurrencyOptionsBsd  `json:"bsd"`
	Btn  *CouponCurrencyOptionsBtn  `json:"btn"`
	Bwp  *CouponCurrencyOptionsBwp  `json:"bwp"`
	Byn  *CouponCurrencyOptionsByn  `json:"byn"`
	Bzd  *CouponCurrencyOptionsBzd  `json:"bzd"`
	CAD  *CouponCurrencyOptionsCAD  `json:"cad"`
	Cdf  *CouponCurrencyOptionsCdf  `json:"cdf"`
	CHF  *CouponCurrencyOptionsCHF  `json:"chf"`
	Clp  *CouponCurrencyOptionsClp  `json:"clp"`
	Cny  *CouponCurrencyOptionsCny  `json:"cny"`
	Cop  *CouponCurrencyOptionsCop  `json:"cop"`
	Crc  *CouponCurrencyOptionsCrc  `json:"crc"`
	Cve  *CouponCurrencyOptionsCve  `json:"cve"`
	CZK  *CouponCurrencyOptionsCZK  `json:"czk"`
	Djf  *CouponCurrencyOptionsDjf  `json:"djf"`
	DKK  *CouponCurrencyOptionsDKK  `json:"dkk"`
	Dop  *CouponCurrencyOptionsDop  `json:"dop"`
	Dzd  *CouponCurrencyOptionsDzd  `json:"dzd"`
	Egp  *CouponCurrencyOptionsEgp  `json:"egp"`
	Etb  *CouponCurrencyOptionsEtb  `json:"etb"`
	EUR  *CouponCurrencyOptionsEUR  `json:"eur"`
	Fjd  *CouponCurrencyOptionsFjd  `json:"fjd"`
	Fkp  *CouponCurrencyOptionsFkp  `json:"fkp"`
	GBP  *CouponCurrencyOptionsGBP  `json:"gbp"`
	Gel  *CouponCurrencyOptionsGel  `json:"gel"`
	Ghs  *CouponCurrencyOptionsGhs  `json:"ghs"`
	Gip  *CouponCurrencyOptionsGip  `json:"gip"`
	Gmd  *CouponCurrencyOptionsGmd  `json:"gmd"`
	Gnf  *CouponCurrencyOptionsGnf  `json:"gnf"`
	Gtq  *CouponCurrencyOptionsGtq  `json:"gtq"`
	Gyd  *CouponCurrencyOptionsGyd  `json:"gyd"`
	HKD  *CouponCurrencyOptionsHKD  `json:"hkd"`
	Hnl  *CouponCurrencyOptionsHnl  `json:"hnl"`
	Hrk  *CouponCurrencyOptionsHrk  `json:"hrk"`
	Htg  *CouponCurrencyOptionsHtg  `json:"htg"`
	Huf  *CouponCurrencyOptionsHuf  `json:"huf"`
	Idr  *CouponCurrencyOptionsIdr  `json:"idr"`
	Ils  *CouponCurrencyOptionsIls  `json:"ils"`
	Inr  *CouponCurrencyOptionsInr  `json:"inr"`
	Isk  *CouponCurrencyOptionsIsk  `json:"isk"`
	Jmd  *CouponCurrencyOptionsJmd  `json:"jmd"`
	Jod  *CouponCurrencyOptionsJod  `json:"jod"`
	Jpy  *CouponCurrencyOptionsJpy  `json:"jpy"`
	Kes  *CouponCurrencyOptionsKes  `json:"kes"`
	Kgs  *CouponCurrencyOptionsKgs  `json:"kgs"`
	Khr  *CouponCurrencyOptionsKhr  `json:"khr"`
	Kmf  *CouponCurrencyOptionsKmf  `json:"kmf"`
	Krw  *CouponCurrencyOptionsKrw  `json:"krw"`
	Kwd  *CouponCurrencyOptionsKwd  `json:"kwd"`
	Kyd  *CouponCurrencyOptionsKyd  `json:"kyd"`
	Kzt  *CouponCurrencyOptionsKzt  `json:"kzt"`
	Lak  *CouponCurrencyOptionsLak  `json:"lak"`
	Lbp  *CouponCurrencyOptionsLbp  `json:"lbp"`
	Lkr  *CouponCurrencyOptionsLkr  `json:"lkr"`
	Lrd  *CouponCurrencyOptionsLrd  `json:"lrd"`
	Lsl  *CouponCurrencyOptionsLsl  `json:"lsl"`
	Mad  *CouponCurrencyOptionsMad  `json:"mad"`
	Mdl  *CouponCurrencyOptionsMdl  `json:"mdl"`
	Mga  *CouponCurrencyOptionsMga  `json:"mga"`
	Mkd  *CouponCurrencyOptionsMkd  `json:"mkd"`
	Mmk  *CouponCurrencyOptionsMmk  `json:"mmk"`
	Mnt  *CouponCurrencyOptionsMnt  `json:"mnt"`
	Mop  *CouponCurrencyOptionsMop  `json:"mop"`
	Mro  *CouponCurrencyOptionsMro  `json:"mro"`
	Mur  *CouponCurrencyOptionsMur  `json:"mur"`
	Mvr  *CouponCurrencyOptionsMvr  `json:"mvr"`
	Mwk  *CouponCurrencyOptionsMwk  `json:"mwk"`
	Mxn  *CouponCurrencyOptionsMxn  `json:"mxn"`
	MYR  *CouponCurrencyOptionsMYR  `json:"myr"`
	Mzn  *CouponCurrencyOptionsMzn  `json:"mzn"`
	Nad  *CouponCurrencyOptionsNad  `json:"nad"`
	Ngn  *CouponCurrencyOptionsNgn  `json:"ngn"`
	Nio  *CouponCurrencyOptionsNio  `json:"nio"`
	NOK  *CouponCurrencyOptionsNOK  `json:"nok"`
	Npr  *CouponCurrencyOptionsNpr  `json:"npr"`
	NZD  *CouponCurrencyOptionsNZD  `json:"nzd"`
	Omr  *CouponCurrencyOptionsOmr  `json:"omr"`
	Pab  *CouponCurrencyOptionsPab  `json:"pab"`
	Pen  *CouponCurrencyOptionsPen  `json:"pen"`
	Pgk  *CouponCurrencyOptionsPgk  `json:"pgk"`
	Php  *CouponCurrencyOptionsPhp  `json:"php"`
	Pkr  *CouponCurrencyOptionsPkr  `json:"pkr"`
	Pln  *CouponCurrencyOptionsPln  `json:"pln"`
	Pyg  *CouponCurrencyOptionsPyg  `json:"pyg"`
	Qar  *CouponCurrencyOptionsQar  `json:"qar"`
	Ron  *CouponCurrencyOptionsRon  `json:"ron"`
	Rsd  *CouponCurrencyOptionsRsd  `json:"rsd"`
	Rub  *CouponCurrencyOptionsRub  `json:"rub"`
	Rwf  *CouponCurrencyOptionsRwf  `json:"rwf"`
	Sar  *CouponCurrencyOptionsSar  `json:"sar"`
	Sbd  *CouponCurrencyOptionsSbd  `json:"sbd"`
	Scr  *CouponCurrencyOptionsScr  `json:"scr"`
	SEK  *CouponCurrencyOptionsSEK  `json:"sek"`
	SGD  *CouponCurrencyOptionsSGD  `json:"sgd"`
	Shp  *CouponCurrencyOptionsShp  `json:"shp"`
	Sll  *CouponCurrencyOptionsSll  `json:"sll"`
	Sos  *CouponCurrencyOptionsSos  `json:"sos"`
	Srd  *CouponCurrencyOptionsSrd  `json:"srd"`
	Std  *CouponCurrencyOptionsStd  `json:"std"`
	Szl  *CouponCurrencyOptionsSzl  `json:"szl"`
	Thb  *CouponCurrencyOptionsThb  `json:"thb"`
	Tjs  *CouponCurrencyOptionsTjs  `json:"tjs"`
	Tnd  *CouponCurrencyOptionsTnd  `json:"tnd"`
	Top  *CouponCurrencyOptionsTop  `json:"top"`
	Try  *CouponCurrencyOptionsTry  `json:"try"`
	Ttd  *CouponCurrencyOptionsTtd  `json:"ttd"`
	Twd  *CouponCurrencyOptionsTwd  `json:"twd"`
	Tzs  *CouponCurrencyOptionsTzs  `json:"tzs"`
	Uah  *CouponCurrencyOptionsUah  `json:"uah"`
	Ugx  *CouponCurrencyOptionsUgx  `json:"ugx"`
	USD  *CouponCurrencyOptionsUSD  `json:"usd"`
	Usdc *CouponCurrencyOptionsUsdc `json:"usdc"`
	Uyu  *CouponCurrencyOptionsUyu  `json:"uyu"`
	Uzs  *CouponCurrencyOptionsUzs  `json:"uzs"`
	Vnd  *CouponCurrencyOptionsVnd  `json:"vnd"`
	Vuv  *CouponCurrencyOptionsVuv  `json:"vuv"`
	Wst  *CouponCurrencyOptionsWst  `json:"wst"`
	Xaf  *CouponCurrencyOptionsXaf  `json:"xaf"`
	Xcd  *CouponCurrencyOptionsXcd  `json:"xcd"`
	Xof  *CouponCurrencyOptionsXof  `json:"xof"`
	Xpf  *CouponCurrencyOptionsXpf  `json:"xpf"`
	Yer  *CouponCurrencyOptionsYer  `json:"yer"`
	Zar  *CouponCurrencyOptionsZar  `json:"zar"`
	Zmw  *CouponCurrencyOptionsZmw  `json:"zmw"`
}

// A coupon contains information about a percent-off or amount-off discount you
// might want to apply to a customer. Coupons may be applied to [subscriptions](https://stripe.com/docs/api#subscriptions), [invoices](https://stripe.com/docs/api#invoices),
// [checkout sessions](https://stripe.com/docs/api/checkout/sessions), [quotes](https://stripe.com/docs/api#quotes), and more. Coupons do not work with conventional one-off [charges](https://stripe.com/docs/api#create_charge) or [payment intents](https://stripe.com/docs/api/payment_intents).
type Coupon struct {
	APIResource
	// Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff int64            `json:"amount_off"`
	AppliesTo *CouponAppliesTo `json:"applies_to"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// If `amount_off` has been set, the three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the amount to take off.
	Currency        Currency               `json:"currency"`
	CurrencyOptions *CouponCurrencyOptions `json:"currency_options"`
	Deleted         bool                   `json:"deleted"`
	// One of `forever`, `once`, and `repeating`. Describes how long a customer who applies this coupon will get the discount.
	Duration CouponDuration `json:"duration"`
	// If `duration` is `repeating`, the number of months the coupon applies. Null if coupon `duration` is `forever` or `once`.
	DurationInMonths int64 `json:"duration_in_months"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions int64 `json:"max_redemptions"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Name of the coupon displayed to customers on for instance invoices or receipts.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a %s100 invoice %s50 instead.
	PercentOff float64 `json:"percent_off"`
	// Date after which the coupon can no longer be redeemed.
	RedeemBy int64 `json:"redeem_by"`
	// Number of times this coupon has been applied to a customer.
	TimesRedeemed int64 `json:"times_redeemed"`
	// Taking account of the above properties, whether this coupon can still be applied to a customer.
	Valid bool `json:"valid"`
}

// CouponList is a list of Coupons as retrieved from a list endpoint.
type CouponList struct {
	APIResource
	ListMeta
	Data []*Coupon `json:"data"`
}

// UnmarshalJSON handles deserialization of a Coupon.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Coupon) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type coupon Coupon
	var v coupon
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Coupon(v)
	return nil
}
