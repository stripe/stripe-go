//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of the card wallet, one of `apple_pay` or `google_pay`.
type CryptoCustomerPaymentTokenCardWalletType string

// List of values that CryptoCustomerPaymentTokenCardWalletType can take
const (
	CryptoCustomerPaymentTokenCardWalletTypeApplePay  CryptoCustomerPaymentTokenCardWalletType = "apple_pay"
	CryptoCustomerPaymentTokenCardWalletTypeGooglePay CryptoCustomerPaymentTokenCardWalletType = "google_pay"
)

// Type of the Payment Token.
type CryptoCustomerPaymentTokenType string

// List of values that CryptoCustomerPaymentTokenType can take
const (
	CryptoCustomerPaymentTokenTypeCard          CryptoCustomerPaymentTokenType = "card"
	CryptoCustomerPaymentTokenTypeUSBankAccount CryptoCustomerPaymentTokenType = "us_bank_account"
)

// Account type: `checkings` or `savings`.
type CryptoCustomerPaymentTokenUSBankAccountAccountType string

// List of values that CryptoCustomerPaymentTokenUSBankAccountAccountType can take
const (
	CryptoCustomerPaymentTokenUSBankAccountAccountTypeChecking CryptoCustomerPaymentTokenUSBankAccountAccountType = "checking"
	CryptoCustomerPaymentTokenUSBankAccountAccountTypeSavings  CryptoCustomerPaymentTokenUSBankAccountAccountType = "savings"
)

// Lists the Payment Tokens for a Crypto Customer.
type CryptoCustomerPaymentTokenListParams struct {
	ListParams `form:"*"`
	ID         *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoCustomerPaymentTokenListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type CryptoCustomerPaymentTokenCardWallet struct {
	// The type of the card wallet, one of `apple_pay` or `google_pay`.
	Type CryptoCustomerPaymentTokenCardWalletType `json:"type"`
}

// A `card` PaymentToken, this hash contains details of the card PaymentToken.
type CryptoCustomerPaymentTokenCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *CryptoCustomerPaymentTokenCardWallet `json:"wallet"`
}

// A `us_bank_account` PaymentToken, this hash contains details of the US bank account PaymentToken.
type CryptoCustomerPaymentTokenUSBankAccount struct {
	// Account type: `checkings` or `savings`.
	AccountType CryptoCustomerPaymentTokenUSBankAccountAccountType `json:"account_type"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
}

// A read-only representation of a user's PaymentMethod for use in Crypto On Ramp transactions.
type CryptoCustomerPaymentToken struct {
	// A `card` PaymentToken, this hash contains details of the card PaymentToken.
	Card *CryptoCustomerPaymentTokenCard `json:"card"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Type of the Payment Token.
	Type CryptoCustomerPaymentTokenType `json:"type"`
	// A `us_bank_account` PaymentToken, this hash contains details of the US bank account PaymentToken.
	USBankAccount *CryptoCustomerPaymentTokenUSBankAccount `json:"us_bank_account"`
}

// CryptoCustomerPaymentTokenList is a list of CustomerPaymentTokens as retrieved from a list endpoint.
type CryptoCustomerPaymentTokenList struct {
	APIResource
	ListMeta
	Data []*CryptoCustomerPaymentToken `json:"data"`
}
