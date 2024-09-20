package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	socketio_client "github.com/hesh915/go-socket.io-client"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start chatting with other users",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to the chat server
		serverURL := "http://localhost:8000"
		opts := &socketio_client.Options{
			Transport: "websocket",
		}

		client, err := socketio_client.NewClient(serverURL, opts)
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}

		// Mutex to handle message output and input safely
		var mu sync.Mutex

		// Handle incoming chat messages
		client.On("chat message", func(msg string) {
			mu.Lock()
			fmt.Print("\r\033[K") // Clear current line before printing new message
			fmt.Println("Message from chat:", msg)
			fmt.Print("Enter message: ") // Redisplay the prompt after message
			mu.Unlock()
		})

		// Handle connection event
		client.On("connect", func() {
			fmt.Println("Connected to chat!")
		})

		// Keep reading user input and sending messages
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			if message == "exit" {
				break
			}

			// Send the message to the chat server
			client.Emit("chat message", message)
		}

		// Since no explicit disconnect method is available, the program will
		// automatically end the connection when the program exits
		log.Println("Exiting chat...")
	},
}

func init() {
	RootCmd.AddCommand(chatCmd)
}
