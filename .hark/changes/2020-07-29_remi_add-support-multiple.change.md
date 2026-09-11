---
title: Add support for multiple coupons on Billing APIs
pr_link: https://github.com/stripe/stripe-go/pull/1136
released_in_version: 71.40.0
---

* Add support for arrays of expandable API resources otherwise returning an array of strings by default
* Add custom deserialization to `Discount` to support expansion of the object
* Add support for `Id`, `Invoice` and `InvoiceItem` on `Discount`.
* Add support for `Discounts` on `Invoice`, `InvoiceItem` and `InvoiceLineItem`
* Add support for `DiscountAmounts` on `CreditNote`, `CreditNoteLineItem`, `InvoiceLineItem`
* Add support for `TotalDiscountAmounts` on `Invoice`
* Add `Object` to `Invoice`, `InvoiceLine`, `Discount` and `Coupon`
