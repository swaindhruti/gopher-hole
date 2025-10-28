package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DB is the global database connection pool
var DB *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() error {
	// Get the database connection string from environment variables
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Open the database connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Verify the connection
	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connection established")
	return nil
}

// GetDB returns the database connection pool
func GetDB() *sql.DB {
	return DB
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
