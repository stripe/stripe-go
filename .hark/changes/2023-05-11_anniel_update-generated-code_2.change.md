---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1653
is_stripe_api_change: true
released_in_version: 74.18.0
---

* Add support for `Paypal` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `NetworkToken` on `ChargePaymentMethodDetailsCard`
* Add support for `TaxabilityReason` and `TaxableAmount` on `CheckoutSessionShippingCostTaxes`, `CheckoutSessionTotalDetailsBreakdownTaxes`, `CreditNoteShippingCostTaxes`, `CreditNoteTaxAmounts`, `InvoiceShippingCostTaxes`, `InvoiceTotalTaxAmounts`, `LineItemTaxes`, `QuoteComputedRecurringTotalDetailsBreakdownTaxes`, `QuoteComputedUpfrontTotalDetailsBreakdownTaxes`, and `QuoteTotalDetailsBreakdownTaxes`
* Add support for new value `paypal` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for new value `eftpos_au` on enums `PaymentIntentPaymentMethodOptionsCardNetwork`, `SetupIntentPaymentMethodOptionsCardNetwork`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCardNetwork`
* Add support for new value `paypal` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `Brand`, `CardholderName`, `Country`, `ExpMonth`, `ExpYear`, `Fingerprint`, `Funding`, `Last4`, `Networks`, and `ReadMethod` on `PaymentMethodCardPresent` and `PaymentMethodInteracPresent`
* Add support for `PreferredLocales` on `PaymentMethodInteracPresent`
* Add support for new value `paypal` on enum `PaymentMethodType`
* Add support for `EffectivePercentage` on `TaxRate`
* Add support for `GBBankTransfer` and `JPBankTransfer` on `CustomerCashBalanceTransactionFundedBankTransfer `
