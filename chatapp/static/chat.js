var ws;

function connect() {
    var url = 'wss://fdatgtaewatr.execute-api.AWS-REGION.amazonaws.com/prod'; // Replace with your API Gateway URL
    ws = new WebSocket(url);

    ws.onopen = function(event) {
        document.getElementById('status').innerHTML = 'Connection is established';
    };
    
    ws.onmessage = function(event) {
        var message = JSON.parse(event.data);
        if (message.message) {
            document.getElementById('status').innerHTML = 'Message received: ' + message.message;
        }
    };

    ws.onerror = function(event) {
        console.log('WebSocket error: ', event);
    };
    
    ws.onclose = function() {
        document.getElementById('status').innerHTML = 'Connection is closed';
        ws = null;
    };
}

document.getElementById('open-button').addEventListener('click', function() {
    if(ws == null) {
        connect();
    } else {
        document.getElementById('status').innerHTML = 'Connection is already open';
    }
});

document.getElementById('send-button').addEventListener('click', function() {
    if(ws != null) {
        var messageInput = document.getElementById('message-input'),
            message = messageInput.value;
        ws.send(JSON.stringify({
            action: 'sendMessage',
            data: message
        }));
        messageInput.value = '';
        document.getElementById('status').innerHTML = 'Message sent: ' + message;
    } else {
        document.getElementById('status').innerHTML = 'Connection is closed. Open the connection first to send a message';
    }
});

document.getElementById('close-button').addEventListener('click', function() {
    if(ws != null) {
        ws.close();
        document.getElementById('status').innerHTML = 'Connection closed by user';
    } else {
        document.getElementById('status').innerHTML = 'Connection is already closed';
    }
});
