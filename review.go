//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The reason the review was closed, or null if it has not yet been closed. One of `approved`, `refunded`, `refunded_as_fraud`, `disputed`, or `redacted`.
type ReviewClosedReason string

// List of values that ReviewClosedReason can take
const (
	ReviewClosedReasonApproved        ReviewClosedReason = "approved"
	ReviewClosedReasonDisputed        ReviewClosedReason = "disputed"
	ReviewClosedReasonRedacted        ReviewClosedReason = "redacted"
	ReviewClosedReasonRefunded        ReviewClosedReason = "refunded"
	ReviewClosedReasonRefundedAsFraud ReviewClosedReason = "refunded_as_fraud"
)

// The reason the review was opened. One of `rule` or `manual`.
type ReviewOpenedReason string

// List of values that ReviewOpenedReason can take
const (
	ReviewOpenedReasonManual ReviewOpenedReason = "manual"
	ReviewOpenedReasonRule   ReviewOpenedReason = "rule"
)

// The reason the review is currently open or closed. One of `rule`, `manual`, `approved`, `refunded`, `refunded_as_fraud`, `disputed`, or `redacted`.
type ReviewReason string

// Deprecated: we preserve this name for backwards-compatibility, prefer `ReviewReason`.
type ReviewReasonType = ReviewReason

// List of values that ReviewReason can take
const (
	ReviewReasonApproved        ReviewReason = "approved"
	ReviewReasonDisputed        ReviewReason = "disputed"
	ReviewReasonManual          ReviewReason = "manual"
	ReviewReasonRefunded        ReviewReason = "refunded"
	ReviewReasonRefundedAsFraud ReviewReason = "refunded_as_fraud"
	ReviewReasonRule            ReviewReason = "rule"
)

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type ReviewListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Retrieves a Review object.
type ReviewParams struct {
	Params `form:"*"`
}

// Approves a Review object, closing it and removing it from the list of reviews.
type ReviewApproveParams struct {
	Params `form:"*"`
}

// Information related to the location of the payment. Note that this information is an approximation and attempts to locate the nearest population center - it should not be used to determine a specific address.
type ReviewIPAddressLocation struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Region    string  `json:"region"`
}

// Information related to the browsing session of the user who initiated the payment.
type ReviewSession struct {
	Browser  string `json:"browser"`
	Device   string `json:"device"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

// Reviews can be used to supplement automated fraud detection with human expertise.
//
// Learn more about [Radar](https://stripe.com/radar) and reviewing payments
// [here](https://stripe.com/docs/radar/reviews).
type Review struct {
	APIResource
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
	Reason            ReviewReason             `json:"reason"`
	Session           *ReviewSession           `json:"session"`
}

// ReviewList is a list of Reviews as retrieved from a list endpoint.
type ReviewList struct {
	APIResource
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
