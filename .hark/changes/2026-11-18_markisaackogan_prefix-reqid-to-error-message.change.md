---
title: Prefix request ID to the error message as `[req_xxx]` if a request ID exists
---

`Error.Msg` and `V2RawError.Message` are now prefixed with `[req_xxx]`, so the prefix appears in the `message` field of the JSON returned by `Error()`.
