# JWT App Authorization

**NOTE:** this should only be used for communication with CI. Otherwise, JWT authentication is _not_ used.

## JWT Setup

Generate a certificate and add it to the SF app:

1. `jwt-cert/generate`
2. Upload public_key.cer to the SF org managing the app. This is the Billing BPO.
   - Salesforce > Setup > Manage Connected apps > View > Edit
   - Use Digital Signature: true
   - Upload jwt-cert/public_key.cer
   - There are some unique permissions that need to be added to the app even if "full access" is requested.
3. Commit `public_key.pem` to the repo

Note that any changes to the app can take some time to persist on the Salesforce side.

## App Authorization

In order for JWT authorization to work on CI you need to 'preauthorize' the app against each SF account you want to use it in.

Visiting https://stripe-force.herokuapp.com/auth/salesforce on the account should be all you need to do.

## Related:

- https://medium.com/@tou_sfdx/salesforce-oauth-jwt-bearer-flow-cc70bfc626c2
- https://github.com/mars/make-connected-app-jwt
- https://github.com/lekkimworld/salesforce-jwt-generator
- https://mannharleen.github.io/2020-03-03-salesforce-jwt/
- https://medium.com/@tou_sfdx/salesforce-oauth-jwt-bearer-flow-cc70bfc626c2
