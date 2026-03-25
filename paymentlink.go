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

// Type of the account referenced.
type PaymentLinkAutomaticTaxLiabilityType string

// List of values that PaymentLinkAutomaticTaxLiabilityType can take
const (
	PaymentLinkAutomaticTaxLiabilityTypeAccount PaymentLinkAutomaticTaxLiabilityType = "account"
	PaymentLinkAutomaticTaxLiabilityTypeSelf    PaymentLinkAutomaticTaxLiabilityType = "self"
)

// Configuration for collecting the customer's billing address. Defaults to `auto`.
type PaymentLinkBillingAddressCollection string

// List of values that PaymentLinkBillingAddressCollection can take
const (
	PaymentLinkBillingAddressCollectionAuto     PaymentLinkBillingAddressCollection = "auto"
	PaymentLinkBillingAddressCollectionRequired PaymentLinkBillingAddressCollection = "required"
)

// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.
//
// When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
type PaymentLinkConsentCollectionPaymentMethodReuseAgreementPosition string

// List of values that PaymentLinkConsentCollectionPaymentMethodReuseAgreementPosition can take
const (
	PaymentLinkConsentCollectionPaymentMethodReuseAgreementPositionAuto   PaymentLinkConsentCollectionPaymentMethodReuseAgreementPosition = "auto"
	PaymentLinkConsentCollectionPaymentMethodReuseAgreementPositionHidden PaymentLinkConsentCollectionPaymentMethodReuseAgreementPosition = "hidden"
)

// If set to `auto`, enables the collection of customer consent for promotional communications.
type PaymentLinkConsentCollectionPromotions string

// List of values that PaymentLinkConsentCollectionPromotions can take
const (
	PaymentLinkConsentCollectionPromotionsAuto PaymentLinkConsentCollectionPromotions = "auto"
	PaymentLinkConsentCollectionPromotionsNone PaymentLinkConsentCollectionPromotions = "none"
)

// If set to `required`, it requires cutomers to accept the terms of service before being able to pay. If set to `none`, customers won't be shown a checkbox to accept the terms of service.
type PaymentLinkConsentCollectionTermsOfService string

// List of values that PaymentLinkConsentCollectionTermsOfService can take
const (
	PaymentLinkConsentCollectionTermsOfServiceNone     PaymentLinkConsentCollectionTermsOfService = "none"
	PaymentLinkConsentCollectionTermsOfServiceRequired PaymentLinkConsentCollectionTermsOfService = "required"
)

// The type of the label.
type PaymentLinkCustomFieldLabelType string

// List of values that PaymentLinkCustomFieldLabelType can take
const (
	PaymentLinkCustomFieldLabelTypeCustom PaymentLinkCustomFieldLabelType = "custom"
)

// The type of the field.
type PaymentLinkCustomFieldType string

// List of values that PaymentLinkCustomFieldType can take
const (
	PaymentLinkCustomFieldTypeDropdown PaymentLinkCustomFieldType = "dropdown"
	PaymentLinkCustomFieldTypeNumeric  PaymentLinkCustomFieldType = "numeric"
	PaymentLinkCustomFieldTypeText     PaymentLinkCustomFieldType = "text"
)

// Configuration for Customer creation during checkout.
type PaymentLinkCustomerCreation string

// List of values that PaymentLinkCustomerCreation can take
const (
	PaymentLinkCustomerCreationAlways     PaymentLinkCustomerCreation = "always"
	PaymentLinkCustomerCreationIfRequired PaymentLinkCustomerCreation = "if_required"
)

// Type of the account referenced.
type PaymentLinkInvoiceCreationInvoiceDataIssuerType string

// List of values that PaymentLinkInvoiceCreationInvoiceDataIssuerType can take
const (
	PaymentLinkInvoiceCreationInvoiceDataIssuerTypeAccount PaymentLinkInvoiceCreationInvoiceDataIssuerType = "account"
	PaymentLinkInvoiceCreationInvoiceDataIssuerTypeSelf    PaymentLinkInvoiceCreationInvoiceDataIssuerType = "self"
)

// Indicates when the funds will be captured from the customer's account.
type PaymentLinkPaymentIntentDataCaptureMethod string

// List of values that PaymentLinkPaymentIntentDataCaptureMethod can take
const (
	PaymentLinkPaymentIntentDataCaptureMethodAutomatic      PaymentLinkPaymentIntentDataCaptureMethod = "automatic"
	PaymentLinkPaymentIntentDataCaptureMethodAutomaticAsync PaymentLinkPaymentIntentDataCaptureMethod = "automatic_async"
	PaymentLinkPaymentIntentDataCaptureMethodManual         PaymentLinkPaymentIntentDataCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with the payment method collected during checkout.
type PaymentLinkPaymentIntentDataSetupFutureUsage string

// List of values that PaymentLinkPaymentIntentDataSetupFutureUsage can take
const (
	PaymentLinkPaymentIntentDataSetupFutureUsageOffSession PaymentLinkPaymentIntentDataSetupFutureUsage = "off_session"
	PaymentLinkPaymentIntentDataSetupFutureUsageOnSession  PaymentLinkPaymentIntentDataSetupFutureUsage = "on_session"
)

// Configuration for collecting a payment method during checkout. Defaults to `always`.
type PaymentLinkPaymentMethodCollection string

// List of values that PaymentLinkPaymentMethodCollection can take
const (
	PaymentLinkPaymentMethodCollectionAlways     PaymentLinkPaymentMethodCollection = "always"
	PaymentLinkPaymentMethodCollectionIfRequired PaymentLinkPaymentMethodCollection = "if_required"
)

// The list of payment method types that customers can use. When `null`, Stripe will dynamically show relevant payment methods you've enabled in your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
type PaymentLinkPaymentMethodType string

// List of values that PaymentLinkPaymentMethodType can take
const (
	PaymentLinkPaymentMethodTypeAffirm           PaymentLinkPaymentMethodType = "affirm"
	PaymentLinkPaymentMethodTypeAfterpayClearpay PaymentLinkPaymentMethodType = "afterpay_clearpay"
	PaymentLinkPaymentMethodTypeAlipay           PaymentLinkPaymentMethodType = "alipay"
	PaymentLinkPaymentMethodTypeAlma             PaymentLinkPaymentMethodType = "alma"
	PaymentLinkPaymentMethodTypeAUBECSDebit      PaymentLinkPaymentMethodType = "au_becs_debit"
	PaymentLinkPaymentMethodTypeBACSDebit        PaymentLinkPaymentMethodType = "bacs_debit"
	PaymentLinkPaymentMethodTypeBancontact       PaymentLinkPaymentMethodType = "bancontact"
	PaymentLinkPaymentMethodTypeBillie           PaymentLinkPaymentMethodType = "billie"
	PaymentLinkPaymentMethodTypeBLIK             PaymentLinkPaymentMethodType = "blik"
	PaymentLinkPaymentMethodTypeBoleto           PaymentLinkPaymentMethodType = "boleto"
	PaymentLinkPaymentMethodTypeCard             PaymentLinkPaymentMethodType = "card"
	PaymentLinkPaymentMethodTypeCashApp          PaymentLinkPaymentMethodType = "cashapp"
	PaymentLinkPaymentMethodTypeEPS              PaymentLinkPaymentMethodType = "eps"
	PaymentLinkPaymentMethodTypeFPX              PaymentLinkPaymentMethodType = "fpx"
	PaymentLinkPaymentMethodTypeGiropay          PaymentLinkPaymentMethodType = "giropay"
	PaymentLinkPaymentMethodTypeGopay            PaymentLinkPaymentMethodType = "gopay"
	PaymentLinkPaymentMethodTypeGrabpay          PaymentLinkPaymentMethodType = "grabpay"
	PaymentLinkPaymentMethodTypeIDEAL            PaymentLinkPaymentMethodType = "ideal"
	PaymentLinkPaymentMethodTypeKlarna           PaymentLinkPaymentMethodType = "klarna"
	PaymentLinkPaymentMethodTypeKonbini          PaymentLinkPaymentMethodType = "konbini"
	PaymentLinkPaymentMethodTypeLink             PaymentLinkPaymentMethodType = "link"
	PaymentLinkPaymentMethodTypeMbWay            PaymentLinkPaymentMethodType = "mb_way"
	PaymentLinkPaymentMethodTypeMobilepay        PaymentLinkPaymentMethodType = "mobilepay"
	PaymentLinkPaymentMethodTypeMultibanco       PaymentLinkPaymentMethodType = "multibanco"
	PaymentLinkPaymentMethodTypeOXXO             PaymentLinkPaymentMethodType = "oxxo"
	PaymentLinkPaymentMethodTypeP24              PaymentLinkPaymentMethodType = "p24"
	PaymentLinkPaymentMethodTypePayByBank        PaymentLinkPaymentMethodType = "pay_by_bank"
	PaymentLinkPaymentMethodTypePayNow           PaymentLinkPaymentMethodType = "paynow"
	PaymentLinkPaymentMethodTypePaypal           PaymentLinkPaymentMethodType = "paypal"
	PaymentLinkPaymentMethodTypePaypay           PaymentLinkPaymentMethodType = "paypay"
	PaymentLinkPaymentMethodTypePayto            PaymentLinkPaymentMethodType = "payto"
	PaymentLinkPaymentMethodTypePix              PaymentLinkPaymentMethodType = "pix"
	PaymentLinkPaymentMethodTypePromptPay        PaymentLinkPaymentMethodType = "promptpay"
	PaymentLinkPaymentMethodTypeQris             PaymentLinkPaymentMethodType = "qris"
	PaymentLinkPaymentMethodTypeRechnung         PaymentLinkPaymentMethodType = "rechnung"
	PaymentLinkPaymentMethodTypeSatispay         PaymentLinkPaymentMethodType = "satispay"
	PaymentLinkPaymentMethodTypeSEPADebit        PaymentLinkPaymentMethodType = "sepa_debit"
	PaymentLinkPaymentMethodTypeShopeepay        PaymentLinkPaymentMethodType = "shopeepay"
	PaymentLinkPaymentMethodTypeSofort           PaymentLinkPaymentMethodType = "sofort"
	PaymentLinkPaymentMethodTypeSwish            PaymentLinkPaymentMethodType = "swish"
	PaymentLinkPaymentMethodTypeTWINT            PaymentLinkPaymentMethodType = "twint"
	PaymentLinkPaymentMethodTypeUSBankAccount    PaymentLinkPaymentMethodType = "us_bank_account"
	PaymentLinkPaymentMethodTypeWeChatPay        PaymentLinkPaymentMethodType = "wechat_pay"
	PaymentLinkPaymentMethodTypeZip              PaymentLinkPaymentMethodType = "zip"
)

// Indicates the type of transaction being performed which customizes relevant text on the page, such as the submit button.
type PaymentLinkSubmitType string

// List of values that PaymentLinkSubmitType can take
const (
	PaymentLinkSubmitTypeAuto      PaymentLinkSubmitType = "auto"
	PaymentLinkSubmitTypeBook      PaymentLinkSubmitType = "book"
	PaymentLinkSubmitTypeDonate    PaymentLinkSubmitType = "donate"
	PaymentLinkSubmitTypePay       PaymentLinkSubmitType = "pay"
	PaymentLinkSubmitTypeSubscribe PaymentLinkSubmitType = "subscribe"
)

// Type of the account referenced.
type PaymentLinkSubscriptionDataInvoiceSettingsIssuerType string

// List of values that PaymentLinkSubscriptionDataInvoiceSettingsIssuerType can take
const (
	PaymentLinkSubscriptionDataInvoiceSettingsIssuerTypeAccount PaymentLinkSubscriptionDataInvoiceSettingsIssuerType = "account"
	PaymentLinkSubscriptionDataInvoiceSettingsIssuerTypeSelf    PaymentLinkSubscriptionDataInvoiceSettingsIssuerType = "self"
)

// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
type PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod string

// List of values that PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod can take
const (
	PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethodCancel        PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod = "cancel"
	PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethodCreateInvoice PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod = "create_invoice"
	PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethodPause         PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod = "pause"
)

type PaymentLinkTaxIDCollectionRequired string

// List of values that PaymentLinkTaxIDCollectionRequired can take
const (
	PaymentLinkTaxIDCollectionRequiredIfSupported PaymentLinkTaxIDCollectionRequired = "if_supported"
	PaymentLinkTaxIDCollectionRequiredNever       PaymentLinkTaxIDCollectionRequired = "never"
)

// Returns a list of your payment links.
type PaymentLinkListParams struct {
	ListParams `form:"*"`
	// Only return payment links that are active or inactive (e.g., pass `false` to list all inactive payment links).
	Active *bool `form:"active" json:"active,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration when `type=hosted_confirmation`.
type PaymentLinkAfterCompletionHostedConfirmationParams struct {
	// A custom message to display to the customer after the purchase is complete.
	CustomMessage *string `form:"custom_message" json:"custom_message,omitempty"`
}

// Configuration when `type=redirect`.
type PaymentLinkAfterCompletionRedirectParams struct {
	// The URL the customer will be redirected to after the purchase is complete. You can embed `{CHECKOUT_SESSION_ID}` into the URL to have the `id` of the completed [checkout session](https://docs.stripe.com/api/checkout/sessions/object#checkout_session_object-id) included.
	URL *string `form:"url" json:"url"`
}

// Behavior after the purchase is complete.
type PaymentLinkAfterCompletionParams struct {
	// Configuration when `type=hosted_confirmation`.
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmationParams `form:"hosted_confirmation" json:"hosted_confirmation,omitempty"`
	// Configuration when `type=redirect`.
	Redirect *PaymentLinkAfterCompletionRedirectParams `form:"redirect" json:"redirect,omitempty"`
	// The specified behavior after the purchase is complete. Either `redirect` or `hosted_confirmation`.
	Type *string `form:"type" json:"type"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type PaymentLinkAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Configuration for automatic tax collection.
type PaymentLinkAutomaticTaxParams struct {
	// Set to `true` to [calculate tax automatically](https://docs.stripe.com/tax) using the customer's location.
	//
	// Enabling this parameter causes the payment link to collect any billing address information necessary for tax calculation.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *PaymentLinkAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
type PaymentLinkConsentCollectionPaymentMethodReuseAgreementParams struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's
	// defaults will be used. When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position *string `form:"position" json:"position"`
}

// Configure fields to gather active consent from customers.
type PaymentLinkConsentCollectionParams struct {
	// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
	PaymentMethodReuseAgreement *PaymentLinkConsentCollectionPaymentMethodReuseAgreementParams `form:"payment_method_reuse_agreement" json:"payment_method_reuse_agreement,omitempty"`
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions" json:"promotions,omitempty"`
	// If set to `required`, it requires customers to check a terms of service checkbox before being able to pay.
	// There must be a valid terms of service URL set in your [Dashboard settings](https://dashboard.stripe.com/settings/public).
	TermsOfService *string `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The options available for the customer to select. Up to 200 options allowed.
type PaymentLinkCustomFieldDropdownOptionParams struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label *string `form:"label" json:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value *string `form:"value" json:"value"`
}

// Configuration for `type=dropdown` fields.
type PaymentLinkCustomFieldDropdownParams struct {
	// The value that pre-fills the field on the payment page.Must match a `value` in the `options` array.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*PaymentLinkCustomFieldDropdownOptionParams `form:"options" json:"options"`
}

// The label for the field, displayed to the customer.
type PaymentLinkCustomFieldLabelParams struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom *string `form:"custom" json:"custom"`
	// The type of the label.
	Type *string `form:"type" json:"type"`
}

// Configuration for `type=numeric` fields.
type PaymentLinkCustomFieldNumericParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Configuration for `type=text` fields.
type PaymentLinkCustomFieldTextParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkCustomFieldParams struct {
	// Configuration for `type=dropdown` fields.
	Dropdown *PaymentLinkCustomFieldDropdownParams `form:"dropdown" json:"dropdown,omitempty"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key *string `form:"key" json:"key"`
	// The label for the field, displayed to the customer.
	Label *PaymentLinkCustomFieldLabelParams `form:"label" json:"label"`
	// Configuration for `type=numeric` fields.
	Numeric *PaymentLinkCustomFieldNumericParams `form:"numeric" json:"numeric,omitempty"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
	// Configuration for `type=text` fields.
	Text *PaymentLinkCustomFieldTextParams `form:"text" json:"text,omitempty"`
	// The type of the field.
	Type *string `form:"type" json:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type PaymentLinkCustomTextAfterSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type PaymentLinkCustomTextShippingAddressParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type PaymentLinkCustomTextSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type PaymentLinkCustomTextTermsOfServiceAcceptanceParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkCustomTextParams struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *PaymentLinkCustomTextAfterSubmitParams `form:"after_submit" json:"after_submit,omitempty"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *PaymentLinkCustomTextShippingAddressParams `form:"shipping_address" json:"shipping_address,omitempty"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *PaymentLinkCustomTextSubmitParams `form:"submit" json:"submit,omitempty"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *PaymentLinkCustomTextTermsOfServiceAcceptanceParams `form:"terms_of_service_acceptance" json:"terms_of_service_acceptance,omitempty"`
	UnsetFields              []PaymentLinkCustomTextParamsUnsetField              `form:"-" json:"-"`
}

// PaymentLinkCustomTextParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkCustomTextParams.
type PaymentLinkCustomTextParamsUnsetField string

const (
	PaymentLinkCustomTextParamsUnsetFieldAfterSubmit              PaymentLinkCustomTextParamsUnsetField = "after_submit"
	PaymentLinkCustomTextParamsUnsetFieldShippingAddress          PaymentLinkCustomTextParamsUnsetField = "shipping_address"
	PaymentLinkCustomTextParamsUnsetFieldSubmit                   PaymentLinkCustomTextParamsUnsetField = "submit"
	PaymentLinkCustomTextParamsUnsetFieldTermsOfServiceAcceptance PaymentLinkCustomTextParamsUnsetField = "terms_of_service_acceptance"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkCustomTextParams) AddUnsetField(field PaymentLinkCustomTextParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Default custom fields to be displayed on invoices for this customer.
type PaymentLinkInvoiceCreationInvoiceDataCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters.
	Value *string `form:"value" json:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkInvoiceCreationInvoiceDataIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Default options for invoice PDF rendering for this customer.
type PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for this invoice.
	Template    *string                                                                 `form:"template" json:"template,omitempty"`
	UnsetFields []PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField `form:"-" json:"-"`
}

// PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParams.
type PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField string

const (
	PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetFieldAmountTaxDisplay PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField = "amount_tax_display"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParams) AddUnsetField(field PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Invoice PDF configuration.
type PaymentLinkInvoiceCreationInvoiceDataParams struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*PaymentLinkInvoiceCreationInvoiceDataCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkInvoiceCreationInvoiceDataIssuerParams `form:"issuer" json:"issuer,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *PaymentLinkInvoiceCreationInvoiceDataRenderingOptionsParams `form:"rendering_options" json:"rendering_options,omitempty"`
	UnsetFields      []PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField      `form:"-" json:"-"`
}

// PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkInvoiceCreationInvoiceDataParams.
type PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField string

const (
	PaymentLinkInvoiceCreationInvoiceDataParamsUnsetFieldAccountTaxIDs    PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField = "account_tax_ids"
	PaymentLinkInvoiceCreationInvoiceDataParamsUnsetFieldCustomFields     PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField = "custom_fields"
	PaymentLinkInvoiceCreationInvoiceDataParamsUnsetFieldMetadata         PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField = "metadata"
	PaymentLinkInvoiceCreationInvoiceDataParamsUnsetFieldRenderingOptions PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField = "rendering_options"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkInvoiceCreationInvoiceDataParams) AddUnsetField(field PaymentLinkInvoiceCreationInvoiceDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkInvoiceCreationInvoiceDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Generate a post-purchase Invoice for one-time payments.
type PaymentLinkInvoiceCreationParams struct {
	// Whether the feature is enabled
	Enabled *bool `form:"enabled" json:"enabled"`
	// Invoice PDF configuration.
	InvoiceData *PaymentLinkInvoiceCreationInvoiceDataParams `form:"invoice_data" json:"invoice_data,omitempty"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
type PaymentLinkLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative Integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity the customer can purchase. By default this value is 0. If there is only one item in the cart then that item's quantity cannot go down to 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes) and an optional performance location.
type PaymentLinkLineItemPriceDataProductDataTaxDetailsParams struct {
	// A tax location ID. Depending on the [tax code](https://docs.stripe.com/tax/tax-for-tickets/reference/tax-location-performance), this is required, optional, or not supported.
	PerformanceLocation *string `form:"performance_location" json:"performance_location,omitempty"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
type PaymentLinkLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images" json:"images,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name" json:"name"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
	// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes) and an optional performance location.
	TaxDetails *PaymentLinkLineItemPriceDataProductDataTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkLineItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The recurring components of a price such as `interval` and `interval_count`.
type PaymentLinkLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type PaymentLinkLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product" json:"product,omitempty"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
	ProductData *PaymentLinkLineItemPriceDataProductDataParams `form:"product_data" json:"product_data,omitempty"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *PaymentLinkLineItemPriceDataRecurringParams `form:"recurring" json:"recurring,omitempty"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount" json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision" json:"unit_amount_decimal,string,omitempty"`
}

// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
type PaymentLinkLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
	AdjustableQuantity *PaymentLinkLineItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of an existing line item on the payment link.
	ID *string `form:"id" json:"id,omitempty"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *PaymentLinkLineItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// The quantity of the line item being purchased.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Controls settings applied for collecting the customer's business name.
type PaymentLinkNameCollectionBusinessParams struct {
	// Enable business name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their business name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's individual name.
type PaymentLinkNameCollectionIndividualParams struct {
	// Enable individual name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their full name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's name.
type PaymentLinkNameCollectionParams struct {
	// Controls settings applied for collecting the customer's business name.
	Business *PaymentLinkNameCollectionBusinessParams `form:"business" json:"business,omitempty"`
	// Controls settings applied for collecting the customer's individual name.
	Individual *PaymentLinkNameCollectionIndividualParams `form:"individual" json:"individual,omitempty"`
}

// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
type PaymentLinkOptionalItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
// There is a maximum of 20 combined line items and optional items.
type PaymentLinkOptionalItemParams struct {
	// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
	AdjustableQuantity *PaymentLinkOptionalItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object.
	Price *string `form:"price" json:"price"`
	// The initial quantity of the line item created when a customer chooses to add this optional item to their order.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type PaymentLinkPaymentIntentDataParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method" json:"capture_method,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Indicates that you intend to [make future payments](https://docs.stripe.com/payments/payment-intents#future-usage) with the payment method collected by this Checkout Session.
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
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// Setting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup *string                                        `form:"transfer_group" json:"transfer_group,omitempty"`
	UnsetFields   []PaymentLinkPaymentIntentDataParamsUnsetField `form:"-" json:"-"`
}

// PaymentLinkPaymentIntentDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkPaymentIntentDataParams.
type PaymentLinkPaymentIntentDataParamsUnsetField string

const (
	PaymentLinkPaymentIntentDataParamsUnsetFieldDescription               PaymentLinkPaymentIntentDataParamsUnsetField = "description"
	PaymentLinkPaymentIntentDataParamsUnsetFieldMetadata                  PaymentLinkPaymentIntentDataParamsUnsetField = "metadata"
	PaymentLinkPaymentIntentDataParamsUnsetFieldStatementDescriptor       PaymentLinkPaymentIntentDataParamsUnsetField = "statement_descriptor"
	PaymentLinkPaymentIntentDataParamsUnsetFieldStatementDescriptorSuffix PaymentLinkPaymentIntentDataParamsUnsetField = "statement_descriptor_suffix"
	PaymentLinkPaymentIntentDataParamsUnsetFieldTransferGroup             PaymentLinkPaymentIntentDataParamsUnsetField = "transfer_group"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkPaymentIntentDataParams) AddUnsetField(field PaymentLinkPaymentIntentDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkPaymentIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls phone number collection settings during checkout.
//
// We recommend that you review your privacy policy and check with your legal contacts.
type PaymentLinkPhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Configuration for the `completed_sessions` restriction type.
type PaymentLinkRestrictionsCompletedSessionsParams struct {
	// The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met.
	Limit *int64 `form:"limit" json:"limit"`
}

// Settings that restrict the usage of a payment link.
type PaymentLinkRestrictionsParams struct {
	// Configuration for the `completed_sessions` restriction type.
	CompletedSessions *PaymentLinkRestrictionsCompletedSessionsParams `form:"completed_sessions" json:"completed_sessions"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations.
	AllowedCountries []*string `form:"allowed_countries" json:"allowed_countries"`
}

// The shipping rate options to apply to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
type PaymentLinkShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate" json:"shipping_rate,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkSubscriptionDataInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type PaymentLinkSubscriptionDataInvoiceSettingsParams struct {
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkSubscriptionDataInvoiceSettingsIssuerParams `form:"issuer" json:"issuer,omitempty"`
}

// Defines how the subscription should behave when the user's free trial ends.
type PaymentLinkSubscriptionDataTrialSettingsEndBehaviorParams struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod *string `form:"missing_payment_method" json:"missing_payment_method"`
}

// Settings related to subscription trials.
type PaymentLinkSubscriptionDataTrialSettingsParams struct {
	// Defines how the subscription should behave when the user's free trial ends.
	EndBehavior *PaymentLinkSubscriptionDataTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionDataParams struct {
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *PaymentLinkSubscriptionDataInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days" json:"trial_period_days,omitempty"`
	// Settings related to subscription trials.
	TrialSettings *PaymentLinkSubscriptionDataTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
	UnsetFields   []PaymentLinkSubscriptionDataParamsUnsetField   `form:"-" json:"-"`
}

// PaymentLinkSubscriptionDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkSubscriptionDataParams.
type PaymentLinkSubscriptionDataParamsUnsetField string

const (
	PaymentLinkSubscriptionDataParamsUnsetFieldMetadata        PaymentLinkSubscriptionDataParamsUnsetField = "metadata"
	PaymentLinkSubscriptionDataParamsUnsetFieldTrialPeriodDays PaymentLinkSubscriptionDataParamsUnsetField = "trial_period_days"
	PaymentLinkSubscriptionDataParamsUnsetFieldTrialSettings   PaymentLinkSubscriptionDataParamsUnsetField = "trial_settings"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkSubscriptionDataParams) AddUnsetField(field PaymentLinkSubscriptionDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkSubscriptionDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls tax ID collection during checkout.
type PaymentLinkTaxIDCollectionParams struct {
	// Enable tax ID collection during checkout. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Describes whether a tax ID is required during checkout. Defaults to `never`. You can't set this parameter if `ui_mode` is `custom`.
	Required *string `form:"required" json:"required,omitempty"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination" json:"destination"`
}

// Creates a payment link.
type PaymentLinkParams struct {
	Params `form:"*"`
	// Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.
	Active *bool `form:"active" json:"active,omitempty"`
	// Behavior after the purchase is complete.
	AfterCompletion *PaymentLinkAfterCompletionParams `form:"after_completion" json:"after_completion,omitempty"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes" json:"allow_promotion_codes,omitempty"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Can only be applied when there are no line items with recurring prices.
	ApplicationFeeAmount *int64 `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. There must be at least 1 line item with a recurring price to use this field.
	ApplicationFeePercent *float64 `form:"application_fee_percent" json:"application_fee_percent,omitempty"`
	// Configuration for automatic tax collection.
	AutomaticTax *PaymentLinkAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Configuration for collecting the customer's billing address. Defaults to `auto`.
	BillingAddressCollection *string `form:"billing_address_collection" json:"billing_address_collection,omitempty"`
	// Configure fields to gather active consent from customers.
	ConsentCollection *PaymentLinkConsentCollectionParams `form:"consent_collection" json:"consent_collection,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies) and supported by each line item's price.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Configures whether [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link create a [Customer](https://docs.stripe.com/api/customers).
	CustomerCreation *string `form:"customer_creation" json:"customer_creation,omitempty"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*PaymentLinkCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
	CustomText *PaymentLinkCustomTextParams `form:"custom_text" json:"custom_text,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The custom message to be displayed to a customer when a payment link is no longer active.
	InactiveMessage *string `form:"inactive_message" json:"inactive_message,omitempty"`
	// Generate a post-purchase Invoice for one-time payments.
	InvoiceCreation *PaymentLinkInvoiceCreationParams `form:"invoice_creation" json:"invoice_creation,omitempty"`
	// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
	LineItems []*PaymentLinkLineItemParams `form:"line_items" json:"line_items,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`. Metadata associated with this Payment Link will automatically be copied to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Controls settings applied for collecting the customer's name.
	NameCollection *PaymentLinkNameCollectionParams `form:"name_collection" json:"name_collection,omitempty"`
	// The account on behalf of which to charge.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
	// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
	// There is a maximum of 20 combined line items and optional items.
	OptionalItems []*PaymentLinkOptionalItemParams `form:"optional_items" json:"optional_items,omitempty"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *PaymentLinkPaymentIntentDataParams `form:"payment_intent_data" json:"payment_intent_data,omitempty"`
	// Specify whether Checkout should collect a payment method. When set to `if_required`, Checkout will not collect a payment method when the total due for the session is 0.This may occur if the Checkout Session includes a free trial or a discount.
	//
	// Can only be set in `subscription` mode. Defaults to `always`.
	//
	// If you'd like information on how to collect a payment method outside of Checkout, read the guide on [configuring subscriptions with a free trial](https://docs.stripe.com/payments/checkout/free-trials).
	PaymentMethodCollection *string `form:"payment_method_collection" json:"payment_method_collection,omitempty"`
	// The list of payment method types that customers can use. If no value is passed, Stripe will dynamically show relevant payment methods from your [payment method settings](https://dashboard.stripe.com/settings/payment_methods) (20+ payment methods [supported](https://docs.stripe.com/payments/payment-methods/integration-options#payment-method-product-support)).
	PaymentMethodTypes []*string `form:"payment_method_types" json:"payment_method_types,omitempty"`
	// Controls phone number collection settings during checkout.
	//
	// We recommend that you review your privacy policy and check with your legal contacts.
	PhoneNumberCollection *PaymentLinkPhoneNumberCollectionParams `form:"phone_number_collection" json:"phone_number_collection,omitempty"`
	// Settings that restrict the usage of a payment link.
	Restrictions *PaymentLinkRestrictionsParams `form:"restrictions" json:"restrictions,omitempty"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkShippingAddressCollectionParams `form:"shipping_address_collection" json:"shipping_address_collection,omitempty"`
	// The shipping rate options to apply to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
	ShippingOptions []*PaymentLinkShippingOptionParams `form:"shipping_options" json:"shipping_options,omitempty"`
	// Describes the type of transaction being performed in order to customize relevant text on the page, such as the submit button. Changing this value will also affect the hostname in the [url](https://docs.stripe.com/api/payment_links/payment_links/object#url) property (example: `donate.stripe.com`).
	SubmitType *string `form:"submit_type" json:"submit_type,omitempty"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkSubscriptionDataParams `form:"subscription_data" json:"subscription_data,omitempty"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *PaymentLinkTaxIDCollectionParams `form:"tax_id_collection" json:"tax_id_collection,omitempty"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
	UnsetFields  []PaymentLinkParamsUnsetField  `form:"-" json:"-"`
}

// PaymentLinkParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkParams.
type PaymentLinkParamsUnsetField string

const (
	PaymentLinkParamsUnsetFieldCustomFields              PaymentLinkParamsUnsetField = "custom_fields"
	PaymentLinkParamsUnsetFieldInactiveMessage           PaymentLinkParamsUnsetField = "inactive_message"
	PaymentLinkParamsUnsetFieldNameCollection            PaymentLinkParamsUnsetField = "name_collection"
	PaymentLinkParamsUnsetFieldOptionalItems             PaymentLinkParamsUnsetField = "optional_items"
	PaymentLinkParamsUnsetFieldPaymentMethodTypes        PaymentLinkParamsUnsetField = "payment_method_types"
	PaymentLinkParamsUnsetFieldRestrictions              PaymentLinkParamsUnsetField = "restrictions"
	PaymentLinkParamsUnsetFieldShippingAddressCollection PaymentLinkParamsUnsetField = "shipping_address_collection"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkParams) AddUnsetField(field PaymentLinkParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type PaymentLinkListLineItemsParams struct {
	ListParams  `form:"*"`
	PaymentLink *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkListLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration when `type=hosted_confirmation`.
type PaymentLinkCreateAfterCompletionHostedConfirmationParams struct {
	// A custom message to display to the customer after the purchase is complete.
	CustomMessage *string `form:"custom_message" json:"custom_message,omitempty"`
}

// Configuration when `type=redirect`.
type PaymentLinkCreateAfterCompletionRedirectParams struct {
	// The URL the customer will be redirected to after the purchase is complete. You can embed `{CHECKOUT_SESSION_ID}` into the URL to have the `id` of the completed [checkout session](https://docs.stripe.com/api/checkout/sessions/object#checkout_session_object-id) included.
	URL *string `form:"url" json:"url"`
}

// Behavior after the purchase is complete.
type PaymentLinkCreateAfterCompletionParams struct {
	// Configuration when `type=hosted_confirmation`.
	HostedConfirmation *PaymentLinkCreateAfterCompletionHostedConfirmationParams `form:"hosted_confirmation" json:"hosted_confirmation,omitempty"`
	// Configuration when `type=redirect`.
	Redirect *PaymentLinkCreateAfterCompletionRedirectParams `form:"redirect" json:"redirect,omitempty"`
	// The specified behavior after the purchase is complete. Either `redirect` or `hosted_confirmation`.
	Type *string `form:"type" json:"type"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type PaymentLinkCreateAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Configuration for automatic tax collection.
type PaymentLinkCreateAutomaticTaxParams struct {
	// Set to `true` to [calculate tax automatically](https://docs.stripe.com/tax) using the customer's location.
	//
	// Enabling this parameter causes the payment link to collect any billing address information necessary for tax calculation.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *PaymentLinkCreateAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
type PaymentLinkCreateConsentCollectionPaymentMethodReuseAgreementParams struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's
	// defaults will be used. When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position *string `form:"position" json:"position"`
}

// Configure fields to gather active consent from customers.
type PaymentLinkCreateConsentCollectionParams struct {
	// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
	PaymentMethodReuseAgreement *PaymentLinkCreateConsentCollectionPaymentMethodReuseAgreementParams `form:"payment_method_reuse_agreement" json:"payment_method_reuse_agreement,omitempty"`
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions" json:"promotions,omitempty"`
	// If set to `required`, it requires customers to check a terms of service checkbox before being able to pay.
	// There must be a valid terms of service URL set in your [Dashboard settings](https://dashboard.stripe.com/settings/public).
	TermsOfService *string `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The options available for the customer to select. Up to 200 options allowed.
type PaymentLinkCreateCustomFieldDropdownOptionParams struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label *string `form:"label" json:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value *string `form:"value" json:"value"`
}

// Configuration for `type=dropdown` fields.
type PaymentLinkCreateCustomFieldDropdownParams struct {
	// The value that pre-fills the field on the payment page.Must match a `value` in the `options` array.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*PaymentLinkCreateCustomFieldDropdownOptionParams `form:"options" json:"options"`
}

// The label for the field, displayed to the customer.
type PaymentLinkCreateCustomFieldLabelParams struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom *string `form:"custom" json:"custom"`
	// The type of the label.
	Type *string `form:"type" json:"type"`
}

// Configuration for `type=numeric` fields.
type PaymentLinkCreateCustomFieldNumericParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Configuration for `type=text` fields.
type PaymentLinkCreateCustomFieldTextParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkCreateCustomFieldParams struct {
	// Configuration for `type=dropdown` fields.
	Dropdown *PaymentLinkCreateCustomFieldDropdownParams `form:"dropdown" json:"dropdown,omitempty"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key *string `form:"key" json:"key"`
	// The label for the field, displayed to the customer.
	Label *PaymentLinkCreateCustomFieldLabelParams `form:"label" json:"label"`
	// Configuration for `type=numeric` fields.
	Numeric *PaymentLinkCreateCustomFieldNumericParams `form:"numeric" json:"numeric,omitempty"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
	// Configuration for `type=text` fields.
	Text *PaymentLinkCreateCustomFieldTextParams `form:"text" json:"text,omitempty"`
	// The type of the field.
	Type *string `form:"type" json:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type PaymentLinkCreateCustomTextAfterSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type PaymentLinkCreateCustomTextShippingAddressParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type PaymentLinkCreateCustomTextSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type PaymentLinkCreateCustomTextTermsOfServiceAcceptanceParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkCreateCustomTextParams struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *PaymentLinkCreateCustomTextAfterSubmitParams `form:"after_submit" json:"after_submit,omitempty"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *PaymentLinkCreateCustomTextShippingAddressParams `form:"shipping_address" json:"shipping_address,omitempty"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *PaymentLinkCreateCustomTextSubmitParams `form:"submit" json:"submit,omitempty"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *PaymentLinkCreateCustomTextTermsOfServiceAcceptanceParams `form:"terms_of_service_acceptance" json:"terms_of_service_acceptance,omitempty"`
	UnsetFields              []PaymentLinkCreateCustomTextParamsUnsetField              `form:"-" json:"-"`
}

// PaymentLinkCreateCustomTextParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkCreateCustomTextParams.
type PaymentLinkCreateCustomTextParamsUnsetField string

const (
	PaymentLinkCreateCustomTextParamsUnsetFieldAfterSubmit              PaymentLinkCreateCustomTextParamsUnsetField = "after_submit"
	PaymentLinkCreateCustomTextParamsUnsetFieldShippingAddress          PaymentLinkCreateCustomTextParamsUnsetField = "shipping_address"
	PaymentLinkCreateCustomTextParamsUnsetFieldSubmit                   PaymentLinkCreateCustomTextParamsUnsetField = "submit"
	PaymentLinkCreateCustomTextParamsUnsetFieldTermsOfServiceAcceptance PaymentLinkCreateCustomTextParamsUnsetField = "terms_of_service_acceptance"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkCreateCustomTextParams) AddUnsetField(field PaymentLinkCreateCustomTextParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Default custom fields to be displayed on invoices for this customer.
type PaymentLinkCreateInvoiceCreationInvoiceDataCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters.
	Value *string `form:"value" json:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkCreateInvoiceCreationInvoiceDataIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Default options for invoice PDF rendering for this customer.
type PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for this invoice.
	Template    *string                                                                       `form:"template" json:"template,omitempty"`
	UnsetFields []PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField `form:"-" json:"-"`
}

// PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParams.
type PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField string

const (
	PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetFieldAmountTaxDisplay PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField = "amount_tax_display"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParams) AddUnsetField(field PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Invoice PDF configuration.
type PaymentLinkCreateInvoiceCreationInvoiceDataParams struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*PaymentLinkCreateInvoiceCreationInvoiceDataCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkCreateInvoiceCreationInvoiceDataIssuerParams `form:"issuer" json:"issuer,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *PaymentLinkCreateInvoiceCreationInvoiceDataRenderingOptionsParams `form:"rendering_options" json:"rendering_options,omitempty"`
	UnsetFields      []PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField      `form:"-" json:"-"`
}

// PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkCreateInvoiceCreationInvoiceDataParams.
type PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField string

const (
	PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetFieldAccountTaxIDs    PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField = "account_tax_ids"
	PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetFieldCustomFields     PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField = "custom_fields"
	PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetFieldMetadata         PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField = "metadata"
	PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetFieldRenderingOptions PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField = "rendering_options"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkCreateInvoiceCreationInvoiceDataParams) AddUnsetField(field PaymentLinkCreateInvoiceCreationInvoiceDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkCreateInvoiceCreationInvoiceDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Generate a post-purchase Invoice for one-time payments.
type PaymentLinkCreateInvoiceCreationParams struct {
	// Whether the feature is enabled
	Enabled *bool `form:"enabled" json:"enabled"`
	// Invoice PDF configuration.
	InvoiceData *PaymentLinkCreateInvoiceCreationInvoiceDataParams `form:"invoice_data" json:"invoice_data,omitempty"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
type PaymentLinkCreateLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative Integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity the customer can purchase. By default this value is 0. If there is only one item in the cart then that item's quantity cannot go down to 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes) and an optional performance location.
type PaymentLinkCreateLineItemPriceDataProductDataTaxDetailsParams struct {
	// A tax location ID. Depending on the [tax code](https://docs.stripe.com/tax/tax-for-tickets/reference/tax-location-performance), this is required, optional, or not supported.
	PerformanceLocation *string `form:"performance_location" json:"performance_location,omitempty"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
type PaymentLinkCreateLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images" json:"images,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name" json:"name"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
	// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes) and an optional performance location.
	TaxDetails *PaymentLinkCreateLineItemPriceDataProductDataTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkCreateLineItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The recurring components of a price such as `interval` and `interval_count`.
type PaymentLinkCreateLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type PaymentLinkCreateLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product" json:"product,omitempty"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
	ProductData *PaymentLinkCreateLineItemPriceDataProductDataParams `form:"product_data" json:"product_data,omitempty"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *PaymentLinkCreateLineItemPriceDataRecurringParams `form:"recurring" json:"recurring,omitempty"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount" json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision" json:"unit_amount_decimal,string,omitempty"`
}

// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
type PaymentLinkCreateLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
	AdjustableQuantity *PaymentLinkCreateLineItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *PaymentLinkCreateLineItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// The quantity of the line item being purchased.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// Controls settings applied for collecting the customer's business name.
type PaymentLinkCreateNameCollectionBusinessParams struct {
	// Enable business name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their business name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's individual name.
type PaymentLinkCreateNameCollectionIndividualParams struct {
	// Enable individual name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their full name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's name.
type PaymentLinkCreateNameCollectionParams struct {
	// Controls settings applied for collecting the customer's business name.
	Business *PaymentLinkCreateNameCollectionBusinessParams `form:"business" json:"business,omitempty"`
	// Controls settings applied for collecting the customer's individual name.
	Individual *PaymentLinkCreateNameCollectionIndividualParams `form:"individual" json:"individual,omitempty"`
}

// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
type PaymentLinkCreateOptionalItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
// There is a maximum of 20 combined line items and optional items.
type PaymentLinkCreateOptionalItemParams struct {
	// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
	AdjustableQuantity *PaymentLinkCreateOptionalItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object.
	Price *string `form:"price" json:"price"`
	// The initial quantity of the line item created when a customer chooses to add this optional item to their order.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type PaymentLinkCreatePaymentIntentDataParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method" json:"capture_method,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Indicates that you intend to [make future payments](https://docs.stripe.com/payments/payment-intents#future-usage) with the payment method collected by this Checkout Session.
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
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// Setting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup *string `form:"transfer_group" json:"transfer_group,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkCreatePaymentIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls phone number collection settings during checkout.
//
// We recommend that you review your privacy policy and check with your legal contacts.
type PaymentLinkCreatePhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Configuration for the `completed_sessions` restriction type.
type PaymentLinkCreateRestrictionsCompletedSessionsParams struct {
	// The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met.
	Limit *int64 `form:"limit" json:"limit"`
}

// Settings that restrict the usage of a payment link.
type PaymentLinkCreateRestrictionsParams struct {
	// Configuration for the `completed_sessions` restriction type.
	CompletedSessions *PaymentLinkCreateRestrictionsCompletedSessionsParams `form:"completed_sessions" json:"completed_sessions"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkCreateShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations.
	AllowedCountries []*string `form:"allowed_countries" json:"allowed_countries"`
}

// The shipping rate options to apply to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
type PaymentLinkCreateShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate" json:"shipping_rate,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkCreateSubscriptionDataInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type PaymentLinkCreateSubscriptionDataInvoiceSettingsParams struct {
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkCreateSubscriptionDataInvoiceSettingsIssuerParams `form:"issuer" json:"issuer,omitempty"`
}

// Defines how the subscription should behave when the user's free trial ends.
type PaymentLinkCreateSubscriptionDataTrialSettingsEndBehaviorParams struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod *string `form:"missing_payment_method" json:"missing_payment_method"`
}

// Settings related to subscription trials.
type PaymentLinkCreateSubscriptionDataTrialSettingsParams struct {
	// Defines how the subscription should behave when the user's free trial ends.
	EndBehavior *PaymentLinkCreateSubscriptionDataTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkCreateSubscriptionDataParams struct {
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *PaymentLinkCreateSubscriptionDataInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days" json:"trial_period_days,omitempty"`
	// Settings related to subscription trials.
	TrialSettings *PaymentLinkCreateSubscriptionDataTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkCreateSubscriptionDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls tax ID collection during checkout.
type PaymentLinkCreateTaxIDCollectionParams struct {
	// Enable tax ID collection during checkout. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Describes whether a tax ID is required during checkout. Defaults to `never`. You can't set this parameter if `ui_mode` is `custom`.
	Required *string `form:"required" json:"required,omitempty"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkCreateTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination" json:"destination"`
}

// Creates a payment link.
type PaymentLinkCreateParams struct {
	Params `form:"*"`
	// Behavior after the purchase is complete.
	AfterCompletion *PaymentLinkCreateAfterCompletionParams `form:"after_completion" json:"after_completion,omitempty"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes" json:"allow_promotion_codes,omitempty"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Can only be applied when there are no line items with recurring prices.
	ApplicationFeeAmount *int64 `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. There must be at least 1 line item with a recurring price to use this field.
	ApplicationFeePercent *float64 `form:"application_fee_percent" json:"application_fee_percent,omitempty"`
	// Configuration for automatic tax collection.
	AutomaticTax *PaymentLinkCreateAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Configuration for collecting the customer's billing address. Defaults to `auto`.
	BillingAddressCollection *string `form:"billing_address_collection" json:"billing_address_collection,omitempty"`
	// Configure fields to gather active consent from customers.
	ConsentCollection *PaymentLinkCreateConsentCollectionParams `form:"consent_collection" json:"consent_collection,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies) and supported by each line item's price.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Configures whether [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link create a [Customer](https://docs.stripe.com/api/customers).
	CustomerCreation *string `form:"customer_creation" json:"customer_creation,omitempty"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*PaymentLinkCreateCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
	CustomText *PaymentLinkCreateCustomTextParams `form:"custom_text" json:"custom_text,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The custom message to be displayed to a customer when a payment link is no longer active.
	InactiveMessage *string `form:"inactive_message" json:"inactive_message,omitempty"`
	// Generate a post-purchase Invoice for one-time payments.
	InvoiceCreation *PaymentLinkCreateInvoiceCreationParams `form:"invoice_creation" json:"invoice_creation,omitempty"`
	// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
	LineItems []*PaymentLinkCreateLineItemParams `form:"line_items" json:"line_items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`. Metadata associated with this Payment Link will automatically be copied to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Controls settings applied for collecting the customer's name.
	NameCollection *PaymentLinkCreateNameCollectionParams `form:"name_collection" json:"name_collection,omitempty"`
	// The account on behalf of which to charge.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
	// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
	// There is a maximum of 20 combined line items and optional items.
	OptionalItems []*PaymentLinkCreateOptionalItemParams `form:"optional_items" json:"optional_items,omitempty"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *PaymentLinkCreatePaymentIntentDataParams `form:"payment_intent_data" json:"payment_intent_data,omitempty"`
	// Specify whether Checkout should collect a payment method. When set to `if_required`, Checkout will not collect a payment method when the total due for the session is 0.This may occur if the Checkout Session includes a free trial or a discount.
	//
	// Can only be set in `subscription` mode. Defaults to `always`.
	//
	// If you'd like information on how to collect a payment method outside of Checkout, read the guide on [configuring subscriptions with a free trial](https://docs.stripe.com/payments/checkout/free-trials).
	PaymentMethodCollection *string `form:"payment_method_collection" json:"payment_method_collection,omitempty"`
	// The list of payment method types that customers can use. If no value is passed, Stripe will dynamically show relevant payment methods from your [payment method settings](https://dashboard.stripe.com/settings/payment_methods) (20+ payment methods [supported](https://docs.stripe.com/payments/payment-methods/integration-options#payment-method-product-support)).
	PaymentMethodTypes []*string `form:"payment_method_types" json:"payment_method_types,omitempty"`
	// Controls phone number collection settings during checkout.
	//
	// We recommend that you review your privacy policy and check with your legal contacts.
	PhoneNumberCollection *PaymentLinkCreatePhoneNumberCollectionParams `form:"phone_number_collection" json:"phone_number_collection,omitempty"`
	// Settings that restrict the usage of a payment link.
	Restrictions *PaymentLinkCreateRestrictionsParams `form:"restrictions" json:"restrictions,omitempty"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkCreateShippingAddressCollectionParams `form:"shipping_address_collection" json:"shipping_address_collection,omitempty"`
	// The shipping rate options to apply to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
	ShippingOptions []*PaymentLinkCreateShippingOptionParams `form:"shipping_options" json:"shipping_options,omitempty"`
	// Describes the type of transaction being performed in order to customize relevant text on the page, such as the submit button. Changing this value will also affect the hostname in the [url](https://docs.stripe.com/api/payment_links/payment_links/object#url) property (example: `donate.stripe.com`).
	SubmitType *string `form:"submit_type" json:"submit_type,omitempty"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkCreateSubscriptionDataParams `form:"subscription_data" json:"subscription_data,omitempty"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *PaymentLinkCreateTaxIDCollectionParams `form:"tax_id_collection" json:"tax_id_collection,omitempty"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkCreateTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a payment link.
type PaymentLinkRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration when `type=hosted_confirmation`.
type PaymentLinkUpdateAfterCompletionHostedConfirmationParams struct {
	// A custom message to display to the customer after the purchase is complete.
	CustomMessage *string `form:"custom_message" json:"custom_message,omitempty"`
}

// Configuration when `type=redirect`.
type PaymentLinkUpdateAfterCompletionRedirectParams struct {
	// The URL the customer will be redirected to after the purchase is complete. You can embed `{CHECKOUT_SESSION_ID}` into the URL to have the `id` of the completed [checkout session](https://docs.stripe.com/api/checkout/sessions/object#checkout_session_object-id) included.
	URL *string `form:"url" json:"url"`
}

// Behavior after the purchase is complete.
type PaymentLinkUpdateAfterCompletionParams struct {
	// Configuration when `type=hosted_confirmation`.
	HostedConfirmation *PaymentLinkUpdateAfterCompletionHostedConfirmationParams `form:"hosted_confirmation" json:"hosted_confirmation,omitempty"`
	// Configuration when `type=redirect`.
	Redirect *PaymentLinkUpdateAfterCompletionRedirectParams `form:"redirect" json:"redirect,omitempty"`
	// The specified behavior after the purchase is complete. Either `redirect` or `hosted_confirmation`.
	Type *string `form:"type" json:"type"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type PaymentLinkUpdateAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Configuration for automatic tax collection.
type PaymentLinkUpdateAutomaticTaxParams struct {
	// Set to `true` to [calculate tax automatically](https://docs.stripe.com/tax) using the customer's location.
	//
	// Enabling this parameter causes the payment link to collect any billing address information necessary for tax calculation.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *PaymentLinkUpdateAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// The options available for the customer to select. Up to 200 options allowed.
type PaymentLinkUpdateCustomFieldDropdownOptionParams struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label *string `form:"label" json:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value *string `form:"value" json:"value"`
}

// Configuration for `type=dropdown` fields.
type PaymentLinkUpdateCustomFieldDropdownParams struct {
	// The value that pre-fills the field on the payment page.Must match a `value` in the `options` array.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*PaymentLinkUpdateCustomFieldDropdownOptionParams `form:"options" json:"options"`
}

// The label for the field, displayed to the customer.
type PaymentLinkUpdateCustomFieldLabelParams struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom *string `form:"custom" json:"custom"`
	// The type of the label.
	Type *string `form:"type" json:"type"`
}

// Configuration for `type=numeric` fields.
type PaymentLinkUpdateCustomFieldNumericParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Configuration for `type=text` fields.
type PaymentLinkUpdateCustomFieldTextParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value" json:"default_value,omitempty"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length" json:"maximum_length,omitempty"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length" json:"minimum_length,omitempty"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkUpdateCustomFieldParams struct {
	// Configuration for `type=dropdown` fields.
	Dropdown *PaymentLinkUpdateCustomFieldDropdownParams `form:"dropdown" json:"dropdown,omitempty"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key *string `form:"key" json:"key"`
	// The label for the field, displayed to the customer.
	Label *PaymentLinkUpdateCustomFieldLabelParams `form:"label" json:"label"`
	// Configuration for `type=numeric` fields.
	Numeric *PaymentLinkUpdateCustomFieldNumericParams `form:"numeric" json:"numeric,omitempty"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
	// Configuration for `type=text` fields.
	Text *PaymentLinkUpdateCustomFieldTextParams `form:"text" json:"text,omitempty"`
	// The type of the field.
	Type *string `form:"type" json:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type PaymentLinkUpdateCustomTextAfterSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type PaymentLinkUpdateCustomTextShippingAddressParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type PaymentLinkUpdateCustomTextSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type PaymentLinkUpdateCustomTextTermsOfServiceAcceptanceParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message" json:"message"`
}

// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkUpdateCustomTextParams struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *PaymentLinkUpdateCustomTextAfterSubmitParams `form:"after_submit" json:"after_submit,omitempty"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *PaymentLinkUpdateCustomTextShippingAddressParams `form:"shipping_address" json:"shipping_address,omitempty"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *PaymentLinkUpdateCustomTextSubmitParams `form:"submit" json:"submit,omitempty"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *PaymentLinkUpdateCustomTextTermsOfServiceAcceptanceParams `form:"terms_of_service_acceptance" json:"terms_of_service_acceptance,omitempty"`
	UnsetFields              []PaymentLinkUpdateCustomTextParamsUnsetField              `form:"-" json:"-"`
}

// PaymentLinkUpdateCustomTextParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdateCustomTextParams.
type PaymentLinkUpdateCustomTextParamsUnsetField string

const (
	PaymentLinkUpdateCustomTextParamsUnsetFieldAfterSubmit              PaymentLinkUpdateCustomTextParamsUnsetField = "after_submit"
	PaymentLinkUpdateCustomTextParamsUnsetFieldShippingAddress          PaymentLinkUpdateCustomTextParamsUnsetField = "shipping_address"
	PaymentLinkUpdateCustomTextParamsUnsetFieldSubmit                   PaymentLinkUpdateCustomTextParamsUnsetField = "submit"
	PaymentLinkUpdateCustomTextParamsUnsetFieldTermsOfServiceAcceptance PaymentLinkUpdateCustomTextParamsUnsetField = "terms_of_service_acceptance"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdateCustomTextParams) AddUnsetField(field PaymentLinkUpdateCustomTextParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Default custom fields to be displayed on invoices for this customer.
type PaymentLinkUpdateInvoiceCreationInvoiceDataCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters.
	Value *string `form:"value" json:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkUpdateInvoiceCreationInvoiceDataIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Default options for invoice PDF rendering for this customer.
type PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for this invoice.
	Template    *string                                                                       `form:"template" json:"template,omitempty"`
	UnsetFields []PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField `form:"-" json:"-"`
}

// PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParams.
type PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField string

const (
	PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetFieldAmountTaxDisplay PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField = "amount_tax_display"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParams) AddUnsetField(field PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Invoice PDF configuration.
type PaymentLinkUpdateInvoiceCreationInvoiceDataParams struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*PaymentLinkUpdateInvoiceCreationInvoiceDataCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkUpdateInvoiceCreationInvoiceDataIssuerParams `form:"issuer" json:"issuer,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *PaymentLinkUpdateInvoiceCreationInvoiceDataRenderingOptionsParams `form:"rendering_options" json:"rendering_options,omitempty"`
	UnsetFields      []PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField      `form:"-" json:"-"`
}

// PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdateInvoiceCreationInvoiceDataParams.
type PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField string

const (
	PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetFieldAccountTaxIDs    PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField = "account_tax_ids"
	PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetFieldCustomFields     PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField = "custom_fields"
	PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetFieldMetadata         PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField = "metadata"
	PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetFieldRenderingOptions PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField = "rendering_options"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdateInvoiceCreationInvoiceDataParams) AddUnsetField(field PaymentLinkUpdateInvoiceCreationInvoiceDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkUpdateInvoiceCreationInvoiceDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Generate a post-purchase Invoice for one-time payments.
type PaymentLinkUpdateInvoiceCreationParams struct {
	// Whether the feature is enabled
	Enabled *bool `form:"enabled" json:"enabled"`
	// Invoice PDF configuration.
	InvoiceData *PaymentLinkUpdateInvoiceCreationInvoiceDataParams `form:"invoice_data" json:"invoice_data,omitempty"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
type PaymentLinkUpdateLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative Integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity the customer can purchase. By default this value is 0. If there is only one item in the cart then that item's quantity cannot go down to 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
type PaymentLinkUpdateLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
	AdjustableQuantity *PaymentLinkUpdateLineItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of an existing line item on the payment link.
	ID *string `form:"id" json:"id"`
	// The quantity of the line item being purchased.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Controls settings applied for collecting the customer's business name.
type PaymentLinkUpdateNameCollectionBusinessParams struct {
	// Enable business name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their business name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's individual name.
type PaymentLinkUpdateNameCollectionIndividualParams struct {
	// Enable individual name collection on the payment link. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Whether the customer is required to provide their full name before checking out. Defaults to `false`.
	Optional *bool `form:"optional" json:"optional,omitempty"`
}

// Controls settings applied for collecting the customer's name.
type PaymentLinkUpdateNameCollectionParams struct {
	// Controls settings applied for collecting the customer's business name.
	Business *PaymentLinkUpdateNameCollectionBusinessParams `form:"business" json:"business,omitempty"`
	// Controls settings applied for collecting the customer's individual name.
	Individual *PaymentLinkUpdateNameCollectionIndividualParams `form:"individual" json:"individual,omitempty"`
}

// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
type PaymentLinkUpdateOptionalItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99.
	Maximum *int64 `form:"maximum" json:"maximum,omitempty"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum *int64 `form:"minimum" json:"minimum,omitempty"`
}

// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
// There is a maximum of 20 combined line items and optional items.
type PaymentLinkUpdateOptionalItemParams struct {
	// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
	AdjustableQuantity *PaymentLinkUpdateOptionalItemAdjustableQuantityParams `form:"adjustable_quantity" json:"adjustable_quantity,omitempty"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object.
	Price *string `form:"price" json:"price"`
	// The initial quantity of the line item created when a customer chooses to add this optional item to their order.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type PaymentLinkUpdatePaymentIntentDataParams struct {
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// Setting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup *string                                              `form:"transfer_group" json:"transfer_group,omitempty"`
	UnsetFields   []PaymentLinkUpdatePaymentIntentDataParamsUnsetField `form:"-" json:"-"`
}

// PaymentLinkUpdatePaymentIntentDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdatePaymentIntentDataParams.
type PaymentLinkUpdatePaymentIntentDataParamsUnsetField string

const (
	PaymentLinkUpdatePaymentIntentDataParamsUnsetFieldDescription               PaymentLinkUpdatePaymentIntentDataParamsUnsetField = "description"
	PaymentLinkUpdatePaymentIntentDataParamsUnsetFieldMetadata                  PaymentLinkUpdatePaymentIntentDataParamsUnsetField = "metadata"
	PaymentLinkUpdatePaymentIntentDataParamsUnsetFieldStatementDescriptor       PaymentLinkUpdatePaymentIntentDataParamsUnsetField = "statement_descriptor"
	PaymentLinkUpdatePaymentIntentDataParamsUnsetFieldStatementDescriptorSuffix PaymentLinkUpdatePaymentIntentDataParamsUnsetField = "statement_descriptor_suffix"
	PaymentLinkUpdatePaymentIntentDataParamsUnsetFieldTransferGroup             PaymentLinkUpdatePaymentIntentDataParamsUnsetField = "transfer_group"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdatePaymentIntentDataParams) AddUnsetField(field PaymentLinkUpdatePaymentIntentDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkUpdatePaymentIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls phone number collection settings during checkout.
//
// We recommend that you review your privacy policy and check with your legal contacts.
type PaymentLinkUpdatePhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Configuration for the `completed_sessions` restriction type.
type PaymentLinkUpdateRestrictionsCompletedSessionsParams struct {
	// The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met.
	Limit *int64 `form:"limit" json:"limit"`
}

// Settings that restrict the usage of a payment link.
type PaymentLinkUpdateRestrictionsParams struct {
	// Configuration for the `completed_sessions` restriction type.
	CompletedSessions *PaymentLinkUpdateRestrictionsCompletedSessionsParams `form:"completed_sessions" json:"completed_sessions"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkUpdateShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations.
	AllowedCountries []*string `form:"allowed_countries" json:"allowed_countries"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkUpdateSubscriptionDataInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type PaymentLinkUpdateSubscriptionDataInvoiceSettingsParams struct {
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkUpdateSubscriptionDataInvoiceSettingsIssuerParams `form:"issuer" json:"issuer,omitempty"`
}

// Defines how the subscription should behave when the user's free trial ends.
type PaymentLinkUpdateSubscriptionDataTrialSettingsEndBehaviorParams struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod *string `form:"missing_payment_method" json:"missing_payment_method"`
}

// Settings related to subscription trials.
type PaymentLinkUpdateSubscriptionDataTrialSettingsParams struct {
	// Defines how the subscription should behave when the user's free trial ends.
	EndBehavior *PaymentLinkUpdateSubscriptionDataTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkUpdateSubscriptionDataParams struct {
	// All invoices will be billed using the specified settings.
	InvoiceSettings *PaymentLinkUpdateSubscriptionDataInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will declaratively set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link. Unlike object-level metadata, this field is declarative. Updates will clear prior values.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days" json:"trial_period_days,omitempty"`
	// Settings related to subscription trials.
	TrialSettings *PaymentLinkUpdateSubscriptionDataTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
	UnsetFields   []PaymentLinkUpdateSubscriptionDataParamsUnsetField   `form:"-" json:"-"`
}

// PaymentLinkUpdateSubscriptionDataParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdateSubscriptionDataParams.
type PaymentLinkUpdateSubscriptionDataParamsUnsetField string

const (
	PaymentLinkUpdateSubscriptionDataParamsUnsetFieldMetadata        PaymentLinkUpdateSubscriptionDataParamsUnsetField = "metadata"
	PaymentLinkUpdateSubscriptionDataParamsUnsetFieldTrialPeriodDays PaymentLinkUpdateSubscriptionDataParamsUnsetField = "trial_period_days"
	PaymentLinkUpdateSubscriptionDataParamsUnsetFieldTrialSettings   PaymentLinkUpdateSubscriptionDataParamsUnsetField = "trial_settings"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdateSubscriptionDataParams) AddUnsetField(field PaymentLinkUpdateSubscriptionDataParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkUpdateSubscriptionDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls tax ID collection during checkout.
type PaymentLinkUpdateTaxIDCollectionParams struct {
	// Enable tax ID collection during checkout. Defaults to `false`.
	Enabled *bool `form:"enabled" json:"enabled"`
	// Describes whether a tax ID is required during checkout. Defaults to `never`. You can't set this parameter if `ui_mode` is `custom`.
	Required *string `form:"required" json:"required,omitempty"`
}

// Updates a payment link.
type PaymentLinkUpdateParams struct {
	Params `form:"*"`
	// Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.
	Active *bool `form:"active" json:"active,omitempty"`
	// Behavior after the purchase is complete.
	AfterCompletion *PaymentLinkUpdateAfterCompletionParams `form:"after_completion" json:"after_completion,omitempty"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes" json:"allow_promotion_codes,omitempty"`
	// Configuration for automatic tax collection.
	AutomaticTax *PaymentLinkUpdateAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Configuration for collecting the customer's billing address. Defaults to `auto`.
	BillingAddressCollection *string `form:"billing_address_collection" json:"billing_address_collection,omitempty"`
	// Configures whether [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link create a [Customer](https://docs.stripe.com/api/customers).
	CustomerCreation *string `form:"customer_creation" json:"customer_creation,omitempty"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*PaymentLinkUpdateCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
	CustomText *PaymentLinkUpdateCustomTextParams `form:"custom_text" json:"custom_text,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The custom message to be displayed to a customer when a payment link is no longer active.
	InactiveMessage *string `form:"inactive_message" json:"inactive_message,omitempty"`
	// Generate a post-purchase Invoice for one-time payments.
	InvoiceCreation *PaymentLinkUpdateInvoiceCreationParams `form:"invoice_creation" json:"invoice_creation,omitempty"`
	// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
	LineItems []*PaymentLinkUpdateLineItemParams `form:"line_items" json:"line_items,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`. Metadata associated with this Payment Link will automatically be copied to [checkout sessions](https://docs.stripe.com/api/checkout/sessions) created by this payment link.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Controls settings applied for collecting the customer's name.
	NameCollection *PaymentLinkUpdateNameCollectionParams `form:"name_collection" json:"name_collection,omitempty"`
	// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
	// There is a maximum of 10 optional items allowed on a payment link, and the existing limits on the number of line items allowed on a payment link apply to the combined number of line items and optional items.
	// There is a maximum of 20 combined line items and optional items.
	OptionalItems []*PaymentLinkUpdateOptionalItemParams `form:"optional_items" json:"optional_items,omitempty"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *PaymentLinkUpdatePaymentIntentDataParams `form:"payment_intent_data" json:"payment_intent_data,omitempty"`
	// Specify whether Checkout should collect a payment method. When set to `if_required`, Checkout will not collect a payment method when the total due for the session is 0.This may occur if the Checkout Session includes a free trial or a discount.
	//
	// Can only be set in `subscription` mode. Defaults to `always`.
	//
	// If you'd like information on how to collect a payment method outside of Checkout, read the guide on [configuring subscriptions with a free trial](https://docs.stripe.com/payments/checkout/free-trials).
	PaymentMethodCollection *string `form:"payment_method_collection" json:"payment_method_collection,omitempty"`
	// The list of payment method types that customers can use. Pass an empty string to enable dynamic payment methods that use your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes []*string `form:"payment_method_types" json:"payment_method_types,omitempty"`
	// Controls phone number collection settings during checkout.
	//
	// We recommend that you review your privacy policy and check with your legal contacts.
	PhoneNumberCollection *PaymentLinkUpdatePhoneNumberCollectionParams `form:"phone_number_collection" json:"phone_number_collection,omitempty"`
	// Settings that restrict the usage of a payment link.
	Restrictions *PaymentLinkUpdateRestrictionsParams `form:"restrictions" json:"restrictions,omitempty"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkUpdateShippingAddressCollectionParams `form:"shipping_address_collection" json:"shipping_address_collection,omitempty"`
	// Describes the type of transaction being performed in order to customize relevant text on the page, such as the submit button. Changing this value will also affect the hostname in the [url](https://docs.stripe.com/api/payment_links/payment_links/object#url) property (example: `donate.stripe.com`).
	SubmitType *string `form:"submit_type" json:"submit_type,omitempty"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkUpdateSubscriptionDataParams `form:"subscription_data" json:"subscription_data,omitempty"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *PaymentLinkUpdateTaxIDCollectionParams `form:"tax_id_collection" json:"tax_id_collection,omitempty"`
	UnsetFields     []PaymentLinkUpdateParamsUnsetField     `form:"-" json:"-"`
}

// PaymentLinkUpdateParamsUnsetField is the list of fields that can be cleared/unset on PaymentLinkUpdateParams.
type PaymentLinkUpdateParamsUnsetField string

const (
	PaymentLinkUpdateParamsUnsetFieldCustomFields              PaymentLinkUpdateParamsUnsetField = "custom_fields"
	PaymentLinkUpdateParamsUnsetFieldInactiveMessage           PaymentLinkUpdateParamsUnsetField = "inactive_message"
	PaymentLinkUpdateParamsUnsetFieldNameCollection            PaymentLinkUpdateParamsUnsetField = "name_collection"
	PaymentLinkUpdateParamsUnsetFieldOptionalItems             PaymentLinkUpdateParamsUnsetField = "optional_items"
	PaymentLinkUpdateParamsUnsetFieldPaymentMethodTypes        PaymentLinkUpdateParamsUnsetField = "payment_method_types"
	PaymentLinkUpdateParamsUnsetFieldRestrictions              PaymentLinkUpdateParamsUnsetField = "restrictions"
	PaymentLinkUpdateParamsUnsetFieldShippingAddressCollection PaymentLinkUpdateParamsUnsetField = "shipping_address_collection"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLinkUpdateParams) AddUnsetField(field PaymentLinkUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *PaymentLinkUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentLinkUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
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
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmation `json:"hosted_confirmation,omitempty"`
	Redirect           *PaymentLinkAfterCompletionRedirect           `json:"redirect,omitempty"`
	// The specified behavior after the purchase is complete.
	Type PaymentLinkAfterCompletionType `json:"type"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type PaymentLinkAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account,omitempty"`
	// Type of the account referenced.
	Type PaymentLinkAutomaticTaxLiabilityType `json:"type"`
}
type PaymentLinkAutomaticTax struct {
	// If `true`, tax will be calculated automatically using the customer's location.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *PaymentLinkAutomaticTaxLiability `json:"liability"`
}

// Settings related to the payment method reuse text shown in the Checkout UI.
type PaymentLinkConsentCollectionPaymentMethodReuseAgreement struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.
	//
	// When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position PaymentLinkConsentCollectionPaymentMethodReuseAgreementPosition `json:"position"`
}

// When set, provides configuration to gather active consent from customers.
type PaymentLinkConsentCollection struct {
	// Settings related to the payment method reuse text shown in the Checkout UI.
	PaymentMethodReuseAgreement *PaymentLinkConsentCollectionPaymentMethodReuseAgreement `json:"payment_method_reuse_agreement"`
	// If set to `auto`, enables the collection of customer consent for promotional communications.
	Promotions PaymentLinkConsentCollectionPromotions `json:"promotions"`
	// If set to `required`, it requires cutomers to accept the terms of service before being able to pay. If set to `none`, customers won't be shown a checkbox to accept the terms of service.
	TermsOfService PaymentLinkConsentCollectionTermsOfService `json:"terms_of_service"`
}

// The options available for the customer to select. Up to 200 options allowed.
type PaymentLinkCustomFieldDropdownOption struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label string `json:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value string `json:"value"`
}
type PaymentLinkCustomFieldDropdown struct {
	// The value that pre-fills on the payment page.
	DefaultValue string `json:"default_value"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*PaymentLinkCustomFieldDropdownOption `json:"options"`
}
type PaymentLinkCustomFieldLabel struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom string `json:"custom"`
	// The type of the label.
	Type PaymentLinkCustomFieldLabelType `json:"type"`
}
type PaymentLinkCustomFieldNumeric struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue string `json:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength int64 `json:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength int64 `json:"minimum_length"`
}
type PaymentLinkCustomFieldText struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue string `json:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength int64 `json:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength int64 `json:"minimum_length"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type PaymentLinkCustomField struct {
	Dropdown *PaymentLinkCustomFieldDropdown `json:"dropdown,omitempty"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key     string                         `json:"key"`
	Label   *PaymentLinkCustomFieldLabel   `json:"label"`
	Numeric *PaymentLinkCustomFieldNumeric `json:"numeric,omitempty"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional bool                        `json:"optional"`
	Text     *PaymentLinkCustomFieldText `json:"text,omitempty"`
	// The type of the field.
	Type PaymentLinkCustomFieldType `json:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type PaymentLinkCustomTextAfterSubmit struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type PaymentLinkCustomTextShippingAddress struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type PaymentLinkCustomTextSubmit struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type PaymentLinkCustomTextTermsOfServiceAcceptance struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}
type PaymentLinkCustomText struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *PaymentLinkCustomTextAfterSubmit `json:"after_submit"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *PaymentLinkCustomTextShippingAddress `json:"shipping_address"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *PaymentLinkCustomTextSubmit `json:"submit"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *PaymentLinkCustomTextTermsOfServiceAcceptance `json:"terms_of_service_acceptance"`
}

// A list of up to 4 custom fields to be displayed on the invoice.
type PaymentLinkInvoiceCreationInvoiceDataCustomField struct {
	// The name of the custom field.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type PaymentLinkInvoiceCreationInvoiceDataIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account,omitempty"`
	// Type of the account referenced.
	Type PaymentLinkInvoiceCreationInvoiceDataIssuerType `json:"type"`
}

// Options for invoice PDF rendering.
type PaymentLinkInvoiceCreationInvoiceDataRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
	// ID of the invoice rendering template to be used for the generated invoice.
	Template string `json:"template"`
}

// Configuration for the invoice. Default invoice values will be used if unspecified.
type PaymentLinkInvoiceCreationInvoiceData struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// A list of up to 4 custom fields to be displayed on the invoice.
	CustomFields []*PaymentLinkInvoiceCreationInvoiceDataCustomField `json:"custom_fields"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Footer to be displayed on the invoice.
	Footer string `json:"footer"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *PaymentLinkInvoiceCreationInvoiceDataIssuer `json:"issuer"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Options for invoice PDF rendering.
	RenderingOptions *PaymentLinkInvoiceCreationInvoiceDataRenderingOptions `json:"rendering_options"`
}

// Configuration for creating invoice for payment mode payment links.
type PaymentLinkInvoiceCreation struct {
	// Enable creating an invoice on successful payment.
	Enabled bool `json:"enabled"`
	// Configuration for the invoice. Default invoice values will be used if unspecified.
	InvoiceData *PaymentLinkInvoiceCreationInvoiceData `json:"invoice_data"`
}
type PaymentLinkNameCollectionBusiness struct {
	// Indicates whether business name collection is enabled for the payment link.
	Enabled bool `json:"enabled"`
	// Whether the customer is required to complete the field before checking out. Defaults to `false`.
	Optional bool `json:"optional"`
}
type PaymentLinkNameCollectionIndividual struct {
	// Indicates whether individual name collection is enabled for the payment link.
	Enabled bool `json:"enabled"`
	// Whether the customer is required to complete the field before checking out. Defaults to `false`.
	Optional bool `json:"optional"`
}
type PaymentLinkNameCollection struct {
	Business   *PaymentLinkNameCollectionBusiness   `json:"business,omitempty"`
	Individual *PaymentLinkNameCollectionIndividual `json:"individual,omitempty"`
}
type PaymentLinkOptionalItemAdjustableQuantity struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled bool `json:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99.
	Maximum int64 `json:"maximum"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum int64 `json:"minimum"`
}

// The optional items presented to the customer at checkout.
type PaymentLinkOptionalItem struct {
	AdjustableQuantity *PaymentLinkOptionalItemAdjustableQuantity `json:"adjustable_quantity"`
	Price              string                                     `json:"price"`
	Quantity           int64                                      `json:"quantity"`
}

// Indicates the parameters to be passed to PaymentIntent creation during checkout.
type PaymentLinkPaymentIntentData struct {
	// Indicates when the funds will be captured from the customer's account.
	CaptureMethod PaymentLinkPaymentIntentDataCaptureMethod `json:"capture_method"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link.
	Metadata map[string]string `json:"metadata"`
	// Indicates that you intend to make future payments with the payment method collected during checkout.
	SetupFutureUsage PaymentLinkPaymentIntentDataSetupFutureUsage `json:"setup_future_usage"`
	// For a non-card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge.
	StatementDescriptor string `json:"statement_descriptor"`
	// For a card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge. Concatenated with the account's statement descriptor prefix to form the complete statement descriptor.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup string `json:"transfer_group"`
}
type PaymentLinkPhoneNumberCollection struct {
	// If `true`, a phone number will be collected during checkout.
	Enabled bool `json:"enabled"`
}
type PaymentLinkRestrictionsCompletedSessions struct {
	// The current number of checkout sessions that have been completed on the payment link which count towards the `completed_sessions` restriction to be met.
	Count int64 `json:"count"`
	// The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met.
	Limit int64 `json:"limit"`
}

// Settings that restrict the usage of a payment link.
type PaymentLinkRestrictions struct {
	CompletedSessions *PaymentLinkRestrictionsCompletedSessions `json:"completed_sessions"`
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
type PaymentLinkSubscriptionDataInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account,omitempty"`
	// Type of the account referenced.
	Type PaymentLinkSubscriptionDataInvoiceSettingsIssuerType `json:"type"`
}
type PaymentLinkSubscriptionDataInvoiceSettings struct {
	Issuer *PaymentLinkSubscriptionDataInvoiceSettingsIssuer `json:"issuer"`
}

// Defines how a subscription behaves when a free trial ends.
type PaymentLinkSubscriptionDataTrialSettingsEndBehavior struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod PaymentLinkSubscriptionDataTrialSettingsEndBehaviorMissingPaymentMethod `json:"missing_payment_method"`
}

// Settings related to subscription trials.
type PaymentLinkSubscriptionDataTrialSettings struct {
	// Defines how a subscription behaves when a free trial ends.
	EndBehavior *PaymentLinkSubscriptionDataTrialSettingsEndBehavior `json:"end_behavior"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionData struct {
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description     string                                      `json:"description"`
	InvoiceSettings *PaymentLinkSubscriptionDataInvoiceSettings `json:"invoice_settings"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link.
	Metadata map[string]string `json:"metadata"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int64 `json:"trial_period_days"`
	// Settings related to subscription trials.
	TrialSettings *PaymentLinkSubscriptionDataTrialSettings `json:"trial_settings"`
}
type PaymentLinkTaxIDCollection struct {
	// Indicates whether tax ID collection is enabled for the session.
	Enabled  bool                               `json:"enabled"`
	Required PaymentLinkTaxIDCollectionRequired `json:"required"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferData struct {
	// The amount in cents (or local equivalent) that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// The connected account receiving the transfer.
	Destination *Account `json:"destination"`
}

// A payment link is a shareable URL that will take your customers to a hosted payment page. A payment link can be shared and used multiple times.
//
// When a customer opens a payment link it will open a new [checkout session](https://docs.stripe.com/api/checkout/sessions) to render the payment page. You can use [checkout session events](https://docs.stripe.com/api/events/types#event_types-checkout.session.completed) to track payments through payment links.
//
// Related guide: [Payment Links API](https://docs.stripe.com/payment-links)
type PaymentLink struct {
	APIResource
	// Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.
	Active          bool                        `json:"active"`
	AfterCompletion *PaymentLinkAfterCompletion `json:"after_completion"`
	// Whether user redeemable promotion codes are enabled.
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// The ID of the Connect application that created the Payment Link.
	Application *Application `json:"application"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account.
	ApplicationFeePercent float64                  `json:"application_fee_percent"`
	AutomaticTax          *PaymentLinkAutomaticTax `json:"automatic_tax"`
	// Configuration for collecting the customer's billing address. Defaults to `auto`.
	BillingAddressCollection PaymentLinkBillingAddressCollection `json:"billing_address_collection"`
	// When set, provides configuration to gather active consent from customers.
	ConsentCollection *PaymentLinkConsentCollection `json:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Configuration for Customer creation during checkout.
	CustomerCreation PaymentLinkCustomerCreation `json:"customer_creation"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*PaymentLinkCustomField `json:"custom_fields"`
	CustomText   *PaymentLinkCustomText    `json:"custom_text"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The custom message to be displayed to a customer when a payment link is no longer active.
	InactiveMessage string `json:"inactive_message"`
	// Configuration for creating invoice for payment mode payment links.
	InvoiceCreation *PaymentLinkInvoiceCreation `json:"invoice_creation"`
	// The line items representing what is being sold.
	LineItems *LineItemList `json:"line_items,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata       map[string]string          `json:"metadata"`
	NameCollection *PaymentLinkNameCollection `json:"name_collection,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The optional items presented to the customer at checkout.
	OptionalItems []*PaymentLinkOptionalItem `json:"optional_items,omitempty"`
	// Indicates the parameters to be passed to PaymentIntent creation during checkout.
	PaymentIntentData *PaymentLinkPaymentIntentData `json:"payment_intent_data"`
	// Configuration for collecting a payment method during checkout. Defaults to `always`.
	PaymentMethodCollection PaymentLinkPaymentMethodCollection `json:"payment_method_collection"`
	// The list of payment method types that customers can use. When `null`, Stripe will dynamically show relevant payment methods you've enabled in your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes    []PaymentLinkPaymentMethodType    `json:"payment_method_types"`
	PhoneNumberCollection *PaymentLinkPhoneNumberCollection `json:"phone_number_collection"`
	// Settings that restrict the usage of a payment link.
	Restrictions *PaymentLinkRestrictions `json:"restrictions"`
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
