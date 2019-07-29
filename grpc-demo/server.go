package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
)

type  StudentProtoService struct {
	
}

func (s *StudentProtoService)GetName(ctx context.Context,r *pb.MyRequest)(*pb.MyResponse,error)  {
	return &pb.MyResponse{Realname:r.GetName()+ "Server "},nil
}

func main() {
	// 创建 gRPC Server 对象
	server := grpc.NewServer()
	// 将方法的rpc server注册到 gRPC Server 的内部注册中心
	pb.RegisterStudentProtoServiceServer(server, &StudentProtoService{})
	// 创建 Listen，监听 TCP 端口
	lis, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	// 开始 lis.Accept，直到 Stop 或 GracefulStop
	server.Serve(lis)
}