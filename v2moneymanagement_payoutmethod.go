//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A set of available payout speeds for this payout method.
type V2MoneyManagementPayoutMethodAvailablePayoutSpeed string

// List of values that V2MoneyManagementPayoutMethodAvailablePayoutSpeed can take
const (
	V2MoneyManagementPayoutMethodAvailablePayoutSpeedInstant  V2MoneyManagementPayoutMethodAvailablePayoutSpeed = "instant"
	V2MoneyManagementPayoutMethodAvailablePayoutSpeedStandard V2MoneyManagementPayoutMethodAvailablePayoutSpeed = "standard"
)

// Closed Enum. The type of payout method.
type V2MoneyManagementPayoutMethodType string

// List of values that V2MoneyManagementPayoutMethodType can take
const (
	V2MoneyManagementPayoutMethodTypeBankAccount  V2MoneyManagementPayoutMethodType = "bank_account"
	V2MoneyManagementPayoutMethodTypeCard         V2MoneyManagementPayoutMethodType = "card"
	V2MoneyManagementPayoutMethodTypeCryptoWallet V2MoneyManagementPayoutMethodType = "crypto_wallet"
)

// Payments status - used when sending OutboundPayments (sending funds to recipients).
type V2MoneyManagementPayoutMethodUsageStatusPayments string

// List of values that V2MoneyManagementPayoutMethodUsageStatusPayments can take
const (
	V2MoneyManagementPayoutMethodUsageStatusPaymentsEligible       V2MoneyManagementPayoutMethodUsageStatusPayments = "eligible"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsInvalid        V2MoneyManagementPayoutMethodUsageStatusPayments = "invalid"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsRequiresAction V2MoneyManagementPayoutMethodUsageStatusPayments = "requires_action"
)

// Transfers status - used when making an OutboundTransfer (sending funds to yourself).
type V2MoneyManagementPayoutMethodUsageStatusTransfers string

// List of values that V2MoneyManagementPayoutMethodUsageStatusTransfers can take
const (
	V2MoneyManagementPayoutMethodUsageStatusTransfersEligible       V2MoneyManagementPayoutMethodUsageStatusTransfers = "eligible"
	V2MoneyManagementPayoutMethodUsageStatusTransfersInvalid        V2MoneyManagementPayoutMethodUsageStatusTransfers = "invalid"
	V2MoneyManagementPayoutMethodUsageStatusTransfersRequiresAction V2MoneyManagementPayoutMethodUsageStatusTransfers = "requires_action"
)

// The type of bank account (checking or savings).
type V2MoneyManagementPayoutMethodBankAccountBankAccountType string

// List of values that V2MoneyManagementPayoutMethodBankAccountBankAccountType can take
const (
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeChecking V2MoneyManagementPayoutMethodBankAccountBankAccountType = "checking"
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeSavings  V2MoneyManagementPayoutMethodBankAccountBankAccountType = "savings"
)

// Which rail is being used to make an outbound money movement to this wallet.
type V2MoneyManagementPayoutMethodCryptoWalletNetwork string

// List of values that V2MoneyManagementPayoutMethodCryptoWalletNetwork can take
const (
	V2MoneyManagementPayoutMethodCryptoWalletNetworkArbitrum        V2MoneyManagementPayoutMethodCryptoWalletNetwork = "arbitrum"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkAvalancheCChain V2MoneyManagementPayoutMethodCryptoWalletNetwork = "avalanche_c_chain"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkBase            V2MoneyManagementPayoutMethodCryptoWalletNetwork = "base"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkEthereum        V2MoneyManagementPayoutMethodCryptoWalletNetwork = "ethereum"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkOptimism        V2MoneyManagementPayoutMethodCryptoWalletNetwork = "optimism"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkPolygon         V2MoneyManagementPayoutMethodCryptoWalletNetwork = "polygon"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkSolana          V2MoneyManagementPayoutMethodCryptoWalletNetwork = "solana"
	V2MoneyManagementPayoutMethodCryptoWalletNetworkStellar         V2MoneyManagementPayoutMethodCryptoWalletNetwork = "stellar"
)

// Indicates whether the payout method has met the necessary requirements for outbound money movement.
type V2MoneyManagementPayoutMethodUsageStatus struct {
	// Payments status - used when sending OutboundPayments (sending funds to recipients).
	Payments V2MoneyManagementPayoutMethodUsageStatusPayments `json:"payments"`
	// Transfers status - used when making an OutboundTransfer (sending funds to yourself).
	Transfers V2MoneyManagementPayoutMethodUsageStatusTransfers `json:"transfers"`
}

// The PayoutMethodBankAccount object details.
type V2MoneyManagementPayoutMethodBankAccount struct {
	// Whether this PayoutMethodBankAccount object was archived. PayoutMethodBankAccount objects can be archived through
	// the /archive API, and they will not be automatically archived by Stripe. Archived PayoutMethodBankAccount objects
	// cannot be used as payout methods and will not appear in the payout method list.
	Archived bool `json:"archived"`
	// The type of bank account (checking or savings).
	BankAccountType V2MoneyManagementPayoutMethodBankAccountBankAccountType `json:"bank_account_type"`
	// The name of the bank this bank account is in. This field is populated automatically by Stripe.
	BankName string `json:"bank_name"`
	// The country code of the bank account.
	Country string `json:"country"`
	// List of enabled flows for this bank account (wire or local).
	EnabledDeliveryOptions []string `json:"enabled_delivery_options"`
	// The last 4 digits of the account number.
	Last4 string `json:"last4"`
	// The routing number of the bank account, if present.
	RoutingNumber string `json:"routing_number,omitempty"`
	// The list of currencies supported by this bank account.
	SupportedCurrencies []Currency `json:"supported_currencies"`
}

// The PayoutMethodCard object details.
type V2MoneyManagementPayoutMethodCard struct {
	// Whether the PayoutMethodCard object was archived. PayoutMethodCard objects can be archived through
	// the /archive API, and they will not be automatically archived by Stripe. Archived PayoutMethodCard objects
	// cannot be used as payout methods and will not appear in the payout method list.
	Archived bool `json:"archived"`
	// The month the card expires.
	ExpMonth string `json:"exp_month"`
	// The year the card expires.
	ExpYear string `json:"exp_year"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
}

// The PayoutMethodCryptoWallet object details.
type V2MoneyManagementPayoutMethodCryptoWallet struct {
	// Destination wallet address.
	Address string `json:"address"`
	// Whether the crypto wallet was archived. Crypto wallets can be archived through the /archive API,
	// and they will not be automatically archived by Stripe. Archived crypto wallets cannot be used as
	// payout method and will not appear in the payout method list.
	Archived bool `json:"archived"`
	// Optional field, required if network supports memos (only "stellar" currently).
	Memo string `json:"memo,omitempty"`
	// Which rail is being used to make an outbound money movement to this wallet.
	Network V2MoneyManagementPayoutMethodCryptoWalletNetwork `json:"network"`
}

// Use the PayoutMethods API to list and interact with PayoutMethod objects.
type V2MoneyManagementPayoutMethod struct {
	APIResource
	// A set of available payout speeds for this payout method.
	AvailablePayoutSpeeds []V2MoneyManagementPayoutMethodAvailablePayoutSpeed `json:"available_payout_speeds"`
	// The PayoutMethodBankAccount object details.
	BankAccount *V2MoneyManagementPayoutMethodBankAccount `json:"bank_account,omitempty"`
	// The PayoutMethodCard object details.
	Card *V2MoneyManagementPayoutMethodCard `json:"card,omitempty"`
	// Created timestamp.
	Created time.Time `json:"created"`
	// The PayoutMethodCryptoWallet object details.
	CryptoWallet *V2MoneyManagementPayoutMethodCryptoWallet `json:"crypto_wallet,omitempty"`
	// ID of the PayoutMethod object.
	ID string `json:"id"`
	// ID of the underlying active OutboundSetupIntent object, if any.
	LatestOutboundSetupIntent string `json:"latest_outbound_setup_intent,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. The type of payout method.
	Type V2MoneyManagementPayoutMethodType `json:"type"`
	// Indicates whether the payout method has met the necessary requirements for outbound money movement.
	UsageStatus *V2MoneyManagementPayoutMethodUsageStatus `json:"usage_status"`
}
