package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dimple278/go-chat-app/commands"
	"github.com/dimple278/go-chat-app/db"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Create a new Socket.io server
	server := socketio.NewServer(nil)

	// Handle connection event
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("New connection: ", s.ID())
		s.SetContext(map[string]interface{}{}) // Set empty context initially
		return nil
	})

	// Handle event where the user sends their ID and username to the server after login
	server.OnEvent("/", "set_user_info", func(s socketio.Conn, userID int, username string) {
		fmt.Printf("User %s (ID: %d) has connected with socket ID: %s\n", username, userID, s.ID())

		// Set the user info in the socket's context
		s.SetContext(map[string]interface{}{
			"userID":   userID,
			"username": username,
		})

		// Add user to ActiveUsers map
		commands.ActiveUsers[s.ID()] = username

		// Broadcast a message to notify all users about the new user joining
		server.BroadcastToNamespace("/", "chat message", fmt.Sprintf("%s joined the chat", username))

		// Send chat history to the new user
		chatHistory, err := db.FetchChatHistory(20, 0) // Fetch the last 20 messages
		if err != nil {
			fmt.Println("Error fetching chat history:", err)
			return
		}
		for _, message := range chatHistory {
			s.Emit("chat message", message)
		}
	})

	// Handle "chat message" event
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		// Check if the message is a command
		if strings.HasPrefix(msg, "/") {
			switch {
			case strings.HasPrefix(msg, "/listusers"):
				// List active users
				response := commands.ListUsersCommand()
				s.Emit("chat message", response)
			case strings.HasPrefix(msg, "/logout"):
				// Handle logout command
				userInfo, ok := s.Context().(map[string]interface{})
				if ok {
					username := userInfo["username"].(string)
					// Broadcast the logout message before closing the connection
					server.BroadcastToNamespace("/", "chat message", fmt.Sprintf("%s left the chat", username))
				}

				// Remove the user from the active users and disconnect
				// response := commands.LogoutCommand(s)
				s.Close()
				return
			case strings.HasPrefix(msg, "/history"):
				// Handle history command
				response := commands.HistoryCommand(msg)
				s.Emit("chat message", response)
			default:
				s.Emit("chat message", "Unknown command")
			}
			return
		}

		// Retrieve user info from the connection context
		userInfo, ok := s.Context().(map[string]interface{})
		if !ok || userInfo["username"] == nil || userInfo["userID"] == nil {
			fmt.Println("Error: user info not found in context")
			s.Emit("chat message", "Error: user info not found. Please log in again.")
			return
		}

		username := userInfo["username"].(string)
		userID := userInfo["userID"].(int)

		fmt.Printf("Received message from %s (userID: %d): %s\n", username, userID, msg)

		// Save the chat message to the database
		err := db.SaveChatMessage(userID, msg)
		if err != nil {
			fmt.Println("Error saving chat message to DB:", err)
		}

		// Broadcast the message to all clients
		server.BroadcastToNamespace("/", "chat message", fmt.Sprintf("%s: %s", username, msg))
	})

	// Handle disconnect event for unexpected disconnects (not from logout)
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Retrieve the username from ActiveUsers
		username, exists := commands.ActiveUsers[s.ID()]
		if exists {
			fmt.Printf("User %s disconnected: %s\n", username, reason)

			// Remove user from the ActiveUsers map
			delete(commands.ActiveUsers, s.ID())
		} else {
			fmt.Printf("Unknown user disconnected with socket ID %s: %s\n", s.ID(), reason)
		}
	})

	// Serve the socket.io server on port 8000
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	fmt.Println("Chat server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
