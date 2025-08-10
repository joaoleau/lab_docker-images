FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o yaml-checker ./src/yamlChecker/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o file-extract ./src/fileExtract/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o comment-on-pr ./src/commentOnPr/cmd/main.go

FROM alpine:latest
RUN apk update && apk add --no-cache git
COPY --from=builder /app/yaml-checker /usr/local/bin/yaml-checker
COPY --from=builder /app/file-extract /usr/local/bin/file-extract
COPY --from=builder /app/comment-on-pr /usr/local/bin/comment-on-pr
COPY test.md .