package main
import (
	"context"
	"log"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
)



func main() {
	// 创建与给定目标（服务端）的连接交互
	conn, err := grpc.Dial(":8899", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	// 创建 StudentProtoService 的客户端对象
	client := pb.NewStudentProtoServiceClient(conn)
	// 发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	resp, err := client.GetName(context.Background(), &pb.MyRequest{Name: "zhangsan ",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	// 输出响应结果
	log.Printf("服务端返回: %s", resp.GetRealname())
}