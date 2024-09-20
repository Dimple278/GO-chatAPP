package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dimple278/go-chat-app/db"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Email: ")
		email, _ := reader.ReadString('\n')

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Enter Password: ")
		password, _ := reader.ReadString('\n')

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.DB.Exec(context.Background(), "INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)", email, username, hashedPassword)
		if err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}

		fmt.Println("User registered successfully!")
	},
}

func init() {
	RootCmd.AddCommand(registerCmd) // Adding the register command to RootCmd
}
