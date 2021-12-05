FROM --platform=linux/x86_64 golang:1.16.0-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

RUN apk update && apk add git 

ENV TZ Asia/Tokyo
ENV GO111MODULE=on
# データベース接続のため
ENV CGO_ENABLED=0

CMD ["go", "run", "main.go"]