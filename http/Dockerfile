
FROM golang:1.16

MAINTAINER mxd "mxdblcf@163.com"

WORKDIR $GOPATH/src/httpTest
ADD httpTest.go $GOPATH/src/httpTest
RUN go build httpTest.go

EXPOSE 9999

ENTRYPOINT ["./httpTest"]