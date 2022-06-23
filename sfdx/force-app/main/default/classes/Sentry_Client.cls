/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-03-18.
 * Modified by Stripe on 2022-02-10 (added error handeling)
 */

public with sharing class Sentry_Client {
    private static Sentry_Dsn dsn = Sentry_Config.getDsn();

    @future(callout=true)
    public static void sendEventToSentry(String serializedEvent) {
        HttpRequest req = new HttpRequest();
        req.setEndpoint(Sentry_Config.STRIPE_CONNECTOR_SENTRY_BASE_URL + dsn.projectId + '/store/');
        req.setHeader('content-type', 'application/json');
        req.setHeader('Content-Encoding', 'User-Agent: sentry-ruby/11');
        addRequestHeaders(req);
        Map<String,Object> deserializedEvent = (Map<String,Object>)JSON.deserializeUntyped(serializedEvent);
        req.setHeader('event_id', (String)deserializedEvent.get('event_id'));

        req.setMethod('POST');
        req.setBody(serializedEvent);

        Sentry_Log.logSentry('[Sentry_Client.sendEventToSentry] dsn: ' + dsn);
        Sentry_Log.logSentry('[Sentry_Client.sendEventToSentry] body: ' + serializedEvent);
        Sentry_Log.logSentry('[Sentry_Client.sendEventToSentry] request: ' + req);

        HttpResponse resp = new Http().send(req);

        Sentry_Log.logSentry('[Sentry_Client.sendEventToSentry] response: ' + resp);

        if(resp.getStatusCode() != 200) {
             Map<String,Object> responseBody = (Map<String,Object>)JSON.deserializeUntyped(resp.getBody());
             String causes;
             String detail;
             if(responseBody.get('causes') != null) {
                 causes = String.valueOf(responseBody.get('causes'));
             } else {
                causes = 'no cause(s) retruned';
             }
             if(responseBody.get('detail') != null) {
                detail = String.valueOf(responseBody.get('detail'));
            } else {
                detail = 'no detail(s) retruned';
            }

            // TODO what value does logging this give us? This should probably be removed.
            errorLogger.create('SendToSentryError', String.valueOf(resp.getStatusCode()), 'Detail: '+detail, 'Causes: '+causes);
        }
    }

    private static void addRequestHeaders(HttpRequest req) {
        Map<String, String> headers = new Map<String, String> {
            'sentry_version' => '7',
            'sentry_sdk' => 'sentry.salesforce',
            'sentry_key' => dsn.publicKey
        };

        if (dsn.privateKey != null) {
            headers.put('sentry_secret', dsn.privateKey);
        }

        String[] headerSets = new List<String>();
        for (String key : headers.keySet()) {
            headerSets.add(key + '=' + headers.get(key));
        }

        String header = 'Sentry ' + String.join(headerSets, ', ');
        Sentry_Log.logSentry('[Sentry_Client.addRequestHeaders] header: ' + header);
        req.setHeader('X-Sentry-Auth', header);
    }
}