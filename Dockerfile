# Use the official Golang image
FROM golang:1.22-alpine AS go-builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use Node.js for SvelteKit
FROM node:18-alpine AS node-builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the code
COPY . .

# Build the SvelteKit app
RUN npm run build

# Final stage
FROM alpine:latest

# Install Node.js
RUN apk add --no-cache nodejs

# Copy the Go binary from the go-builder stage
COPY --from=go-builder /app/main /app/main

# Copy the built SvelteKit app from the node-builder stage
COPY --from=node-builder /app/build /app/client

# Expose ports
EXPOSE 3000 5173

# Start the Go app
CMD ["/app/main"]