#base go image
FROM golang:1.22-alpine3.19 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api cmd/api
COPY pkg pkg
COPY api-gateway.env api-gateway.env

RUN go build -o gatewayApp ./cmd/api

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/gatewayApp .
COPY api-gateway.env .

CMD ["sh","-c","echo $PORT && ./gatewayApp"]





