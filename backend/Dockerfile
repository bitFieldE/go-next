# ベースとなるDockerイメージ指定
FROM golang:1.20.3-alpine as dev

COPY . /github.com/bitFieldE/go-next/backend
WORKDIR /github.com/bitFieldE/go-next/backend

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

COPY *.go ./
COPY go.mod go.sum ./
RUN apk update && \
  apk add --no-cache git && \
  apk add --update gcc musl-dev && \
  go mod download && \
  go get -u github.com/gin-gonic/gin && \
  go install github.com/pressly/goose/v3/cmd/goose@latest && \
  go install github.com/cosmtrek/air@latest && \
  go install github.com/volatiletech/sqlboiler/v4@latest && \
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

CMD ["air", "-c", ".air.toml"]
