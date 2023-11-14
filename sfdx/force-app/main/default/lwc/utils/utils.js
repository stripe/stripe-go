function getErrorMessage(error) {
    if (error.message) {
        return error.message;
    }
    return JSON.stringify(error);
}

function openWindow(url) {
    return window.open(url, '"_blank"');
}

function createToast(message, variant, mode) {
    return new CustomEvent('showtoast', {
        bubbles: true,
        composed: true,
        detail: {
            toast: {
                message: message,
                variant: variant ? variant : 'info',
                mode: mode ? mode : 'dismissible'
            }
        }
    });
}


export {
    getErrorMessage,
    createToast,
    openWindow,
}