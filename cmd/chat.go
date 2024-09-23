package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	socketio_client "github.com/hesh915/go-socket.io-client"
	"github.com/spf13/cobra"
)

// Helper function for reading input from the user
func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading message:", err)
		return ""
	}
	return strings.TrimSpace(message)
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chat session",
	Run: func(cmd *cobra.Command, args []string) {
		userID := GetLoggedInUserID()
		username := GetLoggedInUsername()
		if userID == 0 || username == "" {
			fmt.Println("You need to log in first!")
			return
		}

		fmt.Printf("Starting chat for user: %s (ID: %d)\n", username, userID)

		// Connect to the chat server using Socket.IO client
		opts := &socketio_client.Options{Transport: "websocket"}
		client, err := socketio_client.NewClient("http://localhost:8000", opts)
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}

		// Handle receiving messages
		client.On("chat message", func(msg string) {
			fmt.Println(msg)
		})

		// Send user info after login
		client.Emit("set_user_info", userID, username)

		// Goroutine to handle user input and send messages
		go func() {
			for {
				message := readUserInput()
				if message != "" {
					client.Emit("chat message", message)
				}
			}
		}()

		// Handle OS interrupt signals for graceful shutdown
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("\nDisconnecting...")

	},
}

func init() {
	RootCmd.AddCommand(chatCmd)
}
