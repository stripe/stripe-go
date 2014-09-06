package stripe

import "encoding/json"

// FeeRefundParams is the set of parameters that can be used when refunding a fee.
// For more details see https://stripe.com/docs/api#fee_refund.
type FeeRefundParams struct {
	Params
	Fee    string
	Amount uint64
	Meta   map[string]string
}

// FeeRefundListParams is the set of parameters that can be used when listing fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
type FeeRefundListParams struct {
	ListParams
	Fee string
}

// FeeRefund is the resource representing a Stripe fee refund.
// For more details see https://stripe.com/docs/api#fee_refunds.
type FeeRefund struct {
	Id       string            `json:"id"`
	Amount   uint64            `json:"amount"`
	Created  int64             `json:"created"`
	Currency Currency          `json:"currency"`
	Tx       *Transaction      `json:"balance_transaction"`
	Fee      string            `json:"fee"`
	Meta     map[string]string `json:"metadata"`
}

// FeeRefundist is a list object for fee refunds.
type FeeRefundList struct {
	ListMeta
	Values []*FeeRefund `json:"data"`
}

// FeeRefundIter is a iterator for list responses.
type FeeRefundIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *FeeRefundIter) Next() (*FeeRefund, error) {
	f, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return f.(*FeeRefund), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *FeeRefundIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *FeeRefundIter) Meta() *ListMeta {
	return i.Iter.Meta()
}
func (f *FeeRefund) UnmarshalJSON(data []byte) error {
	type feerefund FeeRefund
	var ff feerefund
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = FeeRefund(ff)
	} else {
		// the id is surrounded by escaped \, so ignore those
		f.Id = string(data[1 : len(data)-1])
	}

	return nil
}
