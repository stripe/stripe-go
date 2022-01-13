//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of the token: `account`, `bank_account`, `card`, or `pii`.
type TokenType string

// List of values that TokenType can take
const (
	TokenTypeAccount     TokenType = "account"
	TokenTypeBankAccount TokenType = "bank_account"
	TokenTypeCard        TokenType = "card"
	TokenTypeCVCUpdate   TokenType = "cvc_update"
	TokenTypePII         TokenType = "pii"
)

// Retrieves the token with the given ID.
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

// Information for the account this token will represent.
type TokenAccountParams struct {
	BusinessType        *string               `form:"business_type"`
	Company             *AccountCompanyParams `form:"company"`
	Individual          *PersonParams         `form:"individual"`
	TOSShownAndAccepted *bool                 `form:"tos_shown_and_accepted"`
}

// The updated CVC value this token will represent.
type TokenCVCUpdateParams struct {
	CVC *string `form:"cvc"`
}

// The PII this token will represent.
type PIIParams struct {
	Params   `form:"*"`
	IDNumber *string `form:"id_number"`
}

// Tokenization is the process Stripe uses to collect sensitive card or bank
// account details, or personally identifiable information (PII), directly from
// your customers in a secure manner. A token representing this information is
// returned to your server to use. You should use our
// [recommended payments integrations](https://stripe.com/docs/payments) to perform this process
// client-side. This ensures that no sensitive card data touches your server,
// and allows your integration to operate in a PCI-compliant way.
//
// If you cannot use client-side tokenization, you can also create tokens using
// the API with either your publishable or secret API key. Keep in mind that if
// your integration uses this method, you are responsible for any PCI compliance
// that may be required, and you must keep your secret API key safe. Unlike with
// client-side tokenization, your customer's information is not sent directly to
// Stripe, so we cannot determine how it is handled or stored.
//
// Tokens cannot be stored or used more than once. To store card or bank account
// information for later use, you can create [Customer](https://stripe.com/docs/api#customers)
// objects or [Custom accounts](https://stripe.com/docs/api#external_accounts). Note that
// [Radar](https://stripe.com/docs/radar), our integrated solution for automatic fraud protection,
// performs best with integrations that use client-side tokenization.
//
// Related guide: [Accept a payment](https://stripe.com/docs/payments/accept-a-payment-charges#web-create-token)
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
