/**
 * @description       :
 * @author            : jmather
 * @group             :
 * @last modified on  : 02-28-2022
 * @last modified by  : Stripe
 * Modifications Log
 * Ver   Date         Author    Modification
 * 1.0   08-10-2019   jmather   Initial Version
 * 1.2   02-10-2022   Stripe    removed trigger and added logic to call handler directly
**/
public with sharing class Sentry {
    public static void record(Exception e) {
        Sentry_Event err = convertExceptionToError(e);
        sendEvent(err);
    }

    public static void record(String issueTitle, Exception e, List<Sentry_LogMessage> lExtraMessages) {
        Sentry_Event err = convertExceptionToError(issueTitle, e, lExtraMessages);
        sendEvent(err);
    }

    private static Sentry_Event convertExceptionToError(Exception e) {
        return new Sentry_Event(e);
    }

    private static Sentry_Event convertExceptionToError(String issueTitle, Exception e, List<Sentry_LogMessage> lExtraMessages) {
        return new Sentry_Event(issueTitle, e, lExtraMessages);
    }

    private static void sendEvent(Sentry_Event err) {
        Sentry_Log.logSentry('[Sentry.sendEvent] event: ' + JSON.serializePretty(err));
        Sentry_Log.logSentry('[Sentry.sendEvent] message: ' + JSON.serializePretty(err.toMessage()));

        Sentry_Error__e formattedSentryError = new Sentry_Error__e(Content__c = JSON.serialize(err.toMessage()));
        Sentry_Error_Handler sentryErrorHandler = new Sentry_Error_Handler();
        List<Sentry_Error__e> listOfSentryErrors = new List<Sentry_Error__e>{formattedSentryError};
        sentryErrorHandler.run(listOfSentryErrors);
    }
}