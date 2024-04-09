package main

import (
	"log"
	"net/http"

	"github.com/fyfirman/auth-management-go/internal/app"
	"github.com/fyfirman/auth-management-go/internal/repository"
	"github.com/fyfirman/auth-management-go/internal/service"
)

func main() {
	_, err := repository.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database")
		return
	}

	userRepository := repository.NewUserRepository()
	tokenRepository := repository.NewTokenRepository()

	userService := service.NewUserService(userRepository, tokenRepository)
	userHandler := app.NewUserHandler(userService)

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/forgot-password", userHandler.ForgotPasswordPassword)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
