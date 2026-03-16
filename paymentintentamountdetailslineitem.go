//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of the line item.
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory string

// List of values that PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory can take
const (
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryDigitalGoods  PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "digital_goods"
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryDonation      PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "donation"
	PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategoryPhysicalGoods PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory = "physical_goods"
)

// Lists all LineItems of a given PaymentIntent.
type PaymentIntentAmountDetailsLineItemListParams struct {
	ListParams `form:"*"`
	Intent     *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentIntentAmountDetailsLineItemListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard struct {
	CommodityCode string `json:"commodity_code"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardPresent struct {
	CommodityCode string `json:"commodity_code"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna struct {
	ImageURL              string `json:"image_url"`
	ProductURL            string `json:"product_url"`
	Reference             string `json:"reference"`
	SubscriptionReference string `json:"subscription_reference"`
}
type PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypal struct {
	// Type of the line item.
	Category PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypalCategory `json:"category"`
	// Description of the line item.
	Description string `json:"description"`
	// The Stripe account ID of the connected account that sells the item. This is only needed when using [Separate Charges and Transfers](https://docs.stripe.com/connect/separate-charges-and-transfers).
	SoldBy string `json:"sold_by"`
}

// Payment method-specific information for line items.
type PaymentIntentAmountDetailsLineItemPaymentMethodOptions struct {
	Card        *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard        `json:"card"`
	CardPresent *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardPresent `json:"card_present"`
	Klarna      *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna      `json:"klarna"`
	Paypal      *PaymentIntentAmountDetailsLineItemPaymentMethodOptionsPaypal      `json:"paypal"`
}

// Contains information about the tax on the item.
type PaymentIntentAmountDetailsLineItemTax struct {
	// The total amount of tax on the transaction represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Required for L2 rates. An integer greater than or equal to 0.
	//
	// This field is mutually exclusive with the `amount_details[line_items][#][tax][total_tax_amount]` field.
	TotalTaxAmount int64 `json:"total_tax_amount"`
}
type PaymentIntentAmountDetailsLineItem struct {
	// The discount applied on this line item represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). An integer greater than 0.
	//
	// This field is mutually exclusive with the `amount_details[discount_amount]` field.
	DiscountAmount int64 `json:"discount_amount"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Payment method-specific information for line items.
	PaymentMethodOptions *PaymentIntentAmountDetailsLineItemPaymentMethodOptions `json:"payment_method_options"`
	// The product code of the line item, such as an SKU. Required for L3 rates. At most 12 characters long.
	ProductCode string `json:"product_code"`
	// The product name of the line item. Required for L3 rates. At most 1024 characters long.
	//
	// For Cards, this field is truncated to 26 alphanumeric characters before being sent to the card networks. For PayPal, this field is truncated to 127 characters.
	ProductName string `json:"product_name"`
	// The quantity of items. Required for L3 rates. An integer greater than 0.
	Quantity int64 `json:"quantity"`
	// Contains information about the tax on the item.
	Tax *PaymentIntentAmountDetailsLineItemTax `json:"tax"`
	// The unit cost of the line item represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Required for L3 rates. An integer greater than or equal to 0.
	UnitCost int64 `json:"unit_cost"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc. Required for L3 rates. At most 12 alphanumeric characters long.
	UnitOfMeasure string `json:"unit_of_measure"`
}

// PaymentIntentAmountDetailsLineItemList is a list of PaymentIntentAmountDetailsLineItems as retrieved from a list endpoint.
type PaymentIntentAmountDetailsLineItemList struct {
	APIResource
	ListMeta
	Data []*PaymentIntentAmountDetailsLineItem `json:"data"`
}
