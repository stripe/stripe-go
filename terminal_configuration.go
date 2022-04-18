//
//
// File generated from our OpenAPI spec
//
//

package stripe

// An object containing device type specific settings for BBPOS WisePOS E readers
type TerminalConfigurationBBPOSWisePOSEParams struct {
	// A File ID representing an image you would like displayed on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Tipping configuration for AUD
type TerminalConfigurationTippingAudParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CAD
type TerminalConfigurationTippingCADParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CHF
type TerminalConfigurationTippingChfParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for DKK
type TerminalConfigurationTippingDkkParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for EUR
type TerminalConfigurationTippingEurParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GBP
type TerminalConfigurationTippingGbpParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HKD
type TerminalConfigurationTippingHkdParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MYR
type TerminalConfigurationTippingMyrParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NOK
type TerminalConfigurationTippingNokParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NZD
type TerminalConfigurationTippingNzdParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SEK
type TerminalConfigurationTippingSekParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SGD
type TerminalConfigurationTippingSgdParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for USD
type TerminalConfigurationTippingUSDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configurations for readers supporting on-reader tips
type TerminalConfigurationTippingParams struct {
	// Tipping configuration for AUD
	Aud *TerminalConfigurationTippingAudParams `form:"aud"`
	// Tipping configuration for CAD
	CAD *TerminalConfigurationTippingCADParams `form:"cad"`
	// Tipping configuration for CHF
	Chf *TerminalConfigurationTippingChfParams `form:"chf"`
	// Tipping configuration for DKK
	Dkk *TerminalConfigurationTippingDkkParams `form:"dkk"`
	// Tipping configuration for EUR
	Eur *TerminalConfigurationTippingEurParams `form:"eur"`
	// Tipping configuration for GBP
	Gbp *TerminalConfigurationTippingGbpParams `form:"gbp"`
	// Tipping configuration for HKD
	Hkd *TerminalConfigurationTippingHkdParams `form:"hkd"`
	// Tipping configuration for MYR
	Myr *TerminalConfigurationTippingMyrParams `form:"myr"`
	// Tipping configuration for NOK
	Nok *TerminalConfigurationTippingNokParams `form:"nok"`
	// Tipping configuration for NZD
	Nzd *TerminalConfigurationTippingNzdParams `form:"nzd"`
	// Tipping configuration for SEK
	Sek *TerminalConfigurationTippingSekParams `form:"sek"`
	// Tipping configuration for SGD
	Sgd *TerminalConfigurationTippingSgdParams `form:"sgd"`
	// Tipping configuration for USD
	USD *TerminalConfigurationTippingUSDParams `form:"usd"`
}

// An object containing device type specific settings for Verifone P400 readers
type TerminalConfigurationVerifoneP400Params struct {
	// A File ID representing an image you would like displayed on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Creates a new Configuration object.
type TerminalConfigurationParams struct {
	Params `form:"*"`
	// An object containing device type specific settings for BBPOS WisePOS E readers
	BBPOSWisePOSE *TerminalConfigurationBBPOSWisePOSEParams `form:"bbpos_wisepos_e"`
	// Tipping configurations for readers supporting on-reader tips
	Tipping *TerminalConfigurationTippingParams `form:"tipping"`
	// An object containing device type specific settings for Verifone P400 readers
	VerifoneP400 *TerminalConfigurationVerifoneP400Params `form:"verifone_p400"`
}

// Returns a list of Configuration objects.
type TerminalConfigurationListParams struct {
	ListParams `form:"*"`
	// if present, only return the account default or non-default configurations.
	IsAccountDefault *bool `form:"is_account_default"`
}
type TerminalConfigurationBBPOSWisePOSE struct {
	// A File ID representing an image you would like displayed on the reader.
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationTippingAud struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingCAD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingChf struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingDkk struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingEur struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingGbp struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingHkd struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingMyr struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingNok struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingNzd struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingSek struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingSgd struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingUSD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTipping struct {
	Aud *TerminalConfigurationTippingAud `json:"aud"`
	CAD *TerminalConfigurationTippingCAD `json:"cad"`
	Chf *TerminalConfigurationTippingChf `json:"chf"`
	Dkk *TerminalConfigurationTippingDkk `json:"dkk"`
	Eur *TerminalConfigurationTippingEur `json:"eur"`
	Gbp *TerminalConfigurationTippingGbp `json:"gbp"`
	Hkd *TerminalConfigurationTippingHkd `json:"hkd"`
	Myr *TerminalConfigurationTippingMyr `json:"myr"`
	Nok *TerminalConfigurationTippingNok `json:"nok"`
	Nzd *TerminalConfigurationTippingNzd `json:"nzd"`
	Sek *TerminalConfigurationTippingSek `json:"sek"`
	Sgd *TerminalConfigurationTippingSgd `json:"sgd"`
	USD *TerminalConfigurationTippingUSD `json:"usd"`
}
type TerminalConfigurationVerifoneP400 struct {
	// A File ID representing an image you would like displayed on the reader.
	Splashscreen *File `json:"splashscreen"`
}

// A Configurations object represents how features should be configured for terminal readers.
type TerminalConfiguration struct {
	APIResource
	BBPOSWisePOSE *TerminalConfigurationBBPOSWisePOSE `json:"bbpos_wisepos_e"`
	Deleted       bool                                `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether this Configuration is the default for your account
	IsAccountDefault bool `json:"is_account_default"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object       string                             `json:"object"`
	Tipping      *TerminalConfigurationTipping      `json:"tipping"`
	VerifoneP400 *TerminalConfigurationVerifoneP400 `json:"verifone_p400"`
}

// TerminalConfigurationList is a list of Configurations as retrieved from a list endpoint.
type TerminalConfigurationList struct {
	APIResource
	ListMeta
	Data []*TerminalConfiguration `json:"data"`
}
