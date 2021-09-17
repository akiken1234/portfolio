FROM golang:1.16.0

WORKDIR /go/src/github.com/akimotokensaku/portfolio/go

COPY . .

ENV GO111MODULE=on
# データベース接続のため
ENV CGO_ENABLED=0

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
