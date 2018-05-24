package stripe

import "encoding/json"

// ReviewReasonType describes the reason why the review is open or closed.
type ReviewReasonType string

const (
	ReviewReasonApproved        ReviewReasonType = "approved"
	ReviewReasonDisputed        ReviewReasonType = "disputed"
	ReviewReasonManual          ReviewReasonType = "manual"
	ReviewReasonRefunded        ReviewReasonType = "refunded"
	ReviewReasonRefundedAsFraud ReviewReasonType = "refunded_as_fraud"
	ReviewReasonRule            ReviewReasonType = "rule"
)

type Review struct {
	Charge   *Charge          `json:"charge"`
	Created  int64            `json:"created"`
	ID       string           `json:"id"`
	Livemode bool             `json:"livemode"`
	Open     bool             `json:"open"`
	Reason   ReviewReasonType `json:"reason"`
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
