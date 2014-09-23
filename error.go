package stripe

import "encoding/json"

// ErrorType is the list of allowed values for the error's type.
// Allowed values are "invalid_request_error", "api_error", "card_error".
type ErrorType string

// ErrorCode is the list of allowed values for the error's code.
// Allowed values are "incorrect_number", "invalid_number", "invalid_expiry_month",
// "invalid_expiry_year", "invalid_cvc", "expired_card", "incorrect_cvc", "incorrect_zip",
// "card_declined", "missing", "processing_error", "rate_limit".
type ErrorCode string

const (
	InvalidRequest ErrorType = "invalid_request_error"
	APIErr         ErrorType = "api_error"
	CardErr        ErrorType = "card_error"

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

// Error is the response returned when a call is unsuccessful.
// For more details see  https://stripe.com/docs/api#errors.
type Error struct {
	Type           ErrorType `json:"type"`
	Msg            string    `json:"message"`
	Code           ErrorCode `json:"code,omitempty"`
	Param          string    `json:"param,omitempty"`
	HTTPStatusCode int       `json:"-"`
}

// Error serializes the Error object and prints the JSON string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
