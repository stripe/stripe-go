package stripe

import "encoding/json"

// ReviewReasonType describes the reason why the review is open or closed.
type ReviewReasonType string

// List of values that ReviewReasonType can take.
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
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type review Review
	var v review
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = Review(v)
	return nil
}
