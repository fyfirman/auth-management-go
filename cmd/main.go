package main

import (
	"log"
	"net/http"

	"github.com/fyfirman/auth-management-go/internal/auth/app"
)

func main() {
	mux := http.NewServeMux()

	// Initialize the handler
	handler := app.NewHandler()

	// Setup the API routes
	handler.SetupRoutes(mux)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
