FROM golang:1.15
WORKDIR /project
COPY . /project
RUN go build main.go
CMD ["websrvc"]