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
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Currency     *string           `form:"currency"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type ShippingRateDeliveryEstimateMaximumParams struct {
	Unit  *string `form:"unit"`
	Value *int64  `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type ShippingRateDeliveryEstimateMinimumParams struct {
	Unit  *string `form:"unit"`
	Value *int64  `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type ShippingRateDeliveryEstimateParams struct {
	Maximum *ShippingRateDeliveryEstimateMaximumParams `form:"maximum"`
	Minimum *ShippingRateDeliveryEstimateMinimumParams `form:"minimum"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type ShippingRateFixedAmountParams struct {
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}

// Creates a new shipping rate object.
type ShippingRateParams struct {
	Params           `form:"*"`
	Active           *bool                               `form:"active"`
	DeliveryEstimate *ShippingRateDeliveryEstimateParams `form:"delivery_estimate"`
	DisplayName      *string                             `form:"display_name"`
	FixedAmount      *ShippingRateFixedAmountParams      `form:"fixed_amount"`
	TaxBehavior      *string                             `form:"tax_behavior"`
	TaxCode          *string                             `form:"tax_code"`
	Type             *string                             `form:"type"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type ShippingRateDeliveryEstimateMaximum struct {
	Unit  ShippingRateDeliveryEstimateMaximumUnit `json:"unit"`
	Value int64                                   `json:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type ShippingRateDeliveryEstimateMinimum struct {
	Unit  ShippingRateDeliveryEstimateMinimumUnit `json:"unit"`
	Value int64                                   `json:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type ShippingRateDeliveryEstimate struct {
	Maximum *ShippingRateDeliveryEstimateMaximum `json:"maximum"`
	Minimum *ShippingRateDeliveryEstimateMinimum `json:"minimum"`
}
type ShippingRateFixedAmount struct {
	Amount   int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// Shipping rates describe the price of shipping presented to your customers and can be
// applied to [Checkout Sessions](https://stripe.com/docs/payments/checkout/shipping) to collect shipping costs.
type ShippingRate struct {
	APIResource
	Active           bool                          `json:"active"`
	Created          int64                         `json:"created"`
	DeliveryEstimate *ShippingRateDeliveryEstimate `json:"delivery_estimate"`
	DisplayName      string                        `json:"display_name"`
	FixedAmount      *ShippingRateFixedAmount      `json:"fixed_amount"`
	ID               string                        `json:"id"`
	Livemode         bool                          `json:"livemode"`
	Metadata         map[string]string             `json:"metadata"`
	Object           string                        `json:"object"`
	TaxBehavior      ShippingRateTaxBehavior       `json:"tax_behavior"`
	TaxCode          *TaxCode                      `json:"tax_code"`
	Type             ShippingRateType              `json:"type"`
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
