//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List the OutboundSetupIntent objects.
type V2MoneyManagementOutboundSetupIntentListParams struct {
	Params `form:"*"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// The type specific details of the bank account payout method.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataBankAccountParams struct {
	// The account number or IBAN of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// The branch number of the bank account, if present.
	BranchNumber *string `form:"branch_number" json:"branch_number,omitempty"`
	// The country code of the bank account.
	Country *string `form:"country" json:"country"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams struct {
	// The expiration month of the card.
	ExpMonth *string `form:"exp_month" json:"exp_month,omitempty"`
	// The expiration year of the card.
	ExpYear *string `form:"exp_year" json:"exp_year,omitempty"`
	// The card number. This can only be passed when creating a new credential on an outbound setup intent in the requires_payout_method state.
	Number *string `form:"number" json:"number,omitempty"`
}

// The type specific details of the crypto wallet payout method.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataCryptoWalletParams struct {
	// Crypto wallet address.
	Address *string `form:"address" json:"address"`
	// Optional field, required if network supports memos (only "stellar" currently).
	Memo *string `form:"memo" json:"memo,omitempty"`
	// Which rail we should use to make an Outbound money movement to this wallet.
	Network *string `form:"network" json:"network"`
}

// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
// If a payout_method provided, used to update data on the credential linked to this setup intent.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams struct {
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentPayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type specific details of the crypto wallet payout method.
	CryptoWallet *V2MoneyManagementOutboundSetupIntentPayoutMethodDataCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// Closed Enum. The type of payout method to be created/updated.
	Type *string `form:"type" json:"type"`
}

// Create an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentParams struct {
	Params `form:"*"`
	// If provided, the existing payout method resource to link to this setup intent.
	// Any payout_method_data provided is used to update information on this linked payout method resource.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
	// If a payout_method provided, used to update data on the credential linked to this setup intent.
	PayoutMethodData *V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams `form:"payout_method_data" json:"payout_method_data,omitempty"`
	// Specify which type of outbound money movement this credential should be set up for (payment | transfer).
	// If not provided, defaults to payment.
	UsageIntent *string `form:"usage_intent" json:"usage_intent,omitempty"`
}

// Cancel an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentCancelParams struct {
	Params `form:"*"`
}

// The type specific details of the bank account payout method.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataBankAccountParams struct {
	// The account number or IBAN of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// The branch number of the bank account, if present.
	BranchNumber *string `form:"branch_number" json:"branch_number,omitempty"`
	// The country code of the bank account.
	Country *string `form:"country" json:"country"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCardParams struct {
	// The expiration month of the card.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The expiration year of the card.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The card number.
	Number *string `form:"number" json:"number"`
}

// The type specific details of the crypto wallet payout method.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCryptoWalletParams struct {
	// Crypto wallet address.
	Address *string `form:"address" json:"address"`
	// Optional field, required if network supports memos (only "stellar" currently).
	Memo *string `form:"memo" json:"memo,omitempty"`
	// Which rail we should use to make an Outbound money movement to this wallet.
	Network *string `form:"network" json:"network"`
}

// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
// If a payout_method provided, used to update data on the credential linked to this setup intent.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataParams struct {
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type specific details of the crypto wallet payout method.
	CryptoWallet *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// Closed Enum. The type of payout method to be created.
	Type *string `form:"type" json:"type"`
}

// Create an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentCreateParams struct {
	Params `form:"*"`
	// If provided, the existing payout method resource to link to this setup intent.
	// Any payout_method_data provided is used to update information on this linked payout method resource.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
	// If a payout_method provided, used to update data on the credential linked to this setup intent.
	PayoutMethodData *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataParams `form:"payout_method_data" json:"payout_method_data,omitempty"`
	// Specify which type of outbound money movement this credential should be set up for (payment | transfer).
	// If not provided, defaults to payment.
	UsageIntent *string `form:"usage_intent" json:"usage_intent,omitempty"`
}

// Retrieve an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentRetrieveParams struct {
	Params `form:"*"`
}

// The type specific details of the bank account payout method.
type V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataBankAccountParams struct {
	// The account number or IBAN of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// The branch number of the bank account, if present.
	BranchNumber *string `form:"branch_number" json:"branch_number,omitempty"`
	// The country code of the bank account.
	Country *string `form:"country" json:"country"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataCardParams struct {
	// The expiration month of the card.
	ExpMonth *string `form:"exp_month" json:"exp_month,omitempty"`
	// The expiration year of the card.
	ExpYear *string `form:"exp_year" json:"exp_year,omitempty"`
	// The card number. This can only be passed when creating a new credential on an outbound setup intent in the requires_payout_method state.
	Number *string `form:"number" json:"number,omitempty"`
}

// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
// If a payout_method provided, used to update data on the credential linked to this setup intent.
type V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataParams struct {
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// Closed Enum. The type of payout method to be created/updated.
	Type *string `form:"type" json:"type"`
}

// Update an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentUpdateParams struct {
	Params `form:"*"`
	// If provided, the existing payout method resource to link to this outbound setup intent.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
	// If a payout_method provided, used to update data on the credential linked to this setup intent.
	PayoutMethodData *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataParams `form:"payout_method_data" json:"payout_method_data,omitempty"`
}
