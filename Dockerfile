FROM golang:1.13-buster as builder

EXPOSE 53/udp 80

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

RUN go build -o iepee cmd/main.go

CMD ./iepee
