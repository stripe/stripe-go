package stripe

import (
	"encoding/json"
)

// SubscriptionScheduleRevisionParams is used when retrieving a schedule revision.
type SubscriptionScheduleRevisionParams struct {
	Params   `form:"*"`
	Schedule *string `form:"-"` // Included in the URL
}

// SubscriptionScheduleRevisionListParams is the set of parameters that can be used when listing
// subscription schedules.
type SubscriptionScheduleRevisionListParams struct {
	ListParams `form:"*"`
	Schedule   *string `form:"-"` // Included in the URL
}

// SubscriptionScheduleRevision is the resource representing a Stripe subscription schedule revision.
type SubscriptionScheduleRevision struct {
	Created          int64                                `json:"created"`
	ID               string                               `json:"id"`
	Livemode         bool                                 `json:"livemode"`
	Metadata         map[string]string                    `json:"metadata"`
	Object           string                               `json:"object"`
	Phases           []*SubscriptionSchedulePhase         `json:"phases"`
	PreviousRevision *SubscriptionSchedule                `json:"previous_revision"`
	RenewalBehavior  SubscriptionScheduleRenewalBehavior  `json:"renewal_behavior"`
	RenewalInterval  *SubscriptionScheduleRenewalInterval `json:"renewal_interval"`
	Schedule         *SubscriptionSchedule                `json:"schedule"`
}

// SubscriptionScheduleRevisionList is a list object for subscription schedule revisions.
type SubscriptionScheduleRevisionList struct {
	ListMeta
	Data []*SubscriptionScheduleRevision `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionScheduleRevision.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionScheduleRevision) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type revision SubscriptionScheduleRevision
	var v revision
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionScheduleRevision(v)
	return nil
}
