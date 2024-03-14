/**
 * Created by jmather-c on 9/1/23.
 */

 public with sharing class UriTargetsV2 {
    public static final String RUBY_SERVICE_BASE_URI = 'https://salesforce-connector.stripe.com';
    public static final String RUBY_SERVICE_LOCAL_BASE_URI = 'http://localhost:3100';
    public static final String RUBY_SERVICE_CONFIG_URI = '/api/configuration';
    public static final String RUBY_SERVICE_CONNECTION_STATUS_URI = '/api/configuration/connection_statuses';
    public static final String RUBY_MULTI_ACCT_SERVICE_CONNECTION_STATUS_URI = '/api/configuration/multi_connection_statuses';
    public static final String RUBY_SERVICE_TRANSLATE_URI = '/api/translate';
    public static final String RUBY_SERVICE_TRANSLATE_ALL_URI = '/api/translate_all';
    public static final String RUBY_SERVICE_SALESFORCE_URI = '/api/auth/salesforce';
    public static final String RUBY_SERVICE_STRIPE_URI = '/api/auth/stripe';
    public static final String RUBY_SERVICE_POST_INSTALL_URI = '/api/auth/salesforce/postinstall';
    public static final String RUBY_SERVICE_OPENAPI_URI = '/openapi.json';

    static final Debug_Helper debugger = new Debug_Helper();

    public static String getTranslateAllEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_TRANSLATE_ALL_URI;
    }

    public static String getTranslateEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_TRANSLATE_URI;
    }

    public static String getConfigEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_CONFIG_URI;
    }

    public static String getConnectionStatusesEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_CONNECTION_STATUS_URI;
    }

    public static String getMultiAccountConnectionStatusesEndpoint() {
        return getPlatformBaseUri() + RUBY_MULTI_ACCT_SERVICE_CONNECTION_STATUS_URI;
    }

    public static String getOpenApiSpecEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_OPENAPI_URI;
    }

    public static String getPostInstallEndpoint() {
        return getPlatformBaseUri() + RUBY_SERVICE_POST_INSTALL_URI;
    }

    public static String getAccountSetDefaultEndpoint(String accountId, Boolean isLive) {
        String liveMode = isLive ? 'live' : 'test';
        return getAccountSetDefaultEndpoint(accountId, liveMode);
    }

    public static String getAccountSetDefaultEndpoint(String accountId, String liveMode) {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_ACCOUNTS_URI + accountId + '/' + liveMode + constants.RUBY_SERVICE_ACCOUNTS_SET_DEFAULT_URI;
    }

    public static String getAccountEndpoint(String accountId, Boolean isLive) {
        String liveMode = isLive ? 'live' : 'test';
        return getAccountEndpoint(accountId, liveMode);
    }

    public static String getAccountEndpoint(String accountId, String liveMode) {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_ACCOUNTS_URI + accountId + '/' + liveMode;
    }

    public static String getStripeAuthorizationUri(Boolean liveMode, StripeOAuthState state) {
        String params = '?salesforceNamespace=' + constants.NAMESPACE;
        params += '&state=' + state.toString();
        debugger.debug('UriTargetsV2.getStripeAuthorizationUri state', state.data);

        state.auth_live_mode = liveMode;

        return getAuthorizationBaseUri() + RUBY_SERVICE_STRIPE_URI + params;
    }

    public static String getIntegrationUserAuthorizationUrl(StripeOAuthState state, Boolean forceLogin) {
        String params = '?salesforceNamespace=' + constants.NAMESPACE;
        if (forceLogin) {
            params += '&prompt=login%20consent';
        }

        params += '&state=' + state.toString();
        debugger.debug('UriTargetsV2.getIntegrationUserAuthorizationUrl state', state.data);

        return getAuthorizationBaseUri() + RUBY_SERVICE_SALESFORCE_URI + params;
    }

    public static String getAuthorizationBaseUri() {
        Setup_Connection_Data__mdt config = utilities.getConnectionData();

        if (config != null && config.Local_Authorization_Target__c) {
            return RUBY_SERVICE_LOCAL_BASE_URI;
        }

        return RUBY_SERVICE_BASE_URI;
    }

    public static String getPlatformBaseUri() {
        Setup_Connection_Data__mdt config = utilities.getConnectionData();

        if (config != null && String.isNotEmpty(config.Platform_Target__c)) {
            return config.Platform_Target__c;
        }

        return RUBY_SERVICE_BASE_URI;
    }

}