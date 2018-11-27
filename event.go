package stripe

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	Account         string        `json:"account"`
	Created         int64         `json:"created"`
	Data            *EventData    `json:"data"`
	ID              string        `json:"id"`
	Livemode        bool          `json:"livemode"`
	PendingWebhooks int64         `json:"pending_webhooks"`
	Request         *EventRequest `json:"request"`
	Type            string        `json:"type"`
}

// internalEvent is used for custom unmarshalling so Data is
// unmarshalled into the correct object (Customer, Invoice, etc...)
type internalEvent struct {
	Event
	Data rawEventData `json:"data"`
}

// rawEventData is an interim struct holding the event data, using the event type
// we can determine what the actual EventData object should be and unmarshal Raw
type rawEventData struct {
	Raw                json.RawMessage        `json:"object"`
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
}

// EventRequest contains information on a request that created an event.
type EventRequest struct {
	// ID is the request ID of the request that created an event, if the event
	// was created by a request.
	ID string `json:"id"`

	// IdempotencyKey is the idempotency key of the request that created an
	// event, if the event was created by a request and if an idempotency key
	// was specified for that request.
	IdempotencyKey string `json:"idempotency_key"`
}

// EventData is the unmarshalled object as a map.
type EventData struct {
	Object               map[string]interface{}
	Account              *Account
	Application          *Application
	Card                 *Card
	BankAccount          *BankAccount
	ApplicationFee       *ApplicationFee
	FeeRefund            *FeeRefund
	Balance              *Balance
	Charge               *Charge
	Dispute              *Dispute
	Refund               *Refund
	Coupon               *Coupon
	Customer             *Customer
	Discount             *Discount
	Subscription         *Subscription
	File                 *File
	Invoice              *Invoice
	InvoiceItem          *InvoiceItem
	IssuingAuthorization *IssuingAuthorization
	IssuingCard          *IssuingCard
	IssuingCardholder    *IssuingCardholder
	IssuingDispute       *IssuingDispute
	IssuingTransaction   *IssuingTransaction
	Order                *Order
	PaymentIntent        *PaymentIntent
	Payout               *Payout
	Plan                 *Plan
	Product              *Product
	Recipient            *Recipient
	ReportRun            *ReportRun
	Review               *Review
	SKU                  *SKU
	Source               *Source
	SourceTransaction    *SourceTransaction
	Topup                *Topup
	Transfer             *Transfer
	PreviousAttributes   map[string]interface{} `json:"previous_attributes"`
	Raw                  json.RawMessage        `json:"object"`
}

// EventParams is the set of parameters that can be used when retrieving events.
// For more details see https://stripe.com/docs/api#retrieve_events.
type EventParams struct {
	Params `form:"*"`
}

// EventList is a list of events as retrieved from a list endpoint.
type EventList struct {
	ListMeta
	Data []*Event `json:"data"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams      `form:"*"`
	Created         *int64            `form:"created"`
	CreatedRange    *RangeQueryParams `form:"created"`
	DeliverySuccess *bool             `form:"delivery_success"`
	Type            *string           `form:"type"`
	Types           []*string         `form:"types"`
}

// GetObjectValue returns the value from the e.Data.Object bag based on the keys hierarchy.
func (e *Event) GetObjectValue(keys ...string) string {
	return getValue(e.Data.Object, keys)
}

// GetPreviousValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPreviousValue(keys ...string) string {
	return getValue(e.Data.PreviousAttributes, keys)
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var internal internalEvent
	if err := json.Unmarsal(data, &internal); err != nil {
		return err
	}
	switch internal.Type {
	case "account.updated":
		var account Account
		if err := json.Unmarshal([]byte(internal.Data.Raw), &account); err != nil {
			return err
		}
		*e = internal.Event
		e.Data.Account = &account
	case "account.application.authorized", "account.application.deauthorized":
		var application Application
		if err := json.Unmarshal([]byte(internal.Data.Raw), &application); err != nil {
			return err
		}
		*e = internal.Event
		e.Data.Application = &application
	case "account.external_account.created", "account.external_account.deleted", "account.external_account.updated":
		// find bank_account, get the type
		if strings.Contains(string(internal.Data.Raw), "bank_account") {
			var bankAccount BankAccount
			if err := json.Unmarshal([]byte(internal.Data.Raw), &bankAccount); err != nil {
				return err
			}
			*e = internal.Event
			e.Data.BankAccount = &bankAccount
		} else {
			var card Card
			if err := json.Unmarshal([]byte(internal.Data.Raw), &card); err != nil {
				return err
			}
			*e = internal.Event
			e.Data.Card = &card
		}
	}
	// transfer other fields
	e.Data.Raw = internal.Data.Raw
	e.Data.PreviousAttributes = internal.Data.PreviousAttributes
	return nil
}

// UnmarshalJSON handles deserialization of the EventData.
// This custom unmarshaling exists so that we can keep both the map and raw data.
func (e *EventData) UnmarshalJSON(data []byte) error {
	type eventdata EventData
	var ee eventdata
	err := json.Unmarshal(data, &ee)
	if err != nil {
		return err
	}

	*e = EventData(ee)
	return json.Unmarshal(e.Raw, &e.Object)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		key := keys[i]

		sliceNode, ok := node.([]interface{})
		if ok {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				panic(fmt.Sprintf(
					"Cannot access nested slice element with non-integer key: %s",
					key))
			}
			node = sliceNode[intKey]
			continue
		}

		mapNode, ok := node.(map[string]interface{})
		if ok {
			node = mapNode[key]
			continue
		}

		panic(fmt.Sprintf(
			"Cannot descend into non-map non-slice object with key: %s", key))
	}

	if node == nil {
		return ""
	}

	return fmt.Sprintf("%v", node)
}
