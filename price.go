//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
type PriceBillingScheme string

// List of values that PriceBillingScheme can take
const (
	PriceBillingSchemePerUnit PriceBillingScheme = "per_unit"
	PriceBillingSchemeTiered  PriceBillingScheme = "tiered"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAedTaxBehavior string

// List of values that PriceCurrencyOptionsAedTaxBehavior can take
const (
	PriceCurrencyOptionsAedTaxBehaviorExclusive   PriceCurrencyOptionsAedTaxBehavior = "exclusive"
	PriceCurrencyOptionsAedTaxBehaviorInclusive   PriceCurrencyOptionsAedTaxBehavior = "inclusive"
	PriceCurrencyOptionsAedTaxBehaviorUnspecified PriceCurrencyOptionsAedTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAfnTaxBehavior string

// List of values that PriceCurrencyOptionsAfnTaxBehavior can take
const (
	PriceCurrencyOptionsAfnTaxBehaviorExclusive   PriceCurrencyOptionsAfnTaxBehavior = "exclusive"
	PriceCurrencyOptionsAfnTaxBehaviorInclusive   PriceCurrencyOptionsAfnTaxBehavior = "inclusive"
	PriceCurrencyOptionsAfnTaxBehaviorUnspecified PriceCurrencyOptionsAfnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAllTaxBehavior string

// List of values that PriceCurrencyOptionsAllTaxBehavior can take
const (
	PriceCurrencyOptionsAllTaxBehaviorExclusive   PriceCurrencyOptionsAllTaxBehavior = "exclusive"
	PriceCurrencyOptionsAllTaxBehaviorInclusive   PriceCurrencyOptionsAllTaxBehavior = "inclusive"
	PriceCurrencyOptionsAllTaxBehaviorUnspecified PriceCurrencyOptionsAllTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAmdTaxBehavior string

// List of values that PriceCurrencyOptionsAmdTaxBehavior can take
const (
	PriceCurrencyOptionsAmdTaxBehaviorExclusive   PriceCurrencyOptionsAmdTaxBehavior = "exclusive"
	PriceCurrencyOptionsAmdTaxBehaviorInclusive   PriceCurrencyOptionsAmdTaxBehavior = "inclusive"
	PriceCurrencyOptionsAmdTaxBehaviorUnspecified PriceCurrencyOptionsAmdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAngTaxBehavior string

// List of values that PriceCurrencyOptionsAngTaxBehavior can take
const (
	PriceCurrencyOptionsAngTaxBehaviorExclusive   PriceCurrencyOptionsAngTaxBehavior = "exclusive"
	PriceCurrencyOptionsAngTaxBehaviorInclusive   PriceCurrencyOptionsAngTaxBehavior = "inclusive"
	PriceCurrencyOptionsAngTaxBehaviorUnspecified PriceCurrencyOptionsAngTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAoaTaxBehavior string

// List of values that PriceCurrencyOptionsAoaTaxBehavior can take
const (
	PriceCurrencyOptionsAoaTaxBehaviorExclusive   PriceCurrencyOptionsAoaTaxBehavior = "exclusive"
	PriceCurrencyOptionsAoaTaxBehaviorInclusive   PriceCurrencyOptionsAoaTaxBehavior = "inclusive"
	PriceCurrencyOptionsAoaTaxBehaviorUnspecified PriceCurrencyOptionsAoaTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsArsTaxBehavior string

// List of values that PriceCurrencyOptionsArsTaxBehavior can take
const (
	PriceCurrencyOptionsArsTaxBehaviorExclusive   PriceCurrencyOptionsArsTaxBehavior = "exclusive"
	PriceCurrencyOptionsArsTaxBehaviorInclusive   PriceCurrencyOptionsArsTaxBehavior = "inclusive"
	PriceCurrencyOptionsArsTaxBehaviorUnspecified PriceCurrencyOptionsArsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAUDTaxBehavior string

// List of values that PriceCurrencyOptionsAUDTaxBehavior can take
const (
	PriceCurrencyOptionsAUDTaxBehaviorExclusive   PriceCurrencyOptionsAUDTaxBehavior = "exclusive"
	PriceCurrencyOptionsAUDTaxBehaviorInclusive   PriceCurrencyOptionsAUDTaxBehavior = "inclusive"
	PriceCurrencyOptionsAUDTaxBehaviorUnspecified PriceCurrencyOptionsAUDTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAwgTaxBehavior string

// List of values that PriceCurrencyOptionsAwgTaxBehavior can take
const (
	PriceCurrencyOptionsAwgTaxBehaviorExclusive   PriceCurrencyOptionsAwgTaxBehavior = "exclusive"
	PriceCurrencyOptionsAwgTaxBehaviorInclusive   PriceCurrencyOptionsAwgTaxBehavior = "inclusive"
	PriceCurrencyOptionsAwgTaxBehaviorUnspecified PriceCurrencyOptionsAwgTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsAznTaxBehavior string

// List of values that PriceCurrencyOptionsAznTaxBehavior can take
const (
	PriceCurrencyOptionsAznTaxBehaviorExclusive   PriceCurrencyOptionsAznTaxBehavior = "exclusive"
	PriceCurrencyOptionsAznTaxBehaviorInclusive   PriceCurrencyOptionsAznTaxBehavior = "inclusive"
	PriceCurrencyOptionsAznTaxBehaviorUnspecified PriceCurrencyOptionsAznTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBamTaxBehavior string

// List of values that PriceCurrencyOptionsBamTaxBehavior can take
const (
	PriceCurrencyOptionsBamTaxBehaviorExclusive   PriceCurrencyOptionsBamTaxBehavior = "exclusive"
	PriceCurrencyOptionsBamTaxBehaviorInclusive   PriceCurrencyOptionsBamTaxBehavior = "inclusive"
	PriceCurrencyOptionsBamTaxBehaviorUnspecified PriceCurrencyOptionsBamTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBbdTaxBehavior string

// List of values that PriceCurrencyOptionsBbdTaxBehavior can take
const (
	PriceCurrencyOptionsBbdTaxBehaviorExclusive   PriceCurrencyOptionsBbdTaxBehavior = "exclusive"
	PriceCurrencyOptionsBbdTaxBehaviorInclusive   PriceCurrencyOptionsBbdTaxBehavior = "inclusive"
	PriceCurrencyOptionsBbdTaxBehaviorUnspecified PriceCurrencyOptionsBbdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBdtTaxBehavior string

// List of values that PriceCurrencyOptionsBdtTaxBehavior can take
const (
	PriceCurrencyOptionsBdtTaxBehaviorExclusive   PriceCurrencyOptionsBdtTaxBehavior = "exclusive"
	PriceCurrencyOptionsBdtTaxBehaviorInclusive   PriceCurrencyOptionsBdtTaxBehavior = "inclusive"
	PriceCurrencyOptionsBdtTaxBehaviorUnspecified PriceCurrencyOptionsBdtTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBgnTaxBehavior string

// List of values that PriceCurrencyOptionsBgnTaxBehavior can take
const (
	PriceCurrencyOptionsBgnTaxBehaviorExclusive   PriceCurrencyOptionsBgnTaxBehavior = "exclusive"
	PriceCurrencyOptionsBgnTaxBehaviorInclusive   PriceCurrencyOptionsBgnTaxBehavior = "inclusive"
	PriceCurrencyOptionsBgnTaxBehaviorUnspecified PriceCurrencyOptionsBgnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBhdTaxBehavior string

// List of values that PriceCurrencyOptionsBhdTaxBehavior can take
const (
	PriceCurrencyOptionsBhdTaxBehaviorExclusive   PriceCurrencyOptionsBhdTaxBehavior = "exclusive"
	PriceCurrencyOptionsBhdTaxBehaviorInclusive   PriceCurrencyOptionsBhdTaxBehavior = "inclusive"
	PriceCurrencyOptionsBhdTaxBehaviorUnspecified PriceCurrencyOptionsBhdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBifTaxBehavior string

// List of values that PriceCurrencyOptionsBifTaxBehavior can take
const (
	PriceCurrencyOptionsBifTaxBehaviorExclusive   PriceCurrencyOptionsBifTaxBehavior = "exclusive"
	PriceCurrencyOptionsBifTaxBehaviorInclusive   PriceCurrencyOptionsBifTaxBehavior = "inclusive"
	PriceCurrencyOptionsBifTaxBehaviorUnspecified PriceCurrencyOptionsBifTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBmdTaxBehavior string

// List of values that PriceCurrencyOptionsBmdTaxBehavior can take
const (
	PriceCurrencyOptionsBmdTaxBehaviorExclusive   PriceCurrencyOptionsBmdTaxBehavior = "exclusive"
	PriceCurrencyOptionsBmdTaxBehaviorInclusive   PriceCurrencyOptionsBmdTaxBehavior = "inclusive"
	PriceCurrencyOptionsBmdTaxBehaviorUnspecified PriceCurrencyOptionsBmdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBndTaxBehavior string

// List of values that PriceCurrencyOptionsBndTaxBehavior can take
const (
	PriceCurrencyOptionsBndTaxBehaviorExclusive   PriceCurrencyOptionsBndTaxBehavior = "exclusive"
	PriceCurrencyOptionsBndTaxBehaviorInclusive   PriceCurrencyOptionsBndTaxBehavior = "inclusive"
	PriceCurrencyOptionsBndTaxBehaviorUnspecified PriceCurrencyOptionsBndTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBobTaxBehavior string

// List of values that PriceCurrencyOptionsBobTaxBehavior can take
const (
	PriceCurrencyOptionsBobTaxBehaviorExclusive   PriceCurrencyOptionsBobTaxBehavior = "exclusive"
	PriceCurrencyOptionsBobTaxBehaviorInclusive   PriceCurrencyOptionsBobTaxBehavior = "inclusive"
	PriceCurrencyOptionsBobTaxBehaviorUnspecified PriceCurrencyOptionsBobTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBrlTaxBehavior string

// List of values that PriceCurrencyOptionsBrlTaxBehavior can take
const (
	PriceCurrencyOptionsBrlTaxBehaviorExclusive   PriceCurrencyOptionsBrlTaxBehavior = "exclusive"
	PriceCurrencyOptionsBrlTaxBehaviorInclusive   PriceCurrencyOptionsBrlTaxBehavior = "inclusive"
	PriceCurrencyOptionsBrlTaxBehaviorUnspecified PriceCurrencyOptionsBrlTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBsdTaxBehavior string

// List of values that PriceCurrencyOptionsBsdTaxBehavior can take
const (
	PriceCurrencyOptionsBsdTaxBehaviorExclusive   PriceCurrencyOptionsBsdTaxBehavior = "exclusive"
	PriceCurrencyOptionsBsdTaxBehaviorInclusive   PriceCurrencyOptionsBsdTaxBehavior = "inclusive"
	PriceCurrencyOptionsBsdTaxBehaviorUnspecified PriceCurrencyOptionsBsdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBtnTaxBehavior string

// List of values that PriceCurrencyOptionsBtnTaxBehavior can take
const (
	PriceCurrencyOptionsBtnTaxBehaviorExclusive   PriceCurrencyOptionsBtnTaxBehavior = "exclusive"
	PriceCurrencyOptionsBtnTaxBehaviorInclusive   PriceCurrencyOptionsBtnTaxBehavior = "inclusive"
	PriceCurrencyOptionsBtnTaxBehaviorUnspecified PriceCurrencyOptionsBtnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBwpTaxBehavior string

// List of values that PriceCurrencyOptionsBwpTaxBehavior can take
const (
	PriceCurrencyOptionsBwpTaxBehaviorExclusive   PriceCurrencyOptionsBwpTaxBehavior = "exclusive"
	PriceCurrencyOptionsBwpTaxBehaviorInclusive   PriceCurrencyOptionsBwpTaxBehavior = "inclusive"
	PriceCurrencyOptionsBwpTaxBehaviorUnspecified PriceCurrencyOptionsBwpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBynTaxBehavior string

// List of values that PriceCurrencyOptionsBynTaxBehavior can take
const (
	PriceCurrencyOptionsBynTaxBehaviorExclusive   PriceCurrencyOptionsBynTaxBehavior = "exclusive"
	PriceCurrencyOptionsBynTaxBehaviorInclusive   PriceCurrencyOptionsBynTaxBehavior = "inclusive"
	PriceCurrencyOptionsBynTaxBehaviorUnspecified PriceCurrencyOptionsBynTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsBzdTaxBehavior string

// List of values that PriceCurrencyOptionsBzdTaxBehavior can take
const (
	PriceCurrencyOptionsBzdTaxBehaviorExclusive   PriceCurrencyOptionsBzdTaxBehavior = "exclusive"
	PriceCurrencyOptionsBzdTaxBehaviorInclusive   PriceCurrencyOptionsBzdTaxBehavior = "inclusive"
	PriceCurrencyOptionsBzdTaxBehaviorUnspecified PriceCurrencyOptionsBzdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCADTaxBehavior string

// List of values that PriceCurrencyOptionsCADTaxBehavior can take
const (
	PriceCurrencyOptionsCADTaxBehaviorExclusive   PriceCurrencyOptionsCADTaxBehavior = "exclusive"
	PriceCurrencyOptionsCADTaxBehaviorInclusive   PriceCurrencyOptionsCADTaxBehavior = "inclusive"
	PriceCurrencyOptionsCADTaxBehaviorUnspecified PriceCurrencyOptionsCADTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCdfTaxBehavior string

// List of values that PriceCurrencyOptionsCdfTaxBehavior can take
const (
	PriceCurrencyOptionsCdfTaxBehaviorExclusive   PriceCurrencyOptionsCdfTaxBehavior = "exclusive"
	PriceCurrencyOptionsCdfTaxBehaviorInclusive   PriceCurrencyOptionsCdfTaxBehavior = "inclusive"
	PriceCurrencyOptionsCdfTaxBehaviorUnspecified PriceCurrencyOptionsCdfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCHFTaxBehavior string

// List of values that PriceCurrencyOptionsCHFTaxBehavior can take
const (
	PriceCurrencyOptionsCHFTaxBehaviorExclusive   PriceCurrencyOptionsCHFTaxBehavior = "exclusive"
	PriceCurrencyOptionsCHFTaxBehaviorInclusive   PriceCurrencyOptionsCHFTaxBehavior = "inclusive"
	PriceCurrencyOptionsCHFTaxBehaviorUnspecified PriceCurrencyOptionsCHFTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsClpTaxBehavior string

// List of values that PriceCurrencyOptionsClpTaxBehavior can take
const (
	PriceCurrencyOptionsClpTaxBehaviorExclusive   PriceCurrencyOptionsClpTaxBehavior = "exclusive"
	PriceCurrencyOptionsClpTaxBehaviorInclusive   PriceCurrencyOptionsClpTaxBehavior = "inclusive"
	PriceCurrencyOptionsClpTaxBehaviorUnspecified PriceCurrencyOptionsClpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCnyTaxBehavior string

// List of values that PriceCurrencyOptionsCnyTaxBehavior can take
const (
	PriceCurrencyOptionsCnyTaxBehaviorExclusive   PriceCurrencyOptionsCnyTaxBehavior = "exclusive"
	PriceCurrencyOptionsCnyTaxBehaviorInclusive   PriceCurrencyOptionsCnyTaxBehavior = "inclusive"
	PriceCurrencyOptionsCnyTaxBehaviorUnspecified PriceCurrencyOptionsCnyTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCopTaxBehavior string

// List of values that PriceCurrencyOptionsCopTaxBehavior can take
const (
	PriceCurrencyOptionsCopTaxBehaviorExclusive   PriceCurrencyOptionsCopTaxBehavior = "exclusive"
	PriceCurrencyOptionsCopTaxBehaviorInclusive   PriceCurrencyOptionsCopTaxBehavior = "inclusive"
	PriceCurrencyOptionsCopTaxBehaviorUnspecified PriceCurrencyOptionsCopTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCrcTaxBehavior string

// List of values that PriceCurrencyOptionsCrcTaxBehavior can take
const (
	PriceCurrencyOptionsCrcTaxBehaviorExclusive   PriceCurrencyOptionsCrcTaxBehavior = "exclusive"
	PriceCurrencyOptionsCrcTaxBehaviorInclusive   PriceCurrencyOptionsCrcTaxBehavior = "inclusive"
	PriceCurrencyOptionsCrcTaxBehaviorUnspecified PriceCurrencyOptionsCrcTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCveTaxBehavior string

// List of values that PriceCurrencyOptionsCveTaxBehavior can take
const (
	PriceCurrencyOptionsCveTaxBehaviorExclusive   PriceCurrencyOptionsCveTaxBehavior = "exclusive"
	PriceCurrencyOptionsCveTaxBehaviorInclusive   PriceCurrencyOptionsCveTaxBehavior = "inclusive"
	PriceCurrencyOptionsCveTaxBehaviorUnspecified PriceCurrencyOptionsCveTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsCZKTaxBehavior string

// List of values that PriceCurrencyOptionsCZKTaxBehavior can take
const (
	PriceCurrencyOptionsCZKTaxBehaviorExclusive   PriceCurrencyOptionsCZKTaxBehavior = "exclusive"
	PriceCurrencyOptionsCZKTaxBehaviorInclusive   PriceCurrencyOptionsCZKTaxBehavior = "inclusive"
	PriceCurrencyOptionsCZKTaxBehaviorUnspecified PriceCurrencyOptionsCZKTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsDjfTaxBehavior string

// List of values that PriceCurrencyOptionsDjfTaxBehavior can take
const (
	PriceCurrencyOptionsDjfTaxBehaviorExclusive   PriceCurrencyOptionsDjfTaxBehavior = "exclusive"
	PriceCurrencyOptionsDjfTaxBehaviorInclusive   PriceCurrencyOptionsDjfTaxBehavior = "inclusive"
	PriceCurrencyOptionsDjfTaxBehaviorUnspecified PriceCurrencyOptionsDjfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsDKKTaxBehavior string

// List of values that PriceCurrencyOptionsDKKTaxBehavior can take
const (
	PriceCurrencyOptionsDKKTaxBehaviorExclusive   PriceCurrencyOptionsDKKTaxBehavior = "exclusive"
	PriceCurrencyOptionsDKKTaxBehaviorInclusive   PriceCurrencyOptionsDKKTaxBehavior = "inclusive"
	PriceCurrencyOptionsDKKTaxBehaviorUnspecified PriceCurrencyOptionsDKKTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsDopTaxBehavior string

// List of values that PriceCurrencyOptionsDopTaxBehavior can take
const (
	PriceCurrencyOptionsDopTaxBehaviorExclusive   PriceCurrencyOptionsDopTaxBehavior = "exclusive"
	PriceCurrencyOptionsDopTaxBehaviorInclusive   PriceCurrencyOptionsDopTaxBehavior = "inclusive"
	PriceCurrencyOptionsDopTaxBehaviorUnspecified PriceCurrencyOptionsDopTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsDzdTaxBehavior string

// List of values that PriceCurrencyOptionsDzdTaxBehavior can take
const (
	PriceCurrencyOptionsDzdTaxBehaviorExclusive   PriceCurrencyOptionsDzdTaxBehavior = "exclusive"
	PriceCurrencyOptionsDzdTaxBehaviorInclusive   PriceCurrencyOptionsDzdTaxBehavior = "inclusive"
	PriceCurrencyOptionsDzdTaxBehaviorUnspecified PriceCurrencyOptionsDzdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsEgpTaxBehavior string

// List of values that PriceCurrencyOptionsEgpTaxBehavior can take
const (
	PriceCurrencyOptionsEgpTaxBehaviorExclusive   PriceCurrencyOptionsEgpTaxBehavior = "exclusive"
	PriceCurrencyOptionsEgpTaxBehaviorInclusive   PriceCurrencyOptionsEgpTaxBehavior = "inclusive"
	PriceCurrencyOptionsEgpTaxBehaviorUnspecified PriceCurrencyOptionsEgpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsEtbTaxBehavior string

// List of values that PriceCurrencyOptionsEtbTaxBehavior can take
const (
	PriceCurrencyOptionsEtbTaxBehaviorExclusive   PriceCurrencyOptionsEtbTaxBehavior = "exclusive"
	PriceCurrencyOptionsEtbTaxBehaviorInclusive   PriceCurrencyOptionsEtbTaxBehavior = "inclusive"
	PriceCurrencyOptionsEtbTaxBehaviorUnspecified PriceCurrencyOptionsEtbTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsEURTaxBehavior string

// List of values that PriceCurrencyOptionsEURTaxBehavior can take
const (
	PriceCurrencyOptionsEURTaxBehaviorExclusive   PriceCurrencyOptionsEURTaxBehavior = "exclusive"
	PriceCurrencyOptionsEURTaxBehaviorInclusive   PriceCurrencyOptionsEURTaxBehavior = "inclusive"
	PriceCurrencyOptionsEURTaxBehaviorUnspecified PriceCurrencyOptionsEURTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsFjdTaxBehavior string

// List of values that PriceCurrencyOptionsFjdTaxBehavior can take
const (
	PriceCurrencyOptionsFjdTaxBehaviorExclusive   PriceCurrencyOptionsFjdTaxBehavior = "exclusive"
	PriceCurrencyOptionsFjdTaxBehaviorInclusive   PriceCurrencyOptionsFjdTaxBehavior = "inclusive"
	PriceCurrencyOptionsFjdTaxBehaviorUnspecified PriceCurrencyOptionsFjdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsFkpTaxBehavior string

// List of values that PriceCurrencyOptionsFkpTaxBehavior can take
const (
	PriceCurrencyOptionsFkpTaxBehaviorExclusive   PriceCurrencyOptionsFkpTaxBehavior = "exclusive"
	PriceCurrencyOptionsFkpTaxBehaviorInclusive   PriceCurrencyOptionsFkpTaxBehavior = "inclusive"
	PriceCurrencyOptionsFkpTaxBehaviorUnspecified PriceCurrencyOptionsFkpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGBPTaxBehavior string

// List of values that PriceCurrencyOptionsGBPTaxBehavior can take
const (
	PriceCurrencyOptionsGBPTaxBehaviorExclusive   PriceCurrencyOptionsGBPTaxBehavior = "exclusive"
	PriceCurrencyOptionsGBPTaxBehaviorInclusive   PriceCurrencyOptionsGBPTaxBehavior = "inclusive"
	PriceCurrencyOptionsGBPTaxBehaviorUnspecified PriceCurrencyOptionsGBPTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGelTaxBehavior string

// List of values that PriceCurrencyOptionsGelTaxBehavior can take
const (
	PriceCurrencyOptionsGelTaxBehaviorExclusive   PriceCurrencyOptionsGelTaxBehavior = "exclusive"
	PriceCurrencyOptionsGelTaxBehaviorInclusive   PriceCurrencyOptionsGelTaxBehavior = "inclusive"
	PriceCurrencyOptionsGelTaxBehaviorUnspecified PriceCurrencyOptionsGelTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGhsTaxBehavior string

// List of values that PriceCurrencyOptionsGhsTaxBehavior can take
const (
	PriceCurrencyOptionsGhsTaxBehaviorExclusive   PriceCurrencyOptionsGhsTaxBehavior = "exclusive"
	PriceCurrencyOptionsGhsTaxBehaviorInclusive   PriceCurrencyOptionsGhsTaxBehavior = "inclusive"
	PriceCurrencyOptionsGhsTaxBehaviorUnspecified PriceCurrencyOptionsGhsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGipTaxBehavior string

// List of values that PriceCurrencyOptionsGipTaxBehavior can take
const (
	PriceCurrencyOptionsGipTaxBehaviorExclusive   PriceCurrencyOptionsGipTaxBehavior = "exclusive"
	PriceCurrencyOptionsGipTaxBehaviorInclusive   PriceCurrencyOptionsGipTaxBehavior = "inclusive"
	PriceCurrencyOptionsGipTaxBehaviorUnspecified PriceCurrencyOptionsGipTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGmdTaxBehavior string

// List of values that PriceCurrencyOptionsGmdTaxBehavior can take
const (
	PriceCurrencyOptionsGmdTaxBehaviorExclusive   PriceCurrencyOptionsGmdTaxBehavior = "exclusive"
	PriceCurrencyOptionsGmdTaxBehaviorInclusive   PriceCurrencyOptionsGmdTaxBehavior = "inclusive"
	PriceCurrencyOptionsGmdTaxBehaviorUnspecified PriceCurrencyOptionsGmdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGnfTaxBehavior string

// List of values that PriceCurrencyOptionsGnfTaxBehavior can take
const (
	PriceCurrencyOptionsGnfTaxBehaviorExclusive   PriceCurrencyOptionsGnfTaxBehavior = "exclusive"
	PriceCurrencyOptionsGnfTaxBehaviorInclusive   PriceCurrencyOptionsGnfTaxBehavior = "inclusive"
	PriceCurrencyOptionsGnfTaxBehaviorUnspecified PriceCurrencyOptionsGnfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGtqTaxBehavior string

// List of values that PriceCurrencyOptionsGtqTaxBehavior can take
const (
	PriceCurrencyOptionsGtqTaxBehaviorExclusive   PriceCurrencyOptionsGtqTaxBehavior = "exclusive"
	PriceCurrencyOptionsGtqTaxBehaviorInclusive   PriceCurrencyOptionsGtqTaxBehavior = "inclusive"
	PriceCurrencyOptionsGtqTaxBehaviorUnspecified PriceCurrencyOptionsGtqTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsGydTaxBehavior string

// List of values that PriceCurrencyOptionsGydTaxBehavior can take
const (
	PriceCurrencyOptionsGydTaxBehaviorExclusive   PriceCurrencyOptionsGydTaxBehavior = "exclusive"
	PriceCurrencyOptionsGydTaxBehaviorInclusive   PriceCurrencyOptionsGydTaxBehavior = "inclusive"
	PriceCurrencyOptionsGydTaxBehaviorUnspecified PriceCurrencyOptionsGydTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsHKDTaxBehavior string

// List of values that PriceCurrencyOptionsHKDTaxBehavior can take
const (
	PriceCurrencyOptionsHKDTaxBehaviorExclusive   PriceCurrencyOptionsHKDTaxBehavior = "exclusive"
	PriceCurrencyOptionsHKDTaxBehaviorInclusive   PriceCurrencyOptionsHKDTaxBehavior = "inclusive"
	PriceCurrencyOptionsHKDTaxBehaviorUnspecified PriceCurrencyOptionsHKDTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsHnlTaxBehavior string

// List of values that PriceCurrencyOptionsHnlTaxBehavior can take
const (
	PriceCurrencyOptionsHnlTaxBehaviorExclusive   PriceCurrencyOptionsHnlTaxBehavior = "exclusive"
	PriceCurrencyOptionsHnlTaxBehaviorInclusive   PriceCurrencyOptionsHnlTaxBehavior = "inclusive"
	PriceCurrencyOptionsHnlTaxBehaviorUnspecified PriceCurrencyOptionsHnlTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsHrkTaxBehavior string

// List of values that PriceCurrencyOptionsHrkTaxBehavior can take
const (
	PriceCurrencyOptionsHrkTaxBehaviorExclusive   PriceCurrencyOptionsHrkTaxBehavior = "exclusive"
	PriceCurrencyOptionsHrkTaxBehaviorInclusive   PriceCurrencyOptionsHrkTaxBehavior = "inclusive"
	PriceCurrencyOptionsHrkTaxBehaviorUnspecified PriceCurrencyOptionsHrkTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsHtgTaxBehavior string

// List of values that PriceCurrencyOptionsHtgTaxBehavior can take
const (
	PriceCurrencyOptionsHtgTaxBehaviorExclusive   PriceCurrencyOptionsHtgTaxBehavior = "exclusive"
	PriceCurrencyOptionsHtgTaxBehaviorInclusive   PriceCurrencyOptionsHtgTaxBehavior = "inclusive"
	PriceCurrencyOptionsHtgTaxBehaviorUnspecified PriceCurrencyOptionsHtgTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsHufTaxBehavior string

// List of values that PriceCurrencyOptionsHufTaxBehavior can take
const (
	PriceCurrencyOptionsHufTaxBehaviorExclusive   PriceCurrencyOptionsHufTaxBehavior = "exclusive"
	PriceCurrencyOptionsHufTaxBehaviorInclusive   PriceCurrencyOptionsHufTaxBehavior = "inclusive"
	PriceCurrencyOptionsHufTaxBehaviorUnspecified PriceCurrencyOptionsHufTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsIdrTaxBehavior string

// List of values that PriceCurrencyOptionsIdrTaxBehavior can take
const (
	PriceCurrencyOptionsIdrTaxBehaviorExclusive   PriceCurrencyOptionsIdrTaxBehavior = "exclusive"
	PriceCurrencyOptionsIdrTaxBehaviorInclusive   PriceCurrencyOptionsIdrTaxBehavior = "inclusive"
	PriceCurrencyOptionsIdrTaxBehaviorUnspecified PriceCurrencyOptionsIdrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsIlsTaxBehavior string

// List of values that PriceCurrencyOptionsIlsTaxBehavior can take
const (
	PriceCurrencyOptionsIlsTaxBehaviorExclusive   PriceCurrencyOptionsIlsTaxBehavior = "exclusive"
	PriceCurrencyOptionsIlsTaxBehaviorInclusive   PriceCurrencyOptionsIlsTaxBehavior = "inclusive"
	PriceCurrencyOptionsIlsTaxBehaviorUnspecified PriceCurrencyOptionsIlsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsInrTaxBehavior string

// List of values that PriceCurrencyOptionsInrTaxBehavior can take
const (
	PriceCurrencyOptionsInrTaxBehaviorExclusive   PriceCurrencyOptionsInrTaxBehavior = "exclusive"
	PriceCurrencyOptionsInrTaxBehaviorInclusive   PriceCurrencyOptionsInrTaxBehavior = "inclusive"
	PriceCurrencyOptionsInrTaxBehaviorUnspecified PriceCurrencyOptionsInrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsIskTaxBehavior string

// List of values that PriceCurrencyOptionsIskTaxBehavior can take
const (
	PriceCurrencyOptionsIskTaxBehaviorExclusive   PriceCurrencyOptionsIskTaxBehavior = "exclusive"
	PriceCurrencyOptionsIskTaxBehaviorInclusive   PriceCurrencyOptionsIskTaxBehavior = "inclusive"
	PriceCurrencyOptionsIskTaxBehaviorUnspecified PriceCurrencyOptionsIskTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsJmdTaxBehavior string

// List of values that PriceCurrencyOptionsJmdTaxBehavior can take
const (
	PriceCurrencyOptionsJmdTaxBehaviorExclusive   PriceCurrencyOptionsJmdTaxBehavior = "exclusive"
	PriceCurrencyOptionsJmdTaxBehaviorInclusive   PriceCurrencyOptionsJmdTaxBehavior = "inclusive"
	PriceCurrencyOptionsJmdTaxBehaviorUnspecified PriceCurrencyOptionsJmdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsJodTaxBehavior string

// List of values that PriceCurrencyOptionsJodTaxBehavior can take
const (
	PriceCurrencyOptionsJodTaxBehaviorExclusive   PriceCurrencyOptionsJodTaxBehavior = "exclusive"
	PriceCurrencyOptionsJodTaxBehaviorInclusive   PriceCurrencyOptionsJodTaxBehavior = "inclusive"
	PriceCurrencyOptionsJodTaxBehaviorUnspecified PriceCurrencyOptionsJodTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsJpyTaxBehavior string

// List of values that PriceCurrencyOptionsJpyTaxBehavior can take
const (
	PriceCurrencyOptionsJpyTaxBehaviorExclusive   PriceCurrencyOptionsJpyTaxBehavior = "exclusive"
	PriceCurrencyOptionsJpyTaxBehaviorInclusive   PriceCurrencyOptionsJpyTaxBehavior = "inclusive"
	PriceCurrencyOptionsJpyTaxBehaviorUnspecified PriceCurrencyOptionsJpyTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKesTaxBehavior string

// List of values that PriceCurrencyOptionsKesTaxBehavior can take
const (
	PriceCurrencyOptionsKesTaxBehaviorExclusive   PriceCurrencyOptionsKesTaxBehavior = "exclusive"
	PriceCurrencyOptionsKesTaxBehaviorInclusive   PriceCurrencyOptionsKesTaxBehavior = "inclusive"
	PriceCurrencyOptionsKesTaxBehaviorUnspecified PriceCurrencyOptionsKesTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKgsTaxBehavior string

// List of values that PriceCurrencyOptionsKgsTaxBehavior can take
const (
	PriceCurrencyOptionsKgsTaxBehaviorExclusive   PriceCurrencyOptionsKgsTaxBehavior = "exclusive"
	PriceCurrencyOptionsKgsTaxBehaviorInclusive   PriceCurrencyOptionsKgsTaxBehavior = "inclusive"
	PriceCurrencyOptionsKgsTaxBehaviorUnspecified PriceCurrencyOptionsKgsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKhrTaxBehavior string

// List of values that PriceCurrencyOptionsKhrTaxBehavior can take
const (
	PriceCurrencyOptionsKhrTaxBehaviorExclusive   PriceCurrencyOptionsKhrTaxBehavior = "exclusive"
	PriceCurrencyOptionsKhrTaxBehaviorInclusive   PriceCurrencyOptionsKhrTaxBehavior = "inclusive"
	PriceCurrencyOptionsKhrTaxBehaviorUnspecified PriceCurrencyOptionsKhrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKmfTaxBehavior string

// List of values that PriceCurrencyOptionsKmfTaxBehavior can take
const (
	PriceCurrencyOptionsKmfTaxBehaviorExclusive   PriceCurrencyOptionsKmfTaxBehavior = "exclusive"
	PriceCurrencyOptionsKmfTaxBehaviorInclusive   PriceCurrencyOptionsKmfTaxBehavior = "inclusive"
	PriceCurrencyOptionsKmfTaxBehaviorUnspecified PriceCurrencyOptionsKmfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKrwTaxBehavior string

// List of values that PriceCurrencyOptionsKrwTaxBehavior can take
const (
	PriceCurrencyOptionsKrwTaxBehaviorExclusive   PriceCurrencyOptionsKrwTaxBehavior = "exclusive"
	PriceCurrencyOptionsKrwTaxBehaviorInclusive   PriceCurrencyOptionsKrwTaxBehavior = "inclusive"
	PriceCurrencyOptionsKrwTaxBehaviorUnspecified PriceCurrencyOptionsKrwTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKwdTaxBehavior string

// List of values that PriceCurrencyOptionsKwdTaxBehavior can take
const (
	PriceCurrencyOptionsKwdTaxBehaviorExclusive   PriceCurrencyOptionsKwdTaxBehavior = "exclusive"
	PriceCurrencyOptionsKwdTaxBehaviorInclusive   PriceCurrencyOptionsKwdTaxBehavior = "inclusive"
	PriceCurrencyOptionsKwdTaxBehaviorUnspecified PriceCurrencyOptionsKwdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKydTaxBehavior string

// List of values that PriceCurrencyOptionsKydTaxBehavior can take
const (
	PriceCurrencyOptionsKydTaxBehaviorExclusive   PriceCurrencyOptionsKydTaxBehavior = "exclusive"
	PriceCurrencyOptionsKydTaxBehaviorInclusive   PriceCurrencyOptionsKydTaxBehavior = "inclusive"
	PriceCurrencyOptionsKydTaxBehaviorUnspecified PriceCurrencyOptionsKydTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsKztTaxBehavior string

// List of values that PriceCurrencyOptionsKztTaxBehavior can take
const (
	PriceCurrencyOptionsKztTaxBehaviorExclusive   PriceCurrencyOptionsKztTaxBehavior = "exclusive"
	PriceCurrencyOptionsKztTaxBehaviorInclusive   PriceCurrencyOptionsKztTaxBehavior = "inclusive"
	PriceCurrencyOptionsKztTaxBehaviorUnspecified PriceCurrencyOptionsKztTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsLakTaxBehavior string

// List of values that PriceCurrencyOptionsLakTaxBehavior can take
const (
	PriceCurrencyOptionsLakTaxBehaviorExclusive   PriceCurrencyOptionsLakTaxBehavior = "exclusive"
	PriceCurrencyOptionsLakTaxBehaviorInclusive   PriceCurrencyOptionsLakTaxBehavior = "inclusive"
	PriceCurrencyOptionsLakTaxBehaviorUnspecified PriceCurrencyOptionsLakTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsLbpTaxBehavior string

// List of values that PriceCurrencyOptionsLbpTaxBehavior can take
const (
	PriceCurrencyOptionsLbpTaxBehaviorExclusive   PriceCurrencyOptionsLbpTaxBehavior = "exclusive"
	PriceCurrencyOptionsLbpTaxBehaviorInclusive   PriceCurrencyOptionsLbpTaxBehavior = "inclusive"
	PriceCurrencyOptionsLbpTaxBehaviorUnspecified PriceCurrencyOptionsLbpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsLkrTaxBehavior string

// List of values that PriceCurrencyOptionsLkrTaxBehavior can take
const (
	PriceCurrencyOptionsLkrTaxBehaviorExclusive   PriceCurrencyOptionsLkrTaxBehavior = "exclusive"
	PriceCurrencyOptionsLkrTaxBehaviorInclusive   PriceCurrencyOptionsLkrTaxBehavior = "inclusive"
	PriceCurrencyOptionsLkrTaxBehaviorUnspecified PriceCurrencyOptionsLkrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsLrdTaxBehavior string

// List of values that PriceCurrencyOptionsLrdTaxBehavior can take
const (
	PriceCurrencyOptionsLrdTaxBehaviorExclusive   PriceCurrencyOptionsLrdTaxBehavior = "exclusive"
	PriceCurrencyOptionsLrdTaxBehaviorInclusive   PriceCurrencyOptionsLrdTaxBehavior = "inclusive"
	PriceCurrencyOptionsLrdTaxBehaviorUnspecified PriceCurrencyOptionsLrdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsLslTaxBehavior string

// List of values that PriceCurrencyOptionsLslTaxBehavior can take
const (
	PriceCurrencyOptionsLslTaxBehaviorExclusive   PriceCurrencyOptionsLslTaxBehavior = "exclusive"
	PriceCurrencyOptionsLslTaxBehaviorInclusive   PriceCurrencyOptionsLslTaxBehavior = "inclusive"
	PriceCurrencyOptionsLslTaxBehaviorUnspecified PriceCurrencyOptionsLslTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMadTaxBehavior string

// List of values that PriceCurrencyOptionsMadTaxBehavior can take
const (
	PriceCurrencyOptionsMadTaxBehaviorExclusive   PriceCurrencyOptionsMadTaxBehavior = "exclusive"
	PriceCurrencyOptionsMadTaxBehaviorInclusive   PriceCurrencyOptionsMadTaxBehavior = "inclusive"
	PriceCurrencyOptionsMadTaxBehaviorUnspecified PriceCurrencyOptionsMadTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMdlTaxBehavior string

// List of values that PriceCurrencyOptionsMdlTaxBehavior can take
const (
	PriceCurrencyOptionsMdlTaxBehaviorExclusive   PriceCurrencyOptionsMdlTaxBehavior = "exclusive"
	PriceCurrencyOptionsMdlTaxBehaviorInclusive   PriceCurrencyOptionsMdlTaxBehavior = "inclusive"
	PriceCurrencyOptionsMdlTaxBehaviorUnspecified PriceCurrencyOptionsMdlTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMgaTaxBehavior string

// List of values that PriceCurrencyOptionsMgaTaxBehavior can take
const (
	PriceCurrencyOptionsMgaTaxBehaviorExclusive   PriceCurrencyOptionsMgaTaxBehavior = "exclusive"
	PriceCurrencyOptionsMgaTaxBehaviorInclusive   PriceCurrencyOptionsMgaTaxBehavior = "inclusive"
	PriceCurrencyOptionsMgaTaxBehaviorUnspecified PriceCurrencyOptionsMgaTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMkdTaxBehavior string

// List of values that PriceCurrencyOptionsMkdTaxBehavior can take
const (
	PriceCurrencyOptionsMkdTaxBehaviorExclusive   PriceCurrencyOptionsMkdTaxBehavior = "exclusive"
	PriceCurrencyOptionsMkdTaxBehaviorInclusive   PriceCurrencyOptionsMkdTaxBehavior = "inclusive"
	PriceCurrencyOptionsMkdTaxBehaviorUnspecified PriceCurrencyOptionsMkdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMmkTaxBehavior string

// List of values that PriceCurrencyOptionsMmkTaxBehavior can take
const (
	PriceCurrencyOptionsMmkTaxBehaviorExclusive   PriceCurrencyOptionsMmkTaxBehavior = "exclusive"
	PriceCurrencyOptionsMmkTaxBehaviorInclusive   PriceCurrencyOptionsMmkTaxBehavior = "inclusive"
	PriceCurrencyOptionsMmkTaxBehaviorUnspecified PriceCurrencyOptionsMmkTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMntTaxBehavior string

// List of values that PriceCurrencyOptionsMntTaxBehavior can take
const (
	PriceCurrencyOptionsMntTaxBehaviorExclusive   PriceCurrencyOptionsMntTaxBehavior = "exclusive"
	PriceCurrencyOptionsMntTaxBehaviorInclusive   PriceCurrencyOptionsMntTaxBehavior = "inclusive"
	PriceCurrencyOptionsMntTaxBehaviorUnspecified PriceCurrencyOptionsMntTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMopTaxBehavior string

// List of values that PriceCurrencyOptionsMopTaxBehavior can take
const (
	PriceCurrencyOptionsMopTaxBehaviorExclusive   PriceCurrencyOptionsMopTaxBehavior = "exclusive"
	PriceCurrencyOptionsMopTaxBehaviorInclusive   PriceCurrencyOptionsMopTaxBehavior = "inclusive"
	PriceCurrencyOptionsMopTaxBehaviorUnspecified PriceCurrencyOptionsMopTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMroTaxBehavior string

// List of values that PriceCurrencyOptionsMroTaxBehavior can take
const (
	PriceCurrencyOptionsMroTaxBehaviorExclusive   PriceCurrencyOptionsMroTaxBehavior = "exclusive"
	PriceCurrencyOptionsMroTaxBehaviorInclusive   PriceCurrencyOptionsMroTaxBehavior = "inclusive"
	PriceCurrencyOptionsMroTaxBehaviorUnspecified PriceCurrencyOptionsMroTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMurTaxBehavior string

// List of values that PriceCurrencyOptionsMurTaxBehavior can take
const (
	PriceCurrencyOptionsMurTaxBehaviorExclusive   PriceCurrencyOptionsMurTaxBehavior = "exclusive"
	PriceCurrencyOptionsMurTaxBehaviorInclusive   PriceCurrencyOptionsMurTaxBehavior = "inclusive"
	PriceCurrencyOptionsMurTaxBehaviorUnspecified PriceCurrencyOptionsMurTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMvrTaxBehavior string

// List of values that PriceCurrencyOptionsMvrTaxBehavior can take
const (
	PriceCurrencyOptionsMvrTaxBehaviorExclusive   PriceCurrencyOptionsMvrTaxBehavior = "exclusive"
	PriceCurrencyOptionsMvrTaxBehaviorInclusive   PriceCurrencyOptionsMvrTaxBehavior = "inclusive"
	PriceCurrencyOptionsMvrTaxBehaviorUnspecified PriceCurrencyOptionsMvrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMwkTaxBehavior string

// List of values that PriceCurrencyOptionsMwkTaxBehavior can take
const (
	PriceCurrencyOptionsMwkTaxBehaviorExclusive   PriceCurrencyOptionsMwkTaxBehavior = "exclusive"
	PriceCurrencyOptionsMwkTaxBehaviorInclusive   PriceCurrencyOptionsMwkTaxBehavior = "inclusive"
	PriceCurrencyOptionsMwkTaxBehaviorUnspecified PriceCurrencyOptionsMwkTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMxnTaxBehavior string

// List of values that PriceCurrencyOptionsMxnTaxBehavior can take
const (
	PriceCurrencyOptionsMxnTaxBehaviorExclusive   PriceCurrencyOptionsMxnTaxBehavior = "exclusive"
	PriceCurrencyOptionsMxnTaxBehaviorInclusive   PriceCurrencyOptionsMxnTaxBehavior = "inclusive"
	PriceCurrencyOptionsMxnTaxBehaviorUnspecified PriceCurrencyOptionsMxnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMYRTaxBehavior string

// List of values that PriceCurrencyOptionsMYRTaxBehavior can take
const (
	PriceCurrencyOptionsMYRTaxBehaviorExclusive   PriceCurrencyOptionsMYRTaxBehavior = "exclusive"
	PriceCurrencyOptionsMYRTaxBehaviorInclusive   PriceCurrencyOptionsMYRTaxBehavior = "inclusive"
	PriceCurrencyOptionsMYRTaxBehaviorUnspecified PriceCurrencyOptionsMYRTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsMznTaxBehavior string

// List of values that PriceCurrencyOptionsMznTaxBehavior can take
const (
	PriceCurrencyOptionsMznTaxBehaviorExclusive   PriceCurrencyOptionsMznTaxBehavior = "exclusive"
	PriceCurrencyOptionsMznTaxBehaviorInclusive   PriceCurrencyOptionsMznTaxBehavior = "inclusive"
	PriceCurrencyOptionsMznTaxBehaviorUnspecified PriceCurrencyOptionsMznTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNadTaxBehavior string

// List of values that PriceCurrencyOptionsNadTaxBehavior can take
const (
	PriceCurrencyOptionsNadTaxBehaviorExclusive   PriceCurrencyOptionsNadTaxBehavior = "exclusive"
	PriceCurrencyOptionsNadTaxBehaviorInclusive   PriceCurrencyOptionsNadTaxBehavior = "inclusive"
	PriceCurrencyOptionsNadTaxBehaviorUnspecified PriceCurrencyOptionsNadTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNgnTaxBehavior string

// List of values that PriceCurrencyOptionsNgnTaxBehavior can take
const (
	PriceCurrencyOptionsNgnTaxBehaviorExclusive   PriceCurrencyOptionsNgnTaxBehavior = "exclusive"
	PriceCurrencyOptionsNgnTaxBehaviorInclusive   PriceCurrencyOptionsNgnTaxBehavior = "inclusive"
	PriceCurrencyOptionsNgnTaxBehaviorUnspecified PriceCurrencyOptionsNgnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNioTaxBehavior string

// List of values that PriceCurrencyOptionsNioTaxBehavior can take
const (
	PriceCurrencyOptionsNioTaxBehaviorExclusive   PriceCurrencyOptionsNioTaxBehavior = "exclusive"
	PriceCurrencyOptionsNioTaxBehaviorInclusive   PriceCurrencyOptionsNioTaxBehavior = "inclusive"
	PriceCurrencyOptionsNioTaxBehaviorUnspecified PriceCurrencyOptionsNioTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNOKTaxBehavior string

// List of values that PriceCurrencyOptionsNOKTaxBehavior can take
const (
	PriceCurrencyOptionsNOKTaxBehaviorExclusive   PriceCurrencyOptionsNOKTaxBehavior = "exclusive"
	PriceCurrencyOptionsNOKTaxBehaviorInclusive   PriceCurrencyOptionsNOKTaxBehavior = "inclusive"
	PriceCurrencyOptionsNOKTaxBehaviorUnspecified PriceCurrencyOptionsNOKTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNprTaxBehavior string

// List of values that PriceCurrencyOptionsNprTaxBehavior can take
const (
	PriceCurrencyOptionsNprTaxBehaviorExclusive   PriceCurrencyOptionsNprTaxBehavior = "exclusive"
	PriceCurrencyOptionsNprTaxBehaviorInclusive   PriceCurrencyOptionsNprTaxBehavior = "inclusive"
	PriceCurrencyOptionsNprTaxBehaviorUnspecified PriceCurrencyOptionsNprTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsNZDTaxBehavior string

// List of values that PriceCurrencyOptionsNZDTaxBehavior can take
const (
	PriceCurrencyOptionsNZDTaxBehaviorExclusive   PriceCurrencyOptionsNZDTaxBehavior = "exclusive"
	PriceCurrencyOptionsNZDTaxBehaviorInclusive   PriceCurrencyOptionsNZDTaxBehavior = "inclusive"
	PriceCurrencyOptionsNZDTaxBehaviorUnspecified PriceCurrencyOptionsNZDTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsOmrTaxBehavior string

// List of values that PriceCurrencyOptionsOmrTaxBehavior can take
const (
	PriceCurrencyOptionsOmrTaxBehaviorExclusive   PriceCurrencyOptionsOmrTaxBehavior = "exclusive"
	PriceCurrencyOptionsOmrTaxBehaviorInclusive   PriceCurrencyOptionsOmrTaxBehavior = "inclusive"
	PriceCurrencyOptionsOmrTaxBehaviorUnspecified PriceCurrencyOptionsOmrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPabTaxBehavior string

// List of values that PriceCurrencyOptionsPabTaxBehavior can take
const (
	PriceCurrencyOptionsPabTaxBehaviorExclusive   PriceCurrencyOptionsPabTaxBehavior = "exclusive"
	PriceCurrencyOptionsPabTaxBehaviorInclusive   PriceCurrencyOptionsPabTaxBehavior = "inclusive"
	PriceCurrencyOptionsPabTaxBehaviorUnspecified PriceCurrencyOptionsPabTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPenTaxBehavior string

// List of values that PriceCurrencyOptionsPenTaxBehavior can take
const (
	PriceCurrencyOptionsPenTaxBehaviorExclusive   PriceCurrencyOptionsPenTaxBehavior = "exclusive"
	PriceCurrencyOptionsPenTaxBehaviorInclusive   PriceCurrencyOptionsPenTaxBehavior = "inclusive"
	PriceCurrencyOptionsPenTaxBehaviorUnspecified PriceCurrencyOptionsPenTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPgkTaxBehavior string

// List of values that PriceCurrencyOptionsPgkTaxBehavior can take
const (
	PriceCurrencyOptionsPgkTaxBehaviorExclusive   PriceCurrencyOptionsPgkTaxBehavior = "exclusive"
	PriceCurrencyOptionsPgkTaxBehaviorInclusive   PriceCurrencyOptionsPgkTaxBehavior = "inclusive"
	PriceCurrencyOptionsPgkTaxBehaviorUnspecified PriceCurrencyOptionsPgkTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPhpTaxBehavior string

// List of values that PriceCurrencyOptionsPhpTaxBehavior can take
const (
	PriceCurrencyOptionsPhpTaxBehaviorExclusive   PriceCurrencyOptionsPhpTaxBehavior = "exclusive"
	PriceCurrencyOptionsPhpTaxBehaviorInclusive   PriceCurrencyOptionsPhpTaxBehavior = "inclusive"
	PriceCurrencyOptionsPhpTaxBehaviorUnspecified PriceCurrencyOptionsPhpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPkrTaxBehavior string

// List of values that PriceCurrencyOptionsPkrTaxBehavior can take
const (
	PriceCurrencyOptionsPkrTaxBehaviorExclusive   PriceCurrencyOptionsPkrTaxBehavior = "exclusive"
	PriceCurrencyOptionsPkrTaxBehaviorInclusive   PriceCurrencyOptionsPkrTaxBehavior = "inclusive"
	PriceCurrencyOptionsPkrTaxBehaviorUnspecified PriceCurrencyOptionsPkrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPlnTaxBehavior string

// List of values that PriceCurrencyOptionsPlnTaxBehavior can take
const (
	PriceCurrencyOptionsPlnTaxBehaviorExclusive   PriceCurrencyOptionsPlnTaxBehavior = "exclusive"
	PriceCurrencyOptionsPlnTaxBehaviorInclusive   PriceCurrencyOptionsPlnTaxBehavior = "inclusive"
	PriceCurrencyOptionsPlnTaxBehaviorUnspecified PriceCurrencyOptionsPlnTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsPygTaxBehavior string

// List of values that PriceCurrencyOptionsPygTaxBehavior can take
const (
	PriceCurrencyOptionsPygTaxBehaviorExclusive   PriceCurrencyOptionsPygTaxBehavior = "exclusive"
	PriceCurrencyOptionsPygTaxBehaviorInclusive   PriceCurrencyOptionsPygTaxBehavior = "inclusive"
	PriceCurrencyOptionsPygTaxBehaviorUnspecified PriceCurrencyOptionsPygTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsQarTaxBehavior string

// List of values that PriceCurrencyOptionsQarTaxBehavior can take
const (
	PriceCurrencyOptionsQarTaxBehaviorExclusive   PriceCurrencyOptionsQarTaxBehavior = "exclusive"
	PriceCurrencyOptionsQarTaxBehaviorInclusive   PriceCurrencyOptionsQarTaxBehavior = "inclusive"
	PriceCurrencyOptionsQarTaxBehaviorUnspecified PriceCurrencyOptionsQarTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsRonTaxBehavior string

// List of values that PriceCurrencyOptionsRonTaxBehavior can take
const (
	PriceCurrencyOptionsRonTaxBehaviorExclusive   PriceCurrencyOptionsRonTaxBehavior = "exclusive"
	PriceCurrencyOptionsRonTaxBehaviorInclusive   PriceCurrencyOptionsRonTaxBehavior = "inclusive"
	PriceCurrencyOptionsRonTaxBehaviorUnspecified PriceCurrencyOptionsRonTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsRsdTaxBehavior string

// List of values that PriceCurrencyOptionsRsdTaxBehavior can take
const (
	PriceCurrencyOptionsRsdTaxBehaviorExclusive   PriceCurrencyOptionsRsdTaxBehavior = "exclusive"
	PriceCurrencyOptionsRsdTaxBehaviorInclusive   PriceCurrencyOptionsRsdTaxBehavior = "inclusive"
	PriceCurrencyOptionsRsdTaxBehaviorUnspecified PriceCurrencyOptionsRsdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsRubTaxBehavior string

// List of values that PriceCurrencyOptionsRubTaxBehavior can take
const (
	PriceCurrencyOptionsRubTaxBehaviorExclusive   PriceCurrencyOptionsRubTaxBehavior = "exclusive"
	PriceCurrencyOptionsRubTaxBehaviorInclusive   PriceCurrencyOptionsRubTaxBehavior = "inclusive"
	PriceCurrencyOptionsRubTaxBehaviorUnspecified PriceCurrencyOptionsRubTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsRwfTaxBehavior string

// List of values that PriceCurrencyOptionsRwfTaxBehavior can take
const (
	PriceCurrencyOptionsRwfTaxBehaviorExclusive   PriceCurrencyOptionsRwfTaxBehavior = "exclusive"
	PriceCurrencyOptionsRwfTaxBehaviorInclusive   PriceCurrencyOptionsRwfTaxBehavior = "inclusive"
	PriceCurrencyOptionsRwfTaxBehaviorUnspecified PriceCurrencyOptionsRwfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSarTaxBehavior string

// List of values that PriceCurrencyOptionsSarTaxBehavior can take
const (
	PriceCurrencyOptionsSarTaxBehaviorExclusive   PriceCurrencyOptionsSarTaxBehavior = "exclusive"
	PriceCurrencyOptionsSarTaxBehaviorInclusive   PriceCurrencyOptionsSarTaxBehavior = "inclusive"
	PriceCurrencyOptionsSarTaxBehaviorUnspecified PriceCurrencyOptionsSarTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSbdTaxBehavior string

// List of values that PriceCurrencyOptionsSbdTaxBehavior can take
const (
	PriceCurrencyOptionsSbdTaxBehaviorExclusive   PriceCurrencyOptionsSbdTaxBehavior = "exclusive"
	PriceCurrencyOptionsSbdTaxBehaviorInclusive   PriceCurrencyOptionsSbdTaxBehavior = "inclusive"
	PriceCurrencyOptionsSbdTaxBehaviorUnspecified PriceCurrencyOptionsSbdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsScrTaxBehavior string

// List of values that PriceCurrencyOptionsScrTaxBehavior can take
const (
	PriceCurrencyOptionsScrTaxBehaviorExclusive   PriceCurrencyOptionsScrTaxBehavior = "exclusive"
	PriceCurrencyOptionsScrTaxBehaviorInclusive   PriceCurrencyOptionsScrTaxBehavior = "inclusive"
	PriceCurrencyOptionsScrTaxBehaviorUnspecified PriceCurrencyOptionsScrTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSEKTaxBehavior string

// List of values that PriceCurrencyOptionsSEKTaxBehavior can take
const (
	PriceCurrencyOptionsSEKTaxBehaviorExclusive   PriceCurrencyOptionsSEKTaxBehavior = "exclusive"
	PriceCurrencyOptionsSEKTaxBehaviorInclusive   PriceCurrencyOptionsSEKTaxBehavior = "inclusive"
	PriceCurrencyOptionsSEKTaxBehaviorUnspecified PriceCurrencyOptionsSEKTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSGDTaxBehavior string

// List of values that PriceCurrencyOptionsSGDTaxBehavior can take
const (
	PriceCurrencyOptionsSGDTaxBehaviorExclusive   PriceCurrencyOptionsSGDTaxBehavior = "exclusive"
	PriceCurrencyOptionsSGDTaxBehaviorInclusive   PriceCurrencyOptionsSGDTaxBehavior = "inclusive"
	PriceCurrencyOptionsSGDTaxBehaviorUnspecified PriceCurrencyOptionsSGDTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsShpTaxBehavior string

// List of values that PriceCurrencyOptionsShpTaxBehavior can take
const (
	PriceCurrencyOptionsShpTaxBehaviorExclusive   PriceCurrencyOptionsShpTaxBehavior = "exclusive"
	PriceCurrencyOptionsShpTaxBehaviorInclusive   PriceCurrencyOptionsShpTaxBehavior = "inclusive"
	PriceCurrencyOptionsShpTaxBehaviorUnspecified PriceCurrencyOptionsShpTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSllTaxBehavior string

// List of values that PriceCurrencyOptionsSllTaxBehavior can take
const (
	PriceCurrencyOptionsSllTaxBehaviorExclusive   PriceCurrencyOptionsSllTaxBehavior = "exclusive"
	PriceCurrencyOptionsSllTaxBehaviorInclusive   PriceCurrencyOptionsSllTaxBehavior = "inclusive"
	PriceCurrencyOptionsSllTaxBehaviorUnspecified PriceCurrencyOptionsSllTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSosTaxBehavior string

// List of values that PriceCurrencyOptionsSosTaxBehavior can take
const (
	PriceCurrencyOptionsSosTaxBehaviorExclusive   PriceCurrencyOptionsSosTaxBehavior = "exclusive"
	PriceCurrencyOptionsSosTaxBehaviorInclusive   PriceCurrencyOptionsSosTaxBehavior = "inclusive"
	PriceCurrencyOptionsSosTaxBehaviorUnspecified PriceCurrencyOptionsSosTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSrdTaxBehavior string

// List of values that PriceCurrencyOptionsSrdTaxBehavior can take
const (
	PriceCurrencyOptionsSrdTaxBehaviorExclusive   PriceCurrencyOptionsSrdTaxBehavior = "exclusive"
	PriceCurrencyOptionsSrdTaxBehaviorInclusive   PriceCurrencyOptionsSrdTaxBehavior = "inclusive"
	PriceCurrencyOptionsSrdTaxBehaviorUnspecified PriceCurrencyOptionsSrdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsStdTaxBehavior string

// List of values that PriceCurrencyOptionsStdTaxBehavior can take
const (
	PriceCurrencyOptionsStdTaxBehaviorExclusive   PriceCurrencyOptionsStdTaxBehavior = "exclusive"
	PriceCurrencyOptionsStdTaxBehaviorInclusive   PriceCurrencyOptionsStdTaxBehavior = "inclusive"
	PriceCurrencyOptionsStdTaxBehaviorUnspecified PriceCurrencyOptionsStdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsSzlTaxBehavior string

// List of values that PriceCurrencyOptionsSzlTaxBehavior can take
const (
	PriceCurrencyOptionsSzlTaxBehaviorExclusive   PriceCurrencyOptionsSzlTaxBehavior = "exclusive"
	PriceCurrencyOptionsSzlTaxBehaviorInclusive   PriceCurrencyOptionsSzlTaxBehavior = "inclusive"
	PriceCurrencyOptionsSzlTaxBehaviorUnspecified PriceCurrencyOptionsSzlTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsThbTaxBehavior string

// List of values that PriceCurrencyOptionsThbTaxBehavior can take
const (
	PriceCurrencyOptionsThbTaxBehaviorExclusive   PriceCurrencyOptionsThbTaxBehavior = "exclusive"
	PriceCurrencyOptionsThbTaxBehaviorInclusive   PriceCurrencyOptionsThbTaxBehavior = "inclusive"
	PriceCurrencyOptionsThbTaxBehaviorUnspecified PriceCurrencyOptionsThbTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTjsTaxBehavior string

// List of values that PriceCurrencyOptionsTjsTaxBehavior can take
const (
	PriceCurrencyOptionsTjsTaxBehaviorExclusive   PriceCurrencyOptionsTjsTaxBehavior = "exclusive"
	PriceCurrencyOptionsTjsTaxBehaviorInclusive   PriceCurrencyOptionsTjsTaxBehavior = "inclusive"
	PriceCurrencyOptionsTjsTaxBehaviorUnspecified PriceCurrencyOptionsTjsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTndTaxBehavior string

// List of values that PriceCurrencyOptionsTndTaxBehavior can take
const (
	PriceCurrencyOptionsTndTaxBehaviorExclusive   PriceCurrencyOptionsTndTaxBehavior = "exclusive"
	PriceCurrencyOptionsTndTaxBehaviorInclusive   PriceCurrencyOptionsTndTaxBehavior = "inclusive"
	PriceCurrencyOptionsTndTaxBehaviorUnspecified PriceCurrencyOptionsTndTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTopTaxBehavior string

// List of values that PriceCurrencyOptionsTopTaxBehavior can take
const (
	PriceCurrencyOptionsTopTaxBehaviorExclusive   PriceCurrencyOptionsTopTaxBehavior = "exclusive"
	PriceCurrencyOptionsTopTaxBehaviorInclusive   PriceCurrencyOptionsTopTaxBehavior = "inclusive"
	PriceCurrencyOptionsTopTaxBehaviorUnspecified PriceCurrencyOptionsTopTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTryTaxBehavior string

// List of values that PriceCurrencyOptionsTryTaxBehavior can take
const (
	PriceCurrencyOptionsTryTaxBehaviorExclusive   PriceCurrencyOptionsTryTaxBehavior = "exclusive"
	PriceCurrencyOptionsTryTaxBehaviorInclusive   PriceCurrencyOptionsTryTaxBehavior = "inclusive"
	PriceCurrencyOptionsTryTaxBehaviorUnspecified PriceCurrencyOptionsTryTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTtdTaxBehavior string

// List of values that PriceCurrencyOptionsTtdTaxBehavior can take
const (
	PriceCurrencyOptionsTtdTaxBehaviorExclusive   PriceCurrencyOptionsTtdTaxBehavior = "exclusive"
	PriceCurrencyOptionsTtdTaxBehaviorInclusive   PriceCurrencyOptionsTtdTaxBehavior = "inclusive"
	PriceCurrencyOptionsTtdTaxBehaviorUnspecified PriceCurrencyOptionsTtdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTwdTaxBehavior string

// List of values that PriceCurrencyOptionsTwdTaxBehavior can take
const (
	PriceCurrencyOptionsTwdTaxBehaviorExclusive   PriceCurrencyOptionsTwdTaxBehavior = "exclusive"
	PriceCurrencyOptionsTwdTaxBehaviorInclusive   PriceCurrencyOptionsTwdTaxBehavior = "inclusive"
	PriceCurrencyOptionsTwdTaxBehaviorUnspecified PriceCurrencyOptionsTwdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsTzsTaxBehavior string

// List of values that PriceCurrencyOptionsTzsTaxBehavior can take
const (
	PriceCurrencyOptionsTzsTaxBehaviorExclusive   PriceCurrencyOptionsTzsTaxBehavior = "exclusive"
	PriceCurrencyOptionsTzsTaxBehaviorInclusive   PriceCurrencyOptionsTzsTaxBehavior = "inclusive"
	PriceCurrencyOptionsTzsTaxBehaviorUnspecified PriceCurrencyOptionsTzsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUahTaxBehavior string

// List of values that PriceCurrencyOptionsUahTaxBehavior can take
const (
	PriceCurrencyOptionsUahTaxBehaviorExclusive   PriceCurrencyOptionsUahTaxBehavior = "exclusive"
	PriceCurrencyOptionsUahTaxBehaviorInclusive   PriceCurrencyOptionsUahTaxBehavior = "inclusive"
	PriceCurrencyOptionsUahTaxBehaviorUnspecified PriceCurrencyOptionsUahTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUgxTaxBehavior string

// List of values that PriceCurrencyOptionsUgxTaxBehavior can take
const (
	PriceCurrencyOptionsUgxTaxBehaviorExclusive   PriceCurrencyOptionsUgxTaxBehavior = "exclusive"
	PriceCurrencyOptionsUgxTaxBehaviorInclusive   PriceCurrencyOptionsUgxTaxBehavior = "inclusive"
	PriceCurrencyOptionsUgxTaxBehaviorUnspecified PriceCurrencyOptionsUgxTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUSDTaxBehavior string

// List of values that PriceCurrencyOptionsUSDTaxBehavior can take
const (
	PriceCurrencyOptionsUSDTaxBehaviorExclusive   PriceCurrencyOptionsUSDTaxBehavior = "exclusive"
	PriceCurrencyOptionsUSDTaxBehaviorInclusive   PriceCurrencyOptionsUSDTaxBehavior = "inclusive"
	PriceCurrencyOptionsUSDTaxBehaviorUnspecified PriceCurrencyOptionsUSDTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUsdcTaxBehavior string

// List of values that PriceCurrencyOptionsUsdcTaxBehavior can take
const (
	PriceCurrencyOptionsUsdcTaxBehaviorExclusive   PriceCurrencyOptionsUsdcTaxBehavior = "exclusive"
	PriceCurrencyOptionsUsdcTaxBehaviorInclusive   PriceCurrencyOptionsUsdcTaxBehavior = "inclusive"
	PriceCurrencyOptionsUsdcTaxBehaviorUnspecified PriceCurrencyOptionsUsdcTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUyuTaxBehavior string

// List of values that PriceCurrencyOptionsUyuTaxBehavior can take
const (
	PriceCurrencyOptionsUyuTaxBehaviorExclusive   PriceCurrencyOptionsUyuTaxBehavior = "exclusive"
	PriceCurrencyOptionsUyuTaxBehaviorInclusive   PriceCurrencyOptionsUyuTaxBehavior = "inclusive"
	PriceCurrencyOptionsUyuTaxBehaviorUnspecified PriceCurrencyOptionsUyuTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsUzsTaxBehavior string

// List of values that PriceCurrencyOptionsUzsTaxBehavior can take
const (
	PriceCurrencyOptionsUzsTaxBehaviorExclusive   PriceCurrencyOptionsUzsTaxBehavior = "exclusive"
	PriceCurrencyOptionsUzsTaxBehaviorInclusive   PriceCurrencyOptionsUzsTaxBehavior = "inclusive"
	PriceCurrencyOptionsUzsTaxBehaviorUnspecified PriceCurrencyOptionsUzsTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsVndTaxBehavior string

// List of values that PriceCurrencyOptionsVndTaxBehavior can take
const (
	PriceCurrencyOptionsVndTaxBehaviorExclusive   PriceCurrencyOptionsVndTaxBehavior = "exclusive"
	PriceCurrencyOptionsVndTaxBehaviorInclusive   PriceCurrencyOptionsVndTaxBehavior = "inclusive"
	PriceCurrencyOptionsVndTaxBehaviorUnspecified PriceCurrencyOptionsVndTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsVuvTaxBehavior string

// List of values that PriceCurrencyOptionsVuvTaxBehavior can take
const (
	PriceCurrencyOptionsVuvTaxBehaviorExclusive   PriceCurrencyOptionsVuvTaxBehavior = "exclusive"
	PriceCurrencyOptionsVuvTaxBehaviorInclusive   PriceCurrencyOptionsVuvTaxBehavior = "inclusive"
	PriceCurrencyOptionsVuvTaxBehaviorUnspecified PriceCurrencyOptionsVuvTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsWstTaxBehavior string

// List of values that PriceCurrencyOptionsWstTaxBehavior can take
const (
	PriceCurrencyOptionsWstTaxBehaviorExclusive   PriceCurrencyOptionsWstTaxBehavior = "exclusive"
	PriceCurrencyOptionsWstTaxBehaviorInclusive   PriceCurrencyOptionsWstTaxBehavior = "inclusive"
	PriceCurrencyOptionsWstTaxBehaviorUnspecified PriceCurrencyOptionsWstTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsXafTaxBehavior string

// List of values that PriceCurrencyOptionsXafTaxBehavior can take
const (
	PriceCurrencyOptionsXafTaxBehaviorExclusive   PriceCurrencyOptionsXafTaxBehavior = "exclusive"
	PriceCurrencyOptionsXafTaxBehaviorInclusive   PriceCurrencyOptionsXafTaxBehavior = "inclusive"
	PriceCurrencyOptionsXafTaxBehaviorUnspecified PriceCurrencyOptionsXafTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsXcdTaxBehavior string

// List of values that PriceCurrencyOptionsXcdTaxBehavior can take
const (
	PriceCurrencyOptionsXcdTaxBehaviorExclusive   PriceCurrencyOptionsXcdTaxBehavior = "exclusive"
	PriceCurrencyOptionsXcdTaxBehaviorInclusive   PriceCurrencyOptionsXcdTaxBehavior = "inclusive"
	PriceCurrencyOptionsXcdTaxBehaviorUnspecified PriceCurrencyOptionsXcdTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsXofTaxBehavior string

// List of values that PriceCurrencyOptionsXofTaxBehavior can take
const (
	PriceCurrencyOptionsXofTaxBehaviorExclusive   PriceCurrencyOptionsXofTaxBehavior = "exclusive"
	PriceCurrencyOptionsXofTaxBehaviorInclusive   PriceCurrencyOptionsXofTaxBehavior = "inclusive"
	PriceCurrencyOptionsXofTaxBehaviorUnspecified PriceCurrencyOptionsXofTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsXpfTaxBehavior string

// List of values that PriceCurrencyOptionsXpfTaxBehavior can take
const (
	PriceCurrencyOptionsXpfTaxBehaviorExclusive   PriceCurrencyOptionsXpfTaxBehavior = "exclusive"
	PriceCurrencyOptionsXpfTaxBehaviorInclusive   PriceCurrencyOptionsXpfTaxBehavior = "inclusive"
	PriceCurrencyOptionsXpfTaxBehaviorUnspecified PriceCurrencyOptionsXpfTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsYerTaxBehavior string

// List of values that PriceCurrencyOptionsYerTaxBehavior can take
const (
	PriceCurrencyOptionsYerTaxBehaviorExclusive   PriceCurrencyOptionsYerTaxBehavior = "exclusive"
	PriceCurrencyOptionsYerTaxBehaviorInclusive   PriceCurrencyOptionsYerTaxBehavior = "inclusive"
	PriceCurrencyOptionsYerTaxBehaviorUnspecified PriceCurrencyOptionsYerTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsZarTaxBehavior string

// List of values that PriceCurrencyOptionsZarTaxBehavior can take
const (
	PriceCurrencyOptionsZarTaxBehaviorExclusive   PriceCurrencyOptionsZarTaxBehavior = "exclusive"
	PriceCurrencyOptionsZarTaxBehaviorInclusive   PriceCurrencyOptionsZarTaxBehavior = "inclusive"
	PriceCurrencyOptionsZarTaxBehaviorUnspecified PriceCurrencyOptionsZarTaxBehavior = "unspecified"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceCurrencyOptionsZmwTaxBehavior string

// List of values that PriceCurrencyOptionsZmwTaxBehavior can take
const (
	PriceCurrencyOptionsZmwTaxBehaviorExclusive   PriceCurrencyOptionsZmwTaxBehavior = "exclusive"
	PriceCurrencyOptionsZmwTaxBehaviorInclusive   PriceCurrencyOptionsZmwTaxBehavior = "inclusive"
	PriceCurrencyOptionsZmwTaxBehaviorUnspecified PriceCurrencyOptionsZmwTaxBehavior = "unspecified"
)

// Specifies a usage aggregation strategy for prices of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
type PriceRecurringAggregateUsage string

// List of values that PriceRecurringAggregateUsage can take
const (
	PriceRecurringAggregateUsageLastDuringPeriod PriceRecurringAggregateUsage = "last_during_period"
	PriceRecurringAggregateUsageLastEver         PriceRecurringAggregateUsage = "last_ever"
	PriceRecurringAggregateUsageMax              PriceRecurringAggregateUsage = "max"
	PriceRecurringAggregateUsageSum              PriceRecurringAggregateUsage = "sum"
)

// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
type PriceRecurringInterval string

// List of values that PriceRecurringInterval can take
const (
	PriceRecurringIntervalDay   PriceRecurringInterval = "day"
	PriceRecurringIntervalMonth PriceRecurringInterval = "month"
	PriceRecurringIntervalWeek  PriceRecurringInterval = "week"
	PriceRecurringIntervalYear  PriceRecurringInterval = "year"
)

// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
type PriceRecurringUsageType string

// List of values that PriceRecurringUsageType can take
const (
	PriceRecurringUsageTypeLicensed PriceRecurringUsageType = "licensed"
	PriceRecurringUsageTypeMetered  PriceRecurringUsageType = "metered"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceTaxBehavior string

// List of values that PriceTaxBehavior can take
const (
	PriceTaxBehaviorExclusive   PriceTaxBehavior = "exclusive"
	PriceTaxBehaviorInclusive   PriceTaxBehavior = "inclusive"
	PriceTaxBehaviorUnspecified PriceTaxBehavior = "unspecified"
)

// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
type PriceTiersMode string

// List of values that PriceTiersMode can take
const (
	PriceTiersModeGraduated PriceTiersMode = "graduated"
	PriceTiersModeVolume    PriceTiersMode = "volume"
)

// After division, either round the result `up` or `down`.
type PriceTransformQuantityRound string

// List of values that PriceTransformQuantityRound can take
const (
	PriceTransformQuantityRoundDown PriceTransformQuantityRound = "down"
	PriceTransformQuantityRoundUp   PriceTransformQuantityRound = "up"
)

// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
type PriceType string

// List of values that PriceType can take
const (
	PriceTypeOneTime   PriceType = "one_time"
	PriceTypeRecurring PriceType = "recurring"
)

// Search for prices you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type PriceSearchParams struct {
	SearchParams `form:"*"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// Only return prices with these recurring fields.
type PriceRecurringListParams struct {
	// Filter by billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// Filter by the usage type for this price. Can be either `metered` or `licensed`.
	UsageType *string `form:"usage_type"`
}

// Returns a list of your prices.
type PriceListParams struct {
	ListParams `form:"*"`
	// Only return prices that are active or inactive (e.g., pass `false` to list all inactive prices).
	Active *bool `form:"active"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return prices for the given currency.
	Currency *string `form:"currency"`
	// Only return the price with these lookup_keys, if any exist.
	LookupKeys []*string `form:"lookup_keys"`
	// Only return prices for the given product.
	Product *string `form:"product"`
	// Only return prices with these recurring fields.
	Recurring *PriceRecurringListParams `form:"recurring"`
	// Only return prices of type `recurring` or `one_time`.
	Type *string `form:"type"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAedCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAedTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAedTierParams.
func (p *PriceCurrencyOptionsAedTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AED.
type PriceCurrencyOptionsAedParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAedCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAedTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAfnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAfnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAfnTierParams.
func (p *PriceCurrencyOptionsAfnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AFN.
type PriceCurrencyOptionsAfnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAfnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAfnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAllCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAllTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAllTierParams.
func (p *PriceCurrencyOptionsAllTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ALL.
type PriceCurrencyOptionsAllParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAllCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAllTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAmdTierParams.
func (p *PriceCurrencyOptionsAmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AMD.
type PriceCurrencyOptionsAmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAngCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAngTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAngTierParams.
func (p *PriceCurrencyOptionsAngTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ANG.
type PriceCurrencyOptionsAngParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAngCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAngTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAoaCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAoaTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAoaTierParams.
func (p *PriceCurrencyOptionsAoaTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AOA.
type PriceCurrencyOptionsAoaParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAoaCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAoaTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsArsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsArsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsArsTierParams.
func (p *PriceCurrencyOptionsArsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ARS.
type PriceCurrencyOptionsArsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsArsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsArsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAUDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAUDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAUDTierParams.
func (p *PriceCurrencyOptionsAUDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AUD.
type PriceCurrencyOptionsAUDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAUDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAUDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAwgCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAwgTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAwgTierParams.
func (p *PriceCurrencyOptionsAwgTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AWG.
type PriceCurrencyOptionsAwgParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAwgCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAwgTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAznCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAznTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsAznTierParams.
func (p *PriceCurrencyOptionsAznTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AZN.
type PriceCurrencyOptionsAznParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAznCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAznTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBamCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBamTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBamTierParams.
func (p *PriceCurrencyOptionsBamTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BAM.
type PriceCurrencyOptionsBamParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBamCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBamTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBbdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBbdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBbdTierParams.
func (p *PriceCurrencyOptionsBbdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BBD.
type PriceCurrencyOptionsBbdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBbdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBbdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBdtCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBdtTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBdtTierParams.
func (p *PriceCurrencyOptionsBdtTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BDT.
type PriceCurrencyOptionsBdtParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBdtCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBdtTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBgnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBgnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBgnTierParams.
func (p *PriceCurrencyOptionsBgnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BGN.
type PriceCurrencyOptionsBgnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBgnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBgnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBhdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBhdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBhdTierParams.
func (p *PriceCurrencyOptionsBhdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BHD.
type PriceCurrencyOptionsBhdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBhdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBhdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBifCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBifTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBifTierParams.
func (p *PriceCurrencyOptionsBifTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BIF.
type PriceCurrencyOptionsBifParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBifCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBifTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBmdTierParams.
func (p *PriceCurrencyOptionsBmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BMD.
type PriceCurrencyOptionsBmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBndTierParams.
func (p *PriceCurrencyOptionsBndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BND.
type PriceCurrencyOptionsBndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBobCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBobTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBobTierParams.
func (p *PriceCurrencyOptionsBobTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BOB.
type PriceCurrencyOptionsBobParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBobCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBobTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBrlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBrlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBrlTierParams.
func (p *PriceCurrencyOptionsBrlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BRL.
type PriceCurrencyOptionsBrlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBrlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBrlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBsdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBsdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBsdTierParams.
func (p *PriceCurrencyOptionsBsdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BSD.
type PriceCurrencyOptionsBsdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBsdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBsdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBtnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBtnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBtnTierParams.
func (p *PriceCurrencyOptionsBtnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BTN.
type PriceCurrencyOptionsBtnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBtnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBtnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBwpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBwpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBwpTierParams.
func (p *PriceCurrencyOptionsBwpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BWP.
type PriceCurrencyOptionsBwpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBwpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBwpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBynCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBynTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBynTierParams.
func (p *PriceCurrencyOptionsBynTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BYN.
type PriceCurrencyOptionsBynParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBynCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBynTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBzdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBzdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsBzdTierParams.
func (p *PriceCurrencyOptionsBzdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BZD.
type PriceCurrencyOptionsBzdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBzdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBzdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCADCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCADTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCADTierParams.
func (p *PriceCurrencyOptionsCADTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CAD.
type PriceCurrencyOptionsCADParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCADCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCADTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCdfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCdfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCdfTierParams.
func (p *PriceCurrencyOptionsCdfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CDF.
type PriceCurrencyOptionsCdfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCdfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCdfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCHFCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCHFTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCHFTierParams.
func (p *PriceCurrencyOptionsCHFTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CHF.
type PriceCurrencyOptionsCHFParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCHFCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCHFTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsClpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsClpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsClpTierParams.
func (p *PriceCurrencyOptionsClpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CLP.
type PriceCurrencyOptionsClpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsClpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsClpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCnyCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCnyTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCnyTierParams.
func (p *PriceCurrencyOptionsCnyTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CNY.
type PriceCurrencyOptionsCnyParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCnyCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCnyTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCopTierParams.
func (p *PriceCurrencyOptionsCopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in COP.
type PriceCurrencyOptionsCopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCrcCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCrcTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCrcTierParams.
func (p *PriceCurrencyOptionsCrcTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CRC.
type PriceCurrencyOptionsCrcParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCrcCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCrcTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCveCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCveTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCveTierParams.
func (p *PriceCurrencyOptionsCveTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CVE.
type PriceCurrencyOptionsCveParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCveCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCveTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCZKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCZKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsCZKTierParams.
func (p *PriceCurrencyOptionsCZKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CZK.
type PriceCurrencyOptionsCZKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCZKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCZKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDjfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDjfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsDjfTierParams.
func (p *PriceCurrencyOptionsDjfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DJF.
type PriceCurrencyOptionsDjfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDjfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDjfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDKKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDKKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsDKKTierParams.
func (p *PriceCurrencyOptionsDKKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DKK.
type PriceCurrencyOptionsDKKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDKKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDKKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsDopTierParams.
func (p *PriceCurrencyOptionsDopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DOP.
type PriceCurrencyOptionsDopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDzdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDzdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsDzdTierParams.
func (p *PriceCurrencyOptionsDzdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DZD.
type PriceCurrencyOptionsDzdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDzdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDzdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEgpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEgpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsEgpTierParams.
func (p *PriceCurrencyOptionsEgpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in EGP.
type PriceCurrencyOptionsEgpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEgpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEgpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEtbCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEtbTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsEtbTierParams.
func (p *PriceCurrencyOptionsEtbTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ETB.
type PriceCurrencyOptionsEtbParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEtbCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEtbTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEURCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEURTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsEURTierParams.
func (p *PriceCurrencyOptionsEURTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in EUR.
type PriceCurrencyOptionsEURParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEURCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEURTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsFjdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsFjdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsFjdTierParams.
func (p *PriceCurrencyOptionsFjdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in FJD.
type PriceCurrencyOptionsFjdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsFjdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsFjdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsFkpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsFkpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsFkpTierParams.
func (p *PriceCurrencyOptionsFkpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in FKP.
type PriceCurrencyOptionsFkpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsFkpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsFkpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGBPCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGBPTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGBPTierParams.
func (p *PriceCurrencyOptionsGBPTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GBP.
type PriceCurrencyOptionsGBPParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGBPCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGBPTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGelCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGelTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGelTierParams.
func (p *PriceCurrencyOptionsGelTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GEL.
type PriceCurrencyOptionsGelParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGelCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGelTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGhsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGhsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGhsTierParams.
func (p *PriceCurrencyOptionsGhsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GHS.
type PriceCurrencyOptionsGhsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGhsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGhsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGipCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGipTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGipTierParams.
func (p *PriceCurrencyOptionsGipTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GIP.
type PriceCurrencyOptionsGipParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGipCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGipTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGmdTierParams.
func (p *PriceCurrencyOptionsGmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GMD.
type PriceCurrencyOptionsGmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGnfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGnfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGnfTierParams.
func (p *PriceCurrencyOptionsGnfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GNF.
type PriceCurrencyOptionsGnfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGnfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGnfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGtqCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGtqTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGtqTierParams.
func (p *PriceCurrencyOptionsGtqTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GTQ.
type PriceCurrencyOptionsGtqParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGtqCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGtqTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGydCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGydTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsGydTierParams.
func (p *PriceCurrencyOptionsGydTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GYD.
type PriceCurrencyOptionsGydParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGydCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGydTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHKDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHKDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsHKDTierParams.
func (p *PriceCurrencyOptionsHKDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HKD.
type PriceCurrencyOptionsHKDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHKDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHKDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHnlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHnlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsHnlTierParams.
func (p *PriceCurrencyOptionsHnlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HNL.
type PriceCurrencyOptionsHnlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHnlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHnlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHrkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHrkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsHrkTierParams.
func (p *PriceCurrencyOptionsHrkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HRK.
type PriceCurrencyOptionsHrkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHrkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHrkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHtgCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHtgTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsHtgTierParams.
func (p *PriceCurrencyOptionsHtgTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HTG.
type PriceCurrencyOptionsHtgParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHtgCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHtgTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHufCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHufTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsHufTierParams.
func (p *PriceCurrencyOptionsHufTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HUF.
type PriceCurrencyOptionsHufParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHufCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHufTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIdrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIdrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsIdrTierParams.
func (p *PriceCurrencyOptionsIdrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in IDR.
type PriceCurrencyOptionsIdrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIdrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIdrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIlsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIlsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsIlsTierParams.
func (p *PriceCurrencyOptionsIlsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ILS.
type PriceCurrencyOptionsIlsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIlsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIlsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsInrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsInrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsInrTierParams.
func (p *PriceCurrencyOptionsInrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in INR.
type PriceCurrencyOptionsInrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsInrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsInrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIskCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIskTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsIskTierParams.
func (p *PriceCurrencyOptionsIskTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ISK.
type PriceCurrencyOptionsIskParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIskCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIskTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsJmdTierParams.
func (p *PriceCurrencyOptionsJmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JMD.
type PriceCurrencyOptionsJmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJodCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJodTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsJodTierParams.
func (p *PriceCurrencyOptionsJodTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JOD.
type PriceCurrencyOptionsJodParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJodCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJodTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJpyCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJpyTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsJpyTierParams.
func (p *PriceCurrencyOptionsJpyTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JPY.
type PriceCurrencyOptionsJpyParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJpyCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJpyTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKesCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKesTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKesTierParams.
func (p *PriceCurrencyOptionsKesTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KES.
type PriceCurrencyOptionsKesParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKesCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKesTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKgsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKgsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKgsTierParams.
func (p *PriceCurrencyOptionsKgsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KGS.
type PriceCurrencyOptionsKgsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKgsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKgsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKhrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKhrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKhrTierParams.
func (p *PriceCurrencyOptionsKhrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KHR.
type PriceCurrencyOptionsKhrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKhrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKhrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKmfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKmfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKmfTierParams.
func (p *PriceCurrencyOptionsKmfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KMF.
type PriceCurrencyOptionsKmfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKmfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKmfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKrwCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKrwTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKrwTierParams.
func (p *PriceCurrencyOptionsKrwTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KRW.
type PriceCurrencyOptionsKrwParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKrwCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKrwTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKwdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKwdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKwdTierParams.
func (p *PriceCurrencyOptionsKwdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KWD.
type PriceCurrencyOptionsKwdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKwdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKwdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKydCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKydTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKydTierParams.
func (p *PriceCurrencyOptionsKydTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KYD.
type PriceCurrencyOptionsKydParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKydCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKydTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKztCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKztTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsKztTierParams.
func (p *PriceCurrencyOptionsKztTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KZT.
type PriceCurrencyOptionsKztParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKztCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKztTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLakCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLakTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsLakTierParams.
func (p *PriceCurrencyOptionsLakTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LAK.
type PriceCurrencyOptionsLakParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLakCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLakTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLbpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLbpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsLbpTierParams.
func (p *PriceCurrencyOptionsLbpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LBP.
type PriceCurrencyOptionsLbpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLbpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLbpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLkrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLkrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsLkrTierParams.
func (p *PriceCurrencyOptionsLkrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LKR.
type PriceCurrencyOptionsLkrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLkrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLkrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLrdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLrdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsLrdTierParams.
func (p *PriceCurrencyOptionsLrdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LRD.
type PriceCurrencyOptionsLrdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLrdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLrdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLslCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLslTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsLslTierParams.
func (p *PriceCurrencyOptionsLslTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LSL.
type PriceCurrencyOptionsLslParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLslCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLslTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMadCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMadTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMadTierParams.
func (p *PriceCurrencyOptionsMadTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MAD.
type PriceCurrencyOptionsMadParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMadCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMadTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMdlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMdlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMdlTierParams.
func (p *PriceCurrencyOptionsMdlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MDL.
type PriceCurrencyOptionsMdlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMdlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMdlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMgaCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMgaTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMgaTierParams.
func (p *PriceCurrencyOptionsMgaTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MGA.
type PriceCurrencyOptionsMgaParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMgaCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMgaTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMkdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMkdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMkdTierParams.
func (p *PriceCurrencyOptionsMkdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MKD.
type PriceCurrencyOptionsMkdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMkdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMkdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMmkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMmkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMmkTierParams.
func (p *PriceCurrencyOptionsMmkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MMK.
type PriceCurrencyOptionsMmkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMmkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMmkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMntCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMntTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMntTierParams.
func (p *PriceCurrencyOptionsMntTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MNT.
type PriceCurrencyOptionsMntParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMntCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMntTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMopTierParams.
func (p *PriceCurrencyOptionsMopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MOP.
type PriceCurrencyOptionsMopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMroCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMroTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMroTierParams.
func (p *PriceCurrencyOptionsMroTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MRO.
type PriceCurrencyOptionsMroParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMroCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMroTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMurCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMurTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMurTierParams.
func (p *PriceCurrencyOptionsMurTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MUR.
type PriceCurrencyOptionsMurParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMurCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMurTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMvrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMvrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMvrTierParams.
func (p *PriceCurrencyOptionsMvrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MVR.
type PriceCurrencyOptionsMvrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMvrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMvrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMwkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMwkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMwkTierParams.
func (p *PriceCurrencyOptionsMwkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MWK.
type PriceCurrencyOptionsMwkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMwkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMwkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMxnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMxnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMxnTierParams.
func (p *PriceCurrencyOptionsMxnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MXN.
type PriceCurrencyOptionsMxnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMxnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMxnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMYRCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMYRTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMYRTierParams.
func (p *PriceCurrencyOptionsMYRTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MYR.
type PriceCurrencyOptionsMYRParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMYRCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMYRTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMznCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMznTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsMznTierParams.
func (p *PriceCurrencyOptionsMznTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MZN.
type PriceCurrencyOptionsMznParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMznCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMznTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNadCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNadTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNadTierParams.
func (p *PriceCurrencyOptionsNadTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NAD.
type PriceCurrencyOptionsNadParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNadCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNadTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNgnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNgnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNgnTierParams.
func (p *PriceCurrencyOptionsNgnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NGN.
type PriceCurrencyOptionsNgnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNgnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNgnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNioCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNioTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNioTierParams.
func (p *PriceCurrencyOptionsNioTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NIO.
type PriceCurrencyOptionsNioParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNioCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNioTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNOKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNOKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNOKTierParams.
func (p *PriceCurrencyOptionsNOKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NOK.
type PriceCurrencyOptionsNOKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNOKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNOKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNprCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNprTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNprTierParams.
func (p *PriceCurrencyOptionsNprTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NPR.
type PriceCurrencyOptionsNprParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNprCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNprTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNZDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNZDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsNZDTierParams.
func (p *PriceCurrencyOptionsNZDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NZD.
type PriceCurrencyOptionsNZDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNZDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNZDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsOmrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsOmrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsOmrTierParams.
func (p *PriceCurrencyOptionsOmrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in OMR.
type PriceCurrencyOptionsOmrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsOmrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsOmrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPabCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPabTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPabTierParams.
func (p *PriceCurrencyOptionsPabTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PAB.
type PriceCurrencyOptionsPabParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPabCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPabTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPenCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPenTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPenTierParams.
func (p *PriceCurrencyOptionsPenTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PEN.
type PriceCurrencyOptionsPenParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPenCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPenTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPgkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPgkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPgkTierParams.
func (p *PriceCurrencyOptionsPgkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PGK.
type PriceCurrencyOptionsPgkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPgkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPgkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPhpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPhpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPhpTierParams.
func (p *PriceCurrencyOptionsPhpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PHP.
type PriceCurrencyOptionsPhpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPhpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPhpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPkrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPkrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPkrTierParams.
func (p *PriceCurrencyOptionsPkrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PKR.
type PriceCurrencyOptionsPkrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPkrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPkrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPlnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPlnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPlnTierParams.
func (p *PriceCurrencyOptionsPlnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PLN.
type PriceCurrencyOptionsPlnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPlnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPlnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPygCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPygTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsPygTierParams.
func (p *PriceCurrencyOptionsPygTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PYG.
type PriceCurrencyOptionsPygParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPygCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPygTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsQarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsQarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsQarTierParams.
func (p *PriceCurrencyOptionsQarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in QAR.
type PriceCurrencyOptionsQarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsQarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsQarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRonCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRonTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsRonTierParams.
func (p *PriceCurrencyOptionsRonTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RON.
type PriceCurrencyOptionsRonParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRonCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRonTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRsdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRsdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsRsdTierParams.
func (p *PriceCurrencyOptionsRsdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RSD.
type PriceCurrencyOptionsRsdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRsdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRsdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRubCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRubTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsRubTierParams.
func (p *PriceCurrencyOptionsRubTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RUB.
type PriceCurrencyOptionsRubParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRubCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRubTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRwfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRwfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsRwfTierParams.
func (p *PriceCurrencyOptionsRwfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RWF.
type PriceCurrencyOptionsRwfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRwfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRwfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSarTierParams.
func (p *PriceCurrencyOptionsSarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SAR.
type PriceCurrencyOptionsSarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSbdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSbdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSbdTierParams.
func (p *PriceCurrencyOptionsSbdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SBD.
type PriceCurrencyOptionsSbdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSbdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSbdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsScrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsScrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsScrTierParams.
func (p *PriceCurrencyOptionsScrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SCR.
type PriceCurrencyOptionsScrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsScrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsScrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSEKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSEKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSEKTierParams.
func (p *PriceCurrencyOptionsSEKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SEK.
type PriceCurrencyOptionsSEKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSEKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSEKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSGDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSGDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSGDTierParams.
func (p *PriceCurrencyOptionsSGDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SGD.
type PriceCurrencyOptionsSGDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSGDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSGDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsShpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsShpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsShpTierParams.
func (p *PriceCurrencyOptionsShpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SHP.
type PriceCurrencyOptionsShpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsShpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsShpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSllCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSllTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSllTierParams.
func (p *PriceCurrencyOptionsSllTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SLL.
type PriceCurrencyOptionsSllParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSllCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSllTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSosCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSosTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSosTierParams.
func (p *PriceCurrencyOptionsSosTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SOS.
type PriceCurrencyOptionsSosParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSosCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSosTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSrdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSrdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSrdTierParams.
func (p *PriceCurrencyOptionsSrdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SRD.
type PriceCurrencyOptionsSrdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSrdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSrdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsStdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsStdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsStdTierParams.
func (p *PriceCurrencyOptionsStdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in STD.
type PriceCurrencyOptionsStdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsStdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsStdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSzlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSzlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsSzlTierParams.
func (p *PriceCurrencyOptionsSzlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SZL.
type PriceCurrencyOptionsSzlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSzlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSzlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsThbCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsThbTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsThbTierParams.
func (p *PriceCurrencyOptionsThbTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in THB.
type PriceCurrencyOptionsThbParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsThbCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsThbTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTjsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTjsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTjsTierParams.
func (p *PriceCurrencyOptionsTjsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TJS.
type PriceCurrencyOptionsTjsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTjsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTjsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTndTierParams.
func (p *PriceCurrencyOptionsTndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TND.
type PriceCurrencyOptionsTndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTopTierParams.
func (p *PriceCurrencyOptionsTopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TOP.
type PriceCurrencyOptionsTopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTryCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTryTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTryTierParams.
func (p *PriceCurrencyOptionsTryTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TRY.
type PriceCurrencyOptionsTryParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTryCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTryTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTtdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTtdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTtdTierParams.
func (p *PriceCurrencyOptionsTtdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TTD.
type PriceCurrencyOptionsTtdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTtdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTtdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTwdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTwdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTwdTierParams.
func (p *PriceCurrencyOptionsTwdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TWD.
type PriceCurrencyOptionsTwdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTwdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTwdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTzsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTzsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsTzsTierParams.
func (p *PriceCurrencyOptionsTzsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TZS.
type PriceCurrencyOptionsTzsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTzsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTzsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUahCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUahTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUahTierParams.
func (p *PriceCurrencyOptionsUahTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UAH.
type PriceCurrencyOptionsUahParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUahCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUahTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUgxCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUgxTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUgxTierParams.
func (p *PriceCurrencyOptionsUgxTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UGX.
type PriceCurrencyOptionsUgxParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUgxCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUgxTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUSDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUSDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUSDTierParams.
func (p *PriceCurrencyOptionsUSDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in USD.
type PriceCurrencyOptionsUSDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUSDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUSDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUsdcCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUsdcTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUsdcTierParams.
func (p *PriceCurrencyOptionsUsdcTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in USDC.
type PriceCurrencyOptionsUsdcParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUsdcCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUsdcTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUyuCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUyuTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUyuTierParams.
func (p *PriceCurrencyOptionsUyuTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UYU.
type PriceCurrencyOptionsUyuParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUyuCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUyuTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUzsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUzsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsUzsTierParams.
func (p *PriceCurrencyOptionsUzsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UZS.
type PriceCurrencyOptionsUzsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUzsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUzsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsVndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsVndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsVndTierParams.
func (p *PriceCurrencyOptionsVndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in VND.
type PriceCurrencyOptionsVndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsVndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsVndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsVuvCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsVuvTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsVuvTierParams.
func (p *PriceCurrencyOptionsVuvTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in VUV.
type PriceCurrencyOptionsVuvParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsVuvCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsVuvTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsWstCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsWstTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsWstTierParams.
func (p *PriceCurrencyOptionsWstTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in WST.
type PriceCurrencyOptionsWstParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsWstCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsWstTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXafCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXafTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsXafTierParams.
func (p *PriceCurrencyOptionsXafTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XAF.
type PriceCurrencyOptionsXafParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXafCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXafTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXcdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXcdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsXcdTierParams.
func (p *PriceCurrencyOptionsXcdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XCD.
type PriceCurrencyOptionsXcdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXcdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXcdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXofCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXofTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsXofTierParams.
func (p *PriceCurrencyOptionsXofTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XOF.
type PriceCurrencyOptionsXofParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXofCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXofTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXpfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXpfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsXpfTierParams.
func (p *PriceCurrencyOptionsXpfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XPF.
type PriceCurrencyOptionsXpfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXpfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXpfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsYerCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsYerTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsYerTierParams.
func (p *PriceCurrencyOptionsYerTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in YER.
type PriceCurrencyOptionsYerParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsYerCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsYerTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsZarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsZarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsZarTierParams.
func (p *PriceCurrencyOptionsZarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ZAR.
type PriceCurrencyOptionsZarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsZarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsZarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsZmwCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsZmwTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceCurrencyOptionsZmwTierParams.
func (p *PriceCurrencyOptionsZmwTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ZMW.
type PriceCurrencyOptionsZmwParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsZmwCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsZmwTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type PriceCurrencyOptionsParams struct {
	// The price defined in AED.
	Aed *PriceCurrencyOptionsAedParams `form:"aed"`
	// The price defined in AFN.
	Afn *PriceCurrencyOptionsAfnParams `form:"afn"`
	// The price defined in ALL.
	All *PriceCurrencyOptionsAllParams `form:"all"`
	// The price defined in AMD.
	Amd *PriceCurrencyOptionsAmdParams `form:"amd"`
	// The price defined in ANG.
	Ang *PriceCurrencyOptionsAngParams `form:"ang"`
	// The price defined in AOA.
	Aoa *PriceCurrencyOptionsAoaParams `form:"aoa"`
	// The price defined in ARS.
	Ars *PriceCurrencyOptionsArsParams `form:"ars"`
	// The price defined in AUD.
	AUD *PriceCurrencyOptionsAUDParams `form:"aud"`
	// The price defined in AWG.
	Awg *PriceCurrencyOptionsAwgParams `form:"awg"`
	// The price defined in AZN.
	Azn *PriceCurrencyOptionsAznParams `form:"azn"`
	// The price defined in BAM.
	Bam *PriceCurrencyOptionsBamParams `form:"bam"`
	// The price defined in BBD.
	Bbd *PriceCurrencyOptionsBbdParams `form:"bbd"`
	// The price defined in BDT.
	Bdt *PriceCurrencyOptionsBdtParams `form:"bdt"`
	// The price defined in BGN.
	Bgn *PriceCurrencyOptionsBgnParams `form:"bgn"`
	// The price defined in BHD.
	Bhd *PriceCurrencyOptionsBhdParams `form:"bhd"`
	// The price defined in BIF.
	Bif *PriceCurrencyOptionsBifParams `form:"bif"`
	// The price defined in BMD.
	Bmd *PriceCurrencyOptionsBmdParams `form:"bmd"`
	// The price defined in BND.
	Bnd *PriceCurrencyOptionsBndParams `form:"bnd"`
	// The price defined in BOB.
	Bob *PriceCurrencyOptionsBobParams `form:"bob"`
	// The price defined in BRL.
	Brl *PriceCurrencyOptionsBrlParams `form:"brl"`
	// The price defined in BSD.
	Bsd *PriceCurrencyOptionsBsdParams `form:"bsd"`
	// The price defined in BTN.
	Btn *PriceCurrencyOptionsBtnParams `form:"btn"`
	// The price defined in BWP.
	Bwp *PriceCurrencyOptionsBwpParams `form:"bwp"`
	// The price defined in BYN.
	Byn *PriceCurrencyOptionsBynParams `form:"byn"`
	// The price defined in BZD.
	Bzd *PriceCurrencyOptionsBzdParams `form:"bzd"`
	// The price defined in CAD.
	CAD *PriceCurrencyOptionsCADParams `form:"cad"`
	// The price defined in CDF.
	Cdf *PriceCurrencyOptionsCdfParams `form:"cdf"`
	// The price defined in CHF.
	CHF *PriceCurrencyOptionsCHFParams `form:"chf"`
	// The price defined in CLP.
	Clp *PriceCurrencyOptionsClpParams `form:"clp"`
	// The price defined in CNY.
	Cny *PriceCurrencyOptionsCnyParams `form:"cny"`
	// The price defined in COP.
	Cop *PriceCurrencyOptionsCopParams `form:"cop"`
	// The price defined in CRC.
	Crc *PriceCurrencyOptionsCrcParams `form:"crc"`
	// The price defined in CVE.
	Cve *PriceCurrencyOptionsCveParams `form:"cve"`
	// The price defined in CZK.
	CZK *PriceCurrencyOptionsCZKParams `form:"czk"`
	// The price defined in DJF.
	Djf *PriceCurrencyOptionsDjfParams `form:"djf"`
	// The price defined in DKK.
	DKK *PriceCurrencyOptionsDKKParams `form:"dkk"`
	// The price defined in DOP.
	Dop *PriceCurrencyOptionsDopParams `form:"dop"`
	// The price defined in DZD.
	Dzd *PriceCurrencyOptionsDzdParams `form:"dzd"`
	// The price defined in EGP.
	Egp *PriceCurrencyOptionsEgpParams `form:"egp"`
	// The price defined in ETB.
	Etb *PriceCurrencyOptionsEtbParams `form:"etb"`
	// The price defined in EUR.
	EUR *PriceCurrencyOptionsEURParams `form:"eur"`
	// The price defined in FJD.
	Fjd *PriceCurrencyOptionsFjdParams `form:"fjd"`
	// The price defined in FKP.
	Fkp *PriceCurrencyOptionsFkpParams `form:"fkp"`
	// The price defined in GBP.
	GBP *PriceCurrencyOptionsGBPParams `form:"gbp"`
	// The price defined in GEL.
	Gel *PriceCurrencyOptionsGelParams `form:"gel"`
	// The price defined in GHS.
	Ghs *PriceCurrencyOptionsGhsParams `form:"ghs"`
	// The price defined in GIP.
	Gip *PriceCurrencyOptionsGipParams `form:"gip"`
	// The price defined in GMD.
	Gmd *PriceCurrencyOptionsGmdParams `form:"gmd"`
	// The price defined in GNF.
	Gnf *PriceCurrencyOptionsGnfParams `form:"gnf"`
	// The price defined in GTQ.
	Gtq *PriceCurrencyOptionsGtqParams `form:"gtq"`
	// The price defined in GYD.
	Gyd *PriceCurrencyOptionsGydParams `form:"gyd"`
	// The price defined in HKD.
	HKD *PriceCurrencyOptionsHKDParams `form:"hkd"`
	// The price defined in HNL.
	Hnl *PriceCurrencyOptionsHnlParams `form:"hnl"`
	// The price defined in HRK.
	Hrk *PriceCurrencyOptionsHrkParams `form:"hrk"`
	// The price defined in HTG.
	Htg *PriceCurrencyOptionsHtgParams `form:"htg"`
	// The price defined in HUF.
	Huf *PriceCurrencyOptionsHufParams `form:"huf"`
	// The price defined in IDR.
	Idr *PriceCurrencyOptionsIdrParams `form:"idr"`
	// The price defined in ILS.
	Ils *PriceCurrencyOptionsIlsParams `form:"ils"`
	// The price defined in INR.
	Inr *PriceCurrencyOptionsInrParams `form:"inr"`
	// The price defined in ISK.
	Isk *PriceCurrencyOptionsIskParams `form:"isk"`
	// The price defined in JMD.
	Jmd *PriceCurrencyOptionsJmdParams `form:"jmd"`
	// The price defined in JOD.
	Jod *PriceCurrencyOptionsJodParams `form:"jod"`
	// The price defined in JPY.
	Jpy *PriceCurrencyOptionsJpyParams `form:"jpy"`
	// The price defined in KES.
	Kes *PriceCurrencyOptionsKesParams `form:"kes"`
	// The price defined in KGS.
	Kgs *PriceCurrencyOptionsKgsParams `form:"kgs"`
	// The price defined in KHR.
	Khr *PriceCurrencyOptionsKhrParams `form:"khr"`
	// The price defined in KMF.
	Kmf *PriceCurrencyOptionsKmfParams `form:"kmf"`
	// The price defined in KRW.
	Krw *PriceCurrencyOptionsKrwParams `form:"krw"`
	// The price defined in KWD.
	Kwd *PriceCurrencyOptionsKwdParams `form:"kwd"`
	// The price defined in KYD.
	Kyd *PriceCurrencyOptionsKydParams `form:"kyd"`
	// The price defined in KZT.
	Kzt *PriceCurrencyOptionsKztParams `form:"kzt"`
	// The price defined in LAK.
	Lak *PriceCurrencyOptionsLakParams `form:"lak"`
	// The price defined in LBP.
	Lbp *PriceCurrencyOptionsLbpParams `form:"lbp"`
	// The price defined in LKR.
	Lkr *PriceCurrencyOptionsLkrParams `form:"lkr"`
	// The price defined in LRD.
	Lrd *PriceCurrencyOptionsLrdParams `form:"lrd"`
	// The price defined in LSL.
	Lsl *PriceCurrencyOptionsLslParams `form:"lsl"`
	// The price defined in MAD.
	Mad *PriceCurrencyOptionsMadParams `form:"mad"`
	// The price defined in MDL.
	Mdl *PriceCurrencyOptionsMdlParams `form:"mdl"`
	// The price defined in MGA.
	Mga *PriceCurrencyOptionsMgaParams `form:"mga"`
	// The price defined in MKD.
	Mkd *PriceCurrencyOptionsMkdParams `form:"mkd"`
	// The price defined in MMK.
	Mmk *PriceCurrencyOptionsMmkParams `form:"mmk"`
	// The price defined in MNT.
	Mnt *PriceCurrencyOptionsMntParams `form:"mnt"`
	// The price defined in MOP.
	Mop *PriceCurrencyOptionsMopParams `form:"mop"`
	// The price defined in MRO.
	Mro *PriceCurrencyOptionsMroParams `form:"mro"`
	// The price defined in MUR.
	Mur *PriceCurrencyOptionsMurParams `form:"mur"`
	// The price defined in MVR.
	Mvr *PriceCurrencyOptionsMvrParams `form:"mvr"`
	// The price defined in MWK.
	Mwk *PriceCurrencyOptionsMwkParams `form:"mwk"`
	// The price defined in MXN.
	Mxn *PriceCurrencyOptionsMxnParams `form:"mxn"`
	// The price defined in MYR.
	MYR *PriceCurrencyOptionsMYRParams `form:"myr"`
	// The price defined in MZN.
	Mzn *PriceCurrencyOptionsMznParams `form:"mzn"`
	// The price defined in NAD.
	Nad *PriceCurrencyOptionsNadParams `form:"nad"`
	// The price defined in NGN.
	Ngn *PriceCurrencyOptionsNgnParams `form:"ngn"`
	// The price defined in NIO.
	Nio *PriceCurrencyOptionsNioParams `form:"nio"`
	// The price defined in NOK.
	NOK *PriceCurrencyOptionsNOKParams `form:"nok"`
	// The price defined in NPR.
	Npr *PriceCurrencyOptionsNprParams `form:"npr"`
	// The price defined in NZD.
	NZD *PriceCurrencyOptionsNZDParams `form:"nzd"`
	// The price defined in OMR.
	Omr *PriceCurrencyOptionsOmrParams `form:"omr"`
	// The price defined in PAB.
	Pab *PriceCurrencyOptionsPabParams `form:"pab"`
	// The price defined in PEN.
	Pen *PriceCurrencyOptionsPenParams `form:"pen"`
	// The price defined in PGK.
	Pgk *PriceCurrencyOptionsPgkParams `form:"pgk"`
	// The price defined in PHP.
	Php *PriceCurrencyOptionsPhpParams `form:"php"`
	// The price defined in PKR.
	Pkr *PriceCurrencyOptionsPkrParams `form:"pkr"`
	// The price defined in PLN.
	Pln *PriceCurrencyOptionsPlnParams `form:"pln"`
	// The price defined in PYG.
	Pyg *PriceCurrencyOptionsPygParams `form:"pyg"`
	// The price defined in QAR.
	Qar *PriceCurrencyOptionsQarParams `form:"qar"`
	// The price defined in RON.
	Ron *PriceCurrencyOptionsRonParams `form:"ron"`
	// The price defined in RSD.
	Rsd *PriceCurrencyOptionsRsdParams `form:"rsd"`
	// The price defined in RUB.
	Rub *PriceCurrencyOptionsRubParams `form:"rub"`
	// The price defined in RWF.
	Rwf *PriceCurrencyOptionsRwfParams `form:"rwf"`
	// The price defined in SAR.
	Sar *PriceCurrencyOptionsSarParams `form:"sar"`
	// The price defined in SBD.
	Sbd *PriceCurrencyOptionsSbdParams `form:"sbd"`
	// The price defined in SCR.
	Scr *PriceCurrencyOptionsScrParams `form:"scr"`
	// The price defined in SEK.
	SEK *PriceCurrencyOptionsSEKParams `form:"sek"`
	// The price defined in SGD.
	SGD *PriceCurrencyOptionsSGDParams `form:"sgd"`
	// The price defined in SHP.
	Shp *PriceCurrencyOptionsShpParams `form:"shp"`
	// The price defined in SLL.
	Sll *PriceCurrencyOptionsSllParams `form:"sll"`
	// The price defined in SOS.
	Sos *PriceCurrencyOptionsSosParams `form:"sos"`
	// The price defined in SRD.
	Srd *PriceCurrencyOptionsSrdParams `form:"srd"`
	// The price defined in STD.
	Std *PriceCurrencyOptionsStdParams `form:"std"`
	// The price defined in SZL.
	Szl *PriceCurrencyOptionsSzlParams `form:"szl"`
	// The price defined in THB.
	Thb *PriceCurrencyOptionsThbParams `form:"thb"`
	// The price defined in TJS.
	Tjs *PriceCurrencyOptionsTjsParams `form:"tjs"`
	// The price defined in TND.
	Tnd *PriceCurrencyOptionsTndParams `form:"tnd"`
	// The price defined in TOP.
	Top *PriceCurrencyOptionsTopParams `form:"top"`
	// The price defined in TRY.
	Try *PriceCurrencyOptionsTryParams `form:"try"`
	// The price defined in TTD.
	Ttd *PriceCurrencyOptionsTtdParams `form:"ttd"`
	// The price defined in TWD.
	Twd *PriceCurrencyOptionsTwdParams `form:"twd"`
	// The price defined in TZS.
	Tzs *PriceCurrencyOptionsTzsParams `form:"tzs"`
	// The price defined in UAH.
	Uah *PriceCurrencyOptionsUahParams `form:"uah"`
	// The price defined in UGX.
	Ugx *PriceCurrencyOptionsUgxParams `form:"ugx"`
	// The price defined in USD.
	USD *PriceCurrencyOptionsUSDParams `form:"usd"`
	// The price defined in USDC.
	Usdc *PriceCurrencyOptionsUsdcParams `form:"usdc"`
	// The price defined in UYU.
	Uyu *PriceCurrencyOptionsUyuParams `form:"uyu"`
	// The price defined in UZS.
	Uzs *PriceCurrencyOptionsUzsParams `form:"uzs"`
	// The price defined in VND.
	Vnd *PriceCurrencyOptionsVndParams `form:"vnd"`
	// The price defined in VUV.
	Vuv *PriceCurrencyOptionsVuvParams `form:"vuv"`
	// The price defined in WST.
	Wst *PriceCurrencyOptionsWstParams `form:"wst"`
	// The price defined in XAF.
	Xaf *PriceCurrencyOptionsXafParams `form:"xaf"`
	// The price defined in XCD.
	Xcd *PriceCurrencyOptionsXcdParams `form:"xcd"`
	// The price defined in XOF.
	Xof *PriceCurrencyOptionsXofParams `form:"xof"`
	// The price defined in XPF.
	Xpf *PriceCurrencyOptionsXpfParams `form:"xpf"`
	// The price defined in YER.
	Yer *PriceCurrencyOptionsYerParams `form:"yer"`
	// The price defined in ZAR.
	Zar *PriceCurrencyOptionsZarParams `form:"zar"`
	// The price defined in ZMW.
	Zmw *PriceCurrencyOptionsZmwParams `form:"zmw"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// These fields can be used to create a new product that this price will belong to.
type PriceProductDataParams struct {
	// Whether the product is currently available for purchase. Defaults to `true`.
	Active *bool `form:"active"`
	// The identifier for the product. Must be unique. If not provided, an identifier will be randomly generated.
	ID *string `form:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.
	//
	// This may be up to 22 characters. The statement description may not include `<`, `>`, `\`, `"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.
	StatementDescriptor *string `form:"statement_descriptor"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A label that represents units of this product in Stripe and on customers' receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel *string `form:"unit_label"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type PriceRecurringParams struct {
	// Specifies a usage aggregation strategy for prices of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
	AggregateUsage *string `form:"aggregate_usage"`
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
	// Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).
	TrialPeriodDays *int64 `form:"trial_period_days"`
	// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
	UsageType *string `form:"usage_type"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceTierParams.
func (p *PriceTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.
type PriceTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by"`
	// After division, either round the result `up` or `down`.
	Round *string `form:"round"`
}

// Creates a new price for an existing product. The price can be recurring or one-time.
type PriceParams struct {
	Params `form:"*"`
	// Whether the price can be used for new purchases. Defaults to `true`.
	Active *bool `form:"active"`
	// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
	BillingScheme *string `form:"billing_scheme"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *PriceCurrencyOptionsParams `form:"currency_options"`
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCustomUnitAmountParams `form:"custom_unit_amount"`
	// A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.
	LookupKey *string `form:"lookup_key"`
	// A brief description of the price, hidden from customers.
	Nickname *string `form:"nickname"`
	// The ID of the product that this price will belong to.
	Product *string `form:"product"`
	// These fields can be used to create a new product that this price will belong to.
	ProductData *PriceProductDataParams `form:"product_data"`
	// The recurring components of a price such as `interval` and `usage_type`.
	Recurring *PriceRecurringParams `form:"recurring"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceTierParams `form:"tiers"`
	// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price, in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode *string `form:"tiers_mode"`
	// If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.
	TransferLookupKey *bool `form:"transfer_lookup_key"`
	// Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.
	TransformQuantity *PriceTransformQuantityParams `form:"transform_quantity"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAedCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAedTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAed struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAedCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAedTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAedTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAfnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAfnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAfn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAfnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAfnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAfnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAllCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAllTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAll struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAllCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAllTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAllTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAmdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAmdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAmd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAmdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAmdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAmdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAngCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAngTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAng struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAngCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAngTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAngTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAoaCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAoaTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAoa struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAoaCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAoaTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAoaTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsArsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsArsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsArs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsArsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsArsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsArsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAUDCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAUDTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAUD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAUDCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAUDTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAUDTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAwgCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAwgTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAwg struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAwgCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAwgTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAwgTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsAznCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsAznTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsAzn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsAznCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsAznTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsAznTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBamCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBamTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBam struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBamCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBamTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBamTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBbdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBbdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBbd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBbdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBbdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBbdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBdtCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBdtTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBdt struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBdtCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBdtTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBdtTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBgnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBgnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBgn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBgnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBgnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBgnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBhdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBhdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBhd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBhdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBhdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBhdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBifCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBifTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBif struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBifCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBifTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBifTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBmdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBmdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBmd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBmdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBmdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBmdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBndCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBndTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBnd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBndCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBndTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBndTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBobCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBobTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBob struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBobCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBobTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBobTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBrlCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBrlTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBrl struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBrlCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBrlTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBrlTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBsdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBsdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBsd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBsdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBsdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBsdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBtnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBtnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBtn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBtnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBtnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBtnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBwpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBwpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBwp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBwpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBwpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBwpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBynCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBynTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsByn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBynCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBynTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBynTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsBzdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsBzdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsBzd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsBzdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsBzdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsBzdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCADCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCADTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCAD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCADCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCADTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCADTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCdfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCdfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCdf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCdfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCdfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCdfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCHFCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCHFTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCHF struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCHFCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCHFTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCHFTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsClpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsClpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsClp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsClpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsClpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsClpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCnyCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCnyTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCny struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCnyCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCnyTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCnyTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCopCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCopTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCop struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCopCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCopTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCopTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCrcCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCrcTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCrc struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCrcCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCrcTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCrcTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCveCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCveTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCve struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCveCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCveTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCveTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsCZKCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsCZKTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsCZK struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsCZKCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsCZKTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsCZKTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDjfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDjfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsDjf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDjfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsDjfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDjfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDKKCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDKKTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsDKK struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDKKCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsDKKTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDKKTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDopCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDopTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsDop struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDopCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsDopTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDopTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsDzdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsDzdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsDzd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsDzdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsDzdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsDzdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEgpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEgpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsEgp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEgpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsEgpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEgpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEtbCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEtbTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsEtb struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEtbCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsEtbTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEtbTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsEURCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsEURTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsEUR struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsEURCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsEURTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsEURTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsFjdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsFjdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsFjd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsFjdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsFjdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsFjdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsFkpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsFkpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsFkp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsFkpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsFkpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsFkpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGBPCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGBPTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGBP struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGBPCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGBPTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGBPTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGelCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGelTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGel struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGelCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGelTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGelTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGhsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGhsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGhs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGhsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGhsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGhsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGipCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGipTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGip struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGipCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGipTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGipTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGmdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGmdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGmd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGmdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGmdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGmdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGnfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGnfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGnf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGnfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGnfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGnfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGtqCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGtqTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGtq struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGtqCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGtqTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGtqTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsGydCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsGydTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsGyd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsGydCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsGydTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsGydTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHKDCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHKDTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsHKD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHKDCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsHKDTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHKDTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHnlCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHnlTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsHnl struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHnlCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsHnlTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHnlTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHrkCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHrkTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsHrk struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHrkCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsHrkTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHrkTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHtgCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHtgTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsHtg struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHtgCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsHtgTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHtgTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsHufCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsHufTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsHuf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsHufCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsHufTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsHufTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIdrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIdrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsIdr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIdrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsIdrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIdrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIlsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIlsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsIls struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIlsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsIlsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIlsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsInrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsInrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsInr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsInrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsInrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsInrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsIskCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsIskTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsIsk struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsIskCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsIskTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsIskTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJmdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJmdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsJmd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJmdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsJmdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJmdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJodCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJodTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsJod struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJodCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsJodTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJodTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsJpyCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsJpyTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsJpy struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsJpyCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsJpyTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsJpyTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKesCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKesTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKes struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKesCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKesTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKesTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKgsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKgsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKgs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKgsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKgsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKgsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKhrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKhrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKhr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKhrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKhrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKhrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKmfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKmfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKmf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKmfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKmfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKmfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKrwCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKrwTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKrw struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKrwCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKrwTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKrwTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKwdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKwdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKwd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKwdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKwdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKwdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKydCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKydTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKyd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKydCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKydTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKydTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsKztCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsKztTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsKzt struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsKztCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsKztTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsKztTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLakCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLakTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsLak struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLakCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsLakTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLakTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLbpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLbpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsLbp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLbpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsLbpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLbpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLkrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLkrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsLkr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLkrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsLkrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLkrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLrdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLrdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsLrd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLrdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsLrdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLrdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsLslCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsLslTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsLsl struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsLslCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsLslTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsLslTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMadCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMadTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMad struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMadCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMadTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMadTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMdlCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMdlTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMdl struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMdlCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMdlTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMdlTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMgaCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMgaTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMga struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMgaCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMgaTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMgaTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMkdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMkdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMkd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMkdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMkdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMkdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMmkCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMmkTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMmk struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMmkCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMmkTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMmkTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMntCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMntTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMnt struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMntCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMntTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMntTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMopCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMopTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMop struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMopCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMopTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMopTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMroCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMroTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMro struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMroCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMroTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMroTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMurCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMurTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMur struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMurCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMurTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMurTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMvrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMvrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMvr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMvrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMvrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMvrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMwkCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMwkTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMwk struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMwkCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMwkTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMwkTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMxnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMxnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMxn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMxnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMxnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMxnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMYRCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMYRTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMYR struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMYRCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMYRTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMYRTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsMznCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsMznTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsMzn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsMznCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsMznTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsMznTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNadCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNadTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNad struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNadCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNadTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNadTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNgnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNgnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNgn struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNgnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNgnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNgnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNioCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNioTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNio struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNioCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNioTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNioTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNOKCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNOKTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNOK struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNOKCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNOKTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNOKTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNprCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNprTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNpr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNprCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNprTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNprTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsNZDCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsNZDTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsNZD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsNZDCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsNZDTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsNZDTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsOmrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsOmrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsOmr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsOmrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsOmrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsOmrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPabCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPabTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPab struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPabCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPabTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPabTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPenCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPenTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPen struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPenCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPenTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPenTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPgkCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPgkTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPgk struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPgkCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPgkTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPgkTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPhpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPhpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPhp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPhpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPhpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPhpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPkrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPkrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPkr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPkrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPkrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPkrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPlnCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPlnTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPln struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPlnCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPlnTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPlnTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsPygCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsPygTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsPyg struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsPygCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsPygTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsPygTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsQarCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsQarTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsQar struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsQarCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsQarTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsQarTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRonCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRonTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsRon struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRonCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsRonTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRonTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRsdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRsdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsRsd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRsdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsRsdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRsdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRubCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRubTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsRub struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRubCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsRubTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRubTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsRwfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsRwfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsRwf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsRwfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsRwfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsRwfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSarCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSarTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSar struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSarCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSarTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSarTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSbdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSbdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSbd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSbdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSbdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSbdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsScrCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsScrTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsScr struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsScrCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsScrTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsScrTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSEKCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSEKTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSEK struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSEKCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSEKTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSEKTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSGDCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSGDTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSGD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSGDCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSGDTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSGDTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsShpCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsShpTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsShp struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsShpCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsShpTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsShpTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSllCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSllTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSll struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSllCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSllTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSllTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSosCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSosTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSos struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSosCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSosTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSosTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSrdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSrdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSrd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSrdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSrdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSrdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsStdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsStdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsStd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsStdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsStdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsStdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsSzlCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsSzlTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsSzl struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsSzlCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsSzlTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsSzlTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsThbCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsThbTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsThb struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsThbCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsThbTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsThbTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTjsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTjsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTjs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTjsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTjsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTjsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTndCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTndTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTnd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTndCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTndTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTndTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTopCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTopTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTop struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTopCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTopTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTopTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTryCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTryTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTry struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTryCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTryTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTryTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTtdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTtdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTtd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTtdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTtdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTtdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTwdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTwdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTwd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTwdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTwdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTwdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsTzsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsTzsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsTzs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsTzsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsTzsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsTzsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUahCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUahTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUah struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUahCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUahTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUahTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUgxCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUgxTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUgx struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUgxCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUgxTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUgxTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUSDCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUSDTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUSD struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUSDCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUSDTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUSDTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUsdcCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUsdcTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUsdc struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUsdcCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUsdcTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUsdcTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUyuCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUyuTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUyu struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUyuCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUyuTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUyuTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsUzsCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsUzsTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsUzs struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsUzsCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsUzsTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsUzsTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsVndCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsVndTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsVnd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsVndCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsVndTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsVndTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsVuvCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsVuvTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsVuv struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsVuvCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsVuvTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsVuvTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsWstCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsWstTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsWst struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsWstCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsWstTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsWstTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXafCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXafTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsXaf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXafCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsXafTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXafTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXcdCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXcdTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsXcd struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXcdCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsXcdTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXcdTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXofCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXofTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsXof struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXofCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsXofTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXofTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsXpfCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsXpfTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsXpf struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsXpfCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsXpfTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsXpfTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsYerCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsYerTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsYer struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsYerCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsYerTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsYerTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsZarCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsZarTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsZar struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsZarCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsZarTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsZarTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCurrencyOptionsZmwCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceCurrencyOptionsZmwTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}
type PriceCurrencyOptionsZmw struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCurrencyOptionsZmwCustomUnitAmount `json:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceCurrencyOptionsZmwTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceCurrencyOptionsZmwTier `json:"tiers"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}
type PriceCurrencyOptions struct {
	Aed  *PriceCurrencyOptionsAed  `json:"aed"`
	Afn  *PriceCurrencyOptionsAfn  `json:"afn"`
	All  *PriceCurrencyOptionsAll  `json:"all"`
	Amd  *PriceCurrencyOptionsAmd  `json:"amd"`
	Ang  *PriceCurrencyOptionsAng  `json:"ang"`
	Aoa  *PriceCurrencyOptionsAoa  `json:"aoa"`
	Ars  *PriceCurrencyOptionsArs  `json:"ars"`
	AUD  *PriceCurrencyOptionsAUD  `json:"aud"`
	Awg  *PriceCurrencyOptionsAwg  `json:"awg"`
	Azn  *PriceCurrencyOptionsAzn  `json:"azn"`
	Bam  *PriceCurrencyOptionsBam  `json:"bam"`
	Bbd  *PriceCurrencyOptionsBbd  `json:"bbd"`
	Bdt  *PriceCurrencyOptionsBdt  `json:"bdt"`
	Bgn  *PriceCurrencyOptionsBgn  `json:"bgn"`
	Bhd  *PriceCurrencyOptionsBhd  `json:"bhd"`
	Bif  *PriceCurrencyOptionsBif  `json:"bif"`
	Bmd  *PriceCurrencyOptionsBmd  `json:"bmd"`
	Bnd  *PriceCurrencyOptionsBnd  `json:"bnd"`
	Bob  *PriceCurrencyOptionsBob  `json:"bob"`
	Brl  *PriceCurrencyOptionsBrl  `json:"brl"`
	Bsd  *PriceCurrencyOptionsBsd  `json:"bsd"`
	Btn  *PriceCurrencyOptionsBtn  `json:"btn"`
	Bwp  *PriceCurrencyOptionsBwp  `json:"bwp"`
	Byn  *PriceCurrencyOptionsByn  `json:"byn"`
	Bzd  *PriceCurrencyOptionsBzd  `json:"bzd"`
	CAD  *PriceCurrencyOptionsCAD  `json:"cad"`
	Cdf  *PriceCurrencyOptionsCdf  `json:"cdf"`
	CHF  *PriceCurrencyOptionsCHF  `json:"chf"`
	Clp  *PriceCurrencyOptionsClp  `json:"clp"`
	Cny  *PriceCurrencyOptionsCny  `json:"cny"`
	Cop  *PriceCurrencyOptionsCop  `json:"cop"`
	Crc  *PriceCurrencyOptionsCrc  `json:"crc"`
	Cve  *PriceCurrencyOptionsCve  `json:"cve"`
	CZK  *PriceCurrencyOptionsCZK  `json:"czk"`
	Djf  *PriceCurrencyOptionsDjf  `json:"djf"`
	DKK  *PriceCurrencyOptionsDKK  `json:"dkk"`
	Dop  *PriceCurrencyOptionsDop  `json:"dop"`
	Dzd  *PriceCurrencyOptionsDzd  `json:"dzd"`
	Egp  *PriceCurrencyOptionsEgp  `json:"egp"`
	Etb  *PriceCurrencyOptionsEtb  `json:"etb"`
	EUR  *PriceCurrencyOptionsEUR  `json:"eur"`
	Fjd  *PriceCurrencyOptionsFjd  `json:"fjd"`
	Fkp  *PriceCurrencyOptionsFkp  `json:"fkp"`
	GBP  *PriceCurrencyOptionsGBP  `json:"gbp"`
	Gel  *PriceCurrencyOptionsGel  `json:"gel"`
	Ghs  *PriceCurrencyOptionsGhs  `json:"ghs"`
	Gip  *PriceCurrencyOptionsGip  `json:"gip"`
	Gmd  *PriceCurrencyOptionsGmd  `json:"gmd"`
	Gnf  *PriceCurrencyOptionsGnf  `json:"gnf"`
	Gtq  *PriceCurrencyOptionsGtq  `json:"gtq"`
	Gyd  *PriceCurrencyOptionsGyd  `json:"gyd"`
	HKD  *PriceCurrencyOptionsHKD  `json:"hkd"`
	Hnl  *PriceCurrencyOptionsHnl  `json:"hnl"`
	Hrk  *PriceCurrencyOptionsHrk  `json:"hrk"`
	Htg  *PriceCurrencyOptionsHtg  `json:"htg"`
	Huf  *PriceCurrencyOptionsHuf  `json:"huf"`
	Idr  *PriceCurrencyOptionsIdr  `json:"idr"`
	Ils  *PriceCurrencyOptionsIls  `json:"ils"`
	Inr  *PriceCurrencyOptionsInr  `json:"inr"`
	Isk  *PriceCurrencyOptionsIsk  `json:"isk"`
	Jmd  *PriceCurrencyOptionsJmd  `json:"jmd"`
	Jod  *PriceCurrencyOptionsJod  `json:"jod"`
	Jpy  *PriceCurrencyOptionsJpy  `json:"jpy"`
	Kes  *PriceCurrencyOptionsKes  `json:"kes"`
	Kgs  *PriceCurrencyOptionsKgs  `json:"kgs"`
	Khr  *PriceCurrencyOptionsKhr  `json:"khr"`
	Kmf  *PriceCurrencyOptionsKmf  `json:"kmf"`
	Krw  *PriceCurrencyOptionsKrw  `json:"krw"`
	Kwd  *PriceCurrencyOptionsKwd  `json:"kwd"`
	Kyd  *PriceCurrencyOptionsKyd  `json:"kyd"`
	Kzt  *PriceCurrencyOptionsKzt  `json:"kzt"`
	Lak  *PriceCurrencyOptionsLak  `json:"lak"`
	Lbp  *PriceCurrencyOptionsLbp  `json:"lbp"`
	Lkr  *PriceCurrencyOptionsLkr  `json:"lkr"`
	Lrd  *PriceCurrencyOptionsLrd  `json:"lrd"`
	Lsl  *PriceCurrencyOptionsLsl  `json:"lsl"`
	Mad  *PriceCurrencyOptionsMad  `json:"mad"`
	Mdl  *PriceCurrencyOptionsMdl  `json:"mdl"`
	Mga  *PriceCurrencyOptionsMga  `json:"mga"`
	Mkd  *PriceCurrencyOptionsMkd  `json:"mkd"`
	Mmk  *PriceCurrencyOptionsMmk  `json:"mmk"`
	Mnt  *PriceCurrencyOptionsMnt  `json:"mnt"`
	Mop  *PriceCurrencyOptionsMop  `json:"mop"`
	Mro  *PriceCurrencyOptionsMro  `json:"mro"`
	Mur  *PriceCurrencyOptionsMur  `json:"mur"`
	Mvr  *PriceCurrencyOptionsMvr  `json:"mvr"`
	Mwk  *PriceCurrencyOptionsMwk  `json:"mwk"`
	Mxn  *PriceCurrencyOptionsMxn  `json:"mxn"`
	MYR  *PriceCurrencyOptionsMYR  `json:"myr"`
	Mzn  *PriceCurrencyOptionsMzn  `json:"mzn"`
	Nad  *PriceCurrencyOptionsNad  `json:"nad"`
	Ngn  *PriceCurrencyOptionsNgn  `json:"ngn"`
	Nio  *PriceCurrencyOptionsNio  `json:"nio"`
	NOK  *PriceCurrencyOptionsNOK  `json:"nok"`
	Npr  *PriceCurrencyOptionsNpr  `json:"npr"`
	NZD  *PriceCurrencyOptionsNZD  `json:"nzd"`
	Omr  *PriceCurrencyOptionsOmr  `json:"omr"`
	Pab  *PriceCurrencyOptionsPab  `json:"pab"`
	Pen  *PriceCurrencyOptionsPen  `json:"pen"`
	Pgk  *PriceCurrencyOptionsPgk  `json:"pgk"`
	Php  *PriceCurrencyOptionsPhp  `json:"php"`
	Pkr  *PriceCurrencyOptionsPkr  `json:"pkr"`
	Pln  *PriceCurrencyOptionsPln  `json:"pln"`
	Pyg  *PriceCurrencyOptionsPyg  `json:"pyg"`
	Qar  *PriceCurrencyOptionsQar  `json:"qar"`
	Ron  *PriceCurrencyOptionsRon  `json:"ron"`
	Rsd  *PriceCurrencyOptionsRsd  `json:"rsd"`
	Rub  *PriceCurrencyOptionsRub  `json:"rub"`
	Rwf  *PriceCurrencyOptionsRwf  `json:"rwf"`
	Sar  *PriceCurrencyOptionsSar  `json:"sar"`
	Sbd  *PriceCurrencyOptionsSbd  `json:"sbd"`
	Scr  *PriceCurrencyOptionsScr  `json:"scr"`
	SEK  *PriceCurrencyOptionsSEK  `json:"sek"`
	SGD  *PriceCurrencyOptionsSGD  `json:"sgd"`
	Shp  *PriceCurrencyOptionsShp  `json:"shp"`
	Sll  *PriceCurrencyOptionsSll  `json:"sll"`
	Sos  *PriceCurrencyOptionsSos  `json:"sos"`
	Srd  *PriceCurrencyOptionsSrd  `json:"srd"`
	Std  *PriceCurrencyOptionsStd  `json:"std"`
	Szl  *PriceCurrencyOptionsSzl  `json:"szl"`
	Thb  *PriceCurrencyOptionsThb  `json:"thb"`
	Tjs  *PriceCurrencyOptionsTjs  `json:"tjs"`
	Tnd  *PriceCurrencyOptionsTnd  `json:"tnd"`
	Top  *PriceCurrencyOptionsTop  `json:"top"`
	Try  *PriceCurrencyOptionsTry  `json:"try"`
	Ttd  *PriceCurrencyOptionsTtd  `json:"ttd"`
	Twd  *PriceCurrencyOptionsTwd  `json:"twd"`
	Tzs  *PriceCurrencyOptionsTzs  `json:"tzs"`
	Uah  *PriceCurrencyOptionsUah  `json:"uah"`
	Ugx  *PriceCurrencyOptionsUgx  `json:"ugx"`
	USD  *PriceCurrencyOptionsUSD  `json:"usd"`
	Usdc *PriceCurrencyOptionsUsdc `json:"usdc"`
	Uyu  *PriceCurrencyOptionsUyu  `json:"uyu"`
	Uzs  *PriceCurrencyOptionsUzs  `json:"uzs"`
	Vnd  *PriceCurrencyOptionsVnd  `json:"vnd"`
	Vuv  *PriceCurrencyOptionsVuv  `json:"vuv"`
	Wst  *PriceCurrencyOptionsWst  `json:"wst"`
	Xaf  *PriceCurrencyOptionsXaf  `json:"xaf"`
	Xcd  *PriceCurrencyOptionsXcd  `json:"xcd"`
	Xof  *PriceCurrencyOptionsXof  `json:"xof"`
	Xpf  *PriceCurrencyOptionsXpf  `json:"xpf"`
	Yer  *PriceCurrencyOptionsYer  `json:"yer"`
	Zar  *PriceCurrencyOptionsZar  `json:"zar"`
	Zmw  *PriceCurrencyOptionsZmw  `json:"zmw"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type PriceCustomUnitAmount struct {
	// The maximum unit amount the customer can specify for this item.
	Maximum int64 `json:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum int64 `json:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset int64 `json:"preset"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type PriceRecurring struct {
	// Specifies a usage aggregation strategy for prices of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
	AggregateUsage PriceRecurringAggregateUsage `json:"aggregate_usage"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval PriceRecurringInterval `json:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int64 `json:"interval_count"`
	// Default number of trial days when subscribing a customer to this price using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).
	TrialPeriodDays int64 `json:"trial_period_days"`
	// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
	UsageType PriceRecurringUsageType `json:"usage_type"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
type PriceTransformQuantity struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by"`
	// After division, either round the result `up` or `down`.
	Round PriceTransformQuantityRound `json:"round"`
}

// Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.
// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.
//
// For example, you might have a single "gold" product that has prices for $10/month, $100/year, and €9 once.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription), [create an invoice](https://stripe.com/docs/billing/invoices/create), and more about [products and prices](https://stripe.com/docs/products-prices/overview).
type Price struct {
	APIResource
	// Whether the price can be used for new purchases.
	Active bool `json:"active"`
	// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
	BillingScheme PriceBillingScheme `json:"billing_scheme"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency        Currency              `json:"currency"`
	CurrencyOptions *PriceCurrencyOptions `json:"currency_options"`
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *PriceCustomUnitAmount `json:"custom_unit_amount"`
	Deleted          bool                   `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A brief description of the price, hidden from customers.
	Nickname string `json:"nickname"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the product this price is associated with.
	Product *Product `json:"product"`
	// The recurring components of a price such as `interval` and `usage_type`.
	Recurring *PriceRecurring `json:"recurring"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior PriceTaxBehavior `json:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PriceTier `json:"tiers"`
	// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
	TiersMode PriceTiersMode `json:"tiers_mode"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
	TransformQuantity *PriceTransformQuantity `json:"transform_quantity"`
	// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
	Type PriceType `json:"type"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	UnitAmount int64 `json:"unit_amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// PriceList is a list of Prices as retrieved from a list endpoint.
type PriceList struct {
	APIResource
	ListMeta
	Data []*Price `json:"data"`
}

// PriceSearchResult is a list of Price search results as retrieved from a search endpoint.
type PriceSearchResult struct {
	APIResource
	SearchMeta
	Data []*Price `json:"data"`
}

// UnmarshalJSON handles deserialization of a Price.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Price) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type price Price
	var v price
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Price(v)
	return nil
}
