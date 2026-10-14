---
title: Support pagination for V2 list fields on resources
semver_level: minor
jira_tickets_closed:
- DEVSDK-3129
---

V2 list fields now deserialize their included page and support automatic pagination through `next_page_url` after being returned by a client method.
