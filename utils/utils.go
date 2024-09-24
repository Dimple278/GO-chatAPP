package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadUserInput is a helper function to read user input from the terminal.
func ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading message:", err)
		return ""
	}
	return strings.TrimSpace(message)
}
