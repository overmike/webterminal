class channel {
    constructor(endpoint, onmessage) {
        this.ws = new WebSocket(endpoint)
        this.ws.onmessage = this.onmessage
        this.notifyexternal = onmessage
    }

    send = (data) => {
        const text = JSON.stringify({message:data})
        this.ws.send(text)
    }

    onmessage = (event) => {
        const text = JSON.parse(event.data);
        this.notifyexternal(text);
    }

}

export default channel