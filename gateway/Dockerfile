FROM golang:1.22-alpine

# Install Node.js and npm
RUN apk add --no-cache nodejs npm

# Install Air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy package.json and package-lock.json
COPY package*.json ./

# Install Node.js dependencies
RUN npm install

# Copy the rest of the code
COPY . .

# Build the SvelteKit app
RUN npm run build

# Expose ports
EXPOSE 8000 5173

# Start Air for hot-reloading and the SvelteKit dev server
CMD air & npm run dev -- --host