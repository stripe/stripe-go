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
	Params `form:"*"`
	// Information for the account this token will represent.
	Account *TokenAccountParams `form:"account"`
	// The bank account this token will represent.
	BankAccount *BankAccountParams `form:"bank_account"`
	Card        *CardParams        `form:"card"`
	// The customer (owned by the application's account) for which to create a token. This can be used only with an [OAuth access token](https://stripe.com/docs/connect/standard-accounts) or [Stripe-Account header](https://stripe.com/docs/connect/authentication). For more details, see [Cloning Saved Payment Methods](https://stripe.com/docs/connect/cloning-saved-payment-methods).
	Customer *string `form:"customer"`
	// The updated CVC value this token will represent.
	CVCUpdate *TokenCVCUpdateParams `form:"cvc_update"`
	// Email is an undocumented parameter used by Stripe Checkout
	// It may be removed from the API without notice.
	Email *string `form:"email"`
	// Information for the person this token will represent.
	Person *PersonParams `form:"person"`
	// The PII this token will represent.
	PII *PIIParams `form:"pii"`
}

// Information for the account this token will represent.
type TokenAccountParams struct {
	// The business type.
	BusinessType *string `form:"business_type"`
	// Information about the company or business.
	Company *AccountCompanyParams `form:"company"`
	// Information about the person represented by the account.
	Individual *PersonParams `form:"individual"`
	// Whether the user described by the data in the token has been shown [the Stripe Connected Account Agreement](https://stripe.com/docs/connect/account-tokens#stripe-connected-account-agreement). When creating an account token to create a new Connect account, this value must be `true`.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// The updated CVC value this token will represent.
type TokenCVCUpdateParams struct {
	// The CVC value, in string form.
	CVC *string `form:"cvc"`
}

// The PII this token will represent.
type PIIParams struct {
	Params `form:"*"`
	// The `id_number` for the PII, in string form.
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
	// These bank accounts are payment methods on `Customer` objects.
	//
	// On the other hand [External Accounts](https://stripe.com/docs/api#external_accounts) are transfer
	// destinations on `Account` objects for [Custom accounts](https://stripe.com/docs/connect/custom-accounts).
	// They can be bank accounts or debit cards as well, and are documented in the links above.
	//
	// Related guide: [Bank Debits and Transfers](https://stripe.com/docs/payments/bank-debits-transfers).
	BankAccount *BankAccount `json:"bank_account"`
	// You can store multiple cards on a customer in order to charge the customer
	// later. You can also store multiple debit cards on a recipient in order to
	// transfer to those cards later.
	//
	// Related guide: [Card Payments with Sources](https://stripe.com/docs/sources/cards).
	Card *Card `json:"card"`
	// IP address of the client that generated the token.
	ClientIP string `json:"client_ip"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Email is an undocumented field but included for all tokens created
	// with Stripe Checkout.
	Email string `json:"email"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Type of the token: `account`, `bank_account`, `card`, or `pii`.
	Type TokenType `json:"type"`
	// Whether this token has already been used (tokens can be used only once).
	Used bool `json:"used"`
}
