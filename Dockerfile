# Stage 1: build your API server
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Run go mod tidy to ensure all dependencies are resolved
RUN go mod tidy

# Now copy the rest of the application code
COPY . ./

# Build your Go application
RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# Stage 2: runtime with Go compiler installed
FROM alpine:3.18

# Install bash, curl—and the Go toolchain so user snippets can compile
RUN apk add --no-cache bash curl go

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
