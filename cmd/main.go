package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatapp",
	Short: "CLI Chat Application",
	Long:  "A Command Line Interface Chat Application using Golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI Chat Application!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
