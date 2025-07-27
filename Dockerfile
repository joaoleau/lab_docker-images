FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o yaml-checker


FROM alpine:latest
RUN apk update && apk add --no-cache git
COPY --from=builder /app/yaml-checker /usr/local/bin/yaml-checker