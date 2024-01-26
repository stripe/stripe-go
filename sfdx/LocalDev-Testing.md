# How to test locally

1. Follow normal setup routine.
   - Create a scratch org, refresh auth token, etc
   - Deploy sfdx package to your Scratch org
1. Follow the instructions below to setup ngrok.
1. Run ruby locally... what works for me is: `dotenv bundle exec puma`
1. Make the local instance internet accessible. You can use ngrok, or my custom localtunnel server if you like. Details below.
1. Add a Remote Site Config for this url to enable Salesforce to call it.  
   (Setup -> Remote Site Settings), /lightning/setup/SecurityRemoteProxy/home
   Select New. Name the remote site whatever you want. The URL is found in your terminal: https:/<>.ngrok-free.app
1. Go to /lightning/setup/CustomMetadata/home, click on `Manage Records` next to `Setup Connection Data`.
   1. Click `Edit` on the `Default` Setup_Connection_Data\_\_mdt record
   2. turn on `Local Authorization Target`
   3. set `Platform Target` to the URL from step 3.
   4. Add .ngrok-free.app to your STRIPEFORCE_HOSTS in your .env
   5. [optional] If you have defined `OAUTH_STATE_SIGNING_KEY` in your `.env` file, set the value of `OAuth State Signing Key` to that value.
1. If something seems wonky, set OAuth_State_Signing_Key\_\_c to `98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a` in Salesforce, but it should default to it even if not set.
1. Restart your service
1. Test!

Note: make sure you have run the Ruby migrations have run locally

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
ngrok config add-authtoken<your_token>
```

#### Running

To run ngrok, execute:

```
ngrok http 3100
```

This will product a URL that you will use above.

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
