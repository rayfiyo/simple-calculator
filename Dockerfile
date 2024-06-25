# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /usr/src/app

COPY server/ ../app
RUN go build -v -o ./app ./server.go

CMD ["./app"]
