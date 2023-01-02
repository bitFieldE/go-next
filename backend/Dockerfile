# ベースとなるDockerイメージ指定
FROM golang:1.18.1-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD [ "/main" ]