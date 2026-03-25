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
	// Information for the account this token represents.
	Account *TokenAccountParams `form:"account" json:"account,omitempty"`
	// The bank account this token will represent.
	BankAccount *BankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The card this token will represent. If you also pass in a customer, the card must be the ID of a card belonging to the customer. Otherwise, if you do not pass in a customer, this is a dictionary containing a user's credit card details, with the options described below.
	Card *CardParams `form:"card" json:"card,omitempty"`
	// Create a token for the customer, which is owned by the application's account. You can only use this with an [OAuth access token](https://docs.stripe.com/connect/standard-accounts) or [Stripe-Account header](https://docs.stripe.com/connect/authentication). Learn more about [cloning saved payment methods](https://docs.stripe.com/connect/cloning-saved-payment-methods).
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The updated CVC value this token represents.
	CVCUpdate *TokenCVCUpdateParams `form:"cvc_update" json:"cvc_update,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Information for the person this token represents.
	Person *PersonParams `form:"person" json:"person,omitempty"`
	// The PII this token represents.
	PII *TokenPIIParams `form:"pii" json:"pii,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Information for the account this token represents.
type TokenAccountParams struct {
	// The business type.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// Information about the company or business.
	Company *AccountCompanyParams `form:"company" json:"company,omitempty"`
	// Information about the person represented by the account.
	Individual *PersonParams `form:"individual" json:"individual,omitempty"`
	// Whether the user described by the data in the token has been shown [the Stripe Connected Account Agreement](https://docs.stripe.com/connect/account-tokens#stripe-connected-account-agreement). When creating an account token to create a new Connect account, this value must be `true`.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted" json:"tos_shown_and_accepted,omitempty"`
}

// The updated CVC value this token represents.
type TokenCVCUpdateParams struct {
	// The CVC value, in string form.
	CVC *string `form:"cvc" json:"cvc"`
}

// The PII this token represents.
type TokenPIIParams struct {
	// The `id_number` for the PII, in string form.
	IDNumber *string `form:"id_number" json:"id_number,omitempty"`
}

// Retrieves the token with the given ID.
type TokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TokenRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Information for the account this token represents.
type TokenCreateAccountParams struct {
	// The business type.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// Information about the company or business.
	Company *AccountCompanyParams `form:"company" json:"company,omitempty"`
	// Information about the person represented by the account.
	Individual *PersonParams `form:"individual" json:"individual,omitempty"`
	// Whether the user described by the data in the token has been shown [the Stripe Connected Account Agreement](https://docs.stripe.com/connect/account-tokens#stripe-connected-account-agreement). When creating an account token to create a new Connect account, this value must be `true`.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted" json:"tos_shown_and_accepted,omitempty"`
}

// The updated CVC value this token represents.
type TokenCreateCVCUpdateParams struct {
	// The CVC value, in string form.
	CVC *string `form:"cvc" json:"cvc"`
}

// The PII this token represents.
type TokenCreatePIIParams struct {
	// The `id_number` for the PII, in string form.
	IDNumber *string `form:"id_number" json:"id_number,omitempty"`
}

// Creates a single-use token that represents a bank account's details.
// You can use this token with any v1 API method in place of a bank account dictionary. You can only use this token once. To do so, attach it to a [connected account](https://docs.stripe.com/api#accounts) where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is application, which includes Custom accounts.
type TokenCreateParams struct {
	Params `form:"*"`
	// Information for the account this token represents.
	Account *TokenCreateAccountParams `form:"account" json:"account,omitempty"`
	// The bank account this token will represent.
	BankAccount *BankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The card this token will represent. If you also pass in a customer, the card must be the ID of a card belonging to the customer. Otherwise, if you do not pass in a customer, this is a dictionary containing a user's credit card details, with the options described below.
	Card *CardParams `form:"card" json:"card,omitempty"`
	// Create a token for the customer, which is owned by the application's account. You can only use this with an [OAuth access token](https://docs.stripe.com/connect/standard-accounts) or [Stripe-Account header](https://docs.stripe.com/connect/authentication). Learn more about [cloning saved payment methods](https://docs.stripe.com/connect/cloning-saved-payment-methods).
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The updated CVC value this token represents.
	CVCUpdate *TokenCreateCVCUpdateParams `form:"cvc_update" json:"cvc_update,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Information for the person this token represents.
	Person *PersonParams `form:"person" json:"person,omitempty"`
	// The PII this token represents.
	PII *TokenCreatePIIParams `form:"pii" json:"pii,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TokenCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Tokenization is the process Stripe uses to collect sensitive card or bank
// account details, or personally identifiable information (PII), directly from
// your customers in a secure manner. A token representing this information is
// returned to your server to use. Use our
// [recommended payments integrations](https://docs.stripe.com/payments) to perform this process
// on the client-side. This guarantees that no sensitive card data touches your server,
// and allows your integration to operate in a PCI-compliant way.
//
// If you can't use client-side tokenization, you can also create tokens using
// the API with either your publishable or secret API key. If
// your integration uses this method, you're responsible for any PCI compliance
// that it might require, and you must keep your secret API key safe. Unlike with
// client-side tokenization, your customer's information isn't sent directly to
// Stripe, so we can't determine how it's handled or stored.
//
// You can't store or use tokens more than once. To store card or bank account
// information for later use, create [Customer](https://docs.stripe.com/api#customers)
// objects or [External accounts](https://docs.stripe.com/api#external_accounts).
// [Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,
// performs best with integrations that use client-side tokenization.
type Token struct {
	APIResource
	// These bank accounts are payment methods on `Customer` objects.
	//
	// On the other hand [External Accounts](https://docs.stripe.com/api#external_accounts) are transfer
	// destinations on `Account` objects for connected accounts.
	// They can be bank accounts or debit cards as well, and are documented in the links above.
	//
	// Related guide: [Bank debits and transfers](https://docs.stripe.com/payments/bank-debits-transfers)
	BankAccount *BankAccount `json:"bank_account,omitempty"`
	// You can store multiple cards on a customer in order to charge the customer
	// later. You can also store multiple debit cards on a recipient in order to
	// transfer to those cards later.
	//
	// Related guide: [Card payments with Sources](https://docs.stripe.com/sources/cards)
	Card *Card `json:"card,omitempty"`
	// IP address of the client that generates the token.
	ClientIP string `json:"client_ip"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Type of the token: `account`, `bank_account`, `card`, or `pii`.
	Type TokenType `json:"type"`
	// Determines if you have already used this token (you can only use tokens once).
	Used bool `json:"used"`
}
