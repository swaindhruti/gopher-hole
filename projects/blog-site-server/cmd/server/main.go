package main

import (
	"blog-app/internal/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Loading environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	// Initialize database connection
	if err := db.InitializeDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.CloseDB()

	// Setup routes
	mux := http.NewServeMux()

	// Statring the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting blog site server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
