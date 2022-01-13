//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Surfaces if automatic tax computation is possible given the current customer location information.
type CustomerTaxAutomaticTax string

// List of values that CustomerTaxAutomaticTax can take
const (
	CustomerTaxAutomaticTaxFailed               CustomerTaxAutomaticTax = "failed"
	CustomerTaxAutomaticTaxNotCollecting        CustomerTaxAutomaticTax = "not_collecting"
	CustomerTaxAutomaticTaxSupported            CustomerTaxAutomaticTax = "supported"
	CustomerTaxAutomaticTaxUnrecognizedLocation CustomerTaxAutomaticTax = "unrecognized_location"
)

// The data source used to infer the customer's location.
type CustomerTaxLocationSource string

// List of values that CustomerTaxLocationSource can take
const (
	CustomerTaxLocationSourceBillingAddress      CustomerTaxLocationSource = "billing_address"
	CustomerTaxLocationSourceIPAddress           CustomerTaxLocationSource = "ip_address"
	CustomerTaxLocationSourcePaymentMethod       CustomerTaxLocationSource = "payment_method"
	CustomerTaxLocationSourceShippingDestination CustomerTaxLocationSource = "shipping_destination"
)

// Describes the customer's tax exemption status. One of `none`, `exempt`, or `reverse`. When set to `reverse`, invoice and receipt PDFs include the text **"Reverse charge"**.
type CustomerTaxExempt string

// List of values that CustomerTaxExempt can take
const (
	CustomerTaxExemptExempt  CustomerTaxExempt = "exempt"
	CustomerTaxExemptNone    CustomerTaxExempt = "none"
	CustomerTaxExemptReverse CustomerTaxExempt = "reverse"
)

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
type CustomerListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
}

// Default custom fields to be displayed on invoices for this customer. When updating, pass an empty string to remove previously-defined fields.
type CustomerInvoiceCustomFieldParams struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}

// Default invoice settings for this customer.
type CustomerInvoiceSettingsParams struct {
	CustomFields         []*CustomerInvoiceCustomFieldParams `form:"custom_fields"`
	DefaultPaymentMethod *string                             `form:"default_payment_method"`
	Footer               *string                             `form:"footer"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type CustomerShippingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// Tax details about the customer.
type CustomerTaxParams struct {
	IPAddress *string `form:"ip_address"`
}

// The customer's tax IDs.
type CustomerTaxIDDataParams struct {
	Type  *string `form:"type"`
	Value *string `form:"value"`
}

// SetSource adds valid sources to a CustomerParams object,
// returning an error for unsupported sources.
func (cp *CustomerParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	cp.Source = source
	return err
}

// Creates a new customer object.
type CustomerParams struct {
	Params              `form:"*"`
	Address             *AddressParams                 `form:"address"`
	Balance             *int64                         `form:"balance"`
	Coupon              *string                        `form:"coupon"`
	DefaultSource       *string                        `form:"default_source"`
	Description         *string                        `form:"description"`
	Email               *string                        `form:"email"`
	InvoicePrefix       *string                        `form:"invoice_prefix"`
	InvoiceSettings     *CustomerInvoiceSettingsParams `form:"invoice_settings"`
	Name                *string                        `form:"name"`
	NextInvoiceSequence *int64                         `form:"next_invoice_sequence"`
	PaymentMethod       *string                        `form:"payment_method"`
	Phone               *string                        `form:"phone"`
	PreferredLocales    []*string                      `form:"preferred_locales"`
	PromotionCode       *string                        `form:"promotion_code"`
	Shipping            *CustomerShippingDetailsParams `form:"shipping"`
	Source              *SourceParams                  `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	Tax                 *CustomerTaxParams             `form:"tax"`
	TaxExempt           *string                        `form:"tax_exempt"`
	TaxIDData           []*CustomerTaxIDDataParams     `form:"tax_id_data"`
	Token               *string                        `form:"-"` // This doesn't seem to be used?
}

// Returns a list of PaymentMethods for a given Customer
type CustomerListPaymentMethodsParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	Type       *string `form:"type"`
}

// Default custom fields to be displayed on invoices for this customer.
type CustomerInvoiceCustomField struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}
type CustomerInvoiceSettings struct {
	CustomFields         []*CustomerInvoiceCustomField `json:"custom_fields"`
	DefaultPaymentMethod *PaymentMethod                `json:"default_payment_method"`
	Footer               string                        `json:"footer"`
}

// Mailing and shipping address for the customer. Appears on invoices emailed to this customer.
type CustomerShippingDetails struct {
	Address Address `json:"address"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
}

// The customer's location as identified by Stripe Tax.
type CustomerTaxLocation struct {
	Country string                    `json:"country"`
	Source  CustomerTaxLocationSource `json:"source"`
	State   string                    `json:"state"`
}
type CustomerTax struct {
	AutomaticTax CustomerTaxAutomaticTax `json:"automatic_tax"`
	IPAddress    string                  `json:"ip_address"`
	Location     *CustomerTaxLocation    `json:"location"`
}

// This object represents a customer of your business. It lets you create recurring charges and track payments that belong to the same customer.
//
// Related guide: [Save a card during payment](https://stripe.com/docs/payments/save-during-payment).
type Customer struct {
	APIResource
	Address             Address                  `json:"address"`
	Balance             int64                    `json:"balance"`
	Created             int64                    `json:"created"`
	Currency            Currency                 `json:"currency"`
	DefaultSource       *PaymentSource           `json:"default_source"`
	Deleted             bool                     `json:"deleted"`
	Delinquent          bool                     `json:"delinquent"`
	Description         string                   `json:"description"`
	Discount            *Discount                `json:"discount"`
	Email               string                   `json:"email"`
	ID                  string                   `json:"id"`
	InvoicePrefix       string                   `json:"invoice_prefix"`
	InvoiceSettings     *CustomerInvoiceSettings `json:"invoice_settings"`
	Livemode            bool                     `json:"livemode"`
	Metadata            map[string]string        `json:"metadata"`
	Name                string                   `json:"name"`
	NextInvoiceSequence int64                    `json:"next_invoice_sequence"`
	Object              string                   `json:"object"`
	Phone               string                   `json:"phone"`
	PreferredLocales    []string                 `json:"preferred_locales"`
	Shipping            *CustomerShippingDetails `json:"shipping"`
	Sources             *SourceList              `json:"sources"`
	Subscriptions       *SubscriptionList        `json:"subscriptions"`
	Tax                 *CustomerTax             `json:"tax"`
	TaxExempt           CustomerTaxExempt        `json:"tax_exempt"`
	TaxIDs              *TaxIDList               `json:"tax_ids"`
}

// CustomerList is a list of Customers as retrieved from a list endpoint.
type CustomerList struct {
	APIResource
	ListMeta
	Data []*Customer `json:"data"`
}

// UnmarshalJSON handles deserialization of a Customer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Customer) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type customer Customer
	var v customer
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Customer(v)
	return nil
}
