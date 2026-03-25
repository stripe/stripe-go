//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
type TestHelpersIssuingAuthorizationAmountDetailsParams struct {
	// The ATM withdrawal fee.
	ATMFee *int64 `form:"atm_fee" json:"atm_fee,omitempty"`
	// The amount of cash requested by the cardholder.
	CashbackAmount *int64 `form:"cashback_amount" json:"cashback_amount,omitempty"`
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingAuthorizationFleetCardholderPromptDataParams struct {
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
type TestHelpersIssuingAuthorizationFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingAuthorizationFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingAuthorizationFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingAuthorizationFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingAuthorizationFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingAuthorizationFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingAuthorizationFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for authorizations using Fleet cards.
type TestHelpersIssuingAuthorizationFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingAuthorizationFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingAuthorizationFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingAuthorizationFuelParams struct {
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

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type TestHelpersIssuingAuthorizationMerchantDataParams struct {
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
	// Postal code where the seller is located
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State where the seller is located
	State *string `form:"state" json:"state,omitempty"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id" json:"terminal_id,omitempty"`
	// URL provided by the merchant on a 3DS request
	URL *string `form:"url" json:"url,omitempty"`
}

// Details about the authorization, such as identifiers, set by the card network.
type TestHelpersIssuingAuthorizationNetworkDataParams struct {
	// Identifier assigned to the acquirer by the card network.
	AcquiringInstitutionID *string `form:"acquiring_institution_id" json:"acquiring_institution_id,omitempty"`
}

// Stripe's assessment of this authorization's likelihood of being card testing activity.
type TestHelpersIssuingAuthorizationRiskAssessmentCardTestingRiskParams struct {
	// The % of declines due to a card number not existing in the past hour, taking place at the same merchant. Higher rates correspond to a greater probability of card testing activity, meaning bad actors may be attempting different card number combinations to guess a correct one. Takes on values between 0 and 100.
	InvalidAccountNumberDeclineRatePastHour *int64 `form:"invalid_account_number_decline_rate_past_hour" json:"invalid_account_number_decline_rate_past_hour,omitempty"`
	// The % of declines due to incorrect verification data (like CVV or expiry) in the past hour, taking place at the same merchant. Higher rates correspond to a greater probability of bad actors attempting to utilize valid card credentials at merchants with verification requirements. Takes on values between 0 and 100.
	InvalidCredentialsDeclineRatePastHour *int64 `form:"invalid_credentials_decline_rate_past_hour" json:"invalid_credentials_decline_rate_past_hour,omitempty"`
	// The likelihood that this authorization is associated with card testing activity. This is assessed by evaluating decline activity over the last hour.
	RiskLevel *string `form:"risk_level" json:"risk_level"`
}

// Stripe's assessment of this authorization's likelihood to be fraudulent.
type TestHelpersIssuingAuthorizationRiskAssessmentFraudRiskParams struct {
	// Stripe's assessment of the likelihood of fraud on an authorization.
	Level *string `form:"level" json:"level"`
	// Stripe's numerical model score assessing the likelihood of fraudulent activity. A higher score means a higher likelihood of fraudulent activity, and anything above 25 is considered high risk.
	Score *float64 `form:"score" json:"score,omitempty"`
}

// The dispute risk of the merchant (the seller on a purchase) on an authorization based on all Stripe Issuing activity.
type TestHelpersIssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams struct {
	// The dispute rate observed across all Stripe Issuing authorizations for this merchant. For example, a value of 50 means 50% of authorizations from this merchant on Stripe Issuing have resulted in a dispute. Higher values mean a higher likelihood the authorization is disputed. Takes on values between 0 and 100.
	DisputeRate *int64 `form:"dispute_rate" json:"dispute_rate,omitempty"`
	// The likelihood that authorizations from this merchant will result in a dispute based on their history on Stripe Issuing.
	RiskLevel *string `form:"risk_level" json:"risk_level"`
}

// Stripe's assessment of the fraud risk for this authorization.
type TestHelpersIssuingAuthorizationRiskAssessmentParams struct {
	// Stripe's assessment of this authorization's likelihood of being card testing activity.
	CardTestingRisk *TestHelpersIssuingAuthorizationRiskAssessmentCardTestingRiskParams `form:"card_testing_risk" json:"card_testing_risk,omitempty"`
	// Stripe's assessment of this authorization's likelihood to be fraudulent.
	FraudRisk *TestHelpersIssuingAuthorizationRiskAssessmentFraudRiskParams `form:"fraud_risk" json:"fraud_risk,omitempty"`
	// The dispute risk of the merchant (the seller on a purchase) on an authorization based on all Stripe Issuing activity.
	MerchantDisputeRisk *TestHelpersIssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams `form:"merchant_dispute_risk" json:"merchant_dispute_risk,omitempty"`
}

// The exemption applied to this authorization.
type TestHelpersIssuingAuthorizationVerificationDataAuthenticationExemptionParams struct {
	// The entity that requested the exemption, either the acquiring merchant or the Issuing user.
	ClaimedBy *string `form:"claimed_by" json:"claimed_by"`
	// The specific exemption claimed for this authorization.
	Type *string `form:"type" json:"type"`
}

// 3D Secure details.
type TestHelpersIssuingAuthorizationVerificationDataThreeDSecureParams struct {
	// The outcome of the 3D Secure authentication request.
	Result *string `form:"result" json:"result"`
}

// Verifications that Stripe performed on information that the cardholder provided to the merchant.
type TestHelpersIssuingAuthorizationVerificationDataParams struct {
	// Whether the cardholder provided an address first line and if it matched the cardholder's `billing.address.line1`.
	AddressLine1Check *string `form:"address_line1_check" json:"address_line1_check,omitempty"`
	// Whether the cardholder provided a postal code and if it matched the cardholder's `billing.address.postal_code`.
	AddressPostalCodeCheck *string `form:"address_postal_code_check" json:"address_postal_code_check,omitempty"`
	// The exemption applied to this authorization.
	AuthenticationExemption *TestHelpersIssuingAuthorizationVerificationDataAuthenticationExemptionParams `form:"authentication_exemption" json:"authentication_exemption,omitempty"`
	// Whether the cardholder provided a CVC and if it matched Stripe's record.
	CVCCheck *string `form:"cvc_check" json:"cvc_check,omitempty"`
	// Whether the cardholder provided an expiry date and if it matched Stripe's record.
	ExpiryCheck *string `form:"expiry_check" json:"expiry_check,omitempty"`
	// 3D Secure details.
	ThreeDSecure *TestHelpersIssuingAuthorizationVerificationDataThreeDSecureParams `form:"three_d_secure" json:"three_d_secure,omitempty"`
}

// Create a test-mode authorization.
type TestHelpersIssuingAuthorizationParams struct {
	Params `form:"*"`
	// The total amount to attempt to authorize. This amount is in the provided currency, or defaults to the card's currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	AmountDetails *TestHelpersIssuingAuthorizationAmountDetailsParams `form:"amount_details" json:"amount_details,omitempty"`
	// How the card details were provided. Defaults to online.
	AuthorizationMethod *string `form:"authorization_method" json:"authorization_method,omitempty"`
	// Card associated with this authorization.
	Card *string `form:"card" json:"card"`
	// The currency of the authorization. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Fleet-specific information for authorizations using Fleet cards.
	Fleet *TestHelpersIssuingAuthorizationFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Probability that this transaction can be disputed in the event of fraud. Assessed by comparing the characteristics of the authorization to card network rules.
	FraudDisputabilityLikelihood *string `form:"fraud_disputability_likelihood" json:"fraud_disputability_likelihood,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingAuthorizationFuelParams `form:"fuel" json:"fuel,omitempty"`
	// If set `true`, you may provide [amount](https://docs.stripe.com/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization.
	IsAmountControllable *bool `form:"is_amount_controllable" json:"is_amount_controllable,omitempty"`
	// The total amount to attempt to authorize. This amount is in the provided merchant currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	MerchantAmount *int64 `form:"merchant_amount" json:"merchant_amount,omitempty"`
	// The currency of the authorization. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	MerchantCurrency *string `form:"merchant_currency" json:"merchant_currency,omitempty"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantData *TestHelpersIssuingAuthorizationMerchantDataParams `form:"merchant_data" json:"merchant_data,omitempty"`
	// Details about the authorization, such as identifiers, set by the card network.
	NetworkData *TestHelpersIssuingAuthorizationNetworkDataParams `form:"network_data" json:"network_data,omitempty"`
	// Stripe's assessment of the fraud risk for this authorization.
	RiskAssessment *TestHelpersIssuingAuthorizationRiskAssessmentParams `form:"risk_assessment" json:"risk_assessment,omitempty"`
	// Verifications that Stripe performed on information that the cardholder provided to the merchant.
	VerificationData *TestHelpersIssuingAuthorizationVerificationDataParams `form:"verification_data" json:"verification_data,omitempty"`
	// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`. Will populate as `null` when no digital wallet was utilized.
	Wallet *string `form:"wallet" json:"wallet,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetCardholderPromptDataParams struct {
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
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for transactions using Fleet cards.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// The legs of the trip.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightSegmentParams struct {
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
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightParams struct {
	// The time that the flight departed.
	DepartureAt *int64 `form:"departure_at" json:"departure_at,omitempty"`
	// The name of the passenger.
	PassengerName *string `form:"passenger_name" json:"passenger_name,omitempty"`
	// Whether the ticket is refundable.
	Refundable *bool `form:"refundable" json:"refundable,omitempty"`
	// The legs of the trip.
	Segments []*TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightSegmentParams `form:"segments" json:"segments,omitempty"`
	// The travel agency that issued the ticket.
	TravelAgency *string `form:"travel_agency" json:"travel_agency,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams struct {
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
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsLodgingParams struct {
	// The time of checking into the lodging.
	CheckInAt *int64 `form:"check_in_at" json:"check_in_at,omitempty"`
	// The number of nights stayed at the lodging.
	Nights *int64 `form:"nights" json:"nights,omitempty"`
}

// The line items in the purchase.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsReceiptParams struct {
	Description *string  `form:"description" json:"description,omitempty"`
	Quantity    *float64 `form:"quantity,high_precision" json:"quantity,string,omitempty"`
	Total       *int64   `form:"total" json:"total,omitempty"`
	UnitCost    *int64   `form:"unit_cost" json:"unit_cost,omitempty"`
}

// Additional purchase information that is optionally provided by the merchant.
type TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams struct {
	// Fleet-specific information for transactions using Fleet cards.
	Fleet *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Information about the flight that was purchased with this transaction.
	Flight *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFlightParams `form:"flight" json:"flight,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Information about lodging that was purchased with this transaction.
	Lodging *TestHelpersIssuingAuthorizationCapturePurchaseDetailsLodgingParams `form:"lodging" json:"lodging,omitempty"`
	// The line items in the purchase.
	Receipt []*TestHelpersIssuingAuthorizationCapturePurchaseDetailsReceiptParams `form:"receipt" json:"receipt,omitempty"`
	// A merchant-specific order number.
	Reference *string `form:"reference" json:"reference,omitempty"`
}

// Capture a test-mode authorization.
type TestHelpersIssuingAuthorizationCaptureParams struct {
	Params `form:"*"`
	// The amount to capture from the authorization. If not provided, the full amount of the authorization will be captured. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	CaptureAmount *int64 `form:"capture_amount" json:"capture_amount,omitempty"`
	// Whether to close the authorization after capture. Defaults to true. Set to false to enable multi-capture flows.
	CloseAuthorization *bool `form:"close_authorization" json:"close_authorization,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Additional purchase information that is optionally provided by the merchant.
	PurchaseDetails *TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams `form:"purchase_details" json:"purchase_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationCaptureParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Expire a test-mode Authorization.
type TestHelpersIssuingAuthorizationExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingAuthorizationFinalizeAmountFleetCardholderPromptDataParams struct {
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
type TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for authorizations using Fleet cards.
type TestHelpersIssuingAuthorizationFinalizeAmountFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingAuthorizationFinalizeAmountFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingAuthorizationFinalizeAmountFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingAuthorizationFinalizeAmountFuelParams struct {
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

// Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.
type TestHelpersIssuingAuthorizationFinalizeAmountParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The final authorization amount that will be captured by the merchant. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	FinalAmount *int64 `form:"final_amount" json:"final_amount"`
	// Fleet-specific information for authorizations using Fleet cards.
	Fleet *TestHelpersIssuingAuthorizationFinalizeAmountFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingAuthorizationFinalizeAmountFuelParams `form:"fuel" json:"fuel,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationFinalizeAmountParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.
type TestHelpersIssuingAuthorizationRespondParams struct {
	Params `form:"*"`
	// Whether to simulate the user confirming that the transaction was legitimate (true) or telling Stripe that it was fraudulent (false).
	Confirmed *bool `form:"confirmed" json:"confirmed"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationRespondParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Increment a test-mode Authorization.
type TestHelpersIssuingAuthorizationIncrementParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The amount to increment the authorization by. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	IncrementAmount *int64 `form:"increment_amount" json:"increment_amount"`
	// If set `true`, you may provide [amount](https://docs.stripe.com/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization.
	IsAmountControllable *bool `form:"is_amount_controllable" json:"is_amount_controllable,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationIncrementParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reverse a test-mode Authorization.
type TestHelpersIssuingAuthorizationReverseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The amount to reverse from the authorization. If not provided, the full amount of the authorization will be reversed. This amount is in the authorization currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	ReverseAmount *int64 `form:"reverse_amount" json:"reverse_amount,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationReverseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
type TestHelpersIssuingAuthorizationCreateAmountDetailsParams struct {
	// The ATM withdrawal fee.
	ATMFee *int64 `form:"atm_fee" json:"atm_fee,omitempty"`
	// The amount of cash requested by the cardholder.
	CashbackAmount *int64 `form:"cashback_amount" json:"cashback_amount,omitempty"`
}

// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
type TestHelpersIssuingAuthorizationCreateFleetCardholderPromptDataParams struct {
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
type TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownFuelParams struct {
	// Gross fuel amount that should equal Fuel Volume multipled by Fuel Unit Cost, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Breakdown of non-fuel portion of the purchase.
type TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownNonFuelParams struct {
	// Gross non-fuel amount that should equal the sum of the line items, inclusive of taxes.
	GrossAmountDecimal *float64 `form:"gross_amount_decimal,high_precision" json:"gross_amount_decimal,string,omitempty"`
}

// Information about tax included in this transaction.
type TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownTaxParams struct {
	// Amount of state or provincial Sales Tax included in the transaction amount. Null if not reported by merchant or not subject to tax.
	LocalAmountDecimal *float64 `form:"local_amount_decimal,high_precision" json:"local_amount_decimal,string,omitempty"`
	// Amount of national Sales Tax or VAT included in the transaction amount. Null if not reported by merchant or not subject to tax.
	NationalAmountDecimal *float64 `form:"national_amount_decimal,high_precision" json:"national_amount_decimal,string,omitempty"`
}

// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
type TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownParams struct {
	// Breakdown of fuel portion of the purchase.
	Fuel *TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownFuelParams `form:"fuel" json:"fuel,omitempty"`
	// Breakdown of non-fuel portion of the purchase.
	NonFuel *TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownNonFuelParams `form:"non_fuel" json:"non_fuel,omitempty"`
	// Information about tax included in this transaction.
	Tax *TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownTaxParams `form:"tax" json:"tax,omitempty"`
}

// Fleet-specific information for authorizations using Fleet cards.
type TestHelpersIssuingAuthorizationCreateFleetParams struct {
	// Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.
	CardholderPromptData *TestHelpersIssuingAuthorizationCreateFleetCardholderPromptDataParams `form:"cardholder_prompt_data" json:"cardholder_prompt_data,omitempty"`
	// The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.
	PurchaseType *string `form:"purchase_type" json:"purchase_type,omitempty"`
	// More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.
	ReportedBreakdown *TestHelpersIssuingAuthorizationCreateFleetReportedBreakdownParams `form:"reported_breakdown" json:"reported_breakdown,omitempty"`
	// The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.
	ServiceType *string `form:"service_type" json:"service_type,omitempty"`
}

// Information about fuel that was purchased with this transaction.
type TestHelpersIssuingAuthorizationCreateFuelParams struct {
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

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type TestHelpersIssuingAuthorizationCreateMerchantDataParams struct {
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
	// Postal code where the seller is located
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State where the seller is located
	State *string `form:"state" json:"state,omitempty"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id" json:"terminal_id,omitempty"`
	// URL provided by the merchant on a 3DS request
	URL *string `form:"url" json:"url,omitempty"`
}

// Details about the authorization, such as identifiers, set by the card network.
type TestHelpersIssuingAuthorizationCreateNetworkDataParams struct {
	// Identifier assigned to the acquirer by the card network.
	AcquiringInstitutionID *string `form:"acquiring_institution_id" json:"acquiring_institution_id,omitempty"`
}

// Stripe's assessment of this authorization's likelihood of being card testing activity.
type TestHelpersIssuingAuthorizationCreateRiskAssessmentCardTestingRiskParams struct {
	// The % of declines due to a card number not existing in the past hour, taking place at the same merchant. Higher rates correspond to a greater probability of card testing activity, meaning bad actors may be attempting different card number combinations to guess a correct one. Takes on values between 0 and 100.
	InvalidAccountNumberDeclineRatePastHour *int64 `form:"invalid_account_number_decline_rate_past_hour" json:"invalid_account_number_decline_rate_past_hour,omitempty"`
	// The % of declines due to incorrect verification data (like CVV or expiry) in the past hour, taking place at the same merchant. Higher rates correspond to a greater probability of bad actors attempting to utilize valid card credentials at merchants with verification requirements. Takes on values between 0 and 100.
	InvalidCredentialsDeclineRatePastHour *int64 `form:"invalid_credentials_decline_rate_past_hour" json:"invalid_credentials_decline_rate_past_hour,omitempty"`
	// The likelihood that this authorization is associated with card testing activity. This is assessed by evaluating decline activity over the last hour.
	RiskLevel *string `form:"risk_level" json:"risk_level"`
}

// Stripe's assessment of this authorization's likelihood to be fraudulent.
type TestHelpersIssuingAuthorizationCreateRiskAssessmentFraudRiskParams struct {
	// Stripe's assessment of the likelihood of fraud on an authorization.
	Level *string `form:"level" json:"level"`
	// Stripe's numerical model score assessing the likelihood of fraudulent activity. A higher score means a higher likelihood of fraudulent activity, and anything above 25 is considered high risk.
	Score *float64 `form:"score" json:"score,omitempty"`
}

// The dispute risk of the merchant (the seller on a purchase) on an authorization based on all Stripe Issuing activity.
type TestHelpersIssuingAuthorizationCreateRiskAssessmentMerchantDisputeRiskParams struct {
	// The dispute rate observed across all Stripe Issuing authorizations for this merchant. For example, a value of 50 means 50% of authorizations from this merchant on Stripe Issuing have resulted in a dispute. Higher values mean a higher likelihood the authorization is disputed. Takes on values between 0 and 100.
	DisputeRate *int64 `form:"dispute_rate" json:"dispute_rate,omitempty"`
	// The likelihood that authorizations from this merchant will result in a dispute based on their history on Stripe Issuing.
	RiskLevel *string `form:"risk_level" json:"risk_level"`
}

// Stripe's assessment of the fraud risk for this authorization.
type TestHelpersIssuingAuthorizationCreateRiskAssessmentParams struct {
	// Stripe's assessment of this authorization's likelihood of being card testing activity.
	CardTestingRisk *TestHelpersIssuingAuthorizationCreateRiskAssessmentCardTestingRiskParams `form:"card_testing_risk" json:"card_testing_risk,omitempty"`
	// Stripe's assessment of this authorization's likelihood to be fraudulent.
	FraudRisk *TestHelpersIssuingAuthorizationCreateRiskAssessmentFraudRiskParams `form:"fraud_risk" json:"fraud_risk,omitempty"`
	// The dispute risk of the merchant (the seller on a purchase) on an authorization based on all Stripe Issuing activity.
	MerchantDisputeRisk *TestHelpersIssuingAuthorizationCreateRiskAssessmentMerchantDisputeRiskParams `form:"merchant_dispute_risk" json:"merchant_dispute_risk,omitempty"`
}

// The exemption applied to this authorization.
type TestHelpersIssuingAuthorizationCreateVerificationDataAuthenticationExemptionParams struct {
	// The entity that requested the exemption, either the acquiring merchant or the Issuing user.
	ClaimedBy *string `form:"claimed_by" json:"claimed_by"`
	// The specific exemption claimed for this authorization.
	Type *string `form:"type" json:"type"`
}

// 3D Secure details.
type TestHelpersIssuingAuthorizationCreateVerificationDataThreeDSecureParams struct {
	// The outcome of the 3D Secure authentication request.
	Result *string `form:"result" json:"result"`
}

// Verifications that Stripe performed on information that the cardholder provided to the merchant.
type TestHelpersIssuingAuthorizationCreateVerificationDataParams struct {
	// Whether the cardholder provided an address first line and if it matched the cardholder's `billing.address.line1`.
	AddressLine1Check *string `form:"address_line1_check" json:"address_line1_check,omitempty"`
	// Whether the cardholder provided a postal code and if it matched the cardholder's `billing.address.postal_code`.
	AddressPostalCodeCheck *string `form:"address_postal_code_check" json:"address_postal_code_check,omitempty"`
	// The exemption applied to this authorization.
	AuthenticationExemption *TestHelpersIssuingAuthorizationCreateVerificationDataAuthenticationExemptionParams `form:"authentication_exemption" json:"authentication_exemption,omitempty"`
	// Whether the cardholder provided a CVC and if it matched Stripe's record.
	CVCCheck *string `form:"cvc_check" json:"cvc_check,omitempty"`
	// Whether the cardholder provided an expiry date and if it matched Stripe's record.
	ExpiryCheck *string `form:"expiry_check" json:"expiry_check,omitempty"`
	// 3D Secure details.
	ThreeDSecure *TestHelpersIssuingAuthorizationCreateVerificationDataThreeDSecureParams `form:"three_d_secure" json:"three_d_secure,omitempty"`
}

// Create a test-mode authorization.
type TestHelpersIssuingAuthorizationCreateParams struct {
	Params `form:"*"`
	// The total amount to attempt to authorize. This amount is in the provided currency, or defaults to the card's currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	AmountDetails *TestHelpersIssuingAuthorizationCreateAmountDetailsParams `form:"amount_details" json:"amount_details,omitempty"`
	// How the card details were provided. Defaults to online.
	AuthorizationMethod *string `form:"authorization_method" json:"authorization_method,omitempty"`
	// Card associated with this authorization.
	Card *string `form:"card" json:"card"`
	// The currency of the authorization. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Fleet-specific information for authorizations using Fleet cards.
	Fleet *TestHelpersIssuingAuthorizationCreateFleetParams `form:"fleet" json:"fleet,omitempty"`
	// Probability that this transaction can be disputed in the event of fraud. Assessed by comparing the characteristics of the authorization to card network rules.
	FraudDisputabilityLikelihood *string `form:"fraud_disputability_likelihood" json:"fraud_disputability_likelihood,omitempty"`
	// Information about fuel that was purchased with this transaction.
	Fuel *TestHelpersIssuingAuthorizationCreateFuelParams `form:"fuel" json:"fuel,omitempty"`
	// If set `true`, you may provide [amount](https://docs.stripe.com/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization.
	IsAmountControllable *bool `form:"is_amount_controllable" json:"is_amount_controllable,omitempty"`
	// The total amount to attempt to authorize. This amount is in the provided merchant currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	MerchantAmount *int64 `form:"merchant_amount" json:"merchant_amount,omitempty"`
	// The currency of the authorization. If not provided, defaults to the currency of the card. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	MerchantCurrency *string `form:"merchant_currency" json:"merchant_currency,omitempty"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantData *TestHelpersIssuingAuthorizationCreateMerchantDataParams `form:"merchant_data" json:"merchant_data,omitempty"`
	// Details about the authorization, such as identifiers, set by the card network.
	NetworkData *TestHelpersIssuingAuthorizationCreateNetworkDataParams `form:"network_data" json:"network_data,omitempty"`
	// Stripe's assessment of the fraud risk for this authorization.
	RiskAssessment *TestHelpersIssuingAuthorizationCreateRiskAssessmentParams `form:"risk_assessment" json:"risk_assessment,omitempty"`
	// Verifications that Stripe performed on information that the cardholder provided to the merchant.
	VerificationData *TestHelpersIssuingAuthorizationCreateVerificationDataParams `form:"verification_data" json:"verification_data,omitempty"`
	// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`. Will populate as `null` when no digital wallet was utilized.
	Wallet *string `form:"wallet" json:"wallet,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingAuthorizationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
