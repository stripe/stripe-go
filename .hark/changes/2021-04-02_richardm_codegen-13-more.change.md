---
title: "codegen: 13 more files"
pr_link: https://github.com/stripe/stripe-go/pull/1269
released_in_version: 72.41.0
---

* Add missing `Object` property to several structs
* Add support for `ExpiresAtNow` to `FileLinkParams`
* Add support for `SubscriptionItem` to `InvoiceItem`
* Add enum definitions for `TerminalReader.DeviceType`
* Add enum definitions for `Topup.status`
* Add support for `Amount`, `AmountRange`, and `Status` to `TopupListParams`
* Added custom `UnmarshalJSON` method for `Topup`
