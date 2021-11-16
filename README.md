# stripe-salesforce

https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3

https://stripe-force.herokuapp.com/auth/salesforce

http://localhost:3100/

# Development

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
