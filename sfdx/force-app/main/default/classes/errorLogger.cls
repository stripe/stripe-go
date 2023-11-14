/*** Strike Error Logger v2.0.0 ***/
public virtual with sharing class errorLogger {
    @TestVisible
    static Debug_Helper debugger = new Debug_Helper();

    public static void create(String actionTaken, String message, String lineNumber, String dump) {

        Map<String, String> err = new Map<String, String>{
            'dump' => dump,
            'message' => message,
            'line_number' => lineNumber,
            'action_taken' => actionTaken
        };

        debugger.debug('errorLogger', err);
    }
    public static void create(String errorName, Exception excep) {
        debugger.debug(errorName, excep);
        if(!Test.isRunningTest() && !utilities.isDeveloperModeEnabled())Sentry.record(excep);
    }
}
