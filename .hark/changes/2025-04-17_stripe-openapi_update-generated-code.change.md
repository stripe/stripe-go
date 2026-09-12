---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2023
is_stripe_api_change: true
released_in_version: 82.1.0-beta.3
---

* Add support for new resources `FxQuote` and `PaymentIntentAmountDetailsLineItem`
* Add support for new services `fxquote.Client` (accessed by `client.API.FxQuotes`) and `paymentintentamountdetailslineitem.Client` (accessed by `client.API.PaymentIntentAmountDetailsLineItems`)
* Remove support for service `invoicepayment.Client` (accessed by `client.API.InvoicePayments`)
* Add support for `Get`, `List`, and `New` methods on resource `FxQuote`
* Remove support for `AttachPaymentIntent` method on resource `Invoice`
* Remove support for `Get` and `List` methods on resource `InvoicePayment`
* Add support for `List` method on resource `PaymentIntentAmountDetailsLineItem`
* Add support for `List` method on service `paymentintentamountdetailslineitem.Client`
* Add support for `Get`, `List`, and `New` methods on service `fxquote.Client`
* Remove support for `Get` and `List` methods on service `invoicepayment.Client`
* Remove support for `AttachPaymentIntent` method on service `invoice.Client`
* Add support for `RegistrationDate` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `USCfpbData` on `AccountParams`, `PersonParams`, `Person`, and `TokenPersonParams`
* Add support for `CustomerReference` and `OrderReference` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
* Add support for `TaxID` on `ChargeBillingDetails`, `ConfirmationTokenPaymentMethodDataBillingDetailsParams`, `ConfirmationTokenPaymentMethodPreviewBillingDetails`, `PaymentIntentConfirmPaymentMethodDataBillingDetailsParams`, `PaymentIntentPaymentMethodDataBillingDetailsParams`, `PaymentMethodBillingDetailsParams`, `PaymentMethodBillingDetails`, `SetupIntentConfirmPaymentMethodDataBillingDetailsParams`, `SetupIntentPaymentMethodDataBillingDetailsParams`, `TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams`, and `TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams`
* Add support for `PriceData` on `CheckoutSessionLineItemParams`
* Add support for `Script` on `CouponParams` and `Coupon`
* Add support for `Type` on `Coupon`
* Add support for new value `fx_quote.expired` on enum `Event.Type`
* Add support for new value `affirm` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `FxQuote` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `PaymentIntent`, `TransferParams`, and `Transfer`
* Add support for `DiscountAmount`, `LineItems`, `Shipping`, and `Tax` on `PaymentIntentAmountDetails`
* Add support for `Pix` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `PendingReason` on `Refund`
* Add support for `Aw`, `Az`, `Bd`, `Bj`, `ET`, `Kg`, `La`, and `Ph` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for snapshot event `EventTypeFxQuoteExpired` with resource `FxQuote`
