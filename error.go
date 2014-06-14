package stripe

import "encoding/json"

type ErrorType string
type ErrorCode string

const (
	InvalidRequest ErrorType = "invalid_request_error"
	ApiErr         ErrorType = "api_error"
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

type Error struct {
	Type  ErrorType `json:"type"`
	Msg   string    `json:"message"`
	Code  ErrorCode `json:"code,omitempty"`
	Param string    `json:"param,omitempty"`
}

func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
