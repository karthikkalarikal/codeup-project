FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go /app
COPY pkg /app/pkg
# COPY .env .env

RUN go build -o main .

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env


CMD ["./main"]
