
FROM golang:1.16
#author info

MAINTAINER mxd "mxdblcf@163.com"
# 容器环境变量添加，会覆盖默认的变量值

WORKDIR /go
COPY httpTest.go  /go
RUN go build httpTest.go
EXPOSE 9999
CMD ["./go/httptest"]





