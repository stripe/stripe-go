//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The payment networks supported by this FinancialAddress
type FundingInstructionsBankTransferFinancialAddressSupportedNetwork string

// List of values that FundingInstructionsBankTransferFinancialAddressSupportedNetwork can take
const (
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkBACS   FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "bacs"
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkFPS    FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "fps"
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkSEPA   FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "sepa"
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkSpei   FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "spei"
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkZengin FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "zengin"
)

// The type of financial address
type FundingInstructionsBankTransferFinancialAddressType string

// List of values that FundingInstructionsBankTransferFinancialAddressType can take
const (
	FundingInstructionsBankTransferFinancialAddressTypeIBAN     FundingInstructionsBankTransferFinancialAddressType = "iban"
	FundingInstructionsBankTransferFinancialAddressTypeSortCode FundingInstructionsBankTransferFinancialAddressType = "sort_code"
	FundingInstructionsBankTransferFinancialAddressTypeSpei     FundingInstructionsBankTransferFinancialAddressType = "spei"
	FundingInstructionsBankTransferFinancialAddressTypeZengin   FundingInstructionsBankTransferFinancialAddressType = "zengin"
)

// The bank_transfer type
type FundingInstructionsBankTransferType string

// List of values that FundingInstructionsBankTransferType can take
const (
	FundingInstructionsBankTransferTypeEUBankTransfer FundingInstructionsBankTransferType = "eu_bank_transfer"
	FundingInstructionsBankTransferTypeJPBankTransfer FundingInstructionsBankTransferType = "jp_bank_transfer"
)

// The `funding_type` of the returned instructions
type FundingInstructionsFundingType string

// List of values that FundingInstructionsFundingType can take
const (
	FundingInstructionsFundingTypeBankTransfer FundingInstructionsFundingType = "bank_transfer"
)

// Iban Records contain E.U. bank account details per the SEPA format.
type FundingInstructionsBankTransferFinancialAddressIBAN struct {
	// The name of the person or business that owns the bank account
	AccountHolderName string `json:"account_holder_name"`
	// The BIC/SWIFT code of the account.
	BIC string `json:"bic"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// The IBAN of the account.
	IBAN string `json:"iban"`
}

// Sort Code Records contain U.K. bank account details per the sort code format.
type FundingInstructionsBankTransferFinancialAddressSortCode struct {
	// The name of the person or business that owns the bank account
	AccountHolderName string `json:"account_holder_name"`
	// The account number
	AccountNumber string `json:"account_number"`
	// The six-digit sort code
	SortCode string `json:"sort_code"`
}

// SPEI Records contain Mexico bank account details per the SPEI format.
type FundingInstructionsBankTransferFinancialAddressSpei struct {
	// The three-digit bank code
	BankCode string `json:"bank_code"`
	// The short banking institution name
	BankName string `json:"bank_name"`
	// The CLABE number
	Clabe string `json:"clabe"`
}

// Zengin Records contain Japan bank account details per the Zengin format.
type FundingInstructionsBankTransferFinancialAddressZengin struct {
	// The account holder name
	AccountHolderName string `json:"account_holder_name"`
	// The account number
	AccountNumber string `json:"account_number"`
	// The bank account type. In Japan, this can only be `futsu` or `toza`.
	AccountType string `json:"account_type"`
	// The bank code of the account
	BankCode string `json:"bank_code"`
	// The bank name of the account
	BankName string `json:"bank_name"`
	// The branch code of the account
	BranchCode string `json:"branch_code"`
	// The branch name of the account
	BranchName string `json:"branch_name"`
}

// A list of financial addresses that can be used to fund a particular balance
type FundingInstructionsBankTransferFinancialAddress struct {
	// Iban Records contain E.U. bank account details per the SEPA format.
	IBAN *FundingInstructionsBankTransferFinancialAddressIBAN `json:"iban"`
	// Sort Code Records contain U.K. bank account details per the sort code format.
	SortCode *FundingInstructionsBankTransferFinancialAddressSortCode `json:"sort_code"`
	// SPEI Records contain Mexico bank account details per the SPEI format.
	Spei *FundingInstructionsBankTransferFinancialAddressSpei `json:"spei"`
	// The payment networks supported by this FinancialAddress
	SupportedNetworks []FundingInstructionsBankTransferFinancialAddressSupportedNetwork `json:"supported_networks"`
	// The type of financial address
	Type FundingInstructionsBankTransferFinancialAddressType `json:"type"`
	// Zengin Records contain Japan bank account details per the Zengin format.
	Zengin *FundingInstructionsBankTransferFinancialAddressZengin `json:"zengin"`
}
type FundingInstructionsBankTransfer struct {
	// The country of the bank account to fund
	Country string `json:"country"`
	// A list of financial addresses that can be used to fund a particular balance
	FinancialAddresses []*FundingInstructionsBankTransferFinancialAddress `json:"financial_addresses"`
	// The bank_transfer type
	Type FundingInstructionsBankTransferType `json:"type"`
}

// Each customer has a [`balance`](https://stripe.com/docs/api/customers/object#customer_object-balance) that is
// automatically applied to future invoices and payments using the `customer_balance` payment method.
// Customers can fund this balance by initiating a bank transfer to any account in the
// `financial_addresses` field.
// Related guide: [Customer Balance - Funding Instructions](https://stripe.com/docs/payments/customer-balance/funding-instructions) to learn more
type FundingInstructions struct {
	APIResource
	BankTransfer *FundingInstructionsBankTransfer `json:"bank_transfer"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The `funding_type` of the returned instructions
	FundingType FundingInstructionsFundingType `json:"funding_type"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
