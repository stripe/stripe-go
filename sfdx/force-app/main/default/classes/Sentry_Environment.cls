/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-06-25.
 */

public with sharing class Sentry_Environment {
    public enum InstanceType { PRODUCTION, TRIAL, SANDBOX, SCRATCH_ORG }

    public static String getInstanceLabel() {
        String orgUrl = Url.getOrgDomainUrl().toExternalForm();
        Sentry_Log.logSentry('[Sentry_Environment.getInstanceLabel] orgUrl: ' + orgUrl);
        List<String> protoAndRemainder = orgUrl.split('://');
        Sentry_Log.logSentry('[Sentry_Environment.getInstanceLabel] protoAndRemainder: ' + protoAndRemainder);
        Sentry_Log.logSentry('[Sentry_Environment.getInstanceLabel] target: ' + protoAndRemainder[1]);

        List<String> pieces = protoAndRemainder[1].split('\\.');
        Sentry_Log.logSentry('[Sentry_Environment.getInstanceLabel] pieces: ' + pieces);
        String url_fragment = pieces[0];

        if (url_fragment.contains('--') == false) {
            if (Schema.getGlobalDescribe().containsKey('ActiveScratchOrg')) {
                return 'production';
            }
            return url_fragment.toLowerCase();
        }
        return url_fragment.split('--')[0].toLowerCase();
    }

    public static InstanceType getInstanceType() {
        Organization org = [
            SELECT InstanceName, IsSandbox, TrialExpirationDate
            FROM Organization
            LIMIT 1
        ];

        if (org.IsSandbox == true && org.TrialExpirationDate == null) {
            return InstanceType.SANDBOX;
        } else if (org.IsSandbox == true) {
            return InstanceType.SCRATCH_ORG;
        } else if (org.TrialExpirationDate != null) {
            return InstanceType.TRIAL;
        } else {
            return InstanceType.PRODUCTION;
        }
    }
}