package stripe

import "encoding/json"

// ErrorType is the list of allowed values for the error's type.
type ErrorType string

// List of values that ErrorType can take.
const (
	ErrorTypeAPI            ErrorType = "api_error"
	ErrorTypeAPIConnection  ErrorType = "api_connection_error"
	ErrorTypeAuthentication ErrorType = "authentication_error"
	ErrorTypeCard           ErrorType = "card_error"
	ErrorTypeInvalidRequest ErrorType = "invalid_request_error"
	ErrorTypePermission     ErrorType = "more_permissions_required"
	ErrorTypeRateLimit      ErrorType = "rate_limit_error"
)

// ErrorCode is the list of allowed values for the error's code.
type ErrorCode string

// DeclineCode is the list of reasons provided by card issuers for decline of payment.
type DeclineCode string

// List of values that ErrorCode can take.
const (
	ErrorCodeAccountAlreadyExists                   ErrorCode = "account_already_exists"
	ErrorCodeAccountCountryInvalidAddress           ErrorCode = "account_country_invalid_address"
	ErrorCodeAccountInvalid                         ErrorCode = "account_invalid"
	ErrorCodeAccountNumberInvalid                   ErrorCode = "account_number_invalid"
	ErrorCodeAlipayUpgradeRequired                  ErrorCode = "alipay_upgrade_required"
	ErrorCodeAmountTooLarge                         ErrorCode = "amount_too_large"
	ErrorCodeAmountTooSmall                         ErrorCode = "amount_too_small"
	ErrorCodeAPIKeyExpired                          ErrorCode = "api_key_expired"
	ErrorCodeAuthenticationRequired                 ErrorCode = "authentication_required"
	ErrorCodeBalanceInsufficient                    ErrorCode = "balance_insufficient"
	ErrorCodeBankAccountExists                      ErrorCode = "bank_account_exists"
	ErrorCodeBankAccountUnusable                    ErrorCode = "bank_account_unusable"
	ErrorCodeBankAccountUnverified                  ErrorCode = "bank_account_unverified"
	ErrorCodeBitcoinUpgradeRequired                 ErrorCode = "bitcoin_upgrade_required"
	ErrorCodeCardDeclined                           ErrorCode = "card_declined"
	ErrorCodeChargeAlreadyCaptured                  ErrorCode = "charge_already_captured"
	ErrorCodeChargeAlreadyRefunded                  ErrorCode = "charge_already_refunded"
	ErrorCodeChargeDisputed                         ErrorCode = "charge_disputed"
	ErrorCodeChargeExceedsSourceLimit               ErrorCode = "charge_exceeds_source_limit"
	ErrorCodeChargeExpiredForCapture                ErrorCode = "charge_expired_for_capture"
	ErrorCodeCountryUnsupported                     ErrorCode = "country_unsupported"
	ErrorCodeCouponExpired                          ErrorCode = "coupon_expired"
	ErrorCodeCustomerMaxSubscriptions               ErrorCode = "customer_max_subscriptions"
	ErrorCodeEmailInvalid                           ErrorCode = "email_invalid"
	ErrorCodeExpiredCard                            ErrorCode = "expired_card"
	ErrorCodeIdempotencyKeyInUse                    ErrorCode = "idempotency_key_in_use"
	ErrorCodeIncorrectAddress                       ErrorCode = "incorrect_address"
	ErrorCodeIncorrectCVC                           ErrorCode = "incorrect_cvc"
	ErrorCodeIncorrectNumber                        ErrorCode = "incorrect_number"
	ErrorCodeIncorrectZip                           ErrorCode = "incorrect_zip"
	ErrorCodeInstantPayoutsUnsupported              ErrorCode = "instant_payouts_unsupported"
	ErrorCodeInvalidCardType                        ErrorCode = "invalid_card_type"
	ErrorCodeInvalidChargeAmount                    ErrorCode = "invalid_charge_amount"
	ErrorCodeInvalidCVC                             ErrorCode = "invalid_cvc"
	ErrorCodeInvalidExpiryMonth                     ErrorCode = "invalid_expiry_month"
	ErrorCodeInvalidExpiryYear                      ErrorCode = "invalid_expiry_year"
	ErrorCodeInvalidNumber                          ErrorCode = "invalid_number"
	ErrorCodeInvalidSourceUsage                     ErrorCode = "invalid_source_usage"
	ErrorCodeInvoiceNoCustomerLineItems             ErrorCode = "invoice_no_customer_line_items"
	ErrorCodeInvoiceNoSubscriptionLineItems         ErrorCode = "invoice_no_subscription_line_items"
	ErrorCodeInvoiceNotEditable                     ErrorCode = "invoice_not_editable"
	ErrorCodeInvoiceUpcomingNone                    ErrorCode = "invoice_upcoming_none"
	ErrorCodeLivemodeMismatch                       ErrorCode = "livemode_mismatch"
	ErrorCodeLockTimeout                            ErrorCode = "lock_timeout"
	ErrorCodeMissing                                ErrorCode = "missing"
	ErrorCodeNotAllowedOnStandardAccount            ErrorCode = "not_allowed_on_standard_account"
	ErrorCodeOrderCreationFailed                    ErrorCode = "order_creation_failed"
	ErrorCodeOrderRequiredSettings                  ErrorCode = "order_required_settings"
	ErrorCodeOrderStatusInvalid                     ErrorCode = "order_status_invalid"
	ErrorCodeOrderUpstreamTimeout                   ErrorCode = "order_upstream_timeout"
	ErrorCodeOutOfInventory                         ErrorCode = "out_of_inventory"
	ErrorCodeParameterInvalidEmpty                  ErrorCode = "parameter_invalid_empty"
	ErrorCodeParameterInvalidInteger                ErrorCode = "parameter_invalid_integer"
	ErrorCodeParameterInvalidStringBlank            ErrorCode = "parameter_invalid_string_blank"
	ErrorCodeParameterInvalidStringEmpty            ErrorCode = "parameter_invalid_string_empty"
	ErrorCodeParameterMissing                       ErrorCode = "parameter_missing"
	ErrorCodeParameterUnknown                       ErrorCode = "parameter_unknown"
	ErrorCodeParametersExclusive                    ErrorCode = "parameters_exclusive"
	ErrorCodePaymentIntentAuthenticationFailure     ErrorCode = "payment_intent_authentication_failure"
	ErrorCodePaymentIntentIncompatiblePaymentMethod ErrorCode = "payment_intent_incompatible_payment_method"
	ErrorCodePaymentIntentInvalidParameter          ErrorCode = "payment_intent_invalid_parameter"
	ErrorCodePaymentIntentPaymentAttemptFailed      ErrorCode = "payment_intent_payment_attempt_failed"
	ErrorCodePaymentIntentUnexpectedState           ErrorCode = "payment_intent_unexpected_state"
	ErrorCodePaymentMethodUnactivated               ErrorCode = "payment_method_unactivated"
	ErrorCodePaymentMethodUnexpectedState           ErrorCode = "payment_method_unexpected_state"
	ErrorCodePayoutsNotAllowed                      ErrorCode = "payouts_not_allowed"
	ErrorCodePlatformAPIKeyExpired                  ErrorCode = "platform_api_key_expired"
	ErrorCodePostalCodeInvalid                      ErrorCode = "postal_code_invalid"
	ErrorCodeProcessingError                        ErrorCode = "processing_error"
	ErrorCodeProductInactive                        ErrorCode = "product_inactive"
	ErrorCodeRateLimit                              ErrorCode = "rate_limit"
	ErrorCodeResourceAlreadyExists                  ErrorCode = "resource_already_exists"
	ErrorCodeResourceMissing                        ErrorCode = "resource_missing"
	ErrorCodeRoutingNumberInvalid                   ErrorCode = "routing_number_invalid"
	ErrorCodeSecretKeyRequired                      ErrorCode = "secret_key_required"
	ErrorCodeSepaUnsupportedAccount                 ErrorCode = "sepa_unsupported_account"
	ErrorCodeSetupAttemptFailed                     ErrorCode = "setup_attempt_failed"
	ErrorCodeSetupIntentAuthenticationFailure       ErrorCode = "setup_intent_authentication_failure"
	ErrorCodeSetupIntentUnexpectedState             ErrorCode = "setup_intent_unexpected_state"
	ErrorCodeShippingCalculationFailed              ErrorCode = "shipping_calculation_failed"
	ErrorCodeSkuInactive                            ErrorCode = "sku_inactive"
	ErrorCodeStateUnsupported                       ErrorCode = "state_unsupported"
	ErrorCodeTaxIDInvalid                           ErrorCode = "tax_id_invalid"
	ErrorCodeTaxesCalculationFailed                 ErrorCode = "taxes_calculation_failed"
	ErrorCodeTestmodeChargesOnly                    ErrorCode = "testmode_charges_only"
	ErrorCodeTLSVersionUnsupported                  ErrorCode = "tls_version_unsupported"
	ErrorCodeTokenAlreadyUsed                       ErrorCode = "token_already_used"
	ErrorCodeTokenInUse                             ErrorCode = "token_in_use"
	ErrorCodeTransfersNotAllowed                    ErrorCode = "transfers_not_allowed"
	ErrorCodeUpstreamOrderCreationFailed            ErrorCode = "upstream_order_creation_failed"
	ErrorCodeURLInvalid                             ErrorCode = "url_invalid"

	// The following error code can be returned though is undocumented
	ErrorCodeInvalidSwipeData ErrorCode = "invalid_swipe_data"
)

// List of DeclineCode values.
const (
	DeclineCodeAuthenticationRequired         DeclineCode = "authentication_required"
	DeclineCodeApproveWithID                  DeclineCode = "approve_with_id"
	DeclineCodeCallIssuer                     DeclineCode = "call_issuer"
	DeclineCodeCardNotSupported               DeclineCode = "card_not_supported"
	DeclineCodeCardVelocityExceeded           DeclineCode = "card_velocity_exceeded"
	DeclineCodeCurrencyNotSupported           DeclineCode = "currency_not_supported"
	DeclineCodeDoNotHonor                     DeclineCode = "do_not_honor"
	DeclineCodeDoNotTryAgain                  DeclineCode = "do_not_try_again"
	DeclineCodeDuplicateTransaction           DeclineCode = "duplicate_transaction"
	DeclineCodeExpiredCard                    DeclineCode = "expired_card"
	DeclineCodeFraudulent                     DeclineCode = "fraudulent"
	DeclineCodeGenericDecline                 DeclineCode = "generic_decline"
	DeclineCodeIncorrectNumber                DeclineCode = "incorrect_number"
	DeclineCodeIncorrectCVC                   DeclineCode = "incorrect_cvc"
	DeclineCodeIncorrectPIN                   DeclineCode = "incorrect_pin"
	DeclineCodeIncorrectZip                   DeclineCode = "incorrect_zip"
	DeclineCodeInsufficientFunds              DeclineCode = "insufficient_funds"
	DeclineCodeInvalidAccount                 DeclineCode = "invalid_account"
	DeclineCodeInvalidAmount                  DeclineCode = "invalid_amount"
	DeclineCodeInvalidCVC                     DeclineCode = "invalid_cvc"
	DeclineCodeInvalidExpiryYear              DeclineCode = "invalid_expiry_year"
	DeclineCodeInvalidNumber                  DeclineCode = "invalid_number"
	DeclineCodeInvalidPIN                     DeclineCode = "invalid_pin"
	DeclineCodeIssuerNotAvailable             DeclineCode = "issuer_not_available"
	DeclineCodeLostCard                       DeclineCode = "lost_card"
	DeclineCodeMerchantBlacklist              DeclineCode = "merchant_blacklist"
	DeclineCodeNewAccountInformationAvailable DeclineCode = "new_account_information_available"
	DeclineCodeNoActionTaken                  DeclineCode = "no_action_taken"
	DeclineCodeNotPermitted                   DeclineCode = "not_permitted"
	DeclineCodePickupCard                     DeclineCode = "pickup_card"
	DeclineCodePINTryExceeded                 DeclineCode = "pin_try_exceeded"
	DeclineCodeProcessingError                DeclineCode = "processing_error"
	DeclineCodeReenterTransaction             DeclineCode = "reenter_transaction"
	DeclineCodeRestrictedCard                 DeclineCode = "restricted_card"
	DeclineCodeRevocationOfAllAuthorizations  DeclineCode = "revocation_of_all_authorizations"
	DeclineCodeRevocationOfAuthorization      DeclineCode = "revocation_of_authorization"
	DeclineCodeSecurityViolation              DeclineCode = "security_violation"
	DeclineCodeServiceNotAllowed              DeclineCode = "service_not_allowed"
	DeclineCodeStolenCard                     DeclineCode = "stolen_card"
	DeclineCodeStopPaymentOrder               DeclineCode = "stop_payment_order"
	DeclineCodeTestModeDecline                DeclineCode = "testmode_decline"
	DeclineCodeTransactionNotAllowed          DeclineCode = "transaction_not_allowed"
	DeclineCodeTryAgainLater                  DeclineCode = "try_again_later"
	DeclineCodeWithdrawalCountLimitExceeded   DeclineCode = "withdrawal_count_limit_exceeded"
)

// Error is the response returned when a call is unsuccessful.
// For more details see  https://stripe.com/docs/api#errors.
type Error struct {
	ChargeID    string      `json:"charge,omitempty"`
	Code        ErrorCode   `json:"code,omitempty"`
	DeclineCode DeclineCode `json:"decline_code,omitempty"`
	DocURL      string      `json:"doc_url,omitempty"`

	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong. For example, Err may hold a CardError that indicates
	// exactly what went wrong during charging a card.
	Err error `json:"-"`

	HTTPStatusCode int            `json:"status,omitempty"`
	Msg            string         `json:"message"`
	Param          string         `json:"param,omitempty"`
	PaymentIntent  *PaymentIntent `json:"payment_intent,omitempty"`
	PaymentMethod  *PaymentMethod `json:"payment_method,omitempty"`
	RequestID      string         `json:"request_id,omitempty"`
	SetupIntent    *SetupIntent   `json:"setup_intent,omitempty"`
	Source         *PaymentSource `json:"source,omitempty"`
	Type           ErrorType      `json:"type"`

	// OAuth specific Error properties. Named OAuthError because of name conflict.
	OAuthError            string `json:"error,omitempty"`
	OAuthErrorDescription string `json:"error_description,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// APIConnectionError is a failure to connect to the Stripe API.
type APIConnectionError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIConnectionError) Error() string {
	return e.stripeErr.Error()
}

// APIError is a catch all for any errors not covered by other types (and
// should be extremely uncommon).
type APIError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	return e.stripeErr.Error()
}

// AuthenticationError is a failure to properly authenticate during a request.
type AuthenticationError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	return e.stripeErr.Error()
}

// PermissionError results when you attempt to make an API request
// for which your API key doesn't have the right permissions.
type PermissionError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *PermissionError) Error() string {
	return e.stripeErr.Error()
}

// CardError are the most common type of error you should expect to handle.
// They result when the user enters a card that can't be charged for some
// reason.
type CardError struct {
	stripeErr *Error
	// DeclineCode is a code indicating a card issuer's reason for declining a
	// card (if they provided one).
	DeclineCode DeclineCode `json:"decline_code,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *CardError) Error() string {
	return e.stripeErr.Error()
}

// InvalidRequestError is an error that occurs when a request contains invalid
// parameters.
type InvalidRequestError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *InvalidRequestError) Error() string {
	return e.stripeErr.Error()
}

// RateLimitError occurs when the Stripe API is hit to with too many requests
// too quickly and indicates that the current request has been rate limited.
type RateLimitError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *RateLimitError) Error() string {
	return e.stripeErr.Error()
}

// rawError deserializes the outer JSON object returned in an error response
// from the API.
type rawError struct {
	E *rawErrorInternal `json:"error,omitempty"`
}

// rawErrorInternal embeds Error to deserialize all the standard error fields,
// but also adds other fields that may or may not be present depending on error
// type to help with deserialization. (e.g. DeclineCode).
type rawErrorInternal struct {
	*Error
	DeclineCode *DeclineCode `json:"decline_code,omitempty"`
}
