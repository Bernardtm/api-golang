# Variables
APP_NAME := infratech-backend
SRC_DIR := ./cmd/$(APP_NAME)
BIN_DIR := ./bin
BIN_PATH := $(BIN_DIR)/$(APP_NAME)
PKG := ./...
GOFILES := $(shell find . -name '*.go' -not -path './vendor/*')
POSTGRES_MIGRATIONS_DIR=./migrations/postgres

# Default commands
.PHONY: build run fmt lint test test-coverage clean mod vet deps swagger releaser-check release create-migration-postgres create-seed-postgres help

# Default target
all: build

# Command to compile the project
build:
	@echo "Compiling project..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH) $(SRC_DIR)

# Command to execute the project
run: build
	@echo "Starting project..."
	$(BIN_PATH)

# Format code
fmt:
	go fmt $(PKG)

# Lint code
lint:
	golangci-lint run

# Command to clean the generated binaries
clean:
	@echo "Cleaning files..."
	rm -rf $(BIN_DIR)

# Generate Go modules
mod:
	go mod tidy

# Run static analysis
vet:
	go vet $(PKG)

# Install dependencies
deps:
	go mod download

test:
	@echo "Running tests..."
	go test ./...

# Run tests with coverage
test-coverage:
	go test -cover $(PKG)

# Command to generate Swagger documentation
swagger:
	@echo "Generating Swagger docs..."
	swag init --parseDependency -g main.go

releaser-check:
	@echo "Checking if releaser config is valid..."
	goreleaser check

release:
	@echo "Releasing..."
	goreleaser release --clean

create-migration-postgres:
	@echo "Creating Postgres migration..."
	migrate create -ext sql -dir $(POSTGRES_MIGRATIONS_DIR) -seq $(name)

create-seed-postgres:
	@echo "Creating Postgres seed..."
	migrate create -ext sql -dir $(POSTGRES_MIGRATIONS_DIR) -seq seed_$(name)

migrate-up-postgres:
	@echo "Running Postgres migrations up..."
	migrate -path=$(POSTGRES_MIGRATIONS_DIR) -database "$(database)" -verbose up

migrate-down-postgres:
	@echo "Running Postgres migrations down..."
	migrate -path=$(POSTGRES_MIGRATIONS_DIR) -database "$(database)" -verbose down

# Help command
help:
	@echo "Available targets:"
	@echo "  build          Build the application"
	@echo "  run            Run the application"
	@echo "  fmt            Format the code"
	@echo "  lint           Lint the code using golangci-lint"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage report"
	@echo "  clean          Clean up build files"
	@echo "  mod            Update Go modules"
	@echo "  vet            Run static analysis"
	@echo "  deps           Install dependencies"
	@echo "  swagger        Generate Swagger documentation"
	@echo "  releaser-check Check if releaser config is valid"
	@echo "  release        Release the project"
	@echo "  create-migration-postgres name=<migration_name>  Create a new PostgreSQL migration"
	@echo "  create-seed-postgres name=<migration_seed_name>  Create a new PostgreSQL seed"
