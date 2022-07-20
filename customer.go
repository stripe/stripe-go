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

// Search for customers you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type CustomerSearchParams struct {
	SearchParams `form:"*"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
type CustomerListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// A case-sensitive filter on the list based on the customer's `email` field. The value must be a string.
	Email *string `form:"email"`
	// Provides a list of customers that are associated with the specified test clock. The response will not include customers with test clocks if this parameter is not set.
	TestClock *string `form:"test_clock"`
}

// Settings controlling the behavior of the customer's cash balance,
// such as reconciliation of funds received.
type CustomerCashBalanceSettingsParams struct {
	// Controls how funds transferred by the customer are applied to payment intents and invoices. Valid options are `automatic` or `manual`. For more information about these reconciliation modes, see [Reconciliation](https://stripe.com/docs/payments/customer-balance/reconciliation).
	ReconciliationMode *string `form:"reconciliation_mode"`
}

// Balance information and default balance settings for this customer.
type CustomerCashBalanceParams struct {
	// Settings controlling the behavior of the customer's cash balance,
	// such as reconciliation of funds received.
	Settings *CustomerCashBalanceSettingsParams `form:"settings"`
}

// Default custom fields to be displayed on invoices for this customer. When updating, pass an empty string to remove previously-defined fields.
type CustomerInvoiceCustomFieldParams struct {
	// The name of the custom field. This may be up to 30 characters.
	Name *string `form:"name"`
	// The value of the custom field. This may be up to 30 characters.
	Value *string `form:"value"`
}

// Default options for invoice PDF rendering for this customer.
type CustomerInvoiceSettingsRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display"`
}

// Default invoice settings for this customer.
type CustomerInvoiceSettingsParams struct {
	// Default custom fields to be displayed on invoices for this customer. When updating, pass an empty string to remove previously-defined fields.
	CustomFields []*CustomerInvoiceCustomFieldParams `form:"custom_fields"`
	// ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *CustomerInvoiceSettingsRenderingOptionsParams `form:"rendering_options"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type CustomerShippingDetailsParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address"`
	// Customer name.
	Name *string `form:"name"`
	// Customer phone (including extension).
	Phone *string `form:"phone"`
}

// Tax details about the customer.
type CustomerTaxParams struct {
	// A recent IP address of the customer used for tax reporting and tax location inference. Stripe recommends updating the IP address when a new PaymentMethod is attached or the address field on the customer is updated. We recommend against updating this field more frequently since it could result in unexpected tax location/reporting outcomes.
	IPAddress *string `form:"ip_address"`
}

// The customer's tax IDs.
type CustomerTaxIDDataParams struct {
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`
	Type *string `form:"type"`
	// Value of the tax ID.
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
	Params `form:"*"`
	// The customer's address.
	Address *AddressParams `form:"address"`
	// An integer amount in cents (or local equivalent) that represents the customer's current balance, which affect the customer's future invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases the amount due on an invoice.
	Balance *int64 `form:"balance"`
	// Balance information and default balance settings for this customer.
	CashBalance *CustomerCashBalanceParams `form:"cash_balance"`
	Coupon      *string                    `form:"coupon"`
	// If you are using payment methods created via the PaymentMethods API, see the [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/update#update_customer-invoice_settings-default_payment_method) parameter.
	//
	// Provide the ID of a payment source already attached to this customer to make it this customer's default payment source.
	//
	// If you want to add a new payment source and make it the default, see the [source](https://stripe.com/docs/api/customers/update#update_customer-source) property.
	DefaultSource *string `form:"default_source"`
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description *string `form:"description"`
	// Customer's email address. It's displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to *512 characters*.
	Email *string `form:"email"`
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix *string `form:"invoice_prefix"`
	// Default invoice settings for this customer.
	InvoiceSettings *CustomerInvoiceSettingsParams `form:"invoice_settings"`
	// The customer's full name or business name.
	Name *string `form:"name"`
	// The sequence to be used on the customer's next invoice. Defaults to 1.
	NextInvoiceSequence *int64  `form:"next_invoice_sequence"`
	PaymentMethod       *string `form:"payment_method"`
	// The customer's phone number.
	Phone *string `form:"phone"`
	// Customer's preferred languages, ordered by preference.
	PreferredLocales []*string `form:"preferred_locales"`
	// The API ID of a promotion code to apply to the customer. The customer will have a discount applied on all recurring payments. Charges you create through the API will not have the discount.
	PromotionCode *string `form:"promotion_code"`
	// The customer's shipping information. Appears on invoices emailed to this customer.
	Shipping *CustomerShippingDetailsParams `form:"shipping"`
	Source   *SourceParams                  `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	// Tax details about the customer.
	Tax *CustomerTaxParams `form:"tax"`
	// The customer's tax exemption. One of `none`, `exempt`, or `reverse`.
	TaxExempt *string `form:"tax_exempt"`
	// The customer's tax IDs.
	TaxIDData []*CustomerTaxIDDataParams `form:"tax_id_data"`
	// ID of the test clock to attach to the customer.
	TestClock *string `form:"test_clock"`
	Token     *string `form:"-"` // This doesn't seem to be used?
}

// Returns a list of PaymentMethods for a given Customer
type CustomerListPaymentMethodsParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	// A required filter on the list, based on the object `type` field.
	Type *string `form:"type"`
}

// Retrieves a PaymentMethod object for a given Customer.
type CustomerRetrievePaymentMethodParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
}

// Configuration for eu_bank_transfer funding type.
type CustomerCreateFundingInstructionsBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Additional parameters for `bank_transfer` funding types
type CustomerCreateFundingInstructionsBankTransferParams struct {
	// Configuration for eu_bank_transfer funding type.
	EUBankTransfer *CustomerCreateFundingInstructionsBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []*string `form:"requested_address_types"`
	// The type of the `bank_transfer`
	Type *string `form:"type"`
}

// Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
// funding instructions will be created. If funding instructions have already been created for a given customer, the same
// funding instructions will be retrieved. In other words, we will return the same funding instructions each time.
type CustomerCreateFundingInstructionsParams struct {
	Params `form:"*"`
	// Additional parameters for `bank_transfer` funding types
	BankTransfer *CustomerCreateFundingInstructionsBankTransferParams `form:"bank_transfer"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The `funding_type` to get the instructions for.
	FundingType *string `form:"funding_type"`
}

// Default custom fields to be displayed on invoices for this customer.
type CustomerInvoiceCustomField struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}

// Default options for invoice PDF rendering for this customer.
type CustomerInvoiceSettingsRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
}
type CustomerInvoiceSettings struct {
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*CustomerInvoiceCustomField `json:"custom_fields"`
	// ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Default footer to be displayed on invoices for this customer.
	Footer string `json:"footer"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *CustomerInvoiceSettingsRenderingOptions `json:"rendering_options"`
}

// Mailing and shipping address for the customer. Appears on invoices emailed to this customer.
type CustomerShippingDetails struct {
	Address Address `json:"address"`
	// Recipient name.
	Name string `json:"name"`
	// Recipient phone (including extension).
	Phone string `json:"phone"`
}

// The customer's location as identified by Stripe Tax.
type CustomerTaxLocation struct {
	// The customer's country as identified by Stripe Tax.
	Country string `json:"country"`
	// The data source used to infer the customer's location.
	Source CustomerTaxLocationSource `json:"source"`
	// The customer's state, county, province, or region as identified by Stripe Tax.
	State string `json:"state"`
}
type CustomerTax struct {
	// Surfaces if automatic tax computation is possible given the current customer location information.
	AutomaticTax CustomerTaxAutomaticTax `json:"automatic_tax"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IPAddress string `json:"ip_address"`
	// The customer's location as identified by Stripe Tax.
	Location *CustomerTaxLocation `json:"location"`
}

// This object represents a customer of your business. It lets you create recurring charges and track payments that belong to the same customer.
//
// Related guide: [Save a card during payment](https://stripe.com/docs/payments/save-during-payment).
type Customer struct {
	APIResource
	// The customer's address.
	Address Address `json:"address"`
	// Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized.
	Balance int64 `json:"balance"`
	// The current funds being held by Stripe on behalf of the customer. These funds can be applied towards payment intents with source "cash_balance".The settings[reconciliation_mode] field describes whether these funds are applied to such payment intents manually or automatically.
	CashBalance *CashBalance `json:"cash_balance"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) the customer can be charged in for recurring billing purposes.
	Currency Currency `json:"currency"`
	// The default three-letter [ISO code for the currency](https://stripe.com/docs/currencies) that the customer will be charged in for billing purposes.
	DefaultCurrency Currency `json:"default_currency"`
	// ID of the default payment source for the customer.
	//
	// If you are using payment methods created via the PaymentMethods API, see the [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.
	DefaultSource *PaymentSource `json:"default_source"`
	Deleted       bool           `json:"deleted"`
	// When the customer's latest invoice is billed by charging automatically, `delinquent` is `true` if the invoice's latest charge failed. When the customer's latest invoice is billed by sending an invoice, `delinquent` is `true` if the invoice isn't paid by its due date.
	//
	// If an invoice is marked uncollectible by [dunning](https://stripe.com/docs/billing/automatic-collection), `delinquent` doesn't get reset to `false`.
	Delinquent bool `json:"delinquent"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Describes the current discount active on the customer, if there is one.
	Discount *Discount `json:"discount"`
	// The customer's email address.
	Email string `json:"email"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The current multi-currency balances, if any, being stored on the customer.If positive in a currency, the customer has a credit to apply to their next invoice denominated in that currency.If negative, the customer has an amount owed that will be added to their next invoice denominated in that currency. These balances do not refer to any unpaid invoices.They solely track amounts that have yet to be successfully applied to any invoice. A balance in a particular currency is only applied to any invoice as an invoice in that currency is finalized.
	InvoiceCreditBalance map[string]int64 `json:"invoice_credit_balance"`
	// The prefix for the customer used to generate unique invoice numbers.
	InvoicePrefix   string                   `json:"invoice_prefix"`
	InvoiceSettings *CustomerInvoiceSettings `json:"invoice_settings"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The customer's full name or business name.
	Name string `json:"name"`
	// The suffix of the customer's next invoice number, e.g., 0001.
	NextInvoiceSequence int64 `json:"next_invoice_sequence"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The customer's phone number.
	Phone string `json:"phone"`
	// The customer's preferred locales (languages), ordered by preference.
	PreferredLocales []string `json:"preferred_locales"`
	// Mailing and shipping address for the customer. Appears on invoices emailed to this customer.
	Shipping *CustomerShippingDetails `json:"shipping"`
	Sources  *SourceList              `json:"sources"`
	// The customer's current subscriptions, if any.
	Subscriptions *SubscriptionList `json:"subscriptions"`
	Tax           *CustomerTax      `json:"tax"`
	// Describes the customer's tax exemption status. One of `none`, `exempt`, or `reverse`. When set to `reverse`, invoice and receipt PDFs include the text **"Reverse charge"**.
	TaxExempt CustomerTaxExempt `json:"tax_exempt"`
	// The customer's tax IDs.
	TaxIDs *TaxIDList `json:"tax_ids"`
	// ID of the test clock this customer belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// CustomerList is a list of Customers as retrieved from a list endpoint.
type CustomerList struct {
	APIResource
	ListMeta
	Data []*Customer `json:"data"`
}

// CustomerSearchResult is a list of Customer search results as retrieved from a search endpoint.
type CustomerSearchResult struct {
	APIResource
	SearchMeta
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
