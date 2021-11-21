
echo "Connect to Dev Hub Og"
sfdx force:auth:web:login -d -a "stripescratch"
echo "Log in"
sfdx force:config:set defaultdevhubusername=test-6c7dn5uulbli@example.com
/Users/arnoldezeolisa/Dropbox/Projects/stripe-salesforce/sfdx/config/circleCliBuild.sh
