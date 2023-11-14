// utilities that are NOT locked to our package to allow for anon-apex to handle various support issues
global with sharing class maintenanceUtilities {
  // resets the organizational-level API key
  // this is not accessible at all once the package is deployed
  global static void resetServiceConnection() {
    Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();
    utilities.generatePackageKey(stripeConnectRec);
  }

  global static void resetEnvironment() {
    maintenanceUtilities.resetSetupData();
    maintenanceUtilities.resetConnection();
  }

  @Future
  global static void resetSetupData() {
    List<Setup_Data__c> setupData = [SELECT Id FROM Setup_Data__c];
    delete setupData;
  }

  @Future
  global static void resetConnection() {
    Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();
    if (stripeConnectRec.Id != null) {
      delete stripeConnectRec;
    }
  }
}