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

// An existing resource to use as the source for setting up outbound credentials.
type V2MoneyManagementOutboundSetupIntentFromResourceParams struct {
	// The identifier of the source resource.
	ID *string `form:"id" json:"id"`
	// The type of the source resource.
	Type *string `form:"type" json:"type"`
}

// The type specific details of the Apple Pay payout method.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataApplePayParams struct {
	// The paymentData property of the Apple-provided PKPaymentToken (or ApplePayPaymentToken, for Apple Pay on the Web) as a UTF-8 encoded serialization of a JSON dictionary.
	PkToken *string `form:"pk_token" json:"pk_token,omitempty"`
	// The paymentMethod.displayName property of the Apple-provided PKPaymentToken (or ApplePayPaymentToken, for Apple Pay on the Web), e.g. "Visa 1234".
	PkTokenDisplayName *string `form:"pk_token_display_name" json:"pk_token_display_name"`
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
	// The currency of the bank account.
	Currency *string `form:"currency" json:"currency"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams struct {
	// The currency of the card.
	Currency *string `form:"currency" json:"currency,omitempty"`
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
	// The type specific details of the Apple Pay payout method.
	ApplePay *V2MoneyManagementOutboundSetupIntentPayoutMethodDataApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentPayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type specific details of the crypto wallet payout method.
	CryptoWallet *V2MoneyManagementOutboundSetupIntentPayoutMethodDataCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// Open Enum. The type of payout method to be created/updated.
	Type *string `form:"type" json:"type"`
}

// Create an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentParams struct {
	Params `form:"*"`
	// An existing resource to use as the source for setting up outbound credentials.
	FromResource *V2MoneyManagementOutboundSetupIntentFromResourceParams `form:"from_resource" json:"from_resource,omitempty"`
	// If provided, the existing payout method resource to link to this setup intent.
	// Any payout_method_data provided is used to update information on this linked payout method resource.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
	// If a payout_method provided, used to update data on the credential linked to this setup intent.
	// Only card expiry (exp_month, exp_year) can be updated in the case where payout_method is provided.
	PayoutMethodData *V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams `form:"payout_method_data" json:"payout_method_data,omitempty"`
	// Specify which type of outbound money movement this credential should be set up for (payment | transfer).
	// If not provided, defaults to payment.
	UsageIntent *string `form:"usage_intent" json:"usage_intent,omitempty"`
}

// Cancel an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentCancelParams struct {
	Params `form:"*"`
}

// An existing resource to use as the source for setting up outbound credentials.
type V2MoneyManagementOutboundSetupIntentCreateFromResourceParams struct {
	// The identifier of the source resource.
	ID *string `form:"id" json:"id"`
	// The type of the source resource.
	Type *string `form:"type" json:"type"`
}

// The type specific details of the Apple Pay payout method.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataApplePayParams struct {
	// The paymentData property of the Apple-provided PKPaymentToken (or ApplePayPaymentToken, for Apple Pay on the Web) as a UTF-8 encoded serialization of a JSON dictionary.
	PkToken *string `form:"pk_token" json:"pk_token,omitempty"`
	// The paymentMethod.displayName property of the Apple-provided PKPaymentToken (or ApplePayPaymentToken, for Apple Pay on the Web), e.g. "Visa 1234".
	PkTokenDisplayName *string `form:"pk_token_display_name" json:"pk_token_display_name"`
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
	// The currency of the bank account.
	Currency *string `form:"currency" json:"currency"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCardParams struct {
	// The currency of the card.
	Currency *string `form:"currency" json:"currency"`
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
	// The type specific details of the Apple Pay payout method.
	ApplePay *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type specific details of the crypto wallet payout method.
	CryptoWallet *V2MoneyManagementOutboundSetupIntentCreatePayoutMethodDataCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// Open Enum. The type of payout method to be created.
	Type *string `form:"type" json:"type"`
}

// Create an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentCreateParams struct {
	Params `form:"*"`
	// An existing resource to use as the source for setting up outbound credentials.
	FromResource *V2MoneyManagementOutboundSetupIntentCreateFromResourceParams `form:"from_resource" json:"from_resource,omitempty"`
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
	// The currency of the bank account.
	Currency *string `form:"currency" json:"currency"`
	// The routing number of the bank account, if present.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
	// The swift code of the bank account, if present.
	SwiftCode *string `form:"swift_code" json:"swift_code,omitempty"`
}

// The type specific details of the card payout method.
type V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataCardParams struct {
	// The currency of the card.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The expiration month of the card.
	ExpMonth *string `form:"exp_month" json:"exp_month,omitempty"`
	// The expiration year of the card.
	ExpYear *string `form:"exp_year" json:"exp_year,omitempty"`
	// The card number. This can only be passed when creating a new credential on an outbound setup intent in the requires_payout_method state.
	Number *string `form:"number" json:"number,omitempty"`
}

// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
// If a payout_method provided, used to update data on the credential linked to this setup intent.
// Only card expiry (exp_month, exp_year) can be updated in the case where payout_method is provided.
type V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataParams struct {
	// The type specific details of the bank account payout method.
	BankAccount *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// The type specific details of the card payout method.
	Card *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataCardParams `form:"card" json:"card,omitempty"`
	// Open Enum. The type of payout method to be created/updated.
	Type *string `form:"type" json:"type"`
}

// Update an OutboundSetupIntent object.
type V2MoneyManagementOutboundSetupIntentUpdateParams struct {
	Params `form:"*"`
	// If provided, the existing payout method resource to link to this outbound setup intent.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.
	// If a payout_method provided, used to update data on the credential linked to this setup intent.
	// Only card expiry (exp_month, exp_year) can be updated in the case where payout_method is provided.
	PayoutMethodData *V2MoneyManagementOutboundSetupIntentUpdatePayoutMethodDataParams `form:"payout_method_data" json:"payout_method_data,omitempty"`
}
