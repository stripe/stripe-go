//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

type BillingAlertTriggered struct {
	// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
	Alert *BillingAlert `json:"alert"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// ID of customer for which the alert triggered
	Customer string `json:"customer"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The value triggering the alert
	Value int64 `json:"value"`
}

// UnmarshalJSON handles deserialization of a BillingAlertTriggered.
// This custom unmarshaling is needed to handle the time fields correctly.
func (b *BillingAlertTriggered) UnmarshalJSON(data []byte) error {
	type billingAlertTriggered BillingAlertTriggered
	v := struct {
		Created int64 `json:"created"`
		*billingAlertTriggered
	}{
		billingAlertTriggered: (*billingAlertTriggered)(b),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	b.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of a BillingAlertTriggered.
// This custom marshaling is needed to handle the time fields correctly.
func (b BillingAlertTriggered) MarshalJSON() ([]byte, error) {
	type billingAlertTriggered BillingAlertTriggered
	v := struct {
		Created int64 `json:"created"`
		billingAlertTriggered
	}{
		billingAlertTriggered: (billingAlertTriggered)(b),
		Created:               b.Created.Unix(),
	}
	bb, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bb, err
}
