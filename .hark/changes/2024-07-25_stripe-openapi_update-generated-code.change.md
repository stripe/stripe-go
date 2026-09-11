---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1893
is_stripe_api_change: true
released_in_version: 79.6.1-beta.1
---

* Add support for new resources `Billing.AlertTriggered` and `Billing.Alert`
* Add support for `Activate`, `Archive`, `Deactivate`, `Get`, `List`, and `New` methods on resource `Alert`
* Add support for new values `issuing.account_closed_for_not_providing_business_model_clarification`, `issuing.account_closed_for_not_providing_url_clarification`, and `issuing.account_closed_for_not_providing_use_case_clarification` on enum `AccountNoticeReason`
* Add support for `DisplayName` on `TreasuryFinancialAccountParams` and `TreasuryFinancialAccount`
