import { LightningElement, api, track } from 'lwc';

export default class Toast extends LightningElement {
    @track timeout;
    @track message;
    @track variant;
    @track mode;
    @track showCloseButton = false;
    @api
    show(message, variant, mode) {
        this.message = message;
        this.variant = variant;
        this.mode = mode;

        let toast = this.template.querySelector('.ap-toast');
        toast.classList.remove('slds-hide');
        setTimeout(function(){
            toast.classList.add('ap-toast_shown');
        }, 200)

        if(mode == 'sticky') {
            toast.classList.add('ap-toast_shown');
            this.showCloseButton = true;
        }

        if(mode !== 'sticky') {
            this.showCloseButton = false;
            this.timeout = setTimeout(function(){
            toast.classList.remove('ap-toast_shown');
            setTimeout(function(){
                toast.classList.add('slds-hide');
            }, 400)
            }, 4000);
        }
    }

    @api
    hide() {
        let toast = this.template.querySelector('.ap-toast');
        toast.classList.remove('ap-toast_shown');
        setTimeout(function(){
            toast.classList.add('slds-hide');
        }, 400)
        clearTimeout(this.timeout);
    }

    /* get showCloseButton() {
        return this.mode !== 'sticky';
    }
 */
    get iconName() {
        return `utility:${this.variant}`;
    }

    get toastClassList() {
        return `ap-toast slds-media slds-media_center slds-notify slds-notify_toast slds-theme_${this.variant}`;
    }
}
