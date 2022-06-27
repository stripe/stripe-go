//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// A unit of time.
type ShippingRateDeliveryEstimateMaximumUnit string

// List of values that ShippingRateDeliveryEstimateMaximumUnit can take
const (
	ShippingRateDeliveryEstimateMaximumUnitBusinessDay ShippingRateDeliveryEstimateMaximumUnit = "business_day"
	ShippingRateDeliveryEstimateMaximumUnitDay         ShippingRateDeliveryEstimateMaximumUnit = "day"
	ShippingRateDeliveryEstimateMaximumUnitHour        ShippingRateDeliveryEstimateMaximumUnit = "hour"
	ShippingRateDeliveryEstimateMaximumUnitMonth       ShippingRateDeliveryEstimateMaximumUnit = "month"
	ShippingRateDeliveryEstimateMaximumUnitWeek        ShippingRateDeliveryEstimateMaximumUnit = "week"
)

// A unit of time.
type ShippingRateDeliveryEstimateMinimumUnit string

// List of values that ShippingRateDeliveryEstimateMinimumUnit can take
const (
	ShippingRateDeliveryEstimateMinimumUnitBusinessDay ShippingRateDeliveryEstimateMinimumUnit = "business_day"
	ShippingRateDeliveryEstimateMinimumUnitDay         ShippingRateDeliveryEstimateMinimumUnit = "day"
	ShippingRateDeliveryEstimateMinimumUnitHour        ShippingRateDeliveryEstimateMinimumUnit = "hour"
	ShippingRateDeliveryEstimateMinimumUnitMonth       ShippingRateDeliveryEstimateMinimumUnit = "month"
	ShippingRateDeliveryEstimateMinimumUnitWeek        ShippingRateDeliveryEstimateMinimumUnit = "week"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAedTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAedTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAedTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAfnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAfnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAfnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAllTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAllTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAllTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAmdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAmdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAmdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAngTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAngTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAngTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAoaTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAoaTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAoaTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsArsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsArsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsArsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAUDTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAUDTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAUDTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAwgTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAwgTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAwgTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsAznTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsAznTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsAznTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBamTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBamTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBamTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBbdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBbdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBbdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBdtTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBdtTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBdtTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBgnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBgnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBgnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBhdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBhdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBhdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBifTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBifTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBifTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBmdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBmdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBmdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBndTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBndTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBndTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBobTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBobTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBobTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBrlTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBrlTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBrlTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBsdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBsdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBsdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBtnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBtnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBtnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBwpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBwpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBwpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBynTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBynTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBynTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsBzdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsBzdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsBzdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCADTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCADTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCADTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCdfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCdfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCdfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCHFTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCHFTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCHFTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsClpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsClpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsClpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCnyTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCnyTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCnyTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCopTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCopTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCopTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCrcTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCrcTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCrcTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCveTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCveTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCveTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsCZKTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsCZKTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsCZKTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsDjfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsDjfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsDjfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsDKKTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsDKKTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsDKKTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsDopTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsDopTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsDopTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsDzdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsDzdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsDzdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsEgpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsEgpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsEgpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsEtbTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsEtbTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsEtbTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsEURTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsEURTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsEURTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsFjdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsFjdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsFjdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsFkpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsFkpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsFkpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGBPTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGBPTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGBPTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGelTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGelTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGelTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGhsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGhsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGhsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGipTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGipTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGipTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGmdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGmdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGmdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGnfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGnfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGnfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGtqTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGtqTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGtqTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsGydTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsGydTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsGydTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsHKDTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsHKDTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsHKDTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsHnlTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsHnlTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsHnlTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsHrkTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsHrkTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsHrkTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsHtgTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsHtgTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsHtgTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsHufTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsHufTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsHufTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsIdrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsIdrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsIdrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsIlsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsIlsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsIlsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsInrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsInrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsInrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsIskTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsIskTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsIskTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsJmdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsJmdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsJmdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsJodTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsJodTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsJodTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsJpyTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsJpyTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsJpyTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKesTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKesTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKesTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKgsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKgsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKgsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKhrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKhrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKhrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKmfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKmfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKmfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKrwTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKrwTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKrwTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKwdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKwdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKwdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKydTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKydTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKydTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsKztTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsKztTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsKztTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsLakTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsLakTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsLakTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsLbpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsLbpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsLbpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsLkrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsLkrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsLkrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsLrdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsLrdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsLrdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsLslTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsLslTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsLslTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMadTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMadTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMadTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMdlTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMdlTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMdlTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMgaTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMgaTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMgaTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMkdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMkdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMkdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMmkTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMmkTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMmkTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMntTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMntTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMntTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMopTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMopTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMopTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMroTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMroTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMroTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMurTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMurTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMurTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMvrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMvrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMvrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMwkTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMwkTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMwkTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMxnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMxnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMxnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMYRTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMYRTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMYRTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsMznTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsMznTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsMznTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNadTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNadTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNadTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNgnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNgnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNgnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNioTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNioTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNioTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNOKTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNOKTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNOKTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNprTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNprTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNprTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsNZDTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsNZDTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsNZDTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsOmrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsOmrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsOmrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPabTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPabTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPabTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPenTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPenTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPenTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPgkTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPgkTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPgkTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPhpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPhpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPhpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPkrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPkrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPkrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPlnTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPlnTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPlnTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsPygTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsPygTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsPygTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsQarTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsQarTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsQarTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsRonTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsRonTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsRonTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsRsdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsRsdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsRsdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsRubTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsRubTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsRubTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsRwfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsRwfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsRwfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSarTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSarTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSarTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSbdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSbdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSbdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsScrTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsScrTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsScrTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSEKTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSEKTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSEKTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSGDTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSGDTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSGDTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsShpTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsShpTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsShpTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSllTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSllTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSllTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSosTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSosTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSosTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSrdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSrdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSrdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsStdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsStdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsStdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsSzlTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsSzlTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsSzlTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsThbTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsThbTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsThbTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTjsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTjsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTjsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTndTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTndTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTndTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTopTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTopTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTopTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTryTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTryTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTryTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTtdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTtdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTtdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTwdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTwdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTwdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsTzsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsTzsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsTzsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUahTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUahTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUahTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUgxTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUgxTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUgxTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUSDTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUSDTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUSDTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUyuTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUyuTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUyuTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsUzsTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsUzsTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsUzsTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsVndTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsVndTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsVndTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsVuvTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsVuvTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsVuvTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsWstTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsWstTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsWstTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsXafTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsXafTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsXafTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsXcdTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsXcdTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsXcdTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsXofTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsXofTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsXofTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsXpfTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsXpfTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsXpfTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsYerTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsYerTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsYerTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsZarTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsZarTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsZarTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior string

// List of values that ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior can take
const (
	ShippingRateFixedAmountCurrencyOptionsZmwTaxBehaviorExclusive   ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior = "exclusive"
	ShippingRateFixedAmountCurrencyOptionsZmwTaxBehaviorInclusive   ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior = "inclusive"
	ShippingRateFixedAmountCurrencyOptionsZmwTaxBehaviorUnspecified ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior = "unspecified"
)

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
type ShippingRateTaxBehavior string

// List of values that ShippingRateTaxBehavior can take
const (
	ShippingRateTaxBehaviorExclusive   ShippingRateTaxBehavior = "exclusive"
	ShippingRateTaxBehaviorInclusive   ShippingRateTaxBehavior = "inclusive"
	ShippingRateTaxBehaviorUnspecified ShippingRateTaxBehavior = "unspecified"
)

// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
type ShippingRateType string

// List of values that ShippingRateType can take
const (
	ShippingRateTypeFixedAmount ShippingRateType = "fixed_amount"
)

// Returns a list of your shipping rates.
type ShippingRateListParams struct {
	ListParams `form:"*"`
	// Only return shipping rates that are active or inactive.
	Active *bool `form:"active"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return shipping rates for the given currency.
	Currency *string `form:"currency"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type ShippingRateDeliveryEstimateMaximumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type ShippingRateDeliveryEstimateMinimumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type ShippingRateDeliveryEstimateParams struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *ShippingRateDeliveryEstimateMaximumParams `form:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *ShippingRateDeliveryEstimateMinimumParams `form:"minimum"`
}

// Shipping rate defined in AED.
type ShippingRateFixedAmountCurrencyOptionsAedParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AFN.
type ShippingRateFixedAmountCurrencyOptionsAfnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ALL.
type ShippingRateFixedAmountCurrencyOptionsAllParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AMD.
type ShippingRateFixedAmountCurrencyOptionsAmdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ANG.
type ShippingRateFixedAmountCurrencyOptionsAngParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AOA.
type ShippingRateFixedAmountCurrencyOptionsAoaParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ARS.
type ShippingRateFixedAmountCurrencyOptionsArsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AUD.
type ShippingRateFixedAmountCurrencyOptionsAUDParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AWG.
type ShippingRateFixedAmountCurrencyOptionsAwgParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in AZN.
type ShippingRateFixedAmountCurrencyOptionsAznParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BAM.
type ShippingRateFixedAmountCurrencyOptionsBamParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BBD.
type ShippingRateFixedAmountCurrencyOptionsBbdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BDT.
type ShippingRateFixedAmountCurrencyOptionsBdtParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BGN.
type ShippingRateFixedAmountCurrencyOptionsBgnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BHD.
type ShippingRateFixedAmountCurrencyOptionsBhdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BIF.
type ShippingRateFixedAmountCurrencyOptionsBifParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BMD.
type ShippingRateFixedAmountCurrencyOptionsBmdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BND.
type ShippingRateFixedAmountCurrencyOptionsBndParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BOB.
type ShippingRateFixedAmountCurrencyOptionsBobParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BRL.
type ShippingRateFixedAmountCurrencyOptionsBrlParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BSD.
type ShippingRateFixedAmountCurrencyOptionsBsdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BTN.
type ShippingRateFixedAmountCurrencyOptionsBtnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BWP.
type ShippingRateFixedAmountCurrencyOptionsBwpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BYN.
type ShippingRateFixedAmountCurrencyOptionsBynParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in BZD.
type ShippingRateFixedAmountCurrencyOptionsBzdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CAD.
type ShippingRateFixedAmountCurrencyOptionsCADParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CDF.
type ShippingRateFixedAmountCurrencyOptionsCdfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CHF.
type ShippingRateFixedAmountCurrencyOptionsCHFParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CLP.
type ShippingRateFixedAmountCurrencyOptionsClpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CNY.
type ShippingRateFixedAmountCurrencyOptionsCnyParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in COP.
type ShippingRateFixedAmountCurrencyOptionsCopParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CRC.
type ShippingRateFixedAmountCurrencyOptionsCrcParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CVE.
type ShippingRateFixedAmountCurrencyOptionsCveParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in CZK.
type ShippingRateFixedAmountCurrencyOptionsCZKParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in DJF.
type ShippingRateFixedAmountCurrencyOptionsDjfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in DKK.
type ShippingRateFixedAmountCurrencyOptionsDKKParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in DOP.
type ShippingRateFixedAmountCurrencyOptionsDopParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in DZD.
type ShippingRateFixedAmountCurrencyOptionsDzdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in EGP.
type ShippingRateFixedAmountCurrencyOptionsEgpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ETB.
type ShippingRateFixedAmountCurrencyOptionsEtbParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in EUR.
type ShippingRateFixedAmountCurrencyOptionsEURParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in FJD.
type ShippingRateFixedAmountCurrencyOptionsFjdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in FKP.
type ShippingRateFixedAmountCurrencyOptionsFkpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GBP.
type ShippingRateFixedAmountCurrencyOptionsGBPParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GEL.
type ShippingRateFixedAmountCurrencyOptionsGelParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GHS.
type ShippingRateFixedAmountCurrencyOptionsGhsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GIP.
type ShippingRateFixedAmountCurrencyOptionsGipParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GMD.
type ShippingRateFixedAmountCurrencyOptionsGmdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GNF.
type ShippingRateFixedAmountCurrencyOptionsGnfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GTQ.
type ShippingRateFixedAmountCurrencyOptionsGtqParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in GYD.
type ShippingRateFixedAmountCurrencyOptionsGydParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in HKD.
type ShippingRateFixedAmountCurrencyOptionsHKDParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in HNL.
type ShippingRateFixedAmountCurrencyOptionsHnlParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in HRK.
type ShippingRateFixedAmountCurrencyOptionsHrkParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in HTG.
type ShippingRateFixedAmountCurrencyOptionsHtgParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in HUF.
type ShippingRateFixedAmountCurrencyOptionsHufParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in IDR.
type ShippingRateFixedAmountCurrencyOptionsIdrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ILS.
type ShippingRateFixedAmountCurrencyOptionsIlsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in INR.
type ShippingRateFixedAmountCurrencyOptionsInrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ISK.
type ShippingRateFixedAmountCurrencyOptionsIskParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in JMD.
type ShippingRateFixedAmountCurrencyOptionsJmdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in JOD.
type ShippingRateFixedAmountCurrencyOptionsJodParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in JPY.
type ShippingRateFixedAmountCurrencyOptionsJpyParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KES.
type ShippingRateFixedAmountCurrencyOptionsKesParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KGS.
type ShippingRateFixedAmountCurrencyOptionsKgsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KHR.
type ShippingRateFixedAmountCurrencyOptionsKhrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KMF.
type ShippingRateFixedAmountCurrencyOptionsKmfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KRW.
type ShippingRateFixedAmountCurrencyOptionsKrwParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KWD.
type ShippingRateFixedAmountCurrencyOptionsKwdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KYD.
type ShippingRateFixedAmountCurrencyOptionsKydParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in KZT.
type ShippingRateFixedAmountCurrencyOptionsKztParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in LAK.
type ShippingRateFixedAmountCurrencyOptionsLakParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in LBP.
type ShippingRateFixedAmountCurrencyOptionsLbpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in LKR.
type ShippingRateFixedAmountCurrencyOptionsLkrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in LRD.
type ShippingRateFixedAmountCurrencyOptionsLrdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in LSL.
type ShippingRateFixedAmountCurrencyOptionsLslParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MAD.
type ShippingRateFixedAmountCurrencyOptionsMadParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MDL.
type ShippingRateFixedAmountCurrencyOptionsMdlParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MGA.
type ShippingRateFixedAmountCurrencyOptionsMgaParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MKD.
type ShippingRateFixedAmountCurrencyOptionsMkdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MMK.
type ShippingRateFixedAmountCurrencyOptionsMmkParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MNT.
type ShippingRateFixedAmountCurrencyOptionsMntParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MOP.
type ShippingRateFixedAmountCurrencyOptionsMopParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MRO.
type ShippingRateFixedAmountCurrencyOptionsMroParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MUR.
type ShippingRateFixedAmountCurrencyOptionsMurParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MVR.
type ShippingRateFixedAmountCurrencyOptionsMvrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MWK.
type ShippingRateFixedAmountCurrencyOptionsMwkParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MXN.
type ShippingRateFixedAmountCurrencyOptionsMxnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MYR.
type ShippingRateFixedAmountCurrencyOptionsMYRParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in MZN.
type ShippingRateFixedAmountCurrencyOptionsMznParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NAD.
type ShippingRateFixedAmountCurrencyOptionsNadParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NGN.
type ShippingRateFixedAmountCurrencyOptionsNgnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NIO.
type ShippingRateFixedAmountCurrencyOptionsNioParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NOK.
type ShippingRateFixedAmountCurrencyOptionsNOKParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NPR.
type ShippingRateFixedAmountCurrencyOptionsNprParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in NZD.
type ShippingRateFixedAmountCurrencyOptionsNZDParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in OMR.
type ShippingRateFixedAmountCurrencyOptionsOmrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PAB.
type ShippingRateFixedAmountCurrencyOptionsPabParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PEN.
type ShippingRateFixedAmountCurrencyOptionsPenParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PGK.
type ShippingRateFixedAmountCurrencyOptionsPgkParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PHP.
type ShippingRateFixedAmountCurrencyOptionsPhpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PKR.
type ShippingRateFixedAmountCurrencyOptionsPkrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PLN.
type ShippingRateFixedAmountCurrencyOptionsPlnParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in PYG.
type ShippingRateFixedAmountCurrencyOptionsPygParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in QAR.
type ShippingRateFixedAmountCurrencyOptionsQarParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in RON.
type ShippingRateFixedAmountCurrencyOptionsRonParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in RSD.
type ShippingRateFixedAmountCurrencyOptionsRsdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in RUB.
type ShippingRateFixedAmountCurrencyOptionsRubParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in RWF.
type ShippingRateFixedAmountCurrencyOptionsRwfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SAR.
type ShippingRateFixedAmountCurrencyOptionsSarParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SBD.
type ShippingRateFixedAmountCurrencyOptionsSbdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SCR.
type ShippingRateFixedAmountCurrencyOptionsScrParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SEK.
type ShippingRateFixedAmountCurrencyOptionsSEKParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SGD.
type ShippingRateFixedAmountCurrencyOptionsSGDParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SHP.
type ShippingRateFixedAmountCurrencyOptionsShpParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SLL.
type ShippingRateFixedAmountCurrencyOptionsSllParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SOS.
type ShippingRateFixedAmountCurrencyOptionsSosParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SRD.
type ShippingRateFixedAmountCurrencyOptionsSrdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in STD.
type ShippingRateFixedAmountCurrencyOptionsStdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in SZL.
type ShippingRateFixedAmountCurrencyOptionsSzlParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in THB.
type ShippingRateFixedAmountCurrencyOptionsThbParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TJS.
type ShippingRateFixedAmountCurrencyOptionsTjsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TND.
type ShippingRateFixedAmountCurrencyOptionsTndParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TOP.
type ShippingRateFixedAmountCurrencyOptionsTopParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TRY.
type ShippingRateFixedAmountCurrencyOptionsTryParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TTD.
type ShippingRateFixedAmountCurrencyOptionsTtdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TWD.
type ShippingRateFixedAmountCurrencyOptionsTwdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in TZS.
type ShippingRateFixedAmountCurrencyOptionsTzsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in UAH.
type ShippingRateFixedAmountCurrencyOptionsUahParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in UGX.
type ShippingRateFixedAmountCurrencyOptionsUgxParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in USD.
type ShippingRateFixedAmountCurrencyOptionsUSDParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in USDC.
type ShippingRateFixedAmountCurrencyOptionsUsdcParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in UYU.
type ShippingRateFixedAmountCurrencyOptionsUyuParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in UZS.
type ShippingRateFixedAmountCurrencyOptionsUzsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in VND.
type ShippingRateFixedAmountCurrencyOptionsVndParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in VUV.
type ShippingRateFixedAmountCurrencyOptionsVuvParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in WST.
type ShippingRateFixedAmountCurrencyOptionsWstParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in XAF.
type ShippingRateFixedAmountCurrencyOptionsXafParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in XCD.
type ShippingRateFixedAmountCurrencyOptionsXcdParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in XOF.
type ShippingRateFixedAmountCurrencyOptionsXofParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in XPF.
type ShippingRateFixedAmountCurrencyOptionsXpfParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in YER.
type ShippingRateFixedAmountCurrencyOptionsYerParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ZAR.
type ShippingRateFixedAmountCurrencyOptionsZarParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rate defined in ZMW.
type ShippingRateFixedAmountCurrencyOptionsZmwParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type ShippingRateFixedAmountCurrencyOptionsParams struct {
	// Shipping rate defined in AED.
	Aed *ShippingRateFixedAmountCurrencyOptionsAedParams `form:"aed"`
	// Shipping rate defined in AFN.
	Afn *ShippingRateFixedAmountCurrencyOptionsAfnParams `form:"afn"`
	// Shipping rate defined in ALL.
	All *ShippingRateFixedAmountCurrencyOptionsAllParams `form:"all"`
	// Shipping rate defined in AMD.
	Amd *ShippingRateFixedAmountCurrencyOptionsAmdParams `form:"amd"`
	// Shipping rate defined in ANG.
	Ang *ShippingRateFixedAmountCurrencyOptionsAngParams `form:"ang"`
	// Shipping rate defined in AOA.
	Aoa *ShippingRateFixedAmountCurrencyOptionsAoaParams `form:"aoa"`
	// Shipping rate defined in ARS.
	Ars *ShippingRateFixedAmountCurrencyOptionsArsParams `form:"ars"`
	// Shipping rate defined in AUD.
	AUD *ShippingRateFixedAmountCurrencyOptionsAUDParams `form:"aud"`
	// Shipping rate defined in AWG.
	Awg *ShippingRateFixedAmountCurrencyOptionsAwgParams `form:"awg"`
	// Shipping rate defined in AZN.
	Azn *ShippingRateFixedAmountCurrencyOptionsAznParams `form:"azn"`
	// Shipping rate defined in BAM.
	Bam *ShippingRateFixedAmountCurrencyOptionsBamParams `form:"bam"`
	// Shipping rate defined in BBD.
	Bbd *ShippingRateFixedAmountCurrencyOptionsBbdParams `form:"bbd"`
	// Shipping rate defined in BDT.
	Bdt *ShippingRateFixedAmountCurrencyOptionsBdtParams `form:"bdt"`
	// Shipping rate defined in BGN.
	Bgn *ShippingRateFixedAmountCurrencyOptionsBgnParams `form:"bgn"`
	// Shipping rate defined in BHD.
	Bhd *ShippingRateFixedAmountCurrencyOptionsBhdParams `form:"bhd"`
	// Shipping rate defined in BIF.
	Bif *ShippingRateFixedAmountCurrencyOptionsBifParams `form:"bif"`
	// Shipping rate defined in BMD.
	Bmd *ShippingRateFixedAmountCurrencyOptionsBmdParams `form:"bmd"`
	// Shipping rate defined in BND.
	Bnd *ShippingRateFixedAmountCurrencyOptionsBndParams `form:"bnd"`
	// Shipping rate defined in BOB.
	Bob *ShippingRateFixedAmountCurrencyOptionsBobParams `form:"bob"`
	// Shipping rate defined in BRL.
	Brl *ShippingRateFixedAmountCurrencyOptionsBrlParams `form:"brl"`
	// Shipping rate defined in BSD.
	Bsd *ShippingRateFixedAmountCurrencyOptionsBsdParams `form:"bsd"`
	// Shipping rate defined in BTN.
	Btn *ShippingRateFixedAmountCurrencyOptionsBtnParams `form:"btn"`
	// Shipping rate defined in BWP.
	Bwp *ShippingRateFixedAmountCurrencyOptionsBwpParams `form:"bwp"`
	// Shipping rate defined in BYN.
	Byn *ShippingRateFixedAmountCurrencyOptionsBynParams `form:"byn"`
	// Shipping rate defined in BZD.
	Bzd *ShippingRateFixedAmountCurrencyOptionsBzdParams `form:"bzd"`
	// Shipping rate defined in CAD.
	CAD *ShippingRateFixedAmountCurrencyOptionsCADParams `form:"cad"`
	// Shipping rate defined in CDF.
	Cdf *ShippingRateFixedAmountCurrencyOptionsCdfParams `form:"cdf"`
	// Shipping rate defined in CHF.
	CHF *ShippingRateFixedAmountCurrencyOptionsCHFParams `form:"chf"`
	// Shipping rate defined in CLP.
	Clp *ShippingRateFixedAmountCurrencyOptionsClpParams `form:"clp"`
	// Shipping rate defined in CNY.
	Cny *ShippingRateFixedAmountCurrencyOptionsCnyParams `form:"cny"`
	// Shipping rate defined in COP.
	Cop *ShippingRateFixedAmountCurrencyOptionsCopParams `form:"cop"`
	// Shipping rate defined in CRC.
	Crc *ShippingRateFixedAmountCurrencyOptionsCrcParams `form:"crc"`
	// Shipping rate defined in CVE.
	Cve *ShippingRateFixedAmountCurrencyOptionsCveParams `form:"cve"`
	// Shipping rate defined in CZK.
	CZK *ShippingRateFixedAmountCurrencyOptionsCZKParams `form:"czk"`
	// Shipping rate defined in DJF.
	Djf *ShippingRateFixedAmountCurrencyOptionsDjfParams `form:"djf"`
	// Shipping rate defined in DKK.
	DKK *ShippingRateFixedAmountCurrencyOptionsDKKParams `form:"dkk"`
	// Shipping rate defined in DOP.
	Dop *ShippingRateFixedAmountCurrencyOptionsDopParams `form:"dop"`
	// Shipping rate defined in DZD.
	Dzd *ShippingRateFixedAmountCurrencyOptionsDzdParams `form:"dzd"`
	// Shipping rate defined in EGP.
	Egp *ShippingRateFixedAmountCurrencyOptionsEgpParams `form:"egp"`
	// Shipping rate defined in ETB.
	Etb *ShippingRateFixedAmountCurrencyOptionsEtbParams `form:"etb"`
	// Shipping rate defined in EUR.
	EUR *ShippingRateFixedAmountCurrencyOptionsEURParams `form:"eur"`
	// Shipping rate defined in FJD.
	Fjd *ShippingRateFixedAmountCurrencyOptionsFjdParams `form:"fjd"`
	// Shipping rate defined in FKP.
	Fkp *ShippingRateFixedAmountCurrencyOptionsFkpParams `form:"fkp"`
	// Shipping rate defined in GBP.
	GBP *ShippingRateFixedAmountCurrencyOptionsGBPParams `form:"gbp"`
	// Shipping rate defined in GEL.
	Gel *ShippingRateFixedAmountCurrencyOptionsGelParams `form:"gel"`
	// Shipping rate defined in GHS.
	Ghs *ShippingRateFixedAmountCurrencyOptionsGhsParams `form:"ghs"`
	// Shipping rate defined in GIP.
	Gip *ShippingRateFixedAmountCurrencyOptionsGipParams `form:"gip"`
	// Shipping rate defined in GMD.
	Gmd *ShippingRateFixedAmountCurrencyOptionsGmdParams `form:"gmd"`
	// Shipping rate defined in GNF.
	Gnf *ShippingRateFixedAmountCurrencyOptionsGnfParams `form:"gnf"`
	// Shipping rate defined in GTQ.
	Gtq *ShippingRateFixedAmountCurrencyOptionsGtqParams `form:"gtq"`
	// Shipping rate defined in GYD.
	Gyd *ShippingRateFixedAmountCurrencyOptionsGydParams `form:"gyd"`
	// Shipping rate defined in HKD.
	HKD *ShippingRateFixedAmountCurrencyOptionsHKDParams `form:"hkd"`
	// Shipping rate defined in HNL.
	Hnl *ShippingRateFixedAmountCurrencyOptionsHnlParams `form:"hnl"`
	// Shipping rate defined in HRK.
	Hrk *ShippingRateFixedAmountCurrencyOptionsHrkParams `form:"hrk"`
	// Shipping rate defined in HTG.
	Htg *ShippingRateFixedAmountCurrencyOptionsHtgParams `form:"htg"`
	// Shipping rate defined in HUF.
	Huf *ShippingRateFixedAmountCurrencyOptionsHufParams `form:"huf"`
	// Shipping rate defined in IDR.
	Idr *ShippingRateFixedAmountCurrencyOptionsIdrParams `form:"idr"`
	// Shipping rate defined in ILS.
	Ils *ShippingRateFixedAmountCurrencyOptionsIlsParams `form:"ils"`
	// Shipping rate defined in INR.
	Inr *ShippingRateFixedAmountCurrencyOptionsInrParams `form:"inr"`
	// Shipping rate defined in ISK.
	Isk *ShippingRateFixedAmountCurrencyOptionsIskParams `form:"isk"`
	// Shipping rate defined in JMD.
	Jmd *ShippingRateFixedAmountCurrencyOptionsJmdParams `form:"jmd"`
	// Shipping rate defined in JOD.
	Jod *ShippingRateFixedAmountCurrencyOptionsJodParams `form:"jod"`
	// Shipping rate defined in JPY.
	Jpy *ShippingRateFixedAmountCurrencyOptionsJpyParams `form:"jpy"`
	// Shipping rate defined in KES.
	Kes *ShippingRateFixedAmountCurrencyOptionsKesParams `form:"kes"`
	// Shipping rate defined in KGS.
	Kgs *ShippingRateFixedAmountCurrencyOptionsKgsParams `form:"kgs"`
	// Shipping rate defined in KHR.
	Khr *ShippingRateFixedAmountCurrencyOptionsKhrParams `form:"khr"`
	// Shipping rate defined in KMF.
	Kmf *ShippingRateFixedAmountCurrencyOptionsKmfParams `form:"kmf"`
	// Shipping rate defined in KRW.
	Krw *ShippingRateFixedAmountCurrencyOptionsKrwParams `form:"krw"`
	// Shipping rate defined in KWD.
	Kwd *ShippingRateFixedAmountCurrencyOptionsKwdParams `form:"kwd"`
	// Shipping rate defined in KYD.
	Kyd *ShippingRateFixedAmountCurrencyOptionsKydParams `form:"kyd"`
	// Shipping rate defined in KZT.
	Kzt *ShippingRateFixedAmountCurrencyOptionsKztParams `form:"kzt"`
	// Shipping rate defined in LAK.
	Lak *ShippingRateFixedAmountCurrencyOptionsLakParams `form:"lak"`
	// Shipping rate defined in LBP.
	Lbp *ShippingRateFixedAmountCurrencyOptionsLbpParams `form:"lbp"`
	// Shipping rate defined in LKR.
	Lkr *ShippingRateFixedAmountCurrencyOptionsLkrParams `form:"lkr"`
	// Shipping rate defined in LRD.
	Lrd *ShippingRateFixedAmountCurrencyOptionsLrdParams `form:"lrd"`
	// Shipping rate defined in LSL.
	Lsl *ShippingRateFixedAmountCurrencyOptionsLslParams `form:"lsl"`
	// Shipping rate defined in MAD.
	Mad *ShippingRateFixedAmountCurrencyOptionsMadParams `form:"mad"`
	// Shipping rate defined in MDL.
	Mdl *ShippingRateFixedAmountCurrencyOptionsMdlParams `form:"mdl"`
	// Shipping rate defined in MGA.
	Mga *ShippingRateFixedAmountCurrencyOptionsMgaParams `form:"mga"`
	// Shipping rate defined in MKD.
	Mkd *ShippingRateFixedAmountCurrencyOptionsMkdParams `form:"mkd"`
	// Shipping rate defined in MMK.
	Mmk *ShippingRateFixedAmountCurrencyOptionsMmkParams `form:"mmk"`
	// Shipping rate defined in MNT.
	Mnt *ShippingRateFixedAmountCurrencyOptionsMntParams `form:"mnt"`
	// Shipping rate defined in MOP.
	Mop *ShippingRateFixedAmountCurrencyOptionsMopParams `form:"mop"`
	// Shipping rate defined in MRO.
	Mro *ShippingRateFixedAmountCurrencyOptionsMroParams `form:"mro"`
	// Shipping rate defined in MUR.
	Mur *ShippingRateFixedAmountCurrencyOptionsMurParams `form:"mur"`
	// Shipping rate defined in MVR.
	Mvr *ShippingRateFixedAmountCurrencyOptionsMvrParams `form:"mvr"`
	// Shipping rate defined in MWK.
	Mwk *ShippingRateFixedAmountCurrencyOptionsMwkParams `form:"mwk"`
	// Shipping rate defined in MXN.
	Mxn *ShippingRateFixedAmountCurrencyOptionsMxnParams `form:"mxn"`
	// Shipping rate defined in MYR.
	MYR *ShippingRateFixedAmountCurrencyOptionsMYRParams `form:"myr"`
	// Shipping rate defined in MZN.
	Mzn *ShippingRateFixedAmountCurrencyOptionsMznParams `form:"mzn"`
	// Shipping rate defined in NAD.
	Nad *ShippingRateFixedAmountCurrencyOptionsNadParams `form:"nad"`
	// Shipping rate defined in NGN.
	Ngn *ShippingRateFixedAmountCurrencyOptionsNgnParams `form:"ngn"`
	// Shipping rate defined in NIO.
	Nio *ShippingRateFixedAmountCurrencyOptionsNioParams `form:"nio"`
	// Shipping rate defined in NOK.
	NOK *ShippingRateFixedAmountCurrencyOptionsNOKParams `form:"nok"`
	// Shipping rate defined in NPR.
	Npr *ShippingRateFixedAmountCurrencyOptionsNprParams `form:"npr"`
	// Shipping rate defined in NZD.
	NZD *ShippingRateFixedAmountCurrencyOptionsNZDParams `form:"nzd"`
	// Shipping rate defined in OMR.
	Omr *ShippingRateFixedAmountCurrencyOptionsOmrParams `form:"omr"`
	// Shipping rate defined in PAB.
	Pab *ShippingRateFixedAmountCurrencyOptionsPabParams `form:"pab"`
	// Shipping rate defined in PEN.
	Pen *ShippingRateFixedAmountCurrencyOptionsPenParams `form:"pen"`
	// Shipping rate defined in PGK.
	Pgk *ShippingRateFixedAmountCurrencyOptionsPgkParams `form:"pgk"`
	// Shipping rate defined in PHP.
	Php *ShippingRateFixedAmountCurrencyOptionsPhpParams `form:"php"`
	// Shipping rate defined in PKR.
	Pkr *ShippingRateFixedAmountCurrencyOptionsPkrParams `form:"pkr"`
	// Shipping rate defined in PLN.
	Pln *ShippingRateFixedAmountCurrencyOptionsPlnParams `form:"pln"`
	// Shipping rate defined in PYG.
	Pyg *ShippingRateFixedAmountCurrencyOptionsPygParams `form:"pyg"`
	// Shipping rate defined in QAR.
	Qar *ShippingRateFixedAmountCurrencyOptionsQarParams `form:"qar"`
	// Shipping rate defined in RON.
	Ron *ShippingRateFixedAmountCurrencyOptionsRonParams `form:"ron"`
	// Shipping rate defined in RSD.
	Rsd *ShippingRateFixedAmountCurrencyOptionsRsdParams `form:"rsd"`
	// Shipping rate defined in RUB.
	Rub *ShippingRateFixedAmountCurrencyOptionsRubParams `form:"rub"`
	// Shipping rate defined in RWF.
	Rwf *ShippingRateFixedAmountCurrencyOptionsRwfParams `form:"rwf"`
	// Shipping rate defined in SAR.
	Sar *ShippingRateFixedAmountCurrencyOptionsSarParams `form:"sar"`
	// Shipping rate defined in SBD.
	Sbd *ShippingRateFixedAmountCurrencyOptionsSbdParams `form:"sbd"`
	// Shipping rate defined in SCR.
	Scr *ShippingRateFixedAmountCurrencyOptionsScrParams `form:"scr"`
	// Shipping rate defined in SEK.
	SEK *ShippingRateFixedAmountCurrencyOptionsSEKParams `form:"sek"`
	// Shipping rate defined in SGD.
	SGD *ShippingRateFixedAmountCurrencyOptionsSGDParams `form:"sgd"`
	// Shipping rate defined in SHP.
	Shp *ShippingRateFixedAmountCurrencyOptionsShpParams `form:"shp"`
	// Shipping rate defined in SLL.
	Sll *ShippingRateFixedAmountCurrencyOptionsSllParams `form:"sll"`
	// Shipping rate defined in SOS.
	Sos *ShippingRateFixedAmountCurrencyOptionsSosParams `form:"sos"`
	// Shipping rate defined in SRD.
	Srd *ShippingRateFixedAmountCurrencyOptionsSrdParams `form:"srd"`
	// Shipping rate defined in STD.
	Std *ShippingRateFixedAmountCurrencyOptionsStdParams `form:"std"`
	// Shipping rate defined in SZL.
	Szl *ShippingRateFixedAmountCurrencyOptionsSzlParams `form:"szl"`
	// Shipping rate defined in THB.
	Thb *ShippingRateFixedAmountCurrencyOptionsThbParams `form:"thb"`
	// Shipping rate defined in TJS.
	Tjs *ShippingRateFixedAmountCurrencyOptionsTjsParams `form:"tjs"`
	// Shipping rate defined in TND.
	Tnd *ShippingRateFixedAmountCurrencyOptionsTndParams `form:"tnd"`
	// Shipping rate defined in TOP.
	Top *ShippingRateFixedAmountCurrencyOptionsTopParams `form:"top"`
	// Shipping rate defined in TRY.
	Try *ShippingRateFixedAmountCurrencyOptionsTryParams `form:"try"`
	// Shipping rate defined in TTD.
	Ttd *ShippingRateFixedAmountCurrencyOptionsTtdParams `form:"ttd"`
	// Shipping rate defined in TWD.
	Twd *ShippingRateFixedAmountCurrencyOptionsTwdParams `form:"twd"`
	// Shipping rate defined in TZS.
	Tzs *ShippingRateFixedAmountCurrencyOptionsTzsParams `form:"tzs"`
	// Shipping rate defined in UAH.
	Uah *ShippingRateFixedAmountCurrencyOptionsUahParams `form:"uah"`
	// Shipping rate defined in UGX.
	Ugx *ShippingRateFixedAmountCurrencyOptionsUgxParams `form:"ugx"`
	// Shipping rate defined in USD.
	USD *ShippingRateFixedAmountCurrencyOptionsUSDParams `form:"usd"`
	// Shipping rate defined in USDC.
	Usdc *ShippingRateFixedAmountCurrencyOptionsUsdcParams `form:"usdc"`
	// Shipping rate defined in UYU.
	Uyu *ShippingRateFixedAmountCurrencyOptionsUyuParams `form:"uyu"`
	// Shipping rate defined in UZS.
	Uzs *ShippingRateFixedAmountCurrencyOptionsUzsParams `form:"uzs"`
	// Shipping rate defined in VND.
	Vnd *ShippingRateFixedAmountCurrencyOptionsVndParams `form:"vnd"`
	// Shipping rate defined in VUV.
	Vuv *ShippingRateFixedAmountCurrencyOptionsVuvParams `form:"vuv"`
	// Shipping rate defined in WST.
	Wst *ShippingRateFixedAmountCurrencyOptionsWstParams `form:"wst"`
	// Shipping rate defined in XAF.
	Xaf *ShippingRateFixedAmountCurrencyOptionsXafParams `form:"xaf"`
	// Shipping rate defined in XCD.
	Xcd *ShippingRateFixedAmountCurrencyOptionsXcdParams `form:"xcd"`
	// Shipping rate defined in XOF.
	Xof *ShippingRateFixedAmountCurrencyOptionsXofParams `form:"xof"`
	// Shipping rate defined in XPF.
	Xpf *ShippingRateFixedAmountCurrencyOptionsXpfParams `form:"xpf"`
	// Shipping rate defined in YER.
	Yer *ShippingRateFixedAmountCurrencyOptionsYerParams `form:"yer"`
	// Shipping rate defined in ZAR.
	Zar *ShippingRateFixedAmountCurrencyOptionsZarParams `form:"zar"`
	// Shipping rate defined in ZMW.
	Zmw *ShippingRateFixedAmountCurrencyOptionsZmwParams `form:"zmw"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type ShippingRateFixedAmountParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *ShippingRateFixedAmountCurrencyOptionsParams `form:"currency_options"`
}

// Creates a new shipping rate object.
type ShippingRateParams struct {
	Params `form:"*"`
	// Whether the shipping rate can be used for new purchases. Defaults to `true`.
	Active *bool `form:"active"`
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *ShippingRateDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *ShippingRateFixedAmountParams `form:"fixed_amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type *string `form:"type"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type ShippingRateDeliveryEstimateMaximum struct {
	// A unit of time.
	Unit ShippingRateDeliveryEstimateMaximumUnit `json:"unit"`
	// Must be greater than 0.
	Value int64 `json:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type ShippingRateDeliveryEstimateMinimum struct {
	// A unit of time.
	Unit ShippingRateDeliveryEstimateMinimumUnit `json:"unit"`
	// Must be greater than 0.
	Value int64 `json:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type ShippingRateDeliveryEstimate struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *ShippingRateDeliveryEstimateMaximum `json:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *ShippingRateDeliveryEstimateMinimum `json:"minimum"`
}
type ShippingRateFixedAmountCurrencyOptionsAed struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAedTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAfn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAfnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAll struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAllTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAmd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAmdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAng struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAngTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAoa struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAoaTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsArs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsArsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAUD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAUDTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAwg struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAwgTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsAzn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsAznTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBam struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBamTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBbd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBbdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBdt struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBdtTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBgn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBgnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBhd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBhdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBif struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBifTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBmd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBmdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBnd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBndTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBob struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBobTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBrl struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBrlTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBsd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBsdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBtn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBtnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBwp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBwpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsByn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBynTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsBzd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsBzdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCAD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCADTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCdf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCdfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCHF struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCHFTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsClp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsClpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCny struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCnyTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCop struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCopTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCrc struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCrcTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCve struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCveTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsCZK struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsCZKTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsDjf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsDjfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsDKK struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsDKKTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsDop struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsDopTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsDzd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsDzdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsEgp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsEgpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsEtb struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsEtbTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsEUR struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsEURTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsFjd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsFjdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsFkp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsFkpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGBP struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGBPTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGel struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGelTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGhs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGhsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGip struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGipTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGmd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGmdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGnf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGnfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGtq struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGtqTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsGyd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsGydTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsHKD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsHKDTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsHnl struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsHnlTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsHrk struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsHrkTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsHtg struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsHtgTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsHuf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsHufTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsIdr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsIdrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsIls struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsIlsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsInr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsInrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsIsk struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsIskTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsJmd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsJmdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsJod struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsJodTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsJpy struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsJpyTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKes struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKesTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKgs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKgsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKhr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKhrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKmf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKmfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKrw struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKrwTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKwd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKwdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKyd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKydTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsKzt struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsKztTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsLak struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsLakTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsLbp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsLbpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsLkr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsLkrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsLrd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsLrdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsLsl struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsLslTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMad struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMadTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMdl struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMdlTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMga struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMgaTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMkd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMkdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMmk struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMmkTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMnt struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMntTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMop struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMopTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMro struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMroTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMur struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMurTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMvr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMvrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMwk struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMwkTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMxn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMxnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMYR struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMYRTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsMzn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsMznTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNad struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNadTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNgn struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNgnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNio struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNioTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNOK struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNOKTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNpr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNprTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsNZD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsNZDTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsOmr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsOmrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPab struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPabTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPen struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPenTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPgk struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPgkTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPhp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPhpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPkr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPkrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPln struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPlnTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsPyg struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsPygTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsQar struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsQarTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsRon struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsRonTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsRsd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsRsdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsRub struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsRubTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsRwf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsRwfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSar struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSarTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSbd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSbdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsScr struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsScrTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSEK struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSEKTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSGD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSGDTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsShp struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsShpTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSll struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSllTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSos struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSosTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSrd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSrdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsStd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsStdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsSzl struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsSzlTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsThb struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsThbTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTjs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTjsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTnd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTndTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTop struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTopTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTry struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTryTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTtd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTtdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTwd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTwdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsTzs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsTzsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUah struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUahTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUgx struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUgxTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUSD struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUSDTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUsdc struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUsdcTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUyu struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUyuTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsUzs struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsUzsTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsVnd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsVndTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsVuv struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsVuvTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsWst struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsWstTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsXaf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsXafTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsXcd struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsXcdTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsXof struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsXofTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsXpf struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsXpfTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsYer struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsYerTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsZar struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsZarTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptionsZmw struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateFixedAmountCurrencyOptionsZmwTaxBehavior `json:"tax_behavior"`
}
type ShippingRateFixedAmountCurrencyOptions struct {
	Aed  *ShippingRateFixedAmountCurrencyOptionsAed  `json:"aed"`
	Afn  *ShippingRateFixedAmountCurrencyOptionsAfn  `json:"afn"`
	All  *ShippingRateFixedAmountCurrencyOptionsAll  `json:"all"`
	Amd  *ShippingRateFixedAmountCurrencyOptionsAmd  `json:"amd"`
	Ang  *ShippingRateFixedAmountCurrencyOptionsAng  `json:"ang"`
	Aoa  *ShippingRateFixedAmountCurrencyOptionsAoa  `json:"aoa"`
	Ars  *ShippingRateFixedAmountCurrencyOptionsArs  `json:"ars"`
	AUD  *ShippingRateFixedAmountCurrencyOptionsAUD  `json:"aud"`
	Awg  *ShippingRateFixedAmountCurrencyOptionsAwg  `json:"awg"`
	Azn  *ShippingRateFixedAmountCurrencyOptionsAzn  `json:"azn"`
	Bam  *ShippingRateFixedAmountCurrencyOptionsBam  `json:"bam"`
	Bbd  *ShippingRateFixedAmountCurrencyOptionsBbd  `json:"bbd"`
	Bdt  *ShippingRateFixedAmountCurrencyOptionsBdt  `json:"bdt"`
	Bgn  *ShippingRateFixedAmountCurrencyOptionsBgn  `json:"bgn"`
	Bhd  *ShippingRateFixedAmountCurrencyOptionsBhd  `json:"bhd"`
	Bif  *ShippingRateFixedAmountCurrencyOptionsBif  `json:"bif"`
	Bmd  *ShippingRateFixedAmountCurrencyOptionsBmd  `json:"bmd"`
	Bnd  *ShippingRateFixedAmountCurrencyOptionsBnd  `json:"bnd"`
	Bob  *ShippingRateFixedAmountCurrencyOptionsBob  `json:"bob"`
	Brl  *ShippingRateFixedAmountCurrencyOptionsBrl  `json:"brl"`
	Bsd  *ShippingRateFixedAmountCurrencyOptionsBsd  `json:"bsd"`
	Btn  *ShippingRateFixedAmountCurrencyOptionsBtn  `json:"btn"`
	Bwp  *ShippingRateFixedAmountCurrencyOptionsBwp  `json:"bwp"`
	Byn  *ShippingRateFixedAmountCurrencyOptionsByn  `json:"byn"`
	Bzd  *ShippingRateFixedAmountCurrencyOptionsBzd  `json:"bzd"`
	CAD  *ShippingRateFixedAmountCurrencyOptionsCAD  `json:"cad"`
	Cdf  *ShippingRateFixedAmountCurrencyOptionsCdf  `json:"cdf"`
	CHF  *ShippingRateFixedAmountCurrencyOptionsCHF  `json:"chf"`
	Clp  *ShippingRateFixedAmountCurrencyOptionsClp  `json:"clp"`
	Cny  *ShippingRateFixedAmountCurrencyOptionsCny  `json:"cny"`
	Cop  *ShippingRateFixedAmountCurrencyOptionsCop  `json:"cop"`
	Crc  *ShippingRateFixedAmountCurrencyOptionsCrc  `json:"crc"`
	Cve  *ShippingRateFixedAmountCurrencyOptionsCve  `json:"cve"`
	CZK  *ShippingRateFixedAmountCurrencyOptionsCZK  `json:"czk"`
	Djf  *ShippingRateFixedAmountCurrencyOptionsDjf  `json:"djf"`
	DKK  *ShippingRateFixedAmountCurrencyOptionsDKK  `json:"dkk"`
	Dop  *ShippingRateFixedAmountCurrencyOptionsDop  `json:"dop"`
	Dzd  *ShippingRateFixedAmountCurrencyOptionsDzd  `json:"dzd"`
	Egp  *ShippingRateFixedAmountCurrencyOptionsEgp  `json:"egp"`
	Etb  *ShippingRateFixedAmountCurrencyOptionsEtb  `json:"etb"`
	EUR  *ShippingRateFixedAmountCurrencyOptionsEUR  `json:"eur"`
	Fjd  *ShippingRateFixedAmountCurrencyOptionsFjd  `json:"fjd"`
	Fkp  *ShippingRateFixedAmountCurrencyOptionsFkp  `json:"fkp"`
	GBP  *ShippingRateFixedAmountCurrencyOptionsGBP  `json:"gbp"`
	Gel  *ShippingRateFixedAmountCurrencyOptionsGel  `json:"gel"`
	Ghs  *ShippingRateFixedAmountCurrencyOptionsGhs  `json:"ghs"`
	Gip  *ShippingRateFixedAmountCurrencyOptionsGip  `json:"gip"`
	Gmd  *ShippingRateFixedAmountCurrencyOptionsGmd  `json:"gmd"`
	Gnf  *ShippingRateFixedAmountCurrencyOptionsGnf  `json:"gnf"`
	Gtq  *ShippingRateFixedAmountCurrencyOptionsGtq  `json:"gtq"`
	Gyd  *ShippingRateFixedAmountCurrencyOptionsGyd  `json:"gyd"`
	HKD  *ShippingRateFixedAmountCurrencyOptionsHKD  `json:"hkd"`
	Hnl  *ShippingRateFixedAmountCurrencyOptionsHnl  `json:"hnl"`
	Hrk  *ShippingRateFixedAmountCurrencyOptionsHrk  `json:"hrk"`
	Htg  *ShippingRateFixedAmountCurrencyOptionsHtg  `json:"htg"`
	Huf  *ShippingRateFixedAmountCurrencyOptionsHuf  `json:"huf"`
	Idr  *ShippingRateFixedAmountCurrencyOptionsIdr  `json:"idr"`
	Ils  *ShippingRateFixedAmountCurrencyOptionsIls  `json:"ils"`
	Inr  *ShippingRateFixedAmountCurrencyOptionsInr  `json:"inr"`
	Isk  *ShippingRateFixedAmountCurrencyOptionsIsk  `json:"isk"`
	Jmd  *ShippingRateFixedAmountCurrencyOptionsJmd  `json:"jmd"`
	Jod  *ShippingRateFixedAmountCurrencyOptionsJod  `json:"jod"`
	Jpy  *ShippingRateFixedAmountCurrencyOptionsJpy  `json:"jpy"`
	Kes  *ShippingRateFixedAmountCurrencyOptionsKes  `json:"kes"`
	Kgs  *ShippingRateFixedAmountCurrencyOptionsKgs  `json:"kgs"`
	Khr  *ShippingRateFixedAmountCurrencyOptionsKhr  `json:"khr"`
	Kmf  *ShippingRateFixedAmountCurrencyOptionsKmf  `json:"kmf"`
	Krw  *ShippingRateFixedAmountCurrencyOptionsKrw  `json:"krw"`
	Kwd  *ShippingRateFixedAmountCurrencyOptionsKwd  `json:"kwd"`
	Kyd  *ShippingRateFixedAmountCurrencyOptionsKyd  `json:"kyd"`
	Kzt  *ShippingRateFixedAmountCurrencyOptionsKzt  `json:"kzt"`
	Lak  *ShippingRateFixedAmountCurrencyOptionsLak  `json:"lak"`
	Lbp  *ShippingRateFixedAmountCurrencyOptionsLbp  `json:"lbp"`
	Lkr  *ShippingRateFixedAmountCurrencyOptionsLkr  `json:"lkr"`
	Lrd  *ShippingRateFixedAmountCurrencyOptionsLrd  `json:"lrd"`
	Lsl  *ShippingRateFixedAmountCurrencyOptionsLsl  `json:"lsl"`
	Mad  *ShippingRateFixedAmountCurrencyOptionsMad  `json:"mad"`
	Mdl  *ShippingRateFixedAmountCurrencyOptionsMdl  `json:"mdl"`
	Mga  *ShippingRateFixedAmountCurrencyOptionsMga  `json:"mga"`
	Mkd  *ShippingRateFixedAmountCurrencyOptionsMkd  `json:"mkd"`
	Mmk  *ShippingRateFixedAmountCurrencyOptionsMmk  `json:"mmk"`
	Mnt  *ShippingRateFixedAmountCurrencyOptionsMnt  `json:"mnt"`
	Mop  *ShippingRateFixedAmountCurrencyOptionsMop  `json:"mop"`
	Mro  *ShippingRateFixedAmountCurrencyOptionsMro  `json:"mro"`
	Mur  *ShippingRateFixedAmountCurrencyOptionsMur  `json:"mur"`
	Mvr  *ShippingRateFixedAmountCurrencyOptionsMvr  `json:"mvr"`
	Mwk  *ShippingRateFixedAmountCurrencyOptionsMwk  `json:"mwk"`
	Mxn  *ShippingRateFixedAmountCurrencyOptionsMxn  `json:"mxn"`
	MYR  *ShippingRateFixedAmountCurrencyOptionsMYR  `json:"myr"`
	Mzn  *ShippingRateFixedAmountCurrencyOptionsMzn  `json:"mzn"`
	Nad  *ShippingRateFixedAmountCurrencyOptionsNad  `json:"nad"`
	Ngn  *ShippingRateFixedAmountCurrencyOptionsNgn  `json:"ngn"`
	Nio  *ShippingRateFixedAmountCurrencyOptionsNio  `json:"nio"`
	NOK  *ShippingRateFixedAmountCurrencyOptionsNOK  `json:"nok"`
	Npr  *ShippingRateFixedAmountCurrencyOptionsNpr  `json:"npr"`
	NZD  *ShippingRateFixedAmountCurrencyOptionsNZD  `json:"nzd"`
	Omr  *ShippingRateFixedAmountCurrencyOptionsOmr  `json:"omr"`
	Pab  *ShippingRateFixedAmountCurrencyOptionsPab  `json:"pab"`
	Pen  *ShippingRateFixedAmountCurrencyOptionsPen  `json:"pen"`
	Pgk  *ShippingRateFixedAmountCurrencyOptionsPgk  `json:"pgk"`
	Php  *ShippingRateFixedAmountCurrencyOptionsPhp  `json:"php"`
	Pkr  *ShippingRateFixedAmountCurrencyOptionsPkr  `json:"pkr"`
	Pln  *ShippingRateFixedAmountCurrencyOptionsPln  `json:"pln"`
	Pyg  *ShippingRateFixedAmountCurrencyOptionsPyg  `json:"pyg"`
	Qar  *ShippingRateFixedAmountCurrencyOptionsQar  `json:"qar"`
	Ron  *ShippingRateFixedAmountCurrencyOptionsRon  `json:"ron"`
	Rsd  *ShippingRateFixedAmountCurrencyOptionsRsd  `json:"rsd"`
	Rub  *ShippingRateFixedAmountCurrencyOptionsRub  `json:"rub"`
	Rwf  *ShippingRateFixedAmountCurrencyOptionsRwf  `json:"rwf"`
	Sar  *ShippingRateFixedAmountCurrencyOptionsSar  `json:"sar"`
	Sbd  *ShippingRateFixedAmountCurrencyOptionsSbd  `json:"sbd"`
	Scr  *ShippingRateFixedAmountCurrencyOptionsScr  `json:"scr"`
	SEK  *ShippingRateFixedAmountCurrencyOptionsSEK  `json:"sek"`
	SGD  *ShippingRateFixedAmountCurrencyOptionsSGD  `json:"sgd"`
	Shp  *ShippingRateFixedAmountCurrencyOptionsShp  `json:"shp"`
	Sll  *ShippingRateFixedAmountCurrencyOptionsSll  `json:"sll"`
	Sos  *ShippingRateFixedAmountCurrencyOptionsSos  `json:"sos"`
	Srd  *ShippingRateFixedAmountCurrencyOptionsSrd  `json:"srd"`
	Std  *ShippingRateFixedAmountCurrencyOptionsStd  `json:"std"`
	Szl  *ShippingRateFixedAmountCurrencyOptionsSzl  `json:"szl"`
	Thb  *ShippingRateFixedAmountCurrencyOptionsThb  `json:"thb"`
	Tjs  *ShippingRateFixedAmountCurrencyOptionsTjs  `json:"tjs"`
	Tnd  *ShippingRateFixedAmountCurrencyOptionsTnd  `json:"tnd"`
	Top  *ShippingRateFixedAmountCurrencyOptionsTop  `json:"top"`
	Try  *ShippingRateFixedAmountCurrencyOptionsTry  `json:"try"`
	Ttd  *ShippingRateFixedAmountCurrencyOptionsTtd  `json:"ttd"`
	Twd  *ShippingRateFixedAmountCurrencyOptionsTwd  `json:"twd"`
	Tzs  *ShippingRateFixedAmountCurrencyOptionsTzs  `json:"tzs"`
	Uah  *ShippingRateFixedAmountCurrencyOptionsUah  `json:"uah"`
	Ugx  *ShippingRateFixedAmountCurrencyOptionsUgx  `json:"ugx"`
	USD  *ShippingRateFixedAmountCurrencyOptionsUSD  `json:"usd"`
	Usdc *ShippingRateFixedAmountCurrencyOptionsUsdc `json:"usdc"`
	Uyu  *ShippingRateFixedAmountCurrencyOptionsUyu  `json:"uyu"`
	Uzs  *ShippingRateFixedAmountCurrencyOptionsUzs  `json:"uzs"`
	Vnd  *ShippingRateFixedAmountCurrencyOptionsVnd  `json:"vnd"`
	Vuv  *ShippingRateFixedAmountCurrencyOptionsVuv  `json:"vuv"`
	Wst  *ShippingRateFixedAmountCurrencyOptionsWst  `json:"wst"`
	Xaf  *ShippingRateFixedAmountCurrencyOptionsXaf  `json:"xaf"`
	Xcd  *ShippingRateFixedAmountCurrencyOptionsXcd  `json:"xcd"`
	Xof  *ShippingRateFixedAmountCurrencyOptionsXof  `json:"xof"`
	Xpf  *ShippingRateFixedAmountCurrencyOptionsXpf  `json:"xpf"`
	Yer  *ShippingRateFixedAmountCurrencyOptionsYer  `json:"yer"`
	Zar  *ShippingRateFixedAmountCurrencyOptionsZar  `json:"zar"`
	Zmw  *ShippingRateFixedAmountCurrencyOptionsZmw  `json:"zmw"`
}
type ShippingRateFixedAmount struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency        Currency                                `json:"currency"`
	CurrencyOptions *ShippingRateFixedAmountCurrencyOptions `json:"currency_options"`
}

// Shipping rates describe the price of shipping presented to your customers and can be
// applied to [Checkout Sessions](https://stripe.com/docs/payments/checkout/shipping)
// and [Orders](https://stripe.com/docs/orders/shipping) to collect shipping costs.
type ShippingRate struct {
	APIResource
	// Whether the shipping rate can be used for new purchases. Defaults to `true`.
	Active bool `json:"active"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *ShippingRateDeliveryEstimate `json:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName string                   `json:"display_name"`
	FixedAmount *ShippingRateFixedAmount `json:"fixed_amount"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior ShippingRateTaxBehavior `json:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *TaxCode `json:"tax_code"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type ShippingRateType `json:"type"`
}

// ShippingRateList is a list of ShippingRates as retrieved from a list endpoint.
type ShippingRateList struct {
	APIResource
	ListMeta
	Data []*ShippingRate `json:"data"`
}

// UnmarshalJSON handles deserialization of a ShippingRate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *ShippingRate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type shippingRate ShippingRate
	var v shippingRate
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = ShippingRate(v)
	return nil
}
