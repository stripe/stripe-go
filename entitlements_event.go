//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Who/what created this grant entitlement event.
type EntitlementsEventGrantGrantedBy string

// List of values that EntitlementsEventGrantGrantedBy can take
const (
	EntitlementsEventGrantGrantedBySubscriptionItem EntitlementsEventGrantGrantedBy = "subscription_item"
	EntitlementsEventGrantGrantedByUser             EntitlementsEventGrantGrantedBy = "user"
)

// Who/what created this revoke entitlement event.
type EntitlementsEventRevokeRevokedBy string

// List of values that EntitlementsEventRevokeRevokedBy can take
const (
	EntitlementsEventRevokeRevokedBySubscriptionItem EntitlementsEventRevokeRevokedBy = "subscription_item"
	EntitlementsEventRevokeRevokedByUser             EntitlementsEventRevokeRevokedBy = "user"
)

// Whether the event is a grant or revocation of the feature.
type EntitlementsEventType string

// List of values that EntitlementsEventType can take
const (
	EntitlementsEventTypeGrant  EntitlementsEventType = "grant"
	EntitlementsEventTypeRevoke EntitlementsEventType = "revoke"
)

// Contains information about type=grant entitlement event.
type EntitlementsEventGrantParams struct {
	// When manually creating a grant entitlement event, you can make it a temporary grant by setting it to expire.
	ExpiresAt *int64 `form:"expires_at"`
}

// Contains information about entitlement events relating to features with type=quantity. Required when the feature has type=quantity.
type EntitlementsEventQuantityParams struct {
	// When granting or revoking an entitlement to a type=quantity feature, you must specify the number of units you're granting/revoking. When the entitlement event type=grant, this number increments the total quantity available to the customer, and when type=revoke it decrements the total quantity available to the customer.
	Units *int64 `form:"units"`
}

// Create an entitlement event manually, outside of the entitlement events automatically created by Stripe lifecycle events.
type EntitlementsEventParams struct {
	Params `form:"*"`
	// The customer that is being granted or revoked entitlement to/from a feature.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The feature that the customer is being granted/revoked entitlement to/from.
	Feature *string `form:"feature"`
	// Contains information about type=grant entitlement event.
	Grant *EntitlementsEventGrantParams `form:"grant"`
	// Contains information about entitlement events relating to features with type=quantity. Required when the feature has type=quantity.
	Quantity *EntitlementsEventQuantityParams `form:"quantity"`
	// Whether the event is a grant or revocation of the feature.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *EntitlementsEventParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// If this entitlement event was created by a subscription_item, this will contains information about the subscription_item.
type EntitlementsEventGrantSubscriptionItem struct {
	// The subscription line item quantity.
	ItemQuantity int64 `json:"item_quantity"`
	// The price for the product that contains the feature for this entitlement event.
	Price string `json:"price"`
	// The product that contains the feature for this entitlement event.
	Product string `json:"product"`
	// The subscription that created this entitlement event.
	Subscription string `json:"subscription"`
	// The subscription item that created this entitlement event.
	SubscriptionItem string `json:"subscription_item"`
}

// Contains information about type=grant entitlement event.
type EntitlementsEventGrant struct {
	// When manually creating a grant entitlement event, you can make it a temporary grant by setting it to expire.
	ExpiresAt int64 `json:"expires_at"`
	// Who/what created this grant entitlement event.
	GrantedBy EntitlementsEventGrantGrantedBy `json:"granted_by"`
	// If this entitlement event was created by a subscription_item, this will contains information about the subscription_item.
	SubscriptionItem *EntitlementsEventGrantSubscriptionItem `json:"subscription_item"`
}

// Contains information about entitlement events relating to features with type=quantity. Required when the feature has type=quantity.
type EntitlementsEventQuantity struct {
	// When granting or revoking an entitlement to a type=quantity feature, you must specify the number of units you're granting/revoking. When the entitlement event type=grant, this number increments the total quantity available to the customer, and when type=revoke it decrements the total quantity available to the customer.
	Units int64 `json:"units"`
}

// If this entitlement event was created by a subscription_item, this will contains information about the subscription_item.
type EntitlementsEventRevokeSubscriptionItem struct {
	// The subscription line item quantity.
	ItemQuantity int64 `json:"item_quantity"`
	// The price for the product that contains the feature for this entitlement event.
	Price string `json:"price"`
	// The product that contains the feature for this entitlement event.
	Product string `json:"product"`
	// The subscription that created this entitlement event.
	Subscription string `json:"subscription"`
	// The subscription item that created this entitlement event.
	SubscriptionItem string `json:"subscription_item"`
}

// Contains information about type=revoke entitlement event.
type EntitlementsEventRevoke struct {
	// A revoke entitlement event will always expire at the same time as the grant it is revoking.
	ExpiresAt int64 `json:"expires_at"`
	// Who/what created this revoke entitlement event.
	RevokedBy EntitlementsEventRevokeRevokedBy `json:"revoked_by"`
	// If this entitlement event was created by a subscription_item, this will contains information about the subscription_item.
	SubscriptionItem *EntitlementsEventRevokeSubscriptionItem `json:"subscription_item"`
}

// An entitlement event either grants or revokes an entitlement to a feature for a customer.
type EntitlementsEvent struct {
	APIResource
	// The customer that is being granted or revoked entitlement to/from a feature.
	Customer string `json:"customer"`
	// The feature that the customer is being granted/revoked entitlement to/from.
	Feature string `json:"feature"`
	// Contains information about type=grant entitlement event.
	Grant *EntitlementsEventGrant `json:"grant"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Contains information about entitlement events relating to features with type=quantity. Required when the feature has type=quantity.
	Quantity *EntitlementsEventQuantity `json:"quantity"`
	// Contains information about type=revoke entitlement event.
	Revoke *EntitlementsEventRevoke `json:"revoke"`
	// Whether the event is a grant or revocation of the feature.
	Type EntitlementsEventType `json:"type"`
}
