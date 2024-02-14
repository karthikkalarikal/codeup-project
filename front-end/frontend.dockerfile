FROM golang:1.22-alpine3.19 AS builder

RUN mkdir /app

WORKDIR /app  

COPY frontApp .
COPY cmd/web/templates /app/cmd/web/templates

CMD ["/app/frontApp"]