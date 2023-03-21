/**
 * Created by jmather-c on 1/4/23.
 */

public with sharing class Debug_Helper {
    // This should always be null going to production.
    // We leave this configurable for developers to expose debug logs during development.
    private Boolean enableDebug = false;

    @TestVisible
    private OutputService outputService = new DefaultOutputService();
    @TestVisible
    private static final Integer lengthLimit = 100;

    // for testability
    public interface OutputService { void debug(String str); }
    public class DefaultOutputService implements OutputService { public void debug(String str) { System.debug(str); } }
    public class DebugHelperException extends Sentry_Exception {}

    // empty constructor to use the default above to make it easy to flip for debugging.
    public Debug_Helper() { }

    public Debug_Helper(Boolean enableDebug) {
        this.enableDebug = enableDebug;
    }

    public Debug_Helper(Boolean enableDebug, OutputService outputService) {
        this.enableDebug = enableDebug;
        if (outputService == null) {
            DebugHelperException e = (DebugHelperException) Sentry_ExceptionFactory.build(DebugHelperException.class);
            e.setMessage('Instantiated Debug_Helper with null ouputService');
            throw e;
        } else {
            this.outputService = outputService;
        }
    }

    public void debug(String source, Object data) {
        debug(source, data.toString());
    }

    public void debug(String source, String message) {
        if (enableDebug == false) {
            return;
        }

        if (message.length() < lengthLimit) {
            outputService.debug('[' + source + '] ' + message);
            return;
        }

        while (message.length() > 0) {
//            Integer nextLen = message.length() < lengthLimit ? message.length() : lengthLimit;
            if (message.length() > lengthLimit) {
                String nextMsg = message.substring(0, lengthLimit);
                message = message.substring(lengthLimit);
                outputService.debug('[' + source + '] ' + nextMsg.trim());
            } else {
                outputService.debug('[' + source + '] ' + message.trim());
                message = '';
            }
        }
    }
}