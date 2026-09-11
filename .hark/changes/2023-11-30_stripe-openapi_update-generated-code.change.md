---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1772
is_stripe_api_change: true
released_in_version: 76.7.0
---

* Add support for new resources `Climate.Order`, `Climate.Product`, and `Climate.Supplier`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `Order`
* Add support for `Get` and `List` methods on resources `Product` and `Supplier`
* Add support for new value `financial_connections_account_inactive` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for new values `climate_order_purchase` and `climate_order_refund` on enum `BalanceTransactionType`
* Add support for `Created` on `CheckoutSessionListParams`
* Add support for `ValidateLocation` on `CustomerTaxParams`
* Add support for new values `climate.order.canceled`, `climate.order.created`, `climate.order.delayed`, `climate.order.delivered`, `climate.order.product_substituted`, `climate.product.created`, and `climate.product.pricing_updated` on enum `EventType`
* Add support for new value `challenge` on enums `PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure` and `SetupIntentPaymentMethodOptionsCardRequestThreeDSecure`
