# ベースとなるDockerイメージ指定
FROM golang:1.18.1-alpine
WORKDIR /app
RUN go mod init backend
RUN go get github.com/gin-gonic/gin \
    github.com/jinzhu/gorm \
# RUN mkdir /go/src/backend
# WORKDIR /go/src/backend
# COPY ./app /go/src/backend
# VOLUME /go/src/backend
# RUN go mod download
# EXPOSE 3000
CMD "go" "run" "main.go"