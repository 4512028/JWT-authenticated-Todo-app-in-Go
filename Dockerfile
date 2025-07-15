FROM golang:1.24-alpine

# Install deps
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build the binary
RUN go build -o todo ./cmd/main.go

# Use .env file in container
ENV GIN_MODE=release

# Expose app port
EXPOSE 8080

# Start app
CMD ["./todo"]