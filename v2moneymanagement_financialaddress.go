//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of Credentials that are provisioned for the FinancialAddress.
type V2MoneyManagementFinancialAddressCredentialsType string

// List of values that V2MoneyManagementFinancialAddressCredentialsType can take
const (
	V2MoneyManagementFinancialAddressCredentialsTypeGBBankAccount V2MoneyManagementFinancialAddressCredentialsType = "gb_bank_account"
	V2MoneyManagementFinancialAddressCredentialsTypeUSBankAccount V2MoneyManagementFinancialAddressCredentialsType = "us_bank_account"
)

// Closed Enum. An enum representing the status of the FinancialAddress. This indicates whether or not the FinancialAddress can be used for any money movement flows.
type V2MoneyManagementFinancialAddressStatus string

// List of values that V2MoneyManagementFinancialAddressStatus can take
const (
	V2MoneyManagementFinancialAddressStatusActive   V2MoneyManagementFinancialAddressStatus = "active"
	V2MoneyManagementFinancialAddressStatusArchived V2MoneyManagementFinancialAddressStatus = "archived"
	V2MoneyManagementFinancialAddressStatusFailed   V2MoneyManagementFinancialAddressStatus = "failed"
	V2MoneyManagementFinancialAddressStatusPending  V2MoneyManagementFinancialAddressStatus = "pending"
)

// The credentials of the UK Bank Account for the FinancialAddress. This contains unique banking details such as the sort code, account number, etc. of a UK bank account.
type V2MoneyManagementFinancialAddressCredentialsGBBankAccount struct {
	// The account holder name to be used during bank transference.
	AccountHolderName string `json:"account_holder_name"`
	// The account number of the UK Bank Account.
	AccountNumber string `json:"account_number,omitempty"`
	// The last four digits of the UK Bank Account number. This will always be returned.
	// To view the full account number when retrieving or listing FinancialAddresses, use the `include` request parameter.
	Last4 string `json:"last4"`
	// The sort code of the UK Bank Account.
	SortCode string `json:"sort_code"`
}

// The credentials of the SEPA Bank Account for the FinancialAddress. This contains unique banking details such as the IBAN, BIC, etc. of a SEPA bank account.
type V2MoneyManagementFinancialAddressCredentialsSEPABankAccount struct {
	// The account holder name to be used during bank transfers.
	AccountHolderName string `json:"account_holder_name"`
	// The name of the Bank.
	BankName string `json:"bank_name"`
	// The BIC of the SEPA Bank Account.
	BIC string `json:"bic"`
	// The originating country of the SEPA Bank account.
	Country string `json:"country"`
	// The IBAN of the SEPA Bank Account.
	IBAN string `json:"iban"`
	// The last four digits of the SEPA Bank Account number. This will always be returned.
	// To view the full account number when retrieving or listing FinancialAddresses, use the `include` request parameter.
	Last4 string `json:"last4"`
}

// The credentials of the US Bank Account for the FinancialAddress. This contains unique banking details such as the routing number, account number, etc. of a US bank account.
type V2MoneyManagementFinancialAddressCredentialsUSBankAccount struct {
	// The account number of the US Bank Account.
	AccountNumber string `json:"account_number,omitempty"`
	// The name of the Bank.
	BankName string `json:"bank_name,omitempty"`
	// The last four digits of the US Bank Account number. This will always be returned.
	// To view the full account number when retrieving or listing FinancialAddresses, use the `include` request parameter.
	Last4 string `json:"last4"`
	// The routing number of the US Bank Account.
	RoutingNumber string `json:"routing_number"`
	// The swift code of the bank or financial institution.
	SwiftCode string `json:"swift_code,omitempty"`
}

// Object indicates the type of credentials that have been allocated and attached to the FinancialAddress.
// It contains all necessary banking details with which to perform money movements with the FinancialAddress.
// This field is only available for FinancialAddresses with an active status.
type V2MoneyManagementFinancialAddressCredentials struct {
	// The credentials of the UK Bank Account for the FinancialAddress. This contains unique banking details such as the sort code, account number, etc. of a UK bank account.
	GBBankAccount *V2MoneyManagementFinancialAddressCredentialsGBBankAccount `json:"gb_bank_account,omitempty"`
	// The credentials of the SEPA Bank Account for the FinancialAddress. This contains unique banking details such as the IBAN, BIC, etc. of a SEPA bank account.
	SEPABankAccount *V2MoneyManagementFinancialAddressCredentialsSEPABankAccount `json:"sepa_bank_account,omitempty"`
	// Open Enum. The type of Credentials that are provisioned for the FinancialAddress.
	Type V2MoneyManagementFinancialAddressCredentialsType `json:"type"`
	// The credentials of the US Bank Account for the FinancialAddress. This contains unique banking details such as the routing number, account number, etc. of a US bank account.
	USBankAccount *V2MoneyManagementFinancialAddressCredentialsUSBankAccount `json:"us_bank_account,omitempty"`
}

// A FinancialAddress contains information needed to transfer money to a Financial Account. A Financial Account can have more than one Financial Address.
type V2MoneyManagementFinancialAddress struct {
	APIResource
	// The creation timestamp of the FinancialAddress.
	Created time.Time `json:"created"`
	// Object indicates the type of credentials that have been allocated and attached to the FinancialAddress.
	// It contains all necessary banking details with which to perform money movements with the FinancialAddress.
	// This field is only available for FinancialAddresses with an active status.
	Credentials *V2MoneyManagementFinancialAddressCredentials `json:"credentials,omitempty"`
	// Open Enum. The currency the FinancialAddress supports.
	Currency Currency `json:"currency"`
	// A ID of the FinancialAccount this FinancialAddress corresponds to.
	FinancialAccount string `json:"financial_account"`
	// The ID of a FinancialAddress.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount.
	SettlementCurrency Currency `json:"settlement_currency,omitempty"`
	// Closed Enum. An enum representing the status of the FinancialAddress. This indicates whether or not the FinancialAddress can be used for any money movement flows.
	Status V2MoneyManagementFinancialAddressStatus `json:"status"`
}
