FROM golang:1.18 AS builder

WORKDIR /build

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

ADD go.mod .
COPY . .
RUN go build -o webshell webshell.go

RUN ls -all
CMD ["./webshell"]