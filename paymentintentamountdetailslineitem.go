//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of the line item.
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory string

// List of values that PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory can take
const (
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryDigitalGoods  PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "digital_goods"
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryDonation      PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "donation"
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryPhysicalGoods PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "physical_goods"
)

// Lists all LineItems of a given PaymentIntent.
type PaymentIntentAmountDetailsLineItemListParams struct {
	ListParams `form:"*"`
	Intent     *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentIntentAmountDetailsLineItemListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardEvCharging struct {
	// The carbon footprint avoided by the charging session, in grams of CO2.
	CarbonFootprintAvoidedGramsCo2 int64 `json:"carbon_footprint_avoided_grams_co2,omitempty"`
	// The time the charging session ended, measured in seconds since the Unix epoch.
	ChargingEndedAt int64 `json:"charging_ended_at,omitempty"`
	// The power output capacity of the charging station, in kilowatts (kW).
	ChargingPowerOutputCapacityKw int64 `json:"charging_power_output_capacity_kw,omitempty"`
	// The time the charging session started, measured in seconds since the Unix epoch.
	ChargingStartedAt int64 `json:"charging_started_at,omitempty"`
	// The type of connector used for the charging session.
	ConnectorType string `json:"connector_type,omitempty"`
	// The estimated distance in kilometers or miles added to the vehicle during the charging session.
	EstimatedRangeAdded int64 `json:"estimated_range_added,omitempty"`
	// The estimated distance in kilometers or miles remaining in the vehicle after the charging session.
	EstimatedRangeLeft int64 `json:"estimated_range_left,omitempty"`
	// The maximum power dispensed during the charging session, in kilowatts (kW).
	MaximumPowerDispensedKw int64 `json:"maximum_power_dispensed_kw,omitempty"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardFleetData struct {
	// The type of product being purchased at this line item.
	ProductType string `json:"product_type,omitempty"`
	// The type of service received at the acceptor location.
	ServiceType string `json:"service_type,omitempty"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard struct {
	CommodityCode string                                                                `json:"commodity_code"`
	EvCharging    *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardEvCharging `json:"ev_charging,omitempty"`
	FleetData     *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardFleetData  `json:"fleet_data,omitempty"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardPresent struct {
	CommodityCode string `json:"commodity_code"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna struct {
	ImageURL              string `json:"image_url"`
	ProductURL            string `json:"product_url"`
	Reference             string `json:"reference"`
	SubscriptionReference string `json:"subscription_reference"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypal struct {
	// Type of the line item.
	Category PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory `json:"category,omitempty"`
	// Description of the line item.
	Description string `json:"description,omitempty"`
	// The Stripe account ID of the connected account that sells the item. This is only needed when using [Separate Charges and Transfers](https://docs.stripe.com/connect/separate-charges-and-transfers).
	SoldBy string `json:"sold_by,omitempty"`
}

// Payment method-specific information for line items.
type PaymentIntentAmountDetailsLineItemPaymentMethodOptions struct {
	Card        *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard        `json:"card,omitempty"`
	CardPresent *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardPresent `json:"card_present,omitempty"`
	Klarna      *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna      `json:"klarna,omitempty"`
	Paypal      *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypal      `json:"paypal,omitempty"`
}

// Contains information about the tax on the item.
type PaymentIntentAmountDetailsLineItemTax struct {
	// The total amount of tax on the transaction represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Required for L2 rates. An integer greater than or equal to 0.
	//
	// This field is mutually exclusive with the `amount_details[line_items][#][tax][total_tax_amount]` field.
	TotalTaxAmount int64 `json:"total_tax_amount"`
}
type PaymentIntentAmountDetailsLineItem struct {
	// The discount applied on this line item represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). An integer greater than 0.
	//
	// This field is mutually exclusive with the `amount_details[discount_amount]` field.
	DiscountAmount int64 `json:"discount_amount"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Payment method-specific information for line items.
	PaymentMethodOptions *PaymentIntentAmountDetailsLineItemPaymentMethodOptions `json:"payment_method_options"`
	// The product code of the line item, such as an SKU. Required for L3 rates. At most 12 characters long.
	ProductCode string `json:"product_code"`
	// The product name of the line item. Required for L3 rates. At most 1024 characters long.
	//
	// For Cards, this field is truncated to 26 alphanumeric characters before being sent to the card networks. For PayPal, this field is truncated to 127 characters.
	ProductName string `json:"product_name"`
	// The quantity of items. Required for L3 rates. An integer greater than 0.
	Quantity int64 `json:"quantity"`
	// The number of decimal places implied in the quantity. For example, if quantity is 10000 and quantity_precision is 2, the actual quantity is 100.00. Defaults to 0 if not provided.
	QuantityPrecision int64 `json:"quantity_precision,omitempty"`
	// Contains information about the tax on the item.
	Tax *PaymentIntentAmountDetailsLineItemTax `json:"tax"`
	// The unit cost of the line item represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Required for L3 rates. An integer greater than or equal to 0.
	UnitCost int64 `json:"unit_cost"`
	// The number of decimal places implied in the unit_cost. For example, if unit_cost is 10000 and unit_cost_precision is 1, the actual unit cost is 1000.0. Defaults to 0 if not provided.
	UnitCostPrecision int64 `json:"unit_cost_precision,omitempty"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc. Required for L3 rates. At most 12 alphanumeric characters long.
	UnitOfMeasure string `json:"unit_of_measure"`
}

// PaymentIntentAmountDetailsLineItemList is a list of PaymentIntentAmountDetailsLineItems as retrieved from a list endpoint.
type PaymentIntentAmountDetailsLineItemList struct {
	APIResource
	ListMeta
	Data []*PaymentIntentAmountDetailsLineItem `json:"data"`
}
