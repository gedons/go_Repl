FROM golang:1.21-alpine as build

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code and build the Go binary
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Second stage: Use a smaller base image for the actual running container
FROM alpine:3.18

WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/main .

# Expose the port the app will listen on
EXPOSE 8080

# Run the Go binary
CMD ["./main"]

