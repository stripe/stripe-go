# How to test locally

1. Follow normal setup routine.
2. Run ruby locally... what works for me is: `dotenv bundle exec puma`
2. Make the local instance internet accessible. You can use ngrok, or my custom localtunnel server if you like. Details below.
3. Add a Remote Site Config for this url to enable Salesforce to call it.  
   (Setup -> Remote Site Settings), /lightning/setup/SecurityRemoteProxy/home
4. Edit the Default Setup_Connection_Data__mdt record, turn on "Local Authorization  
   Target" and set "Platform Target" to the URL from step 2. /lightning/setup/CustomMetadata/home
5. If something seems wonky, set OAuth_State_Signing_Key__c to `98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a` in Salesforce, but it should default to it even if not set.
6. Test!

## making localdev accessible

We can (and do!) use http://localhost:3100 for some local testing, but in order for Salesforce to be able
to call it, we need to expose it somewhere that Salesforce can call it.

The simplest way to do this is with some sort of local tunnel system.

### ngrok

#### Installation

Be sure to install from NPM, not the website.

```
npm install -g ngrok
```

#### Setup

Create an account with [ngrok](https://ngrok.com/), and configure your authtoken.

```
ngrok authtoken <your_token>
``` 

#### Running

To run ngrok, execute:

```
ngrok http 3100
```

### My secure localtunnel server

I run a dedicated localtunnel server for just such purposes as well.

It keeps no logs or anything, and just does what it says on the box to provide public HTTPS transit.

For ease of use, I have added [bin/expose-for-salesforce.sh](bin/expose-for-salesforce.sh) to the repo.

Simply run `./bin/expose-for-salesforce.sh`, and it will give you a URL to use.

```
st-jmather-c1:stripe-salesforce jmather-c$ ./bin/expose-for-salesforce.sh
your url is: https://jmather-c.lt.jmather.com
```

Please note: this is for developer use only. It is not a production service, and I make no guarantees about it.

For use with CI, or other purposes, Stripe should set up their own localtunnel secure server using
[my localtunnel-server fork](https://github.com/jmather/localtunnel-secure-server).

Thanks!