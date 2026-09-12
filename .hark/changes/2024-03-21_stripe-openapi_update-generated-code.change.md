---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1827
is_stripe_api_change: true
released_in_version: 76.23.0-beta.1
---

* Add support for new resources `Entitlements.ActiveEntitlementSummary` and `Entitlements.ActiveEntitlement`
* Add support for `List` method on resource `ActiveEntitlement`
* Add support for `Mobilepay` on `ConfirmationTokenPaymentMethodDataParams` and `ConfirmationTokenPaymentMethodPreview`
* Add support for `UseStripeSDK` on `ConfirmationToken`
* Remove support for `PaymentMethod` on `ConfirmationToken`
* Change type of `ConfirmationTokenMandateData` from `ConfirmationTokensResourceMandateData` to `nullable(ConfirmationTokensResourceMandateData)`
* Add support for new value `mobilepay` on enum `ConfirmationTokenPaymentMethodPreviewType`
* Add support for `Metadata` on `EntitlementsFeatureParams` and `EntitlementsFeature`
* Add support for `Active` on `EntitlementsFeature`
* Add support for new value `entitlements.active_entitlement_summary.updated` on enum `EventType`
* Remove support for value `customer.entitlement_summary.updated` from enum `EventType`
