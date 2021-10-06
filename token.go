//
//
// File generated from our OpenAPI spec
//
//

package stripe

// TokenType is the list of allowed values for a token's type.
type TokenType string

// List of values that TokenType can take.
const (
	TokenTypeAccount     TokenType = "account"
	TokenTypeBankAccount TokenType = "bank_account"
	TokenTypeCard        TokenType = "card"
	TokenTypeCVCUpdate   TokenType = "cvc_update"
	TokenTypePII         TokenType = "pii"
)

// TokenCVCUpdateParams is the set of parameters that can be used when creating a CVC token.
type TokenCVCUpdateParams struct {
	CVC *string `form:"cvc"`
}

// TokenAccountParams is the set of parameters that can be used when creating a Account token.
type TokenAccountParams struct {
	BusinessType        *string               `form:"business_type"`
	Company             *AccountCompanyParams `form:"company"`
	Individual          *PersonParams         `form:"individual"`
	TOSShownAndAccepted *bool                 `form:"tos_shown_and_accepted"`
}

// TokenParams is the set of parameters that can be used when creating a token.
// For more details see:
// - https://stripe.com/docs/api#create_card_token
// - https://stripe.com/docs/api#create_bank_account_token
// - https://stripe.com/docs/api/tokens/create_account
// - https://stripe.com/docs/api/tokens/create_person
type TokenParams struct {
	Params      `form:"*"`
	Account     *TokenAccountParams   `form:"account"`
	BankAccount *BankAccountParams    `form:"bank_account"`
	Card        *CardParams           `form:"card"`
	Customer    *string               `form:"customer"`
	CVCUpdate   *TokenCVCUpdateParams `form:"cvc_update"`
	// Email is an undocumented parameter used by Stripe Checkout
	// It may be removed from the API without notice.
	Email  *string       `form:"email"`
	Person *PersonParams `form:"person"`
	PII    *PIIParams    `form:"pii"`
}

// PIIParams are parameters for personal identifiable information (PII).
type PIIParams struct {
	Params   `form:"*"`
	IDNumber *string `form:"id_number"`
}

// Token is the resource representing a Stripe token.
// For more details see https://stripe.com/docs/api#tokens.
type Token struct {
	APIResource
	BankAccount *BankAccount `json:"bank_account"`
	Card        *Card        `json:"card"`
	ClientIP    string       `json:"client_ip"`
	Created     int64        `json:"created"`
	// Email is an undocumented field but included for all tokens created
	// with Stripe Checkout.
	Email    string    `json:"email"`
	ID       string    `json:"id"`
	Livemode bool      `json:"livemode"`
	Object   string    `json:"object"`
	Type     TokenType `json:"type"`
	Used     bool      `json:"used"`
}
