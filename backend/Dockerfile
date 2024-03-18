FROM golang:1.21 AS builder

ADD ./backend /build

WORKDIR /build

RUN go build -o app cmd/main.go

CMD ["/build/app"]