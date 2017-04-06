// Package recipienttransfer provides the /recipient_transfers APIs
package recipienttransfer

import (
	stripe "github.com/stripe/stripe-go"
)

const (
	Failed  stripe.RecipientTransferStatus = "failed"
	Paid    stripe.RecipientTransferStatus = "paid"
	Pending stripe.RecipientTransferStatus = "pending"
	Transit stripe.RecipientTransferStatus = "in_transit"

	Bank stripe.RecipientTransferType = "bank_account"
	Card stripe.RecipientTransferType = "card"

	SourceAlipay  stripe.RecipientTransferSourceType = "alipay_account"
	SourceBank    stripe.RecipientTransferSourceType = "bank_account"
	SourceBitcoin stripe.RecipientTransferSourceType = "bitcoin_receiver"
	SourceCard    stripe.RecipientTransferSourceType = "card"

	AccountClosed        stripe.RecipientTransferFailCode = "account_closed"
	AccountFrozen        stripe.RecipientTransferFailCode = "account_frozen"
	BankAccountRestrict  stripe.RecipientTransferFailCode = "bank_account_restricted"
	BankOwnerChanged     stripe.RecipientTransferFailCode = "bank_ownership_changed"
	DebitNotAuth         stripe.RecipientTransferFailCode = "debit_not_authorized"
	CouldNotProcess      stripe.RecipientTransferFailCode = "could_not_process"
	InsufficientFunds    stripe.RecipientTransferFailCode = "insufficient_funds"
	InvalidAccountNumber stripe.RecipientTransferFailCode = "invalid_account_number"
	InvalidCurrency      stripe.RecipientTransferFailCode = "invalid_currency"
	NoAccount            stripe.RecipientTransferFailCode = "no_account"
)

// This object can only be returned/expanded on Balance Transactions
// The API doesn't support retrieve/update/list for those.
