//
//
// File generated from our OpenAPI spec
//
//

package stripe

// How the card details were provided.
type RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod string

// List of values that RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod can take
const (
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodChip        RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "chip"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodContactless RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "contactless"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodKeyedIn     RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "keyed_in"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodOnline      RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "online"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodSwipe       RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "swipe"
)

// Defines how the card's information was entered for the authorization.
type RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode string

// List of values that RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode can take
const (
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeContactless           RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "contactless"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeContactlessMagstripe  RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "contactless_magstripe"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeCredentialOnFile      RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "credential_on_file"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeIntegratedCircuitCard RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "integrated_circuit_card"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeMagstripe             RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "magstripe"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeMagstripeNoCvv        RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "magstripe_no_cvv"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeManual                RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "manual"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeOther                 RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "other"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryModeUnknown               RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode = "unknown"
)

// Defines how the card was read at the point of sale.
type RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition string

// List of values that RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition can take
const (
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionAccountVerification RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "account_verification"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionCardNotPresent      RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "card_not_present"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionCardPresent         RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "card_present"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionECommerce           RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "e_commerce"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionKeyEnteredPos       RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "key_entered_pos"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionMissing             RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "missing"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionMOTO                RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "moto"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionOther               RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "other"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionPINEntered          RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "pin_entered"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleConditionRecurring           RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition = "recurring"
)

// The type of the card.
type RadarIssuingAuthorizationEvaluationCardDetailsCardType string

// List of values that RadarIssuingAuthorizationEvaluationCardDetailsCardType can take
const (
	RadarIssuingAuthorizationEvaluationCardDetailsCardTypePhysical RadarIssuingAuthorizationEvaluationCardDetailsCardType = "physical"
	RadarIssuingAuthorizationEvaluationCardDetailsCardTypeVirtual  RadarIssuingAuthorizationEvaluationCardDetailsCardType = "virtual"
)

// The card network over which Stripe received the authorization.
type RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork string

// List of values that RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork can take
const (
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkCirrus     RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "cirrus"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkInterlink  RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "interlink"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkMaestro    RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "maestro"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkMastercard RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "mastercard"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkOther      RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "other"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkPlus       RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "plus"
	RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetworkVisa       RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork = "visa"
)

// Risk level, based on the score.
type RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel string

// List of values that RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel can take
const (
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelElevated    RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "elevated"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelHighest     RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "highest"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelLow         RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "low"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelNormal      RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "normal"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelNotAssessed RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "not_assessed"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevelUnknown     RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel = "unknown"
)

// The status of this signal.
type RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatus string

// List of values that RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatus can take
const (
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatusError   RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatus = "error"
	RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatusSuccess RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatus = "success"
)

// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`.
type RadarIssuingAuthorizationEvaluationTokenDetailsWallet string

// List of values that RadarIssuingAuthorizationEvaluationTokenDetailsWallet can take
const (
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletApplePay   RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "apple_pay"
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletGooglePay  RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "google_pay"
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletSamsungPay RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "samsung_pay"
)

// The outcome of the 3D Secure authentication request.
type RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult string

// List of values that RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult can take
const (
	RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResultAttemptAcknowledged RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult = "attempt_acknowledged"
	RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResultAuthenticated       RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult = "authenticated"
	RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResultExempted            RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult = "exempted"
	RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResultFailed              RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult = "failed"
	RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResultRequired            RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult = "required"
)

// Details about the authorization.
type RadarIssuingAuthorizationEvaluationAuthorizationDetailsParams struct {
	// The total amount of the authorization in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// How the card details were provided.
	AuthorizationMethod *string `form:"authorization_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Defines how the card's information was entered for the authorization.
	EntryMode *string `form:"entry_mode"`
	// Raw code indicating the entry mode from the network message.
	EntryModeRawCode *string `form:"entry_mode_raw_code"`
	// The timestamp of the authorization initiated in seconds.
	InitiatedAt *int64 `form:"initiated_at"`
	// Defines how the card was read at the point of sale.
	PointOfSaleCondition *string `form:"point_of_sale_condition"`
	// Raw code indicating the point of sale condition from the network message.
	PointOfSaleConditionRawCode *string `form:"point_of_sale_condition_raw_code"`
	// User's specified unique ID for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCardDetailsParams struct {
	// The Bank Identification Number (BIN) of the card.
	Bin *string `form:"bin"`
	// The two-letter country code of the BIN issuer.
	BinCountry *string `form:"bin_country"`
	// The type of the card.
	CardType *string `form:"card_type"`
	// The timestamp when the card was created.
	CreatedAt *int64 `form:"created_at"`
	// The last 4 digits of the card number.
	Last4 *string `form:"last4"`
	// User's specified unique ID of the card for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCardholderDetailsParams struct {
	// The timestamp when the cardholder was created.
	CreatedAt *int64 `form:"created_at"`
	// User's specified unique ID of the cardholder for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type RadarIssuingAuthorizationEvaluationMerchantDetailsParams struct {
	// The merchant category code for the seller's business.
	CategoryCode *string `form:"category_code"`
	// Country where the seller is located.
	Country *string `form:"country"`
	// Name of the seller.
	Name *string `form:"name"`
	// Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant.
	NetworkID *string `form:"network_id"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id"`
}

// Details about the authorization, such as identifiers, set by the card network.
type RadarIssuingAuthorizationEvaluationNetworkDetailsParams struct {
	// Identifier assigned to the acquirer by the card network. Sometimes this value is not provided by the network; in this case, the value will be null.
	AcquiringInstitutionID *string `form:"acquiring_institution_id"`
	// The card network over which Stripe received the authorization.
	RoutedNetwork *string `form:"routed_network"`
}

// Details about the token, if a tokenized payment method was used for the authorization.
type RadarIssuingAuthorizationEvaluationTokenDetailsParams struct {
	// The timestamp when the network token was created.
	CreatedAt *int64 `form:"created_at"`
	// User's specified unique ID of the card token for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
	// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`.
	Wallet *string `form:"wallet"`
}

// Details about verification data for the authorization.
type RadarIssuingAuthorizationEvaluationVerificationDetailsParams struct {
	// The outcome of the 3D Secure authentication request.
	ThreeDSecureResult *string `form:"three_d_secure_result"`
}

// Request a fraud risk assessment from Stripe for an Issuing card authorization.
type RadarIssuingAuthorizationEvaluationParams struct {
	Params `form:"*"`
	// Details about the authorization.
	AuthorizationDetails *RadarIssuingAuthorizationEvaluationAuthorizationDetailsParams `form:"authorization_details"`
	// Details about the card used in the authorization.
	CardDetails *RadarIssuingAuthorizationEvaluationCardDetailsParams `form:"card_details"`
	// Details about the cardholder.
	CardholderDetails *RadarIssuingAuthorizationEvaluationCardholderDetailsParams `form:"cardholder_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantDetails *RadarIssuingAuthorizationEvaluationMerchantDetailsParams `form:"merchant_details"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the authorization, such as identifiers, set by the card network.
	NetworkDetails *RadarIssuingAuthorizationEvaluationNetworkDetailsParams `form:"network_details"`
	// Details about the token, if a tokenized payment method was used for the authorization.
	TokenDetails *RadarIssuingAuthorizationEvaluationTokenDetailsParams `form:"token_details"`
	// Details about verification data for the authorization.
	VerificationDetails *RadarIssuingAuthorizationEvaluationVerificationDetailsParams `form:"verification_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarIssuingAuthorizationEvaluationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarIssuingAuthorizationEvaluationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the authorization.
type RadarIssuingAuthorizationEvaluationCreateAuthorizationDetailsParams struct {
	// The total amount of the authorization in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// How the card details were provided.
	AuthorizationMethod *string `form:"authorization_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Defines how the card's information was entered for the authorization.
	EntryMode *string `form:"entry_mode"`
	// Raw code indicating the entry mode from the network message.
	EntryModeRawCode *string `form:"entry_mode_raw_code"`
	// The timestamp of the authorization initiated in seconds.
	InitiatedAt *int64 `form:"initiated_at"`
	// Defines how the card was read at the point of sale.
	PointOfSaleCondition *string `form:"point_of_sale_condition"`
	// Raw code indicating the point of sale condition from the network message.
	PointOfSaleConditionRawCode *string `form:"point_of_sale_condition_raw_code"`
	// User's specified unique ID for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCreateCardDetailsParams struct {
	// The Bank Identification Number (BIN) of the card.
	Bin *string `form:"bin"`
	// The two-letter country code of the BIN issuer.
	BinCountry *string `form:"bin_country"`
	// The type of the card.
	CardType *string `form:"card_type"`
	// The timestamp when the card was created.
	CreatedAt *int64 `form:"created_at"`
	// The last 4 digits of the card number.
	Last4 *string `form:"last4"`
	// User's specified unique ID of the card for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCreateCardholderDetailsParams struct {
	// The timestamp when the cardholder was created.
	CreatedAt *int64 `form:"created_at"`
	// User's specified unique ID of the cardholder for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
}

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type RadarIssuingAuthorizationEvaluationCreateMerchantDetailsParams struct {
	// The merchant category code for the seller's business.
	CategoryCode *string `form:"category_code"`
	// Country where the seller is located.
	Country *string `form:"country"`
	// Name of the seller.
	Name *string `form:"name"`
	// Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant.
	NetworkID *string `form:"network_id"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID *string `form:"terminal_id"`
}

// Details about the authorization, such as identifiers, set by the card network.
type RadarIssuingAuthorizationEvaluationCreateNetworkDetailsParams struct {
	// Identifier assigned to the acquirer by the card network. Sometimes this value is not provided by the network; in this case, the value will be null.
	AcquiringInstitutionID *string `form:"acquiring_institution_id"`
	// The card network over which Stripe received the authorization.
	RoutedNetwork *string `form:"routed_network"`
}

// Details about the token, if a tokenized payment method was used for the authorization.
type RadarIssuingAuthorizationEvaluationCreateTokenDetailsParams struct {
	// The timestamp when the network token was created.
	CreatedAt *int64 `form:"created_at"`
	// User's specified unique ID of the card token for this authorization attempt (e.g., RRN or internal reference).
	Reference *string `form:"reference"`
	// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`.
	Wallet *string `form:"wallet"`
}

// Details about verification data for the authorization.
type RadarIssuingAuthorizationEvaluationCreateVerificationDetailsParams struct {
	// The outcome of the 3D Secure authentication request.
	ThreeDSecureResult *string `form:"three_d_secure_result"`
}

// Request a fraud risk assessment from Stripe for an Issuing card authorization.
type RadarIssuingAuthorizationEvaluationCreateParams struct {
	Params `form:"*"`
	// Details about the authorization.
	AuthorizationDetails *RadarIssuingAuthorizationEvaluationCreateAuthorizationDetailsParams `form:"authorization_details"`
	// Details about the card used in the authorization.
	CardDetails *RadarIssuingAuthorizationEvaluationCreateCardDetailsParams `form:"card_details"`
	// Details about the cardholder.
	CardholderDetails *RadarIssuingAuthorizationEvaluationCreateCardholderDetailsParams `form:"cardholder_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantDetails *RadarIssuingAuthorizationEvaluationCreateMerchantDetailsParams `form:"merchant_details"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the authorization, such as identifiers, set by the card network.
	NetworkDetails *RadarIssuingAuthorizationEvaluationCreateNetworkDetailsParams `form:"network_details"`
	// Details about the token, if a tokenized payment method was used for the authorization.
	TokenDetails *RadarIssuingAuthorizationEvaluationCreateTokenDetailsParams `form:"token_details"`
	// Details about verification data for the authorization.
	VerificationDetails *RadarIssuingAuthorizationEvaluationCreateVerificationDetailsParams `form:"verification_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarIssuingAuthorizationEvaluationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarIssuingAuthorizationEvaluationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the authorization.
type RadarIssuingAuthorizationEvaluationAuthorizationDetails struct {
	// The total amount of the authorization in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// How the card details were provided.
	AuthorizationMethod RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod `json:"authorization_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Defines how the card's information was entered for the authorization.
	EntryMode RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode `json:"entry_mode"`
	// Raw code indicating the entry mode from the network message.
	EntryModeRawCode string `json:"entry_mode_raw_code"`
	// The timestamp of the authorization initiated in seconds.
	InitiatedAt int64 `json:"initiated_at"`
	// Defines how the card was read at the point of sale.
	PointOfSaleCondition RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition `json:"point_of_sale_condition"`
	// Raw code indicating the point of sale condition from the network message.
	PointOfSaleConditionRawCode string `json:"point_of_sale_condition_raw_code"`
	// User's specified unique ID for this authorization attempt (e.g., RRN or internal reference).
	Reference string `json:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCardDetails struct {
	// The Bank Identification Number (BIN) of the card.
	Bin string `json:"bin"`
	// The two-letter country code of the BIN issuer.
	BinCountry string `json:"bin_country"`
	// The type of the card.
	CardType RadarIssuingAuthorizationEvaluationCardDetailsCardType `json:"card_type"`
	// The timestamp when the card was created.
	CreatedAt int64 `json:"created_at"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
	// User's specified unique ID of the card for this authorization attempt (e.g., RRN or internal reference).
	Reference string `json:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCardholderDetails struct {
	// The timestamp when the cardholder was created.
	CreatedAt int64 `json:"created_at"`
	// User's specified unique ID of the cardholder for this authorization attempt (e.g., RRN or internal reference).
	Reference string `json:"reference"`
}

// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
type RadarIssuingAuthorizationEvaluationMerchantDetails struct {
	// The merchant category code for the seller's business.
	CategoryCode string `json:"category_code"`
	// Country where the seller is located.
	Country string `json:"country"`
	// Name of the seller.
	Name string `json:"name"`
	// Identifier assigned to the seller by the card network. Different card networks may assign different network_id fields to the same merchant.
	NetworkID string `json:"network_id"`
	// An ID assigned by the seller to the location of the sale.
	TerminalID string `json:"terminal_id"`
}

// Details about the authorization, such as identifiers, set by the card network.
type RadarIssuingAuthorizationEvaluationNetworkDetails struct {
	// Identifier assigned to the acquirer by the card network. Sometimes this value is not provided by the network; in this case, the value will be null.
	AcquiringInstitutionID string `json:"acquiring_institution_id"`
	// The card network over which Stripe received the authorization.
	RoutedNetwork RadarIssuingAuthorizationEvaluationNetworkDetailsRoutedNetwork `json:"routed_network"`
}

// Signal evaluation data, including score and level.
type RadarIssuingAuthorizationEvaluationSignalsFraudRiskData struct {
	// The time when this signal was evaluated.
	EvaluatedAt int64 `json:"evaluated_at"`
	// Risk level, based on the score.
	Level RadarIssuingAuthorizationEvaluationSignalsFraudRiskDataLevel `json:"level"`
	// Fraud risk score for this evaluation. Score in the range [0,100], formatted as a 2 decimal float, with higher scores indicating a higher likelihood of fraud.
	Score float64 `json:"score"`
}

// A fraud risk signal with status, error, and data fields.
type RadarIssuingAuthorizationEvaluationSignalsFraudRisk struct {
	// Signal evaluation data, including score and level.
	Data *RadarIssuingAuthorizationEvaluationSignalsFraudRiskData `json:"data"`
	// Details of an error that prevented reporting this signal.
	Error map[string]string `json:"error"`
	// The status of this signal.
	Status RadarIssuingAuthorizationEvaluationSignalsFraudRiskStatus `json:"status"`
}

// Collection of fraud risk signals for this authorization evaluation.
type RadarIssuingAuthorizationEvaluationSignals struct {
	// A fraud risk signal with status, error, and data fields.
	FraudRisk *RadarIssuingAuthorizationEvaluationSignalsFraudRisk `json:"fraud_risk"`
}

// Details about the token, if a tokenized payment method was used for the authorization.
type RadarIssuingAuthorizationEvaluationTokenDetails struct {
	// The timestamp when the network token was created.
	CreatedAt int64 `json:"created_at"`
	// User's specified unique ID of the card token for this authorization attempt (e.g., RRN or internal reference).
	Reference string `json:"reference"`
	// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`.
	Wallet RadarIssuingAuthorizationEvaluationTokenDetailsWallet `json:"wallet"`
}

// Details about verification data for the authorization.
type RadarIssuingAuthorizationEvaluationVerificationDetails struct {
	// The outcome of the 3D Secure authentication request.
	ThreeDSecureResult RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult `json:"three_d_secure_result"`
}

// Authorization Evaluations represent fraud risk assessments for Issuing card authorizations. They include fraud risk signals and contextual details about the authorization.
type RadarIssuingAuthorizationEvaluation struct {
	APIResource
	// Details about the authorization.
	AuthorizationDetails *RadarIssuingAuthorizationEvaluationAuthorizationDetails `json:"authorization_details"`
	// Details about the card used in the authorization.
	CardDetails *RadarIssuingAuthorizationEvaluationCardDetails `json:"card_details"`
	// Details about the cardholder.
	CardholderDetails *RadarIssuingAuthorizationEvaluationCardholderDetails `json:"cardholder_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Details about the seller (grocery store, e-commerce website, etc.) where the card authorization happened.
	MerchantDetails *RadarIssuingAuthorizationEvaluationMerchantDetails `json:"merchant_details"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Details about the authorization, such as identifiers, set by the card network.
	NetworkDetails *RadarIssuingAuthorizationEvaluationNetworkDetails `json:"network_details"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Collection of fraud risk signals for this authorization evaluation.
	Signals *RadarIssuingAuthorizationEvaluationSignals `json:"signals"`
	// Details about the token, if a tokenized payment method was used for the authorization.
	TokenDetails *RadarIssuingAuthorizationEvaluationTokenDetails `json:"token_details"`
	// Details about verification data for the authorization.
	VerificationDetails *RadarIssuingAuthorizationEvaluationVerificationDetails `json:"verification_details"`
}
