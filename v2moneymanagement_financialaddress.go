//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of bank account details.
type V2MoneyManagementFinancialAddressBankAccountType string

// List of values that V2MoneyManagementFinancialAddressBankAccountType can take
const (
	V2MoneyManagementFinancialAddressBankAccountTypeABA      V2MoneyManagementFinancialAddressBankAccountType = "aba"
	V2MoneyManagementFinancialAddressBankAccountTypeIBAN     V2MoneyManagementFinancialAddressBankAccountType = "iban"
	V2MoneyManagementFinancialAddressBankAccountTypeSortCode V2MoneyManagementFinancialAddressBankAccountType = "sort_code"
)

// Closed Enum. The status of the FinancialAddress.
type V2MoneyManagementFinancialAddressStatus string

// List of values that V2MoneyManagementFinancialAddressStatus can take
const (
	V2MoneyManagementFinancialAddressStatusActive   V2MoneyManagementFinancialAddressStatus = "active"
	V2MoneyManagementFinancialAddressStatusArchived V2MoneyManagementFinancialAddressStatus = "archived"
	V2MoneyManagementFinancialAddressStatusFailed   V2MoneyManagementFinancialAddressStatus = "failed"
	V2MoneyManagementFinancialAddressStatusPending  V2MoneyManagementFinancialAddressStatus = "pending"
)

// Open Enum. The type of FinancialAddress.
type V2MoneyManagementFinancialAddressType string

// List of values that V2MoneyManagementFinancialAddressType can take
const (
	V2MoneyManagementFinancialAddressTypeBankAccount V2MoneyManagementFinancialAddressType = "bank_account"
)

// The address of the account holder.
type V2MoneyManagementFinancialAddressBankAccountABAAccountHolderAddress struct {
	// City.
	City string `json:"city"`
	// Country.
	Country string `json:"country"`
	// Address line 1.
	Line1 string `json:"line1"`
	// Address line 2.
	Line2 string `json:"line2"`
	// Postal code.
	PostalCode string `json:"postal_code"`
	// State or province.
	State string `json:"state"`
	// Town or suburb.
	Town string `json:"town"`
}

// ABA bank account details (US).
type V2MoneyManagementFinancialAddressBankAccountABA struct {
	// The address of the account holder.
	AccountHolderAddress *V2MoneyManagementFinancialAddressBankAccountABAAccountHolderAddress `json:"account_holder_address,omitempty"`
	// The name of the account holder.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The full account number.
	AccountNumber string `json:"account_number,omitempty"`
	// The name of the bank.
	BankName string `json:"bank_name,omitempty"`
	// The last four digits of the account number.
	Last4 string `json:"last4"`
	// The ABA routing number.
	RoutingNumber string `json:"routing_number"`
}

// IBAN bank account details.
type V2MoneyManagementFinancialAddressBankAccountIBAN struct {
	// The name of the account holder.
	AccountHolderName string `json:"account_holder_name"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The country of the bank account.
	Country string `json:"country"`
	// The full IBAN.
	IBAN string `json:"iban,omitempty"`
	// The last four digits of the IBAN.
	Last4 string `json:"last4"`
}

// Sort code bank account details (UK).
type V2MoneyManagementFinancialAddressBankAccountSortCode struct {
	// The name of the account holder.
	AccountHolderName string `json:"account_holder_name"`
	// The full account number.
	AccountNumber string `json:"account_number,omitempty"`
	// The last four digits of the account number.
	Last4 string `json:"last4"`
	// The sort code.
	SortCode string `json:"sort_code"`
}

// Bank account details for this FinancialAddress.
type V2MoneyManagementFinancialAddressBankAccount struct {
	// ABA bank account details (US).
	ABA *V2MoneyManagementFinancialAddressBankAccountABA `json:"aba,omitempty"`
	// The country of the bank account.
	Country string `json:"country,omitempty"`
	// Open Enum. The currency of the bank account.
	Currency Currency `json:"currency"`
	// IBAN bank account details.
	IBAN *V2MoneyManagementFinancialAddressBankAccountIBAN `json:"iban,omitempty"`
	// Sort code bank account details (UK).
	SortCode *V2MoneyManagementFinancialAddressBankAccountSortCode `json:"sort_code,omitempty"`
	// Open Enum. The type of bank account details.
	Type V2MoneyManagementFinancialAddressBankAccountType `json:"type"`
}

// A FinancialAddress contains information needed to transfer money to a Financial Account. A Financial Account can have more than one Financial Address.
type V2MoneyManagementFinancialAddress struct {
	APIResource
	// Bank account details for this FinancialAddress.
	BankAccount *V2MoneyManagementFinancialAddressBankAccount `json:"bank_account,omitempty"`
	// The creation timestamp of the FinancialAddress.
	Created time.Time `json:"created"`
	// The ID of the FinancialAccount this FinancialAddress corresponds to.
	FinancialAccount string `json:"financial_account"`
	// The ID of the FinancialAddress.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount.
	SettlementCurrency Currency `json:"settlement_currency,omitempty"`
	// Closed Enum. The status of the FinancialAddress.
	Status V2MoneyManagementFinancialAddressStatus `json:"status"`
	// Open Enum. The type of FinancialAddress.
	Type V2MoneyManagementFinancialAddressType `json:"type"`
}
