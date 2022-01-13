//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// [Learn more](https://stripe.com/docs/reports/reporting-categories) about how reporting categories can help you understand balance transactions from an accounting perspective.
type BalanceTransactionReportingCategory string

// List of values that BalanceTransactionReportingCategory can take
const (
	BalanceTransactionReportingCategoryAdvance                     BalanceTransactionReportingCategory = "advance"
	BalanceTransactionReportingCategoryAdvanceFunding              BalanceTransactionReportingCategory = "advance_funding"
	BalanceTransactionReportingCategoryCharge                      BalanceTransactionReportingCategory = "charge"
	BalanceTransactionReportingCategoryChargeFailure               BalanceTransactionReportingCategory = "charge_failure"
	BalanceTransactionReportingCategoryConnectCollectionTransfer   BalanceTransactionReportingCategory = "connect_collection_transfer"
	BalanceTransactionReportingCategoryConnectReservedFunds        BalanceTransactionReportingCategory = "connect_reserved_funds"
	BalanceTransactionReportingCategoryDispute                     BalanceTransactionReportingCategory = "dispute"
	BalanceTransactionReportingCategoryDisputeReversal             BalanceTransactionReportingCategory = "dispute_reversal"
	BalanceTransactionReportingCategoryFee                         BalanceTransactionReportingCategory = "fee"
	BalanceTransactionReportingCategoryIssuingAuthorizationHold    BalanceTransactionReportingCategory = "issuing_authorization_hold"
	BalanceTransactionReportingCategoryIssuingAuthorizationRelease BalanceTransactionReportingCategory = "issuing_authorization_release"
	BalanceTransactionReportingCategoryIssuingTransaction          BalanceTransactionReportingCategory = "issuing_transaction"
	BalanceTransactionReportingCategoryOtherAdjustment             BalanceTransactionReportingCategory = "other_adjustment"
	BalanceTransactionReportingCategoryPartialCaptureReversal      BalanceTransactionReportingCategory = "partial_capture_reversal"
	BalanceTransactionReportingCategoryPayout                      BalanceTransactionReportingCategory = "payout"
	BalanceTransactionReportingCategoryPayoutReversal              BalanceTransactionReportingCategory = "payout_reversal"
	BalanceTransactionReportingCategoryPlatformEarning             BalanceTransactionReportingCategory = "platform_earning"
	BalanceTransactionReportingCategoryPlatformEarningRefund       BalanceTransactionReportingCategory = "platform_earning_refund"
	BalanceTransactionReportingCategoryRefund                      BalanceTransactionReportingCategory = "refund"
	BalanceTransactionReportingCategoryRefundFailure               BalanceTransactionReportingCategory = "refund_failure"
	BalanceTransactionReportingCategoryRiskReservedFunds           BalanceTransactionReportingCategory = "risk_reserved_funds"
	BalanceTransactionReportingCategoryTax                         BalanceTransactionReportingCategory = "tax"
	BalanceTransactionReportingCategoryTopup                       BalanceTransactionReportingCategory = "topup"
	BalanceTransactionReportingCategoryTopupReversal               BalanceTransactionReportingCategory = "topup_reversal"
	BalanceTransactionReportingCategoryTransfer                    BalanceTransactionReportingCategory = "transfer"
	BalanceTransactionReportingCategoryTransferReversal            BalanceTransactionReportingCategory = "transfer_reversal"
)

type BalanceTransactionSourceType string

// List of values that BalanceTransactionSourceType can take
const (
	BalanceTransactionSourceTypeApplicationFee       BalanceTransactionSourceType = "application_fee"
	BalanceTransactionSourceTypeCharge               BalanceTransactionSourceType = "charge"
	BalanceTransactionSourceTypeDispute              BalanceTransactionSourceType = "dispute"
	BalanceTransactionSourceTypeFeeRefund            BalanceTransactionSourceType = "fee_refund"
	BalanceTransactionSourceTypeIssuingAuthorization BalanceTransactionSourceType = "issuing.authorization"
	BalanceTransactionSourceTypeIssuingDispute       BalanceTransactionSourceType = "issuing.dispute"
	BalanceTransactionSourceTypeIssuingTransaction   BalanceTransactionSourceType = "issuing.transaction"
	BalanceTransactionSourceTypePayout               BalanceTransactionSourceType = "payout"
	BalanceTransactionSourceTypeRefund               BalanceTransactionSourceType = "refund"
	BalanceTransactionSourceTypeReversal             BalanceTransactionSourceType = "reversal"
	BalanceTransactionSourceTypeTopup                BalanceTransactionSourceType = "topup"
	BalanceTransactionSourceTypeTransfer             BalanceTransactionSourceType = "transfer"
)

// List of values that BalanceTransactionStatus can take
const (
	BalanceTransactionStatusAvailable BalanceTransactionStatus = "available"
	BalanceTransactionStatusPending   BalanceTransactionStatus = "pending"
)

// Transaction type: `adjustment`, `advance`, `advance_funding`, `anticipation_repayment`, `application_fee`, `application_fee_refund`, `charge`, `connect_collection_transfer`, `contribution`, `issuing_authorization_hold`, `issuing_authorization_release`, `issuing_dispute`, `issuing_transaction`, `payment`, `payment_failure_refund`, `payment_refund`, `payout`, `payout_cancel`, `payout_failure`, `refund`, `refund_failure`, `reserve_transaction`, `reserved_funds`, `stripe_fee`, `stripe_fx_fee`, `tax_fee`, `topup`, `topup_reversal`, `transfer`, `transfer_cancel`, `transfer_failure`, or `transfer_refund`. [Learn more](https://stripe.com/docs/reports/balance-transaction-types) about balance transaction types and what they represent. If you are looking to classify transactions for accounting purposes, you might want to consider `reporting_category` instead.
type BalanceTransactionType string

// List of values that BalanceTransactionType can take
const (
	BalanceTransactionTypeAdjustment                      BalanceTransactionType = "adjustment"
	BalanceTransactionTypeAdvance                         BalanceTransactionType = "advance"
	BalanceTransactionTypeAdvanceFunding                  BalanceTransactionType = "advance_funding"
	BalanceTransactionTypeAnticipationRepayment           BalanceTransactionType = "anticipation_repayment"
	BalanceTransactionTypeApplicationFee                  BalanceTransactionType = "application_fee"
	BalanceTransactionTypeApplicationFeeRefund            BalanceTransactionType = "application_fee_refund"
	BalanceTransactionTypeCharge                          BalanceTransactionType = "charge"
	BalanceTransactionTypeConnectCollectionTransfer       BalanceTransactionType = "connect_collection_transfer"
	BalanceTransactionTypeContribution                    BalanceTransactionType = "contribution"
	BalanceTransactionTypeIssuingAuthorizationHold        BalanceTransactionType = "issuing_authorization_hold"
	BalanceTransactionTypeIssuingAuthorizationRelease     BalanceTransactionType = "issuing_authorization_release"
	BalanceTransactionTypeIssuingAuthorizationDispute     BalanceTransactionType = "issuing_dispute"
	BalanceTransactionTypeIssuingAuthorizationTransaction BalanceTransactionType = "issuing_transaction"
	BalanceTransactionTypePayment                         BalanceTransactionType = "payment"
	BalanceTransactionTypePaymentFailureRefund            BalanceTransactionType = "payment_failure_refund"
	BalanceTransactionTypePaymentRefund                   BalanceTransactionType = "payment_refund"
	BalanceTransactionTypePayout                          BalanceTransactionType = "payout"
	BalanceTransactionTypePayoutCancel                    BalanceTransactionType = "payout_cancel"
	BalanceTransactionTypePayoutFailure                   BalanceTransactionType = "payout_failure"
	BalanceTransactionTypeRefund                          BalanceTransactionType = "refund"
	BalanceTransactionTypeRefundFailure                   BalanceTransactionType = "refund_failure"
	BalanceTransactionTypeReserveTransaction              BalanceTransactionType = "reserve_transaction"
	BalanceTransactionTypeReservedFunds                   BalanceTransactionType = "reserved_funds"
	BalanceTransactionTypeStripeFee                       BalanceTransactionType = "stripe_fee"
	BalanceTransactionTypeStripeFxFee                     BalanceTransactionType = "stripe_fx_fee"
	BalanceTransactionTypeTaxFee                          BalanceTransactionType = "tax_fee"
	BalanceTransactionTypeTopup                           BalanceTransactionType = "topup"
	BalanceTransactionTypeTopupReversal                   BalanceTransactionType = "topup_reversal"
	BalanceTransactionTypeTransfer                        BalanceTransactionType = "transfer"
	BalanceTransactionTypeTransferCancel                  BalanceTransactionType = "transfer_cancel"
	BalanceTransactionTypeTransferFailure                 BalanceTransactionType = "transfer_failure"
	BalanceTransactionTypeTransferRefund                  BalanceTransactionType = "transfer_refund"
)

// Returns a list of transactions that have contributed to the Stripe account balance (e.g., charges, transfers, and so forth). The transactions are returned in sorted order, with the most recent transactions appearing first.
//
// Note that this endpoint was previously called “Balance history” and used the path /v1/balance/history.
type BalanceTransactionListParams struct {
	ListParams       `form:"*"`
	AvailableOn      *int64            `form:"available_on"`
	AvailableOnRange *RangeQueryParams `form:"available_on"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Currency         *string           `form:"currency"`
	Payout           *string           `form:"payout"`
	Source           *string           `form:"source"`
	Type             *string           `form:"type"`
}

// Retrieves the balance transaction with the given ID.
//
// Note that this endpoint previously used the path /v1/balance/history/:id.
type BalanceTransactionParams struct {
	Params `form:"*"`
}

// Fees (in %s) paid for this transaction.
type BalanceTransactionFee struct {
	Amount      int64    `json:"amount"`
	Application string   `json:"application"`
	Currency    Currency `json:"currency"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
}

// Balance transactions represent funds moving through your Stripe account.
// They're created for every type of transaction that comes into or flows out of your Stripe account balance.
//
// Related guide: [Balance Transaction Types](https://stripe.com/docs/reports/balance-transaction-types).
type BalanceTransaction struct {
	APIResource
	Amount            int64                               `json:"amount"`
	AvailableOn       int64                               `json:"available_on"`
	Created           int64                               `json:"created"`
	Currency          Currency                            `json:"currency"`
	Description       string                              `json:"description"`
	ExchangeRate      float64                             `json:"exchange_rate"`
	Fee               int64                               `json:"fee"`
	FeeDetails        []*BalanceTransactionFee            `json:"fee_details"`
	ID                string                              `json:"id"`
	Net               int64                               `json:"net"`
	Object            string                              `json:"object"`
	ReportingCategory BalanceTransactionReportingCategory `json:"reporting_category"`
	Source            *BalanceTransactionSource           `json:"source"`
	Status            BalanceTransactionStatus            `json:"status"`
	Type              BalanceTransactionType              `json:"type"`
}
type BalanceTransactionSource struct {
	ID   string                       `json:"id"`
	Type BalanceTransactionSourceType `json:"object"`

	ApplicationFee       *ApplicationFee       `json:"-"`
	Charge               *Charge               `json:"-"`
	Dispute              *Dispute              `json:"-"`
	FeeRefund            *FeeRefund            `json:"-"`
	IssuingAuthorization *IssuingAuthorization `json:"-"`
	IssuingDispute       *IssuingDispute       `json:"-"`
	IssuingTransaction   *IssuingAuthorization `json:"-"`
	Payout               *Payout               `json:"-"`
	Refund               *Refund               `json:"-"`
	Reversal             *Reversal             `json:"-"`
	Topup                *Topup                `json:"-"`
	Transfer             *Transfer             `json:"-"`
}

// BalanceTransactionList is a list of BalanceTransactions as retrieved from a list endpoint.
type BalanceTransactionList struct {
	APIResource
	ListMeta
	Data []*BalanceTransaction `json:"data"`
}

// UnmarshalJSON handles deserialization of a BalanceTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BalanceTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type balanceTransaction BalanceTransaction
	var v balanceTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BalanceTransaction(v)
	return nil
}

// UnmarshalJSON handles deserialization of a BalanceTransactionSource.
// This custom unmarshaling is needed because the specific type of
// BalanceTransactionSource it refers to is specified in the JSON
func (b *BalanceTransactionSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type balanceTransactionSource BalanceTransactionSource
	var v balanceTransactionSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BalanceTransactionSource(v)
	var err error

	switch b.Type {
	case BalanceTransactionSourceTypeApplicationFee:
		err = json.Unmarshal(data, &b.ApplicationFee)
	case BalanceTransactionSourceTypeCharge:
		err = json.Unmarshal(data, &b.Charge)
	case BalanceTransactionSourceTypeDispute:
		err = json.Unmarshal(data, &b.Dispute)
	case BalanceTransactionSourceTypeFeeRefund:
		err = json.Unmarshal(data, &b.FeeRefund)
	case BalanceTransactionSourceTypeIssuingAuthorization:
		err = json.Unmarshal(data, &b.IssuingAuthorization)
	case BalanceTransactionSourceTypeIssuingDispute:
		err = json.Unmarshal(data, &b.IssuingDispute)
	case BalanceTransactionSourceTypeIssuingTransaction:
		err = json.Unmarshal(data, &b.IssuingTransaction)
	case BalanceTransactionSourceTypePayout:
		err = json.Unmarshal(data, &b.Payout)
	case BalanceTransactionSourceTypeRefund:
		err = json.Unmarshal(data, &b.Refund)
	case BalanceTransactionSourceTypeReversal:
		err = json.Unmarshal(data, &b.Reversal)
	case BalanceTransactionSourceTypeTopup:
		err = json.Unmarshal(data, &b.Topup)
	case BalanceTransactionSourceTypeTransfer:
		err = json.Unmarshal(data, &b.Transfer)
	}
	return err
}
