# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app
COPY . .

ENV EXEC_TIMEOUT=3s

CMD ["go", "run", "temp.go"]
