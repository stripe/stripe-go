//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Enum describing the Stripe product that is managing this FinancialAccount.
type V2MoneyManagementFinancialAccountManagedByType string

// List of values that V2MoneyManagementFinancialAccountManagedByType can take
const (
	V2MoneyManagementFinancialAccountManagedByTypeMultiprocessorSettlement V2MoneyManagementFinancialAccountManagedByType = "multiprocessor_settlement"
)

// Closed Enum. An enum representing the status of the FinancialAccount. This indicates whether or not the FinancialAccount can be used for any money movement flows.
type V2MoneyManagementFinancialAccountStatus string

// List of values that V2MoneyManagementFinancialAccountStatus can take
const (
	V2MoneyManagementFinancialAccountStatusClosed  V2MoneyManagementFinancialAccountStatus = "closed"
	V2MoneyManagementFinancialAccountStatusOpen    V2MoneyManagementFinancialAccountStatus = "open"
	V2MoneyManagementFinancialAccountStatusPending V2MoneyManagementFinancialAccountStatus = "pending"
)

type V2MoneyManagementFinancialAccountStatusDetailsClosedReason string

// List of values that V2MoneyManagementFinancialAccountStatusDetailsClosedReason can take
const (
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonAccountClosed    V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "account_closed"
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonClosedByPlatform V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "closed_by_platform"
	V2MoneyManagementFinancialAccountStatusDetailsClosedReasonOther            V2MoneyManagementFinancialAccountStatusDetailsClosedReason = "other"
)

// Type of the FinancialAccount. An additional hash is included on the FinancialAccount with a name matching this value.
// It contains additional information specific to the FinancialAccount type.
type V2MoneyManagementFinancialAccountType string

// List of values that V2MoneyManagementFinancialAccountType can take
const (
	V2MoneyManagementFinancialAccountTypeOther    V2MoneyManagementFinancialAccountType = "other"
	V2MoneyManagementFinancialAccountTypePayments V2MoneyManagementFinancialAccountType = "payments"
	V2MoneyManagementFinancialAccountTypeStorage  V2MoneyManagementFinancialAccountType = "storage"
)

// Balance that can be used for money movement.
type V2MoneyManagementFinancialAccountBalanceAvailable struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Balance of inbound funds that will later transition to the `available` balance.
type V2MoneyManagementFinancialAccountBalanceInboundPending struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Balance of funds that are being used for a pending outbound money movement.
type V2MoneyManagementFinancialAccountBalanceOutboundPending struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Multi-currency balance of this FinancialAccount, split by availability state. Each balance is represented as a hash where the key is the three-letter ISO currency code, in lowercase, and the value is the amount for that currency.
type V2MoneyManagementFinancialAccountBalance struct {
	// Balance that can be used for money movement.
	Available map[string]*V2MoneyManagementFinancialAccountBalanceAvailable `json:"available"`
	// Balance of inbound funds that will later transition to the `available` balance.
	InboundPending map[string]*V2MoneyManagementFinancialAccountBalanceInboundPending `json:"inbound_pending"`
	// Balance of funds that are being used for a pending outbound money movement.
	OutboundPending map[string]*V2MoneyManagementFinancialAccountBalanceOutboundPending `json:"outbound_pending"`
}

// If this is a managed FinancialAccount, `managed_by` indicates the product that created and manages this FinancialAccount. For managed FinancialAccounts,
// creation of money management resources can only be orchestrated by the managing product.
type V2MoneyManagementFinancialAccountManagedBy struct {
	// Enum describing the Stripe product that is managing this FinancialAccount.
	Type V2MoneyManagementFinancialAccountManagedByType `json:"type"`
}

// If this is a `other` FinancialAccount, this hash indicates what the actual type is. Upgrade your API version to see it reflected in `type`.
type V2MoneyManagementFinancialAccountOther struct {
	// The type of the FinancialAccount, represented as a string. Upgrade your API version to see the type reflected in `financial_account.type`.
	Type string `json:"type"`
}

// If this is a `payments` FinancialAccount, this hash include details specific to `payments` FinancialAccount.
type V2MoneyManagementFinancialAccountPayments struct {
	// The currency that non-settlement currency payments will be converted to.
	DefaultCurrency Currency `json:"default_currency"`
	// Settlement currencies enabled for this FinancialAccount. Payments in other currencies will be automatically converted to `default_currency`.
	SettlementCurrencies []Currency `json:"settlement_currencies"`
}
type V2MoneyManagementFinancialAccountStatusDetailsClosedForwardingSettings struct {
	// The address to send forwarded payments to.
	PaymentMethod string `json:"payment_method,omitempty"`
	// The address to send forwarded payouts to.
	PayoutMethod string `json:"payout_method,omitempty"`
}
type V2MoneyManagementFinancialAccountStatusDetailsClosed struct {
	ForwardingSettings *V2MoneyManagementFinancialAccountStatusDetailsClosedForwardingSettings `json:"forwarding_settings,omitempty"`
	Reason             V2MoneyManagementFinancialAccountStatusDetailsClosedReason              `json:"reason"`
}
type V2MoneyManagementFinancialAccountStatusDetails struct {
	Closed *V2MoneyManagementFinancialAccountStatusDetailsClosed `json:"closed,omitempty"`
}

// If this is a `storage` FinancialAccount, this hash includes details specific to `storage` FinancialAccounts.
type V2MoneyManagementFinancialAccountStorage struct {
	// The currencies that this FinancialAccount can hold.
	HoldsCurrencies []Currency `json:"holds_currencies"`
}

// A FinancialAccount represents a balance and can be used as the source or destination for the money management (`/v2/money_management`) APIs.
type V2MoneyManagementFinancialAccount struct {
	APIResource
	// Multi-currency balance of this FinancialAccount, split by availability state. Each balance is represented as a hash where the key is the three-letter ISO currency code, in lowercase, and the value is the amount for that currency.
	Balance *V2MoneyManagementFinancialAccountBalance `json:"balance"`
	// Open Enum. Two-letter country code that represents the country where the LegalEntity associated with the FinancialAccount is based in.
	Country string `json:"country"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// If this is a managed FinancialAccount, `managed_by` indicates the product that created and manages this FinancialAccount. For managed FinancialAccounts,
	// creation of money management resources can only be orchestrated by the managing product.
	ManagedBy *V2MoneyManagementFinancialAccountManagedBy `json:"managed_by,omitempty"`
	// Metadata associated with the FinancialAccount.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// If this is a `other` FinancialAccount, this hash indicates what the actual type is. Upgrade your API version to see it reflected in `type`.
	Other *V2MoneyManagementFinancialAccountOther `json:"other,omitempty"`
	// If this is a `payments` FinancialAccount, this hash include details specific to `payments` FinancialAccount.
	Payments *V2MoneyManagementFinancialAccountPayments `json:"payments,omitempty"`
	// Closed Enum. An enum representing the status of the FinancialAccount. This indicates whether or not the FinancialAccount can be used for any money movement flows.
	Status        V2MoneyManagementFinancialAccountStatus         `json:"status"`
	StatusDetails *V2MoneyManagementFinancialAccountStatusDetails `json:"status_details,omitempty"`
	// If this is a `storage` FinancialAccount, this hash includes details specific to `storage` FinancialAccounts.
	Storage *V2MoneyManagementFinancialAccountStorage `json:"storage,omitempty"`
	// Type of the FinancialAccount. An additional hash is included on the FinancialAccount with a name matching this value.
	// It contains additional information specific to the FinancialAccount type.
	Type V2MoneyManagementFinancialAccountType `json:"type"`
}
