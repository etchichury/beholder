FROM golang:1.27.4-alpine

WORKDIR /app

COPY go.mod .env ./
COPY cmd/watcher ./cmd/watcher

RUN go build -o main ./cmd/watcher

RUN chmod +x ./main

ENTRYPOINT ["./main"]