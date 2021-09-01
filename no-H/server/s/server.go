package main

import (
	"GoStudy/micro/grpc/simple/proto/pb"
	"GoStudy/micro/grpc/simple/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	addres := "localhost:8889"                 // 服务地址
	listener, err := net.Listen("tcp", addres) // 监听地址
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()                              // 初始化grpc服务
	pb.RegisterHelloServiceServer(s, &service.Hello{}) // 注册服务
	log.Println("start gRPC listen on addres " + addres)
	if err := s.Serve(listener); err != nil { // 监听服务，如启动失败则抛出异常
		log.Println("failed to serve...", err)
		return
	}
}
