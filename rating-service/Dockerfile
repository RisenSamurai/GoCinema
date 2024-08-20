# Use the official Golang image for building the application
FROM golang:1.22-alpine AS builder

# Set the working directory for the build context
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY app/rating-service/go.mod app/rating-service/go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire microservice directory
COPY app/rating-service /app/rating-service

# Set the working directory inside the microservice folder
WORKDIR /app/rating-service

# Build the Go application
RUN go build -o main .

# Use a minimal base image for running the application
FROM alpine:latest

# Set the working directory in the second stage
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/rating-service/main .

# Expose the port that the app runs on (adjust as necessary)
EXPOSE 8081

# Command to run the binary
CMD ["./main"]
