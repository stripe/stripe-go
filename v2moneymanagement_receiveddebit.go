//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of balance transfer that originated the ReceivedDebit.
type V2MoneyManagementReceivedDebitBalanceTransferType string

// List of values that V2MoneyManagementReceivedDebitBalanceTransferType can take
const (
	V2MoneyManagementReceivedDebitBalanceTransferTypeTopup V2MoneyManagementReceivedDebitBalanceTransferType = "topup"
)

// Open Enum. The bank network the debit was originated on.
type V2MoneyManagementReceivedDebitBankTransferGBBankAccountNetwork string

// List of values that V2MoneyManagementReceivedDebitBankTransferGBBankAccountNetwork can take
const (
	V2MoneyManagementReceivedDebitBankTransferGBBankAccountNetworkBACS V2MoneyManagementReceivedDebitBankTransferGBBankAccountNetwork = "bacs"
)

// Open Enum. The standard entry class code for the ACH debit.
type V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode string

// List of values that V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode can take
const (
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeCcd V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "ccd"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeCie V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "cie"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeCtx V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "ctx"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeIat V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "iat"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodePos V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "pos"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodePpd V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "ppd"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeTel V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "tel"
	V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCodeWeb V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode = "web"
)

// Open Enum. Indicates the origin type through which this debit was initiated.
type V2MoneyManagementReceivedDebitBankTransferOriginType string

// List of values that V2MoneyManagementReceivedDebitBankTransferOriginType can take
const (
	V2MoneyManagementReceivedDebitBankTransferOriginTypeGBBankAccount V2MoneyManagementReceivedDebitBankTransferOriginType = "gb_bank_account"
	V2MoneyManagementReceivedDebitBankTransferOriginTypeUSBankAccount V2MoneyManagementReceivedDebitBankTransferOriginType = "us_bank_account"
)

// Open Enum. The type of the payment method used to originate the debit.
type V2MoneyManagementReceivedDebitBankTransferPaymentMethodType string

// List of values that V2MoneyManagementReceivedDebitBankTransferPaymentMethodType can take
const (
	V2MoneyManagementReceivedDebitBankTransferPaymentMethodTypeGBBankAccount V2MoneyManagementReceivedDebitBankTransferPaymentMethodType = "gb_bank_account"
	V2MoneyManagementReceivedDebitBankTransferPaymentMethodTypeUSBankAccount V2MoneyManagementReceivedDebitBankTransferPaymentMethodType = "us_bank_account"
)

// Open Enum. The bank network the debit was originated on.
type V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork string

// List of values that V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork can take
const (
	V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetworkACH V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork = "ach"
)

// Open Enum. The status of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatus string

// List of values that V2MoneyManagementReceivedDebitStatus can take
const (
	V2MoneyManagementReceivedDebitStatusCanceled  V2MoneyManagementReceivedDebitStatus = "canceled"
	V2MoneyManagementReceivedDebitStatusFailed    V2MoneyManagementReceivedDebitStatus = "failed"
	V2MoneyManagementReceivedDebitStatusPending   V2MoneyManagementReceivedDebitStatus = "pending"
	V2MoneyManagementReceivedDebitStatusReturned  V2MoneyManagementReceivedDebitStatus = "returned"
	V2MoneyManagementReceivedDebitStatusScheduled V2MoneyManagementReceivedDebitStatus = "scheduled"
	V2MoneyManagementReceivedDebitStatusSucceeded V2MoneyManagementReceivedDebitStatus = "succeeded"
)

// Open Enum. The reason for the failure of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatusDetailsFailedReason string

// List of values that V2MoneyManagementReceivedDebitStatusDetailsFailedReason can take
const (
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonCapabilityInactive       V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "capability_inactive"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonFinancialAddressInactive V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "financial_address_inactive"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonInsufficientFunds        V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "insufficient_funds"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonNoMandate                V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "no_mandate"
	V2MoneyManagementReceivedDebitStatusDetailsFailedReasonStripeRejected           V2MoneyManagementReceivedDebitStatusDetailsFailedReason = "stripe_rejected"
)

// Open Enum. The reason the ReceivedDebit was returned.
type V2MoneyManagementReceivedDebitStatusDetailsReturnedReason string

// List of values that V2MoneyManagementReceivedDebitStatusDetailsReturnedReason can take
const (
	V2MoneyManagementReceivedDebitStatusDetailsReturnedReasonOriginatorInitiated V2MoneyManagementReceivedDebitStatusDetailsReturnedReason = "originator_initiated"
)

// Open Enum. The type of the ReceivedDebit.
type V2MoneyManagementReceivedDebitType string

// List of values that V2MoneyManagementReceivedDebitType can take
const (
	V2MoneyManagementReceivedDebitTypeBalanceTransfer      V2MoneyManagementReceivedDebitType = "balance_transfer"
	V2MoneyManagementReceivedDebitTypeBankTransfer         V2MoneyManagementReceivedDebitType = "bank_transfer"
	V2MoneyManagementReceivedDebitTypeCardSpend            V2MoneyManagementReceivedDebitType = "card_spend"
	V2MoneyManagementReceivedDebitTypeExternalDebit        V2MoneyManagementReceivedDebitType = "external_debit"
	V2MoneyManagementReceivedDebitTypeStripeBalancePayment V2MoneyManagementReceivedDebitType = "stripe_balance_payment"
)

// This object stores details about the balance transfer object that resulted in the ReceivedDebit.
type V2MoneyManagementReceivedDebitBalanceTransfer struct {
	// The ID of the v1 account that received the balance transfer.
	ToAccount string `json:"to_account,omitempty"`
	// The ID of the topup object that originated the ReceivedDebit.
	Topup string `json:"topup,omitempty"`
	// Open Enum. The type of balance transfer that originated the ReceivedDebit.
	Type V2MoneyManagementReceivedDebitBalanceTransferType `json:"type"`
}

// Object containing details of the GB Bank Account that originated the debit.
// Present when the debit was originated via BACS.
type V2MoneyManagementReceivedDebitBankTransferGBBankAccount struct {
	// The name of the account holder that originated the debit.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The name of the bank the debit originated from.
	BankName string `json:"bank_name,omitempty"`
	// Last 4 digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The bank network the debit was originated on.
	Network V2MoneyManagementReceivedDebitBankTransferGBBankAccountNetwork `json:"network"`
	// The ID of the mandate associated with this debit.
	ReceivedDebitMandate string `json:"received_debit_mandate,omitempty"`
	// The sort code of the bank that originated the debit.
	SortCode string `json:"sort_code,omitempty"`
}

// ACH-specific network details.
type V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACH struct {
	// Additional information included with the ACH debit.
	Addenda string `json:"addenda,omitempty"`
	// The entry description supplied by the company that originated the ACH debit.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,omitempty"`
	// The identifier of the company that originated the ACH debit.
	OriginatorCompanyID string `json:"originator_company_id,omitempty"`
	// The name of the company that originated the ACH debit.
	OriginatorCompanyName string `json:"originator_company_name,omitempty"`
	// The identifier assigned to the receiver of the ACH debit.
	ReceiverIDNumber string `json:"receiver_id_number,omitempty"`
	// The name of the receiver of the ACH debit.
	ReceiverName string `json:"receiver_name,omitempty"`
	// Open Enum. The standard entry class code for the ACH debit.
	StandardEntryClassCode V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACHStandardEntryClassCode `json:"standard_entry_class_code,omitempty"`
	// The trace identifier for the ACH debit.
	TraceID string `json:"trace_id,omitempty"`
}

// Network-specific details about the bank transfer.
type V2MoneyManagementReceivedDebitBankTransferNetworkDetails struct {
	// ACH-specific network details.
	ACH *V2MoneyManagementReceivedDebitBankTransferNetworkDetailsACH `json:"ach,omitempty"`
}

// Object containing details of the US Bank Account that originated the debit.
// Present when the debit was originated via ACH.
type V2MoneyManagementReceivedDebitBankTransferUSBankAccount struct {
	// The name of the bank the debit originated from.
	BankName string `json:"bank_name,omitempty"`
	// Open Enum. The bank network the debit was originated on.
	Network V2MoneyManagementReceivedDebitBankTransferUSBankAccountNetwork `json:"network"`
	// The routing number of the bank that originated the debit.
	RoutingNumber string `json:"routing_number,omitempty"`
}

// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
type V2MoneyManagementReceivedDebitBankTransfer struct {
	// The Financial Address that was debited.
	FinancialAddress string `json:"financial_address"`
	// Object containing details of the GB Bank Account that originated the debit.
	// Present when the debit was originated via BACS.
	GBBankAccount *V2MoneyManagementReceivedDebitBankTransferGBBankAccount `json:"gb_bank_account,omitempty"`
	// Network-specific details about the bank transfer.
	NetworkDetails *V2MoneyManagementReceivedDebitBankTransferNetworkDetails `json:"network_details,omitempty"`
	// Open Enum. Indicates the origin type through which this debit was initiated.
	OriginType V2MoneyManagementReceivedDebitBankTransferOriginType `json:"origin_type"`
	// Open Enum. The type of the payment method used to originate the debit.
	PaymentMethodType V2MoneyManagementReceivedDebitBankTransferPaymentMethodType `json:"payment_method_type"`
	// The statement descriptor set by the originator of the debit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Object containing details of the US Bank Account that originated the debit.
	// Present when the debit was originated via ACH.
	USBankAccount *V2MoneyManagementReceivedDebitBankTransferUSBankAccount `json:"us_bank_account,omitempty"`
}

// The Issuing Authorization for this card_spend. Contains the reference id and the amount.
type V2MoneyManagementReceivedDebitCardSpendAuthorization struct {
	// Amount associated with this issuing authorization.
	Amount Amount `json:"amount"`
	// The reference to the v1 issuing authorization ID.
	IssuingAuthorizationV1 string `json:"issuing_authorization_v1"`
}

// The list of card spend transactions. These contain the transaction reference ID and the amount.
type V2MoneyManagementReceivedDebitCardSpendCardTransaction struct {
	// Amount associated with this issuing transaction.
	Amount Amount `json:"amount"`
	// The reference to the v1 issuing transaction ID.
	IssuingTransactionV1 string `json:"issuing_transaction_v1"`
}

// This object stores details about the issuing transactions that resulted in the ReceivedDebit. Present if `type` field value is `card_spend`.
type V2MoneyManagementReceivedDebitCardSpend struct {
	// The Issuing Authorization for this card_spend. Contains the reference id and the amount.
	Authorization *V2MoneyManagementReceivedDebitCardSpendAuthorization `json:"authorization,omitempty"`
	// The list of card spend transactions. These contain the transaction reference ID and the amount.
	CardTransactions []*V2MoneyManagementReceivedDebitCardSpendCardTransaction `json:"card_transactions"`
	// The reference to the card object that resulted in the debit.
	CardV1ID string `json:"card_v1_id"`
}

// The dispute details.
type V2MoneyManagementReceivedDebitDisputeDetails struct {
	// The ID of the debit dispute, if one has been created.
	DebitDispute string `json:"debit_dispute,omitempty"`
	// The time at which the dispute window closes.
	DisputeWindowClosesAt time.Time `json:"dispute_window_closes_at,omitempty"`
}

// Information that elaborates on the `failed` status of a ReceivedDebit.
// It is only present when the ReceivedDebit status is `failed`.
type V2MoneyManagementReceivedDebitStatusDetailsFailed struct {
	// Open Enum. The reason for the failure of the ReceivedDebit.
	Reason V2MoneyManagementReceivedDebitStatusDetailsFailedReason `json:"reason"`
}

// Information that elaborates on the `returned` status of a ReceivedDebit.
// It is only present when the ReceivedDebit status is `returned`.
type V2MoneyManagementReceivedDebitStatusDetailsReturned struct {
	// Open Enum. The reason the ReceivedDebit was returned.
	Reason V2MoneyManagementReceivedDebitStatusDetailsReturnedReason `json:"reason"`
}

// Detailed information about the status of the ReceivedDebit.
type V2MoneyManagementReceivedDebitStatusDetails struct {
	// Information that elaborates on the `failed` status of a ReceivedDebit.
	// It is only present when the ReceivedDebit status is `failed`.
	Failed *V2MoneyManagementReceivedDebitStatusDetailsFailed `json:"failed"`
	// Information that elaborates on the `returned` status of a ReceivedDebit.
	// It is only present when the ReceivedDebit status is `returned`.
	Returned *V2MoneyManagementReceivedDebitStatusDetailsReturned `json:"returned"`
}

// The time at which the ReceivedDebit transitioned to a particular status.
type V2MoneyManagementReceivedDebitStatusTransitions struct {
	// The time when the ReceivedDebit was marked as `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// The time when the ReceivedDebit was marked as `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// The time when the ReceivedDebit was marked as `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	ReturnedAt time.Time `json:"returned_at,omitempty"`
	// The time when the ReceivedDebit was marked as `succeeded`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// This object stores details about the Stripe Balance Payment that resulted in the ReceivedDebit.
type V2MoneyManagementReceivedDebitStripeBalancePayment struct {
	// ID of the debit agreement associated with this payment.
	DebitAgreement string `json:"debit_agreement,omitempty"`
	// Statement descriptor for the Stripe Balance Payment.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}

// ReceivedDebit resource
type V2MoneyManagementReceivedDebit struct {
	APIResource
	// Amount and currency of the ReceivedDebit.
	Amount Amount `json:"amount"`
	// This object stores details about the balance transfer object that resulted in the ReceivedDebit.
	BalanceTransfer *V2MoneyManagementReceivedDebitBalanceTransfer `json:"balance_transfer,omitempty"`
	// This object stores details about the originating banking transaction that resulted in the ReceivedDebit. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedDebitBankTransfer `json:"bank_transfer,omitempty"`
	// This object stores details about the issuing transactions that resulted in the ReceivedDebit. Present if `type` field value is `card_spend`.
	CardSpend *V2MoneyManagementReceivedDebitCardSpend `json:"card_spend,omitempty"`
	// The time at which the ReceivedDebit was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	Created time.Time `json:"created"`
	// Freeform string sent by the originator of the ReceivedDebit.
	Description string `json:"description,omitempty"`
	// The dispute details.
	DisputeDetails *V2MoneyManagementReceivedDebitDisputeDetails `json:"dispute_details,omitempty"`
	// The amount and currency of the original/external debit request.
	ExternalAmount Amount `json:"external_amount,omitempty"`
	// Financial Account on which funds for ReceivedDebit were debited.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the ReceivedDebit.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A link to the Stripe-hosted receipt for this ReceivedDebit.
	ReceiptURL string `json:"receipt_url,omitempty"`
	// The time at which the scheduled ReceivedDebit is expected to settle.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: `2022-09-18T13:22:18.123Z`.
	// Only present when status is `scheduled`.
	SettlesAt time.Time `json:"settles_at,omitempty"`
	// Open Enum. The status of the ReceivedDebit.
	Status V2MoneyManagementReceivedDebitStatus `json:"status"`
	// Detailed information about the status of the ReceivedDebit.
	StatusDetails *V2MoneyManagementReceivedDebitStatusDetails `json:"status_details,omitempty"`
	// The time at which the ReceivedDebit transitioned to a particular status.
	StatusTransitions *V2MoneyManagementReceivedDebitStatusTransitions `json:"status_transitions,omitempty"`
	// This object stores details about the Stripe Balance Payment that resulted in the ReceivedDebit.
	StripeBalancePayment *V2MoneyManagementReceivedDebitStripeBalancePayment `json:"stripe_balance_payment,omitempty"`
	// Open Enum. The type of the ReceivedDebit.
	Type V2MoneyManagementReceivedDebitType `json:"type"`
}
