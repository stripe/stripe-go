import { LightningElement, api, track } from 'lwc';

export default class Wizard extends LightningElement {
    @api setup = false;
    @api showClose = false;
    @api allowJump = false;
    @api title = 'Wizard Title';
    @api description = '';
    @track toast = {};
    @track shown = false;
    @track currentStepIndex = 0;

    @api init(completedSteps) {
        this.completedSteps = JSON.parse(JSON.stringify(completedSteps));
    }

    @api get progress() {
        let progress = 0;

        this.steps.forEach((step) => {
            if(this.completedSteps.hasOwnProperty(step.tagName)) {
                progress++;
            }
        });

        return progress / this.steps.length;
    }

    @api hide(event) {
        this.classList.add('slds-hide');
        this.shown = false;
    }

    @api show(index) {
        if(index === undefined) {
            index = 0;
        }

        this.showStep(this.steps[index]);
        this.currentStepIndex = index;
        this.classList.remove('slds-hide');
        this.shown = true;
    }

    get completedSteps() {
        return this._completedSteps || {};
    }

    set completedSteps(completedSteps) {
        this._completedSteps = JSON.parse(JSON.stringify(completedSteps));
    }

    get steps() {
        return (
            this._steps ||
            Array.from(this.querySelectorAll('*')).filter((el) => {
                return el.assignedSlot && el.assignedSlot.name === '';
            })
        );
    }

    set steps(_steps) {
        this._steps = _steps;
    }
    
    get showProgressIndicator() {
        return this.shown && this.steps.length > 1;
    }

    // look for a modal component passed into the modal slot. if none found, use default modal provided
    get modal() {
        let _modal = Array.from(this.querySelectorAll('c-modal')).filter(
            (el) => {
                return el.assignedSlot && el.assignedSlot.name === 'modal';
            }
        );
        return _modal.length
            ? _modal[0]
            : this.template.querySelector('c-modal');
    }

    constructor() {
        super();
        this.template.addEventListener('back', this.back.bind(this));
        this.template.addEventListener('exit', this.exit.bind(this));
        this.template.addEventListener('next', this.next.bind(this));
        this.template.addEventListener(
            'primary',
            this.handlePrimary.bind(this)
        );
        this.template.addEventListener('showModal', this.showModal.bind(this));
        this.template.addEventListener('toast', this.showToast.bind(this));
    }

    connectedCallback() {
        this.hide();
    }

    renderedCallback() {
        if(this._init !== true) {
            this._init = true;
            let steps = Array.from(this.querySelectorAll('*')).filter((el) => {
                if(el.assignedSlot && el.assignedSlot.name === '') {
                    this.hideStep(el);
                    return true;
                }
                
                return false;
            });
            this.steps = steps;
            
            if(!this.setup) {
                this.show();
            }
        }
    }

    showToast(event) {
        clearTimeout(this.toastTimeout);

        let iconName,
            type = event.detail.type;

        switch (type) {
            case 'error':
                iconName = 'utility:error';
                break;
            case 'warning':
                iconName = 'utility:warning';
                break;
            case 'success':
                iconName = 'utility:success';
                break;
            default:
                iconName = 'utility:info';
                break;
        }

        this.toast = {
            class: 'slds-notify slds-notify_toast slds-theme_' + type,
            dismissible: event.detail.mode !== 'sticky',
            icon: {
                class:
                    'slds-icon_container slds-m-right_small slds-no-flex slds-align-top slds-icon-utility-' +
                    type,
                name: iconName
            },
            title: event.detail.title
                ? event.detail.title
                : event.detail.message,
            message: event.detail.title ? event.detail.message : null,
            type: type
        };

        if(event.detail.mode !== 'pester') {
            let duration = event.detail.duration || 3000;

            this.toastTimeout = setTimeout(() => {
                this.toast = {};
            }, duration);
        }
    }

    hideToast() {
        this.toast = {};
    }

    exit(event) {
        event.stopPropagation();

        let currentStep;
        let index = this.steps.indexOf(event.target);

        if(index !== -1) {
            currentStep = this.steps[index];
        } else {
            this.steps.forEach((step) => {
                if(!step.classList.contains('slds-hide')) {
                    currentStep = step;
                }
            });
        }

        if(currentStep) {
            this.hideStep(currentStep);
        }

        this.hide();
        this.dispatchEvent(
            new CustomEvent('exit', {
                bubbles: true,
                composed: true
            })
        );
    }

    back(event) {
        event.stopPropagation();

        let index = this.steps.indexOf(event.target);

        if(index !== 0) {
            this.showStep(this.steps[index - 1]);
            this.currentStepIndex = index - 1;
            this.hideStep(this.steps[index]);
        }
    }

    next(event) {
        event.stopPropagation();

        let index = this.steps.indexOf(event.target);

        if(index > -1) {
            let completedSteps = JSON.parse(
                JSON.stringify(this.completedSteps)
            );
            let currentStep = this.steps[index];

            completedSteps[currentStep.tagName] = 1;
            this.completedSteps = completedSteps;

            if(index >= this.steps.length - 1) {
                this.hideStep(currentStep);
                this.hide();
                this.dispatchEvent(
                    new CustomEvent('finish', {
                        bubbles: true,
                        composed: true
                    })
                );
            } else {                
                this.showStep(this.steps[index + 1]);
                this.currentStepIndex = index + 1;
                this.hideStep(currentStep);
            }
        }
    }

    navigateTo(event) {
        event.preventDefault();

        if(!this.allowJump) {
            return;
        }

        let currentStep;
        let index = parseInt(event.currentTarget.dataset.index, 10);
        let nextStep = this.steps[index];

        this.steps.forEach((step) => {
            if(!step.classList.contains('slds-hide')) {
                if(step == nextStep) return;
                currentStep = step;
            }
        });
        
        if(currentStep) this.hideStep(currentStep);
        this.showStep(nextStep);
        this.currentStepIndex = index;
    }

    handlePrimary(event) {
        event.preventDefault();
        event.stopPropagation();

        let currentStep;
        this.steps.forEach((step) => {
            if(!step.classList.contains('slds-hide')) {
                currentStep = step;
            }
        });

        this.hideStep(currentStep);

        if(
            this._modalStepIndex !== null &&
            this._modalStepIndex >= 0 &&
            this._modalStepIndex < this.steps.length
        ) {
            this.showStep(this.steps[this._modalStepIndex]);
            this.currentStepIndex = this._modalStepIndex;
        } else {
            this.hide();
            this.dispatchEvent(
                new CustomEvent('exit', {
                    bubbles: true,
                    composed: true
                })
            );
        }
    }

    hideStep(step) {
        step.classList.add('slds-hide');
        if(typeof step.hide === 'function') {
            step.hide();
        }
    }

    showStep(step) {
        this.template.querySelector('.strike-wizard__content').scrollTop = 0;
        
        if(typeof step.show === 'function') {
            step.show().then(() => {
                step.classList.remove('slds-hide');
            });
        } else {
            step.classList.remove('slds-hide');
        }
    }

    showModal(event) {
        this._modalStepIndex = this.steps.indexOf(event.target);

        switch (event.detail) {
            case 'back':
                this._modalStepIndex--;
                break;
            case 'next':
                this._modalStepIndex++;
                break;
            default:
                this._modalStepIndex = null;
                break;
        }

        this.modal.show();
    }
}
