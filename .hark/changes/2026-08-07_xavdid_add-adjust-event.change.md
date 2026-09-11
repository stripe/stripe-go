---
title: add/adjust event parsing helpers
pr_link: https://github.com/stripe/stripe-go/pull/2399
released_in_version: 86.3.0
---

- Added methods that return their respective `Event`/`EventNotification` structs without verifying authenticity. Use them when you've previously verified an event (e.g. you verified, put the event in a queue, and are now processing). Supports events from [AWS EventBridge](https://docs.stripe.com/event-destinations/eventbridge) and [Azure Event Grid](https://docs.stripe.com/event-destinations/eventgrid) natively.
  - `stripe.ConstructEventWithoutVerification(payload, opts ...WebhookOption)`
  - `Client.ConstructEventWithoutVerification(payload, opts ...WebhookOption)`
  - `Client.ParseEventNotificationWithoutVerification(payload)`
