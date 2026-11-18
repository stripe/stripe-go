---
title: Keep API keys out of the printed output of clients and event notifications
semver_level: patch
jira_tickets_closed:
- RUN_DEVSDK-3729
---

- Printing a `Client` or a v2 event notification with `fmt` (any verb, including `%v`, `%+v`, `%#v`, and `%s`), `log`, or `slog`'s `TextHandler` no longer writes the client's API key.
- `Client` now implements `fmt.Formatter` and prints as `stripe.Client{...}`.
- Every v2 event notification type now implements `fmt.Formatter` and prints its fields as before, except that its internal client is omitted.
- A struct that embeds `stripe.Client`, `*stripe.Client`, or an event notification type inherits that type's `Format` method, so printing the struct uses it too.
