
import {
    hterm,
    lib
} from 'hterm-umdjs';

import channel from './protocol';

hterm.defaultStorage = new lib.Storage.Memory();
const term = new hterm.Terminal();
term.decorate(document.querySelector('#terminal'))
const io = term.io.push()
const chan = new channel("ws://localhost:8081/terminal", (text) => {
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
