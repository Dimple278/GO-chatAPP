package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

// Initialize the database connection pool
func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to the database!")
}

// SaveChatMessage saves a chat message t the chat_history table
func SaveChatMessage(userID int, message string) error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `INSERT INTO chat_history (user_id, message, timestamp) VALUES ($1, $2, $3)`
	_, err := DB.Exec(context.Background(), query, userID, message, time.Now())
	if err != nil {
		return fmt.Errorf("failed to save message: %v", err)
	}
	return nil
}

// FetchChatHistory fetches chat history from the database with limit and offset
func FetchChatHistory(limit int, offset int) ([]string, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `
        SELECT u.username, ch.message, ch.timestamp 
        FROM chat_history ch
        JOIN users u ON ch.user_id = u.id
        ORDER BY ch.timestamp DESC
        LIMIT $1 OFFSET $2;
    `

	rows, err := DB.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch chat history: %v", err)
	}
	defer rows.Close()

	var chatMessages []string
	for rows.Next() {
		var username, message string
		var timestamp time.Time
		err := rows.Scan(&username, &message, &timestamp)
		if err != nil {
			return nil, err
		}
		formattedMessage := fmt.Sprintf("[%s] %s: %s", timestamp.Format("15:04:05"), username, message)
		chatMessages = append(chatMessages, formattedMessage)
	}

	// Reverse the order to get the correct chronological order
	for i, j := 0, len(chatMessages)-1; i < j; i, j = i+1, j-1 {
		chatMessages[i], chatMessages[j] = chatMessages[j], chatMessages[i]
	}

	return chatMessages, nil
}
