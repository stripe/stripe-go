import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';
import { LightningElement, track, api, wire } from 'lwc';

export default class SystemConnectionsStep extends LightningElement {
    @track salesforceComplete = false;
    @track stripeComplete = false;
    @track connectWindow;
    @track isSandbox;
    @track rubyBaseURI = 'https://stripe-force.herokuapp.com'; 
    @api hideAction = false;

    connectedCallback() {
        this.stripeConnectedAppCallback(true);
    }

    stripeConnectedAppCallback() {
        this.validateConnectionStatus(true);
        this.postMessageListener = (event) => {
            if(event.origin === this.rubyBaseURI && event.data === 'connectionSuccessful') {
                this.connectWindow.close();
                this.validateConnectionStatus(false);
            }  
        }  
        window.addEventListener("message", this.postMessageListener.bind(this));
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }

    connectToStripe(event) {
        this.connectWindow = this.isSandbox ? window.open(this.rubyBaseURI + '/auth/salesforcesandbox', '"_blank"') : window.open(this.rubyBaseURI + '/auth/salesforce', '"_blank"');
    }

    async validateConnectionStatus(isConnectedCallback) {
        this.loading = true;
        try {
            const validateConnection = await validateConnectionStatus({
                isConnectedCallback : isConnectedCallback
            }); 
            this.data =  JSON.parse(validateConnection);
            if(this.data.isSuccess) {
                let isConnected = this.data.results.isConnected;
                this.isSandbox = this.data.results.isSandbox
                if (isConnected === 'fresh') {
                    this.salesforceComplete = true;
                    this.stripeComplete = true;
                    this.dispatchEvent(new CustomEvent('getmappingconfigurations'));
                    this.showToast('Authorization successfully completed', 'success', 'dismissable');
                } else if (isConnected === 'notConnected') {
                    this.salesforceComplete = false;
                    this.stripeComplete = false;
                    this.showToast('Authorization expired. Please reauthorize your Stripe and Salesforce accounts.', 'error', 'sticky');
                } else if (isConnected === 'failed') {
                    this.salesforceComplete = false;
                    this.stripeComplete = false;
                    this.showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
                } else if (isConnected) {
                    this.salesforceComplete = true;
                    this.stripeComplete = true;
                    this.dispatchEvent(new CustomEvent("enablenext"));
                } 
            } else { 
                this.showToast(this.data.error, 'error', 'sticky');
            }
        } catch (error) {
            this.showToast(error.message, 'error', 'sticky');
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
