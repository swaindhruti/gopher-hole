package main

import (
	"blog-app/internal/db"
	"blog-app/internal/handlers"
	"blog-app/internal/repository"
	"blog-app/internal/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	// Get database instance
	database := db.GetDB()

	// Initialize repositories
	blogRepo := repository.NewBlogRepository(database)
	userRepo := repository.NewUserRepository(database)
	commentRepo := repository.NewCommentRepository(database)

	// Initialize handlers
	blogHandler := handlers.NewBlogHandler(blogRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	commentHandler := handlers.NewCommentHandler(commentRepo)

	// Setup routes
	mux := routes.Setup(blogHandler, userHandler, commentHandler)

	// Starting the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting blog site server on port %s...\n", port)
	log.Println("Available endpoints:")
	log.Println("  POST   /blogs")
	log.Println("  GET    /blogs")
	log.Println("  GET    /blogs/{id}")
	log.Println("  PUT    /blogs/{id}")
	log.Println("  DELETE /blogs/{id}")
	log.Println("  POST   /users")
	log.Println("  GET    /users")
	log.Println("  GET    /users/{id}")
	log.Println("  PUT    /users/{id}")
	log.Println("  DELETE /users/{id}")
	log.Println("  POST   /comments")
	log.Println("  GET    /blogs/{blogID}/comments")
	log.Println("  GET    /comments/{id}")
	log.Println("  PUT    /comments/{id}")
	log.Println("  DELETE /comments/{id}")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
