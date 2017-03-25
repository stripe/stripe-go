// Package recipienttransfer provides the /recipient_transfers APIs
package recipienttransfer

import (
	stripe "github.com/stripe/stripe-go"
)

const (
	Paid    stripe.RecipientTransferStatus = "paid"
	Pending stripe.RecipientTransferStatus = "pending"
	Transit stripe.RecipientTransferStatus = "in_transit"
	Failed  stripe.RecipientTransferStatus = "failed"

	Card stripe.RecipientTransferType = "card"
	Bank stripe.RecipientTransferType = "bank_account"

	SourceAlipay  stripe.RecipientTransferSourceType = "alipay_account"
	SourceBank    stripe.RecipientTransferSourceType = "bank_account"
	SourceBitcoin stripe.RecipientTransferSourceType = "bitcoin_receiver"
	SourceCard    stripe.RecipientTransferSourceType = "card"

	InsufficientFunds    stripe.RecipientTransferFailCode = "insufficient_funds"
	AccountClosed        stripe.RecipientTransferFailCode = "account_closed"
	NoAccount            stripe.RecipientTransferFailCode = "no_account"
	InvalidAccountNumber stripe.RecipientTransferFailCode = "invalid_account_number"
	DebitNotAuth         stripe.RecipientTransferFailCode = "debit_not_authorized"
	BankOwnerChanged     stripe.RecipientTransferFailCode = "bank_ownership_changed"
	AccountFrozen        stripe.RecipientTransferFailCode = "account_frozen"
	CouldNotProcess      stripe.RecipientTransferFailCode = "could_not_process"
	BankAccountRestrict  stripe.RecipientTransferFailCode = "bank_account_restricted"
	InvalidCurrency      stripe.RecipientTransferFailCode = "invalid_currency"
)

// This object can only be returned/expanded on Balance Transactions
// The API doesn't support retrieve/update/list for those.
