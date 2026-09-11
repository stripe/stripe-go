---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2184
is_stripe_api_change: true
released_in_version: 83.1.0-alpha.6
---

* Add support for new resource `V2BillingPricingPlanSubscriptionComponents`
* Add support for `Get` method on resource `V2BillingPricingPlanSubscriptionComponents`
* Add support for `DimensionPayloadKeys` on `BillingMeterParams` and `BillingMeter`
* Add support for `DimensionFilters` and `DimensionGroupByKeys` on `BillingBillingMeterMeterEventSummaryListParams`
* Add support for `Dimensions` on `BillingMeterEventSummary`
* Add support for `FulfillmentDetails` and `PaymentMethodData` on `DelegatedCheckoutRequestedSessionParams`
* Add support for `LineItemDetails`, `Metadata`, `PaymentMethod`, and `SharedMetadata` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
* Add support for `Currency`, `Customer`, and `RiskDetails` on `DelegatedCheckoutRequestedSessionParams`
* Add support for `SellerDetails` and `SetupFutureUsage` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
* Add support for `AmountSubtotal`, `AmountTotal`, `CreatedAt`, `ExpiresAt`, `OrderDetails`, `SharedPaymentIssuedToken`, `Status`, `TotalDetails`, and `UpdatedAt` on `DelegatedCheckoutRequestedSession`
* Add support for `Address`, `Email`, `FulfillmentOptions`, `Name`, `Phone`, and `SelectedFulfillmentOption` on `DelegatedCheckoutRequestedSessionFulfillmentDetails`
* Add support for new values `billie`, `crypto`, `kr_card`, `kriya`, `mb_way`, `mondu`, `ng_bank_transfer`, `ng_bank`, `ng_card`, `ng_market`, `ng_ussd`, `ng_wallet`, `payco`, `paypay`, `rechnung`, `samsung_pay`, `satispay`, `scalapay`, `sequra`, `sunbit`, `us_bank_account`, and `vipps` on enums `EventsV2CoreHealthAuthorizationRateDropFiringEventImpact.PaymentMethodType`, `EventsV2CoreHealthAuthorizationRateDropResolvedEventImpact.PaymentMethodType`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact.PaymentMethodType`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact.PaymentMethodType`
