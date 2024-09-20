package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func connectDB() *pgxpool.Pool {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	conn, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to the database!")
	return conn
}

func main() {
	db = connectDB()
	defer db.Close()

	// Test the connection
	var greeting string
	err := db.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Println(greeting)
}
