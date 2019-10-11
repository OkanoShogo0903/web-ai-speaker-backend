# 開発環境
FROM golang:1.12.7

WORKDIR /api
COPY . .
ENV GO111MODULE=on

CMD go run main.go
