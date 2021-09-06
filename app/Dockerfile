FROM golang:1.17.0

RUN mkdir /portfolio

WORKDIR /portfolio/app

COPY . /portfolio

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
