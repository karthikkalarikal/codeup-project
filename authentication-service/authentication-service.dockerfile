FROM golang:1.22-alpine3.19 AS builder

RUN mkdir /app
WORKDIR /app  

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api cmd/api
COPY pkg pkg

RUN go build -o authApp ./cmd/api

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/authApp .

CMD ["./authApp"]