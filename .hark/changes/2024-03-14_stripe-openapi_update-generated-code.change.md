---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1825
is_stripe_api_change: true
released_in_version: 76.22.0-beta.1
---

* Add support for new resources `Billing.MeterEventAdjustment`, `Billing.MeterEvent`, and `Billing.Meter`
* Add support for `Deactivate`, `Get`, `List`, `New`, `Reactivate`, and `Update` methods on resource `Meter`
* Add support for `New` method on resources `MeterEventAdjustment` and `MeterEvent`
* Add support for `New` test helper method on resource `ConfirmationToken`
* Add support for `AddLines`, `RemoveLines`, and `UpdateLines` methods on resource `Invoice`
* Add support for `Multibanco` on `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `multibanco` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `SecondLine` on `IssuingPhysicalBundleFeatures`
* Add support for `MultibancoDisplayDetails` on `PaymentIntentNextAction`
* Add support for `Meter` on `PlanParams`, `Plan`, `PriceListRecurringParams`, `PriceRecurringParams`, and `PriceRecurring`
