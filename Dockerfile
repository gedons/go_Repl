# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/server  

# Stage 2: Final image
FROM alpine:3.18

# Install bash and curl (if needed)
RUN apk add --no-cache bash curl

# Copy the compiled binary from the build stage to /usr/local/bin/
COPY --from=build /app/server /usr/local/bin/server

# Set the working directory to where the binary is
WORKDIR /usr/local/bin

# Expose the application port
EXPOSE 8080

# Set the entrypoint to run the binary
CMD ["./server"]
