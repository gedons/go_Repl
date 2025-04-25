FROM golang:1.21-alpine as build

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code into the container
COPY . ./

# Build the Go binary from the correct entrypoint
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Second stage: Use a smaller base image
FROM alpine:3.18

WORKDIR /root/

# Copy the compiled binary from build stage
COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
