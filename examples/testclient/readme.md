# AudioSocket Example

This repository contains an example of using AudioSocket to send audio data from a client to a server.

## Files

- `server.go`: This is the server that listens for incoming connections and processes the audio data received from the client.
- `client.go`: This is the client that connects to the server, sends audio data, and then disconnects.

## Usage

### Server

To run the server, use the `go run server.go` command:

By default, the server listens on port 9092. You can change this by modifying the `listenAddr` constant in `server.go`.

### Client

To run the client, use the `go run client.go` command:

By default, the client connects to the server at `localhost:9092` and sends the audio data from the file `test.slin`. You can change these by modifying the `serverAddr` and `fileName` constants in `client.go`.

## How It Works

The client first connects to the server and sends a call ID message. It then reads the audio data from the specified file and sends it to the server in chunks. After all the audio data is sent, it sends a hangup message to the server and then closes the connection.

The server listens for incoming connections. When a connection is accepted, it reads the call ID from the connection and then processes the incoming audio data. When it receives a hangup message, it closes the connection.

## Note

This is a basic example and does not handle many edge cases. Depending on your needs, you might need to add more functionality to the client and server.
