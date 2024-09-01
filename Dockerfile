# Build stage
FROM golang:1.23-alpine AS build

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest


WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Entrypoint command
CMD ["air"]
