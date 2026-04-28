//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Refund a test-mode Transaction.
type TestHelpersIssuingTransactionRefundParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The total amount to attempt to refund. This amount is in the provided currency, or defaults to the cards currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	RefundAmount *int64 `form:"refund_amount" json:"refund_amount,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingTransactionRefundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams struct {
	// A categorization of the seller's type of business. See our [merchant categories guide](https://docs.stripe.com/issuing/merchant-categories) for a list of possible values.
	Category *string `form:"category" json:"category,omitempty"`
	// City where the seller is located
	City *string `form:"city" json:"city,omitempty"`
	// Country where the seller is located
	Country *string `form:"country" json:"country,omitempty"`
	// Name of the seller
	Name *string `form:"name" json:"name,omitempty"`
	// Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant.
	NetworkID *string `form:"network_id" json:"network_id,omitempty"`
	// The identifier of the payment facilitator (PayFac) that processed this authorization, as assigned by the card network.
	PaymentFacilitatorID *string `form:"payment_facilitator_id" json:"payment_facilitator_id,omitempty"`
	// Postal code where the seller is located
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State where the seller is located
	State *string `form:"state" json:"state,omitempty"`
	// The identifier of the sub-merchant involved in this authorization, as assigned by the payment facilitator.
	SubMerchantID *string `form:"sub_merchant_id" json:"sub_merchant_id,omitempty"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id" json:"terminal_id,omitempty"`
	// URL provided by the merchant on a 3DS request
	URL *string `form:"url" json:"url,omitempty"`
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetCardholderPromptDataParams struct {
	// Driver ID.
	DriverID *string `form:"driver_id" json:"driver_id,omitempty"`
	// Odometer reading.
	Odometer *int64 `form:"odometer" json:"odometer,omitempty"`
	// An alphanumeric ID. This field is used when a vehicle ID, driver ID, or generic ID is entered by the cardholder, but the merchant or card network did not specify the prompt type.
	UnspecifiedID *string `form:"unspecified_id" json:"unspecified_id,omitempty"`
	// User ID.
	UserID *string `form:"user_id" json:"user_id,omitempty"`
	// Vehicle number.
	VehicleNumber *string `form:"vehicle_number" json:"vehicle_number,omitempty"`
}

// Breakdown of fuel portion of the purchase.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for transactions using Fleet cards.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// The legs of the trip.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightSegmentParams struct {
	// The three-letter IATA airport code of the flight's destination.
	ArrivalAirportCode *string `form:"arrival_airport_code" json:"arrival_airport_code,omitempty"`
	// The airline carrier code.
	Carrier *string `form:"carrier" json:"carrier,omitempty"`
	// The three-letter IATA airport code that the flight departed from.
	DepartureAirportCode *string `form:"departure_airport_code" json:"departure_airport_code,omitempty"`
	// The flight number.
	FlightNumber *string `form:"flight_number" json:"flight_number,omitempty"`
	// The flight's service class.
	ServiceClass *string `form:"service_class" json:"service_class,omitempty"`
	// Whether a stopover is allowed on this flight.
	StopoverAllowed *bool `form:"stopover_allowed" json:"stopover_allowed,omitempty"`
}

// Information about the flight that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightParams struct {
	// The time that the flight departed.
	DepartureAt *int64 `form:"departure_at" json:"departure_at,omitempty"`
	// The name of the passenger.
	PassengerName *string `form:"passenger_name" json:"passenger_name,omitempty"`
	// Whether the ticket is refundable.
	Refundable *bool `form:"refundable" json:"refundable,omitempty"`
	// The legs of the trip.
	Segments []*TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightSegmentParams `form:"segments" json:"segments,omitempty"`
	// The travel agency that issued the ticket.
	TravelAgency *string `form:"travel_agency" json:"travel_agency,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams struct {
	// [Conexxus Payment System Product Code](https://www.conexxus.org/conexxus-payment-system-product-codes) identifying the primary fuel product purchased.
	IndustryProductCode *string `form:"industry_product_code" json:"industry_product_code,omitempty"`
	// The quantity of `unit`s of fuel that was dispensed, represented as a decimal string with at most 12 decimal places.
	QuantityDecimal *float64 `form:"quantity_decimal,high_precision" json:"quantity_decimal,string,omitempty"`
	// The type of fuel that was purchased. One of `diesel`, `unleaded_plus`, `unleaded_regular`, `unleaded_super`, or `other`.
	Type *string `form:"type" json:"type,omitempty"`
	// The units for `quantity_decimal`. One of `charging_minute`, `imperial_gallon`, `kilogram`, `kilowatt_hour`, `liter`, `pound`, `us_gallon`, or `other`.
	Unit *string `form:"unit" json:"unit,omitempty"`
	// The cost in cents per each unit of fuel, represented as a decimal string with at most 12 decimal places.
	UnitCostDecimal *float64 `form:"unit_cost_decimal,high_precision" json:"unit_cost_decimal,string,omitempty"`
}

// Information about lodging that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsLodgingParams struct {
	// The time of checking into the lodging.
	CheckInAt *int64 `form:"check_in_at" json:"check_in_at,omitempty"`
	// The number of nights stayed at the lodging.
	Nights *int64 `form:"nights" json:"nights,omitempty"`
}

// The line items in the purchase.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsReceiptParams struct {
	Description *string  `form:"description" json:"description,omitempty"`
	Quantity    *float64 `form:"quantity,high_precision" json:"quantity,string,omitempty"`
	Total       *int64   `form:"total" json:"total,omitempty"`
	UnitCost    *int64   `form:"unit_cost" json:"unit_cost,omitempty"`
}

// Additional purchase information that is optionally provided by the merchant.
type TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams struct {
	// Fleet-specific information for transactions using Fleet cards.
	Fleet *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Information about the flight that was purchased with this transaction.
	Flight *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFlightParams `form:"flight" json:"flight,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Information about lodging that was purchased with this transaction.
	Lodging *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsLodgingParams `form:"lodging" json:"lodging,omitempty"`
	// The line items in the purchase.
	Receipt []*TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsReceiptParams `form:"receipt" json:"receipt,omitempty"`
	// A merchant-specific order number.
	Reference *string `form:"reference" json:"reference,omitempty"`
}

// Allows the user to capture an arbitrary amount, also known as a forced capture.
type TestHelpersIssuingTransactionCreateForceCaptureParams struct {
	Params `form:"*"`
	// The total amount to attempt to capture. This amount is in the provided currency, or defaults to the cards currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Card associated with this transaction.
	Card *string `form:"card" json:"card"`
	// The currency of the capture. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantData *TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams `form:"merchant_data" json:"merchant_data,omitempty"`
	// Additional purchase information that is optionally provided by the merchant.
	PurchaseDetails *TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams `form:"purchase_details" json:"purchase_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingTransactionCreateForceCaptureParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams struct {
	// A categorization of the seller's type of business. See our [merchant categories guide](https://docs.stripe.com/issuing/merchant-categories) for a list of possible values.
	Category *string `form:"category" json:"category,omitempty"`
	// City where the seller is located
	City *string `form:"city" json:"city,omitempty"`
	// Country where the seller is located
	Country *string `form:"country" json:"country,omitempty"`
	// Name of the seller
	Name *string `form:"name" json:"name,omitempty"`
	// Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant.
	NetworkID *string `form:"network_id" json:"network_id,omitempty"`
	// The identifier of the payment facilitator (PayFac) that processed this authorization, as assigned by the card network.
	PaymentFacilitatorID *string `form:"payment_facilitator_id" json:"payment_facilitator_id,omitempty"`
	// Postal code where the seller is located
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State where the seller is located
	State *string `form:"state" json:"state,omitempty"`
	// The identifier of the sub-merchant involved in this authorization, as assigned by the payment facilitator.
	SubMerchantID *string `form:"sub_merchant_id" json:"sub_merchant_id,omitempty"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id" json:"terminal_id,omitempty"`
	// URL provided by the merchant on a 3DS request
	URL *string `form:"url" json:"url,omitempty"`
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetCardholderPromptDataParams struct {
	// Driver ID.
	DriverID *string `form:"driver_id" json:"driver_id,omitempty"`
	// Odometer reading.
	Odometer *int64 `form:"odometer" json:"odometer,omitempty"`
	// An alphanumeric ID. This field is used when a vehicle ID, driver ID, or generic ID is entered by the cardholder, but the merchant or card network did not specify the prompt type.
	UnspecifiedID *string `form:"unspecified_id" json:"unspecified_id,omitempty"`
	// User ID.
	UserID *string `form:"user_id" json:"user_id,omitempty"`
	// Vehicle number.
	VehicleNumber *string `form:"vehicle_number" json:"vehicle_number,omitempty"`
}

// Breakdown of fuel portion of the purchase.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for transactions using Fleet cards.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// The legs of the trip.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightSegmentParams struct {
	// The three-letter IATA airport code of the flight's destination.
	ArrivalAirportCode *string `form:"arrival_airport_code" json:"arrival_airport_code,omitempty"`
	// The airline carrier code.
	Carrier *string `form:"carrier" json:"carrier,omitempty"`
	// The three-letter IATA airport code that the flight departed from.
	DepartureAirportCode *string `form:"departure_airport_code" json:"departure_airport_code,omitempty"`
	// The flight number.
	FlightNumber *string `form:"flight_number" json:"flight_number,omitempty"`
	// The flight's service class.
	ServiceClass *string `form:"service_class" json:"service_class,omitempty"`
	// Whether a stopover is allowed on this flight.
	StopoverAllowed *bool `form:"stopover_allowed" json:"stopover_allowed,omitempty"`
}

// Information about the flight that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightParams struct {
	// The time that the flight departed.
	DepartureAt *int64 `form:"departure_at" json:"departure_at,omitempty"`
	// The name of the passenger.
	PassengerName *string `form:"passenger_name" json:"passenger_name,omitempty"`
	// Whether the ticket is refundable.
	Refundable *bool `form:"refundable" json:"refundable,omitempty"`
	// The legs of the trip.
	Segments []*TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightSegmentParams `form:"segments" json:"segments,omitempty"`
	// The travel agency that issued the ticket.
	TravelAgency *string `form:"travel_agency" json:"travel_agency,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams struct {
	// [Conexxus Payment System Product Code](https://www.conexxus.org/conexxus-payment-system-product-codes) identifying the primary fuel product purchased.
	IndustryProductCode *string `form:"industry_product_code" json:"industry_product_code,omitempty"`
	// The quantity of `unit`s of fuel that was dispensed, represented as a decimal string with at most 12 decimal places.
	QuantityDecimal *float64 `form:"quantity_decimal,high_precision" json:"quantity_decimal,string,omitempty"`
	// The type of fuel that was purchased. One of `diesel`, `unleaded_plus`, `unleaded_regular`, `unleaded_super`, or `other`.
	Type *string `form:"type" json:"type,omitempty"`
	// The units for `quantity_decimal`. One of `charging_minute`, `imperial_gallon`, `kilogram`, `kilowatt_hour`, `liter`, `pound`, `us_gallon`, or `other`.
	Unit *string `form:"unit" json:"unit,omitempty"`
	// The cost in cents per each unit of fuel, represented as a decimal string with at most 12 decimal places.
	UnitCostDecimal *float64 `form:"unit_cost_decimal,high_precision" json:"unit_cost_decimal,string,omitempty"`
}

// Information about lodging that was purchased with this transaction.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsLodgingParams struct {
	// The time of checking into the lodging.
	CheckInAt *int64 `form:"check_in_at" json:"check_in_at,omitempty"`
	// The number of nights stayed at the lodging.
	Nights *int64 `form:"nights" json:"nights,omitempty"`
}

// The line items in the purchase.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsReceiptParams struct {
	Description *string  `form:"description" json:"description,omitempty"`
	Quantity    *float64 `form:"quantity,high_precision" json:"quantity,string,omitempty"`
	Total       *int64   `form:"total" json:"total,omitempty"`
	UnitCost    *int64   `form:"unit_cost" json:"unit_cost,omitempty"`
}

// Additional purchase information that is optionally provided by the merchant.
type TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams struct {
	// Fleet-specific information for transactions using Fleet cards.
	Fleet *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Information about the flight that was purchased with this transaction.
	Flight *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFlightParams `form:"flight" json:"flight,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Information about lodging that was purchased with this transaction.
	Lodging *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsLodgingParams `form:"lodging" json:"lodging,omitempty"`
	// The line items in the purchase.
	Receipt []*TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsReceiptParams `form:"receipt" json:"receipt,omitempty"`
	// A merchant-specific order number.
	Reference *string `form:"reference" json:"reference,omitempty"`
}

// Allows the user to refund an arbitrary amount, also known as a unlinked refund.
type TestHelpersIssuingTransactionCreateUnlinkedRefundParams struct {
	Params `form:"*"`
	// The total amount to attempt to refund. This amount is in the provided currency, or defaults to the cards currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Card associated with this unlinked refund transaction.
	Card *string `form:"card" json:"card"`
	// The currency of the unlinked refund. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantData *TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams `form:"merchant_data" json:"merchant_data,omitempty"`
	// Additional purchase information that is optionally provided by the merchant.
	PurchaseDetails *TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams `form:"purchase_details" json:"purchase_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingTransactionCreateUnlinkedRefundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
