import { LightningElement, api } from 'lwc';

const AlertTypes = {
    warning: {
        default_title: 'Warning',
        default_classes: 'slds-notify slds-notify_alert slds-alert_warning',
        default_icon: 'utility:warning',
    },
    alert: {
        default_title: 'Alert',
        default_classes: 'slds-notify slds-notify_alert',
        default_icon: 'utility:question_mark',
    },
    error: {
        default_title: 'Error',
        default_classes: 'slds-notify slds-notify_alert slds-alert_error',
        default_icon: 'utility:error',
    },
    offline: {
        default_title: 'Offline',
        default_classes: 'slds-notify slds-notify_alert slds-alert_offline',
        default_icon: 'utility:offline',
    },
};

export default class Alert extends LightningElement {
    @api type;
    @api title;
    @api classes;
    @api icon;
    @api message;
    @api iconSize = 'x-small';

    get getTitle() {
        return (this.title) ? this.title : AlertTypes[this.type].default_title;
    }

    get getClasses() {
        return (this.classes) ? this.classes : AlertTypes[this.type].default_classes;
    }

    get getIcon() {
        return (this.icon) ? this.icon : AlertTypes[this.type].default_icon;
    }

    get getIconSize() {
        return this.iconSize;
    }

    get getMessage() {
        return this.message;
    }
}

const typeStrings = {};
const typeKeys = Object.keys(AlertTypes);
for (let i = 0; i < typeKeys.length; i++) {
    const typeKey = typeKeys[i];
    typeStrings[typeKey] = typeKey;
}

Alert.Types = typeStrings;