# Golang Backend

## Project Overview
Golang Backend is a robust REST API.

## Features
- Golang Gin REST API
- JWT Authentication.

## Golang Layer Model (Universal Reference Architecture)

```sh

.github/                # GitHub Actions configuration or other workflow settings
.vscode/                # Visual Studio Code-specific configuration
cmd/
    app/                # Main application (entry point)
configs/                # Configuration files
docs/                   # Project documentation
internal/               # Private project code, not exported
    core/               # Core domain logic
        shareds/        # Core-shared logic
        users/          # User-specific logic (e.g., authentication, permissions)
    infra/              # Infrastructure implementations (e.g., DB, external APIs)
        middleware/     # HTTP Middlewares
    utils/              # Utilities and helper functions
migrations/             # Database migration scripts
tests/                  # Infrastructure automated tests
pkg/                    # Shared code among multiple projects or components
api/                    # API definitions (e.g., gRPC Protobuf or Swagger/OpenAPI)
scripts/                # Automation, CI/CD scripts
.env                    # Environment variables file
.env.example            # Environment variables file example
docker-compose.yml      # Docker Compose file for multi-container configuration
Dockerfile              # Dockerfile for image building
go.mod                  # Go dependencies checksum file
go.sum                  # Checksums of Go dependencies
main.go                 # Main application entry point
Makefile                # Build automation using make
README.md               # Project description


```

---
### Relevant Golang Libraries
Update packages:

```shell
$ go mod tidy
```

Install dependencies:

```shell
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/go-playground/validator/v10
$ go get github.com/gin-gonic/gin
$ go get go.mongodb.org/mongo-driver/mongo
$ go get github.com/joho/godotenv
$ go get github.com/gin-contrib/gzip
$ go get github.com/go-faker/faker/v4
```

Install Swagger (swaggo):

```shell
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
$ go install github.com/swaggo/swag/cmd/swag@latest
```

Install AWS SDK for S3:

```shell
$ go get -u github.com/aws/aws-sdk-go-v2
$ go get -u github.com/aws/aws-sdk-go-v2/config
$ go get -u github.com/aws/aws-sdk-go-v2/service/s3
```

Install Debugger (delve):

```shell
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

Install 'Testing Schema Validation' for Mocking:

```shell
$ go get github.com/stretchr/testify
```

Install Go Releaser:

```shell
$ go install github.com/goreleaser/goreleaser/v2@latest
```


Install golang migrate for migrations:

```shell
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Check Swagger version:

```shell
$ swag --version
```

### Generate Swagger:

```shell
$ swag init --parseDependency -g cmd/main.go
# or
$ swag init --parseDependency
```
--parseDependency flag is needed to make sure Swag correctly analyzes and includes all code dependencies, like encoding/json

### Create the App and database container via docker-compose and Dockerfile:

```shell
$ docker-compose up --build
```

### Application deployment (local testing):

```shell
$ go run main.go
```

### Swagger URL:

```shell
http://localhost:8081/swagger/index.html
```

---

