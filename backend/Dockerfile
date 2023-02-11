# ベースとなるDockerイメージ指定
FROM golang:1.18.1-alpine

WORKDIR /app
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

COPY *.go ./
COPY go.mod go.sum ./
RUN go mod download
# RUN go get -u github.com/gin-gonic/gin && \
#   go get github.com/jinzhu/gorm && \
#   go get github.com/jinzhu/gorm/dialects/postgres && \
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]