#base go image
FROM golang:1.22-alpine3.19 AS builder

#to create directory
RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o api-gateway ./cmd/api

#give permission 
RUN chmod +x /app/api-gateway

FROM alpine:3.19

RUN mkdir /app

COPY --from=builder /app/api-gateway /app

CMD ["/app/api-gateway"]


