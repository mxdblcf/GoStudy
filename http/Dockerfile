

# source
FROM golang:1.18.3-alpine


#maintainer
MAINTAINER mxd "mxdblcf@163.com"
#docker 内部命令
RUN mkdir -p /www/webapp
#设置工作路径
WORKDIR /www/webapp

#复制代码
COPY . /www/webapp

RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on
RUN go build -o httpTest ./hello/client
RUN go build -o hello-server  ./hello
RUN go build httpTest.go

EXPOSE 9999

