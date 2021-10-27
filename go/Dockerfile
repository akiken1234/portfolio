FROM golang:1.16.0-alpine

WORKDIR /go/src/github.com/akimotokensaku/portfolio/go

COPY . .

ENV GO111MODULE=on
# データベース接続のため
ENV CGO_ENABLED=0

RUN apk add git

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
