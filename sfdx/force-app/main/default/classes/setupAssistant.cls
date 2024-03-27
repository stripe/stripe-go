public with sharing class setupAssistant {
    @TestVisible
    static final String STATUS_CONNECTED = 'connected';
    @TestVisible
    static final String STATUS_DISCONNECTED = 'disconnected';
    @TestVisible
    static final String STATUS_FAILED = 'failed';
    @TestVisible
    static final String STATUS_FRESH = 'fresh';
    @TestVisible
    static final String KEY_SALESFORCE = 'salesforce';
    @TestVisible
    static final String KEY_STRIPE = 'stripe';
    @TestVisible
    static final String KEY_STATUS = 'status';
    static final String KEY_SERVICE = 'service';
    @TestVisible
    static final String KEY_SERVICES = 'services';
    static final String KEY_CONNECTION_STATUS = 'connection_status';
    @TestVisible
    static final String SERVICE_DELIMINATOR = '|';
    @TestVisible
    static final String KEY_SALESFORCE_SERVICE = KEY_SALESFORCE + SERVICE_DELIMINATOR + UserInfo.getOrganizationId();
    @TestVisible
    static final String KEY_STRIPE_SERVICE = KEY_STRIPE + SERVICE_DELIMINATOR + KEY_STRIPE;

    @TestVisible
    private static ConfigPayloadHelper payloadHelper = new ConfigPayloadHelper();

    @TestVisible
    private static Debug_Helper debugger = new Debug_Helper();

    public class ConnectionNotInitializedException extends Sentry_Exception {}

    static Boolean isConnected() {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        return utilities.isConnected(stripeConnectRec);
    }

    @AuraEnabled
    public static String getNamespacePrefix() {
        return constants.NAMESPACE_API;
    }

    @AuraEnabled
    public static List<Stripe_Account__c> getStripeAccounts() {
        String deletedStatus = constants.STRIPE_ACCOUNT_CONN_STATUS_DELETED;
        return [
            SELECT Id, Name, Stripe_ID__c, Is_Live_Mode__c, External_ID__c, Connection_Status__c, Is_Primary__c
            FROM Stripe_Account__c
            WHERE Connection_Status__c != :deletedStatus
            WITH SECURITY_ENFORCED
            ORDER BY Is_Primary__c DESC, Stripe_ID__c ASC, Is_Live_Mode__c DESC
        ];
    }

    @AuraEnabled
    public static Boolean setStripeAccountAsDefault(String stripeAccountId, Boolean isLive) {
         if(!isConnected()) {
            throw new CalloutException('Failed to update the primary Stripe account.');
        }
        if (utilities.isUsingV2Middleware()) {
            // Ruby platform should take care of setting the primary account as not primary anymore
            // we should just send it the new primary account
            Map<String, Object> requestBody = new Map<String, Object> {
                'is_default_account_config' => true
            };
            
            Map<String,String> headers = new Map<String,String> {
                'Stripe-Primary-Account-Id' => stripeAccountId,
                'Stripe-Primary-Account-Livemode' => isLive.toString()
            };

            String route = UriTargets.getConfigEndpoint();
            HttpResponse response = utilities.makeCallout(route, 'PUT', JSON.serialize(requestBody), headers);
            utilities.validateCalloutResponse(response, 'setStripeAccountAsDefault');
        } else {
            String uri = UriTargets.getAccountSetDefaultEndpoint(stripeAccountId, isLive);
            HttpResponse response = utilities.makeCallout(uri, 'POST');

            if (response.getStatusCode() > 299) {
                throw new CalloutException('Failed to update the primary Stripe account.');
            }
        }

        // Update the Salesforce custom objects representing this Stripe account
        Boolean found = false;
        List<Stripe_Account__c> stripeAccounts = getStripeAccounts();
        for (Stripe_Account__c stripeAccount : stripeAccounts) {
            if (stripeAccount.Stripe_ID__c == stripeAccountId && stripeAccount.Is_Live_Mode__c == isLive) {
                stripeAccount.Is_Primary__c = true;
                found = true;
            } else {
                stripeAccount.Is_Primary__c = false;
            }
        }
        if (found) {
            update stripeAccounts;
            Stripe_Connection__c conn = utilities.getStripeConnection();
            conn.Primary_Stripe_Account_Id__c = stripeAccountId;
            conn.Primary_Stripe_Account_Livemode__c = isLive ? 'live' : 'test';
            update conn;
        }
        return found;
    }

    @AuraEnabled
    public static Boolean deleteStripeAccount(String stripeAccountId, Boolean isLive) {
        if(!isConnected()) {
            throw new CalloutException('Failed to remove Stripe account');
        }
        if (utilities.isUsingV2Middleware()) {
            // in v2, to "delete" an account, we just set enabled to "disabled"
            try {
                Map<String, Object> requestBody = new Map<String, Object> {
                    'enabled' => false
                };

                Map<String,String> headers = new Map<String,String> {
                    'Stripe-Primary-Account-Id' => stripeAccountId,
                    'Stripe-Primary-Account-Livemode' => isLive.toString()
                };

                String route = UriTargets.getConfigEndpoint();
                HttpResponse response = utilities.makeCallout(route, 'PUT', JSON.serialize(requestBody), headers);
                utilities.validateCalloutResponse(response, 'deleteStripeAccount');
            } catch (Exception e) {
                debugger.debug('deleteStripeAccount exception', e);
                return false;
            }
        } else {
            String uri = UriTargets.getAccountEndpoint(stripeAccountId, isLive);
            HttpResponse response = utilities.makeCallout(uri, 'DELETE');

            if (response.getStatusCode() > 299) {
                throw new CalloutException('Failed to remove Stripe account.');
            }
        }
    
        // Delete the Salesforce custom object representing this Stripe account
        Stripe_Account__c found = null;
        List<Stripe_Account__c> stripeAccounts = getStripeAccounts();
        for (Stripe_Account__c stripeAccount : stripeAccounts) {
            if (stripeAccount.Stripe_ID__c == stripeAccountId && stripeAccount.Is_Live_Mode__c == isLive) {
                found = stripeAccount;
            }
        }
        if (found != null) {
            debugger.debug('deleteStripeAccount found', found);
            found.Connection_Status__c = constants.STRIPE_ACCOUNT_CONN_STATUS_DELETED;
            update found;
        }
        return found != null;
    }

    // downloads sync record and config data from quick action button on sync record component
    @AuraEnabled(Cacheable=true)
    public static String getDownloadForSupportUrl(Id recordId) {
        PageReference ref = Page.DownloadForSupportConfig;
        ref.getParameters().put('recordId', recordId);
        return ref.getUrl();
    }

    @AuraEnabled(Cacheable=true)
    public static Boolean isJSDebugEnabled() {
        return utilities.isJSDebuggingEnabled();
    }

    @AuraEnabled(Cacheable=true)
    public static Boolean isOAuthAutoCloseWindowEnabled() {
        return utilities.isOAuthAutoCloseEnabled();
    }

    @AuraEnabled(Cacheable=false)
    public static Map<String, Object> validateSharedState(String state) {
        debugger.debug('validateSharedState state', state);
        try {
            Map<String, Object> data = StripeOAuthState.from_encrypted_state(state).data;
            debugger.debug('validateSharedState state-data', data);
            return data;
        } catch (Exception e) {
            errorLogger.create('OAuth state decryption failed', e);
            throw e;
        }
    }

    @AuraEnabled(Cacheable=false)
    public static Stripe_Account__c addStripeAccount(String state) {
        try {
            StripeOAuthState data = StripeOAuthState.from_encrypted_state(state);
            debugger.debug('addStripeAccount state-data', data);
            Stripe_Connection__c conn = utilities.getStripeConnection();

            String csac_id = data.csac_id;
            String stripe_account_id = data.stripe_account_id;
            String stripe_account_livemode = data.stripe_account_livemode;
            String stripe_account_name = data.stripe_account_name;
            if (String.isBlank(stripe_account_name)) {
                stripe_account_name = stripe_account_id;
            }

            if (String.isBlank(conn.Primary_Stripe_Account_Id__c)) {
                conn.Primary_Stripe_Account_Id__c = stripe_account_id;
                conn.Primary_Stripe_Account_Livemode__c = stripe_account_livemode;

                if (utilities.isUsingV2Middleware()) {
                    String salesforce_package_key = data.salesforce_package_key;
                    conn.API_Key__c = salesforce_package_key;
                }

                upsert conn;
            }
        
            Stripe_Account__c acct = new Stripe_Account__c(
                    Name = stripe_account_name,
                    Stripe_ID__c = stripe_account_id,
                    Is_Live_Mode__c = stripe_account_livemode == 'live',
                    External_ID__c = stripe_account_id + '|' + stripe_account_livemode,
                    CSAC_ID__c = csac_id,
                    Connection_Status__c = constants.STRIPE_ACCOUNT_CONN_STATUS_CONNECTED
            );

            acct.Is_Primary__c = utilities.getPrimaryStripeAccount() == null;
            debugger.debug('addStripeAccount acct-data', acct);


            Database.upsert(acct, Stripe_Account__c.External_ID__c);

            return acct;
        } catch (Exception e) {
            errorLogger.create('Adding stripe account failed', e);
            throw e;
        }
    }

    @AuraEnabled(Cacheable=true)
    public static String getExportableConfigDownloadUrl() {
        return Page.ConfigExportDownload.getUrl();
    }

    // gets steps completed from setup data to manage/presist states
    @AuraEnabled
    public static String getSetupData() {
        responseData rd = new responseData();
        Boolean isConnected = isConnected();
        try {
            List<Setup_Data__c> setupDataQuery = [
                SELECT Id, Steps_Completed__c, isSetupComplete__c
                FROM Setup_Data__c
                WHERE Name = :constants.SETUP_DATA_RECORD_NAME
                WITH SECURITY_ENFORCED
            ];
            Setup_Data__c setupData;

            if (!setupDataQuery.isEmpty()) {
                setupData = setupDataQuery[0];
            } else {
                // default values must be set explicitly since this record is NOT upserted and does not inherit database defaults
                setupData = new Setup_Data__c(
                    Steps_Completed__c = '{}',
                    Name = constants.SETUP_DATA_RECORD_NAME,
                    isSetupComplete__c = false
                );
            }

            rd.put('setupData', setupData);
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected', isConnected);
        return rd.getJsonString();
    }

    //saves steps completed as a setup data record to manage/presist states
    @AuraEnabled
    public static String saveData(Map<String, Object> newSetupDataRec, Boolean isSetupComplete) {
        responseData rd = new responseData();
        try {
            if (newSetupDataRec != null) {
                Setup_Data__c setupData = (Setup_Data__c)utilities.getSobjectByMap(newSetupDataRec, constants.NAMESPACE_API + 'Setup_Data__c');

                List<Setup_Data__c> data = [SELECT Id, Steps_Completed__c
                                            FROM Setup_Data__c
                                            WHERE Name = :constants.SETUP_DATA_RECORD_NAME
                                            WITH SECURITY_ENFORCED];

                if (!data.isEmpty()) {
                    // override id
                    setupData.Id = data[0].Id;

                    // merge completed steps
                    if (!String.isEmpty(setupData.Steps_Completed__c) && !String.isEmpty(data[0].Steps_Completed__c)) {
                        Map<String, Object> stepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(data[0].Steps_Completed__c);
                        Map<String, Object> newStepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(setupData.Steps_Completed__c);
                        newStepsCompleted.putAll(stepsCompleted);
                        setupData.Steps_Completed__c = JSON.serialize(newStepsCompleted);
                    }

                    //update setup complete flag if true
                    if(isSetupComplete) {
                        setupData.isSetupComplete__c = true;
                    }
                }

                // override name
                setupData.Name = constants.SETUP_DATA_RECORD_NAME;

                upsert setupData;
            }
            return getSetupData();
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    @AuraEnabled
    public static String revokeIntegrationUserAuthorization() {
        responseData rd = new responseData();
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        stripeConnectRec.Salesforce_Connected__c = false;
        rd.put(KEY_STATUS,     STATUS_DISCONNECTED);
        rd.put(KEY_SERVICE, KEY_SALESFORCE_SERVICE);
        update stripeConnectRec;
        return rd.getJsonString();
    }

    // validates connection status
    @AuraEnabled(Cacheable=false)
    public static String getIntegrationUserAuthorizationUri(Boolean forceLogin) {
        StripeOAuthState state = StripeOAuthState.create();
        responseData rd = new responseData();
        rd.put('authorization_uri', UriTargets.getIntegrationUserAuthorizationUrl(state, forceLogin));
        rd.put('message_origin_uri', UriTargets.getAuthorizationBaseUri());
        return rd.getJsonString();
    }


    // validates connection status
    @AuraEnabled(Cacheable=false)
    public static String getStripeAuthorizationUri(Boolean liveMode) {
        StripeOAuthState state = StripeOAuthState.create();
        responseData rd = new responseData();
        rd.put('authorization_uri', UriTargets.getStripeAuthorizationUri(liveMode, state));
        rd.put('message_origin_uri', UriTargets.getAuthorizationBaseUri());
        return rd.getJsonString();
    }


    // validates connection status
    @AuraEnabled
    public static String validateConnectionStatus(String systemToCheck) {
        responseData rd = new responseData();
        Map<String, Object> config = new Map<String, Object>();
        config.put('platform_uri', UriTargets.getAuthorizationBaseUri());
        config.put('salesforce_sandbox_uri_fragment', '/auth/v2/salesforcesandbox');
        config.put('salesforce_production_uri_fragment', '/auth/v2/salesforce');
        config.put('stripe_live_uri_fragment', '/auth/v2/stripelivemode');
        config.put('stripe_test_uri_fragment', '/auth/v2/stripetestmode');
        config.put('is_sandbox', constants.IS_SANDBOX);
        config.put('salesforce_namespace', constants.NAMESPACE);
        rd.put('config', config);

        StripeOAuthState state = StripeOAuthState.create();
        config.put('state', state.toString());

        try {
            // gets custom setting record to determin if we are in a sandbox or production org

            Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();

            // checks if we have a org-level API key generated, if we don't we should generate one
            // saving this conditional for later down the line...
            if (stripeConnectRec.Id == null) {
                // NOTE this is a VERY important line! `generatePackageKey` makes another callout and persists the organization-specific API key
                // on the ruby service. This key is also persisted locally within the salesforce account.
                if (utilities.shouldInitializeMiddlewareConnection()) {
                    utilities.generatePackageKey(stripeConnectRec);
                }

                rd.put(KEY_STRIPE,     STATUS_DISCONNECTED);
                rd.put(KEY_SALESFORCE, STATUS_DISCONNECTED);
                return rd.getJsonString();
            }

            // make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response = utilities.makeCallout(route, 'GET');

            Map<String, Object> responseBody;

            /*
            if we do not get a 200 status code back we create an Error Log record to show we did not get a success
            response from the ruby service and send a status of failed to signify to show an error toast
            */
            if(response.getStatusCode() != 200) {
                rd.put(KEY_STRIPE,     STATUS_FAILED);
                rd.put(KEY_SALESFORCE, STATUS_FAILED);
                errorLogger.create('validateConnectionStatus', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get connection status from Ruby service.');
                return rd.getJsonString();
            }

            /*
            if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not
            */
            responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

            //get connection status values from response
            Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get(KEY_CONNECTION_STATUS);
            stripeConnectRec.Stripe_Connected__c = (Boolean)connectionStatus.get(KEY_STRIPE);
            stripeConnectRec.Salesforce_Connected__c = (Boolean)connectionStatus.get(KEY_SALESFORCE);

            // Question: why do we care if it's a connected callback? Why can't we just return the same data regardless?
            // Answer: returning 'failed' tells us if there is a problem between the ruby service and salesforce.
            // Answer 2: It was the tell for if we were checking a single service or both.
            // Question: why do we need distinct strings here for the status? Can't we just pass a boolean?
            // Answer: The value is an enum: connected, disconnected, failed, or fresh.

            rd.put(KEY_SALESFORCE, (stripeConnectRec.Salesforce_Connected__c) ? STATUS_CONNECTED : STATUS_DISCONNECTED);
            rd.put(KEY_STRIPE,     (stripeConnectRec.Stripe_Connected__c)     ? STATUS_CONNECTED : STATUS_DISCONNECTED);

            // We return a connection status of 'fresh' to determine if we should show a toast message of success or not
            if (systemToCheck == KEY_SALESFORCE) {
                rd.put(KEY_SALESFORCE, (stripeConnectRec.Salesforce_Connected__c) ? STATUS_FRESH : STATUS_DISCONNECTED);
            } else if (systemToCheck == KEY_STRIPE) {
                rd.put(KEY_STRIPE,     (stripeConnectRec.Stripe_Connected__c)     ? STATUS_FRESH : STATUS_DISCONNECTED);
            }

            upsert stripeConnectRec;
            return rd.getJsonString();
        } catch (Exception e) {
            rd.addError(e);
        }

        return rd.getJsonString();
    }

    @AuraEnabled
    public static String checkServiceConnectionStatuses() {
        responseData rd = new responseData();
        // gets custom setting record to determine if we are in a sandbox or production org
        Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();

        // checks if we have a org-level API key generated, if we don't we should generate one
        if (stripeConnectRec.Id == null) {
            // NOTE this is a VERY important line! `generatePackageKey` makes another callout and persists the organization-specific API key
            // on the ruby service. This key is also persisted locally within the salesforce account.
            if (utilities.shouldInitializeMiddlewareConnection()) {
                utilities.generatePackageKey(stripeConnectRec);
            }

            rd.put(KEY_SERVICES, new Map<String, String>{
                    KEY_SALESFORCE_SERVICE => STATUS_DISCONNECTED
            });
            rd.put('message', 'Not yet connected.');
        } else {
            Stripe_Connection__c conn = utilities.getStripeConnection();
            String salesforceStatus = (stripeConnectRec.Salesforce_Connected__c) ? STATUS_CONNECTED : STATUS_DISCONNECTED;
            Map<String, Object> services = new Map<String, String>{
                    KEY_SALESFORCE_SERVICE => salesforceStatus
            };
            rd.put(KEY_SERVICES, services);

            List<Stripe_Account__c> stripeAccounts = getStripeAccounts();
            for (Stripe_Account__c stripeAccount : stripeAccounts) {
                String id = KEY_STRIPE + SERVICE_DELIMINATOR + stripeAccount.External_ID__c;
                String status = stripeAccount.Connection_Status__c;
                services.put(id, status);
            }
        }

        return rd.getJsonString();
    }

    // Gets the current connection statuses (connected, disconnected, failed, or fresh) for both Stripe and Salesforce
    // This forces checks of the tokens platform side.
    @AuraEnabled(Continuation=true)
    public static Object getServiceConnectionStatuses(String serviceToCheck) {
        debugger.debug('getServiceConnectionStatuses', serviceToCheck);

        // gets custom setting record to determin if we are in a sandbox or production org
        Stripe_Connection__c stripeConnectRec = getStripeConnection();

        // make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
        String route = UriTargets.getMultiAccountConnectionStatusesEndpoint();
        Integer timeout = 30000;
        HttpRequest request = utilities.buildRequest(route, 'GET', timeout);
        debugger.debug('request', request);
        Continuation con = new Continuation(timeout / 1000);
        con.addHttpRequest(request);
        con.state = serviceToCheck;
        con.continuationMethod = 'handleServiceConnectionStatusesResponse';
        if (Test.isRunningTest()) {
            con.continuationMethod = 'handleServiceConnectionStatusesResponseTest';
        }
        debugger.debug('getServiceConnectionStatuses','returning continuation...');
        return con;
    }

    // this is for tests... I know not why it's like this.
    public static Object handleServiceConnectionStatusesResponseTest() {
        List<String> labels = new List<String>();
        labels.add('continuation-1');
        String stripeAcct = 'acct_12345';
        String stripeLivemode = 'test';
        String stripeSvc = KEY_STRIPE + SERVICE_DELIMINATOR + stripeAcct + SERVICE_DELIMINATOR + stripeLivemode;

        return handleServiceConnectionStatusesResponse(labels, stripeSvc);
    }

    @AuraEnabled
    public static Object handleServiceConnectionStatusesResponse(List<String> labels, Object state) {
        debugger.debug('handleServiceConnectionStatusesResponse','resuming continuation...');
        debugger.debug('handleServiceConnectionStatusesResponse.labels',labels);
        debugger.debug('handleServiceConnectionStatusesResponse.state',state);
        String serviceToCheck = String.valueOf(state);

        return handleServiceConnectionStatusesResponseProcessor(Continuation.getResponse(labels[0]), serviceToCheck);
    }

    // Gets the current connection statuses (connected, disconnected, failed, or fresh) for both Stripe and Salesforce
    // This forces checks of the tokens platform side.
    public static String getServiceConnectionStatusesSync(String serviceToCheck) {
        debugger.debug('getServiceConnectionStatusesSync', serviceToCheck);

        // make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
        String route = UriTargets.getMultiAccountConnectionStatusesEndpoint();
        Integer timeout = 30000;
        HttpRequest request = utilities.buildRequest(route, 'GET', timeout);
        debugger.debug('request', request);

        responseData rd = new responseData();

        try {
            Http http = new Http();
            HttpResponse response = http.send(request);
            return (String) handleServiceConnectionStatusesResponseProcessor(response, serviceToCheck);
        } catch (Exception e) {
            rd.addError(e);
        }

        return rd.getJsonString();
    }

    public static Object handleServiceConnectionStatusesResponseProcessor(HttpResponse response, String serviceToCheck) {
        responseData rd = new responseData();
        Stripe_Connection__c stripeConnectRec = getStripeConnection();

        try {
            Map<String, Object> responseBody;
            Stripe_Connection__c conn = utilities.getStripeConnection();
            String STRIPE_SERVICE = String.isBlank(conn.Primary_Stripe_Account_Id__c) ? null : KEY_STRIPE + SERVICE_DELIMINATOR + conn.Primary_Stripe_Account_Id__c;

            /*
            if we do not get a 200 status code back we create an Error Log record to show we did not get a success
            response from the ruby service and send a status of failed to signify to show an error toast
            */
            if(response.getStatusCode() != 200) {
                rd.put(KEY_SERVICES, new Map<String, String> {
                        KEY_SALESFORCE_SERVICE => STATUS_DISCONNECTED
                });
                errorLogger.create('validateConnectionStatus', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get connection status from Ruby service.');
                rd.put('message', 'Failed to get connection status from Ruby service.' + response.getBody());
                return rd.getJsonString();
            }

            debugger.debug('handleServiceConnectionStatusesResponse response body', response.getBody());
            debugger.debug('handleServiceConnectionStatusesResponse svcToCheck', serviceToCheck);

            /*
            if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not
            */
            responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
            if (utilities.isUsingV2Middleware()) {
                responseBody = (Map<String,Object>)responseBody.get('statuses');
            }
            Boolean salesforceConnected;
            Boolean stripeConnected;
            Map<String, String> services = new Map<String, String>();

            //get connection status values from response
            debugger.debug('Got connection status result', responseBody);

            // Question: why do we care if it's a connected callback? Why can't we just return the same data regardless?
            // Answer: returning 'failed' tells us if there is a problem between the ruby service and salesforce.
            // Answer 2: It was the tell for if we were checking a single service or both.
            // Question: why do we need distinct strings here for the status? Can't we just pass a boolean?
            // Answer: The value is an enum: connected, disconnected, failed, or fresh.
            for (String key : responseBody.keySet()) {
                if (key.startsWith(KEY_SALESFORCE)) {
                    salesforceConnected = (Boolean)responseBody.get(key);
                    if (stripeConnectRec.Salesforce_Connected__c == false || stripeConnectRec.Salesforce_Connected__c == null) {
                        stripeConnectRec.Salesforce_Connected__c = salesforceConnected;
                    }
                    if (serviceToCheck == key) {
                        services.put(key, (salesforceConnected) ? STATUS_FRESH : STATUS_DISCONNECTED);
                    } else {
                        services.put(key, (salesforceConnected) ? STATUS_CONNECTED : STATUS_DISCONNECTED);
                    }
                } else if (key.startsWith(KEY_STRIPE)) {
                    stripeConnected = (Boolean)responseBody.get(key);
                    if (stripeConnectRec.Stripe_Connected__c == false || stripeConnectRec.Stripe_Connected__c == null) {
                        stripeConnectRec.Stripe_Connected__c = stripeConnected;
                    }

                    if (serviceToCheck == key) {
                        services.put(key, (stripeConnected) ? STATUS_FRESH : STATUS_DISCONNECTED);
                    } else {
                        services.put(key, (stripeConnected) ? STATUS_CONNECTED : STATUS_DISCONNECTED);
                    }
                }
            }

            debugger.debug('stripeConnectRec', stripeConnectRec);
            debugger.debug('services', services);

            rd.put(KEY_SERVICES, services);

            upsert stripeConnectRec;
            return rd.getJsonString();
        } catch (Exception e) {
            rd.addError(e);
        }

        return rd.getJsonString();
    }

    //gets object fields to populate picklist in mapper
    @AuraEnabled
    public static String getPicklistValuesForMapper(Boolean isConnectedCallback, String ObjectApiName) {
        responseData rd = new responseData();
        Boolean isConnected = false;
        try {
            isConnected = isConnected();

            if (isConnected) {
                if (isConnectedCallback == true) {
                    //builds a map of fields lists by object
                    Map<String, Set<Object>> fieldListByObjectMap = new Map<String, Set<Object>>{
                        'OrderItem' => utilities.getListOfFieldsByObject('OrderItem'),
                        'Order' => utilities.getListOfFieldsByObject('Order'),
                        'Account' => utilities.getListOfFieldsByObject('Account'),
                        'Product2' => utilities.getListOfFieldsByObject('Product2'),
                        'PricebookEntry' => utilities.getListOfFieldsByObject('PricebookEntry'),
                        'Order_Stripe_Coupon__c' => utilities.getListOfFieldsByObject(constants.NAMESPACE_API + 'Order_Stripe_Coupon__c')
                    };
                    rd.put('fieldListByObjectMap', fieldListByObjectMap);
                } else {
                    rd.put('ObjectApiName', ObjectApiName);
                    rd.put('listOfObjectFields', utilities.getListOfFieldsByObject(ObjectApiName));
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected', isConnected);
        return rd.getJsonString();
    }


    //gets formatted stripe object fields
    @AuraEnabled
    public static String getFormattedStripeObjectFields() {
        responseData rd = new responseData();
        try {
            //using a blank string here to avoid error for no `Salesforce-Key` in make callout function
            Map<String,String> headers = new Map<String,String> {
                'Salesforce-Key' => ''
            };

            // this has to stay the ruby base service uri because the request is being made from apex
            String route = UriTargets.getOpenApiSpecEndpoint();
            Map<String, Object> responseBody = null;

            HttpResponse response;
            response = utilities.makeCallout(route,'GET', headers);
            if (response.getStatusCode() == 200) {
                responseBody = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
                rd.put('formattedStripeCustomerFields', (List<Object>)responseBody.get('formattedStripeCustomerFields'));
                rd.put('formattedStripeProductItemFields', (List<Object>)responseBody.get('formattedStripeProductItemFields'));
                rd.put('formattedStripeSubscriptionFields', (List<Object>)responseBody.get('formattedStripeSubscriptionFields'));
                rd.put('formattedStripeSubscriptionSchedulePhaseFields', (List<Object>)responseBody.get('formattedStripeSubscriptionSchedulePhaseFields'));
                rd.put('formattedStripeSubscriptionItemFields', (List<Object>)responseBody.get('formattedStripeSubscriptionItemFields'));
                rd.put('formattedStripePriceFields', (List<Object>)responseBody.get('formattedStripePriceFields'));
                rd.put('formattedStripePriceOrderItemFields', (List<Object>)responseBody.get('formattedStripePriceFields'));
                rd.put('formattedStripeCouponFields', (List<Object>)responseBody.get('formattedStripeCouponFields'));
            } else {
                errorLogger.create('getFormattedStripeObjectFields', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //gets all saved mappings from ruby service
    @AuraEnabled(Cacheable = false)
    public static String getMappingConfigurations() {
        responseData rd = new responseData();
        Boolean isConnected = false;
        try {
            isConnected = isConnected();

            if (isConnected) {
                //used to store configuration maps from response
                Map<String, Object> allMappingConfigurations = new Map<String, Object>();

                //construct call out to rubys configuration endpoint
                String route = UriTargets.getConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route,'GET');

                //if the call out is successful pull out all mapping objects and add to list
                Map<String, Object> responseBody;
                if (response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

                    payloadHelper.extractConfigHash(responseBody, rd);

                    payloadHelper.extractFieldMappingData(responseBody, rd, true);
                } else {
                    errorLogger.create('getMappingConfigurations', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
                }
            }
        } catch (Exception e) {
            debugger.debug('getMappingConfigurations error', e);
            rd.addError(e);
        }
        rd.put('isConnected',isConnected);
        return rd.getJsonString();
    }

    //saves mappings and sends them to the ruby service to store
    @AuraEnabled
    public static String saveMappingConfigurations(String jsonMappingConfigurationsObject) {
        responseData rd = new responseData();
        try {
            if(isConnected()) {
                String route = UriTargets.getConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', jsonMappingConfigurationsObject);
                rd = utilities.validateCalloutResponse(response, 'saveMappingConfigurations');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

     //gets all saved sync prefrences data
     @AuraEnabled(Cacheable = false)
     public static String getSyncPreferences() {
         responseData rd = new responseData();
         try {
            rd.put('isSandbox', constants.IS_SANDBOX);
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if(utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'GET');

            Map<String, Object> errorBody;
            Map<String, Object> responseBody;

            /*if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not*/
            if (response.getStatusCode() == 200) {
                responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

                payloadHelper.extractConnectionStatus(responseBody, stripeConnectRec);

                //ensure we are still connected
                Boolean isConnected = stripeConnectRec.Stripe_Connected__c && stripeConnectRec.Salesforce_Connected__c;
                if (isConnected == false) {
                    rd.put('isConnected', false);
                    upsert stripeConnectRec;
                    return rd.getJsonString();
                }

                payloadHelper.extractConfigHash(responseBody, rd);
                payloadHelper.extractSettingsData(responseBody, rd, true);

                String recRetention = rd.getString(ConfigPayloadHelper.SYNC_RECORD_RETENTION);
                if (stripeConnectRec.Sync_Record_Retention__c != recRetention) {
                    stripeConnectRec.Sync_Record_Retention__c = recRetention;
                    upsert stripeConnectRec;
                }
            }
            rd.put('isConnected', true);
        } catch (Exception e) {
             rd.addError(e);
        }

        return rd.getJsonString();
    }

    //gets multi currency options if it is enablec in the org
    @AuraEnabled
    public static String getMulticurrencySelectionOptions() {
        responseData rd = new responseData();
        try {
            Boolean multiCurrencyEnabled = UserInfo.isMultiCurrencyOrganization();
            Map<String,System.OrgLimit> limitsMap = OrgLimits.getMap();
            System.OrgLimit apiRequestsLimit = limitsMap.get('DailyApiRequests');
            rd.put('orgMaxApiLimit', apiRequestsLimit.getLimit());
            if (multiCurrencyEnabled) {
                List<Object> supportedISOCodes = new List<Object>{
                    new Map<String, Object> {
                        'label' => 'Select default currency...',
                        'value' => ''
                    }
                };

                /* using a dynamic query here because the currencyType object can only be referenced directly
                when the multicurrency feature is enabled in the org otherwise we get a type error*/
                String supportedISOCodesQuery = 'SELECT ISOCode FROM CurrencyType WHERE IsActive = true';
                List<sObject> supportedISOCodesQueryResults = Database.query(supportedISOCodesQuery);

                //Iterate through the supported currencies to build list of supported ISO code
                for(sObject supportedISOCode : supportedISOCodesQueryResults) {
                    supportedISOCodes.add(new Map<String, Object> {
                        'label' => supportedISOCode.get('ISOCode'),
                        'value' => supportedISOCode.get('ISOCode')
                    });
                }
                rd.put('supportedISOCodes', supportedISOCodes);
            }
            rd.put('isMultiCurrencyEnabled', multiCurrencyEnabled);
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends updated sync preference data to the ruby service and saves the values
    @AuraEnabled
    public static String saveSyncPreferences(String defaultCurrency, String syncRecordRetention, String syncStartDate, String apiPercentageLimit, String cpqTermUnit, String cpqProratePrecision, Boolean pollingEnabled, String configurationHash) {
        responseData rd = new responseData();
        try {
            if (isConnected()) {
                Map<String, Object> settings = new Map<String, Object> {
                    'polling_enabled' => pollingEnabled,
                    'sync_start_date' => syncStartDate,
                    'sync_record_retention' => syncRecordRetention,
                    'default_currency' => UserInfo.getDefaultCurrency(),
                    'cpq_prorate_precision' => cpqProratePrecision
                };

                if (apiPercentageLimit != null) {
                    settings.put('api_percentage_limit', apiPercentageLimit);
                }

                Map<String, Object> requestBody = new Map<String, Object> {
                    'configuration_hash' => configurationHash,
                    'settings' => settings
                };

                String route = UriTargets.getConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(requestBody));
                rd = utilities.validateCalloutResponse(response, 'saveSyncPreferences');

            }
        } catch (Exception e) {
           rd.addError(e);
        }
        return rd.getJsonString();
    }

    //activates polling in the ruby service
    @AuraEnabled
    public static String activatePolling(String syncStartDate, Boolean isConfigEnabled, String configurationHash) {
        responseData rd = new responseData();
        debugger.debug('activatePolling isConfigEnabled', isConfigEnabled);
        try {
            if (isConnected()) {
                Map<String, Object> resquestBody = new Map<String, Object> {
                        'enabled' => isConfigEnabled,
                        'configuration_hash' => configurationHash,
                        'settings' => new Map<String, Object> {
                                'polling_enabled' => true,
                                'sync_start_date' => syncStartDate
                        }
                };

                String route = UriTargets.getConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(resquestBody));
                rd = utilities.validateCalloutResponse(response, 'activatePolling');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends records to be synced to ruby service called from quick action button on record detail page
    @AuraEnabled
    public static String manualRetry(Id recordId) {
        responseData rd = new responseData();
        Boolean isSyncRecordDispactched = false;
        try {
            if (isConnected()) {
                //get sync record information for ruby service
                List<Sync_Record__c> syncRecordList = [SELECT Id, Primary_Record_ID__c, Secondary_Record_ID__c, Resolution_Status__c, Secondary_Object_Type__c, Primary_Object_Type__c
                                                        FROM Sync_Record__c
                                                        WHERE Resolution_Status__c != 'Resolved'
                                                        AND Id = :recordId
                                                        WITH SECURITY_ENFORCED
                                                        LIMIT 1];

                if (syncRecordList.isEmpty()) {
                    rd.put('isSyncRecordDispactched', isSyncRecordDispactched);
                    return rd.getJsonString();
                }

                Sync_Record__c recordToSync = syncRecordList[0];
                
                String route = UriTargets.getTranslateEndpoint();

                Map<String, Object> body = new Map<String, Object>{
                    'object_type' => (String)recordToSync.Primary_Object_Type__c,
                    'object_ids' => new List<String> {
                        (String)recordToSync.Primary_Record_ID__c
                    }
                };

                HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));

                if (response.getStatusCode() == 200) {
                    isSyncRecordDispactched = true;
                } else {
                    rd.put('isSyncRecordDispactched', false);
                    errorLogger.create('manualRetry', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Sync_Record ID that failed: '+(String)recordId);
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isSyncRecordDispactched', isSyncRecordDispactched);
        return rd.getJsonString();
    }

     //gets all saved filter settings
     @AuraEnabled
    public static String getFilterSettings() {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'GET');

            if (response.getStatusCode() != 200) {
                errorLogger.create('getFilterSettings', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            Map<String, Object> responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

            //get connection status values from response
            Boolean isConnected = payloadHelper.extractConnectionStatus(responseBody, stripeConnectRec);

            //ensure we are still connected
            if (isConnected == false) {
                rd.put('isConnected', false);
                upsert stripeConnectRec;
                return rd.getJsonString();
            }

            payloadHelper.extractFilterSettingsData(responseBody, rd);
        } catch (Exception e) {
            rd.addError(e);
        }

        rd.put('isConnected', true);
        return rd.getJsonString();
    }

    //sends updated sync preference data to the ruby service and saves the values
    @AuraEnabled
    public static String saveFilterSettings(String productFilter, String orderFilter, String accountFilter, String priceBookFilter) {
        responseData rd = new responseData();
        try {
            Map<String, String> queryFilterMap = new Map<String, String> {
                'Product2' => productFilter,
                'Order' => orderFilter,
                'Account' => accountFilter,
                'PricebookEntry' => priceBookFilter
            };

            List<Map<String, Object>> validationErrorMapList = new List<Map<String, Object>> ();
            //Iterate through filters and run query to ensure they are valid
            for (String objectName : queryFilterMap.keySet()) {
                if (queryFilterMap.get(objectName) == null || String.isBlank(queryFilterMap.get(objectName))) {
                    continue;
                }

                String whereClause = queryFilterMap.get(objectName);
                try {
                    /*
                    Using limit zero in queries so there are no records are returned therefore it
                    will not count against the total number of rows retrieved salesforce limitation
                    */
                    List<sObject> sobjList = Database.query('SELECT Id FROM ' + objectName + ' WHERE ' + whereClause + ' WITH SECURITY_ENFORCED LIMIT 0');
                } catch (QueryException e) {
                    validationErrorMapList.add(new Map<String, Object> {
                        'Object' => objectName,
                        'Error' => String.valueOf(e)
                    });
                }
            }

            if (!validationErrorMapList.isEmpty()) {
                rd.put('isValidationError', true);
                rd.put('isFiltersSaved', false);
                rd.put('ValidationErrors', validationErrorMapList);
                return rd.getJsonString();
            }

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isFiltersSaved', false);
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            Map<String, Object> requestBody = new Map<String, Object> {
                'settings' => new Map<String, Object> {
                    'filters' => queryFilterMap
                }
            };
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(requestBody));
            rd = utilities.validateCalloutResponse(response, 'saveFilterSettings');

            if(response.getStatusCode() != 200) {
                rd.put('isValidationError', false);
                rd.put('isFiltersSaved', false);
                errorLogger.create('saveFilterSettings', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
            }

        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isFiltersSaved', true);
        rd.put('isValidationError', false);
        return rd.getJsonString();
    }

    @AuraEnabled(Cacheable=true)
    public static String getCoreData() {
        responseData rd = new responseData();
        try {
            if (utilities.isPackagedEnvironment()) {
                rd.put('packageVersion', System.requestVersion().toString());
            } else {
                rd.put('packageVersion', '0.0.0');
            }

            rd.put('configDownloadUri', Page.ConfigExportDownload.getUrl());

            rd.put('fieldListByObjectMap', new Map<String, Set<Object>>{
                'OrderItem' => utilities.getListOfFieldsByObject('OrderItem'),
                'Order' => utilities.getListOfFieldsByObject('Order'),
                'Account' => utilities.getListOfFieldsByObject('Account'),
                'Product2' => utilities.getListOfFieldsByObject('Product2'),
                'PricebookEntry' => utilities.getListOfFieldsByObject('PricebookEntry'),
                'Order_Stripe_Coupon__c' => utilities.getListOfFieldsByObject(constants.NAMESPACE_API + 'Order_Stripe_Coupon__c')
            });

            List<Map<String, String>> supportedIsoCodes = utilities.getMultiCurrencyOptions();
            rd.put('isMultiCurrencyEnabled', supportedIsoCodes.isEmpty() == false);
            rd.put('supportedISOCodes', supportedIsoCodes);
            rd.put('defaultCurrency', UserInfo.getDefaultCurrency());
            rd.put('isCpqInstalled', utilities.isCpqEnabled());
            rd.put('isSandbox', constants.IS_SANDBOX);
            Map<String,System.OrgLimit> limitsMap = OrgLimits.getMap();
            System.OrgLimit apiRequestsLimit = limitsMap.get('DailyApiRequests');
            rd.put('orgMaxApiLimit', apiRequestsLimit.getLimit());
        } catch (Exception e) {
            rd.addError('get core data', e);
        }
        return rd.getJsonString();
    }

    @AuraEnabled
    public static String getTranslationConfig() {
        responseData rd = new responseData();
        if (isConnected() == false) {
            rd.put('isConnected', false);
            rd.error = 'The connector is not yet registered with the platform. Please continue setup.';
            return rd.getJsonString();
        }
        rd.put('isConnected', true);

        try {
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'GET');
            Map<String, Object> responseBody = (Map<String, Object>) JSON.deserializeUntyped(response.getBody());
            debugger.debug('getTranslationConfig getStatusCode', response.getStatusCode());
            debugger.debug('getTranslationConfig response', response.getBody());
            if (response.getStatusCode() != 200) {
                String errorMessage = (String)responseBody.get('error');
                throw new CalloutException(errorMessage);
            }
            if (utilities.isUsingV2Middleware()) {
                responseBody = (Map<String,Object>)responseBody.get(ConfigPayloadHelper.V2_ACCOUNT_CONFIG);
            }
            payloadHelper.extractConfigHash(responseBody, rd);
            if (utilities.isUsingV2Middleware()) {
                payloadHelper.extractV2SettingsData(responseBody, rd, true);
            } else {
                payloadHelper.extractSettingsData(responseBody, rd, true);
            }
            payloadHelper.extractFieldMappingData(responseBody, rd, true);
            payloadHelper.extractFilterSettingsData(responseBody, rd);
        } catch (Exception e) {
            rd.addError('get configuration from platform', e);
        }

        return rd.getJsonString();
    }

    @AuraEnabled(Cacheable=true)
    public static String getDataMapperLayout() {
        return getFormattedStripeObjectFields();
    }

    @AuraEnabled
    public static String saveTranslationConfig(String config) {
        responseData rd = new responseData();
        if (isConnected() == false) {
            rd.put('isConnected', false);
            rd.put('isConfigSaved', false);
            rd.error = 'The connector is not yet registered with the platform. Please continue setup.';
            return rd.getJsonString();
        }
        rd.put('isConnected', true);
        debugger.debug('saveTranslationConfig config', config);
        Map<String, Object> conf = (Map<String, Object>) JSON.deserializeUntyped(config);
        debugger.debug('saveTranslationConfig conf', conf);
        Map<String, Object> payload = new Map<String, Object>();
        ConfigPayloadHelper.ErrorTracker errors = new ConfigPayloadHelper.ErrorTracker();
        try {
            Boolean configHashErr = payloadHelper.validateConfigHash(conf, payload, errors);
            Boolean dataMappingErr = payloadHelper.validateDataMappingConfig(conf, payload, errors);
            Boolean syncPrefsErr = payloadHelper.validateSyncPreferences(conf, payload, errors);
            Boolean filtersErr = payloadHelper.validateFilterConfig(conf, payload, errors);
            debugger.debug('saveTranslationConfig payload', payload);

            if (configHashErr || dataMappingErr || filtersErr || syncPrefsErr) {
                rd.put('isValidationError', true);
                rd.put('isFiltersSaved', false);
                rd.put('ValidationErrors', errors.errors);
                return rd.getJsonString();
            }
            
            if (conf.containsKey('enabled')) {
                payload.put('enabled', conf.get('enabled'));
            }

            String jsonPayload = JSON.serialize(payload);
            debugger.debug('saveTranslationConfig jsonPayload', jsonPayload);
            String route = UriTargets.getConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'PUT', jsonPayload);
            rd = utilities.validateCalloutResponse(response, 'saveConfiguration');
            debugger.debug('saveTranslationConfig response', response.getBody());
            rd.put('isConfigSaved', true);
            rd.put('isValidationError', false);
            rd.put('isFiltersSaved', true);
        } catch (Exception e) {
            rd.addError('save translation config', e);
            rd.put('isConfigSaved', false);
            rd.put('isValidationError', false);
            rd.put('isFiltersSaved', false);
        }

        return rd.getJsonString();
    }

    //sends object name for all records syncing to the ruby service
    @AuraEnabled
    public static String syncAllRecords(String objectType) {
        debugger.debug('setupAssistant.syncAllRecords', 'objectType: ' + objectType);
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            debugger.debug('setupAssistant.syncAllRecords', 'is connected');

            String route = UriTargets.getTranslateAllEndpoint();

            Map<String, Object> body = new Map<String, Object>{
                'object_type' => objectType
            };

            debugger.debug('setupAssistant.syncAllRecords', 'route: ' + route);
            debugger.debug('setupAssistant.syncAllRecords', 'body: ' + JSON.serialize(body));
            HttpResponse response = utilities.makeCallout(route, 'POST', JSON.serialize(body));

            debugger.debug('setupAssistant.syncAllRecords', 'Got response status code: ' + response.getStatusCode());

            if (response.getStatusCode() != 200) {
                rd.put('syncAllRecordsDispatched', false);
                errorLogger.create('syncAllRecords', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Object type that failed: ' + objectType);
                return rd.getJsonString();
            }

            rd.put('syncAllRecordsDispatched', true);
            rd.put('isConnected', true);

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends object name for all records syncing to the ruby service
    @AuraEnabled
    public static String manualTranslation(String translationRecordID) {
        responseData rd = new responseData();
        try {
            List<String> listOfAcceptableObjects = new List<String> {
                'Account',
                'Order',
                'Pricebook2',
                'Product2'
            };
            rd.put('isRecordValid', false);

            Id recordId = Id.valueOf(translationRecordID);
            String objectName = string.valueOf(recordId.getsobjecttype());
            if(!listOfAcceptableObjects.contains(objectName)) {
                rd.put('errorMessage', 'this is an id for a record on the ' + objectName + ' object which is not suppoted, full list of supported objects: ' + listOfAcceptableObjects);
                return rd.getJsonString();
            }

            String validationQuery = 'SELECT Id FROM ' + objectName + ' WHERE Id = :recordId WITH SECURITY_ENFORCED LIMIT 1';
            List<SObject> listOfFoundRecords = Database.query(validationQuery);
            if (listOfFoundRecords.isEmpty()) {
                rd.put('errorMessage', 'No ' + objectName + ' objects found for this Salesforce record Id');
                return rd.getJsonString();
            }

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //here is where we need to validate if its record or not
            String route = UriTargets.getTranslateEndpoint();

            Map<String, Object> body = new Map<String, Object> {
                'object_type' => objectName,
                'object_ids' => new List<String> {
                    translationRecordID
                }
            };

            debugger.debug('manualTranslation body', body);

            HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));
            if(response.getStatusCode() == 200) {
                rd.put('isRecordValid', true);
            } else {
                errorLogger.create('manualTranslation', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Translation Record ID That Failed: '+(String)recordId);
                throw new CalloutException('There was an error in queueing this record for translation, check sync logs for more detail');
            }

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //performs a permission check on user before granting access to the application
    @AuraEnabled
    public static String checkUserPermissions() {
        responseData rd = new responseData();
        try {
            //Map for front end state setting
            Map<String, Object> permissionIssueMap = new Map<String, Object> {
                'isPermSetAssigned' => false,
                'isSystemPermissionMissing' => false,
                'isObjectPermissionMissing' => false,
                'missingPermissionSets' => new List<String>(),
                'missingObjectPermissionList' => new List<String>(),
                'missingSystemPermissionList' => new List<String>()
            };

            //List of Objects to check access on
            List<String> listOfObjectsToCheckAccessOn = new List<String> {
                'Account',
                'Product2',
                'Order',
                'OrderItem',
                'PricebookEntry'
            };

            if ((Boolean)utilities.isCpqEnabled()) {
                listOfObjectsToCheckAccessOn.add('SBQQ__Quote__c');
                listOfObjectsToCheckAccessOn.add('SBQQ__OrderItemConsumptionSchedule__c');
                listOfObjectsToCheckAccessOn.add('SBQQ__OrderItemConsumptionRate__c');
                listOfObjectsToCheckAccessOn.add('ConsumptionSchedule');
                listOfObjectsToCheckAccessOn.add('ConsumptionRate');
            }

            List<Map<String,Object>> missingObjectPermissionList = validateObjectLevelPermission(listOfObjectsToCheckAccessOn);


            //add list of missing permissions if any were found
            if (!missingObjectPermissionList.isEmpty()) {
                permissionIssueMap.put('isObjectPermissionMissing', true);
                permissionIssueMap.put('missingObjectPermissionList', missingObjectPermissionList);
            }

            /* this list holds system permission fields which hold boolean denoting if the user has access or not.
            This is incase there are other system permissions need to be checked in the future. Full list of Profile fields
            here https://developer.salesforce.com/docs/atlas.en-us.sfFieldRef.meta/sfFieldRef/salesforce_field_reference_Profile.htm.*/
            List<String> permissionList = new List<String> {
                'Name',
                'PermissionsEditActivatedOrders'
            };

            //Used to get label names of permissions
            Map<String, Schema.SObjectType> schemaMap = Schema.getGlobalDescribe();
            Schema.SObjectType profileSchema = schemaMap.get('Profile');
            Map<String, Schema.SObjectField> fieldMap = profileSchema.getDescribe().fields.getMap();

            //Map of system permissions by label
            Map<String, Object> activePermissionMap = new Map<String, Object>();
            for (String permisson : permissionList) {
                if (permisson == 'Name') {
                    continue;
                }
                activePermissionMap.put(fieldMap.get(permisson).getDescribe().getLabel(), false);
            }

            List<String> userFieldsList = new List<String> {
                'Id',
                'Name'
            };

            for (String permission : permissionList) {
                userFieldsList.add('Profile.' + permission);
            }

            List<String> permissionSetAssignmentFieldsList = new List<String>();
            for (String permission : permissionList) {
                permissionSetAssignmentFieldsList.add('PermissionSet.' + permission);
            }

            List<SObject> userPermissions = Database.Query (
                String.join (
                    new List<String> {
                        'SELECT',
                        String.join(userFieldsList, ', '),
                        ',',
                        '(SELECT Id,',
                        String.join(permissionSetAssignmentFieldsList, ', '),
                        'FROM PermissionSetAssignments) FROM User WHERE Id = \'' +
                        UserInfo.getUserId() + '\' WITH SECURITY_ENFORCED LIMIT 1'
                    },
                    ' '
                )
            );

            if (userPermissions.isEmpty()) {
                rd.put('permissionIssueMap', permissionIssueMap);
                return rd.getJsonString();
            }
            SObject userProfileSystemPermissions = userPermissions[0].getSobject('Profile');

            //check profile level permissions
            for (String permisson : permissionList) {
                String permissionLabel = fieldMap.get(permisson).getDescribe().getLabel();
                if (permisson != 'Name') {
                    activePermissionMap.put(permissionLabel, (Boolean)userProfileSystemPermissions.get(permisson));
                }
            }

            Map<String, Object> validateUserPermissionsMap = validateUserPermissions(userPermissions, activePermissionMap, permissionIssueMap, fieldMap, permissionList);
            activePermissionMap = (Map<String, Object>)validateUserPermissionsMap.get('activePermissionMap');
            permissionIssueMap = (Map<String, Object>)validateUserPermissionsMap.get('permissionIssueMap');


            List<String> missingSystemPermissionList = new List<String>();
            for (String permission : activePermissionMap.keySet()) {
                if (!(Boolean)activePermissionMap.get(permission)) {
                    permissionIssueMap.put('isSystemPermissionMissing', true);
                    missingSystemPermissionList.add(permission);
                }
            }

            if (!missingSystemPermissionList.isEmpty()) {
                permissionIssueMap.put('missingSystemPermissionList', missingSystemPermissionList);
            }

            rd.put('permissionIssueMap', permissionIssueMap);

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    // builds map based on object permissioning
    private static List<Map<String,Object>> validateObjectLevelPermission(List<String> listOfObjectsToCheckAccessOn) {
        List<Map<String,Object>> missingObjectPermissionList = new List<Map<String,Object>>();
        for (String objectApiName : listOfObjectsToCheckAccessOn) {
            Boolean hasObjectAccess = true;
            DescribeSObjectResult objectDescribe = dmlManager.getSObjectDescribeByObjectName(objectApiName);

            Map<String,Object> missingPermissionsMap = new Map<String,Object> {
                'Object' => objectDescribe.getLabel(),
                'Access' => true,
                'Create' => true,
                'Edit' => true
            };

            if (!(Boolean)objectDescribe.isAccessible()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Access', false);
            }

            if (!(Boolean)objectDescribe.isCreateable()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Create', false);
            }

            if (!(Boolean)objectDescribe.isUpdateable()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Edit', false);
            }

            if (!hasObjectAccess) {
                missingObjectPermissionList.add(missingPermissionsMap);
            }
        }
        return missingObjectPermissionList;
    }

    // check each permset assigned to the user
    private static Map<String, Object> validateUserPermissions(List<SObject> userPermissions, Map<String, Object> activePermissionMap, Map<String, Object> permissionIssueMap, Map<String, Schema.SObjectField> fieldMap, List<String> permissionList) {
        List<SObject> userPermissionsAssignments = userPermissions[0].getSObjects('PermissionSetAssignments');
        Boolean isPermSetMissing = true;
        if (!userPermissionsAssignments.isEmpty()) {
            List<String> listOfAssignedPermissionSets = new List<String> {
                constants.PACKAGED_PERMISSION_SET_NAME
            };
            for (SObject userPermissionsAssignment : userPermissionsAssignments) {
                SObject userPermSetSystemPermissions = userPermissionsAssignment.getSobject('PermissionSet');

                if(userPermSetSystemPermissions.get('Name') == constants.PACKAGED_PERMISSION_SET_NAME.replace(' ','_')) {
                    isPermSetMissing = false;
                }

                for (String permisson : permissionList) {
                    String permissionLabel = fieldMap.get(permisson).getDescribe().getLabel();
                    if(permisson != 'Name') {
                        Boolean systemPermissionValue = (Boolean)userPermSetSystemPermissions.get(permisson);
                        //only update the permission in the map if it is true in the permset and false in the map
                        if (!(Boolean)activePermissionMap.get(permissionLabel) && systemPermissionValue) {
                            activePermissionMap.put(permissionLabel, systemPermissionValue);
                        }
                    }
                }
            }
        }

        //set missing permission set state
        if (isPermSetMissing) {
            permissionIssueMap.put('missingPermissionSets', new List<String> {
                constants.PACKAGED_PERMISSION_SET_NAME
            });
        } else {
            permissionIssueMap.put('isPermSetAssigned', true);
        }

        return new Map<String, Object> {
           'activePermissionMap' => activePermissionMap,
           'permissionIssueMap' => permissionIssueMap
        };
    }

    private static Stripe_Connection__c getStripeConnection() {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

        // checks if we have a org-level API key generated, if we don't we should generate one
        if (stripeConnectRec.Id == null && utilities.isUsingV1Middleware()) {
            throw new ConnectionNotInitializedException();
        }

        return stripeConnectRec;
    }

    //Creates a test metadata record in test context
    @TestVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata {
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = utilities.getStripeConnectionData();
            return setupConfigMetadata;
        } set; }
}