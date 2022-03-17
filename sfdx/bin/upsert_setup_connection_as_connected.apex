// Description: upsert connection status to look as if the account is connected
//              useful for testing
// sfdx force:apex:execute -u mbianco+standardcpq@stripe.com -f bin/upsert_setup_connection_as_connected.apex

Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
stripeConnectRec.Salesforce_Connected__c = true;
stripeConnectRec.Stripe_Connected__c = true;
upsert stripeConnectRec;