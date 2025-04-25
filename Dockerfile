# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS build

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy app source and build
COPY . .
RUN go build -o main ./cmd/server

# Stage 2: Final image without DinD
FROM alpine:3.18

# Install Docker client only, and other required tools
RUN apk add --no-cache bash curl docker-cli

WORKDIR /root/

# Copy the compiled binary
COPY --from=build /app/main .

# Expose the application port
EXPOSE 8080

# Run the Go app
CMD ["./main"]
