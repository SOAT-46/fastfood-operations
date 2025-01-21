FROM golang:1.23.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/google/wire/cmd/wire@latest
RUN go generate ./...
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main/main.go # Adjust based on your main.go location

RUN go build -o main ./cmd/main

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
