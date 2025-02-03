FROM golang:1.23.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go install github.com/google/wire/cmd/wire@latest
RUN go generate ./...
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

RUN go build -o main ./cmd

FROM alpine:3.21.2

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
