import { LightningElement, track, api, wire } from 'lwc';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import isCpqEnabled from '@salesforce/apex/utilities.isCpqEnabled';
import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';

export default class OutboundStep extends LightningElement {
    @api isAuthComplete = false;
    @track connectWindow;

    connectedCallback() {
        this.checkIfCPQIsEnabled();
    }

   stripeConnectedAppCallback(isCpqEnabled) {
        const rubyServiceAuthOriginURI = 'https://stripe-force.herokuapp.com'
        if(isCpqEnabled === true) {
            this.validateConnectionStatus(true);
            this.postMessageListener = (event) => {
                if(event.origin === rubyServiceAuthOriginURI && event.data === 'connectionSuccessful') {
                    this.connectWindow.close()
                    this.validateConnectionStatus(false);
                }  
                console.log('Can we see dis in connectedCallback??')
            } 
            window.addEventListener("message", this.postMessageListener.bind(this));
        }
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }

    connectToStripe() {
        //'https://stripe-force.herokuapp.com/auth/salesforce'; //production 
        const rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforcesandbox'; //sandbox
        this.connectWindow = window.open(rubyAuthURI, '"_blank"');
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
                if(isConnected === 'fresh') {
                    this.isAuthComplete = true;
                    this.showToast('Authorization successfully completed', 'success');
                } else if(isConnected === 'notConnected') {
                    this.isAuthComplete = false;
                    this.showToast('Your authorization expired, please reauthorize your stripe & salesforce accounts', 'error');
                }else if(isConnected === 'failed') {
                    this.isAuthComplete = false;
                    this.showToast('Please refresh the page there was an error when checking your connection status', 'error');
                } else if (isConnected === true) {
                    this.isAuthComplete = true;
                } 
                this.loading = false;
            } else {
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error, 'error');
        }
    }

    checkIfCPQIsEnabled() {
       /* This is the cpq check to determine if we should show an error toast or not
       commented this out for now in case we decide we want to add this check back  
        try {
            const checkIfCpqIsEnabled = await isCpqEnabled();
            this.data =  JSON.parse(checkIfCpqIsEnabled);
            if(this.data.isSuccess) {
                let isCpqInstalled = this.data.isCpqInstalled;
                if(isCpqInstalled === true) {
                    this.stripeConnectedAppCallback(true);
                } else {
                    this.showToast('The CPQ package is not installed in this org', 'error');
                }
            } else {
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error, 'error');
        }*/
        this.stripeConnectedAppCallback(true);
    }

    async next(event) {
        if(this.isAuthComplete === true) {
            event.stopPropagation();
            try {
                const saveSetupData = await saveData({
                    setupData: {
                        Steps_Completed__c: JSON.stringify({
                            'C-OUTBOUND-STEP': 1
                        })
                    }
                });
                this.data =  JSON.parse(saveSetupData);
                if(this.data.isSuccess) {
                    this.setupData = this.data.setupData;
                    this.dispatchEvent(new CustomEvent('next', {
                        bubbles: true,
                        composed: true
                    }));

                } else {
                    this.showToast(this.data.error, 'error');
                }
            } catch (error) {
                this.showToast(error, 'error');
            }
        }
    }


    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: { 
                toast: {
                    message: message,
                    variant: variant,
                    mode: mode
                }
            }
        })); 
    }
}
