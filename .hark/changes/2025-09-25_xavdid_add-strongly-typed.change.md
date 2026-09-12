---
title: Add strongly typed EventNotifications
pr_url: https://github.com/stripe/stripe-go/pull/2121
is_breaking: true
released_in_version: 83.0.0
---

We've overhauled how V2 Events are handled in the SDK! This approach should provide a lot more information at authoring and compile time, leading to more robust integrations. As part of this process, there are a number of changes to be aware of.
- ⚠️ Rename function `Client.ParseThinEvent` to `Client.ParseEventNotification` and remove the `ThinEvent` struct.
    - This function now returns a `EventNotificationContainer` (which is an interface that all `EventNotification`s adhere to) instead of `ThinEvent`. When applicable, these event notifications will have the `RelatedObject` field and a function `FetchRelatedObject()`. They also have a `FetchEvent()` method to retrieve their corresponding event.
    - If you parse an event the SDK doesn't have types for (e.g. it's newer than the SDK you're using), you'll get an instance of `UnknownEventNotification` instead of a more specific type. It has both the `RelatedObject` field and the function `FetchRelatedObject()` (but they may be `nil`)
- ⚠️ Removed `API.parseThinEvent`. Use `Client.ParseEventNotification` instead (referring to the [migration guide](https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client) if necessary).
