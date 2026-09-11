---
title: Updates beta branch with changes in master
pr_link: https://github.com/stripe/stripe-go/pull/1930
released_in_version: 80.2.0-beta.1
---

* Add support for `ReportingChart` on `AccountSessionComponentsParams`
* Remove support for `FromSchedule` on `QuoteSubscriptionData`
* Add support for `AllowRedisplay` on `TerminalReaderCollectPaymentMethodCollectConfigParams`
* Moved raw request functionality from `preview` package to `rawrequest` package, and removed `preview`
