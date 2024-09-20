package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	// Handle new connections
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("New connection: ", s.ID())
		return nil
	})

	// Handle incoming messages
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println("Received message: ", msg)
		// Broadcast the message to all connected clients
		server.BroadcastToNamespace("/", "chat message", msg)
	})

	// Handle disconnects
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Disconnected: ", s.ID(), " Reason: ", reason)
	})

	go server.Serve()
	defer server.Close()

	// Serve the socket.io server on port 8000
	http.Handle("/socket.io/", server)
	fmt.Println("Chat server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
