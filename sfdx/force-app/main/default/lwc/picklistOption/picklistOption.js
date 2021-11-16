import { LightningElement, api, track } from 'lwc';

export default class picklistOption extends LightningElement {
    @api option;
    @api parentValue;
    @api index;

    @track hasFocus = false;

    @api
    select(){
        if(!this.option.disabled){
            let evt = new CustomEvent('select',{
                detail:{
                    selectedOption: this.option
                }
            });
            this.dispatchEvent(evt);
        }
    }

    @api
    addFocus(){
        this.hasFocus = true;
    }

    @api
    removeFocus(){
        this.hasFocus = false;
    }

    @api
    getValue(){
        return this.option.value;
    }

    updateFocusIndex(event){
        event.preventDefault();
        this.dispatchEvent(new CustomEvent('updatefocusindex', {detail: this.option.value}));
    }

    preventDefault(event){
        event.preventDefault();
    }

    get listItemClassList() {
        return 'slds-listbox__item' + (this.option.disabled ? ' so-listbox__item_disabled' : '');
    }

    get isSelected(){
        return this.option.value === this.parentValue;
    }

    get iconClasslist(){
        return this.option.value === this.parentValue ? '' : 'slds-is-selected';
    }

    get optionClasslist(){
        return 'slds-media slds-listbox__option slds-listbox__option_plain slds-media_small slds-media_center' + 
               (this.isSelected ? ' slds-is-selected' : '') + 
               (this.option.disabled ? ' so-disabled' : '') + 
               (this.hasFocus ? ' slds-has-focus' : '');
    }
}