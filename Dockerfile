# Build stage
FROM golang:nanoserver-ltsc2022 AS builder

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copy the Go dependency files and download the modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Compile the application
RUN go build -o main ./cmd/main.go

# Final stage
FROM alpine:3.18

# Create a non-root user to run the application (for security)
RUN adduser -D -g '' appuser

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=builder /app/main .

# Adjust permissions
RUN chown -R appuser /app

# Switch to the non-root user
USER appuser

# Expose the port that the application will use
EXPOSE 8080

# Command to run the application
CMD ["air", "-c", ".air.toml"]

