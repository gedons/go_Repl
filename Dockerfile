# Stage 1: build your API server
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# Stage 2: runtime with Go compiler installed
FROM alpine:3.18

# install bash, curl—and the Go toolchain so user snippets can compile
RUN apk add --no-cache bash curl go

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
