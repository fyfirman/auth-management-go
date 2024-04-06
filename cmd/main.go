package main

import (
	"log"
	"net/http"

	"github.com/fyfirman/auth-management-go/internal/app"
	"github.com/fyfirman/auth-management-go/internal/repository"
	"github.com/fyfirman/auth-management-go/internal/service"
)

func main() {
	repository.ConnectDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := app.NewUserHandler(userService)

	http.HandleFunc("/register", userHandler.Register)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
