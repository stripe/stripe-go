/**
 * Created by jmather-c on 8/18/23.
 */

public with sharing class StripeOAuthState {

    public class OAuthInvalidStateException extends Sentry_Exception {}

    private static Map<String, String> JWTHeader = new Map<String, String>{
        'alg' => 'HS256',
        'typ' => 'JWT'
    };
    private static String JWTAlgo = 'hmacSHA256';


    public Map<String, Object> data;

    public String org_id {
        get { return (String) data.get('org_id'); }
        set { data.put('org_id', value); }
    }

    public String csrf {
        get { return (String) data.get('csrf'); }
        set { data.put('csrf', value); }
    }

    public String user_id {
        get { return (String) data.get('user_id'); }
        set { data.put('user_id', value); }
    }

    public String namespace {
        get { return (String) data.get('salesforce_namespace'); }
        set { data.put('salesforce_namespace', value); }
    }

    public String instance_type {
        get { return (String) data.get('salesforce_instance_type'); }
        set { data.put('salesforce_instance_type', value); }
    }

    public String primary_stripe_account_id {
        get { return (String) data.get('primary_stripe_account_id'); }
        set { data.put('primary_stripe_account_id', value); }
    }

    public String salesforce_instance_subdomain {
        get { return (String) data.get('salesforce_instance_subdomain'); }
        set { data.put('salesforce_instance_subdomain', value); }
    }


    public static StripeOAuthState create() {
        Stripe_Connection__c conn = utilities.getStripeConnection();
        Url instanceUrl = new Url(Url.getOrgDomainUrl().toExternalForm());
        String subdomain = instanceUrl.getHost().split('\\.')[0];
        return new StripeOAuthState(
                UserInfo.getOrganizationId(),
                EncodingUtil.convertToHex(Crypto.generateAesKey(256)),
                subdomain,
                conn.Primary_Stripe_Account_Id__c);
    }

    public static StripeOAuthState from_encrypted_state(String encrypted_state) {
        Map<String, Object> data = StripeOAuthState.decrypt(encrypted_state);
        if (data == null) {
            OAuthInvalidStateException ex = (OAuthInvalidStateException) Sentry_ExceptionFactory.build(OAuthInvalidStateException.class);
            ex.setMessage('Invalid OAuth State');
            ex.context.put('encrypted_state', encrypted_state);

            throw ex;
        }
        return new StripeOAuthState(data);
    }

    private StripeOAuthState() {
    }

    public StripeOAuthState(Map<String, Object> data) {
        this.data = data;
    }

    public StripeOAuthState(String org_id, String csrf, String salesforce_instance_subdomain, String primary_stripe_account_id) {
        this.data = new Map<String, Object>();
        this.org_id = org_id;
        this.csrf = csrf;
        this.namespace = constants.NAMESPACE;
        this.instance_type = Sentry_Environment.getInstanceType().name();
        this.salesforce_instance_subdomain = salesforce_instance_subdomain;
        this.primary_stripe_account_id = primary_stripe_account_id;
    }

    override public String toString() {
        return StripeOAuthState.encrypt(data);
    }

    public String toJSON() {
        return JSON.serialize(data);
    }

    // jwt encoding reffed from
    // https://help.salesforce.com/s/articleView?id=sfdo.EL_Create_a_JWT.htm&type=5

    public static String toSanitizedBase64(Blob b) {
        String base64 = EncodingUtil.base64Encode(b);
        base64 = base64.replaceAll('=', '');
        base64 = base64.replaceAll('\\+', '-');
        base64 = base64.replaceAll('\\/', '_');

        return base64;
    }

    public static Blob fromSanitizedBase64(String s) {
        String base64 = s.replaceAll('-', '+');
        base64 = base64.replaceAll('_', '/');
        // '=' * (4 - b64.length % 4)
        base64 = base64 + '='.repeat(4 - Math.mod(base64.length(), 4));
        return EncodingUtil.base64Decode(base64);
    }

    public static String encrypt(Map<String, Object> state) {
        try {
            String header_b64 = toSanitizedBase64(Blob.valueOf(JSON.serialize(JWTHeader)));
            String state_b64 = toSanitizedBase64(Blob.valueOf(JSON.serialize(state)));
            Blob hash = Crypto.generateMac(JWTAlgo, Blob.valueOf(state_b64), utilities.getOAuthStateSigningKey());
            String hash_b64 = toSanitizedBase64(hash);
            return header_b64 + '.' + state_b64 + '.' + hash_b64;
        } catch (Exception e) {
            errorLogger.create('StripeOAuthState.encrypt failed', e);
            return null;
        }
    }

    public static Map<String, Object> decrypt(String state) {
        try {
            List<String> parts = state.split('\\.');
            String input = parts[1];
            Blob hash = fromSanitizedBase64(parts[2]);
            Boolean verified = Crypto.verifyHmac(JWTAlgo, Blob.valueOf(input), utilities.getOAuthStateSigningKey(), hash);
            if (!verified) {
                errorLogger.create('StripeOAuthState.decrypt failed', 'hash verification failed', '120', 'StripeOAuthState.decrypt');
                return null;
            }
            return (Map<String, Object>) JSON.deserializeUntyped(fromSanitizedBase64(parts[1]).toString());
        } catch (Exception e) {
            errorLogger.create('StripeOAuthState.decrypt failed', e);
            return null;
        }
    }
}