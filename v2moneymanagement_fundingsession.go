//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The types of FinancialAddress that can be funded in this session.
type V2MoneyManagementFundingSessionFinancialAddressType string

// List of values that V2MoneyManagementFundingSessionFinancialAddressType can take
const (
	V2MoneyManagementFundingSessionFinancialAddressTypeBankAccount  V2MoneyManagementFundingSessionFinancialAddressType = "bank_account"
	V2MoneyManagementFundingSessionFinancialAddressTypeCryptoWallet V2MoneyManagementFundingSessionFinancialAddressType = "crypto_wallet"
)

// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
type V2MoneyManagementFundingSessionFinancialAddressOptionsCryptoWallet struct {
	// Open Enum. The currency the crypto wallet FinancialAddress settles into the FinancialAccount. Required.
	SettlementCurrency Currency `json:"settlement_currency"`
}

// Per-type options used when creating the FinancialAddress.
type V2MoneyManagementFundingSessionFinancialAddressOptions struct {
	// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
	CryptoWallet *V2MoneyManagementFundingSessionFinancialAddressOptionsCryptoWallet `json:"crypto_wallet,omitempty"`
}

// A FundingSession is a hosted funding surface for a customer to fund a FinancialAccount.
type V2MoneyManagementFundingSession struct {
	APIResource
	// The ID of the Account that owns the FinancialAccount.
	Account string `json:"account"`
	// The creation timestamp of the FundingSession.
	Created time.Time `json:"created"`
	// The ID of the FinancialAccount this FundingSession funds.
	FinancialAccount string `json:"financial_account"`
	// Per-type options used when creating the FinancialAddress.
	FinancialAddressOptions *V2MoneyManagementFundingSessionFinancialAddressOptions `json:"financial_address_options"`
	// Open Enum. The types of FinancialAddress that can be funded in this session.
	FinancialAddressTypes []V2MoneyManagementFundingSessionFinancialAddressType `json:"financial_address_types"`
	// The ID of the FundingSession. ID prefix: `fndsess`.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The URL the customer is redirected to after completing (or abandoning) the funding session.
	ReturnURL string `json:"return_url"`
	// The short-lived hosted funding URL the customer visits to fund the FinancialAccount.
	URL string `json:"url"`
}
