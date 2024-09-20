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
		var passwordHash string
		err := db.DB.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE LOWER(username) = LOWER($1)", username).Scan(&passwordHash)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Username not found!")
				return
			}
			log.Fatalf("Error fetching user: %v", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
		if err != nil {
			fmt.Println("Invalid credentials!")
			return
		}

		fmt.Println("Login successful!")
		// Proceed to chat functionality (to be implemented later)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
