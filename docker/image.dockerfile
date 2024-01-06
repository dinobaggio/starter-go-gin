FROM golang:1.20

WORKDIR /app

COPY ../ .

RUN go mod tidy

COPY ../.env.dev .env

RUN go install github.com/cosmtrek/air@latest

