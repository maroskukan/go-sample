FROM golang:1.20
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./app.go
CMD ["./app"]