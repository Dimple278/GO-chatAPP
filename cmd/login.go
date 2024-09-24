package cmd

import (
	"fmt"
	"strings"

	"github.com/dimple278/go-chat-app/db"
	"github.com/dimple278/go-chat-app/utils"
	"github.com/spf13/cobra"
)

var loggedInUserID int
var loggedInUsername string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Enter Username: ")
		username := utils.ReadUserInput()
		username = strings.TrimSpace(username)

		fmt.Print("Enter Password: ")
		password := utils.ReadUserInput()
		password = strings.TrimSpace(password)

		// Use db package to authenticate the user
		userID, err := db.AuthenticateUser(username, password)
		if err != nil {
			fmt.Println("Login failed:", err)
			return
		}

		fmt.Println("Login successful!")
		fmt.Printf("Welcome, User: %s (ID: %d)\n", username, userID)

		// Set the logged-in user ID and username for the chat functionality
		SetLoggedInUserID(userID)
		SetLoggedInUsername(username)

		// Start the chat by calling the chat command directly
		chatCmd.Run(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}

// SetLoggedInUserID sets the user ID of the logged-in user
func SetLoggedInUserID(userID int) {
	loggedInUserID = userID
}

// SetLoggedInUsername sets the username of the logged-in user
func SetLoggedInUsername(username string) {
	loggedInUsername = username
}

// GetLoggedInUserID returns the user ID of the logged-in user
func GetLoggedInUserID() int {
	return loggedInUserID
}

// GetLoggedInUsername returns the username of the logged-in user
func GetLoggedInUsername() string {
	return loggedInUsername
}
