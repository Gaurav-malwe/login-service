# Use the official Golang image
FROM golang:1.22.6-alpine3.20

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build ./cmd/main.go

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
