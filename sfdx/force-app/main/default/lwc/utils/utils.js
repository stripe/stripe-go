function getErrorMessage(error) {
    if (error.message) {
        return error.message;
    }
    return JSON.stringify(error);
}


export {
    getErrorMessage
}