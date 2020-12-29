FROM golang:1.15

RUN mkdir -p /project
WORKDIR /project
ADD . /project
RUN go build ./main.go
CMD ["./websrvc"]