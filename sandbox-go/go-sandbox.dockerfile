FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go /app
COPY pkg /app/pkg
# COPY golang-app-seccomp.json .


# COPY .env .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-executer .

FROM alpine:3.19

# USER root 
# RUN mkdir -p /etc/docker/seccomp
# COPY golang-app-seccomp.json .


RUN mkdir /app && \
    apk --no-cache add ca-certificates go && \   
    adduser -D -s /bin/sh appuser
RUN mkdir -p /app/etc/docker/seccomp
# RUN chmod +r /etc/docker/seccomp/golang-app-seccomp.json

WORKDIR /app


USER appuser                        


COPY --from=builder /app/go-executer .
COPY .env .env

COPY golang-app-seccomp.json /app/etc/docker/seccomp/golang-app-seccomp.json

# COPY golang-app-seccomp.json /etc/docker/seccomp/

RUN echo "$PATH"
CMD ["./go-executer"]
