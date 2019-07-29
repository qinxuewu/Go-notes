package main

import (
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"io"
	"log"
	"net"
)


type StreamService struct {

}


func main() {
	// 创建 gRPC Server 对象
	server := grpc.NewServer()
	// 将方法的rpc server注册到 gRPC Server 的内部注册中心
	pb.RegisterStreamServiceServer(server,&StreamService{})
	// 创建 Listen，监听 TCP 端口
	lis,err:=net.Listen("tcp",":8899")
	if err !=nil {
		log.Fatal("net.Listen  err: %v",err)
	}
	// 开始 lis.Accept，直到 Stop 或 GracefulStop
	server.Serve(lis)

}

// 服务器端流返回流式数据
func (s *StreamService) List(r *pb.StreamRequest,stream pb.StreamService_ListServer) error  {
	for n := 0; n <= 6; n++ {
		// 循环返回流式数据给客户端
		//  protoc 在生成时，根据定义生成了各式各样符合标准的接口方法。最终再统一调度内部的 SendMsg 方法
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
// 接收客户端流式 RPC
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error  {
	for  {
		r,err:=stream.Recv()
		if err !=nil {
			return  err
		}
		// 流关闭 后
		if err == io.EOF{
			// 将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Recv
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		log.Printf("接收客户端流式 RPC pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return  nil
}
// 双向流式 RPC
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	return nil
}