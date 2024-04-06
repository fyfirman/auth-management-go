# Makefile

# Include .env file and export its variables
-include .env
export

# Default values for variables
DB_USER ?= default_user
DB_PASSWORD ?= default_password
DB_NAME ?= default_db_name
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_SSLMODE ?= disable

DB_DSN := "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)"

MIGRATIONS_DIR := db/migrations

GOOSE_CMD := goose -dir=$(MIGRATIONS_DIR) postgres $(DB_DSN)

BINARY_NAME := auth_management


.PHONY: migrate-create migrate-up migrate-down

# Create a new migration
migrate-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE_CMD) create $$name sql

# Apply all available migrations
migrate-up:
	$(GOOSE_CMD) up

# Roll back the last migration
migrate-down:
	$(GOOSE_CMD) down

# Build the Go application
build:
	go build -o $(BINARY_NAME) ./cmd

# Run the Go application
run:
	go run ./cmd/main.go