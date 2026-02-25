package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enzomendez315/sentinel/internal/db"
	"github.com/enzomendez315/sentinel/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	pool, err := db.Init(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := server.Run(pool); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
