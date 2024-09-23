package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dimple278/go-chat-app/db"
)

// HistoryCommand retrieves chat history with a limit (default is 20)
func HistoryCommand(msg string) string {
	parts := strings.Split(msg, " ")
	limit := 20
	if len(parts) == 2 {
		if l, err := strconv.Atoi(parts[1]); err == nil {
			limit = l
		}
	}

	chatHistory, err := db.FetchChatHistory(limit, 0)
	if err != nil {
		return "Error fetching chat history"
	}

	if len(chatHistory) == 0 {
		return "No chat history available"
	}

	result := "Chat history:\n"
	for _, history := range chatHistory {
		result += fmt.Sprintf("%s\n", history)
	}

	return result
}
