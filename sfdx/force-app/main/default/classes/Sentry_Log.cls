/**
 * Created by jmather on 2019-08-08.
 */

global virtual with sharing class Sentry_Log {
    public static Boolean debugSentry = false;
    public static List<String> logLevels = new List<String> {
            Sentry_LogMessage.LEVEL_TRACE,
            Sentry_LogMessage.LEVEL_DEBUG,
            Sentry_LogMessage.LEVEL_INFO,
            Sentry_LogMessage.LEVEL_WARN,
            Sentry_LogMessage.LEVEL_ERROR
    };

    global static List<Sentry_LogMessage> messages = new List<Sentry_LogMessage>();

    global static void logSentry(String message) {
        System.debug('message: ' + message);
        log(Sentry_LogMessage.LEVEL_SENTRY_DEBUG, message);
    }

    global static void logSentry(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_SENTRY_DEBUG, message, data);
    }

    global static void logTrace(String message) {
        log(Sentry_LogMessage.LEVEL_TRACE, message);
    }

    global static void logTrace(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_TRACE, message, data);
    }

    global static void logDebug(String message) {
        log(Sentry_LogMessage.LEVEL_DEBUG, message);
    }

    global static void logDebug(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_DEBUG, message, data);
    }

    global static void logInfo(String message) {
        log(Sentry_LogMessage.LEVEL_INFO, message);
    }

    global static void logInfo(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_INFO, message, data);
    }

    global static void logWarn(String message) {
        log(Sentry_LogMessage.LEVEL_WARN, message);
    }

    global static void logWarn(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_WARN, message, data);
    }

    global static void logError(String message) {
        log(Sentry_LogMessage.LEVEL_ERROR, message);
    }

    global static void logError(String message, Map<String, String> data) {
        log(Sentry_LogMessage.LEVEL_ERROR, message, data);
    }

    global static void log(String level, String message) {
        addLogMessage(new Sentry_LogMessage(level, message));
    }

    global static void log(String level, String message, Map<String, String> data) {
        addLogMessage(new Sentry_LogMessage(level, message, data));
    }

    global static void addLogMessage(Sentry_LogMessage message) {
        if (message.level == Sentry_LogMessage.LEVEL_SENTRY_DEBUG && debugSentry) {
            showDebug(message);
        }

        if (logLevels.contains(message.level)) {
            showDebug(message);
            messages.add(message);
        }
    }
    
    global static void logException(Exception e) {
        if (e.getCause() != null) {
            logException(e);
        }

        Map<String, String> data = new Map<String, String>();

        data.put('message', e.getMessage());
        data.put('line_number', String.valueOf(e.getLineNumber()));
        data.put('stack_trace', e.getStackTraceString());
        data.put('type_name', e.getTypeName());

        logError(e.getMessage(), data);
    }

    public static void showDebug(Sentry_LogMessage message) {
        System.debug(message.message);

        if (message.data.isEmpty()) {
            return;
        }

        for (String key : message.data.keySet()) {
            System.debug(key + ' = ' + JSON.serializePretty(message.data.get(key)));
        }
    }
}