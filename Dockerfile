# Development Dockerfile
FROM golang:1.22-alpine

# Install Node.js and npm
RUN apk add --no-cache nodejs npm

# Set working directory
WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install Node.js dependencies
COPY package*.json ./
RUN npm install

RUN npm install -D tailwindcss postcss autoprefixer svelte-preprocess

RUN npx tailwindcss init tailwind.config.cjs -p

# Copy the rest of the code
COPY . .

# Install Air for hot-reloading Go code
RUN go install github.com/air-verse/air@latest

# Expose ports for Go and SvelteKit
EXPOSE 8000 5173

# Start development servers
CMD air & npm run dev -- --host
