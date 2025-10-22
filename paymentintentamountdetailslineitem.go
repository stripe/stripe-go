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
	// Total portion of the amount that is for tax.
	TotalTaxAmount int64 `json:"total_tax_amount"`
}
type PaymentIntentAmountDetailsLineItem struct {
	// The amount an item was discounted for. Positive integer.
	DiscountAmount int64 `json:"discount_amount"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Payment method-specific information for line items.
	PaymentMethodOptions *PaymentIntentAmountDetailsLineItemPaymentMethodOptions `json:"payment_method_options"`
	// Unique identifier of the product. At most 12 characters long.
	ProductCode string `json:"product_code"`
	// Name of the product. At most 100 characters long.
	ProductName string `json:"product_name"`
	// Number of items of the product. Positive integer.
	Quantity int64 `json:"quantity"`
	// Contains information about the tax on the item.
	Tax *PaymentIntentAmountDetailsLineItemTax `json:"tax"`
	// Cost of the product. Non-negative integer.
	UnitCost int64 `json:"unit_cost"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc.
	UnitOfMeasure string `json:"unit_of_measure"`
}

// PaymentIntentAmountDetailsLineItemList is a list of PaymentIntentAmountDetailsLineItems as retrieved from a list endpoint.
type PaymentIntentAmountDetailsLineItemList struct {
	APIResource
	ListMeta
	Data []*PaymentIntentAmountDetailsLineItem `json:"data"`
}
