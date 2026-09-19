//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Encryption scheme used for these credentials.
type V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncryptedType string

// List of values that V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncryptedType can take
const (
	V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncryptedTypeHpke V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncryptedType = "hpke"
)

// Tempo network configured for each stablecoin currency. Keys are lowercase currency codes.
type V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCurrencyNetworks string

// List of values that V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCurrencyNetworks can take
const (
	V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCurrencyNetworksTempo V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCurrencyNetworks = "tempo"
)

// Network family for the wallet address.
type V2MoneyManagementFinancialAccountWalletExportCredentialsWalletNetworkType string

// List of values that V2MoneyManagementFinancialAccountWalletExportCredentialsWalletNetworkType can take
const (
	V2MoneyManagementFinancialAccountWalletExportCredentialsWalletNetworkTypeEthereum V2MoneyManagementFinancialAccountWalletExportCredentialsWalletNetworkType = "ethereum"
)

// Credentials encrypted to the supplied recipient public key.
type V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncrypted struct {
	// Base64url-encoded encrypted wallet credentials. Stripe does not persist this response.
	Ciphertext string `json:"ciphertext"`
	// Base64url-encoded HPKE encapsulated key.
	EncapsulatedKey string `json:"encapsulated_key"`
	// Encryption scheme used for these credentials.
	Type V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncryptedType `json:"type"`
}

// Exported wallets and credentials encrypted to the supplied recipient public key.
type V2MoneyManagementFinancialAccountWalletExportCredentialsWallet struct {
	// Public address of the exported wallet.
	Address string `json:"address"`
	// Credentials encrypted to the supplied recipient public key.
	CredentialsEncrypted *V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCredentialsEncrypted `json:"credentials_encrypted"`
	// Tempo network configured for each stablecoin currency. Keys are lowercase currency codes.
	CurrencyNetworks map[string]V2MoneyManagementFinancialAccountWalletExportCredentialsWalletCurrencyNetworks `json:"currency_networks"`
	// Network family for the wallet address.
	NetworkType V2MoneyManagementFinancialAccountWalletExportCredentialsWalletNetworkType `json:"network_type"`
}

// Credentials exported from a FinancialAccount wallet export.
type V2MoneyManagementFinancialAccountWalletExportCredentials struct {
	APIResource
	// End of the fixed one-hour credentials retrieval window.
	CredentialsAvailableUntil time.Time `json:"credentials_available_until"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Exported wallets and credentials encrypted to the supplied recipient public key.
	Wallets []*V2MoneyManagementFinancialAccountWalletExportCredentialsWallet `json:"wallets"`
}
