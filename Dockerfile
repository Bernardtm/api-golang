# Start from the official Golang base image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all Go modules
RUN go mod download

# Copy the rest of the application code (this is where it copies all the source files)
COPY . .

# Build the Go app inside the cmd/myapp directory
RUN go build -o /bernardtm/backend ./cmd/myapp

# Create a smaller image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /bernardtm/backend .

# Expose the application port (replace with your app's port)
EXPOSE 8080

# Command to run the Go program
CMD ["./bernardtm/backend"]
