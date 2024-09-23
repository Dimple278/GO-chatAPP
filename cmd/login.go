package cmd

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dimple278/go-chat-app/db"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var loggedInUserID int
var loggedInUsername string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		// Fetch user from the database
		var userID int
		var passwordHash string
		err := db.DB.QueryRow(context.Background(), "SELECT id, password_hash FROM users WHERE LOWER(username) = LOWER($1)", username).Scan(&userID, &passwordHash)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Username not found!")
				return
			}
			log.Fatalf("Error fetching user: %v", err)
		}

		// Compare provided password with the stored password hash
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
		if err != nil {
			fmt.Println("Invalid credentials!")
			return
		}

		fmt.Println("Login successful!")
		fmt.Printf("Welcome, User: %s (ID: %d)\n", username, userID)

		// Set the logged-in user ID for the chat functionality
		SetLoggedInUserID(userID)

		// Set the logged-in username for the chat functionality
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

func SetLoggedInUsername(username string) {
	loggedInUsername = username
}

// GetLoggedInUserID returns the user ID of the logged-in user
func GetLoggedInUserID() int {
	return loggedInUserID
}

func GetLoggedInUsername() string {
	return loggedInUsername
}
