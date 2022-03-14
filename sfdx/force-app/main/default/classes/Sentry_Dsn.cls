/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-03-25.
 */

public with sharing class Sentry_Dsn {
    public String privateKey;
    public String publicKey;
    public String projectId;
    public Url url;

    public Sentry_Dsn(String dsn) {
        url = new Url(dsn);
        this.publicKey = getPublicKey(url);
        this.privateKey = getPrivateKey(url);
        this.projectId = getProjectId(url);
    }

    String getProjectId(Url url) {
        String[] pieces = url.getPath().split('/');

        if (pieces.size() == 0) {
            return null;
        }
        return pieces[pieces.size() - 1];
    }

    String getPublicKey(Url url) {
        String userInfo = url.getUserInfo();
        if (String.isBlank(userInfo)) {
            return null;
        }

        if (userInfo.contains(':')) {
            return url.getUserInfo().split(':')[0];
        }
        return userInfo;
    }

    String getPrivateKey(Url url) {
        String[] pieces = url.getUserInfo().split(':');

        if (pieces.size() == 2) {
            return pieces[1];
        }
        return null;
    }
}