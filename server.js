const net = require('net');
const server = new net.Server();
server.listen(8081, function() {
    console.log('Launching server...');
});

server.on('connection', function(socket) {
    socket.on('data', function(chunk) {
        console.log(`Data received from client: ${chunk.toString()}`);
        socket.write(chunk.toString().toUpperCase() + "\n");
    });
});