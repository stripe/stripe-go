//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of bank account (checking or savings).
type V2MoneyManagementPayoutMethodBankAccountBankAccountType string

// List of values that V2MoneyManagementPayoutMethodBankAccountBankAccountType can take
const (
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeChecking V2MoneyManagementPayoutMethodBankAccountBankAccountType = "checking"
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeFutsu    V2MoneyManagementPayoutMethodBankAccountBankAccountType = "futsu"
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeSavings  V2MoneyManagementPayoutMethodBankAccountBankAccountType = "savings"
	V2MoneyManagementPayoutMethodBankAccountBankAccountTypeToza     V2MoneyManagementPayoutMethodBankAccountBankAccountType = "toza"
)

// Open Enum. The type of payout method.
type V2MoneyManagementPayoutMethodType string

// List of values that V2MoneyManagementPayoutMethodType can take
const (
	V2MoneyManagementPayoutMethodTypeBankAccount V2MoneyManagementPayoutMethodType = "bank_account"
	V2MoneyManagementPayoutMethodTypeCard        V2MoneyManagementPayoutMethodType = "card"
)

// Payments status - used when sending OutboundPayments (sending funds to recipients).
// If disabled, enable the payout method by creating an OutboundSetupIntent using [`POST /v2/money_management/outbound_setup_intents`](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
type V2MoneyManagementPayoutMethodUsageStatusPayments string

// List of values that V2MoneyManagementPayoutMethodUsageStatusPayments can take
const (
	V2MoneyManagementPayoutMethodUsageStatusPaymentsDisabled       V2MoneyManagementPayoutMethodUsageStatusPayments = "disabled"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsEligible       V2MoneyManagementPayoutMethodUsageStatusPayments = "eligible"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsIneligible     V2MoneyManagementPayoutMethodUsageStatusPayments = "ineligible"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsInvalid        V2MoneyManagementPayoutMethodUsageStatusPayments = "invalid"
	V2MoneyManagementPayoutMethodUsageStatusPaymentsRequiresAction V2MoneyManagementPayoutMethodUsageStatusPayments = "requires_action"
)

// Transfers status - used when making an OutboundTransfer (sending funds to yourself).
// If disabled, enable the payout method by creating an OutboundSetupIntent using [`POST /v2/money_management/outbound_setup_intents`](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
type V2MoneyManagementPayoutMethodUsageStatusTransfers string

// List of values that V2MoneyManagementPayoutMethodUsageStatusTransfers can take
const (
	V2MoneyManagementPayoutMethodUsageStatusTransfersDisabled       V2MoneyManagementPayoutMethodUsageStatusTransfers = "disabled"
	V2MoneyManagementPayoutMethodUsageStatusTransfersEligible       V2MoneyManagementPayoutMethodUsageStatusTransfers = "eligible"
	V2MoneyManagementPayoutMethodUsageStatusTransfersIneligible     V2MoneyManagementPayoutMethodUsageStatusTransfers = "ineligible"
	V2MoneyManagementPayoutMethodUsageStatusTransfersInvalid        V2MoneyManagementPayoutMethodUsageStatusTransfers = "invalid"
	V2MoneyManagementPayoutMethodUsageStatusTransfersRequiresAction V2MoneyManagementPayoutMethodUsageStatusTransfers = "requires_action"
)

// The PayoutMethodBankAccount object details.
type V2MoneyManagementPayoutMethodBankAccount struct {
	// The type of bank account (checking or savings).
	BankAccountType V2MoneyManagementPayoutMethodBankAccountBankAccountType `json:"bank_account_type"`
	// The name of the bank this bank account is in. This field is populated automatically by Stripe.
	BankName string `json:"bank_name"`
	// The branch number of the bank account, if present.
	BranchNumber string `json:"branch_number,omitempty"`
	// The country code of the bank account.
	Country string `json:"country"`
	// The ID of the Financial Connections Account used to create the bank account.
	FinancialConnectionsAccount string `json:"financial_connections_account,omitempty"`
	// The last 4 digits of the account number.
	Last4 string `json:"last4"`
	// The routing number of the bank account, if present.
	RoutingNumber string `json:"routing_number,omitempty"`
	// The list of currencies supported by this bank account.
	SupportedCurrencies []Currency `json:"supported_currencies"`
	// The swift code of the bank or financial institution.
	SwiftCode string `json:"swift_code,omitempty"`
}

// The PayoutMethodCard object details.
type V2MoneyManagementPayoutMethodCard struct {
	// The month the card expires.
	ExpMonth string `json:"exp_month"`
	// The year the card expires.
	ExpYear string `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two
	// recipients who've signed up with you are using the same card number, for example.
	Fingerprint string `json:"fingerprint"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
	// The list of currencies supported by this bank account.
	SupportedCurrencies []Currency `json:"supported_currencies"`
}

// Indicates whether the payout method has met the necessary requirements for outbound money movement.
type V2MoneyManagementPayoutMethodUsageStatus struct {
	// Payments status - used when sending OutboundPayments (sending funds to recipients).
	// If disabled, enable the payout method by creating an OutboundSetupIntent using [`POST /v2/money_management/outbound_setup_intents`](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
	Payments V2MoneyManagementPayoutMethodUsageStatusPayments `json:"payments"`
	// Transfers status - used when making an OutboundTransfer (sending funds to yourself).
	// If disabled, enable the payout method by creating an OutboundSetupIntent using [`POST /v2/money_management/outbound_setup_intents`](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
	Transfers V2MoneyManagementPayoutMethodUsageStatusTransfers `json:"transfers"`
}

// Use the PayoutMethods API to list and interact with PayoutMethod objects.
type V2MoneyManagementPayoutMethod struct {
	APIResource
	// Whether the payout method was archived. Payout methods can be archived through the /archive API,
	// and they will not be automatically archived by Stripe. Archived payout methods cannot be used
	// for outbound money movement.
	Archived bool `json:"archived"`
	// The PayoutMethodBankAccount object details.
	BankAccount *V2MoneyManagementPayoutMethodBankAccount `json:"bank_account,omitempty"`
	// The PayoutMethodCard object details.
	Card *V2MoneyManagementPayoutMethodCard `json:"card,omitempty"`
	// Created timestamp.
	Created time.Time `json:"created"`
	// ID of the PayoutMethod object.
	ID string `json:"id"`
	// ID of the underlying active OutboundSetupIntent object, if any.
	LatestOutboundSetupIntent string `json:"latest_outbound_setup_intent,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Whether the Payout Method is currently unusable for money movement, despite potentially being correctly set up.
	// Please reach out to Stripe Support for more information.
	Restricted bool `json:"restricted"`
	// Open Enum. The type of payout method.
	Type V2MoneyManagementPayoutMethodType `json:"type"`
	// Indicates whether the payout method has met the necessary requirements for outbound money movement.
	UsageStatus *V2MoneyManagementPayoutMethodUsageStatus `json:"usage_status"`
}
