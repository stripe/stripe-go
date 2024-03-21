/**
 * Created by rpena on 3/11/2024.
 */

global class PostInstallClass implements InstallHandler {
    global void onInstall(InstallContext context) {
        if((constants.NAMESPACE == 'QaStripeConnect' && context.previousVersion()?.compareTo(new Version(1,96)) < 0)
                || (constants.NAMESPACE == 'stripeConnector' && context.previousVersion()?.compareTo(new Version(2,8)) < 0) || constants.NAMESPACE == 'c') {
            //get response structure
            String route = UriTargetsV1.getConfigEndpoint();
            HttpResponse response = utilities.makeCallout(route, 'GET');

            //get stripe_account_id from response JSON
            Map<String, Object> result = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
            Map<String, Object> connectionStatus = (Map<String, Object>)result.get('connection_status');
            String stripe_id = connectionStatus.get('stripe_account_id').toString();
            Boolean liveMode = Sentry_Environment.getInstanceType() == Sentry_Environment.InstanceType.valueOf('PRODUCTION');

            //create Stripe_Account__c record if there are no other Stripe Accounts
            List<Stripe_Account__c> existingStripeAccounts = [SELECT Id, External_ID__c FROM Stripe_Account__c WHERE External_ID__c != NULL];
            if(existingStripeAccounts.size() == 0){
                Stripe_Account__c stripeAccount = new Stripe_Account__c();
                stripeAccount.Stripe_ID__c = stripe_id;
                stripeAccount.Is_Live_Mode__c = liveMode;
                stripeAccount.External_ID__c = stripe_id + '|' + liveMode;
                stripeAccount.Name = stripe_id;
                stripeAccount.Is_Primary__c = true;
                insert(stripeAccount);
            }
        }
    }
}