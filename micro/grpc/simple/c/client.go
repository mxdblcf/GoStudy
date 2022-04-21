package main

import (
	"GoStudy/micro/grpc/simple/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	//address = "localhost:8889"
	address = "localhost:8888"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 建立链接
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)                 // 初始化客户端
	ctx := context.Background()                              // 初始化元数据
	helloName := &pb.HelloInfoRequest{Name: "chengjian.liu"} // 构造请求体
	info, err := client.HelloInfo(ctx, helloName)            // 调用helloInfo rpc服务，就像调用普通函数一样
	if err != nil {
		return
	}
	fmt.Println("hello", info.Name)
}
