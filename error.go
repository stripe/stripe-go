package stripe

import (
	"encoding/json"
	"net/http"
)

// errorEnums: The beginning of the section generated from our OpenAPI spec
// errorEnums: The end of the section generated from our OpenAPI spec

// ErrorType is the list of allowed values for the error's type.
type ErrorType string

// List of values that ErrorType can take.
// errorTypes: The beginning of the section generated from our OpenAPI spec
// errorTypes: The end of the section generated from our OpenAPI spec

// DeclineCode is the list of reasons provided by card issuers for decline of payment.
type DeclineCode string

// ErrorCode is the list of allowed values for the error's code.
type ErrorCode string

// List of values that ErrorCode can take.
// For descriptions see https://stripe.com/docs/error-codes
// v1ErrorCodes: The beginning of the section generated from our OpenAPI spec
// v1ErrorCodes: The end of the section generated from our OpenAPI spec

// List of DeclineCode values.
// For descriptions see https://stripe.com/docs/declines/codes
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
	DeclineCodeInvalidExpiryMonth             DeclineCode = "invalid_expiry_month"
	DeclineCodeInvalidExpiryYear              DeclineCode = "invalid_expiry_year"
	DeclineCodeInvalidNumber                  DeclineCode = "invalid_number"
	DeclineCodeInvalidPIN                     DeclineCode = "invalid_pin"
	DeclineCodeIssuerNotAvailable             DeclineCode = "issuer_not_available"
	DeclineCodeLostCard                       DeclineCode = "lost_card"
	DeclineCodeMerchantBlacklist              DeclineCode = "merchant_blacklist"
	DeclineCodeNewAccountInformationAvailable DeclineCode = "new_account_information_available"
	DeclineCodeNoActionTaken                  DeclineCode = "no_action_taken"
	DeclineCodeNotPermitted                   DeclineCode = "not_permitted"
	DeclineCodeOfflinePINRequired             DeclineCode = "offline_pin_required"
	DeclineCodeOnlineOrOfflinePINRequired     DeclineCode = "online_or_offline_pin_required"
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

type retrier interface {
	canRetry() bool
}

type redacter interface {
	redact() error
}

// Error is the response returned when a call is unsuccessful.
// For more details see https://stripe.com/docs/api#errors.
type Error struct {
	APIResource

	ChargeID    string      `json:"charge,omitempty"`
	Code        ErrorCode   `json:"code,omitempty"`
	DeclineCode DeclineCode `json:"decline_code,omitempty"`
	DocURL      string      `json:"doc_url,omitempty"`

	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong. For example, Err may hold a CardError that indicates
	// exactly what went wrong during charging a card.
	Err error `json:"-"`

	HTTPStatusCode    int               `json:"status,omitempty"`
	Msg               string            `json:"message"`
	DeveloperMsg      string            `json:"developer_message,omitempty"`
	Param             string            `json:"param,omitempty"`
	PaymentIntent     *PaymentIntent    `json:"payment_intent,omitempty"`
	PaymentMethod     *PaymentMethod    `json:"payment_method,omitempty"`
	PaymentMethodType PaymentMethodType `json:"payment_method_type,omitempty"`
	RequestID         string            `json:"request_id,omitempty"`
	RequestLogURL     string            `json:"request_log_url,omitempty"`
	SetupIntent       *SetupIntent      `json:"setup_intent,omitempty"`
	Source            *PaymentSource    `json:"source,omitempty"`
	Type              ErrorType         `json:"type"`

	// OAuth specific Error properties. Named OAuthError because of name conflict.
	OAuthError            string `json:"error,omitempty"`
	OAuthErrorDescription string `json:"error_description,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// Unwrap returns the wrapped typed error.
func (e *Error) Unwrap() error {
	return e.Err
}

// canRetry implements the retrier interface.
func (e *Error) canRetry() bool {
	if e == nil {
		return false
	}

	// 429 Too Many Requests
	//
	// There are a few different problems that can lead to a 429. The most
	// common is rate limiting, on which we *don't* want to retry because
	// that'd likely contribute to more contention problems. However, some 429s
	// are lock timeouts, which is when a request conflicted with another
	// request or an internal process on some particular object. These 429s are
	// safe to retry.
	if e.HTTPStatusCode == http.StatusTooManyRequests && e.Code == ErrorCodeLockTimeout {
		return true
	}

	return false
}

// redact returns a copy of the error object with sensitive fields replaced with
// a placeholder value. This implements the redacter interface.
func (e *Error) redact() error {
	// Fast path, since this applies to most cases
	if e.PaymentIntent == nil && e.SetupIntent == nil {
		return e
	}
	errCopy := *e
	if e.PaymentIntent != nil {
		pi := *e.PaymentIntent
		errCopy.PaymentIntent = &pi
		errCopy.PaymentIntent.ClientSecret = "REDACTED"
	}
	if e.SetupIntent != nil {
		si := *e.SetupIntent
		errCopy.SetupIntent = &si
		errCopy.SetupIntent.ClientSecret = "REDACTED"
	}
	return &errCopy
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

// IdempotencyError occurs when an Idempotency-Key is re-used on a request
// that does not match the first request's API endpoint and parameters.
type IdempotencyError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *IdempotencyError) Error() string {
	return e.stripeErr.Error()
}

// errorStructs: The beginning of the section generated from our OpenAPI spec
// errorStructs: The end of the section generated from our OpenAPI spec

// V2RawError is a catch-all for any errors not covered by other types
type V2RawError struct {
	Code        string     `json:"code"`
	Type        *ErrorType `json:"type,omitempty"`
	Message     string     `json:"message"`
	UserMessage *string    `json:"user_message,omitempty"`
	// HTTPStatusCode is the HTTP status code of the response returned by Stripe.
	HTTPStatusCode int `json:"status,omitempty"`
	// RequestID is the ID of the request, as returned by Stripe.
	RequestID string `json:"request_id,omitempty"`
}

func (e *V2RawError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

func (e *V2RawError) redact() error {
	return e
}

func (e *V2RawError) canRetry() bool {
	return false
}

// rawError deserializes the outer JSON object returned in an error response
// from the API.
type rawError struct {
	Error *Error `json:"error,omitempty"`
}
