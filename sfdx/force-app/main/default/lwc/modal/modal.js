import { LightningElement, api, track } from 'lwc';

export default class Modal extends LightningElement {
    @api title = '';
    @api hideHeader = false;
    @api hideFooter = false;
    @api large = false;
    @api medium = false;
    @api small = false;
    @api contentPadding = 'medium';
    @api overflowVisible = false;
    @api modalStyle = '';
    @api modalContentStyle = '';
    @api backdropStyle = '';
    @api variant = '';

    @track showModal = false;
    @track tabIndex = '-1';

    handleKeyUp = (function(event) {
        if(event.key === 'Escape') {
            this.hide();
        }
    }).bind(this);
    
    @api
    show() {
        this.showModal = true;
        this.tabIndex = '0';
        
        this.dispatchEvent(new CustomEvent('modalshow'));

        window.setTimeout(() => {
            let body = this.template.querySelector('[data-id="body"]');

            if(body){
                body.scrollTop = 0;
            }
            
            window.addEventListener('keyup', this.handleKeyUp);
            this.dispatchEvent(new CustomEvent('modalshown'));
        }, 400);
    }
    
    @api
    hide() {
        this.showModal = false;
        this.tabIndex = '-1';
        
        window.removeEventListener('keyup', this.handleKeyUp);
        this.dispatchEvent(new CustomEvent('modalhide'));

        window.setTimeout(() => {
            this.dispatchEvent(new CustomEvent('modalhidden'));
        }, 400);
    }

    close() {
        this.dispatchEvent(new CustomEvent('modalclose'));
        this.hide();
    }

    get modalClass() {
        return 'slds-modal' + (this.variant === 'error' || this.variant === 'warning' ? ' slds-modal_prompt' : '') + (this.large ? ' slds-modal_large' : '') + (this.medium ? ' slds-modal_medium' : '') + (this.small ? ' slds-modal_small' : '') + (this.showModal ? ' slds-fade-in-open' : ' sm-pointer-events_none');
    }

    get backdropClass() {
        return 'slds-backdrop' + (this.showModal ? ' slds-backdrop_open' : ' sm-pointer-events_none');
    }

    get modalHeaderClass() {
        return 'slds-modal__header' + (this.hideHeader ? ' slds-modal__header_empty' : '') + (this.variant === 'error' || this.variant === 'warning' ? ' slds-theme_alert-texture slds-theme_' + this.variant : '');
    }

    get modalFooterClass() {
        return 'slds-modal__footer' + (this.variant === 'error' || this.variant === 'warning' ? ' slds-theme_default' : '');
    }

    get modalContentClass() {
        return 'slds-modal__content slds-p-around_' + this.contentPadding + (this.overflowVisible ? ' slds-expanded' : '');
    }
}
