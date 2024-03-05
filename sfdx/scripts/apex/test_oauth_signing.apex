// Description: Adds global key metadata so you don't have to.
//              Never run in production accounts. ONly for scratch accounts.
// Usage:  sfdx force:apex:execute -f scripts/apex/test_oauth_signing.apex |  grep USER_DEBUG | sed "s/.*|DEBUG|//"
public static List<Setup_Connection_Data__mdt> mdRecs = [
        SELECT Global_Key__c, Local_Authorization_Target__c, Enable_Apex_Debug__c, OAuth_AutoClose_Window__c,
                Enable_JS_Debug__c, Platform_Target__c, OAuth_State_Signing_Key__c,
                Middleware_Infrastructure_Version__c
        FROM Setup_Connection_Data__mdt
        LIMIT 1
];
public static Boolean getSampleState() {
    if (mdRecs.size() > 0) {
        Setup_Connection_Data__mdt mdRec = mdRecs[0];
        StripeOAuthState state = StripeOAuthState.create();
        System.debug('KEY: ' + mdRec.OAuth_State_Signing_Key__c);
        System.debug('STATEDATA: ' + JSON.serialize(state.data));
        System.debug('STATE: ' + state.toString());
        return true;
    }

    System.debug(LoggingLevel.ERROR, 'Run create_global_key_metadata before this script');
    return false;
}

getSampleState();