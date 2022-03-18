// Description: delete connection status to look as if the account is fresh
//              useful for testing
// usage: sfdx force:apex:execute -f bin/delete_setup_connection_data.apex -u mbianco+standardcpq@stripe.com

Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
delete stripeConnectRec;

