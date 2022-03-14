/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-06-25.
 */

global virtual with sharing class Sentry_Exception extends Exception {
    global Map<String, Object> context = new Map<String, Object>();
    private List<String> stackTrace;

    global void setStackTrace(List<String> lines) {
        stackTrace = lines;
    }

    global String getCustomStackTraceAsString() {
        return String.join(getCustomStackTrace(), '\n');
    }

    global List<String> getCustomStackTrace() {
        if (stackTrace != null) {
            return stackTrace;
        }
        return new List<String>();
    }
}