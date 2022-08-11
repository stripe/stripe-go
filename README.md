# Installation Links

Production Package Install:
https://login.salesforce.com/packaging/installPackage.apexp?p0=04t5f000000n8RNAAY

Update with `sfdx force:package1:version:list --json -u mbianco+stripeconnector@stripe.com | jq -r '.result[-1].MetadataPackageVersionId'`

QA Package Install:
https://test.salesforce.com/packaging/installPackage.apexp?p0=04t5f00000074nmAAA

Update with `sfdx force:package1:version:list --json -u mbianco+newstripeconnectorqa@stripe.com | jq -r '.result[-1].MetadataPackageVersionId'`

# stripe-salesforce

- [user documentation](https://docs.google.com/document/d/1GKfHhp0FwGHiAPxDv8AzJfnf0vJSbvF3ypMuzvhvZdw/edit#heading=h.rib60y3xmt51)
- [technical design](https://docs.google.com/document/d/1_7kf3oUTuNqQuopfapTP30Di1XwpuY7VYS_APYiGuPY/edit)
- [high-level architecture for security](https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3)
- [setup guide by appiphony](https://docs.google.com/document/d/17vE7-lL0DTnwRoVcYjgjDboQt72YmOVf3Mc4IS1mbWU/edit)
- [feature roadmap](https://docs.google.com/spreadsheets/d/136PUl_U7bMW7uMSwcqujisasAJNQJIOPQimGGG7iG00/edit#gid=0)
- [google drive folder with walkthrough videos and lots of docs](https://drive.google.com/drive/folders/14XhQGty83lqdMhLJduI9pwRI2mWAcrCl)

# URLs

- Production URL: https://salesforce.suitesync.io/auth/salesforce
- Dev URL: http://localhost:3100/
- [SF dev & test environments](https://docs.google.com/spreadsheets/d/136PUl_U7bMW7uMSwcqujisasAJNQJIOPQimGGG7iG00/edit#gid=0)
- [high-level architecture for security](https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3)
- [setup guide by appiphony](https://docs.google.com/document/d/17vE7-lL0DTnwRoVcYjgjDboQt72YmOVf3Mc4IS1mbWU/edit)

# Stripe Systems

- Sentry errors: https://sentry.corp.stripe.com/organizations/stripe/issues/?project=610
- Papertrail: https://my.papertrailapp.com/groups/16772762/events
  - Throw Stripe IDs or Salesforce IDs in the search to see a papertrail for them
  - Order ID from Salesforce are best

# Dev Snippets

Loud logging everywhere:

```
Stripe.log_level = 'debug'; Restforce.log = true; ENV['LOG_LEVEL'] = 'debug'
```

Or, via shell:

```
STRIPE_LOG=debug SALESFORCE_LOG=true LOG_LEVEL=DEBUG bundle exec ruby test/...
```

Include debugging tools locally. This is done automatically in dev.

```
require_relative './test/support/salesforce_debugging.rb'; include StripeForce::Constants; include SalesforceDebugging
```

Pull user reference. Helpful for console debugging:

```
@user = u = StripeForce::User[90]
```

Check SQL generated:

```ruby
locker = Integrations::Locker.new(user)
poller = StripeForce::OrderPoller.new(user)
poller.send(:generate_soql, DateTime.parse('2022-06-13 12:36:59 +0000'), DateTime.parse('2022-06-13 12:41:29'))
locker.clear_locked_resources
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

In order for the ruby service to update Stripe ID fields, it needs to have the right permissions. Here are the two permissions required:

- "Stripe Connector Integration User" permission set
- Update Activated Orders

In the org, go to setup>permission sets> stripe connector integration user> manage assignments>add assignments and check your user in the org > then click assign.

## Creating a stratch org

- Specify your dev hub via `sfdx force:config:set --defaultdevusername`. This can be any salesforce dev org.
- Then use `sfdx force:org:create -s -f config/project-scratch-def.json -a cool-alias -d 30` to create a stratch org. Lasts for a max of 30d.
  `alias sf=sfdx` is going to make your CLI-life easier.

Apex triggers run synchronously, TODO add to notes

## Finding an existing package

Here's how to find the URL of the latest package we have:

- Login to the packaging org (we have QA & Prod)
- Setup > Package Manager > Click on package name
- Versions > click on version number
- "Installation URL" is what you are looking for. The ID on the end can be used via SFDX.

Here's the shell command to install a package:

```shell
sfdx force:package:install --upgradetype DELETE -p 04t5f000000aSFv
```

Each version of a package has a unique URL.

## Production Package Deployment

[Here's a video walkthrough](https://drive.google.com/file/d/1Ok4Gl2rBJwy3w4AieeIa9PkWLxyK_4rP/view?usp=sharing)

First, deploy the source to our [production environment](https://docs.google.com/spreadsheets/d/136PUl_U7bMW7uMSwcqujisasAJNQJIOPQimGGG7iG00/edit#gid=0). Right now, this is the subdomain `appiphonycom5-dev-ed`. This org will _not_ expire since it is tied to a production package.

Here's how to deploy the source:

```shell
cd sfdx
sfdx force:source:deploy -p force-app/main/default -u mbianco+stripeconnector@stripe.com
```

If you don't have access to that account, add access:

```
sfdx auth:web:login
```

Then create a new package in the production packaging org.

- Setup > Package Manager (https://appiphonycom5-dev-ed.lightning.force.com/lightning/setup/Package/home)
- Click on the name of the package
- Copy the package name field (`stripeConnector`), you'll need it in the next step
- Click on upload, then:
  - Release Type > Managed Released
  - Version Name: `QaStripeConnect`
  - New version number will be determined automatically
  - New Order Save Behavior: unchecked (this should be done by default)
  - Do not worry about description or release notes

Some notes:

- globalGuidGenerator should not be included in the production package
- If you are adding new apex classes, JS code, remote sites, etc you need to click "add" and MANUALLY go through each category of items that you are adding and add them individually.
- You can automatically add stuff to a package, but it is not recommended since it is very to add items you don't want, which is hard to debug (and remove!)
- Removing something from a package is very hard. You need to (a) open up a case with SF and get them to switch the package to beta (b) deploy the latest version of the source (c) we remove the component.
  - Partner Community cases: "View My Cases". We need to provide the reason we are reverting back to beta. Although we need to remove the actual components.
- If something is causing an error that you deleted in your source code, you can edit the offending file to resolve whatever is causing the error. This is often easier than asking SF to move your package to beta.
- Ensure the "old order save behavior" is NOT checked

### Production Package Release

- https://security.secure.force.com/security/tools/forcecom/scanner use username and password of the production packaging org.
- Include sec rev in the username of another account which contains the latest production package.
- Make sure to write the Release Notes

### Production Package Testing

- We only wish to test the currently deployed production package. Therefore we do NOT run the salesforce tests like we do in QA, this is because those tests involve pushing a new source install.
- If you accidentally run these tests against production, you must follow the below steps to clean up the prod test account:

  - `sfdx force:source:delete -p force-app/main/default/ --apiversion=54.0 -u brennen+prodtest@stripe.com`
  - If you are running into issues, try deleting the conflicts via the salesforce UI or via SFDX but specify down the path to just the problematic files.
  - Some records may need to be manually altered:
    - Go to the [Lightning App Builder](https://productiontest2-dev-ed.lightning.force.com/lightning/setup/FlexiPageList/page?address=%2F_ui%2Fflexipage%2Fui%2FFlexiPageFilterListPage) and select the 'Sync Record Record Page' that lacks a Namespace Prefix and has an 'Edit' button. Click edit and then 'Activation' in the top right hand corner. Under 'Org Default' click 'Remove as Org Default' and select both Desktop and Phone. Then click Save.

- If you are running into permission errors related to Orders, you are most likely missing the standard CPQ setup permissions that a real account would have. We have some permission sets created for us in the force-app/main/scratchSetup folder. If you are creating a new testing environment you will have to deploy this folder and assign the permissions to your test user.

  - ie: `sfdx force:source:deploy -p force-app/main/scratchSetup -u brennen+prodtest@stripe.com` followed by `sfdx force:user:permset:assign -n "Order_Permissions" -u brennen`+prodtest@stripe.com
  - Right now, OrderPermissions is the only permission set we have defined but there may be additional ones in the future.

- If you are seeing Order upsert failures like "Can't save order products with negative quantities: Quantity"
  - Check the following boxes on this page: `sfdx force:org:open -u brennen+prodtest@stripe.com -p "/lightning/setup/OrderSettings/home"`
    - Enable Negative Quantity
    - Enable Zero Quantity

## Creating a new QA package

[Here's the full guide to creating a new package](https://developer.salesforce.com/docs/atlas.en-us.packagingGuide.meta/packagingGuide/uploading_packages.htm). Below is the short guide.

First, deploy the source to our [QA environment](https://docs.google.com/spreadsheets/d/136PUl_U7bMW7uMSwcqujisasAJNQJIOPQimGGG7iG00/edit#gid=0). You want to do this in stages to avoid errors.

```shell
cd sfdx
sfdx force:source:deploy -p force-app/main/default --apiversion=54.0 -u mbianco+newstripeconnectorqa@stripe.com
```

If you don't have access to that account, add access:

```
sfdx auth:web:login
```

Then create a new package in the QA org.

- Setup > Package Manager (https://appiphonycom7-dev-ed.lightning.force.com/lightning/setup/Package/home)
- Click on the name of the package
- Copy the package name field, you'll need it in the next step
- Click on upload, then:
  - Release Type > Managed Released
  - Version Name: `QaStripeConnect`
  - New version number will be determined automatically
  - New Order Save Behavior: unchecked (this should be done by default)

After rolling, make sure to install it on the package org test account (distinct from the packaging org, which is used to generate the package):

https://appiphony92-dev-ed.my.salesforce.com/

**Warning:** the QA package is automatically updated on each CI build, so it is possible that CI runs the deploy command _right_ after your deploy command finishes running, which could cause strange issues where QA includes a WIP branch.

**Note:** Everytime you upload a package you'll see a "Are you sure?" message letting you know you can't edit components. Ignore it.

## Removing Components from QA or Production Packages

As referenced above, it is a slightly convoluted process to remove a component from a package once it has been added. You must:

- Log into the partner community (https://help.salesforce.com/)
- Make sure you select technical support from the options drop down when creating a new case
- Provide them the org ID, package name and version number you would like rolled back to Beta.
- They will prompt you to do so but you can pre-emptively grant Salesforce Support access via:
  - Click your name -> My settings -> Personal -> Grant Account Login Access
  - Set the access duration for Salesforce.com Support to 1 Week
  - Click Save

For deleting objects, you may need to remove all the dependencies first by making a shell of the component and then deploying that.

- Example commit here: https://github.com/stripe/stripe-salesforce/commit/f

To delete lightning web components, you must delete them from the Lightning Components Page

- https://appiphonycom7-dev-ed.lightning.force.com/lightning/setup/LightningComponentBundles/home

## Manually Creating Global Key Metadata

This needs to be done when you are installing the source and not the package:

- Setup > Custom Metadata Types
- Setup Connection Data > Manage Records > New
- Label: Default, Global Key: $SF_MANAGED_PACKAGE_API_KEY

## Creating Salesforce Metadata

- List all metadata within the account `sfdx force:mdapi:listmetadata -m CustomMetadata -u dev`
- Get field information on a metadata record:
  - `sfdx force:schema:sobject:describe -s Setup_Connection_Data__mdt -u dev | jq '.fields[] | [.name,.label,.type]'`
  - `sfdx force:schema:sobject:describe -s Setup_Connection_Data__mdt -u dev | jq '.fields | map(.name)'`
- Get the metadata value `sfdx force:data:soql:query -q "SELECT Id, Label, Global_Key__c FROM Setup_Connection_Data__mdt" -u dev`
- `sfdx force:cmdt:record:create --typename Setup_Connection_Data__mdt --recordname SetupConnectionData API_Key__c="" Global_Key__c="$SF_MANAGED_PACKAGE_API_KEY" Salesforce_Connected__c="" Stripe_Connected__c="" --protected=true`
- setup> custom metadata types> Stripe Connection Data > new and enter ‘Default’ for the label and developer name and add the key and save

## Platform OAuth Tokens

- The Billing PBO Salesforce org holds our connected application
- Setup > App Manager > Stripe Connector > Edit to change scopes or add additional valid callback URLs
- Same flow with a 'view' action to view consumer tokens

## Batch Service Example

```
http -vvv POST https://appiphony15-dev-ed.my.salesforce.com/services/apexrest/batchApi 'Authorization: OAuth TOKEN' 'order_ids:=["8015f000000UNriAAG"]'
```

`TOKEN` and domain can be pulled from `.env` after refreshing tokens locally:

```
bundle exec ruby scripts/refresh-tokens.rb mbianco+biancodevorg@stripe.com
```

## Disable MFA (email code verification) on test accounts

Note this _should_ be handled automatically by `scratchSetup/settings/Security` when your scratch org is setup. You may need to deploy these settings on non-scratch orgs.

Disable Multi Factor Authentication (aka mfa, 2fa, email code verification):

- Setup > Session Settings > Session Security Levels
- remove Multi Factor Authentication

There a google plugin called Whitelist All IPs for Salesforce that our contractors use. Install that plugin and go to Setup, switch to Classic, then go to Network Access and there should be a Whitelist All IPs button that you click and will allow anyone to log in without having sent an email verification code.

## Connecting to a different org

TODO this may not be applicable anymore since we've removed the `.sfdx` folder from source control:

- `sfdx/sfdx-project.json` edit the `sfdcLoginUrl` to match the URL you are logging in against
- `sfdx/.sfdx/sfdx-config.json` edit the `defaultdevhubusername` / `defaultusername` TODO confirm if this is needed
- `sf alias:set standardcpq=mbianco+standardcpq@stripe.com` adds an entry to `~/.sfdx/alias.json`

## Resetting Salesforce Org Keys

- Debug console
- Anon Apex
- maintenanceUtilities.resetServiceConnection();

# Development

## Deployment

`git remote add heroku https://git.heroku.com/stripe-force.git`

## Helpful Scripts

Checkout `scripts/` and `bin/` for helpful scripts to aid development.

## API

```shell
http POST https://salesforce.suitesync.io/v1/post-install Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:ORGANIZATION_KEY key=123123

http PUT https://salesforce.suitesync.io/v1/configuration field_defaults:='{"subscription_schedule":{"STRIPE_FIELD_2":"value"}}' field_mappings:='{}' Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:THE_KEY

http https://salesforce.suitesync.io/v1/configuration Salesforce-Account-Id:00D5e000003V0C7EAK Salesforce-Key:THE_KEY
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

# Debugging

Install all gems locally to easily grep through them:

```shell
BUNDLE_DISABLE_SHARED_GEMS=1 BUNDLE_PATH=vendor/bundle bundle
```

# Tests

- Ruby:

  - Before running tests, you'll need a valid oauth token: `bundle exec ruby scripts/refresh-tokens.rb`
    - If the above script is not granting you a new access token for your scratch org, you can also generate a new one via `sfdx force:org:open -u brennen-scratch`
    - This will have refreshed your access token for SFDX, so you can re-run the refresh-tokens script above to replace it in your ENV.
  - `NO_RESCUE=true bundle exec rails test "test/**/test*.rb"` will run the entire test suite
  - `NO_RESCUE=1` to avoid autoloading pry-rescue in the test suite

- SFDX:
  - Run a single test: `sfdx force:apex:test:run --classnames <class_name> -u brennen-scratch --synchronous`
    - ie `sfdx force:apex:test:run --classnames test_updateSyncRecordsStatusesTrigger -u brennen-scratch --synchronous`
    - `--synchronous` will print the results of the tests instead of async execution, and a seperate command to get results.

# Heroku

```
heroku drains:add syslog+tls://logs5.papertrailapp.com:28081
```

After adding to papertrail, you'll need to rename the system to something legible. [More instructions here.](https://papertrailapp.com/systems/setup?type=system&platform=heroku)

- To run ruby arbitrary files on a dyno:
  - Add the file to /bin
  - Deploy to canary
  - `heroku run bash -a stripe-force-canary`
  - `bundle exec ruby bin/<script_file_name.rb>`
