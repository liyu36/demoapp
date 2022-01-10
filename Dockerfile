FROM golang:1.16.12 as base
WORKDIR /demo
ENV GOPROXY="https://goproxy.cn,direct"
COPY * ./
RUN go mod tidy
RUN go build -o demoapp . 

FROM ubuntu:18.04
WORKDIR /demo
COPY --from=base /demo/demoapp .
EXPOSE 80
CMD "/demo/demoapp"