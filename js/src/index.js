
import {
    hterm,
    lib
} from 'hterm-umdjs';

import channel from './protocol';

var loc = window.location,
    wsUri;
if (loc.protocol === "https:") {
    wsUri = "wss:";
} else {
    wsUri = "ws:";
}
wsUri += "//" + loc.host;
wsUri += loc.pathname + "terminal";

hterm.defaultStorage = new lib.Storage.Memory();
const term = new hterm.Terminal();
term.decorate(document.querySelector('#terminal'))
const io = term.io.push()
const chan = new channel(wsUri, (text) => {
    io.print(text)
});
io.onVTKeystroke = (data) => {
    chan.send(data)
}
io.sendString = (data) => {
    chan.send(data)
}
io.onTerminalResize = (columns, rows) => {
    chan.resize(columns, rows)
};
term.installKeyboard()


const webterminal = document.getElementById("webterminal");
