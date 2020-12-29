FROM golang:1.15

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./app.go

CMD ["./app"]