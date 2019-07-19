package stripe

import "encoding/json"

// CreditNoteReason is the reason why a given credit note was created.
type CreditNoteReason string

// List of values that CreditNoteReason can take.
const (
	CreditNoteReasonDuplicate             CreditNoteReason = "duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
)

// CreditNoteStatus is the list of allowed values for the credit note's status.
type CreditNoteStatus string

// List of values that CreditNoteStatus can take.
const (
	CreditNoteStatusIssued CreditNoteStatus = "issued"
	CreditNoteStatusVoid   CreditNoteStatus = "void"
)

// CreditNoteType is the list of allowed values for the credit note's type.
type CreditNoteType string

// List of values that CreditNoteType can take.
const (
	CreditNoteTypePostPayment CreditNoteType = "post_payment"
	CreditNoteTypePrePayment  CreditNoteType = "pre_payment"
)

// CreditNoteParams is the set of parameters that can be used when creating or updating a credit note.
// For more details see https://stripe.com/docs/api/credit_notes/create, https://stripe.com/docs/api/credit_notes/update.
type CreditNoteParams struct {
	Params       `form:"*"`
	Amount       *int64  `form:"amount"`
	CreditAmount *int64  `form:"credit_amount"`
	Invoice      *string `form:"invoice"`
	Memo         *string `form:"memo"`
	Reason       *string `form:"reason"`
	Refund       *string `form:"refund"`
	RefundAmount *int64  `form:"refund_amount"`
}

// CreditNoteListParams is the set of parameters that can be used when listing credit notes.
// For more details see https://stripe.com/docs/api/credit_notes/list.
type CreditNoteListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Invoice    *string `form:"invoice"`
}

// CreditNoteVoidParams is the set of parameters that can be used when voiding invoices.
type CreditNoteVoidParams struct {
	Params `form:"*"`
}

// CreditNote is the resource representing a Stripe credit note.
// For more details see https://stripe.com/docs/api/credit_notes/object.
type CreditNote struct {
	Amount                     int64                       `json:"amount"`
	Created                    int64                       `json:"created"`
	Currency                   Currency                    `json:"currency"`
	Customer                   *Customer                   `json:"customer"`
	CustomerBalanceTransaction *CustomerBalanceTransaction `json:"customer_balance_transaction"`
	Invoice                    *Invoice                    `json:"invoice"`
	ID                         string                      `json:"id"`
	Livemode                   bool                        `json:"livemode"`
	Memo                       string                      `json:"memo"`
	Metadata                   map[string]string           `json:"metadata"`
	Number                     string                      `json:"number"`
	Object                     string                      `json:"object"`
	PDF                        string                      `json:"pdf"`
	Reason                     CreditNoteReason            `json:"reason"`
	Refund                     *Refund                     `json:"refund"`
	Status                     CreditNoteStatus            `json:"status"`
	Type                       CreditNoteType              `json:"type"`
	VoidedAt                   int64                       `json:"voided_at"`
}

// CreditNoteList is a list of credit notes as retrieved from a list endpoint.
type CreditNoteList struct {
	ListMeta
	Data []*CreditNote `json:"data"`
}

// UnmarshalJSON handles deserialization of a CreditNote.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *CreditNote) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type note CreditNote
	var v note
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = CreditNote(v)
	return nil
}
