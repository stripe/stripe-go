---
title: Fix passing context to EventNotificationHandler callbacks & update example
pr_url: https://github.com/stripe/stripe-go/pull/2238
is_breaking: true
released_in_version: 84.4.0-beta.1
---

- Fixes a bug where the first argument to event registration functions (e.g. `stripe.EventNotificationHandler.OnV1BillingMeterErrorReportTriggered`) didn't take a `context.Context` argument. To fix, we added the `ctx` argument:
    - before: `func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error`
    - after: `func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error`
- this is a breaking change if you're already using the new `EventNotificationHandler`. You'll need to update the function you're registering.
