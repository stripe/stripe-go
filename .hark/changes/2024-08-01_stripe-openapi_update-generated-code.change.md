---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1898
is_stripe_api_change: true
released_in_version: 79.7.0-beta.1
---

* Add support for `AttachPayment` method on resource `Invoice`
* Add support for `AppInstall` and `AppViewport` on `AccountSessionComponentsParams`
* Remove support for `PartnerRejectedDetails` on `DisputeEvidenceDetailsEnhancedEligibilityVisaCompellingEvidence3`
* Add support for `LinesInvalid` on `QuoteStatusDetailsStaleLastReason`
* Add support for new value `lines_invalid` on enum `QuoteStatusDetailsStaleLastReasonType`
* Add support for `LastPriceMigrationError` on `SubscriptionSchedule` and `Subscription`
