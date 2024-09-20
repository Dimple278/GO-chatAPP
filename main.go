package main

import (
	"log"

	"github.com/dimple278/go-chat-app/cmd"
	"github.com/dimple278/go-chat-app/db"
)

func main() {
	db.Connect()
	defer db.DB.Close()

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
