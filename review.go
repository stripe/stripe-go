package stripe

import "encoding/json"

// ReviewClosedReason describes the reason why the review is closed.
type ReviewClosedReason string

// List of values that ReviewClosedReason can take.
const (
	ReviewClosedReasonApproved        ReviewClosedReason = "approved"
	ReviewClosedReasonDisputed        ReviewClosedReason = "disputed"
	ReviewClosedReasonRefunded        ReviewClosedReason = "refunded"
	ReviewClosedReasonRefundedAsFraud ReviewClosedReason = "refunded_as_fraud"
)

// ReviewOpenedReason describes the reason why the review is opened.
type ReviewOpenedReason string

// List of values that ReviewOpenedReason can take.
const (
	ReviewOpenedReasonManual ReviewOpenedReason = "manual"
	ReviewOpenedReasonRule   ReviewOpenedReason = "rule"
)

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

// ReviewParams is the set of parameters that can be used when approving a review.
type ReviewParams struct {
	Params `form:"*"`
}

// ReviewApproveParams is the set of parameters that can be used when approving a review.
type ReviewApproveParams struct {
	Params `form:"*"`
}

// ReviewListParams is the set of parameters that can be used when listing reviews.
type ReviewListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// ReviewIPAddressLocation represents information about the IP associated with a review.
type ReviewIPAddressLocation struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Region    string  `json:"region"`
}

// ReviewSession represents information about the browser session associated with a review.
type ReviewSession struct {
	Browser  string `json:"browser"`
	Device   string `json:"device"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

// Review is the resource representing a Radar review.
// For more details see https://stripe.com/docs/api#reviews.
type Review struct {
	BillingZip        string                   `json:"billing_zip"`
	Charge            *Charge                  `json:"charge"`
	ClosedReason      ReviewClosedReason       `json:"closed_reason"`
	Created           int64                    `json:"created"`
	ID                string                   `json:"id"`
	IPAddress         string                   `json:"ip_address"`
	IPAddressLocation *ReviewIPAddressLocation `json:"ip_address_location"`
	Livemode          bool                     `json:"livemode"`
	Object            string                   `json:"object"`
	Open              bool                     `json:"open"`
	OpenedReason      ReviewOpenedReason       `json:"opened_reason"`
	PaymentIntent     *PaymentIntent           `json:"payment_intent"`
	Reason            ReviewReasonType         `json:"reason"`
	Session           *ReviewSession           `json:"session"`
}

// ReviewList is a list of reviews as retrieved from a list endpoint.
type ReviewList struct {
	ListMeta
	Data []*Review `json:"data"`
}

// UnmarshalJSON handles deserialization of a Review.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
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
