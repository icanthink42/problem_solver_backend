# WebSocket Server with Gin

This is a simple WebSocket server implementation using Go, Gin, and Gorilla WebSocket.

## Features

- WebSocket server implementation
- Real-time message broadcasting to all connected clients
- Thread-safe client management
- Simple HTTP endpoint for server status

## Setup

1. Make sure you have Go installed (version 1.21 or later)
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## Endpoints

- `GET /`: Returns a simple JSON message indicating the server is running
- `GET /ws`: WebSocket endpoint for client connections

## Testing

You can test the WebSocket connection using tools like [websocat](https://github.com/vi/websocat) or browser-based WebSocket clients:

```bash
# Using websocat
websocat ws://localhost:8080/ws
```

## Implementation Details

- The server uses a hub pattern to manage WebSocket connections
- Each client runs two goroutines: one for reading messages and one for writing
- Messages sent by any client are broadcasted to all connected clients
- The implementation includes proper connection cleanup and error handling
