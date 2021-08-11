
FROM golang:latest

MAINTAINER mxd "raymond_2008@yahoo.com"

WORKDIR $GOPATH/src/httpTest
ADD . $GOPATH/src/httpTest
RUN go build httpTest.go

EXPOSE 9090

ENTRYPOINT ["./httpTest"]