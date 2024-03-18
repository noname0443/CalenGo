FROM golang:1.21 AS builder

ADD ./backend /build

WORKDIR /build

RUN go build -o /build/cmd/app cmd/main.go

WORKDIR /build/cmd/

CMD ["./app"]