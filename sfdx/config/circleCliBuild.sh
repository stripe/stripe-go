echo "Connect to Dev Hub Og"
sfdx force:config:set defaultdevhubusername=test-6c7dn5uulbli@example.com
sfdx auth:jwt:grant --clientid 3MVG9_7ddP9KqTzco5f8yk9A4U6BZW4SbOYWnBpGFyCu.8SSXLAZASV.Y4Iq4dredXg78NiRh4.Gkwltu05vw --jwtkeyfile /Users/arnoldezeolisa/Dropbox/Projects/stripe-salesforce/sfdx/JWT/server.key --username test-6c7dn5uulbli@example.com --instanceurl https://test.salesforce.com