package routes

import (
	"blog-app/internal/handlers"
	"net/http"
)

// Setup initializes all routes for the application
func Setup(
	blogHandler *handlers.BlogHandler,
	userHandler *handlers.UserHandler,
	commentHandler *handlers.CommentHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	// Blog routes
	mux.HandleFunc("POST /blogs", blogHandler.CreateBlog)
	mux.HandleFunc("GET /blogs", blogHandler.GetAllBlogs)
	mux.HandleFunc("GET /blogs/{id}", blogHandler.GetBlog)
	mux.HandleFunc("PUT /blogs/{id}", blogHandler.UpdateBlog)
	mux.HandleFunc("DELETE /blogs/{id}", blogHandler.DeleteBlog)

	// User routes
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Comment routes
	mux.HandleFunc("POST /comments", commentHandler.CreateComment)
	mux.HandleFunc("GET /blogs/{blogID}/comments", commentHandler.GetCommentsForBlog)
	mux.HandleFunc("GET /comments/{id}", commentHandler.GetComment)
	mux.HandleFunc("PUT /comments/{id}", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /comments/{id}", commentHandler.DeleteComment)

	return mux
}
