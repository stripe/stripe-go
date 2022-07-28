// utilities that are NOT locked to our package to allow for anon-apex to handle various support issues
global with sharing class maintenanceUtilities {
  // resets the organizational-level API key
  // this is not accessible at all once the package is deployed
  global static void resetServiceConnection() {
    Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
    utilities.generatePackageKey(stripeConnectRec);    
  }
}