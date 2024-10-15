//
//
// File generated from our OpenAPI spec
//
//

package stripe

// CA-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferCaBankTransferParams struct{}

// EU-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferEUBankTransferParams struct {
	BIC        *string `form:"bic"`
	IBANLast4  *string `form:"iban_last4"`
	Network    *string `form:"network"`
	SenderName *string `form:"sender_name"`
}

// GB-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferGBBankTransferParams struct {
	AccountNumberLast4 *string `form:"account_number_last4"`
	SenderName         *string `form:"sender_name"`
	SortCode           *string `form:"sort_code"`
}

// JP-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferJPBankTransferParams struct {
	SenderBank   *string `form:"sender_bank"`
	SenderBranch *string `form:"sender_branch"`
	SenderName   *string `form:"sender_name"`
}

// MX-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferMXBankTransferParams struct{}

// US-specific details of the bank transfer funding.
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferUSBankTransferParams struct {
	Network    *string `form:"network"`
	SenderName *string `form:"sender_name"`
}
type TestHelpersCustomerCashBalanceTransactionFundedBankTransferParams struct {
	// CA-specific details of the bank transfer funding.
	CaBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferCaBankTransferParams `form:"ca_bank_transfer"`
	// EU-specific details of the bank transfer funding.
	EUBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// GB-specific details of the bank transfer funding.
	GBBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferGBBankTransferParams `form:"gb_bank_transfer"`
	// JP-specific details of the bank transfer funding.
	JPBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferJPBankTransferParams `form:"jp_bank_transfer"`
	// MX-specific details of the bank transfer funding.
	MXBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferMXBankTransferParams `form:"mx_bank_transfer"`
	Reference      *string                                                                          `form:"reference"`
	Type           *string                                                                          `form:"type"`
	// US-specific details of the bank transfer funding.
	USBankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferUSBankTransferParams `form:"us_bank_transfer"`
}

// If this is a `type=funded` transaction, contains information about the funding.
type TestHelpersCustomerCashBalanceTransactionFundedParams struct {
	BankTransfer *TestHelpersCustomerCashBalanceTransactionFundedBankTransferParams `form:"bank_transfer"`
}

// If this is a `type=funding_reversed` transaction, contains information about the reversal of a funding.
type TestHelpersCustomerCashBalanceTransactionFundingReversedParams struct {
	// The ID of the `funded` cash balance transaction to be reversed.
	ReversedCustomerCashBalanceTransaction *string `form:"reversed_customer_cash_balance_transaction"`
}

// Simulate various customer cash balance side-effects by creating synthetic cash balance transactions in testmode.
type TestHelpersCustomerCashBalanceTransactionParams struct {
	Params `form:"*"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If this is a `type=funded` transaction, contains information about the funding.
	Funded *TestHelpersCustomerCashBalanceTransactionFundedParams `form:"funded"`
	// If this is a `type=funding_reversed` transaction, contains information about the reversal of a funding.
	FundingReversed *TestHelpersCustomerCashBalanceTransactionFundingReversedParams `form:"funding_reversed"`
	// The amount associated with the cash balance transaction. Only applicable to transactions of type `funded`.
	NetAmount *int64 `form:"net_amount"`
	// The type of cash balance transaction to generate.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCustomerCashBalanceTransactionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
