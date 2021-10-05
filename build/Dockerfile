FROM golang:latest

MAINTAINER	liming<2572597150@qq.com>
ENV WORKPATH /go/src/app
WORKDIR $WORKPATH

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# 重定向日志输出
CMD go run $WORKPATH/main.go > ./server.log