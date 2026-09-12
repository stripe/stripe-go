---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1911
is_stripe_api_change: true
released_in_version: 79.12.0-beta.1
---

* Add support for new resources `Issuing.DisputeSettlementDetail` and `Issuing.Settlement`
* Add support for `Get` and `List` methods on resource `DisputeSettlementDetail`
* Remove support for `List` method on resource `QuotePhase`
* Add support for new values `issuing_dispute_settlement_detail.created`, `issuing_dispute_settlement_detail.updated`, `issuing_settlement.created`, and `issuing_settlement.updated` on enum `EventType`
* Add support for `Settlement` on `IssuingTransactionListParams` and `IssuingTransaction`
