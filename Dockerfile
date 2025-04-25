# Use a multi-stage build to compile the Go binary
FROM golang:1.21-alpine as build

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code and build the Go binary
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Use Docker-in-Docker for running Docker commands
FROM docker:20.10.7-dind

WORKDIR /root/

# Install Go (to run the Go binary) and dependencies
RUN apk add --no-cache bash git curl

# Copy the compiled Go binary from the build stage
COPY --from=build /app/main .

# Expose the port the app will listen on
EXPOSE 8080

# Start Docker daemon and your Go application
CMD ["sh", "-c", "dockerd & ./main"]
