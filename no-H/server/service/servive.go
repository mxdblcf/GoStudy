package service

import (
	"GoStudy/micro/grpc/simple/proto/pb"
	"context"
	"fmt"
	"reflect"
)

type Hello struct{}

func (cls *Hello) HelloInfo(ctx context.Context, pkg *pb.HelloInfoRequest) (*pb.HelloInfoReply, error) {
	fmt.Println("mxd", pkg.GetName(), reflect.TypeOf(pkg.Name))
	return &pb.HelloInfoReply{
		Name: pkg.Name,
	}, nil
}
