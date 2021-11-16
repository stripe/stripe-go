import { LightningElement, track, api } from 'lwc';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';

export default class OutboundStep extends LightningElement {
    @api isAuthComplete = false;
    

    connectedCallback() {
        this.template.addEventListener('next', this.next.bind(this));
        //this.stripeConnectedAppCallback();
        console.log('IS this being hit');
        this.postMessageListener = (event) => {
            console.log('event')
            console.log(event)
            if(event.origin === 'https://connect.stripe.com' ) {
                
                //this.stripeConnectedAppCallback();
            } 
        } 
       window.addEventListener('message', this.postMessageListener);
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }
    connectToStripe() {
        //const rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforce'; //production
        const rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforcesandbox'; //sandbox
        window.open(rubyAuthURI, '"_blank"').focus();;
    }

    stripeConnectedAppCallback() {
        this.loading = true;
        validateConnectionStatus()
        .then(response => {
           let responseData = JSON.parse(response);
            if(responseData.isSuccess) {
                let isConnected = responseData.results.isConnected;
                if(isConnected === 'fresh') {
                    this.isAuthComplete = true;
                    this.showToast('Authorization successfully completed', 'success');
                }
                if(isConnected === true)this.isAuthComplete = true;
            } else { 
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error.body.message, 'error');
        }).finally(() => {
            this.loading = false;
        });
    }

    next(event) {
        event.stopPropagation();
        saveData({
            setupData: {
                Steps_Completed__c: JSON.stringify({
                    'C-OUTBOUND-STEP': 1
                }),
                Live_Mode__c: this.setupData.Live_Mode__c
            }
        }).then((responseDataString) => {
            let responseData = JSON.parse(responseDataString);

            if(responseData.isSuccess) {
                this.setupData = responseData.results.setupData;
                this.dispatchEvent(new CustomEvent('next', {
                    bubbles: true,
                    composed: true
                }));

            } else {
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error, 'error');
        });
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