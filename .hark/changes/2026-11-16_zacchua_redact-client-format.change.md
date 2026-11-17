---
title: Keep API keys out of the printed output of clients and event notifications
semver_level: patch
jira_tickets_closed:
- RUN_DEVSDK-3729
---

- Printing a `Client` or a v2 event notification with `fmt` (any verb, including `%v`, `%+v`, `%#v`, and `%s`), `log`, or `slog`'s `TextHandler` no longer writes the client's API key.
- `Client` now implements `fmt.Formatter` and prints as `stripe.Client{...}`. A struct that embeds `stripe.Client` or `*stripe.Client` inherits this method, so printing that struct also uses this format.
