FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go /app
COPY pkg /app/pkg
# COPY .env .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-executer .

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

RUN apk --no-cache add ca-certificates go

COPY --from=builder /app/go-executer .
COPY .env .env

RUN echo $"PATH"
CMD ["./go-executer"]
