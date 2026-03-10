//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The method used for authorization.
type RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod string

// List of values that RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod can take
const (
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodChip        RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "chip"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodContactless RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "contactless"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodKeyedIn     RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "keyed_in"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodOnline      RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "online"
	RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethodSwipe       RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod = "swipe"
)

// The card entry mode.
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

// The point of sale condition.
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

// The type of card (physical or virtual).
type RadarIssuingAuthorizationEvaluationCardDetailsCardType string

// List of values that RadarIssuingAuthorizationEvaluationCardDetailsCardType can take
const (
	RadarIssuingAuthorizationEvaluationCardDetailsCardTypePhysical RadarIssuingAuthorizationEvaluationCardDetailsCardType = "physical"
	RadarIssuingAuthorizationEvaluationCardDetailsCardTypeVirtual  RadarIssuingAuthorizationEvaluationCardDetailsCardType = "virtual"
)

// The card network that processed the authorization.
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

// The wallet provider, if applicable.
type RadarIssuingAuthorizationEvaluationTokenDetailsWallet string

// List of values that RadarIssuingAuthorizationEvaluationTokenDetailsWallet can take
const (
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletApplePay   RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "apple_pay"
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletGooglePay  RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "google_pay"
	RadarIssuingAuthorizationEvaluationTokenDetailsWalletSamsungPay RadarIssuingAuthorizationEvaluationTokenDetailsWallet = "samsung_pay"
)

// The result of the 3D Secure verification.
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
	// The authorization amount in the smallest currency unit.
	Amount *int64 `form:"amount"`
	// The method used for authorization.
	AuthorizationMethod *string `form:"authorization_method"`
	// Three-letter ISO currency code in lowercase.
	Currency *string `form:"currency"`
	// The card entry mode.
	EntryMode *string `form:"entry_mode"`
	// The raw code for the card entry mode.
	EntryModeRawCode *string `form:"entry_mode_raw_code"`
	// The time when the authorization was initiated (Unix timestamp).
	InitiatedAt *int64 `form:"initiated_at"`
	// The point of sale condition.
	PointOfSaleCondition *string `form:"point_of_sale_condition"`
	// The raw code for the point of sale condition.
	PointOfSaleConditionRawCode *string `form:"point_of_sale_condition_raw_code"`
	// External reference for the authorization.
	Reference *string `form:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCardDetailsParams struct {
	// Bank Identification Number (BIN) of the card.
	Bin *string `form:"bin"`
	// Two-letter ISO country code of the card's issuing bank.
	BinCountry *string `form:"bin_country"`
	// The type of card (physical or virtual).
	CardType *string `form:"card_type"`
	// The time when the card was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// Last 4 digits of the card number.
	Last4 *string `form:"last4"`
	// External reference for the card.
	Reference *string `form:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCardholderDetailsParams struct {
	// The time when the cardholder was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// External reference for the cardholder.
	Reference *string `form:"reference"`
}

// Details about the merchant where the authorization occurred.
type RadarIssuingAuthorizationEvaluationMerchantDetailsParams struct {
	// Merchant Category Code (MCC).
	CategoryCode *string `form:"category_code"`
	// Two-letter ISO country code of the merchant.
	Country *string `form:"country"`
	// Name of the merchant.
	Name *string `form:"name"`
	// Network merchant identifier.
	NetworkID *string `form:"network_id"`
	// Terminal identifier.
	TerminalID *string `form:"terminal_id"`
}

// Details about the card network processing.
type RadarIssuingAuthorizationEvaluationNetworkDetailsParams struct {
	// The acquiring institution identifier.
	AcquiringInstitutionID *string `form:"acquiring_institution_id"`
	// The card network that routed the authorization.
	RoutedNetwork *string `form:"routed_network"`
}

// Details about the token, if a tokenized payment method was used.
type RadarIssuingAuthorizationEvaluationTokenDetailsParams struct {
	// The time when the token was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// External reference for the token.
	Reference *string `form:"reference"`
	// The wallet provider for the tokenized payment method.
	Wallet *string `form:"wallet"`
}

// Details about verification checks performed.
type RadarIssuingAuthorizationEvaluationVerificationDetailsParams struct {
	// The result of 3D Secure verification.
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
	// Details about the merchant where the authorization occurred.
	MerchantDetails *RadarIssuingAuthorizationEvaluationMerchantDetailsParams `form:"merchant_details"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the card network processing.
	NetworkDetails *RadarIssuingAuthorizationEvaluationNetworkDetailsParams `form:"network_details"`
	// Details about the token, if a tokenized payment method was used.
	TokenDetails *RadarIssuingAuthorizationEvaluationTokenDetailsParams `form:"token_details"`
	// Details about verification checks performed.
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
	// The authorization amount in the smallest currency unit.
	Amount *int64 `form:"amount"`
	// The method used for authorization.
	AuthorizationMethod *string `form:"authorization_method"`
	// Three-letter ISO currency code in lowercase.
	Currency *string `form:"currency"`
	// The card entry mode.
	EntryMode *string `form:"entry_mode"`
	// The raw code for the card entry mode.
	EntryModeRawCode *string `form:"entry_mode_raw_code"`
	// The time when the authorization was initiated (Unix timestamp).
	InitiatedAt *int64 `form:"initiated_at"`
	// The point of sale condition.
	PointOfSaleCondition *string `form:"point_of_sale_condition"`
	// The raw code for the point of sale condition.
	PointOfSaleConditionRawCode *string `form:"point_of_sale_condition_raw_code"`
	// External reference for the authorization.
	Reference *string `form:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCreateCardDetailsParams struct {
	// Bank Identification Number (BIN) of the card.
	Bin *string `form:"bin"`
	// Two-letter ISO country code of the card's issuing bank.
	BinCountry *string `form:"bin_country"`
	// The type of card (physical or virtual).
	CardType *string `form:"card_type"`
	// The time when the card was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// Last 4 digits of the card number.
	Last4 *string `form:"last4"`
	// External reference for the card.
	Reference *string `form:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCreateCardholderDetailsParams struct {
	// The time when the cardholder was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// External reference for the cardholder.
	Reference *string `form:"reference"`
}

// Details about the merchant where the authorization occurred.
type RadarIssuingAuthorizationEvaluationCreateMerchantDetailsParams struct {
	// Merchant Category Code (MCC).
	CategoryCode *string `form:"category_code"`
	// Two-letter ISO country code of the merchant.
	Country *string `form:"country"`
	// Name of the merchant.
	Name *string `form:"name"`
	// Network merchant identifier.
	NetworkID *string `form:"network_id"`
	// Terminal identifier.
	TerminalID *string `form:"terminal_id"`
}

// Details about the card network processing.
type RadarIssuingAuthorizationEvaluationCreateNetworkDetailsParams struct {
	// The acquiring institution identifier.
	AcquiringInstitutionID *string `form:"acquiring_institution_id"`
	// The card network that routed the authorization.
	RoutedNetwork *string `form:"routed_network"`
}

// Details about the token, if a tokenized payment method was used.
type RadarIssuingAuthorizationEvaluationCreateTokenDetailsParams struct {
	// The time when the token was created (Unix timestamp).
	CreatedAt *int64 `form:"created_at"`
	// External reference for the token.
	Reference *string `form:"reference"`
	// The wallet provider for the tokenized payment method.
	Wallet *string `form:"wallet"`
}

// Details about verification checks performed.
type RadarIssuingAuthorizationEvaluationCreateVerificationDetailsParams struct {
	// The result of 3D Secure verification.
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
	// Details about the merchant where the authorization occurred.
	MerchantDetails *RadarIssuingAuthorizationEvaluationCreateMerchantDetailsParams `form:"merchant_details"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the card network processing.
	NetworkDetails *RadarIssuingAuthorizationEvaluationCreateNetworkDetailsParams `form:"network_details"`
	// Details about the token, if a tokenized payment method was used.
	TokenDetails *RadarIssuingAuthorizationEvaluationCreateTokenDetailsParams `form:"token_details"`
	// Details about verification checks performed.
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

// Details about the authorization transaction.
type RadarIssuingAuthorizationEvaluationAuthorizationDetails struct {
	// The authorization amount in the smallest currency unit.
	Amount int64 `json:"amount"`
	// The method used for authorization.
	AuthorizationMethod RadarIssuingAuthorizationEvaluationAuthorizationDetailsAuthorizationMethod `json:"authorization_method"`
	// Three-letter ISO currency code in lowercase.
	Currency Currency `json:"currency"`
	// The card entry mode.
	EntryMode RadarIssuingAuthorizationEvaluationAuthorizationDetailsEntryMode `json:"entry_mode"`
	// The raw code for the card entry mode.
	EntryModeRawCode string `json:"entry_mode_raw_code"`
	// The time when the authorization was initiated.
	InitiatedAt int64 `json:"initiated_at"`
	// The point of sale condition.
	PointOfSaleCondition RadarIssuingAuthorizationEvaluationAuthorizationDetailsPointOfSaleCondition `json:"point_of_sale_condition"`
	// The raw code for the point of sale condition.
	PointOfSaleConditionRawCode string `json:"point_of_sale_condition_raw_code"`
	// External reference for the authorization.
	Reference string `json:"reference"`
}

// Details about the card used in the authorization.
type RadarIssuingAuthorizationEvaluationCardDetails struct {
	// The Bank Identification Number (BIN) of the card.
	Bin string `json:"bin"`
	// The country code associated with the card BIN.
	BinCountry string `json:"bin_country"`
	// The type of card (physical or virtual).
	CardType RadarIssuingAuthorizationEvaluationCardDetailsCardType `json:"card_type"`
	// The time when the card was created.
	CreatedAt int64 `json:"created_at"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
	// External reference for the card.
	Reference string `json:"reference"`
}

// Details about the cardholder.
type RadarIssuingAuthorizationEvaluationCardholderDetails struct {
	// The time when the cardholder was created.
	CreatedAt int64 `json:"created_at"`
	// External reference for the cardholder.
	Reference string `json:"reference"`
}

// Details about the merchant where the authorization occurred.
type RadarIssuingAuthorizationEvaluationMerchantDetails struct {
	// The merchant category code (MCC).
	CategoryCode string `json:"category_code"`
	// The merchant country code.
	Country string `json:"country"`
	// The merchant name.
	Name string `json:"name"`
	// The merchant identifier from the card network.
	NetworkID string `json:"network_id"`
	// The terminal identifier.
	TerminalID string `json:"terminal_id"`
}

// Details about the card network processing.
type RadarIssuingAuthorizationEvaluationNetworkDetails struct {
	// The acquiring institution identifier.
	AcquiringInstitutionID string `json:"acquiring_institution_id"`
	// The card network that processed the authorization.
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

// Details about the token, if a tokenized payment method was used.
type RadarIssuingAuthorizationEvaluationTokenDetails struct {
	// The time when the token was created.
	CreatedAt int64 `json:"created_at"`
	// External reference for the token.
	Reference string `json:"reference"`
	// The wallet provider, if applicable.
	Wallet RadarIssuingAuthorizationEvaluationTokenDetailsWallet `json:"wallet"`
}

// Details about verification checks performed.
type RadarIssuingAuthorizationEvaluationVerificationDetails struct {
	// The result of the 3D Secure verification.
	ThreeDSecureResult RadarIssuingAuthorizationEvaluationVerificationDetailsThreeDSecureResult `json:"three_d_secure_result"`
}

// Authorization Evaluations represent fraud risk assessments for Issuing card authorizations. They include fraud risk signals and contextual details about the authorization.
type RadarIssuingAuthorizationEvaluation struct {
	APIResource
	// Details about the authorization transaction.
	AuthorizationDetails *RadarIssuingAuthorizationEvaluationAuthorizationDetails `json:"authorization_details"`
	// Details about the card used in the authorization.
	CardDetails *RadarIssuingAuthorizationEvaluationCardDetails `json:"card_details"`
	// Details about the cardholder.
	CardholderDetails *RadarIssuingAuthorizationEvaluationCardholderDetails `json:"cardholder_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Details about the merchant where the authorization occurred.
	MerchantDetails *RadarIssuingAuthorizationEvaluationMerchantDetails `json:"merchant_details"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Details about the card network processing.
	NetworkDetails *RadarIssuingAuthorizationEvaluationNetworkDetails `json:"network_details"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Collection of fraud risk signals for this authorization evaluation.
	Signals *RadarIssuingAuthorizationEvaluationSignals `json:"signals"`
	// Details about the token, if a tokenized payment method was used.
	TokenDetails *RadarIssuingAuthorizationEvaluationTokenDetails `json:"token_details"`
	// Details about verification checks performed.
	VerificationDetails *RadarIssuingAuthorizationEvaluationVerificationDetails `json:"verification_details"`
}
