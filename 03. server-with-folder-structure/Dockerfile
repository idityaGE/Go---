# Use the latest stable version of Golang (or a more recent one than 1.12.0 if possible)
FROM golang:1.20-alpine3.17

# Set environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working directory
WORKDIR /app

# Install git and other necessary tools in a single step to reduce layers
RUN apk add --no-cache git

# Copy dependency files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies without tidying to avoid introducing unintended changes
RUN go mod download

# Copy application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Define the default command to run the application
CMD ["./main"]
