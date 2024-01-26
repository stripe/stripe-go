/**
 * Created by jmather-c on 9/1/23.
 */

public with sharing class UriTargets {
    static final Debug_Helper debugger = new Debug_Helper();

    public static String getTranslateAllEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_TRANSLATE_ALL_URI;
    }

    public static String getTranslateEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_TRANSLATE_URI;
    }

    public static String getConfigEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_CONFIG_URI;
    }

    public static String getConnectionStatusesEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_CONNECTION_STATUS_URI;
    }

    public static String getMultiAccountConnectionStatusesEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_MULTI_ACCT_SERVICE_CONNECTION_STATUS_URI;
    }

    public static String getOpenApiSpecEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_OPENAPI_URI;
    }

    public static String getPostInstallEndpoint() {
        return getPlatformBaseUri() + constants.RUBY_SERVICE_POST_INSTALL_URI;
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
        debugger.debug('UriTargets.getStripeAuthorizationUri state', state.data);

        if (liveMode) {
            return getAuthorizationBaseUri() + constants.RUBY_SERVICE_STRIPE_LIVE_URI + params;
        }

        return getAuthorizationBaseUri() + constants.RUBY_SERVICE_STRIPE_TEST_URI + params;
    }

    public static String getIntegrationUserAuthorizationUrl(StripeOAuthState state, Boolean forceLogin) {
        String params = '?salesforceNamespace=' + constants.NAMESPACE;
        if (forceLogin) {
            params += '&prompt=login%20consent';
        }

        params += '&state=' + state.toString();
        debugger.debug('UriTargets.getIntegrationUserAuthorizationUrl state', state.data);

        if (constants.IS_SANDBOX) {
            return getAuthorizationBaseUri() + constants.RUBY_SERVICE_SALESFORCE_TEST_URI + params;
        }

        return getAuthorizationBaseUri() + constants.RUBY_SERVICE_SALESFORCE_PROD_URI + params;
    }

    public static String getAuthorizationBaseUri() {
        Setup_Connection_Data__mdt config = utilities.getConnectionData();

        if (config != null && config.Local_Authorization_Target__c) {
            return constants.RUBY_SERVICE_LOCAL_BASE_URI;
        }

        return constants.RUBY_SERVICE_BASE_URI;
    }

    public static String getPlatformBaseUri() {
        Setup_Connection_Data__mdt config = utilities.getConnectionData();

        if (config != null && String.isNotEmpty(config.Platform_Target__c)) {
            return config.Platform_Target__c;
        }

        return constants.RUBY_SERVICE_BASE_URI;
    }

}