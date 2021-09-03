FROM golang:1.17.0

RUN mkdir /portfolio

WORKDIR /portfolio

COPY . /portfolio

RUN go env -w GO111MODULE=on
RUN go get -u github.com/gin-gonic/gin

# RUN curl -fLo /go/bin/air https://git.io/linux_air \
#   && chmod +x /go/bin/air

# CMD air
