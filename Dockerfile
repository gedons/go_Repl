# Stage 1: build your API server
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && go mod verify

RUN go build -v -o server ./cmd/server

FROM alpine:3.18

RUN apk add --no-cache bash curl go

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

# Start the application
CMD ["./server"]
