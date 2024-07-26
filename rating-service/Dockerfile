FROM ubuntu:latest
LABEL authors="Samurai"

FROM golang:1.22-alpine

# Install Node.js and npm
RUN apk add --no-cache nodejs npm

# Set working directory
WORKDIR /app/rating-service

# Install Go dependencies
COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest


CMD air & --host