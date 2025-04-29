# Stage 1: build your API server
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy -v && go mod verify

RUN go build -v -o server ./cmd/server

# Stage 2: create lightweight runtime image
FROM alpine:3.18

RUN apk add --no-cache bash curl

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]