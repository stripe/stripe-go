//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The type of Stripe Money Movement that originated the ReceivedCredit.
type V2MoneyManagementReceivedCreditBalanceTransferType string

// List of values that V2MoneyManagementReceivedCreditBalanceTransferType can take
const (
	V2MoneyManagementReceivedCreditBalanceTransferTypeOutboundPayment  V2MoneyManagementReceivedCreditBalanceTransferType = "outbound_payment"
	V2MoneyManagementReceivedCreditBalanceTransferTypeOutboundTransfer V2MoneyManagementReceivedCreditBalanceTransferType = "outbound_transfer"
	V2MoneyManagementReceivedCreditBalanceTransferTypePayout           V2MoneyManagementReceivedCreditBalanceTransferType = "payout"
	V2MoneyManagementReceivedCreditBalanceTransferTypeTransfer         V2MoneyManagementReceivedCreditBalanceTransferType = "transfer"
	V2MoneyManagementReceivedCreditBalanceTransferTypePayoutV1         V2MoneyManagementReceivedCreditBalanceTransferType = "payout_v1"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferCaBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferCaBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferCaBankAccountNetworkACSS V2MoneyManagementReceivedCreditBankTransferCaBankAccountNetwork = "acss"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferEUBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferEUBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferEUBankAccountNetworkSEPA V2MoneyManagementReceivedCreditBankTransferEUBankAccountNetwork = "sepa"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetworkChaps V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork = "chaps"
	V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetworkFPS   V2MoneyManagementReceivedCreditBankTransferGBBankAccountNetwork = "fps"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferMXBankAccountNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferMXBankAccountNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferMXBankAccountNetworkSpei V2MoneyManagementReceivedCreditBankTransferMXBankAccountNetwork = "spei"
)

// Open Enum. Standard Entry Class code of the ACH entry.
type V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode string

// List of values that V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode can take
const (
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeCcd V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "ccd"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeCie V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "cie"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeCtx V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "ctx"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeIat V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "iat"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodePos V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "pos"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodePpd V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "ppd"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeTel V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "tel"
	V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCodeWeb V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode = "web"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetworkACH            V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork = "ach"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetworkRTP            V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork = "rtp"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetworkUSDomesticWire V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork = "us_domestic_wire"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabeNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabeNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabeNetworkSpei V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabeNetwork = "spei"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpaNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpaNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpaNetworkACSS V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpaNetwork = "acss"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBANNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBANNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBANNetworkSEPACreditTransfer V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBANNetwork = "sepa_credit_transfer"
)

// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetwork string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetwork can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetworkChaps V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetwork = "chaps"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetworkFPS   V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetwork = "fps"
)

// Open Enum. The type of bank transfer that originated this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType string

// List of values that V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType can take
const (
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountTypeABA      V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType = "aba"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountTypeClabe    V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType = "clabe"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountTypeCpa      V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType = "cpa"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountTypeIBAN     V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType = "iban"
	V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountTypeSortCode V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType = "sort_code"
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

// The network the crypto was received from.
type V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork string

// List of values that V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork can take
const (
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkArbitrum        V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "arbitrum"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkAvalancheCChain V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "avalanche_c_chain"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkBase            V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "base"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkEthereum        V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "ethereum"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkOptimism        V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "optimism"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkPolygon         V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "polygon"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkSolana          V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "solana"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkStellar         V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "stellar"
	V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetworkTempo           V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork = "tempo"
)

// Open Enum. The type of crypto wallet transfer that originated this ReceivedCredit.
type V2MoneyManagementReceivedCreditCryptoWalletTransferType string

// List of values that V2MoneyManagementReceivedCreditCryptoWalletTransferType can take
const (
	V2MoneyManagementReceivedCreditCryptoWalletTransferTypeCryptoWallet V2MoneyManagementReceivedCreditCryptoWalletTransferType = "crypto_wallet"
)

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

// Open Enum. The type of the sender.
type V2MoneyManagementReceivedCreditStripeNetworkTransferFromType string

// List of values that V2MoneyManagementReceivedCreditStripeNetworkTransferFromType can take
const (
	V2MoneyManagementReceivedCreditStripeNetworkTransferFromTypeNetworkBusinessProfile V2MoneyManagementReceivedCreditStripeNetworkTransferFromType = "network_business_profile"
)

// Open Enum. The type of flow that caused the ReceivedCredit.
type V2MoneyManagementReceivedCreditType string

// List of values that V2MoneyManagementReceivedCreditType can take
const (
	V2MoneyManagementReceivedCreditTypeBalanceTransfer       V2MoneyManagementReceivedCreditType = "balance_transfer"
	V2MoneyManagementReceivedCreditTypeBankTransfer          V2MoneyManagementReceivedCreditType = "bank_transfer"
	V2MoneyManagementReceivedCreditTypeCardSpend             V2MoneyManagementReceivedCreditType = "card_spend"
	V2MoneyManagementReceivedCreditTypeCryptoWalletTransfer  V2MoneyManagementReceivedCreditType = "crypto_wallet_transfer"
	V2MoneyManagementReceivedCreditTypeExternalCredit        V2MoneyManagementReceivedCreditType = "external_credit"
	V2MoneyManagementReceivedCreditTypeStripeBalancePayment  V2MoneyManagementReceivedCreditType = "stripe_balance_payment"
	V2MoneyManagementReceivedCreditTypeStripeNetworkTransfer V2MoneyManagementReceivedCreditType = "stripe_network_transfer"
)

// This object stores details about the originating Stripe transaction that resulted in the ReceivedCredit. Present if `type` field value is `balance_transfer`.
type V2MoneyManagementReceivedCreditBalanceTransfer struct {
	// The ID of the account that owns the source object originated the ReceivedCredit.
	FromAccount string `json:"from_account,omitempty"`
	// The ID of the outbound payment object that originated the ReceivedCredit.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// The ID of the outbound transfer object that originated the ReceivedCredit.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// The ID of the payout object that originated the ReceivedCredit.
	Payout string `json:"payout,omitempty"`
	// The ID of the v1 transfer object that originated the ReceivedCredit.
	Transfer string `json:"transfer,omitempty"`
	// Open Enum. The type of Stripe Money Movement that originated the ReceivedCredit.
	Type V2MoneyManagementReceivedCreditBalanceTransferType `json:"type"`
}

// Deprecated. Use `originating_bank_account.cpa` instead.
type V2MoneyManagementReceivedCreditBankTransferCaBankAccount struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	// Depending on the bank, this may instead be the last 4 digits of the return account number.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferCaBankAccountNetwork `json:"network"`
}

// Deprecated. Use `originating_bank_account.iban` instead.
type V2MoneyManagementReceivedCreditBankTransferEUBankAccount struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The bic of the account that originated the transfer.
	BIC string `json:"bic,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferEUBankAccountNetwork `json:"network"`
}

// Deprecated. Use `originating_bank_account.sort_code` instead.
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

// Deprecated. Use `originating_bank_account.clabe` instead.
type V2MoneyManagementReceivedCreditBankTransferMXBankAccount struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferMXBankAccountNetwork `json:"network"`
}

// NACHA details for the ACH entry that created this ReceivedCredit.
type V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACH struct {
	// Payment-related information from the ACH addenda record, up to 80 characters.
	Addenda string `json:"addenda,omitempty"`
	// Company Entry Description from the ACH batch header, e.g. "HCCLAIMPMT".
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,omitempty"`
	// Company Identification from the ACH batch header.
	OriginatorCompanyID string `json:"originator_company_id,omitempty"`
	// Company Name from the ACH batch header -- the business that sent the funds.
	OriginatorCompanyName string `json:"originator_company_name,omitempty"`
	// Identification Number from the ACH entry detail record.
	ReceiverIDNumber string `json:"receiver_id_number,omitempty"`
	// Individual Name from the ACH entry detail record.
	ReceiverName string `json:"receiver_name,omitempty"`
	// Open Enum. Standard Entry Class code of the ACH entry.
	StandardEntryClassCode V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACHStandardEntryClassCode `json:"standard_entry_class_code,omitempty"`
	// Trace Number from the ACH entry detail record.
	TraceID string `json:"trace_id,omitempty"`
}

// Network-level detail for the transfer that created this ReceivedCredit. Present only for ACH.
type V2MoneyManagementReceivedCreditBankTransferNetworkDetails struct {
	// NACHA details for the ACH entry that created this ReceivedCredit.
	ACH *V2MoneyManagementReceivedCreditBankTransferNetworkDetailsACH `json:"ach"`
}

// Hash containing the transaction bank details. Present if `type` field value is `aba`.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABA struct {
	// The name of the account holder that sent the payment.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABANetwork `json:"network"`
	// The routing number of the account that originated the transfer.
	RoutingNumber string `json:"routing_number,omitempty"`
}

// Hash containing the transaction bank details. Present if `type` field value is `clabe`.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabe struct {
	// The name of the account holder that sent the payment.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The BIC/SWIFT code of the account that originated the transfer.
	BIC string `json:"bic,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabeNetwork `json:"network"`
}

// Hash containing the transaction bank details. Present if `type` field value is `cpa`.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpa struct {
	// The name of the account holder that sent the payment.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The BIC/SWIFT code of the account that originated the transfer.
	BIC string `json:"bic,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpaNetwork `json:"network"`
}

// Hash containing the transaction bank details. Present if `type` field value is `iban`.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBAN struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The BIC/SWIFT code of the account that originated the transfer.
	BIC string `json:"bic,omitempty"`
	// The origination country of the bank transfer.
	Country string `json:"country,omitempty"`
	// The IBAN that originated the transfer.
	IBAN string `json:"iban,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBANNetwork `json:"network"`
}

// Hash containing the transaction bank details. Present if `type` field value is `sort_code`.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCode struct {
	// The account holder name of the bank account the transfer was received from.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The bank name the transfer was received from.
	BankName string `json:"bank_name,omitempty"`
	// The last 4 digits of the account number that originated the transfer.
	Last4 string `json:"last4,omitempty"`
	// Open Enum. The money transmission network used to send funds for this ReceivedCredit.
	Network V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCodeNetwork `json:"network"`
	// The sort code of the account that originated the transfer.
	SortCode string `json:"sort_code,omitempty"`
}

// Hash containing the originating bank account details and type for this bank transfer.
type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccount struct {
	// Hash containing the transaction bank details. Present if `type` field value is `aba`.
	ABA *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountABA `json:"aba,omitempty"`
	// Hash containing the transaction bank details. Present if `type` field value is `clabe`.
	Clabe *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabe `json:"clabe,omitempty"`
	// Hash containing the transaction bank details. Present if `type` field value is `cpa`.
	Cpa *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpa `json:"cpa,omitempty"`
	// Hash containing the transaction bank details. Present if `type` field value is `iban`.
	IBAN *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIBAN `json:"iban,omitempty"`
	// Hash containing the transaction bank details. Present if `type` field value is `sort_code`.
	SortCode *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCode `json:"sort_code,omitempty"`
	// Open Enum. The type of bank transfer that originated this ReceivedCredit.
	Type V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountType `json:"type"`
}

// Deprecated. Use `originating_bank_account.iban` instead.
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

// Deprecated. Use `originating_bank_account.aba` instead.
type V2MoneyManagementReceivedCreditBankTransferUSBankAccount struct {
	// The name of the account holder that sent the payment.
	AccountHolderName string `json:"account_holder_name,omitempty"`
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
	// Deprecated. Use `originating_bank_account.cpa` instead.
	CaBankAccount *V2MoneyManagementReceivedCreditBankTransferCaBankAccount `json:"ca_bank_account,omitempty"`
	// Deprecated. Use `originating_bank_account.iban` instead.
	EUBankAccount *V2MoneyManagementReceivedCreditBankTransferEUBankAccount `json:"eu_bank_account,omitempty"`
	// Financial Address on which funds for ReceivedCredit were received.
	FinancialAddress string `json:"financial_address"`
	// Deprecated. Use `originating_bank_account.sort_code` instead.
	GBBankAccount *V2MoneyManagementReceivedCreditBankTransferGBBankAccount `json:"gb_bank_account,omitempty"`
	// Deprecated. Use `originating_bank_account.clabe` instead.
	MXBankAccount *V2MoneyManagementReceivedCreditBankTransferMXBankAccount `json:"mx_bank_account,omitempty"`
	// Network-level detail for the transfer that created this ReceivedCredit. Present only for ACH.
	NetworkDetails *V2MoneyManagementReceivedCreditBankTransferNetworkDetails `json:"network_details,omitempty"`
	// Hash containing the originating bank account details and type for this bank transfer.
	OriginatingBankAccount *V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccount `json:"originating_bank_account"`
	// Deprecated. Use `originating_bank_account.iban` instead.
	SEPABankAccount *V2MoneyManagementReceivedCreditBankTransferSEPABankAccount `json:"sepa_bank_account,omitempty"`
	// Freeform string set by originator of the external ReceivedCredit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Deprecated. Use `originating_bank_account.aba` instead.
	USBankAccount *V2MoneyManagementReceivedCreditBankTransferUSBankAccount `json:"us_bank_account,omitempty"`
}

// Hash containing information about the Dispute that triggered this credit.
type V2MoneyManagementReceivedCreditCardSpendDispute struct {
	// The reference to the v1 issuing dispute ID.
	IssuingDisputeV1 string `json:"issuing_dispute_v1"`
}

// Hash containing information about the Refund that triggered this credit.
type V2MoneyManagementReceivedCreditCardSpendRefund struct {
	// The reference to the v1 issuing transaction ID.
	IssuingTransactionV1 string `json:"issuing_transaction_v1"`
}

// This object stores details about the originating issuing card spend that resulted in the ReceivedCredit. Present if `type` field value is `card_spend`.
type V2MoneyManagementReceivedCreditCardSpend struct {
	// The reference to the issuing card object.
	CardV1ID string `json:"card_v1_id"`
	// Hash containing information about the Dispute that triggered this credit.
	Dispute *V2MoneyManagementReceivedCreditCardSpendDispute `json:"dispute,omitempty"`
	// Hash containing information about the Refund that triggered this credit.
	Refund *V2MoneyManagementReceivedCreditCardSpendRefund `json:"refund,omitempty"`
}

// Hash containing the transaction crypto wallet details.
type V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWallet struct {
	// The address of the wallet the crypto was received from.
	Address string `json:"address"`
	// A memo also for identifying the recipient for memo-based blockchains (e.g., Stellar),.
	Memo string `json:"memo"`
	// The network the crypto was received from.
	Network V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWalletNetwork `json:"network"`
}

// This object stores details about the originating crypto transaction that resulted in the ReceivedCredit. Present if `type` field value is `crypto_wallet_transfer`.
type V2MoneyManagementReceivedCreditCryptoWalletTransfer struct {
	// Hash containing the transaction crypto wallet details.
	CryptoWallet *V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWallet `json:"crypto_wallet"`
	// Financial Address on which funds for ReceivedCredit were received.
	FinancialAddress string `json:"financial_address"`
	// Freeform string set by originator of the external ReceivedCredit.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Open Enum. The type of crypto wallet transfer that originated this ReceivedCredit.
	Type V2MoneyManagementReceivedCreditCryptoWalletTransferType `json:"type"`
}

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

// This object stores details about the stripe balance pay refund that resulted in the ReceivedCredit. Present if `type` field value is `stripe_balance_payment`.
type V2MoneyManagementReceivedCreditStripeBalancePayment struct {
	// ID of the debit agreement associated with this payment.
	DebitAgreement string `json:"debit_agreement,omitempty"`
	// Statement descriptor for the Stripe Balance Payment.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}

// Information about the sender of the network transfer.
type V2MoneyManagementReceivedCreditStripeNetworkTransferFrom struct {
	// The network ID of the sender.
	NetworkBusinessProfile string `json:"network_business_profile"`
	// Open Enum. The type of the sender.
	Type V2MoneyManagementReceivedCreditStripeNetworkTransferFromType `json:"type"`
}

// This object stores details about the Stripe network transfer that resulted in the ReceivedCredit. Present if `type` field value is `stripe_network_transfer`.
type V2MoneyManagementReceivedCreditStripeNetworkTransfer struct {
	// Information about the sender of the network transfer.
	From *V2MoneyManagementReceivedCreditStripeNetworkTransferFrom `json:"from"`
}

// Use ReceivedCredits API to retrieve information on when, where, and how funds are sent into your FinancialAccount.
type V2MoneyManagementReceivedCredit struct {
	APIResource
	// The amount and currency of the ReceivedCredit.
	Amount Amount `json:"amount"`
	// The amount and currency of the ReceivedCredit that was received.
	AmountReceived Amount `json:"amount_received"`
	// This object stores details about the originating Stripe transaction that resulted in the ReceivedCredit. Present if `type` field value is `balance_transfer`.
	BalanceTransfer *V2MoneyManagementReceivedCreditBalanceTransfer `json:"balance_transfer,omitempty"`
	// This object stores details about the originating banking transaction that resulted in the ReceivedCredit. Present if `type` field value is `bank_transfer`.
	BankTransfer *V2MoneyManagementReceivedCreditBankTransfer `json:"bank_transfer,omitempty"`
	// This object stores details about the originating issuing card spend that resulted in the ReceivedCredit. Present if `type` field value is `card_spend`.
	CardSpend *V2MoneyManagementReceivedCreditCardSpend `json:"card_spend,omitempty"`
	// Time at which the ReceivedCredit was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// This object stores details about the originating crypto transaction that resulted in the ReceivedCredit. Present if `type` field value is `crypto_wallet_transfer`.
	CryptoWalletTransfer *V2MoneyManagementReceivedCreditCryptoWalletTransfer `json:"crypto_wallet_transfer,omitempty"`
	// Freeform string set by originator of the ReceivedCredit.
	Description string `json:"description,omitempty"`
	// The amount and currency of the original/external credit request.
	ExternalAmount Amount `json:"external_amount,omitempty"`
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
	// This object stores details about the stripe balance pay refund that resulted in the ReceivedCredit. Present if `type` field value is `stripe_balance_payment`.
	StripeBalancePayment *V2MoneyManagementReceivedCreditStripeBalancePayment `json:"stripe_balance_payment,omitempty"`
	// This object stores details about the Stripe network transfer that resulted in the ReceivedCredit. Present if `type` field value is `stripe_network_transfer`.
	StripeNetworkTransfer *V2MoneyManagementReceivedCreditStripeNetworkTransfer `json:"stripe_network_transfer,omitempty"`
	// Open Enum. The type of flow that caused the ReceivedCredit.
	Type V2MoneyManagementReceivedCreditType `json:"type"`
}
