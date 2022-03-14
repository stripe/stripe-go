/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-06-25.
 */

public with sharing class Sentry_Error_Handler {
    public void run(List<Sentry_Error__e> listOfSentryErrors) {
        List<String> errors = new List<String>();

        for (Sentry_Error__e e : listOfSentryErrors) {
            Sentry_Log.logSentry('[Sentry_Error_Handler.run] System event: ' + JSON.serializePretty(e));
            Sentry_Log.logSentry('[Sentry_Error_Handler.run] Sentry event: ' + e.Content__c);
            errors.add(e.Content__c);
        }
        if (errors.isEmpty() || !Sentry_Config.canSendToSentry()) {
            if (Sentry_Config.canSendToSentry()) {
                Sentry_Log.logSentry('[Sentry_Error_Handler.run] no errors, exiting.');
            } else {
                Sentry_Log.logSentry('[Sentry_Error_Handler.run] sending to sentry is disabled. Not sending ' + errors.size() + ' errors.');
            }
            return;
        }
        Sentry_Log.logSentry('[Sentry_Error_Handler.run] calling sendErrors with ' + errors.size() + ' errors.');
        sendErrors(errors);
    }

    public static void sendErrors(List<String> errors) {
        for (String error : errors) {
            Sentry_Client.sendEventToSentry(error);
        }
    }
}