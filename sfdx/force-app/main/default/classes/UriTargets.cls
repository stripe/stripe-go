/**
 * Created by jmather-c on 9/1/23.
 */

 public with sharing class UriTargets {
    public static String getTranslateAllEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getTranslateAllEndpoint();
        }

        return UriTargetsV2.getTranslateAllEndpoint();
    }

    public static String getTranslateEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getTranslateEndpoint();
        }

        return UriTargetsV2.getTranslateEndpoint();
    }

    public static String getConfigEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getConfigEndpoint();
        }

        return UriTargetsV2.getConfigEndpoint();
    }

    public static String getConnectionStatusesEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getConnectionStatusesEndpoint();
        }

        return UriTargetsV2.getConnectionStatusesEndpoint();
    }

    public static String getMultiAccountConnectionStatusesEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getMultiAccountConnectionStatusesEndpoint();
        }

        return UriTargetsV2.getMultiAccountConnectionStatusesEndpoint();
    }

    public static String getOpenApiSpecEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getOpenApiSpecEndpoint();
        }

        return UriTargetsV2.getOpenApiSpecEndpoint();
    }

    public static String getPostInstallEndpoint() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getPostInstallEndpoint();
        }

        return UriTargetsV2.getPostInstallEndpoint();
    }

    public static String getAccountSetDefaultEndpoint(String accountId, Boolean isLive) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getAccountSetDefaultEndpoint(accountId, isLive);
        }

        return UriTargetsV2.getAccountSetDefaultEndpoint(accountId, isLive);
    }

    public static String getAccountSetDefaultEndpoint(String accountId, String liveMode) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getAccountSetDefaultEndpoint(accountId, liveMode);
        }

        return UriTargetsV2.getAccountSetDefaultEndpoint(accountId, liveMode);
    }

    public static String getAccountEndpoint(String accountId, Boolean isLive) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getAccountEndpoint(accountId, isLive);
        }

        return UriTargetsV2.getAccountEndpoint(accountId, isLive);
    }

    public static String getAccountEndpoint(String accountId, String liveMode) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getAccountEndpoint(accountId, liveMode);
        }

        return UriTargetsV2.getAccountEndpoint(accountId, liveMode);
    }

    public static String getStripeAuthorizationUri(Boolean liveMode, StripeOAuthState state) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getStripeAuthorizationUri(liveMode, state);
        }

        return UriTargetsV2.getStripeAuthorizationUri(liveMode, state);
    }

    public static String getIntegrationUserAuthorizationUrl(StripeOAuthState state, Boolean forceLogin) {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getIntegrationUserAuthorizationUrl(state, forceLogin);
        }

        return UriTargetsV2.getIntegrationUserAuthorizationUrl(state, forceLogin);
    }

    public static String getAuthorizationBaseUri() {
        if (utilities.isUsingV1Middleware()) {
            return UriTargetsV1.getAuthorizationBaseUri();
        }

        return UriTargetsV2.getAuthorizationBaseUri();
    }
}