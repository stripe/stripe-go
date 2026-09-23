//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Current wallet export status. The lifecycle is pending, ready, then complete.
type V2MoneyManagementFinancialAccountWalletExportStatus string

// List of values that V2MoneyManagementFinancialAccountWalletExportStatus can take
const (
	V2MoneyManagementFinancialAccountWalletExportStatusComplete V2MoneyManagementFinancialAccountWalletExportStatus = "complete"
	V2MoneyManagementFinancialAccountWalletExportStatusPending  V2MoneyManagementFinancialAccountWalletExportStatus = "pending"
	V2MoneyManagementFinancialAccountWalletExportStatusReady    V2MoneyManagementFinancialAccountWalletExportStatus = "ready"
)

// Network on which each stablecoin currency is stored. Keys are lowercase currency codes.
type V2MoneyManagementFinancialAccountWalletExportWalletCurrencyNetworks string

// List of values that V2MoneyManagementFinancialAccountWalletExportWalletCurrencyNetworks can take
const (
	V2MoneyManagementFinancialAccountWalletExportWalletCurrencyNetworksTempo V2MoneyManagementFinancialAccountWalletExportWalletCurrencyNetworks = "tempo"
)

// Network family for the wallet address.
type V2MoneyManagementFinancialAccountWalletExportWalletNetworkType string

// List of values that V2MoneyManagementFinancialAccountWalletExportWalletNetworkType can take
const (
	V2MoneyManagementFinancialAccountWalletExportWalletNetworkTypeEthereum V2MoneyManagementFinancialAccountWalletExportWalletNetworkType = "ethereum"
)

// Public wallet metadata. Null while pending or ready, and retained after the credential window expires.
type V2MoneyManagementFinancialAccountWalletExportWallet struct {
	// Public address of the exported wallet.
	Address string `json:"address"`
	// Network on which each stablecoin currency is stored. Keys are lowercase currency codes.
	CurrencyNetworks map[string]V2MoneyManagementFinancialAccountWalletExportWalletCurrencyNetworks `json:"currency_networks"`
	// Network family for the wallet address.
	NetworkType V2MoneyManagementFinancialAccountWalletExportWalletNetworkType `json:"network_type"`
}

// The singleton wallet export for a FinancialAccount.
type V2MoneyManagementFinancialAccountWalletExport struct {
	APIResource
	// End of the fixed one-hour credentials retrieval window. Null until the first successful credential export; remains readable after expiry.
	CredentialsAvailableUntil time.Time `json:"credentials_available_until,omitempty"`
	// FinancialAccount whose wallet is being exported.
	FinancialAccount string `json:"financial_account"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Current wallet export status. The lifecycle is pending, ready, then complete.
	Status V2MoneyManagementFinancialAccountWalletExportStatus `json:"status"`
	// Public wallet metadata. Null while pending or ready, and retained after the credential window expires.
	Wallets []*V2MoneyManagementFinancialAccountWalletExportWallet `json:"wallets,omitempty"`
}
