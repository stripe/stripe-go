/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-06-25.
 */

public with sharing class Sentry_Error_Handler {
    public void run(List<Sentry_Error__e> listOfSentryErrors) {

        if (listOfSentryErrors.isEmpty() || !Sentry_Config.canSendToSentry()) {
            if (Sentry_Config.canSendToSentry()) {
                Sentry_Log.logSentry('[Sentry_Error_Handler.run] no errors, exiting.');
            } else {
                Sentry_Log.logSentry('[Sentry_Error_Handler.run] sending to sentry is disabled. Not sending ' + listOfSentryErrors.size() + ' errors.');
            }
            return;
        }

        Sentry_Error__e exceptionContent = listOfSentryErrors[0];
        Sentry_Log.logSentry('[Sentry_Error_Handler.run] System event: ' + JSON.serializePretty(exceptionContent));
        Sentry_Log.logSentry('[Sentry_Error_Handler.run] Sentry event: ' + exceptionContent.Content__c);
        Sentry_Log.logSentry('[Sentry_Error_Handler.run] calling sendErrors with ' + listOfSentryErrors.size() + ' errors.');
        Sentry_Client.sendEventToSentry(exceptionContent.Content__c);
        
    }
}