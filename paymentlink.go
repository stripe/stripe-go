//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The specified behavior after the purchase is complete.
type PaymentLinkAfterCompletionType string

// List of values that PaymentLinkAfterCompletionType can take
const (
	PaymentLinkAfterCompletionTypeHostedConfirmation PaymentLinkAfterCompletionType = "hosted_confirmation"
	PaymentLinkAfterCompletionTypeRedirect           PaymentLinkAfterCompletionType = "redirect"
)

// Configuration for collecting the customer's billing address.
type PaymentLinkBillingAddressCollection string

// List of values that PaymentLinkBillingAddressCollection can take
const (
	PaymentLinkBillingAddressCollectionAuto     PaymentLinkBillingAddressCollection = "auto"
	PaymentLinkBillingAddressCollectionRequired PaymentLinkBillingAddressCollection = "required"
)

// If set to `auto`, enables the collection of customer consent for promotional communications.
type PaymentLinkConsentCollectionPromotions string

// List of values that PaymentLinkConsentCollectionPromotions can take
const (
	PaymentLinkConsentCollectionPromotionsAuto PaymentLinkConsentCollectionPromotions = "auto"
	PaymentLinkConsentCollectionPromotionsNone PaymentLinkConsentCollectionPromotions = "none"
)

// Configuration for Customer creation during checkout.
type PaymentLinkCustomerCreation string

// List of values that PaymentLinkCustomerCreation can take
const (
	PaymentLinkCustomerCreationAlways     PaymentLinkCustomerCreation = "always"
	PaymentLinkCustomerCreationIfRequired PaymentLinkCustomerCreation = "if_required"
)

// Indicates when the funds will be captured from the customer's account.
type PaymentLinkPaymentIntentDataCaptureMethod string

// List of values that PaymentLinkPaymentIntentDataCaptureMethod can take
const (
	PaymentLinkPaymentIntentDataCaptureMethodAutomatic PaymentLinkPaymentIntentDataCaptureMethod = "automatic"
	PaymentLinkPaymentIntentDataCaptureMethodManual    PaymentLinkPaymentIntentDataCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with the payment method collected during checkout.
type PaymentLinkPaymentIntentDataSetupFutureUsage string

// List of values that PaymentLinkPaymentIntentDataSetupFutureUsage can take
const (
	PaymentLinkPaymentIntentDataSetupFutureUsageOffSession PaymentLinkPaymentIntentDataSetupFutureUsage = "off_session"
	PaymentLinkPaymentIntentDataSetupFutureUsageOnSession  PaymentLinkPaymentIntentDataSetupFutureUsage = "on_session"
)

// The list of payment method types that customers can use. When `null`, Stripe will dynamically show relevant payment methods you've enabled in your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
type PaymentLinkPaymentMethodType string

// List of values that PaymentLinkPaymentMethodType can take
const (
	PaymentLinkPaymentMethodTypeAffirm           PaymentLinkPaymentMethodType = "affirm"
	PaymentLinkPaymentMethodTypeAfterpayClearpay PaymentLinkPaymentMethodType = "afterpay_clearpay"
	PaymentLinkPaymentMethodTypeAlipay           PaymentLinkPaymentMethodType = "alipay"
	PaymentLinkPaymentMethodTypeAUBECSDebit      PaymentLinkPaymentMethodType = "au_becs_debit"
	PaymentLinkPaymentMethodTypeBACSDebit        PaymentLinkPaymentMethodType = "bacs_debit"
	PaymentLinkPaymentMethodTypeBancontact       PaymentLinkPaymentMethodType = "bancontact"
	PaymentLinkPaymentMethodTypeBLIK             PaymentLinkPaymentMethodType = "blik"
	PaymentLinkPaymentMethodTypeBoleto           PaymentLinkPaymentMethodType = "boleto"
	PaymentLinkPaymentMethodTypeCard             PaymentLinkPaymentMethodType = "card"
	PaymentLinkPaymentMethodTypeEPS              PaymentLinkPaymentMethodType = "eps"
	PaymentLinkPaymentMethodTypeFPX              PaymentLinkPaymentMethodType = "fpx"
	PaymentLinkPaymentMethodTypeGiropay          PaymentLinkPaymentMethodType = "giropay"
	PaymentLinkPaymentMethodTypeGrabpay          PaymentLinkPaymentMethodType = "grabpay"
	PaymentLinkPaymentMethodTypeIdeal            PaymentLinkPaymentMethodType = "ideal"
	PaymentLinkPaymentMethodTypeKlarna           PaymentLinkPaymentMethodType = "klarna"
	PaymentLinkPaymentMethodTypeKonbini          PaymentLinkPaymentMethodType = "konbini"
	PaymentLinkPaymentMethodTypeOXXO             PaymentLinkPaymentMethodType = "oxxo"
	PaymentLinkPaymentMethodTypeP24              PaymentLinkPaymentMethodType = "p24"
	PaymentLinkPaymentMethodTypePayNow           PaymentLinkPaymentMethodType = "paynow"
	PaymentLinkPaymentMethodTypePromptPay        PaymentLinkPaymentMethodType = "promptpay"
	PaymentLinkPaymentMethodTypeSepaDebit        PaymentLinkPaymentMethodType = "sepa_debit"
	PaymentLinkPaymentMethodTypeSofort           PaymentLinkPaymentMethodType = "sofort"
	PaymentLinkPaymentMethodTypeUSBankAccount    PaymentLinkPaymentMethodType = "us_bank_account"
	PaymentLinkPaymentMethodTypeWechatPay        PaymentLinkPaymentMethodType = "wechat_pay"
)

// Indicates the type of transaction being performed which customizes relevant text on the page, such as the submit button.
type PaymentLinkSubmitType string

// List of values that PaymentLinkSubmitType can take
const (
	PaymentLinkSubmitTypeAuto   PaymentLinkSubmitType = "auto"
	PaymentLinkSubmitTypeBook   PaymentLinkSubmitType = "book"
	PaymentLinkSubmitTypeDonate PaymentLinkSubmitType = "donate"
	PaymentLinkSubmitTypePay    PaymentLinkSubmitType = "pay"
)

// Returns a list of your payment links.
type PaymentLinkListParams struct {
	ListParams `form:"*"`
	// Only return payment links that are active or inactive (e.g., pass `false` to list all inactive payment links).
	Active *bool `form:"active"`
}

// Configuration when `type=hosted_confirmation`.
type PaymentLinkAfterCompletionHostedConfirmationParams struct {
	// A custom message to display to the customer after the purchase is complete.
	CustomMessage *string `form:"custom_message"`
}

// Configuration when `type=redirect`.
type PaymentLinkAfterCompletionRedirectParams struct {
	// The URL the customer will be redirected to after the purchase is complete. You can embed `{CHECKOUT_SESSION_ID}` into the URL to have the `id` of the completed [checkout session](https://stripe.com/docs/api/checkout/sessions/object#checkout_session_object-id) included.
	URL *string `form:"url"`
}

// Behavior after the purchase is complete.
type PaymentLinkAfterCompletionParams struct {
	// Configuration when `type=hosted_confirmation`.
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmationParams `form:"hosted_confirmation"`
	// Configuration when `type=redirect`.
	Redirect *PaymentLinkAfterCompletionRedirectParams `form:"redirect"`
	// The specified behavior after the purchase is complete. Either `redirect` or `hosted_confirmation`.
	Type *string `form:"type"`
}

// Configuration for automatic tax collection.
type PaymentLinkAutomaticTaxParams struct {
	// If `true`, tax will be calculated automatically using the customer's location.
	Enabled *bool `form:"enabled"`
}

// Configure fields to gather active consent from customers.
type PaymentLinkConsentCollectionParams struct {
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
type PaymentLinkLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative Integer.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 99.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer can purchase. By default this value is 0. You can specify a value up to 98. If there is only one item in the cart then that item's quantity cannot go down to 0.
	Minimum *int64 `form:"minimum"`
}

// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
type PaymentLinkLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
	AdjustableQuantity *PaymentLinkLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The ID of an existing line item on the payment link.
	ID *string `form:"id"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) or [Plan](https://stripe.com/docs/api/plans) object.
	Price *string `form:"price"`
	// The quantity of the line item being purchased.
	Quantity *int64 `form:"quantity"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type PaymentLinkPaymentIntentDataParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to [make future payments](https://stripe.com/docs/payments/payment-intents#future-usage) with the payment method collected by this Checkout Session.
	//
	// When setting this to `on_session`, Checkout will show a notice to the customer that their payment details will be saved.
	//
	// When setting this to `off_session`, Checkout will show a notice to the customer that their payment details will be saved and used for future payments.
	//
	// If a Customer has been provided or Checkout creates a new Customer,Checkout will attach the payment method to the Customer.
	//
	// If Checkout does not create a Customer, the payment method is not attached to a Customer. To reuse the payment method, you can retrieve it from the Checkout Session's PaymentIntent.
	//
	// When processing card payments, Checkout also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as SCA.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Controls phone number collection settings during checkout.
//
// We recommend that you review your privacy policy and check with your legal contacts.
type PaymentLinkPhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	Enabled *bool `form:"enabled"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []*string `form:"allowed_countries"`
}

// The shipping rate options to apply to [checkout sessions](https://stripe.com/docs/api/checkout/sessions) created by this payment link.
type PaymentLinkShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionDataParams struct {
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days"`
}

// Controls tax ID collection during checkout.
type PaymentLinkTaxIDCollectionParams struct {
	// Set to `true` to enable tax ID collection.
	Enabled *bool `form:"enabled"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination"`
}

// Creates a payment link.
type PaymentLinkParams struct {
	Params `form:"*"`
	// Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.
	Active *bool `form:"active"`
	// Behavior after the purchase is complete.
	AfterCompletion *PaymentLinkAfterCompletionParams `form:"after_completion"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Can only be applied when there are no line items with recurring prices.
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. There must be at least 1 line item with a recurring price to use this field.
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Configuration for automatic tax collection.
	AutomaticTax *PaymentLinkAutomaticTaxParams `form:"automatic_tax"`
	// Configuration for collecting the customer's billing address.
	BillingAddressCollection *string `form:"billing_address_collection"`
	// Configure fields to gather active consent from customers.
	ConsentCollection *PaymentLinkConsentCollectionParams `form:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies) and supported by each line item's price.
	Currency *string `form:"currency"`
	// Configures whether [checkout sessions](https://stripe.com/docs/api/checkout/sessions) created by this payment link create a [Customer](https://stripe.com/docs/api/customers).
	CustomerCreation *string `form:"customer_creation"`
	// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
	LineItems []*PaymentLinkLineItemParams `form:"line_items"`
	// The account on behalf of which to charge.
	OnBehalfOf *string `form:"on_behalf_of"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *PaymentLinkPaymentIntentDataParams `form:"payment_intent_data"`
	// The list of payment method types that customers can use. Pass an empty string to enable automatic payment methods that use your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// Controls phone number collection settings during checkout.
	//
	// We recommend that you review your privacy policy and check with your legal contacts.
	PhoneNumberCollection *PaymentLinkPhoneNumberCollectionParams `form:"phone_number_collection"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkShippingAddressCollectionParams `form:"shipping_address_collection"`
	// The shipping rate options to apply to [checkout sessions](https://stripe.com/docs/api/checkout/sessions) created by this payment link.
	ShippingOptions []*PaymentLinkShippingOptionParams `form:"shipping_options"`
	// Describes the type of transaction being performed in order to customize relevant text on the page, such as the submit button. Changing this value will also affect the hostname in the [url](https://stripe.com/docs/api/payment_links/payment_links/object#url) property (example: `donate.stripe.com`).
	SubmitType *string `form:"submit_type"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkSubscriptionDataParams `form:"subscription_data"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *PaymentLinkTaxIDCollectionParams `form:"tax_id_collection"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkTransferDataParams `form:"transfer_data"`
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type PaymentLinkListLineItemsParams struct {
	ListParams  `form:"*"`
	PaymentLink *string `form:"-"` // Included in URL
}
type PaymentLinkAfterCompletionHostedConfirmation struct {
	// The custom message that is displayed to the customer after the purchase is complete.
	CustomMessage string `json:"custom_message"`
}
type PaymentLinkAfterCompletionRedirect struct {
	// The URL the customer will be redirected to after the purchase is complete.
	URL string `json:"url"`
}
type PaymentLinkAfterCompletion struct {
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmation `json:"hosted_confirmation"`
	Redirect           *PaymentLinkAfterCompletionRedirect           `json:"redirect"`
	// The specified behavior after the purchase is complete.
	Type PaymentLinkAfterCompletionType `json:"type"`
}
type PaymentLinkAutomaticTax struct {
	// If `true`, tax will be calculated automatically using the customer's location.
	Enabled bool `json:"enabled"`
}

// When set, provides configuration to gather active consent from customers.
type PaymentLinkConsentCollection struct {
	// If set to `auto`, enables the collection of customer consent for promotional communications.
	Promotions PaymentLinkConsentCollectionPromotions `json:"promotions"`
}

// Indicates the parameters to be passed to PaymentIntent creation during checkout.
type PaymentLinkPaymentIntentData struct {
	// Indicates when the funds will be captured from the customer's account.
	CaptureMethod PaymentLinkPaymentIntentDataCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with the payment method collected during checkout.
	SetupFutureUsage PaymentLinkPaymentIntentDataSetupFutureUsage `json:"setup_future_usage"`
}
type PaymentLinkPhoneNumberCollection struct {
	// If `true`, a phone number will be collected during checkout.
	Enabled bool `json:"enabled"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkShippingAddressCollection struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []string `json:"allowed_countries"`
}

// The shipping rate options applied to the session.
type PaymentLinkShippingOption struct {
	// A non-negative integer in cents representing how much to charge.
	ShippingAmount int64 `json:"shipping_amount"`
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *ShippingRate `json:"shipping_rate"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionData struct {
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int64 `json:"trial_period_days"`
}
type PaymentLinkTaxIDCollection struct {
	// Indicates whether tax ID collection is enabled for the session.
	Enabled bool `json:"enabled"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferData struct {
	// The amount in %s that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// The connected account receiving the transfer.
	Destination *Account `json:"destination"`
}

// A payment link is a shareable URL that will take your customers to a hosted payment page. A payment link can be shared and used multiple times.
//
// When a customer opens a payment link it will open a new [checkout session](https://stripe.com/docs/api/checkout/sessions) to render the payment page. You can use [checkout session events](https://stripe.com/docs/api/events/types#event_types-checkout.session.completed) to track payments through payment links.
//
// Related guide: [Payment Links API](https://stripe.com/docs/payments/payment-links/api)
type PaymentLink struct {
	APIResource
	// Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.
	Active          bool                        `json:"active"`
	AfterCompletion *PaymentLinkAfterCompletion `json:"after_completion"`
	// Whether user redeemable promotion codes are enabled.
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account.
	ApplicationFeePercent float64                  `json:"application_fee_percent"`
	AutomaticTax          *PaymentLinkAutomaticTax `json:"automatic_tax"`
	// Configuration for collecting the customer's billing address.
	BillingAddressCollection PaymentLinkBillingAddressCollection `json:"billing_address_collection"`
	// When set, provides configuration to gather active consent from customers.
	ConsentCollection *PaymentLinkConsentCollection `json:"consent_collection"`
	// Configuration for Customer creation during checkout.
	CustomerCreation PaymentLinkCustomerCreation `json:"customer_creation"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The line items representing what is being sold.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Indicates the parameters to be passed to PaymentIntent creation during checkout.
	PaymentIntentData *PaymentLinkPaymentIntentData `json:"payment_intent_data"`
	// The list of payment method types that customers can use. When `null`, Stripe will dynamically show relevant payment methods you've enabled in your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes    []PaymentLinkPaymentMethodType    `json:"payment_method_types"`
	PhoneNumberCollection *PaymentLinkPhoneNumberCollection `json:"phone_number_collection"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkShippingAddressCollection `json:"shipping_address_collection"`
	// The shipping rate options applied to the session.
	ShippingOptions []*PaymentLinkShippingOption `json:"shipping_options"`
	// Indicates the type of transaction being performed which customizes relevant text on the page, such as the submit button.
	SubmitType PaymentLinkSubmitType `json:"submit_type"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkSubscriptionData `json:"subscription_data"`
	TaxIDCollection  *PaymentLinkTaxIDCollection  `json:"tax_id_collection"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkTransferData `json:"transfer_data"`
	// The public URL that can be shared with customers.
	URL string `json:"url"`
}

// PaymentLinkList is a list of PaymentLinks as retrieved from a list endpoint.
type PaymentLinkList struct {
	APIResource
	ListMeta
	Data []*PaymentLink `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentLink.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentLink) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentLink PaymentLink
	var v paymentLink
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentLink(v)
	return nil
}
