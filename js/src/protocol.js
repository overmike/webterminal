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

    resize = (cols, rows) => {
        const resizeCmd = JSON.stringify({
            resize: {
                columns: cols,
                rows: rows
            }
        })
        this.ws.send(resizeCmd)
    }

    onmessage = (event) => {
        const text = JSON.parse(event.data);
        this.notifyexternal(text.result.message);
    }

}

export default channel