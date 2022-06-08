import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';
import { LightningElement, track, api, wire } from 'lwc';
import { getErrorMessage } from 'c/utils'

export default class SystemConnectionsStep extends LightningElement {
    @track salesforceComplete = false;
    @track stripeComplete = false;
    @track connectWindow;
    @track isSandbox;
    @track rubyBaseURI = 'https://stripe-force.herokuapp.com';
    @track salesforceNamespace;
    @api hideAction = false;

    connectedCallback() {
        this.stripeConnectedAppCallback(true);
    }

    stripeConnectedAppCallback() {
        this.validateConnectionStatus(true, '');
        this.postMessageListener = (event) => {
            if (event.origin === this.rubyBaseURI && event.data === 'stripeConnectionSuccessful') {
                this.validateConnectionStatus(false, 'salesforce');
            } else if (event.origin === this.rubyBaseURI && event.data === 'salesforceConnectionSuccessful') {
                this.connectWindow.close();
                this.validateConnectionStatus(false, 'stripe');
            }
        }
        window.addEventListener("message", this.postMessageListener.bind(this));
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }

    connectToStripe(event) {
        let oauthConnectionURL = this.rubyBaseURI;

        if(this.isSandbox) {
            oauthConnectionURL += '/auth/salesforcesandbox'
        } else {
            oauthConnectionURL += '/auth/salesforce'
        }

        oauthConnectionURL += "?salesforceNamespace=" + this.salesforceNamespace

        this.connectWindow = window.open(oauthConnectionURL, '"_blank"');
    }

    async validateConnectionStatus(isConnectedCallback, systemConnected) {
        this.loading = true;
        try {
            const validateConnection = await validateConnectionStatus({
                isConnectedCallback : isConnectedCallback,
                systemConnected : systemConnected
            });
            const responseData = JSON.parse(validateConnection);

            if(responseData.isSuccess) {
                const isStripeConnected = responseData.results.isStripeConnected;
                const isSalesforceConnected = responseData.results.isSalesforceConnected;
                this.isSandbox = responseData.results.isSandbox
                this.salesforceNamespace = responseData.results.salesforceNamespace

                if (isStripeConnected === 'freshConnection') {
                    this.stripeComplete = true;
                    this.showToast('Stripe to salesforce authorization successfully completed', 'success', 'dismissable');
                } 

                if (isSalesforceConnected === 'freshConnection') {
                    this.salesforceComplete = true;
                    this.showToast('Salesforce to stripe Authorization successfully completed', 'success', 'dismissable');
                }

                if (isStripeConnected === 'stripeDisconnected') {
                    this.stripeComplete = false;
                } else if (isStripeConnected === 'stripeConnected') {
                    this.stripeComplete = true;
                }

                if (isSalesforceConnected === 'salesforceDisconnected') {
                    this.salesforceComplete = false;
                } else if (isSalesforceConnected === 'salesforceConnected') {
                    this.salesforceComplete = true;
                }

                if (this.salesforceComplete === true && this.stripeComplete === true) {
                    this.dispatchEvent(new CustomEvent('getmappingconfigurations'));
                }

                if (isStripeConnected === 'failed' && isSalesforceConnected === 'failed') {
                    this.salesforceComplete = false;
                    this.stripeComplete = false;
                    this.showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
                } 
            } else {
                this.showToast(responseData.error, 'error', 'sticky');
            }
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        }
    }

    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: {
                toast: {
                    message: message,
                    variant: variant ? variant : 'info',
                    mode: mode ? mode : 'dismissible'
                }
            }
        }));
    }

    get actionLabel() {
        return (this.salesforceComplete && this.stripeComplete) ? 'Reauthorize' : 'Authorize';
    }
}
