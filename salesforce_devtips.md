## General

- List of system level permissions https://developer.salesforce.com/docs/atlas.en-us.sfFieldRef.meta/sfFieldRef/salesforce_field_reference_PermissionSet.htm

## Metadata

- [You cannot specify a default ListView](https://salesforce.stackexchange.com/questions/124447/default-listviews-in-lightning)
- You can't modify most types of metadata once it is created. Be very careful when creating fields. https://developer.salesforce.com/docs/atlas.en-us.236.0.packagingGuide.meta/packagingGuide/modifying_custom_fields.htm
- What metadata is available on the account?
- `sfdx force:mdapi:listmetadata -m Layout` to get all layouts on an account. If you want to pull a namespaced layout: `sfdx force:source:retrieve -m "Layout:Account-SBQQ__CPQ Account Layout"`
- Pull custom field from an account: `sfdx force:source:retrieve -m CustomField:Order.Stripe_Transaction_ID__c`

## Lighting Web Componetns

- `#lightning` in `pages/setup.page` is where the primary component is mounted

## Apex

- If `Apex CPU time limit exceeded` is encountered all DB operations are not committed. They are all wrapped into a transaction that is committed at the end of the Apex call.
- You can't do a callout after a DML (DB) operation. TODO I may be getting the order of operations wrong here, look this up
- You can't use variables in the `FROM` clause of a SOQL query `[...]`
- `--loglevel` seems to mess with some commands
- If you run apex anon, `sfdx` seems to truncate logs. The top of the logs won't come through.

## SOQL

- `IS NOT NULL` == `field != null`
- There is some fancy relationship syntax. You can use 1:1 lookups in `SELECT` and `WHERE` without a join
- You only `SELECT` via SOQL. All other CRUD operations are done through a ORM-like flow.

## CPQ

- `SBQQ__PriceEditable__c` must be true on the line to customize the price later on

## Tools

Most of these are blocked, but still interesting finds:

- https://www.pocketsoap.com/osx/soqlx/#Download
- https://chrome.google.com/webstore/detail/salesforce-inspector/aodjmnfhjibkcdimpodiifdjnnncaafh?hl=en
- https://github.com/motiko/sfdc-debug-logs
