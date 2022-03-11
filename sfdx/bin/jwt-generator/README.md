# salesforce-jwt-generator

Using the JWT OAuth Flow requires you to:

- Generate a public/private key pair
- Create a Connected App
- Generate a JWT signing it with the private key
- Exchange the JWT for an access token
- Use the access token as a Bearer token

Below there is a section for each of the above steps.

## Generate public/private key pair

Generate a public/private keypair using openssl and fill in the required info when you generate the certificate.

```
openssl req -newkey rsa:2048 -nodes -keyout private_key.pem -x509 -days 365 -out certificate.pem
openssl x509 -outform der -in certificate.pem -out public_key.der
openssl x509 -in certificate.pem -pubkey > public_key.pem
```

## Create a Connected App

In Salesforce create a Connected App through the App Manager in Setup and upload the public key (public_key.cer from the above steps) to the app. Be sure to select the offline_access scope as well as other required scopes. For testing the `openid` scope is always good. Save the Connected App and make a note of the consumer key (client_id).

**Please note**: If you plan on using the JWT to issue an access token the user must have authorized the Connected App _or_ it must be marked admin approved and the Connected App added to the user profile or assigned with a permission set. You must also ensure that the `refresh_token, offline_access` scope gets assigned.

**Please note**: If you plan on using the JWT to create an access token that may be used to open the org using `frontdoor.jsp` ensure that the Connected App assigns the `web` scope.

## Generate a JWT

Use the node.js app in this repo to create a JWT. Once you've cloned the repo create a `.env` file with the following
keys:

- SUBJECT (the username of the user to impersonate)
- CLIENT_ID (the consumer key (client_id) of the Connected App you created)
- **optional** AUDIENCE (https://login.salesforce.com or https://test.salesforce.com as appropriate)
- PATH_PRIVATE_KEY (path to the pem-file with the private key (`private_key.pem`from above))
- PATH_PUBLIC_KEY (path to the pem-file with the public key (`public_key.pem`from above))

**Please note:** The JWT expires in 5 minutes so be quick about exchanging it for an access token!

```
npm install
npm start
```

## Exchange the JWT for an access token

Using Postman or similar post to the OAuth token-endpoint of Salesforce specifying a `grant_type`-parameter of `urn:ietf:params:oauth:grant-type:jwt-bearer` and specify the JWT in the `assertion` parameter.

```
POST /services/oauth2/token HTTP/1.1
Host: login.salesforce.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 731
Connection: close

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzQzNDQzNDcsImV4cCI6MTU3NDM0NDY0NywiYXVkIjoiaHR0cHM6Ly9sb2dpbi5zYWxlc2ZvcmNlLmNvbSIsImlzcyI6Inh5ejEyMyIsInN1YiI6Impkb2VAZm9vLmRlbW8ifQ.jpEPDj_9DEhzvCUGwvEefZvd63IPvtBAZCSJ_-RJ-nlAqktbwoMoCrUFb_S1u0xRuWKBhwY7Mg58claQN2UTyxhjjDYzchIRsTbrRB-KNxzd6J_ew0of8IpB8NWN_1245KuO9clfo_Yoq8wwZUTBSSt55jh4-TyjpRg4UjIikus76GZL0xvWBWfGD2zxgshOgWMk-sewJE5REGP8FPz-SqV6L_o_ua82FbBvpchwRavFmK-y0E8kDNtoOhJyW-P8jvTMfZog1hslqPQBF6-z9EBUGFb482DrEh1vspwIGV-ioLHTmJo5kBhsJXrDG6hwODVVe2G_1eSl-52k4gOvTw
```

## Use frontdoor.jsp

Instead of using Postman as outlined about you can also invoke `npm run access_token` from the command line to get an access token directly. That will also create a link to open the org directly using the `frontdoor.jsp` approach. In most shells you can Cmd-click or similar to open the link directly. If need be you can append `&retURL=<relative URL>` (i.e. `&retURL=/lightning/r/Account/00109000007fVDSAA2/view`) to open a specific record or page directly.

**Please note**: If using `frontdoor.jsp` for a community you should set the `AUDIENCE` in `.env` to the community url (i.e. `https://isvsi-14ddd2ecd93-656fd6e55fe.force.com/customer`) and the `SUBJECT` to a community username. For a community using the retURL-parameter is also required but that is automatically attempted extracted from the community url by the script. The resulting frontdoor-url will be something like `https://isvsi-14ddd2ecd93-656fd6e55fe.force.com/customer/secur/frontdoor.jsp?sid=00D09000006JtKK!ARQADDdyuTBanM0oB71MT.qjbDirRlPDTFrZ2UAYJsvfZjBDD0k36NJoNeCwaNUrWx54vjSG3c10UKaQTB8RGcitOUrJYzKR&retURL=/customer/s`.

## Use the access token as a Bearer token

```
GET /id/00D3X000002KFdlUAG/0053X00000AdY37QAF HTTP/1.1
Host: login.salesforce.com
Authorization: Bearer 00D3X0000...zLwRJ3AzGgXa
Connection: close
```

## Using with Azure

Using JWT's are also possible with Micrsoft Azure using the client_credentials flow specifying a JWT instead of a client_secret. To do this you must have uploaded the public key to the App Registation in Azure AD and you must specify a certificate thumbprint in the `x5t` key in the JWT header. The thumbprint is specified using `CERTIFICATE_THUMBPRINT` in the `.env` file and may be generated using `openssl` using a somewhat special process as it is not simply the sha1 hash as described in the documentation (https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials).

The hash to specify may be generated as follows as per https://stackoverflow.com/a/52625165

```
echo $(openssl x509 -in your.cert.pem -fingerprint -noout) | sed 's/SHA1 Fingerprint=//g' | sed 's/://g' | xxd -r -ps | base64
```

Using the above exchange the JWT for an access token using a POST like below:

```
POST /b34deb2b-232f-4322-af4d-c732d5d885d0/oauth2/v2.0/token HTTP/1.1
Host: login.microsoftonline.com
Content-Type: application/x-www-form-urlencoded
Connection: close

client_id=43d816a5-0cf4-888a-f8a0-7c88e6fc254e
&grant_type=client_credentials
&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer
&client_assertion=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsIng1dCI6IjRyRWxzREZUeXNyYktoQjB6VHNyUk5TeFQ2cz0ifQ.eyJpYXQiOjE1NzQzNDgzMzEsImV4cCI6MTU3NDM0ODYzMSwiYXVkIjoiaHR0cHM6Ly9sb2dpbi5taWNyb3NvZnRvbmxpbmUuY29tL2IzNGRlYjJiLTIzMmYtNDMyMi1hZjRkLWM3MzJkNWQ4ODVkMC9vYXV0aDIvdjIuMC90b2tlbiIsImlzcyI6IjQzZDgxNmE1LTBjZjQtODg4YS1mOGEwLTdjODhlNmZjMjU0ZSIsInN1YiI6IjQzZDgxNmE1LTBjZjQtODg4YS1mOGEwLTdjODhlNmZjMjU0ZSJ9.CV7YCZ4Oak-g8b0nBTYweZDSp6lvYH48US02dWMV1Nie7wkYaqmTlTSRD5HGH5Jt5xGc9g0mNnX3p13m0AcbXTmZJ0MOfjnrAPvXJxtXEMEQXnHhIt_IExQ7NTNQWXvLRmlHydDFHMd-ss9QQt2BTwqPl6Lqlt4mgT9RfSd-6W3pTLyFsB21-WSCH1j7ykR9j5A5wfTpBtj_h4-kz3gq6VlFTVg2Mph4KlNYkssGSRd74qY1_olKeMFbI6Wft4Ige79D1qIcbC9DFITKMoEaSFSWS1Pg9pxkHoyOPFihO51SCXzFNRLGvA1nEQFRkV0raUokdWmoi7u_S_mUZe3qYQ&tenant=b34deb2b-232f-4322-af4d-c732d5d885d0&scope=https%3A%2F%2Fgraph.microsoft.com%2F.default
```
