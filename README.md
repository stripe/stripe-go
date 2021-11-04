# stripe-salesforce

https://paper.dropbox.com/doc/SalesForceStripe-Connector-Architecture-A6jDl31hXxE2DOp9QKjl3

https://stripe-force.herokuapp.com/
http://localhost:3100/

## Rubocop Daemon

- Strip out the logic which checks for where rubocop is located and add `bundle exec` to the command prefix
- Start daemon in a separate terminal (`bundle exec rubocop-daemon start`)
