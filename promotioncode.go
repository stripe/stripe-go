//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://stripe.com/docs/api/promotion_codes/list) with the desired code.
type PromotionCodeParams struct {
	Params `form:"*"`
	// Whether the promotion code is currently active.
	Active *bool `form:"active"`
	// The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code *string `form:"code"`
	// The coupon for this promotion code.
	Coupon *string `form:"coupon"`
	// The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer *string `form:"customer"`
	// The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.
	ExpiresAt *int64 `form:"expires_at"`
	// A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`.
	MaxRedemptions *int64 `form:"max_redemptions"`
	// Settings that restrict the redemption of the promotion code.
	Restrictions *PromotionCodeRestrictionsParams `form:"restrictions"`
}

// Promotion code defined in AED.
type PromotionCodeRestrictionsCurrencyOptionsAedParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AFN.
type PromotionCodeRestrictionsCurrencyOptionsAfnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ALL.
type PromotionCodeRestrictionsCurrencyOptionsAllParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AMD.
type PromotionCodeRestrictionsCurrencyOptionsAmdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ANG.
type PromotionCodeRestrictionsCurrencyOptionsAngParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AOA.
type PromotionCodeRestrictionsCurrencyOptionsAoaParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ARS.
type PromotionCodeRestrictionsCurrencyOptionsArsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AUD.
type PromotionCodeRestrictionsCurrencyOptionsAUDParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AWG.
type PromotionCodeRestrictionsCurrencyOptionsAwgParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in AZN.
type PromotionCodeRestrictionsCurrencyOptionsAznParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BAM.
type PromotionCodeRestrictionsCurrencyOptionsBamParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BBD.
type PromotionCodeRestrictionsCurrencyOptionsBbdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BDT.
type PromotionCodeRestrictionsCurrencyOptionsBdtParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BGN.
type PromotionCodeRestrictionsCurrencyOptionsBgnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BHD.
type PromotionCodeRestrictionsCurrencyOptionsBhdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BIF.
type PromotionCodeRestrictionsCurrencyOptionsBifParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BMD.
type PromotionCodeRestrictionsCurrencyOptionsBmdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BND.
type PromotionCodeRestrictionsCurrencyOptionsBndParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BOB.
type PromotionCodeRestrictionsCurrencyOptionsBobParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BRL.
type PromotionCodeRestrictionsCurrencyOptionsBrlParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BSD.
type PromotionCodeRestrictionsCurrencyOptionsBsdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BTN.
type PromotionCodeRestrictionsCurrencyOptionsBtnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BWP.
type PromotionCodeRestrictionsCurrencyOptionsBwpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BYN.
type PromotionCodeRestrictionsCurrencyOptionsBynParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in BZD.
type PromotionCodeRestrictionsCurrencyOptionsBzdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CAD.
type PromotionCodeRestrictionsCurrencyOptionsCADParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CDF.
type PromotionCodeRestrictionsCurrencyOptionsCdfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CHF.
type PromotionCodeRestrictionsCurrencyOptionsCHFParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CLP.
type PromotionCodeRestrictionsCurrencyOptionsClpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CNY.
type PromotionCodeRestrictionsCurrencyOptionsCnyParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in COP.
type PromotionCodeRestrictionsCurrencyOptionsCopParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CRC.
type PromotionCodeRestrictionsCurrencyOptionsCrcParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CVE.
type PromotionCodeRestrictionsCurrencyOptionsCveParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in CZK.
type PromotionCodeRestrictionsCurrencyOptionsCZKParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in DJF.
type PromotionCodeRestrictionsCurrencyOptionsDjfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in DKK.
type PromotionCodeRestrictionsCurrencyOptionsDKKParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in DOP.
type PromotionCodeRestrictionsCurrencyOptionsDopParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in DZD.
type PromotionCodeRestrictionsCurrencyOptionsDzdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in EGP.
type PromotionCodeRestrictionsCurrencyOptionsEgpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ETB.
type PromotionCodeRestrictionsCurrencyOptionsEtbParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in EUR.
type PromotionCodeRestrictionsCurrencyOptionsEURParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in FJD.
type PromotionCodeRestrictionsCurrencyOptionsFjdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in FKP.
type PromotionCodeRestrictionsCurrencyOptionsFkpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GBP.
type PromotionCodeRestrictionsCurrencyOptionsGBPParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GEL.
type PromotionCodeRestrictionsCurrencyOptionsGelParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GHS.
type PromotionCodeRestrictionsCurrencyOptionsGhsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GIP.
type PromotionCodeRestrictionsCurrencyOptionsGipParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GMD.
type PromotionCodeRestrictionsCurrencyOptionsGmdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GNF.
type PromotionCodeRestrictionsCurrencyOptionsGnfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GTQ.
type PromotionCodeRestrictionsCurrencyOptionsGtqParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in GYD.
type PromotionCodeRestrictionsCurrencyOptionsGydParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in HKD.
type PromotionCodeRestrictionsCurrencyOptionsHKDParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in HNL.
type PromotionCodeRestrictionsCurrencyOptionsHnlParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in HRK.
type PromotionCodeRestrictionsCurrencyOptionsHrkParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in HTG.
type PromotionCodeRestrictionsCurrencyOptionsHtgParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in HUF.
type PromotionCodeRestrictionsCurrencyOptionsHufParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in IDR.
type PromotionCodeRestrictionsCurrencyOptionsIdrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ILS.
type PromotionCodeRestrictionsCurrencyOptionsIlsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in INR.
type PromotionCodeRestrictionsCurrencyOptionsInrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ISK.
type PromotionCodeRestrictionsCurrencyOptionsIskParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in JMD.
type PromotionCodeRestrictionsCurrencyOptionsJmdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in JOD.
type PromotionCodeRestrictionsCurrencyOptionsJodParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in JPY.
type PromotionCodeRestrictionsCurrencyOptionsJpyParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KES.
type PromotionCodeRestrictionsCurrencyOptionsKesParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KGS.
type PromotionCodeRestrictionsCurrencyOptionsKgsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KHR.
type PromotionCodeRestrictionsCurrencyOptionsKhrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KMF.
type PromotionCodeRestrictionsCurrencyOptionsKmfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KRW.
type PromotionCodeRestrictionsCurrencyOptionsKrwParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KWD.
type PromotionCodeRestrictionsCurrencyOptionsKwdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KYD.
type PromotionCodeRestrictionsCurrencyOptionsKydParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in KZT.
type PromotionCodeRestrictionsCurrencyOptionsKztParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in LAK.
type PromotionCodeRestrictionsCurrencyOptionsLakParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in LBP.
type PromotionCodeRestrictionsCurrencyOptionsLbpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in LKR.
type PromotionCodeRestrictionsCurrencyOptionsLkrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in LRD.
type PromotionCodeRestrictionsCurrencyOptionsLrdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in LSL.
type PromotionCodeRestrictionsCurrencyOptionsLslParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MAD.
type PromotionCodeRestrictionsCurrencyOptionsMadParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MDL.
type PromotionCodeRestrictionsCurrencyOptionsMdlParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MGA.
type PromotionCodeRestrictionsCurrencyOptionsMgaParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MKD.
type PromotionCodeRestrictionsCurrencyOptionsMkdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MMK.
type PromotionCodeRestrictionsCurrencyOptionsMmkParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MNT.
type PromotionCodeRestrictionsCurrencyOptionsMntParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MOP.
type PromotionCodeRestrictionsCurrencyOptionsMopParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MRO.
type PromotionCodeRestrictionsCurrencyOptionsMroParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MUR.
type PromotionCodeRestrictionsCurrencyOptionsMurParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MVR.
type PromotionCodeRestrictionsCurrencyOptionsMvrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MWK.
type PromotionCodeRestrictionsCurrencyOptionsMwkParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MXN.
type PromotionCodeRestrictionsCurrencyOptionsMxnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MYR.
type PromotionCodeRestrictionsCurrencyOptionsMYRParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in MZN.
type PromotionCodeRestrictionsCurrencyOptionsMznParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NAD.
type PromotionCodeRestrictionsCurrencyOptionsNadParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NGN.
type PromotionCodeRestrictionsCurrencyOptionsNgnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NIO.
type PromotionCodeRestrictionsCurrencyOptionsNioParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NOK.
type PromotionCodeRestrictionsCurrencyOptionsNOKParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NPR.
type PromotionCodeRestrictionsCurrencyOptionsNprParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in NZD.
type PromotionCodeRestrictionsCurrencyOptionsNZDParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in OMR.
type PromotionCodeRestrictionsCurrencyOptionsOmrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PAB.
type PromotionCodeRestrictionsCurrencyOptionsPabParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PEN.
type PromotionCodeRestrictionsCurrencyOptionsPenParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PGK.
type PromotionCodeRestrictionsCurrencyOptionsPgkParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PHP.
type PromotionCodeRestrictionsCurrencyOptionsPhpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PKR.
type PromotionCodeRestrictionsCurrencyOptionsPkrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PLN.
type PromotionCodeRestrictionsCurrencyOptionsPlnParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in PYG.
type PromotionCodeRestrictionsCurrencyOptionsPygParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in QAR.
type PromotionCodeRestrictionsCurrencyOptionsQarParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in RON.
type PromotionCodeRestrictionsCurrencyOptionsRonParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in RSD.
type PromotionCodeRestrictionsCurrencyOptionsRsdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in RUB.
type PromotionCodeRestrictionsCurrencyOptionsRubParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in RWF.
type PromotionCodeRestrictionsCurrencyOptionsRwfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SAR.
type PromotionCodeRestrictionsCurrencyOptionsSarParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SBD.
type PromotionCodeRestrictionsCurrencyOptionsSbdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SCR.
type PromotionCodeRestrictionsCurrencyOptionsScrParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SEK.
type PromotionCodeRestrictionsCurrencyOptionsSEKParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SGD.
type PromotionCodeRestrictionsCurrencyOptionsSGDParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SHP.
type PromotionCodeRestrictionsCurrencyOptionsShpParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SLL.
type PromotionCodeRestrictionsCurrencyOptionsSllParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SOS.
type PromotionCodeRestrictionsCurrencyOptionsSosParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SRD.
type PromotionCodeRestrictionsCurrencyOptionsSrdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in STD.
type PromotionCodeRestrictionsCurrencyOptionsStdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in SZL.
type PromotionCodeRestrictionsCurrencyOptionsSzlParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in THB.
type PromotionCodeRestrictionsCurrencyOptionsThbParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TJS.
type PromotionCodeRestrictionsCurrencyOptionsTjsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TND.
type PromotionCodeRestrictionsCurrencyOptionsTndParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TOP.
type PromotionCodeRestrictionsCurrencyOptionsTopParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TRY.
type PromotionCodeRestrictionsCurrencyOptionsTryParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TTD.
type PromotionCodeRestrictionsCurrencyOptionsTtdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TWD.
type PromotionCodeRestrictionsCurrencyOptionsTwdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in TZS.
type PromotionCodeRestrictionsCurrencyOptionsTzsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in UAH.
type PromotionCodeRestrictionsCurrencyOptionsUahParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in UGX.
type PromotionCodeRestrictionsCurrencyOptionsUgxParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in USD.
type PromotionCodeRestrictionsCurrencyOptionsUSDParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in USDC.
type PromotionCodeRestrictionsCurrencyOptionsUsdcParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in UYU.
type PromotionCodeRestrictionsCurrencyOptionsUyuParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in UZS.
type PromotionCodeRestrictionsCurrencyOptionsUzsParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in VND.
type PromotionCodeRestrictionsCurrencyOptionsVndParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in VUV.
type PromotionCodeRestrictionsCurrencyOptionsVuvParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in WST.
type PromotionCodeRestrictionsCurrencyOptionsWstParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in XAF.
type PromotionCodeRestrictionsCurrencyOptionsXafParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in XCD.
type PromotionCodeRestrictionsCurrencyOptionsXcdParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in XOF.
type PromotionCodeRestrictionsCurrencyOptionsXofParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in XPF.
type PromotionCodeRestrictionsCurrencyOptionsXpfParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in YER.
type PromotionCodeRestrictionsCurrencyOptionsYerParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ZAR.
type PromotionCodeRestrictionsCurrencyOptionsZarParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion code defined in ZMW.
type PromotionCodeRestrictionsCurrencyOptionsZmwParams struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
}

// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PromotionCodeRestrictionsCurrencyOptionsParams struct {
	// Promotion code defined in AED.
	Aed *PromotionCodeRestrictionsCurrencyOptionsAedParams `form:"aed"`
	// Promotion code defined in AFN.
	Afn *PromotionCodeRestrictionsCurrencyOptionsAfnParams `form:"afn"`
	// Promotion code defined in ALL.
	All *PromotionCodeRestrictionsCurrencyOptionsAllParams `form:"all"`
	// Promotion code defined in AMD.
	Amd *PromotionCodeRestrictionsCurrencyOptionsAmdParams `form:"amd"`
	// Promotion code defined in ANG.
	Ang *PromotionCodeRestrictionsCurrencyOptionsAngParams `form:"ang"`
	// Promotion code defined in AOA.
	Aoa *PromotionCodeRestrictionsCurrencyOptionsAoaParams `form:"aoa"`
	// Promotion code defined in ARS.
	Ars *PromotionCodeRestrictionsCurrencyOptionsArsParams `form:"ars"`
	// Promotion code defined in AUD.
	AUD *PromotionCodeRestrictionsCurrencyOptionsAUDParams `form:"aud"`
	// Promotion code defined in AWG.
	Awg *PromotionCodeRestrictionsCurrencyOptionsAwgParams `form:"awg"`
	// Promotion code defined in AZN.
	Azn *PromotionCodeRestrictionsCurrencyOptionsAznParams `form:"azn"`
	// Promotion code defined in BAM.
	Bam *PromotionCodeRestrictionsCurrencyOptionsBamParams `form:"bam"`
	// Promotion code defined in BBD.
	Bbd *PromotionCodeRestrictionsCurrencyOptionsBbdParams `form:"bbd"`
	// Promotion code defined in BDT.
	Bdt *PromotionCodeRestrictionsCurrencyOptionsBdtParams `form:"bdt"`
	// Promotion code defined in BGN.
	Bgn *PromotionCodeRestrictionsCurrencyOptionsBgnParams `form:"bgn"`
	// Promotion code defined in BHD.
	Bhd *PromotionCodeRestrictionsCurrencyOptionsBhdParams `form:"bhd"`
	// Promotion code defined in BIF.
	Bif *PromotionCodeRestrictionsCurrencyOptionsBifParams `form:"bif"`
	// Promotion code defined in BMD.
	Bmd *PromotionCodeRestrictionsCurrencyOptionsBmdParams `form:"bmd"`
	// Promotion code defined in BND.
	Bnd *PromotionCodeRestrictionsCurrencyOptionsBndParams `form:"bnd"`
	// Promotion code defined in BOB.
	Bob *PromotionCodeRestrictionsCurrencyOptionsBobParams `form:"bob"`
	// Promotion code defined in BRL.
	Brl *PromotionCodeRestrictionsCurrencyOptionsBrlParams `form:"brl"`
	// Promotion code defined in BSD.
	Bsd *PromotionCodeRestrictionsCurrencyOptionsBsdParams `form:"bsd"`
	// Promotion code defined in BTN.
	Btn *PromotionCodeRestrictionsCurrencyOptionsBtnParams `form:"btn"`
	// Promotion code defined in BWP.
	Bwp *PromotionCodeRestrictionsCurrencyOptionsBwpParams `form:"bwp"`
	// Promotion code defined in BYN.
	Byn *PromotionCodeRestrictionsCurrencyOptionsBynParams `form:"byn"`
	// Promotion code defined in BZD.
	Bzd *PromotionCodeRestrictionsCurrencyOptionsBzdParams `form:"bzd"`
	// Promotion code defined in CAD.
	CAD *PromotionCodeRestrictionsCurrencyOptionsCADParams `form:"cad"`
	// Promotion code defined in CDF.
	Cdf *PromotionCodeRestrictionsCurrencyOptionsCdfParams `form:"cdf"`
	// Promotion code defined in CHF.
	CHF *PromotionCodeRestrictionsCurrencyOptionsCHFParams `form:"chf"`
	// Promotion code defined in CLP.
	Clp *PromotionCodeRestrictionsCurrencyOptionsClpParams `form:"clp"`
	// Promotion code defined in CNY.
	Cny *PromotionCodeRestrictionsCurrencyOptionsCnyParams `form:"cny"`
	// Promotion code defined in COP.
	Cop *PromotionCodeRestrictionsCurrencyOptionsCopParams `form:"cop"`
	// Promotion code defined in CRC.
	Crc *PromotionCodeRestrictionsCurrencyOptionsCrcParams `form:"crc"`
	// Promotion code defined in CVE.
	Cve *PromotionCodeRestrictionsCurrencyOptionsCveParams `form:"cve"`
	// Promotion code defined in CZK.
	CZK *PromotionCodeRestrictionsCurrencyOptionsCZKParams `form:"czk"`
	// Promotion code defined in DJF.
	Djf *PromotionCodeRestrictionsCurrencyOptionsDjfParams `form:"djf"`
	// Promotion code defined in DKK.
	DKK *PromotionCodeRestrictionsCurrencyOptionsDKKParams `form:"dkk"`
	// Promotion code defined in DOP.
	Dop *PromotionCodeRestrictionsCurrencyOptionsDopParams `form:"dop"`
	// Promotion code defined in DZD.
	Dzd *PromotionCodeRestrictionsCurrencyOptionsDzdParams `form:"dzd"`
	// Promotion code defined in EGP.
	Egp *PromotionCodeRestrictionsCurrencyOptionsEgpParams `form:"egp"`
	// Promotion code defined in ETB.
	Etb *PromotionCodeRestrictionsCurrencyOptionsEtbParams `form:"etb"`
	// Promotion code defined in EUR.
	EUR *PromotionCodeRestrictionsCurrencyOptionsEURParams `form:"eur"`
	// Promotion code defined in FJD.
	Fjd *PromotionCodeRestrictionsCurrencyOptionsFjdParams `form:"fjd"`
	// Promotion code defined in FKP.
	Fkp *PromotionCodeRestrictionsCurrencyOptionsFkpParams `form:"fkp"`
	// Promotion code defined in GBP.
	GBP *PromotionCodeRestrictionsCurrencyOptionsGBPParams `form:"gbp"`
	// Promotion code defined in GEL.
	Gel *PromotionCodeRestrictionsCurrencyOptionsGelParams `form:"gel"`
	// Promotion code defined in GHS.
	Ghs *PromotionCodeRestrictionsCurrencyOptionsGhsParams `form:"ghs"`
	// Promotion code defined in GIP.
	Gip *PromotionCodeRestrictionsCurrencyOptionsGipParams `form:"gip"`
	// Promotion code defined in GMD.
	Gmd *PromotionCodeRestrictionsCurrencyOptionsGmdParams `form:"gmd"`
	// Promotion code defined in GNF.
	Gnf *PromotionCodeRestrictionsCurrencyOptionsGnfParams `form:"gnf"`
	// Promotion code defined in GTQ.
	Gtq *PromotionCodeRestrictionsCurrencyOptionsGtqParams `form:"gtq"`
	// Promotion code defined in GYD.
	Gyd *PromotionCodeRestrictionsCurrencyOptionsGydParams `form:"gyd"`
	// Promotion code defined in HKD.
	HKD *PromotionCodeRestrictionsCurrencyOptionsHKDParams `form:"hkd"`
	// Promotion code defined in HNL.
	Hnl *PromotionCodeRestrictionsCurrencyOptionsHnlParams `form:"hnl"`
	// Promotion code defined in HRK.
	Hrk *PromotionCodeRestrictionsCurrencyOptionsHrkParams `form:"hrk"`
	// Promotion code defined in HTG.
	Htg *PromotionCodeRestrictionsCurrencyOptionsHtgParams `form:"htg"`
	// Promotion code defined in HUF.
	Huf *PromotionCodeRestrictionsCurrencyOptionsHufParams `form:"huf"`
	// Promotion code defined in IDR.
	Idr *PromotionCodeRestrictionsCurrencyOptionsIdrParams `form:"idr"`
	// Promotion code defined in ILS.
	Ils *PromotionCodeRestrictionsCurrencyOptionsIlsParams `form:"ils"`
	// Promotion code defined in INR.
	Inr *PromotionCodeRestrictionsCurrencyOptionsInrParams `form:"inr"`
	// Promotion code defined in ISK.
	Isk *PromotionCodeRestrictionsCurrencyOptionsIskParams `form:"isk"`
	// Promotion code defined in JMD.
	Jmd *PromotionCodeRestrictionsCurrencyOptionsJmdParams `form:"jmd"`
	// Promotion code defined in JOD.
	Jod *PromotionCodeRestrictionsCurrencyOptionsJodParams `form:"jod"`
	// Promotion code defined in JPY.
	Jpy *PromotionCodeRestrictionsCurrencyOptionsJpyParams `form:"jpy"`
	// Promotion code defined in KES.
	Kes *PromotionCodeRestrictionsCurrencyOptionsKesParams `form:"kes"`
	// Promotion code defined in KGS.
	Kgs *PromotionCodeRestrictionsCurrencyOptionsKgsParams `form:"kgs"`
	// Promotion code defined in KHR.
	Khr *PromotionCodeRestrictionsCurrencyOptionsKhrParams `form:"khr"`
	// Promotion code defined in KMF.
	Kmf *PromotionCodeRestrictionsCurrencyOptionsKmfParams `form:"kmf"`
	// Promotion code defined in KRW.
	Krw *PromotionCodeRestrictionsCurrencyOptionsKrwParams `form:"krw"`
	// Promotion code defined in KWD.
	Kwd *PromotionCodeRestrictionsCurrencyOptionsKwdParams `form:"kwd"`
	// Promotion code defined in KYD.
	Kyd *PromotionCodeRestrictionsCurrencyOptionsKydParams `form:"kyd"`
	// Promotion code defined in KZT.
	Kzt *PromotionCodeRestrictionsCurrencyOptionsKztParams `form:"kzt"`
	// Promotion code defined in LAK.
	Lak *PromotionCodeRestrictionsCurrencyOptionsLakParams `form:"lak"`
	// Promotion code defined in LBP.
	Lbp *PromotionCodeRestrictionsCurrencyOptionsLbpParams `form:"lbp"`
	// Promotion code defined in LKR.
	Lkr *PromotionCodeRestrictionsCurrencyOptionsLkrParams `form:"lkr"`
	// Promotion code defined in LRD.
	Lrd *PromotionCodeRestrictionsCurrencyOptionsLrdParams `form:"lrd"`
	// Promotion code defined in LSL.
	Lsl *PromotionCodeRestrictionsCurrencyOptionsLslParams `form:"lsl"`
	// Promotion code defined in MAD.
	Mad *PromotionCodeRestrictionsCurrencyOptionsMadParams `form:"mad"`
	// Promotion code defined in MDL.
	Mdl *PromotionCodeRestrictionsCurrencyOptionsMdlParams `form:"mdl"`
	// Promotion code defined in MGA.
	Mga *PromotionCodeRestrictionsCurrencyOptionsMgaParams `form:"mga"`
	// Promotion code defined in MKD.
	Mkd *PromotionCodeRestrictionsCurrencyOptionsMkdParams `form:"mkd"`
	// Promotion code defined in MMK.
	Mmk *PromotionCodeRestrictionsCurrencyOptionsMmkParams `form:"mmk"`
	// Promotion code defined in MNT.
	Mnt *PromotionCodeRestrictionsCurrencyOptionsMntParams `form:"mnt"`
	// Promotion code defined in MOP.
	Mop *PromotionCodeRestrictionsCurrencyOptionsMopParams `form:"mop"`
	// Promotion code defined in MRO.
	Mro *PromotionCodeRestrictionsCurrencyOptionsMroParams `form:"mro"`
	// Promotion code defined in MUR.
	Mur *PromotionCodeRestrictionsCurrencyOptionsMurParams `form:"mur"`
	// Promotion code defined in MVR.
	Mvr *PromotionCodeRestrictionsCurrencyOptionsMvrParams `form:"mvr"`
	// Promotion code defined in MWK.
	Mwk *PromotionCodeRestrictionsCurrencyOptionsMwkParams `form:"mwk"`
	// Promotion code defined in MXN.
	Mxn *PromotionCodeRestrictionsCurrencyOptionsMxnParams `form:"mxn"`
	// Promotion code defined in MYR.
	MYR *PromotionCodeRestrictionsCurrencyOptionsMYRParams `form:"myr"`
	// Promotion code defined in MZN.
	Mzn *PromotionCodeRestrictionsCurrencyOptionsMznParams `form:"mzn"`
	// Promotion code defined in NAD.
	Nad *PromotionCodeRestrictionsCurrencyOptionsNadParams `form:"nad"`
	// Promotion code defined in NGN.
	Ngn *PromotionCodeRestrictionsCurrencyOptionsNgnParams `form:"ngn"`
	// Promotion code defined in NIO.
	Nio *PromotionCodeRestrictionsCurrencyOptionsNioParams `form:"nio"`
	// Promotion code defined in NOK.
	NOK *PromotionCodeRestrictionsCurrencyOptionsNOKParams `form:"nok"`
	// Promotion code defined in NPR.
	Npr *PromotionCodeRestrictionsCurrencyOptionsNprParams `form:"npr"`
	// Promotion code defined in NZD.
	NZD *PromotionCodeRestrictionsCurrencyOptionsNZDParams `form:"nzd"`
	// Promotion code defined in OMR.
	Omr *PromotionCodeRestrictionsCurrencyOptionsOmrParams `form:"omr"`
	// Promotion code defined in PAB.
	Pab *PromotionCodeRestrictionsCurrencyOptionsPabParams `form:"pab"`
	// Promotion code defined in PEN.
	Pen *PromotionCodeRestrictionsCurrencyOptionsPenParams `form:"pen"`
	// Promotion code defined in PGK.
	Pgk *PromotionCodeRestrictionsCurrencyOptionsPgkParams `form:"pgk"`
	// Promotion code defined in PHP.
	Php *PromotionCodeRestrictionsCurrencyOptionsPhpParams `form:"php"`
	// Promotion code defined in PKR.
	Pkr *PromotionCodeRestrictionsCurrencyOptionsPkrParams `form:"pkr"`
	// Promotion code defined in PLN.
	Pln *PromotionCodeRestrictionsCurrencyOptionsPlnParams `form:"pln"`
	// Promotion code defined in PYG.
	Pyg *PromotionCodeRestrictionsCurrencyOptionsPygParams `form:"pyg"`
	// Promotion code defined in QAR.
	Qar *PromotionCodeRestrictionsCurrencyOptionsQarParams `form:"qar"`
	// Promotion code defined in RON.
	Ron *PromotionCodeRestrictionsCurrencyOptionsRonParams `form:"ron"`
	// Promotion code defined in RSD.
	Rsd *PromotionCodeRestrictionsCurrencyOptionsRsdParams `form:"rsd"`
	// Promotion code defined in RUB.
	Rub *PromotionCodeRestrictionsCurrencyOptionsRubParams `form:"rub"`
	// Promotion code defined in RWF.
	Rwf *PromotionCodeRestrictionsCurrencyOptionsRwfParams `form:"rwf"`
	// Promotion code defined in SAR.
	Sar *PromotionCodeRestrictionsCurrencyOptionsSarParams `form:"sar"`
	// Promotion code defined in SBD.
	Sbd *PromotionCodeRestrictionsCurrencyOptionsSbdParams `form:"sbd"`
	// Promotion code defined in SCR.
	Scr *PromotionCodeRestrictionsCurrencyOptionsScrParams `form:"scr"`
	// Promotion code defined in SEK.
	SEK *PromotionCodeRestrictionsCurrencyOptionsSEKParams `form:"sek"`
	// Promotion code defined in SGD.
	SGD *PromotionCodeRestrictionsCurrencyOptionsSGDParams `form:"sgd"`
	// Promotion code defined in SHP.
	Shp *PromotionCodeRestrictionsCurrencyOptionsShpParams `form:"shp"`
	// Promotion code defined in SLL.
	Sll *PromotionCodeRestrictionsCurrencyOptionsSllParams `form:"sll"`
	// Promotion code defined in SOS.
	Sos *PromotionCodeRestrictionsCurrencyOptionsSosParams `form:"sos"`
	// Promotion code defined in SRD.
	Srd *PromotionCodeRestrictionsCurrencyOptionsSrdParams `form:"srd"`
	// Promotion code defined in STD.
	Std *PromotionCodeRestrictionsCurrencyOptionsStdParams `form:"std"`
	// Promotion code defined in SZL.
	Szl *PromotionCodeRestrictionsCurrencyOptionsSzlParams `form:"szl"`
	// Promotion code defined in THB.
	Thb *PromotionCodeRestrictionsCurrencyOptionsThbParams `form:"thb"`
	// Promotion code defined in TJS.
	Tjs *PromotionCodeRestrictionsCurrencyOptionsTjsParams `form:"tjs"`
	// Promotion code defined in TND.
	Tnd *PromotionCodeRestrictionsCurrencyOptionsTndParams `form:"tnd"`
	// Promotion code defined in TOP.
	Top *PromotionCodeRestrictionsCurrencyOptionsTopParams `form:"top"`
	// Promotion code defined in TRY.
	Try *PromotionCodeRestrictionsCurrencyOptionsTryParams `form:"try"`
	// Promotion code defined in TTD.
	Ttd *PromotionCodeRestrictionsCurrencyOptionsTtdParams `form:"ttd"`
	// Promotion code defined in TWD.
	Twd *PromotionCodeRestrictionsCurrencyOptionsTwdParams `form:"twd"`
	// Promotion code defined in TZS.
	Tzs *PromotionCodeRestrictionsCurrencyOptionsTzsParams `form:"tzs"`
	// Promotion code defined in UAH.
	Uah *PromotionCodeRestrictionsCurrencyOptionsUahParams `form:"uah"`
	// Promotion code defined in UGX.
	Ugx *PromotionCodeRestrictionsCurrencyOptionsUgxParams `form:"ugx"`
	// Promotion code defined in USD.
	USD *PromotionCodeRestrictionsCurrencyOptionsUSDParams `form:"usd"`
	// Promotion code defined in USDC.
	Usdc *PromotionCodeRestrictionsCurrencyOptionsUsdcParams `form:"usdc"`
	// Promotion code defined in UYU.
	Uyu *PromotionCodeRestrictionsCurrencyOptionsUyuParams `form:"uyu"`
	// Promotion code defined in UZS.
	Uzs *PromotionCodeRestrictionsCurrencyOptionsUzsParams `form:"uzs"`
	// Promotion code defined in VND.
	Vnd *PromotionCodeRestrictionsCurrencyOptionsVndParams `form:"vnd"`
	// Promotion code defined in VUV.
	Vuv *PromotionCodeRestrictionsCurrencyOptionsVuvParams `form:"vuv"`
	// Promotion code defined in WST.
	Wst *PromotionCodeRestrictionsCurrencyOptionsWstParams `form:"wst"`
	// Promotion code defined in XAF.
	Xaf *PromotionCodeRestrictionsCurrencyOptionsXafParams `form:"xaf"`
	// Promotion code defined in XCD.
	Xcd *PromotionCodeRestrictionsCurrencyOptionsXcdParams `form:"xcd"`
	// Promotion code defined in XOF.
	Xof *PromotionCodeRestrictionsCurrencyOptionsXofParams `form:"xof"`
	// Promotion code defined in XPF.
	Xpf *PromotionCodeRestrictionsCurrencyOptionsXpfParams `form:"xpf"`
	// Promotion code defined in YER.
	Yer *PromotionCodeRestrictionsCurrencyOptionsYerParams `form:"yer"`
	// Promotion code defined in ZAR.
	Zar *PromotionCodeRestrictionsCurrencyOptionsZarParams `form:"zar"`
	// Promotion code defined in ZMW.
	Zmw *PromotionCodeRestrictionsCurrencyOptionsZmwParams `form:"zmw"`
}

// Settings that restrict the redemption of the promotion code.
type PromotionCodeRestrictionsParams struct {
	// Promotion codes defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *PromotionCodeRestrictionsCurrencyOptionsParams `form:"currency_options"`
	// A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices
	FirstTimeTransaction *bool `form:"first_time_transaction"`
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount *int64 `form:"minimum_amount"`
	// Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount
	MinimumAmountCurrency *string `form:"minimum_amount_currency"`
}

// Returns a list of your promotion codes.
type PromotionCodeListParams struct {
	ListParams `form:"*"`
	// Filter promotion codes by whether they are active.
	Active *bool `form:"active"`
	// Only return promotion codes that have this case-insensitive code.
	Code *string `form:"code"`
	// Only return promotion codes for this coupon.
	Coupon *string `form:"coupon"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return promotion codes that are restricted to this customer.
	Customer *string `form:"customer"`
}
type PromotionCodeRestrictionsCurrencyOptionsAed struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAfn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAll struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAmd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAng struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAoa struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsArs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAUD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAwg struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsAzn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBam struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBbd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBdt struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBgn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBhd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBif struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBmd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBnd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBob struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBrl struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBsd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBtn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBwp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsByn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsBzd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCAD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCdf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCHF struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsClp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCny struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCop struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCrc struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCve struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsCZK struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsDjf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsDKK struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsDop struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsDzd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsEgp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsEtb struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsEUR struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsFjd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsFkp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGBP struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGel struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGhs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGip struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGmd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGnf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGtq struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsGyd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsHKD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsHnl struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsHrk struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsHtg struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsHuf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsIdr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsIls struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsInr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsIsk struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsJmd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsJod struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsJpy struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKes struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKgs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKhr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKmf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKrw struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKwd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKyd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsKzt struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsLak struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsLbp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsLkr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsLrd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsLsl struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMad struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMdl struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMga struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMkd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMmk struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMnt struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMop struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMro struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMur struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMvr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMwk struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMxn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMYR struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsMzn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNad struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNgn struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNio struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNOK struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNpr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsNZD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsOmr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPab struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPen struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPgk struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPhp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPkr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPln struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsPyg struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsQar struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsRon struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsRsd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsRub struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsRwf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSar struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSbd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsScr struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSEK struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSGD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsShp struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSll struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSos struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSrd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsStd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsSzl struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsThb struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTjs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTnd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTop struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTry struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTtd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTwd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsTzs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUah struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUgx struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUSD struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUsdc struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUyu struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsUzs struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsVnd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsVuv struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsWst struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsXaf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsXcd struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsXof struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsXpf struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsYer struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsZar struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptionsZmw struct {
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
}
type PromotionCodeRestrictionsCurrencyOptions struct {
	Aed  *PromotionCodeRestrictionsCurrencyOptionsAed  `json:"aed"`
	Afn  *PromotionCodeRestrictionsCurrencyOptionsAfn  `json:"afn"`
	All  *PromotionCodeRestrictionsCurrencyOptionsAll  `json:"all"`
	Amd  *PromotionCodeRestrictionsCurrencyOptionsAmd  `json:"amd"`
	Ang  *PromotionCodeRestrictionsCurrencyOptionsAng  `json:"ang"`
	Aoa  *PromotionCodeRestrictionsCurrencyOptionsAoa  `json:"aoa"`
	Ars  *PromotionCodeRestrictionsCurrencyOptionsArs  `json:"ars"`
	AUD  *PromotionCodeRestrictionsCurrencyOptionsAUD  `json:"aud"`
	Awg  *PromotionCodeRestrictionsCurrencyOptionsAwg  `json:"awg"`
	Azn  *PromotionCodeRestrictionsCurrencyOptionsAzn  `json:"azn"`
	Bam  *PromotionCodeRestrictionsCurrencyOptionsBam  `json:"bam"`
	Bbd  *PromotionCodeRestrictionsCurrencyOptionsBbd  `json:"bbd"`
	Bdt  *PromotionCodeRestrictionsCurrencyOptionsBdt  `json:"bdt"`
	Bgn  *PromotionCodeRestrictionsCurrencyOptionsBgn  `json:"bgn"`
	Bhd  *PromotionCodeRestrictionsCurrencyOptionsBhd  `json:"bhd"`
	Bif  *PromotionCodeRestrictionsCurrencyOptionsBif  `json:"bif"`
	Bmd  *PromotionCodeRestrictionsCurrencyOptionsBmd  `json:"bmd"`
	Bnd  *PromotionCodeRestrictionsCurrencyOptionsBnd  `json:"bnd"`
	Bob  *PromotionCodeRestrictionsCurrencyOptionsBob  `json:"bob"`
	Brl  *PromotionCodeRestrictionsCurrencyOptionsBrl  `json:"brl"`
	Bsd  *PromotionCodeRestrictionsCurrencyOptionsBsd  `json:"bsd"`
	Btn  *PromotionCodeRestrictionsCurrencyOptionsBtn  `json:"btn"`
	Bwp  *PromotionCodeRestrictionsCurrencyOptionsBwp  `json:"bwp"`
	Byn  *PromotionCodeRestrictionsCurrencyOptionsByn  `json:"byn"`
	Bzd  *PromotionCodeRestrictionsCurrencyOptionsBzd  `json:"bzd"`
	CAD  *PromotionCodeRestrictionsCurrencyOptionsCAD  `json:"cad"`
	Cdf  *PromotionCodeRestrictionsCurrencyOptionsCdf  `json:"cdf"`
	CHF  *PromotionCodeRestrictionsCurrencyOptionsCHF  `json:"chf"`
	Clp  *PromotionCodeRestrictionsCurrencyOptionsClp  `json:"clp"`
	Cny  *PromotionCodeRestrictionsCurrencyOptionsCny  `json:"cny"`
	Cop  *PromotionCodeRestrictionsCurrencyOptionsCop  `json:"cop"`
	Crc  *PromotionCodeRestrictionsCurrencyOptionsCrc  `json:"crc"`
	Cve  *PromotionCodeRestrictionsCurrencyOptionsCve  `json:"cve"`
	CZK  *PromotionCodeRestrictionsCurrencyOptionsCZK  `json:"czk"`
	Djf  *PromotionCodeRestrictionsCurrencyOptionsDjf  `json:"djf"`
	DKK  *PromotionCodeRestrictionsCurrencyOptionsDKK  `json:"dkk"`
	Dop  *PromotionCodeRestrictionsCurrencyOptionsDop  `json:"dop"`
	Dzd  *PromotionCodeRestrictionsCurrencyOptionsDzd  `json:"dzd"`
	Egp  *PromotionCodeRestrictionsCurrencyOptionsEgp  `json:"egp"`
	Etb  *PromotionCodeRestrictionsCurrencyOptionsEtb  `json:"etb"`
	EUR  *PromotionCodeRestrictionsCurrencyOptionsEUR  `json:"eur"`
	Fjd  *PromotionCodeRestrictionsCurrencyOptionsFjd  `json:"fjd"`
	Fkp  *PromotionCodeRestrictionsCurrencyOptionsFkp  `json:"fkp"`
	GBP  *PromotionCodeRestrictionsCurrencyOptionsGBP  `json:"gbp"`
	Gel  *PromotionCodeRestrictionsCurrencyOptionsGel  `json:"gel"`
	Ghs  *PromotionCodeRestrictionsCurrencyOptionsGhs  `json:"ghs"`
	Gip  *PromotionCodeRestrictionsCurrencyOptionsGip  `json:"gip"`
	Gmd  *PromotionCodeRestrictionsCurrencyOptionsGmd  `json:"gmd"`
	Gnf  *PromotionCodeRestrictionsCurrencyOptionsGnf  `json:"gnf"`
	Gtq  *PromotionCodeRestrictionsCurrencyOptionsGtq  `json:"gtq"`
	Gyd  *PromotionCodeRestrictionsCurrencyOptionsGyd  `json:"gyd"`
	HKD  *PromotionCodeRestrictionsCurrencyOptionsHKD  `json:"hkd"`
	Hnl  *PromotionCodeRestrictionsCurrencyOptionsHnl  `json:"hnl"`
	Hrk  *PromotionCodeRestrictionsCurrencyOptionsHrk  `json:"hrk"`
	Htg  *PromotionCodeRestrictionsCurrencyOptionsHtg  `json:"htg"`
	Huf  *PromotionCodeRestrictionsCurrencyOptionsHuf  `json:"huf"`
	Idr  *PromotionCodeRestrictionsCurrencyOptionsIdr  `json:"idr"`
	Ils  *PromotionCodeRestrictionsCurrencyOptionsIls  `json:"ils"`
	Inr  *PromotionCodeRestrictionsCurrencyOptionsInr  `json:"inr"`
	Isk  *PromotionCodeRestrictionsCurrencyOptionsIsk  `json:"isk"`
	Jmd  *PromotionCodeRestrictionsCurrencyOptionsJmd  `json:"jmd"`
	Jod  *PromotionCodeRestrictionsCurrencyOptionsJod  `json:"jod"`
	Jpy  *PromotionCodeRestrictionsCurrencyOptionsJpy  `json:"jpy"`
	Kes  *PromotionCodeRestrictionsCurrencyOptionsKes  `json:"kes"`
	Kgs  *PromotionCodeRestrictionsCurrencyOptionsKgs  `json:"kgs"`
	Khr  *PromotionCodeRestrictionsCurrencyOptionsKhr  `json:"khr"`
	Kmf  *PromotionCodeRestrictionsCurrencyOptionsKmf  `json:"kmf"`
	Krw  *PromotionCodeRestrictionsCurrencyOptionsKrw  `json:"krw"`
	Kwd  *PromotionCodeRestrictionsCurrencyOptionsKwd  `json:"kwd"`
	Kyd  *PromotionCodeRestrictionsCurrencyOptionsKyd  `json:"kyd"`
	Kzt  *PromotionCodeRestrictionsCurrencyOptionsKzt  `json:"kzt"`
	Lak  *PromotionCodeRestrictionsCurrencyOptionsLak  `json:"lak"`
	Lbp  *PromotionCodeRestrictionsCurrencyOptionsLbp  `json:"lbp"`
	Lkr  *PromotionCodeRestrictionsCurrencyOptionsLkr  `json:"lkr"`
	Lrd  *PromotionCodeRestrictionsCurrencyOptionsLrd  `json:"lrd"`
	Lsl  *PromotionCodeRestrictionsCurrencyOptionsLsl  `json:"lsl"`
	Mad  *PromotionCodeRestrictionsCurrencyOptionsMad  `json:"mad"`
	Mdl  *PromotionCodeRestrictionsCurrencyOptionsMdl  `json:"mdl"`
	Mga  *PromotionCodeRestrictionsCurrencyOptionsMga  `json:"mga"`
	Mkd  *PromotionCodeRestrictionsCurrencyOptionsMkd  `json:"mkd"`
	Mmk  *PromotionCodeRestrictionsCurrencyOptionsMmk  `json:"mmk"`
	Mnt  *PromotionCodeRestrictionsCurrencyOptionsMnt  `json:"mnt"`
	Mop  *PromotionCodeRestrictionsCurrencyOptionsMop  `json:"mop"`
	Mro  *PromotionCodeRestrictionsCurrencyOptionsMro  `json:"mro"`
	Mur  *PromotionCodeRestrictionsCurrencyOptionsMur  `json:"mur"`
	Mvr  *PromotionCodeRestrictionsCurrencyOptionsMvr  `json:"mvr"`
	Mwk  *PromotionCodeRestrictionsCurrencyOptionsMwk  `json:"mwk"`
	Mxn  *PromotionCodeRestrictionsCurrencyOptionsMxn  `json:"mxn"`
	MYR  *PromotionCodeRestrictionsCurrencyOptionsMYR  `json:"myr"`
	Mzn  *PromotionCodeRestrictionsCurrencyOptionsMzn  `json:"mzn"`
	Nad  *PromotionCodeRestrictionsCurrencyOptionsNad  `json:"nad"`
	Ngn  *PromotionCodeRestrictionsCurrencyOptionsNgn  `json:"ngn"`
	Nio  *PromotionCodeRestrictionsCurrencyOptionsNio  `json:"nio"`
	NOK  *PromotionCodeRestrictionsCurrencyOptionsNOK  `json:"nok"`
	Npr  *PromotionCodeRestrictionsCurrencyOptionsNpr  `json:"npr"`
	NZD  *PromotionCodeRestrictionsCurrencyOptionsNZD  `json:"nzd"`
	Omr  *PromotionCodeRestrictionsCurrencyOptionsOmr  `json:"omr"`
	Pab  *PromotionCodeRestrictionsCurrencyOptionsPab  `json:"pab"`
	Pen  *PromotionCodeRestrictionsCurrencyOptionsPen  `json:"pen"`
	Pgk  *PromotionCodeRestrictionsCurrencyOptionsPgk  `json:"pgk"`
	Php  *PromotionCodeRestrictionsCurrencyOptionsPhp  `json:"php"`
	Pkr  *PromotionCodeRestrictionsCurrencyOptionsPkr  `json:"pkr"`
	Pln  *PromotionCodeRestrictionsCurrencyOptionsPln  `json:"pln"`
	Pyg  *PromotionCodeRestrictionsCurrencyOptionsPyg  `json:"pyg"`
	Qar  *PromotionCodeRestrictionsCurrencyOptionsQar  `json:"qar"`
	Ron  *PromotionCodeRestrictionsCurrencyOptionsRon  `json:"ron"`
	Rsd  *PromotionCodeRestrictionsCurrencyOptionsRsd  `json:"rsd"`
	Rub  *PromotionCodeRestrictionsCurrencyOptionsRub  `json:"rub"`
	Rwf  *PromotionCodeRestrictionsCurrencyOptionsRwf  `json:"rwf"`
	Sar  *PromotionCodeRestrictionsCurrencyOptionsSar  `json:"sar"`
	Sbd  *PromotionCodeRestrictionsCurrencyOptionsSbd  `json:"sbd"`
	Scr  *PromotionCodeRestrictionsCurrencyOptionsScr  `json:"scr"`
	SEK  *PromotionCodeRestrictionsCurrencyOptionsSEK  `json:"sek"`
	SGD  *PromotionCodeRestrictionsCurrencyOptionsSGD  `json:"sgd"`
	Shp  *PromotionCodeRestrictionsCurrencyOptionsShp  `json:"shp"`
	Sll  *PromotionCodeRestrictionsCurrencyOptionsSll  `json:"sll"`
	Sos  *PromotionCodeRestrictionsCurrencyOptionsSos  `json:"sos"`
	Srd  *PromotionCodeRestrictionsCurrencyOptionsSrd  `json:"srd"`
	Std  *PromotionCodeRestrictionsCurrencyOptionsStd  `json:"std"`
	Szl  *PromotionCodeRestrictionsCurrencyOptionsSzl  `json:"szl"`
	Thb  *PromotionCodeRestrictionsCurrencyOptionsThb  `json:"thb"`
	Tjs  *PromotionCodeRestrictionsCurrencyOptionsTjs  `json:"tjs"`
	Tnd  *PromotionCodeRestrictionsCurrencyOptionsTnd  `json:"tnd"`
	Top  *PromotionCodeRestrictionsCurrencyOptionsTop  `json:"top"`
	Try  *PromotionCodeRestrictionsCurrencyOptionsTry  `json:"try"`
	Ttd  *PromotionCodeRestrictionsCurrencyOptionsTtd  `json:"ttd"`
	Twd  *PromotionCodeRestrictionsCurrencyOptionsTwd  `json:"twd"`
	Tzs  *PromotionCodeRestrictionsCurrencyOptionsTzs  `json:"tzs"`
	Uah  *PromotionCodeRestrictionsCurrencyOptionsUah  `json:"uah"`
	Ugx  *PromotionCodeRestrictionsCurrencyOptionsUgx  `json:"ugx"`
	USD  *PromotionCodeRestrictionsCurrencyOptionsUSD  `json:"usd"`
	Usdc *PromotionCodeRestrictionsCurrencyOptionsUsdc `json:"usdc"`
	Uyu  *PromotionCodeRestrictionsCurrencyOptionsUyu  `json:"uyu"`
	Uzs  *PromotionCodeRestrictionsCurrencyOptionsUzs  `json:"uzs"`
	Vnd  *PromotionCodeRestrictionsCurrencyOptionsVnd  `json:"vnd"`
	Vuv  *PromotionCodeRestrictionsCurrencyOptionsVuv  `json:"vuv"`
	Wst  *PromotionCodeRestrictionsCurrencyOptionsWst  `json:"wst"`
	Xaf  *PromotionCodeRestrictionsCurrencyOptionsXaf  `json:"xaf"`
	Xcd  *PromotionCodeRestrictionsCurrencyOptionsXcd  `json:"xcd"`
	Xof  *PromotionCodeRestrictionsCurrencyOptionsXof  `json:"xof"`
	Xpf  *PromotionCodeRestrictionsCurrencyOptionsXpf  `json:"xpf"`
	Yer  *PromotionCodeRestrictionsCurrencyOptionsYer  `json:"yer"`
	Zar  *PromotionCodeRestrictionsCurrencyOptionsZar  `json:"zar"`
	Zmw  *PromotionCodeRestrictionsCurrencyOptionsZmw  `json:"zmw"`
}
type PromotionCodeRestrictions struct {
	CurrencyOptions *PromotionCodeRestrictionsCurrencyOptions `json:"currency_options"`
	// A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices
	FirstTimeTransaction bool `json:"first_time_transaction"`
	// Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
	MinimumAmount int64 `json:"minimum_amount"`
	// Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount
	MinimumAmountCurrency Currency `json:"minimum_amount_currency"`
}

// A Promotion Code represents a customer-redeemable code for a [coupon](https://stripe.com/docs/api#coupons). It can be used to
// create multiple codes for a single coupon.
type PromotionCode struct {
	APIResource
	// Whether the promotion code is currently active. A promotion code is only active if the coupon is also valid.
	Active bool `json:"active"`
	// The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for each customer.
	Code string `json:"code"`
	// A coupon contains information about a percent-off or amount-off discount you
	// might want to apply to a customer. Coupons may be applied to [subscriptions](https://stripe.com/docs/api#subscriptions), [invoices](https://stripe.com/docs/api#invoices),
	// [checkout sessions](https://stripe.com/docs/api/checkout/sessions), [quotes](https://stripe.com/docs/api#quotes), and more. Coupons do not work with conventional one-off [charges](https://stripe.com/docs/api#create_charge) or [payment intents](https://stripe.com/docs/api/payment_intents).
	Coupon *Coupon `json:"coupon"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The customer that this promotion code can be used by.
	Customer *Customer `json:"customer"`
	// Date at which the promotion code can no longer be redeemed.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Maximum number of times this promotion code can be redeemed.
	MaxRedemptions int64 `json:"max_redemptions"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object       string                     `json:"object"`
	Restrictions *PromotionCodeRestrictions `json:"restrictions"`
	// Number of times this promotion code has been used.
	TimesRedeemed int64 `json:"times_redeemed"`
}

// PromotionCodeList is a list of PromotionCodes as retrieved from a list endpoint.
type PromotionCodeList struct {
	APIResource
	ListMeta
	Data []*PromotionCode `json:"data"`
}

// UnmarshalJSON handles deserialization of a PromotionCode.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PromotionCode) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type promotionCode PromotionCode
	var v promotionCode
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PromotionCode(v)
	return nil
}
