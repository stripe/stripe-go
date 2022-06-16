import jwt from "jsonwebtoken";
import fetch from "node-fetch";

export const generateJWT = async (args) => {
    // validate
    if (!args.issuer) throw Error("You must supply the issuer key (client ID)");
    if (!args.subject) throw Error("You must supply the subject key (username)");
    if (!args.privateKey) throw Error("You must supply a privateKey for signing");

    // additonal payload to add
    const additionalPayload = args.additionalPayload || {
        "scopes": ["refresh_token", "api"].join(" ")
    };

    // specify main payload
    var signOptions = {
        "issuer": args.issuer,
        "subject": args.subject,
        "audience": args.audience || "https://login.salesforce.com",
        "expiresIn": 5 * 60,
        "algorithm": "RS256",
    };

    // return promise
    return new Promise((resolve, reject) => {
        try {
            // sign token with private key
            const token = jwt.sign(additionalPayload, args.privateKey, signOptions);
            resolve(token);
        } catch (err) {
            reject(err);
        }
    })
}

export const verifyJWT = async (args) => {
    // validate
    if (!args.token) throw Error("You must supply the token key");
    if (!args.issuer) throw Error("You must supply the issuer key (client ID)");
    if (!args.subject) throw Error("You must supply the subject key (username)");
    if (!args.publicKey) throw Error("You must supply a publicKey for verification");

    // verify that it was signed correctly
    const verifyOptions = {
        "issuer": args.issuer,
        "subject": args.subject,
        "audience": args.audience || "https://login.salesforce.com",
        "expiresIn": 5 * 60,
        "algorithm": "RS256"
    };

    return new Promise((resolve, reject) => {
        try {
            const legit = jwt.verify(args.token, args.publicKey, verifyOptions);
            resolve(legit);
        } catch (err) {
            reject(err);
        }
    })
}

export const accessTokenFromJWT = async (args) => {
    // validate
    if (!args.token) throw Error("You must supply the token key");

    const url = args.audience || 'https://login.salesforce.com';
    
    return fetch(`${url}/services/oauth2/token`, {
        "method": "post",
        "headers": {
            "content-type": "application/x-www-form-urlencoded"
        },
        "body": `grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=${args.token}`
    }).then(resp => resp.json()).then(data => {
        if (data.error) throw Error(`Error caught (${data.error})`);
        return data.access_token;
    })
}

import * as fs from "fs";

const issuer = process.env.SF_CONSUMER_KEY;
const subject = process.env.SF_USERNAME;
const audience = process.env.SF_URL;
const privateKey = fs.readFileSync(process.env.SF_JWT_PRIVATE_KEY_PATH, 'utf8');

generateJWT({
    issuer,
    subject,
    privateKey,
    audience
}).then(token => {
    return accessTokenFromJWT({
        token,
        audience
    })
}).then(accessToken => {
    console.log(`${accessToken}`);
})
