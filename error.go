package stripe

import "encoding/json"

// ErrorType is the list of allowed values for the error's type.
type ErrorType string

// ErrorCode is the list of allowed values for the error's code.
type ErrorCode string

const (
	ErrorTypeAPI            ErrorType = "api_error"
	ErrorTypeAPIConnection  ErrorType = "api_connection_error"
	ErrorTypeAuthentication ErrorType = "authentication_error"
	ErrorTypeCard           ErrorType = "card_error"
	ErrorTypeInvalidRequest ErrorType = "invalid_request_error"
	ErrorTypeRateLimit      ErrorType = "rate_limit_error"

	IncorrectNum  ErrorCode = "incorrect_number"
	InvalidNum    ErrorCode = "invalid_number"
	InvalidExpM   ErrorCode = "invalid_expiry_month"
	InvalidExpY   ErrorCode = "invalid_expiry_year"
	InvalidCvc    ErrorCode = "invalid_cvc"
	ExpiredCard   ErrorCode = "expired_card"
	IncorrectCvc  ErrorCode = "incorrect_cvc"
	IncorrectZip  ErrorCode = "incorrect_zip"
	CardDeclined  ErrorCode = "card_declined"
	Missing       ErrorCode = "missing"
	ProcessingErr ErrorCode = "processing_error"
	RateLimit     ErrorCode = "rate_limit"
)

// StripeError is the response returned when a call is unsuccessful.
// For more details see  https://stripe.com/docs/api#errors.
type StripeError struct {
	Type           ErrorType `json:"type"`
	Msg            string    `json:"message"`
	Code           ErrorCode `json:"code,omitempty"`
	Param          string    `json:"param,omitempty"`
	RequestID      string    `json:"request_id,omitempty"`
	HTTPStatusCode int       `json:"status,omitempty"`
	ChargeID       string    `json:"charge,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *StripeError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// APIConnectionError is a failure to connect to the Stripe API.
type APIConnectionError struct {
	StripeError
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIConnectionError) Error() string {
	return e.StripeError.Error()
}

// APIError is a catch all for any errors not covered by other types (and
// should be extremely uncommon).
type APIError struct {
	StripeError
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	return e.StripeError.Error()
}

// AuthenticationError is a failure to properly authenticate during a request.
type AuthenticationError struct {
	StripeError
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	return e.StripeError.Error()
}

// CardError are the most common type of error you should expect to handle.
// They result when the user enters a card that can't be charged for some
// reason.
type CardError struct {
	StripeError
	DeclineCode string `json:"decline_code,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *CardError) Error() string {
	return e.StripeError.Error()
}

// InvalidRequestError is an error that occurs when a request contains invalid
// parameters.
type InvalidRequestError struct {
	StripeError
}

// Error serializes the error object to JSON and returns it as a string.
func (e *InvalidRequestError) Error() string {
	return e.StripeError.Error()
}

// RateLimitError occurs when the Stripe API is hit to with too many requests
// too quickly and indicates that the current request has been rate limited.
type RateLimitError struct {
	StripeError
}

// Error serializes the error object to JSON and returns it as a string.
func (e *RateLimitError) Error() string {
	return e.StripeError.Error()
}
