package stripe

import "encoding/json"

// Reason describes the reason why the review is open or closed.
// Allowed values are "rule", "manual", "approved", "refunded",
// "refunded_as_fraud", "disputed".
type ReasonType string

const (
	ReasonApproved        ReasonType = "approved"
	ReasonDisputed        ReasonType = "disputed"
	ReasonManual          ReasonType = "manual"
	ReasonRefunded        ReasonType = "refunded"
	ReasonRefundedAsFraud ReasonType = "refunded_as_fraud"
	ReasonRule            ReasonType = "rule"
)

type Review struct {
	Charge   *Charge    `json:"charge"`
	Created  int64      `json:"created"`
	ID       string     `json:"id"`
	Livemode bool       `json:"livemode"`
	Open     bool       `json:"open"`
	Reason   ReasonType `json:"reason"`
}

func (r *Review) UnmarshalJSON(data []byte) error {
	type review Review
	var rr review

	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Review(rr)
	} else {
		// Otherwise...we have to strip the escaping
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
