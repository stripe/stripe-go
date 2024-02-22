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

    // maps to line ~89 in state_encryption_algo.rb
    private static final String FIELD_SALESFORCE_ACCOUNT_ID = 'org_id';
    private static final String FIELD_SALESFORCE_NAMESPACE = 'sf_ns';
    private static final String FIELD_SALESFORCE_INSTANCE_TYPE = 'sf_it';
    private static final String FIELD_SALESFORCE_INSTANCE_SUBDOMAIN = 'sf_sd';
    private static final String FIELD_USER_ID = 'uid';
    private static final String FIELD_CASC_ID = 'csac';
    private static final String FIELD_CSRF = 'csrf';
    private static final String FIELD_STRIPE_ACCOUNT_ID = 'st_id';
    private static final String FIELD_STRIPE_ACCOUNT_LIVEMODE = 'st_lm';
    private static final String FIELD_STRIPE_ACCOUNT_NAME = 'st_na';
    private static final String FIELD_PRIMARY_STRIPE_ACCOUNT_ID = 'pst_id';
    private static final String FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE = 'pst_lm';



    public Map<String, Object> data;

    public String org_id {
        get { return (String) data.get(FIELD_SALESFORCE_ACCOUNT_ID); }
        set { data.put(FIELD_SALESFORCE_ACCOUNT_ID, value); }
    }

    public String namespace {
        get { return (String) data.get(FIELD_SALESFORCE_NAMESPACE); }
        set { data.put(FIELD_SALESFORCE_NAMESPACE, value); }
    }

    public String instance_type {
        get { return (String) data.get(FIELD_SALESFORCE_INSTANCE_TYPE); }
        set { data.put(FIELD_SALESFORCE_INSTANCE_TYPE, value); }
    }

    public String instance_subdomain {
        get { return (String) data.get(FIELD_SALESFORCE_INSTANCE_SUBDOMAIN); }
        set { data.put(FIELD_SALESFORCE_INSTANCE_SUBDOMAIN, value); }
    }

    public String user_id {
        get { return (String) data.get(FIELD_USER_ID); }
        set { data.put(FIELD_USER_ID, value); }
    }

    public String casc_id {
        get { return (String) data.get(FIELD_CASC_ID); }
        set { data.put(FIELD_CASC_ID, value); }
    }

    public String csrf {
        get { return (String) data.get(FIELD_CSRF); }
        set { data.put(FIELD_CSRF, value); }
    }

    public String stripe_account_id {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_ID); }
        set { data.put(FIELD_STRIPE_ACCOUNT_ID, value); }
    }

    public String stripe_account_livemode {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_LIVEMODE); }
        set { data.put(FIELD_STRIPE_ACCOUNT_LIVEMODE, value); }
    }

    public String stripe_account_name {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_NAME); }
        set { data.put(FIELD_STRIPE_ACCOUNT_NAME, value); }
    }

    public String primary_stripe_account_id {
        get { return (String) data.get(FIELD_PRIMARY_STRIPE_ACCOUNT_ID); }
        set { data.put(FIELD_PRIMARY_STRIPE_ACCOUNT_ID, value); }
    }

    public String primary_stripe_account_livemode {
        get { return (String) data.get(FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE); }
        set { data.put(FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE, value); }
    }

    public static StripeOAuthState create() {
        Stripe_Connection__c conn = utilities.getStripeConnection();
        Url instanceUrl = new Url(Url.getOrgDomainUrl().toExternalForm());
        String subdomain = instanceUrl.getHost().split('\\.')[0];
        return new StripeOAuthState(
                UserInfo.getOrganizationId(),
                EncodingUtil.convertToHex(Crypto.generateAesKey(256)),
                subdomain,
                conn.Primary_Stripe_Account_Id__c,
                conn.Primary_Stripe_Account_Livemode__c);
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

    public StripeOAuthState(String org_id, String csrf, String salesforce_instance_subdomain, String primary_stripe_account_id, String primary_stripe_account_livemode) {
        this.data = new Map<String, Object>();
        this.org_id = org_id;
        this.csrf = csrf;
        this.namespace = constants.NAMESPACE;
        this.instance_type = Sentry_Environment.getInstanceType().name();
        this.instance_subdomain = salesforce_instance_subdomain;
        this.primary_stripe_account_id = primary_stripe_account_id;
        this.primary_stripe_account_livemode = primary_stripe_account_livemode;
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