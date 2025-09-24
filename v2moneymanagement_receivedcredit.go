//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The status of the ReceivedCredit.
type V2MoneyManagementReceivedCreditStatus string

// List of values that V2MoneyManagementReceivedCreditStatus can take
const (
	V2MoneyManagementReceivedCreditStatusFailed    V2MoneyManagementReceivedCreditStatus = "failed"
	V2MoneyManagementReceivedCreditStatusPending   V2MoneyManagementReceivedCreditStatus = "pending"
	V2MoneyManagementReceivedCreditStatusReturned  V2MoneyManagementReceivedCreditStatus = "returned"
	V2MoneyManagementReceivedCreditStatusSucceeded V2MoneyManagementReceivedCreditStatus = "succeeded"
)

// Open Enum. The `failed` status reason.
type V2MoneyManagementReceivedCreditStatusDetailsFailedReason string

// List of values that V2MoneyManagementReceivedCreditStatusDetailsFailedReason can take
const (
	V2MoneyManagementReceivedCreditStatusDetailsFailedReasonCapabilityInactive                    V2MoneyManagementReceivedCreditStatusDetailsFailedReason = "capability_inactive"
	V2MoneyManagementReceivedCreditStatusDetailsFailedReasonCurrencyUnsupportedOnFinancialAddress V2MoneyManagementReceivedCreditStatusDetailsFailedReason = "currency_unsupported_on_financial_address"
	V2MoneyManagementReceivedCreditStatusDetailsFailedReasonFinancialAddressInactive              V2MoneyManagementReceivedCreditStatusDetailsFailedReason = "financial_address_inactive"
	V2MoneyManagementReceivedCreditStatusDetailsFailedReasonStripeRejected                        V2MoneyManagementReceivedCreditStatusDetailsFailedReason = "stripe_rejected"
)

// Open Enum. The `returned` status reason.
type V2MoneyManagementReceivedCreditStatusDetailsReturnedReason string

// List of values that V2MoneyManagementReceivedCreditStatusDetailsReturnedReason can take
const (
	V2MoneyManagementReceivedCreditStatusDetailsReturnedReasonOriginatorInitiatedReversal V2MoneyManagementReceivedCreditStatusDetailsReturnedReason = "originator_initiated_reversal"
)

// Open Enum. The type of flow that caused the ReceivedCredit.
type V2MoneyManagementReceivedCreditType string

// List of values that V2MoneyManagementReceivedCreditType can take
const (
	V2MoneyManagementReceivedCreditTypeBalanceTransfer V2MoneyManagementReceivedCreditType = "balance_transfer"
	V2MoneyManagementReceivedCreditTypeBankTransfer    V2MoneyManagementReceivedCreditType = "bank_transfer"
	V2MoneyManagementReceivedCreditTypeExternalCredit  V2MoneyManagementReceivedCreditType = "external_credit"
)

// Open Enum. The type of Stripe Money Movement that originated the ReceivedCredit.
type V2MoneyManagementReceivedCreditBalanceTransferType string

// List of values that V2MoneyManagementReceivedCreditBalanceTransferType can take
const (
	V2MoneyManagementReceivedCreditBalanceTransferTypeOutboundPayment  V2MoneyManagementReceivedCreditBalanceTransferType = "outbound_payment"
	V2MoneyManagementReceivedCreditBalanceTransferTypeOutboundTransfer V2MoneyManagementReceivedCreditBalanceTransferType = "outbound_transfer"
	V2MoneyManagementReceivedCreditBalanceTransferTypePayoutV1         V2MoneyManagementReceivedCreditBalanceTransferType = "payout_v1"
)

// Open Enum. Indicates the origin of source from which external funds originated from.
type V2MoneyManagementReceivedCreditBankTransferOriginType string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginType can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginTypeGBBankAccount   V2MoneyManagementReceivedCreditBankTransferOriginType = "gb_bank_account"
	V2MoneyManagementReceivedCreditBankTransferOriginTypeSEPABankAccount V2MoneyManagementReceivedCreditBankTransferOriginType = "sepa_bank_account"
	V2MoneyManagementReceivedCreditBankTransferOriginTypeUSBankAccount   V2MoneyManagementReceivedCreditBankTransferOriginType = "us_bank_account"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetworkFPS V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork = "fps"
)

// The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferSEPABankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferSEPABankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferSEPABankAccountNetworkSEPACreditTransfer V2MoneyManagementReceivedCreditBankTransferSEPABankAccountNetwork = "sepa_credit_transfer"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetworkACH            V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork = "ach"
	V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetworkRTP            V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork = "rtp"
	V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetworkUSDomesticWire V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork = "us_domestic_wire"
)

// Hash that provides additional information regarding the reason behind a `failed` ReceivedCredit status. It is only present when the ReceivedCredit status is `failed`.
type V2MoneyManagementReceivedCreditStatusDetailsFailed struct {
	// Open Enum. The `failed` status reason.
	Reason V2MoneyManagementReceivedCreditStatusDetailsFailedReason `json:"reason"`
}

// Hash that provides additional information regarding the reason behind a `returned` ReceivedCredit status. It is only present when the ReceivedCredit status is `returned`.
type V2MoneyManagementReceivedCreditStatusDetailsReturned struct {
	// Open Enum. The `returned` status reason.
	Reason V2MoneyManagementReceivedCreditStatusDetailsReturnedReason `json:"reason"`
}

// This hash contains detailed information that elaborates on the specific status of the ReceivedCredit. e.g the reason behind a failure if the status is marked as `failed`.
type V2MoneyManagementReceivedCreditStatusDetails struct {
	// Hash that provides additional information regarding the reason behind a `failed` ReceivedCredit status. It is only present when the ReceivedCredit status is `failed`.
	Failed *V2MoneyManagementReceivedCreditStatusDetailsFailed `json:"failed,omitempty"`
	// Hash that provides additional information regarding the reason behind a `returned` ReceivedCredit status. It is only present when the ReceivedCredit status is `returned`.
	Returned *V2MoneyManagementReceivedCreditStatusDetailsReturned `json:"returned,omitempty"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementReceivedCreditStatusTransitions struct {
	// Timestamp describing when the ReceivedCredit was marked as `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// Timestamp describing when the ReceivedCredit changed status to `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ReturnedAt time.Time `json:"returned_at,omitempty"`
	// Timestamp describing when the ReceivedCredit was marked as `succeeded`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// This object stores details about the originating Stripe transaction that resulted in the ReceivedCredit. Present if `type` field value is `balance_transfer`.
type V2MoneyManagementReceivedCreditBalanceTransfer struct {
	// The ID of the account that owns the source object originated the ReceivedCredit.
	FromAccount string `json:"from_account,omitempty"`
	// The ID of the outbound payment object that originated the ReceivedCredit.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// The ID of the outbound transfer object that originated the ReceivedCredit.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// The ID of the payout object that originated the ReceivedCredit.
	PayoutV1 string `json:"payout_v1,omitempty"`
	// Open Enum. The type of Stripe Money Movement that originated the ReceivedCredit.
	Type V2MoneyManagementReceivedCreditBalanceTransferType `json:"type"`
}

// Hash containing the transaction bank details. Present if `origin_type` field value is `gb_bank_account`.
type V2MoneyManagementReceivedCreditBankTransferGBBankAccount struct {
	// The bank name the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork `json:"network"`
	// The sort code of the account that originated the transfer.
	SortCode string `json:"sort_code,omitempty"`
}

// Hash containing the transaction bank details. Present if `origin_type` field value is `sepa_bank_account`.
type V2MoneyManagementReceivedCreditBankTransferSEPABankAccount struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The BIC of the SEPA account.
	BIC string `json:"bic,omitempty"`
	// The origination country of the bank transfer.
	Country string `json:"country,omitempty"`
	// The IBAN that originated the transfer.
	IBAN string `json:"iban,omitempty"`
	// The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferSEPABankAccountNetwork `json:"network"`
}

// Hash containing the transaction bank details. Present if `origin_type` field value is `us_bank_account`.
type V2MoneyManagementReceivedCreditBankTransferUSBankAccount struct {
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferUSBankAccountNetwork `json:"network"`
	// The routing number of the account that originated the transfer.
	RoutingNumber string `json:"routing_number,omitempty"`
}

// This object stores details about the originating banking transaction that resulted in the ReceivedCredit. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementReceivedCreditBankTransfer struct {
	// Financial Address on which funds for ReceivedCredit were received.
	FinancialAddress string `json:"financial_address"`
	// Hash containing the transaction bank details. Present if `origin_type` field value is `gb_bank_account`.
	GBBankAccount *V2MoneyManagementReceivedCreditBankTransferGBBankAccount `json:"gb_bank_account,omitempty"`
	// Open Enum. Indicates the origin of source from which external funds originated from.
	OriginType V2MoneyManagementReceivedCreditBankTransferOriginType `json:"origin_type"`
	// Hash containing the transaction bank details. Present if `origin_type` field value is `sepa_bank_account`.
	SEPABankAccount *V2MoneyManagementReceivedCreditBankTransferSEPABankAccount `json:"sepa_bank_account,omitempty"`
	// Freeform string set by originator of the external ReceivedCredit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Hash containing the transaction bank details. Present if `origin_type` field value is `us_bank_account`.
	USBankAccount *V2MoneyManagementReceivedCreditBankTransferUSBankAccount `json:"us_bank_account,omitempty"`
}

// Use ReceivedCredits API to retrieve information on when, where, and how funds are sent into your FinancialAccount.
type V2MoneyManagementReceivedCredit struct {
	APIResource
	// The amount and currency of the ReceivedCredit.
	Amount Amount `json:"amount"`
	// This object stores details about the originating Stripe transaction that resulted in the ReceivedCredit. Present if `type` field value is `balance_transfer`.
	BalanceTransfer *V2MoneyManagementReceivedCreditBalanceTransfer `json:"balance_transfer,omitempty"`
	// This object stores details about the originating banking transaction that resulted in the ReceivedCredit. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedCreditBankTransfer `json:"bank_transfer,omitempty"`
	// Time at which the ReceivedCredit was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Freeform string set by originator of the ReceivedCredit.
	Description string `json:"description,omitempty"`
	// Financial Account ID on which funds for ReceivedCredit were received.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the ReceivedCredit.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A hosted transaction receipt URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	ReceiptURL string `json:"receipt_url,omitempty"`
	// Open Enum. The status of the ReceivedCredit.
	Status V2MoneyManagementReceivedCreditStatus `json:"status"`
	// This hash contains detailed information that elaborates on the specific status of the ReceivedCredit. e.g the reason behind a failure if the status is marked as `failed`.
	StatusDetails *V2MoneyManagementReceivedCreditStatusDetails `json:"status_details,omitempty"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementReceivedCreditStatusTransitions `json:"status_transitions,omitempty"`
	// Open Enum. The type of flow that caused the ReceivedCredit.
	Type V2MoneyManagementReceivedCreditType `json:"type"`
}
