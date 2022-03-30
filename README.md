# stripe-salesforce

https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3

https://stripe-force.herokuapp.com/auth/salesforce

http://localhost:3100/

```
Stripe.log_level = 'debug'; Restforce.log = true;
```

# PR Review Conventions

Here are some of the 'tags' that PR comments will be prefixed with to provide more context:

- no tag: requested change that requires either an explanation or fix
- `fyi` sharing some sort of domain or product knowledge that may be helpful
- `nit` non-blocking style or small change
- `followup` something to definitely implement in the future
- `future` something to consider in the future
- `curious`
- `suggestion`
- `issue`

# SalesForce

## Permission assignments

in the org I went to setup>permission sets> stripe connector integration user> manage assignments>add assignments then I checked your user in the org > then click assign (edited)

## Creating a stratch org

- Specify your dev hub via `sfdx force:config:set --defaultdevusername`. This can be any salesforce dev org.
- Then use `sfdx force:org:create -s -f config/project-scratch-def.json -a cool-alias -d 30` to create a stratch org. Lasts for a max of 30d.
  `alias sf=sfdx` is going to make your CLI-life easier.

Apex triggers run synchronously, TODO add to notes

## Creating a new package

[Here's the full guide to creating a new package](https://developer.salesforce.com/docs/atlas.en-us.packagingGuide.meta/packagingGuide/uploading_packages.htm). Here's how to find the URL of the latest package we have:

- Login to the packaging org (we have QA & Prod)
- Setup > Package Manager > Click on package name
- Versions > click on version number
- "Installation URL" is what you are looking for. The ID on the end can be used via SFDX.

Here's the shell command to install a package:

```shell
sfdx force:package:install --upgradetype DELETE -p 04t5f000000aSFv
```

Each version of a package has a unique URL.

## Manually Creating Global Key Metadata

This needs to be done when you are installing the source and not the package:

- Setup > Custom Metadata Types
- Setup Connection Data > Manage Records > New
- Label: Default, Global Key: $SF_MANAGED_PACKAGE_API_KEY

## Creating Metadata

- List all metadata within the account `sfdx force:mdapi:listmetadata -m CustomMetadata -u dev`
- Get field information on a metadata record:
  - `sfdx force:schema:sobject:describe -s Setup_Connection_Data__mdt -u dev | jq '.fields[] | [.name,.label,.type]'`
  - `sfdx force:schema:sobject:describe -s Setup_Connection_Data__mdt -u dev | jq '.fields | map(.name)'`
- Get the metadata value `sfdx force:data:soql:query -q "SELECT Id, Label, Global_Key__c FROM Setup_Connection_Data__mdt" -u dev`
- `sfdx force:cmdt:record:create --typename Setup_Connection_Data__mdt --recordname SetupConnectionData API_Key__c="" Global_Key__c="$SF_MANAGED_PACKAGE_API_KEY" Salesforce_Connected__c="" Stripe_Connected__c="" --protected=true`
- setup> custom metadata types> Stripe Connection Data > new and enter ‘Default’ for the label and developer name and add the key and save

## OAuth Tokens

- The Billing PBO Salesforce org holds our connected application
- Setup > App Manager > Stripe Connector > Edit to change scopes
- Same flow with a 'view' action to view consumer tokens

## Connecting to a different org

- `sfdx/sfdx-project.json` edit the `sfdcLoginUrl` to match the URL you are logging in against
- `sfdx/.sfdx/sfdx-config.json` edit the `defaultdevhubusername` / `defaultusername` TODO confirm if this is needed
- `sf alias:set standardcpq=mbianco+standardcpq@stripe.com` adds an entry to `~/.sfdx/alias.json`

## Pulling Custom Fields

sfdx force:source:retrieve -m CustomField:Order.Stripe_Transaction_ID\_\_c

- "A restricted picklist's values are limited to only those values defined by a Salesforce admin, which prevents users from loading redundant or erroneous values through the API." per the SF documentation on it"

## SOQL

- `IS NOT NULL` == `field != null`

## Tools

https://www.pocketsoap.com/osx/soqlx/#Download

# Development

## Deployment

`git remote add heroku https://git.heroku.com/stripe-force.git`

## Helpful Scripts

Checkout `scripts/` and `bin/` for helpful scripts to aid development.

## API

```shell
http POST https://stripe-force.herokuapp.com/v1/post-install Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:ORGANIZATION_KEY key=123123

http PUT https://stripe-force.herokuapp.com/v1/configuration field_defaults:='{"subscription_schedule":{"STRIPE_FIELD_2":"value"}}' field_mappings:='{}' Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:THE_KEY

http https://stripe-force.herokuapp.com/v1/configuration Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:THE_KEY
```

## Processes

`foreman start` will run all necessary processes with one command, or you can start them individually via `foreman run web`.

## Rubocop Daemon

- Strip out the logic which checks for where rubocop is located and add `bundle exec` to the command prefix
- Start daemon in a separate terminal (`bundle exec rubocop-daemon start`)

## Typechecking / Sorbet

Want to reference a local copy?

```shell
export SRB_SORBET_TYPED_REPO="/Users/mbianco/Projects/sorbet-typed"
export SRB_SORBET_TYPED_REVISION="mbianco/rails-fixes"
```

# Tests

- `bundle exec rake` will run the entire test suite
- `NO_RESCUE=1` to avoid autoloading pry-rescue in the test suite
