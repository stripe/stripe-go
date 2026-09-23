//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves the wallet export metadata for a closed FinancialAccount. Credentials are returned only by the export_credentials action.
type V2MoneyManagementFinancialAccountsWalletExportParams struct {
	Params `form:"*"`
}

// Encryption parameters for the exported credentials.
type V2MoneyManagementFinancialAccountsWalletExportExportCredentialsEncryptionParams struct {
	// Base64url-encoded raw P-256 recipient public key. Stripe does not persist this key material.
	RecipientPublicKey *string `form:"recipient_public_key" json:"recipient_public_key"`
	// Encryption scheme for the response. HPKE uses BASE mode, DHKEM_P256_HKDF_SHA256, HKDF_SHA256, and CHACHA20_POLY1305.
	Type *string `form:"type" json:"type"`
}

// Exports wallet credentials encrypted to the supplied recipient key. The first successful request starts one fixed one-hour retrieval window; later requests may use a different recipient key without extending it.
type V2MoneyManagementFinancialAccountsWalletExportExportCredentialsParams struct {
	Params `form:"*"`
	// Encryption parameters for the exported credentials.
	Encryption *V2MoneyManagementFinancialAccountsWalletExportExportCredentialsEncryptionParams `form:"encryption" json:"encryption"`
}

// Retrieves the wallet export metadata for a closed FinancialAccount. Credentials are returned only by the export_credentials action.
type V2MoneyManagementFinancialAccountsWalletExportRetrieveParams struct {
	Params `form:"*"`
}
