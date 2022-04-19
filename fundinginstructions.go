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
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkSepa   FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "sepa"
	FundingInstructionsBankTransferFinancialAddressSupportedNetworkZengin FundingInstructionsBankTransferFinancialAddressSupportedNetwork = "zengin"
)

// The type of financial address
type FundingInstructionsBankTransferFinancialAddressType string

// List of values that FundingInstructionsBankTransferFinancialAddressType can take
const (
	FundingInstructionsBankTransferFinancialAddressTypeIban   FundingInstructionsBankTransferFinancialAddressType = "iban"
	FundingInstructionsBankTransferFinancialAddressTypeZengin FundingInstructionsBankTransferFinancialAddressType = "zengin"
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

// Zengin Records contain Japan bank account details per the Zengin format.
type FundingInstructionsBankTransferFinancialAddressZengin struct{}

// A list of financial addresses that can be used to fund a particular balance
type FundingInstructionsBankTransferFinancialAddress struct {
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
