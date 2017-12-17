var wsUri;
var btnConnect;
var btnDisconnect;
var btnRefresh;
var ws;

var sUri = 'ws://127.0.0.1:7772/region';
var commands;
var playerKey = '';

window.addEventListener('load', onPageLoaded, false)

function loadJSON(callback) {   
    var xobj = new XMLHttpRequest();
    xobj.overrideMimeType("application/json");
    xobj.open('GET', './command.json', true);
    xobj.onreadystatechange = function () {
        if (xobj.readyState == 4) {
            callback(xobj.responseText);
        }
    };
    xobj.send(null);
}
    
function onPageLoaded() {
    wsUri = document.getElementById('wsUri');
    wsUri.value = sUri;

    btnConnect = document.getElementById('connect');
    btnConnect.onclick = doConnect;
    btnConnect.disabled = false;
    btnDisconnect = document.getElementById('disconnect');
    btnDisconnect.onclick = doDisconnect;
    btnDisconnect.disabled = true;

    btnRefresh = document.getElementById('refresh');
    btnRefresh.onclick = doRefresh;

    loadJSON(function(response) {
        commands = JSON.parse(response);
        var body = document.getElementsByTagName('body')[0];
        for (i = 0; i < commands.length; i++) {
            var txtCmdName = document.createElement('div');
            txtCmdName.id = 'cmdName' + i;
            txtCmdName.className = "cmdName";
            txtCmdName.textContent = commands[i].name;
            body.appendChild(txtCmdName);
            var br = document.createElement('br');
            body.appendChild(br);
    
            var txtCmd = document.createElement('input');
            txtCmd.type = "text";
            txtCmd.id = 'command' + i;
            txtCmd.className = "command";
            body.appendChild(txtCmd);
    
            txtCmd.value = JSON.stringify(commands[i].command);
    
            var btnSend = document.createElement('button');
            btnSend.id = 'send' + i;
            btnSend.innerHTML = "Send"
            body.appendChild(btnSend);
    
            var br = document.createElement('br');
            body.appendChild(br);
    
            btnSend.btnIndex = i;
            btnSend.onclick = doSend;
        }
    });
}

function doConnect() {
    ws = new WebSocket(wsUri.value);
    ws.onopen = function(e) { onWebSocketOpen(e) };
    ws.onclose = function(e) { onWebSocketClose(e) };
    ws.onmessage = function(e) { onWebSocketMessage(e) };
    ws.onerror = function(e) { onWebSocketError(e) };
}
function doDisconnect() {
    ws.close();
}
function doRefresh() {
    loadJSON(function(response) {
        commands = JSON.parse(response);
        for (i = 0; i < commands.length; i++) {
            var txtCmdName = document.getElementById('cmdName' + i);
            txtCmdName.textContent = commands[i].name

            var txtCmd = document.getElementById('command' + i);
            txtCmd.value = JSON.stringify(commands[i].command);
        }
    });
}
function doSend() {
    var txtCmd = document.getElementById('command' + this.btnIndex);
    ws.send(txtCmd.value);
}

function onWebSocketOpen(e) {
    btnConnect.disabled = true
    btnDisconnect.disabled = false
}
function onWebSocketClose(e) {
    btnConnect.disabled = false
    btnDisconnect.disabled = true
}
function onWebSocketMessage(e) {

}
function onWebSocketError(e) {
    window.alert('ERROR: ' + e.data)
}