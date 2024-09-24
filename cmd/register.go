package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/dimple278/go-chat-app/db"
	"github.com/dimple278/go-chat-app/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Enter Email: ")
		email := utils.ReadUserInput()
		email = strings.TrimSpace(email)

		fmt.Print("Enter Username: ")
		username := utils.ReadUserInput()
		username = strings.TrimSpace(username)

		fmt.Print("Enter Password: ")
		password := utils.ReadUserInput()
		password = strings.TrimSpace(password)

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		// Call the db package to handle user registration
		err = db.RegisterUser(email, username, hashedPassword)
		if err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}

		fmt.Println("User registered successfully! Please log in to continue.")
		loginCmd.Run(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd) // Adding the register command to RootCmd
}
