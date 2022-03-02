# stripe-salesforce

https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3

https://stripe-force.herokuapp.com/auth/salesforce

http://localhost:3100/

# PR Review Conventions

Here are some of the 'tags' that PR comments will be prefixed with to provide more context:

- no tag: requested change that requires either an explanation or fix
- `fyi` sharing some sort of domain or product knowledge that may be helpful
- `nit` non-blocking style or small change
- `future` something to consider or implement in the future

# SalesForce

## Connecting to a different org

- `sfdx/sfdx-project.json` edit the `sfdcLoginUrl` to match the URL you are logging in against
- `sfdx/.sfdx/sfdx-config.json` edit the `defaultdevhubusername` / `defaultusername` TODO confirm if this is needed

## Pulling Custom Fields

sfdx force:source:retrieve -m CustomField:Order.Stripe_Transaction_ID\_\_c

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
